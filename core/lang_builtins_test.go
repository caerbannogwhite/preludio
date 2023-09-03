package preludiocore

import (
	"bytefeeder"
	"fmt"
	"gandalff"
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

	bytecode, _, _ = bytefeeder.CompileSource(source)
	new(ByteEater).InitVM().RunBytecode(bytecode)

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
	new(ByteEater).InitVM().RunBytecode(bytecode)

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
	new(ByteEater).InitVM().RunBytecode(bytecode)

	fmt.Println(err)

	// if be.__currentResult == nil {
	// 	t.Error("Expected result, got nil")
	// } else if be.__currentResult.isDataframe() == false {
	// 	t.Error("Expected dataframe, got", be.__currentResult)
	// } else if df, err = be.__currentResult.getDataframe(); err == nil {
	// }

}
