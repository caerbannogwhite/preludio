package gandalff

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
		nullMap = make([]uint8, len(data)/8+1)
	} else {
		nullMap = make([]uint8, 0)
	}

	actualData := make([]*string, len(data))
	for i, v := range data {
		actualData[i] = pool.Get(v)
	}

	return GSeriesString{isNullable: isNullable, name: name, data: actualData, nullMap: nullMap, pool: pool}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////////

func (s GSeriesString) Len() int {
	return len(s.data)
}

func (s GSeriesString) IsNullable() bool {
	return s.isNullable
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

func (s GSeriesString) SetNull(i int) {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
	}
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

func (s GSeriesString) SetNullMask(mask []bool) {
	for k, v := range mask {
		if v {
			s.nullMap[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMap[k/8] &= ^(1 << uint(k%8))
		}
	}
}

func (s GSeriesString) Get(i int) interface{} {
	return *s.data[i]
}

func (s GSeriesString) Set(i int, v interface{}) {
	s.data[i] = s.pool.Get(v.(string))
}

///////////////////////////////		ALL DATA ACCESSORS		/////////////////////////////////

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

/////////////////////////////// 		SERIES OPERATIONS		/////////////////////////////////

func (s GSeriesString) Filter(mask []bool) GSeries {
	data := make([]string, 0)
	for i, v := range s.data {
		if mask[i] {
			data = append(data, *v)
		}
	}
	return NewGSeriesString(s.name, s.isNullable, data, s.pool)
}
