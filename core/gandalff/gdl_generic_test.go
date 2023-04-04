package gandalff

import (
	"sync"
	"testing"
)

const ___1K = 1_000
const __10K = 10_000
const _100K = 100_000
const ___1M = 1_000_000
const __10M = 10_000_000
const _100M = 100_000_000

func BenchmarkGen_Perf___1K____Baseline(b *testing.B) {

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < ___1K; j++ {
			v1[j] *= v2[j]
		}
	}
}

func BenchmarkGen_Perf___1K__2_Goroutines(b *testing.B) {

	GO_ROUTINES := 2

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, ___1K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1K__4_Goroutines(b *testing.B) {

	GO_ROUTINES := 4

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, ___1K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1K__8_Goroutines(b *testing.B) {

	GO_ROUTINES := 8

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, ___1K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1K_16_Goroutines(b *testing.B) {

	GO_ROUTINES := 16

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, splitSize*8)

		// 9
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*8, splitSize*9)

		// 10
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*9, splitSize*10)

		// 11
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*10, splitSize*11)

		// 12
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*11, splitSize*12)

		// 13
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*12, splitSize*13)

		// 14
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*13, splitSize*14)

		// 15
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*14, splitSize*15)

		// 16
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*15, ___1K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1K__2_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1K - ___1K%2
		for j := 0; j < upper; j += 2 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
		}

		if ___1K%2 > 0 {
			v1[___1K-1] *= v2[___1K-1]
		}
	}
}

func BenchmarkGen_Perf___1K__4_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1K - ___1K%4
		for j := 0; j < upper; j += 4 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
		}

		switch ___1K % 4 {
		case 3:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
		case 2:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
		case 1:
			v1[___1K-1] *= v2[___1K-1]
		}
	}
}

func BenchmarkGen_Perf___1K__8_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1K - ___1K%8
		for j := 0; j < upper; j += 8 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
		}

		switch ___1K % 8 {
		case 7:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
		case 6:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
		case 5:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
		case 4:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
		case 3:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
		case 2:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
		case 1:
			v1[___1K-1] *= v2[___1K-1]
		}
	}
}

func BenchmarkGen_Perf___1K_16_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1K)
	v2 := make([]float64, ___1K)

	for i := 0; i < ___1K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1K - ___1K%16
		for j := 0; j < upper; j += 16 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
			v1[j+8] *= v2[j+8]
			v1[j+9] *= v2[j+9]
			v1[j+10] *= v2[j+10]
			v1[j+11] *= v2[j+11]
			v1[j+12] *= v2[j+12]
			v1[j+13] *= v2[j+13]
			v1[j+14] *= v2[j+14]
			v1[j+15] *= v2[j+15]
		}

		switch ___1K % 16 {
		case 15:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
			v1[___1K-9] *= v2[___1K-9]
			v1[___1K-10] *= v2[___1K-10]
			v1[___1K-11] *= v2[___1K-11]
			v1[___1K-12] *= v2[___1K-12]
			v1[___1K-13] *= v2[___1K-13]
			v1[___1K-14] *= v2[___1K-14]
			v1[___1K-15] *= v2[___1K-15]
		case 14:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
			v1[___1K-9] *= v2[___1K-9]
			v1[___1K-10] *= v2[___1K-10]
			v1[___1K-11] *= v2[___1K-11]
			v1[___1K-12] *= v2[___1K-12]
			v1[___1K-13] *= v2[___1K-13]
			v1[___1K-14] *= v2[___1K-14]
		case 13:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
			v1[___1K-9] *= v2[___1K-9]
			v1[___1K-10] *= v2[___1K-10]
			v1[___1K-11] *= v2[___1K-11]
			v1[___1K-12] *= v2[___1K-12]
			v1[___1K-13] *= v2[___1K-13]
		case 12:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
			v1[___1K-9] *= v2[___1K-9]
			v1[___1K-10] *= v2[___1K-10]
			v1[___1K-11] *= v2[___1K-11]
			v1[___1K-12] *= v2[___1K-12]
		case 11:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
			v1[___1K-9] *= v2[___1K-9]
			v1[___1K-10] *= v2[___1K-10]
			v1[___1K-11] *= v2[___1K-11]
		case 10:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
			v1[___1K-9] *= v2[___1K-9]
			v1[___1K-10] *= v2[___1K-10]
		case 9:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
			v1[___1K-9] *= v2[___1K-9]
		case 8:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
			v1[___1K-8] *= v2[___1K-8]
		case 7:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
			v1[___1K-7] *= v2[___1K-7]
		case 6:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
			v1[___1K-6] *= v2[___1K-6]
		case 5:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
			v1[___1K-5] *= v2[___1K-5]
		case 4:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
			v1[___1K-4] *= v2[___1K-4]
		case 3:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
			v1[___1K-3] *= v2[___1K-3]
		case 2:
			v1[___1K-1] *= v2[___1K-1]
			v1[___1K-2] *= v2[___1K-2]
		case 1:
			v1[___1K-1] *= v2[___1K-1]
		}
	}
}

