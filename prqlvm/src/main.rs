use clap::Parser;
use polars::export::num::ToPrimitive;
use polars::prelude::*;
use std::collections::HashMap;
use std::fmt::Display;
use std::fmt::Formatter;
use std::fs;
use std::fs::File;
use std::str;

mod prql_std;
mod prql_vm;

const TERM_NULL: u16 = 0;
const TERM_BOOL: u16 = 1;
const TERM_NUMERIC: u16 = 2;
const TERM_STRING: u16 = 3;
const TERM_INTERVAL: u16 = 5;
const TERM_RANGE: u16 = 6;
const TERM_LIST: u16 = 7;
const TERM_PIPELINE: u16 = 8;
const TERM_IDENT: u16 = 10;

const TYPE_COLUMN_NULL: u16 = 20;
const TYPE_COLUMN_BOOL: u16 = 21;
const TYPE_COLUMN_NUMERIC: u16 = 22;
const TYPE_COLUMN_STRING: u16 = 23;

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
const OP_BINARY_PLUS: u16 = 103;
const OP_BINARY_MINUS: u16 = 104;

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

#[derive(Clone, Debug)]
pub enum PrqlDataType {
    Null,
    Bool,
    Numeric,
    String,
}

impl Default for PrqlDataType {
    fn default() -> Self {
        PrqlDataType::Null
    }
}

impl PrqlDataType {
    /// Convert to the physical data type
    #[must_use]
    pub fn to_physical(&self) -> DataType {
        use PrqlDataType::*;
        match self {
            Null => DataType::Null,
            Bool => DataType::Boolean,
            Numeric => DataType::Float64,
            String => DataType::Utf8,
        }
    }

    pub fn polars_to_prql(dtype: &DataType) -> PrqlDataType {
        match dtype {
            DataType::Null => PrqlDataType::Null,
            DataType::Unknown => PrqlDataType::Null,
            DataType::Boolean => PrqlDataType::Bool,
            DataType::UInt8 => PrqlDataType::Numeric,
            DataType::UInt16 => PrqlDataType::Numeric,
            DataType::UInt32 => PrqlDataType::Numeric,
            DataType::UInt64 => PrqlDataType::Numeric,
            DataType::Int8 => PrqlDataType::Numeric,
            DataType::Int16 => PrqlDataType::Numeric,
            DataType::Int32 => PrqlDataType::Numeric,
            DataType::Int64 => PrqlDataType::Numeric,
            DataType::Float32 => PrqlDataType::Numeric,
            DataType::Float64 => PrqlDataType::Numeric,
            DataType::Utf8 => PrqlDataType::String,
            _ => PrqlDataType::Null,
        }
    }

    pub fn prql_to_polars(self) -> DataType {
        match self {
            PrqlDataType::Null => DataType::Null,
            PrqlDataType::Bool => DataType::Boolean,
            PrqlDataType::Numeric => DataType::Float64,
            PrqlDataType::String => DataType::Utf8,
        }
    }
}

impl Display for PrqlDataType {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        let s = match self {
            PrqlDataType::Null => "Null",
            PrqlDataType::Bool => "Bool",
            PrqlDataType::Numeric => "Numeric",
            PrqlDataType::String => "String",
        };
        f.write_str(s)
    }
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

pub struct Operation {
    opcode: u16,
    param1: u16,
    param2: [u8; 8],
}

pub struct PRQLVirtualMachine {
    __verbosity_level: u8,
    __debug_level: u8,
    __counter: u64,
    __current_directory: String,
    __current_table: DataFrame,
    __current_temp_column: String,
    __symbol_table: Vec<String>,
    __stack: Vec<Operation>,
    __function_num_params: u64,
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

            __current_table: DataFrame::default(),
            __current_temp_column: String::from(""),

            __symbol_table: Vec::new(),
            __stack: Vec::new(),
            __function_num_params: 0,
            __functions: HashMap::new(),
            __variables: HashMap::new(),

