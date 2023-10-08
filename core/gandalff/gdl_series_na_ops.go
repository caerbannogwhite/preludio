package gandalff

import "fmt"

func (s SeriesNA) Mul(other Series) Series {
	switch o := other.(type) {
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
	case SeriesBool:
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
	case SeriesInt:
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
	case SeriesInt64:
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
	case SeriesFloat64:
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

func (s SeriesNA) Div(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesNA) Mod(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesNA) Exp(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesNA) Add(other Series) Series {
	switch o := other.(type) {
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
	case SeriesBool:
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
	case SeriesInt:
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
	case SeriesInt64:
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
	case SeriesFloat64:
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
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				result := make([]*string, resultSize)
				resultNullMask := __binVecInit(0, false)
				result[0] = s.pool.Put(NULL_STRING + *o.data[0])
				return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
			} else {
				resultSize := o.Len()
				result := make([]*string, resultSize)
				resultNullMask := __binVecInit(0, false)
				for i := 0; i < resultSize; i++ {
					result[i] = s.pool.Put(NULL_STRING + *o.data[i])
				}
				return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				result := make([]*string, resultSize)
				resultNullMask := __binVecInit(0, false)
				for i := 0; i < resultSize; i++ {
					result[i] = s.pool.Put(NULL_STRING + *o.data[0])
				}
				return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				result := make([]*string, resultSize)
				resultNullMask := __binVecInit(0, false)
				for i := 0; i < resultSize; i++ {
					result[i] = s.pool.Put(NULL_STRING + *o.data[i])
				}
				return SeriesString{isNullable: false, nullMask: resultNullMask, pool: o.pool, data: result}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
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
	case SeriesDuration:
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

func (s SeriesNA) Sub(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
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
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
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
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesNA) Eq(other Series) Series {
	switch o := other.(type) {
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
	case SeriesBool:
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
	case SeriesInt:
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
	case SeriesInt64:
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
	case SeriesFloat64:
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
	case SeriesString:
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
	case SeriesTime:
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
	case SeriesDuration:
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

func (s SeriesNA) Ne(other Series) Series {
	switch o := other.(type) {
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
	case SeriesBool:
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
	case SeriesInt:
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
	case SeriesInt64:
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
	case SeriesFloat64:
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
	case SeriesString:
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
	case SeriesTime:
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
	case SeriesDuration:
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

func (s SeriesNA) Gt(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesString:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesNA) Ge(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesString:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
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
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesNA) Lt(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesString:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesNA) Le(other Series) Series {
	switch o := other.(type) {
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesBool:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesInt64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesFloat64:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesString:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
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
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}
