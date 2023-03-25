package gandalff

import (
	"fmt"
	"typesys"
)

// GDLSeriesBool represents a series of bools.
// The data is stored as a byte array, with each bit representing a bool.
type GDLSeriesBool struct {
	isNullable bool
	name       string
	size       int
	data       []uint8
	nullMask   []uint8
}

func NewGDLSeriesBool(name string, isNullable bool, data []bool) GDLSeriesBool {
	size := len(data)
	var actualData []uint8
	if size%8 == 0 {
		actualData = make([]uint8, size/8)
		for i := 0; i < size; i++ {
			if data[i] {
				actualData[i/8] |= 1 << uint(i%8)
			}
		}
	} else {
		actualData = make([]uint8, size/8+1)
		for i := 0; i < size; i++ {
			if data[i] {
				actualData[i/8] |= 1 << uint(i%8)
			}
		}
	}

	var nullMask []uint8
	if isNullable {
		nullMask = make([]uint8, len(actualData))

	} else {
		nullMask = make([]uint8, 0)
	}

	return GDLSeriesBool{isNullable: isNullable, name: name, size: size, data: actualData, nullMask: nullMask}
}

///////////////////////////////		BASIC ACCESSORS		/////////////////////////////////

func (s GDLSeriesBool) Len() int {
	return s.size
}

func (s GDLSeriesBool) IsNullable() bool {
	return s.isNullable
}

func (s GDLSeriesBool) MakeNullable() {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = make([]uint8, len(s.data))
	}
}

func (s GDLSeriesBool) Name() string {
	return s.name
}

func (s GDLSeriesBool) Type() typesys.BaseType {
	return typesys.BoolType
}

func (s GDLSeriesBool) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesBool) NullCount() int {
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

func (s GDLSeriesBool) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesBool) SetNull(i int) error {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	}
	return fmt.Errorf("GDLSeriesBool.SetNull: series is not nullable")
}

