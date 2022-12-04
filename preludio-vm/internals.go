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
	PRQL_INTERNAL_TAG_ERROR       PrqlInternalTag = 0
	PRQL_INTERNAL_TAG_EXPRESSION  PrqlInternalTag = 1
	PRQL_INTERNAL_TAG_NAMED_PARAM PrqlInternalTag = 2
	PRQL_INTERNAL_TAG_ASSIGNMENT  PrqlInternalTag = 3
	PRQL_INTERNAL_TAG_START_FUNC  PrqlInternalTag = 4
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

func NewPrqlInternalStartFunc() *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_START_FUNC}
}

func NewPrqlInternalTerm(val interface{}) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_EXPRESSION, Expr: NewPrqlExpr(val)}
}

func (i *PrqlInternal) SetParamName(name string) {
	i.Tag = PRQL_INTERNAL_TAG_NAMED_PARAM
	i.Name = name
}

func (i *PrqlInternal) SetAssignment(name string) {
	i.Tag = PRQL_INTERNAL_TAG_ASSIGNMENT
	i.Name = name
}

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

type PrqlExpr []interface{}

func NewPrqlExpr(t interface{}) *PrqlExpr {
	e := make(PrqlExpr, 1)
	e[0] = t
	return &e
}

func (i *PrqlInternal) GetValue() interface{} {
	return (*i.Expr)[0]
}

func (i *PrqlInternal) GetValueBool() (bool, error) {
	switch v := (*i.Expr)[0].(type) {
	case bool:
		return v, nil
	default:
		return false, errors.New(fmt.Sprintf("expecting bool, got %T", v))
	}
}

func (i *PrqlInternal) GetValueInteger() (int64, error) {
	switch v := (*i.Expr)[0].(type) {
	case int64:
		return v, nil
	default:
		return 0, errors.New(fmt.Sprintf("expecting integer, got %T", v))
	}
}

func (i *PrqlInternal) GetValueFloat() (float64, error) {
	switch v := (*i.Expr)[0].(type) {
	case float64:
		return v, nil
	default:
		return 0, errors.New(fmt.Sprintf("expecting float, got %T", v))
	}
}

func (i *PrqlInternal) GetValueString() (string, error) {
	switch v := (*i.Expr)[0].(type) {
	case string:
		return v, nil
	default:
		return "", errors.New(fmt.Sprintf("expecting string, got %T", v))
	}
}

func (i *PrqlInternal) GetValueSeries() (series.Series, error) {
	switch v := (*i.Expr)[0].(type) {
	case series.Series:
		return v, nil
	default:
		return series.Series{}, errors.New(fmt.Sprintf("expecting series, got %T", v))
	}
}

func (i *PrqlInternal) GetValueDataframe() (dataframe.DataFrame, error) {
	switch v := (*i.Expr)[0].(type) {
	case dataframe.DataFrame:
		return v, nil
	default:
		return dataframe.DataFrame{}, errors.New(fmt.Sprintf("expecting dataframe, got %T", v))
	}
}

func (l *PrqlInternal) Mul(r *PrqlInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_MUL)
}

func (l *PrqlInternal) Div(r *PrqlInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_DIV)
}

func (l *PrqlInternal) Mod(r *PrqlInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_MOD)
}

func (l *PrqlInternal) Add(r *PrqlInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_ADD)
}

func (l *PrqlInternal) Sub(r *PrqlInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_SUB)
}

func (l *PrqlInternal) Pow(r *PrqlInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_POW)
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

func (e *PrqlInternal) Solve() error {
	tmp := make([]interface{}, 1)

	for len((*e)) > 1 {
		t1 := (*e)[0]
		t2 := (*e)[1]

		var result interface{}

		// UNARY
		if op, ok := IsOperator(t2); ok {
			(*e) = (*e)[2:len((*e))]

			switch op {
			case UN_EXPR_ADD:
			case UN_EXPR_SUB:
			case UN_EXPR_NOT:
			}
		} else

		// BINARY
		{
			op, _ := IsOperator((*e)[2])
			(*e) = (*e)[3:len((*e))]

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
					// case dataframe.DataFrame:
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
					// case dataframe.DataFrame:
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
					// case dataframe.DataFrame:
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
					// case dataframe.DataFrame:
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
				// case dataframe.DataFrame:
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
					// case dataframe.DataFrame:
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
					// case dataframe.DataFrame:
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
					// case dataframe.DataFrame:
					default:
						return errors.New(fmt.Sprintf("Binary Addition not implemented for %T and %T", val1, val2))
					}
				case string:
					switch val2 := t2.(type) {
					case bool:
						result = fmt.Sprintf("%s%v", val1, val2)
					case int64:
						result = fmt.Sprintf("%s%v", val1, val2)
					case float64:
						result = fmt.Sprintf("%s%v", val1, val2)
					case string:
						result = val1 + val2
					case series.Series:
					// case dataframe.DataFrame:
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
				// case dataframe.DataFrame:
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

		tmp[0] = result
		(*e) = append(tmp, (*e)...)
	}
	return nil
}
