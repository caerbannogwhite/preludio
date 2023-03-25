package gandalff

import (
	"strconv"
	"sync"
	"typesys"
)

///////////////////////////////		TO STRING		/////////////////////////////////

const NULL_STRING = "NA"
const BOOL_TRUE_STRING = "true"
const BOOL_FALSE_STRING = "false"

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}

func intToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

///////////////////////////////		NULLABLE TYPES		/////////////////////////////////

type NullableBool struct {
	Valid bool
	Value bool
}

type NullableInt struct {
	Valid bool
	Value int
}

type NullableFloat struct {
	Valid bool
	Value float64
}

type NullableString struct {
	Valid bool
	Value string
}

///////////////////////////////		GDLSERIES		/////////////////////////////////

type GDLSeries interface {

	// Basic accessors.

	// Returns the length of the series.
	Len() int
	// Returns if the series admits null values.
	IsNullable() bool
	// Makes the series nullable.
	MakeNullable()
	// Returns the name of the series.
	Name() string
	// Returns the type of the series.
	Type() typesys.BaseType
	// Returns if the series has null values.
	HasNull() bool
	// Returns the number of null values in the series.
	NullCount() int
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

	// Append appends a value or a slice of values to the series.
	Append(v interface{}) error
	// Append nullable elements to the series.
	AppendNullable(v interface{}) error

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
	Filter(mask []bool) GDLSeries

	// Group the elements in the series.
	Group() GDLSeriesPartition
	SubGroup(gp GDLSeriesPartition) GDLSeriesPartition
}

type GDLSeriesPartition interface {
	// Returns the number of groups.
	GetGroupsCount() int
	// Returns the non-null groups.
	GetNonNullGroups() [][]int
	// Returns the null group.
	GetNullGroup() []int
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
