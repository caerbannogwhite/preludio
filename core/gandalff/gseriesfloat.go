package gandalff

import "errors"

// GSeriesFloat represents a series of floats.
type GSeriesFloat struct {
	isNullable bool
	name       string
	data       []float64
	nullMap    []uint8
}

func NewGSeriesFloat(name string, isNullable bool, makeCopy bool, data []float64) GSeriesFloat {
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

	return GSeriesFloat{isNullable: isNullable, name: name, data: data, nullMap: nullMap}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////////

func (s GSeriesFloat) Len() int {
	return len(s.data)
}

func (s GSeriesFloat) IsNullable() bool {
	return s.isNullable
}

func (s GSeriesFloat) MakeNullable() {
	if !s.isNullable {
		s.isNullable = true
		if len(s.data)%8 == 0 {
			s.nullMap = make([]uint8, len(s.data)/8)
		} else {
			s.nullMap = make([]uint8, len(s.data)/8+1)
		}
	}
}

func (s GSeriesFloat) Name() string {
	return s.name
}

func (s GSeriesFloat) Type() GSeriesType {
	return FloatType
}

func (s GSeriesFloat) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GSeriesFloat) NullCount() int {
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

func (s GSeriesFloat) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GSeriesFloat) SetNull(i int) error {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
		return nil
	}
	return errors.New("GSeriesFloat.SetNull: series is not nullable")
}

func (s GSeriesFloat) GetNullMask() []bool {
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

func (s GSeriesFloat) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return errors.New("GSeriesFloat.SetNullMask: series is not nullable")
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

func (s GSeriesFloat) Get(i int) interface{} {
	return s.data[i]
}

func (s GSeriesFloat) Set(i int, v interface{}) {
	s.data[i] = v.(float64)
}

// Append appends a value or a slice of values to the series.
func (s GSeriesFloat) Append(v interface{}) error {
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
			return errors.New("GSeriesFloat.Append: invalid type")
		}
	} else {
		if b, ok := v.(float64); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]float64); ok {
			s.data = append(s.data, bv...)
		} else {
			return errors.New("GSeriesFloat.Append: invalid type")
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GSeriesFloat) AppendNullable(v interface{}) error {
	if !s.isNullable {
		return errors.New("GSeriesFloat.AppendNullable: series is not nullable")
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
		return errors.New("GSeriesFloat.AppendNullable: invalid type")
	}

	return nil
}

///////////////////////////////		ALL DATA ACCESSORS			/////////////////////////

func (s GSeriesFloat) Data() interface{} {
	return s.data
}

func (s GSeriesFloat) NullableData() interface{} {
	data := make([]NullableFloat, len(s.data))
	for i, v := range s.data {
		data[i] = NullableFloat{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GSeriesFloat) StringData() []string {
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

func (s GSeriesFloat) Copy() GSeries {
	data := make([]float64, len(s.data))
	copy(data, s.data)
	nullMap := make([]uint8, len(s.nullMap))
	copy(nullMap, s.nullMap)

	return GSeriesFloat{isNullable: s.isNullable, name: s.name, data: data, nullMap: s.nullMap}
}

///////////////////////////////		SERIES OPERATIONS			/////////////////////////

func (s GSeriesFloat) Filter(mask []bool) GSeries {
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
	return GSeriesFloat{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}
