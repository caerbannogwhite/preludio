package gandalff

import (
	"fmt"
	"typesys"
)

type Series interface {

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
	IsSorted() SeriesSortOrder

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
	SetNull(i int) Series
	// Returns the null mask of the series.
	GetNullMask() []bool
	// Sets the null mask of the series.
	SetNullMask(mask []bool) Series
	// Makes the series nullable.
	MakeNullable() Series

	// Get the element at index i.
	Get(i int) any
	// Get the element at index i as a string.
	GetString(i int) string
	// Set the element at index i.
	Set(i int, v any) Series
	// Take the elements according to the given interval.
	Take(start, end, step int) Series

	// Sort Interface.
	Less(i, j int) bool
	Swap(i, j int)

	// Append elements to the series.
	Append(v any) Series
	// AppendRaw appends a value or a slice of values to the series.
	AppendRaw(v any) Series
	// Append nullable elements to the series.
	AppendNullable(v any) Series
	// Append a series to the series.
	AppendSeries(other Series) Series

	// All-data accessors.

	// Returns the actual data of the series.
	Data() any
	// Returns the nullable data of the series.
	DataAsNullable() any
	// Returns the data of the series as a slice of strings.
	DataAsString() []string

	// Casts the series to a given type.
	Cast(t typesys.BaseType, stringPool *StringPool) Series
	// Copies the series.
	Copy() Series

	// Series operations.

	// Filters out the elements by the given mask series.
	Filter(mask SeriesBool) Series
	// Filters out the elements by the given mask.
	FilterByMask(mask []bool) Series
	// Filters out the elements by the given indices.
	FilterByIndeces(indices []int) Series

	// Maps the elements of the series.
	Map(f GDLMapFunc, stringPool *StringPool) Series

	// Group the elements in the series.
	Group() Series
	SubGroup(gp SeriesPartition) Series

	// Get the partition of the series.
	GetPartition() SeriesPartition

	// Sorts the elements of the series.
	Sort() Series
	SortRev() Series
}

func NewSeries(name string, t typesys.BaseType, nullable bool, makeCopy bool, data any, pool *StringPool) Series {
	switch t {
	case typesys.BoolType:
		return NewSeriesBool(name, nullable, data.([]bool))
	case typesys.Int32Type:
		return NewSeriesInt32(name, nullable, makeCopy, data.([]int32))
	case typesys.Int64Type:
		return NewSeriesInt64(name, nullable, makeCopy, data.([]int64))
	case typesys.Float64Type:
		return NewSeriesFloat64(name, nullable, makeCopy, data.([]float64))
	case typesys.StringType:
		return NewSeriesString(name, nullable, data.([]string), pool)
	default:
		return SeriesError{fmt.Sprintf("NewSeries: unknown type: %v", t)}
	}
}

type SeriesPartition interface {
	// Returns the number partitions.
	GetSize() int
	// Returns the indices of the groups.
	GetMap() map[int64][]int
	// Returns the indices for a given value. The value must be of the same type as the series.
	// If val is nil then the indices of the null values are returned.
	GetValueIndices(val any) []int
	// Returns  the keys of the groups.
	GetKeys() any

	debugPrint()
}
