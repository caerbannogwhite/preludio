package gandalff

import (
	"fmt"
	"typesys"
)

// GDLSeriesFloat64 represents a series of floats.
type GDLSeriesFloat64 struct {
	isGrouped  bool
	isNullable bool
	sorted     GDLSeriesSortOrder
	name       string
	data       []float64
	nullMask   []uint8
	partition  *GDLSeriesFloat64Partition
}

func NewGDLSeriesFloat64(name string, isNullable bool, makeCopy bool, data []float64) GDLSeriesFloat64 {
	var nullMask []uint8
	if isNullable {
		nullMask = __initNullMask(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GDLSeriesFloat64{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////////

// Returns the number of elements in the series.
func (s GDLSeriesFloat64) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s GDLSeriesFloat64) Name() string {
	return s.name
}

// Returns the type of the series.
func (s GDLSeriesFloat64) Type() typesys.BaseType {
	return typesys.Float64Type
}

// Returns if the series is grouped.
func (s GDLSeriesFloat64) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s GDLSeriesFloat64) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s GDLSeriesFloat64) IsSorted() GDLSeriesSortOrder {
	return s.sorted
}

// Returns if the series has null values.
func (s GDLSeriesFloat64) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s GDLSeriesFloat64) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8; i++ {
			count += int((v & (1 << uint(i))) >> uint(i))
		}
	}
	return count
}

// Returns the number of non-null values in the series.
func (s GDLSeriesFloat64) NonNullCount() int {
	return s.Len() - s.NullCount()
}

// Returns if the element at index i is null.
func (s GDLSeriesFloat64) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s GDLSeriesFloat64) SetNull(i int) GDLSeries {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	} else {
		nullMask := __initNullMask(len(s.data))
		nullMask[i/8] |= 1 << uint(i%8)

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s GDLSeriesFloat64) GetNullMask() []bool {
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
func (s GDLSeriesFloat64) SetNullMask(mask []bool) GDLSeries {
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
		nullMask := __initNullMask(len(s.data))
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
func (s GDLSeriesFloat64) MakeNullable() GDLSeries {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __initNullMask(len(s.data))
	}
	return s
}

// Get the element at index i.
func (s GDLSeriesFloat64) Get(i int) any {
	return s.data[i]
}

// Get the element at index i as a string.
func (s GDLSeriesFloat64) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return floatToString(s.data[i])
}

// Set the element at index i. The type of v must be float64 or NullableFloat64.
func (s GDLSeriesFloat64) Set(i int, v any) GDLSeries {
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
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.Set: provided value %t is not of type float64 or NullableFloat64", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s GDLSeriesFloat64) Take(start, end, step int) GDLSeries {
	return s
}

