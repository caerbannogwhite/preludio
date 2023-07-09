package gandalff

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"typesys"
	"unsafe"
)

// SeriesString represents a series of strings.
type SeriesString struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	name       string
	data       []*string
	nullMask   []uint8
	pool       *StringPool
	partition  *SeriesStringPartition
}

func NewSeriesString(name string, isNullable bool, data []string, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	actualData := make([]*string, len(data))
	for i, v := range data {
		actualData[i] = pool.Put(v)
	}

	return SeriesString{isNullable: isNullable, name: name, data: actualData, nullMask: nullMask, pool: pool}
}

////////////////////////			BASIC ACCESSORS

func (s SeriesString) SetStringPool(pool *StringPool) Series {
	for i, v := range s.data {
		s.data[i] = pool.Put(*v)
	}
	s.pool = pool
	return s
}

// Returns the number of elements in the series.
func (s SeriesString) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s SeriesString) Name() string {
	return s.name
}

// Returns the type of the series.
func (s SeriesString) Type() typesys.BaseType {
	return typesys.StringType
}

// Returns if the series is grouped.
func (s SeriesString) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s SeriesString) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s SeriesString) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series has null values.
func (s SeriesString) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s SeriesString) NullCount() int {
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
func (s SeriesString) NonNullCount() int {
	return s.Len() - s.NullCount()
}

// Returns if the element at index i is null.
func (s SeriesString) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s SeriesString) SetNull(i int) Series {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return s
	} else {
		nullMask := __binVecInit(len(s.data))
		nullMask[i/8] |= 1 << uint(i%8)

		return SeriesString{
			isGrouped:  s.isGrouped,
			isNullable: true,
			sorted:     s.sorted,
			name:       s.name,
			data:       s.data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  s.partition,
		}
	}
}

// Returns the null mask of the series.
func (s SeriesString) GetNullMask() []bool {
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
func (s SeriesString) SetNullMask(mask []bool) Series {
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

		return SeriesString{
			isGrouped:  s.isGrouped,
			isNullable: true,
			sorted:     s.sorted,
			name:       s.name,
			data:       s.data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  s.partition,
		}
	}
}

// Makes the series nullable.
func (s SeriesString) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __binVecInit(len(s.data))
	}
	return s
}

