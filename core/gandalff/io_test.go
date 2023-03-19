package gandalff

import (
	"os"
	"strings"
	"testing"
)

func Test_TypeGuesser(t *testing.T) {

	// Create a new type guesser.
	guesser := NewTypeGuesser()

	// Test the bool type.
	if guesser.GuessType("true") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("true").ToString())
	}

	if guesser.GuessType("false") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("false").ToString())
	}

	if guesser.GuessType("True") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("True").ToString())
	}

	if guesser.GuessType("False") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("False").ToString())
	}

	if guesser.GuessType("TRUE") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("TRUE").ToString())
	}

	if guesser.GuessType("FALSE") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("FALSE").ToString())
	}

	if guesser.GuessType("t") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("t").ToString())
	}

	if guesser.GuessType("f") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("f").ToString())
	}

	if guesser.GuessType("T") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("T").ToString())
	}

	if guesser.GuessType("F") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("F").ToString())
	}

	if guesser.GuessType("TrUe") != BoolType {
		t.Error("Expected BOOL, got", guesser.GuessType("TrUe").ToString())
	}

	// Test the int type.
	if guesser.GuessType("0") != IntType {
		t.Error("Expected INT, got", guesser.GuessType("0").ToString())
	}

	if guesser.GuessType("1") != IntType {
		t.Error("Expected INT, got", guesser.GuessType("1").ToString())
	}

	if guesser.GuessType("10000") != IntType {
		t.Error("Expected INT, got", guesser.GuessType("10000").ToString())
	}

	if guesser.GuessType("-1") != IntType {
		t.Error("Expected INT, got", guesser.GuessType("-1").ToString())
	}

	if guesser.GuessType("-10000") != IntType {
		t.Error("Expected INT, got", guesser.GuessType("-10000").ToString())
	}

	// Test the float type.
	if guesser.GuessType("0.0") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("0.0").ToString())
	}

	if guesser.GuessType("1.0") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("1.0").ToString())
	}

	if guesser.GuessType("10000.0") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("10000.0").ToString())
	}

	if guesser.GuessType("-1.0") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("-1.0").ToString())
	}

	if guesser.GuessType("-1e3") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("-1e3").ToString())
	}

	if guesser.GuessType("-1e-3") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("-1e-3").ToString())
	}

	if guesser.GuessType("2.0E4") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("2.0E4").ToString())
	}

	if guesser.GuessType("2.0e4") != FloatType {
		t.Error("Expected FLOAT, got", guesser.GuessType("2.0e4").ToString())
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
	if df.Types()[0] != StringType {
		t.Error("Expected STRING, got", df.Types()[0].ToString())
	}

	if df.Types()[1] != IntType {
		t.Error("Expected INT, got", df.Types()[1].ToString())
	}

	if df.Types()[2] != FloatType {
		t.Error("Expected FLOAT, got", df.Types()[2].ToString())
	}

	if df.Types()[3] != BoolType {
		t.Error("Expected BOOL, got", df.Types()[3].ToString())
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

func Test_FromCSV_100000Rows(t *testing.T) {

	f, err := os.OpenFile("testFiles\\organizations-100000.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Error(err)
	}

	// Create a new dataframe from the CSV data.
	df := FromCSV(f, ',', true, 100)
	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	// Check the number of rows.
	if df.NRows() != 100000 {
		t.Error("Expected 100000 rows, got", df.NRows())
	}

	// Check the number of columns.
	if df.NCols() != 9 {
		t.Error("Expected 9 columns, got", df.NCols())
	}

	names := []string{"Index", "Organization Id", "Name", "Website", "Country", "Description", "Founded", "Industry", "Number of employees"}

	// Check the column names.
	for i := 0; i < len(names); i++ {
		if df.Names()[i] != names[i] {
			t.Error("Expected ", names[i], ", got", df.Names()[i])
		}
	}

	// Check the column types.
	if df.Types()[0] != IntType {
		t.Error("Expected INT, got", df.Types()[0].ToString())
	}

	if df.Types()[1] != StringType {
		t.Error("Expected STRING, got", df.Types()[1].ToString())
	}

	if df.Types()[2] != StringType {
		t.Error("Expected STRING, got", df.Types()[2].ToString())
	}

	if df.Types()[3] != StringType {
		t.Error("Expected STRING, got", df.Types()[3].ToString())
	}

	if df.Types()[4] != StringType {
		t.Error("Expected STRING, got", df.Types()[4].ToString())
	}

	if df.Types()[5] != StringType {
		t.Error("Expected STRING, got", df.Types()[5].ToString())
	}

	if df.Types()[6] != IntType {
		t.Error("Expected INT, got", df.Types()[6].ToString())
	}

	if df.Types()[7] != StringType {
		t.Error("Expected STRING, got", df.Types()[7].ToString())
	}

	if df.Types()[8] != IntType {
		t.Error("Expected INT, got", df.Types()[8].ToString())
	}
}

func Test_FromCSV_500000Rows(t *testing.T) {
	f, err := os.OpenFile("testFiles\\organizations-500000.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Error(err)
	}

	// Create a new dataframe from the CSV data.
	df := FromCSV(f, ',', true, 100)
	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	// Check the number of rows.
	if df.NRows() != 500000 {
		t.Error("Expected 100000 rows, got", df.NRows())
	}

	// Check the number of columns.
	if df.NCols() != 9 {
		t.Error("Expected 9 columns, got", df.NCols())
	}

	names := []string{"Index", "Organization Id", "Name", "Website", "Country", "Description", "Founded", "Industry", "Number of employees"}

	// Check the column names.
	for i := 0; i < len(names); i++ {
		if df.Names()[i] != names[i] {
			t.Error("Expected ", names[i], ", got", df.Names()[i])
		}
	}

	// Check the column types.
	if df.Types()[0] != IntType {
		t.Error("Expected INT, got", df.Types()[0].ToString())
	}

	if df.Types()[1] != StringType {
		t.Error("Expected STRING, got", df.Types()[1].ToString())
	}

	if df.Types()[2] != StringType {
		t.Error("Expected STRING, got", df.Types()[2].ToString())
	}

	if df.Types()[3] != StringType {
		t.Error("Expected STRING, got", df.Types()[3].ToString())
	}

	if df.Types()[4] != StringType {
		t.Error("Expected STRING, got", df.Types()[4].ToString())
	}

	if df.Types()[5] != StringType {
		t.Error("Expected STRING, got", df.Types()[5].ToString())
	}

	if df.Types()[6] != IntType {
		t.Error("Expected INT, got", df.Types()[6].ToString())
	}

	if df.Types()[7] != StringType {
		t.Error("Expected STRING, got", df.Types()[7].ToString())
	}

	if df.Types()[8] != IntType {
		t.Error("Expected INT, got", df.Types()[8].ToString())
	}
}