func (s GDLSeriesBool) GetNullMask() []bool {
	mask := make([]bool, s.size)
	idx := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8 && idx < s.size; i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

func (s GDLSeriesBool) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return fmt.Errorf("GDLSeriesBool.SetNullMask: series is not nullable")
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

func (s GDLSeriesBool) Get(i int) interface{} {
	return s.data[i/8]&(1<<uint(i%8)) != 0
}

func (s GDLSeriesBool) Set(i int, v interface{}) {
	if b, ok := v.(bool); ok {
		if b {
			s.data[i/8] |= 1 << uint(i%8)
		} else {
			s.data[i/8] &= ^(1 << uint(i%8))
		}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesBool) Append(v interface{}) GDLSeries {
	switch v.(type) {
	case bool:
		return s.AppendRaw(v)
	case []bool:
		return s.AppendRaw(v)
	case NullableBool:
		return s.AppendNullable(v)
	case []NullableBool:
		return s.AppendNullable(v)
	case GDLSeriesBool:
		return s.AppendSeries(v.(GDLSeriesBool))
	case GDLSeriesError:
		return v.(GDLSeriesError)
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.Append: invalid type %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesBool) AppendRaw(v interface{}) GDLSeries {
	if s.isNullable {
		if b, ok := v.(bool); ok {

			// adjust size of data and nullMask if necessary
			s.size++
			if s.size/8 > len(s.data) {
				s.data = append(s.data, 0)
				s.nullMask = append(s.nullMask, 0)
			}

			// set value
			if b {
				s.data[s.size/8] |= 1 << uint(s.size%8)
			}
			// this should not be necessary, because the data slice is initialized with 0
			// else {
			// 	s.data[s.size/8] &= ^(1 << uint(s.size%8))
			// }
		} else if bv, ok := v.([]bool); ok {

			// adjust size of data and nullMask if necessary
			s.size += len(bv)
			if s.size/8 > len(s.data) {
				s.data = append(s.data, make([]uint8, s.size/8-len(s.data))...)
				s.nullMask = append(s.nullMask, make([]uint8, s.size/8-len(s.nullMask))...)
			}

			// set values
			idx := s.size
			for _, v := range bv {
				if v {
					s.data[idx/8] |= 1 << uint(idx%8)
				}
				// this should not be necessary, because the data slice is initialized with 0
				// else {
				// 	s.data[idx/8] &= ^(1 << uint(idx%8))
				// }
				idx++
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.Append: invalid type %T", v)}
		}
	} else {
		if b, ok := v.(bool); ok {

			// adjust size of data if necessary
			s.size++
			if s.size/8 > len(s.data) {
				s.data = append(s.data, 0)
			}

			// set value
			if b {
				s.data[s.size/8] |= 1 << uint(s.size%8)
			}
			// this should not be necessary, because the data slice is initialized with 0
			// else {
			// 	s.data[s.size/8] &= ^(1 << uint(s.size%8))
			// }
		} else if bv, ok := v.([]bool); ok {

			// adjust size of data if necessary
			s.size += len(bv)
			if s.size/8 > len(s.data) {
				s.data = append(s.data, make([]uint8, s.size/8-len(s.data))...)
			}

			// set values
			idx := s.size
			for _, v := range bv {
				if v {
					s.data[idx/8] |= 1 << uint(idx%8)
				}
				// this should not be necessary, because the data slice is initialized with 0
				// else {
				// 	s.data[idx/8] &= ^(1 << uint(idx%8))
				// }
				idx++
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.Append: invalid type %T", v)}
		}
	}
	return nil
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesBool) AppendNullable(v interface{}) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesBool.AppendNullable: series is not nullable"}
	}

	if b, ok := v.(NullableBool); ok {
		// adjust size of data and nullMask if necessary
		s.size++
		if s.size/8 > len(s.data) {
			s.data = append(s.data, 0)
			s.nullMask = append(s.nullMask, 0)
		}

		// set value
		if b.Valid {
			if b.Value {
				s.data[s.size/8] |= 1 << uint(s.size%8)
			}
			// this should not be necessary, because the data slice is initialized with 0
			// else {
			// 	s.data[s.size/8] &= ^(1 << uint(s.size%8))
			// }
		} else {
			s.nullMask[s.size/8] |= 1 << uint(s.size%8)
		}
	} else if bv, ok := v.([]NullableBool); ok {
		// adjust size of data and nullMask if necessary
		s.size += len(bv)
		if s.size/8 > len(s.data) {
			s.data = append(s.data, make([]uint8, s.size/8-len(s.data))...)
			s.nullMask = append(s.nullMask, make([]uint8, s.size/8-len(s.nullMask))...)
		}

		// set values
		idx := s.size
		for _, v := range bv {
			if v.Valid {
				if v.Value {
					s.data[idx/8] |= 1 << uint(idx%8)
				}
				// this should not be necessary, because the data slice is initialized with 0
				// else {
				// 	s.data[idx/8] &= ^(1 << uint(idx%8))
				// }
			} else {
				s.nullMask[idx/8] |= 1 << uint(idx%8)
			}
			idx++
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.AppendNullable: invalid type %T", v)}
	}

	return nil
}

// AppendSeries appends a series to the series.
func (s GDLSeriesBool) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesBool
	if o, ok = other.(GDLSeriesBool); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesBool.AppendSeries: invalid type %T", other)}
	}

	size := s.size + o.size

	if s.isNullable {
		// adjust size of data and nullMask if necessary
		if size > len(s.data)*8 {
			s.data = append(s.data, make([]uint8, size/8-len(s.data)+1)...)
			s.nullMask = append(s.nullMask, make([]uint8, size/8-len(s.nullMask)+1)...)
		}

		// both series are nullable
		if o.isNullable {
			sIdx := s.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8; j++ {
					if o.nullMask[oIdx/8]&(1<<uint(j)) != 0 {
						s.nullMask[sIdx/8] |= 1 << uint(sIdx%8)
					} else {
						if v&(1<<uint(j)) != 0 {
							s.data[sIdx/8] |= 1 << uint(sIdx%8)
						}
					}
					sIdx++
					oIdx++
				}
			}

			return GDLSeriesBool{
				isNullable: true,
				name:       s.name,
				size:       size,
				data:       s.data,
				nullMask:   s.nullMask,
			}
		} else

		// s is nullable, o is not nullable
		{
			sIdx := s.size - o.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8; j++ {
					if v&(1<<uint(j)) != 0 {
						s.data[sIdx/8] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}

			return GDLSeriesBool{
				isNullable: true,
				name:       s.name,
				size:       size,
				data:       s.data,
				nullMask:   s.nullMask,
			}
		}
	} else {
		// s is not nullable, o is nullable
		if o.isNullable {
			if s.size > len(s.data)*8 {
				s.data = append(s.data, make([]uint8, s.size/8-len(s.data)+1)...)
				s.nullMask = make([]uint8, len(s.data))
			}

			// set values
			sIdx := s.size - o.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8; j++ {
					if o.nullMask[oIdx/8]&(1<<uint(j)) != 0 {
						s.nullMask[sIdx/8] |= 1 << uint(sIdx%8)
					} else {
						if v&(1<<uint(j)) != 0 {
							s.data[sIdx/8] |= 1 << uint(sIdx%8)
						}
					}
					sIdx++
					oIdx++
				}
			}

			return GDLSeriesBool{
				isNullable: true,
				name:       s.name,
				size:       size,
				data:       s.data,
				nullMask:   s.nullMask,
			}
		} else

		// both series are not nullable
		{
			if s.size > len(s.data)*8 {
				s.data = append(s.data, make([]uint8, s.size/8-len(s.data)+1)...)
			}

			// set values
			sIdx := s.size - o.size
			oIdx := 0
			for _, v := range o.data {
				for j := 0; j < 8; j++ {
					if v&(1<<uint(j)) != 0 {
						s.data[sIdx/8] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}

			return GDLSeriesBool{
				isNullable: false,
				name:       s.name,
				size:       size,
				data:       s.data,
				nullMask:   s.nullMask,
			}
		}
	}
}

/////////////////////////////// 		ALL DATA ACCESSORS		///////////////////////////////

func (s GDLSeriesBool) Data() interface{} {
	data := make([]bool, s.size)
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < s.size; j++ {
			data[i*8+j] = v&(1<<uint(j)) != 0
		}
	}
	return data
}

// NullableData returns a slice of NullableBool.
func (s GDLSeriesBool) NullableData() interface{} {
	data := make([]NullableBool, len(s.data))
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < len(s.data); j++ {
			if s.nullMask[i]&(1<<uint(j)) != 0 {
				data[i*8+j] = NullableBool{false, false}
			} else {
				data[i*8+j] = NullableBool{v&(1<<uint(j)) != 0, true}
			}
		}
	}
	return data
}

// StringData returns a slice of strings.
func (s GDLSeriesBool) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		for j := 0; j < 8 && i*8+j < len(s.data); j++ {
			if s.nullMask[i]&(1<<uint(j)) != 0 {
				data[i*8+j] = NULL_STRING
			} else {
				if v&(1<<uint(j)) != 0 {
					data[i*8+j] = BOOL_TRUE_STRING
				} else {
					data[i*8+j] = BOOL_FALSE_STRING
				}
			}
		}
	}
	return data
}

