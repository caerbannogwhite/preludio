## GANDALFF: Golang, ANother DAta Library For Fun

Or Gdl: Golang Data Library

### What is it?

Gandalff is a library for data manipulation in Go. It is inspired by the R language and the dplyr package.
It supports nullable types: null data is optimized for memory usage.
`GDLSeriesBool` stores the boolean data as bits, and `GDLSeriesString` stores the string data in a string pool.

### Why?

The primary purpose of this library is to have an easy-to-use data manipulation library for Go and to be as close as possible to the R language and the Dplyr package.
It also has to be performant; hopefully, it will become as fast as Polars.

### Examples

```go
func Example01() {
	data := `
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

	NewBaseDataFrame().
		FromCSV().
		SetReader(strings.NewReader(data)).
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
```

### Supported data types

The data types not checked are not yet supported, but might be in the future.

- [ ] Bool
- [x] Bool (memory optimized)
- [ ] Int16
- [x] Int32 (Golang int)
- [x] Int64
- [ ] Float32
- [x] Float64
- [ ] Complex64
- [ ] Complex128
- [x] String
- [ ] Time
- [ ] Duration

### Supported operations for Series

- [x] Filter

  - [x] Filter (by Bool series)
  - [x] FilterByMask
  - [x] FilterByIndex

- [ ] Group

  - [x] Group
  - [x] SubGroup
  - [ ] Group with nulls
  - [ ] SubGroup with nulls

- [x] Map
- [ ] Sort

  - [x] Sort
  - [ ] SortRev
  - [ ] Sort with nulls
  - [ ] SortRev with nulls


- [ ] Take

### Supported operations for DataFrame

- [x] Agg
- [x] Filter
- [x] GroupBy
- [ ] Join

	- [ ] Inner
	- [ ] Left
	- [ ] Right
	- [ ] Outer
	- [ ] Inner with nulls
	- [ ] Left with nulls
	- [ ] Right with nulls
	- [ ] Outer with nulls

- [ ] Map
- [ ] OrderBy
- [x] Select
- [ ] Take

### Supported stats functions

- [x] Count
- [x] Sum
- [x] Mean
- [ ] Median
- [x] Min
- [x] Max
- [x] StdDev
- [ ] Variance
- [ ] Quantile

### Implementation details

This is how the interface for the Series type currently looks like
with all the methods that are currently implemented.

```go
type GDLSeries interface {

	// Basic accessors.

	// Returns the number of elements in the series.
	Len() int
	// Returns the name of the series.
	Name() string
	// Returns the type of the series.
	Type() typesys.BaseType

	// Returns if the series is grouped.
	IsGrouped() bool
	// Returns if the series admits null values.
	IsNullable() bool
	// Returns if the series is sorted.
	IsSorted() GDLSeriesSortOrder

	// Nullability operations.

	// Returns if the series has null values.
	HasNull() bool
	// Returns the number of null values in the series.
	NullCount() int
	// Returns the number of non-null values in the series.
	NonNullCount() int
	// Returns if the element at index i is null.
	IsNull(i int) bool
	// Sets the element at index i to null.
	SetNull(i int) GDLSeries
	// Returns the null mask of the series.
	GetNullMask() []bool
	// Sets the null mask of the series.
	SetNullMask(mask []bool) GDLSeries
	// Makes the series nullable.
	MakeNullable() GDLSeries

	// Get the element at index i.
	Get(i int) any
	// Get the element at index i as a string.
	GetString(i int) string
	// Set the element at index i.
	Set(i int, v any) GDLSeries
	// Take the elements according to the given interval.
	Take(start, end, step int) GDLSeries

	// Sort Interface.
	Less(i, j int) bool
	Swap(i, j int)

	// Append elements to the series.
	Append(v any) GDLSeries
	// AppendRaw appends a value or a slice of values to the series.
	AppendRaw(v any) GDLSeries
	// Append nullable elements to the series.
	AppendNullable(v any) GDLSeries
	// Append a series to the series.
	AppendSeries(other GDLSeries) GDLSeries

	// All-data accessors.

	// Returns the actual data of the series.
	Data() any
	// Returns the nullable data of the series.
	DataAsNullable() any
	// Returns the data of the series as a slice of strings.
	DataAsString() []string

	// Casts the series to a given type.
	Cast(t typesys.BaseType, stringPool *StringPool) GDLSeries
	// Copies the series.
	Copy() GDLSeries

	// Series operations.

	// Filters out the elements by the given mask series.
	Filter(mask GDLSeriesBool) GDLSeries
	// Filters out the elements by the given mask.
	FilterByMask(mask []bool) GDLSeries
	// Filters out the elements by the given indices.
	FilterByIndeces(indices []int) GDLSeries

	// Maps the elements of the series.
	Map(f GDLMapFunc, stringPool *StringPool) GDLSeries

	// Group the elements in the series.
	Group() GDLSeries
	SubGroup(gp GDLSeriesPartition) GDLSeries

	// Get the partition of the series.
	GetPartition() GDLSeriesPartition

	// Sorts the elements of the series.
	Sort() GDLSeries
	SortRev() GDLSeries
}
```

### TODO

- [ ] Implement non memory optimized Bool series.
- [ ] Implement memory optimized Bool series with uint64.
- [ ] Implement int32 series with int32.
- [ ] Using uint64 for null mask.
- [ ] Implement and test grouped sorting for all types.
- [ ] Implement and test grouped with nulls.
- [ ] Implement chunked series.
- [ ] Implement CSV writer.
- [ ] Implement Excel reader and writer (https://github.com/tealeg/xlsx).
- [ ] Implement JSON reader and writer.