// Get the element at index i.
func (s SeriesString) Get(i int) any {
	return *s.data[i]
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
	if ss, ok := v.(string); ok {
		s.data[i] = s.pool.Put(ss)
	} else if ns, ok := v.(NullableString); ok {
		if ns.Valid {
			s.data[i] = s.pool.Put(ns.Value)
		} else {
			s.data[i] = nil
			s.nullMask[i/8] |= 1 << uint(i%8)
		}
	} else {
		return SeriesError{fmt.Sprintf("SeriesString.Set: provided value %t is not of type string or NullableString", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesString) Take(start, end, step int) Series {
	return s
}

func (s SeriesString) Less(i, j int) bool {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			return false
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			return true
		}
	}
	return strings.Compare(*s.data[i], *s.data[j]) < 0
}

func (s SeriesString) Swap(i, j int) {
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
func (s SeriesString) Append(v any) Series {
	switch v := v.(type) {
	case string:
		return s.AppendRaw(v)
	case []string:
		return s.AppendRaw(v)
	case NullableString:
		return s.AppendNullable(v)
	case []NullableString:
		return s.AppendNullable(v)
	case SeriesString:
		return s.AppendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesString.Append: invalid type, %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesString) AppendRaw(v any) Series {
	if s.isNullable {
		if str, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Put(str))
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if strv, ok := v.([]string); ok {
			for _, str := range strv {
				s.data = append(s.data, s.pool.Put(str))
			}
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
			}
		} else {
			return SeriesError{fmt.Sprintf("SeriesString.AppendRaw: invalid type %T", v)}
		}
	} else {
		if str, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Put(str))
		} else if strv, ok := v.([]string); ok {
			for _, str := range strv {
				s.data = append(s.data, s.pool.Put(str))
			}
		} else {
			return SeriesError{fmt.Sprintf("SeriesString.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s SeriesString) AppendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesString.AppendNullable: series is not nullable"}
	}

	if str, ok := v.(NullableString); ok {
		s.data = append(s.data, s.pool.Put(str.Value))
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !str.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}
	} else if strv, ok := v.([]NullableString); ok {
		if len(s.data) > len(s.nullMask)<<8 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for _, str := range strv {
			s.data = append(s.data, s.pool.Put(str.Value))
			if !str.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return SeriesError{fmt.Sprintf("SeriesString.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s SeriesString) AppendSeries(other Series) Series {
	var ok bool
	var o SeriesString
	if o, ok = other.(SeriesString); !ok {
		return SeriesError{fmt.Sprintf("SeriesString.AppendSeries: invalid type %T", other)}
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

func (s SeriesString) Data() any {
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
func (s SeriesString) Cast(t typesys.BaseType, stringPool *StringPool) Series {
	switch t {
	case typesys.BoolType:
		data := __binVecInit(len(s.data))
		nullMask := __binVecInit(len(s.data))
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
					} else if b {
						data[i>>3] |= (1 << uint(i%8))
					}
				}
			}
		} else {
			for i, v := range s.data {
				b, err := typeGuesser.atoBool(*v)
				if err != nil {
					nullMask[i>>3] |= (1 << uint(i%8))
				} else if b {
					data[i>>3] |= (1 << uint(i%8))
				}
			}
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: true,
			sorted:     SORTED_NONE,
			size:       len(s.data),
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}

	case typesys.Int32Type:
		data := make([]int32, len(s.data))
		nullMask := __binVecInit(len(s.data))
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
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}

	case typesys.Int64Type:
		data := make([]int64, len(s.data))
		nullMask := __binVecInit(len(s.data))
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
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}

	case typesys.Float64Type:
		data := make([]float64, len(s.data))
		nullMask := __binVecInit(len(s.data))
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
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}

	case typesys.StringType:
		return s

	default:
		return SeriesError{fmt.Sprintf("SeriesString.Cast: invalid type %s", t.ToString())}
	}
}

func (s SeriesString) Copy() Series {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return NewSeriesString(s.name, s.isNullable, data, s.pool)
}