            input_file: String::new(),
        };

        // Load functions
        vm.__functions.insert(String::from("derive"), prql_derive);
        vm.__functions.insert(String::from("from"), prql_from);
        vm.__functions.insert(String::from("import"), prql_import);
        vm.__functions.insert(String::from("export"), prql_export);
        vm.__functions.insert(String::from("select"), prql_select);

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
                if self.__debug_level > 10 {
                    let param_name =
                        self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();
                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_PUSH_NAMED_PARAM", "", param_name
                    );
                }

                self.__stack.push(Operation {
                    opcode: OP_PUSH_NAMED_PARAM,
                    param1: 0,
                    param2: param2,
                });
            }

            OP_PUSH_ASSIGN_IDENT => {
                if self.__debug_level > 10 {
                    let param_name =
                        self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();
                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_PUSH_ASSIGN_IDENT", "", param_name
                    );
                }

                self.__stack.push(Operation {
                    opcode: OP_PUSH_ASSIGN_IDENT,
                    param1: 0,
                    param2: param2,
                });
            }

            OP_PUSH_TERM => {
                if self.__debug_level > 10 {
                    let mut term_tERM_str = String::from("UNKNOWN");
                    let mut term_val = String::from("");
                    match param1 {
                        TERM_NULL => {
                            term_tERM_str = String::from("NULL");
                        }
                        TERM_BOOL => {
                            term_tERM_str = String::from("BOOL");
                            if param2[7] == 1 {
                                term_val = String::from("true");
                            } else {
                                term_val = String::from("false");
                            };
                        }
                        TERM_NUMERIC => {
                            term_tERM_str = String::from("NUMERIC");
                            term_val = f64::from_le_bytes(param2).to_string();
                        }
                        TERM_STRING => {
                            term_tERM_str = String::from("STRING");
                            term_val =
                                self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();
                        }
                        TERM_IDENT => {
                            term_tERM_str = String::from("IDENT");
                            term_val =
                                self.__symbol_table[(u64::from_be_bytes(param2) as usize)].clone();
                        }
                        _ => {}
                    }

                    println!(
                        "{:<25} | {:<20} | {:<20}",
                        "OP_PUSH_TERM", term_tERM_str, term_val
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
                self.__function_num_params += 1;
            }

            OP_GOTO => {}

            /////////////////////////////////////////////////////////
            ////                MULTIPLICATION
            OP_BINARY_MUL => {
                let term2 = self.__stack.pop().unwrap();
                let term1 = self.__stack.pop().unwrap();

                let mut result = Operation {
                    opcode: OP_PUSH_TERM,
                    param1: 0,
                    param2: [0, 0, 0, 0, 0, 0, 0, 0],
                };

                match term1.param1 {
                    TERM_NULL => match term2.param1 {
                        TERM_NULL => result.param1 = TERM_NULL,
                        TERM_BOOL => result.param1 = TERM_BOOL,
                        TERM_NUMERIC => {
                            result.param1 = TERM_NUMERIC;
                            result.param2 = f64::to_le_bytes(0.0);
                        }
                        TERM_STRING => {
                            result.param1 = TERM_STRING;
                            result.param2 = self.__insert_symbol(String::from(""));
                        }
                        TERM_IDENT => {
                            let col_name = self.__symbol_table
                                [(u64::from_be_bytes(term2.param2) as usize)]
                                .clone();

                            let tmp_col_name = "tmp";
                            result.param1 = TERM_IDENT;
                            result.param2 = self.__insert_symbol(String::from(tmp_col_name));

                            // match self
                            //     .__current_table
                            //     .schema()
                            //     .unwrap()
                            //     .get(&col_name)
                            //     .unwrap()
                            // {
                            //     DataType::Null => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| Ok(s),
                            //                         GetOutput::from_type(DataType::Null),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Boolean => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| Ok(s * 0),
                            //                         GetOutput::from_type(DataType::Boolean),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Float64 => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| Ok(s * 0),
                            //                         GetOutput::from_type(DataType::Float64),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Utf8 => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| {
                            //                             s.len();
                            //                             Ok(Series::new("", &[""]))
                            //                         },
                            //                         GetOutput::from_type(DataType::Utf8),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     _ => {}
                            // }
                        }
                        _ => {}
                    },
                    TERM_BOOL => match term2.param1 {
                        TERM_NULL => result.param1 = TERM_BOOL,
                        TERM_BOOL => {
                            result.param1 = TERM_BOOL;
                            result.param2[7] = term1.param2[7] * term2.param2[7];
                        }
                        TERM_NUMERIC => {
                            result.param1 = TERM_NUMERIC;
                            if term1.param2[7] == 1 {
                                result.param2 = term2.param2;
                            } else {
                                result.param2 = f64::to_le_bytes(0.0);
                            }
                        }
                        TERM_STRING => {
                            result.param1 = TERM_STRING;
                            if term1.param2[7] == 1 {
                                result.param2 = term2.param2;
                            } else {
                                result.param2 = self.__insert_symbol(String::from(""));
                            }
                        }
                        TERM_IDENT => {
                            // let col_name = self.__symbol_table
                            //     [(u64::from_be_bytes(term2.param2) as usize)]
                            //     .clone();

                            // let tmp_col_name = "tmp";
                            // result.param1 = TERM_IDENT;
                            // result.param2 = self.__insert_symbol(String::from(tmp_col_name));

                            // let value = term1.param2[7] as i64;
                            // match self
                            //     .__current_table
                            //     .schema()
                            //     .unwrap()
                            //     .get(&col_name)
                            //     .unwrap()
                            // {
                            //     DataType::Null => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| Ok(s),
                            //                         GetOutput::from_type(DataType::Boolean),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Boolean => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| Ok(s * value),
                            //                         GetOutput::from_type(DataType::Boolean),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Float64 => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| Ok(s * value),
                            //                         GetOutput::from_type(DataType::Float64),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Utf8 => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| {
                            //                             s.len();
                            //                             Ok(Series::new("", &[""]))
                            //                         },
                            //                         GetOutput::from_type(DataType::Utf8),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     _ => {}
                            // }
                        }
                        _ => {}
                    },
                    TERM_NUMERIC => match term2.param1 {
                        TERM_NULL => {
                            result.param1 = TERM_NUMERIC;
                            result.param2 = f64::to_le_bytes(0.0);
                        }
                        TERM_BOOL => {
                            result.param1 = TERM_NUMERIC;
                            if term2.param2[7] == 1 {
                                result.param2 = term1.param2;
                            } else {
                                result.param2 = f64::to_le_bytes(0.0);
                            }
                        }
                        TERM_NUMERIC => {
                            result.param1 = TERM_NUMERIC;
                            result.param2 = f64::to_le_bytes(
                                f64::from_le_bytes(term1.param2) * f64::from_le_bytes(term2.param2),
                            );
                        }
                        TERM_STRING => {
                            result.param1 = TERM_STRING;
                            result.param2 = self.__insert_symbol(
                                self.__symbol_table[(u64::from_be_bytes(term2.param2) as usize)]
                                    .repeat(f64::from_le_bytes(term1.param2).to_usize().unwrap()),
                            );
                        }
                        TERM_IDENT => {}
                        _ => {}
                    },
                    TERM_STRING => match term2.param1 {
                        TERM_NULL => {
                            result.param1 = TERM_STRING;
                            result.param2 = self.__insert_symbol(String::from(""));
                        }
                        TERM_BOOL => {
                            result.param1 = TERM_STRING;
                            if term2.param2[7] == 1 {
                                result.param2 = term1.param2;
                            } else {
                                result.param2 = self.__insert_symbol(String::from(""));
                            }
                        }
                        TERM_NUMERIC => {
                            result.param1 = TERM_STRING;
                            result.param2 = self.__insert_symbol(
                                self.__symbol_table[(u64::from_be_bytes(term1.param2) as usize)]
                                    .repeat(f64::from_le_bytes(term2.param2).to_usize().unwrap()),
                            );
                        }
                        TERM_STRING => {}
                        TERM_IDENT => {}
                        _ => {}
                    },
                    TERM_IDENT => match term2.param1 {
                        TERM_NULL => {}
                        TERM_BOOL => {}
                        TERM_NUMERIC => {
                            // let col_name = self.__symbol_table
                            //     [(u64::from_be_bytes(term1.param2) as usize)]
                            //     .clone();

                            // let tmp_col_name = "tmp";
                            // result.param1 = TERM_IDENT;
                            // result.param2 = self.__insert_symbol(String::from(tmp_col_name));

                            // let value = term1.param2[7] as i64;
                            // match self
                            //     .__current_table
                            //     .schema()
                            //     .unwrap()
                            //     .get(&col_name)
                            //     .unwrap()
                            // {
                            //     DataType::Null => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| Ok(s),
                            //                         GetOutput::from_type(DataType::Boolean),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Boolean => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         move |s| Ok(s * value),
                            //                         GetOutput::from_type(DataType::Boolean),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Float64 => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         move |s| Ok(s * value),
                            //                         GetOutput::from_type(DataType::Float64),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     DataType::Utf8 => {
                            //         self.__current_table =
                            //             self.__current_table.clone().with_column(
                            //                 col(&col_name)
                            //                     .map(
                            //                         |s| {
                            //                             s.len();
                            //                             Ok(Series::new("", &[""]))
                            //                         },
                            //                         GetOutput::from_type(DataType::Utf8),
                            //                     )
                            //                     .alias(tmp_col_name),
                            //             );
                            //     }
                            //     _ => {}
                            // }
                        }
                        TERM_STRING => {}
                        TERM_IDENT => {}
                        _ => {}
                    },
                    _ => {}
                }

                println!(
                    "{}",
                    self.__current_table
                        .clone()
                        .lazy()
                        .to_owned()
                        .limit(5)
                        .collect()
                        .unwrap()
                );

                self.__stack.push(result);
            }

            /////////////////////////////////////////////////////////
            ////                DIVISION
            OP_BINARY_DIV => {}

            /////////////////////////////////////////////////////////
            ////                MODULUS
            OP_BINARY_MOD => {}

            /////////////////////////////////////////////////////////
            ////                ADDITION
            OP_BINARY_PLUS => {
                let term2 = self.__stack.pop().unwrap();
                let term1 = self.__stack.pop().unwrap();

                let mut result = Operation {
                    opcode: OP_PUSH_TERM,
                    param1: 0,
                    param2: [0, 0, 0, 0, 0, 0, 0, 0],
                };

                match term1.param1 {
                    TERM_NULL => match term2.param1 {
                        TERM_NULL => result.param1 = TERM_NULL,
                        TERM_BOOL => {
                            result.param1 = TERM_BOOL;
                            result.param2 = term2.param2;
                        }
                        TERM_NUMERIC => {
                            result.param1 = TERM_NUMERIC;
                            result.param2 = term2.param2;
                        }
                        TERM_STRING => {
                            result.param1 = TERM_STRING;
                            result.param2 = term2.param2;
                        }
                        TERM_IDENT => {}
                        _ => {}
                    },
                    TERM_BOOL => {}
                    TERM_NUMERIC => {}
                    TERM_STRING => match term2.param1 {
                        TERM_NULL => {
                            result.param1 = TERM_STRING;
                            result.param2 = term1.param2;
                        }
                        TERM_BOOL => {}
                        TERM_NUMERIC => {}
                        TERM_STRING => {
                            result.param1 = TERM_STRING;
                            result.param2 = self.__insert_symbol(
                                self.__symbol_table[(u64::from_be_bytes(term1.param2) as usize)]
                                    .clone()
                                    + &self.__symbol_table
                                        [(u64::from_be_bytes(term2.param2) as usize)],
                            );
                        }
                        TERM_IDENT => {}
                        _ => {}
                    },
                    TERM_IDENT => {}
                    _ => {}
                }

                self.__stack.push(result);
            }

            /////////////////////////////////////////////////////////
            ////                SUBTRACTION
            OP_BINARY_MINUS => {}

            _ => println!("[💣] Byte-Code Error - Unknown op code: {}", opcode),
        }
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
        position_params: &mut Vec<FunctionParam>,
        named_params: &mut HashMap<String, FunctionParam>,
    ) {
        let mut counter: u64 = 0;
        while counter < self.__function_num_params {
            let op = self.__stack.pop().unwrap();
            match op.opcode {
                OP_PUSH_TERM => {
                    position_params.push(FunctionParam {
                        type_: op.param1,
                        value: op.param2,
                    });
                    counter += 1;
                }
                OP_PUSH_NAMED_PARAM => {
                    let term = self.__stack.pop().unwrap();
                    named_params.insert(
                        self.__symbol_table[u64::from_be_bytes(op.param2) as usize].clone(),
                        FunctionParam {
                            type_: term.param1,
                            value: term.param2,
                        },
                    );
                    counter += 1;
                }
                _ => {}
            }
        }
    }
}

