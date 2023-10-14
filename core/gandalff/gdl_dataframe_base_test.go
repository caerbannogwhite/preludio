package gandalff

import (
	"math"
	"os"
	"runtime"
	"strings"
	"testing"
)

const (
	data1 = `
name,age,weight,junior,department,salary band
Alice C,29,75.0,F,HR,4
John Doe,30,80.5,true,IT,2
Bob,31,85.0,T,IT,4
Jane H,25,60.0,false,IT,4
Mary,28,70.0,false,IT,3
Oliver,32,90.0,true,HR,1
Ursula,27,65.0,f,Business,4
Charlie,33,60.0,t,Business,2
`

	data2 = `
department,number of employees,budget
IT,4,100000
HR,2,50000
Business,2,50000
Operations,4,250000
`
)

func equalFloats(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func Test_BaseDataFrame_Base(t *testing.T) {

}

func Test_BaseDataFrame_Filter(t *testing.T) {
	// Create a new dataframe from the CSV data.
	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	mask := df.Series("department").
		Map(func(v any) any {
			return v.(string) == "IT"
		}).(SeriesBool).
		And(
			df.Series("age").Map(func(v any) any {
				return v.(int64) >= 30
			}).(SeriesBool),
		)

	res := df.Filter(mask.(SeriesBool))
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	if res.NRows() != 2 {
		t.Errorf("Expected 2 rows, got %d", res.NRows())
	}

	names := res.Series("name").Data().([]string)

	if names[0] != "John Doe" {
		t.Errorf("Expected John Doe, got %s", names[0])
	}
	if names[1] != "Bob" {
		t.Errorf("Expected Bob, got %s", names[1])
	}
}

func Benchmark_100000Rows_Filter(b *testing.B) {
	f, err := os.OpenFile("testdata\\organizations-100000.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}

	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(f).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("Country").Map(func(v any) any { return v.(string) == "United States of America" }).(SeriesBool).
				And(
					df.Series("Founded").Map(func(v any) any { return v.(int64) >= 2000 })).(SeriesBool).
				And(
					df.Series("Number of employees").Map(func(v any) any { return v.(int64) < 1000 })).(SeriesBool),
		)
	}
	b.StopTimer()
}

