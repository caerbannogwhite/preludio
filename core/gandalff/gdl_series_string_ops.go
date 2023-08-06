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
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + boolToString(o.data[0]))
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[0]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + boolToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + intToString(int64(o.data[0])))
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[0])))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(int64(o.data[i])))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + intToString(o.data[0]))
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[0]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + intToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + floatToString(o.data[0]))
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[0]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + floatToString(o.data[i]))
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		}
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.pool.Put(*s.data[0] + *o.data[0])
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[0] + *o.data[i])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[0])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.pool.Put(*s.data[i] + *o.data[i])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: s.pool, data: result}
					}
				}
			}
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
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] == *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] == *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] == *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
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
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] != *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] != *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] != *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
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
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] > *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] > *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] > *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
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
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] >= *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] >= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] >= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
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
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] < *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] < *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] < *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
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
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = *s.data[0] <= *o.data[0]
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[0] <= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[0]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = *s.data[i] <= *o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}