func BenchmarkGen_Perf__10K____Baseline(b *testing.B) {

	v1 := make([]float64, __10K)
	v2 := make([]float64, __10K)

	for i := 0; i < __10K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < __10K; j++ {
			v1[j] *= v2[j]
		}
	}
}

func BenchmarkGen_Perf__10K__2_Goroutines(b *testing.B) {

	GO_ROUTINES := 2

	v1 := make([]float64, __10K)
	v2 := make([]float64, __10K)

	for i := 0; i < __10K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, __10K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf__10K__4_Goroutines(b *testing.B) {

	GO_ROUTINES := 4

	v1 := make([]float64, __10K)
	v2 := make([]float64, __10K)

	for i := 0; i < __10K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, __10K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf__10K__8_Goroutines(b *testing.B) {

	GO_ROUTINES := 8

	v1 := make([]float64, __10K)
	v2 := make([]float64, __10K)

	for i := 0; i < __10K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, __10K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf__10K_16_Goroutines(b *testing.B) {

	GO_ROUTINES := 16

	v1 := make([]float64, __10K)
	v2 := make([]float64, __10K)

	for i := 0; i < __10K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, splitSize*8)

		// 9
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*8, splitSize*9)

		// 10
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*9, splitSize*10)

		// 11
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*10, splitSize*11)

		// 12
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*11, splitSize*12)

		// 13
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*12, splitSize*13)

		// 14
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*13, splitSize*14)

		// 15
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*14, splitSize*15)

		// 16
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*15, __10K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100K____Baseline(b *testing.B) {

	v1 := make([]float64, _100K)
	v2 := make([]float64, _100K)

	for i := 0; i < _100K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < _100K; j++ {
			v1[j] *= v2[j]
		}
	}
}

func BenchmarkGen_Perf_100K__2_Goroutines(b *testing.B) {

	GO_ROUTINES := 2

	v1 := make([]float64, _100K)
	v2 := make([]float64, _100K)

	for i := 0; i < _100K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, _100K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100K__4_Goroutines(b *testing.B) {

	GO_ROUTINES := 4

	v1 := make([]float64, _100K)
	v2 := make([]float64, _100K)

	for i := 0; i < _100K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, _100K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100K__8_Goroutines(b *testing.B) {

	GO_ROUTINES := 8

	v1 := make([]float64, _100K)
	v2 := make([]float64, _100K)

	for i := 0; i < _100K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, _100K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100K_16_Goroutines(b *testing.B) {

	GO_ROUTINES := 16

	v1 := make([]float64, _100K)
	v2 := make([]float64, _100K)

	for i := 0; i < _100K; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100K - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100K / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, splitSize*8)

		// 9
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*8, splitSize*9)

		// 10
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*9, splitSize*10)

		// 11
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*10, splitSize*11)

		// 12
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*11, splitSize*12)

		// 13
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*12, splitSize*13)

		// 14
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*13, splitSize*14)

		// 15
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*14, splitSize*15)

		// 16
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*15, _100K)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1M____Baseline(b *testing.B) {

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < ___1M; j++ {
			v1[j] *= v2[j]
		}
	}
}

