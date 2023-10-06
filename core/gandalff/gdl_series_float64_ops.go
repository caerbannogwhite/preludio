package gandalff

import (
	"fmt"
	"math"
)

func (s SeriesFloat64) Neg() Series {
	for i, v := range s.data {
		s.data[i] = -v
	}

	return s
}

func (s SeriesFloat64) Mul(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] * float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] * o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] * o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] * o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] * o.data[0]
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesFloat64) Div(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] / b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] / b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] / b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] / b2
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] / b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] / b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] / float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] / o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] / o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] / o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] / o.data[0]
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] / o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] / o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesFloat64) Mod(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(s.data[0], b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(s.data[0], b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(s.data[0], b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(s.data[0], b2)
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesFloat64) Pow(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Pow(s.data[0], b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Pow(s.data[0], b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Pow(s.data[0], b2)
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Pow(s.data[0], b2)
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[0], b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Pow(s.data[i], b2)
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = math.Pow(s.data[0], float64(o.data[0]))
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[0], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(s.data[i], float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesFloat64) Add(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] + float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] + o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] + o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] + o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] + o.data[0]
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						result[0] = o.pool.Put(floatToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = o.pool.Put(floatToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = o.pool.Put(floatToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = o.pool.Put(floatToString(s.data[0]) + *o.data[0])
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
							result[i] = o.pool.Put(floatToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[0]) + *o.data[i])
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
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[0])
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
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(floatToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesFloat64) Sub(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] - float64(o.data[0])
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] - o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] - o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] - o.data[0]
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] - o.data[0]
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesFloat64{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesFloat64{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesFloat64) Eq(other Series) Series {
	switch o := other.(type) {
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] == float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] == float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] == float64(o.data[0])
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
							result[i] = s.data[0] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == float64(o.data[i])
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
							result[i] = s.data[i] == float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[0])
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
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] == float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] == float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] == float64(o.data[0])
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
							result[i] = s.data[0] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == float64(o.data[i])
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
							result[i] = s.data[i] == float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[0])
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
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] == o.data[0]
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
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
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
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
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
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
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

func (s SeriesFloat64) Ne(other Series) Series {
	switch o := other.(type) {
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] != float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] != float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] != float64(o.data[0])
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
							result[i] = s.data[0] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != float64(o.data[i])
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
							result[i] = s.data[i] != float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[0])
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
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] != float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] != float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] != float64(o.data[0])
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
							result[i] = s.data[0] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != float64(o.data[i])
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
							result[i] = s.data[i] != float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[0])
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
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] != o.data[0]
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
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
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
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
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
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
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

func (s SeriesFloat64) Gt(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
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
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] > float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] > float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] > float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] > float64(o.data[0])
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
							result[i] = s.data[0] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > float64(o.data[i])
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
							result[i] = s.data[i] > float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[0])
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
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] > float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] > float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] > float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] > float64(o.data[0])
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
							result[i] = s.data[0] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > float64(o.data[i])
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
							result[i] = s.data[i] > float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[0])
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
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] > o.data[0]
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
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
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
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
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
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
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

func (s SeriesFloat64) Ge(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
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
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] >= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] >= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] >= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] >= float64(o.data[0])
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
							result[i] = s.data[0] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= float64(o.data[i])
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
							result[i] = s.data[i] >= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[0])
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
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] >= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] >= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] >= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] >= float64(o.data[0])
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
							result[i] = s.data[0] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= float64(o.data[i])
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
							result[i] = s.data[i] >= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[0])
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
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] >= o.data[0]
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
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
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
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
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
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
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

func (s SeriesFloat64) Lt(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
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
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] < float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] < float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] < float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] < float64(o.data[0])
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
							result[i] = s.data[0] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < float64(o.data[i])
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
							result[i] = s.data[i] < float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[0])
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
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] < float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] < float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] < float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] < float64(o.data[0])
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
							result[i] = s.data[0] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < float64(o.data[i])
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
							result[i] = s.data[i] < float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[0])
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
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] < o.data[0]
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
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
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
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
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
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
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

func (s SeriesFloat64) Le(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
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
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
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
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] <= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] <= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] <= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] <= float64(o.data[0])
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
							result[i] = s.data[0] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= float64(o.data[i])
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
							result[i] = s.data[i] <= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[0])
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
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] <= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] <= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] <= float64(o.data[0])
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] <= float64(o.data[0])
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
							result[i] = s.data[0] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= float64(o.data[i])
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
							result[i] = s.data[i] <= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[0])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[0])
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
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= float64(o.data[i])
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result}
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
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						result[0] = s.data[0] <= o.data[0]
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
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
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
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
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
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, true)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, true)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
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
