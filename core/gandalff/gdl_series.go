package gandalff

import (
	"preludiometa"
)

type Series interface {
	// Utility functions.
	printInfo()

	// Basic accessors.

	// Return the context of the series.
	GetContext() *Context
	// Return the number of elements in the series.
	Len() int
	// Return the type of the series.
	Type() preludiometa.BaseType
	// Return the type and cardinality of the series.
	TypeCard() preludiometa.BaseTypeCard
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
	Cast(t preludiometa.BaseType) Series
	// Copie the series.
	Copy() Series

	// Series operations.

	// Filter out the elements by the given mask.
	// Mask can be a bool series, a slice of bools or a slice of ints.
	Filter(mask any) Series
	filterIntSlice(mask []int, check bool) Series

	// Apply the given function to each element of the series.
	Map(f MapFunc) Series
	MapNull(f MapFuncNull) Series

	// Group the elements in the series.
	group() Series
	GroupBy(gp SeriesPartition) Series
	UnGroup() Series

	// Get the partition of the series.
	GetPartition() SeriesPartition

	// Sort Interface.
	Less(i, j int) bool
	equal(i, j int) bool
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
}

type SeriesPartition interface {
	// Return the number partitions.
	getSize() int

	// Return the indices of the groups.
	getMap() map[int64][]int
}
