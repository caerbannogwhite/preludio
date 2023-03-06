package preludio

import (
	"fmt"
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
	Tag      __p_intern_tag__
	expr     __p_expr_items__
	Name     string
	ErrorMsg string
}

func newPInternError(msg string) *__p_intern__ {
	return &__p_intern__{Tag: PRELUDIO_INTERNAL_TAG_ERROR, ErrorMsg: msg}
}

func newPInternBeginFrame() *__p_intern__ {
	return &__p_intern__{Tag: PRELUDIO_INTERNAL_TAG_BEGIN_FRAME}
}

func newPInternTerm(val interface{}) *__p_intern__ {
	return &__p_intern__{Tag: PRELUDIO_INTERNAL_TAG_EXPRESSION, expr: *NewPExprItems(val)}
}

func (i *__p_intern__) setParamName(name string) {
	i.Tag = PRELUDIO_INTERNAL_TAG_NAMED_PARAM
	i.Name = name
}

func (i *__p_intern__) setAssignment(name string) {
	i.Tag = PRELUDIO_INTERNAL_TAG_ASSIGNMENT
	i.Name = name
}

type __p_expr_items_op__ uint8

const (
	BIN_EXPR_MUL __p_expr_items_op__ = 0
	BIN_EXPR_DIV __p_expr_items_op__ = 1
	BIN_EXPR_MOD __p_expr_items_op__ = 2
	BIN_EXPR_ADD __p_expr_items_op__ = 3
	BIN_EXPR_SUB __p_expr_items_op__ = 4
	BIN_EXPR_POW __p_expr_items_op__ = 5

	UN_EXPR_ADD __p_expr_items_op__ = 30
	UN_EXPR_SUB __p_expr_items_op__ = 31
	UN_EXPR_NOT __p_expr_items_op__ = 32

	NO_OP __p_expr_items_op__ = 50
)

type __p_expr_items__ []interface{}

func NewPExprItems(t interface{}) *__p_expr_items__ {
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

func (i *__p_intern__) isScalarFloat() bool {
	if v, ok := i.expr[0].([]float64); ok && len(v) == 1 {
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

func (i *__p_intern__) ListToBoolVector() ([]bool, error) {
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

func (l *__p_list__) ListToBoolVector() ([]bool, error) {
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

func (i *__p_intern__) ListToIntegerVector() ([]int, error) {
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

func (l *__p_list__) ListToIntegerVector() ([]int, error) {
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

func (i *__p_intern__) ListToFloatVector() ([]float64, error) {
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

func (l *__p_list__) ListToFloatVector() ([]float64, error) {
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

func (i *__p_intern__) ListToStringVector() ([]string, error) {
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

func (l *__p_list__) ListToStringVector() ([]string, error) {
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

func (lhs *__p_intern__) Mul(rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, BIN_EXPR_MUL)
}

func (lhs *__p_intern__) Div(rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, BIN_EXPR_DIV)
}

func (lhs *__p_intern__) Mod(rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, BIN_EXPR_MOD)
}

func (lhs *__p_intern__) Add(rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, BIN_EXPR_ADD)
}

func (lhs *__p_intern__) Sub(rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, BIN_EXPR_SUB)
}

func (lhs *__p_intern__) Pow(rhs *__p_intern__) {
	lhs.expr = append(lhs.expr, rhs.expr...)
	lhs.expr = append(lhs.expr, BIN_EXPR_POW)
}

func isOperator(t interface{}) (__p_expr_items_op__, bool) {
	if v, ok := t.(__p_expr_items_op__); ok {
		return v, true
	}
	return NO_OP, false
}

func (i *__p_intern__) Solve(vm *ByteEater) error {
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
				if err := l[idx].Solve(vm); err != nil {
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
			case UN_EXPR_ADD:
			case UN_EXPR_SUB:
			case UN_EXPR_NOT:
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

			case BIN_EXPR_MUL:
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					DIVISION

			case BIN_EXPR_DIV:
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					MODULUS

			case BIN_EXPR_MOD:
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					ADDITION

			case BIN_EXPR_ADD:
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
					}
				}

			///////////////////////////////////////////////////////////////////
			////////					SUBTRACTION

			case BIN_EXPR_SUB:
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
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
						return fmt.Errorf("Binary Multiplication not implemented for %T and %T", val1, val2)
					}
				}
			}
		}

		tmp[0] = result
		i.expr = append(tmp, i.expr...)
	}
	return nil
}
