package preludiocore

import (
	"fmt"
	"preludiometa"
	"time"

	"gandalff"
)

type pIntern struct {
	tag   pInternTagType
	name  string
	vm    *ByteEater
	value interface{}
}

func (vm *ByteEater) newPInternBeginFrame() *pIntern {
	return &pIntern{tag: PRELUDIO_INTERNAL_TAG_BEGIN_FRAME}
}

func (vm *ByteEater) newPInternTerm(val interface{}) *pIntern {
	switch v := val.(type) {
	case bool:
		val = gandalff.NewSeriesBool([]bool{v}, nil, false, vm.__context)
	case []bool:
		val = gandalff.NewSeriesBool(v, nil, false, vm.__context)
	case int64:
		val = gandalff.NewSeriesInt64([]int64{v}, nil, false, vm.__context)
	case []int64:
		val = gandalff.NewSeriesInt64(v, nil, false, vm.__context)
	case float64:
		val = gandalff.NewSeriesFloat64([]float64{v}, nil, false, vm.__context)
	case []float64:
		val = gandalff.NewSeriesFloat64(v, nil, false, vm.__context)
	case string:
		val = gandalff.NewSeriesString([]string{v}, nil, false, vm.__context)
	case []string:
		val = gandalff.NewSeriesString(v, nil, false, vm.__context)
	case time.Time:
		val = gandalff.NewSeriesTime([]time.Time{v}, nil, false, vm.__context)
	case []time.Time:
		val = gandalff.NewSeriesTime(v, nil, false, vm.__context)
	case time.Duration:
		val = gandalff.NewSeriesDuration([]time.Duration{v}, nil, false, vm.__context)
	case []time.Duration:
		val = gandalff.NewSeriesDuration(v, nil, false, vm.__context)
	default:
		val = v
	}

	return &pIntern{tag: PRELUDIO_INTERNAL_TAG_EXPRESSION, vm: vm, value: val}
}

func (i *pIntern) setParamName(name string) {
	i.tag = PRELUDIO_INTERNAL_TAG_NAMED_PARAM
	i.name = name
}

func (i *pIntern) setAssignment(name string) {
	i.tag = PRELUDIO_INTERNAL_TAG_ASSIGNMENT
	i.name = name
}

func (i *pIntern) toResult(res *[]preludiometa.Columnar, fullOutput bool, outputSnippetLength int) error {
	switch i.tag {
	case PRELUDIO_INTERNAL_TAG_EXPRESSION, PRELUDIO_INTERNAL_TAG_NAMED_PARAM, PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
		switch v := i.value.(type) {
		case gandalff.Series:
			*res = append(*res, seriesToColumnar(fullOutput, outputSnippetLength, i.name, v))

		case gandalff.DataFrame:
			df := dataFrameToColumnar(fullOutput, outputSnippetLength, &v)
			*res = append(*res, df...)
		}
	}
	return nil
}

func (i *pIntern) getValue() interface{} {
	return i.value
}

func (i *pIntern) isBoolScalar() bool {
	if s, ok := i.value.(gandalff.SeriesBool); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *pIntern) isBoolVector() bool {
	if _, ok := i.value.(gandalff.SeriesBool); ok {
		return true
	}
	return false
}

func (i *pIntern) getBoolScalar() (bool, error) {
	if s, ok := i.value.(gandalff.SeriesBool); ok && s.Len() == 1 {
		return s.Get(0).(bool), nil
	}
	return false, fmt.Errorf("expecting bool scalar, got %T", i.value)
}

func (i *pIntern) getBoolVector() ([]bool, error) {
	if s, ok := i.value.(gandalff.SeriesBool); ok {
		return s.Data().([]bool), nil
	}
	return []bool{}, fmt.Errorf("expecting bool vector, got %T", i.value)
}

func (i *pIntern) isIntScalar() bool {
	if s, ok := i.value.(gandalff.SeriesInt); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *pIntern) isIntVector() bool {
	if _, ok := i.value.(gandalff.SeriesInt); ok {
		return true
	}
	return false
}

func (i *pIntern) getIntScalar() (int, error) {
	if s, ok := i.value.(gandalff.SeriesInt); ok && s.Len() == 1 {
		return s.Get(0).(int), nil
	}
	return 0, fmt.Errorf("expecting int scalar, got %T", i.value)
}

func (i *pIntern) getIntVector() ([]int, error) {
	if s, ok := i.value.(gandalff.SeriesInt); ok {
		return s.Data().([]int), nil
	}
	return []int{}, fmt.Errorf("expecting int vector, got %T", i.value)
}

func (i *pIntern) isInt64Scalar() bool {
	if s, ok := i.value.(gandalff.SeriesInt64); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *pIntern) isInt64Vector() bool {
	if _, ok := i.value.(gandalff.SeriesInt64); ok {
		return true
	}
	return false
}

func (i *pIntern) getInt64Scalar() (int64, error) {
	if s, ok := i.value.(gandalff.SeriesInt64); ok && s.Len() == 1 {
		return s.Get(0).(int64), nil
	}
	return 0, fmt.Errorf("expecting int64 scalar, got %T", i.value)
}

func (i *pIntern) getInt64Vector() ([]int64, error) {
	if s, ok := i.value.(gandalff.SeriesInt64); ok {
		return s.Data().([]int64), nil
	}
	return []int64{}, fmt.Errorf("expecting int64 vector, got %T", i.value)
}

func (i *pIntern) isFloat64Scalar() bool {
	if s, ok := i.value.(gandalff.SeriesFloat64); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *pIntern) isFloat64Vector() bool {
	if _, ok := i.value.(gandalff.SeriesFloat64); ok {
		return true
	}
	return false
}

