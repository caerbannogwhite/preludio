package gandalff

import (
	"fmt"
	"sort"
	"typesys"
)

// GDLSeriesInt32 represents a series of ints.
type GDLSeriesInt32 struct {
	isNullable bool
	name       string
	data       []int
	nullMask   []uint8
}

func NewGDLSeriesInt32(name string, isNullable bool, makeCopy bool, data []int) GDLSeriesInt32 {
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

func (s GDLSeriesInt32) Len() int {
	return len(s.data)
}

func (s GDLSeriesInt32) IsNullable() bool {
	return s.isNullable
}

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

func (s GDLSeriesInt32) Name() string {
	return s.name
}

func (s GDLSeriesInt32) Type() typesys.BaseType {
	return typesys.Int32Type
}

func (s GDLSeriesInt32) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesInt32) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8; i++ {
			count += int((v & (1 << uint(i))) >> uint(i))
		}
	}
	return count
}

func (s GDLSeriesInt32) NonNullCount() int {
	return s.Len() - s.NullCount()
}

func (s GDLSeriesInt32) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesInt32) SetNull(i int) error {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	}
	return fmt.Errorf("GDLSeriesInt32.SetNull: series is not nullable")
}

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

func (s GDLSeriesInt32) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return fmt.Errorf("GDLSeriesInt32.SetNullMask: series is not nullable")
	}

	for k, v := range mask {
		if v {
			s.nullMask[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMask[k/8] &= ^(1 << uint(k%8))
		}
	}

	return nil
}

func (s GDLSeriesInt32) Get(i int) interface{} {
	return s.data[i]
}

func (s GDLSeriesInt32) Set(i int, v interface{}) {
	s.data[i] = v.(int)
}

func (s GDLSeriesInt32) Less(i, j int) bool {
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

func (s GDLSeriesInt32) Append(v interface{}) GDLSeries {
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
func (s GDLSeriesInt32) AppendRaw(v interface{}) GDLSeries {
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
func (s GDLSeriesInt32) AppendNullable(v interface{}) GDLSeries {
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

func (s GDLSeriesInt32) Data() interface{} {
	return s.data
}

func (s GDLSeriesInt32) NullableData() interface{} {
	data := make([]NullableInt32, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt32{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GDLSeriesInt32) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		if s.IsNull(i) {
			data[i] = NULL_STRING
		} else {
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

	return GDLSeriesInt32{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
	}
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

	return GDLSeriesInt32{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
	}
}

func (s GDLSeriesInt32) Map(f GDLMapFunc, stringPool *StringPool) GDLSeries {
	return s
}

/////////////////////////////// 		GROUPING OPERATIONS		/////////////////////////

type GDLSeriesInt32Partition struct {
	partition []map[int][]int
	nullGroup [][]int
}

func (p GDLSeriesInt32Partition) GetSize() int {
	return len(p.partition)
}

func (p GDLSeriesInt32Partition) GetGroupsCount() int {
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

func (p GDLSeriesInt32Partition) GetIndices() [][]int {
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

func (p GDLSeriesInt32Partition) GetValueIndices(sub int, val interface{}) []int {
	if sub >= len(p.partition) {
		return nil
	}

	if v, ok := val.(int); ok {
		return p.partition[sub][v]
	}

	return nil
}

func (s GDLSeriesInt32Partition) GetNullIndices(sub int) []int {
	if sub >= len(s.nullGroup) {
		return nil
	}

	return s.nullGroup[sub]
}

func (s GDLSeriesInt32) Group() GDLSeriesPartition {
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
	return GDLSeriesInt32Partition{
		partition: []map[int][]int{groups},
		nullGroup: nullGroup,
	}
}

func (s GDLSeriesInt32) SubGroup(partition GDLSeriesPartition) GDLSeriesPartition {
	var nullGroup [][]int

	groups := make([]map[int][]int, 0)
	indices := partition.GetIndices()
	if s.isNullable {
		nullGroup = make([][]int, partition.GetGroupsCount())

		for _, g := range indices {
			groups = append(groups, make(map[int][]int))
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
			groups = append(groups, make(map[int][]int))
			for _, idx := range g {
				if groups[gi][s.data[idx]] == nil {
					groups[gi][s.data[idx]] = make([]int, 0)
				}
				groups[gi][s.data[idx]] = append(groups[gi][s.data[idx]], idx)
			}
		}
	}
	return GDLSeriesInt32Partition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

func (s GDLSeriesInt32) Sort() GDLSeries {
	data := make([]int, len(s.data))
	copy(data, s.data)
	sort.Ints(data)
	return GDLSeriesInt32{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   s.nullMask,
	}
}

func (s GDLSeriesInt32) SortDesc() GDLSeries {
	data := make([]int, len(s.data))
	copy(data, s.data)
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	return GDLSeriesInt32{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   s.nullMask,
	}
}

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////

func (s GDLSeriesInt32) Mul(other GDLSeries) GDLSeries {
	switch o := other.(type) {
	case GDLSeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						result := s.data
						result[0] = s.data[0] * o.data[0]
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						result := s.data
						result[0] = s.data[0] * o.data[0]
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						result := s.data
						result[0] = s.data[0] * o.data[0]
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						result := s.data
						result[0] = s.data[0] * o.data[0]
						return GDLSeriesInt32{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int, resultSize)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(o.data)
						result := make([]int, resultSize)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int, resultSize)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(o.data)
						result := make([]int, resultSize)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return GDLSeriesInt32{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(s.data)
						result := s.data
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return GDLSeriesInt32{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			}
		}
	case GDLSeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						result := o.data
						result[0] = float64(s.data[0]) * o.data[0]
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						result := o.data
						result[0] = float64(s.data[0]) * o.data[0]
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						result := o.data
						result[0] = float64(s.data[0]) * o.data[0]
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						result := o.data
						result[0] = float64(s.data[0]) * o.data[0]
						return GDLSeriesFloat64{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, len(s.data))
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(s.data)
						result := make([]float64, len(s.data))
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, len(s.data))
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(s.data)
						result := make([]float64, len(s.data))
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return GDLSeriesFloat64{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: true, name: s.name, data: result, nullMask: s.nullMask}
					} else {
						resultSize := len(o.data)
						result := o.data
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return GDLSeriesFloat64{isNullable: false, name: s.name, data: result, nullMask: s.nullMask}
					}
				}
			}
		}
	}
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