func Test_BaseDataFrame_GroupBy_Count(t *testing.T) {
	// Create a new dataframe from the CSV data.
	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	// Group by department
	res := df.GroupBy("department").Agg(Count())
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp1 := map[string]int64{
		"HR":       2,
		"IT":       4,
		"Business": 2,
	}

	if res.NRows() != len(exp1) {
		t.Errorf("Expected %d rows, got %d", len(exp1), res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		dept := res.Series("department").Get(i).(string)
		n := res.Series("n").Get(i).(int64)

		if n != exp1[dept] {
			t.Errorf("Expected %d, got %d", exp1[dept], n)
		}
	}

	// Group by department and junior
	res = df.Ungroup().GroupBy("junior", "department").Agg(Count())
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp2 := map[bool]map[string]int64{
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
		n := res.Series("n").Get(i).(int64)

		if n != exp2[junior][dept] {
			t.Errorf("Expected %d, got %d", exp2[junior][dept], n)
		}
	}

	// Group by department and junior
	res = df.Ungroup().GroupBy("department", "junior").Agg(Count())
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp3 := map[string]map[bool]int64{
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
		n := res.Series("n").Get(i).(int64)

		if n != exp3[dept][junior] {
			t.Errorf("Expected %d, got %d", exp3[dept][junior], n)
		}
	}

	// Group by department and salary band
	res = df.Ungroup().GroupBy("department", "salary band").Agg(Count())
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp4 := map[string]map[int64]int64{
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
		salaryBand := res.Series("salary band").Get(i).(int64)
		dept := res.Series("department").Get(i).(string)
		n := res.Series("n").Get(i).(int64)

		if n != exp4[dept][salaryBand] {
			t.Errorf("Expected %d, got %d", exp4[dept][salaryBand], n)
		}
	}

	// Group by weight
	res = df.Ungroup().GroupBy("weight").Agg(Count())
	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp5 := map[float64]int64{
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
		n := res.Series("n").Get(i).(int64)

		if n != exp5[weight] {
			t.Errorf("Expected %d, got %d", exp5[weight], n)
		}
	}
}

func Benchmark_100000Rows_GroupBy_Count(b *testing.B) {

	f, err := os.OpenFile("testdata\\organizations-100000.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}

	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(f).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	f.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Ungroup().GroupBy("Country", "Industry").Agg(Count())
	}
	b.StopTimer()
}

func Test_BaseDataFrame_GroupBy_Sum(t *testing.T) {
	// Create a new dataframe from the CSV data.
	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	res := df.GroupBy("department").
		Agg(Sum("age"), Sum("weight"), Sum("junior"), Sum("salary band"))

	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp := map[string][]float64{
		"HR":       {61.0, 165.0, 1.0, 5.0},
		"IT":       {114.0, 295.5, 2.0, 13.0},
		"Business": {60.0, 125.0, 1.0, 6.0},
	}

	if res.NRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		dept := res.Series("department").Get(i).(string)
		age := res.Series("age").Get(i).(float64)
		weight := res.Series("weight").Get(i).(float64)
		junior := res.Series("junior").Get(i).(float64)
		salary := res.Series("salary band").Get(i).(float64)

		if age != exp[dept][0] {
			t.Errorf("Expected 'age' %f, got %f", exp[dept][0], age)
		}

		if weight != exp[dept][1] {
			t.Errorf("Expected 'weight' %f, got %f", exp[dept][1], weight)
		}

		if junior != exp[dept][2] {
			t.Errorf("Expected 'junior' %f, got %f", exp[dept][2], junior)
		}

		if salary != exp[dept][3] {
			t.Errorf("Expected 'salary band' %f, got %f", exp[dept][3], salary)
		}
	}
}

// func Test_BaseDataFrame_GroupBy_Min(t *testing.T) {
// 	// Create a new dataframe from the CSV data.
// 	df := NewBaseDataFrame(ctx).FromCSV().
// 		SetReader(strings.NewReader(data1)).
// 		SetDelimiter(',').
// 		SetHeader(true).
// 		SetGuessDataTypeLen(3).
// 		Read()

// 	if df.GetError() != nil {
// 		t.Error(df.GetError())
// 	}

// 	res := df.GroupBy("department").
// 		Agg(Min("age"), Min("weight"), Min("junior"), Min("salary band"))

// 	if res.GetError() != nil {
// 		t.Error(res.GetError())
// 	}

// 	exp := map[string][]float64{
// 		"HR":       {29.0, 75.0, 0.0, 1.0},
// 		"IT":       {25.0, 60.0, 0.0, 2.0},
// 		"Business": {27.0, 60.0, 0.0, 2.0},
// 	}

// 	if res.NRows() != 3 {
// 		t.Errorf("Expected 3 rows, got %d", res.NRows())
// 	}

// 	for i := 0; i < res.NRows(); i++ {
// 		dept := res.Series("department").Get(i).(string)
// 		age := res.Series("age").Get(i).(float64)
// 		weight := res.Series("weight").Get(i).(float64)
// 		junior := res.Series("junior").Get(i).(float64)
// 		salary := res.Series("salary band").Get(i).(float64)

// 		if age != exp[dept][0] {
// 			t.Errorf("Expected 'age' %f, got %f", exp[dept][0], age)
// 		}

// 		if weight != exp[dept][1] {
// 			t.Errorf("Expected 'weight' %f, got %f", exp[dept][1], weight)
// 		}

// 		if junior != exp[dept][2] {
// 			t.Errorf("Expected 'junior' %f, got %f", exp[dept][2], junior)
// 		}

// 		if salary != exp[dept][3] {
// 			t.Errorf("Expected 'salary band' %f, got %f", exp[dept][3], salary)
// 		}
// 	}
// }

