package gandalff

import (
	"strconv"
	"sync"
)

type GSeriesType uint8

const (
	BoolType GSeriesType = iota
	IntType
	FloatType
	StringType
	TimeType
	DurationType
	ComplexType
	InterfaceType
)

func (t GSeriesType) ToString() string {
	switch t {
	case BoolType:
		return "bool"
	case IntType:
		return "int"
	case FloatType:
		return "float"
	case StringType:
		return "string"
	case TimeType:
		return "time"
	case DurationType:
		return "duration"
	case ComplexType:
		return "complex"
	case InterfaceType:
		return "interface"
	}
	return "unknown"
}

///////////////////////////////		TO STRING		/////////////////////////////////

const NULL_STRING = "NA"

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}

func intToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func timeToString(t int64) string {
	return strconv.FormatInt(t, 10)
}

// func durationToString(d int64) string {
// 	return ""
// }

func complexToString(c complex128) string {
	return strconv.FormatComplex(c, 'f', -1, 128)
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

type NullableTime struct {
	Valid bool
	Value int64
}

type NullableDuration struct {
	Valid bool
	Value int64
}

type NullableComplex struct {
	Valid bool
	Value complex128
}

type NullableInterface struct {
	Valid bool
	Value interface{}
}

///////////////////////////////		GSERIES		/////////////////////////////////

type GSeries interface {
	// Returns the length of the series.
	Len() int
	// Returns if the series admits null values.
	IsNullable() bool
	// Returns if the series has null values.
	HasNull() bool
	// Returns if the element at index i is null.
	IsNull(i int) bool
	// Sets the element at index i to null.
	SetNull(i int)
	// Returns the name of the series.
	Name() string
	// Returns the type of the series.
	Type() GSeriesType
	// Returns the actual data of the series.
	Data() interface{}
	// Returns the null mask of the series.
	NullMask() []bool
	// Sets the null mask of the series.
	SetNullMask(mask []bool)
	// Returns the nullable data of the series.
	NullableData() interface{}
	// Returns the data of the series as a slice of strings.
	StringData() []string
	// Copies the series.
	Copy() GSeries
	// Filters out the elements by the given mask.
	Filter(mask []bool) GSeries
	// Filters out the elements by the given mask in place.
	FilterInPlace(mask []bool)
	// Filters out the elements by the given indeces.
	FilterByIndex(indexes []int) GSeries
	// Filters out the elements by the given indeces in place.
	FilterByIndexInPlace(indexes []int)
}

type StringPool struct {
	sync.RWMutex
	pool map[string]*string
}

func NewStringPool() *StringPool {
	return &StringPool{pool: make(map[string]*string)}
}

func (sp *StringPool) Get(s string) *string {
	sp.RLock()
	strPtr, ok := sp.pool[s]
	sp.RUnlock()
	if ok {
		return strPtr
	}

	sp.Lock()
	defer sp.Unlock()
	if strPtr, ok := sp.pool[s]; ok {
		// Someone else inserted the string while we were waiting
		return strPtr
	}

	// Create a new string and add it to the pool
	strPtr = &s
	sp.pool[s] = strPtr
	return strPtr
}

// GSeriesString represents a series of strings.
// type GSeriesString struct {
// 	isNullable bool
// 	name       string
// 	data       []*string
// 	nullMap    []uint8
// 	pool       *StringPool
// }

// func NewGSeriesString(name string, isNullable bool, data []string, pool *StringPool) GSeriesString {
// 	var nullMap []uint8
// 	if isNullable {
// 		nullMap = make([]uint8, len(data)/8+1)
// 	} else {
// 		nullMap = make([]uint8, 0)
// 	}

// 	actualData := make([]*string, len(data))
// 	for i, v := range data {
// 		actualData[i] = pool.Get(v)
// 	}

// 	return GSeriesString{isNullable: isNullable, name: name, data: actualData, nullMap: nullMap, pool: pool}
// }

// func (s GSeriesString) Len() int {
// 	return len(s.data)
// }

// func (s GSeriesString) IsNullable() bool {
// 	return s.isNullable
// }

// func (s GSeriesString) HasNull() bool {
// 	for _, v := range s.nullMap {
// 		if v != 0 {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (s GSeriesString) IsNull(i int) bool {
// 	if s.isNullable {
// 		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
// 	}
// 	return false
// }

// func (s GSeriesString) SetNull(i int) {
// 	if s.isNullable {
// 		s.nullMap[i/8] |= 1 << uint(i%8)
// 	}
// }

// func (s GSeriesString) Name() string {
// 	return s.name
// }

// func (s GSeriesString) Type() GSeriesType {
// 	return StringType
// }

// func (s GSeriesString) Data() interface{} {
// 	data := make([]string, len(s.data))
// 	for i, v := range s.data {
// 		data[i] = *v
// 	}
// 	return data
// }

// func (s GSeriesString) NullMask() []bool {
// 	mask := make([]bool, len(s.data))
// 	for k, v := range s.nullMap {
// 		for i := 0; i < 8; i++ {
// 			mask[k*8+i] = v&(1<<uint(i)) != 0
// 		}
// 	}
// 	return mask
// }

// func (s GSeriesString) SetNullMask(mask []bool) {
// 	for k, v := range mask {
// 		if v {
// 			s.nullMap[k/8] |= 1 << uint(k%8)
// 		} else {
// 			s.nullMap[k/8] &= ^(1 << uint(k%8))
// 		}
// 	}
// }

// func (s GSeriesString) NullableData() interface{} {
// 	data := make([]NullableString, len(s.data))
// 	for i, v := range s.data {
// 		data[i] = NullableString{Valid: !s.IsNull(i), Value: *v}
// 	}
// 	return data
// }

// func (s GSeriesString) StringData() []string {
// 	data := make([]string, len(s.data))
// 	for i, v := range s.data {
// 		if s.IsNull(i) {
// 			data[i] = NULL_STRING
// 		} else {
// 			data[i] = *v
// 		}
// 	}
// 	return data
// }
