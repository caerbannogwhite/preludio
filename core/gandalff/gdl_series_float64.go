package gandalff

import (
	"fmt"
	"preludiometa"
	"sort"
	"time"
	"unsafe"
)

// SeriesFloat64 represents a series of floats.
type SeriesFloat64 struct {
	isNullable bool
	sorted     SeriesSortOrder
	data       []float64
	nullMask   []uint8
	partition  *SeriesFloat64Partition
	ctx        *Context
}

// Get the element at index i as a string.
func (s SeriesFloat64) GetAsString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return floatToString(s.data[i])
}

// Set the element at index i. The value v can be any belonging to types:
// int8, int16, int, int, int64, float32, float64 and their nullable versions.
func (s SeriesFloat64) Set(i int, v any) Series {
	if s.partition != nil {
		return SeriesError{"SeriesFloat64.Set: cannot set values in a grouped series"}
	}

	switch val := v.(type) {
	case nil:
		s = s.MakeNullable().(SeriesFloat64)
		s.nullMask[i>>3] |= 1 << uint(i%8)

	case int8:
		s.data[i] = float64(val)

	case int16:
		s.data[i] = float64(val)

	case int:
		s.data[i] = float64(val)

	case int32:
		s.data[i] = float64(val)

	case int64:
		s.data[i] = float64(val)

	case float32:
		s.data[i] = float64(val)

	case float64:
		s.data[i] = val

	case NullableInt8:
		s = s.MakeNullable().(SeriesFloat64)
		if v.(NullableInt8).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt16:
		s = s.MakeNullable().(SeriesFloat64)
		if v.(NullableInt16).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt:
		s = s.MakeNullable().(SeriesFloat64)
		if v.(NullableInt).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt64:
		s = s.MakeNullable().(SeriesFloat64)
		if v.(NullableInt64).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableFloat32:
		s = s.MakeNullable().(SeriesFloat64)
		if v.(NullableFloat32).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableFloat64:
		s = s.MakeNullable().(SeriesFloat64)
		if v.(NullableFloat64).Valid {
			s.data[i] = val.Value
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesFloat64.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

// Return the underlying data as a slice of float64.
func (s SeriesFloat64) Float64s() []float64 {
	return s.data
}

// Return the underlying data as a slice of NullableFloat64.
func (s SeriesFloat64) DataAsNullable() any {
	data := make([]NullableFloat64, len(s.data))
	for i, v := range s.data {
		data[i] = NullableFloat64{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// Return the underlying data as a slice of strings.
func (s SeriesFloat64) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = floatToString(v)
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = floatToString(v)
		}
	}
	return data
}

// Casts the series to a given type.
func (s SeriesFloat64) Cast(t preludiometa.BaseType) Series {
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
		data := make([]int64, len(s.data))
		for i, v := range s.data {
			data[i] = int64(v)
		}

		return SeriesInt64{
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Float64Type:
		return s

	case preludiometa.StringType:
		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = s.ctx.stringPool.Put(NULL_STRING)
				} else {
					data[i] = s.ctx.stringPool.Put(floatToString(v))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = s.ctx.stringPool.Put(floatToString(v))
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
			data[i] = time.Unix(0, int64(v))
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
		return SeriesError{fmt.Sprintf("SeriesFloat64.Cast: invalid type %s", t.ToString())}
	}
}

////////////////////////			GROUPING OPERATIONS

// A SeriesFloat64Partition is a partition of a SeriesFloat64.
// Each key is a hash of a bool value, and each value is a slice of indices
// of the original series that are set to that value.
type SeriesFloat64Partition struct {
	partition    map[int64][]int
	indexToGroup []int
}

func (gp *SeriesFloat64Partition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesFloat64Partition) getMap() map[int64][]int {
	return gp.partition
}

func (s SeriesFloat64) group() Series {

	// Define the worker callback
	worker := func(threadNum, start, end int, map_ map[int64][]int) {
		for i := start; i < end; i++ {
			map_[*(*int64)(unsafe.Pointer((&s.data[i])))] = append(map_[*(*int64)(unsafe.Pointer((&s.data[i])))], i)
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		for i := start; i < end; i++ {
			if s.IsNull(i) {
				(*nulls) = append((*nulls), i)
			} else {
				map_[*(*int64)(unsafe.Pointer((&s.data[i])))] = append(map_[*(*int64)(unsafe.Pointer((&s.data[i])))], i)
			}
		}
	}

	partition := SeriesFloat64Partition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_2, len(s.data), s.HasNull(),
			worker, workerNulls),
	}

	s.partition = &partition

	return s
}

func (s SeriesFloat64) GroupBy(partition SeriesPartition) Series {
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
				newHash = *(*int64)(unsafe.Pointer((&(s.data)[index]))) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
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
					newHash = *(*int64)(unsafe.Pointer((&(s.data)[index]))) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	newPartition := SeriesFloat64Partition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(keys), s.HasNull(),
			worker, workerNulls),
	}

	s.partition = &newPartition

	return s
}

////////////////////////			SORTING OPERATIONS

func (s SeriesFloat64) Less(i, j int) bool {
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

func (s SeriesFloat64) equal(i, j int) bool {
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

func (s SeriesFloat64) Swap(i, j int) {
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

func (s SeriesFloat64) Sort() Series {
	if s.sorted != SORTED_ASC {
		sort.Sort(s)
		s.sorted = SORTED_ASC
	}
	return s
}

func (s SeriesFloat64) SortRev() Series {
	if s.sorted != SORTED_DESC {
		sort.Sort(sort.Reverse(s))
		s.sorted = SORTED_DESC
	}
	return s
}
