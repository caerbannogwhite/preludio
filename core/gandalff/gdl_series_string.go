package gandalff

import (
	"fmt"
	"typesys"
)

// GDLSeriesString represents a series of strings.
type GDLSeriesString struct {
	isNullable bool
	name       string
	data       []*string
	nullMask   []uint8
	pool       *StringPool
}

func NewGDLSeriesString(name string, isNullable bool, data []string, pool *StringPool) GDLSeriesString {
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

func (s GDLSeriesString) Len() int {
	return len(s.data)
}

func (s GDLSeriesString) IsNullable() bool {
	return s.isNullable
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

func (s GDLSeriesString) Name() string {
	return s.name
}

func (s GDLSeriesString) Type() typesys.BaseType {
	return typesys.StringType
}

func (s GDLSeriesString) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesString) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8; i++ {
			if v&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s GDLSeriesString) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesString) SetNull(i int) error {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	}
	return fmt.Errorf("GDLSeriesString.SetNull: series is not nullable")
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

func (s GDLSeriesString) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return fmt.Errorf("GDLSeriesString.SetNull: series is not nullable")
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

func (s GDLSeriesString) Get(i int) interface{} {
	return *s.data[i]
}

func (s GDLSeriesString) Set(i int, v interface{}) {
	s.data[i] = s.pool.Get(v.(string))
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesString) Append(v interface{}) GDLSeries {
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
func (s GDLSeriesString) AppendRaw(v interface{}) GDLSeries {
	if s.isNullable {
		if b, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Get(b))
			if len(s.data)/8 > len(s.nullMask) {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if bv, ok := v.([]string); ok {
			for _, b := range bv {
				s.data = append(s.data, s.pool.Get(b))
			}
			if len(s.data)/8 > len(s.nullMask) {
				s.nullMask = append(s.nullMask, make([]uint8, len(s.data)/8-len(s.nullMask))...)
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesString.Append: invalid type, %T", v)}
		}
	} else {
		if b, ok := v.(string); ok {
			s.data = append(s.data, s.pool.Get(b))
		} else if bv, ok := v.([]string); ok {
			for _, b := range bv {
				s.data = append(s.data, s.pool.Get(b))
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesString.Append: invalid type, %T", v)}
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesString) AppendNullable(v interface{}) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.AppendNullable: series is not nullable")}
	}

	if b, ok := v.(NullableString); ok {
		s.data = append(s.data, s.pool.Get(b.Value))
		if len(s.data)/8 > len(s.nullMask) {
			s.nullMask = append(s.nullMask, 0)
		}
		if !b.Valid {
			s.nullMask[len(s.data)/8] |= 1 << uint(len(s.data)%8)
		}
	} else if bv, ok := v.([]NullableString); ok {
		for _, b := range bv {
			s.data = append(s.data, s.pool.Get(b.Value))
			if len(s.data)/8 > len(s.nullMask) {
				s.nullMask = append(s.nullMask, 0)
			}
			if !b.Valid {
				s.nullMask[len(s.data)/8] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesString.AppendNullable: invalid type")}
	}

	return nil
}

// AppendSeries appends a series to the series.
func (s GDLSeriesString) AppendSeries(v GDLSeries) GDLSeries {
	return nil
}

///////////////////////////////		ALL DATA ACCESSORS		/////////////////////////////

func (s GDLSeriesString) Data() interface{} {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		data[i] = *v
	}
	return data
}

func (s GDLSeriesString) NullableData() interface{} {
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

func (s GDLSeriesString) Filter(mask []bool) GDLSeries {
	data := make([]string, 0)
	for i, v := range s.data {
		if mask[i] {
			data = append(data, *v)
		}
	}
	return NewGDLSeriesString(s.name, s.isNullable, data, s.pool)
}

/////////////////////////////// 		GRAPH OPERATIONS		/////////////////////////

type GDLSeriesStringPartition struct {
	partition map[*string][]int
	nullGroup []int
}

func (p GDLSeriesStringPartition) GetGroupsCount() int {
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

func (p GDLSeriesStringPartition) GetNonNullGroups() [][]int {
	partition := make([][]int, 0)
	for _, v := range p.partition {
		if len(v) > 0 {
			partition = append(partition, v)
		}
	}
	return partition
}

func (s GDLSeriesStringPartition) GetNullGroup() []int {
	return s.nullGroup
}

func (s GDLSeriesString) Group() GDLSeriesPartition {
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
	return GDLSeriesStringPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

func (s GDLSeriesString) SubGroup(partition GDLSeriesPartition) GDLSeriesPartition {
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
	return GDLSeriesStringPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}
