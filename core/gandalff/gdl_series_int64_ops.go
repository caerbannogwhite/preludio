package gandalff

import (
	"fmt"
	"math"
)

func (s SeriesInt64) Neg() Series {
	for i, v := range s.data {
		s.data[i] = -v
	}
	return s
}

func (s SeriesInt64) Mul(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] * int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] * int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] * int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] * int64(o.data[0])
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[0])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * int64(o.data[i])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] * o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] * o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] * o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] * o.data[0]
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) * o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) * o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) * o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) * o.data[0]
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesInt64) Div(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) / o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) / o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) / o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) / o.data[0]
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesInt64) Mod(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesInt64) Pow(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesInt64) Add(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] + int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] + int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] + int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] + int64(o.data[0])
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[0])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + int64(o.data[i])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] + o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] + o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] + o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] + o.data[0]
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) + o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) + o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) + o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) + o.data[0]
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result[0] = o.pool.Put(intToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = o.pool.Put(intToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = o.pool.Put(intToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = o.pool.Put(intToString(s.data[0]) + *o.data[0])
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
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
							result[i] = o.pool.Put(intToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[0]) + *o.data[i])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
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
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[0])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
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
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: true, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = o.pool.Put(intToString(s.data[i]) + *o.data[i])
						}
						return SeriesString{isNullable: false, name: s.name, nullMask: resultNullMask, pool: o.pool, data: result}
					}
				}
			}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesInt64) Sub(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] - int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] - int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] - int64(o.data[0])
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] - int64(o.data[0])
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[0])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[0])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - int64(o.data[i])
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] - o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] - o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] - o.data[0]
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] - o.data[0]
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesInt64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesInt64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) - o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) - o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) - o.data[0]
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) - o.data[0]
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return SeriesFloat64{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return SeriesFloat64{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
					}
				}
			}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesInt64) Eq(other Series) Series {
	switch o := other.(type) {
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] == int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] == int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] == int64(o.data[0])
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
							result[i] = s.data[0] == int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == int64(o.data[i])
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
							result[i] = s.data[0] == int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == int64(o.data[i])
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
							result[i] = s.data[i] == int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == int64(o.data[0])
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
							result[i] = s.data[i] == int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == int64(o.data[0])
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
							result[i] = s.data[i] == int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == int64(o.data[i])
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
							result[i] = s.data[i] == int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == int64(o.data[i])
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] == o.data[0]
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
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
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
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
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
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
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
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
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
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
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
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) == o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) == o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) == o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) == o.data[0]
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
							result[i] = float64(s.data[0]) == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) == o.data[i]
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
							result[i] = float64(s.data[0]) == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) == o.data[i]
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
							result[i] = float64(s.data[i]) == o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[0]
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
							result[i] = float64(s.data[i]) == o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[0]
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
							result[i] = float64(s.data[i]) == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[i]
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
							result[i] = float64(s.data[i]) == o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[i]
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

func (s SeriesInt64) Ne(other Series) Series {
	switch o := other.(type) {
	case SeriesInt32:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] != int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] != int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] != int64(o.data[0])
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
							result[i] = s.data[0] != int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != int64(o.data[i])
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
							result[i] = s.data[0] != int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != int64(o.data[i])
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
							result[i] = s.data[i] != int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != int64(o.data[0])
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
							result[i] = s.data[i] != int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != int64(o.data[0])
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
							result[i] = s.data[i] != int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != int64(o.data[i])
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
							result[i] = s.data[i] != int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != int64(o.data[i])
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] != o.data[0]
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
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
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
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
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
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
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
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
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
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
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
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) != o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) != o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) != o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) != o.data[0]
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
							result[i] = float64(s.data[0]) != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) != o.data[i]
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
							result[i] = float64(s.data[0]) != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) != o.data[i]
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
							result[i] = float64(s.data[i]) != o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[0]
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
							result[i] = float64(s.data[i]) != o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[0]
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
							result[i] = float64(s.data[i]) != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[i]
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
							result[i] = float64(s.data[i]) != o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[i]
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

