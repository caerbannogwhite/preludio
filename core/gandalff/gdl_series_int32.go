package gandalff

import (
	"fmt"
	"sort"
	"typesys"
)

// GDLSeriesInt32 represents a series of ints.
type GDLSeriesInt32 struct {
	isGrouped  bool
	isNullable bool
	isSorted   bool
	name       string
	data       []int
	nullMask   []uint8
	partition  *GDLSeriesInt32Grouping
}

func NewGDLSeriesInt32(name string, isNullable bool, makeCopy bool, data []int) GDLSeries {
	var nullMask []uint8
	if isNullable {
		if len(data)%8 == 0 {
			nullMask = make([]uint8, (len(data) >> 3))
		} else {
			nullMask = make([]uint8, (len(data)>>3)+1)
		}
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GDLSeriesInt32{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////

// Returns the number of elements in the series.
func (s GDLSeriesInt32) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s GDLSeriesInt32) Name() string {
	return s.name
}

// Returns the type of the series.
func (s GDLSeriesInt32) Type() typesys.BaseType {
	return typesys.Int32Type
}

// Returns if the series is grouped.
func (s GDLSeriesInt32) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s GDLSeriesInt32) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s GDLSeriesInt32) IsSorted() bool {
	return s.isSorted
}

// Returns if the series has null values.
func (s GDLSeriesInt32) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s GDLSeriesInt32) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8; i++ {
			count += int((v & (1 << uint(i))) >> uint(i))
		}
	}
	return count
}

// Returns the number of non-null values in the series.
func (s GDLSeriesInt32) NonNullCount() int {
	return s.Len() - s.NullCount()
}

