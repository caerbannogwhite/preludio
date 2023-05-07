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
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

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
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

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
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

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
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

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
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

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
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

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

////////////////////////			GROUP BY
//
// GroupBy challege: more info here https://github.com/h2oai/db-benchmark/tree/master

func Test_GroupBy_Q1_1e4(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1").Agg(Sum("v1"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 30027 {
		t.Errorf("Expected 30027, got %f", check)
	}
}

func Test_GroupBy_Q1_1e5(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1").Agg(Sum("v1"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 300292 {
		t.Errorf("Expected 300292, got %f", check)
	}
}

func Test_GroupBy_Q1_1e6(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1").Agg(Sum("v1"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 3000297 {
		t.Errorf("Expected 3000297, got %f", check)
	}
}

func Test_GroupBy_Q1_1e7(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1").Agg(Sum("v1"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 29998789 {
		t.Errorf("Expected 29998789, got %f", check)
	}
}

func Test_GroupBy_Q2_1e4(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1", "id2").Agg(Sum("v1"))

	if df.NRows() != 6272 {
		t.Errorf("Expected 6272 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 30027 {
		t.Errorf("Expected 30027, got %f", check)
	}
}

func Test_GroupBy_Q2_1e5(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1", "id2").Agg(Sum("v1"))

	if df.NRows() != 9999 {
		t.Errorf("Expected 9999 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 300292 {
		t.Errorf("Expected 300292, got %f", check)
	}
}

func Test_GroupBy_Q2_1e6(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1", "id2").Agg(Sum("v1"))

	if df.NRows() != 10000 {
		t.Errorf("Expected 10000 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 3000297 {
		t.Errorf("Expected 3000297, got %f", check)
	}
}

func Test_GroupBy_Q2_1e7(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id1", "id2").Agg(Sum("v1"))

	if df.NRows() != 10000 {
		t.Errorf("Expected 10000 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 29998789 {
		t.Errorf("Expected 29998789, got %f", check)
	}
}

func Test_GroupBy_Q3_1e4(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id3").Agg(Sum("v1"), Mean("v3"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check1 != 30027 {
		t.Errorf("Expected 30027, got %f", check1)
	}

	check2 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check2, 4999.383247863238, 10e-6) {
		t.Errorf("Expected 4999.383247863238, got %f", check2)
	}
}

func Test_GroupBy_Q3_1e5(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id3").Agg(Sum("v1"), Mean("v3"))

	if df.NRows() != 1000 {
		t.Errorf("Expected 1000 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check1 != 300292 {
		t.Errorf("Expected 300292, got %f", check1)
	}

	check2 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check2, 50095.22836212861, 10e-6) {
		t.Errorf("Expected 50095.22836212861, got %f", check2)
	}
}

func Test_GroupBy_Q3_1e6(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id3").Agg(Sum("v1"), Mean("v3"))

	if df.NRows() != 10000 {
		t.Errorf("Expected 10000 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check1 != 3000297 {
		t.Errorf("Expected 3000297, got %f", check1)
	}

	check2 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check2, 500393.46150263766, 10e-6) {
		t.Errorf("Expected 500393.46150263766, got %f", check2)
	}
}

func Test_GroupBy_Q3_1e7(t *testing.T) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		t.Skip(err)
	}

	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read().
		GroupBy("id3").Agg(Sum("v1"), Mean("v3"))

	if df.NRows() != 100000 {
		t.Errorf("Expected 100000 rows, got %d", df.NRows())
	}

	if df.NCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check1 != 29998789 {
		t.Errorf("Expected 29998789, got %f", check1)
	}

	check2 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check2, 4999719.62234443, 10e-6) {
		t.Errorf("Expected 4999719.62234443, got %f", check2)
	}
}

func Benchmark_GroupBy_Q1_1e4(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e4(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1", "id2").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1", "id2").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1", "id2").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id1", "id2").Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e4(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id3").Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id3").Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id3").Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id3").Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e4(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id4").Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id4").Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id4").Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id4").Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e4(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e4_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id6").Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e5(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e5_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id6").Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e6(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e6_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id6").Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e7(b *testing.B) {
	f, err := os.OpenFile("testdata\\G1_1e7_1e2_0_0.csv", os.O_RDONLY, 0666)
	if err != nil {
		b.Error(err)
	}
	df := NewBaseDataFrame().
		FromCSV().
		SetDelimiter(',').
		SetReader(f).
		Read()

	f.Close()

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.GroupBy("id6").Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}
