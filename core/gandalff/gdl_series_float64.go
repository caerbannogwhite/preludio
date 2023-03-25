package gandalff

import (
	"errors"
	"typesys"
)

// GDLSeriesFloat64 represents a series of floats.
type GDLSeriesFloat64 struct {
	isNullable bool
	name       string
	data       []float64
	nullMap    []uint8
}

func NewGDLSeriesFloat64(name string, isNullable bool, makeCopy bool, data []float64) GDLSeriesFloat64 {
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
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GDLSeriesFloat64{isNullable: isNullable, name: name, data: data, nullMap: nullMap}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////////

func (s GDLSeriesFloat64) Len() int {
	return len(s.data)
}

func (s GDLSeriesFloat64) IsNullable() bool {
	return s.isNullable
}

func (s GDLSeriesFloat64) MakeNullable() {
	if !s.isNullable {
		s.isNullable = true
		if len(s.data)%8 == 0 {
			s.nullMap = make([]uint8, len(s.data)/8)
		} else {
			s.nullMap = make([]uint8, len(s.data)/8+1)
		}
	}
}

func (s GDLSeriesFloat64) Name() string {
	return s.name
}

func (s GDLSeriesFloat64) Type() typesys.BaseType {
	return typesys.Float64Type
}

func (s GDLSeriesFloat64) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesFloat64) NullCount() int {
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

func (s GDLSeriesFloat64) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesFloat64) SetNull(i int) error {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
		return nil
	}
	return errors.New("GDLSeriesFloat64.SetNull: series is not nullable")
}

func (s GDLSeriesFloat64) GetNullMask() []bool {
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

func (s GDLSeriesFloat64) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return errors.New("GDLSeriesFloat64.SetNullMask: series is not nullable")
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

func (s GDLSeriesFloat64) Get(i int) interface{} {
	return s.data[i]
}

func (s GDLSeriesFloat64) Set(i int, v interface{}) {
	s.data[i] = v.(float64)
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesFloat64) Append(v interface{}) error {
	if s.isNullable {
		if b, ok := v.(float64); ok {
			s.data = append(s.data, b)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
		} else if bv, ok := v.([]float64); ok {
			s.data = append(s.data, bv...)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, make([]uint8, len(s.data)/8-len(s.nullMap))...)
			}
		} else {
			return errors.New("GDLSeriesFloat64.Append: invalid type")
		}
	} else {
		if b, ok := v.(float64); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]float64); ok {
			s.data = append(s.data, bv...)
		} else {
			return errors.New("GDLSeriesFloat64.Append: invalid type")
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesFloat64) AppendNullable(v interface{}) error {
	if !s.isNullable {
		return errors.New("GDLSeriesFloat64.AppendNullable: series is not nullable")
	}

	if b, ok := v.(NullableFloat); ok {
		s.data = append(s.data, b.Value)
		if len(s.data)/8 > len(s.nullMap) {
			s.nullMap = append(s.nullMap, 0)
		}
		if !b.Valid {
			s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableFloat); ok {
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
		return errors.New("GDLSeriesFloat64.AppendNullable: invalid type")
	}

	return nil
}

///////////////////////////////		ALL DATA ACCESSORS			/////////////////////////

func (s GDLSeriesFloat64) Data() interface{} {
	return s.data
}

func (s GDLSeriesFloat64) NullableData() interface{} {
	data := make([]NullableFloat, len(s.data))
	for i, v := range s.data {
		data[i] = NullableFloat{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GDLSeriesFloat64) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		if s.IsNull(i) {
			data[i] = NULL_STRING
		} else {
			data[i] = floatToString(v)
		}
	}
	return data
}

func (s GDLSeriesFloat64) Copy() GSeries {
	data := make([]float64, len(s.data))
	copy(data, s.data)
	nullMap := make([]uint8, len(s.nullMap))
	copy(nullMap, s.nullMap)

	return GDLSeriesFloat64{isNullable: s.isNullable, name: s.name, data: data, nullMap: s.nullMap}
}

///////////////////////////////		SERIES OPERATIONS			/////////////////////////

func (s GDLSeriesFloat64) Filter(mask []bool) GSeries {
	data := make([]float64, 0)
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range mask {
		if v {
			data = append(data, s.data[i])
			if s.isNullable {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}
	}
	return GDLSeriesFloat64{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

///////////////////////////////		GROUPING OPERATIONS			/////////////////////////

func (s GDLSeriesFloat64) Group() GSeriesPartition {
	return nil
}

func (s GDLSeriesFloat64) SubGroup(gp GSeriesPartition) GSeriesPartition {
	return nil
}
