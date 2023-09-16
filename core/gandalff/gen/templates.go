package main

var TEMPLATE_BASIC_ACCESSORS = `
package gandalff

import "typesys"

////////////////////////			BASIC ACCESSORS

// Returns the number of elements in the series.
func (s {{.SeriesName}}) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s {{.SeriesName}}) Name() string {
	return s.name
}

// Sets the name of the series.
func (s {{.SeriesName}}) SetName(name string) Series {
	s.name = name
	return s
}

// Returns the type of the series.
func (s {{.SeriesName}}) Type() typesys.BaseType {
	return typesys.{{.SeriesType}}
}

// Returns the type and cardinality of the series.
func (s {{.SeriesName}}) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{Base: typesys.{{.SeriesType}}, Card: s.Len()}
}

// Returns if the series is grouped.
func (s {{.SeriesName}}) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s {{.SeriesName}}) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s {{.SeriesName}}) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series is error.
func (s {{.SeriesName}}) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s {{.SeriesName}}) GetError() string {
	return ""
}

// Returns if the series has null values.
func (s {{.SeriesName}}) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s {{.SeriesName}}) NullCount() int {
	count := 0
	for _, x := range s.nullMask {
		for ; x != 0; x >>= 1 {
			count += int(x & 1)
		}
	}
	return count
}

// Returns if the element at index i is null.
func (s {{.SeriesName}}) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s {{.SeriesName}}) SetNull(i int) Series {
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

// Returns the null mask of the series.
func (s {{.SeriesName}}) GetNullMask() []bool {
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

// Sets the null mask of the series.
func (s {{.SeriesName}}) SetNullMask(mask []bool) Series {
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

// Makes the series nullable.
func (s {{.SeriesName}}) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __binVecInit(len(s.data))
	}
	return s
}

// Get the element at index i.
func (s {{.SeriesName}}) Get(i int) any {
	return {{if .IsGoTypePtr}}*{{end}}s.data[i]
}
`

var TEMPLATE_FILTERS = `
// Filters out the elements by the given mask.
// Mask can be SeriesBool, SeriesBoolMemOpt, bool slice or a int slice.
func (s {{.SeriesName}}) Filter(mask any) Series {
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
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.Filter: invalid type %T", mask)}
	}
}

func (s {{.SeriesName}}) filterBool(mask SeriesBool) Series {
	return s.filterBoolSlice(mask.data)
}

func (s {{.SeriesName}}) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
	if mask.size != s.Len() {
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return SeriesError{"{{.SeriesName}}.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()
	var nullMask []uint8

	data := make([]{{.SeriesGoType}}, elementCount)
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

func (s {{.SeriesName}}) filterBoolSlice(mask []bool) Series {
	if len(mask) != len(s.data) {
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []{{.SeriesGoType}}
	var nullMask []uint8

	data = make([]{{.SeriesGoType}}, elementCount)

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

func (s {{.SeriesName}}) filterIntSlice(indexes []int) Series {
	if len(indexes) == 0 {
		s.data = make([]{{.SeriesGoType}}, 0)
		s.nullMask = make([]uint8, 0)
		return s
	}

	var data []{{.SeriesGoType}}
	var nullMask []uint8

	size := len(indexes)
	data = make([]{{.SeriesGoType}}, size)

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
`
