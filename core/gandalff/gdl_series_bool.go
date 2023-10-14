package gandalff

import (
	"fmt"
	"preludiometa"
	"sort"
)

// SeriesBool represents a series of bools.
// The data is stored as a byte array, with each bit representing a bool.
type SeriesBool struct {
	isNullable bool
	sorted     SeriesSortOrder
	data       []bool
	nullMask   []uint8
	partition  *SeriesBoolPartition
	ctx        *Context
}

// Get the element at index i as a string.
func (s SeriesBool) GetAsString(i int) string {
	if s.isNullable && s.nullMask[i>>3]&(1<<uint(i%8)) != 0 {
		return NULL_STRING
	} else if s.data[i] {
		return BOOL_TRUE_STRING
	} else {
		return BOOL_FALSE_STRING
	}
}

// Set the element at index i. The value must be of type bool or NullableBool.
func (s SeriesBool) Set(i int, v any) Series {
	if s.partition != nil {
		return SeriesError{"SeriesBool.Set: cannot set values in a grouped series"}
	}

	switch v := v.(type) {
	case nil:
		s = s.MakeNullable().(SeriesBool)
		s.nullMask[i>>3] |= 1 << uint(i%8)

	case bool:
		s.data[i] = v

	case NullableBool:
		s = s.MakeNullable().(SeriesBool)
		if v.Valid {
			s.data[i] = v.Value
		} else {
			s.nullMask[i>>3] |= 1 << uint(i%8)
			s.data[i] = false
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesBool.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

// Return the underlying data as a slice of bools.
func (s SeriesBool) Bools() []bool {
	return s.data
}

// Return the underlying data as a slice of NullableBool.
func (s SeriesBool) DataAsNullable() any {
	data := make([]NullableBool, len(s.data))
	for i, v := range s.data {
		data[i] = NullableBool{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// Return the data as a slice of strings.
func (s SeriesBool) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else if v {
				data[i] = BOOL_TRUE_STRING
			} else {
				data[i] = BOOL_FALSE_STRING
			}
		}
	} else {
		for i, v := range s.data {
			if v {
				data[i] = BOOL_TRUE_STRING
			} else {
				data[i] = BOOL_FALSE_STRING
			}
		}
	}
	return data
}

// Cast the series to a given type.
func (s SeriesBool) Cast(t preludiometa.BaseType) Series {
	switch t {
	case preludiometa.BoolType:
		return s

	case preludiometa.IntType:
		data := make([]int, len(s.data))
		for i, v := range s.data {
			if v {
				data[i] = 1
			}
		}

		return SeriesInt{
			isNullable: s.isNullable,
			sorted:     s.sorted,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Int64Type:
		data := make([]int64, len(s.data))
		for i, v := range s.data {
			if v {
				data[i] = 1
			}
		}

		return SeriesInt64{
			isNullable: s.isNullable,
			sorted:     s.sorted,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Float64Type:
		data := make([]float64, len(s.data))
		for i, v := range s.data {
			if v {
				data[i] = 1
			}
		}

		return SeriesFloat64{
			isNullable: s.isNullable,
			sorted:     s.sorted,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.StringType:
		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = s.ctx.stringPool.Put(NULL_STRING)
				} else if v {
					data[i] = s.ctx.stringPool.Put(BOOL_TRUE_STRING)
				} else {
					data[i] = s.ctx.stringPool.Put(BOOL_FALSE_STRING)
				}
			}
		} else {
			for i, v := range s.data {
				if v {
					data[i] = s.ctx.stringPool.Put(BOOL_TRUE_STRING)
				} else {
					data[i] = s.ctx.stringPool.Put(BOOL_FALSE_STRING)
				}
			}
		}

		return SeriesString{
			isNullable: s.isNullable,
			sorted:     s.sorted,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesBool.Cast: invalid type %s", t.ToString())}
	}
}

////////////////////////			GROUPING OPERATIONS

// A SeriesBoolPartition is a partition of a SeriesBool.
// Each key is a hash of a bool value, and each value is a slice of indices
// of the original series that are set to that value.
type SeriesBoolPartition struct {
	partition map[int64][]int
}

func (gp *SeriesBoolPartition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesBoolPartition) getMap() map[int64][]int {
	return gp.partition
}

func (s SeriesBool) group() Series {

	// Define the worker callback
	worker := func(threadNum, start, end int, map_ map[int64][]int) {
		for i := start; i < end; i++ {
			if s.data[i] {
				map_[1] = append(map_[1], i)
			} else {
				map_[0] = append(map_[0], i)
			}
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		for i := start; i < end; i++ {
			if s.IsNull(i) {
				(*nulls) = append((*nulls), i)
			} else if s.data[i] {
				map_[1] = append(map_[1], i)
			} else {
				map_[0] = append(map_[0], i)
			}

		}
	}

	partition := SeriesBoolPartition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, s.Len(), s.HasNull(),
			worker, workerNulls),
	}

	s.partition = &partition

	return s
}

func (s SeriesBool) GroupBy(partition SeriesPartition) Series {
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
				if s.data[index] {
					newHash = (1 + HASH_MAGIC_NUMBER) + (h << 13) + (h >> 4)
				} else {
					newHash = HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
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
				} else if s.data[index] {
					newHash = (1 + HASH_MAGIC_NUMBER) + (h << 13) + (h >> 4)
				} else {
					newHash = HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	newPartition := SeriesBoolPartition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(keys), s.HasNull(),
			worker, workerNulls),
	}

	s.partition = &newPartition

	return s
}

////////////////////////			SORTING OPERATIONS

func (s SeriesBool) Less(i, j int) bool {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			return false
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			return true
		}
	}
	return !s.data[i] && s.data[j]
}

func (s SeriesBool) equal(i, j int) bool {
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

func (s SeriesBool) Swap(i, j int) {
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

func (s SeriesBool) Sort() Series {
	if s.sorted != SORTED_ASC {
		sort.Sort(s)
		s.sorted = SORTED_ASC
	}
	return s
}

func (s SeriesBool) SortRev() Series {
	if s.sorted != SORTED_DESC {
		sort.Sort(sort.Reverse(s))
		s.sorted = SORTED_DESC
	}
	return s
}
