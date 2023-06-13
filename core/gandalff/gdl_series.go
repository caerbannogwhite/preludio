package gandalff

import (
	"fmt"
	"sync"
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
	// case typesys.Int16Type:
	// 	return NewSeriesInt16(name, nullable, data.([]int16))
	case typesys.Int32Type:
		return NewSeriesInt32(name, nullable, makeCopy, data.([]int))
	// case typesys.Int64Type:
	// 	return NewSeriesInt64(name, nullable, data.([]int64))
	// case typesys.Float32Type:
	// 	return NewSeriesFloat32(name, nullable, data.([]float32))
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
}

type StringPool struct {
	sync.RWMutex
	pool map[string]*string
}

func NewStringPool() *StringPool {
	return &StringPool{pool: make(map[string]*string)}
}

// Get returns the address of the string if it exists in the pool, otherwise nil.
func (sp *StringPool) Get(s string) *string {
	if entry, ok := sp.pool[s]; ok {
		return entry
	}
	return nil
}

// Put returns the address of the string if it exists in the pool, otherwise it adds it to the pool and returns its address.
func (sp *StringPool) Put(s string) *string {
	if entry, ok := sp.pool[s]; ok {
		return entry
	}

	// Create a new string and add it to the pool
	addr := &s
	sp.pool[s] = addr
	return addr
}

// PutSync returns the address of the string if it exists in the pool, otherwise it adds it to the pool and returns its address.
// This version is thread-safe.
func (sp *StringPool) PutSync(s string) *string {
	sp.RLock()
	entry, ok := sp.pool[s]
	sp.RUnlock()
	if ok {
		return entry
	}

	sp.Lock()
	defer sp.Unlock()
	if entry, ok := sp.pool[s]; ok {
		// Someone else inserted the string while we were waiting
		return entry
	}

	// Create a new string and add it to the pool
	sp.pool[s] = &s
	return sp.pool[s]
}

// Dummy series for error handling.
type SeriesError struct {
	msg string
}

func (s SeriesError) Error() string {
	return s.msg
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

// Makes the series nullable.
func (s SeriesError) MakeNullable() Series {
	return s
}

// Returns the name of the series.
func (s SeriesError) Name() string {
	return ""
}

// Returns the type of the series.
func (s SeriesError) Type() typesys.BaseType {
	return typesys.ErrorType
}

// Returns if the series has null values.
func (s SeriesError) HasNull() bool {
	return false
}

// Returns the number of null values in the series.
func (s SeriesError) NullCount() int {
	return 0
}

// Returns the number of non-null values in the series.
func (s SeriesError) NonNullCount() int {
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
func (s SeriesError) Take(start, end, step int) Series {
	return s
}

// Sort interface.
func (s SeriesError) Less(i, j int) bool {
	return false
}

func (s SeriesError) Swap(i, j int) {}

// Append elements to the series.
func (s SeriesError) Append(v any) Series {
	return s
}

// AppendRaw appends a value or a slice of values to the series.
func (s SeriesError) AppendRaw(v any) Series {
	return s
}

// Append nullable elements to the series.
func (s SeriesError) AppendNullable(v any) Series {
	return s
}

// Append a series to the series.
func (s SeriesError) AppendSeries(other Series) Series {
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
func (s SeriesError) Cast(t typesys.BaseType, stringPool *StringPool) Series {
	return s
}

// Copies the series.
func (s SeriesError) Copy() Series {
	return s
}

// Series operations.

// Filters out the elements by the given mask series.
func (s SeriesError) Filter(mask SeriesBool) Series {
	return s
}

// FilterByMask returns a new series with elements filtered by the mask.
func (s SeriesError) FilterByMask(mask []bool) Series {
	return s
}

// FilterByIndeces returns a new series with elements filtered by the indeces.
func (s SeriesError) FilterByIndeces(indeces []int) Series {
	return s
}

func (s SeriesError) Map(f GDLMapFunc, stringPool *StringPool) Series {
	return s
}

// Group the elements in the series.
func (s SeriesError) Group() Series {
	return nil
}

func (s SeriesError) SubGroup(gp SeriesPartition) Series {
	return nil
}

func (s SeriesError) GetPartition() SeriesPartition {
	return nil
}

func (s SeriesError) Sort() Series {
	return s
}

func (s SeriesError) SortRev() Series {
	return s
}
