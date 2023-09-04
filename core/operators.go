package preludiocore

import (
	"fmt"
	"gandalff"
	"typesys"
)

func (vm *ByteEater) processList(p *__p_intern__) error {
	var series gandalff.Series

	list := p.expr[0].(__p_list__)
	for i := range list {

		// Skip assignments
		if list[i].tag == PRELUDIO_INTERNAL_TAG_ASSIGNMENT {
			return nil
		}

		if symb, ok := list[i].expr[0].(__p_symbol__); ok {
			list[i].expr[0] = vm.symbolResolution(symb)
		}

		if _, ok := list[i].expr[0].(__p_list__); ok {
			err := vm.processList(&list[i])
			if err != nil {
				return err
			}
		}

		if v, ok := list[i].expr[0].(gandalff.Series); ok {
			err := vm.solveExpr(&list[i])
			if err != nil {
				return err
			}

			if series == nil {
				series = v
			} else if v.Len() > 1 {
				// only append if the elements in the list are scalars
				return nil
			} else if series.Type() == v.Type() {
				series = series.Append(v)
			} else if series.Type().CanCoerceTo(v.Type()) {
				series = series.Cast(v.Type(), p.vm.__stringPool).Append(v)
			} else if v.Type().CanCoerceTo(series.Type()) {
				series = series.Append(v.Cast(series.Type(), p.vm.__stringPool))
			} else {
				return fmt.Errorf("cannot append %s to %s", v.Type().ToString(), series.Type().ToString())
			}
		}
	}

	p.expr[0] = series

	return nil
}

func (vm *ByteEater) solveExpr(p *__p_intern__) error {
	// Preprocess the expression
	// Check if elements in the expression are:
	//  - symbols: resolve them
	//  - lists: recursively solve all the sub-expressions
	for i := range p.expr {
		if symb, ok := p.expr[i].(__p_symbol__); ok {
			p.expr[i] = vm.symbolResolution(symb)
		}

		if _, ok := p.expr[i].(__p_list__); ok {
			err := vm.processList(p)
			if err != nil {
				return err
			}
		}
	}

	stack := make([]interface{}, 0)

	var op typesys.OPCODE
	var ok, errorMode bool
	var exprIdx int
	var result interface{}

	for len(p.expr) > 1 {

		// Load the stack until we find an operators
		ok = false
		for exprIdx = 0; !ok; op, ok = p.expr[exprIdx].(typesys.OPCODE) {
			exprIdx++
		}
		stack = append(stack, p.expr[0:exprIdx]...)
		p.expr = p.expr[exprIdx+1 : len(p.expr)]

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

		p.expr = append([]interface{}{result}, p.expr...)
	}

	return nil
}
