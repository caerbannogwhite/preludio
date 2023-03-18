package gandalff

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
		nullMap = make([]uint8, len(data)/8+1)
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

func (s GSeriesInt) SetNull(i int) {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
	}
}

func (s GSeriesInt) Get(i int) interface{} {
	return s.data[i]
}

func (s GSeriesInt) Set(i int, v interface{}) {
	s.data[i] = v.(int)
}

///////////////////////////////  	ALL DATA ACCESSORS  /////////////////////////////////

func (s GSeriesInt) Data() interface{} {
	return s.data
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

func (s GSeriesInt) SetNullMask(mask []bool) {
	for k, v := range mask {
		if v {
			s.nullMap[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMap[k/8] &= ^(1 << uint(k%8))
		}
	}
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

func (s GSeriesInt) FilterInPlace(filter []bool) {
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
	s.data = data
	s.nullMap = nullMap
}

func (s GSeriesInt) FilterByIndex(index []int) GSeries {
	data := make([]int, len(index))
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range index {
		data[i] = s.data[v]
		if s.isNullable {
			nullMap[i/8] |= 1 << uint(i%8)
		}
	}
	return GSeriesInt{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

func (s GSeriesInt) FilterByIndexInPlace(index []int) {
	data := make([]int, len(index))
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range index {
		data[i] = s.data[v]
		if s.isNullable {
			nullMap[i/8] |= 1 << uint(i%8)
		}
	}
	s.data = data
	s.nullMap = nullMap
}
