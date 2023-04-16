package gandalff

import (
	"os"
	"strings"
	"testing"
	"typesys"
)

func Test_TypeGuesser(t *testing.T) {

	// Create a new type guesser.
	guesser := NewTypeGuesser()

	// Test the bool type.
	if guesser.GuessType("true") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("true").ToString())
	}

	if guesser.GuessType("false") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("false").ToString())
	}

	if guesser.GuessType("True") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("True").ToString())
	}

	if guesser.GuessType("False") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("False").ToString())
	}

	if guesser.GuessType("TRUE") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("TRUE").ToString())
	}

	if guesser.GuessType("FALSE") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("FALSE").ToString())
	}

	if guesser.GuessType("t") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("t").ToString())
	}

	if guesser.GuessType("f") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("f").ToString())
	}

	if guesser.GuessType("T") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("T").ToString())
	}

	if guesser.GuessType("F") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("F").ToString())
	}

	if guesser.GuessType("TrUe") != typesys.BoolType {
		t.Error("Expected Bool, got", guesser.GuessType("TrUe").ToString())
	}

	// Test the int type.
	if guesser.GuessType("0") != typesys.Int32Type {
		t.Error("Expected Int32, got", guesser.GuessType("0").ToString())
	}

	if guesser.GuessType("1") != typesys.Int32Type {
		t.Error("Expected Int32, got", guesser.GuessType("1").ToString())
	}

	if guesser.GuessType("10000") != typesys.Int32Type {
		t.Error("Expected Int32, got", guesser.GuessType("10000").ToString())
	}

	if guesser.GuessType("-1") != typesys.Int32Type {
		t.Error("Expected Int32, got", guesser.GuessType("-1").ToString())
	}

	if guesser.GuessType("-10000") != typesys.Int32Type {
		t.Error("Expected Int32, got", guesser.GuessType("-10000").ToString())
	}

	// Test the float type.
	if guesser.GuessType("0.0") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("0.0").ToString())
	}

	if guesser.GuessType("1.0") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("1.0").ToString())
	}

	if guesser.GuessType("10000.0") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("10000.0").ToString())
	}

	if guesser.GuessType("-1.0") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("-1.0").ToString())
	}

	if guesser.GuessType("-1e3") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("-1e3").ToString())
	}

	if guesser.GuessType("-1e-3") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("-1e-3").ToString())
	}

	if guesser.GuessType("2.0E4") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("2.0E4").ToString())
	}

	if guesser.GuessType("2.0e4") != typesys.Float64Type {
		t.Error("Expected Float64, got", guesser.GuessType("2.0e4").ToString())
	}
}

func Test_FromCSV(t *testing.T) {

	data := `name,age,weight,junior
Alice C,29,75.0,F
John Doe,30,80.5,true
Bob,31,85.0,T
Jane H,25,60.0,false
Mary,28,70.0,false
Oliver,32,90.0,true
Ursula,27,65.0,f
Charlie,33,95.0,t
`

	// Create a new dataframe from the CSV data.
	df := FromCSV(strings.NewReader(data), ',', true, 3)
	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	// Check the number of rows.
	if df.NRows() != 8 {
		t.Error("Expected 8 rows, got", df.NRows())
	}

	// Check the number of columns.
	if df.NCols() != 4 {
		t.Error("Expected 4 columns, got", df.NCols())
	}

	// Check the column names.
	if df.Names()[0] != "name" {
		t.Error("Expected 'name', got", df.Names()[0])
	}

	if df.Names()[1] != "age" {
		t.Error("Expected 'age', got", df.Names()[1])
	}

	if df.Names()[2] != "weight" {
		t.Error("Expected 'weight', got", df.Names()[2])
	}

	if df.Names()[3] != "junior" {
		t.Error("Expected 'junior', got", df.Names()[3])
	}

	// Check the column types.
	if df.Types()[0] != typesys.StringType {
		t.Error("Expected String, got", df.Types()[0].ToString())
	}

	if df.Types()[1] != typesys.Int32Type {
		t.Error("Expected Int32, got", df.Types()[1].ToString())
	}

	if df.Types()[2] != typesys.Float64Type {
		t.Error("Expected Float64, got", df.Types()[2].ToString())
	}

	if df.Types()[3] != typesys.BoolType {
		t.Error("Expected Bool, got", df.Types()[3].ToString())
	}

	// Check the values.
	if df.SeriesAt(0).Data().([]string)[0] != "Alice C" {
		t.Error("Expected 'Alice C', got", df.SeriesAt(0).Data().([]string)[0])
	}

	if df.SeriesAt(0).Data().([]string)[1] != "John Doe" {
		t.Error("Expected 'John Doe', got", df.SeriesAt(0).Data().([]string)[1])
	}

	if df.SeriesAt(0).Data().([]string)[2] != "Bob" {
		t.Error("Expected 'Bob', got", df.SeriesAt(0).Data().([]string)[2])
	}

	if df.SeriesAt(0).Data().([]string)[3] != "Jane H" {
		t.Error("Expected 'Jane H', got", df.SeriesAt(0).Data().([]string)[3])
	}

	if df.SeriesAt(1).Data().([]int)[4] != 28 {
		t.Error("Expected 28, got", df.SeriesAt(1).Data().([]int)[4])
	}

	if df.SeriesAt(1).Data().([]int)[5] != 32 {
		t.Error("Expected 32, got", df.SeriesAt(1).Data().([]int)[5])
	}

	if df.SeriesAt(1).Data().([]int)[6] != 27 {
		t.Error("Expected 27, got", df.SeriesAt(1).Data().([]int)[6])
	}

	if df.SeriesAt(1).Data().([]int)[7] != 33 {
		t.Error("Expected 33, got", df.SeriesAt(1).Data().([]int)[7])
	}

	if df.SeriesAt(2).Data().([]float64)[0] != 75.0 {
		t.Error("Expected 75.0, got", df.SeriesAt(2).Data().([]float64)[0])
	}

	if df.SeriesAt(2).Data().([]float64)[1] != 80.5 {
		t.Error("Expected 80.5, got", df.SeriesAt(2).Data().([]float64)[1])
	}

	if df.SeriesAt(2).Data().([]float64)[2] != 85.0 {
		t.Error("Expected 85.0, got", df.SeriesAt(2).Data().([]float64)[2])
	}

	if df.SeriesAt(2).Data().([]float64)[3] != 60.0 {
		t.Error("Expected 60.0, got", df.SeriesAt(2).Data().([]float64)[3])
	}

	if df.SeriesAt(3).Data().([]bool)[4] != false {
		t.Error("Expected false, got", df.SeriesAt(3).Data().([]bool)[4])
	}

	if df.SeriesAt(3).Data().([]bool)[5] != true {
		t.Error("Expected true, got", df.SeriesAt(3).Data().([]bool)[5])
	}

	if df.SeriesAt(3).Data().([]bool)[6] != false {
		t.Error("Expected false, got", df.SeriesAt(3).Data().([]bool)[6])
	}

	if df.SeriesAt(3).Data().([]bool)[7] != true {
		t.Error("Expected true, got", df.SeriesAt(3).Data().([]bool)[7])
	}
}

