package gandalff

import (
	"math"
	"testing"
)

var SIZE = 1_000_000

type MyInterface interface {
	Times2() MyInterface
	PlusOne() MyInterface
}

type MyStruct struct {
	b bool
	s string
	v []float64
}

func (s MyStruct) Mod7() MyStruct {
	for i := range s.v {
		s.v[i] = math.Mod(s.v[i], 7.0)
	}

	s.b = true
	s.s += "Mod7"
	return s
}

func (s MyStruct) Times3() MyStruct {
	for i := range s.v {
		s.v[i] *= 3
	}

	s.b = true
	s.s += "Times3"
	return s
}

func (s MyStruct) PlusOne() MyStruct {
	for i := range s.v {
		s.v[i] += 1
	}

	s.b = true
	s.s += "PlusOne"
	return s
}

type MyStructPtr struct {
	b bool
	s string
	v *[]float64
}

func (s MyStructPtr) Mod7() MyStructPtr {
	v := *s.v
	for i := range v {
		v[i] = math.Mod(v[i], 7.0)
	}

	s.b = true
	s.s += "Mod7"
	s.v = &v
	return s
}

func (s MyStructPtr) Times3() MyStructPtr {
	v := *s.v
	for i := range v {
		v[i] *= 3
	}

	s.b = true
	s.s += "Times3"
	s.v = &v
	return s
}

func (s MyStructPtr) PlusOne() MyStructPtr {
	v := *s.v
	for i := range v {
		v[i] += 1
	}

	s.b = true
	s.s += "PlusOne"
	s.v = &v
	return s
}

func Test_Gen_MyStruct(t *testing.T) {
	v := make([]float64, SIZE)

	for i := 0; i < SIZE; i++ {
		v[i] = float64(i)
	}

	res := (MyStruct{v: v}).Mod7().Times3().PlusOne().Mod7().PlusOne().Times3().PlusOne().Times3().Mod7()
	if res.b != true {
		t.Fail()
	}

	if res.s != "Mod7Times3PlusOneMod7PlusOneTimes3PlusOneTimes3Mod7" {
		t.Fail()
	}

	if res.v[0] != 0 || res.v[1] != 6 || res.v[2] != 5 || res.v[3] != 4 || res.v[4] != 3 {
		t.Fail()
	}
}

func Test_Gen_MyStructPtr(t *testing.T) {
	v := make([]float64, SIZE)

	for i := 0; i < SIZE; i++ {
		v[i] = float64(i)
	}

	res := (MyStructPtr{v: &v}).Mod7().Times3().PlusOne().Mod7().PlusOne().Times3().PlusOne().Times3().Mod7()
	if res.b != true {
		t.Fail()
	}

	if res.s != "Mod7Times3PlusOneMod7PlusOneTimes3PlusOneTimes3Mod7" {
		t.Fail()
	}

	if (*res.v)[0] != 0 || (*res.v)[1] != 6 || (*res.v)[2] != 5 || (*res.v)[3] != 4 || (*res.v)[4] != 3 {
		t.Fail()
	}
}

func Benchmark_Gen_MyStruct(b *testing.B) {
	v := make([]float64, SIZE)

	for i := 0; i < SIZE; i++ {
		v[i] = float64(i)
	}

	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(MyStruct{v: v}).Mod7().Times3().PlusOne().Mod7().PlusOne().Times3().PlusOne().Times3().Mod7()
	}
}

func Benchmark_Gen_MyStructPtr(b *testing.B) {
	v := make([]float64, SIZE)

	for i := 0; i < SIZE; i++ {
		v[i] = float64(i)
	}

	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(MyStructPtr{v: &v}).Mod7().Times3().PlusOne().Mod7().PlusOne().Times3().PlusOne().Times3().Mod7()
	}
}
