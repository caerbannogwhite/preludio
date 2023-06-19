package gandalff

import (
	"testing"
	"unsafe"
)

func Test_Unsafe_Pointer(t *testing.T) {
	s := "hello world"

	p1 := unsafe.Pointer(&s)
	p2 := (*int64)(unsafe.Pointer(&p1))

	i := *p2
	p4 := (*string)(unsafe.Pointer(uintptr(i)))

	if s != *p4 {
		t.Errorf("String mismatch")
	}
}
