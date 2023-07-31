package gandalff

import (
	"fmt"
	"typesys"
)

// SeriesBoolMemOpt represents a series of bools.
// The data is stored as a byte array, with each bit representing a bool.
type SeriesBoolMemOpt struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	size       int
	name       string
	data       []uint8
	nullMask   []uint8
	partition  *SeriesBoolMemOptPartition
}

////////////////////////			BASIC ACCESSORS

func (s SeriesBoolMemOpt) __trueCount() int {
	count := 0
	for _, v := range s.data {
		for i := 0; i < 8; i++ {
			count += int((v & (1 << uint(i))) >> uint(i))
		}
	}
	return count
}

// Returns the number of elements in the series.
func (s SeriesBoolMemOpt) Len() int {
	return s.size
}

// Returns the name of the series.
func (s SeriesBoolMemOpt) Name() string {
	return s.name
}

// Returns the type of the series.
func (s SeriesBoolMemOpt) Type() typesys.BaseType {
	return typesys.BoolType
}

// Returns the type and cardinality of the series.
func (s SeriesBoolMemOpt) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{typesys.BoolType, s.Len()}
}

// Returns if the series is grouped.
func (s SeriesBoolMemOpt) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s SeriesBoolMemOpt) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s SeriesBoolMemOpt) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series is error.
func (s SeriesBoolMemOpt) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s SeriesBoolMemOpt) GetError() string {
	return ""
}

