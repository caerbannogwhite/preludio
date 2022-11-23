use clap::Parser;
use internal::PrqlInternalDim;
use polars::export::num::ToPrimitive;
use std::collections::HashMap;
use std::fs;
use std::str;

mod internal;
mod prql_std;
mod type_system;
mod vm;

const TERM_NULL: u16 = 0;
const TERM_BOOL: u16 = 1;
const TERM_NUMERIC: u16 = 2;
const TERM_STRING: u16 = 3;
// const TERM_INTERVAL: u16 = 5;
// const TERM_RANGE: u16 = 6;
// const TERM_LIST: u16 = 7;
// const TERM_PIPELINE: u16 = 8;
const TERM_IDENT: u16 = 10;

const OP_BEGIN_PIPELINE: u16 = 0;
const OP_END_PIPELINE: u16 = 1;
const OP_ASSIGN_TABLE: u16 = 2;
// const OP_BEGIN_FUNC_CALL: u16 = 3;
const OP_MAKE_FUNC_CALL: u16 = 4;
const OP_BEGIN_LIST: u16 = 5;
const OP_END_LIST: u16 = 6;
const OP_ADD_FUNC_PARAM: u16 = 7;
const OP_ADD_EXPR_TERM: u16 = 8;
const OP_PUSH_NAMED_PARAM: u16 = 9;
const OP_PUSH_ASSIGN_IDENT: u16 = 10;
const OP_PUSH_TERM: u16 = 11;
const OP_END_FUNC_CALL_PARAM: u16 = 12;
const OP_GOTO: u16 = 50;

const OP_BINARY_MUL: u16 = 100;
const OP_BINARY_DIV: u16 = 101;
const OP_BINARY_MOD: u16 = 102;
const OP_BINARY_ADD: u16 = 103;
const OP_BINARY_SUB: u16 = 104;

// const OP_BINARY_EQ: u16 = 110;
// const OP_BINARY_NE: u16 = 111;
// const OP_BINARY_GE: u16 = 112;
// const OP_BINARY_LE: u16 = 113;
// const OP_BINARY_GT: u16 = 114;
// const OP_BINARY_LT: u16 = 115;

// const OP_BINARY_AND: u16 = 120;
// const OP_BINARY_OR: u16 = 121;
// const OP_BINARY_COALESCE: u16 = 122;

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
        println!("=======");
        println!("Debug level:     {}", vm.__debug_level);
        println!("Verbosity level: {}", vm.__verbosity_level);
        println!("Input file:      {}\n", vm.input_file);
    }

    let input = fs::read(&vm.input_file).unwrap();
    vm.read_prql_bytecode(&input);
}

pub struct FunctionParam {
    type_: u16,
    value: [u8; 8],
}

pub struct PRQLVirtualMachine {
    __verbosity_level: u8,
    __debug_level: u8,
    __counter: u64,
    __current_directory: String,
    __current_result: Option<internal::PrqlInternal>,
    __current_temp_column: String,
    __symbol_table: Vec<String>,
    __stack: Vec<internal::PrqlInternal>,
    __function_num_params: u64,
    __functions: HashMap<String, fn(&mut PRQLVirtualMachine)>,
    __variables: HashMap<String, internal::PrqlInternal>,

    input_file: String,
}

