use clap::Parser;
use polars::{functions, prelude::*};
use std::collections::HashMap;
use std::env;
use std::fs::File;
use std::io::BufReader;
use std::path::Path;
use std::str;

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
const OP_BEGIN_LIST: u64 = 4;
const OP_END_LIST: u64 = 5;
const OP_ADD_FUNC_PARAM: u64 = 6;
const OP_ADD_EXPR_TERM: u64 = 7;
const OP_CALL_FUNC: u64 = 8;
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

const OP_BINARY_EQ: u64 = 110;
const OP_BINARY_NE: u64 = 111;
const OP_BINARY_GE: u64 = 112;
const OP_BINARY_LE: u64 = 113;
const OP_BINARY_GT: u64 = 114;
const OP_BINARY_LT: u64 = 115;

const OP_BINARY_AND: u64 = 120;
const OP_BINARY_OR: u64 = 121;
const OP_BINARY_COALESCE: u64 = 122;

#[derive(Parser)]
#[command(author = "Massimo Meneghello", version, about, long_about = None)]
struct Cli {
    #[arg(short, long, action = clap::ArgAction::Count, default_value_t=0)]
    debug: u8,
    #[arg(short, long, action = clap::ArgAction::Count, default_value_t=0)]
    verbosity: u8,
    #[arg(short, long, default_value_t="")]
    input_file: String,
}

fn main() {
    let cli = Cli::parse();

    let mut vm = PRQLVirtualMachine::new();
    vm.__debug_level__ = cli.verbose;
    vm.__verbosity_level__ = cli.debug;

    let args: Vec<String> = env::args().collect();
    print!("{}", args[0]);

    let fp = File::open(Path::new(cli.input_file)).unwrap();
    let file = BufReader::new(&fp);

    vm.read_prql_bytecode(file.buffer());
}

pub struct PRQLVirtualMachine {
    __verbosity_level__: u8,
    __debug_level__: u8,
    __counter__: u64,
    __current_directory__: String,
    __current_table__: DataFrame,
    __symbol_table__: Vec<String>,
    __functions__: HashMap<String, fn(&mut PRQLVirtualMachine)>,
    __variables__: HashMap<String, DataFrame>,
}

impl PRQLVirtualMachine {
    pub fn new() -> PRQLVirtualMachine {
        return PRQLVirtualMachine {
            __verbosity_level__: 0,
            __debug_level__: 0,
            __counter__: 0,
            __current_directory__: String::new(),
            __current_table__: DataFrame::empty(),
            __symbol_table__: Vec::new(),
            __functions__: HashMap::new(),
            __variables__: HashMap::new(),
        };
    }

    pub fn read_prql_bytecode(&mut self, bytes: &[u8]) {
        // check signature
        if bytes[0] != 0x11 || bytes[1] != 0x01 || bytes[2] != 0x19 || bytes[3] != 0x93 {
            panic!("Wrong bytecode format.")
        }

        // skip bytes 4 to 8 and read the number of elements
        // in the symbol table
        let mut buff = [
            bytes[8], bytes[9], bytes[10], bytes[11], bytes[12], bytes[13], bytes[14], bytes[15],
        ];

        let table_length: u64 = u64::from_be_bytes(buff);
        let mut offset: u64 = 16;
        for i in 0..table_length {
            // copy length of string into the buffer
            for j in 0..8 {
                buff[j] = bytes[(offset as usize) + j];
            }

            let symbol_length: u64 = u64::from_be_bytes(buff);
            offset += 16;

            // read symbol and insert into the symbol table
            let res = str::from_utf8(&bytes[(offset as usize)..(offset + symbol_length) as usize])
                .unwrap();
            self.__symbol_table__.push(res.to_string());

            offset += symbol_length;
        }

        if self.__debug_level__ > 5 {
            println!(
                "BYTE MARK: {:x}{:x}{:x}{:x}",
                bytes[0], bytes[1], bytes[2], bytes[3]
            );
            println!("SYMBOL NUM: {}\n", table_length);

            println!("TABLE");
            println!("-----");
            for symb in self.__symbol_table__.iter() {
                println!("{}", symb)
            }
        }
    }

    pub fn push_instruction(&self, op_code: u64, param1: u64, param2: u64) {
        match op_code {
            // PIPELINE
            OP_BEGIN_PIPELINE => {
                if self.__verbosity_level__ > 10 {
                    println!("OP_BEGIN_PIPELINE");
                }
            }

            OP_END_PIPELINE => {
                if self.__verbosity_level__ > 10 {
                    println!("OP_END_PIPELINE");
                }
            }

            OP_ASSIGN_TABLE => {}
            // OP_BEGIN_LIST => {}
            // OP_END_LIST => {}
            OP_ADD_FUNC_PARAM => {}
            OP_ADD_EXPR_TERM => {}
            OP_CALL_FUNC => {}
            OP_PUSH_NAMED_PARAM => {}
            OP_PUSH_ASSIGN_IDENT => {}
            OP_PUSH_TERM => {}
            OP_END_FUNC_CALL_PARAM => {}
            // OP_GOTO => {}
            _ => println!("Unknown op code: {}", op_code),
        }
    }
}

fn prql_derive(vm: &mut PRQLVirtualMachine) {}

fn prql_from(vm: &mut PRQLVirtualMachine) {}

fn prql_import(vm: &mut PRQLVirtualMachine) {
    vm.__current_table__ = CsvReader::from_path("iris_csv")
}
