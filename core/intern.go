package preludiocore

import (
	"fmt"
	"typesys"

	"gandalff"

	"github.com/go-gota/gota/dataframe"
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
	expr []interface{}
	name string
}

func newPInternBeginFrame() *__p_intern__ {
	return &__p_intern__{tag: PRELUDIO_INTERNAL_TAG_BEGIN_FRAME}
}

func newPInternTerm(val interface{}) *__p_intern__ {
	e := make([]interface{}, 1)
	e[0] = val
	return &__p_intern__{tag: PRELUDIO_INTERNAL_TAG_EXPRESSION, expr: e}
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
		case []bool:
			*res = append(*res, NewColumnarBool(i.name, fullOutput, outputSnippetLength, v))
		case []int:
			*res = append(*res, NewColumnarInt(i.name, fullOutput, outputSnippetLength, v))
		case []float64:
			*res = append(*res, NewColumnarFloat(i.name, fullOutput, outputSnippetLength, v))
		case []string:
			*res = append(*res, NewColumnarString(i.name, fullOutput, outputSnippetLength, v))
		case dataframe.DataFrame:
			df, err := DataFrameToColumnar(fullOutput, outputSnippetLength, &v)
			if err != nil {
				return err
			}
			*res = append(*res, df...)
		}

	case PRELUDIO_INTERNAL_TAG_NAMED_PARAM:
		switch v := i.expr[0].(type) {
		case []bool:
			*res = append(*res, NewColumnarBool(i.name, fullOutput, outputSnippetLength, v))
		case []int:
			*res = append(*res, NewColumnarInt(i.name, fullOutput, outputSnippetLength, v))
		case []float64:
			*res = append(*res, NewColumnarFloat(i.name, fullOutput, outputSnippetLength, v))
		case []string:
			*res = append(*res, NewColumnarString(i.name, fullOutput, outputSnippetLength, v))
		case dataframe.DataFrame:
			df, err := DataFrameToColumnar(fullOutput, outputSnippetLength, &v)
			if err != nil {
				return err
			}
			*res = append(*res, df...)
		}

	case PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
		switch v := i.expr[0].(type) {
		case []bool:
			*res = append(*res, NewColumnarBool(i.name, fullOutput, outputSnippetLength, v))
		case []int:
			*res = append(*res, NewColumnarInt(i.name, fullOutput, outputSnippetLength, v))
		case []float64:
			*res = append(*res, NewColumnarFloat(i.name, fullOutput, outputSnippetLength, v))
		case []string:
			*res = append(*res, NewColumnarString(i.name, fullOutput, outputSnippetLength, v))
		case dataframe.DataFrame:
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
	_, ok := i.expr[0].(dataframe.DataFrame)
	return ok
}

func (i *__p_intern__) getDataframe() (dataframe.DataFrame, error) {
	switch v := i.expr[0].(type) {
	case dataframe.DataFrame:
		return v, nil
	default:
		return dataframe.DataFrame{}, fmt.Errorf("expecting dataframe, got %T", v)
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
		return __p_list__{}, fmt.Errorf("expecting list, got %T", v)
	}
}

func (i *__p_intern__) listToBoolVector() ([]bool, error) {
	if l, ok := i.expr[0].(__p_list__); ok {
		res := make([]bool, len(l))
		for j, e := range l {
			v, err := e.getBoolScalar()
			if err != nil {
				return []bool{}, fmt.Errorf("expecting list of bools")
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []bool{}, fmt.Errorf("expecting list of bools")
	}
}

func (l *__p_list__) listToBoolVector() ([]bool, error) {
	res := make([]bool, len(*l))
	for j, e := range *l {
		v, err := e.getBoolScalar()
		if err != nil {
			return []bool{}, fmt.Errorf("expecting list of bools")
		}
		res[j] = v
	}
	return res, nil
}

func (i *__p_intern__) listToIntegerVector() ([]int, error) {
	if l, ok := i.expr[0].(__p_list__); ok {
		res := make([]int, len(l))
		for j, e := range l {
			v, err := e.getIntegerScalar()
			if err != nil {
				return []int{}, fmt.Errorf("expecting list of integers")
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []int{}, fmt.Errorf("expecting list of integers")
	}
}

func (l *__p_list__) listToIntegerVector() ([]int, error) {
	res := make([]int, len(*l))
	for j, e := range *l {
		v, err := e.getIntegerScalar()
		if err != nil {
			return []int{}, fmt.Errorf("expecting list of integers")
		}
		res[j] = v
	}
	return res, nil
}

func (i *__p_intern__) listToFloatVector() ([]float64, error) {
	if l, ok := i.expr[0].(__p_list__); ok {
		res := make([]float64, len(l))
		for j, e := range l {
			v, err := e.getFloatScalar()
			if err != nil {
				return []float64{}, fmt.Errorf("expecting list of floats")
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []float64{}, fmt.Errorf("expecting list of floats")
	}
}

func (l *__p_list__) listToFloatVector() ([]float64, error) {
	res := make([]float64, len(*l))
	for j, e := range *l {
		v, err := e.getFloatScalar()
		if err != nil {
			return []float64{}, fmt.Errorf("expecting list of floats")
		}
		res[j] = v
	}
	return res, nil
}

func (i *__p_intern__) listToStringVector() ([]string, error) {
	if l, ok := i.expr[0].(__p_list__); ok {
		res := make([]string, len(l))
		for j, e := range l {
			v, err := e.getStringScalar()
			if err != nil {
				return []string{}, fmt.Errorf("expecting list of strings")
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []string{}, fmt.Errorf("expecting list of strings")
	}
}

func (l *__p_list__) listToStringVector() ([]string, error) {
	res := make([]string, len(*l))
	for j, e := range *l {
		v, err := e.getStringScalar()
		if err != nil {
			return []string{}, fmt.Errorf("expecting list of strings")
		}
		res[j] = v
	}
	return res, nil
}

func (lhs *__p_intern__) appendOperand(op typesys.OPCODE, rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, op)
}

func isOperator(t interface{}) (typesys.OPCODE, bool) {
	if v, ok := t.(typesys.OPCODE); ok {
		return v, true
	}
	return typesys.NO_OP, false
}
