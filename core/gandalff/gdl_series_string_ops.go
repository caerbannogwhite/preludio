package gandalff

import (
	"fmt"
)

func (s SeriesString) Mul(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Div(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Mod(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Pow(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Add(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
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
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
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
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
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
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.pool.Put(*s.data[0] + o.data[0].String())
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + o.data[i].String())
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[0].String())
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + o.data[i].String())
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Sub(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Eq(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Ne(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Gt(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Ge(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Lt(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesString) Le(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}
