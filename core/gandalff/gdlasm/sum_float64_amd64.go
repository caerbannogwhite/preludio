package gdlasm

import "unsafe"

//go:noescape
func _sum_float64_avx_intrinsics(vec, size, res unsafe.Pointer)

func _sum_float64_c(vec, size, res unsafe.Pointer)

func _sum_grouped_float64_c(vec, vecSize, indeces, indecesSize, res unsafe.Pointer)