func (i *pIntern) getFloat64Scalar() (float64, error) {
	if s, ok := i.value.(gandalff.SeriesFloat64); ok && s.Len() == 1 {
		return s.Get(0).(float64), nil
	}
	return 0, fmt.Errorf("expecting float scalar, got %T", i.value)
}

func (i *pIntern) getFloat64Vector() ([]float64, error) {
	if s, ok := i.value.(gandalff.SeriesFloat64); ok {
		return s.Data().([]float64), nil
	}
	return []float64{}, fmt.Errorf("expecting float vector, got %T", i.value)
}

func (i *pIntern) isStringScalar() bool {
	if s, ok := i.value.(gandalff.SeriesString); ok && s.Len() == 1 {
		return true
	}
	return false
}

func (i *pIntern) isStringVector() bool {
	if _, ok := i.value.(gandalff.SeriesString); ok {
		return true
	}
	return false
}

func (i *pIntern) getStringScalar() (string, error) {
	if s, ok := i.value.(gandalff.SeriesString); ok && s.Len() == 1 {
		return s.Get(0).(string), nil
	}
	return "", fmt.Errorf("expecting string scalar, got %T", i.value)
}

func (i *pIntern) getStringVector() ([]string, error) {
	if s, ok := i.value.(gandalff.SeriesString); ok {
		return s.Data().([]string), nil
	}
	return []string{}, fmt.Errorf("expecting string vector, got %T", i.value)
}

func (i *pIntern) isSymbol() bool {
	_, ok := i.value.(pSymbol)
	return ok
}

func (i *pIntern) getSymbol() (pSymbol, error) {
	if v, ok := i.value.(pSymbol); ok {
		return v, nil
	}
	return "", fmt.Errorf("expecting symbol, got %T", i.value)
}

func (i *pIntern) isDataframe() bool {
	_, ok := i.value.(gandalff.DataFrame)
	return ok
}

func (i *pIntern) getDataframe() (gandalff.DataFrame, error) {
	switch v := i.value.(type) {
	case gandalff.DataFrame:
		return v, nil
	default:
		return nil, fmt.Errorf("expecting dataframe, got %T", v)
	}
}

func (i *pIntern) isList() bool {
	_, ok := i.value.(pList)
	return ok
}

func (i *pIntern) getList() (pList, error) {
	switch v := i.value.(type) {
	case pList:
		return v, nil
	default:
		return nil, fmt.Errorf("expecting list, got %T", v)
	}
}

func (i *pIntern) listToSeriesBool() (gandalff.Series, error) {
	switch l := i.value.(type) {
	case pList:
		res := make([]bool, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case bool:
				res[j] = v
			default:
				return nil, fmt.Errorf("expecting list of bools, got %T", e.getValue())
			}
		}
		return gandalff.NewSeriesBool(res, nil, false, i.vm.__context), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.value)
	}
}

func (i *pIntern) listToSeriesInt64() (gandalff.Series, error) {
	switch l := i.value.(type) {
	case pList:
		res := make([]int64, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case int64:
				res[j] = v
			default:
				return nil, fmt.Errorf("expecting list of ints, got %T", e.getValue())
			}
		}
		return gandalff.NewSeriesInt64(res, nil, false, i.vm.__context), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.value)
	}
}

func (i *pIntern) listToSeriesFloat64() (gandalff.Series, error) {
	switch l := i.value.(type) {
	case pList:
		res := make([]float64, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case float64:
				res[j] = v
			default:
				return nil, fmt.Errorf("expecting list of floats, got %T", e.getValue())
			}
		}
		return gandalff.NewSeriesFloat64(res, nil, false, i.vm.__context), nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.value)
	}
}

func (i *pIntern) listToStringSlice() ([]string, error) {
	switch l := i.value.(type) {
	case pList:
		res := make([]string, len(l))
		for j, e := range l {
			switch v := e.getValue().(type) {
			case string:
				res[j] = v
			case pSymbol:
				res[j] = string(v)
			default:
				return nil, fmt.Errorf("expecting list of strings or symbols, got %T", e.getValue())
			}
		}
		return res, nil

	default:
		return nil, fmt.Errorf("expecting list, got %T", i.value)
	}
}

func (i *pIntern) listToSeriesString() (gandalff.Series, error) {
	switch l := i.value.(type) {
	case pList:
		res := make([]string, len(l))
		for j, e := range l {
			v, err := e.getStringScalar()
			if err != nil {
				return nil, fmt.Errorf("expecting list of strings")
			}
			res[j] = v
		}
		return gandalff.NewSeriesString(res, nil, false, i.vm.__context), nil

	default:
		return nil, fmt.Errorf("expecting list of strings")
	}
}

// isNeg returns true if the expression is a negative number
// used for special cases like orderBy
// func (i *pIntern) isNeg() bool {
// 	if len(i.expr) == 2 {
// 		if op, ok := i.expr[1].(preludiometa.OPCODE); ok && op == preludiometa.OP_UNARY_SUB {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (lhs *pIntern) appendBinaryOperation(op preludiometa.OPCODE, rhs *pIntern) {
// 	lhs.expr = append(lhs.expr, rhs.expr...)
// 	lhs.expr = append(lhs.expr, op)
// }

// func (rhs *pIntern) appendUnaryOperation(op preludiometa.OPCODE) {
// 	rhs.expr = append(rhs.expr, op)
// }

func isOperator(t interface{}) (preludiometa.OPCODE, bool) {
	if v, ok := t.(preludiometa.OPCODE); ok {
		return v, true
	}
	return preludiometa.NO_OP, false
}
