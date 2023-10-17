package gandalff

import (
	"fmt"
	"preludiometa"
	"time"
)

// SeriesNA represents a series with no data.
type SeriesNA struct {
	size      int
	partition *SeriesNAPartition
	ctx       *Context
}

func (s SeriesNA) printInfo() {}

// Return the context of the series.
func (s SeriesNA) GetContext() *Context {
	return s.ctx
}

// Returns the length of the series.
func (s SeriesNA) Len() int {
	return s.size
}

// Returns if the series is grouped.
func (s SeriesNA) IsGrouped() bool {
	return s.partition != nil
}

// Returns if the series admits null values.
func (s SeriesNA) IsNullable() bool {
	return true
}

func (s SeriesNA) IsSorted() SeriesSortOrder {
	return SORTED_ASC
}

// Returns if the series is error.
func (s SeriesNA) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s SeriesNA) GetError() string {
	return ""
}

// Makes the series nullable.
func (s SeriesNA) MakeNullable() Series {
	return s
}

// Make the series non-nullable.
func (s SeriesNA) MakeNonNullable() Series {
	return s
}

// Returns the type of the series.
func (s SeriesNA) Type() preludiometa.BaseType {
	return preludiometa.NullType
}

// Returns the type and cardinality of the series.
func (s SeriesNA) TypeCard() preludiometa.BaseTypeCard {
	return preludiometa.BaseTypeCard{Base: preludiometa.NullType, Card: s.Len()}
}

// Returns if the series has null values.
func (s SeriesNA) HasNull() bool {
	return true
}

// Returns the number of null values in the series.
func (s SeriesNA) NullCount() int {
	return s.size
}

// Returns if the element at index i is null.
func (s SeriesNA) IsNull(i int) bool {
	return true
}

// Returns the null mask of the series.
func (s SeriesNA) GetNullMask() []bool {
	nullMask := make([]bool, s.size)
	for i := 0; i < s.size; i++ {
		nullMask[i] = true
	}
	return nullMask
}

// Sets the null mask of the series.
func (s SeriesNA) SetNullMask(mask []bool) Series {
	return s
}

// Get the element at index i.
func (s SeriesNA) Get(i int) any {
	return nil
}

func (s SeriesNA) GetAsString(i int) string {
	return NULL_STRING
}

// Set the element at index i.
func (s SeriesNA) Set(i int, v any) Series {
	return s
}

// Take the elements according to the given interval.
func (s SeriesNA) Take(params ...int) Series {
	return s
}