func BenchmarkGen_Perf___1M__2_Goroutines(b *testing.B) {

	GO_ROUTINES := 2

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, ___1M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1M__4_Goroutines(b *testing.B) {

	GO_ROUTINES := 4

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, ___1M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1M__8_Goroutines(b *testing.B) {

	GO_ROUTINES := 8

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, ___1M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1M_16_Goroutines(b *testing.B) {

	GO_ROUTINES := 16

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := ___1M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, splitSize*8)

		// 9
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*8, splitSize*9)

		// 10
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*9, splitSize*10)

		// 11
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*10, splitSize*11)

		// 12
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*11, splitSize*12)

		// 13
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*12, splitSize*13)

		// 14
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*13, splitSize*14)

		// 15
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*14, splitSize*15)

		// 16
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*15, ___1M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf___1M__2_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1M - ___1M%2
		for j := 0; j < upper; j += 2 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
		}

		if ___1M%2 > 0 {
			v1[___1M-1] *= v2[___1M-1]
		}
	}
}

func BenchmarkGen_Perf___1M__4_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1M - ___1M%4
		for j := 0; j < upper; j += 4 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
		}

		switch ___1M % 4 {
		case 3:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
		case 2:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
		case 1:
			v1[___1M-1] *= v2[___1M-1]
		}
	}
}

func BenchmarkGen_Perf___1M__8_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1M - ___1M%8
		for j := 0; j < upper; j += 8 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
		}

		switch ___1M % 8 {
		case 7:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
		case 6:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
		case 5:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
		case 4:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
		case 3:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
		case 2:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
		case 1:
			v1[___1M-1] *= v2[___1M-1]
		}
	}
}

func BenchmarkGen_Perf___1M_16_Unrolling(b *testing.B) {

	v1 := make([]float64, ___1M)
	v2 := make([]float64, ___1M)

	for i := 0; i < ___1M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(___1M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := ___1M - ___1M%16
		for j := 0; j < upper; j += 16 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
			v1[j+8] *= v2[j+8]
			v1[j+9] *= v2[j+9]
			v1[j+10] *= v2[j+10]
			v1[j+11] *= v2[j+11]
			v1[j+12] *= v2[j+12]
			v1[j+13] *= v2[j+13]
			v1[j+14] *= v2[j+14]
			v1[j+15] *= v2[j+15]
		}

		switch ___1M % 16 {
		case 15:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
			v1[___1M-9] *= v2[___1M-9]
			v1[___1M-10] *= v2[___1M-10]
			v1[___1M-11] *= v2[___1M-11]
			v1[___1M-12] *= v2[___1M-12]
			v1[___1M-13] *= v2[___1M-13]
			v1[___1M-14] *= v2[___1M-14]
			v1[___1M-15] *= v2[___1M-15]
		case 14:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
			v1[___1M-9] *= v2[___1M-9]
			v1[___1M-10] *= v2[___1M-10]
			v1[___1M-11] *= v2[___1M-11]
			v1[___1M-12] *= v2[___1M-12]
			v1[___1M-13] *= v2[___1M-13]
			v1[___1M-14] *= v2[___1M-14]
		case 13:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
			v1[___1M-9] *= v2[___1M-9]
			v1[___1M-10] *= v2[___1M-10]
			v1[___1M-11] *= v2[___1M-11]
			v1[___1M-12] *= v2[___1M-12]
			v1[___1M-13] *= v2[___1M-13]
		case 12:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
			v1[___1M-9] *= v2[___1M-9]
			v1[___1M-10] *= v2[___1M-10]
			v1[___1M-11] *= v2[___1M-11]
			v1[___1M-12] *= v2[___1M-12]
		case 11:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
			v1[___1M-9] *= v2[___1M-9]
			v1[___1M-10] *= v2[___1M-10]
			v1[___1M-11] *= v2[___1M-11]
		case 10:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
			v1[___1M-9] *= v2[___1M-9]
			v1[___1M-10] *= v2[___1M-10]
		case 9:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
			v1[___1M-9] *= v2[___1M-9]
		case 8:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
			v1[___1M-8] *= v2[___1M-8]
		case 7:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
			v1[___1M-7] *= v2[___1M-7]
		case 6:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
			v1[___1M-6] *= v2[___1M-6]
		case 5:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
			v1[___1M-5] *= v2[___1M-5]
		case 4:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
			v1[___1M-4] *= v2[___1M-4]
		case 3:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
			v1[___1M-3] *= v2[___1M-3]
		case 2:
			v1[___1M-1] *= v2[___1M-1]
			v1[___1M-2] *= v2[___1M-2]
		case 1:
			v1[___1M-1] *= v2[___1M-1]
		}
	}
}

