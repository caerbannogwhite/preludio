package gandalff

import (
	"fmt"
	"preludiometa"
	"sort"
	"time"
)

// SeriesInt64 represents a series of ints.
type SeriesInt64 struct {
	isNullable bool
	sorted     SeriesSortOrder
	data       []int64
	nullMask   []uint8
	partition  *SeriesInt64Partition
	ctx        *Context
}

// Get the element at index i as a string.
func (s SeriesInt64) GetAsString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return intToString(s.data[i])
}

// Set the element at index i. The value v can be any belonging to types:
// int8, int16, int, int, int64 and their nullable versions.
func (s SeriesInt64) Set(i int, v any) Series {
	if s.partition != nil {
		return SeriesError{"SeriesInt64.Set: cannot set values on a grouped Series"}
	}

	switch val := v.(type) {
	case nil:
		s = s.MakeNullable().(SeriesInt64)
		s.nullMask[i>>3] |= 1 << uint(i%8)

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
		s = s.MakeNullable().(SeriesInt64)
		if v.(NullableInt8).Valid {
			s.data[i] = int64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt16:
		s = s.MakeNullable().(SeriesInt64)
		if v.(NullableInt16).Valid {
			s.data[i] = int64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt:
		s = s.MakeNullable().(SeriesInt64)
		if v.(NullableInt).Valid {
			s.data[i] = int64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt64:
		s = s.MakeNullable().(SeriesInt64)
		if v.(NullableInt64).Valid {
			s.data[i] = val.Value
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesInt64.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

// Return the underlying data as a slice of int64.
func (s SeriesInt64) Int64s() []int64 {
	return s.data
}

// Return the underlying data as a slice of NullableInt64.
func (s SeriesInt64) DataAsNullable() any {
	data := make([]NullableInt64, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt64{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// Return the underlying data as a slice of strings.
func (s SeriesInt64) DataAsString() []string {
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
func (s SeriesInt64) Cast(t preludiometa.BaseType) Series {
	switch t {
	case preludiometa.BoolType:
		data := make([]bool, len(s.data))
		for i, v := range s.data {
			data[i] = v != 0
		}

		return SeriesBool{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.IntType:
		data := make([]int, len(s.data))
		for i, v := range s.data {
			data[i] = int(v)
		}

		return SeriesInt{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Int64Type:
		return s

	case preludiometa.Float64Type:
		data := make([]float64, len(s.data))
		for i, v := range s.data {
			data[i] = float64(v)
		}

		return SeriesFloat64{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.StringType:
		if s.ctx.stringPool == nil {
			return SeriesError{"SeriesInt64.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = s.ctx.stringPool.Put(NULL_STRING)
				} else {
					data[i] = s.ctx.stringPool.Put(intToString(v))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = s.ctx.stringPool.Put(intToString(v))
			}
		}

		return SeriesString{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.TimeType:
		data := make([]time.Time, len(s.data))
		for i, v := range s.data {
			data[i] = time.Unix(0, v)
		}

		return SeriesTime{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.DurationType:
		data := make([]time.Duration, len(s.data))
		for i, v := range s.data {
			data[i] = time.Duration(v)
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
		return SeriesError{fmt.Sprintf("SeriesInt64.Cast: invalid type %s", t.ToString())}
	}
}

////////////////////////			GROUPING OPERATIONS

// A SeriesInt64Partition is a partition of a SeriesInt64.
// Each key is a hash of a bool value, and each value is a slice of indices
// of the original series that are set to that value.
type SeriesInt64Partition struct {
	partition           map[int64][]int
	isDense             bool
	partitionDenseMin   int64
	partitionDense      [][]int
	partitionDenseNulls []int
}

func (gp *SeriesInt64Partition) getSize() int {
	if gp.isDense {
		if gp.partitionDenseNulls != nil && len(gp.partitionDenseNulls) > 0 {
			return len(gp.partitionDense) + 1
		}
		return len(gp.partitionDense)
	}
	return len(gp.partition)
}

func (gp *SeriesInt64Partition) getMap() map[int64][]int {
	if gp.isDense {
		map_ := make(map[int64][]int, len(gp.partitionDense))
		for i, part := range gp.partitionDense {
			map_[int64(i)+gp.partitionDenseMin] = part
		}

		// Merge the nulls to the map
		if gp.partitionDenseNulls != nil && len(gp.partitionDenseNulls) > 0 {
			nullKey := __series_get_nullkey(map_, HASH_NULL_KEY)
			map_[nullKey] = gp.partitionDenseNulls
		}

		return map_
	}

	return gp.partition
}

func (s SeriesInt64) group() Series {
	var useDenseMap bool
	var min, max int64
	var partition SeriesInt64Partition

	// If the number of elements is small,
	// look for the minimum and maximum values
	if len(s.data) < MINIMUM_PARALLEL_SIZE_2 {
		useDenseMap = true
		max = s.data[0]
		min = s.data[0]
		for _, v := range s.data {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
	}

	// If the difference between the maximum and minimum values is acceptable,
	// then we can use a dense map, otherwise we use a sparse map
	if useDenseMap && (max-min >= MINIMUM_PARALLEL_SIZE_1) {
		useDenseMap = false
	}

	// DENSE MAP
	if useDenseMap {
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
			partitionDenseMin:   min,
			partitionDense:      map_,
			partitionDenseNulls: nulls,
		}
	} else

	// SPARSE MAP
	{
		// Define the worker callback
		worker := func(threadNum, start, end int, map_ map[int64][]int) {
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

		// Define the worker callback for nulls
		workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
			for i := start; i < end; i++ {
				if s.IsNull(i) {
					(*nulls) = append((*nulls), i)
				} else {
					map_[s.data[i]] = append(map_[s.data[i]], i)
				}
			}
		}

		partition = SeriesInt64Partition{
			isDense: false,
			partition: __series_groupby(
				THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_2, len(s.data), s.HasNull(),
				worker, workerNulls),
		}
	}

	s.partition = &partition

	return s
}

func (s SeriesInt64) GroupBy(partition SeriesPartition) Series {
	if partition == nil {
		return s
	}

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
				newHash = s.data[index] + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
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
					newHash = s.data[index] + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	newPartition := SeriesInt64Partition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_2, len(keys), s.HasNull(),
			worker, workerNulls),
	}

	s.partition = &newPartition

	return s
}

////////////////////////			SORTING OPERATIONS

func (s SeriesInt64) Less(i, j int) bool {
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

func (s SeriesInt64) equal(i, j int) bool {
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

func (s SeriesInt64) Swap(i, j int) {
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

func (s SeriesInt64) Sort() Series {
	if s.sorted != SORTED_ASC {
		sort.Sort(s)
		s.sorted = SORTED_ASC
	}
	return s
}

func (s SeriesInt64) SortRev() Series {
	if s.sorted != SORTED_DESC {
		sort.Sort(sort.Reverse(s))
		s.sorted = SORTED_DESC
	}
	return s
}
