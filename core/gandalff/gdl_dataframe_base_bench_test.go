package gandalff

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var G1_1e4_1e2_0_0_df *DataFrame
var G1_1e5_1e2_0_0_df *DataFrame
var G1_1e6_1e2_0_0_df *DataFrame
var G1_1e7_1e2_0_0_df *DataFrame
var G1_1e4_1e2_10_0_df *DataFrame
var G1_1e5_1e2_10_0_df *DataFrame
var G1_1e6_1e2_10_0_df *DataFrame
var G1_1e7_1e2_10_0_df *DataFrame

func read_G1_1e4_1e2_0_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e4_1e2_0_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e4_1e2_0_0_df = &df
	} else {
		G1_1e4_1e2_0_0_df = nil
	}
}

func read_G1_1e5_1e2_0_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e5_1e2_0_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e5_1e2_0_0_df = &df
	} else {
		G1_1e5_1e2_0_0_df = nil
	}
}

func read_G1_1e6_1e2_0_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e6_1e2_0_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e6_1e2_0_0_df = &df
	} else {
		G1_1e6_1e2_0_0_df = nil
	}
}

func read_G1_1e7_1e2_0_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e7_1e2_0_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e7_1e2_0_0_df = &df
	} else {
		G1_1e7_1e2_0_0_df = nil
	}
}

func read_G1_1e4_1e2_10_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e4_1e2_10_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e4_1e2_10_0_df = &df
	} else {
		G1_1e4_1e2_10_0_df = nil
	}
}

func read_G1_1e5_1e2_10_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e5_1e2_10_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e5_1e2_10_0_df = &df
	} else {
		G1_1e5_1e2_10_0_df = nil
	}
}

func read_G1_1e6_1e2_10_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e6_1e2_10_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e6_1e2_10_0_df = &df
	} else {
		G1_1e6_1e2_10_0_df = nil
	}
}

func read_G1_1e7_1e2_10_0() {
	f, err := os.OpenFile(filepath.Join("testdata", "G1_1e7_1e2_10_0.csv"), os.O_RDONLY, 0666)
	if err == nil {
		df := NewBaseDataFrame(ctx).
			FromCSV().
			SetDelimiter(',').
			SetReader(f).
			Read()

		f.Close()

		G1_1e7_1e2_10_0_df = &df
	} else {
		G1_1e7_1e2_10_0_df = nil
	}
}

func init() {
	read_G1_1e4_1e2_0_0()
	read_G1_1e5_1e2_0_0()
	read_G1_1e6_1e2_0_0()
	read_G1_1e7_1e2_0_0()
	read_G1_1e4_1e2_10_0()
	read_G1_1e5_1e2_10_0()
	read_G1_1e6_1e2_10_0()
	read_G1_1e7_1e2_10_0()
}

func Benchmark_Filter_Q1_1e5(b *testing.B) {
	if G1_1e5_1e2_0_0_df == nil {
		b.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_0_0_df)
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int64) > 500
			}).(SeriesBool).Or(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}).(SeriesBool)).(SeriesBool))
	}
	b.StopTimer()
}

func Benchmark_Filter_Q1_1e6(b *testing.B) {
	if G1_1e6_1e2_0_0_df == nil {
		b.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_0_0_df)
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int64) > 500
			}).(SeriesBool).Or(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}).(SeriesBool)).(SeriesBool))
	}
	b.StopTimer()
}

func Benchmark_Filter_Q1_1e7(b *testing.B) {
	if G1_1e7_1e2_0_0_df == nil {
		b.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_0_0_df)
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int64) > 500
			}).(SeriesBool).Or(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}).(SeriesBool)).(SeriesBool))
	}
	b.StopTimer()
}