func (s SeriesString) getDataPtr() *[]*string {
	return &s.data
}

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask series.
func (s SeriesString) Filter(mask SeriesBool) Series {
	if mask.size != s.Len() {
		return SeriesError{fmt.Sprintf("SeriesString.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return SeriesError{"SeriesString.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()

	data := make([]*string, elementCount)
	var nullMask []uint8

	if s.isNullable {

		nullMask = __binVecInit(elementCount)

		if s.NullCount() > 0 {
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
func (s SeriesString) FilterByMask(mask []bool) Series {
	if len(mask) != len(s.data) {
		return SeriesError{fmt.Sprintf("SeriesString.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	chunkLen := len(mask) / THREADS_NUMBER
	elementCountVec := make([]int, THREADS_NUMBER)

	for i := 0; i < THREADS_NUMBER-1; i++ {
		for j := chunkLen * i; j < chunkLen*(i+1); j++ {
			if mask[j] {
				elementCountVec[i]++
			}
		}
	}

	for j := chunkLen * (THREADS_NUMBER - 1); j < len(mask); j++ {
		if mask[j] {
			elementCountVec[THREADS_NUMBER-1]++
		}
	}

	elementCountTot := 0
	for _, v := range elementCountVec {
		elementCountTot += v
	}

	var data []*string
	var nullMask []uint8

	data = make([]*string, elementCountTot)

	if s.isNullable {

		if elementCountTot%8 == 0 {
			nullMask = make([]uint8, (elementCountTot >> 3))
		} else {
			nullMask = make([]uint8, (elementCountTot>>3)+1)
		}

		if s.NullCount() > 0 {
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

			// if chunkLen < MINIMUM_PARALLEL_SIZE {
			dstIdx := 0
			for srcIdx, v := range mask {
				if v {
					data[dstIdx] = s.data[srcIdx]
					dstIdx++
				}
			}
			// } else {
			// 	var wg sync.WaitGroup
			// 	wg.Add(THREADS_NUMBER)

			// 	for n := 0; n < THREADS_NUMBER; n++ {
			// 		dstIdx := 0
			// 		if n > 0 {
			// 			dstIdx = elementCountVec[n-1]
			// 		}

			// 		start := n * chunkLen
			// 		end := (n + 1) * chunkLen
			// 		if n == THREADS_NUMBER-1 {
			// 			end = len(s.data)
			// 		}

			// 		go func() {
			// 			for srcIdx, v := range mask[start:end] {
			// 				if v {
			// 					data[dstIdx] = s.data[srcIdx]
			// 					dstIdx++
			// 				}
			// 			}
			// 			wg.Done()
			// 		}()
			// 	}

			// 	wg.Wait()
			// }
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

func (s SeriesString) FilterByIndeces(indeces []int) Series {
	var data []*string
	var nullMask []uint8

	size := len(indeces)
	data = make([]*string, size)

	if s.isNullable {

		nullMask = __binVecInit(size)

		for dstIdx, srcIdx := range indeces {
			data[dstIdx] = s.data[srcIdx]
			if srcIdx%8 > dstIdx%8 {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
			}
		}
	} else {
		for dstIdx, srcIdx := range indeces {
			data[dstIdx] = s.data[srcIdx]
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}

func (s SeriesString) Map(f GDLMapFunc, stringPool *StringPool) Series {
	if len(s.data) == 0 {
		return s
	}

	v := f(*(s.data[0]))
	switch v.(type) {
	case bool:

		data := __binVecInit(len(s.data))

		chunkLen := len(s.data) / THREADS_NUMBER
		if chunkLen < MINIMUM_PARALLEL_SIZE_2 {
			for i := 0; i < len(s.data); i++ {
				if f((*s.data[i])).(bool) {
					data[i>>3] |= (1 << uint(i%8))
				}
			}
		} else {
			var wg sync.WaitGroup
			wg.Add(THREADS_NUMBER)

			for n := 0; n < THREADS_NUMBER; n++ {
				start := n * chunkLen
				end := (n + 1) * chunkLen
				if n == THREADS_NUMBER-1 {
					end = len(s.data)
				}

				go func() {
					for i := start; i < end; i++ {
						if f((*s.data[i])).(bool) {
							data[i>>3] |= (1 << uint(i%8))
						}
					}
					wg.Done()
				}()
			}

			wg.Wait()
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
			data[i] = f((*s.data[i])).(int32)
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
			data[i] = f((*s.data[i])).(int64)
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
			data[i] = f((*s.data[i])).(float64)
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
			return SeriesError{"SeriesString.Map: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = stringPool.Put(f((*s.data[i])).(string))
		}

		s.isGrouped = false
		s.sorted = SORTED_NONE
		s.data = data

		return s
	}

	return SeriesError{fmt.Sprintf("SeriesString.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

type SeriesStringPartition struct {
	seriesSize   int
	partition    map[int64][]int
	indexToGroup []int
	pool         *StringPool
}

func (gp SeriesStringPartition) GetSize() int {
	return len(gp.partition)
}

func (gp SeriesStringPartition) GetMap() map[int64][]int {
	return gp.partition
}

func (gp SeriesStringPartition) GetValueIndices(val any) []int {
	if val == nil {
		if nulls, ok := gp.partition[HASH_NULL_KEY]; ok {
			return nulls
		}
	} else {
		if v, ok := val.(string); ok {
			if addr := gp.pool.Put(v); addr == nil {
				if vals, ok := gp.partition[(*(*int64)(unsafe.Pointer(unsafe.Pointer(addr))))]; ok {
					return vals
				}
			}
		}
	}

	return make([]int, 0)
}

func (gp SeriesStringPartition) GetKeys() any {
	keys := make([]string, 0, len(gp.partition))
	for k := range gp.partition {
		if k != HASH_NULL_KEY {
			keys = append(keys, *(*string)(unsafe.Pointer(&k)))
		}
	}
	return keys
}

func (gp SeriesStringPartition) debugPrint() {
	fmt.Println("SeriesStringPartition")
	map_ := gp.GetMap()
	for k, v := range map_ {
		ptr := (*string)(unsafe.Pointer(uintptr(k)))
		fmt.Printf("%10s: %v\n", *ptr, v)
	}
}

func (s SeriesString) Group() Series {

	var partition SeriesStringPartition
	if len(s.data) < MINIMUM_PARALLEL_SIZE_1 {
		map_ := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

		var ptr unsafe.Pointer
		for i := 0; i < len(s.data); i++ {
			ptr = unsafe.Pointer(s.data[i])
			map_[(*(*int64)(unsafe.Pointer(&ptr)))] = append(map_[(*(*int64)(unsafe.Pointer(&ptr)))], i)
		}

		partition = SeriesStringPartition{
			seriesSize: s.Len(),
			partition:  map_,
			pool:       s.pool,
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
			var ptr unsafe.Pointer
			for i := start; i < end; i++ {
				ptr = unsafe.Pointer(s.data[i])
				map_[(*(*int64)(unsafe.Pointer(&ptr)))] = append(map_[(*(*int64)(unsafe.Pointer(&ptr)))], i)
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(s.data), allMaps, nil, worker)

		partition = SeriesStringPartition{
			seriesSize: s.Len(),
			partition:  allMaps[0],
			pool:       s.pool,
		}
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s SeriesString) SubGroup(partition SeriesPartition) Series {
	var newPartition SeriesStringPartition
	otherIndeces := partition.GetMap()

	if len(s.data) < MINIMUM_PARALLEL_SIZE_1 {

		map_ := make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

		var newHash int64
		var ptr unsafe.Pointer
		for h, v := range otherIndeces {
			for _, index := range v {
				ptr = unsafe.Pointer(s.data[index])
				newHash = *(*int64)(unsafe.Pointer(&ptr)) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}

		newPartition = SeriesStringPartition{
			seriesSize: s.Len(),
			partition:  map_,
			pool:       s.pool,
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
			var ptr unsafe.Pointer

			map_ := allMaps[threadNum]
			for _, h := range keys[start:end] {
				for _, index := range otherIndeces[h] {
					ptr = unsafe.Pointer(s.data[index])
					newHash = *(*int64)(unsafe.Pointer(&ptr)) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
					map_[newHash] = append(map_[newHash], index)
				}
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(keys), allMaps, nil, worker)

		newPartition = SeriesStringPartition{
			seriesSize: s.Len(),
			partition:  allMaps[0],
			pool:       s.pool,
		}
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s SeriesString) GetPartition() SeriesPartition {
	return s.partition
}

func (s SeriesString) Sort() Series {
	return s
}

func (s SeriesString) SortRev() Series {
	return s
}

////////////////////////			SORTING OPERATIONS

////////////////////////			ARITHMETIC OPERATIONS

func (s SeriesString) Add(other Series) Series {
	switch o := other.(type) {
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	}
	return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), other.Type().ToString())}

}

////////////////////////			LOGICAL OPERATIONS

// The body of this function has been generated by the code generator at
// https://github.com/caerbannogwhite/preludio/tree/main/core/gandalff/gen
func (s SeriesString) Eq(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		}
	}
	return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesString) Ne(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		}
	}
	return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesString) Gt(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		}
	}
	return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesString) Ge(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		}
	}
	return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesString) Lt(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		}
	}
	return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s SeriesString) Le(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: boolVecToBinVec(result), size: resultSize}
					}
				}
			}
		}
	}
	return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), other.Type().ToString())}

}
