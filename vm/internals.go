package preludiovm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gota/gota/dataframe"
)

type InternalTag int

const (
	PRELUDIO_INTERNAL_TAG_ERROR       InternalTag = 0
	PRELUDIO_INTERNAL_TAG_EXPRESSION  InternalTag = 1
	PRELUDIO_INTERNAL_TAG_NAMED_PARAM InternalTag = 2
	PRELUDIO_INTERNAL_TAG_ASSIGNMENT  InternalTag = 3
	PRELUDIO_INTERNAL_TAG_BEGIN_FRAME InternalTag = 4
)

type PreludioInternal struct {
	Tag      InternalTag
	Expr     PreludioExpr
	Name     string
	ErrorMsg string
}

func NewPreludioInternalError(msg string) *PreludioInternal {
	return &PreludioInternal{Tag: PRELUDIO_INTERNAL_TAG_ERROR, ErrorMsg: msg}
}

func NewPreludioInternalBeginFrame() *PreludioInternal {
	return &PreludioInternal{Tag: PRELUDIO_INTERNAL_TAG_BEGIN_FRAME}
}

func NewPreludioInternalTerm(val interface{}) *PreludioInternal {
	return &PreludioInternal{Tag: PRELUDIO_INTERNAL_TAG_EXPRESSION, Expr: *NewPreludioExpr(val)}
}

func (i *PreludioInternal) SetParamName(name string) {
	i.Tag = PRELUDIO_INTERNAL_TAG_NAMED_PARAM
	i.Name = name
}

func (i *PreludioInternal) SetAssignment(name string) {
	i.Tag = PRELUDIO_INTERNAL_TAG_ASSIGNMENT
	i.Name = name
}

type PreludioExprOp uint8

const (
	BIN_EXPR_MUL PreludioExprOp = 0
	BIN_EXPR_DIV PreludioExprOp = 1
	BIN_EXPR_MOD PreludioExprOp = 2
	BIN_EXPR_ADD PreludioExprOp = 3
	BIN_EXPR_SUB PreludioExprOp = 4
	BIN_EXPR_POW PreludioExprOp = 5

	UN_EXPR_ADD PreludioExprOp = 30
	UN_EXPR_SUB PreludioExprOp = 31
	UN_EXPR_NOT PreludioExprOp = 32

	NO_OP PreludioExprOp = 50
)

type PreludioExpr []interface{}

func NewPreludioExpr(t interface{}) *PreludioExpr {
	e := make(PreludioExpr, 1)
	e[0] = t
	return &e
}

func (i *PreludioInternal) GetValue() interface{} {
	return i.Expr[0]
}

func (i *PreludioInternal) IsBoolScalar() bool {
	if s, ok := i.Expr[0].([]bool); ok && len(s) == 1 {
		return true
	}
	return false
}

func (i *PreludioInternal) GetBoolScalar() (bool, error) {
	if v, ok := i.Expr[0].([]bool); ok && len(v) == 1 {
		return v[0], nil
	}
	return false, errors.New(fmt.Sprintf("expecting bool scalar, got %T", i.Expr[0]))
}

func (i *PreludioInternal) GetBoolVector() ([]bool, error) {
	if v, ok := i.Expr[0].([]bool); ok {
		return v, nil
	}
	return []bool{}, errors.New(fmt.Sprintf("expecting bool vector, got %T", i.Expr[0]))
}

func (i *PreludioInternal) IsIntegerScalar() bool {
	if v, ok := i.Expr[0].([]int); ok && len(v) == 1 {
		return true
	}
	return false
}

func (i *PreludioInternal) GetIntegerScalar() (int, error) {
	if v, ok := i.Expr[0].([]int); ok && len(v) == 1 {
		return v[0], nil
	}
	return 0, errors.New(fmt.Sprintf("expecting integer scalar, got %T", i.Expr[0]))
}

func (i *PreludioInternal) GetIntegerVector() ([]int, error) {
	if v, ok := i.Expr[0].([]int); ok {
		return v, nil
	}
	return []int{}, errors.New(fmt.Sprintf("expecting integer vector, got %T", i.Expr[0]))
}

func (i *PreludioInternal) IsScalarFloat() bool {
	if v, ok := i.Expr[0].([]float64); ok && len(v) == 1 {
		return true
	}
	return false
}

