package gandalff

import (
	"fmt"
	"typesys"
)

// GDLSeriesBool represents a series of bools.
// The data is stored as a byte array, with each bit representing a bool.
type GDLSeriesBool struct {
	isGrouped  bool
	isNullable bool
	isSorted   bool
	size       int
	name       string
	data       []uint8
	nullMask   []uint8
	partition  GDLSeriesBoolPartition
}

func NewGDLSeriesBool(name string, isNullable bool, data []bool) GDLSeries {
	size := len(data)
	var actualData []uint8
	if size%8 == 0 {
		actualData = make([]uint8, (size >> 3))
		for i := 0; i < size; i++ {
			if data[i] {
				actualData[i>>3] |= 1 << uint(i%8)
			}
		}
	} else {
		actualData = make([]uint8, (size>>3)+1)
		for i := 0; i < size; i++ {
			if data[i] {
				actualData[i>>3] |= 1 << uint(i%8)
			}
		}
	}

	var nullMask []uint8
	if isNullable {
		nullMask = make([]uint8, len(actualData))

	} else {
		nullMask = make([]uint8, 0)
	}

	return GDLSeriesBool{isNullable: isNullable, name: name, size: size, data: actualData, nullMask: nullMask}
}

///////////////////////////////		BASIC ACCESSORS		/////////////////////////////////

// Returns the number of elements in the series.
func (s GDLSeriesBool) Len() int {
	return s.size
}

// Returns the name of the series.
func (s GDLSeriesBool) Name() string {
	return s.name
}

// Returns the type of the series.
func (s GDLSeriesBool) Type() typesys.BaseType {
	return typesys.BoolType
}

// Returns if the series is grouped.
func (s GDLSeriesBool) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s GDLSeriesBool) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s GDLSeriesBool) IsSorted() bool {
	return s.isSorted
}

// Returns if the series has null values.
func (s GDLSeriesBool) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s GDLSeriesBool) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8; i++ {
			count += int((v & (1 << uint(i))) >> uint(i))
		}
	}
	return count
}

// Returns the number of non-null values in the series.
func (s GDLSeriesBool) NonNullCount() int {
	return s.size - s.NullCount()
}

// Returns if the element at index i is null.
func (s GDLSeriesBool) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i>>3]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s GDLSeriesBool) SetNull(i int) GDLSeries {
	if s.isNullable {
		s.nullMask[i>>3] |= 1 << uint(i%8)

		s.isSorted = false
		return s
	} else {
		nullMask := make([]uint8, len(s.data))
		nullMask[i>>3] |= 1 << uint(i%8)

		s.isNullable = true
		s.isSorted = false
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s GDLSeriesBool) GetNullMask() []bool {
	mask := make([]bool, s.size)
	idx := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8 && idx < s.size; i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

// Sets the null mask of the series.
func (s GDLSeriesBool) SetNullMask(mask []bool) GDLSeries {
	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				s.nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}

		s.isSorted = false
		return s
	} else {
		nullMask := make([]uint8, len(s.data))
		for k, v := range mask {
			if v {
				nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}

		s.isNullable = true
		s.isSorted = false
		s.nullMask = nullMask

		return s
	}
}

// Makes the series nullable.
func (s GDLSeriesBool) MakeNullable() GDLSeries {
	if !s.isNullable {
		s.isNullable = true
		s.isSorted = false
		s.nullMask = make([]uint8, len(s.data))
	}
	return s
}

// Get the element at index i.
func (s GDLSeriesBool) Get(i int) any {
	return s.data[i>>3]&(1<<uint(i%8)) != 0
}

// Set the element at index i. The value must be of type bool or NullableBool.
func (s GDLSeriesBool) Set(i int, v any) GDLSeries {
	if b, ok := v.(bool); ok {
		if b {
			s.data[i>>3] |= 1 << uint(i%8)
		} else {
			s.data[i>>3] &= ^(1 << uint(i%8))
		}
	} else if nb, ok := v.(NullableBool); ok {
		if nb.Valid {
			if nb.Value {
				s.data[i>>3] |= 1 << uint(i%8)
			} else {
				s.data[i>>3] &= ^(1 << uint(i%8))
			}
		} else {
			s.data[i>>3] &= ^(1 << uint(i%8))
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.Set: provided value %t is not of type bool or NullableBool", v)}
	}

	s.isSorted = false
	return s
}

func (s GDLSeriesBool) Less(i, j int) bool {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			return false
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			return true
		}
	}
	return s.data[i>>3]&(1<<uint(i%8)) > 0 && s.data[j>>3]&(1<<uint(j%8)) == 0
}

