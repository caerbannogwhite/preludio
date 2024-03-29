package gandalff

import (
	"testing"
)

func Test_BinVec_Count(t *testing.T) {

	var v []uint8

	v = []uint8{0x00}
	if c := __binVecCount(v); c != 0 {
		t.Errorf("Expected 0, got %d", c)
	}

	v = []uint8{0x01}
	if c := __binVecCount(v); c != 1 {
		t.Errorf("Expected 1, got %d", c)
	}

	v = []uint8{0x00, 0x00, 0x00, 0x00, 0x00}
	if c := __binVecCount(v); c != 0 {
		t.Errorf("Expected 0, got %d", c)
	}

	v = []uint8{0x00, 0x01, 0x00, 0x00, 0x00}
	if c := __binVecCount(v); c != 1 {
		t.Errorf("Expected 1, got %d", c)
	}

	v = []uint8{0x00, 0x01, 0x00, 0x00, 0x01}
	if c := __binVecCount(v); c != 2 {
		t.Errorf("Expected 2, got %d", c)
	}

	v = []uint8{0x01, 0x01, 0x01, 0x01, 0x01}
	if c := __binVecCount(v); c != 5 {
		t.Errorf("Expected 5, got %d", c)
	}

	v = []uint8{0x01, 0x02, 0x04, 0x08}
	if c := __binVecCount(v); c != 4 {
		t.Errorf("Expected 4, got %d", c)
	}

	v = []uint8{0x01, 0x03, 0x07, 0x0F}
	if c := __binVecCount(v); c != 10 {
		t.Errorf("Expected 10, got %d", c)
	}
}

func Benchmark_BinVec_Count(b *testing.B) {
	v := make([]uint8, 1_000_000)
	for i := 0; i < len(v); i++ {
		v[i] = uint8(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		__binVecCount(v)
	}
}
