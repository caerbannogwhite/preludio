use clap::Parser;
use polars::export::num::ToPrimitive;
use polars::prelude::*;
use std::collections::HashMap;
use std::fs;
use std::ops::Index;
use std::str;

mod prql_std;
mod prql_vm;

const TYPE_NULL: u64 = 0;
const TYPE_BOOL: u64 = 1;
const TYPE_NUMERIC: u64 = 2;
const TYPE_STRING: u64 = 3;
const TYPE_IDENT: u64 = 4;
const TYPE_INTERVAL: u64 = 5;
const TYPE_RANGE: u64 = 6;
const TYPE_LIST: u64 = 7;
const TYPE_PIPELINE: u64 = 8;

const OP_BEGIN_PIPELINE: u64 = 0;
const OP_END_PIPELINE: u64 = 1;
const OP_ASSIGN_TABLE: u64 = 2;
const OP_BEGIN_FUNC_CALL: u64 = 3;
const OP_END_FUNC_CALL: u64 = 4;
const OP_BEGIN_LIST: u64 = 5;
const OP_END_LIST: u64 = 6;
const OP_ADD_FUNC_PARAM: u64 = 7;
const OP_ADD_EXPR_TERM: u64 = 8;
const OP_PUSH_NAMED_PARAM: u64 = 9;
const OP_PUSH_ASSIGN_IDENT: u64 = 10;
const OP_PUSH_TERM: u64 = 11;
const OP_END_FUNC_CALL_PARAM: u64 = 12;
const OP_GOTO: u64 = 50;

const OP_BINARY_MUL: u64 = 100;
const OP_BINARY_DIV: u64 = 101;
const OP_BINARY_MOD: u64 = 102;
const OP_BINARY_PLUS: u64 = 103;
const OP_BINARY_MINUS: u64 = 104;

// const OP_BINARY_EQ: u64 = 110;
// const OP_BINARY_NE: u64 = 111;
// const OP_BINARY_GE: u64 = 112;
// const OP_BINARY_LE: u64 = 113;
// const OP_BINARY_GT: u64 = 114;
// const OP_BINARY_LT: u64 = 115;

// const OP_BINARY_AND: u64 = 120;
// const OP_BINARY_OR: u64 = 121;
// const OP_BINARY_COALESCE: u64 = 122;

#[derive(Parser)]
#[command(name = "PRQL VM")]
#[command(author = "Massimo Meneghello <massimo.meneghello93@gmail.com>")]
#[command(version = "0.1.0")]
#[command(about = "Execute PRQL bytecode", long_about = None)]
struct Cli {
    #[arg(short, long, default_value_t = 0)]
    debug: u8,
    #[arg(short, long, default_value_t = 0)]
    verbosity: u8,
    #[arg(short, long, default_value_t = String::new())]
    input_file: String,
}

fn main() {
    let cli = Cli::parse();

    let mut vm = PRQLVirtualMachine::new();
    vm.__debug_level = cli.debug;
    vm.__verbosity_level = cli.verbosity;
    vm.input_file = cli.input_file;

    // vm.__debug_level = 10;
    // vm.__verbosity_level = 10;
    // vm.input_file = String::from("C:\\Users\\massi\\source\\repos\\prqlvs\\bytecode");

    if vm.__verbosity_level > 0 {
        println!("PRQL VM");
        println!("-------");
        println!("Debug level:     {}", vm.__debug_level);
        println!("Verbosity level: {}", vm.__verbosity_level);
        println!("Input file:      {}\n", vm.input_file);
    }

    let input = fs::read(&vm.input_file).unwrap();
    vm.read_prql_bytecode(&input);
}

pub struct Operation {
    opcode: u64,
    param1: u64,
    param2: u64,
}

pub struct PRQLVirtualMachine {
    __verbosity_level: u8,
    __debug_level: u8,
    __counter: u64,
    __current_directory: String,
    __current_table: DataFrame,
    __num_symbol_table: Vec<f64>,
    __str_symbol_table: Vec<String>,
    __stack: Vec<Operation>,
    __functions: HashMap<String, fn(&mut PRQLVirtualMachine)>,
    __variables: HashMap<String, DataFrame>,

    input_file: String,
}

impl PRQLVirtualMachine {
    pub fn new() -> PRQLVirtualMachine {
        let mut vm = PRQLVirtualMachine {
            __verbosity_level: 0,
            __debug_level: 0,
            __counter: 0,
            __current_directory: String::new(),

            __current_table: DataFrame::empty(),

            __num_symbol_table: Vec::new(),
            __str_symbol_table: Vec::new(),
            __stack: Vec::new(),
            __functions: HashMap::new(),
            __variables: HashMap::new(),

            input_file: String::new(),
        };

        // Load functions
        vm.__functions.insert(String::from("derive"), prql_derive);
        vm.__functions.insert(String::from("from"), prql_from);
        vm.__functions.insert(String::from("import"), prql_import);

        return vm;
    }

