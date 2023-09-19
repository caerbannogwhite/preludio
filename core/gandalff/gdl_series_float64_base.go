
package gandalff

import (
	"fmt"
	
	"typesys"
)

////////////////////////			BASIC ACCESSORS

// Return the number of elements in the series.
func (s SeriesFloat64) Len() int {
	return len(s.data)
}

// Return the name of the series.
func (s SeriesFloat64) Name() string {
	return s.name
}

// Set the name of the series.
func (s SeriesFloat64) SetName(name string) Series {
	s.name = name
	return s
}

// Return the StringPool of the series.
func (s SeriesFloat64) StringPool() *StringPool {
	return s.pool
}

// Set the StringPool for this series.
func (s SeriesFloat64) SetStringPool(pool *StringPool) Series {
	s.pool = pool
	return s
}

// Return the type of the series.
func (s SeriesFloat64) Type() typesys.BaseType {
	return typesys.Float64Type
}

// Return the type and cardinality of the series.
func (s SeriesFloat64) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{Base: typesys.Float64Type, Card: s.Len()}
}

// Return if the series is grouped.
func (s SeriesFloat64) IsGrouped() bool {
	return s.isGrouped
}

// Return if the series admits null values.
func (s SeriesFloat64) IsNullable() bool {
	return s.isNullable
}

// Return if the series is sorted.
func (s SeriesFloat64) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Return if the series is error.
func (s SeriesFloat64) IsError() bool {
	return false
}

// Return the error message of the series.
func (s SeriesFloat64) GetError() string {
	return ""
}

// Return if the series has null values.
func (s SeriesFloat64) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Return the number of null values in the series.
func (s SeriesFloat64) NullCount() int {
	count := 0
	for _, x := range s.nullMask {
		for ; x != 0; x >>= 1 {
			count += int(x & 1)
		}
	}
	return count
}

// Return if the element at index i is null.
func (s SeriesFloat64) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Set the element at index i to null.
func (s SeriesFloat64) SetNull(i int) Series {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	} else {
		nullMask := __binVecInit(len(s.data))
		nullMask[i/8] |= 1 << uint(i%8)

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Return the null mask of the series.
func (s SeriesFloat64) GetNullMask() []bool {
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

// Set the null mask of the series.
func (s SeriesFloat64) SetNullMask(mask []bool) Series {
	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k/8] |= 1 << uint(k%8)
			} else {
				s.nullMask[k/8] &= ^(1 << uint(k%8))
			}
		}
		return s
	} else {
		nullMask := __binVecInit(len(s.data))
		for k, v := range mask {
			if v {
				nullMask[k/8] |= 1 << uint(k%8)
			} else {
				nullMask[k/8] &= ^(1 << uint(k%8))
			}
		}

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Make the series nullable.
func (s SeriesFloat64) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __binVecInit(len(s.data))
	}
	return s
}

// Get the element at index i.
func (s SeriesFloat64) Get(i int) any {
	return s.data[i]
}

// Take the elements according to the given interval.
func (s SeriesFloat64) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess("SeriesFloat64", s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces, false)
}

// Return the elements of the series as a slice.
func (s SeriesFloat64) Data() any {
	return s.data
}

// Copy the series.
func (s SeriesFloat64) Copy() Series {
	data := make([]float64, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesFloat64{
		isGrouped:  s.isGrouped,
		isNullable: s.isNullable,
		sorted:     s.sorted,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
		pool:       s.pool,
		partition:  s.partition,
	}
}

func (s SeriesFloat64) getDataPtr() *[]float64 {
	return &s.data
}

////////////////////////			FILTER OPERATIONS

// Filters out the elements by the given mask.
// Mask can be SeriesBool, SeriesBoolMemOpt, bool slice or a int slice.
func (s SeriesFloat64) Filter(mask any) Series {
	switch mask := mask.(type) {
	case SeriesBool:
		return s.filterBool(mask)
	case SeriesBoolMemOpt:
		return s.filterBoolMemOpt(mask)
	case []bool:
		return s.filterBoolSlice(mask)
	case []int:
		return s.filterIntSlice(mask, true)
	default:
		return SeriesError{fmt.Sprintf("SeriesFloat64.Filter: invalid type %T", mask)}
	}
}

func (s SeriesFloat64) filterBool(mask SeriesBool) Series {
	return s.filterBoolSlice(mask.data)
}

func (s SeriesFloat64) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
	if mask.size != s.Len() {
		return SeriesError{fmt.Sprintf("SeriesFloat64.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return SeriesError{"SeriesFloat64.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()
	var nullMask []uint8

	data := make([]float64, elementCount)
	if s.isNullable {
		nullMask = __binVecInit(elementCount)
		dstIdx := 0
		for srcIdx := 0; srcIdx < s.Len(); srcIdx++ {
			if mask.data[srcIdx>>3]&(1<<uint(srcIdx%8)) != 0 {
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
		nullMask = make([]uint8, 0)
		dstIdx := 0
		for srcIdx := 0; srcIdx < s.Len(); srcIdx++ {
			if mask.data[srcIdx>>3]&(1<<uint(srcIdx%8)) != 0 {
				data[dstIdx] = s.data[srcIdx]
				dstIdx++
			}
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}

func (s SeriesFloat64) filterBoolSlice(mask []bool) Series {
	if len(mask) != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesFloat64.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []float64
	var nullMask []uint8

	data = make([]float64, elementCount)

	if s.isNullable {
		nullMask = __binVecInit(elementCount)
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
		nullMask = make([]uint8, 0)
		dstIdx := 0
		for srcIdx, v := range mask {
			if v {
				data[dstIdx] = s.data[srcIdx]
				dstIdx++
			}
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}

func (s SeriesFloat64) filterIntSlice(indexes []int, check bool) Series {
	if len(indexes) == 0 {
		s.data = make([]float64, 0)
		s.nullMask = make([]uint8, 0)
		return s
	}

	// check if indexes are in range
	if check {
		for _, v := range indexes {
			if v < 0 || v >= len(s.data) {
				return SeriesError{fmt.Sprintf("SeriesFloat64.Filter: index %d is out of range", v)}
			}
		}
	}

	var data []float64
	var nullMask []uint8

	size := len(indexes)
	data = make([]float64, size)

	if s.isNullable {
		nullMask = __binVecInit(size)
		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
			if srcIdx%8 > dstIdx%8 {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
			}
		}
	} else {
		nullMask = make([]uint8, 0)
		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}
