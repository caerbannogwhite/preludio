## GANDALFF: Golang, ANother DAta Library For Fun

Or Gdl: Golang Data Library

### What is it?

Gandalff is a library for data manipulation in Go. It is inspired by the R language and the dplyr package.
It supports nullable types: null data is optimized for memory usage.
`GDLSeriesBool` stores the boolean data as bits, and `GDLSeriesString` stores the string data in a string pool.

### Why?

I wanted to learn Go and I wanted to learn how to write a library. I also wanted to learn how to write a library that is easy to use and that is easy to extend.

### Supported data types

The data types not checked are not yet supported, but might be in the future.

- [x] Bool
- [ ] Int16
- [x] Int32 (Golang int)
- [ ] Int64
- [ ] Float32
- [x] Float64
- [ ] Complex64
- [ ] Complex128
- [x] String
- [ ] Time
- [ ] Duration

### Supported operations for Series

- [x] Filter
  - [x] FilterByMask
  - [x] FilterByIndex

- [x] Group
  - [x] Group
  - [x] SubGroup

- [x] Map
- [ ] Sort

### Supported operations for DataFrame

- [x] Filter
- [x] GroupBy
- [ ] Map
- [x] Select
- [ ] OrderBy
- [ ] Join

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

	// Returns the length of the series.
	Len() int
	// Returns if the series admits null values.
	IsNullable() bool
	// Makes the series nullable.
	MakeNullable() GDLSeries
	// Returns the name of the series.
	Name() string
	// Returns the type of the series.
	Type() typesys.BaseType
	// Returns if the series has null values.
	HasNull() bool
	// Returns the number of null values in the series.
	NullCount() int
	// Returns the number of non-null values in the series.
	NonNullCount() int
	// Returns if the element at index i is null.
	IsNull(i int) bool
	// Sets the element at index i to null.
	SetNull(i int) error
	// Returns the null mask of the series.
	GetNullMask() []bool
	// Sets the null mask of the series.
	SetNullMask(mask []bool) error

	// Get the element at index i.
	Get(i int) interface{}
	// Set the element at index i.
	Set(i int, v interface{})

	// Append elements to the series.
	Append(v interface{}) GDLSeries
	// AppendRaw appends a value or a slice of values to the series.
	AppendRaw(v interface{}) GDLSeries
	// Append nullable elements to the series.
	AppendNullable(v interface{}) GDLSeries
	// Append a series to the series.
	AppendSeries(other GDLSeries) GDLSeries

	// All-data accessors.

	// Returns the actual data of the series.
	Data() interface{}
	// Returns the nullable data of the series.
	NullableData() interface{}
	// Returns the data of the series as a slice of strings.
	StringData() []string

	// Copies the series.
	Copy() GDLSeries

	// Series operations.

	// Filters out the elements by the given mask.
	FilterByMask(mask []bool) GDLSeries
	// Filters out the elements by the given indices.
	FilterByIndeces(indices []int) GDLSeries

	// Maps the elements of the series.
	Map(f GDLMapFunc, stringPool *StringPool) GDLSeries

	// Group the elements in the series.
	Group() GDLSeriesPartition
	SubGroup(gp GDLSeriesPartition) GDLSeriesPartition
}
```