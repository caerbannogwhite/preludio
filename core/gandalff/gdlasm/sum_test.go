package gdlasm

import (
	"testing"
	"unsafe"
)

func Test_Sum_Float64C(t *testing.T) {
	var vec []float64
	var size int
	var res float64

	vec = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	size = len(vec)

	_sum_float64_c(unsafe.Pointer(&vec[0]), unsafe.Pointer(&size), unsafe.Pointer(&res))

	if res != 55 {
		t.Error("Expected 55, got ", res)
	}
}

func Test_Sum_Grouped_Float64_C(t *testing.T) {
	var vec []float64
	var vecSize int
	var indeces []int
	var indecesSize int
	var res float64

	vec = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vecSize = len(vec)
	indeces = []int{0, 1}
	indecesSize = len(indeces)

	_sum_grouped_float64_c(unsafe.Pointer(&vec[0]), unsafe.Pointer(&vecSize), unsafe.Pointer(&indeces[0]), unsafe.Pointer(&indecesSize), unsafe.Pointer(&res))

	if res != 45 {
		t.Error("Expected 45, got ", res)
	}
}