// func Test_BaseDataFrame_GroupBy_Max(t *testing.T) {
// 	// Create a new dataframe from the CSV data.
// 	df := NewBaseDataFrame(ctx).FromCSV().
// 		SetReader(strings.NewReader(data1)).
// 		SetDelimiter(',').
// 		SetHeader(true).
// 		SetGuessDataTypeLen(3).
// 		Read()

// 	if df.GetError() != nil {
// 		t.Error(df.GetError())
// 	}

// 	res := df.GroupBy("department").
// 		Agg(Max("age"), Max("weight"), Max("junior"), Max("salary band"))

// 	if res.GetError() != nil {
// 		t.Error(res.GetError())
// 	}

// 	exp := map[string][]float64{
// 		"HR":       {32.0, 90.0, 1.0, 4.0},
// 		"IT":       {31.0, 85.0, 1.0, 4.0},
// 		"Business": {33.0, 65.0, 1.0, 4.0},
// 	}

// 	if res.NRows() != 3 {
// 		t.Errorf("Expected 3 rows, got %d", res.NRows())
// 	}

// 	for i := 0; i < res.NRows(); i++ {
// 		dept := res.Series("department").Get(i).(string)
// 		age := res.Series("age").Get(i).(float64)
// 		weight := res.Series("weight").Get(i).(float64)
// 		junior := res.Series("junior").Get(i).(float64)
// 		salary := res.Series("salary band").Get(i).(float64)

// 		if age != exp[dept][0] {
// 			t.Errorf("Expected 'age' %f, got %f", exp[dept][0], age)
// 		}

// 		if weight != exp[dept][1] {
// 			t.Errorf("Expected 'weight' %f, got %f", exp[dept][1], weight)
// 		}

// 		if junior != exp[dept][2] {
// 			t.Errorf("Expected 'junior' %f, got %f", exp[dept][2], junior)
// 		}

// 		if salary != exp[dept][3] {
// 			t.Errorf("Expected 'salary band' %f, got %f", exp[dept][3], salary)
// 		}
// 	}
// }

func Test_BaseDataFrame_GroupBy_Mean(t *testing.T) {
	// Create a new dataframe from the CSV data.
	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	if df.GetError() != nil {
		t.Error(df.GetError())
	}

	res := df.GroupBy("department").
		Agg(Mean("age"), Mean("weight"), Mean("junior"), Mean("salary band"))

	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	exp := map[string][]float64{
		"Business": {30.0, 62.5, 0.5, 3.0},
		"HR":       {30.5, 82.5, 0.5, 2.5},
		"IT":       {28.5, 73.875, 0.5, 3.25},
	}

	if res.NRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", res.NRows())
	}

	for i := 0; i < res.NRows(); i++ {
		dept := res.Series("department").Get(i).(string)
		age := res.Series("age").Get(i).(float64)
		weight := res.Series("weight").Get(i).(float64)
		junior := res.Series("junior").Get(i).(float64)
		salary := res.Series("salary band").Get(i).(float64)

		if age != exp[dept][0] {
			t.Errorf("Expected 'age' %f, got %f", exp[dept][0], age)
		}

		if weight != exp[dept][1] {
			t.Errorf("Expected 'weight' %f, got %f", exp[dept][1], weight)
		}

		if junior != exp[dept][2] {
			t.Errorf("Expected 'junior' %f, got %f", exp[dept][2], junior)
		}

		if salary != exp[dept][3] {
			t.Errorf("Expected 'salary band' %f, got %f", exp[dept][3], salary)
		}
	}
}

func Benchmark_100000Rows_GroupBy_Mean(b *testing.B) {

	f, err := os.OpenFile("testdata\\organizations-100000.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}

	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(f).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	f.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Ungroup().GroupBy("Country", "Industry").Agg(Mean("Number of employees"))
	}
	b.StopTimer()
}

