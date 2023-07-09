package gandalff

import (
	"fmt"
	"typesys"
)

// SeriesBool represents a series of bools.
// The data is stored as a byte array, with each bit representing a bool.
type SeriesBool struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	size       int
	name       string
	data       []uint8
	nullMask   []uint8
	partition  *SeriesBoolPartition
}

func NewSeriesBool(name string, isNullable bool, data []bool) Series {
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

	return SeriesBool{isNullable: isNullable, name: name, size: size, data: actualData, nullMask: nullMask}
}

////////////////////////			BASIC ACCESSORS

func (s SeriesBool) __trueCount() int {
	count := 0
	for _, v := range s.data {
		for i := 0; i < 8; i++ {
			count += int((v & (1 << uint(i))) >> uint(i))
		}
	}
	return count
}

// Returns the number of elements in the series.
func (s SeriesBool) Len() int {
	return s.size
}

// Returns the name of the series.
func (s SeriesBool) Name() string {
	return s.name
}

// Returns the type of the series.
func (s SeriesBool) Type() typesys.BaseType {
	return typesys.BoolType
}

// Returns if the series is grouped.
func (s SeriesBool) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s SeriesBool) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s SeriesBool) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series has null values.
func (s SeriesBool) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s SeriesBool) NullCount() int {
	count := 0
	for _, x := range s.nullMask {
		for x != 0 {
			count += int(x & 1)
			x >>= 1
		}
	}
	return count
}

// Returns the number of non-null values in the series.
func (s SeriesBool) NonNullCount() int {
	return s.size - s.NullCount()
}

// Returns if the element at index i is null.
func (s SeriesBool) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i>>3]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s SeriesBool) SetNull(i int) Series {
	if s.isNullable {
		s.nullMask[i>>3] |= 1 << uint(i%8)

		s.sorted = SORTED_NONE
		return s
	} else {
		nullMask := make([]uint8, len(s.data))
		nullMask[i>>3] |= 1 << uint(i%8)

		s.isNullable = true
		s.sorted = SORTED_NONE
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s SeriesBool) GetNullMask() []bool {
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
func (s SeriesBool) SetNullMask(mask []bool) Series {
	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				s.nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}

		s.sorted = SORTED_NONE
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
		s.sorted = SORTED_NONE
		s.nullMask = nullMask

		return s
	}
}

// Makes the series nullable.
func (s SeriesBool) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.sorted = SORTED_NONE
		s.nullMask = make([]uint8, len(s.data))
	}
	return s
}

// Get the element at index i.
func (s SeriesBool) Get(i int) any {
	return s.data[i>>3]&(1<<uint(i%8)) != 0
}

// Get the element at index i as a string.
func (s SeriesBool) GetString(i int) string {
	if s.isNullable && s.nullMask[i>>3]&(1<<uint(i%8)) != 0 {
		return NULL_STRING
	} else if s.data[i>>3]&(1<<uint(i%8)) != 0 {
		return BOOL_TRUE_STRING
	} else {
		return BOOL_FALSE_STRING
	}
}

// Set the element at index i. The value must be of type bool or NullableBool.
func (s SeriesBool) Set(i int, v any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesBool.Set: provided value %t is not of type bool or NullableBool", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesBool) Take(start, end, step int) Series {
	return s
}

func (s SeriesBool) Less(i, j int) bool {
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

func (s SeriesBool) Swap(i, j int) {
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
func (s SeriesBool) Append(v any) Series {
	switch v := v.(type) {
	case bool:
		return s.AppendRaw(v)
	case []bool:
		return s.AppendRaw(v)
	case NullableBool:
		return s.AppendNullable(v)
	case []NullableBool:
		return s.AppendNullable(v)
	case SeriesBool:
		return s.AppendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesBool.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesBool) AppendRaw(v any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesBool.Append: invalid type %T", v)}
	}

	s.size = size
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s SeriesBool) AppendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesBool.AppendNullable: series is not nullable"}
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
		return SeriesError{fmt.Sprintf("SeriesBool.AppendNullable: invalid type %T", v)}
	}

	s.size = size
	return s
}

