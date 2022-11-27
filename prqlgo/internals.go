package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type PrqlInternalTag int

const (
	PRQL_INTERNAL_TAG_ERROR        PrqlInternalTag = 0
	PRQL_INTERNAL_TAG_EXPRESSION   PrqlInternalTag = 1
	PRQL_INTERNAL_TAG_PARAM_NAME   PrqlInternalTag = 2
	PRQL_INTERNAL_TAG_ASSING_IDENT PrqlInternalTag = 3
)

type PrqlInternal struct {
	Tag      PrqlInternalTag
	Expr     *PrqlExpr
	Name     string
	ErrorMsg string
}

func NewPrqlInternalError(msg string) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_ERROR, ErrorMsg: msg}
}

func NewPrqlInternalTerm(val interface{}) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_EXPRESSION, Expr: NewPrqlExpr(val)}
}

func NewPrqlInternalParamName(name string) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_PARAM_NAME, Name: name}
}

func NewPrqlInternalAssignIdent(ident string) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_ASSING_IDENT, Name: ident}
}

// func (op1 *PrqlInternal) ArithBinaryMul(op2 *PrqlInternal) {
// 	switch val1 := op1.Value.(type) {
// 	case bool:
// 		switch val2 := op2.Value.(type) {
// 		case bool:
// 			op1.Value = val1 && val2
// 		case int:
// 			if val1 {
// 				op1.Value = val2
// 			} else {
// 				op1.Value = int(0)
// 			}
// 		case float64:
// 			if val1 {
// 				op1.Value = val2
// 			} else {
// 				op1.Value = float64(0)
// 			}
// 		case string:
// 			if val1 {
// 				op1.Value = val2
// 			} else {
// 				op1.Value = ""
// 			}

// 		case series.Series:

// 		// case dataframe.DataFrame:

// 		default:
// 			op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 			op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 		}

// 	case int:
// 		switch val2 := op2.Value.(type) {
// 		case bool:
// 			if !val2 {
// 				op1.Value = int(0)
// 			}
// 		case int:
// 			op1.Value = val1 * val2
// 		case float64:
// 			op1.Value = float64(val1) * val2
// 		case string:
// 			op1.Value = strings.Repeat(val2, val1)

// 		case series.Series:
// 		// case dataframe.DataFrame:
// 		default:
// 			op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 			op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 		}

// 	case float64:
// 		switch val2 := op2.Value.(type) {
// 		case bool:
// 			if !val2 {
// 				op1.Value = float64(0)
// 			}
// 		case int:
// 			op1.Value = val1 * float64(val2)
// 		case float64:
// 			op1.Value = val1 * val2
// 		// case string:
// 		case series.Series:
// 		// case dataframe.DataFrame:
// 		default:
// 			op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 			op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 		}

// 	case string:
// 		switch val2 := op2.Value.(type) {
// 		case bool:
// 			if !val2 {
// 				op1.Value = ""
// 			}
// 		case int:
// 			op1.Value = strings.Repeat(val1, val2)
// 		// case float64:
// 		// case string:
// 		case series.Series:
// 		// case dataframe.DataFrame:
// 		default:
// 			op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 			op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 		}

// 	// SERIES
// 	case series.Series:
// 		switch val2 := op2.Value.(type) {
// 		case bool:
// 		case int:
// 		case float64:
// 		case string:
// 		case series.Series:
// 		// case dataframe.DataFrame:
// 		default:
// 			op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 			op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 		}

// 	// DATAFRAME
// 	// case dataframe.DataFrame:
// 	// 	switch val2 := op2.Value.(type) {
// 	// 	case bool:
// 	// 	case int:
// 	// 	case float64:
// 	// 	case string:
// 	// 	case series.Series:
// 	// 	case dataframe.DataFrame:
// 	// 	default:
// 	// 		op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 	// 		op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 	// 	}

// 	default:
// 		op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 		op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T.", val1)
// 	}
// }

// func (op1 *PrqlInternal) ArithBinaryAdd(op2 *PrqlInternal) {
// 	switch val1 := op1.Value.(type) {
// 	case bool:
// 		switch val2 := op2.Value.(type) {
// 		case bool:
// 			op1.Value = val1 || val2
// 		case int:
// 			if val1 {
// 				op1.Value = val2 + 1
// 			}
// 		case float64:
// 			if val1 {
// 				op1.Value = val2 + 1.0
// 			}
// 		case string:
// 			op1.Value = fmt.Sprintf("%v%v", val1, val2)

// 		case series.Series:

// 		// case dataframe.DataFrame:

// 		default:
// 			op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 			op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 		}

// 	// case int:
// 	// 	switch val2 := op2.Value.(type) {
// 	// 	case bool:
// 	// 		if !val2 {
// 	// 			op1.Value = int(0)
// 	// 		}
// 	// 	case int:
// 	// 		op1.Value = val1 * val2
// 	// 	case float64:
// 	// 		op1.Value = float64(val1) * val2
// 	// 	case string:
// 	// 		op1.Value = strings.Repeat(val2, val1)

// 	// 	case series.Series:
// 	// 	// case dataframe.DataFrame:
// 	// 	default:
// 	// 		op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 	// 		op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 	// 	}