func Benchmark_Filter_Q2_1e5(b *testing.B) {
	if G1_1e5_1e2_0_0_df == nil {
		b.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_0_0_df)
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int64) > 500
			}).(SeriesBool).And(
				df.Series("v3").Map(func(v any) any {
					return v.(float64) < 50
				}).(SeriesBool),
			).(SeriesBool).And(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}).(SeriesBool).Or(
					df.Series("id2").Map(func(v any) any {
						return v.(string) == "id024"
					}).(SeriesBool),
				),
			).(SeriesBool).And(
				df.Series("v1").Map(func(v any) any {
					return v.(int64) == 5
				}).(SeriesBool),
			).(SeriesBool).And(
				df.Series("v2").Map(func(v any) any {
					return v.(int64) == 1
				}).(SeriesBool),
			).(SeriesBool),
		)
	}
	b.StopTimer()
}

func Benchmark_Filter_Q2_1e6(b *testing.B) {
	if G1_1e6_1e2_0_0_df == nil {
		b.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_0_0_df)
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int64) > 500
			}).(SeriesBool).And(
				df.Series("v3").Map(func(v any) any {
					return v.(float64) < 50
				}).(SeriesBool),
			).(SeriesBool).And(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}).(SeriesBool).Or(
					df.Series("id2").Map(func(v any) any {
						return v.(string) == "id024"
					}).(SeriesBool),
				),
			).(SeriesBool).And(
				df.Series("v1").Map(func(v any) any {
					return v.(int64) == 5
				}).(SeriesBool),
			).(SeriesBool).And(
				df.Series("v2").Map(func(v any) any {
					return v.(int64) == 1
				}).(SeriesBool),
			).(SeriesBool),
		)
	}
	b.StopTimer()
}

func Benchmark_Filter_Q2_1e7(b *testing.B) {
	if G1_1e7_1e2_0_0_df == nil {
		b.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_0_0_df)
	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df.Filter(
			df.Series("id6").Map(func(v any) any {
				return v.(int64) > 500
			}).(SeriesBool).And(
				df.Series("v3").Map(func(v any) any {
					return v.(float64) < 50
				}).(SeriesBool),
			).(SeriesBool).And(
				df.Series("id1").Map(func(v any) any {
					return v.(string) == "id024"
				}).(SeriesBool).Or(
					df.Series("id2").Map(func(v any) any {
						return v.(string) == "id024"
					}).(SeriesBool),
				),
			).(SeriesBool).And(
				df.Series("v1").Map(func(v any) any {
					return v.(int64) == 5
				}).(SeriesBool),
			).(SeriesBool).And(
				df.Series("v2").Map(func(v any) any {
					return v.(int64) == 1
				}).(SeriesBool),
			).(SeriesBool),
		)
	}
	b.StopTimer()
}

////////////////////////			GROUP BY TESTS
//
// GroupBy challege: more info here https://github.com/h2oai/db-benchmark/tree/master

