package gandalff

import "unsafe"

type SeriesPart struct {
	part map[uint64][]int
}

func HashStrPtr(ptrs *[]*string) map[uint64][]int {
	m := make(map[uint64][]int)
	for i, v := range *ptrs {
		m[*(*uint64)(unsafe.Pointer(unsafe.Pointer(v)))] = append(m[*(*uint64)(unsafe.Pointer(unsafe.Pointer(v)))], i)
	}
	return m
}

func ReHashStrPtr(ptrs *[]*string, m map[uint64][]int) map[uint64][]int {
	for i, v := range *ptrs {
		m[*(*uint64)(unsafe.Pointer(unsafe.Pointer(v)))] = append(m[*(*uint64)(unsafe.Pointer(unsafe.Pointer(v)))], i)
	}
	return m
}
