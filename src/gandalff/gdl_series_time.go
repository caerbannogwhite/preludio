package gandalff

import (
	"fmt"
	"preludiometa"
	"sort"
	"time"
)

// SeriesTime represents a datetime series.
type SeriesTime struct {
	isNullable bool
	sorted     SeriesSortOrder
	data       []time.Time
	nullMask   []uint8
	partition  *SeriesTimePartition
	ctx        *Context
	timeFormat string
}

// Get the time format of the series.
func (s SeriesTime) GetTimeFormat() string {
	return s.timeFormat
}

// Set the time format of the series.
func (s SeriesTime) SetTimeFormat(format string) Series {
	s.timeFormat = format
	return s
}

// Get the element at index i as a string.
func (s SeriesTime) GetAsString(i int) string {
	if s.isNullable && s.nullMask[i>>3]&(1<<uint(i%8)) != 0 {
		return NULL_STRING
	}
	return s.data[i].Format(s.timeFormat)
}

// Set the element at index i. The value v must be of type time.Time or NullableTime.
func (s SeriesTime) Set(i int, v any) Series {
	if s.partition != nil {
		return SeriesError{"SeriesTime.Set: cannot set values on a grouped Series"}
	}

	switch v := v.(type) {
	case nil:
		s = s.MakeNullable().(SeriesTime)
		s.nullMask[i>>3] |= 1 << uint(i%8)

	case time.Time:
		s.data[i] = v

	case NullableTime:
		s = s.MakeNullable().(SeriesTime)
		if v.Valid {
			s.data[i] = v.Value
		} else {
			s.data[i] = time.Time{}
			s.nullMask[i/8] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesTime.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

// Return the underlying data as a slice of time.Time.
func (s SeriesTime) Times() []time.Time {
	return s.data
}

// Return the underlying data as a slice of NullableTime.
func (s SeriesTime) DataAsNullable() any {
	data := make([]NullableTime, len(s.data))
	for i, v := range s.data {
		data[i] = NullableTime{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// Return the underlying data as a slice of strings.
func (s SeriesTime) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = v.Format(s.timeFormat)
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = v.Format(s.timeFormat)
		}
	}
	return data
}

// Casts the series to a given type.
func (s SeriesTime) Cast(t preludiometa.BaseType) Series {
	switch t {
	case preludiometa.BoolType:
		return SeriesError{fmt.Sprintf("SeriesTime.Cast: cannot cast to %s", t.ToString())}

	case preludiometa.IntType:
		data := make([]int, len(s.data))
		for i, v := range s.data {
			data[i] = int(v.UnixNano())
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
			data[i] = v.UnixNano()
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
			data[i] = float64(v.UnixNano())
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
					data[i] = s.ctx.stringPool.Put(v.Format(s.timeFormat))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = s.ctx.stringPool.Put(v.Format(s.timeFormat))
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

	case preludiometa.TimeType:
		return s

	case preludiometa.DurationType:
		data := make([]time.Duration, len(s.data))
		for i, v := range s.data {
			data[i] = v.Sub(time.Time{})
		}

		return SeriesDuration{
			isNullable: s.isNullable,
			sorted:     s.sorted,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesTime.Cast: invalid type %T", t)}
	}
}

////////////////////////			GROUPING OPERATIONS

// A SeriesTimePartition is a partition of a SeriesTime.
// Each key is a hash of a bool value, and each value is a slice of indices
// of the original series that are set to that value.
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

	s.partition = &newPartition

	return s
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
