package gandalff

import (
	"fmt"
	"strconv"
	"sync"
	"typesys"
)

///////////////////////////////		TO STRING		/////////////////////////////////

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

///////////////////////////////		NULLABLE TYPES		/////////////////////////////////

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

///////////////////////////////		GDLSERIES		/////////////////////////////////

type GDLMapFunc func(interface{}) interface{}

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
	IsSorted() bool

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
	// Set the element at index i.
	Set(i int, v any) GDLSeries

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
	NullableData() any
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
	Group() GDLSeries
	SubGroup(gp GDLSeriesPartition) GDLSeries

	// Get the partition of the series.
	GetPartition() GDLSeriesPartition

	// Sorts the elements of the series.
	Sort() GDLSeries
	SortRev() GDLSeries
}

func NewGDLSeries(name string, t typesys.BaseType, nullable bool, makeCopy bool, data any, pool *StringPool) GDLSeries {
	switch t {
	case typesys.BoolType:
		return NewGDLSeriesBool(name, nullable, data.([]bool))
	// case typesys.Int16Type:
	// 	return NewGDLSeriesInt16(name, nullable, data.([]int16))
	case typesys.Int32Type:
		return NewGDLSeriesInt32(name, nullable, makeCopy, data.([]int))
	// case typesys.Int64Type:
	// 	return NewGDLSeriesInt64(name, nullable, data.([]int64))
	// case typesys.Float32Type:
	// 	return NewGDLSeriesFloat32(name, nullable, data.([]float32))
	case typesys.Float64Type:
		return NewGDLSeriesFloat64(name, nullable, makeCopy, data.([]float64))
	case typesys.StringType:
		return NewGDLSeriesString(name, nullable, data.([]string), pool)
	default:
		return GDLSeriesError{fmt.Sprintf("NewGDLSeries: unknown type: %v", t)}
	}
}

type GDLSeriesPartition interface {
	// Returns the number partitions.
	GetSize() int
	// Returns the number of groups.
	GetGroupsCount() int
	// Returns the indices of the groups.
	GetIndices() [][]int
	// Returns the indices for a given value in a given sub-group.
	GetValueIndices(sub int, val interface{}) []int
	// Returns the null group.
	GetNullIndices(sub int) []int
}

type StringPoolEntry struct {
	Addr  *string
	Count uint32
}

type StringPool struct {
	sync.RWMutex
	pool map[string]StringPoolEntry
}

func NewStringPool() *StringPool {
	return &StringPool{pool: make(map[string]StringPoolEntry)}
}

// Add adds the string to the pool if it doesn't exist, otherwise it increments the reference count.
func (sp *StringPool) Add(s string) *string {
	sp.Lock()
	defer sp.Unlock()
	if entry, ok := sp.pool[s]; ok {
		entry.Count++ // Increment the reference count
		return entry.Addr
	}

	// Create a new string and add it to the pool
	strPtr := &s
	sp.pool[s] = StringPoolEntry{Addr: strPtr, Count: 1}
	return strPtr
}

// Remove removes the string from the pool if it exists and decrements the reference count.
func (sp *StringPool) Remove(s string) {
	sp.Lock()
	defer sp.Unlock()
	if entry, ok := sp.pool[s]; ok {
		entry.Count-- // Decrement the reference count
		if entry.Count == 0 {
			delete(sp.pool, s)
		}
	}
}

// Get returns the address of the string if it exists in the pool, otherwise nil.
func (sp *StringPool) Get(s string) *string {
	sp.RLock()
	entry, ok := sp.pool[s]
	sp.RUnlock()
	if ok {
		return entry.Addr
	}

	sp.Lock()
	defer sp.Unlock()
	if entry, ok := sp.pool[s]; ok {
		// Someone else inserted the string while we were waiting
		return entry.Addr
	}

	// Create a new string and add it to the pool
	sp.pool[s] = StringPoolEntry{Addr: &s, Count: 1}
	return sp.pool[s].Addr
}