func BenchmarkGen_Perf__10M____Baseline(b *testing.B) {

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < __10M; j++ {
			v1[j] *= v2[j]
		}
	}
}

func BenchmarkGen_Perf__10M__2_Goroutines(b *testing.B) {

	GO_ROUTINES := 2

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, __10M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf__10M__4_Goroutines(b *testing.B) {

	GO_ROUTINES := 4

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, __10M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf__10M__8_Goroutines(b *testing.B) {

	GO_ROUTINES := 8

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, __10M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf__10M_16_Goroutines(b *testing.B) {

	GO_ROUTINES := 16

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := __10M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, splitSize*8)

		// 9
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*8, splitSize*9)

		// 10
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*9, splitSize*10)

		// 11
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*10, splitSize*11)

		// 12
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*11, splitSize*12)

		// 13
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*12, splitSize*13)

		// 14
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*13, splitSize*14)

		// 15
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*14, splitSize*15)

		// 16
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*15, __10M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf__10M__2_Unrolling(b *testing.B) {

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := __10M - __10M%2
		for j := 0; j < upper; j += 2 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
		}

		if __10M%2 > 0 {
			v1[__10M-1] *= v2[__10M-1]
		}
	}
}

func BenchmarkGen_Perf__10M__4_Unrolling(b *testing.B) {

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := __10M - __10M%4
		for j := 0; j < upper; j += 4 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
		}

		switch __10M % 4 {
		case 3:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
		case 2:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
		case 1:
			v1[__10M-1] *= v2[__10M-1]
		}
	}
}

func BenchmarkGen_Perf__10M__8_Unrolling(b *testing.B) {

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := __10M - __10M%8
		for j := 0; j < upper; j += 8 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
		}

		switch __10M % 8 {
		case 7:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
		case 6:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
		case 5:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
		case 4:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
		case 3:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
		case 2:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
		case 1:
			v1[__10M-1] *= v2[__10M-1]
		}
	}
}

