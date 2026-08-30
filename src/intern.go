package preludiocore

import (
	"fmt"

	"github.com/caerbannogwhite/enchanter/dataframe"
	"github.com/caerbannogwhite/enchanter/meta"
	"github.com/caerbannogwhite/enchanter/series"
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
		e[0] = series.NewSeriesBool([]bool{v}, nil, false, vm.__context)
	case []bool:
		e[0] = series.NewSeriesBool(v, nil, false, vm.__context)
	case int64:
		e[0] = series.NewSeriesInt64([]int64{v}, nil, false, vm.__context)
	case []int64:
		e[0] = series.NewSeriesInt64(v, nil, false, vm.__context)
	case float64:
		e[0] = series.NewSeriesFloat64([]float64{v}, nil, false, vm.__context)
	case []float64:
		e[0] = series.NewSeriesFloat64(v, nil, false, vm.__context)
	case string:
		e[0] = series.NewSeriesString([]string{v}, nil, false, vm.__context)
	case []string:
		e[0] = series.NewSeriesString(v, nil, false, vm.__context)
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

func (i *__p_intern__) toResult(res *[]meta.Columnar, fullOutput bool, outputSnippetLength int) error {
	switch i.tag {
	case PRELUDIO_INTERNAL_TAG_EXPRESSION, PRELUDIO_INTERNAL_TAG_NAMED_PARAM, PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
		switch v := i.expr[0].(type) {
		case series.Series:
			*res = append(*res, seriesToColumnar(fullOutput, outputSnippetLength, i.name, v))

		case dataframe.DataFrame:
			df := dataFrameToColumnar(fullOutput, outputSnippetLength, &v)
			*res = append(*res, df...)
		}
	}
	return nil
}

func (i *__p_intern__) getValue() interface{} {
	return i.expr[0]
}

func (i *__p_intern__) isBoolScalar() bool {
	if s, ok := i.expr[0].(series.Bools); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isBoolVector() bool {
	if _, ok := i.expr[0].(series.Bools); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getBoolScalar() (bool, error) {
	if s, ok := i.expr[0].(series.Bools); ok && s.Len() == 1 {
		return s.Get(0).(bool), nil
	}
	return false, fmt.Errorf("expecting bool scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getBoolVector() ([]bool, error) {
	if s, ok := i.expr[0].(series.Bools); ok {
		return s.Data().([]bool), nil
	}
	return []bool{}, fmt.Errorf("expecting bool vector, got %T", i.expr[0])
}

func (i *__p_intern__) isIntScalar() bool {
	if s, ok := i.expr[0].(series.Ints); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isIntVector() bool {
	if _, ok := i.expr[0].(series.Ints); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getIntScalar() (int, error) {
	if s, ok := i.expr[0].(series.Ints); ok && s.Len() == 1 {
		return s.Get(0).(int), nil
	}
	return 0, fmt.Errorf("expecting int scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getIntVector() ([]int, error) {
	if s, ok := i.expr[0].(series.Ints); ok {
		return s.Data().([]int), nil
	}
	return []int{}, fmt.Errorf("expecting int vector, got %T", i.expr[0])
}

func (i *__p_intern__) isInt64Scalar() bool {
	if s, ok := i.expr[0].(series.Int64s); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isInt64Vector() bool {
	if _, ok := i.expr[0].(series.Int64s); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getInt64Scalar() (int64, error) {
	if s, ok := i.expr[0].(series.Int64s); ok && s.Len() == 1 {
		return s.Get(0).(int64), nil
	}
	return 0, fmt.Errorf("expecting int64 scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getInt64Vector() ([]int64, error) {
	if s, ok := i.expr[0].(series.Int64s); ok {
		return s.Data().([]int64), nil
	}
	return []int64{}, fmt.Errorf("expecting int64 vector, got %T", i.expr[0])
}

func (i *__p_intern__) isFloat64Scalar() bool {
	if s, ok := i.expr[0].(series.Float64s); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isFloat64Vector() bool {
	if _, ok := i.expr[0].(series.Float64s); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getFloat64Scalar() (float64, error) {
	if s, ok := i.expr[0].(series.Float64s); ok && s.Len() == 1 {
		return s.Get(0).(float64), nil
	}
	return 0, fmt.Errorf("expecting float scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getFloat64Vector() ([]float64, error) {
	if s, ok := i.expr[0].(series.Float64s); ok {
		return s.Data().([]float64), nil
	}
	return []float64{}, fmt.Errorf("expecting float vector, got %T", i.expr[0])
}

func (i *__p_intern__) isStringScalar() bool {
	if s, ok := i.expr[0].(series.Strings); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isStringVector() bool {
	if _, ok := i.expr[0].(series.Strings); ok {
		return true
	}
	return false
}

func (i *__p_intern__) getStringScalar() (string, error) {
	if s, ok := i.expr[0].(series.Strings); ok && s.Len() == 1 {
		return s.Get(0).(string), nil
	}
	return "", fmt.Errorf("expecting string scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getStringVector() ([]string, error) {
	if s, ok := i.expr[0].(series.Strings); ok {
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
	_, ok := i.expr[0].(dataframe.DataFrame)
	return ok
}

func (i *__p_intern__) getDataframe() (dataframe.DataFrame, error) {
	switch v := i.expr[0].(type) {
	case dataframe.DataFrame:
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

func (i *__p_intern__) listToSeriesBool() (series.Series, error) {
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
		return series.NewSeriesBool(res, nil, false, i.vm.__context), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.expr[0])
	}
}

func (i *__p_intern__) listToSeriesInt64() (series.Series, error) {
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
		return series.NewSeriesInt64(res, nil, false, i.vm.__context), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.expr[0])
	}
}

func (i *__p_intern__) listToSeriesFloat64() (series.Series, error) {
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
		return series.NewSeriesFloat64(res, nil, false, i.vm.__context), nil

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

func (i *__p_intern__) listToSeriesString() (series.Series, error) {
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
		return series.NewSeriesString(res, nil, false, i.vm.__context), nil

	default:
		return nil, fmt.Errorf("expecting list of strings")
	}
}

// isNeg returns true if the expression is a negative number
// used for special cases like orderBy
func (i *__p_intern__) isNeg() bool {
	if len(i.expr) == 2 {
		if op, ok := i.expr[1].(meta.OPCODE); ok && op == meta.OP_UNARY_SUB {
			return true
		}
	}
	return false
}

func (lhs *__p_intern__) appendBinaryOperation(op meta.OPCODE, rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, op)
}

func (rhs *__p_intern__) appendUnaryOperation(op meta.OPCODE) {
	rhs.expr = append(rhs.expr, op)
}

func isOperator(t interface{}) (meta.OPCODE, bool) {
	if v, ok := t.(meta.OPCODE); ok {
		return v, true
	}
	return meta.NO_OP, false
}