pub fn prql_derive(vm: &mut PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING derive");
    }
}

pub fn prql_from(vm: &mut PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING from");
    }
}

pub fn prql_import(vm: &mut PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING import");
    }

    let help_message = "
    import function
    ===============

    (path)    - the path of the input file
    
    enc       - [ Utf8 | LossyUtf8 ]

    type      - [ csv | json ]
                  ^^^
    delimiter - \",\" the file delimiter
                  ^
    skip      - 0 number of rows to skip
                ^  
    ";

    let mut position_params: Vec<FunctionParam> = Vec::new();
    let mut named_params: HashMap<String, FunctionParam> = HashMap::new();

    vm.__read_params(&mut position_params, &mut named_params);

    let path = vm.__symbol_table[u64::from_be_bytes(position_params[0].value) as usize].as_str();

    let mut input_file_type = String::from("csv");
    if named_params.contains_key("type") {
        input_file_type =
            vm.__symbol_table[u64::from_be_bytes(named_params["type"].value) as usize].clone();
    }

    let mut delimiter: u8 = ',' as u8;
    if named_params.contains_key("delimiter") {
        delimiter = vm.__symbol_table[u64::from_be_bytes(named_params["delimiter"].value) as usize]
            .as_bytes()[0];
    }

    let mut enc_str = String::from("utf8");
    if named_params.contains_key("enc") {
        enc_str = vm.__symbol_table[u64::from_be_bytes(named_params["enc"].value) as usize]
            .to_lowercase();
    }

    let mut skip_rows = 0;
    if named_params.contains_key("skip") {
        skip_rows = f64::from_ne_bytes(named_params["enc"].value) as usize
    }

    match input_file_type.as_str() {
        "csv" => {
            let mut enc = CsvEncoding::Utf8;
            match enc_str.as_str() {
                "lossyutf8" => enc = CsvEncoding::LossyUtf8,
                _ => {}
            }

            // Read the first 10 rows to get the schema
            // Then, coerce the schema to to the PRQL data types
            let head = CsvReader::from_path(path)
                .unwrap()
                .with_delimiter(delimiter)
                .with_encoding(enc)
                .with_skip_rows(skip_rows)
                .with_n_rows(Some(10))
                // .with_quote_char(quote)
                .finish()
                .unwrap();

            let mut schema = Schema::new();
            for col in head.schema().iter() {
                schema.with_column(
                    col.0.to_string(),
                    PrqlDataType::polars_to_prql(col.1).to_physical(),
                );
            }

            vm.__current_table = CsvReader::from_path(path)
                .unwrap()
                .with_delimiter(delimiter)
                .with_encoding(enc)
                .with_skip_rows(skip_rows)
                // .with_n_rows(10)
                .with_dtypes(Some(&schema))
                // .with_quote_char(quote)
                .finish()
                .unwrap();

            for col in vm.__current_table.schema().iter() {
                println!("{} {}", col.0, col.1);
            }
        }
        "json" => {}
        _ => {}
    }
}

