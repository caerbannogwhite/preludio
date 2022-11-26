package main

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

func (e *PrqlExpr) Solve() interface{} {
	return e.stack[0]
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