// Copy returns a copy of the series.
func (s GDLSeriesBool) Copy() GDLSeries {
	data := make([]uint8, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nullMask,
	}
}

///////////////////////////////		SERIES OPERATIONS		/////////////////////////////

// Filter returns a new series with elements at the indices where mask is true.
func (s GDLSeriesBool) Filter(mask []bool) GDLSeries {
	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	if s.isNullable {
		var data []uint8
		var nullMask []uint8
		if elementCount%8 == 0 {
			data = make([]uint8, elementCount/8)
			nullMask = make([]uint8, elementCount/8)
		} else {
			data = make([]uint8, elementCount/8+1)
			nullMask = make([]uint8, elementCount/8+1)
		}

		idx := 0
		for i, v := range mask {
			if v {
				data[idx/8] |= s.data[i/8] & (1 << uint(i%8))
				nullMask[idx/8] |= s.nullMask[i/8] & (1 << uint(i%8))
				idx++
			}
		}

		return GDLSeriesBool{
			isNullable: s.isNullable,
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}
	}

	data := make([]uint8, elementCount/8+1)
	idx := 0
	for i, v := range mask {
		if v {
			data[idx/8] |= s.data[i/8] & (1 << uint(i%8))
			idx++
		}
	}

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nil,
	}
}

func (s GDLSeriesBool) FilterByIndex(indexes []int) GDLSeries {
	if s.isNullable {
		var data []uint8
		var nullMask []uint8
		if len(indexes)%8 == 0 {
			data = make([]uint8, len(indexes)/8)
			nullMask = make([]uint8, len(indexes)/8)
		} else {
			data = make([]uint8, len(indexes)/8+1)
			nullMask = make([]uint8, len(indexes)/8+1)
		}

		for i, v := range indexes {
			data[i/8] |= s.data[v/8] & (1 << uint(v%8))
			nullMask[i/8] |= s.nullMask[v/8] & (1 << uint(v%8))
		}

		return GDLSeriesBool{
			isNullable: s.isNullable,
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}
	}

	data := make([]uint8, len(indexes)/8+1)
	for i, v := range indexes {
		data[i/8] |= s.data[v/8] & (1 << uint(v%8))
	}

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   nil,
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

func (s GDLSeriesBool) Group() GDLSeriesPartition {
	groups := make(map[bool][]int)
	nullGroup := make([]int, 0)
	// if s.isNullable {
	// 	for i, v := range s.data {
	// 		if s.IsNull(i) {
	// 			nullGroup = append(nullGroup, i)
	// 		} else {
	// 			groups[v] = append(groups[v], i)
	// 		}
	// 	}
	// } else {
	// 	for i, v := range s.data {
	// 		groups[v] = append(groups[v], i)
	// 	}
	// }
	return GDLSeriesBoolPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

func (s GDLSeriesBool) SubGroup(gp GDLSeriesPartition) GDLSeriesPartition {
	groups := make(map[bool][]int)
	nullGroup := make([]int, 0)
	// if s.isNullable {
	// 	for _, v := range gp.GetNonNullGroups() {
	// 		for _, idx := range v {
	// 			if s.IsNull(idx) {
	// 				nullGroup = append(nullGroup, idx)
	// 			} else {
	// 				groups[s.data[idx]] = append(groups[s.data[idx]], idx)
	// 			}
	// 		}
	// 	}
	// } else {
	// 	for _, v := range gp.GetNonNullGroups() {
	// 		for _, idx := range v {
	// 			groups[s.data[idx]] = append(groups[s.data[idx]], idx)
	// 		}
	// 	}
	// }
	return GDLSeriesBoolPartition{
		partition: groups,
		nullGroup: nullGroup,
	}
}

///////////////////////////////		LOGIC OPERATIONS		/////////////////////////////

// And performs logical AND operation between two series
// If one of the series is nullable, the result series will be nullable
// If the other series is not a boolean series, the result will be nil
func (s GDLSeriesBool) And(other GDLSeries) GDLSeries {
	if other.Type() != typesys.BoolType {
		return nil
	}
	o := other.(GDLSeriesBool)

	if s.isNullable || other.IsNullable() {
		data := make([]uint8, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			nullMask[i] = s.nullMask[i] | o.nullMask[i]
			data[i] = s.data[i] & o.data[i]
		}

		return GDLSeriesBool{
			isNullable: true,
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}
	}

	data := make([]uint8, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] = s.data[i] & o.data[i]
	}

	return GDLSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMask:   make([]uint8, 0),
	}
}