func (i *PreludioInternal) GetFloatScalar() (float64, error) {
	if v, ok := i.Expr[0].([]float64); ok && len(v) == 1 {
		return v[0], nil
	}
	return 0, errors.New(fmt.Sprintf("expecting float scalar, got %T", i.Expr[0]))
}

func (i *PreludioInternal) GetFloatVector() ([]float64, error) {
	if v, ok := i.Expr[0].([]float64); ok {
		return v, nil
	}
	return []float64{}, errors.New(fmt.Sprintf("expecting float vector, got %T", i.Expr[0]))
}

func (i *PreludioInternal) IsStringScalar() bool {
	if v, ok := i.Expr[0].([]string); ok && len(v) == 1 {
		return true
	}
	return false
}

func (i *PreludioInternal) GetStringScalar() (string, error) {
	if v, ok := i.Expr[0].([]string); ok && len(v) == 1 {
		return v[0], nil
	}
	return "", errors.New(fmt.Sprintf("expecting string scalar, got %T", i.Expr[0]))
}

func (i *PreludioInternal) GetStringVector() ([]string, error) {
	if v, ok := i.Expr[0].([]string); ok {
		return v, nil
	}
	return []string{}, errors.New(fmt.Sprintf("expecting string vector, got %T", i.Expr[0]))
}

func (i *PreludioInternal) IsSymbol() bool {
	_, ok := i.Expr[0].(PreludioSymbol)
	return ok
}

func (i *PreludioInternal) GetSymbol() (PreludioSymbol, error) {
	if v, ok := i.Expr[0].(PreludioSymbol); ok {
		return v, nil
	}
	return "", errors.New(fmt.Sprintf("expecting symbol, got %T", i.Expr[0]))
}

func (i *PreludioInternal) IsDataframe() bool {
	_, ok := i.Expr[0].(dataframe.DataFrame)
	return ok
}

func (i *PreludioInternal) GetDataframe() (dataframe.DataFrame, error) {
	switch v := i.Expr[0].(type) {
	case dataframe.DataFrame:
		return v, nil
	default:
		return dataframe.DataFrame{}, errors.New(fmt.Sprintf("expecting dataframe, got %T", v))
	}
}

func (i *PreludioInternal) IsList() bool {
	_, ok := i.Expr[0].(PreludioList)
	return ok
}

func (i *PreludioInternal) GetList() (PreludioList, error) {
	switch v := i.Expr[0].(type) {
	case PreludioList:
		return v, nil
	default:
		return PreludioList{}, errors.New(fmt.Sprintf("expecting list, got %T", v))
	}
}

