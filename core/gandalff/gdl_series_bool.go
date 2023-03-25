package gandalff

import (
	"errors"
	"typesys"
)

// GDLSeriesBool represents a series of bools.
type GDLSeriesBool struct {
	isNullable bool
	name       string
	data       []bool
	nullMap    []uint8
}

func NewGDLSeriesBool(name string, isNullable bool, makeCopy bool, data []bool) GDLSeriesBool {
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

	return GDLSeriesBool{isNullable: isNullable, name: name, data: data, nullMap: nullMap}
}

///////////////////////////////		BASIC ACCESSORS		/////////////////////////////////

func (s GDLSeriesBool) Len() int {
	return len(s.data)
}

func (s GDLSeriesBool) IsNullable() bool {
	return s.isNullable
}

func (s GDLSeriesBool) MakeNullable() {
	if !s.isNullable {
		s.isNullable = true
		if len(s.data)%8 == 0 {
			s.nullMap = make([]uint8, len(s.data)/8)
		} else {
			s.nullMap = make([]uint8, len(s.data)/8+1)
		}
	}
}

func (s GDLSeriesBool) Name() string {
	return s.name
}

func (s GDLSeriesBool) Type() typesys.BaseType {
	return typesys.BoolType
}

func (s GDLSeriesBool) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesBool) NullCount() int {
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

func (s GDLSeriesBool) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesBool) SetNull(i int) error {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
		return nil
	}
	return errors.New("GDLSeriesBool.SetNull: series is not nullable")
}

func (s GDLSeriesBool) GetNullMask() []bool {
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

func (s GDLSeriesBool) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return errors.New("GDLSeriesBool.SetNullMask: series is not nullable")
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

func (s GDLSeriesBool) Get(i int) interface{} {
	return s.data[i]
}

func (s GDLSeriesBool) Set(i int, v interface{}) {
	s.data[i] = v.(bool)
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesBool) Append(v interface{}) error {
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
			return errors.New("GDLSeriesBool.Append: invalid type")
		}
	} else {
		if b, ok := v.(bool); ok {
			s.data = append(s.data, b)
		} else if bv, ok := v.([]bool); ok {
			s.data = append(s.data, bv...)
		} else {
			return errors.New("GDLSeriesBool.Append: invalid type")
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesBool) AppendNullable(v interface{}) error {
	if !s.isNullable {
		return errors.New("GDLSeriesBool.AppendNullable: series is not nullable")
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
		return errors.New("GDLSeriesBool.AppendNullable: invalid type")
	}

	return nil
}

/////////////////////////////// 		ALL DATA ACCESSORS		///////////////////////////////

func (s GDLSeriesBool) Data() interface{} {
	return s.data
}

// NullableData returns a slice of NullableBool.
func (s GDLSeriesBool) NullableData() interface{} {
	data := make([]NullableBool, len(s.data))
	for i, v := range s.data {
		data[i] = NullableBool{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

// StringData returns a slice of strings.
func (s GDLSeriesBool) StringData() []string {
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
func (s GDLSeriesBool) Copy() GSeries {
	data := make([]bool, len(s.data))
	copy(data, s.data)
	nullMap := make([]uint8, len(s.nullMap))
	copy(nullMap, s.nullMap)

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMap:    nullMap,
	}
}

///////////////////////////////		SERIES OPERATIONS		/////////////////////////////

// Filter returns a new series with elements at the indices where mask is true.
func (s GDLSeriesBool) Filter(mask []bool) GSeries {
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

		return GDLSeriesBool{
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

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

func (s GDLSeriesBool) FilterByIndex(indexes []int) GSeries {
	if s.isNullable {
		data := make([]bool, len(indexes))
		nullMap := make([]uint8, len(s.nullMap))
		for i, v := range indexes {
			data[i] = s.data[v]
			if s.IsNull(v) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}

		return GDLSeriesBool{
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

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

///////////////////////////////		GROUPING OPERATIONS		/////////////////////////////

type GDLSeriesBoolPartition struct {
	partition map[bool][]int
	nullGroup []int
}

func (p GDLSeriesBoolPartition) GetGroupsCount() int {
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

func (p GDLSeriesBoolPartition) GetNonNullGroups() [][]int {
	partition := make([][]int, 0)
	for _, v := range p.partition {
		if len(v) > 0 {
			partition = append(partition, v)
		}
	}
	return partition
}

func (s GDLSeriesBoolPartition) GetNullGroup() []int {
	return s.nullGroup
}

func (s GDLSeriesBool) Group() GSeriesPartition {
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
	return GDLSeriesBoolPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

func (s GDLSeriesBool) SubGroup(gp GSeriesPartition) GSeriesPartition {
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
	return GDLSeriesBoolPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

///////////////////////////////		LOGIC OPERATIONS		/////////////////////////////

// And performs logical AND operation between two series
func (s GDLSeriesBool) And(other GSeries) GSeries {
	if s.isNullable || other.IsNullable() {
		data := make([]bool, len(s.data))
		nullMap := make([]uint8, len(s.nullMap))
		for i := 0; i < len(s.data); i++ {
			if s.IsNull(i) || other.IsNull(i) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
			data[i] = s.data[i] && other.Get(i).(bool)
		}

		return GDLSeriesBool{
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

	return GDLSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

// Or performs logical OR operation between two series
func (s GDLSeriesBool) Or(other GSeries) GSeries {
	if s.isNullable || other.IsNullable() {
		data := make([]bool, len(s.data))
		nullMap := make([]uint8, len(s.nullMap))
		for i := 0; i < len(s.data); i++ {
			if s.IsNull(i) || other.IsNull(i) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
			data[i] = s.data[i] || other.Get(i).(bool)
		}

		return GDLSeriesBool{
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

	return GDLSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

// Not performs logical NOT operation on series
func (s GDLSeriesBool) Not() GSeries {
	if s.isNullable {
		data := make([]bool, len(s.data))
		nullMap := make([]uint8, len(s.nullMap))
		for i := 0; i < len(s.data); i++ {
			if s.IsNull(i) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
			data[i] = !s.data[i]
		}

		return GDLSeriesBool{
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

	return GDLSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMap:    make([]uint8, 0),
	}
}

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////////