func Benchmark_FromCSV_100000Rows(b *testing.B) {

	// Create a new dataframe from the CSV data.
	var df *GDLDataFrame

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f, err := os.OpenFile("testdata\\organizations-100000.csv", os.O_RDONLY, 0666)
		if err != nil {
			b.Error(err)
		}
		df = FromCSV(f, ',', true, 100)
		f.Close()
	}
	b.StopTimer()

	if df.GetError() != nil {
		b.Error(df.GetError())
	}

	// Check the number of rows.
	if df.NRows() != 100000 {
		b.Error("Expected 100000 rows, got", df.NRows())
	}

	// Check the number of columns.
	if df.NCols() != 9 {
		b.Error("Expected 9 columns, got", df.NCols())
	}

	names := []string{"Index", "Organization Id", "Name", "Website", "Country", "Description", "Founded", "Industry", "Number of employees"}

	// Check the column names.
	for i := 0; i < len(names); i++ {
		if df.Names()[i] != names[i] {
			b.Error("Expected ", names[i], ", got", df.Names()[i])
		}
	}

	// Check the column types.
	if df.Types()[0] != typesys.Int32Type {
		b.Error("Expected Int32, got", df.Types()[0].ToString())
	}

	if df.Types()[1] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[1].ToString())
	}

	if df.Types()[2] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[2].ToString())
	}

	if df.Types()[3] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[3].ToString())
	}

	if df.Types()[4] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[4].ToString())
	}

	if df.Types()[5] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[5].ToString())
	}

	if df.Types()[6] != typesys.Int32Type {
		b.Error("Expected Int32, got", df.Types()[6].ToString())
	}

	if df.Types()[7] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[7].ToString())
	}

	if df.Types()[8] != typesys.Int32Type {
		b.Error("Expected Int32, got", df.Types()[8].ToString())
	}
}

func Benchmark_FromCSV_500000Rows(b *testing.B) {
	// Create a new dataframe from the CSV data.
	var df *GDLDataFrame

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f, err := os.OpenFile("testdata\\organizations-500000.csv", os.O_RDONLY, 0666)
		if err != nil {
			b.Error(err)
		}
		df = FromCSV(f, ',', true, 100)
		f.Close()
	}
	b.StopTimer()

	if df.GetError() != nil {
		b.Error(df.GetError())
	}

	// Check the number of rows.
	if df.NRows() != 500000 {
		b.Error("Expected 100000 rows, got", df.NRows())
	}

	// Check the number of columns.
	if df.NCols() != 9 {
		b.Error("Expected 9 columns, got", df.NCols())
	}

	names := []string{"Index", "Organization Id", "Name", "Website", "Country", "Description", "Founded", "Industry", "Number of employees"}

	// Check the column names.
	for i := 0; i < len(names); i++ {
		if df.Names()[i] != names[i] {
			b.Error("Expected ", names[i], ", got", df.Names()[i])
		}
	}

	// Check the column types.
	if df.Types()[0] != typesys.Int32Type {
		b.Error("Expected Int32, got", df.Types()[0].ToString())
	}

	if df.Types()[1] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[1].ToString())
	}

	if df.Types()[2] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[2].ToString())
	}

	if df.Types()[3] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[3].ToString())
	}

	if df.Types()[4] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[4].ToString())
	}

	if df.Types()[5] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[5].ToString())
	}

	if df.Types()[6] != typesys.Int32Type {
		b.Error("Expected Int32, got", df.Types()[6].ToString())
	}

	if df.Types()[7] != typesys.StringType {
		b.Error("Expected String, got", df.Types()[7].ToString())
	}

	if df.Types()[8] != typesys.Int32Type {
		b.Error("Expected Int32, got", df.Types()[8].ToString())
	}
}