func (i *PreludioInternal) ListToBoolVector() ([]bool, error) {
	if l, ok := i.Expr[0].(PreludioList); ok {
		res := make([]bool, len(l))
		for j, e := range l {
			v, err := e.GetBoolScalar()
			if err != nil {
				return []bool{}, errors.New(fmt.Sprintf("expecting list of bools"))
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []bool{}, errors.New(fmt.Sprintf("expecting list of bools"))
	}
}

func (l *PreludioList) ListToBoolVector() ([]bool, error) {
	res := make([]bool, len(*l))
	for j, e := range *l {
		v, err := e.GetBoolScalar()
		if err != nil {
			return []bool{}, errors.New(fmt.Sprintf("expecting list of bools"))
		}
		res[j] = v
	}
	return res, nil
}

func (i *PreludioInternal) ListToIntegerVector() ([]int, error) {
	if l, ok := i.Expr[0].(PreludioList); ok {
		res := make([]int, len(l))
		for j, e := range l {
			v, err := e.GetIntegerScalar()
			if err != nil {
				return []int{}, errors.New(fmt.Sprintf("expecting list of integers"))
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []int{}, errors.New(fmt.Sprintf("expecting list of integers"))
	}
}

func (l *PreludioList) ListToIntegerVector() ([]int, error) {
	res := make([]int, len(*l))
	for j, e := range *l {
		v, err := e.GetIntegerScalar()
		if err != nil {
			return []int{}, errors.New(fmt.Sprintf("expecting list of integers"))
		}
		res[j] = v
	}
	return res, nil
}

func (i *PreludioInternal) ListToFloatVector() ([]float64, error) {
	if l, ok := i.Expr[0].(PreludioList); ok {
		res := make([]float64, len(l))
		for j, e := range l {
			v, err := e.GetFloatScalar()
			if err != nil {
				return []float64{}, errors.New(fmt.Sprintf("expecting list of floats"))
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []float64{}, errors.New(fmt.Sprintf("expecting list of floats"))
	}
}

func (l *PreludioList) ListToFloatVector() ([]float64, error) {
	res := make([]float64, len(*l))
	for j, e := range *l {
		v, err := e.GetFloatScalar()
		if err != nil {
			return []float64{}, errors.New(fmt.Sprintf("expecting list of floats"))
		}
		res[j] = v
	}
	return res, nil
}

func (i *PreludioInternal) ListToStringVector() ([]string, error) {
	if l, ok := i.Expr[0].(PreludioList); ok {
		res := make([]string, len(l))
		for j, e := range l {
			v, err := e.GetStringScalar()
			if err != nil {
				return []string{}, errors.New(fmt.Sprintf("expecting list of strings"))
			}
			res[j] = v
		}
		return res, nil
	} else {
		return []string{}, errors.New(fmt.Sprintf("expecting list of strings"))
	}
}

func (l *PreludioList) ListToStringVector() ([]string, error) {
	res := make([]string, len(*l))
	for j, e := range *l {
		v, err := e.GetStringScalar()
		if err != nil {
			return []string{}, errors.New(fmt.Sprintf("expecting list of strings"))
		}
		res[j] = v
	}
	return res, nil
}

func (lhs *PreludioInternal) Mul(rhs *PreludioInternal) {
	lhs.Expr = append(lhs.Expr, rhs.Expr...)
	lhs.Expr = append(lhs.Expr, BIN_EXPR_MUL)
}

func (lhs *PreludioInternal) Div(rhs *PreludioInternal) {
	lhs.Expr = append(lhs.Expr, rhs.Expr...)
	lhs.Expr = append(lhs.Expr, BIN_EXPR_DIV)
}

func (lhs *PreludioInternal) Mod(rhs *PreludioInternal) {
	lhs.Expr = append(lhs.Expr, rhs.Expr...)
	lhs.Expr = append(lhs.Expr, BIN_EXPR_MOD)
}

func (lhs *PreludioInternal) Add(rhs *PreludioInternal) {
	lhs.Expr = append(lhs.Expr, rhs.Expr...)
	lhs.Expr = append(lhs.Expr, BIN_EXPR_ADD)
}

func (lhs *PreludioInternal) Sub(rhs *PreludioInternal) {
	lhs.Expr = append(lhs.Expr, rhs.Expr...)
	lhs.Expr = append(lhs.Expr, BIN_EXPR_SUB)
}

func (lhs *PreludioInternal) Pow(rhs *PreludioInternal) {
	lhs.Expr = append(lhs.Expr, rhs.Expr...)
	lhs.Expr = append(lhs.Expr, BIN_EXPR_POW)
}

func IsOperator(t interface{}) (PreludioExprOp, bool) {
	if v, ok := t.(PreludioExprOp); ok {
		return v, true
	}
	return NO_OP, false
}

func (i *PreludioInternal) Solve(vm *PreludioVM) error {
	tmp := make([]interface{}, 1)

	// TODO: check if this is possible and
	// if it's the case to raise an error
	if len(i.Expr) == 0 {
		return nil
	}

	// Check if the expression is a list
	// and recursively solve all the expressions
	// in the list
	if len(i.Expr) == 1 {
		switch l := i.Expr[0].(type) {
		case PreludioSymbol:
			i.Expr[0] = vm.SymbolResolution(l)

		case PreludioList:
			for _, t := range l {
				if err := t.Solve(vm); err != nil {
					return err
				}
			}

		default:
			return nil
		}
	}

	for len(i.Expr) > 1 {
		t1 := i.Expr[0]
		t2 := i.Expr[1]

		var result interface{}

		// UNARY
		if op, ok := IsOperator(t2); ok {
			i.Expr = i.Expr[2:len(i.Expr)]

			if s, ok := t1.(PreludioSymbol); ok {
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
			op, _ := IsOperator(i.Expr[2])
			i.Expr = i.Expr[3:len(i.Expr)]

			// Symbo resolution
			if s, ok := t1.(PreludioSymbol); ok {
				t1 = vm.SymbolResolution(s)
			}

			if s, ok := t2.(PreludioSymbol); ok {
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
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
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
					}
				}
			}
		}

		tmp[0] = result
		i.Expr = append(tmp, i.Expr...)
	}
	return nil
}