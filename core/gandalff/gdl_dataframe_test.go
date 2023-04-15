package gandalff

import (
	"strings"
	"testing"
)

func Test_GDataFrame_Base(t *testing.T) {

}

func Test_GDataFrame_GroupBy_Count(t *testing.T) {

	data := `name,age,weight,junior,department,salary band
Alice C,29,75.0,F,HR,4
John Doe,30,80.5,true,IT,2
Bob,31,85.0,T,IT,4
Jane H,25,60.0,false,IT,4
Mary,28,70.0,false,IT,3
Oliver,32,90.0,true,HR,1
Ursula,27,65.0,f,Business,4
Charlie,33,60.0,t,Business,2
`

	// Create a new dataframe from the CSV data.
	df := FromCSV(strings.NewReader(data), ',', true, 3)
	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	// Group by department
	res := df.GroupBy("department").Count("n")
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp1 := map[string]int{
		"HR":       2,
		"IT":       4,
		"Business": 2,
	}

	if res.NRows() != len(exp1) {
		t.Errorf("Expected %d rows, got %d", len(exp1), res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		dept := res.Series("department").Get(i).(string)
		n := res.Series("n").Get(i).(int)

		if n != exp1[dept] {
			t.Errorf("Expected %d, got %d", exp1[dept], n)
		}
	}

	// Group by department and junior
	res = df.Ungroup().GroupBy("junior", "department").Count("n")
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp2 := map[bool]map[string]int{
		true: {
			"HR":       1,
			"IT":       2,
			"Business": 1,
		},
		false: {
			"HR":       1,
			"IT":       2,
			"Business": 1,
		},
	}

	if res.NRows() != 6 {
		t.Errorf("Expected 6 rows, got %d", res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		junior := res.Series("junior").Get(i).(bool)
		dept := res.Series("department").Get(i).(string)
		n := res.Series("n").Get(i).(int)

		if n != exp2[junior][dept] {
			t.Errorf("Expected %d, got %d", exp2[junior][dept], n)
		}
	}

	// Group by department and junior
	res = df.Ungroup().GroupBy("department", "junior").Count("n")
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp3 := map[string]map[bool]int{
		"HR": {
			true:  1,
			false: 1,
		},
		"IT": {
			true:  2,
			false: 2,
		},
		"Business": {
			true:  1,
			false: 1,
		},
	}

	if res.NRows() != 6 {
		t.Errorf("Expected 6 rows, got %d", res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		junior := res.Series("junior").Get(i).(bool)
		dept := res.Series("department").Get(i).(string)
		n := res.Series("n").Get(i).(int)

		if n != exp3[dept][junior] {
			t.Errorf("Expected %d, got %d", exp3[dept][junior], n)
		}
	}

	// Group by department and salary band
	res = df.Ungroup().GroupBy("department", "salary band").Count("n")
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp4 := map[string]map[int]int{
		"HR": {
			1: 1,
			4: 1,
		},
		"IT": {
			2: 1,
			4: 2,
			3: 1,
		},
		"Business": {
			4: 1,
			2: 1,
		},
	}

	if res.NRows() != 7 {
		t.Errorf("Expected 7 rows, got %d", res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		salaryBand := res.Series("salary band").Get(i).(int)
		dept := res.Series("department").Get(i).(string)
		n := res.Series("n").Get(i).(int)

		if n != exp4[dept][salaryBand] {
			t.Errorf("Expected %d, got %d", exp4[dept][salaryBand], n)
		}
	}

	// Group by weight
	res = df.Ungroup().GroupBy("weight").Count("n")
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp5 := map[float64]int{
		75.0: 1,
		80.5: 1,
		85.0: 1,
		60.0: 2,
		70.0: 1,
		90.0: 1,
		65.0: 1,
	}

	if res.NRows() != len(exp5) {
		t.Errorf("Expected %d rows, got %d", len(exp5), res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		weight := res.Series("weight").Get(i).(float64)
		n := res.Series("n").Get(i).(int)

		if n != exp5[weight] {
			t.Errorf("Expected %d, got %d", exp5[weight], n)
		}
	}
}
