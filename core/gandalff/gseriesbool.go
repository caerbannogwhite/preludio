package gandalff

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
		nullMap = make([]uint8, len(data)/8+1)
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

func (s GSeriesBool) SetNull(i int) {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
	}
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

func (s GSeriesBool) SetNullMask(mask []bool) {
	for k, v := range mask {
		if v {
			s.nullMap[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMap[k/8] &= ^(1 << uint(k%8))
		}
	}
}

func (s GSeriesBool) Get(i int) interface{} {
	return s.data[i]
}

func (s GSeriesBool) Set(i int, v interface{}) {
	s.data[i] = v.(bool)
}

/////////////////////////////// 		ALL DATA ACCESSORS		///////////////////////////////

func (s GSeriesBool) Data() interface{} {
	return s.data
}

func (s GSeriesBool) NullableData() interface{} {
	data := make([]NullableBool, len(s.data))
	for i, v := range s.data {
		data[i] = NullableBool{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

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

func (s GSeriesBool) Filter(mask []bool) GSeries {
	if s.isNullable {
		data := make([]bool, 0)
		nullMap := make([]uint8, len(s.nullMap))
		idx := 0
		for i, v := range s.data {
			if mask[i] {
				data = append(data, v)
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

	data := make([]bool, 0)
	nullMap := make([]uint8, 0)
	for i, v := range s.data {
		if mask[i] {
			data = append(data, v)
		}
	}

	return GSeriesBool{
		isNullable: s.isNullable,
		name:       s.name,
		data:       data,
		nullMap:    nullMap,
	}
}

func (s GSeriesBool) FilterInPlace(mask []bool) {
	if s.isNullable {
		data := make([]bool, 0)
		nullMap := make([]uint8, len(s.nullMap))
		idx := 0
		for i, v := range s.data {
			if mask[i] {
				data = append(data, v)
				if s.IsNull(i) {
					nullMap[idx/8] |= 1 << uint(idx%8)
				}
				idx++
			}
		}

		s.data = data
		s.nullMap = nullMap
	} else {
		data := make([]bool, 0)
		for i, v := range s.data {
			if mask[i] {
				data = append(data, v)
			}
		}

		s.data = data
		s.nullMap = make([]uint8, 0)
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

func (s GSeriesBool) FilterByIndexInPlace(indexes []int) {
	if s.isNullable {
		data := make([]bool, len(indexes))
		nullMap := make([]uint8, len(s.nullMap))
		for i, v := range indexes {
			data[i] = s.data[v]
			if s.IsNull(v) {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}

		s.data = data
		s.nullMap = nullMap
	} else {
		data := make([]bool, len(indexes))
		for i, v := range indexes {
			data[i] = s.data[v]
		}

		s.data = data
		s.nullMap = make([]uint8, 0)
	}
}

///////////////////////////////		LOGIC OPERATIONS		/////////////////////////////

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
