package gandalff

import (
	"fmt"
	"sort"
	"sync"
	"typesys"
)

// GDLSeriesInt64 represents a series of ints.
type GDLSeriesInt64 struct {
	isGrouped  bool
	isNullable bool
	sorted     GDLSeriesSortOrder
	name       string
	data       []int64
	nullMask   []uint8
	partition  *SeriesInt64Partition
}

func NewGDLSeriesInt64(name string, isNullable bool, makeCopy bool, data []int64) GDLSeries {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GDLSeriesInt64{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

////////////////////////			BASIC ACCESSORS

// Returns the number of elements in the series.
func (s GDLSeriesInt64) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s GDLSeriesInt64) Name() string {
	return s.name
}

// Returns the type of the series.
func (s GDLSeriesInt64) Type() typesys.BaseType {
	return typesys.Int32Type
}

// Returns if the series is grouped.
func (s GDLSeriesInt64) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s GDLSeriesInt64) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s GDLSeriesInt64) IsSorted() GDLSeriesSortOrder {
	return s.sorted
}

// Returns if the series has null values.
func (s GDLSeriesInt64) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s GDLSeriesInt64) NullCount() int {
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
func (s GDLSeriesInt64) NonNullCount() int {
	return s.Len() - s.NullCount()
}

// Returns if the element at index i is null.
func (s GDLSeriesInt64) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s GDLSeriesInt64) SetNull(i int) GDLSeries {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return s
	} else {
		nullMask := __binVecInit(len(s.data))
		nullMask[i/8] |= 1 << uint(i%8)

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s GDLSeriesInt64) GetNullMask() []bool {
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
func (s GDLSeriesInt64) SetNullMask(mask []bool) GDLSeries {
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
func (s GDLSeriesInt64) MakeNullable() GDLSeries {
	if !s.isNullable {
		s.nullMask = __binVecInit(len(s.data))
		s.isNullable = true
	}
	return s
}

// Get the element at index i.
func (s GDLSeriesInt64) Get(i int) any {
	return s.data[i]
}

// Get the element at index i as a string.
func (s GDLSeriesInt64) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return intToString(s.data[i])
}

// Set the element at index i. The value v can be of type int8, int16, int, int32, int64,
// NullableInt8, NullableInt16, NullableInt32, NullableInt64.
func (s GDLSeriesInt64) Set(i int, v any) GDLSeries {
	switch val := v.(type) {
	case int8:
		s.data[i] = int64(val)
	case int16:
		s.data[i] = int64(val)
	case int:
		s.data[i] = int64(val)
	case int32:
		s.data[i] = int64(val)
	case int64:
		s.data[i] = val
	case NullableInt8:
		if v.(NullableInt8).Valid {
			s.data[i] = int64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	case NullableInt16:
		if v.(NullableInt16).Valid {
			s.data[i] = int64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	case NullableInt32:
		if v.(NullableInt32).Valid {
			s.data[i] = int64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	case NullableInt64:
		if v.(NullableInt64).Valid {
			s.data[i] = val.Value
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.Set: provided value %t is not compatible with type int64 or NullableInt64", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s GDLSeriesInt64) Take(start, end, step int) GDLSeries {
	if start < 0 || start >= s.Len() || end < 0 || end > s.Len() || step == 0 {
		return NewGDLSeriesInt64(s.name, s.isNullable, false, []int64{})
	} else

	// reverse
	if step < 0 {
		return s
	} else

	// normal
	{
		size := end - start
		if size%step != 0 {
			size = size/step + 1
		} else {
			size = size / step
		}

		if s.isNullable {
			data := make([]int64, size)
			nullMask := __binVecInit(size)

			for i, j := start, 0; i < end; i, j = i+step, j+1 {
				data[j] = s.data[i]
				if s.IsNull(i) {
					nullMask[j>>3] |= 1 << uint(j%8)
				}
			}
			return GDLSeriesInt64{
				isGrouped:  false,
				isNullable: true,
				sorted:     SORTED_NONE,
				name:       s.name,
				data:       data,
				nullMask:   nullMask,
			}
		} else {
			data := make([]int64, size)
			for i, j := start, 0; i < end; i, j = i+step, j+1 {
				data[j] = s.data[i]
			}
			return GDLSeriesInt64{
				isGrouped:  false,
				isNullable: false,
				sorted:     SORTED_NONE,
				name:       s.name,
				data:       data,
				nullMask:   nil,
			}
		}
	}
}

func (s GDLSeriesInt64) Less(i, j int) bool {
	if s.isGrouped {
		if s.partition.indexToGroup[i] != s.partition.indexToGroup[j] {
			return s.partition.indexToGroup[i] < s.partition.indexToGroup[j]
		}
		return s.data[i] < s.data[j]
	} else

	// if s is grouped the null element are is the same group
	// so there is no need to check if the element is null
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			return false
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			return true
		}
	}

	return s.data[i] < s.data[j]
}

func (s GDLSeriesInt64) Swap(i, j int) {
	if s.isGrouped {
		s.partition.indexToGroup[i], s.partition.indexToGroup[j] = s.partition.indexToGroup[j], s.partition.indexToGroup[i]
	}

	if s.isNullable {
		// i is null, j is not null
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 && s.nullMask[j>>3]&(1<<uint(j%8)) == 0 {
			s.nullMask[i>>3] &= ^(1 << uint(i%8))
			s.nullMask[j>>3] |= 1 << uint(j%8)
		} else

		// i is not null, j is null
		if s.nullMask[i>>3]&(1<<uint(i%8)) == 0 && s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			s.nullMask[i>>3] |= 1 << uint(i%8)
			s.nullMask[j>>3] &= ^(1 << uint(j%8))
		}
	}

	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s GDLSeriesInt64) Append(v any) GDLSeries {
	switch v := v.(type) {
	case int64:
		return s.AppendRaw(v)
	case []int64:
		return s.AppendRaw(v)
	case NullableInt64:
		return s.AppendNullable(v)
	case []NullableInt64:
		return s.AppendNullable(v)
	case GDLSeriesInt64:
		return s.AppendSeries(v)
	case GDLSeriesError:
		return v
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesInt64) AppendRaw(v any) GDLSeries {
	if s.isNullable {
		if i, ok := v.(int64); ok {
			s.data = append(s.data, i)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if iv, ok := v.([]int64); ok {
			s.data = append(s.data, iv...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.AppendRaw: invalid type %T", v)}
		}
	} else {
		if b, ok := v.(int64); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]int64); ok {
			s.data = append(s.data, bv...)
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesInt64) AppendNullable(v any) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesInt64.AppendNullable: series is not nullable"}
	}

	if b, ok := v.(NullableInt64); ok {
		s.data = append(s.data, b.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !b.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableInt64); ok {
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
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s GDLSeriesInt64) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesInt64
	if o, ok = other.(GDLSeriesInt64); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.AppendSeries: invalid type %T", other)}
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
			s.nullMask = __binVecInit(len(s.data))
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

////////////////////////			ALL DATA ACCESSORS

func (s GDLSeriesInt64) Data() any {
	return s.data
}

func (s GDLSeriesInt64) DataAsNullable() any {
	data := make([]NullableInt64, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt64{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GDLSeriesInt64) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = intToString(v)
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = intToString(v)
		}
	}
	return data
}

// Casts the series to a given type.
func (s GDLSeriesInt64) Cast(t typesys.BaseType, stringPool *StringPool) GDLSeries {
	switch t {
	case typesys.BoolType:
		data := __binVecInit(len(s.data))
		for i, v := range s.data {
			if v != 0 {
				data[i>>3] |= (1 << uint(i%8))
			}
		}

		return GDLSeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			size:       len(s.data),
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case typesys.Int32Type:
		return s

	case typesys.Float64Type:
		data := make([]float64, len(s.data))
		for i, v := range s.data {
			data[i] = float64(v)
		}

		return GDLSeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case typesys.StringType:
		if stringPool == nil {
			return GDLSeriesError{"GDLSeriesInt64.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = stringPool.Put(NULL_STRING)
				} else {
					data[i] = stringPool.Put(intToString(v))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = stringPool.Put(intToString(v))
			}
		}

		return GDLSeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.Cast: invalid type %s", t.ToString())}
	}
}

func (s GDLSeriesInt64) Copy() GDLSeries {
	data := make([]int64, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return GDLSeriesInt64{isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

func (s GDLSeriesInt64) __getDataPtr() *[]int64 {
	return &s.data
}

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask series.
func (s GDLSeriesInt64) Filter(mask GDLSeriesBool) GDLSeries {
	if mask.size != s.Len() {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return GDLSeriesError{"GDLSeriesInt64.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()
	var nullMask []uint8

	data := make([]int64, elementCount)
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

// FilterByMask returns a new series with elements filtered by the mask.
func (s GDLSeriesInt64) FilterByMask(mask []bool) GDLSeries {
	if len(mask) != len(s.data) {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []int64
	var nullMask []uint8

	data = make([]int64, elementCount)

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

func (s GDLSeriesInt64) FilterByIndeces(indexes []int) GDLSeries {
	var data []int64
	var nullMask []uint8

	size := len(indexes)
	data = make([]int64, size)

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
		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}

func (s GDLSeriesInt64) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:

		var data []uint8
		if len(s.data)%8 == 0 {
			data = make([]uint8, (len(s.data) >> 3))
		} else {
			data = make([]uint8, (len(s.data)>>3)+1)
		}

		chunkLen := len(s.data) / THREADS_NUMBER
		if chunkLen < MINIMUM_PARALLEL_SIZE_2 {
			for i := 0; i < len(s.data); i++ {
				if f(s.data[i]).(bool) {
					data[i>>3] |= (1 << uint(i%8))
				}
			}
		} else {
			var wg sync.WaitGroup
			wg.Add(THREADS_NUMBER)

			for n := 0; n < THREADS_NUMBER; n++ {
				start := n * chunkLen
				end := (n + 1) * chunkLen
				if n == THREADS_NUMBER-1 {
					end = len(s.data)
				}

				go func() {
					for i := start; i < end; i++ {
						if f(s.data[i]).(bool) {
							data[i>>3] |= (1 << uint(i%8))
						}
					}
					wg.Done()
				}()
			}

			wg.Wait()
		}

		return GDLSeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			size:       len(s.data),
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case int:
		data := make([]int, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int)
		}

		return GDLSeriesInt32{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case int64:
		data := make([]int64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int64)
		}

		s.isGrouped = false
		s.sorted = SORTED_NONE
		s.data = data

		return s

	case float64:
		data := make([]float64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(float64)
		}

		return GDLSeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case string:
		if stringPool == nil {
			return GDLSeriesError{"GDLSeriesInt64.Map: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = stringPool.Put(f(s.data[i]).(string))
		}

		return GDLSeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       stringPool,
		}
	}

	return GDLSeriesError{fmt.Sprintf("GDLSeriesInt64.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

type SeriesInt64Partition struct {
	isDense             bool
	seriesSize          int
	partition           map[int64][]int
	nullKey             []int
	partitionDense      [][]int
	partitionDenseNulls []int
	indexToGroup        []int
}

func (gp SeriesInt64Partition) GetSize() int {
	if gp.isDense {
		nulls := 0
		if len(gp.partitionDenseNulls) > 0 {
			nulls = 1
		}
		return len(gp.partitionDense) + nulls
	}
	return len(gp.partition)
}

func (gp SeriesInt64Partition) beginSorting() SeriesInt64Partition {
	gp.indexToGroup = make([]int, gp.seriesSize)
	if gp.isDense {
		for i, part := range gp.partitionDense {
			for _, idx := range part {
				gp.indexToGroup[idx] = i
			}
		}

		for _, idx := range gp.partitionDenseNulls {
			gp.indexToGroup[idx] = len(gp.partitionDense)
		}
	} else {
		for i, part := range gp.partition {
			for _, idx := range part {
				gp.indexToGroup[idx] = int(i)
			}
		}
	}
	return gp
}

func (gp SeriesInt64Partition) endSorting() SeriesInt64Partition {
	if gp.isDense {
		newPartitionDense := make([][]int, len(gp.partitionDense))
		newPartitionDenseNulls := make([]int, len(gp.partitionDenseNulls))

		for _, part := range gp.partitionDense {
			newPartitionDense[gp.indexToGroup[part[0]]] = make([]int, len(part))
		}

		for i, idx := range gp.indexToGroup {
			if idx == len(gp.partitionDense) {
				newPartitionDenseNulls = append(newPartitionDenseNulls, i)
			} else {
				newPartitionDense[idx] = append(newPartitionDense[idx], i)
			}
		}

		gp.partitionDense = newPartitionDense
		gp.partitionDenseNulls = newPartitionDenseNulls
	} else {
		newPartition := make(map[int64][]int, len(gp.partition))

		for _, part := range gp.partition {
			newPartition[int64(gp.indexToGroup[part[0]])] = make([]int, len(part))
		}

		for i, idx := range gp.indexToGroup {
			newPartition[int64(idx)] = append(newPartition[int64(idx)], i)
		}

		gp.partition = newPartition
	}

	gp.indexToGroup = nil
	return gp
}

func (gp SeriesInt64Partition) GetMap() map[int64][]int {
	if gp.isDense {
		map_ := make(map[int64][]int, len(gp.partitionDense))
		for i, part := range gp.partitionDense {
			map_[int64(i)] = part
		}
		return map_
	}

	return gp.partition
}

func (gp SeriesInt64Partition) GetValueIndices(val any) []int {
	if val == nil {
		if gp.isDense {
			return gp.partitionDenseNulls
		} else if nulls, ok := gp.partition[HASH_NULL_KEY]; ok {
			return nulls
		}
	} else if v, ok := val.(int32); ok {
		if gp.isDense {
			return gp.partitionDense[v]
		} else if part, ok := gp.partition[int64(v)]; ok {
			return part
		}
	}

	return make([]int, 0)
}

func (gp SeriesInt64Partition) GetKeys() any {
	var keys []int
	if gp.isDense {
		keys = make([]int, 0, len(gp.partitionDense))
		for k, indeces := range gp.partitionDense {
			if len(indeces) > 0 {
				keys = append(keys, k)
			}
		}
	} else {
		keys = make([]int, 0, len(gp.partition))
		for k := range gp.partition {
			if k != HASH_NULL_KEY {
				keys = append(keys, int(k))
			}
		}
	}

	return keys
}

func (gp SeriesInt64Partition) InnerJoin(other SeriesInt64Partition) {

}

func (s GDLSeriesInt64) Group() GDLSeries {

	var partition SeriesInt64Partition
	if len(s.data) < MINIMUM_PARALLEL_SIZE_2 {
		max := s.data[0]
		min := s.data[0]
		for _, v := range s.data {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}

		var nulls []int
		map_ := make([][]int, max-min+1)
		for i := 0; i < len(map_); i++ {
			map_[i] = make([]int, 0, DEFAULT_DENSE_MAP_ARRAY_INITIAL_CAPACITY)
		}

		if s.HasNull() {
			nulls = make([]int, 0, DEFAULT_DENSE_MAP_ARRAY_INITIAL_CAPACITY)
			for i, v := range s.data {
				if s.IsNull(i) {
					nulls = append(nulls, i)
				} else {
					map_[v-min] = append(map_[v-min], i)
				}
			}
		} else {
			for i, v := range s.data {
				map_[v-min] = append(map_[v-min], i)
			}
		}

		partition = SeriesInt64Partition{
			isDense:             true,
			seriesSize:          s.Len(),
			partitionDense:      map_,
			partitionDenseNulls: nulls,
		}
	} else {
		if s.HasNull() {
			// Initialize the maps
			allMaps := make([]map[int64][]int, THREADS_NUMBER)
			for i := 0; i < THREADS_NUMBER; i++ {
				allMaps[i] = make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
			}

			allNulls := make([][]int, THREADS_NUMBER)

			// Define the worker callback
			worker := func(threadNum, start, end int) {
				map_ := allMaps[threadNum]
				for i := start; i < end; i++ {
					if s.IsNull(i) {
						allNulls[threadNum] = append(allNulls[threadNum], i)
					} else {
						map_[s.data[i]] = append(map_[s.data[i]], i)
					}
				}
			}

			__series_groupby_multithreaded(THREADS_NUMBER, len(s.data), allMaps, allNulls, worker)

			partition = SeriesInt64Partition{
				isDense:    false,
				seriesSize: s.Len(),
				partition:  allMaps[0],
			}
		} else {
			// Initialize the maps
			allMaps := make([]map[int64][]int, THREADS_NUMBER)
			for i := 0; i < THREADS_NUMBER; i++ {
				allMaps[i] = make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
			}

			// Define the worker callback
			worker := func(threadNum, start, end int) {
				map_ := allMaps[threadNum]
				up := end - ((end - start) % 8)
				for i := start; i < up; {
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
					map_[s.data[i]] = append(map_[s.data[i]], i)
					i++
				}

				for i := up; i < end; i++ {
					map_[s.data[i]] = append(map_[s.data[i]], i)
				}
			}

			__series_groupby_multithreaded(THREADS_NUMBER, len(s.data), allMaps, nil, worker)

			partition = SeriesInt64Partition{
				isDense:    false,
				seriesSize: s.Len(),
				partition:  allMaps[0],
			}
		}
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s GDLSeriesInt64) SubGroup(partition SeriesPartition) GDLSeries {
	var newPartition SeriesInt64Partition
	otherIndeces := partition.GetMap()

	if len(s.data) < MINIMUM_PARALLEL_SIZE_2 {

		map_ := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

		var newHash int64
		for h, v := range otherIndeces {
			for _, index := range v {
				newHash = s.data[index] + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}

		newPartition = SeriesInt64Partition{
			seriesSize: s.Len(),
			partition:  map_,
		}
	} else {

		// collect all keys
		keys := make([]int64, len(otherIndeces))
		i := 0
		for k := range otherIndeces {
			keys[i] = k
			i++
		}

		// Initialize the maps and the wait groups
		allMaps := make([]map[int64][]int, THREADS_NUMBER)
		for i := 0; i < THREADS_NUMBER; i++ {
			allMaps[i] = make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
		}

		// Define the worker callback
		worker := func(threadNum, start, end int) {
			var newHash int64
			map_ := allMaps[threadNum]
			for _, h := range keys[start:end] {
				for _, index := range otherIndeces[h] {
					newHash = s.data[index] + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
					map_[newHash] = append(map_[newHash], index)
				}
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(keys), allMaps, nil, worker)

		newPartition = SeriesInt64Partition{
			seriesSize: s.Len(),
			partition:  allMaps[0],
		}
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s GDLSeriesInt64) GetPartition() SeriesPartition {
	return s.partition
}

////////////////////////			SORTING OPERATIONS

func (s GDLSeriesInt64) Sort() GDLSeries {
	if s.sorted != SORTED_ASC {
		if s.isGrouped {
			*s.partition = (*s.partition).beginSorting()
			sort.Sort(s)
			*s.partition = (*s.partition).endSorting()
		} else {
			sort.Sort(s)
		}
		s.sorted = SORTED_ASC
	}
	return s
}

func (s GDLSeriesInt64) SortRev() GDLSeries {
	if s.sorted != SORTED_DESC {
		if s.isGrouped {
			*s.partition = (*s.partition).beginSorting()
			sort.Sort(sort.Reverse(s))
			*s.partition = (*s.partition).endSorting()
		} else {
			sort.Sort(sort.Reverse(s))
		}
		s.sorted = SORTED_DESC
	}
	return s
}

////////////////////////			ARITHMETIC OPERATIONS

func (s GDLSeriesInt64) Mul(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt64) Div(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt64) Mod(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt64) Add(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt64) Sub(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), other.Type().ToString())}

}
