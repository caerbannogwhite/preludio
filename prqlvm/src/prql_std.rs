use polars::prelude::*;
use std::collections::HashMap;
use std::fs::File;

pub fn prql_derive(vm: &mut super::PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING derive");
    }
}

pub fn prql_from(vm: &mut super::PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING from");
    }
}

pub fn prql_import(vm: &mut super::PRQLVirtualMachine) {
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

    let mut position_params: Vec<super::FunctionParam> = Vec::new();
    let mut named_params: HashMap<String, super::FunctionParam> = HashMap::new();

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
                    super::type_system::PrqlDataType::polars_to_prql(col.1, 0).to_physical(),
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
        }
        "json" => {}
        _ => {}
    }
}

pub fn prql_new(vm: &mut super::PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING new");
    }

    let help_message = "
    new function
    ===============

    ";

    let mut position_params: Vec<super::FunctionParam> = Vec::new();
    let mut named_params: HashMap<String, super::FunctionParam> = HashMap::new();

    vm.__read_params(&mut position_params, &mut named_params);

    let mut input_file_type = String::from("csv");
    if named_params.contains_key("type") {
        input_file_type =
            vm.__symbol_table[u64::from_be_bytes(named_params["type"].value) as usize].clone();
    }

    vm.__current_table = DataFrame::default();
}

pub fn prql_export(vm: &mut super::PRQLVirtualMachine) {
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

    let mut position_params: Vec<super::FunctionParam> = Vec::new();
    let mut named_params: HashMap<String, super::FunctionParam> = HashMap::new();

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

pub fn prql_select(vm: &mut super::PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING select");
    }

    let help_message = "
    select function
    ===============
    ";
}
