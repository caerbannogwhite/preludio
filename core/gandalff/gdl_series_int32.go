package gandalff

import (
	"fmt"
	"typesys"
)

// GDLSeriesInt32 represents a series of ints.
type GDLSeriesInt32 struct {
	isNullable bool
	name       string
	data       []int
	nullMask   []uint8
}

func NewGDLSeriesInt32(name string, isNullable bool, makeCopy bool, data []int) GDLSeriesInt32 {
	var nullMask []uint8
	if isNullable {
		if len(data)%8 == 0 {
			nullMask = make([]uint8, (len(data) >> 3))
		} else {
			nullMask = make([]uint8, (len(data)>>3)+1)
		}
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GDLSeriesInt32{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////

func (s GDLSeriesInt32) Len() int {
	return len(s.data)
}

func (s GDLSeriesInt32) IsNullable() bool {
	return s.isNullable
}

func (s GDLSeriesInt32) MakeNullable() GDLSeries {
	if !s.isNullable {
		if len(s.data)%8 == 0 {
			s.nullMask = make([]uint8, (len(s.data) >> 3))
		} else {
			s.nullMask = make([]uint8, (len(s.data)>>3)+1)
		}
		s.isNullable = true
	}
	return s
}

func (s GDLSeriesInt32) Name() string {
	return s.name
}

func (s GDLSeriesInt32) Type() typesys.BaseType {
	return typesys.Int32Type
}

func (s GDLSeriesInt32) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesInt32) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
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
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesInt32) SetNull(i int) error {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	}
	return fmt.Errorf("GDLSeriesInt32.SetNull: series is not nullable")
}

func (s GDLSeriesInt32) GetNullMask() []bool {
	mask := make([]bool, len(s.data))
	idx := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8 && idx < len(s.data); i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

func (s GDLSeriesInt32) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return fmt.Errorf("GDLSeriesInt32.SetNullMask: series is not nullable")
	}

	for k, v := range mask {
		if v {
			s.nullMask[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMask[k/8] &= ^(1 << uint(k%8))
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

func (s GDLSeriesInt32) Append(v interface{}) GDLSeries {
	switch v := v.(type) {
	case int:
		return s.AppendRaw(v)
	case []int:
		return s.AppendRaw(v)
	case NullableInt32:
		return s.AppendNullable(v)
	case []NullableInt32:
		return s.AppendNullable(v)
	case GDLSeriesInt32:
		return s.AppendSeries(v)
	case GDLSeriesError:
		return v
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesInt32) AppendRaw(v interface{}) GDLSeries {
	if s.isNullable {
		if i, ok := v.(int); ok {
			s.data = append(s.data, i)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if iv, ok := v.([]int); ok {
			s.data = append(s.data, iv...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendRaw: invalid type %T", v)}
		}
	} else {
		if b, ok := v.(int); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]int); ok {
			s.data = append(s.data, bv...)
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesInt32) AppendNullable(v interface{}) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesInt32.AppendNullable: series is not nullable"}
	}

	if b, ok := v.(NullableInt32); ok {
		s.data = append(s.data, b.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !b.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableInt32); ok {
		if len(s.data) > len(s.nullMask)<<8 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for _, b := range bv {
			s.data = append(s.data, b.Value)
			if !b.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s GDLSeriesInt32) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesInt32
	if o, ok = other.(GDLSeriesInt32); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendSeries: invalid type %T", other)}
	}

	if s.isNullable {
		if o.isNullable {
			s.data = append(s.data, o.data...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
			}

			// merge null masks
			sIdx := len(s.data) - len(o.data)
			oIdx := 0
			for _, v := range o.nullMask {
				for j := 0; j < 8; j++ {
					if v&(1<<uint(j)) != 0 {
						s.nullMask[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}
		} else {
			s.data = append(s.data, o.data...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
			}
		}
	} else {
		if o.isNullable {
			s.data = append(s.data, o.data...)
			if len(s.data)%8 == 0 {
				s.nullMask = make([]uint8, (len(s.data) >> 3))
			} else {
				s.nullMask = make([]uint8, (len(s.data)>>3)+1)
			}
			s.isNullable = true

			// merge null masks
			sIdx := len(s.data) - len(o.data)
			oIdx := 0
			for _, v := range o.nullMask {
				for j := 0; j < 8; j++ {
					if v&(1<<uint(j)) != 0 {
						s.nullMask[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}
		} else {
			s.data = append(s.data, o.data...)
		}
	}

	return s
}

///////////////////////////////  	ALL DATA ACCESSORS  /////////////////////////////////

func (s GDLSeriesInt32) Data() interface{} {
	return s.data
}

func (s GDLSeriesInt32) NullableData() interface{} {
	data := make([]NullableInt32, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt32{Valid: !s.IsNull(i), Value: v}
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

func (s GDLSeriesInt32) Copy() GDLSeries {
	data := make([]int, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return GDLSeriesInt32{isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

///////////////////////////////  	SERIES OPERATIONS  //////////////////////////////////

// FilterByMask returns a new series with elements filtered by the mask.
func (s GDLSeriesInt32) FilterByMask(mask []bool) GDLSeries {
	if len(mask) != len(s.data) {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []int
	var nullMask []uint8

	data = make([]int, elementCount)

	if s.isNullable {

		if elementCount%8 == 0 {
			nullMask = make([]uint8, (elementCount >> 3))
		} else {
			nullMask = make([]uint8, (elementCount>>3)+1)
		}

		dstIdx := 0
		for srcIdx, v := range mask {
			if v {
				data[dstIdx] = s.data[srcIdx]
				if srcIdx%8 > dstIdx%8 {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				} else {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
				}
				dstIdx++
			}
		}
	} else {
		dstIdx := 0
		for srcIdx, v := range mask {
			if v {
				data[dstIdx] = s.data[srcIdx]
				dstIdx++
			}
		}
	}

	return GDLSeriesInt32{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
	}
}

func (s GDLSeriesInt32) FilterByIndeces(indexes []int) GDLSeries {
	var data []int
	var nullMask []uint8

	size := len(indexes)
	data = make([]int, size)

	if s.isNullable {

		if size%8 == 0 {
			nullMask = make([]uint8, (size >> 3))
		} else {
			nullMask = make([]uint8, (size>>3)+1)
		}

		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
			if srcIdx%8 > dstIdx%8 {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
			}
		}
	} else {
		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
		}
	}

	return GDLSeriesInt32{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
	}
}

///////////////////////////////		GROUPING OPERATIONS			/////////////////////////

func (s GDLSeriesInt32) Group() GDLSeriesPartition {
	return nil
}

func (s GDLSeriesInt32) SubGroup(gp GDLSeriesPartition) GDLSeriesPartition {
	return nil
}
