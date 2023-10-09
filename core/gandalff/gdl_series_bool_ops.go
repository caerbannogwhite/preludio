package gandalff

import (
	"fmt"
	"math"
)

// Not performs logical NOT operation on series
func (s SeriesBool) Not() Series {
	for i := 0; i < len(s.data); i++ {
		s.data[i] = !s.data[i]
	}

	return s
}

func (s SeriesBool) And(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] && o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] && o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] && o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] && o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] && o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] && o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] && o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] && o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot and %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot and %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot and %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Or(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] || o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] || o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] || o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] || o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] || o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] || o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] || o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] || o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot or %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, s.nullMask[0] == 1)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				result[0] = s.data[0]
				return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			} else {
				resultSize := o.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, s.nullMask[0] == 1)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				for i := 0; i < resultSize; i++ {
					result[i] = s.data[0]
				}
				return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, false)
					copy(resultNullMask, s.nullMask)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				for i := 0; i < resultSize; i++ {
					result[i] = s.data[i]
				}
				return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, false)
					copy(resultNullMask, s.nullMask)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				for i := 0; i < resultSize; i++ {
					result[i] = s.data[i]
				}
				return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			}
			return SeriesError{fmt.Sprintf("Cannot or %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot or %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Mul(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Div(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Mod(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Exp(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Add(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(boolToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Sub(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesInt64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Eq(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Ne(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Gt(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Ge(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Lt(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesBool) Le(other Series) Series {
	if s.ctx != other.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, other.GetContext())}
	}
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}
