package preludiocore

import (
	"bytefeeder"
	"gandalff"
	"os"
	"testing"
	"typesys"
)

func Test_Builtin_New(t *testing.T) {
	var err error
	var source string
	var bytecode []byte
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

	be := new(ByteEater).InitVM()
	bytecode, _, _ = bytefeeder.CompileSource(source)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {

		// check types
		if df.Series("A").Type() != typesys.BoolType {
			t.Error("Expected bool type, got", df.Series("A").Type())
		}
		if df.Series("B").Type() != typesys.StringType {
			t.Error("Expected string type, got", df.Series("B").Type())
		}
		if df.Series("C").Type() != typesys.Int64Type {
			t.Error("Expected int type, got", df.Series("C").Type())
		}
		if df.Series("D").Type() != typesys.Float64Type {
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

	bytecode, _, _ = bytefeeder.CompileSource(source)
	be.InitVM().RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {

		// check types
		if df.Series("A").Type() != typesys.BoolType {
			t.Error("Expected bool type, got", df.Series("A").Type())
		}
		if df.Series("B").Type() != typesys.StringType {
			t.Error("Expected string type, got", df.Series("B").Type())
		}
		if df.Series("C").Type() != typesys.Int64Type {
			t.Error("Expected int type, got", df.Series("C").Type())
		}
		if df.Series("D").Type() != typesys.Float64Type {
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

	bytecode, _, err = bytefeeder.CompileSource(source)
	be.RunBytecode(bytecode)
}

func Test_Builtin_Pipelines1(t *testing.T) {
	var err error
	var source string
	var bytecode []byte
	var df gandalff.DataFrame

	// basic test
	source = `
let clean = (
	readCSV "..\\test_files\\Cars.csv" delimiter: ";" header:true
	strReplace [MPG, Displacement, Horsepower, Acceleration] old:"," new:"."
	asFloat [MPG, Displacement, Horsepower, Acceleration]
	orderBy [-Origin, Cylinders, -MPG]
)

let europe5Cylinders = (
	from clean
	filter Cylinders == 5 and Origin == "Europe"
)	  
`
	be = new(ByteEater).InitVM()
	bytecode, _, _ = bytefeeder.CompileSource(source)
	be.RunBytecode(bytecode)

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
			if df.SeriesAt(0).Type() != typesys.StringType {
				t.Error("Expected string type, got", df.SeriesAt(0).Type())
			}
			if df.SeriesAt(1).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(1).Type())
			}
			if df.SeriesAt(2).Type() != typesys.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(2).Type())
			}
			if df.SeriesAt(3).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(3).Type())
			}
			if df.SeriesAt(4).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(4).Type())
			}
			if df.SeriesAt(5).Type() != typesys.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(5).Type())
			}
			if df.SeriesAt(6).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(6).Type())
			}
			if df.SeriesAt(7).Type() != typesys.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(7).Type())
			}
			if df.SeriesAt(8).Type() != typesys.StringType {
				t.Error("Expected string type, got", df.SeriesAt(8).Type())
			}

			// check names
			if df.SeriesAt(0).Name() != "Car" {
				t.Error("Expected Car, got", df.SeriesAt(0).Name())
			}
			if df.SeriesAt(1).Name() != "MPG" {
				t.Error("Expected MPG, got", df.SeriesAt(1).Name())
			}
			if df.SeriesAt(2).Name() != "Cylinders" {
				t.Error("Expected Cylinders, got", df.SeriesAt(2).Name())
			}
			if df.SeriesAt(3).Name() != "Displacement" {
				t.Error("Expected Displacement, got", df.SeriesAt(3).Name())
			}
			if df.SeriesAt(4).Name() != "Horsepower" {
				t.Error("Expected Horsepower, got", df.SeriesAt(4).Name())
			}
			if df.SeriesAt(5).Name() != "Weight" {
				t.Error("Expected Weight, got", df.SeriesAt(5).Name())
			}
			if df.SeriesAt(6).Name() != "Acceleration" {
				t.Error("Expected Acceleration, got", df.SeriesAt(6).Name())
			}
			if df.SeriesAt(7).Name() != "Model" {
				t.Error("Expected Model, got", df.SeriesAt(7).Name())
			}
			if df.SeriesAt(8).Name() != "Origin" {
				t.Error("Expected Origin, got", df.SeriesAt(8).Name())
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
			if df.SeriesAt(0).Type() != typesys.StringType {
				t.Error("Expected string type, got", df.SeriesAt(0).Type())
			}
			if df.SeriesAt(1).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(1).Type())
			}
			if df.SeriesAt(2).Type() != typesys.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(2).Type())
			}
			if df.SeriesAt(3).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(3).Type())
			}
			if df.SeriesAt(4).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(4).Type())
			}
			if df.SeriesAt(5).Type() != typesys.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(5).Type())
			}
			if df.SeriesAt(6).Type() != typesys.Float64Type {
				t.Error("Expected float type, got", df.SeriesAt(6).Type())
			}
			if df.SeriesAt(7).Type() != typesys.Int64Type {
				t.Error("Expected int type, got", df.SeriesAt(7).Type())
			}
			if df.SeriesAt(8).Type() != typesys.StringType {
				t.Error("Expected string type, got", df.SeriesAt(8).Type())
			}

			// check names
			if df.SeriesAt(0).Name() != "Car" {
				t.Error("Expected Car, got", df.SeriesAt(0).Name())
			}
			if df.SeriesAt(1).Name() != "MPG" {
				t.Error("Expected MPG, got", df.SeriesAt(1).Name())
			}
			if df.SeriesAt(2).Name() != "Cylinders" {
				t.Error("Expected Cylinders, got", df.SeriesAt(2).Name())
			}
			if df.SeriesAt(3).Name() != "Displacement" {
				t.Error("Expected Displacement, got", df.SeriesAt(3).Name())
			}
			if df.SeriesAt(4).Name() != "Horsepower" {
				t.Error("Expected Horsepower, got", df.SeriesAt(4).Name())
			}
			if df.SeriesAt(5).Name() != "Weight" {
				t.Error("Expected Weight, got", df.SeriesAt(5).Name())
			}
			if df.SeriesAt(6).Name() != "Acceleration" {
				t.Error("Expected Acceleration, got", df.SeriesAt(6).Name())
			}
			if df.SeriesAt(7).Name() != "Model" {
				t.Error("Expected Model, got", df.SeriesAt(7).Name())
			}
			if df.SeriesAt(8).Name() != "Origin" {
				t.Error("Expected Origin, got", df.SeriesAt(8).Name())
			}
		}
	} else {
		t.Error("Expected result, got nil")
	}
}

func Test_Builtin_Pipelines2(t *testing.T) {
	var err error
	var source string
	var bytecode []byte

	// basic test
	source = `
let clean = (
	readCSV "..\\test_files\\Cars.csv" delimiter: ";" header:true
	strReplace [MPG, Displacement, Horsepower, Acceleration] old:"," new:"."
	asFloat [MPG, Displacement, Horsepower, Acceleration]
	orderBy [-Origin, Cylinders, -MPG]
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
	writeCSV "..\\test_files\\CarsRes.csv" delimiter: "\t"
)	  
`

	bytecode, _, _ = bytefeeder.CompileSource(source)
	new(ByteEater).InitVM().RunBytecode(bytecode)

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
