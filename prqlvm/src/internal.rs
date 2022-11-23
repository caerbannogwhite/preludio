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
    scalar_type: super::type_system::PrqlBaseDataType,
    scalar_bool: bool,
    scalar_num: f64,
    scalar_string: String,
    name: String,
    data_frame: DataFrame,
    error_message: String,
}

impl PrqlInternal {
    pub fn new() -> Self {
        Self {
            dim: PrqlInternalDim::Scalar,
            tag: PrqlInternalTag::ExprTerm,
            scalar_type: PrqlBaseDataType::Null,
            scalar_bool: false,
            scalar_num: 0.0,
            scalar_string: String::from(""),
            name: String::from(""),
            data_frame: DataFrame::empty(),
            error_message: String::from(""),
        }
    }

    pub fn new_assign_ident(ident: String) -> Self {
        Self {
            dim: PrqlInternalDim::Scalar,
            tag: PrqlInternalTag::AssignIdent,
            scalar_type: PrqlBaseDataType::Null,
            scalar_bool: false,
            scalar_num: 0.0,
            scalar_string: String::from(""),
            name: ident,
            data_frame: DataFrame::empty(),
            error_message: String::from(""),
        }
    }

    pub fn new_error(msg: String) -> Self {
        Self {
            dim: PrqlInternalDim::Scalar,
            tag: PrqlInternalTag::Error,
            scalar_type: PrqlBaseDataType::Null,
            scalar_bool: false,
            scalar_num: 0.0,
            scalar_string: String::from(""),
            name: String::from(""),
            data_frame: DataFrame::empty(),
            error_message: msg,
        }
    }

    pub fn new_param_name(param_name: String) -> Self {
        Self {
            dim: PrqlInternalDim::Scalar,
            tag: PrqlInternalTag::ParamName,
            scalar_type: PrqlBaseDataType::Null,
            scalar_bool: false,
            scalar_num: 0.0,
            scalar_string: String::from(""),
            name: param_name,
            data_frame: DataFrame::empty(),
            error_message: String::from(""),
        }
    }

    pub fn with_scalar_null(mut self) -> Self {
        self.scalar_type = PrqlBaseDataType::Null;
        return self;
    }

    pub fn with_scalar_bool(mut self, val: bool) -> Self {
        self.scalar_type = PrqlBaseDataType::Bool;
        self.scalar_bool = val;
        return self;
    }

    pub fn with_scalar_numeric(mut self, val: f64) -> Self {
        self.scalar_type = PrqlBaseDataType::Numeric;
        self.scalar_num = val;
        return self;
    }

    pub fn with_scalar_string(mut self, val: String) -> Self {
        self.scalar_type = PrqlBaseDataType::String;
        self.scalar_string = val;
        return self;
    }

    pub fn with_series(mut self, data_frame: DataFrame) -> Self {
        self.dim = PrqlInternalDim::Series;
        self.tag = PrqlInternalTag::ExprTerm;
        self.data_frame = data_frame;
        return self;
    }

    pub fn with_table(mut self, data_frame: DataFrame) -> Self {
        self.dim = PrqlInternalDim::Table;
        self.tag = PrqlInternalTag::ExprTerm;
        self.data_frame = data_frame;
        return self;
    }

    pub fn get_dim(self) -> PrqlInternalDim {
        return self.dim;
    }

    pub fn get_tag(self) -> PrqlInternalTag {
        return self.tag;
    }

    pub fn get_data_frame(self) -> DataFrame {
        return self.data_frame;
    }

    pub fn get_name(self) -> String {
        return self.name;
    }

    pub fn get_scalar_bool(self) -> bool {
        return self.scalar_bool;
    }

    pub fn get_scalar_num(self) -> f64 {
        return self.scalar_num;
    }

    pub fn get_scalar_string(self) -> String {
        return self.scalar_string;
    }

