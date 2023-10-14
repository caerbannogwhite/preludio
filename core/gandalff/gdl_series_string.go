package gandalff

import (
	"fmt"
	"preludiometa"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// SeriesString represents a series of strings.
type SeriesString struct {
	isNullable bool
	sorted     SeriesSortOrder
	data       []*string
	nullMask   []uint8
	partition  *SeriesStringPartition
	ctx        *Context
}

// Get the element at index i as a string.
func (s SeriesString) GetAsString(i int) string {
	return *s.data[i]
}

// Set the element at index i. The value v must be of type string or NullableString.
func (s SeriesString) Set(i int, v any) Series {
	if s.partition != nil {
		return SeriesError{"SeriesString.Set: cannot set values on a grouped Series"}
	}

	switch v := v.(type) {
	case nil:
		s = s.MakeNullable().(SeriesString)
		s.data[i] = s.ctx.stringPool.nullStringPtr
		s.nullMask[i>>3] |= 1 << uint(i%8)

	case string:
		s.data[i] = s.ctx.stringPool.Put(v)

	case NullableString:
		s = s.MakeNullable().(SeriesString)
		if v.Valid {
			s.data[i] = s.ctx.stringPool.Put(v.Value)
		} else {
			s.data[i] = s.ctx.stringPool.nullStringPtr
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesString.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

// Return the underlying data as a slice of string.
func (s SeriesString) Strings() []string {
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

// Return the underlying data as a slice of NullableString.
func (s SeriesString) DataAsNullable() any {
	data := make([]NullableString, len(s.data))
	for i, v := range s.data {
		data[i] = NullableString{Valid: !s.IsNull(i), Value: *v}
	}
	return data
}

// Return the underlying data as a slice of string.
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
func (s SeriesString) Cast(t preludiometa.BaseType) Series {
	switch t {
	case preludiometa.BoolType:
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
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.IntType:
		data := make([]int, len(s.data))
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
						data[i] = int(d)
					}
				}
			}
		} else {
			for i, v := range s.data {
				d, err := strconv.Atoi(*v)
				if err != nil {
					nullMask[i>>3] |= (1 << uint(i%8))
				} else {
					data[i] = int(d)
				}
			}
		}

		return SeriesInt{
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Int64Type:
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
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.Float64Type:
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
			isNullable: true,
			sorted:     SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case preludiometa.StringType:
		return s

	case preludiometa.TimeType:
		return SeriesError{"SeriesString.Cast: cannot cast to Time, use SeriesString.ParseTime(layout) instead"}

	default:
		return SeriesError{fmt.Sprintf("SeriesString.Cast: invalid type %s", t.ToString())}
	}
}

// Parse the series as a time series.
func (s SeriesString) ParseTime(layout string) Series {
	data := make([]time.Time, len(s.data))
	nullMask := __binVecInit(len(s.data), false)
	if s.isNullable {
		copy(nullMask, s.nullMask)
	}

	for i, v := range s.data {
		if s.isNullable && s.IsNull(i) {
			continue
		}
		t, err := time.Parse(layout, *v)
		if err != nil {
			nullMask[i>>3] |= (1 << uint(i%8))
		} else {
			data[i] = t
		}
	}

	return SeriesTime{
		isNullable: true,
		sorted:     SORTED_NONE,
		data:       data,
		nullMask:   nullMask,
		partition:  nil,
		ctx:        s.ctx,
	}
}

////////////////////////			GROUPING OPERATIONS

// A SeriesStringPartition is a partition of a SeriesString.
// Each key is a hash of a bool value, and each value is a slice of indices
// of the original series that are set to that value.
type SeriesStringPartition struct {
	partition map[int64][]int
	ctx       *Context
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
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(s.data), s.HasNull(),
			worker, workerNulls),
		ctx: s.ctx,
	}

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
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_1, len(keys), s.HasNull(),
			worker, workerNulls),
		ctx: s.ctx,
	}

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
	if s.partition != nil {
		return SeriesError{"SeriesString.ToUpper() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.ctx.stringPool.Put(strings.ToUpper(*s.data[i]))
	}

	return s
}

func (s SeriesString) ToLower() Series {
	if s.partition != nil {
		return SeriesError{"SeriesString.ToLower() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.ctx.stringPool.Put(strings.ToLower(*s.data[i]))
	}

	return s
}

func (s SeriesString) TrimSpace() Series {
	if s.partition != nil {
		return SeriesError{"SeriesString.TrimSpace() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.ctx.stringPool.Put(strings.TrimSpace(*s.data[i]))
	}

	return s
}

func (s SeriesString) Trim(cutset string) Series {
	if s.partition != nil {
		return SeriesError{"SeriesString.Trim() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.ctx.stringPool.Put(strings.Trim(*s.data[i], cutset))
	}

	return s
}

func (s SeriesString) Replace(old, new string, n int) Series {
	if s.partition != nil {
		return SeriesError{"SeriesString.Replace() not supported on grouped Series"}
	}

	for i := 0; i < len(s.data); i++ {
		s.data[i] = s.ctx.stringPool.Put(strings.Replace(*s.data[i], old, new, n))
	}

	return s
}
