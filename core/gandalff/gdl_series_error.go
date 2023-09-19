package gandalff

import "typesys"

// Dummy series for error handling.
type SeriesError struct {
	msg string
}

// Returns the length of the series.
func (s SeriesError) Len() int {
	return 0
}

// Returns if the series is grouped.
func (s SeriesError) IsGrouped() bool {
	return false
}

// Returns if the series admits null values.
func (s SeriesError) IsNullable() bool {
	return false
}

func (s SeriesError) IsSorted() SeriesSortOrder {
	return SORTED_NONE
}

// Returns if the series is error.
func (s SeriesError) IsError() bool {
	return true
}

// Returns the error message of the series.
func (s SeriesError) GetError() string {
	return s.msg
}

// Makes the series nullable.
func (s SeriesError) MakeNullable() Series {
	return s
}

// Returns the name of the series.
func (s SeriesError) Name() string {
	return ""
}

// Sets the name of the series.
func (s SeriesError) SetName(name string) Series {
	return s
}

// Return the StringPool of the series.
func (s SeriesError) StringPool() *StringPool {
	return nil
}

// Set the StringPool for this series.
func (s SeriesError) SetStringPool(pool *StringPool) Series {
	return s
}

// Returns the type of the series.
func (s SeriesError) Type() typesys.BaseType {
	return typesys.ErrorType
}

// Returns the type and cardinality of the series.
func (s SeriesError) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{Base: typesys.ErrorType, Card: s.Len()}
}

// Returns if the series has null values.
func (s SeriesError) HasNull() bool {
	return false
}

// Returns the number of null values in the series.
func (s SeriesError) NullCount() int {
	return 0
}

// Returns if the element at index i is null.
func (s SeriesError) IsNull(i int) bool {
	return false
}

// Sets the element at index i to null.
func (s SeriesError) SetNull(i int) Series {
	return nil
}

// Returns the null mask of the series.
func (s SeriesError) GetNullMask() []bool {
	return []bool{}
}

// Sets the null mask of the series.
func (s SeriesError) SetNullMask(mask []bool) Series {
	return nil
}

// Get the element at index i.
func (s SeriesError) Get(i int) any {
	return nil
}

func (s SeriesError) GetString(i int) string {
	return ""
}

// Set the element at index i.
func (s SeriesError) Set(i int, v any) Series {
	return s
}

// Take the elements according to the given interval.
func (s SeriesError) Take(params ...int) Series {
	return s
}

// Append elements to the series.
func (s SeriesError) Append(v any) Series {
	return s
}

// All-data accessors.

// Returns the actual data of the series.
func (s SeriesError) Data() any {
	return nil
}

// Returns the nullable data of the series.
func (s SeriesError) DataAsNullable() any {
	return nil
}

// Returns the data of the series as a slice of strings.
func (s SeriesError) DataAsString() []string {
	return []string{}
}

// Casts the series to a given type.
func (s SeriesError) Cast(t typesys.BaseType) Series {
	return s
}

// Copies the series.
func (s SeriesError) Copy() Series {
	return s
}

// Series operations.

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s SeriesError) Filter(mask any) Series {
	return s
}

func (s SeriesError) filterIntSlice(mask []int, check bool) Series {
	return s
}

func (s SeriesError) Map(f GDLMapFunc) Series {
	return s
}

// Group the elements in the series.
func (s SeriesError) group() Series {
	return nil
}

func (s SeriesError) GroupBy(gp SeriesPartition) Series {
	return nil
}

func (s SeriesError) UnGroup() Series {
	return s
}

func (s SeriesError) GetPartition() SeriesPartition {
	return nil
}

// Sort interface.
func (s SeriesError) Less(i, j int) bool {
	return false
}

func (s SeriesError) equal(i, j int) bool {
	return false
}

func (s SeriesError) Swap(i, j int) {}

func (s SeriesError) Sort() Series {
	return s
}

func (s SeriesError) SortRev() Series {
	return s
}

////////////////////////			ARITHMETIC OPERATIONS

func (s SeriesError) Mul(other Series) Series {
	return s
}

func (s SeriesError) Div(other Series) Series {
	return s
}

func (s SeriesError) Mod(other Series) Series {
	return s
}

func (s SeriesError) Pow(other Series) Series {
	return s
}

func (s SeriesError) Add(other Series) Series {
	return s
}

func (s SeriesError) Sub(other Series) Series {
	return s
}

////////////////////////			LOGICAL OPERATIONS

func (s SeriesError) Eq(other Series) Series {
	return s
}

func (s SeriesError) Ne(other Series) Series {
	return s
}

func (s SeriesError) Gt(other Series) Series {
	return s
}

func (s SeriesError) Ge(other Series) Series {
	return s
}

func (s SeriesError) Lt(other Series) Series {
	return s
}

func (s SeriesError) Le(other Series) Series {
	return s
}
