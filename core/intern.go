package preludiocore

import (
	"fmt"
	"typesys"

	"gandalff"
)

type __p_intern_tag__ uint8

const (
	// PRELUDIO_INTERNAL_TAG_ERROR       __p_intern_tag__ = 0
	PRELUDIO_INTERNAL_TAG_EXPRESSION  __p_intern_tag__ = 1
	PRELUDIO_INTERNAL_TAG_NAMED_PARAM __p_intern_tag__ = 2
	PRELUDIO_INTERNAL_TAG_ASSIGNMENT  __p_intern_tag__ = 3
	PRELUDIO_INTERNAL_TAG_BEGIN_FRAME __p_intern_tag__ = 4
)

type __p_intern__ struct {
	tag  __p_intern_tag__
	vm   *ByteEater
	expr []interface{}
	name string
}

func (vm *ByteEater) newPInternBeginFrame() *__p_intern__ {
	return &__p_intern__{tag: PRELUDIO_INTERNAL_TAG_BEGIN_FRAME}
}

func (vm *ByteEater) newPInternTerm(val interface{}) *__p_intern__ {
	e := make([]interface{}, 1)

	switch v := val.(type) {
	case bool:
		e[0] = gandalff.NewSeriesBool("", true, false, []bool{v})
	case []bool:
		e[0] = gandalff.NewSeriesBool("", true, false, v)
	case int64:
		e[0] = gandalff.NewSeriesInt64("", true, false, []int64{v})
	case []int64:
		e[0] = gandalff.NewSeriesInt64("", true, false, v)
	case float64:
		e[0] = gandalff.NewSeriesFloat64("", true, false, []float64{v})
	case []float64:
		e[0] = gandalff.NewSeriesFloat64("", true, false, v)
	case string:
		e[0] = gandalff.NewSeriesString("", true, []string{v}, vm.__stringPool)
	case []string:
		e[0] = gandalff.NewSeriesString("", true, v, vm.__stringPool)
	default:
		e[0] = v
	}

	return &__p_intern__{tag: PRELUDIO_INTERNAL_TAG_EXPRESSION, vm: vm, expr: e}
}

func (i *__p_intern__) setParamName(name string) {
	i.tag = PRELUDIO_INTERNAL_TAG_NAMED_PARAM
	i.name = name
}

func (i *__p_intern__) setAssignment(name string) {
	i.tag = PRELUDIO_INTERNAL_TAG_ASSIGNMENT
	i.name = name
}

func (i *__p_intern__) toResult(res *[]typesys.Columnar, fullOutput bool, outputSnippetLength int) error {
	switch i.tag {
	case PRELUDIO_INTERNAL_TAG_EXPRESSION:
		switch v := i.expr[0].(type) {
		case gandalff.SeriesBool:
			*res = append(*res, NewColumnarBool(i.name, fullOutput, outputSnippetLength, v.Bools()))
		case gandalff.SeriesInt64:
			*res = append(*res, NewColumnarInt64(i.name, fullOutput, outputSnippetLength, v.Int64s()))
		case gandalff.SeriesFloat64:
			*res = append(*res, NewColumnarFloat(i.name, fullOutput, outputSnippetLength, v.Float64s()))
		case gandalff.SeriesString:
			*res = append(*res, NewColumnarString(i.name, fullOutput, outputSnippetLength, v.Strings()))
		case gandalff.DataFrame:
			df, err := DataFrameToColumnar(fullOutput, outputSnippetLength, &v)
			if err != nil {
				return err
			}
			*res = append(*res, df...)
		}

	case PRELUDIO_INTERNAL_TAG_NAMED_PARAM:
		switch v := i.expr[0].(type) {
		case gandalff.SeriesBool:
			*res = append(*res, NewColumnarBool(i.name, fullOutput, outputSnippetLength, v.Bools()))
		case gandalff.SeriesInt64:
			*res = append(*res, NewColumnarInt64(i.name, fullOutput, outputSnippetLength, v.Int64s()))
		case gandalff.SeriesFloat64:
			*res = append(*res, NewColumnarFloat(i.name, fullOutput, outputSnippetLength, v.Float64s()))
		case gandalff.SeriesString:
			*res = append(*res, NewColumnarString(i.name, fullOutput, outputSnippetLength, v.Strings()))
		case gandalff.DataFrame:
			df, err := DataFrameToColumnar(fullOutput, outputSnippetLength, &v)
			if err != nil {
				return err
			}
			*res = append(*res, df...)
		}

	case PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
		switch v := i.expr[0].(type) {
		case gandalff.SeriesBool:
			*res = append(*res, NewColumnarBool(i.name, fullOutput, outputSnippetLength, v.Bools()))
		case gandalff.SeriesInt64:
			*res = append(*res, NewColumnarInt64(i.name, fullOutput, outputSnippetLength, v.Int64s()))
		case gandalff.SeriesFloat64:
			*res = append(*res, NewColumnarFloat(i.name, fullOutput, outputSnippetLength, v.Float64s()))
		case gandalff.SeriesString:
			*res = append(*res, NewColumnarString(i.name, fullOutput, outputSnippetLength, v.Strings()))
		case gandalff.DataFrame:
			df, err := DataFrameToColumnar(fullOutput, outputSnippetLength, &v)
			if err != nil {
				return err
			}
			*res = append(*res, df...)
		}
	}
	return nil
}

