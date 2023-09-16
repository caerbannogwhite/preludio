package gandalff

import (
	"fmt"
	"sort"
	"time"
	"typesys"
)

// SeriesTime represents a datetime series.
type SeriesTime struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	name       string
	data       []time.Time
	nullMask   []uint8
	partition  *SeriesTimePartition
}

// Get the element at index i as a string.
func (s SeriesTime) GetString(i int) string {
	if s.isNullable && s.nullMask[i>>3]&(1<<uint(i%8)) != 0 {
		return NULL_STRING
	}
	return s.data[i].String()
}

// Set the element at index i. The value v must be of type time.Time or NullableTime.
func (s SeriesTime) Set(i int, v any) Series {
	switch v := v.(type) {
	case time.Time:
		s.data[i] = v

	case NullableTime:
		if v.Valid {
			s.data[i] = v.Value
		} else {
			s.data[i] = time.Time{}
			s.nullMask[i/8] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesTime.Set: provided value %T is not compatible with type time.Time or NullableTime", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesTime) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess(s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces)
}

// Append appends a value or a slice of values to the series.
func (s SeriesTime) Append(v any) Series {
	switch v := v.(type) {
	case bool, []bool:
		return s.appendRaw(v)
	case NullableTime, []NullableTime:
		return s.appendNullable(v)
	case SeriesTime:
		return s.appendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesTime.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesTime) appendRaw(v any) Series {
	switch v := v.(type) {
	case time.Time:
		s.data = append(s.data, v)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []time.Time:
		s.data = append(s.data, v...)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesTime.Append: invalid type %T", v)}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s SeriesTime) appendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesTime.AppendNullable: series is not nullable"}
	}

	switch v := v.(type) {
	case NullableTime:
		s.data = append(s.data, v.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableTime:
		ssize := len(s.data)
		s.data = append(s.data, make([]time.Time, len(v))...)
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
		return SeriesError{fmt.Sprintf("SeriesTime.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s SeriesTime) appendSeries(other Series) Series {
	var ok bool
	var o SeriesTime
	if o, ok = other.(SeriesTime); !ok {
		return SeriesError{fmt.Sprintf("SeriesTime.AppendSeries: invalid type %T", other)}
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

func (s SeriesTime) Times() []time.Time {
	return s.data
}

func (s SeriesTime) Data() any {
	return s.data
}

// NullableData returns a slice of NullableTime.
func (s SeriesTime) DataAsNullable() any {
	data := make([]NullableTime, len(s.data))
	for i, v := range s.data {
		data[i] = NullableTime{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// StringData returns a slice of strings.
func (s SeriesTime) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = v.String()
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = v.String()
		}
	}
	return data
}

// Casts the series to a given type.
func (s SeriesTime) Cast(t typesys.BaseType, stringPool *StringPool) Series {

	switch t {
	case typesys.BoolType:
		return SeriesError{"SeriesTime.Cast: cannot cast to bool"}

	case typesys.Int32Type:
		data := make([]int32, len(s.data))
		for i, v := range s.data {
			data[i] = int32(v.UnixNano())
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
		data := make([]int64, len(s.data))
		for i, v := range s.data {
			data[i] = v.UnixNano()
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
		data := make([]float64, len(s.data))
		for i, v := range s.data {
			data[i] = float64(v.UnixNano())
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
			return SeriesError{"SeriesTime.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = stringPool.Put(NULL_STRING)
				} else {
					data[i] = stringPool.Put(v.String())
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = stringPool.Put(v.String())
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
		return SeriesError{fmt.Sprintf("SeriesTime.Cast: invalid type %s", t.ToString())}
	}
}

// Copy returns a copy of the series.
func (s SeriesTime) Copy() Series {
	data := make([]time.Time, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesTime{
		isGrouped:  s.isGrouped,
		isNullable: s.isNullable,
		data:       data,
		nullMask:   nullMask,
		partition:  s.partition,
	}
}

func (s SeriesTime) getDataPtr() *[]time.Time {
	return &s.data
}

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s SeriesTime) Filter(mask any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesTime.Filter: invalid type %T", mask)}
	}
}

func (s SeriesTime) filterBool(mask SeriesBool) Series {
	return s.filterBoolSlice(mask.data)
}

// Filters out the elements by the given mask series.
func (s SeriesTime) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
	if mask.Len() != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesTime.Filter: mask length (%d) does not match series length (%d)", mask.Len(), len(s.data))}
	}

	if mask.isNullable {
		return SeriesError{"SeriesTime.Filter: mask series cannot be nullable for this operation"}
	}

	return s.filterBoolSlice(mask.Data().([]bool))
}

func (s SeriesTime) filterBoolSlice(mask []bool) Series {
	if len(mask) != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesTime.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	data := make([]time.Time, elementCount)
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
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
					data[dstIdx] = s.data[srcIdx]
				} else {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
					data[dstIdx] = s.data[srcIdx]
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

func (s SeriesTime) filterIntSlice(indexes []int) Series {
	var nullMask []uint8

	size := len(indexes)
	data := make([]time.Time, size)

	if s.isNullable {
		nullMask = __binVecInit(size)
		for dstIdx, srcIdx := range indexes {
			if srcIdx%8 > dstIdx%8 {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))

			}
			data[dstIdx] = s.data[srcIdx]
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

func (s SeriesTime) Map(f GDLMapFunc, stringPool *StringPool) Series {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:
		data := make([]bool, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(bool)
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

		return SeriesInt32{
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
			return SeriesError{"SeriesTime.Map: StringPool is nil"}
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

	return SeriesError{fmt.Sprintf("SeriesTime.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

// A partition is trivially a vector of maps (or boolIndices in this case)
// Each element of the vector represent a sub-group (the default is 1,
// which means no sub-grouping).
// So is for the null group, which has the same size as the partition vector.
type SeriesTimePartition struct {
	partition map[int64][]int
}

func (gp *SeriesTimePartition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesTimePartition) getMap() map[int64][]int {
	return gp.partition
}

func (s SeriesTime) group() Series {

	// Define the worker callback
	worker := func(threadNum, start, end int, map_ map[int64][]int) {
		for i := start; i < end; i++ {
			map_[s.data[i].UnixNano()] = append(map_[s.data[i].UnixNano()], i)
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		for i := start; i < end; i++ {
			if s.IsNull(i) {
				(*nulls) = append((*nulls), i)
			} else {
				map_[s.data[i].UnixNano()] = append(map_[s.data[i].UnixNano()], i)
			}
		}
	}

	partition := SeriesTimePartition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, s.Len(), s.HasNull(),
			worker, workerNulls),
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s SeriesTime) GroupBy(partition SeriesPartition) Series {
	// collect all keys
	otherIndeces := partition.getMap()
	keys := make([]int64, len(otherIndeces))
	i := 0
	for k := range otherIndeces {
		keys[i] = k
		i++
	}

	// Define the worker callback
	worker := func(threadNum, start, end int, map_ map[int64][]int) {
		var newHash int64
		for _, h := range keys[start:end] { // keys is defined outside the function
			for _, index := range otherIndeces[h] { // otherIndeces is defined outside the function
				newHash = s.data[index].UnixNano() + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		var newHash int64
		for _, h := range keys[start:end] { // keys is defined outside the function
			for _, index := range otherIndeces[h] { // otherIndeces is defined outside the function
				if s.IsNull(index) {
					newHash = HASH_MAGIC_NUMBER_NULL + (h << 13) + (h >> 4)
				} else {
					newHash = s.data[index].UnixNano() + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	newPartition := SeriesTimePartition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(keys), s.HasNull(),
			worker, workerNulls),
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s SeriesTime) UnGroup() Series {
	s.isGrouped = false
	s.partition = nil
	return s
}

func (s SeriesTime) GetPartition() SeriesPartition {
	return s.partition
}

////////////////////////			SORTING OPERATIONS

func (s SeriesTime) Less(i, j int) bool {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			return false
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			return true
		}
	}
	return s.data[i].Compare(s.data[j]) < 0
}

func (s SeriesTime) equal(i, j int) bool {
	if s.isNullable {
		if (s.nullMask[i>>3] & (1 << uint(i%8))) > 0 {
			return (s.nullMask[j>>3] & (1 << uint(j%8))) > 0
		}
		if (s.nullMask[j>>3] & (1 << uint(j%8))) > 0 {
			return false
		}
	}

	return s.data[i] == s.data[j]
}

func (s SeriesTime) Swap(i, j int) {
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

func (s SeriesTime) Sort() Series {
	if s.sorted != SORTED_ASC {
		sort.Sort(s)
		s.sorted = SORTED_ASC
	}
	return s
}

func (s SeriesTime) SortRev() Series {
	if s.sorted != SORTED_DESC {
		sort.Sort(sort.Reverse(s))
		s.sorted = SORTED_DESC
	}
	return s
}