func (s GDLSeriesFloat64) Less(i, j int) bool {
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

func (s GDLSeriesFloat64) Swap(i, j int) {
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

func (s GDLSeriesFloat64) Append(v any) GDLSeries {
	switch v := v.(type) {
	case float64:
		return s.AppendRaw(v)
	case []float64:
		return s.AppendRaw(v)
	case NullableFloat64:
		return s.AppendNullable(v)
	case []NullableFloat64:
		return s.AppendNullable(v)
	case GDLSeriesFloat64:
		return s.AppendSeries(v)
	case GDLSeriesError:
		return v
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.Append: invalid type, %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesFloat64) AppendRaw(v any) GDLSeries {
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
			return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendRaw: invalid type %T", v)}
		}
	} else {
		if f, ok := v.(float64); ok {
			s.data = append(s.data, f)
		} else if fv, ok := v.([]float64); ok {
			s.data = append(s.data, fv...)
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesFloat64) AppendNullable(v any) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesFloat64.AppendNullable: series is not nullable"}
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
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendNullable: invalid type %T", v)}
	}

	return s
}

func (s GDLSeriesFloat64) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesFloat64
	if o, ok = other.(GDLSeriesFloat64); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendSeries: invalid type %T", other)}
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

///////////////////////////////		ALL DATA ACCESSORS			/////////////////////////

func (s GDLSeriesFloat64) Data() any {
	return s.data
}

func (s GDLSeriesFloat64) DataAsNullable() any {
	data := make([]NullableFloat64, len(s.data))
	for i, v := range s.data {
		data[i] = NullableFloat64{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GDLSeriesFloat64) DataAsString() []string {
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
func (s GDLSeriesFloat64) Cast(t typesys.BaseType, stringPool *StringPool) GDLSeries {
	switch t {
	case typesys.BoolType:
		data := __initNullMask(len(s.data))
		for i, v := range s.data {
			if v != 0 {
				data[i>>3] |= (1 << uint(i%8))
			}
		}

		return GDLSeriesBool{
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
		data := make([]int, len(s.data))
		for i, v := range s.data {
			data[i] = int(v)
		}

		return GDLSeriesInt32{
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
			return GDLSeriesError{"GDLSeriesFloat64.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = stringPool.Get(NULL_STRING)
				} else {
					data[i] = stringPool.Get(floatToString(v))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = stringPool.Get(floatToString(v))
			}
		}

		return GDLSeriesString{
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
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.Cast: invalid type %s", t.ToString())}
	}
}

func (s GDLSeriesFloat64) Copy() GDLSeries {
	data := make([]float64, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return GDLSeriesFloat64{
		isGrouped: s.isGrouped, sorted: s.sorted, isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

///////////////////////////////		SERIES OPERATIONS			/////////////////////////

// Filters out the elements by the given mask series.
func (s GDLSeriesFloat64) Filter(mask GDLSeriesBool) GDLSeries {
	if mask.size != s.Len() {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return GDLSeriesError{"GDLSeriesFloat64.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()
	var nullMask []uint8

	data := make([]float64, elementCount)
	if s.isNullable {

		nullMask = __initNullMask(elementCount)

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
func (s GDLSeriesFloat64) FilterByMask(mask []bool) GDLSeries {
	if len(mask) != len(s.data) {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
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

		nullMask = __initNullMask(elementCount)

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

func (s GDLSeriesFloat64) FilterByIndeces(indexes []int) GDLSeries {
	var data []float64
	var nullMask []uint8

	size := len(indexes)
	data = make([]float64, size)

	if s.isNullable {

		nullMask = __initNullMask(size)

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

func (s GDLSeriesFloat64) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:

		// Not a null mask but you get the same result
		data := __initNullMask(len(s.data))
		for i := 0; i < len(s.data); i++ {
			if f(s.data[i]).(bool) {
				data[i>>3] |= (1 << uint(i%8))
			}
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
			data[i] = f(s.data[i]).(int)
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
			data[i] = f(s.data[i]).(float64)
		}

		s.isGrouped = false
		s.sorted = SORTED_NONE
		s.data = data

		return s

	case string:
		if stringPool == nil {
			return GDLSeriesError{"GDLSeriesFloat64.Map: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = stringPool.Get(f(s.data[i]).(string))
		}

		return GDLSeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       stringPool,
		}
	}

	return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.Map: Unsupported type %T", v)}
}

/////////////////////////////// 		GROUPING OPERATIONS		/////////////////////////

type GDLSeriesFloat64Partition struct {
	partition []map[float64][]int
	nullGroup [][]int
}

func (p GDLSeriesFloat64Partition) GetSize() int {
	return len(p.partition)
}

func (p GDLSeriesFloat64Partition) GetGroupsCount() int {
	count := 0
	for _, s := range p.partition {
		for _, g := range s {
			if len(g) > 0 {
				count++
			}
		}
	}

	for _, g := range p.nullGroup {
		if len(g) > 0 {
			count++
		}
	}
	return count
}

func (p GDLSeriesFloat64Partition) GetIndices() [][]int {
	indices := make([][]int, 0)

	for _, s := range p.partition {
		for _, g := range s {
			if len(g) > 0 {
				indices = append(indices, g)
			}
		}
	}

	for _, g := range p.nullGroup {
		if len(g) > 0 {
			indices = append(indices, g)
		}
	}

	return indices
}

func (p GDLSeriesFloat64Partition) GetValueIndices(sub int, val interface{}) []int {
	if sub >= len(p.partition) {
		return nil
	}

	if v, ok := val.(float64); ok {
		return p.partition[sub][v]
	}

	return nil
}

func (s GDLSeriesFloat64Partition) GetNullIndices(sub int) []int {
	if sub >= len(s.nullGroup) {
		return nil
	}

	return s.nullGroup[sub]
}

func (s GDLSeriesFloat64) Group() GDLSeries {
	var nullGroup [][]int

	groups := make(map[float64][]int)
	if s.isNullable {
		nullGroup = make([][]int, 1)
		nullGroup[0] = make([]int, 0)

		for i, v := range s.data {
			if s.IsNull(i) {
				nullGroup[0] = append(nullGroup[0], i)
			} else {
				groups[v] = append(groups[v], i)
			}
		}
	} else {
		for i, v := range s.data {
			groups[v] = append(groups[v], i)
		}
	}
	return GDLSeriesFloat64{
		isGrouped:  true,
		isNullable: s.isNullable,
		sorted:     s.sorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition:  &GDLSeriesFloat64Partition{partition: []map[float64][]int{groups}, nullGroup: nullGroup},
	}
}

func (s GDLSeriesFloat64) SubGroup(partition GDLSeriesPartition) GDLSeries {
	var nullGroup [][]int

	groups := make([]map[float64][]int, 0)
	indices := partition.GetIndices()
	if s.isNullable {
		nullGroup = make([][]int, partition.GetGroupsCount())

		for _, g := range indices {
			groups = append(groups, make(map[float64][]int))
			for _, i := range g {
				if s.IsNull(i) {
					nullGroup[i] = append(nullGroup[i], i)
				} else {
					groups[i][s.data[i]] = append(groups[i][s.data[i]], i)
				}
			}
		}
	} else {
		for gi, g := range indices {
			groups = append(groups, make(map[float64][]int))
			for _, idx := range g {
				if groups[gi][s.data[idx]] == nil {
					groups[gi][s.data[idx]] = make([]int, 0)
				}
				groups[gi][s.data[idx]] = append(groups[gi][s.data[idx]], idx)
			}
		}
	}

	return GDLSeriesFloat64{
		isGrouped:  true,
		isNullable: s.isNullable,
		sorted:     s.sorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition:  &GDLSeriesFloat64Partition{groups, nullGroup},
	}
}

func (s GDLSeriesFloat64) GetPartition() GDLSeriesPartition {
	return s.partition
}

func (s GDLSeriesFloat64) Sort() GDLSeries {
	return s
}

func (s GDLSeriesFloat64) SortRev() GDLSeries {
	return s
}

///////////////////////////////		SORTING OPERATIONS		/////////////////////////////

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////

func (s GDLSeriesFloat64) Mul(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), other.Type().ToString())}
}

func (s GDLSeriesFloat64) Div(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesFloat64) Mod(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesFloat64) Add(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesFloat64) Sub(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), other.Type().ToString())}

}
