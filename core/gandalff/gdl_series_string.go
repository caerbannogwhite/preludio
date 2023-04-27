package gandalff

import (
	"fmt"
	"strings"
	"typesys"
)

// GDLSeriesString represents a series of strings.
type GDLSeriesString struct {
	isGrouped  bool
	isNullable bool
	isSorted   bool
	name       string
	data       []*string
	nullMask   []uint8
	pool       *StringPool
	partition  GDLSeriesStringPartition
}

func NewGDLSeriesString(name string, isNullable bool, data []string, pool *StringPool) GDLSeries {
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

	actualData := make([]*string, len(data))
	for i, v := range data {
		actualData[i] = pool.Get(v)
	}

	return GDLSeriesString{isNullable: isNullable, name: name, data: actualData, nullMask: nullMask, pool: pool}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////

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
func (s GDLSeriesString) IsSorted() bool {
	return s.isSorted
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

func (s GDLSeriesString) SetNull(i int) GDLSeries {
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

		return GDLSeriesString{
			isGrouped:  s.isGrouped,
			isNullable: true,
			isSorted:   s.isSorted,
			name:       s.name,
			data:       s.data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  s.partition,
		}
	}
}

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

func (s GDLSeriesString) MakeNullable() GDLSeries {
	if !s.isNullable {
		s.isNullable = true
		if len(s.data)%8 == 0 {
			s.nullMask = make([]uint8, (len(s.data) >> 3))
		} else {
			s.nullMask = make([]uint8, (len(s.data)>>3)+1)
		}
	}
	return s
}

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

		return GDLSeriesString{
			isGrouped:  s.isGrouped,
			isNullable: true,
			isSorted:   s.isSorted,
			name:       s.name,
			data:       s.data,
			nullMask:   nullMask,
			pool:       s.pool,
			partition:  s.partition,
		}
	}
}

func (s GDLSeriesString) Get(i int) any {
	return *s.data[i]
}

func (s GDLSeriesString) Set(i int, v any) {
	s.data[i] = s.pool.Get(v.(string))
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

///////////////////////////////		ALL DATA ACCESSORS		/////////////////////////////

func (s GDLSeriesString) Data() any {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return data
}

func (s GDLSeriesString) NullableData() any {
	data := make([]NullableString, len(s.data))
	for i, v := range s.data {
		data[i] = NullableString{Valid: !s.IsNull(i), Value: *v}
	}
	return data
}

func (s GDLSeriesString) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		if s.IsNull(i) {
			data[i] = NULL_STRING
		} else {
			data[i] = *v
		}
	}
	return data
}

func (s GDLSeriesString) Copy() GDLSeries {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return NewGDLSeriesString(s.name, s.isNullable, data, s.pool)
}

/////////////////////////////// 		SERIES OPERATIONS		/////////////////////////

// FilterByMask returns a new series with elements filtered by the mask.
func (s GDLSeriesString) FilterByMask(mask []bool) GDLSeries {
	if len(mask) != len(s.data) {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.FilterByMask: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []*string
	var nullMask []uint8

	data = make([]*string, elementCount)

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

	return GDLSeriesString{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
	}
}

func (s GDLSeriesString) FilterByIndeces(indexes []int) GDLSeries {
	var data []*string
	var nullMask []uint8

	size := len(indexes)
	data = make([]*string, size)

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

	return GDLSeriesString{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
	}
}

func (s GDLSeriesString) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	return s
}

/////////////////////////////// 		GROUPING OPERATIONS		/////////////////////////

type GDLSeriesStringPartition struct {
	partition []map[*string][]int
	nullGroup [][]int
}

func (p GDLSeriesStringPartition) GetSize() int {
	return len(p.partition)
}

func (p GDLSeriesStringPartition) GetGroupsCount() int {
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

func (p GDLSeriesStringPartition) GetIndices() [][]int {
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

func (p GDLSeriesStringPartition) GetValueIndices(sub int, val interface{}) []int {
	if sub >= len(p.partition) {
		return nil
	}

	if v, ok := val.(*string); ok {
		return p.partition[sub][v]
	}

	return nil
}

func (s GDLSeriesStringPartition) GetNullIndices(sub int) []int {
	if sub >= len(s.nullGroup) {
		return nil
	}

	return s.nullGroup[sub]
}

func (s GDLSeriesString) Group() GDLSeries {
	var nullGroup [][]int

	groups := make(map[*string][]int)
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

	partition := GDLSeriesStringPartition{
		partition: []map[*string][]int{groups},
		nullGroup: nullGroup,
	}

	return GDLSeriesString{
		isGrouped:  true,
		isNullable: s.isNullable,
		isSorted:   s.isSorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition:  partition,
		pool:       s.pool,
	}
}

func (s GDLSeriesString) SubGroup(partition GDLSeriesPartition) GDLSeries {
	var nullGroup [][]int

	groups := make([]map[*string][]int, 0)
	indices := partition.GetIndices()
	if s.isNullable {
		nullGroup = make([][]int, partition.GetGroupsCount())

		for _, g := range indices {
			groups = append(groups, make(map[*string][]int))
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
			groups = append(groups, make(map[*string][]int))
			for _, idx := range g {
				if groups[gi][s.data[idx]] == nil {
					groups[gi][s.data[idx]] = make([]int, 0)
				}
				groups[gi][s.data[idx]] = append(groups[gi][s.data[idx]], idx)
			}
		}
	}

	newPartition := GDLSeriesStringPartition{
		partition: groups,
		nullGroup: nullGroup,
	}

	return GDLSeriesString{
		isGrouped:  true,
		isNullable: s.isNullable,
		isSorted:   s.isSorted,
		name:       s.name,
		data:       s.data,
		nullMask:   s.nullMask,
		partition:  newPartition,
		pool:       s.pool,
	}
}

func (s GDLSeriesString) GetPartition() GDLSeriesPartition {
	return s.partition
}

func (s GDLSeriesString) Sort() GDLSeries {
	return s
}

///////////////////////////////		SORTING OPERATIONS		/////////////////////////////

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////

func (s GDLSeriesString) Mul(other GDLSeries) GDLSeries {
	return s
}

func (s GDLSeriesString) Add(other GDLSeries) GDLSeries {
	return s
}