// AppendSeries appends a series to the series.
func (s SeriesBool) AppendSeries(other Series) Series {
	var ok bool
	var o SeriesBool
	if o, ok = other.(SeriesBool); !ok {
		return SeriesError{fmt.Sprintf("SeriesBool.AppendSeries: invalid type %T", other)}
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

////////////////////////			ALL DATA ACCESSORS

func (s SeriesBool) Data() any {
	data := make([]bool, s.size)
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < s.size; j++ {
			data[i*8+j] = v&(1<<uint(j)) != 0
		}
	}
	return data
}

// NullableData returns a slice of NullableBool.
func (s SeriesBool) DataAsNullable() any {
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
func (s SeriesBool) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
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
	} else {
		for i, v := range s.data {
			for j := 0; j < 8 && i*8+j < len(s.data); j++ {
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

// Casts the series to a given type.
func (s SeriesBool) Cast(t typesys.BaseType, stringPool *StringPool) Series {
	switch t {
	case typesys.BoolType:
		return s

	case typesys.Int32Type:
		data := make([]int32, s.size)
		for i, v := range s.data {
			for j := 0; j < 8 && i*8+j < s.size; j++ {
				if v&(1<<uint(j)) != 0 {
					data[i*8+j] = 1
				}
			}
		}

		return SeriesInt32{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     s.sorted,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
		}

	case typesys.Int64Type:
		data := make([]int64, s.size)
		for i, v := range s.data {
			for j := 0; j < 8 && i*8+j < s.size; j++ {
				if v&(1<<uint(j)) != 0 {
					data[i*8+j] = 1
				}
			}
		}

		return SeriesInt64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     s.sorted,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
		}

	case typesys.Float64Type:
		data := make([]float64, s.size)
		for i, v := range s.data {
			for j := 0; j < 8 && i*8+j < s.size; j++ {
				if v&(1<<uint(j)) != 0 {
					data[i*8+j] = 1.0
				}
			}
		}

		return SeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     s.sorted,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
		}

	case typesys.StringType:
		if stringPool == nil {
			return SeriesError{"SeriesBool.Cast: StringPool is nil"}
		}

		data := make([]*string, s.size)
		if s.isNullable {
			for i, v := range s.data {
				for j := 0; j < 8 && i*8+j < s.size; j++ {
					if s.nullMask[i]&(1<<uint(j)) != 0 {
						data[i*8+j] = stringPool.Put(NULL_STRING)
					} else {
						if v&(1<<uint(j)) != 0 {
							data[i*8+j] = stringPool.Put(BOOL_TRUE_STRING)
						} else {
							data[i*8+j] = stringPool.Put(BOOL_FALSE_STRING)
						}
					}
				}
			}
		} else {
			for i, v := range s.data {
				for j := 0; j < 8 && i*8+j < s.size; j++ {
					if v&(1<<uint(j)) != 0 {
						data[i*8+j] = stringPool.Put(BOOL_TRUE_STRING)
					} else {
						data[i*8+j] = stringPool.Put(BOOL_FALSE_STRING)
					}
				}
			}
		}

		return SeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     s.sorted,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesBool.Cast: invalid type %s", t.ToString())}
	}
}

// Copy returns a copy of the series.
func (s SeriesBool) Copy() Series {
	data := make([]uint8, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesBool{
		isGrouped:  s.isGrouped,
		isNullable: s.isNullable,
		size:       s.size,
		data:       data,
		nullMask:   nullMask,
		partition:  s.partition,
	}
}

func (s SeriesBool) getDataPtr() *[]uint8 {
	return &s.data
}

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask series.
func (s SeriesBool) Filter(mask SeriesBool) Series {
	if mask.size != s.size {
		return SeriesError{fmt.Sprintf("SeriesBool.Filter: mask length (%d) does not match series length (%d)", mask.size, s.size)}
	}

	if mask.isNullable {
		return SeriesError{"SeriesBool.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()

	data := __binVecInit(elementCount)
	var nullMask []uint8

	if s.isNullable {

		nullMask = __binVecInit(elementCount)

		dstIdx := 0
		for srcIdx := 0; srcIdx < s.size; srcIdx++ {
			if mask.data[srcIdx>>3]&(1<<uint(srcIdx%8)) != 0 {

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
		for srcIdx := 0; srcIdx < s.size; srcIdx++ {
			if mask.data[srcIdx>>3]&(1<<uint(srcIdx%8)) != 0 {
				if srcIdx%8 > dstIdx%8 {
					data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				} else {
					data[dstIdx>>3] |= ((s.data[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
				}
				dstIdx++
			}
		}
	}

	s.size = elementCount
	s.data = data
	s.nullMask = nullMask

	return s
}

// FilterByMask returns a new series with elements filtered by the mask.
func (s SeriesBool) FilterByMask(mask []bool) Series {
	if len(mask) != s.size {
		return SeriesError{fmt.Sprintf("SeriesBool.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), s.size)}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	data := __binVecInit(elementCount)
	var nullMask []uint8

	if s.isNullable {

		nullMask = __binVecInit(elementCount)

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

	s.size = elementCount
	s.data = data
	s.nullMask = nullMask

	return s
}

func (s SeriesBool) FilterByIndeces(indexes []int) Series {
	var data []uint8
	var nullMask []uint8

	size := len(indexes)
	data = __binVecInit(len(indexes))

	if s.isNullable {
		nullMask = __binVecInit(size)
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

	s.size = size
	s.data = data
	s.nullMask = nullMask

	return s
}

func (s SeriesBool) Map(f GDLMapFunc, stringPool *StringPool) Series {
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
		s.sorted = SORTED_NONE
		s.data = data

		return s

	case int32:
		data := make([]int32, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = f(s.data[i>>3]&(1<<uint(i%8)) != 0).(int32)
		}

		return SeriesInt32{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case int64:
		data := make([]int64, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = f(s.data[i>>3]&(1<<uint(i%8)) != 0).(int64)
		}

		return SeriesInt64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case float64:
		data := make([]float64, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = f(s.data[i>>3]&(1<<uint(i%8)) != 0).(float64)
		}

		return SeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case string:
		if stringPool == nil {
			return SeriesError{"SeriesBool.Map: StringPool is nil"}
		}

		data := make([]*string, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = stringPool.Put(f(s.data[i>>3]&(1<<uint(i%8)) != 0).(string))
		}

		return SeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       stringPool,
		}
	}

	return SeriesError{fmt.Sprintf("SeriesBool.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

// A partition is trivially a vector of maps (or boolIndices in this case)
// Each element of the vector represent a sub-group (the default is 1,
// which means no sub-grouping).
// So is for the null group, which has the same size as the partition vector.
type SeriesBoolPartition struct {
	series    *SeriesBool
	partition map[int64][]int
	nulls     []int
}

func (p SeriesBoolPartition) GetSize() int {
	return len(p.partition)
}

func (p SeriesBoolPartition) GetMap() map[int64][]int {
	return p.partition
}

func (p SeriesBoolPartition) GetValueIndices(val any) []int {
	if val == nil {
		return p.nulls
	} else if v, ok := val.(bool); ok {
		if v {
			return p.partition[1]
		} else {
			return p.partition[0]
		}
	}

	return make([]int, 0)
}

func (gp SeriesBoolPartition) GetKeys() any {
	keys := make([]bool, 0, 2)
	return keys
}

func (gp SeriesBoolPartition) debugPrint() {
	fmt.Println("SeriesBoolPartition")
	data := gp.series.Data().([]bool)
	for k, v := range gp.partition {
		fmt.Printf("%10v - %5v: %v\n", k, data[v[0]], v)
	}
}

func (s SeriesBool) Group() Series {
	map_ := make(map[int64][]int)
	for index := 0; index < s.size; index++ {
		map_[int64((s.data[index>>3]&(1<<(index%8)))>>int64(index%8))] = append(map_[int64((s.data[index>>3]&(1<<(index%8)))>>int64(index%8))], index)
	}

	return SeriesBool{
		isGrouped:  true,
		isNullable: s.isNullable,
		sorted:     s.sorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition: &SeriesBoolPartition{
			series:    &s,
			partition: map_,
			nulls:     nil,
		}}
}

func (s SeriesBool) SubGroup(partition SeriesPartition) Series {
	newMap := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

	var newHash int64
	for h, indexes := range partition.GetMap() {
		for _, index := range indexes {
			newHash = int64((s.data[index>>3]&(1<<(index%8)))>>int64(index%8)) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
			newMap[newHash] = append(newMap[newHash], index)
		}
	}

	return SeriesBool{
		isGrouped:  true,
		isNullable: s.isNullable,
		sorted:     s.sorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition: &SeriesBoolPartition{
			series:    &s,
			partition: newMap,
			nulls:     nil,
		}}
}

func (s SeriesBool) GetPartition() SeriesPartition {
	return s.partition
}

func (s SeriesBool) Sort() Series {
	return s
}

func (s SeriesBool) SortRev() Series {
	return s
}

////////////////////////			SORTING OPERATIONS

////////////////////////			LOGIC OPERATIONS

// And performs logical AND operation between two series
// If one of the series is nullable, the result series will be nullable
// If the other series is not a boolean series, the result will be nil
func (s SeriesBool) And(other Series) Series {
	if other.Type() != typesys.BoolType {
		return SeriesError{fmt.Sprintf("SeriesBool: cannot perform AND operation between %T and %T", s, other)}
	}

	o := other.(SeriesBool)
	if s.size != o.size {
		return SeriesError{fmt.Sprintf("SeriesBool: cannot perform AND operation between series of different sizes: %d and %d", s.size, o.size)}
	}

	sNullCnt := s.NullCount()
	oNullCnt := o.NullCount()

	if sNullCnt > 0 || oNullCnt > 0 {
		if s.isNullable {
			if o.isNullable {
				// both are nullable
				for i := 0; i < len(s.data); i++ {
					s.nullMask[i] |= o.nullMask[i]
					s.data[i] &= o.data[i]
				}

				return s
			} else
			// s is nullable, o is not nullable
			{
				for i := 0; i < len(s.data); i++ {
					s.data[i] &= o.data[i]
				}

				return s
			}
		} else if o.isNullable {
			// s is not nullable, o is nullable
			for i := 0; i < len(s.data); i++ {
				s.data[i] &= o.data[i]
			}

			s.isNullable = true
			s.nullMask = o.nullMask

			return s
		}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] &= o.data[i]
	}

	return s
}

// Or performs logical OR operation between two series
// If one of the series is nullable, the result series will be nullable
// If the other series is not a boolean series, the result will be nil
func (s SeriesBool) Or(other Series) Series {
	if other.Type() != typesys.BoolType {
		return SeriesError{fmt.Sprintf("SeriesBool: cannot perform OR operation between %T and %T", s, other)}
	}

	o := other.(SeriesBool)
	if s.size != o.size {
		return SeriesError{fmt.Sprintf("SeriesBool: cannot perform OR operation between series of different sizes: %d and %d", s.size, o.size)}
	}

	sNullCnt := s.NullCount()
	oNullCnt := o.NullCount()

	if sNullCnt > 0 || oNullCnt > 0 {
		if s.isNullable {
			if o.isNullable {
				// both are nullable
				for i := 0; i < len(s.data); i++ {
					s.nullMask[i] |= o.nullMask[i]
					s.data[i] |= o.data[i]
				}

				return s
			} else
			// s is nullable, o is not nullable
			{
				for i := 0; i < len(s.data); i++ {
					s.data[i] |= o.data[i]
				}

				return s
			}
		} else if o.isNullable {
			// s is not nullable, o is nullable
			for i := 0; i < len(s.data); i++ {
				s.data[i] |= o.data[i]
			}

			s.isNullable = true
			s.nullMask = o.nullMask

			return s
		}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] |= o.data[i]
	}

	return s
}

// Not performs logical NOT operation on series
func (s SeriesBool) Not() Series {
	for i := 0; i < len(s.data); i++ {
		s.data[i] = ^s.data[i]
	}

	// clear the unused bits
	for i := s.size; i < len(s.data)*8; i++ {
		s.data[i>>3] &^= 1 << (i % 8)
	}

	return s
}

////////////////////////			ARITHMETIC OPERATIONS

// func (s SeriesBool) Mul(other Series) Series {
// 	return s
// }

// func (s SeriesBool) Div(other Series) Series {
// 	return s
// }

// func (s SeriesBool) Mod(other Series) Series {
// 	return s
// }

// func (s SeriesBool) Add(other Series) Series {
// 	return s
// }

// func (s SeriesBool) Sub(other Series) Series {
// 	return s
// }
