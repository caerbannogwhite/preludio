package preludiocore

import (
	"fmt"
	"gandalff"
	"typesys"
)

func solveExpr(vm *ByteEater, i *__p_intern__) error {
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
				if err := solveExpr(vm, &l[idx]); err != nil {
					return err
				}
			}

		default:
		}
	}

	stack := make([]interface{}, 0)

	for len(i.expr) > 1 {

		// Load the stack until we find an operators
		var ok bool
		var op typesys.OPCODE
		for {
			if op, ok = isOperator(i.expr[0]); ok {
				i.expr = i.expr[1:len(i.expr)]
				break
			}

			stack = append(stack, i.expr[0])
			i.expr = i.expr[1:len(i.expr)]
		}

		var result interface{}

		// UNARY
		// if op, ok := isOperator(t2); ok {
		// 	i.expr = i.expr[2:len(i.expr)]

		// 	if s, ok := t1.(__p_symbol__); ok {
		// 		t1 = vm.symbolResolution(s)
		// 	}

		// 	switch op {
		// 	case typesys.OP_UNARY_ADD:
		// 	case typesys.OP_UNARY_SUB:
		// 	case typesys.OP_UNARY_NOT:
		// 	}
		// } else

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

			switch op {
			case typesys.OP_BINARY_MUL:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Mul(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_DIV:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Div(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_MOD:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Mod(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_POW:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Pow(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_ADD:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Add(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_SUB:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Sub(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_AND:
				if s1, ok := t1.(gandalff.SeriesBool); ok {
					result = s1.And(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_OR:
				if s1, ok := t1.(gandalff.SeriesBool); ok {
					result = s1.Or(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_EQ:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Eq(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_NE:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Ne(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_LT:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Lt(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_LE:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Le(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_GT:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Gt(t2.(gandalff.Series))
				}

			case typesys.OP_BINARY_GE:
				if s1, ok := t1.(gandalff.Series); ok {
					result = s1.Ge(t2.(gandalff.Series))
				}
			}

			if _, ok := result.(gandalff.SeriesError); ok {
				return fmt.Errorf("operator %s not supported between %s and %s",
					operatorToString(op),
					t1.(gandalff.Series).Type().ToString(),
					t2.(gandalff.Series).Type().ToString())
			}
		}

		i.expr = append([]interface{}{result}, i.expr...)
	}
	return nil
}
