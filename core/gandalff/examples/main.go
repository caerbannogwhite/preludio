package main

import (
	"fmt"
	. "gandalff"
	"strings"
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

func Example01() {
	NewBaseDataFrame().
		FromCSV().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		Read().
		Select("department", "age", "weight", "junior").
		GroupBy("department").
		Agg(Min("age"), Max("weight"), Mean("junior"), Count()).
		PrettyPrint()

	// Output:
	// +------------+------------+------------+------------+------------+
	// | department |        age |     weight |     junior |          n |
	// +------------+------------+------------+------------+------------+
	// |     String |    Float64 |    Float64 |    Float64 |      Int32 |
	// +------------+------------+------------+------------+------------+
	// |         HR |         29 |         90 |        0.5 |          2 |
	// |         IT |         25 |         85 |        0.5 |          4 |
	// |   Business |         27 |         65 |        0.5 |          2 |
	// +------------+------------+------------+------------+------------+
}

func Example02() {
	pool := NewStringPool()

	employees := NewBaseDataFrame().
		SetStringPool(pool).
		FromCSV().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		Read()

	departments := NewBaseDataFrame().
		SetStringPool(pool).
		FromCSV().
		SetReader(strings.NewReader(data2)).
		SetDelimiter(',').
		SetHeader(true).
		Read()

	departments.PrettyPrint()

	employees.Join(LEFT_JOIN, departments, "department").
		PrettyPrint()
}

func main() {
	// fmt.Println("Example01:")
	// Example01()

	fmt.Println("Example02:")
	Example02()
}
