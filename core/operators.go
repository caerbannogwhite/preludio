package preludiocore

import (
	"fmt"
	"gandalff"
	"typesys"
)

func (vm *ByteEater) solveExpr(i *__p_intern__) error {
	// TODO: check if this is possible and
	// if it's the case to raise an error
	if i == nil || i.expr == nil || len(i.expr) == 0 {
		return fmt.Errorf("invalid expression")
	}

	// Check if the expression is:
	//  - a symbol: resolve it
	//  - a list: recursively solve all the expressions
	if len(i.expr) == 1 {
		switch l := i.expr[0].(type) {
		case __p_symbol__:
			i.expr[0] = vm.symbolResolution(l)

		case __p_list__:
			for idx := range l {
				if err := vm.solveExpr(&l[idx]); err != nil {
					return err
				}
			}

			return i.processList()
		}
	}

	stack := make([]interface{}, 0)

	var op typesys.OPCODE
	var ok, errorMode bool
	var exprIdx int
	var result interface{}

	for len(i.expr) > 1 {

		// Load the stack until we find an operators
		ok = false
		for exprIdx = 0; !ok; op, ok = i.expr[exprIdx].(typesys.OPCODE) {
			exprIdx++
		}
		stack = append(stack, i.expr[0:exprIdx]...)
		i.expr = i.expr[exprIdx+1 : len(i.expr)]

		errorMode = false
		result = gandalff.SeriesError{}

		// UNARY
		if op.IsUnaryOp() {
			t1 := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			if s, ok := t1.(__p_symbol__); ok {
				t1 = vm.symbolResolution(s)
			}

			switch op {
			case typesys.OP_UNARY_ADD:
				result = t1

			case typesys.OP_UNARY_SUB:
				switch s1 := t1.(type) {
				case gandalff.SeriesInt32:
					result = s1.Neg()
				case gandalff.SeriesInt64:
					result = s1.Neg()
				case gandalff.SeriesFloat64:
					result = s1.Neg()
				default:
					errorMode = true
				}

			case typesys.OP_UNARY_NOT:
				if s1, ok := t1.(gandalff.SeriesBool); ok {
					result = s1.Not()
				} else {
					errorMode = true
				}
			}

			// Check for errors
			if _, ok := result.(gandalff.SeriesError); ok || errorMode {
				return fmt.Errorf("unary operator %s not supported for %s",
					operatorToCode(op),
					t1.(gandalff.Series).TypeCard().ToString())
			}
		} else

		// BINARY
		{
			t2 := stack[len(stack)-1]
			t1 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]

			// Symbol resolution
			if s, ok := t1.(__p_symbol__); ok {
				t1 = vm.symbolResolution(s)
			}

			if s, ok := t2.(__p_symbol__); ok {
				t2 = vm.symbolResolution(s)
			}

			// Type check
			s1 := t1.(gandalff.Series)
			s2 := t2.(gandalff.Series)

			switch op {
			case typesys.OP_BINARY_MUL:
				result = s1.Mul(s2)

			case typesys.OP_BINARY_DIV:
				result = s1.Div(s2)

			case typesys.OP_BINARY_MOD:
				result = s1.Mod(s2)

			case typesys.OP_BINARY_POW:
				result = s1.Pow(s2)

			case typesys.OP_BINARY_ADD:
				result = s1.Add(s2)

			case typesys.OP_BINARY_SUB:
				result = s1.Sub(s2)

			case typesys.OP_BINARY_EQ:
				result = s1.Eq(s2)

			case typesys.OP_BINARY_NE:
				result = s1.Ne(s2)

			case typesys.OP_BINARY_LT:
				result = s1.Lt(s2)

			case typesys.OP_BINARY_LE:
				result = s1.Le(s2)

			case typesys.OP_BINARY_GT:
				result = s1.Gt(s2)

			case typesys.OP_BINARY_GE:
				result = s1.Ge(s2)

			case typesys.OP_BINARY_AND:
				if s1, ok := t1.(gandalff.SeriesBool); ok {
					result = s1.And(s2)
				} else {
					errorMode = true
				}

			case typesys.OP_BINARY_OR:
				if s1, ok := t1.(gandalff.SeriesBool); ok {
					result = s1.Or(s2)
				} else {
					errorMode = true
				}
			}

			// Check for errors
			if _, ok := result.(gandalff.SeriesError); ok || errorMode {
				return fmt.Errorf("binary operator %s not supported between %s and %s",
					operatorToString(op),
					s1.TypeCard().ToString(),
					s2.TypeCard().ToString())
			}
		}

		i.expr = append([]interface{}{result}, i.expr...)
	}

	return nil
}
