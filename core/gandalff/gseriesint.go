package gandalff

import "errors"

// GSeriesInt represents a series of ints.
type GSeriesInt struct {
	isNullable bool
	name       string
	data       []int
	nullMap    []uint8
}

func NewGSeriesInt(name string, isNullable bool, makeCopy bool, data []int) GSeriesInt {
	var nullMap []uint8
	if isNullable {
		if len(data)%8 == 0 {
			nullMap = make([]uint8, len(data)/8)
		} else {
			nullMap = make([]uint8, len(data)/8+1)
		}
	} else {
		nullMap = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GSeriesInt{isNullable: isNullable, name: name, data: data, nullMap: nullMap}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////////

func (s GSeriesInt) Len() int {
	return len(s.data)
}

func (s GSeriesInt) IsNullable() bool {
	return s.isNullable
}

func (s GSeriesInt) MakeNullable() {
	if !s.isNullable {
		if len(s.data)%8 == 0 {
			s.nullMap = make([]uint8, len(s.data)/8)
		} else {
			s.nullMap = make([]uint8, len(s.data)/8+1)
		}
		s.isNullable = true
	}
}

func (s GSeriesInt) Name() string {
	return s.name
}

func (s GSeriesInt) Type() GSeriesType {
	return IntType
}

func (s GSeriesInt) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GSeriesInt) NullCount() int {
	count := 0
	for _, v := range s.nullMap {
		for i := 0; i < 8; i++ {
			if v&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s GSeriesInt) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GSeriesInt) SetNull(i int) error {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
		return nil
	}
	return errors.New("GSeriesInt.SetNull: series is not nullable")
}

func (s GSeriesInt) GetNullMask() []bool {
	mask := make([]bool, len(s.data))
	idx := 0
	for _, v := range s.nullMap {
		for i := 0; i < 8 && idx < len(s.data); i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

func (s GSeriesInt) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return errors.New("GSeriesInt.SetNullMask: series is not nullable")
	}

	for k, v := range mask {
		if v {
			s.nullMap[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMap[k/8] &= ^(1 << uint(k%8))
		}
	}

	return nil
}

func (s GSeriesInt) Get(i int) interface{} {
	return s.data[i]
}

func (s GSeriesInt) Set(i int, v interface{}) {
	s.data[i] = v.(int)
}

// Append appends a value or a slice of values to the series.
func (s GSeriesInt) Append(v interface{}) error {
	if s.isNullable {
		if b, ok := v.(int); ok {
			s.data = append(s.data, b)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
		} else if bv, ok := v.([]int); ok {
			s.data = append(s.data, bv...)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, make([]uint8, len(s.data)/8-len(s.nullMap))...)
			}
		} else {
			return errors.New("GSeriesInt.Append: invalid type")
		}
	} else {
		if b, ok := v.(int); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]int); ok {
			s.data = append(s.data, bv...)
		} else {
			return errors.New("GSeriesInt.Append: invalid type")
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GSeriesInt) AppendNullable(v interface{}) error {
	if !s.isNullable {
		return errors.New("GSeriesInt.AppendNullable: series is not nullable")
	}

	if b, ok := v.(NullableInt); ok {
		s.data = append(s.data, b.Value)
		if len(s.data)/8 > len(s.nullMap) {
			s.nullMap = append(s.nullMap, 0)
		}
		if !b.Valid {
			s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableInt); ok {
		for _, b := range bv {
			s.data = append(s.data, b.Value)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
			if !b.Valid {
				s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return errors.New("GSeriesInt.AppendNullable: invalid type")
	}

	return nil
}

///////////////////////////////  	ALL DATA ACCESSORS  /////////////////////////////////

func (s GSeriesInt) Data() interface{} {
	return s.data
}

func (s GSeriesInt) NullableData() interface{} {
	data := make([]NullableInt, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GSeriesInt) StringData() []string {
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

func (s GSeriesInt) Copy() GSeries {
	data := make([]int, len(s.data))
	copy(data, s.data)
	nullMap := make([]uint8, len(s.nullMap))
	copy(nullMap, s.nullMap)

	return GSeriesInt{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

///////////////////////////////  	SERIES OPERATIONS  //////////////////////////////////

func (s GSeriesInt) Filter(filter []bool) GSeries {
	data := make([]int, 0)
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range filter {
		if v {
			data = append(data, s.data[i])
			if s.isNullable {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}
	}
	return GSeriesInt{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

type GSeriesIntPartition struct {
	partition map[int][]int
	nullGroup []int
}

func (p GSeriesIntPartition) GetGroupsCount() int {
	count := 0
	for _, v := range p.partition {
		if len(v) > 0 {
			count++
		}
	}
	if len(p.nullGroup) > 0 {
		count++
	}
	return count
}

func (p GSeriesIntPartition) GetNonNullGroups() [][]int {
	partition := make([][]int, 0)
	for _, v := range p.partition {
		if len(v) > 0 {
			partition = append(partition, v)
		}
	}
	return partition
}

func (s GSeriesIntPartition) GetNullGroup() []int {
	return s.nullGroup
}

func (s GSeriesInt) Group() GSeriesPartition {
	groups := make(map[int][]int)
	nullGroup := make([]int, 0)
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				nullGroup = append(nullGroup, i)
			} else {
				groups[v] = append(groups[v], i)
			}
		}
	} else {
		for i, v := range s.data {
			groups[v] = append(groups[v], i)
		}
	}
	return GSeriesIntPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}
