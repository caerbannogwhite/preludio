package preludio

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
)

type __p_intern_tag__ int

const (
	PRELUDIO_INTERNAL_TAG_ERROR       __p_intern_tag__ = 0
	PRELUDIO_INTERNAL_TAG_EXPRESSION  __p_intern_tag__ = 1
	PRELUDIO_INTERNAL_TAG_NAMED_PARAM __p_intern_tag__ = 2
	PRELUDIO_INTERNAL_TAG_ASSIGNMENT  __p_intern_tag__ = 3
	PRELUDIO_INTERNAL_TAG_BEGIN_FRAME __p_intern_tag__ = 4
)

type __p_intern__ struct {
	tag      __p_intern_tag__
	expr     __p_expr_items__
	name     string
	errorMsg string
}

func newPInternError(msg string) *__p_intern__ {
	return &__p_intern__{tag: PRELUDIO_INTERNAL_TAG_ERROR, errorMsg: msg}
}

func newPInternBeginFrame() *__p_intern__ {
	return &__p_intern__{tag: PRELUDIO_INTERNAL_TAG_BEGIN_FRAME}
}

func newPInternTerm(val interface{}) *__p_intern__ {
	return &__p_intern__{tag: PRELUDIO_INTERNAL_TAG_EXPRESSION, expr: *newPExprItems(val)}
}

func (i *__p_intern__) setParamName(name string) {
	i.tag = PRELUDIO_INTERNAL_TAG_NAMED_PARAM
	i.name = name
}

func (i *__p_intern__) setAssignment(name string) {
	i.tag = PRELUDIO_INTERNAL_TAG_ASSIGNMENT
	i.name = name
}

func (i *__p_intern__) ToString() string {
	switch i.tag {
	case PRELUDIO_INTERNAL_TAG_ERROR:
		return i.errorMsg
	case PRELUDIO_INTERNAL_TAG_EXPRESSION:
		return fmt.Sprintf("%v", i.expr.ToString())
	case PRELUDIO_INTERNAL_TAG_NAMED_PARAM:
		return fmt.Sprintf("%s = %v", i.name, i.expr.ToString())
	case PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
		return fmt.Sprintf("%s = %v", i.name, i.expr.ToString())
	case PRELUDIO_INTERNAL_TAG_BEGIN_FRAME:
		return "BEGIN_FRAME"
	}
	return "UNKNOWN"
}

type __p_expr_items__ []interface{}

func (e *__p_expr_items__) ToString() string {
	if len(*e) > 0 {
		switch (*e)[0].(type) {
		case []bool:
			return fmt.Sprintf("%v", (*e)[0].([]bool))
		case []int:
			return fmt.Sprintf("%v", (*e)[0].([]int))
		case []float64:
			return fmt.Sprintf("%v", (*e)[0].([]float64))
		case []string:
			return fmt.Sprintf("%v", (*e)[0].([]string))
		case []__p_symbol__:
			return fmt.Sprintf("%v", (*e)[0].([]__p_symbol__))
		// case []__p_date__:
		// 	return fmt.Sprintf("%v", (*e)[0].([]__p_date__))
		// case []__p_time__:
		// 	return fmt.Sprintf("%v", (*e)[0].([]__p_time__))
		// case []__p_datetime__:
		// 	return fmt.Sprintf("%v", (*e)[0].([]__p_datetime__))
		// case []__p_duration__:
		// 	return fmt.Sprintf("%v", (*e)[0].([]__p_duration__))
		default:
			return fmt.Sprintf("%v", (*e)[0])
		}
	}
	return ""
}

func newPExprItems(t interface{}) *__p_expr_items__ {
	e := make(__p_expr_items__, 1)
	e[0] = t
	return &e
}

func (i *__p_intern__) getValue() interface{} {
	return i.expr[0]
}