func Test_GroupBy_Q1_1e4(t *testing.T) {
	if G1_1e4_1e2_0_0_df == nil {
		t.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e4_1e2_0_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

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
	if G1_1e5_1e2_0_0_df == nil {
		t.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_0_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

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
	if G1_1e6_1e2_0_0_df == nil {
		t.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_0_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

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
	if G1_1e7_1e2_0_0_df == nil {
		t.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_0_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

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

func Test_GroupBy_Q1_1e4_10PercNAs(t *testing.T) {
	if G1_1e4_1e2_10_0_df == nil {
		t.Skip("G1_1e4_1e2_10_0 dataframe not loaded")
	}

	df := (*G1_1e4_1e2_10_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

	if df.NRows() != 91 {
		t.Errorf("Expected 91 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 27044 {
		t.Errorf("Expected 27044, got %f", check)
	}
}

func Test_GroupBy_Q1_1e5_10PercNAs(t *testing.T) {
	if G1_1e5_1e2_10_0_df == nil {
		t.Skip("G1_1e5_1e2_10_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_10_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

	if df.NRows() != 91 {
		t.Errorf("Expected 91 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 270421 {
		t.Errorf("Expected 270421, got %f", check)
	}
}

func Test_GroupBy_Q1_1e6_10PercNAs(t *testing.T) {
	if G1_1e6_1e2_10_0_df == nil {
		t.Skip("G1_1e6_1e2_10_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_10_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

	if df.NRows() != 91 {
		t.Errorf("Expected 91 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 2700684 {
		t.Errorf("Expected 2700684, got %f", check)
	}
}

func Test_GroupBy_Q1_1e7_10PercNAs(t *testing.T) {
	if G1_1e7_1e2_10_0_df == nil {
		t.Skip("G1_1e7_1e2_10_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_10_0_df).
		GroupBy("id1").
		Agg(Sum("v1"))

	if df.NRows() != 91 {
		t.Errorf("Expected 91 rows, got %d", df.NRows())
	}

	if df.NCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NCols())
	}

	check := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if check != 26998588 {
		t.Errorf("Expected 26998588, got %f", check)
	}
}

func Test_GroupBy_Q2_1e4(t *testing.T) {
	if G1_1e4_1e2_0_0_df == nil {
		t.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e4_1e2_0_0_df).
		GroupBy("id1", "id2").
		Agg(Sum("v1"))

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
	if G1_1e5_1e2_0_0_df == nil {
		t.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_0_0_df).
		GroupBy("id1", "id2").
		Agg(Sum("v1"))

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
	if G1_1e6_1e2_0_0_df == nil {
		t.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_0_0_df).
		GroupBy("id1", "id2").
		Agg(Sum("v1"))

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
	if G1_1e7_1e2_0_0_df == nil {
		t.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_0_0_df).
		GroupBy("id1", "id2").
		Agg(Sum("v1"))

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
	if G1_1e4_1e2_0_0_df == nil {
		t.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e4_1e2_0_0_df).
		GroupBy("id3").
		Agg(Sum("v1"), Mean("v3"))

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
	if G1_1e5_1e2_0_0_df == nil {
		t.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_0_0_df).
		GroupBy("id3").
		Agg(Sum("v1"), Mean("v3"))

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
	if G1_1e6_1e2_0_0_df == nil {
		t.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_0_0_df).
		GroupBy("id3").
		Agg(Sum("v1"), Mean("v3"))

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
	if G1_1e7_1e2_0_0_df == nil {
		t.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_0_0_df).
		GroupBy("id3").
		Agg(Sum("v1"), Mean("v3"))

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

func Test_GroupBy_Q4_1e4(t *testing.T) {
	if G1_1e4_1e2_0_0_df == nil {
		t.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e4_1e2_0_0_df).
		GroupBy("id4").
		Agg(Mean("v1"), Mean("v2"), Mean("v3"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 300.1460223942026, 10e-6) {
		t.Errorf("Expected 300.1460223942026, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 803.8206781360852, 10e-6) {
		t.Errorf("Expected 803.8206781360852, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 5008.9079567058325, 10e-6) {
		t.Errorf("Expected 5008.9079567058325, got %f", check3)
	}
}

func Test_GroupBy_Q4_1e5(t *testing.T) {
	if G1_1e5_1e2_0_0_df == nil {
		t.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_0_0_df).
		GroupBy("id4").
		Agg(Mean("v1"), Mean("v2"), Mean("v3"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 300.29996127903826, 10e-6) {
		t.Errorf("Expected 300.29996127903826, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 800.8632014058803, 10e-6) {
		t.Errorf("Expected 800.8632014058803, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 5009.03811345283, 10e-6) {
		t.Errorf("Expected 5009.03811345283, got %f", check3)
	}
}

func Test_GroupBy_Q4_1e6(t *testing.T) {
	if G1_1e6_1e2_0_0_df == nil {
		t.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_0_0_df).
		GroupBy("id4").
		Agg(Mean("v1"), Mean("v2"), Mean("v3"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 300.0300474405866, 10e-6) {
		t.Errorf("Expected 300.0300474405866, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 799.8113837581368, 10e-6) {
		t.Errorf("Expected 799.8113837581368, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 5003.666447664572, 10e-6) {
		t.Errorf("Expected 5003.666447664572, got %f", check3)
	}
}

func Test_GroupBy_Q4_1e7(t *testing.T) {
	if G1_1e7_1e2_0_0_df == nil {
		t.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_0_0_df).
		GroupBy("id4").
		Agg(Mean("v1"), Mean("v2"), Mean("v3"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 299.9879818750654, 10e-6) {
		t.Errorf("Expected 299.9879818750654, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 799.8941794099782, 10e-6) {
		t.Errorf("Expected 799.8941794099782, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 4999.766872833688, 10e-6) {
		t.Errorf("Expected 4999.766872833688, got %f", check3)
	}
}

func Test_GroupBy_Q5_1e4(t *testing.T) {
	if G1_1e4_1e2_0_0_df == nil {
		t.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e4_1e2_0_0_df).
		GroupBy("id6").
		Agg(Sum("v1"), Sum("v2"), Sum("v3"))

	if df.NRows() != 100 {
		t.Errorf("Expected 100 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 30027, 10e-6) {
		t.Errorf("Expected 30027, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 80396, 10e-6) {
		t.Errorf("Expected 80396, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 500378.166716, 10e-6) {
		t.Errorf("Expected 500378.166716, got %f", check3)
	}
}

func Test_GroupBy_Q5_1e5(t *testing.T) {
	if G1_1e5_1e2_0_0_df == nil {
		t.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e5_1e2_0_0_df).
		GroupBy("id6").
		Agg(Sum("v1"), Sum("v2"), Sum("v3"))

	if df.NRows() != 1000 {
		t.Errorf("Expected 1000 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 300292, 10e-6) {
		t.Errorf("Expected 300292, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 800809, 10e-6) {
		t.Errorf("Expected 800809, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 5009219.2870470015, 10e-6) {
		t.Errorf("Expected 5009219.2870470015, got %f", check3)
	}
}

func Test_GroupBy_Q5_1e6(t *testing.T) {
	if G1_1e6_1e2_0_0_df == nil {
		t.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e6_1e2_0_0_df).
		GroupBy("id6").
		Agg(Sum("v1"), Sum("v2"), Sum("v3"))

	if df.NRows() != 10000 {
		t.Errorf("Expected 10000 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 3000297, 10e-6) {
		t.Errorf("Expected 3000297, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 7998131, 10e-6) {
		t.Errorf("Expected 7998131, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 50037098.685274005, 10e-6) {
		t.Errorf("Expected 50037098.685274005, got %f", check3)
	}
}

func Test_GroupBy_Q5_1e7(t *testing.T) {
	if G1_1e7_1e2_0_0_df == nil {
		t.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	df := (*G1_1e7_1e2_0_0_df).
		GroupBy("id6").
		Agg(Sum("v1"), Sum("v2"), Sum("v3"))

	if df.NRows() != 100000 {
		t.Errorf("Expected 100000 rows, got %d", df.NRows())
	}

	if df.NCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", df.NCols())
	}

	check1 := df.Agg(Sum("v1")).Series("v1").Get(0).(float64)
	if !equalFloats(check1, 29998789, 10e-6) {
		t.Errorf("Expected 29998789, got %f", check1)
	}

	check2 := df.Agg(Sum("v2")).Series("v2").Get(0).(float64)
	if !equalFloats(check2, 79989360, 10e-6) {
		t.Errorf("Expected 79989360, got %f", check2)
	}

	check3 := df.Agg(Sum("v3")).Series("v3").Get(0).(float64)
	if !equalFloats(check3, 499976651.4080609, 10e-6) {
		t.Errorf("Expected 499976651.4080609, got %f", check3)
	}
}

////////////////////////			GROUP BY BENCHMARKS

func Benchmark_GroupBy_Q1_1e4(b *testing.B) {
	if G1_1e4_1e2_0_0_df == nil {
		b.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e4_1e2_0_0_df).GroupBy("id1").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e5(b *testing.B) {
	if G1_1e5_1e2_0_0_df == nil {
		b.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e5_1e2_0_0_df).GroupBy("id1").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e6(b *testing.B) {
	if G1_1e6_1e2_0_0_df == nil {
		b.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e6_1e2_0_0_df).GroupBy("id1").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q1_1e7(b *testing.B) {
	if G1_1e7_1e2_0_0_df == nil {
		b.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e7_1e2_0_0_df).GroupBy("id1").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e4(b *testing.B) {
	if G1_1e4_1e2_0_0_df == nil {
		b.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e4_1e2_0_0_df).GroupBy("id1", "id2").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e5(b *testing.B) {
	if G1_1e5_1e2_0_0_df == nil {
		b.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e5_1e2_0_0_df).GroupBy("id1", "id2").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e6(b *testing.B) {
	if G1_1e6_1e2_0_0_df == nil {
		b.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e6_1e2_0_0_df).GroupBy("id1", "id2").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q2_1e7(b *testing.B) {
	if G1_1e7_1e2_0_0_df == nil {
		b.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e7_1e2_0_0_df).GroupBy("id1", "id2").
			Agg(Sum("v1"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e4(b *testing.B) {
	if G1_1e4_1e2_0_0_df == nil {
		b.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e4_1e2_0_0_df).GroupBy("id3").
			Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e5(b *testing.B) {
	if G1_1e5_1e2_0_0_df == nil {
		b.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e5_1e2_0_0_df).GroupBy("id3").
			Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e6(b *testing.B) {
	if G1_1e6_1e2_0_0_df == nil {
		b.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e6_1e2_0_0_df).GroupBy("id3").
			Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q3_1e7(b *testing.B) {
	if G1_1e7_1e2_0_0_df == nil {
		b.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e7_1e2_0_0_df).GroupBy("id3").
			Agg(Sum("v1"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e4(b *testing.B) {
	if G1_1e4_1e2_0_0_df == nil {
		b.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e4_1e2_0_0_df).
			GroupBy("id4").
			Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e5(b *testing.B) {
	if G1_1e5_1e2_0_0_df == nil {
		b.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e5_1e2_0_0_df).
			GroupBy("id4").
			Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e6(b *testing.B) {
	if G1_1e6_1e2_0_0_df == nil {
		b.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e6_1e2_0_0_df).GroupBy("id4").Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q4_1e7(b *testing.B) {
	if G1_1e7_1e2_0_0_df == nil {
		b.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e7_1e2_0_0_df).GroupBy("id4").Agg(Mean("v1"), Mean("v2"), Mean("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e4(b *testing.B) {
	if G1_1e4_1e2_0_0_df == nil {
		b.Skip("G1_1e4_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e4_1e2_0_0_df).
			GroupBy("id6").
			Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e5(b *testing.B) {
	if G1_1e5_1e2_0_0_df == nil {
		b.Skip("G1_1e5_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e5_1e2_0_0_df).
			GroupBy("id6").
			Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e6(b *testing.B) {
	if G1_1e6_1e2_0_0_df == nil {
		b.Skip("G1_1e6_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e6_1e2_0_0_df).GroupBy("id6").
			Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}

func Benchmark_GroupBy_Q5_1e7(b *testing.B) {
	if G1_1e7_1e2_0_0_df == nil {
		b.Skip("G1_1e7_1e2_0_0 dataframe not loaded")
	}

	runtime.GC()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*G1_1e7_1e2_0_0_df).GroupBy("id6").
			Agg(Sum("v1"), Sum("v2"), Sum("v3"))
	}
	b.StopTimer()
}
