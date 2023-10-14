package preludiocore

import (
	"gandalff"
	"os"
	"preludiometa"
	"testing"
)

func init() {
	be = new(ByteEater).InitVM()
}

func Test_Builtin_readCSV(t *testing.T) {
	var err error
	var df gandalff.DataFrame

	// CSV, comma delimiter, no header
	content := `true,hello,.43403,0
false,world,3e-2,4294
true,,0.000000001,-324
false,this is a string,4E4,3245
false,"hello again",0.000000000001,0`

	err = os.WriteFile("csvtest00_read_comma.csv", []byte(content), 0644)
	if err != nil {
		t.Error("Error writing test file", err)
	}
	defer os.Remove("csvtest00_read_comma.csv")

	be.RunSource(`rcsv "csvtest00_read_comma.csv" del: "," head: false`)
	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {
		records := df.Records(false)

		if len(records) != 5 {
			t.Error("Expected 5 records, got", len(records))
		}

		if records[0][0] != "true" {
			t.Error("Expected \"true\", got", records[0][0])
		}
		if records[0][1] != "hello" {
			t.Error("Expected \"hello\", got", records[0][1])
		}
		if records[3][1] != "this is a string" {
			t.Error("Expected \"this is a string\", got", records[3][1])
		}
	} else {
		t.Error("Expected no error, got", err)
	}

	// CSV, semicolon delimiter, no header
	content = `true;hello;.43403;0
false;world;3e-2;4294
true;;0.000000001;-324
false;this is a string;4E4;3245
false;"hello again";0.000000000001;0`

	err = os.WriteFile("csvtest01_read_semicolon.csv", []byte(content), 0644)
	if err != nil {
		t.Error("Error writing test file", err)
	}
	defer os.Remove("csvtest01_read_semicolon.csv")

	be.RunSource(`rcsv "csvtest01_read_semicolon.csv" del: ";" head: false`)
	if be.__currentResult == nil {
		t.Error("Expected result, got nil", be.__output.Log)
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {
		records := df.Records(false)

		if len(records) != 5 {
			t.Error("Expected 5 records, got", len(records))
		}

		if records[0][0] != "true" {
			t.Error("Expected \"true\", got", records[0][0])
		}
		if records[0][1] != "hello" {
			t.Error("Expected \"hello\", got", records[0][1])
		}
		if records[3][1] != "this is a string" {
			t.Error("Expected \"this is a string\", got", records[3][1])
		}
	} else {
		t.Error("Expected no error, got", err)
	}

	// CSV, tab delimiter, header
	content = `A bool	something	a numeric value	an integer value
true	hello	.43403	0
false	world	3e-2	4294
true	0.000000001	-324	-1
false	this is a string	4E4	3245
false	"hello again"	0.000000000001	0`

	err = os.WriteFile("csvtest02_read_tab_header.csv", []byte(content), 0644)
	if err != nil {
		t.Error("Error writing test file", err)
	}
	defer os.Remove("csvtest02_read_tab_header.csv")

	be.RunSource(`rcsv "csvtest02_read_tab_header.csv" del: "\t" head: true`)
	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err != nil {
		records := df.Records(false)

		if len(records) != 4 {
			t.Error("Expected 4 records, got", len(records))
		}

		if records[0][0] != "true" {
			t.Error("Expected \"true\", got", records[0][0])
		}
		if records[0][1] != "hello" {
			t.Error("Expected \"hello\", got", records[0][1])
		}
		if records[3][1] != "this is a string" {
			t.Error("Expected \"this is a string\", got", records[3][1])
		}
	}
}

func Test_Builtin_New(t *testing.T) {
	var err error
	var source string
	var df gandalff.DataFrame

	// basic test
	source = `
	(
		new [
			A = [true, false, true, false, true],
			B = ["hello", "world", "this is a string", "this is another string", "this is a third string"],
			C = [1, 2, 3, 4, 5],
			D = [1.1, 2.2, 3.3, 4.4, 5.5]
		]
	)
	`

	be.RunSource(source)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {

		// check types
		if df.Series("A").Type() != preludiometa.BoolType {
			t.Error("Expected bool type, got", df.Series("A").Type())
		}
		if df.Series("B").Type() != preludiometa.StringType {
			t.Error("Expected string type, got", df.Series("B").Type())
		}
		if df.Series("C").Type() != preludiometa.Int64Type {
			t.Error("Expected int type, got", df.Series("C").Type())
		}
		if df.Series("D").Type() != preludiometa.Float64Type {
			t.Error("Expected float type, got", df.Series("D").Type())
		}

		// check values
		bools := []bool{true, false, true, false, true}
		if !boolSliceEqual(df.Series("A").(gandalff.SeriesBool).Bools(), bools) {
			t.Error("Expected bool values", bools, "got", df.Series("A").(gandalff.SeriesBool).Bools())
		}

		strings := []string{"hello", "world", "this is a string", "this is another string", "this is a third string"}
		if !stringSliceEqual(df.Series("B").(gandalff.SeriesString).Strings(), strings) {
			t.Error("Expected string values", strings, "got", df.Series("B").(gandalff.SeriesString).Strings())
		}

		ints := []int64{1, 2, 3, 4, 5}
		if !int64SliceEqual(df.Series("C").(gandalff.SeriesInt64).Int64s(), ints) {
			t.Error("Expected int values", ints, "got", df.Series("C").(gandalff.SeriesInt64).Int64s())
		}

		floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
		if !float64SliceEqual(df.Series("D").(gandalff.SeriesFloat64).Float64s(), floats) {
			t.Error("Expected float values", floats, "got", df.Series("D").(gandalff.SeriesFloat64).Float64s())
		}
	} else {
		t.Error("Expected no error, got", err)
	}

	// from lists
	source = `
	let listOfBools = [true, false, true, false, true]
	let listOfStrings = ["hello", "world", "this is a string", "this is another string", "this is a third string"]

	(
		new [
			A = listOfBools or false,
			B = listOfStrings + "!",
			C = [1, 2, 3, 4, 5] * 2,
			D = [1.1, 2.2, 3.3, 4.4, 5.5] / 2
		]
	)
	`

	be.RunSource(source)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {

		// check types
		if df.Series("A").Type() != preludiometa.BoolType {
			t.Error("Expected bool type, got", df.Series("A").Type())
		}
		if df.Series("B").Type() != preludiometa.StringType {
			t.Error("Expected string type, got", df.Series("B").Type())
		}
		if df.Series("C").Type() != preludiometa.Int64Type {
			t.Error("Expected int type, got", df.Series("C").Type())
		}
		if df.Series("D").Type() != preludiometa.Float64Type {
			t.Error("Expected float type, got", df.Series("D").Type())
		}

		// check values
		bools := []bool{true, false, true, false, true}
		if !boolSliceEqual(df.Series("A").(gandalff.SeriesBool).Bools(), bools) {
			t.Error("Expected bool values", bools, "got", df.Series("A").(gandalff.SeriesBool).Bools())
		}

		strings := []string{"hello!", "world!", "this is a string!", "this is another string!", "this is a third string!"}
		if !stringSliceEqual(df.Series("B").(gandalff.SeriesString).Strings(), strings) {
			t.Error("Expected string values", strings, "got", df.Series("B").(gandalff.SeriesString).Strings())
		}

		ints := []int64{2, 4, 6, 8, 10}
		if !int64SliceEqual(df.Series("C").(gandalff.SeriesInt64).Int64s(), ints) {
			t.Error("Expected int values", ints, "got", df.Series("C").(gandalff.SeriesInt64).Int64s())
		}

		floats := []float64{0.55, 1.1, 1.65, 2.2, 2.75}
		if !float64SliceEqual(df.Series("D").(gandalff.SeriesFloat64).Float64s(), floats) {
			t.Error("Expected float values", floats, "got", df.Series("D").(gandalff.SeriesFloat64).Float64s())
		}
	} else {
		t.Error("Expected no error, got", err)
	}

	// only one column
	source = `(new [A = [1, 2, 3, 4, 5]])`
	be.RunSource(source)
	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {

		// check types
		if df.Series("A").Type() != preludiometa.Int64Type {
			t.Error("Expected int type, got", df.Series("A").Type())
		}

		// check values
		ints := []int64{1, 2, 3, 4, 5}
		if !int64SliceEqual(df.Series("A").(gandalff.SeriesInt64).Int64s(), ints) {
			t.Error("Expected int values", ints, "got", df.Series("A").(gandalff.SeriesInt64).Int64s())
		}
	} else {
		t.Error("Expected no error, got", err)
	}

	// different lengths
	source = `
	(
		new [
			A = [true, false, true, false],
			B = ["hello", "world", "this is a string"],
			C = [1, 2, 3, 4, 5],
		]
	)
	`

	be.RunSource(source)
	if be.getLastError() != "new: BaseDataFrame.AddSeries: series length (3) does not match dataframe length (4)" {
		t.Error("Expected error, got", be.getLastError())
	}
}

func Test_Builtin_Join(t *testing.T) {
	var err error
	var source string
	var df gandalff.DataFrame

	// basic test
	source = `
	let df1 = (
		new [
			A = [true, false, true, false, true],
			B = ["one", "two", "three", "four", "five"],
			C = [1, 2, 3, 4, 5],
			D = [1.1, 2.2, 3.3, 4.4, 5.5]
		]
	)

	let df2 = (
		new [
			A = [true, true, false],
			B = ["one", "four", "five"],
			C = [1, 1, 1],
		]
	)

	let df3 = (
		new [
			A = [true, false, true, false, true],
			B = ["four", "five", "six", "seven", "eight"],
			C = [1, 1, 2, 6, 7],
			D = [5.5, 4.4, 1.1, 1.1, 1.1]
		]
	)

	let j1 = (from df1 | join left df2 on: [A])
	let j2 = (from df1 | join right df2 on: [A])
	let j3 = (from df1 | join inner df2 on: [A])
	let j4 = (from df1 | join outer df2 on: [A])
	`

	be.RunSource(source)

	if p, ok := be.__globalNamespace["j1"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B_x").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_x").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("B_y").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_y").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 8 {
				t.Error("Expected 8 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}

	if p, ok := be.__globalNamespace["j2"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B_x").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_x").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("B_y").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_y").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 8 {
				t.Error("Expected 8 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}

	if p, ok := be.__globalNamespace["j3"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B_x").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_x").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("B_y").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_y").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 8 {
				t.Error("Expected 8 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}

	if p, ok := be.__globalNamespace["j4"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B_x").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_x").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("B_y").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B_y").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 8 {
				t.Error("Expected 8 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}

	source = `
	j1 = (from df1 | join left df2 on: [A, B])
	j2 = (from df1 | join right df2 on: [A, B])
	j3 = (from df1 | join inner df2 on: [A, B])
	j4 = (from df1 | join outer df2 on: [A, B])
	`

	be.RunSource(source)

	if p, ok := be.__globalNamespace["j1"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 5 {
				t.Error("Expected 5 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}

	if p, ok := be.__globalNamespace["j2"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 3 {
				t.Error("Expected 3 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}

	if p, ok := be.__globalNamespace["j3"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 1 {
				t.Error("Expected 1 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}

	if p, ok := be.__globalNamespace["j4"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err == nil {

			// check types
			if df.Series("A").Type() != preludiometa.BoolType {
				t.Error("Expected bool type, got", df.Series("A").Type().ToString())
			}
			if df.Series("B").Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.Series("B").Type().ToString())
			}
			if df.Series("C_x").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_x").Type().ToString())
			}
			if df.Series("D").Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.Series("D").Type().ToString())
			}
			if df.Series("C_y").Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.Series("C_y").Type().ToString())
			}

			// check number of rows
			if df.NRows() != 7 {
				t.Error("Expected 7 rows, got", df.NRows())
			}

		} else {
			t.Error("Expected no error, got", err)
		}
	} else {
		t.Error("Expected result, got nil")
	}
}

func Test_Builtin_Pipelines1(t *testing.T) {
	var err error
	var source string
	var df gandalff.DataFrame

	// basic test
	source = `
	let clean = (
		rcsv "..\\test_files\\Cars.csv" del:";" head:true
		strReplace [MPG, Displacement, Horsepower, Acceleration] old:"," new:"."
		asFlt [MPG, Displacement, Horsepower, Acceleration]
		sort [-Origin, Cylinders, -MPG]
	)

	let europe5Cylinders = (
		from clean
		filter Cylinders == 5 and Origin == "Europe"
	)
	`

	be.RunSource(source)

	// check clean
	if p, ok := be.__globalNamespace["clean"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err != nil {
			t.Error("Expected no error, got", err)
		} else {
			if df.NCols() != 9 {
				t.Error("Expected 9 columns, got", df.NCols())
			}

			// check types
			if df.SeriesAt(0).Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.SeriesAt(0).Type())
			}
			if df.SeriesAt(1).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(1).Type())
			}
			if df.SeriesAt(2).Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(2).Type())
			}
			if df.SeriesAt(3).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(3).Type())
			}
			if df.SeriesAt(4).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(4).Type())
			}
			if df.SeriesAt(5).Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(5).Type())
			}
			if df.SeriesAt(6).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(6).Type())
			}
			if df.SeriesAt(7).Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(7).Type())
			}
			if df.SeriesAt(8).Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.SeriesAt(8).Type())
			}

			// check names
			if df.NameAt(0) != "Car" {
				t.Error("Expected Car, got", df.NameAt(0))
			}
			if df.NameAt(1) != "MPG" {
				t.Error("Expected MPG, got", df.NameAt(1))
			}
			if df.NameAt(2) != "Cylinders" {
				t.Error("Expected Cylinders, got", df.NameAt(2))
			}
			if df.NameAt(3) != "Displacement" {
				t.Error("Expected Displacement, got", df.NameAt(3))
			}
			if df.NameAt(4) != "Horsepower" {
				t.Error("Expected Horsepower, got", df.NameAt(4))
			}
			if df.NameAt(5) != "Weight" {
				t.Error("Expected Weight, got", df.NameAt(5))
			}
			if df.NameAt(6) != "Acceleration" {
				t.Error("Expected Acceleration, got", df.NameAt(6))
			}
			if df.NameAt(7) != "Model" {
				t.Error("Expected Model, got", df.NameAt(7))
			}
			if df.NameAt(8) != "Origin" {
				t.Error("Expected Origin, got", df.NameAt(8))
			}
		}
	} else {
		t.Error("Expected result, got nil")
	}

	// check europe5Cylinders
	if p, ok := be.__globalNamespace["europe5Cylinders"]; ok {
		if !p.isDataframe() {
			t.Error("Expected dataframe, got", p)
		} else if df, err = p.getDataframe(); err != nil {
			t.Error("Expected no error, got", err)
		} else {
			if df.NCols() != 9 {
				t.Error("Expected 9 columns, got", df.NCols())
			}
			if df.NRows() != 3 {
				t.Error("Expected 3 rows, got", df.NRows())
			}

			// check types
			if df.SeriesAt(0).Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.SeriesAt(0).Type())
			}
			if df.SeriesAt(1).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(1).Type())
			}
			if df.SeriesAt(2).Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(2).Type())
			}
			if df.SeriesAt(3).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(3).Type())
			}
			if df.SeriesAt(4).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(4).Type())
			}
			if df.SeriesAt(5).Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(5).Type())
			}
			if df.SeriesAt(6).Type() != preludiometa.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(6).Type())
			}
			if df.SeriesAt(7).Type() != preludiometa.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(7).Type())
			}
			if df.SeriesAt(8).Type() != preludiometa.StringType {
				t.Error("Expected string type, got", df.SeriesAt(8).Type())
			}

			// check names
			if df.NameAt(0) != "Car" {
				t.Error("Expected Car, got", df.NameAt(0))
			}
			if df.NameAt(1) != "MPG" {
				t.Error("Expected MPG, got", df.NameAt(1))
			}
			if df.NameAt(2) != "Cylinders" {
				t.Error("Expected Cylinders, got", df.NameAt(2))
			}
			if df.NameAt(3) != "Displacement" {
				t.Error("Expected Displacement, got", df.NameAt(3))
			}
			if df.NameAt(4) != "Horsepower" {
				t.Error("Expected Horsepower, got", df.NameAt(4))
			}
			if df.NameAt(5) != "Weight" {
				t.Error("Expected Weight, got", df.NameAt(5))
			}
			if df.NameAt(6) != "Acceleration" {
				t.Error("Expected Acceleration, got", df.NameAt(6))
			}
			if df.NameAt(7) != "Model" {
				t.Error("Expected Model, got", df.NameAt(7))
			}
			if df.NameAt(8) != "Origin" {
				t.Error("Expected Origin, got", df.NameAt(8))
			}
		}
	} else {
		t.Error("Expected result, got nil")
	}
}

