package main

var TEMPLATE_BASIC_ACCESSORS = `package gandalff

import (
	"fmt"
	"time"
	"preludiometa"
)

func (s {{.SeriesName}}) printInfo() {
	fmt.Println("{{.SeriesName}}")
	fmt.Println("==========")
	fmt.Println("IsNullable:", s.isNullable)
	fmt.Println("Sorted:    ", s.sorted)
	fmt.Println("Data:      ", s.data)
	fmt.Println("NullMask:  ", s.nullMask)
	fmt.Println("Partition: ", s.partition)
	fmt.Println("Context:   ", s.ctx)
}

////////////////////////			BASIC ACCESSORS

// Return the context of the series.
func (s {{.SeriesName}}) GetContext() *Context {
	return s.ctx
}

// Return the number of elements in the series.
func (s {{.SeriesName}}) Len() int {
	return len(s.data)
}

// Return the type of the series.
func (s {{.SeriesName}}) Type() preludiometa.BaseType {
	return preludiometa.{{.SeriesTypeStr}}
}

// Return the type and cardinality of the series.
func (s {{.SeriesName}}) TypeCard() preludiometa.BaseTypeCard {
	return preludiometa.BaseTypeCard{Base: preludiometa.{{.SeriesTypeStr}}, Card: s.Len()}
}

// Return if the series is grouped.
func (s {{.SeriesName}}) IsGrouped() bool {
	return s.partition != nil
}

// Return if the series admits null values.
func (s {{.SeriesName}}) IsNullable() bool {
	return s.isNullable
}

// Return if the series is sorted.
func (s {{.SeriesName}}) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Return if the series is error.
func (s {{.SeriesName}}) IsError() bool {
	return false
}

// Return the error message of the series.
func (s {{.SeriesName}}) GetError() string {
	return ""
}

// Return the partition of the series.
func (s {{.SeriesName}}) GetPartition() SeriesPartition {
	return s.partition
}

// Return if the series has null values.
func (s {{.SeriesName}}) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Return the number of null values in the series.
func (s {{.SeriesName}}) NullCount() int {
	count := 0
	for _, x := range s.nullMask {
		for ; x != 0; x >>= 1 {
			count += int(x & 1)
		}
	}
	return count
}

// Return if the element at index i is null.
func (s {{.SeriesName}}) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i>>3]&(1<<uint(i%8)) != 0
	}
	return false
}

// Return the null mask of the series.
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

// Set the null mask of the series.
func (s {{.SeriesName}}) SetNullMask(mask []bool) Series {
	if s.partition != nil {
		return SeriesError{"{{.SeriesName}}.SetNullMask: cannot set values on a grouped series"}
	}

	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				s.nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}
		return s
	} else {
		nullMask := __binVecInit(len(s.data), false)
		for k, v := range mask {
			if v {
				nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Make the series nullable.
func (s {{.SeriesName}}) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __binVecInit(len(s.data), false)
	}
	return s
}

// Make the series non-nullable.
func (s {{.SeriesName}}) MakeNonNullable() Series {
	if s.isNullable {
		s.isNullable = false
		s.nullMask = make([]uint8, 0)
	}
	return s
}

// Get the element at index i.
func (s {{.SeriesName}}) Get(i int) any {
	return {{if .IsGoTypePtr}}*{{end}}s.data[i]
}

// Append appends a value or a slice of values to the series.
func (s {{.SeriesName}}) Append(v any) Series {
	if s.partition != nil {
		return SeriesError{"{{.SeriesName}}.Append: cannot append values to a grouped series"}
	}

	switch v := v.(type) {
	case nil:
		s.data = append(s.data, {{.DefaultValue}})
		s = s.MakeNullable().({{.SeriesName}})
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		s.nullMask[(len(s.data)-1)>>3] |= 1 << uint8((len(s.data)-1)%8)

	case SeriesNA:
		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, v.Len(), true, __binVecInit(v.Len(), true))
		s.data = append(s.data, make([]{{.SeriesGoTypeStr}}, v.Len())...)

	case {{.SeriesGoOuterTypeStr}}:
		{{if eq .SeriesName "SeriesString" -}}
		s.data = append(s.data, s.ctx.stringPool.Put(v))
		{{- else -}}
		s.data = append(s.data, v)
		{{- end}}
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []{{.SeriesGoOuterTypeStr}}:
		{{if eq .SeriesName "SeriesString" -}}
		s.data = append(s.data, make([]*string, len(v))...)
		for i, str := range v {
			s.data[len(s.data)-len(v)+i] = s.ctx.stringPool.Put(str)
		}
		{{- else -}}
		s.data = append(s.data, v...)
		{{- end}}
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	case {{.SeriesNullableTypeStr}}:
		{{if eq .SeriesName "SeriesString" -}}
		s.data = append(s.data, s.ctx.stringPool.Put(v.Value))
		{{- else -}}
		s.data = append(s.data, v.Value)
		{{- end}}
		s = s.MakeNullable().({{.SeriesName}})
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[(len(s.data)-1)>>3] |= 1 << uint8((len(s.data)-1)%8)
		}

	case []{{.SeriesNullableTypeStr}}:
		ssize := len(s.data)
		s.data = append(s.data, make([]{{.SeriesGoTypeStr}}, len(v))...)
		s = s.MakeNullable().({{.SeriesName}})
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for i, b := range v {
			{{if eq .SeriesName "SeriesString" -}}
			s.data[ssize+i] = s.ctx.stringPool.Put(b.Value)
			{{- else -}}
			s.data[ssize+i] = b.Value
			{{- end}}
			if !b.Valid {
				s.nullMask[(ssize+i)>>3] |= 1 << uint8((ssize+i)%8)
			}
		}

	case {{.SeriesName}}:
		if s.ctx != v.ctx {
			return SeriesError{"{{.SeriesName}}.Append: cannot append {{.SeriesName}} from different contexts"}
		}

		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, len(v.data), v.isNullable, v.nullMask)
		s.data = append(s.data, v.data...)

	default:
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.Append: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s {{.SeriesName}}) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess("{{.SeriesName}}", s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces, false)
}

// Return the elements of the series as a slice.
func (s {{.SeriesName}}) Data() any {
	{{if eq .SeriesName "SeriesString" -}}
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return data
	{{- else -}}
	return s.data
	{{- end}}
}

// Copy the series.
func (s {{.SeriesName}}) Copy() Series {
	data := make([]{{.SeriesGoTypeStr}}, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return {{.SeriesName}}{
		isNullable: s.isNullable,
		sorted:     s.sorted,
		data:       data,
		nullMask:   nullMask,
		partition:  s.partition,
		ctx:        s.ctx,
	}
}

func (s {{.SeriesName}}) getDataPtr() *[]{{.SeriesGoTypeStr}} {
	return &s.data
}

// Ungroup the series.
func (s {{.SeriesName}}) UnGroup() Series {
	s.partition = nil
	return s
}
`

