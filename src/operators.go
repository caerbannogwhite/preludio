package preludiocore

import (
	"fmt"

	"github.com/caerbannogwhite/aargh/meta"
	"github.com/caerbannogwhite/aargh/series"
)

func (vm *ByteEater) processList(list *__p_list__) (interface{}, error) {
	convertToSeries := true

	var series_ series.Series
	for i := range *list {

		if (*list)[i].tag == PRELUDIO_INTERNAL_TAG_ASSIGNMENT {
			convertToSeries = false
			break
		}

		switch v := (*list)[i].expr[0].(type) {
		case __p_list__:
			convertToSeries = false

		case series.Series:
			if series_ == nil {
				series_ = v
			} else if v.Len() > 1 {
				convertToSeries = false
				break
			} else if series_.Type() == v.Type() {
				series_ = series_.Append(v)
			} else if series_.Type().CanCoerceTo(v.Type()) {
				series_ = series_.Cast(v.Type()).Append(v)
			} else if v.Type().CanCoerceTo(series_.Type()) {
				series_ = series_.Append(v.Cast(series_.Type()))
			} else {
				return list, fmt.Errorf("cannot append %s to %s", v.Type().String(), series_.Type().String())
			}
		}
	}

	if convertToSeries {
		return series_, nil
	}
	return *list, nil
}

func (vm *ByteEater) solveExpr(p *__p_intern__) error {
	// Preprocess the expression
	// Check if elements in the expression are:
	//  - symbols: resolve them
	//  - lists: recursively solve all the sub-expressions
	var err error
	for i := range p.expr {
		if symb, ok := p.expr[i].(__p_symbol__); ok {
			p.expr[i] = vm.symbolResolution(symb)
		}

		if list, ok := p.expr[i].(__p_list__); ok {
			for j := range list {
				err = vm.solveExpr(&list[j])
				if err != nil {
					return err
				}
			}

			p.expr[i], err = vm.processList(&list)
			if err != nil {
				return err
			}
		}
	}

	stack := make([]interface{}, 0)

	var op meta.OPCODE
	var ok, errorMode bool
	var exprIdx int
	var result interface{}

	for len(p.expr) > 1 {

		// Load the stack until we find an operators
		ok = false
		for exprIdx = 0; !ok; op, ok = p.expr[exprIdx].(meta.OPCODE) {
			exprIdx++
		}
		stack = append(stack, p.expr[0:exprIdx]...)
		p.expr = p.expr[exprIdx+1 : len(p.expr)]

		errorMode = false
		result = series.Errors{}

		// UNARY
		if op.IsUnaryOp() {
			t1 := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			switch op {
			case meta.OP_UNARY_ADD:
				result = t1

			case meta.OP_UNARY_SUB:
				switch s1 := t1.(type) {
				case series.Ints:
					result = s1.Neg()
				case series.Int64s:
					result = s1.Neg()
				case series.Float64s:
					result = s1.Neg()
				default:
					errorMode = true
				}

			case meta.OP_UNARY_NOT:
				if s1, ok := t1.(series.Bools); ok {
					result = s1.Not()
				} else {
					errorMode = true
				}
			}

			// Check for errors
			if _, ok := result.(series.Errors); ok || errorMode {
				return fmt.Errorf("unary operator %s not supported for %s",
					op.ToCodeString(),
					t1.(series.Series).TypeCard().ToString())
			}
		} else

		// BINARY
		{
			s2 := stack[len(stack)-1].(series.Series)
			s1 := stack[len(stack)-2].(series.Series)
			stack = stack[0 : len(stack)-2]

			switch op {
			case meta.OP_BINARY_MUL:
				result = s1.Mul(s2)

			case meta.OP_BINARY_DIV:
				result = s1.Div(s2)

			case meta.OP_BINARY_MOD:
				result = s1.Mod(s2)

			case meta.OP_BINARY_EXP:
				result = s1.Exp(s2)

			case meta.OP_BINARY_ADD:
				result = s1.Add(s2)

			case meta.OP_BINARY_SUB:
				result = s1.Sub(s2)

			case meta.OP_BINARY_EQ:
				result = s1.Eq(s2)

			case meta.OP_BINARY_NE:
				result = s1.Ne(s2)

			case meta.OP_BINARY_LT:
				result = s1.Lt(s2)

			case meta.OP_BINARY_LE:
				result = s1.Le(s2)

			case meta.OP_BINARY_GT:
				result = s1.Gt(s2)

			case meta.OP_BINARY_GE:
				result = s1.Ge(s2)

			case meta.OP_BINARY_AND:
				if s1, ok := s1.(series.Bools); ok {
					result = s1.And(s2)
				} else {
					errorMode = true
				}

			case meta.OP_BINARY_OR:
				if s1, ok := s1.(series.Bools); ok {
					result = s1.Or(s2)
				} else {
					errorMode = true
				}
			}

			// Check for errors
			if _, ok := result.(series.Errors); ok || errorMode {
				return fmt.Errorf("binary operator %s not supported between %s and %s",
					op.ToString(),
					s1.TypeCard().ToString(),
					s2.TypeCard().ToString())
			}
		}

		p.expr = append([]interface{}{result}, p.expr...)
	}

	return nil
}