func Test_Builtin_Pipelines2(t *testing.T) {
	var err error

	// basic test
	source := `
	let clean = (
		rcsv "..\\test_files\\Cars.csv" del:";" head:true
		strReplace [MPG, Displacement, Horsepower, Acceleration] old:"," new:"."
		asFlt [MPG, Displacement, Horsepower, Acceleration]
		sort [-Origin, Cylinders, -MPG]
	)

	(
		from clean
		derive [
			Stat = ((MPG * Cylinders * Displacement) / Horsepower * Acceleration) / Weight,
			CarOrigin = Car + " - " + Origin
		]
		filter Stat > 1.3
		select [Car, Origin, Stat]
		take 10
		wcsv "..\\test_files\\CarsRes.csv" del:"\t"
	)
	`

	new(ByteEater).InitVM().RunSource(source)

	b, err := os.ReadFile("..\\test_files\\CarsRes.csv")
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	expected := `Car	Origin	Stat
Plymouth Champ	US	1.83352
Plymouth Horizon Miser	US	1.7524705882352942
Ford Fiesta	US	1.71529696969697
Mercury Lynx l	US	1.6412611764705882
Dodge Colt Hatchback Custom	US	1.3154005221932117
Plymouth Horizon 4	US	1.5561474793077505
Plymouth Horizon TC3	US	1.4345581395348839
Ford Escort 4W	US	1.6434362234342674
Chevrolet Cavalier 2-door	US	1.3008920098690453
Chevrolet Chevette	US	1.3142830188679246
`

	if string(b) != expected {
		t.Error("Expected", expected, "got", string(b))
	}

	err = os.Remove("..\\test_files\\CarsRes.csv")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
}
