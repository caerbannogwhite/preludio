package preludiocore

import (
	"fmt"
	"gandalff"
	"preludiometa"
)

func (vm *ByteEater) processList(list *pList) (interface{}, error) {
	convertToSeries := true

	var series gandalff.Series
	for i := range *list {

		if (*list)[i].tag == PRELUDIO_INTERNAL_TAG_ASSIGNMENT {
			convertToSeries = false
			break
		}

		switch v := (*list)[i].value.(type) {
		case pList:
			convertToSeries = false

		case gandalff.Series:
			if series == nil {
				series = v
			} else if v.Len() > 1 {
				convertToSeries = false
				break
			} else if series.Type() == v.Type() {
				series = series.Append(v)
			} else if series.Type().CanCoerceTo(v.Type()) {
				series = series.Cast(v.Type()).Append(v)
			} else if v.Type().CanCoerceTo(series.Type()) {
				series = series.Append(v.Cast(series.Type()))
			} else {
				return list, fmt.Errorf("cannot append %s to %s", v.Type().ToString(), series.Type().ToString())
			}
		}
	}

	if convertToSeries {
		return series, nil
	}
	return *list, nil
}

func (vm *ByteEater) handleOperator(op preludiometa.OPCODE) {

	var operand1, operand2 *pIntern

	switch op {

	///////////////////////////////////////////////////////////////////////
	///////////				ARITHMETIC OPERATIONS

	case preludiometa.OP_BINARY_MUL:
		vm.printDebug(10, "OP_BINARY_MUL", "", "")

		operand2 = vm.stackPop()
		operand1 = vm.stackPop()

		// Check if the operands are series
		if s1, ok := operand1.value.(gandalff.Series); ok {
			if s2, ok := operand2.value.(gandalff.Series); ok {
				operand1.value = s1.Mul(s2)
				vm.stackPush(operand1)
				return
			}
		}

	case preludiometa.OP_BINARY_DIV:
		vm.printDebug(10, "OP_BINARY_DIV", "", "")

		operand2 = vm.stackPop()
		operand1 = vm.stackPop()

		// Check if the operands are series
		if s1, ok := operand1.value.(gandalff.Series); ok {
			if s2, ok := operand2.value.(gandalff.Series); ok {
				operand1.value = s1.Div(s2)
				vm.stackPush(operand1)
				return
			}
		}

	case preludiometa.OP_BINARY_MOD:
		vm.printDebug(10, "OP_BINARY_MOD", "", "")

	case preludiometa.OP_BINARY_EXP:
		vm.printDebug(10, "OP_BINARY_EXP", "", "")

	case preludiometa.OP_BINARY_ADD:
		vm.printDebug(10, "OP_BINARY_ADD", "", "")

	case preludiometa.OP_BINARY_SUB:
		vm.printDebug(10, "OP_BINARY_SUB", "", "")

	case preludiometa.OP_BINARY_EQ:
		vm.printDebug(10, "OP_BINARY_EQ", "", "")

	case preludiometa.OP_BINARY_NE:
		vm.printDebug(10, "OP_BINARY_NE", "", "")

	case preludiometa.OP_BINARY_GE:
		vm.printDebug(10, "OP_BINARY_GE", "", "")

	case preludiometa.OP_BINARY_LE:
		vm.printDebug(10, "OP_BINARY_LE", "", "")

	case preludiometa.OP_BINARY_GT:
		vm.printDebug(10, "OP_BINARY_GT", "", "")

	case preludiometa.OP_BINARY_LT:
		vm.printDebug(10, "OP_BINARY_LT", "", "")

	case preludiometa.OP_BINARY_AND:
		vm.printDebug(10, "OP_BINARY_AND", "", "")

	case preludiometa.OP_BINARY_OR:
		vm.printDebug(10, "OP_BINARY_OR", "", "")

	///////////////////////////////////////////////////////////////////////
	///////////				OTHER OPERATIONS

	case preludiometa.OP_BINARY_COALESCE:
		vm.printDebug(10, "OP_BINARY_COALESCE", "", "")

	case preludiometa.OP_BINARY_MODEL:
		vm.printDebug(10, "OP_BINARY_MODEL", "", "")

	case preludiometa.OP_INDEXING:
		vm.printDebug(10, "OP_INDEXING", "", "")

	case preludiometa.OP_HELP:
		vm.printDebug(10, "OP_HELP", "", "")

	///////////////////////////////////////////////////////////////////////
	///////////				UNARY OPERATIONS

	case preludiometa.OP_UNARY_REV:
		vm.printDebug(10, "OP_UNARY_REV", "", "")

	case preludiometa.OP_UNARY_SUB:
		vm.printDebug(10, "OP_UNARY_SUB", "", "")

	case preludiometa.OP_UNARY_ADD:
		vm.printDebug(10, "OP_UNARY_ADD", "", "")

	case preludiometa.OP_UNARY_NOT:
		vm.printDebug(10, "OP_UNARY_NOT", "", "")

	}
}