    pub fn read_prql_bytecode(&mut self, bytes: &[u8]) {
        /////////////////////////////////////////////////////////////
        ////                    PREAMBLE

        // check signature
        if bytes[0] != 0x11 || bytes[1] != 0x01 || bytes[2] != 0x19 || bytes[3] != 0x93 {
            panic!("Wrong bytecode format.")
        }

        // skip bytes 4 to 8 and read the number of elements
        // in the symbol table
        let mut buff = [
            bytes[8], bytes[9], bytes[10], bytes[11], bytes[12], bytes[13], bytes[14], bytes[15],
        ];

        /////////////////////////////////////////////////////////////
        ////                    NUM SYMBOL TABLE
        let num_table_length: u64 = u64::from_be_bytes(buff);
        let mut offset: u64 = 16;
        for _ in 0..num_table_length {
            // copy length of string into the buffer
            for j in 0..8 {
                buff[j] = bytes[(offset as usize) + j];
            }

            self.__num_symbol_table.push(f64::from_ne_bytes(buff));
            offset += 8;
        }

        /////////////////////////////////////////////////////////////
        ////                    STR SYMBOL TABLE
        for j in 0..8 {
            buff[j] = bytes[(offset as usize) + j];
        }
        let str_table_length: u64 = u64::from_be_bytes(buff);
        offset += 8;
        for _ in 0..str_table_length {
            // copy length of string into the buffer
            for j in 0..8 {
                buff[j] = bytes[(offset as usize) + j];
            }

            let symbol_length: u64 = u64::from_be_bytes(buff);
            offset += 8;

            // read symbol and insert into the symbol table
            let res = str::from_utf8(&bytes[(offset as usize)..(offset + symbol_length) as usize])
                .unwrap();
            self.__str_symbol_table.push(res.to_string());

            offset += symbol_length;
        }

        if self.__debug_level > 5 {
            println!("BYTECODE");
            println!("--------");
            println!(
                "BYTE MARK:       0x{:x} 0x{:x} 0x{:x} 0x{:x}",
                bytes[0], bytes[1], bytes[2], bytes[3]
            );
            println!("NUMERIC SYMBOL NUM: {}\n", num_table_length);

            println!("NUMERIC SYMBOLS");
            println!("---------------");
            for symb in self.__num_symbol_table.iter() {
                println!("{}", symb)
            }
            println!();

            println!("STRING SYMBOL NUM: {}\n", str_table_length);

            println!("STRING SYMBOLS");
            println!("--------------");
            for symb in self.__str_symbol_table.iter() {
                println!("{}", symb)
            }
            println!();

            println!("OPERATIONS");
            println!("----------");
        }

        /////////////////////////////////////////////////////////////
        ////                    OPERATIONS
        while (offset as usize) < bytes.len() {
            for j in 0..8 {
                buff[j] = bytes[(offset as usize) + j];
            }
            let opcode: u64 = u64::from_be_bytes(buff);
            offset += 8;

            for j in 0..8 {
                buff[j] = bytes[(offset as usize) + j];
            }
            let param1: u64 = u64::from_be_bytes(buff);
            offset += 8;

            for j in 0..8 {
                buff[j] = bytes[(offset as usize) + j];
            }
            let param2: u64 = u64::from_be_bytes(buff);
            offset += 8;

            self.read_instruction(opcode, param1, param2)
        }
    }

