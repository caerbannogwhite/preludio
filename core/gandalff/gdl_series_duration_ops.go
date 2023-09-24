package gandalff

import (
	"fmt"
)

func (s SeriesDuration) Mul(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Div(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Mod(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}
}

func (s SeriesDuration) Pow(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use power %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Add(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Sub(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Eq(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Ne(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Gt(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Ge(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Lt(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Le(other Series) Series {
	switch o := other.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}