// 	switch op {
// 	case preludiometa.OP_UNARY_ADD:

// }

// func (vm *ByteEater) solveExpr(p *pIntern) error {
// 	// Preprocess the expression
// 	// Check if elements in the expression are:
// 	//  - symbols: resolve them
// 	//  - lists: recursively solve all the sub-expressions
// 	var err error
// 	for i := range p.expr {
// 		if symb, ok := p.expr[i].(__p_symbol__); ok {
// 			p.expr[i] = vm.symbolResolution(symb)
// 		}

// 		if list, ok := p.expr[i].(pList); ok {
// 			for j := range list {
// 				err = vm.solveExpr(&list[j])
// 				if err != nil {
// 					return err
// 				}
// 			}

// 			p.expr[i], err = vm.processList(&list)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	stack := make([]interface{}, 0)

// 	var op preludiometa.OPCODE
// 	var ok, errorMode bool
// 	var exprIdx int
// 	var result interface{}

// 	for len(p.expr) > 1 {

// 		// Load the stack until we find an operators
// 		ok = false
// 		for exprIdx = 0; !ok; op, ok = p.expr[exprIdx].(preludiometa.OPCODE) {
// 			exprIdx++
// 		}
// 		stack = append(stack, p.expr[0:exprIdx]...)
// 		p.expr = p.expr[exprIdx+1 : len(p.expr)]

// 		errorMode = false
// 		result = gandalff.SeriesError{}

// 		// UNARY
// 		if op.IsUnaryOp() {
// 			t1 := stack[len(stack)-1]
// 			stack = stack[0 : len(stack)-1]

// 			switch op {
// 			case preludiometa.OP_UNARY_ADD:
// 				result = t1

// 			case preludiometa.OP_UNARY_SUB:
// 				switch s1 := t1.(type) {
// 				case gandalff.SeriesInt:
// 					result = s1.Neg()
// 				case gandalff.SeriesInt64:
// 					result = s1.Neg()
// 				case gandalff.SeriesFloat64:
// 					result = s1.Neg()
// 				default:
// 					errorMode = true
// 				}

// 			case preludiometa.OP_UNARY_NOT:
// 				if s1, ok := t1.(gandalff.SeriesBool); ok {
// 					result = s1.Not()
// 				} else {
// 					errorMode = true
// 				}
// 			}

// 			// Check for errors
// 			if _, ok := result.(gandalff.SeriesError); ok || errorMode {
// 				return fmt.Errorf("unary operator %s not supported for %s",
// 					op.ToCodeString(),
// 					t1.(gandalff.Series).TypeCard().ToString())
// 			}
// 		} else

// 		// BINARY
// 		{
// 			s2 := stack[len(stack)-1].(gandalff.Series)
// 			s1 := stack[len(stack)-2].(gandalff.Series)
// 			stack = stack[0 : len(stack)-2]

// 			switch op {
// 			case preludiometa.OP_BINARY_MUL:
// 				result = s1.Mul(s2)

// 			case preludiometa.OP_BINARY_DIV:
// 				result = s1.Div(s2)

// 			case preludiometa.OP_BINARY_MOD:
// 				result = s1.Mod(s2)

// 			case preludiometa.OP_BINARY_EXP:
// 				result = s1.Exp(s2)

// 			case preludiometa.OP_BINARY_ADD:
// 				result = s1.Add(s2)

// 			case preludiometa.OP_BINARY_SUB:
// 				result = s1.Sub(s2)

// 			case preludiometa.OP_BINARY_EQ:
// 				result = s1.Eq(s2)

// 			case preludiometa.OP_BINARY_NE:
// 				result = s1.Ne(s2)

// 			case preludiometa.OP_BINARY_LT:
// 				result = s1.Lt(s2)

// 			case preludiometa.OP_BINARY_LE:
// 				result = s1.Le(s2)

// 			case preludiometa.OP_BINARY_GT:
// 				result = s1.Gt(s2)

// 			case preludiometa.OP_BINARY_GE:
// 				result = s1.Ge(s2)

// 			case preludiometa.OP_BINARY_AND:
// 				if s1, ok := s1.(gandalff.SeriesBool); ok {
// 					result = s1.And(s2)
// 				} else {
// 					errorMode = true
// 				}

// 			case preludiometa.OP_BINARY_OR:
// 				if s1, ok := s1.(gandalff.SeriesBool); ok {
// 					result = s1.Or(s2)
// 				} else {
// 					errorMode = true
// 				}
// 			}

// 			// Check for errors
// 			if _, ok := result.(gandalff.SeriesError); ok || errorMode {
// 				return fmt.Errorf("binary operator %s not supported between %s and %s",
// 					op.ToString(),
// 					s1.TypeCard().ToString(),
// 					s2.TypeCard().ToString())
// 			}
// 		}

// 		p.expr = append([]interface{}{result}, p.expr...)
// 	}

// 	return nil
// }