func (i *__p_intern__) isBoolScalar() bool {
	if s, ok := i.expr[0].([]bool); ok && len(s) == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isBoolVector() bool {
	if s, ok := i.expr[0].([]bool); ok && len(s) > 1 {
		return true
	}
	return false
}

func (i *__p_intern__) getBoolScalar() (bool, error) {
	if v, ok := i.expr[0].([]bool); ok && len(v) == 1 {
		return v[0], nil
	}
	return false, fmt.Errorf("expecting bool scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getBoolVector() ([]bool, error) {
	if v, ok := i.expr[0].([]bool); ok {
		return v, nil
	}
	return []bool{}, fmt.Errorf("expecting bool vector, got %T", i.expr[0])
}

func (i *__p_intern__) isIntegerScalar() bool {
	if v, ok := i.expr[0].([]int); ok && len(v) == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isIntegerVector() bool {
	if v, ok := i.expr[0].([]int); ok && len(v) > 1 {
		return true
	}
	return false
}

func (i *__p_intern__) getIntegerScalar() (int, error) {
	if v, ok := i.expr[0].([]int); ok && len(v) == 1 {
		return v[0], nil
	}
	return 0, fmt.Errorf("expecting integer scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getIntegerVector() ([]int, error) {
	if v, ok := i.expr[0].([]int); ok {
		return v, nil
	}
	return []int{}, fmt.Errorf("expecting integer vector, got %T", i.expr[0])
}

func (i *__p_intern__) isFloatScalar() bool {
	if v, ok := i.expr[0].([]float64); ok && len(v) == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isFloatVector() bool {
	if v, ok := i.expr[0].([]float64); ok && len(v) > 1 {
		return true
	}
	return false
}

func (i *__p_intern__) getFloatScalar() (float64, error) {
	if v, ok := i.expr[0].([]float64); ok && len(v) == 1 {
		return v[0], nil
	}
	return 0, fmt.Errorf("expecting float scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getFloatVector() ([]float64, error) {
	if v, ok := i.expr[0].([]float64); ok {
		return v, nil
	}
	return []float64{}, fmt.Errorf("expecting float vector, got %T", i.expr[0])
}

func (i *__p_intern__) isStringScalar() bool {
	if v, ok := i.expr[0].([]string); ok && len(v) == 1 {
		return true
	}
	return false
}

func (i *__p_intern__) isStringVector() bool {
	if v, ok := i.expr[0].([]string); ok && len(v) > 1 {
		return true
	}
	return false
}

func (i *__p_intern__) getStringScalar() (string, error) {
	if v, ok := i.expr[0].([]string); ok && len(v) == 1 {
		return v[0], nil
	}
	return "", fmt.Errorf("expecting string scalar, got %T", i.expr[0])
}

func (i *__p_intern__) getStringVector() ([]string, error) {
	if v, ok := i.expr[0].([]string); ok {
		return v, nil
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

func (lhs *__p_intern__) appendOperand(op OPCODE, rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, op)
}

func isOperator(t interface{}) (OPCODE, bool) {
	if v, ok := t.(OPCODE); ok {
		return v, true
	}
	return NO_OP, false
}

func (i *__p_intern__) solve(vm *ByteEater) error {
	tmp := make([]interface{}, 1)

	// TODO: check if this is possible and
	// if it's the case to raise an error
	if len(i.expr) == 0 {
		return nil
	}

	// Check if the expression is a list
	// and recursively solve all the expressions
	// in the list
	if len(i.expr) == 1 {
		switch l := i.expr[0].(type) {
		case __p_symbol__:
			i.expr[0] = vm.SymbolResolution(l)

		case __p_list__:
			for idx := range l {
				if err := l[idx].solve(vm); err != nil {
					return err
				}
			}

		default:
			return nil
		}
	}

	for len(i.expr) > 1 {
		t1 := i.expr[0]
		t2 := i.expr[1]

		var result interface{}

		// UNARY
		if op, ok := isOperator(t2); ok {
			i.expr = i.expr[2:len(i.expr)]

			if s, ok := t1.(__p_symbol__); ok {
				t1 = vm.SymbolResolution(s)
			}

			switch op {
			case OP_UNARY_ADD:
			case OP_UNARY_SUB:
			case OP_UNARY_NOT:
			}
		} else

		// BINARY
		{
			op, _ := isOperator(i.expr[2])
			i.expr = i.expr[3:len(i.expr)]

			// Symbo resolution
			if s, ok := t1.(__p_symbol__); ok {
				t1 = vm.SymbolResolution(s)
			}

			if s, ok := t2.(__p_symbol__); ok {
				t2 = vm.SymbolResolution(s)
			}

			switch op {

			///////////////////////////////////////////////////////////////////
			////////					MULTIPLICATION

			case OP_BINARY_MUL:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) * BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) * BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) * BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]int, len(val2))
								for j, _ := range val2 {
									res[j] = 0
								}
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) * n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]float64, len(val2))
								for j, _ := range val2 {
									res[j] = 0.0
								}
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) * n
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]string, len(val2))
								for j, _ := range val2 {
									res[j] = ""
								}
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, b := range val1 {
								if b {
									res[j] = val2[j]
								} else {
									res[j] = ""
								}
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, b := range val1 {
								if b {
									res[j] = val2[j]
								} else {
									res[j] = ""
								}
							}
						}
						result = res

					default:
						return fmt.Errorf("binary * operator not implemented for %T and %T", val1, val2)
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] * BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n * BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] * BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] * n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] * n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) * f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) * f
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = strings.Repeat(s, int(val1[0]))
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, n := range val1 {
								res[j] = strings.Repeat(val2[0], n)
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = strings.Repeat(s, val1[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary * operator not implemented for %T and %T", val1, val2)
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] * BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f * BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] * BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] * float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f * float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] * float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] * f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] * f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary * operator not implemented for %T and %T", val1, val2)
					}

				case []string:
					switch val2 := t2.(type) {
					case []bool:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, b := range val2 {
								if b {
									res[j] = val1[0]
								} else {
									res[j] = ""
								}
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								if val2[0] {
									res[j] = s
								} else {
									res[j] = ""
								}
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								if val2[j] {
									res[j] = s
								} else {
									res[j] = ""
								}
							}
						}
						result = res
					case []int:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, n := range val2 {
								res[j] = strings.Repeat(val1[0], n)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, n := range val2 {
								res[j] = strings.Repeat(val1[0], n)
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = strings.Repeat(s, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary * operator not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					DIVISION

			case OP_BINARY_DIV:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) / BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) / BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) / BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]int, len(val2))
								for j, _ := range val2 {
									res[j] = 0
								}
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) / n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]float64, len(val2))
								for j, _ := range val2 {
									res[j] = 0.0
								}
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) / n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary / operator not implemented for %T and %T", val1, val2)
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] / BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n / BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] / BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] / n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] / n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) / f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) / f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary / operator not implemented for %T and %T", val1, val2)
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] / BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f / BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] / BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] / float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f / float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] / float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] / f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] / f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary / operator not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					MODULUS

			case OP_BINARY_MOD:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) % BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) % BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) % BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]int, len(val2))
								for j, _ := range val2 {
									res[j] = 0
								}
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) % val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) % n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary %% operator not implemented for %T and %T", val1, val2)
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] % BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n % BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] % BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] % n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n % val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] % n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary %% operator not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					ADDITION

			case OP_BINARY_ADD:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) + BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) + BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) + BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[0]) + n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) + n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[0]) + f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) + n
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%t%s", val1[0], s)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%t%s", b, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%t%s", b, val2[j])
							}
						}
						result = res

					default:
						return fmt.Errorf("binary + operator not implemented for %T and %T", val1, val2)
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] + BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n + BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] + BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] + n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] + n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) + f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) + f
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%d%s", val1[0], s)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%d%s", b, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%d%s", b, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary + operator not implemented for %T and %T", val1, val2)
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] + BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f + BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] + BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] + float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f + float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] + float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] + f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] + f
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, f := range val2 {
								res[j] = fmt.Sprintf("%f%s", val1[0], f)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, f := range val1 {
								res[j] = fmt.Sprintf("%f%s", f, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, f := range val1 {
								res[j] = fmt.Sprintf("%f%s", f, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary + operator not implemented for %T and %T", val1, val2)
					}

				case []string:
					switch val2 := t2.(type) {
					case []bool:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, b := range val2 {
								res[j] = fmt.Sprintf("%s%t", val1[0], b)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%t", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%t", s, val2[j])
							}
						}
						result = res
					case []int:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, n := range val2 {
								res[j] = fmt.Sprintf("%s%d", val1[0], n)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%d", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%d", s, val2[j])
							}
						}
						result = res
					case []float64:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, f := range val2 {
								res[j] = fmt.Sprintf("%s%f", val1[0], f)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%f", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%f", s, val2[j])
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%s%s", val1[0], s)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%s", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%s", s, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary + operator not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					SUBTRACTION

			case OP_BINARY_SUB:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) - BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) - BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) - BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[0]) - n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) - n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[0]) - f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) - n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary - operator not implemented for %T and %T", val1, val2)
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] - BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n - BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] - BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] - n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] - n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) - f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) - f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary - operator not implemented for %T and %T", val1, val2)
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] - BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f - BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] - BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] - float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f - float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] - float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] - f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] - f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary - operator not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					EQUAL
			case OP_BINARY_EQ:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == b
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = b == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == b
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[0]) == n
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) == n
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[0]) == f
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[j]) == f
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%t", val1[0]) == s
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%t", b) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%t", val1[j]) == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary == operator not implemented for %T and %T", val1, val2)
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = n == BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] == n
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = n == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] == n
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) == f
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) == f
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = strconv.Itoa(val1[0]) == s
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = strconv.Itoa(n) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = strconv.Itoa(val1[j]) == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary == operator not implemented for %T and %T", val1, val2)
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, f := range val1 {
								res[j] = f == BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] == float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, f := range val1 {
								res[j] = f == float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] == float64(n)
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] == f
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, f := range val1 {
								res[j] = f == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] == f
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = strconv.FormatFloat(val1[0], 'f', -1, 64) == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary == operator not implemented for %T and %T", val1, val2)
					}

				case []string:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == strconv.FormatBool(b)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == strconv.FormatBool(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == strconv.FormatBool(b)
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] == strconv.Itoa(n)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == strconv.Itoa(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] == strconv.Itoa(n)
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] == strconv.FormatFloat(f, 'f', -1, 64)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == strconv.FormatFloat(val2[0], 'f', -1, 64)
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] == strconv.FormatFloat(f, 'f', -1, 64)
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = val1[0] == s
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = val1[j] == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary == operator not implemented for %T and %T", val1, val2)
					}
				}
			}
		}

		tmp[0] = result
		i.expr = append(tmp, i.expr...)
	}
	return nil
}
