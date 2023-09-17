package gandalff

import (
	"fmt"
	"sort"
	"typesys"
	"unsafe"
)

// SeriesFloat64 represents a series of floats.
type SeriesFloat64 struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	name       string
	data       []float64
	nullMask   []uint8
	partition  *SeriesFloat64Partition
}

// Get the element at index i as a string.
func (s SeriesFloat64) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return floatToString(s.data[i])
}

// Set the element at index i. The value v can be any belonging to types:
// int8, int16, int, int32, int64, float32, float64 and their nullable versions.
func (s SeriesFloat64) Set(i int, v any) Series {
	switch val := v.(type) {
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
		s.MakeNullable()
		if v.(NullableInt8).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt16:
		s.MakeNullable()
		if v.(NullableInt16).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt32:
		s.MakeNullable()
		if v.(NullableInt32).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt64:
		s.MakeNullable()
		if v.(NullableInt64).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableFloat32:
		s.MakeNullable()
		if v.(NullableFloat32).Valid {
			s.data[i] = float64(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableFloat64:
		s.MakeNullable()
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

// Take the elements according to the given interval.
func (s SeriesFloat64) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess("SeriesFloat64", s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces, false)
}

// Append appends a value or a slice of values to the series.
func (s SeriesFloat64) Append(v any) Series {
	switch v := v.(type) {
	case float64:
		s.data = append(s.data, v)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []float64:
		s.data = append(s.data, v...)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	case NullableFloat64:
		s.data = append(s.data, v.Value)
		s.isNullable = true
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableFloat64:
		ssize := len(s.data)
		s.data = append(s.data, make([]float64, len(v))...)
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

	case SeriesFloat64:
		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, len(v.data), v.isNullable, v.nullMask)
		s.data = append(s.data, v.data...)

	default:
		return SeriesError{fmt.Sprintf("SeriesFloat64.Append: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

func (s SeriesFloat64) Float64s() []float64 {
	return s.data
}

func (s SeriesFloat64) Data() any {
	return s.data
}

func (s SeriesFloat64) DataAsNullable() any {
	data := make([]NullableFloat64, len(s.data))
	for i, v := range s.data {
		data[i] = NullableFloat64{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

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
func (s SeriesFloat64) Cast(t typesys.BaseType, stringPool *StringPool) Series {
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
			partition:  nil,
		}

	case typesys.Int32Type:
		data := make([]int32, len(s.data))
		for i, v := range s.data {
			data[i] = int32(v)
		}

		return SeriesInt32{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
		}

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
			partition:  nil,
		}

	case typesys.Float64Type:
		return s

	case typesys.StringType:
		if stringPool == nil {
			return SeriesError{"SeriesFloat64.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = stringPool.Put(NULL_STRING)
				} else {
					data[i] = stringPool.Put(floatToString(v))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = stringPool.Put(floatToString(v))
			}
		}

		return SeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       stringPool,
			partition:  nil,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesFloat64.Cast: invalid type %s", t.ToString())}
	}
}

func (s SeriesFloat64) Copy() Series {
	data := make([]float64, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesFloat64{
		isGrouped: s.isGrouped, sorted: s.sorted, isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

func (s SeriesFloat64) getDataPtr() *[]float64 {
	return &s.data
}

func (s SeriesFloat64) Map(f GDLMapFunc, stringPool *StringPool) Series {
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

		s.isGrouped = false
		s.sorted = SORTED_NONE
		s.data = data

		return s

	case string:
		if stringPool == nil {
			return SeriesError{"SeriesFloat64.Map: StringPool is nil"}
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

	return SeriesError{fmt.Sprintf("SeriesFloat64.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

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

	s.isGrouped = true
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

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s SeriesFloat64) UnGroup() Series {
	s.isGrouped = false
	s.partition = nil
	return s
}

func (s SeriesFloat64) GetPartition() SeriesPartition {
	return s.partition
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
