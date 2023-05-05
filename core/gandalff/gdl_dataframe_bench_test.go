package gandalff

import (
	"os"
	"runtime"
	"testing"
)

func Benchmark_Filter_Q1_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int) > 500
			}, nil).(GDLSeriesBool).Or(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}, nil).(GDLSeriesBool)).(GDLSeriesBool))
	}
	b.StopTimer()
}

func Benchmark_Filter_Q1_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int) > 500
			}, nil).(GDLSeriesBool).Or(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}, nil).(GDLSeriesBool)).(GDLSeriesBool))
	}
	b.StopTimer()
}

func Benchmark_Filter_Q1_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int) > 500
			}, nil).(GDLSeriesBool).Or(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}, nil).(GDLSeriesBool)).(GDLSeriesBool))
	}
	b.StopTimer()
}

func Benchmark_Filter_Q2_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int) > 500
			}, nil).(GDLSeriesBool).And(
				df.Series("v3").Map(func(v any) any {
					return v.(float64) < 50
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool).And(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}, nil).(GDLSeriesBool).Or(
					df.Series("id2").Map(func(v any) any {
						return v.(string) == "id024"
					}, nil).(GDLSeriesBool),
				),
			).(GDLSeriesBool).And(
				df.Series("v1").Map(func(v any) any {
					return v.(int) == 5
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool).And(
				df.Series("v2").Map(func(v any) any {
					return v.(int) == 1
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool),
		)
	}
	b.StopTimer()
}

func Benchmark_Filter_Q2_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int) > 500
			}, nil).(GDLSeriesBool).And(
				df.Series("v3").Map(func(v any) any {
					return v.(float64) < 50
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool).And(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}, nil).(GDLSeriesBool).Or(
					df.Series("id2").Map(func(v any) any {
						return v.(string) == "id024"
					}, nil).(GDLSeriesBool),
				),
			).(GDLSeriesBool).And(
				df.Series("v1").Map(func(v any) any {
					return v.(int) == 5
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool).And(
				df.Series("v2").Map(func(v any) any {
					return v.(int) == 1
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool),
		)
	}
	b.StopTimer()
}

func Benchmark_Filter_Q2_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int) > 500
			}, nil).(GDLSeriesBool).And(
				df.Series("v3").Map(func(v any) any {
					return v.(float64) < 50
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool).And(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}, nil).(GDLSeriesBool).Or(
					df.Series("id2").Map(func(v any) any {
						return v.(string) == "id024"
					}, nil).(GDLSeriesBool),
				),
			).(GDLSeriesBool).And(
				df.Series("v1").Map(func(v any) any {
					return v.(int) == 5
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool).And(
				df.Series("v2").Map(func(v any) any {
					return v.(int) == 1
				}, nil).(GDLSeriesBool),
			).(GDLSeriesBool),
		)
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1").Sum().Select("id1", "v1")
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1").Sum().Select("id1", "v1")
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := FromCSV(f, ',', true, 100)
	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1").Sum().Select("id1", "v1")
	}
	b.StopTimer()
}
