package gandalff

import (
	"errors"
)

// GSeriesBool represents a series of bools.
type GSeriesBool struct {
	isNullable bool
	name       string
	data       []bool
	nullMap    []uint8
}

func NewGSeriesBool(name string, isNullable bool, makeCopy bool, data []bool) GSeriesBool {
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
		dataCopy := make([]bool, len(data))
		copy(dataCopy, data)
		data = dataCopy
	}

	return GSeriesBool{isNullable: isNullable, name: name, data: data, nullMap: nullMap}
}

///////////////////////////////		BASIC ACCESSORS		/////////////////////////////////

func (s GSeriesBool) Len() int {
	return len(s.data)
}

func (s GSeriesBool) IsNullable() bool {
	return s.isNullable
}

func (s GSeriesBool) MakeNullable() {
	if !s.isNullable {
		s.isNullable = true
		if len(s.data)%8 == 0 {
			s.nullMap = make([]uint8, len(s.data)/8)
		} else {
			s.nullMap = make([]uint8, len(s.data)/8+1)
		}
	}
}

func (s GSeriesBool) Name() string {
	return s.name
}

func (s GSeriesBool) Type() GSeriesType {
	return BoolType
}

func (s GSeriesBool) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GSeriesBool) NullCount() int {
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

func (s GSeriesBool) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GSeriesBool) SetNull(i int) error {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
		return nil
	}
	return errors.New("GSeriesBool.SetNull: series is not nullable")
}