func (s GDLSeriesBool) Swap(i, j int) {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			s.nullMask[j>>3] |= 1 << uint(j%8)
		} else {
			s.nullMask[j>>3] &= ^(1 << uint(j%8))
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			s.nullMask[i>>3] |= 1 << uint(i%8)
		} else {
			s.nullMask[i>>3] &= ^(1 << uint(i%8))
		}
	}
	if s.data[i>>3]&(1<<uint(i%8)) > 0 {
		s.data[j>>3] |= 1 << uint(j%8)
	} else {
		s.data[j>>3] &= ^(1 << uint(j%8))
	}
	if s.data[j>>3]&(1<<uint(j%8)) > 0 {
		s.data[i>>3] |= 1 << uint(i%8)
	} else {
		s.data[i>>3] &= ^(1 << uint(i%8))
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesBool) Append(v any) GDLSeries {
	switch v := v.(type) {
	case bool:
		return s.AppendRaw(v)
	case []bool:
		return s.AppendRaw(v)
	case NullableBool:
		return s.AppendNullable(v)
	case []NullableBool:
		return s.AppendNullable(v)
	case GDLSeriesBool:
		return s.AppendSeries(v)
	case GDLSeriesError:
		return v
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesBool) AppendRaw(v any) GDLSeries {
	var size int
	if b, ok := v.(bool); ok {

		// adjust size of data and nullMask if necessary
		size = s.size + 1
		if size > len(s.data)<<3 {
			s.data = append(s.data, 0)
			if s.isNullable {
				s.nullMask = append(s.nullMask, 0)
			}
		}

		if b {
			s.data[len(s.data)-1] |= 1 << uint(s.size%8)
		}
	} else if bv, ok := v.([]bool); ok {

		// adjust size of data and nullMask if necessary
		size = s.size + len(bv)
		if size > len(s.data)<<3 {
			s.data = append(s.data, make([]uint8, (size>>3)-len(s.data)+1)...)
			if s.isNullable {
				s.nullMask = append(s.nullMask, make([]uint8, (size>>3)-len(s.nullMask)+1)...)
			}
		}

		idx := s.size
		for _, b := range bv {
			if b {
				s.data[idx>>3] |= 1 << uint(idx%8)
			}
			idx++
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.Append: invalid type %T", v)}
	}

	s.size = size
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesBool) AppendNullable(v any) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesBool.AppendNullable: series is not nullable"}
	}

	var size int
	if b, ok := v.(NullableBool); ok {
		// adjust size of data and nullMask if necessary
		size = s.size + 1
		if size > len(s.data)<<3 {
			s.data = append(s.data, 0)
			s.nullMask = append(s.nullMask, 0)
		}

		if !b.Valid {
			s.nullMask[len(s.nullMask)-1] |= 1 << uint(s.size%8)
		}
		if b.Value {
			s.data[len(s.data)-1] |= 1 << uint(s.size%8)
		}
	} else if bv, ok := v.([]NullableBool); ok {
		// adjust size of data and nullMask if necessary
		size = s.size + len(bv)
		if size > len(s.data)<<3 {
			s.data = append(s.data, make([]uint8, (size>>3)-len(s.data)+1)...)
			s.nullMask = append(s.nullMask, make([]uint8, (size>>3)-len(s.nullMask)+1)...)
		}

		idx := s.size
		for _, b := range bv {
			if !b.Valid {
				s.nullMask[idx>>3] |= 1 << uint(idx%8)
			}
			if b.Value {
				s.data[idx>>3] |= 1 << uint(idx%8)
			}
			idx++
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.AppendNullable: invalid type %T", v)}
	}

	s.size = size
	return s
}

// AppendSeries appends a series to the series.
func (s GDLSeriesBool) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesBool
	if o, ok = other.(GDLSeriesBool); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.AppendSeries: invalid type %T", other)}
	}

	size := s.size + o.size

	if s.isNullable {
		// adjust size of data and nullMask if necessary
		if size > len(s.data)<<3 {
			s.data = append(s.data, make([]uint8, (size>>3)-len(s.data)+1)...)
			s.nullMask = append(s.nullMask, make([]uint8, (size>>3)-len(s.nullMask)+1)...)
		}

		// both series are nullable
		if o.isNullable {
			sIdx := s.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8 && sIdx < size; j++ {
					// TODO: optimize?
					if v&(1<<uint(j)) != 0 {
						s.data[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					if o.nullMask[oIdx>>3]&(1<<uint(j)) != 0 {
						s.nullMask[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}
		} else

		// s is nullable, o is not nullable
		{
			sIdx := s.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8 && sIdx < size; j++ {
					// TODO: optimize?
					if v&(1<<uint(j)) != 0 {
						s.data[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}
		}
	} else {
		// s is not nullable, o is nullable
		if o.isNullable {
			if s.size > len(s.data)<<3 {
				s.data = append(s.data, make([]uint8, (s.size>>3)-len(s.data)+1)...)
				s.nullMask = make([]uint8, len(s.data))
			}

			sIdx := s.size - o.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8; j++ {
					if v&(1<<uint(j)) != 0 {
						s.data[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					if o.nullMask[oIdx>>3]&(1<<uint(j)) != 0 {
						s.nullMask[sIdx>>3] |= 1 << uint(sIdx%8)
					}

					sIdx++
					oIdx++
				}
			}

			s.isNullable = true
		} else

		// both series are not nullable
		{
			if s.size > len(s.data)<<3 {
				s.data = append(s.data, make([]uint8, (s.size>>3)-len(s.data)+1)...)
			}

			sIdx := s.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8 && sIdx < size; j++ {
					// TODO: optimize?
					if v&(1<<uint(j)) != 0 {
						s.data[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					if o.nullMask[oIdx>>3]&(1<<uint(j)) != 0 {
						s.nullMask[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}
		}
	}

	s.size = size
	return s
}

/////////////////////////////// 		ALL DATA ACCESSORS		///////////////////////////////

func (s GDLSeriesBool) Data() any {
	data := make([]bool, s.size)
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < s.size; j++ {
			data[i*8+j] = v&(1<<uint(j)) != 0
		}
	}
	return data
}

// NullableData returns a slice of NullableBool.
func (s GDLSeriesBool) NullableData() any {
	data := make([]NullableBool, len(s.data))
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < len(s.data); j++ {
			if s.nullMask[i]&(1<<uint(j)) != 0 {
				data[i*8+j] = NullableBool{false, false}
			} else {
				data[i*8+j] = NullableBool{v&(1<<uint(j)) != 0, true}
			}
		}
	}
	return data
}

// StringData returns a slice of strings.
func (s GDLSeriesBool) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < len(s.data); j++ {
			if s.nullMask[i]&(1<<uint(j)) != 0 {
				data[i*8+j] = NULL_STRING
			} else {
				if v&(1<<uint(j)) != 0 {
					data[i*8+j] = BOOL_TRUE_STRING
				} else {
					data[i*8+j] = BOOL_FALSE_STRING
				}
			}
		}
	}
	return data
}

// Copy returns a copy of the series.
func (s GDLSeriesBool) Copy() GDLSeries {
	data := make([]uint8, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return GDLSeriesBool{
		isGrouped:  s.isGrouped,
		isNullable: s.isNullable,
		size:       s.size,
		data:       data,
		nullMask:   nullMask,
		partition:  s.partition,
	}
}

///////////////////////////////		SERIES OPERATIONS		/////////////////////////////

// FilterByMask returns a new series with elements filtered by the mask.
func (s GDLSeriesBool) FilterByMask(mask []bool) GDLSeries {
	if len(mask) != s.size {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), s.size)}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []uint8
	var nullMask []uint8

	if elementCount%8 == 0 {
		data = make([]uint8, (elementCount >> 3))
	} else {
		data = make([]uint8, (elementCount>>3)+1)
	}

	if s.isNullable {

		if elementCount%8 == 0 {
			nullMask = make([]uint8, (elementCount >> 3))
		} else {
			nullMask = make([]uint8, (elementCount>>3)+1)
		}

		dstIdx := 0
		for srcIdx, v := range mask {
			if v {

				// s.data[srcIdx>>3] 			-> 	selects the byte in s.data that contains the bit
				// 1 << uint(srcIdx%8)			-> 	shifts a 1 to the position of the bit
				// >> uint(srcIdx%8-dstIdx%8))	-> 	shifts the bit to the position of the bit in the destination byte
				//
				// TODO: optimize? is there a better way to select the destination bit?
				if srcIdx%8 > dstIdx%8 {
					data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				} else {
					data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
				}
				dstIdx++
			}
		}
	} else {
		dstIdx := 0
		for srcIdx, v := range mask {
			if v {
				if srcIdx%8 > dstIdx%8 {
					data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				} else {
					data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
				}
				dstIdx++
			}
		}
	}

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		size:       elementCount,
		data:       data,
		nullMask:   nullMask,
	}
}

func (s GDLSeriesBool) FilterByIndeces(indexes []int) GDLSeries {
	var data []uint8
	var nullMask []uint8

	size := len(indexes)
	if size%8 == 0 {
		data = make([]uint8, (size >> 3))
	} else {
		data = make([]uint8, (size>>3)+1)
	}

	if s.isNullable {

		if size%8 == 0 {
			nullMask = make([]uint8, (size >> 3))
		} else {
			nullMask = make([]uint8, (size>>3)+1)
		}

		for dstIdx, srcIdx := range indexes {
			if srcIdx%8 > dstIdx%8 {
				data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
			}
		}
	} else {
		for dstIdx, srcIdx := range indexes {
			if srcIdx%8 > dstIdx%8 {
				data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
			}
		}
	}

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		size:       size,
		data:       data,
		nullMask:   nullMask,
	}
}

func (s GDLSeriesBool) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	if s.size == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:
		data := make([]uint8, len(s.data))
		for i := 0; i < s.size; i++ {
			if f(s.data[i>>3]&(1<<uint(i%8)) != 0).(bool) {
				data[i>>3] |= (1 << uint(i%8))
			}
		}

		s.isGrouped = false
		s.isSorted = false
		s.data = data

		return s

	case int:
		data := make([]int, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = f(s.data[i>>3]&(1<<uint(i%8)) != 0).(int)
		}

		return GDLSeriesInt32{
			isGrouped:  false,
			isNullable: s.isNullable,
			isSorted:   false,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case float64:
		data := make([]float64, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = f(s.data[i>>3]&(1<<uint(i%8)) != 0).(float64)
		}

		return GDLSeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			isSorted:   false,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case string:
		if stringPool == nil {
			return GDLSeriesError{"GDLSeriesBool.Map: StringPool is nil"}
		}

		data := make([]*string, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = stringPool.Add(f(s.data[i>>3]&(1<<uint(i%8)) != 0).(string))
		}

		return GDLSeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			isSorted:   false,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       stringPool,
		}
	}

	return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.Map: Unsupported type %T", v)}
}

///////////////////////////////		GROUPING OPERATIONS		/////////////////////////////

type boolIndices struct {
	trues  []int
	falses []int
}

// A partition is trivially a vector of maps (or boolIndices in this case)
// Each element of the vector represent a sub-group (the default is 1,
// which means no sub-grouping).
// So is for the null group, which has the same size as the partition vector.
type GDLSeriesBoolPartition struct {
	partition []boolIndices
	nullGroup [][]int
}

func (p GDLSeriesBoolPartition) GetSize() int {
	return len(p.partition)
}

func (p GDLSeriesBoolPartition) GetGroupsCount() int {
	count := 0
	for _, s := range p.partition {
		if len(s.trues) > 0 {
			count++
		}
		if len(s.falses) > 0 {
			count++
		}
	}

	for _, g := range p.nullGroup {
		if len(g) > 0 {
			count++
		}
	}
	return count
}

func (p GDLSeriesBoolPartition) GetIndices() [][]int {
	indices := make([][]int, 0)

	for _, s := range p.partition {
		if len(s.trues) > 0 {
			indices = append(indices, s.trues)
		}
		if len(s.falses) > 0 {
			indices = append(indices, s.falses)
		}
	}

	for _, g := range p.nullGroup {
		if len(g) > 0 {
			indices = append(indices, g)
		}
	}

	return indices
}

func (p GDLSeriesBoolPartition) GetValueIndices(sub int, val interface{}) []int {
	if sub >= len(p.partition) {
		return nil
	}

	if v, ok := val.(bool); ok {
		if v {
			return p.partition[sub].trues
		}
		return p.partition[sub].falses
	}

	return nil
}

func (s GDLSeriesBoolPartition) GetNullIndices(sub int) []int {
	if sub >= len(s.nullGroup) {
		return nil
	}

	return s.nullGroup[sub]
}

func (s GDLSeriesBool) Group() GDLSeries {
	var nullGroup [][]int

	groups := boolIndices{make([]int, 0), make([]int, 0)}
	if s.isNullable {
		nullGroup = make([][]int, 1)
		nullGroup[0] = make([]int, 0)

		for i := 0; i < s.size; i++ {
			if s.IsNull(i) {
				nullGroup[0] = append(nullGroup[0], i)
			} else if s.data[i>>3]&(1<<(i%8)) > 0 {
				groups.trues = append(groups.trues, i)
			} else {
				groups.falses = append(groups.falses, i)
			}
		}
	} else {
		for i := 0; i < s.size; i++ {
			if s.data[i>>3]&(1<<(i%8)) > 0 {
				groups.trues = append(groups.trues, i)
			} else {
				groups.falses = append(groups.falses, i)
			}
		}
	}

	return GDLSeriesBool{
		isGrouped:  true,
		isNullable: s.isNullable,
		isSorted:   s.isSorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition: GDLSeriesBoolPartition{
			partition: []boolIndices{groups},
			nullGroup: nullGroup,
		}}
}

func (s GDLSeriesBool) SubGroup(partition GDLSeriesPartition) GDLSeries {
	var nullGroup [][]int

	groups := make([]boolIndices, 0)
	indices := partition.GetIndices()
	if s.isNullable {
		nullGroup = make([][]int, partition.GetGroupsCount())

		for _, g := range indices {
			bi := boolIndices{make([]int, 0), make([]int, 0)}
			for _, i := range g {
				if s.IsNull(i) {
					nullGroup[i] = append(nullGroup[i], i)
				} else if s.data[i>>3]&(1<<(i%8)) > 0 {
					bi.trues = append(bi.trues, i)
				} else {
					bi.falses = append(bi.falses, i)
				}
			}
			groups = append(groups, bi)
		}
	} else {
		for _, g := range indices {
			bi := boolIndices{make([]int, 0), make([]int, 0)}
			for _, i := range g {
				if s.data[i>>3]&(1<<(i%8)) > 0 {
					bi.trues = append(bi.trues, i)
				} else {
					bi.falses = append(bi.falses, i)
				}
			}
			groups = append(groups, bi)
		}
	}

	return GDLSeriesBool{
		isGrouped:  true,
		isNullable: s.isNullable,
		isSorted:   s.isSorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition: GDLSeriesBoolPartition{
			partition: groups,
			nullGroup: nullGroup,
		}}
}

func (s GDLSeriesBool) GetPartition() GDLSeriesPartition {
	return s.partition
}

func (s GDLSeriesBool) Sort() GDLSeries {
	return s
}

func (s GDLSeriesBool) SortRev() GDLSeries {
	return s
}

///////////////////////////////		SORTING OPERATIONS		/////////////////////////////

///////////////////////////////		LOGIC OPERATIONS		/////////////////////////////

// And performs logical AND operation between two series
// If one of the series is nullable, the result series will be nullable
// If the other series is not a boolean series, the result will be nil
func (s GDLSeriesBool) And(other GDLSeries) GDLSeries {
	if other.Type() != typesys.BoolType {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool: cannot perform AND operation between %T and %T", s, other)}
	}

	o := other.(GDLSeriesBool)
	if s.size != o.size {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool: cannot perform AND operation between series of different sizes: %d and %d", s.size, o.size)}
	}

	if s.isNullable || other.IsNullable() {
		data := make([]uint8, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			nullMask[i] = s.nullMask[i] | o.nullMask[i]
			data[i] = s.data[i] & o.data[i]
		}

		return GDLSeriesBool{
			isNullable: true,
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}
	}

	data := make([]uint8, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] = s.data[i] & o.data[i]
	}

	return GDLSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMask:   make([]uint8, 0),
	}
}