impl PRQLVirtualMachine {
    pub fn new() -> PRQLVirtualMachine {
        let mut vm = PRQLVirtualMachine {
            __verbosity_level: 0,
            __debug_level: 0,
            __counter: 0,
            __current_directory: String::new(),

            __current_result: std::option::Option::None,
            __current_temp_column: String::from(""),

            __symbol_table: Vec::new(),
            __stack: Vec::new(),
            __function_num_params: 0,
            __functions: HashMap::new(),
            __variables: HashMap::new(),

            input_file: String::new(),
        };

        // Load functions
        vm.__functions
            .insert(String::from("derive"), prql_std::prql_derive);
        vm.__functions
            .insert(String::from("from"), prql_std::prql_from);
        vm.__functions
            .insert(String::from("import_csv"), prql_std::prql_import_csv);
        vm.__functions
            .insert(String::from("export_csv"), prql_std::prql_export_csv);
        vm.__functions
            .insert(String::from("select"), prql_std::prql_select);

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
        let mut buff_8 = [
            bytes[8], bytes[9], bytes[10], bytes[11], bytes[12], bytes[13], bytes[14], bytes[15],
        ];

        /////////////////////////////////////////////////////////////
        ////                    SYMBOL TABLE
        let table_length: u64 = u64::from_be_bytes(buff_8);
        let mut offset: u64 = 16;
        for _ in 0..table_length {
            // copy length of string into the buffer
            for j in 0..8 {
                buff_8[j] = bytes[(offset as usize) + j];
            }

            let symbol_length: u64 = u64::from_be_bytes(buff_8);
            offset += 8;

            // read symbol and insert into the symbol table
            let res = str::from_utf8(&bytes[(offset as usize)..(offset + symbol_length) as usize])
                .unwrap();
            self.__symbol_table.push(res.to_string());

            offset += symbol_length;
        }

        if self.__debug_level > 5 {
            println!("BYTECODE");
            println!("========");
            println!(
                "BYTE MARK:       0x{:x} 0x{:x} 0x{:x} 0x{:x}",
                bytes[0], bytes[1], bytes[2], bytes[3]
            );
            println!("STRING SYMBOL NUM: {}\n", table_length);

            println!("STRING SYMBOLS");
            println!("==============");
            for symb in self.__symbol_table.iter() {
                println!("{}", symb)
            }
            println!();

            println!("OPERATIONS");
            println!("==========");
        }

        /////////////////////////////////////////////////////////////
        ////                    OPERATIONS
        let mut buff_2 = [0, 0];
        while (offset as usize) < bytes.len() {
            for j in 0..2 {
                buff_2[j] = bytes[(offset as usize) + j];
            }
            let opcode: u16 = u16::from_be_bytes(buff_2);
            offset += 2;

            for j in 0..2 {
                buff_2[j] = bytes[(offset as usize) + j];
            }
            let param1: u16 = u16::from_be_bytes(buff_2);
            offset += 2;

            for j in 0..8 {
                buff_8[j] = bytes[(offset as usize) + j];
            }
            offset += 8;

            self.read_instruction(opcode, param1, buff_8)
        }
    }

    pub fn read_instruction(&mut self, opcode: u16, param1: u16, param2: [u8; 8]) {
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

            // OP_BEGIN_FUNC_CALL => {
            //     if self.__debug_level > 10 {
            //         println!("{:<25} | {:<20} | {:<20}", "OP_BEGIN_FUNC_CALL", "", "");
            //     }
            // }
            OP_MAKE_FUNC_CALL => {
                let function_name = &self.__symbol_table[(u64::from_be_bytes(param2) as usize)];
                if self.__debug_level > 10 {
                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_MAKE_FUNC_CALL", "", function_name
                    );
                }

                if self.__functions.contains_key(function_name) {
                    self.__functions[function_name](self);
                } else {
                    println!("[😞] Error - Function {} not found.", function_name);
                }

                self.__function_num_params = 0;
            }

            OP_BEGIN_LIST => {}
            OP_END_LIST => {}

            OP_ADD_FUNC_PARAM => {}
            OP_ADD_EXPR_TERM => {}

            OP_PUSH_NAMED_PARAM => {
                let param_name = self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();
                if self.__debug_level > 10 {
                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_PUSH_NAMED_PARAM", "", param_name
                    );
                }

