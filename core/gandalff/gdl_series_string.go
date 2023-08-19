package gandalff

import (
	"fmt"
	"sort"
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

// Sets the name of the series.
func (s SeriesString) SetName(name string) Series {
	s.name = name
	return s
}

// Returns the type of the series.
func (s SeriesString) Type() typesys.BaseType {
	return typesys.StringType
}

// Returns the type and cardinality of the series.
func (s SeriesString) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{Base: typesys.StringType, Card: s.Len()}
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

// Returns if the series is error.
func (s SeriesString) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s SeriesString) GetError() string {
	return ""
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
func (s SeriesString) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess(s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces)
}

// Append appends a value or a slice of values to the series.
func (s SeriesString) Append(v any) Series {
	switch v := v.(type) {
	case string, []string:
		return s.appendRaw(v)
	case NullableString, []NullableString:
		return s.appendNullable(v)
	case SeriesString:
		return s.appendSeries(v)
	case SeriesError:
		return v
	default:
		return SeriesError{fmt.Sprintf("SeriesString.Append: invalid type, %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s SeriesString) appendRaw(v any) Series {
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
func (s SeriesString) appendNullable(v any) Series {
	if !s.isNullable {
		return SeriesError{"SeriesString.AppendNullable: series is not nullable"}
	}

	switch v := v.(type) {
	case NullableString:
		s.data = append(s.data, s.pool.Put(v.Value))
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableString:
		ssize := len(s.data)
		s.data = append(s.data, make([]*string, len(v))...)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for i, b := range v {
			s.data[ssize+i] = s.pool.Put(b.Value)
			if !b.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesString.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s SeriesString) appendSeries(other Series) Series {
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

func (s SeriesString) Strings() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return data
}

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
		data := make([]bool, len(s.data))
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

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s SeriesString) Filter(mask any) Series {
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
		return SeriesError{fmt.Sprintf("SeriesString.Filter: invalid type %T", mask)}
	}
}

func (s SeriesString) filterBool(mask SeriesBool) Series {
	return s.filterBoolSlice(mask.data)
}

func (s SeriesString) filterBoolMemOpt(mask SeriesBoolMemOpt) Series {
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

func (s SeriesString) filterBoolSlice(mask []bool) Series {
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

func (s SeriesString) filterIntSlice(indeces []int) Series {
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
		data := make([]bool, len(s.data))
		chunkLen := len(s.data) / THREADS_NUMBER
		if chunkLen < MINIMUM_PARALLEL_SIZE_2 {
			for i := 0; i < len(s.data); i++ {
				data[i] = f(*(s.data[i])).(bool)
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
						data[i] = f(*(s.data[i])).(bool)
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

func (s SeriesString) UnGroup() Series {
	s.isGrouped = false
	s.partition = nil
	return s
}

func (s SeriesString) GetPartition() SeriesPartition {
	return s.partition
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
	if s.isNullable && (s.nullMask[i>>3]&(1<<uint(i%8))) > 0 && (s.nullMask[j>>3]&(1<<uint(j%8)) > 0) {
		return true
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
