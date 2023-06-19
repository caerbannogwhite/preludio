package gandalff

import (
	"fmt"
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

func NewSeriesFloat64(name string, isNullable bool, makeCopy bool, data []float64) SeriesFloat64 {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesFloat64{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

////////////////////////			BASIC ACCESSORS

// Returns the number of elements in the series.
func (s SeriesFloat64) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s SeriesFloat64) Name() string {
	return s.name
}

// Returns the type of the series.
func (s SeriesFloat64) Type() typesys.BaseType {
	return typesys.Float64Type
}

// Returns if the series is grouped.
func (s SeriesFloat64) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s SeriesFloat64) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s SeriesFloat64) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series has null values.
func (s SeriesFloat64) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s SeriesFloat64) NullCount() int {
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
func (s SeriesFloat64) NonNullCount() int {
	return s.Len() - s.NullCount()
}

// Returns if the element at index i is null.
func (s SeriesFloat64) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s SeriesFloat64) SetNull(i int) Series {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	} else {
		nullMask := __binVecInit(len(s.data))
		nullMask[i/8] |= 1 << uint(i%8)

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s SeriesFloat64) GetNullMask() []bool {
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
func (s SeriesFloat64) SetNullMask(mask []bool) Series {
	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k/8] |= 1 << uint(k%8)
			} else {
				s.nullMask[k/8] &= ^(1 << uint(k%8))
			}
		}
		return s
	} else {
		nullMask := __binVecInit(len(s.data))
		for k, v := range mask {
			if v {
				nullMask[k/8] |= 1 << uint(k%8)
			} else {
				nullMask[k/8] &= ^(1 << uint(k%8))
			}
		}

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Makes the series nullable.
func (s SeriesFloat64) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __binVecInit(len(s.data))
	}
	return s
}

// Get the element at index i.
func (s SeriesFloat64) Get(i int) any {
	return s.data[i]
}

// Get the element at index i as a string.
func (s SeriesFloat64) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return floatToString(s.data[i])
}

