
package gandalff

import "typesys"

////////////////////////			BASIC ACCESSORS

// Returns the number of elements in the series.
func (s SeriesInt64) Len() int {
	return len(s.data)
}

// Returns the name of the series.
func (s SeriesInt64) Name() string {
	return s.name
}

// Sets the name of the series.
func (s SeriesInt64) SetName(name string) Series {
	s.name = name
	return s
}

// Returns the type of the series.
func (s SeriesInt64) Type() typesys.BaseType {
	return typesys.Int64Type
}

// Returns the type and cardinality of the series.
func (s SeriesInt64) TypeCard() typesys.BaseTypeCard {
	return typesys.BaseTypeCard{Base: typesys.Int64Type, Card: s.Len()}
}

// Returns if the series is grouped.
func (s SeriesInt64) IsGrouped() bool {
	return s.isGrouped
}

// Returns if the series admits null values.
func (s SeriesInt64) IsNullable() bool {
	return s.isNullable
}

// Returns if the series is sorted.
func (s SeriesInt64) IsSorted() SeriesSortOrder {
	return s.sorted
}

// Returns if the series is error.
func (s SeriesInt64) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s SeriesInt64) GetError() string {
	return ""
}

// Returns if the series has null values.
func (s SeriesInt64) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Returns the number of null values in the series.
func (s SeriesInt64) NullCount() int {
	count := 0
	for _, x := range s.nullMask {
		for ; x != 0; x >>= 1 {
			count += int(x & 1)
		}
	}
	return count
}

// Returns if the element at index i is null.
func (s SeriesInt64) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

// Sets the element at index i to null.
func (s SeriesInt64) SetNull(i int) Series {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	} else {
		nullMask := __binVecInit(len(s.data))
		nullMask[i/8] |= 1 << uint(i%8)

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Returns the null mask of the series.
func (s SeriesInt64) GetNullMask() []bool {
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

// Sets the null mask of the series.
func (s SeriesInt64) SetNullMask(mask []bool) Series {
	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k/8] |= 1 << uint(k%8)
			} else {
				s.nullMask[k/8] &= ^(1 << uint(k%8))
			}
		}
		return s
	} else {
		nullMask := __binVecInit(len(s.data))
		for k, v := range mask {
			if v {
				nullMask[k/8] |= 1 << uint(k%8)
			} else {
				nullMask[k/8] &= ^(1 << uint(k%8))
			}
		}

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Makes the series nullable.
func (s SeriesInt64) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __binVecInit(len(s.data))
	}
	return s
}

// Get the element at index i.
func (s SeriesInt64) Get(i int) any {
	return s.data[i]
}
