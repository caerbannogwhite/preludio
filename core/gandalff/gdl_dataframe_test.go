package gandalff

import (
	"strings"
	"testing"
)

func Test_GDataFrame_Base(t *testing.T) {

}

func Test_GDataFrame_GroupBy(t *testing.T) {

	data := `name,age,weight,junior,department
Alice C,29,75.0,F,HR
John Doe,30,80.5,true,IT
Bob,31,85.0,T,IT
Jane H,25,60.0,false,IT
Mary,28,70.0,false,IT
Oliver,32,90.0,true,HR
Ursula,27,65.0,f,Business
Charlie,33,95.0,t,Business
`

	// Create a new dataframe from the CSV data.
	df := FromCSV(strings.NewReader(data), ',', true, 3)
	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	// df.GroupBy("department").Count().PrettyPrint()

}