pub fn prql_new(vm: &mut PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING new");
    }

    let help_message = "
    new function
    ===============

    ";

    let mut position_params: Vec<FunctionParam> = Vec::new();
    let mut named_params: HashMap<String, FunctionParam> = HashMap::new();

    vm.__read_params(&mut position_params, &mut named_params);

    let mut input_file_type = String::from("csv");
    if named_params.contains_key("type") {
        input_file_type =
            vm.__symbol_table[u64::from_be_bytes(named_params["type"].value) as usize].clone();
    }

    vm.__current_table = DataFrame::default();
}

pub fn prql_export(vm: &mut PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING output");
    }

    let help_message = "
    output function
    ===============

    (path)    - the path of the input file
    
    type      - [ csv | json ]
                  ^^^
    delimiter - \",\" the file delimiter
                  ^
    ";

    let mut position_params: Vec<FunctionParam> = Vec::new();
    let mut named_params: HashMap<String, FunctionParam> = HashMap::new();

    vm.__read_params(&mut position_params, &mut named_params);

    let path = vm.__symbol_table[u64::from_be_bytes(position_params[0].value) as usize].as_str();

    let mut input_file_type = String::from("csv");
    if named_params.contains_key("type") {
        input_file_type =
            vm.__symbol_table[u64::from_be_bytes(named_params["type"].value) as usize].clone();
    }

    let mut delimiter: u8 = ',' as u8;
    if named_params.contains_key("delimiter") {
        delimiter = vm.__symbol_table[u64::from_be_bytes(named_params["delimiter"].value) as usize]
            .as_bytes()[0];
    }

    match input_file_type.as_str() {
        "csv" => {
            CsvWriter::new(File::create(path).unwrap())
                .with_delimiter(delimiter)
                .finish(&mut vm.__current_table)
                .unwrap();
        }
        "json" => {}
        _ => {}
    }
}

pub fn prql_select(vm: &mut PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING select");
    }

    let help_message = "
    select function
    ===============
    ";
}
