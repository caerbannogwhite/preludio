package gandalff

import (
	"errors"
)

// GSeriesString represents a series of strings.
type GSeriesString struct {
	isNullable bool
	name       string
	data       []*string
	nullMap    []uint8
	pool       *StringPool
}

func NewGSeriesString(name string, isNullable bool, data []string, pool *StringPool) GSeriesString {
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

	actualData := make([]*string, len(data))
	for i, v := range data {
		actualData[i] = pool.Get(v)
	}

	return GSeriesString{isNullable: isNullable, name: name, data: actualData, nullMap: nullMap, pool: pool}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////

func (s GSeriesString) Len() int {
	return len(s.data)
}

func (s GSeriesString) IsNullable() bool {
	return s.isNullable
}

func (s GSeriesString) MakeNullable() {
	if !s.isNullable {
		s.isNullable = true
		if len(s.data)%8 == 0 {
			s.nullMap = make([]uint8, len(s.data)/8)
		} else {
			s.nullMap = make([]uint8, len(s.data)/8+1)
		}
	}
}

func (s GSeriesString) Name() string {
	return s.name
}

func (s GSeriesString) Type() GSeriesType {
	return StringType
}

func (s GSeriesString) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GSeriesString) NullCount() int {
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

func (s GSeriesString) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GSeriesString) SetNull(i int) error {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
		return nil
	}
	return errors.New("GSeriesString.SetNull: series is not nullable")
}

func (s GSeriesString) GetNullMask() []bool {
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

func (s GSeriesString) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return errors.New("GSeriesString.SetNull: series is not nullable")
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

func (s GSeriesString) Get(i int) interface{} {
	return *s.data[i]
}

func (s GSeriesString) Set(i int, v interface{}) {
	s.data[i] = s.pool.Get(v.(string))
}

// Append appends a value or a slice of values to the series.
func (s GSeriesString) Append(v interface{}) error {
	if s.isNullable {
		if b, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Get(b))
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
		} else if bv, ok := v.([]string); ok {
			for _, b := range bv {
				s.data = append(s.data, s.pool.Get(b))
			}
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, make([]uint8, len(s.data)/8-len(s.nullMap))...)
			}
		} else {
			return errors.New("GSeriesString.Append: invalid type")
		}
	} else {
		if b, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Get(b))
		} else if bv, ok := v.([]string); ok {
			for _, b := range bv {
				s.data = append(s.data, s.pool.Get(b))
			}
		} else {
			return errors.New("GSeriesString.Append: invalid type")
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GSeriesString) AppendNullable(v interface{}) error {
	if !s.isNullable {
		return errors.New("GSeriesString.AppendNullable: series is not nullable")
	}

	if b, ok := v.(NullableString); ok {
		s.data = append(s.data, s.pool.Get(b.Value))
		if len(s.data)/8 > len(s.nullMap) {
			s.nullMap = append(s.nullMap, 0)
		}
		if !b.Valid {
			s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableString); ok {
		for _, b := range bv {
			s.data = append(s.data, s.pool.Get(b.Value))
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
			if !b.Valid {
				s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return errors.New("GSeriesString.AppendNullable: invalid type")
	}

	return nil
}

///////////////////////////////		ALL DATA ACCESSORS		/////////////////////////////

func (s GSeriesString) Data() interface{} {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return data
}

func (s GSeriesString) NullableData() interface{} {
	data := make([]NullableString, len(s.data))
	for i, v := range s.data {
		data[i] = NullableString{Valid: !s.IsNull(i), Value: *v}
	}
	return data
}

func (s GSeriesString) StringData() []string {
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

func (s GSeriesString) Copy() GSeries {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return NewGSeriesString(s.name, s.isNullable, data, s.pool)
}

/////////////////////////////// 		SERIES OPERATIONS		/////////////////////////

func (s GSeriesString) Filter(mask []bool) GSeries {
	data := make([]string, 0)
	for i, v := range s.data {
		if mask[i] {
			data = append(data, *v)
		}
	}
	return NewGSeriesString(s.name, s.isNullable, data, s.pool)
}

/////////////////////////////// 		GRAPH OPERATIONS		/////////////////////////

type GSeriesStringPartition struct {
	partition map[*string][]int
	nullGroup []int
}

func (p GSeriesStringPartition) GetGroupsCount() int {
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

func (p GSeriesStringPartition) GetNonNullGroups() [][]int {
	partition := make([][]int, 0)
	for _, v := range p.partition {
		if len(v) > 0 {
			partition = append(partition, v)
		}
	}
	return partition
}

func (s GSeriesStringPartition) GetNullGroup() []int {
	return s.nullGroup
}

func (s GSeriesString) Group() GSeriesPartition {
	groups := make(map[*string][]int)
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
	return GSeriesStringPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

func (s GSeriesString) SubGroup(partition GSeriesPartition) GSeriesPartition {
	groups := make(map[*string][]int)
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
	return GSeriesStringPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}