func (s GSeriesBool) GetNullMask() []bool {
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

func (s GSeriesBool) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return errors.New("GSeriesBool.SetNullMask: series is not nullable")
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

func (s GSeriesBool) Get(i int) interface{} {
	return s.data[i]
}

func (s GSeriesBool) Set(i int, v interface{}) {
	s.data[i] = v.(bool)
}

// Append appends a value or a slice of values to the series.
func (s GSeriesBool) Append(v interface{}) error {
	if s.isNullable {
		if b, ok := v.(bool); ok {
			s.data = append(s.data, b)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, 0)
			}
		} else if bv, ok := v.([]bool); ok {
			s.data = append(s.data, bv...)
			if len(s.data)/8 > len(s.nullMap) {
				s.nullMap = append(s.nullMap, make([]uint8, len(s.data)/8-len(s.nullMap))...)
			}
		} else {
			return errors.New("GSeriesBool.Append: invalid type")
		}
	} else {
		if b, ok := v.(bool); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]bool); ok {
			s.data = append(s.data, bv...)
		} else {
			return errors.New("GSeriesBool.Append: invalid type")
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GSeriesBool) AppendNullable(v interface{}) error {
	if !s.isNullable {
		return errors.New("GSeriesBool.AppendNullable: series is not nullable")
	}

	if b, ok := v.(NullableBool); ok {
		s.data = append(s.data, b.Value)
		if len(s.data)/8 > len(s.nullMap) {
			s.nullMap = append(s.nullMap, 0)
		}
		if !b.Valid {
			s.nullMap[len(s.data)/8] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableBool); ok {
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
		return errors.New("GSeriesBool.AppendNullable: invalid type")
	}

	return nil
}

/////////////////////////////// 		ALL DATA ACCESSORS		///////////////////////////////

func (s GSeriesBool) Data() interface{} {
	return s.data
}

// NullableData returns a slice of NullableBool.
func (s GSeriesBool) NullableData() interface{} {
	data := make([]NullableBool, len(s.data))
	for i, v := range s.data {
		data[i] = NullableBool{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// StringData returns a slice of strings.
func (s GSeriesBool) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		if s.IsNull(i) {
			data[i] = NULL_STRING
		} else {
			data[i] = boolToString(v)
		}
	}
	return data
}

// Copy returns a copy of the series.
func (s GSeriesBool) Copy() GSeries {
	data := make([]bool, len(s.data))
	copy(data, s.data)
	nullMap := make([]uint8, len(s.nullMap))
	copy(nullMap, s.nullMap)

	return GSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMap:    nullMap,
	}
}

///////////////////////////////		SERIES OPERATIONS		/////////////////////////////

// Filter returns a new series with elements at the indices where mask is true.
func (s GSeriesBool) Filter(mask []bool) GSeries {
	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	if s.isNullable {
		data := make([]bool, elementCount)
		var nullMap []uint8
		if len(data)%8 == 0 {
			nullMap = make([]uint8, len(data)/8)
		} else {
			nullMap = make([]uint8, len(data)/8+1)
		}

		idx := 0
		for i, v := range mask {
			if v {
				data[idx] = s.data[i]
				if s.IsNull(i) {
					nullMap[idx/8] |= 1 << uint(idx%8)
				}
				idx++
			}
		}

		return GSeriesBool{
			isNullable: s.isNullable,
			name:       s.name,
			data:       data,
			nullMap:    nullMap,
		}
	}

	data := make([]bool, elementCount)
	idx := 0
	for i, v := range mask {
		if v {
			data[idx] = s.data[i]
			idx++
		}
	}

	return GSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

func (s GSeriesBool) FilterByIndex(indexes []int) GSeries {
	if s.isNullable {
		data := make([]bool, len(indexes))
		nullMap := make([]uint8, len(s.nullMap))
		for i, v := range indexes {
			data[i] = s.data[v]
			if s.IsNull(v) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}

		return GSeriesBool{
			isNullable: s.isNullable,
			name:       s.name,
			data:       data,
			nullMap:    nullMap,
		}
	}

	data := make([]bool, len(indexes))
	for i, v := range indexes {
		data[i] = s.data[v]
	}

	return GSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

///////////////////////////////		GROUPING OPERATIONS		/////////////////////////////

type GSeriesBoolPartition struct {
	partition map[bool][]int
	nullGroup []int
}

func (p GSeriesBoolPartition) GetGroupsCount() int {
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

func (p GSeriesBoolPartition) GetNonNullGroups() [][]int {
	partition := make([][]int, 0)
	for _, v := range p.partition {
		if len(v) > 0 {
			partition = append(partition, v)
		}
	}
	return partition
}

func (s GSeriesBoolPartition) GetNullGroup() []int {
	return s.nullGroup
}

func (s GSeriesBool) Group() GSeriesPartition {
	groups := make(map[bool][]int)
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
	return GSeriesBoolPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

func (s GSeriesBool) SubGroup(gp GSeriesPartition) GSeriesPartition {
	groups := make(map[bool][]int)
	nullGroup := make([]int, 0)
	if s.isNullable {
		for _, v := range gp.GetNonNullGroups() {
			for _, idx := range v {
				if s.IsNull(idx) {
					nullGroup = append(nullGroup, idx)
				} else {
					groups[s.data[idx]] = append(groups[s.data[idx]], idx)
				}
			}
		}
	} else {
		for _, v := range gp.GetNonNullGroups() {
			for _, idx := range v {
				groups[s.data[idx]] = append(groups[s.data[idx]], idx)
			}
		}
	}
	return GSeriesBoolPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

///////////////////////////////		LOGIC OPERATIONS		/////////////////////////////

// And performs logical AND operation between two series
func (s GSeriesBool) And(other GSeries) GSeries {
	if s.isNullable || other.IsNullable() {
		data := make([]bool, len(s.data))
		nullMap := make([]uint8, len(s.nullMap))
		for i := 0; i < len(s.data); i++ {
			if s.IsNull(i) || other.IsNull(i) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
			data[i] = s.data[i] && other.Get(i).(bool)
		}

		return GSeriesBool{
			isNullable: true,
			name:       s.name,
			data:       data,
			nullMap:    nullMap,
		}
	}

	data := make([]bool, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] = s.data[i] && other.Get(i).(bool)
	}

	return GSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

// Or performs logical OR operation between two series
func (s GSeriesBool) Or(other GSeries) GSeries {
	if s.isNullable || other.IsNullable() {
		data := make([]bool, len(s.data))
		nullMap := make([]uint8, len(s.nullMap))
		for i := 0; i < len(s.data); i++ {
			if s.IsNull(i) || other.IsNull(i) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
			data[i] = s.data[i] || other.Get(i).(bool)
		}

		return GSeriesBool{
			isNullable: true,
			name:       s.name,
			data:       data,
			nullMap:    nullMap,
		}
	}

	data := make([]bool, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] = s.data[i] || other.Get(i).(bool)
	}

	return GSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

// Not performs logical NOT operation on series
func (s GSeriesBool) Not() GSeries {
	if s.isNullable {
		data := make([]bool, len(s.data))
		nullMap := make([]uint8, len(s.nullMap))
		for i := 0; i < len(s.data); i++ {
			if s.IsNull(i) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
			data[i] = !s.data[i]
		}

		return GSeriesBool{
			isNullable: true,
			name:       s.name,
			data:       data,
			nullMap:    nullMap,
		}
	}

	data := make([]bool, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] = !s.data[i]
	}

	return GSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////////