func (i *__p_intern__) getValue() interface{} {
	return i.expr[0]
}

func (i *__p_intern__) isBoolScalar() bool {
	if s, ok := i.expr[0].(gandalff.SeriesBool); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isBoolVector() bool {
	if _, ok := i.expr[0].(gandalff.SeriesBool); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getBoolScalar() (bool, error) {
	if s, ok := i.expr[0].(gandalff.SeriesBool); ok && s.Len() == 1 {
		return s.Get(0).(bool), nil
	}
	return false, fmt.Errorf("expecting bool scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getBoolVector() ([]bool, error) {
	if s, ok := i.expr[0].(gandalff.SeriesBool); ok {
		return s.Data().([]bool), nil
	}
	return []bool{}, fmt.Errorf("expecting bool vector, got %T", i.expr[0])
}

func (i *__p_intern__) isInt32Scalar() bool {
	if s, ok := i.expr[0].(gandalff.SeriesInt32); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isInt32Vector() bool {
	if _, ok := i.expr[0].(gandalff.SeriesInt32); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getInt32Scalar() (int32, error) {
	if s, ok := i.expr[0].(gandalff.SeriesInt32); ok && s.Len() == 1 {
		return s.Get(0).(int32), nil
	}
	return 0, fmt.Errorf("expecting int32 scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getInt32Vector() ([]int32, error) {
	if s, ok := i.expr[0].(gandalff.SeriesInt32); ok {
		return s.Data().([]int32), nil
	}
	return []int32{}, fmt.Errorf("expecting int32 vector, got %T", i.expr[0])
}

func (i *__p_intern__) isInt64Scalar() bool {
	if s, ok := i.expr[0].(gandalff.SeriesInt64); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isInt64Vector() bool {
	if _, ok := i.expr[0].(gandalff.SeriesInt64); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getInt64Scalar() (int64, error) {
	if s, ok := i.expr[0].(gandalff.SeriesInt64); ok && s.Len() == 1 {
		return s.Get(0).(int64), nil
	}
	return 0, fmt.Errorf("expecting int64 scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getInt64Vector() ([]int64, error) {
	if s, ok := i.expr[0].(gandalff.SeriesInt64); ok {
		return s.Data().([]int64), nil
	}
	return []int64{}, fmt.Errorf("expecting int64 vector, got %T", i.expr[0])
}

func (i *__p_intern__) isFloat64Scalar() bool {
	if s, ok := i.expr[0].(gandalff.SeriesFloat64); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isFloat64Vector() bool {
	if _, ok := i.expr[0].(gandalff.SeriesFloat64); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getFloat64Scalar() (float64, error) {
	if s, ok := i.expr[0].(gandalff.SeriesFloat64); ok && s.Len() == 1 {
		return s.Get(0).(float64), nil
	}
	return 0, fmt.Errorf("expecting float scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getFloat64Vector() ([]float64, error) {
	if s, ok := i.expr[0].(gandalff.SeriesFloat64); ok {
		return s.Data().([]float64), nil
	}
	return []float64{}, fmt.Errorf("expecting float vector, got %T", i.expr[0])
}

func (i *__p_intern__) isStringScalar() bool {
	if s, ok := i.expr[0].(gandalff.SeriesString); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isStringVector() bool {
	if _, ok := i.expr[0].(gandalff.SeriesString); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getStringScalar() (string, error) {
	if s, ok := i.expr[0].(gandalff.SeriesString); ok && s.Len() == 1 {
		return s.Get(0).(string), nil
	}
	return "", fmt.Errorf("expecting string scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getStringVector() ([]string, error) {
	if s, ok := i.expr[0].(gandalff.SeriesString); ok {
		return s.Data().([]string), nil
	}
	return []string{}, fmt.Errorf("expecting string vector, got %T", i.expr[0])
}

func (i *__p_intern__) isSymbol() bool {
	_, ok := i.expr[0].(__p_symbol__)
	return ok
}

func (i *__p_intern__) getSymbol() (__p_symbol__, error) {
	if v, ok := i.expr[0].(__p_symbol__); ok {
		return v, nil
	}
	return "", fmt.Errorf("expecting symbol, got %T", i.expr[0])
}

func (i *__p_intern__) isDataframe() bool {
	_, ok := i.expr[0].(gandalff.DataFrame)
	return ok
}

func (i *__p_intern__) getDataframe() (gandalff.DataFrame, error) {
	switch v := i.expr[0].(type) {
	case gandalff.DataFrame:
		return v, nil
	default:
		return nil, fmt.Errorf("expecting dataframe, got %T", v)
	}
}

func (i *__p_intern__) isList() bool {
	_, ok := i.expr[0].(__p_list__)
	return ok
}

func (i *__p_intern__) getList() (__p_list__, error) {
	switch v := i.expr[0].(type) {
	case __p_list__:
		return v, nil
	default:
		return nil, fmt.Errorf("expecting list, got %T", v)
	}
}

func (p *__p_intern__) processList() error {
	var series gandalff.Series

	list := p.expr[0].(__p_list__)
	for _, q := range list {
		switch v := q.expr[0].(type) {
		// case __p_list__:

		case gandalff.Series:
			if series == nil {
				series = v
			} else if series.Type() == v.Type() {
				series = series.Append(v)
			} else if series.Type().CanCoerceTo(v.Type()) {
				series = series.Cast(v.Type(), p.vm.__stringPool).Append(v)
			} else {
				return fmt.Errorf("cannot append %s to %s", v.Type().ToString(), series.Type().ToString())
			}
		}
	}

	p.expr[0] = series

	return nil
}

func (i *__p_intern__) listToSeries() (gandalff.Series, error) {
	switch l := i.expr[0].(type) {
	case __p_list__:
		switch l[0].getValue().(type) {
		case gandalff.SeriesBool:
			s := gandalff.NewSeriesBool(i.name, true, false, []bool{})
			for _, e := range l {
				switch v := e.getValue().(type) {
				case gandalff.SeriesBool:
					s.Append(v)
				default:
					return nil, fmt.Errorf("expecting list of bools, got %T", e.getValue())
				}
			}

			return s, nil

		case gandalff.SeriesInt64:
			s := gandalff.NewSeriesInt64(i.name, true, false, []int64{})
			for _, e := range l {
				switch v := e.getValue().(type) {
				case gandalff.SeriesInt64:
					s.Append(v)
				default:
					return nil, fmt.Errorf("expecting list of ints, got %T", e.getValue())
				}
			}

			return s, nil

		case gandalff.SeriesFloat64:
			s := gandalff.NewSeriesFloat64(i.name, true, false, []float64{})
			for _, e := range l {
				switch v := e.getValue().(type) {
				case gandalff.SeriesFloat64:
					s.Append(v)
				default:
					return nil, fmt.Errorf("expecting list of floats, got %T", e.getValue())
				}
			}

			return s, nil

		case gandalff.SeriesString:
			s := gandalff.NewSeriesString(i.name, true, []string{}, i.vm.__stringPool)
			for _, e := range l {
				switch v := e.getValue().(type) {
				case gandalff.SeriesString:
					s.Append(v)
				default:
					return nil, fmt.Errorf("expecting list of strings, got %T", e.getValue())
				}
			}

			return s, nil

		default:
			return nil, fmt.Errorf("expecting list of scalars, got %T", l[0].getValue())
		}

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.expr[0])
	}
}

func (i *__p_intern__) listToSeriesBool() (gandalff.Series, error) {
	switch l := i.expr[0].(type) {
	case __p_list__:
		res := make([]bool, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case bool:
				res[j] = v
			default:
				return nil, fmt.Errorf("expecting list of bools, got %T", e.getValue())
			}
		}
		return gandalff.NewSeriesBool("", true, false, res), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.expr[0])
	}
}

func (i *__p_intern__) listToSeriesInt64() (gandalff.Series, error) {
	switch l := i.expr[0].(type) {
	case __p_list__:
		res := make([]int64, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case int64:
				res[j] = v
			default:
				return nil, fmt.Errorf("expecting list of ints, got %T", e.getValue())
			}
		}
		return gandalff.NewSeriesInt64("", true, false, res), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.expr[0])
	}
}

func (i *__p_intern__) listToSeriesFloat64() (gandalff.Series, error) {
	switch l := i.expr[0].(type) {
	case __p_list__:
		res := make([]float64, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case float64:
				res[j] = v
			default:
				return nil, fmt.Errorf("expecting list of floats, got %T", e.getValue())
			}
		}
		return gandalff.NewSeriesFloat64("", true, false, res), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.expr[0])
	}
}

func (i *__p_intern__) listToStringSlice() ([]string, error) {
	switch l := i.expr[0].(type) {
	case __p_list__:
		res := make([]string, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case string:
				res[j] = v
			case __p_symbol__:
				res[j] = string(v)
			default:
				return nil, fmt.Errorf("expecting list of strings or symbols, got %T", e.getValue())
			}
		}
		return res, nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.expr[0])
	}
}

func (i *__p_intern__) listToSeriesString() (gandalff.Series, error) {
	switch l := i.expr[0].(type) {
	case __p_list__:
		res := make([]string, len(l))
		for j, e := range l {
			v, err := e.getStringScalar()
			if err != nil {
				return nil, fmt.Errorf("expecting list of strings")
			}
			res[j] = v
		}
		return gandalff.NewSeriesString("", true, res, i.vm.__stringPool), nil

	default:
		return nil, fmt.Errorf("expecting list of strings")
	}
}

// isNeg returns true if the expression is a negative number
// used for special cases like orderBy
func (i *__p_intern__) isNeg() bool {
	if len(i.expr) == 2 {
		if op, ok := i.expr[1].(typesys.OPCODE); ok && op == typesys.OP_UNARY_SUB {
			return true
		}
	}
	return false
}

func (lhs *__p_intern__) appendBinaryOperation(op typesys.OPCODE, rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, op)
}

func (rhs *__p_intern__) appendUnaryOperation(op typesys.OPCODE) {
	rhs.expr = append(rhs.expr, op)
}

func isOperator(t interface{}) (typesys.OPCODE, bool) {
	if v, ok := t.(typesys.OPCODE); ok {
		return v, true
	}
	return typesys.NO_OP, false
}
