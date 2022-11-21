use polars::prelude::*;
use std::fmt::Display;
use std::fmt::Formatter;

#[derive(Clone, Debug)]
pub enum PrqlBaseDataType {
    Null,
    Bool,
    Numeric,
    String,
}

// impl PrqlBaseDataType {
//     /// Convert to the physical data type
//     #[must_use]
//     pub fn to_rust(&self) -> DataType {
//         match self {
//             PrqlBaseDataType::Null => DataType::Null,
//             PrqlBaseDataType::Bool => DataType::Boolean,
//             PrqlBaseDataType::Numeric => DataType::Float64,
//             PrqlBaseDataType::String => DataType::Utf8,
//         }
//     }
// }

// #[derive(Clone, Debug)]
pub struct PrqlDataType {
    base: PrqlBaseDataType,
    size: usize,
}

impl Default for PrqlDataType {
    fn default() -> Self {
        PrqlDataType {
            base: PrqlBaseDataType::Null,
            size: 0,
        }
    }
}

impl PrqlDataType {
    /// Convert to the physical data type
    #[must_use]
    pub fn to_physical(&self) -> DataType {
        match self.base {
            PrqlBaseDataType::Null => DataType::Null,
            PrqlBaseDataType::Bool => DataType::Boolean,
            PrqlBaseDataType::Numeric => DataType::Float64,
            PrqlBaseDataType::String => DataType::Utf8,
        }
    }

    pub fn is_scalar(&self) -> bool {
        self.size != 1
    }

    pub fn polars_to_prql(dtype: &DataType, size: usize) -> PrqlDataType {
        match dtype {
            DataType::Null => PrqlDataType {
                base: PrqlBaseDataType::Null,
                size: size,
            },
            DataType::Unknown => PrqlDataType {
                base: PrqlBaseDataType::Null,
                size: size,
            },
            DataType::Boolean => PrqlDataType {
                base: PrqlBaseDataType::Bool,
                size: size,
            },
            DataType::UInt8 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::UInt16 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::UInt32 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::UInt64 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::Int8 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::Int16 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::Int32 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::Int64 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::Float32 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::Float64 => PrqlDataType {
                base: PrqlBaseDataType::Numeric,
                size: size,
            },
            DataType::Utf8 => PrqlDataType {
                base: PrqlBaseDataType::String,
                size: size,
            },
            _ => PrqlDataType {
                base: PrqlBaseDataType::Null,
                size: size,
            },
        }
    }

    pub fn prql_to_polars(self) -> DataType {
        match self.base {
            PrqlBaseDataType::Null => DataType::Null,
            PrqlBaseDataType::Bool => DataType::Boolean,
            PrqlBaseDataType::Numeric => DataType::Float64,
            PrqlBaseDataType::String => DataType::Utf8,
        }
    }
}

impl Display for PrqlDataType {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        let s = match self.base {
            PrqlBaseDataType::Null => format!("Null[{}]", self.size),
            PrqlBaseDataType::Bool => format!("Bool[{}]", self.size),
            PrqlBaseDataType::Numeric => format!("Numeric[{}]", self.size),
            PrqlBaseDataType::String => format!("String[{}]", self.size),
        };
        f.write_str(s.as_str())
    }
}
