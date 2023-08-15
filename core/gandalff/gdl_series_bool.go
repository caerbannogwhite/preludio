package gandalff

import (
	"fmt"
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

////////////////////////			BASIC ACCESSORS

// Returns the number of elements in the series.
func (s SeriesBool) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s SeriesBool) Name() string {
	return s.name
}

// Sets the name of the series.
func (s SeriesBool) SetName(name string) Series {
	s.name = name
	return s
}

// Returns the type of the series.
func (s SeriesBool) Type() typesys.BaseType {
	return typesys.BoolType
}

// Returns the type and cardinality of the series.
func (s SeriesBool) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{Base: typesys.BoolType, Card: s.Len()}
}

// Returns if the series is grouped.
func (s SeriesBool) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s SeriesBool) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s SeriesBool) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series is error.
func (s SeriesBool) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s SeriesBool) GetError() string {
	return ""
}

// Returns if the series has null values.
func (s SeriesBool) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s SeriesBool) NullCount() int {
	count := 0
	for _, x := range s.nullMask {
		for x != 0 {
			count += int(x & 1)
			x >>= 1
		}
	}
	return count
}

// Returns the number of non-null values in the series.
func (s SeriesBool) NonNullCount() int {
	return len(s.data) - s.NullCount()
}

// Returns if the element at index i is null.
func (s SeriesBool) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i>>3]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s SeriesBool) SetNull(i int) Series {
	if s.isNullable {
		s.nullMask[i>>3] |= 1 << uint(i%8)

		s.sorted = SORTED_NONE
		return s
	} else {
		nullMask := make([]uint8, len(s.data))
		nullMask[i>>3] |= 1 << uint(i%8)

		s.isNullable = true
		s.sorted = SORTED_NONE
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s SeriesBool) GetNullMask() []bool {
	mask := make([]bool, len(s.data))
	idx := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8 && idx < len(s.data); i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

// Sets the null mask of the series.
func (s SeriesBool) SetNullMask(mask []bool) Series {
	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				s.nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}

		s.sorted = SORTED_NONE
		return s
	} else {
		nullMask := make([]uint8, len(s.data))
		for k, v := range mask {
			if v {
				nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}

		s.isNullable = true
		s.sorted = SORTED_NONE
		s.nullMask = nullMask

		return s
	}
}

// Makes the series nullable.
func (s SeriesBool) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.sorted = SORTED_NONE
		s.nullMask = make([]uint8, len(s.data))
	}
	return s
}

// Get the element at index i.
func (s SeriesBool) Get(i int) any {
	return s.data[i]
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
	if b, ok := v.(bool); ok {
		s.data[i] = b
	} else if nb, ok := v.(NullableBool); ok {
		if nb.Valid {
			s.data[i] = nb.Value
		} else {
			s.nullMask[i>>3] |= 1 << uint(i%8)
			s.data[i] = false
		}
	} else {
		return SeriesError{fmt.Sprintf("SeriesBool.Set: provided value %t is not of type bool or NullableBool", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesBool) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess(s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces)
}

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

func (s SeriesBool) Swap(i, j int) {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			s.nullMask[j>>3] |= 1 << uint(j%8)
		} else {
			s.nullMask[j>>3] &= ^(1 << uint(j%8))
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			s.nullMask[i>>3] |= 1 << uint(i%8)
		} else {
			s.nullMask[i>>3] &= ^(1 << uint(i%8))
		}
	}

	s.data[i], s.data[j] = s.data[j], s.data[i]
}

// Append appends a value or a slice of values to the series.
func (s SeriesBool) Append(v any) Series {
	switch v := v.(type) {
	case bool, []bool:
		return s.appendRaw(v)
	case NullableBool, []NullableBool:
		return s.appendNullable(v)
	case SeriesBool:
		return s.appendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesBool.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesBool) appendRaw(v any) Series {
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

	default:
		return SeriesError{fmt.Sprintf("SeriesBool.Append: invalid type %T", v)}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s SeriesBool) appendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesBool.AppendNullable: series is not nullable"}
	}

	switch v := v.(type) {
	case NullableBool:
		s.data = append(s.data, v.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableBool:
		ssize := len(s.data)
		s.data = append(s.data, make([]bool, len(v))...)
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
		return SeriesError{fmt.Sprintf("SeriesBool.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s SeriesBool) appendSeries(other Series) Series {
	var ok bool
	var o SeriesBool
	if o, ok = other.(SeriesBool); !ok {
		return SeriesError{fmt.Sprintf("SeriesBool.AppendSeries: invalid type %T", other)}
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
		data[i] = NullableBool{v, s.IsNull(i)}
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

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s SeriesBool) Filter(mask any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesBool.Filter: invalid type %T", mask)}
	}
}

func (s SeriesBool) filterBool(mask SeriesBool) Series {
	return s.filterBoolSlice(mask.data)
}

// Filters out the elements by the given mask series.
func (s SeriesBool) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
	if mask.Len() != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesBool.Filter: mask length (%d) does not match series length (%d)", mask.Len(), len(s.data))}
	}

	if mask.isNullable {
		return SeriesError{"SeriesBool.Filter: mask series cannot be nullable for this operation"}
	}

	return s.filterBoolSlice(mask.Data().([]bool))
}

func (s SeriesBool) filterBoolSlice(mask []bool) Series {
	if len(mask) != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesBool.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	data := make([]bool, elementCount)
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

func (s SeriesBool) filterIntSlice(indexes []int) Series {
	var nullMask []uint8

	size := len(indexes)
	data := make([]bool, size)

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
	seriesList []Series
	partition  map[int64][]int
}

func (gp *SeriesBoolPartition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesBoolPartition) getMap() map[int64][]int {
	return gp.partition
}

func (gp *SeriesBoolPartition) getSeriesList() []Series {
	return gp.seriesList
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

func (s SeriesBool) Sort() Series {
	return s
}

func (s SeriesBool) SortRev() Series {
	return s
}