                self.__stack
                    .push(internal::PrqlInternal::new_param_name(param_name));
            }

            OP_PUSH_ASSIGN_IDENT => {
                let ident = self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();
                if self.__debug_level > 10 {
                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_PUSH_ASSIGN_IDENT", "", ident
                    );
                }

                self.__stack
                    .push(internal::PrqlInternal::new_assign_ident(ident));
            }

            OP_PUSH_TERM => {
                let mut term_type_str = String::from("UNKNOWN");
                let mut term_val = String::from("");
                match param1 {
                    TERM_NULL => {
                        term_type_str = String::from("NULL");
                        self.__stack
                            .push(internal::PrqlInternal::new().with_scalar_null());
                    }
                    TERM_BOOL => {
                        term_type_str = String::from("BOOL");
                        term_val = String::from("true");
                        let mut bool_val = true;
                        if param2[7] == 0 {
                            term_val = String::from("false");
                            bool_val = false;
                        }

                        self.__stack
                            .push(internal::PrqlInternal::new().with_scalar_bool(bool_val));
                    }
                    TERM_NUMERIC => {
                        term_type_str = String::from("NUMERIC");
                        let num_val = f64::from_le_bytes(param2);
                        term_val = num_val.to_string();

                        self.__stack
                            .push(internal::PrqlInternal::new().with_scalar_numeric(num_val));
                    }
                    TERM_STRING => {
                        term_type_str = String::from("STRING");
                        term_val =
                            self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();

                        self.__stack
                            .push(internal::PrqlInternal::new().with_scalar_string(term_val));
                    }
                    TERM_IDENT => {
                        term_type_str = String::from("IDENT");
                        term_val =
                            self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();

                        self.__stack.push(self.__ident_resolution(term_val));
                    }
                    _ => {}
                }

                if self.__debug_level > 10 {
                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_PUSH_TERM", term_type_str, term_val
                    );
                }
            }

            OP_END_FUNC_CALL_PARAM => {
                if self.__debug_level > 10 {
                    println!("{:<25} | {:<20} | {:<20}", "OP_END_FUNC_CALL_PARAM", "", "");
                }
                self.__function_num_params += 1;
            }

            OP_GOTO => {}

            /////////////////////////////////////////////////////////
            ////                MULTIPLICATION
            OP_BINARY_MUL => {
                let term2 = self.__stack.pop().unwrap();
                // let term1 = self.__stack.pop().unwrap();

                self.__stack.last().unwrap().arith_binary_mul(term2);
            }

            /////////////////////////////////////////////////////////
            ////                DIVISION
            OP_BINARY_DIV => {
                let term2 = self.__stack.pop().unwrap();
                // let term1 = self.__stack.pop().unwrap();

                self.__stack.last().unwrap().arith_binary_div(term2);
            }

            /////////////////////////////////////////////////////////
            ////                MODULUS
            OP_BINARY_MOD => {
                let term2 = self.__stack.pop().unwrap();
                // let term1 = self.__stack.pop().unwrap();

                self.__stack.last().unwrap().arith_binary_mod(term2);
            }

            /////////////////////////////////////////////////////////
            ////                ADDITION
            OP_BINARY_ADD => {
                let term2 = self.__stack.pop().unwrap();
                // let term1 = self.__stack.pop().unwrap();

                self.__stack.last().unwrap().arith_binary_add(term2);
            }

            /////////////////////////////////////////////////////////
            ////                SUBTRACTION
            OP_BINARY_SUB => {
                let term2 = self.__stack.pop().unwrap();
                // let term1 = self.__stack.pop().unwrap();

                self.__stack.last().unwrap().arith_binary_sub(term2);
            }

            _ => println!("[💣] Byte-Code Error - Unknown op code: {}", opcode),
        }
    }

    fn __ident_resolution(&self, ident: String) -> internal::PrqlInternal {
        // Look at the current result first
        if self.__current_result.is_some() {
            let res = self.__current_result.as_ref().unwrap();

            match res.get_dim() {
                PrqlInternalDim::Scalar => {
                    return internal::PrqlInternal::new_error(format!(
                        "Cannot find symbol {}",
                        ident
                    ));
                }
                PrqlInternalDim::Series => {
                    return internal::PrqlInternal::new_error(format!(
                        "Cannot find symbol {}",
                        ident
                    ));
                }
                PrqlInternalDim::Table => {
                    if res
                        .get_data_frame()
                        .get_column_names()
                        .contains(&ident.as_str())
                    {
                        return internal::PrqlInternal::new()
                            .with_series(res.get_data_frame().select([ident]).unwrap());
                    }
                    return internal::PrqlInternal::new_error(format!(
                        "Cannot find symbol {}",
                        ident
                    ));
                }
            }
        }

        // If no result, look at the variables in the environment
        if self.__variables.contains_key(&ident) {
            return self.__variables[&ident];
        }

        return internal::PrqlInternal::new_error(format!("Cannot find symbol {}", ident));
    }

    pub fn __float_to_param(f: f64) -> u64 {
        let bytes = f.to_ne_bytes();
        u64::from_be_bytes(bytes)
    }

    pub fn __param_to_float(p: u64) -> f64 {
        let bytes = p.to_be_bytes();
        f64::from_bits(u64::from_be_bytes(bytes))
    }

    fn __insert_symbol(&mut self, symb: String) -> [u8; 8] {
        let res = self.__symbol_table.iter().position(|r| r == &symb);
        if res.is_some() {
            return u64::to_be_bytes(res.unwrap().to_u64().unwrap());
        }

        let l = self.__symbol_table.len();
        self.__symbol_table.push(symb);
        return u64::to_be_bytes(l.to_u64().unwrap());
    }

    fn __read_params(
        &mut self,
        position_params: &mut Vec<internal::PrqlInternal>,
        named_params: &mut HashMap<String, internal::PrqlInternal>,
    ) {
        let mut counter: u64 = 0;
        while counter < self.__function_num_params {
            let term = self.__stack.pop().unwrap();
            match term.get_tag() {
                internal::PrqlInternalTag::ExprTerm => {
                    position_params.push(term);
                    counter += 1;
                }
                internal::PrqlInternalTag::ParamName => {
                    let val = self.__stack.pop().unwrap();
                    named_params.insert(term.get_name(), val);
                    counter += 1;
                }
                internal::PrqlInternalTag::AssignIdent => {}
                internal::PrqlInternalTag::Error => {}
            }
        }
    }
}
