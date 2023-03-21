package preludiocore

import (
	"math"
	"preludiocompiler"
	"testing"
)

var be *ByteEater

func init() {
	be = new(ByteEater).InitVM()
}

func Test_Operator_Mul(t *testing.T) {

	var err error

	b1 := newPInternTerm([]bool{true, false, true, false})
	b2 := newPInternTerm([]bool{false, false, true, true})
	in := newPInternTerm([]int{1, 2, 3, 4})
	fl := newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL * BOOL
		b1.appendOperand(preludiocompiler.OP_BINARY_MUL, b2)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 0 || v[1] != 0 || v[2] != 1 || v[3] != 0 {
				t.Error("Expected [0, 0, 1, 0], got", v)
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * INTEGER
		b1.appendOperand(preludiocompiler.OP_BINARY_MUL, in)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 1 || v[1] != 0 || v[2] != 3 || v[3] != 0 {
				t.Error("Expected [1, 0, 3, 0] got", v)
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * FLOAT
		b1.appendOperand(preludiocompiler.OP_BINARY_MUL, fl)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := b1.getFloatVector()
			if v[0] != 5.0 || v[1] != 0.0 || v[2] != 7.0 || v[3] != 0.0 {
				t.Error("Expected [5.0, 0.0, 7.0, 0.0] got", v)
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * STRING
		b1.appendOperand(preludiocompiler.OP_BINARY_MUL, st)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := b1.getStringVector()
			if v[0] != "a" || v[1] != "" || v[2] != "c" || v[3] != "" {
				t.Error("Expected [\"a\", \"\", \"c\", \"\"] got", v)
			}
		}
	}

	// INTEGER
	{
		// INTEGER * BOOL
		in.appendOperand(preludiocompiler.OP_BINARY_MUL, b2)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 0 || v[1] != 0 || v[2] != 3 || v[3] != 4 {
				t.Error("Expected [0, 0, 3, 4] got", v)
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * INTEGER
		in.appendOperand(preludiocompiler.OP_BINARY_MUL, in)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 1 || v[1] != 4 || v[2] != 9 || v[3] != 16 {
				t.Error("Expected [1, 4, 9, 16] got", v)
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * FLOAT
		in.appendOperand(preludiocompiler.OP_BINARY_MUL, fl)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := in.getFloatVector()
			if v[0] != 5.0 || v[1] != 12.0 || v[2] != 21.0 || v[3] != 32.0 {
				t.Error("Expected [5.0, 12.0, 21.0, 32.0] got", v)
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * STRING
		in.appendOperand(preludiocompiler.OP_BINARY_MUL, st)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := in.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "ccc" || v[3] != "dddd" {
				t.Error("Expected [\"a\", \"bb\", \"ccc\", \"dddd\"]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT * BOOL
		fl.appendOperand(preludiocompiler.OP_BINARY_MUL, b2)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 0.0 || v[1] != 0.0 || v[2] != 7.0 || v[3] != 8.0 {
				t.Error("Expected [0.0, 0.0, 7.0, 8.0] got", v)
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * INTEGER
		fl.appendOperand(preludiocompiler.OP_BINARY_MUL, in)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 5.0 || v[1] != 12.0 || v[2] != 21.0 || v[3] != 32.0 {
				t.Error("Expected [0.0, 0.0, 21.0, 0.0] got", v)
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * FLOAT
		fl.appendOperand(preludiocompiler.OP_BINARY_MUL, fl)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 25.0 || v[1] != 36.0 || v[2] != 49.0 || v[3] != 64.0 {
				t.Error("Expected [25.0, 36.0, 49.0, 64.0] got", v)
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * STRING
		fl.appendOperand(preludiocompiler.OP_BINARY_MUL, st)
		err = solveExpr(be, fl)

		if err == nil || err.Error() != "binary * operator not implemented for []float64 and []string" {
			t.Error(err)
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING * BOOL
		st.appendOperand(preludiocompiler.OP_BINARY_MUL, b2)
		err = solveExpr(be, st)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "" || v[1] != "" || v[2] != "c" || v[3] != "d" {
				t.Error("Expected [\"\", \"\", \"c\", \"d\"] got", v)
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * INTEGER
		st.appendOperand(preludiocompiler.OP_BINARY_MUL, in)
		err = solveExpr(be, st)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "ccc" || v[3] != "dddd" {
				t.Error("Expected [\"a\", \"bb\", \"ccc\", \"dddd\"] got", v)
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * FLOAT
		st.appendOperand(preludiocompiler.OP_BINARY_MUL, fl)
		err = solveExpr(be, st)

		if err == nil || err.Error() != "binary * operator not implemented for []string and []float64" {
			t.Error(err)
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * STRING
		st.appendOperand(preludiocompiler.OP_BINARY_MUL, st)
		err = solveExpr(be, st)

		if err == nil || err.Error() != "binary * operator not implemented for []string and []string" {
			t.Error(err)
		}
	}
}

func Test_Operator_Div(t *testing.T) {

	var err error

	b1 := newPInternTerm([]bool{true, false, true, false})
	b2 := newPInternTerm([]bool{false, false, true, true})
	in := newPInternTerm([]int{1, 2, 3, 4})
	fl := newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL / BOOL
		b1.appendOperand(preludiocompiler.OP_BINARY_DIV, b2)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := b1.getFloatVector()
			if !math.IsInf(v[0], 1) || !math.IsNaN(v[1]) || v[2] != 1.0 || v[3] != 0.0 {
				t.Error("Expected [+Inf, NaN, 1.0, 0.0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL / INTEGER
		b1.appendOperand(preludiocompiler.OP_BINARY_DIV, in)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := b1.getFloatVector()
			if v[0] != 1.0 || v[1] != 0.0 || v[2] != 0.3333333333333333 || v[3] != 0.0 {
				t.Error("Expected [1.0, 0.0, 0.3333333333333333, 0.0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL / FLOAT
		b1.appendOperand(preludiocompiler.OP_BINARY_DIV, fl)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := b1.getFloatVector()
			if v[0] != 0.2 || v[1] != 0.0 || v[2] != 0.14285714285714285 || v[3] != 0.0 {
				t.Error("Expected [0.2, 0.0, 0.14285714285714285, 0.0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL / STRING
		b1.appendOperand(preludiocompiler.OP_BINARY_DIV, st)
		err = solveExpr(be, b1)

		if err == nil || err.Error() != "binary / operator not implemented for []bool and []string" {
			t.Error(err)
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})
	}

	// INTEGER
	{
		// INTEGER / BOOL
		in.appendOperand(preludiocompiler.OP_BINARY_DIV, b1)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := in.getFloatVector()
			if v[0] != 1.0 || !math.IsInf(v[1], 1) || v[2] != 3.0 || !math.IsInf(v[3], 1) {
				t.Error("Expected [1.0, +Inf, 3.0, +Inf]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER / INTEGER
		in.appendOperand(preludiocompiler.OP_BINARY_DIV, in)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := in.getFloatVector()
			if v[0] != 1.0 || v[1] != 1.0 || v[2] != 1.0 || v[3] != 1.0 {
				t.Error("Expected [1.0, 1.0, 1.0, 1.0]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER / FLOAT
		in.appendOperand(preludiocompiler.OP_BINARY_DIV, fl)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := in.getFloatVector()
			if v[0] != 0.2 || v[1] != 0.3333333333333333 || v[2] != 0.42857142857142855 || v[3] != 0.5 {
				t.Error("Expected [0.2, 0.3333333333333333, 0.42857142857142855, 0.5]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER / STRING
		in.appendOperand(preludiocompiler.OP_BINARY_DIV, st)
		err = solveExpr(be, in)

		if err == nil || err.Error() != "binary / operator not implemented for []int and []string" {
			t.Error(err)
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT / BOOL
		fl.appendOperand(preludiocompiler.OP_BINARY_DIV, b1)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 5.0 || !math.IsInf(v[1], 1) || v[2] != 7.0 || !math.IsInf(v[3], 1) {
				t.Error("Expected [5.0, +Inf, 7.0, +Inf]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT / INTEGER
		fl.appendOperand(preludiocompiler.OP_BINARY_DIV, in)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 5.0 || v[1] != 3.0 || v[2] != 2.3333333333333335 || v[3] != 2.0 {
				t.Error("Expected [5.0, 3.0, 2.3333333333333335, 2.0]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT / FLOAT
		fl.appendOperand(preludiocompiler.OP_BINARY_DIV, fl)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 1 || v[1] != 1 || v[2] != 1 || v[3] != 1 {
				t.Error("Expected [1, 1, 1, 1]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT / STRING
		fl.appendOperand(preludiocompiler.OP_BINARY_DIV, st)
		err = solveExpr(be, fl)

		if err == nil || err.Error() != "binary / operator not implemented for []float64 and []string" {
			t.Error(err)
		}
	}
}

func Test_Operator_Add(t *testing.T) {

	var err error

	b1 := newPInternTerm([]bool{true, false, true, false})
	b2 := newPInternTerm([]bool{false, false, true, true})
	in := newPInternTerm([]int{1, 2, 3, 4})
	fl := newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL + BOOL
		b1.appendOperand(preludiocompiler.OP_BINARY_ADD, b2)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 1 || v[1] != 0 || v[2] != 2 || v[3] != 1 {
				t.Error("Expected [1, 0, 2, 1] got", v)
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL + INTEGER
		b1.appendOperand(preludiocompiler.OP_BINARY_ADD, in)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isIntegerVector() {
			t.Error("Expected int vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 2 || v[1] != 2 || v[2] != 4 || v[3] != 4 {
				t.Error("Expected [2, 2, 4, 4] got", v)
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL + FLOAT
		b1.appendOperand(preludiocompiler.OP_BINARY_ADD, fl)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := b1.getFloatVector()
			if v[0] != 6.0 || v[1] != 6.0 || v[2] != 8.0 || v[3] != 8.0 {
				t.Error("Expected [6.0, 6.0, 8.0, 8.0] got", v)
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL + STRING
		b1.appendOperand(preludiocompiler.OP_BINARY_ADD, st)
		err = solveExpr(be, b1)

		if err != nil {
			t.Error(err)
		} else if !b1.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := b1.getStringVector()
			if v[0] != "truea" || v[1] != "falseb" || v[2] != "truec" || v[3] != "falsed" {
				t.Error("Expected [truea, falseb, truec, falsed] got", v)
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})
	}

	// INTEGER
	{
		// INTEGER + BOOL
		in.appendOperand(preludiocompiler.OP_BINARY_ADD, b1)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isIntegerVector() {
			t.Error("Expected int vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 2 || v[1] != 2 || v[2] != 4 || v[3] != 4 {
				t.Error("Expected [2, 2, 4, 4] got", v)
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER + INTEGER
		in.appendOperand(preludiocompiler.OP_BINARY_ADD, in)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isIntegerVector() {
			t.Error("Expected int vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 2 || v[1] != 4 || v[2] != 6 || v[3] != 8 {
				t.Error("Expected [2, 4, 6, 8] got", v)
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER + FLOAT
		in.appendOperand(preludiocompiler.OP_BINARY_ADD, fl)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := in.getFloatVector()
			if v[0] != 6.0 || v[1] != 8.0 || v[2] != 10.0 || v[3] != 12.0 {
				t.Error("Expected [6.0, 8.0, 10.0, 12.0] got", v)
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER + STRING
		in.appendOperand(preludiocompiler.OP_BINARY_ADD, st)
		err = solveExpr(be, in)

		if err != nil {
			t.Error(err)
		} else if !in.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := in.getStringVector()
			if v[0] != "1a" || v[1] != "2b" || v[2] != "3c" || v[3] != "4d" {
				t.Error("Expected [1a, 2b, 3c, 4d] got", v)
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT + BOOL
		fl.appendOperand(preludiocompiler.OP_BINARY_ADD, b1)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 6.0 || v[1] != 6.0 || v[2] != 8.0 || v[3] != 8.0 {
				t.Error("Expected [6.0, 6.0, 8.0, 8.0] got", v)
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT + INTEGER
		fl.appendOperand(preludiocompiler.OP_BINARY_ADD, in)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 6.0 || v[1] != 8.0 || v[2] != 10.0 || v[3] != 12.0 {
				t.Error("Expected [6.0, 8.0, 10.0, 12.0] got", v)
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT + FLOAT
		fl.appendOperand(preludiocompiler.OP_BINARY_ADD, fl)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 10.0 || v[1] != 12.0 || v[2] != 14.0 || v[3] != 16.0 {
				t.Error("Expected [10.0, 12.0, 14.0, 16.0] got", v)
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT + STRING
		fl.appendOperand(preludiocompiler.OP_BINARY_ADD, st)
		err = solveExpr(be, fl)

		if err != nil {
			t.Error(err)
		} else if !fl.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := fl.getStringVector()
			if v[0] != "5.000000a" || v[1] != "6.000000b" || v[2] != "7.000000c" || v[3] != "8.000000d" {
				t.Error("Expected [5.000000a, 6.000000b, 7.000000c, 8.000000d] got", v)
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING + BOOL
		st.appendOperand(preludiocompiler.OP_BINARY_ADD, b1)
		err = solveExpr(be, st)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "atrue" || v[1] != "bfalse" || v[2] != "ctrue" || v[3] != "dfalse" {
				t.Error("Expected [atrue, bfalse, ctrue, dfalse] got", v)
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING + INTEGER
		st.appendOperand(preludiocompiler.OP_BINARY_ADD, in)
		err = solveExpr(be, st)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "a1" || v[1] != "b2" || v[2] != "c3" || v[3] != "d4" {
				t.Error("Expected [a1, b2, c3, d4] got", v)
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING + FLOAT
		st.appendOperand(preludiocompiler.OP_BINARY_ADD, fl)
		err = solveExpr(be, st)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "a5.000000" || v[1] != "b6.000000" || v[2] != "c7.000000" || v[3] != "d8.000000" {
				t.Error("Expected [a5.000000, b6.000000, c7.000000, d8.000000] got", v)
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING + STRING
		st.appendOperand(preludiocompiler.OP_BINARY_ADD, st)
		err = solveExpr(be, st)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "aa" || v[1] != "bb" || v[2] != "cc" || v[3] != "dd" {
				t.Error("Expected [aa, bb, cc, dd] got", v)
			}
		}
	}
}
