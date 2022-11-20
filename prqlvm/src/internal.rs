use polars::export::num::ToPrimitive;
use polars::prelude::*;

use crate::type_system::PrqlBaseDataType;

pub enum PrqlInternalDim {
    Scalar,
    Series,
    Table,
}

pub enum PrqlInternalTag {
    ExprTerm,
    ParamName,
    AssignIdent,
    Error,
}

pub struct PrqlInternal {
    dim: PrqlInternalDim,
    tag: PrqlInternalTag,
    scalar_type: Option<super::type_system::PrqlBaseDataType>,
    scalar_bool: Option<bool>,
    scalar_num: Option<f64>,
    scalar_string: Option<String>,
    name: Option<String>,
    data_frame: Option<DataFrame>,

    error_message: Option<String>,
}

impl PrqlInternal {
    pub fn new_scalar(
        scalar_type: PrqlBaseDataType,
        scalar_bool: Option<bool>,
        scal_num: Option<f64>,
        scal_string: Option<String>,
    ) -> PrqlInternal {
        return PrqlInternal {
            dim: PrqlInternalDim::Scalar,
            tag: PrqlInternalTag::ExprTerm,
            scalar_type: Some(scalar_type),
            scalar_bool: scalar_bool,
            scalar_num: scal_num,
            scalar_string: scal_string,
            name: None,
            data_frame: None,
            error_message: None,
        };
    }

    pub fn new_error(msg: Option<String>) -> PrqlInternal {
        return PrqlInternal {
            dim: PrqlInternalDim::Scalar,
            tag: PrqlInternalTag::Error,
            scalar_type: None,
            scalar_bool: None,
            scalar_num: None,
            scalar_string: None,
            name: None,
            data_frame: None,
            error_message: msg,
        };
    }

    pub fn get_dim(self) -> PrqlInternalDim {
        return self.dim;
    }
    
    pub fn get_data_frame(self) -> Option<DataFrame> {
        return self.data_frame;
    }

    pub fn arith_binary_mul(mut self, term: PrqlInternal) {
        match self.dim {
            PrqlInternalDim::Scalar => match term.dim {
                PrqlInternalDim::Scalar => {}
                PrqlInternalDim::Series => {}
                PrqlInternalDim::Table => {
                    self.tag = PrqlInternalTag::Error;
                    self.error_message = std::option::Option::Some(String::from(
                        "Error: binary mult not implemented for tables.",
                    ));
                }
            },
            PrqlInternalDim::Series => match term.dim {
                PrqlInternalDim::Scalar => {}
                PrqlInternalDim::Series => {}
                PrqlInternalDim::Table => {
                    self.tag = PrqlInternalTag::Error;
                    self.error_message = std::option::Option::Some(String::from(
                        "Error: binary mult not implemented for tables.",
                    ));
                }
            },
            PrqlInternalDim::Table => {
                self.tag = PrqlInternalTag::Error;
                self.error_message = std::option::Option::Some(String::from(
                    "Error: binary mult not implemented for tables.",
                ));
            }
        }

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
                    let col_name =
                        self.__symbol_table[(u64::from_be_bytes(term2.param2) as usize)].clone();

                    let tmp_col_name = "tmp";
                    result.param1 = TERM_IDENT;
                    result.param2 = self.__insert_symbol(String::from(tmp_col_name));

                    // match self
                    //     .__current_result
                    //     .schema()
                    //     .unwrap()
                    //     .get(&col_name)
                    //     .unwrap()
                    // {
                    //     DataType::Null => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         |s| Ok(s),
                    //                         GetOutput::from_type(DataType::Null),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Boolean => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         |s| Ok(s * 0),
                    //                         GetOutput::from_type(DataType::Boolean),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Float64 => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         |s| Ok(s * 0),
                    //                         GetOutput::from_type(DataType::Float64),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Utf8 => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
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
                    //     .__current_result
                    //     .schema()
                    //     .unwrap()
                    //     .get(&col_name)
                    //     .unwrap()
                    // {
                    //     DataType::Null => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         |s| Ok(s),
                    //                         GetOutput::from_type(DataType::Boolean),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Boolean => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         |s| Ok(s * value),
                    //                         GetOutput::from_type(DataType::Boolean),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Float64 => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         |s| Ok(s * value),
                    //                         GetOutput::from_type(DataType::Float64),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Utf8 => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
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
                    //     .__current_result
                    //     .schema()
                    //     .unwrap()
                    //     .get(&col_name)
                    //     .unwrap()
                    // {
                    //     DataType::Null => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         |s| Ok(s),
                    //                         GetOutput::from_type(DataType::Boolean),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Boolean => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         move |s| Ok(s * value),
                    //                         GetOutput::from_type(DataType::Boolean),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Float64 => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
                    //                 col(&col_name)
                    //                     .map(
                    //                         move |s| Ok(s * value),
                    //                         GetOutput::from_type(DataType::Float64),
                    //                     )
                    //                     .alias(tmp_col_name),
                    //             );
                    //     }
                    //     DataType::Utf8 => {
                    //         self.__current_result =
                    //             self.__current_result.clone().with_column(
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
    }

    pub fn arith_binary_div(mut self, term2: PrqlInternal) {}

    pub fn arith_binary_mod(mut self, term2: PrqlInternal) {}

    pub fn arith_binary_add(mut self, term2: PrqlInternal) {
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
                        self.__symbol_table[(u64::from_be_bytes(term1.param2) as usize)].clone()
                            + &self.__symbol_table[(u64::from_be_bytes(term2.param2) as usize)],
                    );
                }
                TERM_IDENT => {}
                _ => {}
            },
            TERM_IDENT => {}
            _ => {}
        }
    }

    pub fn arith_binary_sub(mut self, term2: PrqlInternal) {}
}
