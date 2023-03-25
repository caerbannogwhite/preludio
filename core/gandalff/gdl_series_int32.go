package gandalff

import (
	"errors"
	"typesys"
)

// GDLSeriesInt32 represents a series of ints.
type GDLSeriesInt32 struct {
	isNullable bool
	name       string
	data       []int
	nullMap    []uint8
}

func NewGDLSeriesInt32(name string, isNullable bool, makeCopy bool, data []int) GDLSeriesInt32 {
	var nullMap []uint8
	if isNullable {
		if len(data)%8 == 0 {
			nullMap = make([]uint8, len(data)/8)
		} else {
			nullMap = make([]uint8, len(data)/8+1)
		}
	} else {
		nullMap = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GDLSeriesInt32{isNullable: isNullable, name: name, data: data, nullMap: nullMap}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////

func (s GDLSeriesInt32) Len() int {
	return len(s.data)
}

func (s GDLSeriesInt32) IsNullable() bool {
	return s.isNullable
}

func (s GDLSeriesInt32) MakeNullable() {
	if !s.isNullable {
		if len(s.data)%8 == 0 {
			s.nullMap = make([]uint8, len(s.data)/8)
		} else {
			s.nullMap = make([]uint8, len(s.data)/8+1)
		}
		s.isNullable = true
	}
}

func (s GDLSeriesInt32) Name() string {
	return s.name
}

func (s GDLSeriesInt32) Type() typesys.BaseType {
	return typesys.Int32Type
}

func (s GDLSeriesInt32) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesInt32) NullCount() int {
	count := 0
	for _, v := range s.nullMap {
		for i := 0; i < 8; i++ {
			if v&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s GDLSeriesInt32) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesInt32) SetNull(i int) error {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
		return nil
	}
	return errors.New("GDLSeriesInt32.SetNull: series is not nullable")
}

func (s GDLSeriesInt32) GetNullMask() []bool {
	mask := make([]bool, len(s.data))
	idx := 0
	for _, v := range s.nullMap {
		for i := 0; i < 8 && idx < len(s.data); i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

func (s GDLSeriesInt32) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return errors.New("GDLSeriesInt32.SetNullMask: series is not nullable")
	}

	for k, v := range mask {
		if v {
			s.nullMap[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMap[k/8] &= ^(1 << uint(k%8))
		}
	}

	return nil
}

func (s GDLSeriesInt32) Get(i int) interface{} {
	return s.data[i]
}

func (s GDLSeriesInt32) Set(i int, v interface{}) {
	s.data[i] = v.(int)
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesInt32) Append(v interface{}) error {
	if s.isNullable {
		if b, ok := v.(int); ok {
			s.data = append(s.data, b)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
		} else if bv, ok := v.([]int); ok {
			s.data = append(s.data, bv...)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, make([]uint8, len(s.data)/8-len(s.nullMap))...)
			}
		} else {
			return errors.New("GDLSeriesInt32.Append: invalid type")
		}
	} else {
		if b, ok := v.(int); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]int); ok {
			s.data = append(s.data, bv...)
		} else {
			return errors.New("GDLSeriesInt32.Append: invalid type")
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesInt32) AppendNullable(v interface{}) error {
	if !s.isNullable {
		return errors.New("GDLSeriesInt32.AppendNullable: series is not nullable")
	}

	if b, ok := v.(NullableInt); ok {
		s.data = append(s.data, b.Value)
		if len(s.data)/8 > len(s.nullMap) {
			s.nullMap = append(s.nullMap, 0)
		}
		if !b.Valid {
			s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableInt); ok {
		for _, b := range bv {
			s.data = append(s.data, b.Value)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
			if !b.Valid {
				s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return errors.New("GDLSeriesInt32.AppendNullable: invalid type")
	}

	return nil
}

///////////////////////////////  	ALL DATA ACCESSORS  /////////////////////////////////

func (s GDLSeriesInt32) Data() interface{} {
	return s.data
}

func (s GDLSeriesInt32) NullableData() interface{} {
	data := make([]NullableInt, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GDLSeriesInt32) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		if s.IsNull(i) {
			data[i] = NULL_STRING
		} else {
			data[i] = intToString(v)
		}
	}
	return data
}

func (s GDLSeriesInt32) Copy() GSeries {
	data := make([]int, len(s.data))
	copy(data, s.data)
	nullMap := make([]uint8, len(s.nullMap))
	copy(nullMap, s.nullMap)

	return GDLSeriesInt32{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

///////////////////////////////  	SERIES OPERATIONS  //////////////////////////////////

func (s GDLSeriesInt32) Filter(filter []bool) GSeries {
	data := make([]int, 0)
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range filter {
		if v {
			data = append(data, s.data[i])
			if s.isNullable {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}
	}
	return GDLSeriesInt32{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

///////////////////////////////		GROUPING OPERATIONS			/////////////////////////

func (s GDLSeriesInt32) Group() GSeriesPartition {
	return nil
}

func (s GDLSeriesInt32) SubGroup(gp GSeriesPartition) GSeriesPartition {
	return nil
}
