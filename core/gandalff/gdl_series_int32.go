package gandalff

import (
	"fmt"
	"sort"
	"sync"
	"typesys"
)

// SeriesInt32 represents a series of ints.
type SeriesInt32 struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	name       string
	data       []int32
	nullMask   []uint8
	partition  *SeriesInt32Partition
}

////////////////////////			BASIC ACCESSORS

// Returns the number of elements in the series.
func (s SeriesInt32) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s SeriesInt32) Name() string {
	return s.name
}

// Sets the name of the series.
func (s SeriesInt32) SetName(name string) Series {
	s.name = name
	return s
}

// Returns the type of the series.
func (s SeriesInt32) Type() typesys.BaseType {
	return typesys.Int32Type
}

// Returns the type and cardinality of the series.
func (s SeriesInt32) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{typesys.Int32Type, s.Len()}
}

// Returns if the series is grouped.
func (s SeriesInt32) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s SeriesInt32) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s SeriesInt32) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series is error.
func (s SeriesInt32) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s SeriesInt32) GetError() string {
	return ""
}

// Returns if the series has null values.
func (s SeriesInt32) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s SeriesInt32) NullCount() int {
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
func (s SeriesInt32) NonNullCount() int {
	return s.Len() - s.NullCount()
}

// Returns if the element at index i is null.
func (s SeriesInt32) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s SeriesInt32) SetNull(i int) Series {
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
func (s SeriesInt32) GetNullMask() []bool {
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
func (s SeriesInt32) SetNullMask(mask []bool) Series {
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
func (s SeriesInt32) MakeNullable() Series {
	if !s.isNullable {
		s.nullMask = __binVecInit(len(s.data))
		s.isNullable = true
	}
	return s
}

// Get the element at index i.
func (s SeriesInt32) Get(i int) any {
	return s.data[i]
}

// Get the element at index i as a string.
func (s SeriesInt32) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return intToString(int64(s.data[i]))
}

// Set the element at index i. The value v can be of type int8, int16, int, int32,
// NullableInt8, NullableInt16, NullableInt32.
func (s SeriesInt32) Set(i int, v any) Series {
	switch val := v.(type) {
	case int8:
		s.data[i] = int32(val)
	case int16:
		s.data[i] = int32(val)
	case int:
		s.data[i] = int32(val)
	case int32:
		s.data[i] = int32(val)
	case NullableInt8:
		if v.(NullableInt8).Valid {
			s.data[i] = int32(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	case NullableInt16:
		if v.(NullableInt16).Valid {
			s.data[i] = int32(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	case NullableInt32:
		if v.(NullableInt32).Valid {
			s.data[i] = int32(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	default:
		return SeriesError{fmt.Sprintf("SeriesInt64.Set: provided value %t is not compatible with type int32 or NullableInt32", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesInt32) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess(s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces)
}

func (s SeriesInt32) Less(i, j int) bool {
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

func (s SeriesInt32) Swap(i, j int) {
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

func (s SeriesInt32) Append(v any) Series {
	switch v := v.(type) {
	case int32, []int32:
		return s.appendRaw(v)
	case NullableInt32, []NullableInt32:
		return s.appendNullable(v)
	case SeriesInt32:
		return s.appendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesInt32.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesInt32) appendRaw(v any) Series {
	switch v := v.(type) {
	case int32:
		s.data = append(s.data, v)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []int32:
		s.data = append(s.data, v...)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesInt32.Append: invalid type %T", v)}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s SeriesInt32) appendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesInt32.AppendNullable: series is not nullable"}
	}

	switch v := v.(type) {
	case NullableInt32:
		s.data = append(s.data, v.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableInt32:
		ssize := len(s.data)
		s.data = append(s.data, make([]int32, len(v))...)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for i, b := range v {
			s.data[ssize+i] = b.Value
			if !b.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesInt32.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s SeriesInt32) appendSeries(other Series) Series {
	var ok bool
	var o SeriesInt32
	if o, ok = other.(SeriesInt32); !ok {
		return SeriesError{fmt.Sprintf("SeriesInt32.AppendSeries: invalid type %T", other)}
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

func (s SeriesInt32) Int32s() []int32 {
	return s.data
}

func (s SeriesInt32) Data() any {
	return s.data
}

func (s SeriesInt32) DataAsNullable() any {
	data := make([]NullableInt32, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt32{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s SeriesInt32) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = intToString(int64(v))
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = intToString(int64(v))
		}
	}
	return data
}

// Casts the series to a given type.
func (s SeriesInt32) Cast(t typesys.BaseType, stringPool *StringPool) Series {
	switch t {
	case typesys.BoolType:
		data := make([]bool, len(s.data))
		for i, v := range s.data {
			data[i] = v != 0
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case typesys.Int32Type:
		return s

	case typesys.Int64Type:
		data := make([]int64, len(s.data))
		for i, v := range s.data {
			data[i] = int64(v)
		}

		return SeriesInt64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case typesys.Float64Type:
		data := make([]float64, len(s.data))
		for i, v := range s.data {
			data[i] = float64(v)
		}

		return SeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case typesys.StringType:
		if stringPool == nil {
			return SeriesError{"SeriesInt32.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = stringPool.Put(NULL_STRING)
				} else {
					data[i] = stringPool.Put(intToString(int64(v)))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = stringPool.Put(intToString(int64(v)))
			}
		}

		return SeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesInt32.Cast: invalid type %s", t.ToString())}
	}
}

func (s SeriesInt32) Copy() Series {
	data := make([]int32, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesInt32{isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

func (s SeriesInt32) getDataPtr() *[]int32 {
	return &s.data
}

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s SeriesInt32) Filter(mask any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesInt32.Filter: invalid type %T", mask)}
	}
}

func (s SeriesInt32) filterBool(mask SeriesBool) Series {
	return s.filterBoolSlice(mask.data)
}

func (s SeriesInt32) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
	if mask.size != s.Len() {
		return SeriesError{fmt.Sprintf("SeriesInt32.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return SeriesError{"SeriesInt32.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()
	var nullMask []uint8

	data := make([]int32, elementCount)
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

func (s SeriesInt32) filterBoolSlice(mask []bool) Series {
	if len(mask) != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesInt32.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []int32
	var nullMask []uint8

	data = make([]int32, elementCount)

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

func (s SeriesInt32) filterIntSlice(indexes []int) Series {
	var data []int32
	var nullMask []uint8

	size := len(indexes)
	data = make([]int32, size)

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

func (s SeriesInt32) Map(f GDLMapFunc, stringPool *StringPool) Series {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:

		data := make([]bool, len(s.data))
		chunkLen := len(s.data) / THREADS_NUMBER
		if chunkLen < MINIMUM_PARALLEL_SIZE_2 {
			for i := 0; i < len(s.data); i++ {
				data[i] = f(s.data[i]).(bool)
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
						data[i] = f(s.data[i]).(bool)
					}
					wg.Done()
				}()
			}

			wg.Wait()
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case int32:
		data := make([]int32, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int32)
		}

		s.isGrouped = false
		s.sorted = SORTED_NONE
		s.data = data

		return s

	case int64:
		data := make([]int64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int64)
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
		data := make([]float64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(float64)
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
			return SeriesError{"SeriesInt32.Map: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = stringPool.Put(f(s.data[i]).(string))
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

	return SeriesError{fmt.Sprintf("SeriesInt32.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

type SeriesInt32Partition struct {
	isDense             bool
	seriesSize          int
	partition           map[int64][]int
	partitionDenseMin   int32
	partitionDense      [][]int
	partitionDenseNulls []int
	indexToGroup        []int
}

func (gp SeriesInt32Partition) GetSize() int {
	if gp.isDense {
		nulls := 0
		if len(gp.partitionDenseNulls) > 0 {
			nulls = 1
		}
		return len(gp.partitionDense) + nulls
	}
	return len(gp.partition)
}

func (gp SeriesInt32Partition) beginSorting() SeriesInt32Partition {
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

func (gp SeriesInt32Partition) endSorting() SeriesInt32Partition {
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

func (gp SeriesInt32Partition) GetMap() map[int64][]int {
	if gp.isDense {
		map_ := make(map[int64][]int, len(gp.partitionDense))
		for i, part := range gp.partitionDense {
			map_[int64(i+int(gp.partitionDenseMin))] = part
		}
		return map_
	}

	return gp.partition
}

func (gp SeriesInt32Partition) GetValueIndices(val any) []int {
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

func (gp SeriesInt32Partition) GetKeys() any {
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

func (gp SeriesInt32Partition) debugPrint() {
	fmt.Println("SeriesInt32Partition")
	map_ := gp.GetMap()
	for k, v := range map_ {
		fmt.Printf("%4d: %v\n", k, v)
	}
}

func (s SeriesInt32) Group() Series {
	var partition SeriesInt32Partition
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

		map_ := make([][]int, max-min+1)
		for i := 0; i < len(map_); i++ {
			map_[i] = make([]int, 0, DEFAULT_DENSE_MAP_ARRAY_INITIAL_CAPACITY)
		}

		for i, v := range s.data {
			map_[v-min] = append(map_[v-min], i)
		}

		partition = SeriesInt32Partition{
			isDense:           true,
			seriesSize:        s.Len(),
			partitionDenseMin: min,
			partitionDense:    map_,
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
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
			}

			for i := up; i < end; i++ {
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(s.data), allMaps, nil, worker)

		partition = SeriesInt32Partition{
			isDense:    false,
			seriesSize: s.Len(),
			partition:  allMaps[0],
		}
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s SeriesInt32) SubGroup(partition SeriesPartition) Series {
	var newPartition SeriesInt32Partition
	otherIndeces := partition.GetMap()

	if len(s.data) < MINIMUM_PARALLEL_SIZE_2 {

		map_ := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

		var newHash int64
		for h, v := range otherIndeces {
			for _, index := range v {
				newHash = int64(s.data[index]) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}

		newPartition = SeriesInt32Partition{
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
					newHash = int64(s.data[index]) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
					map_[newHash] = append(map_[newHash], index)
				}
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(keys), allMaps, nil, worker)

		newPartition = SeriesInt32Partition{
			seriesSize: s.Len(),
			partition:  allMaps[0],
		}
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s SeriesInt32) GetPartition() SeriesPartition {
	return s.partition
}

////////////////////////			SORTING OPERATIONS

func (s SeriesInt32) Sort() Series {
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

func (s SeriesInt32) SortRev() Series {
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