func BenchmarkGen_Perf__10M_16_Unrolling(b *testing.B) {

	v1 := make([]float64, __10M)
	v2 := make([]float64, __10M)

	for i := 0; i < __10M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(__10M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := __10M - __10M%16
		for j := 0; j < upper; j += 16 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
			v1[j+8] *= v2[j+8]
			v1[j+9] *= v2[j+9]
			v1[j+10] *= v2[j+10]
			v1[j+11] *= v2[j+11]
			v1[j+12] *= v2[j+12]
			v1[j+13] *= v2[j+13]
			v1[j+14] *= v2[j+14]
			v1[j+15] *= v2[j+15]
		}

		switch __10M % 16 {
		case 15:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
			v1[__10M-9] *= v2[__10M-9]
			v1[__10M-10] *= v2[__10M-10]
			v1[__10M-11] *= v2[__10M-11]
			v1[__10M-12] *= v2[__10M-12]
			v1[__10M-13] *= v2[__10M-13]
			v1[__10M-14] *= v2[__10M-14]
			v1[__10M-15] *= v2[__10M-15]
		case 14:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
			v1[__10M-9] *= v2[__10M-9]
			v1[__10M-10] *= v2[__10M-10]
			v1[__10M-11] *= v2[__10M-11]
			v1[__10M-12] *= v2[__10M-12]
			v1[__10M-13] *= v2[__10M-13]
			v1[__10M-14] *= v2[__10M-14]
		case 13:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
			v1[__10M-9] *= v2[__10M-9]
			v1[__10M-10] *= v2[__10M-10]
			v1[__10M-11] *= v2[__10M-11]
			v1[__10M-12] *= v2[__10M-12]
			v1[__10M-13] *= v2[__10M-13]
		case 12:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
			v1[__10M-9] *= v2[__10M-9]
			v1[__10M-10] *= v2[__10M-10]
			v1[__10M-11] *= v2[__10M-11]
			v1[__10M-12] *= v2[__10M-12]
		case 11:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
			v1[__10M-9] *= v2[__10M-9]
			v1[__10M-10] *= v2[__10M-10]
			v1[__10M-11] *= v2[__10M-11]
		case 10:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
			v1[__10M-9] *= v2[__10M-9]
			v1[__10M-10] *= v2[__10M-10]
		case 9:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
			v1[__10M-9] *= v2[__10M-9]
		case 8:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
			v1[__10M-8] *= v2[__10M-8]
		case 7:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
			v1[__10M-7] *= v2[__10M-7]
		case 6:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
			v1[__10M-6] *= v2[__10M-6]
		case 5:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
			v1[__10M-5] *= v2[__10M-5]
		case 4:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
			v1[__10M-4] *= v2[__10M-4]
		case 3:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
			v1[__10M-3] *= v2[__10M-3]
		case 2:
			v1[__10M-1] *= v2[__10M-1]
			v1[__10M-2] *= v2[__10M-2]
		case 1:
			v1[__10M-1] *= v2[__10M-1]
		}
	}
}

func BenchmarkGen_Perf_100M____Baseline(b *testing.B) {

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < _100M; j++ {
			v1[j] *= v2[j]
		}
	}
}

func BenchmarkGen_Perf_100M__2_Goroutines(b *testing.B) {

	GO_ROUTINES := 2

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, _100M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100M__4_Goroutines(b *testing.B) {

	GO_ROUTINES := 4

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, _100M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100M__8_Goroutines(b *testing.B) {

	GO_ROUTINES := 8

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, _100M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100M_16_Goroutines(b *testing.B) {

	GO_ROUTINES := 16

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		splitSize := _100M / GO_ROUTINES
		wt := sync.WaitGroup{}
		wt.Add(GO_ROUTINES)

		// 1
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(0, splitSize)

		// 2
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize, splitSize*2)

		// 3
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*2, splitSize*3)

		// 4
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*3, splitSize*4)

		// 5
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*4, splitSize*5)

		// 6
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*5, splitSize*6)

		// 7
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*6, splitSize*7)

		// 8
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*7, splitSize*8)

		// 9
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*8, splitSize*9)

		// 10
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*9, splitSize*10)

		// 11
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*10, splitSize*11)

		// 12
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*11, splitSize*12)

		// 13
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*12, splitSize*13)

		// 14
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*13, splitSize*14)

		// 15
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*14, splitSize*15)

		// 16
		go func(start, end int) {
			for k := start; k < end; k++ {
				v1[k] *= v2[k]
			}
			wt.Done()
		}(splitSize*15, _100M)

		wt.Wait()
	}
}

func BenchmarkGen_Perf_100M__2_Unrolling(b *testing.B) {

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := _100M - _100M%2
		for j := 0; j < upper; j += 2 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
		}

		if _100M%2 > 0 {
			v1[_100M-1] *= v2[_100M-1]
		}
	}
}

func BenchmarkGen_Perf_100M__4_Unrolling(b *testing.B) {

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := _100M - _100M%4
		for j := 0; j < upper; j += 4 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
		}

		switch _100M % 4 {
		case 3:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
		case 2:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
		case 1:
			v1[_100M-1] *= v2[_100M-1]
		}
	}
}