// Or performs logical OR operation between two series
// If one of the series is nullable, the result series will be nullable
// If the other series is not a boolean series, the result will be nil
func (s GDLSeriesBool) Or(other GDLSeries) GDLSeries {
	if other.Type() != typesys.BoolType {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool: cannot perform OR operation between %T and %T", s, other)}
	}

	o := other.(GDLSeriesBool)
	if s.size != o.size {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool: cannot perform OR operation between series of different sizes: %d and %d", s.size, o.size)}
	}

	if s.isNullable || other.IsNullable() {
		data := make([]uint8, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			nullMask[i] = s.nullMask[i] | o.nullMask[i]
			data[i] = s.data[i] | o.data[i]
		}

		return GDLSeriesBool{
			isNullable: true,
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}
	}

	data := make([]uint8, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] = s.data[i] | o.data[i]
	}

	return GDLSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMask:   make([]uint8, 0),
	}
}

// Not performs logical NOT operation on series
func (s GDLSeriesBool) Not() GDLSeries {
	data := make([]uint8, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] ^= s.data[i]
	}

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   s.nullMask,
	}
}

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////

func (s GDLSeriesBool) Mul(other GDLSeries) GDLSeries {
	return s
}

func (s GDLSeriesBool) Div(other GDLSeries) GDLSeries {
	return s
}

func (s GDLSeriesBool) Mod(other GDLSeries) GDLSeries {
	return s
}

func (s GDLSeriesBool) Add(other GDLSeries) GDLSeries {
	return s
}

func (s GDLSeriesBool) Sub(other GDLSeries) GDLSeries {
	return s
}
