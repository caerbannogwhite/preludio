use polars::prelude::*;
use std::env;
use std::fs::File;
use std::io::BufReader;
use std::path::Path;
use std::str;

const OP_BEGIN_PIPELINE: i64 = 0;
const OP_END_PIPELINE: i64 = 1;
const OP_BEGIN_LIST: i64 = 4;
const OP_END_LIST: i64 = 5;
const OP_ADD_FUNC_PARAM: i64 = 6;
const OP_ADD_EXPR_TERM: i64 = 7;
const OP_CALL_FUNC: i64 = 8;

fn main() {
    let args: Vec<String> = env::args().collect();
    print!("{}", args[0]);

    let fp = File::open(Path::new(&args[0])).unwrap();
    let file = BufReader::new(&fp);

    let mut vm = PRQLVirtualMachine::new();
    vm.read_bytecode(file.buffer());
}

pub struct PRQLVirtualMachine {
    __verbose_level__: i64,
    __debug_level__: i64,
    __counter__: i64,
    __current_directory__: String,
    __current_table__: DataFrame,
    __symbol_table__: Vec<String>,
}

impl PRQLVirtualMachine {
    pub fn new() -> PRQLVirtualMachine {
        return PRQLVirtualMachine {
            __verbose_level__: 0,
            __debug_level__: 0,
            __counter__: 0,
            __current_directory__: String::new(),
            __current_table__: DataFrame::empty(),
            __symbol_table__: Vec::new(),
        };
    }

    pub fn read_bytecode(&mut self, bytes: &[u8]) {
        // check signature
        if bytes[0] != 0x11 || bytes[1] != 0x01 || bytes[2] != 0x19 || bytes[3] != 0x93 {}

        // skip bytes 4 to 8 and read the number of elements in the
        // symbol table
        let mut buff = [
            bytes[8], bytes[9], bytes[10], bytes[11], bytes[12], bytes[13], bytes[14], bytes[15],
        ];

        let table_length: i64 = i64::from_be_bytes(buff);
        let mut offset: i64 = 16;
        for i in 0..table_length {
            // copy length of string into the buffer
            for j in 0..8 {
                buff[j] = bytes[(offset as usize) + j];
            }

            let symbol_length: i64 = i64::from_be_bytes(buff);
            offset += 16;

            // read symbol and insert into the symbol table
            let res = str::from_utf8(&bytes[(offset as usize)..(offset + symbol_length) as usize])
                .unwrap();
            self.__symbol_table__.push(res.to_string());

            offset += symbol_length;
        }
    }

    pub fn push_instruction(&self, op_code: i64, param1: i64, param2: i64, param3: i64) {
        match op_code {
            OP_BEGIN_PIPELINE => {
                if self.__verbose_level__ > 10 {
                    println!("OP_BEGIN_PIPELINE");
                }
            }
            OP_END_PIPELINE => {
                if self.__verbose_level__ > 10 {
                    println!("OP_END_PIPELINE");
                }
            }
            _ => println!("Unknown op code: {}", op_code),
        }
    }

    pub fn run_binary() {}

    pub fn hello() {
        println!("hello from prqlvm")
    }
}
