## GANDALFF: Golang, ANother DAta Library For Fun

Or, for short, GDL: Golang Data Library

### What is it?

Gandalff is a library for data manipulation in Go.
It supports nullable types: null data is optimized for memory usage.

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

	NewBaseDataFrame(NewContext()).
		FromCSV().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		Read().
		Select("department", "age", "weight", "junior").
		GroupBy("department").
		Agg(Min("age"), Max("weight"), Mean("junior"), Count()).
		PrettyPrint(NewPrettyPrintParams())

	// Output:
	// ╭────────────┬────────────┬────────────┬────────────┬────────────╮
	// │ department │        age │     weight │     junior │          n │
	// ├────────────┼────────────┼────────────┼────────────┼────────────┤
	// │     String │    Float64 │    Float64 │    Float64 │      Int64 │
	// ├────────────┼────────────┼────────────┼────────────┼────────────┤
	// │         HR │         29 │         90 │        0.5 │          2 │
	// │         IT │         25 │         85 │        0.5 │          4 │
	// │   Business │         27 │         65 │        0.5 │          2 │
	// ╰────────────┴────────────┴────────────┴────────────┴────────────╯
}
```

### Community

You can join the [Gandalff community on Discord](https://discord.gg/vPv5bhXY).

### Supported data types

The data types not checked are not yet supported, but might be in the future.

- [x] Bool
- [ ] Bool (memory optimized, not fully implemented yet)
- [ ] Int16
- [x] Int
- [x] Int64
- [ ] Float32
- [x] Float64
- [ ] Complex64
- [ ] Complex128
- [x] String
- [x] Time
- [x] Duration

### Supported operations for Series

- [x] Filter

  - [x] filter by bool slice
  - [x] filter by int slice
  - [x] filter by bool series
  - [x] filter by int series

- [x] Group

  - [x] Group (with nulls)
  - [x] SubGroup (with nulls)

- [x] Map
- [x] Sort

  - [x] Sort (with nulls)
  - [x] SortRev (with nulls)

- [x] Take

### Supported operations for DataFrame

- [x] Agg
- [x] Filter
- [x] GroupBy
- [ ] Join

  - [x] Inner
  - [x] Left
  - [x] Right
  - [x] Outer
  - [ ] Inner with nulls
  - [ ] Left with nulls
  - [ ] Right with nulls
  - [ ] Outer with nulls

- [ ] Map
- [x] OrderBy
- [x] Select
- [x] Take

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

This is how the interface for the Series type currently looks like:

```go
// Basic accessors.

// Return the context of the series.
GetContext() *Context
// Return the number of elements in the series.
Len() int
// Return the type of the series.
Type() typesys.BaseType
// Return the type and cardinality of the series.
TypeCard() typesys.BaseTypeCard
// Return if the series is grouped.
IsGrouped() bool
// Return if the series admits null values.
IsNullable() bool
// Return if the series is sorted.
IsSorted() SeriesSortOrder
// Return if the series is error.
IsError() bool
// Return the error message of the series.
GetError() string

// Nullability operations.

// Return if the series has null values.
HasNull() bool
// Return the number of null values in the series.
NullCount() int
// Return if the element at index i is null.
IsNull(i int) bool
// Return the null mask of the series.
GetNullMask() []bool
// Set the null mask of the series.
SetNullMask(mask []bool) Series
// Make the series nullable.
MakeNullable() Series
// Make the series non-nullable.
MakeNonNullable() Series

// Get the element at index i.
Get(i int) any
// Get the element at index i as a string.
GetAsString(i int) string
// Set the element at index i.
Set(i int, v any) Series
// Take the elements according to the given interval.
Take(params ...int) Series

// Append elements to the series.
// Value can be a single value, slice of values,
// a nullable value, a slice of nullable values or a series.
Append(v any) Series

// All-data accessors.

// Return the actual data of the series.
Data() any
// Return the nullable data of the series.
DataAsNullable() any
// Return the data of the series as a slice of strings.
DataAsString() []string

// Cast the series to a given type.
Cast(t typesys.BaseType) Series
// Copie the series.
Copy() Series

// Series operations.

// Filter out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
Filter(mask any) Series

// Apply the given function to each element of the series.
Map(f MapFunc) Series
MapNull(f MapFuncNull) Series

// Group the elements in the series.
GroupBy(gp SeriesPartition) Series
UnGroup() Series

// Get the partition of the series.
GetPartition() SeriesPartition

// Sort Interface.
Less(i, j int) bool
Swap(i, j int)

// Sort the elements of the series.
Sort() Series
SortRev() Series

// Arithmetic operations.
Mul(other Series) Series
Div(other Series) Series
Mod(other Series) Series
Exp(other Series) Series
Add(other Series) Series
Sub(other Series) Series

// Logical operations.
Eq(other Series) Series
Ne(other Series) Series
Gt(other Series) Series
Ge(other Series) Series
Lt(other Series) Series
Le(other Series) Series
```

### TODO

- [ ] Improve dataframe PrettyPrint: add parameters, optimize data display, use lipgloss.
- [ ] Implement string factors.
- [ ] SeriesTime: set time format.
- [ ] Implement `Set(i []int, v []any) Series`.
- [ ] Add `Slice(i []int) Series` (using filter?).
- [ ] Implement memory optimized Bool series with uint64.
- [ ] Use uint64 for null mask.
- [ ] Implement chunked series.
- [ ] Implement Excel reader and writer (https://github.com/tealeg/xlsx).
- [ ] Implement JSON reader and writer.