// Or performs logical OR operation between two series
// If one of the series is nullable, the result series will be nullable
// If the other series is not a boolean series, the result will be nil
func (s GDLSeriesBool) Or(other GDLSeries) GDLSeries {
	if other.Type() != typesys.BoolType {
		return nil
	}
	o := other.(GDLSeriesBool)

	if s.isNullable || other.IsNullable() {
		data := make([]uint8, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			nullMask[i] = s.nullMask[i] | o.nullMask[i]
			data[i] = s.data[i] | o.data[i]
		}

		return GDLSeriesBool{
			isNullable: true,
			name:       s.name,
			data:       data,
			nullMask:   nullMask,
		}
	}

	data := make([]uint8, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] = s.data[i] | o.data[i]
	}

	return GDLSeriesBool{
		isNullable: false,
		name:       s.name,
		data:       data,
		nullMask:   make([]uint8, 0),
	}
}

// Not performs logical NOT operation on series
func (s GDLSeriesBool) Not() GDLSeries {
	data := make([]uint8, len(s.data))
	for i := 0; i < len(s.data); i++ {
		data[i] ^= s.data[i]
	}

	return GDLSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMask:   s.nullMask,
	}
}

///////////////////////////////		ARITHMETIC OPERATIONS		/////////////////////////////