// Append elements to the series.
func (s SeriesNA) Append(v any) Series {
	var nullMask []byte
	switch v := v.(type) {
	case nil:
		s.size++
		return s

	case SeriesNA:
		s.size += v.size
		return s

	case bool, NullableBool, []bool, []NullableBool, SeriesBool:
		var data []bool
		switch v := v.(type) {
		case bool:
			data = make([]bool, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case NullableBool:
			data = make([]bool, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []bool:
			data = append(make([]bool, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []NullableBool:
			data = make([]bool, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case SeriesBool:
			data = append(make([]bool, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return SeriesBool{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int, NullableInt, []int, []NullableInt, SeriesInt:
		var data []int
		switch v := v.(type) {
		case int:
			data = make([]int, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case NullableInt:
			data = make([]int, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []int:
			data = append(make([]int, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []NullableInt:
			data = make([]int, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case SeriesInt:
			data = append(make([]int, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return SeriesInt{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int64, NullableInt64, []int64, []NullableInt64, SeriesInt64:
		var data []int64
		switch v := v.(type) {
		case int64:
			data = make([]int64, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case NullableInt64:
			data = make([]int64, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []int64:
			data = append(make([]int64, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []NullableInt64:
			data = make([]int64, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case SeriesInt64:
			data = append(make([]int64, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return SeriesInt64{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case float64, NullableFloat64, []float64, []NullableFloat64, SeriesFloat64:
		var data []float64
		switch v := v.(type) {
		case float64:
			data = make([]float64, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case NullableFloat64:
			data = make([]float64, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []float64:
			data = append(make([]float64, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []NullableFloat64:
			data = make([]float64, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case SeriesFloat64:
			data = append(make([]float64, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return SeriesFloat64{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case string, NullableString, []string, []NullableString, SeriesString:
		data := make([]*string, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = s.ctx.stringPool.nullStringPtr
		}

		switch v := v.(type) {
		case string:
			data = append(data, s.ctx.stringPool.Put(v))
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case NullableString:
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data = append(data, s.ctx.stringPool.Put(v.Value))
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			} else {
				data = append(data, s.ctx.stringPool.nullStringPtr)
			}

		case []string:
			data = append(data, make([]*string, len(v))...)
			for i, v := range v {
				data[s.size+i] = s.ctx.stringPool.Put(v)
			}
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []NullableString:
			data = append(data, make([]*string, len(v))...)
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = s.ctx.stringPool.Put(v.Value)
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
					data[s.size+i] = s.ctx.stringPool.nullStringPtr
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case SeriesString:
			data = append(data, v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return SeriesString{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesNA.Append: invalid type %T", v)}
	}
}

// All-data accessors.

// Returns the actual data of the series.
func (s SeriesNA) Data() any {
	return make([]bool, s.size)
}

// Returns the nullable data of the series.
func (s SeriesNA) DataAsNullable() any {
	return make([]NullableBool, s.size)
}

// Returns the data of the series as a slice of strings.
func (s SeriesNA) DataAsString() []string {
	data := make([]string, s.size)
	for i := 0; i < s.size; i++ {
		data[i] = NULL_STRING
	}
	return data
}

// Casts the series to a given type.
func (s SeriesNA) Cast(t preludiometa.BaseType) Series {
	switch t {
	case preludiometa.NullType:
		return s

	case preludiometa.BoolType:
		return SeriesBool{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       make([]bool, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.IntType:
		return SeriesInt{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       make([]int, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Int64Type:
		return SeriesInt64{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       make([]int64, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Float64Type:
		return SeriesFloat64{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       make([]float64, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.StringType:
		return SeriesString{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       make([]*string, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.TimeType:
		return SeriesTime{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       make([]time.Time, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.DurationType:
		return SeriesDuration{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       make([]time.Duration, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesNA.Cast: invalid type %s", t.ToString())}
	}
}

// Copies the series.
func (s SeriesNA) Copy() Series {
	return s
}

// Series operations.

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s SeriesNA) Filter(mask any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesNA.Filter: invalid type %T", mask)}
	}
}

func (s SeriesNA) filterBool(mask SeriesBool) Series {
	elementCount := 0
	for _, v := range mask.data {
		if v {
			elementCount++
		}
	}

	s.size = elementCount
	return s
}

func (s SeriesNA) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
	s.size = mask.__trueCount()
	return s
}

func (s SeriesNA) filterBoolSlice(mask []bool) Series {
	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	s.size = elementCount
	return s
}

func (s SeriesNA) filterIntSlice(indexes []int, check bool) Series {
	// check if indexes are in range
	if check {
		for _, v := range indexes {
			if v < 0 || v >= s.size {
				return SeriesError{fmt.Sprintf("SeriesNA.Filter: index %d is out of range", v)}
			}
		}
	}

	s.size = len(indexes)
	return s
}

func (s SeriesNA) Map(f MapFunc) Series {
	return s
}

func (s SeriesNA) MapNull(f MapFuncNull) Series {
	return s
}

type SeriesNAPartition struct {
	partition map[int64][]int
}

func (gp *SeriesNAPartition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesNAPartition) getMap() map[int64][]int {
	return gp.partition
}

// Group the elements in the series.
func (s SeriesNA) group() Series {
	return s
}

func (s SeriesNA) GroupBy(gp SeriesPartition) Series {
	return s
}

func (s SeriesNA) UnGroup() Series {
	return s
}

func (s SeriesNA) GetPartition() SeriesPartition {
	return s.partition
}

// Sort interface.
func (s SeriesNA) Less(i, j int) bool {
	return false
}

func (s SeriesNA) equal(i, j int) bool {
	return false
}

func (s SeriesNA) Swap(i, j int) {}

func (s SeriesNA) Sort() Series {
	return s
}

func (s SeriesNA) SortRev() Series {
	return s
}
