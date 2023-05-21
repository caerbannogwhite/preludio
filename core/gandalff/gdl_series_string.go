package gandalff

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"typesys"
	"unsafe"
)

// GDLSeriesString represents a series of strings.
type GDLSeriesString struct {
	isGrouped  bool
	isNullable bool
	sorted     GDLSeriesSortOrder
	name       string
	data       []*string
	nullMask   []uint8
	pool       *StringPool
	partition  *SeriesStringPartition
}

func NewGDLSeriesString(name string, isNullable bool, data []string, pool *StringPool) GDLSeries {
	var nullMask []uint8
	if isNullable {
		nullMask = __initPackBinVec(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	actualData := make([]*string, len(data))
	for i, v := range data {
		actualData[i] = pool.Get(v)
	}

	return GDLSeriesString{isNullable: isNullable, name: name, data: actualData, nullMask: nullMask, pool: pool}
}

////////////////////////			BASIC ACCESSORS

// Returns the number of elements in the series.
func (s GDLSeriesString) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s GDLSeriesString) Name() string {
	return s.name
}

// Returns the type of the series.
func (s GDLSeriesString) Type() typesys.BaseType {
	return typesys.StringType
}

// Returns if the series is grouped.
func (s GDLSeriesString) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s GDLSeriesString) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s GDLSeriesString) IsSorted() GDLSeriesSortOrder {
	return s.sorted
}

// Returns if the series has null values.
func (s GDLSeriesString) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s GDLSeriesString) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8; i++ {
			count += int((v & (1 << uint(i))) >> uint(i))
		}
	}
	return count
}

// Returns the number of non-null values in the series.
func (s GDLSeriesString) NonNullCount() int {
	return s.Len() - s.NullCount()
}

