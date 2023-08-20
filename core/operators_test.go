package preludiocore

import (
	"fmt"
	"math"
	"testing"
	"typesys"
)

var be *ByteEater

func init() {
	be = new(ByteEater).InitVM()
}

func Test_Operator_Mul(t *testing.T) {

	b1 := be.newPInternTerm([]bool{true, false, true, false})
	b2 := be.newPInternTerm([]bool{false, false, true, true})
	in := be.newPInternTerm([]int64{1, 2, 3, 4})
	fl := be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := be.newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL * BOOL
		b1.appendBinaryOperation(typesys.OP_BINARY_MUL, b2)
		if err := checkExpression(be, b1, []int64{0, 0, 1, 0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL * INTEGER
		b1.appendBinaryOperation(typesys.OP_BINARY_MUL, in)
		if err := checkExpression(be, b1, []int64{1, 0, 3, 0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL * FLOAT
		b1.appendBinaryOperation(typesys.OP_BINARY_MUL, fl)
		if err := checkExpression(be, b1, []float64{5.0, 0.0, 7.0, 0.0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL * STRING
		b1.appendBinaryOperation(typesys.OP_BINARY_MUL, st)
		if err := checkExpression(be, b1, fmt.Errorf("binary operator * not supported between Bool[4] and String[4]")); err != nil {
			t.Error(err)
		}
	}

	// INTEGER
	{
		// INTEGER * BOOL
		in.appendBinaryOperation(typesys.OP_BINARY_MUL, b2)
		if err := checkExpression(be, in, []int64{0, 0, 3, 4}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER * INTEGER
		in.appendBinaryOperation(typesys.OP_BINARY_MUL, in)
		if err := checkExpression(be, in, []int64{1, 4, 9, 16}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER * FLOAT
		in.appendBinaryOperation(typesys.OP_BINARY_MUL, fl)
		if err := checkExpression(be, in, []float64{5.0, 12.0, 21.0, 32.0}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT * BOOL
		fl.appendBinaryOperation(typesys.OP_BINARY_MUL, b2)
		if err := checkExpression(be, fl, []float64{0.0, 0.0, 7.0, 8.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * INTEGER
		fl.appendBinaryOperation(typesys.OP_BINARY_MUL, in)
		if err := checkExpression(be, fl, []float64{5.0, 12.0, 21.0, 32.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * FLOAT
		fl.appendBinaryOperation(typesys.OP_BINARY_MUL, fl)
		if err := checkExpression(be, fl, []float64{25.0, 36.0, 49.0, 64.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING * BOOL
		st.appendBinaryOperation(typesys.OP_BINARY_MUL, b2)
		if err := checkExpression(be, st, fmt.Errorf("binary operator * not supported between String[4] and Bool[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * INTEGER
		st.appendBinaryOperation(typesys.OP_BINARY_MUL, in)
		if err := checkExpression(be, st, fmt.Errorf("binary operator * not supported between String[4] and Int64[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * FLOAT
		st.appendBinaryOperation(typesys.OP_BINARY_MUL, fl)
		if err := checkExpression(be, st, fmt.Errorf("binary operator * not supported between String[4] and Float64[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * STRING
		st.appendBinaryOperation(typesys.OP_BINARY_MUL, st)
		if err := checkExpression(be, st, fmt.Errorf("binary operator * not supported between String[4] and String[4]")); err != nil {
			t.Error(err)
		}
	}
}

func Test_Operator_Div(t *testing.T) {

	b1 := be.newPInternTerm([]bool{true, false, true, false})
	b2 := be.newPInternTerm([]bool{false, false, true, true})
	in := be.newPInternTerm([]int64{1, 2, 3, 4})
	fl := be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := be.newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL / BOOL
		b1.appendBinaryOperation(typesys.OP_BINARY_DIV, b2)
		if err := checkExpression(be, b1, []float64{math.Inf(1), math.NaN(), 1.0, 0.0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL / INTEGER
		b1.appendBinaryOperation(typesys.OP_BINARY_DIV, in)
		if err := checkExpression(be, b1, []float64{1.0, 0.0, 0.3333333333333333, 0.0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL / FLOAT
		b1.appendBinaryOperation(typesys.OP_BINARY_DIV, fl)
		if err := checkExpression(be, b1, []float64{0.2, 0.0, 0.14285714285714285, 0.0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL / STRING
		b1.appendBinaryOperation(typesys.OP_BINARY_DIV, st)
		if err := checkExpression(be, b1, fmt.Errorf("binary operator / not supported between Bool[4] and String[4]")); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})
	}

	// INTEGER
	{
		// INTEGER / BOOL
		in.appendBinaryOperation(typesys.OP_BINARY_DIV, b1)
		if err := checkExpression(be, in, []float64{1.0, math.Inf(1), 3.0, math.Inf(1)}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER / INTEGER
		in.appendBinaryOperation(typesys.OP_BINARY_DIV, in)
		if err := checkExpression(be, in, []float64{1.0, 1.0, 1.0, 1.0}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER / FLOAT
		in.appendBinaryOperation(typesys.OP_BINARY_DIV, fl)
		if err := checkExpression(be, in, []float64{0.2, 0.3333333333333333, 0.42857142857142855, 0.5}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER / STRING
		in.appendBinaryOperation(typesys.OP_BINARY_DIV, st)
		if err := checkExpression(be, in, fmt.Errorf("binary operator / not supported between Int64[4] and String[4]")); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT / BOOL
		fl.appendBinaryOperation(typesys.OP_BINARY_DIV, b1)
		if err := checkExpression(be, fl, []float64{5.0, math.Inf(1), 7.0, math.Inf(1)}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT / INTEGER
		fl.appendBinaryOperation(typesys.OP_BINARY_DIV, in)
		if err := checkExpression(be, fl, []float64{5.0, 3.0, 2.3333333333333335, 2.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT / FLOAT
		fl.appendBinaryOperation(typesys.OP_BINARY_DIV, fl)
		if err := checkExpression(be, fl, []float64{1.0, 1.0, 1.0, 1.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT / STRING
		fl.appendBinaryOperation(typesys.OP_BINARY_DIV, st)
		if err := checkExpression(be, fl, fmt.Errorf("binary operator / not supported between Float64[4] and String[4]")); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING / BOOL
		st.appendBinaryOperation(typesys.OP_BINARY_DIV, b1)
		if err := checkExpression(be, st, fmt.Errorf("binary operator / not supported between String[4] and Bool[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING / INTEGER
		st.appendBinaryOperation(typesys.OP_BINARY_DIV, in)
		if err := checkExpression(be, st, fmt.Errorf("binary operator / not supported between String[4] and Int64[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING / FLOAT
		st.appendBinaryOperation(typesys.OP_BINARY_DIV, fl)
		if err := checkExpression(be, st, fmt.Errorf("binary operator / not supported between String[4] and Float64[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING / STRING
		st.appendBinaryOperation(typesys.OP_BINARY_DIV, st)
		if err := checkExpression(be, st, fmt.Errorf("binary operator / not supported between String[4] and String[4]")); err != nil {
			t.Error(err)
		}
	}
}

func Test_Operator_Add(t *testing.T) {

	b1 := be.newPInternTerm([]bool{true, false, true, false})
	b2 := be.newPInternTerm([]bool{false, false, true, true})
	in := be.newPInternTerm([]int64{1, 2, 3, 4})
	fl := be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := be.newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL + BOOL
		b1.appendBinaryOperation(typesys.OP_BINARY_ADD, b2)
		if err := checkExpression(be, b1, []int64{1, 0, 2, 1}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL + INTEGER
		b1.appendBinaryOperation(typesys.OP_BINARY_ADD, in)
		if err := checkExpression(be, b1, []int64{2, 2, 4, 4}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL + FLOAT
		b1.appendBinaryOperation(typesys.OP_BINARY_ADD, fl)
		if err := checkExpression(be, b1, []float64{6.0, 6.0, 8.0, 8.0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL + STRING
		b1.appendBinaryOperation(typesys.OP_BINARY_ADD, st)
		if err := checkExpression(be, b1, []string{"truea", "falseb", "truec", "falsed"}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})
	}

	// INTEGER
	{
		// INTEGER + BOOL
		in.appendBinaryOperation(typesys.OP_BINARY_ADD, b1)
		if err := checkExpression(be, in, []int64{2, 2, 4, 4}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER + INTEGER
		in.appendBinaryOperation(typesys.OP_BINARY_ADD, in)
		if err := checkExpression(be, in, []int64{2, 4, 6, 8}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER + FLOAT
		in.appendBinaryOperation(typesys.OP_BINARY_ADD, fl)
		if err := checkExpression(be, in, []float64{6.0, 8.0, 10.0, 12.0}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER + STRING
		in.appendBinaryOperation(typesys.OP_BINARY_ADD, st)
		if err := checkExpression(be, in, []string{"1a", "2b", "3c", "4d"}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT + BOOL
		fl.appendBinaryOperation(typesys.OP_BINARY_ADD, b1)
		if err := checkExpression(be, fl, []float64{6.0, 6.0, 8.0, 8.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT + INTEGER
		fl.appendBinaryOperation(typesys.OP_BINARY_ADD, in)
		if err := checkExpression(be, fl, []float64{6.0, 8.0, 10.0, 12.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT + FLOAT
		fl.appendBinaryOperation(typesys.OP_BINARY_ADD, fl)
		if err := checkExpression(be, fl, []float64{10.0, 12.0, 14.0, 16.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT + STRING
		fl.appendBinaryOperation(typesys.OP_BINARY_ADD, st)
		if err := checkExpression(be, fl, []string{"5a", "6b", "7c", "8d"}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING + BOOL
		st.appendBinaryOperation(typesys.OP_BINARY_ADD, b1)
		if err := checkExpression(be, st, []string{"atrue", "bfalse", "ctrue", "dfalse"}); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING + INTEGER
		st.appendBinaryOperation(typesys.OP_BINARY_ADD, in)
		if err := checkExpression(be, st, []string{"a1", "b2", "c3", "d4"}); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING + FLOAT
		st.appendBinaryOperation(typesys.OP_BINARY_ADD, fl)
		if err := checkExpression(be, st, []string{"a5", "b6", "c7", "d8"}); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING + STRING
		st.appendBinaryOperation(typesys.OP_BINARY_ADD, st)
		if err := checkExpression(be, st, []string{"aa", "bb", "cc", "dd"}); err != nil {
			t.Error(err)
		}
	}
}

func Test_Operator_Sub(t *testing.T) {

	b1 := be.newPInternTerm([]bool{true, false, true, false})
	b2 := be.newPInternTerm([]bool{false, false, true, true})
	in := be.newPInternTerm([]int64{1, 2, 3, 4})
	fl := be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := be.newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL - BOOL
		b1.appendBinaryOperation(typesys.OP_BINARY_SUB, b2)
		if err := checkExpression(be, b1, []int64{1, 0, 0, -1}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL - INTEGER
		b1.appendBinaryOperation(typesys.OP_BINARY_SUB, in)
		if err := checkExpression(be, b1, []int64{0, -2, -2, -4}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL - FLOAT
		b1.appendBinaryOperation(typesys.OP_BINARY_SUB, fl)
		if err := checkExpression(be, b1, []float64{-4.0, -6.0, -6.0, -8.0}); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})

		// BOOL - STRING
		b1.appendBinaryOperation(typesys.OP_BINARY_SUB, st)
		if err := checkExpression(be, b1, fmt.Errorf("binary operator - not supported between Bool[4] and String[4]")); err != nil {
			t.Error(err)
		}

		// reset b1
		b1 = be.newPInternTerm([]bool{true, false, true, false})
	}

	// INTEGER
	{
		// INTEGER - BOOL
		in.appendBinaryOperation(typesys.OP_BINARY_SUB, b1)
		if err := checkExpression(be, in, []int64{0, 2, 2, 4}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER - INTEGER
		in.appendBinaryOperation(typesys.OP_BINARY_SUB, in)
		if err := checkExpression(be, in, []int64{0, 0, 0, 0}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER - FLOAT
		in.appendBinaryOperation(typesys.OP_BINARY_SUB, fl)
		if err := checkExpression(be, in, []float64{-4.0, -4.0, -4.0, -4.0}); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})

		// INTEGER - STRING
		in.appendBinaryOperation(typesys.OP_BINARY_SUB, st)
		if err := checkExpression(be, in, fmt.Errorf("binary operator - not supported between Int64[4] and String[4]")); err != nil {
			t.Error(err)
		}

		// reset in
		in = be.newPInternTerm([]int64{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT - BOOL
		fl.appendBinaryOperation(typesys.OP_BINARY_SUB, b1)
		if err := checkExpression(be, fl, []float64{4.0, 6.0, 6.0, 8.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT - INTEGER
		fl.appendBinaryOperation(typesys.OP_BINARY_SUB, in)
		if err := checkExpression(be, fl, []float64{4.0, 4.0, 4.0, 4.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT - FLOAT
		fl.appendBinaryOperation(typesys.OP_BINARY_SUB, fl)
		if err := checkExpression(be, fl, []float64{0.0, 0.0, 0.0, 0.0}); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT - STRING
		fl.appendBinaryOperation(typesys.OP_BINARY_SUB, st)
		if err := checkExpression(be, fl, fmt.Errorf("binary operator - not supported between Float64[4] and String[4]")); err != nil {
			t.Error(err)
		}

		// reset fl
		fl = be.newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING - BOOL
		st.appendBinaryOperation(typesys.OP_BINARY_SUB, b1)
		if err := checkExpression(be, st, fmt.Errorf("binary operator - not supported between String[4] and Bool[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING - INTEGER
		st.appendBinaryOperation(typesys.OP_BINARY_SUB, in)
		if err := checkExpression(be, st, fmt.Errorf("binary operator - not supported between String[4] and Int64[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING - FLOAT
		st.appendBinaryOperation(typesys.OP_BINARY_SUB, fl)
		if err := checkExpression(be, st, fmt.Errorf("binary operator - not supported between String[4] and Float64[4]")); err != nil {
			t.Error(err)
		}

		// reset st
		st = be.newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING - STRING
		st.appendBinaryOperation(typesys.OP_BINARY_SUB, st)
		if err := checkExpression(be, st, fmt.Errorf("binary operator - not supported between String[4] and String[4]")); err != nil {
			t.Error(err)
		}
	}
}
