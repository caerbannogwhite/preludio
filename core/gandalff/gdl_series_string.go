package gandalff

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"typesys"
	"unsafe"
)

// SeriesString represents a series of strings.
type SeriesString struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	data       []*string
	nullMask   []uint8
	pool       *StringPool
	partition  *SeriesStringPartition
}

// Get the element at index i as a string.
func (s SeriesString) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return *s.data[i]
}

// Set the element at index i. The value v must be of type string or NullableString.
func (s SeriesString) Set(i int, v any) Series {
	switch v := v.(type) {
	case string:
		s.data[i] = s.pool.Put(v)

	case NullableString:
		s.MakeNullable()
		if v.Valid {
			s.data[i] = s.pool.Put(v.Value)
		} else {
			s.data[i] = nil
			s.nullMask[i/8] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesString.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Append appends a value or a slice of values to the series.
func (s SeriesString) Append(v any) Series {
	switch v := v.(type) {
	case string:
		s.data = append(s.data, s.pool.Put(v))
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []string:
		s.data = append(s.data, make([]*string, len(v))...)
		for i, str := range v {
			s.data[len(s.data)-len(v)+i] = s.pool.Put(str)
		}
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	case NullableString:
		s.isNullable = true
		s.data = append(s.data, s.pool.Put(v.Value))
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableString:
		s.isNullable = true
		ssize := len(s.data)
		s.data = append(s.data, make([]*string, len(v))...)
		for i, b := range v {
			s.data[ssize+i] = s.pool.Put(b.Value)
			if !b.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}

	case SeriesString:
		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, len(v.data), v.isNullable, v.nullMask)
		s.data = append(s.data, make([]*string, len(v.data))...)
		for i := 0; i < len(v.data); i++ {
			s.data[len(s.data)-len(v.data)+i] = s.pool.Put(*v.data[i])
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesString.Append: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

func (s SeriesString) Strings() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return data
}

func (s SeriesString) DataAsNullable() any {
	data := make([]NullableString, len(s.data))
	for i, v := range s.data {
		data[i] = NullableString{Valid: !s.IsNull(i), Value: *v}
	}
	return data
}

func (s SeriesString) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = *v
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = *v
		}
	}
	return data
}

// Casts the series to a given type.
func (s SeriesString) Cast(t typesys.BaseType) Series {
	switch t {
	case typesys.BoolType:
		data := make([]bool, len(s.data))
		nullMask := __binVecInit(len(s.data), false)
		if s.isNullable {
			copy(nullMask, s.nullMask)
		}

		typeGuesser := newTypeGuesser()
		if s.isNullable {
			for i, v := range s.data {
				if !s.IsNull(i) {
					b, err := typeGuesser.atoBool(*v)
					if err != nil {
						nullMask[i>>3] |= (1 << uint(i%8))
					}
					data[i] = b
				}
			}
		} else {
			for i, v := range s.data {
				b, err := typeGuesser.atoBool(*v)
				if err != nil {
					nullMask[i>>3] |= (1 << uint(i%8))
				}
				data[i] = b
			}
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case typesys.Int32Type:
		data := make([]int32, len(s.data))
		nullMask := __binVecInit(len(s.data), false)
		if s.isNullable {
			copy(nullMask, s.nullMask)
		}

		if s.isNullable {
			for i, v := range s.data {
				if !s.IsNull(i) {
					d, err := strconv.Atoi(*v)
					if err != nil {
						nullMask[i>>3] |= (1 << uint(i%8))
					} else {
						data[i] = int32(d)
					}
				}
			}
		} else {
			for i, v := range s.data {
				d, err := strconv.Atoi(*v)
				if err != nil {
					nullMask[i>>3] |= (1 << uint(i%8))
				} else {
					data[i] = int32(d)
				}
			}
		}

		return SeriesInt32{
			isGrouped:  false,
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case typesys.Int64Type:
		data := make([]int64, len(s.data))
		nullMask := __binVecInit(len(s.data), false)
		if s.isNullable {
			copy(nullMask, s.nullMask)
		}

		if s.isNullable {
			for i, v := range s.data {
				if !s.IsNull(i) {
					d, err := strconv.ParseInt(*v, 10, 64)
					if err != nil {
						nullMask[i>>3] |= (1 << uint(i%8))
					} else {
						data[i] = d
					}
				}
			}
		} else {
			for i, v := range s.data {
				d, err := strconv.ParseInt(*v, 10, 64)
				if err != nil {
					nullMask[i>>3] |= (1 << uint(i%8))
				} else {
					data[i] = d
				}
			}
		}

		return SeriesInt64{
			isGrouped:  false,
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case typesys.Float64Type:
		data := make([]float64, len(s.data))
		nullMask := __binVecInit(len(s.data), false)
		if s.isNullable {
			copy(nullMask, s.nullMask)
		}

		if s.isNullable {
			for i, v := range s.data {
				if !s.IsNull(i) {
					f, err := strconv.ParseFloat(*v, 64)
					if err != nil {
						nullMask[i>>3] |= (1 << uint(i%8))
					} else {
						data[i] = f
					}
				}
			}
		} else {
			for i, v := range s.data {
				f, err := strconv.ParseFloat(*v, 64)
				if err != nil {
					nullMask[i>>3] |= (1 << uint(i%8))
				} else {
					data[i] = f
				}
			}
		}

		return SeriesFloat64{
			isGrouped:  false,
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case typesys.StringType:
		return s

	default:
		return SeriesError{fmt.Sprintf("SeriesString.Cast: invalid type %s", t.ToString())}
	}
}

////////////////////////			GROUPING OPERATIONS

type SeriesStringPartition struct {
	partition map[int64][]int
	pool      *StringPool
}

func (gp *SeriesStringPartition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesStringPartition) getMap() map[int64][]int {
	return gp.partition
}

func (s SeriesString) group() Series {

	// Define the worker callback
	worker := func(threadNum, start, end int, map_ map[int64][]int) {
		var ptr unsafe.Pointer
		for i := start; i < end; i++ {
			ptr = unsafe.Pointer(s.data[i])
			map_[(*(*int64)(unsafe.Pointer(&ptr)))] = append(map_[(*(*int64)(unsafe.Pointer(&ptr)))], i)
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		var ptr unsafe.Pointer
		for i := start; i < end; i++ {
			if s.IsNull(i) {
				(*nulls) = append((*nulls), i)
			} else {
				ptr = unsafe.Pointer(s.data[i])
				map_[(*(*int64)(unsafe.Pointer(&ptr)))] = append(map_[(*(*int64)(unsafe.Pointer(&ptr)))], i)
			}
		}
	}

	partition := SeriesStringPartition{
		pool: s.pool,
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(s.data), s.HasNull(),
			worker, workerNulls),
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s SeriesString) GroupBy(partition SeriesPartition) Series {
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
		var ptr unsafe.Pointer
		for _, h := range keys[start:end] { // keys is defined outside the function
			for _, index := range otherIndeces[h] { // otherIndeces is defined outside the function
				ptr = unsafe.Pointer(s.data[index])
				newHash = *(*int64)(unsafe.Pointer(&ptr)) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		var newHash int64
		var ptr unsafe.Pointer
		for _, h := range keys[start:end] { // keys is defined outside the function
			for _, index := range otherIndeces[h] { // otherIndeces is defined outside the function
				if s.IsNull(index) {
					newHash = HASH_MAGIC_NUMBER_NULL + (h << 13) + (h >> 4)
				} else {
					ptr = unsafe.Pointer(s.data[index])
					newHash = *(*int64)(unsafe.Pointer(&ptr)) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	newPartition := SeriesStringPartition{
		pool: s.pool,
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(keys), s.HasNull(),
			worker, workerNulls),
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

////////////////////////			SORTING OPERATIONS

func (s SeriesString) Less(i, j int) bool {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			return false
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			return true
		}
	}

	return (*s.data[i]) < (*s.data[j])
}

func (s SeriesString) equal(i, j int) bool {
	if s.isNullable {
		if (s.nullMask[i>>3] & (1 << uint(i%8))) > 0 {
			return (s.nullMask[j>>3] & (1 << uint(j%8))) > 0
		}
		if (s.nullMask[j>>3] & (1 << uint(j%8))) > 0 {
			return false
		}
	}

	return (*s.data[i]) == (*s.data[j])
}

func (s SeriesString) Swap(i, j int) {
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

func (s SeriesString) Sort() Series {
	if s.sorted != SORTED_ASC {
		sort.Sort(s)
		s.sorted = SORTED_ASC
	}
	return s
}

func (s SeriesString) SortRev() Series {
	if s.sorted != SORTED_DESC {
		sort.Sort(sort.Reverse(s))
		s.sorted = SORTED_DESC
	}
	return s
}

////////////////////////			STRING OPERATIONS

func (s SeriesString) ToUpper() Series {
	if s.isGrouped {
		return SeriesError{"SeriesString.ToUpper() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.pool.Put(strings.ToUpper(*s.data[i]))
	}

	return s
}

func (s SeriesString) ToLower() Series {
	if s.isGrouped {
		return SeriesError{"SeriesString.ToLower() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.pool.Put(strings.ToLower(*s.data[i]))
	}

	return s
}

func (s SeriesString) TrimSpace() Series {
	if s.isGrouped {
		return SeriesError{"SeriesString.TrimSpace() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.pool.Put(strings.TrimSpace(*s.data[i]))
	}

	return s
}

func (s SeriesString) Trim(cutset string) Series {
	if s.isGrouped {
		return SeriesError{"SeriesString.Trim() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.pool.Put(strings.Trim(*s.data[i], cutset))
	}

	return s
}

func (s SeriesString) Replace(old, new string, n int) Series {
	if s.isGrouped {
		return SeriesError{"SeriesString.Replace() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.pool.Put(strings.Replace(*s.data[i], old, new, n))
	}

	return s
}