func BenchmarkGen_Perf_100M__8_Unrolling(b *testing.B) {

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := _100M - _100M%8
		for j := 0; j < upper; j += 8 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
		}

		switch _100M % 8 {
		case 7:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
		case 6:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
		case 5:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
		case 4:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
		case 3:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
		case 2:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
		case 1:
			v1[_100M-1] *= v2[_100M-1]
		}
	}
}

func BenchmarkGen_Perf_100M_16_Unrolling(b *testing.B) {

	v1 := make([]float64, _100M)
	v2 := make([]float64, _100M)

	for i := 0; i < _100M; i++ {
		v1[i] = float64(i)
		v2[i] = float64(_100M - i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		upper := _100M - _100M%16
		for j := 0; j < upper; j += 16 {
			v1[j] *= v2[j]
			v1[j+1] *= v2[j+1]
			v1[j+2] *= v2[j+2]
			v1[j+3] *= v2[j+3]
			v1[j+4] *= v2[j+4]
			v1[j+5] *= v2[j+5]
			v1[j+6] *= v2[j+6]
			v1[j+7] *= v2[j+7]
			v1[j+8] *= v2[j+8]
			v1[j+9] *= v2[j+9]
			v1[j+10] *= v2[j+10]
			v1[j+11] *= v2[j+11]
			v1[j+12] *= v2[j+12]
			v1[j+13] *= v2[j+13]
			v1[j+14] *= v2[j+14]
			v1[j+15] *= v2[j+15]
		}

		switch _100M % 16 {
		case 15:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
			v1[_100M-9] *= v2[_100M-9]
			v1[_100M-10] *= v2[_100M-10]
			v1[_100M-11] *= v2[_100M-11]
			v1[_100M-12] *= v2[_100M-12]
			v1[_100M-13] *= v2[_100M-13]
			v1[_100M-14] *= v2[_100M-14]
			v1[_100M-15] *= v2[_100M-15]
		case 14:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
			v1[_100M-9] *= v2[_100M-9]
			v1[_100M-10] *= v2[_100M-10]
			v1[_100M-11] *= v2[_100M-11]
			v1[_100M-12] *= v2[_100M-12]
			v1[_100M-13] *= v2[_100M-13]
			v1[_100M-14] *= v2[_100M-14]
		case 13:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
			v1[_100M-9] *= v2[_100M-9]
			v1[_100M-10] *= v2[_100M-10]
			v1[_100M-11] *= v2[_100M-11]
			v1[_100M-12] *= v2[_100M-12]
			v1[_100M-13] *= v2[_100M-13]
		case 12:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
			v1[_100M-9] *= v2[_100M-9]
			v1[_100M-10] *= v2[_100M-10]
			v1[_100M-11] *= v2[_100M-11]
			v1[_100M-12] *= v2[_100M-12]
		case 11:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
			v1[_100M-9] *= v2[_100M-9]
			v1[_100M-10] *= v2[_100M-10]
			v1[_100M-11] *= v2[_100M-11]
		case 10:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
			v1[_100M-9] *= v2[_100M-9]
			v1[_100M-10] *= v2[_100M-10]
		case 9:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
			v1[_100M-9] *= v2[_100M-9]
		case 8:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
			v1[_100M-8] *= v2[_100M-8]
		case 7:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
			v1[_100M-7] *= v2[_100M-7]
		case 6:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
			v1[_100M-6] *= v2[_100M-6]
		case 5:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
			v1[_100M-5] *= v2[_100M-5]
		case 4:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
			v1[_100M-4] *= v2[_100M-4]
		case 3:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
			v1[_100M-3] *= v2[_100M-3]
		case 2:
			v1[_100M-1] *= v2[_100M-1]
			v1[_100M-2] *= v2[_100M-2]
		case 1:
			v1[_100M-1] *= v2[_100M-1]
		}
	}
}
