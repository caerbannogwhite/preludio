package preludio

import (
	"math"
	"testing"
)

var be *ByteEater

func init() {
	be = new(ByteEater).InitVM()
}

func TestMul(t *testing.T) {

	var err error

	b1 := newPInternTerm([]bool{true, false, true, false})
	b2 := newPInternTerm([]bool{false, false, true, true})
	in := newPInternTerm([]int{1, 2, 3, 4})
	fl := newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL * BOOL
		b1.appendOperand(OP_BINARY_MUL, b2)
		err = b1.solve(be)

		if err != nil {
			t.Error(err)
		} else if !b1.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 0 || v[1] != 0 || v[2] != 1 || v[3] != 0 {
				t.Error("Expected [0, 0, 1, 0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * INTEGER
		b1.appendOperand(OP_BINARY_MUL, in)
		err = b1.solve(be)

		if err != nil {
			t.Error(err)
		} else if !b1.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 1 || v[1] != 0 || v[2] != 3 || v[3] != 0 {
				t.Error("Expected [1, 0, 3, 0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * FLOAT
		b1.appendOperand(OP_BINARY_MUL, fl)
		err = b1.solve(be)

		if err != nil {
			t.Error(err)
		} else if !b1.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := b1.getFloatVector()
			if v[0] != 5.0 || v[1] != 0.0 || v[2] != 7.0 || v[3] != 0.0 {
				t.Error("Expected [5.0, 0.0, 7.0, 0.0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * STRING
		b1.appendOperand(OP_BINARY_MUL, st)
		err = b1.solve(be)

		if err != nil {
			t.Error(err)
		} else if !b1.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := b1.getStringVector()
			if v[0] != "a" || v[1] != "" || v[2] != "c" || v[3] != "" {
				t.Error("Expected [\"a\", \"\", \"c\", \"\"]")
			}
		}
	}

	// INTEGER
	{
		// INTEGER * BOOL
		in.appendOperand(OP_BINARY_MUL, b2)
		err = in.solve(be)

		if err != nil {
			t.Error(err)
		} else if !in.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 0 || v[1] != 0 || v[2] != 3 || v[3] != 4 {
				t.Error("Expected [0, 0, 3, 4]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * INTEGER
		in.appendOperand(OP_BINARY_MUL, in)
		err = in.solve(be)

		if err != nil {
			t.Error(err)
		} else if !in.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 1 || v[1] != 4 || v[2] != 9 || v[3] != 16 {
				t.Error("Expected [1, 4, 9, 16]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * FLOAT
		in.appendOperand(OP_BINARY_MUL, fl)
		err = in.solve(be)

		if err != nil {
			t.Error(err)
		} else if !in.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := in.getFloatVector()
			if v[0] != 5.0 || v[1] != 12.0 || v[2] != 21.0 || v[3] != 32.0 {
				t.Error("Expected [5.0, 12.0, 21.0, 32.0]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * STRING
		in.appendOperand(OP_BINARY_MUL, st)
		err = in.solve(be)

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
		fl.appendOperand(OP_BINARY_MUL, b2)
		err = fl.solve(be)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 0.0 || v[1] != 0.0 || v[2] != 7.0 || v[3] != 8.0 {
				t.Error("Expected [0.0, 0.0, 7.0, 8.0]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * INTEGER
		fl.appendOperand(OP_BINARY_MUL, in)
		err = fl.solve(be)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 5.0 || v[1] != 12.0 || v[2] != 21.0 || v[3] != 32.0 {
				t.Error("Expected [0.0, 0.0, 21.0, 0.0]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * FLOAT
		fl.appendOperand(OP_BINARY_MUL, fl)
		err = fl.solve(be)

		if err != nil {
			t.Error(err)
		} else if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 25.0 || v[1] != 36.0 || v[2] != 49.0 || v[3] != 64.0 {
				t.Error("Expected [25.0, 36.0, 49.0, 64.0]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * STRING
		fl.appendOperand(OP_BINARY_MUL, st)
		err = fl.solve(be)

		if err == nil || err.Error() != "binary * operator not implemented for []float64 and []string" {
			t.Error(err)
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING * BOOL
		st.appendOperand(OP_BINARY_MUL, b2)
		err = st.solve(be)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "" || v[1] != "" || v[2] != "c" || v[3] != "d" {
				t.Error("Expected [\"\", \"\", \"c\", \"d\"]")
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * INTEGER
		st.appendOperand(OP_BINARY_MUL, in)
		err = st.solve(be)

		if err != nil {
			t.Error(err)
		} else if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "ccc" || v[3] != "dddd" {
				t.Error("Expected [\"a\", \"bb\", \"ccc\", \"dddd\"]")
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * FLOAT
		st.appendOperand(OP_BINARY_MUL, fl)
		err = st.solve(be)

		if err == nil || err.Error() != "binary * operator not implemented for []string and []float64" {
			t.Error(err)
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * STRING
		st.appendOperand(OP_BINARY_MUL, st)
		err = st.solve(be)

		if err == nil || err.Error() != "binary * operator not implemented for []string and []string" {
			t.Error(err)
		}
	}
}

func TestDiv(t *testing.T) {

	var err error

	b1 := newPInternTerm([]bool{true, false, true, false})
	b2 := newPInternTerm([]bool{false, false, true, true})
	in := newPInternTerm([]int{1, 2, 3, 4})
	fl := newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL / BOOL
		b1.appendOperand(OP_BINARY_DIV, b2)
		err = b1.solve(be)

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
		b1.appendOperand(OP_BINARY_DIV, in)
		err = b1.solve(be)

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
		b1.appendOperand(OP_BINARY_DIV, fl)
		err = b1.solve(be)

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
		b1.appendOperand(OP_BINARY_DIV, st)
		err = b1.solve(be)

		if err == nil || err.Error() != "binary / operator not implemented for []bool and []string" {
			t.Error(err)
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})
	}

	// INTEGER
	{
		// INTEGER / BOOL
		in.appendOperand(OP_BINARY_DIV, b1)
		err = in.solve(be)

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
		in.appendOperand(OP_BINARY_DIV, in)
		err = in.solve(be)

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
		in.appendOperand(OP_BINARY_DIV, fl)
		err = in.solve(be)

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
		in.appendOperand(OP_BINARY_DIV, st)
		err = in.solve(be)

		if err == nil || err.Error() != "binary / operator not implemented for []int and []string" {
			t.Error(err)
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT / BOOL
		fl.appendOperand(OP_BINARY_DIV, b1)
		err = fl.solve(be)

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
		fl.appendOperand(OP_BINARY_DIV, in)
		err = fl.solve(be)

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
		fl.appendOperand(OP_BINARY_DIV, fl)
		err = fl.solve(be)

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
		fl.appendOperand(OP_BINARY_DIV, st)
		err = fl.solve(be)

		if err == nil || err.Error() != "binary / operator not implemented for []float64 and []string" {
			t.Error(err)
		}
	}
}
