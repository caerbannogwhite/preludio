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
  - [x] Filter (by Bool series)
  - [x] FilterByMask
  - [x] FilterByIndex

- [x] Group
  - [x] Group
  - [x] SubGroup

- [x] Map
- [ ] Sort
  - [x] Sort
  - [ ] SortRev

- [ ] Take

### Supported operations for DataFrame

- [x] Filter
- [x] GroupBy
- [ ] Join
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
```