// Returns if the series has null values.
func (s SeriesBoolMemOpt) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s SeriesBoolMemOpt) NullCount() int {
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
func (s SeriesBoolMemOpt) NonNullCount() int {
	return s.size - s.NullCount()
}

// Returns if the element at index i is null.
func (s SeriesBoolMemOpt) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i>>3]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s SeriesBoolMemOpt) SetNull(i int) Series {
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
func (s SeriesBoolMemOpt) GetNullMask() []bool {
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
func (s SeriesBoolMemOpt) SetNullMask(mask []bool) Series {
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
func (s SeriesBoolMemOpt) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.sorted = SORTED_NONE
		s.nullMask = make([]uint8, len(s.data))
	}
	return s
}

// Get the element at index i.
func (s SeriesBoolMemOpt) Get(i int) any {
	return s.data[i>>3]&(1<<uint(i%8)) != 0
}

// Get the element at index i as a string.
func (s SeriesBoolMemOpt) GetString(i int) string {
	if s.isNullable && s.nullMask[i>>3]&(1<<uint(i%8)) != 0 {
		return NULL_STRING
	} else if s.data[i>>3]&(1<<uint(i%8)) != 0 {
		return BOOL_TRUE_STRING
	} else {
		return BOOL_FALSE_STRING
	}
}

// Set the element at index i. The value must be of type bool or NullableBool.
func (s SeriesBoolMemOpt) Set(i int, v any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.Set: provided value %t is not of type bool or NullableBool", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesBoolMemOpt) Take(start, end, step int) Series {
	return s
}

func (s SeriesBoolMemOpt) Less(i, j int) bool {
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

func (s SeriesBoolMemOpt) Swap(i, j int) {
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
func (s SeriesBoolMemOpt) Append(v any) Series {
	switch v := v.(type) {
	case bool, []bool:
		return s.appendRaw(v)
	case NullableBool, []NullableBool:
		return s.appendNullable(v)
	case SeriesBoolMemOpt:
		return s.appendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesBoolMemOpt) appendRaw(v any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.Append: invalid type %T", v)}
	}

	s.size = size
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s SeriesBoolMemOpt) appendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesBoolMemOpt.AppendNullable: series is not nullable"}
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
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.AppendNullable: invalid type %T", v)}
	}

	s.size = size
	return s
}

// AppendSeries appends a series to the series.
func (s SeriesBoolMemOpt) appendSeries(other Series) Series {
	var ok bool
	var o SeriesBoolMemOpt
	if o, ok = other.(SeriesBoolMemOpt); !ok {
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.AppendSeries: invalid type %T", other)}
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

func (s SeriesBoolMemOpt) Data() any {
	data := make([]bool, s.size)
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < s.size; j++ {
			data[i*8+j] = v&(1<<uint(j)) != 0
		}
	}
	return data
}

// NullableData returns a slice of NullableBool.
func (s SeriesBoolMemOpt) DataAsNullable() any {
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
func (s SeriesBoolMemOpt) DataAsString() []string {
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
func (s SeriesBoolMemOpt) Cast(t typesys.BaseType, stringPool *StringPool) Series {
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
			return SeriesError{"SeriesBoolMemOpt.Cast: StringPool is nil"}
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
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.Cast: invalid type %s", t.ToString())}
	}
}

// Copy returns a copy of the series.
func (s SeriesBoolMemOpt) Copy() Series {
	data := make([]uint8, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesBoolMemOpt{
		isGrouped:  s.isGrouped,
		isNullable: s.isNullable,
		size:       s.size,
		data:       data,
		nullMask:   nullMask,
		partition:  s.partition,
	}
}

func (s SeriesBoolMemOpt) getDataPtr() *[]uint8 {
	return &s.data
}

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s SeriesBoolMemOpt) Filter(mask any) Series {
	switch mask := mask.(type) {
	case SeriesBool:
		return s.filterBool(mask)
	case SeriesBoolMemOpt:
		return s.filterBoolMemOpt(mask)
	case []bool:
		return s.filterBoolSlice(mask)
	case []int:
		return s.filterIntSlice(mask)
	default:
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.Filter: invalid type %T", mask)}
	}
}

func (s SeriesBoolMemOpt) filterBool(mask SeriesBool) Series {
	return s.filterBoolSlice(mask.data)
}

func (s SeriesBoolMemOpt) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
	if mask.size != s.size {
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.Filter: mask length (%d) does not match series length (%d)", mask.size, s.size)}
	}

	if mask.isNullable {
		return SeriesError{"SeriesBoolMemOpt.Filter: mask series cannot be nullable for this operation"}
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

func (s SeriesBoolMemOpt) filterBoolSlice(mask []bool) Series {
	if len(mask) != s.size {
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), s.size)}
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

func (s SeriesBoolMemOpt) filterIntSlice(indexes []int) Series {
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

func (s SeriesBoolMemOpt) Map(f GDLMapFunc, stringPool *StringPool) Series {
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
			return SeriesError{"SeriesBoolMemOpt.Map: StringPool is nil"}
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

	return SeriesError{fmt.Sprintf("SeriesBoolMemOpt.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

// A partition is trivially a vector of maps (or boolIndices in this case)
// Each element of the vector represent a sub-group (the default is 1,
// which means no sub-grouping).
// So is for the null group, which has the same size as the partition vector.
type SeriesBoolMemOptPartition struct {
	series    *SeriesBoolMemOpt
	partition map[int64][]int
	nulls     []int
}

func (p SeriesBoolMemOptPartition) GetSize() int {
	return len(p.partition)
}

func (p SeriesBoolMemOptPartition) GetMap() map[int64][]int {
	return p.partition
}

func (p SeriesBoolMemOptPartition) GetValueIndices(val any) []int {
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

func (gp SeriesBoolMemOptPartition) GetKeys() any {
	keys := make([]bool, 0, 2)
	return keys
}

func (gp SeriesBoolMemOptPartition) debugPrint() {
	fmt.Println("SeriesBoolMemOptPartition")
	data := gp.series.Data().([]bool)
	for k, v := range gp.partition {
		fmt.Printf("%10v - %5v: %v\n", k, data[v[0]], v)
	}
}

func (s SeriesBoolMemOpt) Group() Series {
	map_ := make(map[int64][]int)
	for index := 0; index < s.size; index++ {
		map_[int64((s.data[index>>3]&(1<<(index%8)))>>int64(index%8))] = append(map_[int64((s.data[index>>3]&(1<<(index%8)))>>int64(index%8))], index)
	}

	return SeriesBoolMemOpt{
		isGrouped:  true,
		isNullable: s.isNullable,
		sorted:     s.sorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition: &SeriesBoolMemOptPartition{
			series:    &s,
			partition: map_,
			nulls:     nil,
		}}
}

func (s SeriesBoolMemOpt) SubGroup(partition SeriesPartition) Series {
	newMap := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

	var newHash int64
	for h, indexes := range partition.GetMap() {
		for _, index := range indexes {
			newHash = int64((s.data[index>>3]&(1<<(index%8)))>>int64(index%8)) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
			newMap[newHash] = append(newMap[newHash], index)
		}
	}

	return SeriesBoolMemOpt{
		isGrouped:  true,
		isNullable: s.isNullable,
		sorted:     s.sorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition: &SeriesBoolMemOptPartition{
			series:    &s,
			partition: newMap,
			nulls:     nil,
		}}
}

func (s SeriesBoolMemOpt) GetPartition() SeriesPartition {
	return s.partition
}

func (s SeriesBoolMemOpt) Sort() Series {
	return s
}

func (s SeriesBoolMemOpt) SortRev() Series {
	return s
}

////////////////////////			SORTING OPERATIONS

////////////////////////			LOGIC OPERATIONS

// And performs logical AND operation between two series
// If one of the series is nullable, the result series will be nullable
// If the other series is not a boolean series, the result will be an error
func (s SeriesBoolMemOpt) And(other Series) Series {
	if other.Type() != typesys.BoolType {
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt: cannot perform AND operation between %T and %T", s, other)}
	}

	o := other.(SeriesBoolMemOpt)

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
func (s SeriesBoolMemOpt) Or(other Series) Series {
	if other.Type() != typesys.BoolType {
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt: cannot perform OR operation between %T and %T", s, other)}
	}

	o := other.(SeriesBoolMemOpt)
	if s.size != o.size {
		return SeriesError{fmt.Sprintf("SeriesBoolMemOpt: cannot perform OR operation between series of different sizes: %d and %d", s.size, o.size)}
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
func (s SeriesBoolMemOpt) Not() Series {
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

func (s SeriesBoolMemOpt) Mul(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot multiply SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Div(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot divide SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Mod(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot modulo SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Pow(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot power SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Add(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot add SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Sub(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot subtract SeriesBoolMemOpt and %T", other)}
}

////////////////////////			LOGICAL OPERATIONS

func (s SeriesBoolMemOpt) Eq(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot compare SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Ne(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot compare SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Gt(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot compare SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Ge(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot compare SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Lt(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot compare SeriesBoolMemOpt and %T", other)}
}

func (s SeriesBoolMemOpt) Le(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot compare SeriesBoolMemOpt and %T", other)}
}