// Returns if the element at index i is null.
func (s GDLSeriesInt32) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s GDLSeriesInt32) SetNull(i int) GDLSeries {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return s
	} else {
		var nullMask []uint8
		if len(s.data)%8 == 0 {
			nullMask = make([]uint8, (len(s.data) >> 3))
		} else {
			nullMask = make([]uint8, (len(s.data)>>3)+1)
		}

		nullMask[i/8] |= 1 << uint(i%8)

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s GDLSeriesInt32) GetNullMask() []bool {
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
func (s GDLSeriesInt32) SetNullMask(mask []bool) GDLSeries {
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
		var nullMask []uint8
		if len(s.data)%8 == 0 {
			nullMask = make([]uint8, (len(s.data) >> 3))
		} else {
			nullMask = make([]uint8, (len(s.data)>>3)+1)
		}

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
func (s GDLSeriesInt32) MakeNullable() GDLSeries {
	if !s.isNullable {
		if len(s.data)%8 == 0 {
			s.nullMask = make([]uint8, (len(s.data) >> 3))
		} else {
			s.nullMask = make([]uint8, (len(s.data)>>3)+1)
		}
		s.isNullable = true
	}
	return s
}

// Get the element at index i.
func (s GDLSeriesInt32) Get(i int) any {
	return s.data[i]
}

// Get the element at index i as a string.
func (s GDLSeriesInt32) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return intToString(s.data[i])
}

// Set the element at index i. The value v must be of type int or NullableInt32.
func (s GDLSeriesInt32) Set(i int, v any) GDLSeries {
	if ii, ok := v.(int); ok {
		s.data[i] = ii
	} else if ni, ok := v.(NullableInt32); ok {
		if ni.Valid {
			s.data[i] = ni.Value
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.Set: provided value %t is not of type int or NullableInt32", v)}
	}

	s.isSorted = false
	return s
}

func (s GDLSeriesInt32) Less(i, j int) bool {
	if s.isGrouped {
		if s.partition.indexToGroup[i] != s.partition.indexToGroup[j] {
			return s.partition.indexToGroup[i] < s.partition.indexToGroup[j]
		}
		return s.data[i] < s.data[j]
	} else

	// if s is grouped the null element are is the same group
	// so there is no need to check if the element is null
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

func (s GDLSeriesInt32) Swap(i, j int) {
	if s.isGrouped {
		s.partition.indexToGroup[i], s.partition.indexToGroup[j] = s.partition.indexToGroup[j], s.partition.indexToGroup[i]
	}

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

func (s GDLSeriesInt32) Append(v any) GDLSeries {
	switch v := v.(type) {
	case int:
		return s.AppendRaw(v)
	case []int:
		return s.AppendRaw(v)
	case NullableInt32:
		return s.AppendNullable(v)
	case []NullableInt32:
		return s.AppendNullable(v)
	case GDLSeriesInt32:
		return s.AppendSeries(v)
	case GDLSeriesError:
		return v
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesInt32) AppendRaw(v any) GDLSeries {
	if s.isNullable {
		if i, ok := v.(int); ok {
			s.data = append(s.data, i)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if iv, ok := v.([]int); ok {
			s.data = append(s.data, iv...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendRaw: invalid type %T", v)}
		}
	} else {
		if b, ok := v.(int); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]int); ok {
			s.data = append(s.data, bv...)
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesInt32) AppendNullable(v any) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesInt32.AppendNullable: series is not nullable"}
	}

	if b, ok := v.(NullableInt32); ok {
		s.data = append(s.data, b.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !b.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableInt32); ok {
		if len(s.data) > len(s.nullMask)<<8 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for _, b := range bv {
			s.data = append(s.data, b.Value)
			if !b.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendNullable: invalid type %T", v)}
	}

	return s
}

// AppendSeries appends a series to the series.
func (s GDLSeriesInt32) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesInt32
	if o, ok = other.(GDLSeriesInt32); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.AppendSeries: invalid type %T", other)}
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

///////////////////////////////  	ALL DATA ACCESSORS  /////////////////////////////////

func (s GDLSeriesInt32) Data() any {
	return s.data
}

func (s GDLSeriesInt32) DataAsNullable() any {
	data := make([]NullableInt32, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt32{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GDLSeriesInt32) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = intToString(v)
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = intToString(v)
		}
	}
	return data
}

func (s GDLSeriesInt32) Copy() GDLSeries {
	data := make([]int, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return GDLSeriesInt32{isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

///////////////////////////////  	SERIES OPERATIONS  //////////////////////////////////

// Filters out the elements by the given mask series.
func (s GDLSeriesInt32) Filter(mask GDLSeriesBool) GDLSeries {
	if mask.size != s.Len() {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.Filter: mask length (%d) does not match series length (%d)", mask.size, s.Len())}
	}

	if mask.isNullable {
		return GDLSeriesError{"GDLSeriesInt32.Filter: mask series cannot be nullable for this operation"}
	}

	elementCount := mask.__trueCount()
	var nullMask []uint8

	data := make([]int, elementCount)
	if s.isNullable {

		if elementCount%8 == 0 {
			nullMask = make([]uint8, (elementCount >> 3))
		} else {
			nullMask = make([]uint8, (elementCount>>3)+1)
		}

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
func (s GDLSeriesInt32) FilterByMask(mask []bool) GDLSeries {
	if len(mask) != len(s.data) {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []int
	var nullMask []uint8

	data = make([]int, elementCount)

	if s.isNullable {

		if elementCount%8 == 0 {
			nullMask = make([]uint8, (elementCount >> 3))
		} else {
			nullMask = make([]uint8, (elementCount>>3)+1)
		}

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

func (s GDLSeriesInt32) FilterByIndeces(indexes []int) GDLSeries {
	var data []int
	var nullMask []uint8

	size := len(indexes)
	data = make([]int, size)

	if s.isNullable {

		if size%8 == 0 {
			nullMask = make([]uint8, (size >> 3))
		} else {
			nullMask = make([]uint8, (size>>3)+1)
		}

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

func (s GDLSeriesInt32) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:

		var data []uint8
		if len(s.data)%8 == 0 {
			data = make([]uint8, (len(s.data) >> 3))
		} else {
			data = make([]uint8, (len(s.data)>>3)+1)
		}

		for i := 0; i < len(s.data); i++ {
			if f(s.data[i]).(bool) {
				data[i>>3] |= (1 << uint(i%8))
			}
		}

		return GDLSeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			isSorted:   false,
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

		s.isGrouped = false
		s.isSorted = false
		s.data = data

		return s

	case float64:
		data := make([]float64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(float64)
		}

		return GDLSeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			isSorted:   false,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
		}

	case string:
		if stringPool == nil {
			return GDLSeriesError{"GDLSeriesInt32.Map: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = stringPool.Add(f(s.data[i]).(string))
		}

		return GDLSeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			isSorted:   false,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       stringPool,
		}
	}

	return GDLSeriesError{fmt.Sprintf("GDLSeriesInt32.Map: Unsupported type %T", v)}
}

/////////////////////////////// 		GROUPING OPERATIONS		/////////////////////////

type GDLSeriesInt32Grouping struct {
	seriesSize   int
	partitions   []map[int][]int
	nullGroups   [][]int
	indexToGroup []int
}

func (gp GDLSeriesInt32Grouping) GetSize() int {
	return len(gp.partitions)
}

func (gp GDLSeriesInt32Grouping) beginSorting() GDLSeriesInt32Grouping {
	gp.indexToGroup = make([]int, gp.seriesSize)
	for i, part := range gp.partitions {
		for _, g := range part {
			for _, idx := range g {
				gp.indexToGroup[idx] = i
			}
		}
	}

	for i, g := range gp.nullGroups {
		for _, idx := range g {
			gp.indexToGroup[idx] = i + len(gp.partitions)
		}
	}

	return gp
}

func (gp GDLSeriesInt32Grouping) endSorting() GDLSeriesInt32Grouping {
	newPartitions := make(map[int][]int, len(gp.partitions))
	newNullGroups := make([][]int, len(gp.nullGroups))

	for i, part := range gp.partitions {
		newPartitions[i] = make([]int, 0, len(part))
	}

	for i, g := range gp.nullGroups {
		newNullGroups[i] = make([]int, 0, len(g))
	}

	for i, g := range gp.indexToGroup {
		if g < len(gp.partitions) {
			newPartitions[g] = append(newPartitions[g], i)
		} else {
			newNullGroups[g-len(gp.partitions)] = append(newNullGroups[g-len(gp.partitions)], i)
		}
	}

	gp.indexToGroup = nil
	return gp
}

func (gp GDLSeriesInt32Grouping) GetGroupsCount() int {
	count := 0
	for _, s := range gp.partitions {
		for _, g := range s {
			if len(g) > 0 {
				count++
			}
		}
	}

	for _, g := range gp.nullGroups {
		if len(g) > 0 {
			count++
		}
	}
	return count
}

func (gp GDLSeriesInt32Grouping) GetIndices() [][]int {
	indices := make([][]int, 0)

	for _, s := range gp.partitions {
		for _, g := range s {
			if len(g) > 0 {
				indices = append(indices, g)
			}
		}
	}

	for _, g := range gp.nullGroups {
		if len(g) > 0 {
			indices = append(indices, g)
		}
	}

	return indices
}

func (gp GDLSeriesInt32Grouping) GetValueIndices(sub int, val interface{}) []int {
	if sub >= len(gp.partitions) {
		return nil
	}

	if v, ok := val.(int); ok {
		return gp.partitions[sub][v]
	}

	return nil
}

func (gp GDLSeriesInt32Grouping) GetNullIndices(sub int) []int {
	if sub >= len(gp.nullGroups) {
		return nil
	}

	return gp.nullGroups[sub]
}

func (s GDLSeriesInt32) Group() GDLSeries {
	var nullGroup [][]int

	groups := make(map[int][]int)
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

	partition := GDLSeriesInt32Grouping{
		seriesSize: s.Len(),
		partitions: []map[int][]int{groups},
		nullGroups: nullGroup,
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s GDLSeriesInt32) SubGroup(partitions GDLSeriesPartition) GDLSeries {
	var nullGroups [][]int

	embeddedPartitions := make([]map[int][]int, partitions.GetGroupsCount())
	indices := partitions.GetIndices()
	if s.isNullable {
		nullGroups = make([][]int, partitions.GetGroupsCount())

		for gi, g := range indices {

			// initialize embedded partitions
			embeddedPartitions[gi] = make(map[int][]int)
			nullGroups[gi] = make([]int, 0)

			for _, idx := range g {
				if s.IsNull(idx) {
					nullGroups[gi] = append(nullGroups[gi], idx)
				} else {
					if embeddedPartitions[gi][s.data[idx]] == nil {
						embeddedPartitions[gi][s.data[idx]] = make([]int, 0)
					}
					embeddedPartitions[gi][s.data[idx]] = append(embeddedPartitions[gi][s.data[idx]], idx)
				}
			}
		}
	} else {
		// for gi, g := range indices {
		// 	groups = append(groups, make(map[int][]int))
		// 	for _, idx := range g {
		// 		if groups[gi][s.data[idx]] == nil {
		// 			groups[gi][s.data[idx]] = make([]int, 0)
		// 		}
		// 		groups[gi][s.data[idx]] = append(groups[gi][s.data[idx]], idx)
		// 	}
		// }
	}

	newPartition := GDLSeriesInt32Grouping{
		seriesSize: s.Len(),
		partitions: embeddedPartitions,
		nullGroups: nullGroups,
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s GDLSeriesInt32) GetPartition() GDLSeriesPartition {
	return s.partition
}

/////////////////////////////// 		SORTING OPERATIONS		/////////////////////////

func (s GDLSeriesInt32) Sort() GDLSeries {
	if !s.isSorted {
		if s.isGrouped {
			*s.partition = (*s.partition).beginSorting()
			sort.Sort(s)
			*s.partition = (*s.partition).endSorting()
		} else {
			sort.Sort(s)
		}
		s.isSorted = true
	}
	return s
}

func (s GDLSeriesInt32) SortRev() GDLSeries {
	if !s.isSorted {
		sort.Sort(sort.Reverse(s))
		s.isSorted = true
	}
	return s
}

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////

func (s GDLSeriesInt32) Mul(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt32) Div(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt32) Mod(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt32) Add(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), other.Type().ToString())}

}

func (s GDLSeriesInt32) Sub(other GDLSeries) GDLSeries {
	return GDLSeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), other.Type().ToString())}

}
