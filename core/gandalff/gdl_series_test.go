package gandalff

import (
	"fmt"
	"testing"
	"typesys"
	"unsafe"
)

func Test_GDLSeries(t *testing.T) {

	s := NewGDLSeries("test", typesys.BoolType, true, false, []bool{true, false, true, false, true, false, true, false, true, false}, nil)

	r := s.Append(true).
		Append([]NullableBool{{true, true}, {true, false}}).
		FilterByMask([]bool{true, false, true, false, true, false, true, false, true, false, true, true, false})

	if e, ok := r.(GDLSeriesError); ok {
		t.Errorf("Expected a series, got an error: %s", e.Error())
	}

	if r.Len() != 7 {
		t.Errorf("Expected length of 7, got %d", r.Len())
	}
}

func TestXxx(t *testing.T) {

	// data1 := []string{"A", "A", "B", "B", "C", "C", "D", "D", "E", "E"}
	// data2 := []string{"!", "@", "0", "0", "0", "^", "&", "*", "1", "1"}

	// pool := NewStringPool()

	// s1 := NewGDLSeriesString("test1", false, data1, pool)
	// s2 := NewGDLSeriesString("test2", false, data2, pool)

	// p1 := s1.(GDLSeriesString).__getDataPtr()
	// p2 := s2.(GDLSeriesString).__getDataPtr()

	// m1 := HashStringPtrVec(p1)

	// m2 := CombineStringPtrVec(m1, p2)

	// fmt.Println(m1)
	// fmt.Println(m2)

	v1 := float64(1.61803398874989484820458683436563811772030917980576286213544862270526046281890)
	v2 := float64(1.61803398874989484820458683436563811772030917980576286213544862270526046281890)

	fmt.Println(uint64(v1))
	fmt.Println(*(*uint64)(unsafe.Pointer(&v1)))
	fmt.Println(*(*uint64)(unsafe.Pointer(&v2)))
}