// 	// case float64:
// 	// 	switch val2 := op2.Value.(type) {
// 	// 	case bool:
// 	// 		if !val2 {
// 	// 			op1.Value = float64(0)
// 	// 		}
// 	// 	case int:
// 	// 		op1.Value = val1 * float64(val2)
// 	// 	case float64:
// 	// 		op1.Value = val1 * val2
// 	// 	// case string:
// 	// 	case series.Series:
// 	// 	// case dataframe.DataFrame:
// 	// 	default:
// 	// 		op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 	// 		op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 	// 	}

// 	// case string:
// 	// 	switch val2 := op2.Value.(type) {
// 	// 	case bool:
// 	// 		if !val2 {
// 	// 			op1.Value = ""
// 	// 		}
// 	// 	case int:
// 	// 		op1.Value = strings.Repeat(val1, val2)
// 	// 	// case float64:
// 	// 	// case string:
// 	// 	case series.Series:
// 	// 	// case dataframe.DataFrame:
// 	// 	default:
// 	// 		op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 	// 		op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 	// 	}

// 	// SERIES
// 	case series.Series:
// 		switch val2 := op2.Value.(type) {
// 		case bool:
// 		case int:
// 		case float64:
// 		case string:
// 		case series.Series:
// 		// case dataframe.DataFrame:
// 		default:
// 			op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 			op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 		}

// 	// DATAFRAME
// 	// case dataframe.DataFrame:
// 	// 	switch val2 := op2.Value.(type) {
// 	// 	case bool:
// 	// 	case int:
// 	// 	case float64:
// 	// 	case string:
// 	// 	case series.Series:
// 	// 	case dataframe.DataFrame:
// 	// 	default:
// 	// 		op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 	// 		op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T and %T.", val1, val2)
// 	// 	}

// 	default:
// 		op1.Tag = PRQL_INTERNAL_TAG_ERROR
// 		op1.ErrorMsg = fmt.Sprintf("cannot apply binary multiplication to %T.", val1)
// 	}
// }

type PrqlExprOp uint8

const (
	BIN_EXPR_MUL PrqlExprOp = 0
	BIN_EXPR_DIV PrqlExprOp = 1
	BIN_EXPR_MOD PrqlExprOp = 2
	BIN_EXPR_ADD PrqlExprOp = 3
	BIN_EXPR_SUB PrqlExprOp = 4
	BIN_EXPR_POW PrqlExprOp = 5

	UN_EXPR_ADD PrqlExprOp = 30
	UN_EXPR_SUB PrqlExprOp = 31
	UN_EXPR_NOT PrqlExprOp = 32

	NO_OP PrqlExprOp = 50
)

type PrqlExpr struct {
	stack []interface{}
}

func NewPrqlExpr(t interface{}) *PrqlExpr {
	e := PrqlExpr{}
	e.stack = append(e.stack, t)
	return &e
}

func (e *PrqlExpr) GetValue() interface{} {
	return e.stack[0]
}

func (e *PrqlExpr) GetValueBool() (bool, error) {
	switch v := e.stack[0].(type) {
	case bool:
		return v, nil
	default:
		return false, errors.New(fmt.Sprintf("expecting bool, got %T", v))
	}
}

func (e *PrqlExpr) GetValueInteger() (int64, error) {
	switch v := e.stack[0].(type) {
	case int64:
		return v, nil
	default:
		return 0, errors.New(fmt.Sprintf("expecting integer, got %T", v))
	}
}

func (e *PrqlExpr) GetValueFloat() (float64, error) {
	switch v := e.stack[0].(type) {
	case float64:
		return v, nil
	default:
		return 0, errors.New(fmt.Sprintf("expecting float, got %T", v))
	}
}

func (e *PrqlExpr) GetValueString() (string, error) {
	switch v := e.stack[0].(type) {
	case string:
		return v, nil
	default:
		return "", errors.New(fmt.Sprintf("expecting string, got %T", v))
	}
}

func (e *PrqlExpr) GetValueSeries() (series.Series, error) {
	switch v := e.stack[0].(type) {
	case series.Series:
		return v, nil
	default:
		return series.Series{}, errors.New(fmt.Sprintf("expecting series, got %T", v))
	}
}

func (e *PrqlExpr) GetValueDataframe() (dataframe.DataFrame, error) {
	switch v := e.stack[0].(type) {
	case dataframe.DataFrame:
		return v, nil
	default:
		return dataframe.DataFrame{}, errors.New(fmt.Sprintf("expecting dataframe, got %T", v))
	}
}

func (l *PrqlExpr) Mul(r *PrqlExpr) {
	l.stack = append(l.stack, r.stack...)
	l.stack = append(l.stack, BIN_EXPR_MUL)
}

func (l *PrqlExpr) Div(r *PrqlExpr) {
	l.stack = append(l.stack, r.stack...)
	l.stack = append(l.stack, BIN_EXPR_DIV)
}

func (l *PrqlExpr) Mod(r *PrqlExpr) {
	l.stack = append(l.stack, r.stack...)
	l.stack = append(l.stack, BIN_EXPR_MOD)
}

