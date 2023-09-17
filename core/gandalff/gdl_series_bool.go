package gandalff

import (
	"fmt"
	"sort"
	"typesys"
)

// SeriesBool represents a series of bools.
// The data is stored as a byte array, with each bit representing a bool.
type SeriesBool struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	name       string
	data       []bool
	nullMask   []uint8
	partition  *SeriesBoolPartition
}

// Get the element at index i as a string.
func (s SeriesBool) GetString(i int) string {
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
	switch v := v.(type) {
	case bool:
		s.data[i] = v

	case NullableBool:
		s.MakeNullable()
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

// Take the elements according to the given interval.
func (s SeriesBool) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess("SeriesBool", s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces, false)
}

// Append appends a value or a slice of values to the series.
func (s SeriesBool) Append(v any) Series {
	switch v := v.(type) {
	case bool:
		s.data = append(s.data, v)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []bool:
		s.data = append(s.data, v...)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	case NullableBool:
		s.data = append(s.data, v.Value)
		s.isNullable = true
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableBool:
		ssize := len(s.data)
		s.data = append(s.data, make([]bool, len(v))...)
		s.isNullable = true
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for i, b := range v {
			s.data[ssize+i] = b.Value
			if !b.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}

	case SeriesBool:
		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, len(v.data), v.isNullable, v.nullMask)
		s.data = append(s.data, v.data...)

	default:
		return SeriesError{fmt.Sprintf("SeriesBool.Append: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

func (s SeriesBool) Bools() []bool {
	return s.data
}

func (s SeriesBool) Data() any {
	return s.data
}

// NullableData returns a slice of NullableBool.
func (s SeriesBool) DataAsNullable() any {
	data := make([]NullableBool, len(s.data))
	for i, v := range s.data {
		data[i] = NullableBool{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// StringData returns a slice of strings.
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

// Casts the series to a given type.
func (s SeriesBool) Cast(t typesys.BaseType, stringPool *StringPool) Series {
	switch t {
	case typesys.BoolType:
		return s

	case typesys.Int32Type:
		data := make([]int32, len(s.data))
		for i, v := range s.data {
			if v {
				data[i] = 1
			}
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
			if v {
				data[i] = 1
			}
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
			if v {
				data[i] = 1
			}
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
			return SeriesError{"SeriesBool.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = stringPool.Put(NULL_STRING)
				} else if v {
					data[i] = stringPool.Put(BOOL_TRUE_STRING)
				} else {
					data[i] = stringPool.Put(BOOL_FALSE_STRING)
				}
			}
		} else {
			for i, v := range s.data {
				if v {
					data[i] = stringPool.Put(BOOL_TRUE_STRING)
				} else {
					data[i] = stringPool.Put(BOOL_FALSE_STRING)
				}
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
		return SeriesError{fmt.Sprintf("SeriesBool.Cast: invalid type %s", t.ToString())}
	}
}

// Copy returns a copy of the series.
func (s SeriesBool) Copy() Series {
	data := make([]bool, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesBool{
		isGrouped:  s.isGrouped,
		isNullable: s.isNullable,
		data:       data,
		nullMask:   nullMask,
		partition:  s.partition,
	}
}

func (s SeriesBool) getDataPtr() *[]bool {
	return &s.data
}

func (s SeriesBool) Map(f GDLMapFunc, stringPool *StringPool) Series {
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
			return SeriesError{"SeriesBool.Map: StringPool is nil"}
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

	return SeriesError{fmt.Sprintf("SeriesBool.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

// A partition is trivially a vector of maps (or boolIndices in this case)
// Each element of the vector represent a sub-group (the default is 1,
// which means no sub-grouping).
// So is for the null group, which has the same size as the partition vector.
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

	s.isGrouped = true
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

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s SeriesBool) UnGroup() Series {
	s.isGrouped = false
	s.partition = nil
	return s
}

func (s SeriesBool) GetPartition() SeriesPartition {
	return s.partition
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
