package gandalff

import (
	"fmt"
	"testing"
)

type MyInterface interface {
	Times2() MyInterface
	PlusOne() MyInterface
}

type MyStruct struct {
	v []float64
}

func (s MyStruct) Times2() MyStruct {
	for i := range s.v {
		s.v[i] *= 2
	}
	return s
}

func (s MyStruct) PlusOne() MyStruct {
	for i := range s.v {
		s.v[i] += 1
	}
	return s
}

type MyStructPtr struct {
	v *[]float64
}

func (s MyStructPtr) Times2() MyStructPtr {
	for i := range *s.v {
		(*s.v)[i] *= 2
	}
	return s
}

func (s MyStructPtr) PlusOne() MyStructPtr {
	for i := range *s.v {
		(*s.v)[i] += 1
	}
	return s
}

func Benchmark_Gen_MyStruct(b *testing.B) {

	N := 100_000
	v := make([]float64, N)

	for i := 0; i < N; i++ {
		v[i] = float64(i)
	}

	s := MyStruct{v: v}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Times2().PlusOne()
	}
	b.StopTimer()

	res := s.Times2().PlusOne()
	fmt.Println(res.v[0], res.v[1], res.v[2], res.v[3], res.v[4])
}

func Benchmark_Gen_MyStructPtr(b *testing.B) {

	N := 100_000
	v := make([]float64, N)

	for i := 0; i < N; i++ {
		v[i] = float64(i)
	}

	s := MyStructPtr{v: &v}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Times2().PlusOne()
	}
	b.StopTimer()

	res := s.Times2().PlusOne()
	fmt.Println((*res.v)[0], (*res.v)[1], (*res.v)[2], (*res.v)[3], (*res.v)[4])
}
