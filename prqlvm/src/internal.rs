use polars::prelude::*;
use std::fmt::Display;
use std::fmt::Formatter;

use crate::type_system::PrqlBaseDataType;

#[derive(Clone, Copy, Debug)]
pub enum PrqlInternalDim {
    Scalar,
    Series,
    Table,
}

impl Display for PrqlInternalDim {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        let s = match self {
            PrqlInternalDim::Scalar => String::from("Scalar"),
            PrqlInternalDim::Series => String::from("Series"),
            PrqlInternalDim::Table => String::from("Table"),
        };
        f.write_str(s.as_str())
    }
}

#[derive(Clone, Copy, Debug)]
pub enum PrqlInternalTag {
    ExprTerm,
    ParamName,
    AssignIdent,
    Error,
}

impl Display for PrqlInternalTag {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        let s = match self {
            PrqlInternalTag::ExprTerm => String::from("Term"),
            PrqlInternalTag::ParamName => String::from("Param Name"),
            PrqlInternalTag::AssignIdent => String::from("Assign Ident"),
            PrqlInternalTag::Error => String::from("Error"),
        };
        f.write_str(s.as_str())
    }
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

    pub fn copy() {}

    pub fn get_dim(self) -> PrqlInternalDim {
        return self.dim;
    }

    pub fn get_tag(self) -> PrqlInternalTag {
        return self.tag;
    }

    pub fn get_data_frame(self) -> Option<DataFrame> {
        return self.data_frame;
    }

    pub fn get_name(self) -> Option<String> {
        return self.name;
    }

    pub fn get_scalar_bool(self) -> Option<bool> {
        return self.scalar_bool;
    }

    pub fn get_scalar_num(self) -> Option<f64> {
        return self.scalar_num;
    }

    pub fn get_scalar_string(self) -> Option<String> {
        return self.scalar_string;
    }

    pub fn arith_binary_mul(mut self, term: PrqlInternal) {
        match self.dim {
            PrqlInternalDim::Scalar => match term.dim {
                PrqlInternalDim::Scalar => match self.scalar_type.unwrap() {
                    PrqlBaseDataType::Null => match term.scalar_type.unwrap() {
                        PrqlBaseDataType::Null => {}
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = Some(PrqlBaseDataType::Bool);
                            self.scalar_bool = Some(false);
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = Some(PrqlBaseDataType::Numeric);
                            self.scalar_num = Some(0.0);
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = Some(PrqlBaseDataType::String);
                            self.scalar_string = Some(String::from(""));
                        }
                    },

                    PrqlBaseDataType::Bool => match term.scalar_type.unwrap() {
                        PrqlBaseDataType::Null => {
                            self.scalar_type = Some(PrqlBaseDataType::Bool);
                            self.scalar_bool = Some(false);
                        }
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = Some(PrqlBaseDataType::Bool);
                            self.scalar_bool =
                                Some(self.scalar_bool.unwrap() & term.scalar_bool.unwrap());
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = Some(PrqlBaseDataType::Numeric);
                            if self.scalar_bool.unwrap() {
                                self.scalar_num = term.scalar_num;
                            } else {
                                self.scalar_num = Some(0.0);
                            }
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = Some(PrqlBaseDataType::String);
                            if self.scalar_bool.unwrap() {
                                self.scalar_string = term.scalar_string;
                            } else {
                                self.scalar_string = Some(String::from(""));
                            }
                        }
                    },

                    PrqlBaseDataType::Numeric => match term.scalar_type.unwrap() {
                        PrqlBaseDataType::Null => {
                            self.scalar_type = Some(PrqlBaseDataType::Numeric);
                            self.scalar_num = Some(0.0);
                        }
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = Some(PrqlBaseDataType::Numeric);
                            if !term.scalar_bool.unwrap() {
                                self.scalar_num = Some(0.0);
                            }
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = Some(PrqlBaseDataType::Numeric);
                            self.scalar_num =
                                Some(self.scalar_num.unwrap() * term.scalar_num.unwrap());
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = Some(PrqlBaseDataType::String);
                            self.scalar_string = Some(
                                term.scalar_string
                                    .unwrap()
                                    .repeat(self.scalar_num.unwrap() as usize),
                            );
                        }
                    },

                    PrqlBaseDataType::String => match term.scalar_type.unwrap() {
                        PrqlBaseDataType::Null => {
                            self.scalar_type = Some(PrqlBaseDataType::String);
                            self.scalar_string = Some(String::from(""));
                        }
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = Some(PrqlBaseDataType::String);
                            if !term.scalar_bool.unwrap() {
                                self.scalar_string = Some(String::from(""));
                            }
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = Some(PrqlBaseDataType::String);
                            self.scalar_num =
                                Some(self.scalar_num.unwrap() * term.scalar_num.unwrap());
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = Some(PrqlBaseDataType::String);
                            self.scalar_string = Some(
                                term.scalar_string
                                    .unwrap()
                                    .repeat(self.scalar_num.unwrap() as usize),
                            );
                        }
                    },
                },
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

    pub fn arith_binary_div(mut self, term2: PrqlInternal) {}

    pub fn arith_binary_mod(mut self, term2: PrqlInternal) {}

    pub fn arith_binary_add(mut self, term2: PrqlInternal) {}

    pub fn arith_binary_sub(mut self, term2: PrqlInternal) {}
}

impl Copy for PrqlInternal {}

impl Clone for PrqlInternal {
    fn clone(&self) -> PrqlInternal {
        return PrqlInternal {
            dim: self.dim,
            tag: self.tag,
            scalar_type: Some(self.scalar_type.unwrap().clone()),
            scalar_bool: Some(self.scalar_bool.unwrap().clone()),
            scalar_num: Some(self.scalar_num.unwrap()),
            scalar_string: Some(self.scalar_string.unwrap().clone()),
            name: Some(self.name.unwrap().clone()),
            data_frame: Some(self.data_frame.unwrap().clone()),
            error_message: Some(self.error_message.unwrap().clone()),
        };
    }
}

impl Display for PrqlInternal {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        let s = format!("Internal Dim: {}, Tag: {}", self.dim, self.tag);
        f.write_str(s.as_str())
    }
}
