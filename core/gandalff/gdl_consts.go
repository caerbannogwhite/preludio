package gandalff

import "strconv"

const (
	// The default capacity of a series.
	DEFAULT_SERIES_INITIAL_CAPACITY = 10
	// The default capacity of a hash map.
	DEFAULT_HASH_MAP_INITIAL_CAPACITY = 1024
	// Number of threads to use for parallel operations.
	THREADS_NUMBER = 8
	// Minimum number of elements to use parallel operations.
	MINIMUM_PARALLEL_SIZE = 16_384

	// HASH_MAGIC_NUMBER = 0x9e3779b9
	HASH_MAGIC_NUMBER = 0xa8f4979b77e3f93f
)

////////////////////////////////			ENUMS

type GDLSeriesSortOrder int16

const (
	// The series is not sorted.
	SORTED_NONE GDLSeriesSortOrder = iota
	// The series is sorted in ascending order.
	SORTED_ASC
	// The series is sorted in descending order.
	SORTED_DESC
)

////////////////////////////////			ERRORS

////////////////////////////////			TO STRING

const NULL_STRING = "NA"
const BOOL_TRUE_STRING = "true"
const BOOL_FALSE_STRING = "false"

func intToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

type any interface{}

////////////////////////////////			NULLABLE TYPES

type NullableBool struct {
	Valid bool
	Value bool
}

type NullableInt16 struct {
	Valid bool
	Value int16
}

type NullableInt32 struct {
	Valid bool
	Value int
}

type NullableInt64 struct {
	Valid bool
	Value int64
}

type NullableFloat32 struct {
	Valid bool
	Value float32
}

type NullableFloat64 struct {
	Valid bool
	Value float64
}

type NullableString struct {
	Valid bool
	Value string
}

type GDLMapFunc func(any) any