    pub fn arith_binary_mul(mut self, term: PrqlInternal) {
        match self.dim {
            PrqlInternalDim::Scalar => match term.dim {
                PrqlInternalDim::Scalar => match self.scalar_type {
                    PrqlBaseDataType::Null => match term.scalar_type {
                        PrqlBaseDataType::Null => {}
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = PrqlBaseDataType::Bool;
                            self.scalar_bool = false;
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = PrqlBaseDataType::Numeric;
                            self.scalar_num = 0.0;
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = PrqlBaseDataType::String;
                            self.scalar_string = String::from("");
                        }
                    },

                    PrqlBaseDataType::Bool => match term.scalar_type {
                        PrqlBaseDataType::Null => {
                            self.scalar_type = PrqlBaseDataType::Bool;
                            self.scalar_bool = false;
                        }
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = PrqlBaseDataType::Bool;
                            self.scalar_bool = self.scalar_bool & term.scalar_bool;
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = PrqlBaseDataType::Numeric;
                            if self.scalar_bool {
                                self.scalar_num = term.scalar_num;
                            } else {
                                self.scalar_num = 0.0;
                            }
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = PrqlBaseDataType::String;
                            if self.scalar_bool {
                                self.scalar_string = term.scalar_string;
                            } else {
                                self.scalar_string = String::from("");
                            }
                        }
                    },

                    PrqlBaseDataType::Numeric => match term.scalar_type {
                        PrqlBaseDataType::Null => {
                            self.scalar_type = PrqlBaseDataType::Numeric;
                            self.scalar_num = 0.0;
                        }
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = PrqlBaseDataType::Numeric;
                            if !term.scalar_bool {
                                self.scalar_num = 0.0;
                            }
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = PrqlBaseDataType::Numeric;
                            self.scalar_num = self.scalar_num * term.scalar_num;
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = PrqlBaseDataType::String;
                            self.scalar_string =
                                term.scalar_string.repeat(self.scalar_num as usize);
                        }
                    },

                    PrqlBaseDataType::String => match term.scalar_type {
                        PrqlBaseDataType::Null => {
                            self.scalar_type = PrqlBaseDataType::String;
                            self.scalar_string = String::from("");
                        }
                        PrqlBaseDataType::Bool => {
                            self.scalar_type = PrqlBaseDataType::String;
                            if !term.scalar_bool {
                                self.scalar_string = String::from("");
                            }
                        }
                        PrqlBaseDataType::Numeric => {
                            self.scalar_type = PrqlBaseDataType::String;
                            self.scalar_num = self.scalar_num * term.scalar_num;
                        }
                        PrqlBaseDataType::String => {
                            self.scalar_type = PrqlBaseDataType::String;
                            self.scalar_string =
                                term.scalar_string.repeat(self.scalar_num as usize);
                        }
                    },
                },
                PrqlInternalDim::Series => {}
                PrqlInternalDim::Table => {
                    self.tag = PrqlInternalTag::Error;
                    self.error_message =
                        String::from("Error: binary mult not implemented for tables.");
                }
            },
            PrqlInternalDim::Series => match term.dim {
                PrqlInternalDim::Scalar => {}
                PrqlInternalDim::Series => {}
                PrqlInternalDim::Table => {
                    self.tag = PrqlInternalTag::Error;
                    self.error_message =
                        String::from("Error: binary mult not implemented for tables.");
                }
            },
            PrqlInternalDim::Table => {
                self.tag = PrqlInternalTag::Error;
                self.error_message = String::from("Error: binary mult not implemented for tables.");
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

impl Clone for PrqlInternal {
    fn clone(&self) -> PrqlInternal {
        return PrqlInternal {
            dim: self.dim,
            tag: self.tag,
            scalar_type: self.scalar_type.clone(),
            scalar_bool: self.scalar_bool.clone(),
            scalar_num: self.scalar_num,
            scalar_string: self.scalar_string.clone(),
            name: self.name.clone(),
            data_frame: self.data_frame.clone(),
            error_message: self.error_message.clone(),
        };
    }
}

impl Display for PrqlInternal {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        let s = format!("Internal Dim: {}, Tag: {}", self.dim, self.tag);
        f.write_str(s.as_str())
    }
}