    pub fn read_instruction(&mut self, opcode: u64, param1: u64, param2: u64) {
        match opcode {
            // PIPELINE
            OP_BEGIN_PIPELINE => {
                if self.__debug_level > 10 {
                    println!("{:<25} | {:<20} | {:<20}", "OP_BEGIN_PIPELINE", "", "");
                }
            }

            OP_END_PIPELINE => {
                if self.__debug_level > 10 {
                    println!("{:<25} | {:<20} | {:<20}", "OP_END_PIPELINE", "", "");
                }
            }

            OP_ASSIGN_TABLE => {}

            OP_BEGIN_FUNC_CALL => {
                let function_name = &self.__str_symbol_table[(param1 as usize)];
                if self.__debug_level > 10 {
                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_BEGIN_FUNC_CALL", function_name, ""
                    );
                }
            }

            OP_END_FUNC_CALL => {
                if self.__debug_level > 10 {
                    println!("{:<25} | {:<20} | {:<20}", "OP_END_FUNC_CALL", "", "");
                }
            }

            OP_BEGIN_LIST => {}
            OP_END_LIST => {}

            OP_ADD_FUNC_PARAM => {}
            OP_ADD_EXPR_TERM => {}

            OP_PUSH_NAMED_PARAM => {
                if self.__debug_level > 10 {
                    println!("{:<25} | {:<20} | {:<20}", "OP_PUSH_NAMED_PARAM", "", "");
                }
            }

            OP_PUSH_ASSIGN_IDENT => {}

            OP_PUSH_TERM => {
                if self.__debug_level > 10 {
                    let mut term_type_str = "UNKNOWN";
                    let mut term_val = "";
                    match param1 {
                        TYPE_NULL => {
                            term_type_str = "NULL";
                        }
                        TYPE_BOOL => {
                            term_type_str = "BOOL";
                            if param2 == 1 {
                                term_val = "true";
                            } else {
                                term_val = "false";
                            };
                        }
                        TYPE_NUMERIC => {
                            term_type_str = "NUMERIC";
                            // term_val = &self.__num_symbol_table[(param2 as usize)].to_string();
                        }
                        TYPE_STRING => {
                            term_type_str = "STRING";
                            term_val = &self.__str_symbol_table[(param2 as usize)];
                        }
                        TYPE_IDENT => {
                            term_type_str = "IDENT";
                            term_val = &self.__str_symbol_table[(param2 as usize)];
                        }
                        _ => {}
                    }

                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_PUSH_TERM", term_type_str, term_val
                    );
                }

                self.__stack.push(Operation {
                    opcode: opcode,
                    param1: param1,
                    param2: param2,
                })
            }

            OP_END_FUNC_CALL_PARAM => {
                if self.__debug_level > 10 {
                    println!("{:<25} | {:<20} | {:<20}", "OP_END_FUNC_CALL_PARAM", "", "");
                }
            }

            OP_GOTO => {}

            OP_BINARY_MUL => {
                let term2 = self.__stack.pop().unwrap();
                let term1 = self.__stack.pop().unwrap();

                let mut result = Operation {
                    opcode: OP_PUSH_TERM,
                    param1: 0,
                    param2: 0,
                };

                match term1.param1 {
                    TYPE_NULL => match term2.param2 {
                        TYPE_NULL => result.param1 = TYPE_NULL,
                        TYPE_BOOL => {
                            result.param1 = TYPE_BOOL;
                            result.param2 = 0;
                        }
                        TYPE_NUMERIC => {
                            result.param1 = TYPE_NUMERIC;

                            result.param2 = 0;
                        }
                        TYPE_STRING => result.param1 = TYPE_NULL,
                        TYPE_IDENT => result.param1 = TYPE_NULL,
                        _ => {}
                    },
                    TYPE_BOOL => match term2.param2 {
                        TYPE_NULL => result.param1 = TYPE_NULL,
                        TYPE_BOOL => {
                            result.param1 = TYPE_BOOL;
                            result.param2 = term1.param2 * term2.param2;
                        }
                        TYPE_NUMERIC => {}
                        TYPE_STRING => {}
                        TYPE_IDENT => {}
                        _ => {}
                    },
                    TYPE_NUMERIC => match term2.param2 {
                        TYPE_NULL => result.param1 = TYPE_NULL,
                        TYPE_BOOL => {}
                        TYPE_NUMERIC => {}
                        TYPE_STRING => {}
                        TYPE_IDENT => {}
                        _ => {}
                    },
                    TYPE_STRING => match term2.param2 {
                        TYPE_NULL => {
                            result.param1 = TYPE_STRING;
                            result.param2 = self.insert_string_symbol(String::from(""))
                        }
                        TYPE_BOOL => {}
                        TYPE_NUMERIC => {}
                        TYPE_STRING => {}
                        TYPE_IDENT => {}
                        _ => {}
                    },
                    TYPE_IDENT => match term2.param2 {
                        TYPE_NULL => {},
                        TYPE_BOOL => {}
                        TYPE_NUMERIC => {}
                        TYPE_STRING => {}
                        TYPE_IDENT => {}
                        _ => {}
                    },
                    _ => {}
                }

                self.__stack.push(result);
            }

            OP_BINARY_DIV => {}
            OP_BINARY_MOD => {}
            OP_BINARY_PLUS => {}
            OP_BINARY_MINUS => {}

            _ => println!("Unknown op code: {}", opcode),
        }
    }

    fn insert_string_symbol(&mut self, symb: String) -> u64 {
        let res = self.__str_symbol_table.iter().position(|r| r == &symb);
        if res.is_some() {
            return res.unwrap().to_u64().unwrap();
        }

        let l = self.__str_symbol_table.len();
        self.__str_symbol_table.push(symb);
        return l.to_u64().unwrap();

    }
}

pub fn prql_derive(vm: &mut PRQLVirtualMachine) {}

pub fn prql_from(vm: &mut PRQLVirtualMachine) {}

pub fn prql_import(vm: &mut PRQLVirtualMachine) {
    vm.__current_table = CsvReader::from_path("path.csv").unwrap().finish().unwrap();
}
