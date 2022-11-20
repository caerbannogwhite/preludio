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

pub fn prql_import_csv(vm: &mut super::PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING import_csv");
    }

    let help_message = "
    import_csv function
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

    let mut position_params: Vec<super::internal::PrqlInternal> = Vec::new();
    let mut named_params: HashMap<String, super::internal::PrqlInternal> = HashMap::new();

    vm.__read_params(&mut position_params, &mut named_params);

    let path = position_params[0].get_scalar_string().unwrap();

    let mut delimiter: u8 = ',' as u8;
    if named_params.contains_key("delimiter") {
        delimiter = named_params["delimiter"]
            .get_scalar_string()
            .unwrap()
            .as_bytes()[0];
    }

    let mut enc_str = String::from("utf8");
    if named_params.contains_key("enc") {
        enc_str = named_params["enc"]
            .get_scalar_string()
            .unwrap()
            .to_lowercase();
    }

    let mut skip_rows = 0;
    if named_params.contains_key("skip") {
        skip_rows = named_params["enc"].get_scalar_num().unwrap() as usize;
    }

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

    let df = CsvReader::from_path(path)
        .unwrap()
        .with_delimiter(delimiter)
        .with_encoding(enc)
        .with_skip_rows(skip_rows)
        // .with_n_rows(10)
        .with_dtypes(Some(&schema))
        // .with_quote_char(quote)
        .finish()
        .unwrap();

    vm.__current_result = Some(super::internal::PrqlInternal {
        dim: super::internal::PrqlInternalDim::Table,
        tag: super::internal::PrqlInternalTag::ExprTerm,
        scalar_type: None,
        scalar_bool: None,
        scalar_num: None,
        scalar_string: None,
        name: None,
        data_frame: Some(df),
        error_message: None,
    });

    vm.__stack.push(vm.__current_result.unwrap());
}

pub fn prql_new(vm: &mut super::PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING new");
    }

    let help_message = "
    new function
    ===============

    ";

    let mut position_params: Vec<super::internal::PrqlInternal> = Vec::new();
    let mut named_params: HashMap<String, super::internal::PrqlInternal> = HashMap::new();

    vm.__read_params(&mut position_params, &mut named_params);

    let df = DataFrame::default();

    vm.__current_result = Some(super::internal::PrqlInternal {
        dim: super::internal::PrqlInternalDim::Table,
        tag: super::internal::PrqlInternalTag::ExprTerm,
        scalar_type: None,
        scalar_bool: None,
        scalar_num: None,
        scalar_string: None,
        name: None,
        data_frame: Some(df),
        error_message: None,
    });

    vm.__stack.push(vm.__current_result.unwrap());
}

pub fn prql_export_csv(vm: &mut super::PRQLVirtualMachine) {
    if vm.__debug_level > 5 {
        println!("CALLING export_csv");
    }

    let help_message = "
    export_csv function
    ===============

    (path)    - the path of the input file
    
    type      - [ csv | json ]
                  ^^^
    delimiter - \",\" the file delimiter
                  ^
    ";

    let mut position_params: Vec<super::internal::PrqlInternal> = Vec::new();
    let mut named_params: HashMap<String, super::internal::PrqlInternal> = HashMap::new();

    vm.__read_params(&mut position_params, &mut named_params);

    let path = position_params[0].get_scalar_string().unwrap();

    let mut delimiter: u8 = ',' as u8;
    if named_params.contains_key("delimiter") {
        delimiter = named_params["delimiter"]
            .get_scalar_string()
            .unwrap()
            .as_bytes()[0];
    }

    CsvWriter::new(File::create(path).unwrap())
        .with_delimiter(delimiter)
        .finish(&mut vm.__stack.pop().unwrap().get_data_frame().unwrap())
        .unwrap();
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