// Set the element at index i. The type of v must be float64 or NullableFloat64.
func (s SeriesFloat64) Set(i int, v any) Series {
	if f, ok := v.(float64); ok {
		s.data[i] = f
	} else if nf, ok := v.(NullableFloat64); ok {
		if nf.Valid {
			s.data[i] = nf.Value
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	} else {
		return SeriesError{fmt.Sprintf("SeriesFloat64.Set: provided value %t is not of type float64 or NullableFloat64", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesFloat64) Take(start, end, step int) Series {
	return s
}

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

func (s SeriesFloat64) Swap(i, j int) {
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

func (s SeriesFloat64) Append(v any) Series {
	switch v := v.(type) {
	case float64:
		return s.AppendRaw(v)
	case []float64:
		return s.AppendRaw(v)
	case NullableFloat64:
		return s.AppendNullable(v)
	case []NullableFloat64:
		return s.AppendNullable(v)
	case SeriesFloat64:
		return s.AppendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesFloat64.Append: invalid type, %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesFloat64) AppendRaw(v any) Series {
	if s.isNullable {
		if f, ok := v.(float64); ok {
			s.data = append(s.data, f)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if fv, ok := v.([]float64); ok {
			s.data = append(s.data, fv...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
			}
		} else {
			return SeriesError{fmt.Sprintf("SeriesFloat64.AppendRaw: invalid type %T", v)}
		}
	} else {
		if f, ok := v.(float64); ok {
			s.data = append(s.data, f)
		} else if fv, ok := v.([]float64); ok {
			s.data = append(s.data, fv...)
		} else {
			return SeriesError{fmt.Sprintf("SeriesFloat64.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s SeriesFloat64) AppendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesFloat64.AppendNullable: series is not nullable"}
	}

	if f, ok := v.(NullableFloat64); ok {
		s.data = append(s.data, f.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !f.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}
	} else if fv, ok := v.([]NullableFloat64); ok {
		if len(s.data) > len(s.nullMask)<<8 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for _, f := range fv {
			s.data = append(s.data, f.Value)
			if !f.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return SeriesError{fmt.Sprintf("SeriesFloat64.AppendNullable: invalid type %T", v)}
	}

	return s
}

func (s SeriesFloat64) AppendSeries(other Series) Series {
	var ok bool
	var o SeriesFloat64
	if o, ok = other.(SeriesFloat64); !ok {
		return SeriesError{fmt.Sprintf("SeriesFloat64.AppendSeries: invalid type %T", other)}
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
			if len(s.data)%8 == 0 {
				s.nullMask = make([]uint8, (len(s.data) >> 3))
			} else {
				s.nullMask = make([]uint8, (len(s.data)>>3)+1)
			}
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
		data := __binVecInit(len(s.data))
		for i, v := range s.data {
			if v != 0 {
				data[i>>3] |= (1 << uint(i%8))
			}
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			size:       len(s.data),
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

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask series.
func (s SeriesFloat64) Filter(mask SeriesBool) Series {
	if mask.size != s.Len() {
		return SeriesError{fmt.Sprintf("SeriesFloat64.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return SeriesError{"SeriesFloat64.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()
	var nullMask []uint8

	data := make([]float64, elementCount)
	if s.isNullable {

		nullMask = __binVecInit(elementCount)

		dstIdx := 0
		for srcIdx := 0; srcIdx < s.Len(); srcIdx++ {
			if mask.data[srcIdx>>3]&(1<<uint(srcIdx%8)) != 0 {
				data[dstIdx] = s.data[srcIdx]
				if srcIdx%8 > dstIdx%8 {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				} else {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
				}
				dstIdx++
			}
		}
	} else {
		dstIdx := 0
		for srcIdx := 0; srcIdx < s.Len(); srcIdx++ {
			if mask.data[srcIdx>>3]&(1<<uint(srcIdx%8)) != 0 {
				data[dstIdx] = s.data[srcIdx]
				dstIdx++
			}
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}

// FilterByMask returns a new series with elements filtered by the mask.
func (s SeriesFloat64) FilterByMask(mask []bool) Series {
	if len(mask) != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesFloat64.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []float64
	var nullMask []uint8

	data = make([]float64, elementCount)

	if s.isNullable {

		nullMask = __binVecInit(elementCount)

		dstIdx := 0
		for srcIdx, v := range mask {
			if v {
				data[dstIdx] = s.data[srcIdx]
				if srcIdx%8 > dstIdx%8 {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				} else {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
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

func (s SeriesFloat64) FilterByIndeces(indexes []int) Series {
	var data []float64
	var nullMask []uint8

	size := len(indexes)
	data = make([]float64, size)

	if s.isNullable {

		nullMask = __binVecInit(size)

		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
			if srcIdx%8 > dstIdx%8 {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
			}
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

func (s SeriesFloat64) Map(f GDLMapFunc, stringPool *StringPool) Series {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:

		// Not a null mask but you get the same result
		data := __binVecInit(len(s.data))
		for i := 0; i < len(s.data); i++ {
			if f(s.data[i]).(bool) {
				data[i>>3] |= (1 << uint(i%8))
			}
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			size:       len(s.data),
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
	series       *SeriesFloat64
	seriesSize   int
	partition    map[int64][]int
	indexToGroup []int
}

func (gp SeriesFloat64Partition) GetSize() int {
	return len(gp.partition)
}

func (gp SeriesFloat64Partition) beginSorting() SeriesFloat64Partition {
	gp.indexToGroup = make([]int, gp.seriesSize)
	for i, part := range gp.partition {
		for _, idx := range part {
			gp.indexToGroup[idx] = int(i)
		}
	}

	return gp
}

func (gp SeriesFloat64Partition) endSorting() SeriesFloat64Partition {
	// newPartition := make(map[int64][]int, len(gp.partition))
	// newNullGroup := make([]int, len(gp.nulls))

	// for i, part := range gp.partition {
	// 	newPartition[i] = make([]int, 0, len(part))
	// }

	// for i, g := range gp.indexToGroup {
	// 	if g < len(gp.partition) {
	// 		newPartition[int64(g)] = append(newPartition[int64(g)], i)
	// 	} else {
	// 		newNullGroup[g-len(gp.partition)] = append(newNullGroup[g-len(gp.partition)], i)
	// 	}
	// }

	gp.indexToGroup = nil
	return gp
}

func (gp SeriesFloat64Partition) GetMap() map[int64][]int {
	return gp.partition
}

func (gp SeriesFloat64Partition) GetValueIndices(val any) []int {
	if val == nil {
		if nulls, ok := gp.partition[HASH_NULL_KEY]; ok {
			return nulls
		}
	} else if v, ok := val.(float64); ok {
		if part, ok := gp.partition[int64(v)]; ok {
			return part
		}
	}

	return make([]int, 0)
}

func (gp SeriesFloat64Partition) GetKeys() any {
	keys := make([]float64, 0, len(gp.partition))
	for k := range gp.partition {
		if k != HASH_NULL_KEY {
			keys = append(keys, *(*float64)(unsafe.Pointer(&k)))
		}
	}
	return keys
}

func (gp SeriesFloat64Partition) debugPrint() {
	fmt.Println("SeriesFloat64Partition")
	data := gp.series.Data().([]float64)
	for k, v := range gp.partition {
		// f := *(*float64)(unsafe.Pointer(&k))
		fmt.Printf("%v - %10.4f: %v\n", k, data[v[0]], v)
	}
}

func (s SeriesFloat64) Group() Series {

	var partition SeriesFloat64Partition
	if len(s.data) < MINIMUM_PARALLEL_SIZE_2 {
		map_ := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
		for i, v := range s.data {
			map_[*(*int64)(unsafe.Pointer((&v)))] = append(map_[*(*int64)(unsafe.Pointer((&v)))], i)
		}

		partition = SeriesFloat64Partition{
			series:     &s,
			seriesSize: s.Len(),
			partition:  map_,
		}
	} else {

		// Initialize the maps and the wait groups
		allMaps := make([]map[int64][]int, THREADS_NUMBER)
		for i := 0; i < THREADS_NUMBER; i++ {
			allMaps[i] = make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
		}

		// Define the worker callback
		worker := func(threadNum, start, end int) {
			map_ := allMaps[threadNum]
			for i := start; i < end; i++ {
				map_[*(*int64)(unsafe.Pointer((&s.data[i])))] = append(map_[*(*int64)(unsafe.Pointer((&s.data[i])))], i)
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(s.data), allMaps, nil, worker)

		partition = SeriesFloat64Partition{
			series:     &s,
			seriesSize: s.Len(),
			partition:  allMaps[0],
		}
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s SeriesFloat64) SubGroup(partition SeriesPartition) Series {
	var newPartition SeriesFloat64Partition
	otherIndeces := partition.GetMap()

	if len(s.data) < MINIMUM_PARALLEL_SIZE_2 {

		map_ := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

		var newHash int64
		for h, v := range otherIndeces {
			for _, index := range v {
				newHash = *(*int64)(unsafe.Pointer((&(s.data)[index]))) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}

		newPartition = SeriesFloat64Partition{
			series:     &s,
			seriesSize: s.Len(),
			partition:  map_,
		}
	} else {

		// collect all keys
		keys := make([]int64, len(otherIndeces))
		i := 0
		for k := range otherIndeces {
			keys[i] = k
			i++
		}

		// Initialize the maps and the wait groups
		allMaps := make([]map[int64][]int, THREADS_NUMBER)
		for i := 0; i < THREADS_NUMBER; i++ {
			allMaps[i] = make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
		}

		// Define the worker callback
		worker := func(threadNum, start, end int) {
			var newHash int64
			map_ := allMaps[threadNum]
			for _, h := range keys[start:end] {
				for _, index := range otherIndeces[h] {
					newHash = *(*int64)(unsafe.Pointer((&(s.data)[index]))) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
					map_[newHash] = append(map_[newHash], index)
				}
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(keys), allMaps, nil, worker)

		newPartition = SeriesFloat64Partition{
			series:     &s,
			seriesSize: s.Len(),
			partition:  allMaps[0],
		}
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s SeriesFloat64) GetPartition() SeriesPartition {
	return s.partition
}

func (s SeriesFloat64) Sort() Series {
	return s
}

func (s SeriesFloat64) SortRev() Series {
	return s
}

////////////////////////			SORTING OPERATIONS

////////////////////////			ARITHMETIC OPERATIONS

func (s SeriesFloat64) Mul(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), other.Type().ToString())}
}

func (s SeriesFloat64) Div(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesFloat64) Mod(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesFloat64) Add(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesFloat64) Sub(other Series) Series {
	return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), other.Type().ToString())}

}