// Dummy series for error handling.
type GDLSeriesError struct {
	msg string
}

func (s GDLSeriesError) Error() string {
	return s.msg
}

// Returns the length of the series.
func (s GDLSeriesError) Len() int {
	return 0
}

// Returns if the series is grouped.
func (s GDLSeriesError) IsGrouped() bool {
	return false
}

// Returns if the series admits null values.
func (s GDLSeriesError) IsNullable() bool {
	return false
}

func (s GDLSeriesError) IsSorted() bool {
	return false
}

// Makes the series nullable.
func (s GDLSeriesError) MakeNullable() GDLSeries {
	return s
}

// Returns the name of the series.
func (s GDLSeriesError) Name() string {
	return ""
}

// Returns the type of the series.
func (s GDLSeriesError) Type() typesys.BaseType {
	return typesys.ErrorType
}

// Returns if the series has null values.
func (s GDLSeriesError) HasNull() bool {
	return false
}

// Returns the number of null values in the series.
func (s GDLSeriesError) NullCount() int {
	return 0
}

// Returns the number of non-null values in the series.
func (s GDLSeriesError) NonNullCount() int {
	return 0
}

// Returns if the element at index i is null.
func (s GDLSeriesError) IsNull(i int) bool {
	return false
}

// Sets the element at index i to null.
func (s GDLSeriesError) SetNull(i int) GDLSeries {
	return nil
}

// Returns the null mask of the series.
func (s GDLSeriesError) GetNullMask() []bool {
	return []bool{}
}

// Sets the null mask of the series.
func (s GDLSeriesError) SetNullMask(mask []bool) GDLSeries {
	return nil
}

// Get the element at index i.
func (s GDLSeriesError) Get(i int) any {
	return nil
}

// Set the element at index i.
func (s GDLSeriesError) Set(i int, v any) GDLSeries {
	return s
}

// Sort interface.
func (s GDLSeriesError) Less(i, j int) bool {
	return false
}

func (s GDLSeriesError) Swap(i, j int) {}

// Append elements to the series.
func (s GDLSeriesError) Append(v any) GDLSeries {
	return s
}

// AppendRaw appends a value or a slice of values to the series.
func (s GDLSeriesError) AppendRaw(v any) GDLSeries {
	return s
}

// Append nullable elements to the series.
func (s GDLSeriesError) AppendNullable(v any) GDLSeries {
	return s
}

// Append a series to the series.
func (s GDLSeriesError) AppendSeries(other GDLSeries) GDLSeries {
	return s
}

// All-data accessors.

// Returns the actual data of the series.
func (s GDLSeriesError) Data() any {
	return nil
}

// Returns the nullable data of the series.
func (s GDLSeriesError) NullableData() any {
	return nil
}

// Returns the data of the series as a slice of strings.
func (s GDLSeriesError) StringData() []string {
	return []string{}
}

// Copies the series.
func (s GDLSeriesError) Copy() GDLSeries {
	return s
}

// Series operations.

// FilterByMask returns a new series with elements filtered by the mask.
func (s GDLSeriesError) FilterByMask(mask []bool) GDLSeries {
	return s
}

// FilterByIndeces returns a new series with elements filtered by the indeces.
func (s GDLSeriesError) FilterByIndeces(indeces []int) GDLSeries {
	return s
}

func (s GDLSeriesError) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	return s
}

// Group the elements in the series.
func (s GDLSeriesError) Group() GDLSeries {
	return nil
}

func (s GDLSeriesError) SubGroup(gp GDLSeriesPartition) GDLSeries {
	return nil
}

func (s GDLSeriesError) GetPartition() GDLSeriesPartition {
	return nil
}

func (s GDLSeriesError) Sort() GDLSeries {
	return s
}

func (s GDLSeriesError) SortRev() GDLSeries {
	return s
}