func Benchmark_500000Rows_GroupBy_Mean(b *testing.B) {

	f, err := os.OpenFile("testdata\\organizations-500000.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}

	df := NewBaseDataFrame(ctx).FromCSV().
		SetReader(f).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	f.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Ungroup().GroupBy("Country", "Industry").Agg(Mean("Number of employees"))
	}
	b.StopTimer()
}

// func Test_BaseDataFrame_GroupBy_Std(t *testing.T) {
// 	// Create a new dataframe from the CSV data.
// 	df := NewBaseDataFrame(ctx).FromCSV().
// 		SetReader(strings.NewReader(data1)).
// 		SetDelimiter(',').
// 		SetHeader(true).
// 		SetGuessDataTypeLen(3).
// 		Read()

// 	if df.GetError() != nil {
// 		t.Error(df.GetError())
// 	}

// 	res := df.GroupBy("department").
// 		Agg(Std("age"), Std("weight"), Std("junior"), Std("salary band"))

// 	if res.GetError() != nil {
// 		t.Error(res.GetError())
// 	}

// 	exp := map[string][]float64{
// 		"Business": {3.000000, 2.500000, 0.5, 1.000000},
// 		"HR":       {1.500000, 7.500000, 0.5, 1.500000},
// 		"IT":       {2.2912878475, 9.6848786776, 0.5, 0.8291561976},
// 	}

// 	if res.NRows() != 3 {
// 		t.Errorf("Expected 3 rows, got %d", res.NRows())
// 	}

// 	for i := 0; i < res.NRows(); i++ {
// 		dept := res.Series("department").Get(i).(string)
// 		age := res.Series("age").Get(i).(float64)
// 		weight := res.Series("weight").Get(i).(float64)
// 		junior := res.Series("junior").Get(i).(float64)
// 		salary := res.Series("salary band").Get(i).(float64)

// 		if equalFloats(age, exp[dept][0], 10e-8) == false {
// 			t.Errorf("Expected 'age' %.10f, got %.10f", exp[dept][0], age)
// 		}

// 		if equalFloats(weight, exp[dept][1], 10e-8) == false {
// 			t.Errorf("Expected 'weight' %.10f, got %.10f", exp[dept][1], weight)
// 		}

// 		if equalFloats(junior, exp[dept][2], 10e-8) == false {
// 			t.Errorf("Expected 'junior' %.10f, got %.10f", exp[dept][2], junior)
// 		}

// 		if equalFloats(salary, exp[dept][3], 10e-8) == false {
// 			t.Errorf("Expected 'salary band' %.10f, got %.10f", exp[dept][3], salary)
// 		}
// 	}
// }

