package gandalff

import (
	"fmt"
	"time"
)

func (s SeriesTime) Mul(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesTime) Div(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesTime) Mod(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesTime) Exp(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesTime) Add(other Series) Series {
	switch o := other.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = o.pool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = o.pool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = o.pool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = o.pool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
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
							result[i] = o.pool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
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
							result[i] = o.pool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
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
							result[i] = o.pool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
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
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[0].Year(), int(o.data[0].Month()), o.data[0].Day())
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].AddDate(o.data[i].Year(), int(o.data[i].Month()), o.data[i].Day())
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
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
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Add(o.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Add(o.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Add(o.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Add(o.data[0])
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(o.data[i])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[0])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(o.data[i])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesTime) Sub(other Series) Series {
	switch o := other.(type) {
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Sub(o.data[0])
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Sub(o.data[0])
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Sub(o.data[0])
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Sub(o.data[0])
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[0])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[0])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[0])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[0])
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Sub(o.data[i])
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Add(-o.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Add(-o.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Add(-o.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Add(-o.data[0])
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(-o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(-o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(-o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Add(-o.data[i])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[0])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Add(-o.data[i])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesTime) Eq(other Series) Series {
	switch o := other.(type) {
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Compare(o.data[0]) == 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) == 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) == 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Compare(o.data[0]) == 0
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
							result[i] = s.data[0].Compare(o.data[i]) == 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) == 0
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
							result[i] = s.data[0].Compare(o.data[i]) == 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) == 0
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
							result[i] = s.data[i].Compare(o.data[0]) == 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == 0
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
							result[i] = s.data[i].Compare(o.data[i]) == 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) == 0
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
							result[i] = s.data[i].Compare(o.data[i]) == 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) == 0
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

func (s SeriesTime) Ne(other Series) Series {
	switch o := other.(type) {
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Compare(o.data[0]) != 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) != 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) != 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Compare(o.data[0]) != 0
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
							result[i] = s.data[0].Compare(o.data[i]) != 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) != 0
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
							result[i] = s.data[0].Compare(o.data[i]) != 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) != 0
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
							result[i] = s.data[i].Compare(o.data[0]) != 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) != 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) != 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) != 0
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
							result[i] = s.data[i].Compare(o.data[i]) != 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) != 0
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
							result[i] = s.data[i].Compare(o.data[i]) != 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) != 0
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

func (s SeriesTime) Gt(other Series) Series {
	switch o := other.(type) {
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Compare(o.data[0]) == 1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) == 1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) == 1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Compare(o.data[0]) == 1
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
							result[i] = s.data[0].Compare(o.data[i]) == 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) == 1
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
							result[i] = s.data[0].Compare(o.data[i]) == 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) == 1
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
							result[i] = s.data[i].Compare(o.data[0]) == 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == 1
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
							result[i] = s.data[i].Compare(o.data[i]) == 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) == 1
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
							result[i] = s.data[i].Compare(o.data[i]) == 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) == 1
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

func (s SeriesTime) Ge(other Series) Series {
	switch o := other.(type) {
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Compare(o.data[0]) >= 1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) >= 1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) >= 1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Compare(o.data[0]) >= 1
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
							result[i] = s.data[0].Compare(o.data[i]) >= 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) >= 1
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
							result[i] = s.data[0].Compare(o.data[i]) >= 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) >= 1
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
							result[i] = s.data[i].Compare(o.data[0]) >= 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) >= 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) >= 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) >= 1
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
							result[i] = s.data[i].Compare(o.data[i]) >= 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) >= 1
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
							result[i] = s.data[i].Compare(o.data[i]) >= 1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) >= 1
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

func (s SeriesTime) Lt(other Series) Series {
	switch o := other.(type) {
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Compare(o.data[0]) == -1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) == -1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) == -1
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Compare(o.data[0]) == -1
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
							result[i] = s.data[0].Compare(o.data[i]) == -1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) == -1
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
							result[i] = s.data[0].Compare(o.data[i]) == -1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) == -1
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
							result[i] = s.data[i].Compare(o.data[0]) == -1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == -1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == -1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) == -1
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
							result[i] = s.data[i].Compare(o.data[i]) == -1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) == -1
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
							result[i] = s.data[i].Compare(o.data[i]) == -1
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) == -1
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

func (s SeriesTime) Le(other Series) Series {
	switch o := other.(type) {
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0].Compare(o.data[0]) <= 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) <= 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0].Compare(o.data[0]) <= 0
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0].Compare(o.data[0]) <= 0
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
							result[i] = s.data[0].Compare(o.data[i]) <= 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) <= 0
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
							result[i] = s.data[0].Compare(o.data[i]) <= 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0].Compare(o.data[i]) <= 0
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
							result[i] = s.data[i].Compare(o.data[0]) <= 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) <= 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) <= 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[0]) <= 0
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
							result[i] = s.data[i].Compare(o.data[i]) <= 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) <= 0
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
							result[i] = s.data[i].Compare(o.data[i]) <= 0
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i].Compare(o.data[i]) <= 0
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
