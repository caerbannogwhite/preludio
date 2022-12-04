package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type InternalTag int

const (
	PRELUDIO_INTERNAL_TAG_ERROR       InternalTag = 0
	PRELUDIO_INTERNAL_TAG_EXPRESSION  InternalTag = 1
	PRELUDIO_INTERNAL_TAG_NAMED_PARAM InternalTag = 2
	PRELUDIO_INTERNAL_TAG_ASSIGNMENT  InternalTag = 3
	PRELUDIO_INTERNAL_TAG_START_FUNC  InternalTag = 4
)

type PreludioInternal struct {
	Tag      InternalTag
	Expr     *PreludioExpr
	Name     string
	ErrorMsg string
}

func NewPreludioInternalError(msg string) *PreludioInternal {
	return &PreludioInternal{Tag: PRELUDIO_INTERNAL_TAG_ERROR, ErrorMsg: msg}
}

func NewPreludioInternalStartFunc() *PreludioInternal {
	return &PreludioInternal{Tag: PRELUDIO_INTERNAL_TAG_START_FUNC}
}

func NewPreludioInternalTerm(val interface{}) *PreludioInternal {
	return &PreludioInternal{Tag: PRELUDIO_INTERNAL_TAG_EXPRESSION, Expr: NewPreludioExpr(val)}
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
	return (*i.Expr)[0]
}

func (i *PreludioInternal) GetValueBool() (bool, error) {
	switch v := (*i.Expr)[0].(type) {
	case bool:
		return v, nil
	default:
		return false, errors.New(fmt.Sprintf("expecting bool, got %T", v))
	}
}

func (i *PreludioInternal) GetValueInteger() (int64, error) {
	switch v := (*i.Expr)[0].(type) {
	case int64:
		return v, nil
	default:
		return 0, errors.New(fmt.Sprintf("expecting integer, got %T", v))
	}
}

func (i *PreludioInternal) GetValueFloat() (float64, error) {
	switch v := (*i.Expr)[0].(type) {
	case float64:
		return v, nil
	default:
		return 0, errors.New(fmt.Sprintf("expecting float, got %T", v))
	}
}

func (i *PreludioInternal) GetValueString() (string, error) {
	switch v := (*i.Expr)[0].(type) {
	case string:
		return v, nil
	default:
		return "", errors.New(fmt.Sprintf("expecting string, got %T", v))
	}
}

func (i *PreludioInternal) GetValueSymbol() (PreludioSymbol, error) {
	switch v := (*i.Expr)[0].(type) {
	case PreludioSymbol:
		return v, nil
	default:
		return "", errors.New(fmt.Sprintf("expecting symbol, got %T", v))
	}
}

func (i *PreludioInternal) GetValueSeries() (series.Series, error) {
	switch v := (*i.Expr)[0].(type) {
	case series.Series:
		return v, nil
	default:
		return series.Series{}, errors.New(fmt.Sprintf("expecting series, got %T", v))
	}
}

func (i *PreludioInternal) GetValueDataframe() (dataframe.DataFrame, error) {
	switch v := (*i.Expr)[0].(type) {
	case dataframe.DataFrame:
		return v, nil
	default:
		return dataframe.DataFrame{}, errors.New(fmt.Sprintf("expecting dataframe, got %T", v))
	}
}

func (i *PreludioInternal) GetValueList() ([]*PreludioInternal, error) {
	switch v := (*i.Expr)[0].(type) {
	case []*PreludioInternal:
		return v, nil
	default:
		return []*PreludioInternal{}, errors.New(fmt.Sprintf("expecting list, got %T", v))
	}
}

func (l *PreludioInternal) Mul(r *PreludioInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_MUL)
}

func (l *PreludioInternal) Div(r *PreludioInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_DIV)
}

func (l *PreludioInternal) Mod(r *PreludioInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_MOD)
}

func (l *PreludioInternal) Add(r *PreludioInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_ADD)
}

func (l *PreludioInternal) Sub(r *PreludioInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_SUB)
}

func (l *PreludioInternal) Pow(r *PreludioInternal) {
	(*l.Expr) = append((*l.Expr), (*r.Expr)...)
	(*l.Expr) = append((*l.Expr), BIN_EXPR_POW)
}

func IsOperator(t interface{}) (PreludioExprOp, bool) {
	switch v := t.(type) {
	case PreludioExprOp:
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

func (i *PreludioInternal) Solve() error {
	tmp := make([]interface{}, 1)

	// TODO: check if this is possible and
	// if it's the case to raise an error
	if len(*i.Expr) == 0 {
		return nil
	}

	// Check if the expression is a list
	// and recursively solve all the expressions
	// in the list
	if len(*i.Expr) == 1 {
		switch l := (*i.Expr)[0].(type) {
		case []*PreludioInternal:
			for _, t := range l {
				if err := t.Solve(); err != nil {
					return err
				}
			}
		default:
			return nil
		}
	}

	for len(*i.Expr) > 1 {
		t1 := (*i.Expr)[0]
		t2 := (*i.Expr)[1]

		var result interface{}

		// UNARY
		if op, ok := IsOperator(t2); ok {
			(*i.Expr) = (*i.Expr)[2:len((*i.Expr))]

			switch op {
			case UN_EXPR_ADD:
			case UN_EXPR_SUB:
			case UN_EXPR_NOT:
			}
		} else

		// BINARY
		{
			op, _ := IsOperator((*i.Expr)[2])
			(*i.Expr) = (*i.Expr)[3:len((*i.Expr))]

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
		(*i.Expr) = append(tmp, (*i.Expr)...)
	}
	return nil
}
