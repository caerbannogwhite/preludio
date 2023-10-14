package gandalff

import (
	"fmt"
	"preludiometa"
	"sort"
	"time"
)

// SeriesDuration represents a duration series.
type SeriesDuration struct {
	isNullable bool
	sorted     SeriesSortOrder
	data       []time.Duration
	nullMask   []uint8
	partition  *SeriesDurationPartition
	ctx        *Context
}

// Get the element at index i as a string.
func (s SeriesDuration) GetAsString(i int) string {
	if s.isNullable && s.nullMask[i>>3]&(1<<uint(i%8)) != 0 {
		return NULL_STRING
	}
	return s.data[i].String()
}

// Set the element at index i. The value v must be of type time.Duration or NullableDuration.
func (s SeriesDuration) Set(i int, v any) Series {
	if s.partition != nil {
		return SeriesError{"SeriesDuration.Set: cannot set values on a grouped Series"}
	}

	switch v := v.(type) {
	case nil:
		s = s.MakeNullable().(SeriesDuration)
		s.nullMask[i>>3] |= 1 << uint(i%8)

	case time.Duration:
		s.data[i] = v

	case NullableDuration:
		s = s.MakeNullable().(SeriesDuration)
		if v.Valid {
			s.data[i] = v.Value
		} else {
			s.data[i] = time.Duration(0)
			s.nullMask[i/8] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesDuration.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

// Return the underlying data as a slice of time.Duration.
func (s SeriesDuration) Times() []time.Duration {
	return s.data
}

// Return the underlying data as a slice of NullableDuration.
func (s SeriesDuration) DataAsNullable() any {
	data := make([]NullableDuration, len(s.data))
	for i, v := range s.data {
		data[i] = NullableDuration{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// Return the underlying data as a slice of strings.
func (s SeriesDuration) DataAsString() []string {
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
func (s SeriesDuration) Cast(t preludiometa.BaseType) Series {

	switch t {
	case preludiometa.BoolType:
		return SeriesError{fmt.Sprintf("SeriesDuration.Cast: cannot cast to %s", t.ToString())}

	case preludiometa.IntType:
		data := make([]int, len(s.data))
		for i, v := range s.data {
			data[i] = int(v.Nanoseconds())
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
			data[i] = v.Nanoseconds()
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
			data[i] = float64(v.Nanoseconds())
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
				} else {
					data[i] = s.ctx.stringPool.Put(v.String())
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = s.ctx.stringPool.Put(v.String())
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
		return SeriesError{fmt.Sprintf("SeriesDuration.Cast: invalid type %T", t)}
	}
}

////////////////////////			GROUPING OPERATIONS

// A SeriesDurationPartition is a partition of a SeriesDuration.
// Each key is a hash of a bool value, and each value is a slice of indices
// of the original series that are set to that value.
type SeriesDurationPartition struct {
	partition map[int64][]int
}

func (gp *SeriesDurationPartition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesDurationPartition) getMap() map[int64][]int {
	return gp.partition
}

func (s SeriesDuration) group() Series {

	// Define the worker callback
	worker := func(threadNum, start, end int, map_ map[int64][]int) {
		for i := start; i < end; i++ {
			map_[s.data[i].Nanoseconds()] = append(map_[s.data[i].Nanoseconds()], i)
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		for i := start; i < end; i++ {
			if s.IsNull(i) {
				(*nulls) = append((*nulls), i)
			} else {
				map_[s.data[i].Nanoseconds()] = append(map_[s.data[i].Nanoseconds()], i)
			}
		}
	}

	partition := SeriesDurationPartition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, s.Len(), s.HasNull(),
			worker, workerNulls),
	}

	s.partition = &partition

	return s
}

func (s SeriesDuration) GroupBy(partition SeriesPartition) Series {
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
				newHash = s.data[index].Nanoseconds() + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
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
					newHash = s.data[index].Nanoseconds() + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	newPartition := SeriesDurationPartition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(keys), s.HasNull(),
			worker, workerNulls),
	}

	s.partition = &newPartition

	return s
}

////////////////////////			SORTING OPERATIONS

func (s SeriesDuration) Less(i, j int) bool {
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

func (s SeriesDuration) equal(i, j int) bool {
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

func (s SeriesDuration) Swap(i, j int) {
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

func (s SeriesDuration) Sort() Series {
	if s.sorted != SORTED_ASC {
		sort.Sort(s)
		s.sorted = SORTED_ASC
	}
	return s
}

func (s SeriesDuration) SortRev() Series {
	if s.sorted != SORTED_DESC {
		sort.Sort(sort.Reverse(s))
		s.sorted = SORTED_DESC
	}
	return s
}
