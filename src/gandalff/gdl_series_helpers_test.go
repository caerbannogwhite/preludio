package gandalff

import (
	"testing"
)

func Test_BinVec_Init(t *testing.T) {
	var vecA, vecB []uint8

	// 6
	vecA = __binVecFromBools([]bool{true, true, true, true, true, false})
	vecB = __binVecInit(6, true)
	__binVecSet(vecB, 5, false)
	if vecA[0] != 0x1F {
		t.Errorf("Expected 0x1F, got %x", vecA[0])
	}
	if vecB[0] != 0x1F {
		t.Errorf("Expected 0x1F, got %x", vecB[0])
	}

	vecA = __binVecFromBools([]bool{true, true, true, true, true, true})
	vecB = __binVecInit(6, true)
	if vecA[0] != 0x3F {
		t.Errorf("Expected 0x3F, got %x", vecA[0])
	}
	if vecB[0] != 0x3F {
		t.Errorf("Expected 0x3F, got %x", vecB[0])
	}

	// 7
	vecA = __binVecFromBools([]bool{true, true, true, true, true, true, false})
	vecB = __binVecInit(7, true)
	__binVecSet(vecB, 6, false)
	if vecA[0] != 0x3F {
		t.Errorf("Expected 0x6F, got %x", vecA[0])
	}
	if vecB[0] != 0x3F {
		t.Errorf("Expected 0x6F, got %x", vecB[0])
	}

	vecA = __binVecFromBools([]bool{true, true, true, true, true, true, true})
	vecB = __binVecInit(7, true)
	if vecA[0] != 0x7F {
		t.Errorf("Expected 0x7F, got %x", vecA[0])
	}
	if vecB[0] != 0x7F {
		t.Errorf("Expected 0x7F, got %x", vecB[0])
	}

	// 8
	vecA = __binVecFromBools([]bool{true, true, true, true, true, true, true, false})
	vecB = __binVecInit(8, true)
	__binVecSet(vecB, 7, false)
	if vecA[0] != 0x7F {
		t.Errorf("Expected 0x7F, got %x", vecA[0])
	}
	if vecB[0] != 0x7F {
		t.Errorf("Expected 0x7F, got %x", vecB[0])
	}

	vecA = __binVecFromBools([]bool{true, true, true, true, true, true, true, true})
	vecB = __binVecInit(8, true)
	if vecA[0] != 0xFF {
		t.Errorf("Expected 0xFF, got %x", vecA[0])
	}
	if vecB[0] != 0xFF {
		t.Errorf("Expected 0xFF, got %x", vecB[0])
	}

	// 9
	vecA = __binVecFromBools([]bool{true, true, true, true, true, true, true, true, false})
	vecB = __binVecInit(9, true)
	__binVecSet(vecB, 8, false)
	if vecA[0] != 0xFF || vecA[1] != 0x00 {
		t.Errorf("Expected 0xFF 0x00, got %x %x", vecA[0], vecA[1])
	}
	if vecB[0] != 0xFF || vecB[1] != 0x00 {
		t.Errorf("Expected 0xFF 0x00, got %x %x", vecB[0], vecB[1])
	}

	vecA = __binVecFromBools([]bool{true, true, true, true, true, true, true, true, true})
	vecB = __binVecInit(9, true)
	if vecA[0] != 0xFF || vecA[1] != 0x01 {
		t.Errorf("Expected 0xFF 0x01, got %x %x", vecA[0], vecA[1])
	}
	if vecB[0] != 0xFF || vecB[1] != 0x01 {
		t.Errorf("Expected 0xFF 0x01, got %x %x", vecB[0], vecB[1])
	}
}

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
