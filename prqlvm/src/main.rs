use polars::prelude::*;

const OP_BEGIN_PIPELINE: i8 = 0;
const OP_END_PIPELINE: i8 = 1;
const OP_BEGIN_LIST: i8 = 4;
const OP_END_LIST: i8 = 5;
const OP_ADD_FUNC_PARAM: i8 = 6;
const OP_ADD_EXPR_TERM: i8 = 7;
const OP_CALL_FUNC: i8 = 8;

fn main() {
    println!("Hello, world!");
}

struct VMStatus {
    __verbose_level__: i8,
    __debug_level__: i8,
    __counter__: i64,
    __current_directory__: String,
    __current_table__: DataFrame,
    __symbol_table__: Vec<String>,
}

impl VMStatus {
    pub fn read_prql_bytecode(&self) {}

    pub fn push_instruction(&self, op_code: i8, param1: i8, param2: i128, param3: String) {
        match op_code {
            OP_BEGIN_PIPELINE => {
                if (self.__verbose_level__ > 10) {
                    println!("OP_BEGIN_PIPELINE");
                }
            }
            OP_END_PIPELINE => {
                if (self.__verbose_level__ > 10) {
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