func Test_BaseDataFrame_Join(t *testing.T) {
	dfx := NewBaseDataFrame(ctx).
		AddSeriesFromInt64s("A", []int64{1, 1, 2, 3, 4, 5, 5}, nil, false).
		AddSeriesFromStrings("B", []string{"a", "b", "c", "d", "e", "f", "g"}, nil, false)

	dfy := NewBaseDataFrame(ctx).
		AddSeriesFromInt64s("A", []int64{4, 5, 6, 6}, nil, false).
		AddSeriesFromStrings("C", []string{"h", "i", "j", "k"}, nil, false)

	///////////////////			INNER JOIN

	res := dfx.Join(INNER_JOIN, dfy, "A")

	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	if res.NRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", res.NRows())
	}

	if res.NCols() != 3 {
		t.Errorf("Expected 3 cols, got %d", res.NCols())
	}

	resAexp := []int64{4, 5, 5}
	resBexp := []string{"e", "f", "g"}
	resCexp := []string{"h", "i", "i"}

	if !checkEqSliceInt64(resAexp, res.SeriesAt(0).Data().([]int64), nil, "Inner Join") {
		t.Errorf("Expected %v, got %v", resAexp, res.SeriesAt(0).Data().([]int64))
	}
	if !checkEqSliceString(resBexp, res.SeriesAt(1).Data().([]string), nil, "Inner Join") {
		t.Errorf("Expected %v, got %v", resBexp, res.SeriesAt(1).Data().([]string))
	}
	if !checkEqSliceString(resCexp, res.SeriesAt(2).Data().([]string), nil, "Inner Join") {
		t.Errorf("Expected %v, got %v", resCexp, res.SeriesAt(2).Data().([]string))
	}

	///////////////////			LEFT JOIN

	res = dfx.Join(LEFT_JOIN, dfy, "A")

	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	if res.NRows() != 7 {
		t.Errorf("Expected 7 rows, got %d", res.NRows())
	}

	if res.NCols() != 3 {
		t.Errorf("Expected 3 cols, got %d", res.NCols())
	}

	resAexp = []int64{1, 1, 2, 3, 4, 5, 5}
	resBexp = []string{"a", "b", "c", "d", "e", "f", "g"}
	resCexp = []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "h", "i", "i"}

	if !checkEqSliceInt64(resAexp, res.SeriesAt(0).Data().([]int64), nil, "Left Join") {
		t.Errorf("Expected %v, got %v", resAexp, res.SeriesAt(0).Data().([]int64))
	}
	if !checkEqSliceString(resBexp, res.SeriesAt(1).Data().([]string), nil, "Left Join") {
		t.Errorf("Expected %v, got %v", resBexp, res.SeriesAt(1).Data().([]string))
	}
	if !checkEqSliceString(resCexp, res.SeriesAt(2).Data().([]string), nil, "Left Join") {
		t.Errorf("Expected %v, got %v", resCexp, res.SeriesAt(2).Data().([]string))
	}

	///////////////////			RIGHT JOIN

	res = dfx.Join(RIGHT_JOIN, dfy, "A")

	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	if res.NRows() != 5 {
		t.Errorf("Expected 5 rows, got %d", res.NRows())
	}

	if res.NCols() != 3 {
		t.Errorf("Expected 3 cols, got %d", res.NCols())
	}

	resAexp = []int64{4, 5, 5, 6, 6}
	resBexp = []string{"e", "f", "g", NULL_STRING, NULL_STRING}
	resCexp = []string{"h", "i", "i", "j", "k"}

	if !checkEqSliceInt64(resAexp, res.SeriesAt(0).Data().([]int64), nil, "Right Join") {
		t.Errorf("Expected %v, got %v", resAexp, res.SeriesAt(0).Data().([]int64))
	}
	if !checkEqSliceString(resBexp, res.SeriesAt(1).Data().([]string), nil, "Right Join") {
		t.Errorf("Expected %v, got %v", resBexp, res.SeriesAt(1).Data().([]string))
	}
	if !checkEqSliceString(resCexp, res.SeriesAt(2).Data().([]string), nil, "Right Join") {
		t.Errorf("Expected %v, got %v", resCexp, res.SeriesAt(2).Data().([]string))
	}

	///////////////////			FULL JOIN

	res = dfx.Join(OUTER_JOIN, dfy, "A")

	if res.GetError() != nil {
		t.Error(res.GetError())
	}

	if res.NRows() != 9 {
		t.Errorf("Expected 9 rows, got %d", res.NRows())
	}

	if res.NCols() != 3 {
		t.Errorf("Expected 3 cols, got %d", res.NCols())
	}

	resAexp = []int64{1, 1, 2, 3, 4, 5, 5, 6, 6}
	resBexp = []string{"a", "b", "c", "d", "e", "f", "g", NULL_STRING, NULL_STRING}
	resCexp = []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "h", "i", "i", "j", "k"}

	if !checkEqSliceInt64(resAexp, res.SeriesAt(0).Data().([]int64), nil, "Full Join") {
		t.Errorf("Expected %v, got %v", resAexp, res.SeriesAt(0).Data().([]int64))
	}
	if !checkEqSliceString(resBexp, res.SeriesAt(1).Data().([]string), nil, "Full Join") {
		t.Errorf("Expected %v, got %v", resBexp, res.SeriesAt(1).Data().([]string))
	}
	if !checkEqSliceString(resCexp, res.SeriesAt(2).Data().([]string), nil, "Full Join") {
		t.Errorf("Expected %v, got %v", resCexp, res.SeriesAt(2).Data().([]string))
	}
}