var TEMPLATE_FILTERS = `
////////////////////////			FILTER OPERATIONS

// Filters out the elements by the given mask.
// Mask can be SeriesBool, SeriesBoolMemOpt, SeriesInt, bool slice or a int slice.
func (s {{.SeriesName}}) Filter(mask any) Series {
	switch mask := mask.(type) {
	case SeriesBool:
		return s.filterBoolSlice(mask.data)
	case SeriesBoolMemOpt:
		return s.filterBoolMemOpt(mask)
	case SeriesInt:
		return s.filterIntSlice(mask.data, true)
	case []bool:
		return s.filterBoolSlice(mask)
	case []int:
		return s.filterIntSlice(mask, true)
	default:
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.Filter: invalid type %T", mask)}
	}
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

	data := make([]{{.SeriesGoTypeStr}}, elementCount)
	if s.isNullable {
		nullMask = __binVecInit(elementCount, false)
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
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.Filter: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []{{.SeriesGoTypeStr}}
	var nullMask []uint8

	data = make([]{{.SeriesGoTypeStr}}, elementCount)

	if s.isNullable {
		nullMask = __binVecInit(elementCount, false)
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

func (s {{.SeriesName}}) filterIntSlice(indexes []int, check bool) Series {
	if len(indexes) == 0 {
		s.data = make([]{{.SeriesGoTypeStr}}, 0)
		s.nullMask = make([]uint8, 0)
		return s
	}

	// check if indexes are in range
	if check {
		for _, v := range indexes {
			if v < 0 || v >= len(s.data) {
				return SeriesError{fmt.Sprintf("{{.SeriesName}}.Filter: index %d is out of range", v)}
			}
		}
	}

	var data []{{.SeriesGoTypeStr}}
	var nullMask []uint8

	size := len(indexes)
	data = make([]{{.SeriesGoTypeStr}}, size)

	if s.isNullable {
		nullMask = __binVecInit(size, false)
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

var TEMPLATE_MAPS = `
// Apply the given function to each element of the series.
func (s {{.SeriesName}}) Map(f MapFunc) Series {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:
		data := make([]bool, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f({{if .IsGoTypePtr}}*{{end}}s.data[i]).(bool)
		}

		return SeriesBool{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int:
		data := make([]int, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f({{if .IsGoTypePtr}}*{{end}}s.data[i]).(int)
		}

		return SeriesInt{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int64:
		data := make([]int64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f({{if .IsGoTypePtr}}*{{end}}s.data[i]).(int64)
		}

		return SeriesInt64{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case float64:
		data := make([]float64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f({{if .IsGoTypePtr}}*{{end}}s.data[i]).(float64)
		}

		return SeriesFloat64{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case string:
		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = s.ctx.stringPool.Put(f({{if .IsGoTypePtr}}*{{end}}s.data[i]).(string))
		}

		return SeriesString{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Time:
		data := make([]time.Time, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f({{if .IsGoTypePtr}}*{{end}}s.data[i]).(time.Time)
		}

		return SeriesTime{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Duration:
		data := make([]time.Duration, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(time.Duration)
		}

		return SeriesDuration{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.Map: Unsupported type %T", v)}
	}
}

// Apply the given function to each element of the series.
func (s {{.SeriesName}}) MapNull(f MapFuncNull) Series {
	if len(s.data) == 0 {
		return s
	}

	if !s.isNullable {
		return SeriesError{"{{.SeriesName}}.MapNull: series is not nullable"}
	}

	v, isNull := f(s.Get(0), s.IsNull(0))
	switch v.(type) {
	case bool:
		data := make([]bool, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f({{if .IsGoTypePtr}}*{{end}}s.data[i], s.IsNull(i))
			data[i] = v.(bool)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return SeriesBool{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int:
		data := make([]int, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f({{if .IsGoTypePtr}}*{{end}}s.data[i], s.IsNull(i))
			data[i] = v.(int)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return SeriesInt{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int64:
		data := make([]int64, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f({{if .IsGoTypePtr}}*{{end}}s.data[i], s.IsNull(i))
			data[i] = v.(int64)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return SeriesInt64{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case float64:
		data := make([]float64, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f({{if .IsGoTypePtr}}*{{end}}s.data[i], s.IsNull(i))
			data[i] = v.(float64)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return SeriesFloat64{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case string:
		data := make([]*string, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f({{if .IsGoTypePtr}}*{{end}}s.data[i], s.IsNull(i))
			data[i] = s.ctx.stringPool.Put(v.(string))
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return SeriesString{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Time:
		data := make([]time.Time, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f({{if .IsGoTypePtr}}*{{end}}s.data[i], s.IsNull(i))
			data[i] = v.(time.Time)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return SeriesTime{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Duration:
		data := make([]time.Duration, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = v.(time.Duration)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return SeriesDuration{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return SeriesError{fmt.Sprintf("{{.SeriesName}}.MapNull: Unsupported type %T", v)}
	}
}
`