func (l *PrqlExpr) Add(r *PrqlExpr) {
	l.stack = append(l.stack, r.stack...)
	l.stack = append(l.stack, BIN_EXPR_ADD)
}

func (l *PrqlExpr) Sub(r *PrqlExpr) {
	l.stack = append(l.stack, r.stack...)
	l.stack = append(l.stack, BIN_EXPR_SUB)
}

func (l *PrqlExpr) Pow(r *PrqlExpr) {
	l.stack = append(l.stack, r.stack...)
	l.stack = append(l.stack, BIN_EXPR_POW)
}

func IsOperator(t interface{}) (PrqlExprOp, bool) {
	switch v := t.(type) {
	case PrqlExprOp:
		return v, true
	}
	return NO_OP, false
}

func BoolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func BoolToFloat64(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

func (e *PrqlExpr) Solve() error {

	for len(e.stack) > 1 {
		t1 := e.stack[0]
		t2 := e.stack[1]

		var result interface{}

		// UNARY
		if op, ok := IsOperator(t2); ok {
			e.stack = e.stack[2:len(e.stack)]

			switch op {
			case UN_EXPR_ADD:
			case UN_EXPR_SUB:
			case UN_EXPR_NOT:
			}
		} else

		// BINARY
		{
			op, _ := IsOperator(e.stack[2])
			e.stack = e.stack[3:len(e.stack)]

			switch op {
			case BIN_EXPR_MUL:
				switch val1 := t1.(type) {
				case bool:
					switch val2 := t2.(type) {
					case bool:
						result = BoolToInt64(val1) * BoolToInt64(val2)
					case int64:
						result = BoolToInt64(val1) * val2
					case float64:
						result = BoolToFloat64(val1) * val2
					case string:
						if val1 {
							result = val2
						} else {
							result = ""
						}
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
					}
				case int64:
					switch val2 := t2.(type) {
					case bool:
						result = val1 * BoolToInt64(val2)
					case int64:
						result = val1 * val2
					case float64:
						result = float64(val1) * val2
					case string:
						result = strings.Repeat(val2, int(val1))
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
					}
				case float64:
					switch val2 := t2.(type) {
					case bool:
						result = val1 * BoolToFloat64(val2)
					case int64:
						result = val1 * float64(val2)
					case float64:
						result = val1 * val2
					case string:
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
					}
				case string:
					switch val2 := t2.(type) {
					case bool:
						if val2 {
							result = val1
						} else {
							result = ""
						}
					case int64:
						result = strings.Repeat(val1, int(val2))
					case float64:
					case string:
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Multiplication not implemented for %T and %T", val1, val2))
					}
				case series.Series:
					// switch val2 := t2.(type) {
					// case bool:
					// case int64:
					// case float64:
					// case string:
					// case series.Series:
					// case dataframe.DataFrame:
					// default:
					// }
				case dataframe.DataFrame:
					// switch val2 := t2.(type) {
					// case bool:
					// case int64:
					// case float64:
					// case string:
					// case series.Series:
					// case dataframe.DataFrame:
					// default:
					// }
				default:
				}
			case BIN_EXPR_DIV:
			case BIN_EXPR_MOD:
			case BIN_EXPR_ADD:
				switch val1 := t1.(type) {
				case bool:
					switch val2 := t2.(type) {
					case bool:
						result = BoolToInt64(val1) + BoolToInt64(val2)
					case int64:
						result = BoolToInt64(val1) + val2
					case float64:
						result = BoolToFloat64(val1) + val2
					case string:
						result = fmt.Sprintf("%v%s", val1, val2)
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Addition not implemented for %T and %T", val1, val2))
					}
				case int64:
					switch val2 := t2.(type) {
					case bool:
						result = val1 + BoolToInt64(val2)
					case int64:
						result = val1 + val2
					case float64:
						result = float64(val1) + val2
					case string:
						result = fmt.Sprintf("%v%s", val1, val2)
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Addition not implemented for %T and %T", val1, val2))
					}
				case float64:
					switch val2 := t2.(type) {
					case bool:
						result = val1 + BoolToFloat64(val2)
					case int64:
						result = val1 + float64(val2)
					case float64:
						result = val1 + val2
					case string:
						result = fmt.Sprintf("%v%s", val1, val2)
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Addition not implemented for %T and %T", val1, val2))
					}
				case string:
					switch val2 := t2.(type) {
					case bool:
					case int64:
					case float64:
					case string:
						result = val1 + val2
					case series.Series:
					case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Addition not implemented for %T and %T", val1, val2))
					}
				case series.Series:
					// switch val2 := t2.(type) {
					// case bool:
					// case int64:
					// case float64:
					// case string:
					// case series.Series:
					// case dataframe.DataFrame:
					// default:
					// }
				case dataframe.DataFrame:
					// switch val2 := t2.(type) {
					// case bool:
					// case int64:
					// case float64:
					// case string:
					// case series.Series:
					// case dataframe.DataFrame:
					// default:
					// }
				default:
				}
			case BIN_EXPR_SUB:
			}
		}

		e.stack = append(e.stack, result)
	}
	return nil
}