func Test_BaseDataFrame_Sort(t *testing.T) {
	var res DataFrame

	df := NewBaseDataFrame(ctx).
		AddSeriesFromInt64s("A", []int64{1, 5, 2, 1, 4, 1, 5, 1, 2, 1}, nil, false).
		AddSeriesFromStrings("B", []string{"a", "b", "c", "d", "e", "f", "g", "a", "b", "c"}, nil, false).
		AddSeriesFromFloat64s("C", []float64{1.2, 2.3, 3.4, 4.5, 5.6, 7.8, 8.9, 1.2, 2.3, 3.4}, nil, false).
		AddSeriesFromBools("D", []bool{true, false, true, true, false, true, true, false, true, false}, nil, false)

	res = df.OrderBy(Asc("A"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort, column A asc failed")
	}

	res = df.OrderBy(Desc("A"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{5, 5, 4, 2, 2, 1, 1, 1, 1, 1}, nil, "") {
		t.Error("BaseDataFrame Sort, column A desc failed")
	}

	res = df.OrderBy(Asc("B"))
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"a", "a", "b", "b", "c", "c", "d", "e", "f", "g"}, nil, "") {
		t.Error("BaseDataFrame Sort, column B asc failed")
	}

	res = df.OrderBy(Desc("B"))
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"g", "f", "e", "d", "c", "c", "b", "b", "a", "a"}, nil, "") {
		t.Error("BaseDataFrame Sort, column B desc failed")
	}

	res = df.OrderBy(Asc("C"))
	if !checkEqSliceFloat64(res.Series("C").(SeriesFloat64).Float64s(), []float64{1.2, 1.2, 2.3, 2.3, 3.4, 3.4, 4.5, 5.6, 7.8, 8.9}, nil, "") {
		t.Error("BaseDataFrame Sort, column C asc failed")
	}

	res = df.OrderBy(Desc("C"))
	if !checkEqSliceFloat64(res.Series("C").(SeriesFloat64).Float64s(), []float64{8.9, 7.8, 5.6, 4.5, 3.4, 3.4, 2.3, 2.3, 1.2, 1.2}, nil, "") {
		t.Error("BaseDataFrame Sort, column C desc failed")
	}

	res = df.OrderBy(Asc("D"))
	if !checkEqSliceBool(res.Series("D").(SeriesBool).Bools(), []bool{false, false, false, false, true, true, true, true, true, true}, nil, "") {
		t.Error("BaseDataFrame Sort, column D asc failed")
	}

	res = df.OrderBy(Desc("D"))
	if !checkEqSliceBool(res.Series("D").(SeriesBool).Bools(), []bool{true, true, true, true, true, true, false, false, false, false}, nil, "") {
		t.Error("BaseDataFrame Sort, column D desc failed")
	}

	////////////////////////			.Sort() with 2 columns

	res = df.OrderBy(Asc("A"), Asc("B"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"a", "a", "c", "d", "f", "b", "c", "e", "b", "g"}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: B failed")
	}

	res = df.OrderBy(Asc("A"), Desc("B"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"f", "d", "c", "a", "a", "c", "b", "e", "g", "b"}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc: B failed")
	}

	res = df.OrderBy(Desc("A"), Asc("B"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{5, 5, 4, 2, 2, 1, 1, 1, 1, 1}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, B asc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"b", "g", "e", "b", "c", "a", "a", "c", "d", "f"}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, B asc: B failed")
	}

	res = df.OrderBy(Desc("A"), Desc("B"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{5, 5, 4, 2, 2, 1, 1, 1, 1, 1}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, B desc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"g", "b", "e", "c", "b", "f", "d", "c", "a", "a"}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, B desc: B failed")
	}

	res = df.OrderBy(Asc("A"), Asc("C"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, C asc: A failed")
	}
	if !checkEqSliceFloat64(res.Series("C").(SeriesFloat64).Float64s(), []float64{1.2, 1.2, 3.4, 4.5, 7.8, 2.3, 3.4, 5.6, 2.3, 8.9}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, C asc: C failed")
	}

	res = df.OrderBy(Asc("A"), Desc("C"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, C desc: A failed")
	}
	if !checkEqSliceFloat64(res.Series("C").(SeriesFloat64).Float64s(), []float64{7.8, 4.5, 3.4, 1.2, 1.2, 3.4, 2.3, 5.6, 8.9, 2.3}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, C desc: C failed")
	}

	res = df.OrderBy(Desc("A"), Asc("C"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{5, 5, 4, 2, 2, 1, 1, 1, 1, 1}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, C asc: A failed")
	}
	if !checkEqSliceFloat64(res.Series("C").(SeriesFloat64).Float64s(), []float64{2.3, 8.9, 5.6, 2.3, 3.4, 1.2, 1.2, 3.4, 4.5, 7.8}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, C asc: C failed")
	}

	res = df.OrderBy(Desc("A"), Desc("C"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{5, 5, 4, 2, 2, 1, 1, 1, 1, 1}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, C desc: A failed")
	}
	if !checkEqSliceFloat64(res.Series("C").(SeriesFloat64).Float64s(), []float64{8.9, 2.3, 5.6, 3.4, 2.3, 7.8, 4.5, 3.4, 1.2, 1.2}, nil, "") {
		t.Error("BaseDataFrame Sort A desc, C desc: C failed")
	}

	////////////////////////			.Sort() with 3 columns

	res = df.OrderBy(Asc("A"), Asc("B"), Asc("D"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc, D asc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"a", "a", "c", "d", "f", "b", "c", "e", "b", "g"}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc, D asc: B failed")
	}
	if !checkEqSliceBool(res.Series("D").(SeriesBool).Bools(), []bool{false, true, false, true, true, true, true, false, false, true}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc, D asc: D failed")
	}

	res = df.OrderBy(Asc("A"), Asc("B"), Desc("D"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc, D desc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"a", "a", "c", "d", "f", "b", "c", "e", "b", "g"}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc, D desc: B failed")
	}
	if !checkEqSliceBool(res.Series("D").(SeriesBool).Bools(), []bool{true, false, false, true, true, true, true, false, false, true}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc, D desc: D failed")
	}

	res = df.OrderBy(Asc("A"), Desc("B"), Asc("D"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc, D asc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"f", "d", "c", "a", "a", "c", "b", "e", "g", "b"}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc, D asc: B failed")
	}
	if !checkEqSliceBool(res.Series("D").(SeriesBool).Bools(), []bool{true, true, false, false, true, true, true, false, true, false}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc, D asc: D failed")
	}

	res = df.OrderBy(Asc("A"), Desc("B"), Desc("D"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 1, 1, 2, 2, 4, 5, 5}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc, D desc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"f", "d", "c", "a", "a", "c", "b", "e", "g", "b"}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc, D desc: B failed")
	}
	if !checkEqSliceBool(res.Series("D").(SeriesBool).Bools(), []bool{true, true, false, true, false, true, true, false, true, false}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B desc, D desc: D failed")
	}

	////////////////////////

	res = df.OrderBy(Desc("D"), Asc("C"), Desc("B"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 2, 2, 1, 1, 5, 1, 5, 1, 4}, nil, "") {
		t.Error("BaseDataFrame Sort D desc, C asc, B desc: A failed")
	}
	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"a", "b", "c", "d", "f", "g", "a", "b", "c", "e"}, nil, "") {
		t.Error("BaseDataFrame Sort D desc, C asc, B desc: B failed")
	}
	if !checkEqSliceFloat64(res.Series("C").(SeriesFloat64).Float64s(), []float64{1.2, 2.3, 3.4, 4.5, 7.8, 8.9, 1.2, 2.3, 3.4, 5.6}, nil, "") {
		t.Error("BaseDataFrame Sort D desc, C asc, B desc: C failed")
	}
	if !checkEqSliceBool(res.Series("D").(SeriesBool).Bools(), []bool{true, true, true, true, true, true, false, false, false, false}, nil, "") {
		t.Error("BaseDataFrame Sort D desc, C asc, B desc: D failed")
	}
}