// Returns if the element at index i is null.
func (s GDLSeriesString) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s GDLSeriesString) SetNull(i int) GDLSeries {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return s
	} else {
		nullMask := __initPackBinVec(len(s.data))
		nullMask[i/8] |= 1 << uint(i%8)

		return GDLSeriesString{
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
func (s GDLSeriesString) GetNullMask() []bool {
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
func (s GDLSeriesString) SetNullMask(mask []bool) GDLSeries {
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
		nullMask := __initPackBinVec(len(s.data))

		for k, v := range mask {
			if v {
				nullMask[k/8] |= 1 << uint(k%8)
			} else {
				nullMask[k/8] &= ^(1 << uint(k%8))
			}
		}

		return GDLSeriesString{
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
func (s GDLSeriesString) MakeNullable() GDLSeries {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __initPackBinVec(len(s.data))
	}
	return s
}

// Get the element at index i.
func (s GDLSeriesString) Get(i int) any {
	return *s.data[i]
}

// Get the element at index i as a string.
func (s GDLSeriesString) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return *s.data[i]
}

// Set the element at index i. The value v must be of type string or NullableString.
func (s GDLSeriesString) Set(i int, v any) GDLSeries {
	if ss, ok := v.(string); ok {
		s.data[i] = s.pool.Get(ss)
	} else if ns, ok := v.(NullableString); ok {
		if ns.Valid {
			s.data[i] = s.pool.Get(ns.Value)
		} else {
			s.data[i] = nil
			s.nullMask[i/8] |= 1 << uint(i%8)
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.Set: provided value %t is not of type string or NullableString", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s GDLSeriesString) Take(start, end, step int) GDLSeries {
	return s
}

func (s GDLSeriesString) Less(i, j int) bool {
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

func (s GDLSeriesString) Swap(i, j int) {
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
func (s GDLSeriesString) Append(v any) GDLSeries {
	switch v := v.(type) {
	case string:
		return s.AppendRaw(v)
	case []string:
		return s.AppendRaw(v)
	case NullableString:
		return s.AppendNullable(v)
	case []NullableString:
		return s.AppendNullable(v)
	case GDLSeriesString:
		return s.AppendSeries(v)
	case GDLSeriesError:
		return v
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.Append: invalid type, %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesString) AppendRaw(v any) GDLSeries {
	if s.isNullable {
		if str, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Get(str))
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if strv, ok := v.([]string); ok {
			for _, str := range strv {
				s.data = append(s.data, s.pool.Get(str))
			}
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesString.AppendRaw: invalid type %T", v)}
		}
	} else {
		if str, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Get(str))
		} else if strv, ok := v.([]string); ok {
			for _, str := range strv {
				s.data = append(s.data, s.pool.Get(str))
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesString.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesString) AppendNullable(v any) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesString.AppendNullable: series is not nullable"}
	}

	if str, ok := v.(NullableString); ok {
		s.data = append(s.data, s.pool.Get(str.Value))
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
			s.data = append(s.data, s.pool.Get(str.Value))
			if !str.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s GDLSeriesString) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesString
	if o, ok = other.(GDLSeriesString); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.AppendSeries: invalid type %T", other)}
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

func (s GDLSeriesString) Data() any {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return data
}

func (s GDLSeriesString) DataAsNullable() any {
	data := make([]NullableString, len(s.data))
	for i, v := range s.data {
		data[i] = NullableString{Valid: !s.IsNull(i), Value: *v}
	}
	return data
}

func (s GDLSeriesString) DataAsString() []string {
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
func (s GDLSeriesString) Cast(t typesys.BaseType, stringPool *StringPool) GDLSeries {
	switch t {
	case typesys.BoolType:
		data := __initPackBinVec(len(s.data))
		nullMask := __initPackBinVec(len(s.data))
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

		return GDLSeriesBool{
			isGrouped:  false,
			isNullable: true,
			sorted:     SORTED_NONE,
			size:       len(s.data),
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}

	case typesys.Int32Type:
		data := make([]int, len(s.data))
		nullMask := __initPackBinVec(len(s.data))
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
						data[i] = d
					}
				}
			}
		} else {
			for i, v := range s.data {
				d, err := strconv.Atoi(*v)
				if err != nil {
					nullMask[i>>3] |= (1 << uint(i%8))
				} else {
					data[i] = d
				}
			}
		}

		return GDLSeriesInt32{
			isGrouped:  false,
			isNullable: true,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}

	case typesys.Float64Type:
		data := make([]float64, len(s.data))
		nullMask := __initPackBinVec(len(s.data))
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

		return GDLSeriesFloat64{
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
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.Cast: invalid type %s", t.ToString())}
	}
}

func (s GDLSeriesString) Copy() GDLSeries {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return NewGDLSeriesString(s.name, s.isNullable, data, s.pool)
}

func (s GDLSeriesString) __getDataPtr() *[]*string {
	return &s.data
}

////////////////////////			SERIES OPERATIONS

// Filters out the elements by the given mask series.
func (s GDLSeriesString) Filter(mask GDLSeriesBool) GDLSeries {
	if mask.size != s.Len() {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return GDLSeriesError{"GDLSeriesString.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()

	data := make([]*string, elementCount)
	var nullMask []uint8

	if s.isNullable {

		nullMask = __initPackBinVec(elementCount)

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
func (s GDLSeriesString) FilterByMask(mask []bool) GDLSeries {
	if len(mask) != len(s.data) {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
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

func (s GDLSeriesString) FilterByIndeces(indeces []int) GDLSeries {
	var data []*string
	var nullMask []uint8

	size := len(indeces)
	data = make([]*string, size)

	if s.isNullable {

		nullMask = __initPackBinVec(size)

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

func (s GDLSeriesString) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	if len(s.data) == 0 {
		return s
	}

	v := f(*(s.data[0]))
	switch v.(type) {
	case bool:

		data := __initPackBinVec(len(s.data))

		chunkLen := len(s.data) / THREADS_NUMBER
		if chunkLen < MINIMUM_PARALLEL_SIZE {
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

		return GDLSeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			size:       len(s.data),
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case int:
		data := make([]int, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f((*s.data[i])).(int)
		}

		return GDLSeriesInt32{
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

		return GDLSeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case string:
		if stringPool == nil {
			return GDLSeriesError{"GDLSeriesString.Map: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = stringPool.Get(f((*s.data[i])).(string))
		}

		s.isGrouped = false
		s.sorted = SORTED_NONE
		s.data = data

		return s
	}

	return GDLSeriesError{fmt.Sprintf("GDLSeriesString.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

type SeriesStringPartition struct {
	seriesSize   int
	partition    map[uint64][]int
	nulls        []int
	indexToGroup []int
}

func (gp SeriesStringPartition) GetSize() int {
	return len(gp.partition)
}

func (gp SeriesStringPartition) GetGroupsCount() int {
	// count := 0
	// for _, s := range gp.partitions {
	// 	for _, g := range s {
	// 		if len(g) > 0 {
	// 			count++
	// 		}
	// 	}
	// }

	// for _, g := range gp.nullGroups {
	// 	if len(g) > 0 {
	// 		count++
	// 	}
	// }
	// return count
	return 0
}

func (gp SeriesStringPartition) GetIndices() map[uint64][]int {
	return gp.partition
}

func (gp SeriesStringPartition) GetValueIndices(val any) []int {
	// if sub >= len(gp.partitions) {
	// 	return nil
	// }

	// if v, ok := val.(*string); ok {
	// 	return gp.partitions[sub][v]
	// }

	return nil
}

func (gp SeriesStringPartition) GetNullIndices() []int {
	return gp.nulls
}

func (gp SeriesStringPartition) GetKeys() any {
	// keysMap := make(map[*string]bool)
	// for p := range gp.partitions {
	// 	for k := range gp.partitions[p] {
	// 		keysMap[k] = true
	// 	}
	// }

	keys := make([]string, 0)
	// for k := range keysMap {
	// 	keys = append(keys, *k)
	// }
	return keys
}

func (s GDLSeriesString) Group() GDLSeries {

	var partition SeriesStringPartition
	if len(s.data) < MINIMUM_PARALLEL_SIZE {
		map_ := make(map[uint64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
		for i, v := range s.data {
			map_[(*(*uint64)(unsafe.Pointer(unsafe.Pointer(v))))] = append(map_[(*(*uint64)(unsafe.Pointer(unsafe.Pointer(v))))], i)
		}

		partition = SeriesStringPartition{
			seriesSize:   s.Len(),
			partition:    map_,
			indexToGroup: make([]int, s.Len()),
		}
	} else {

		// Initialize the maps and the wait groups
		allMaps := make([]map[uint64][]int, THREADS_NUMBER)
		for i := 0; i < THREADS_NUMBER; i++ {
			allMaps[i] = make(map[uint64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
		}

		// Define the worker callback
		worker := func(start, end int, map_ map[uint64][]int) {
			for i := start; i < end; i++ {
				map_[(*(*uint64)(unsafe.Pointer(unsafe.Pointer(s.data[i]))))] = append(map_[(*(*uint64)(unsafe.Pointer(unsafe.Pointer(s.data[i]))))], i)
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(s.data), &allMaps, worker)

		partition = SeriesStringPartition{
			seriesSize:   s.Len(),
			partition:    allMaps[0],
			indexToGroup: make([]int, s.Len()),
		}
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s GDLSeriesString) SubGroup(partition SeriesPartition) GDLSeries {
	var newPartition SeriesStringPartition
	otherIndeces := partition.GetIndices()

	if len(s.data) < MINIMUM_PARALLEL_SIZE {

		map_ := make(map[uint64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

		var newHash uint64
		for k, v := range otherIndeces {
			for _, index := range v {
				newHash = *(*uint64)(unsafe.Pointer(unsafe.Pointer((s.data)[index]))) + HASH_MAGIC_NUMBER + (k << 12) + (k >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}

		newPartition = SeriesStringPartition{
			seriesSize:   s.Len(),
			partition:    map_,
			indexToGroup: make([]int, s.Len()),
		}
	} else {

		// collect all keys
		keys := make([]uint64, len(otherIndeces))
		i := 0
		for k := range otherIndeces {
			keys[i] = k
			i++
		}

		// Initialize the maps and the wait groups
		allMaps := make([]map[uint64][]int, THREADS_NUMBER)
		for i := 0; i < THREADS_NUMBER; i++ {
			allMaps[i] = make(map[uint64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
		}

		// Define the worker callback
		worker := func(start, end int, map_ map[uint64][]int) {
			var newHash uint64
			for _, h := range keys[start:end] {
				for _, index := range otherIndeces[h] {
					newHash = *(*uint64)(unsafe.Pointer(unsafe.Pointer((s.data)[index]))) + HASH_MAGIC_NUMBER + (h << 12) + (h >> 4)
					map_[newHash] = append(map_[newHash], index)
				}
			}
		}

		__series_groupby_multithreaded(THREADS_NUMBER, len(keys), &allMaps, worker)

		newPartition = SeriesStringPartition{
			seriesSize:   s.Len(),
			partition:    allMaps[0],
			indexToGroup: make([]int, s.Len()),
		}
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s GDLSeriesString) GetPartition() SeriesPartition {
	return s.partition
}

func (s GDLSeriesString) Sort() GDLSeries {
	return s
}

func (s GDLSeriesString) SortRev() GDLSeries {
	return s
}

////////////////////////			SORTING OPERATIONS

////////////////////////			ARITHMETIC OPERATIONS

func (s GDLSeriesString) Mul(other GDLSeries) GDLSeries {
	return s
}

func (s GDLSeriesString) Add(other GDLSeries) GDLSeries {
	return s
}