func (s SeriesInt64) Gt(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] > int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] > int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] > int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] > int64(o.data[0])
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
							result[i] = s.data[0] > int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > int64(o.data[i])
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
							result[i] = s.data[0] > int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > int64(o.data[i])
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
							result[i] = s.data[i] > int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > int64(o.data[0])
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
							result[i] = s.data[i] > int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > int64(o.data[0])
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
							result[i] = s.data[i] > int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > int64(o.data[i])
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
							result[i] = s.data[i] > int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > int64(o.data[i])
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] > o.data[0]
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
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
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
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
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
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
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
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
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
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
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
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) > o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) > o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) > o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) > o.data[0]
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
							result[i] = float64(s.data[0]) > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) > o.data[i]
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
							result[i] = float64(s.data[0]) > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) > o.data[i]
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
							result[i] = float64(s.data[i]) > o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[0]
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
							result[i] = float64(s.data[i]) > o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[0]
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
							result[i] = float64(s.data[i]) > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[i]
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
							result[i] = float64(s.data[i]) > o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[i]
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

func (s SeriesInt64) Ge(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] >= int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] >= int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] >= int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] >= int64(o.data[0])
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
							result[i] = s.data[0] >= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= int64(o.data[i])
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
							result[i] = s.data[0] >= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= int64(o.data[i])
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
							result[i] = s.data[i] >= int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= int64(o.data[0])
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
							result[i] = s.data[i] >= int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= int64(o.data[0])
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
							result[i] = s.data[i] >= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= int64(o.data[i])
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
							result[i] = s.data[i] >= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= int64(o.data[i])
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] >= o.data[0]
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
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
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
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
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
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
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
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
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
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
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
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) >= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) >= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) >= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) >= o.data[0]
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
							result[i] = float64(s.data[0]) >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) >= o.data[i]
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
							result[i] = float64(s.data[0]) >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) >= o.data[i]
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
							result[i] = float64(s.data[i]) >= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[0]
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
							result[i] = float64(s.data[i]) >= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[0]
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
							result[i] = float64(s.data[i]) >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[i]
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
							result[i] = float64(s.data[i]) >= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[i]
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

func (s SeriesInt64) Lt(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] < int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] < int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] < int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] < int64(o.data[0])
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
							result[i] = s.data[0] < int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < int64(o.data[i])
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
							result[i] = s.data[0] < int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < int64(o.data[i])
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
							result[i] = s.data[i] < int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < int64(o.data[0])
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
							result[i] = s.data[i] < int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < int64(o.data[0])
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
							result[i] = s.data[i] < int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < int64(o.data[i])
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
							result[i] = s.data[i] < int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < int64(o.data[i])
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] < o.data[0]
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
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
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
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
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
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
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
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
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
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
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
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) < o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) < o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) < o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) < o.data[0]
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
							result[i] = float64(s.data[0]) < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) < o.data[i]
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
							result[i] = float64(s.data[0]) < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) < o.data[i]
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
							result[i] = float64(s.data[i]) < o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[0]
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
							result[i] = float64(s.data[i]) < o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[0]
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
							result[i] = float64(s.data[i]) < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[i]
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
							result[i] = float64(s.data[i]) < o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[i]
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

func (s SeriesInt64) Le(other Series) Series {
	switch o := other.(type) {
	case SeriesBool:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						b2 := int64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
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
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
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
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							b2 := int64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] <= int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] <= int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] <= int64(o.data[0])
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] <= int64(o.data[0])
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
							result[i] = s.data[0] <= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= int64(o.data[i])
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
							result[i] = s.data[0] <= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= int64(o.data[i])
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
							result[i] = s.data[i] <= int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= int64(o.data[0])
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
							result[i] = s.data[i] <= int64(o.data[0])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= int64(o.data[0])
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
							result[i] = s.data[i] <= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= int64(o.data[i])
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
							result[i] = s.data[i] <= int64(o.data[i])
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= int64(o.data[i])
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = s.data[0] <= o.data[0]
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
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
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
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
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
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
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
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
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
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
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
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: false, name: s.name, nullMask: resultNullMask, data: result}
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
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) <= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						result[0] = float64(s.data[0]) <= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					}
				} else {
					if o.isNullable {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, o.nullMask)
						result[0] = float64(s.data[0]) <= o.data[0]
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						result[0] = float64(s.data[0]) <= o.data[0]
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
							result[i] = float64(s.data[0]) <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) <= o.data[i]
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
							result[i] = float64(s.data[0]) <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(o.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) <= o.data[i]
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
							result[i] = float64(s.data[i]) <= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[0]
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
							result[i] = float64(s.data[i]) <= o.data[0]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[0]
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
							result[i] = float64(s.data[i]) <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[i]
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
							result[i] = float64(s.data[i]) <= o.data[i]
						}
						return SeriesBool{isNullable: true, name: s.name, nullMask: resultNullMask, data: result}
					} else {
						resultSize := len(s.data)
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[i]
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