func Test_BaseDataFrame_Sort_Nulls(t *testing.T) {
	var res DataFrame

	a := NewSeriesInt64([]int64{1, 4, 2, 1, 4, 1, 4, 1, 2, 1}, nil, true, ctx).
		SetNullMask([]bool{false, false, false, false, false, true, false, false, true, true})
	b := NewSeriesString([]string{"a", "b", "c", "d", "e", "f", "g", "a", "b", "c"}, nil, true, ctx).
		SetNullMask([]bool{true, true, false, false, false, true, false, false, false, false})
	c := NewSeriesFloat64([]float64{1.2, 2.3, 3.4, 4.5, 5.6, 7.8, 8.9, 1.2, 2.3, 3.4}, nil, true, ctx).
		SetNullMask([]bool{false, false, false, false, false, true, false, false, true, true})
	d := NewSeriesBool([]bool{true, false, true, true, false, true, true, false, true, false}, nil, true, ctx).
		SetNullMask([]bool{false, false, false, false, false, true, false, false, true, true})

	df := NewBaseDataFrame(ctx).
		AddSeries("A", a).
		AddSeries("B", b).
		AddSeries("C", c).
		AddSeries("D", d)

	res = df.OrderBy(Asc("A"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 2, 4, 4, 4, 1, 2, 1}, nil, "") {
		t.Error("BaseDataFrame Sort A asc: A failed")
	}
	if !checkEqSliceBool(res.Series("A").GetNullMask(), []bool{false, false, false, false, false, false, false, true, true, true}, nil, "") {
		t.Error("BaseDataFrame Sort A asc: A nullmask failed")
	}

	res = df.OrderBy(Desc("A"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 2, 1, 4, 4, 4, 2, 1, 1, 1}, nil, "") {
		t.Error("BaseDataFrame Sort A desc: A failed")
	}
	if !checkEqSliceBool(res.Series("A").GetNullMask(), []bool{true, true, true, false, false, false, false, false, false, false}, nil, "") {
		t.Error("BaseDataFrame Sort A desc: A nullmask failed")
	}

	res = df.OrderBy(Asc("A"), Asc("B"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 2, 4, 4, 4, 2, 1, 1}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: A failed")
	}
	if !checkEqSliceBool(res.Series("A").GetNullMask(), []bool{false, false, false, false, false, false, false, true, true, true}, nil, "") {
		t.Error("BaseDataFrame Sort A asc: A nullmask failed")
	}

	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{"a", "d", NULL_STRING, "c", "e", "g", NULL_STRING, "b", "c", NULL_STRING}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: B failed")
	}
	if !checkEqSliceBool(res.Series("B").GetNullMask(), []bool{false, false, true, false, false, false, true, false, false, true}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: B nullmask failed")
	}

	res = df.OrderBy(Asc("A"), Desc("B"))
	if !checkEqSliceInt64(res.Series("A").(SeriesInt64).Int64s(), []int64{1, 1, 1, 2, 4, 4, 4, 1, 1, 2}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: A failed")
	}
	if !checkEqSliceBool(res.Series("A").GetNullMask(), []bool{false, false, false, false, false, false, false, true, true, true}, nil, "") {
		t.Error("BaseDataFrame Sort A asc: A nullmask failed")
	}

	if !checkEqSliceString(res.Series("B").(SeriesString).Strings(), []string{NULL_STRING, "d", "a", "c", NULL_STRING, "g", "e", NULL_STRING, "c", "b"}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: B failed")
	}
	if !checkEqSliceBool(res.Series("B").GetNullMask(), []bool{true, false, false, false, true, false, false, true, false, false}, nil, "") {
		t.Error("BaseDataFrame Sort A asc, B asc: B nullmask failed")
	}
}
