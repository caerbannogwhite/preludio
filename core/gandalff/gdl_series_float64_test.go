package gandalff

import (
	"math/rand"
	"testing"
	"typesys"
)

func Test_GDLSeriesFloat64_Base(t *testing.T) {

	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GDLSeriesFloat64.
	s := NewGDLSeriesFloat64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Check the length.
	if s.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", s.Len())
	}

	// Check the name.
	if s.Name() != "test" {
		t.Errorf("Expected name of \"test\", got %s", s.Name())
	}

	// Check the type.
	if s.Type() != typesys.Float64Type {
		t.Errorf("Expected type of Float64Type, got %s", s.Type().ToString())
	}

	// Check the data.
	for i, v := range s.Data().([]float64) {
		if v != data[i] {
			t.Errorf("Expected data of []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}, got %v", s.Data())
		}
	}

	// Check the nullability.
	if !s.IsNullable() {
		t.Errorf("Expected IsNullable() to be true, got false")
	}

	// Check the null mask.
	for i, v := range s.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of []bool{false, false, false, false, true, false, true, false, false, true}, got %v", s.GetNullMask())
		}
	}

	// Check the null values.
	for i := range s.Data().([]float64) {
		if s.IsNull(i) != mask[i] {
			t.Errorf("Expected IsNull(%d) to be %t, got %t", i, mask[i], s.IsNull(i))
		}
	}

	// Check the null count.
	if s.NullCount() != 3 {
		t.Errorf("Expected NullCount() to be 3, got %d", s.NullCount())
	}

	// Check the HasNull() method.
	if !s.HasNull() {
		t.Errorf("Expected HasNull() to be true, got false")
	}

	// Check the SetNull() method.
	for i := range s.Data().([]float64) {
		s.SetNull(i)
	}

	// Check the null values.
	for i := range s.Data().([]float64) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get() method.
	for i, v := range s.Data().([]float64) {
		if s.Get(i).(float64) != v {
			t.Errorf("Expected Get(%d) to be %f, got %f", i, v, s.Get(i).(float64))
		}
	}

	// Check the Set() method.
	for i := range s.Data().([]float64) {
		s.Set(i, float64(i+10))
	}

	// Check the data.
	for i, v := range s.Data().([]float64) {
		if v != float64(i+10) {
			t.Errorf("Expected data of []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}, got %v", s.Data())
		}
	}

	// Check the MakeNullable() method.
	p := NewGDLSeriesFloat64("test", false, true, data)

	// Check the nullability.
	if p.IsNullable() {
		t.Errorf("Expected IsNullable() to be false, got true")
	}

	// Set values to null.
	p.SetNull(1)
	p.SetNull(3)
	p.SetNull(5)

	// Check the null count.
	if p.NullCount() != 0 {
		t.Errorf("Expected NullCount() to be 0, got %d", p.NullCount())
	}

	// Make the series nullable.
	p = p.MakeNullable().(GDLSeriesFloat64)

	// Check the nullability.
	if !p.IsNullable() {
		t.Errorf("Expected IsNullable() to be true, got false")
	}

	// Check the null count.
	if p.NullCount() != 0 {
		t.Errorf("Expected NullCount() to be 0, got %d", p.NullCount())
	}

	p.SetNull(1)
	p.SetNull(3)
	p.SetNull(5)

	// Check the null count.
	if p.NullCount() != 3 {
		t.Errorf("Expected NullCount() to be 3, got %d", p.NullCount())
	}
}

func Test_GDLSeriesFloat64_Append(t *testing.T) {
	dataA := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	dataB := []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	dataC := []float64{21.0, 22.0, 23.0, 24.0, 25.0, 26.0, 27.0, 28.0, 29.0, 30.0}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewGDLSeriesFloat64("testA", true, true, dataA)
	sB := NewGDLSeriesFloat64("testB", true, true, dataB)
	sC := NewGDLSeriesFloat64("testC", true, true, dataC)

	// Set the null masks.
	sA.SetNullMask(maskA)
	sB.SetNullMask(maskB)
	sC.SetNullMask(maskC)

	// Append the series.
	result := sA.AppendSeries(sB).AppendSeries(sC)

	// Check the name.
	if result.Name() != "testA" {
		t.Errorf("Expected name of \"testA\", got %s", result.Name())
	}

	// Check the length.
	if result.Len() != 30 {
		t.Errorf("Expected length of 30, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]float64) {
		if i < 10 {
			if v != dataA[i] {
				t.Errorf("Expected %f, got %f at index %d", dataA[i], v, i)
			}
		} else if i < 20 {
			if v != dataB[i-10] {
				t.Errorf("Expected %f, got %f at index %d", dataB[i-10], v, i)
			}
		} else {
			if v != dataC[i-20] {
				t.Errorf("Expected %f, got %f at index %d", dataC[i-20], v, i)
			}
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if i < 10 {
			if v != maskA[i] {
				t.Errorf("Expected nullMask %t, got %t at index %d", maskA[i], v, i)
			}
		} else if i < 20 {
			if v != maskB[i-10] {
				t.Errorf("Expected nullMask %t, got %t at index %d", maskB[i-10], v, i)
			}
		} else {
			if v != maskC[i-20] {
				t.Errorf("Expected nullMask %t, got %t at index %d", maskC[i-20], v, i)
			}
		}
	}

	// Append random values.
	dataD := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	sD := NewGDLSeriesFloat64("testD", true, true, dataD)

	// Check the original data.
	for i, v := range sD.Data().([]float64) {
		if v != dataD[i] {
			t.Errorf("Expected %f, got %f at index %d", dataD[i], v, i)
		}
	}

	for i := 0; i < 100; i++ {
		r := rand.Float64()
		switch i % 4 {
		case 0:
			sD = sD.Append(r).(GDLSeriesFloat64)
		case 1:
			sD = sD.Append([]float64{r}).(GDLSeriesFloat64)
		case 2:
			sD = sD.Append(NullableFloat64{true, r}).(GDLSeriesFloat64)
		case 3:
			sD = sD.Append([]NullableFloat64{{false, r}}).(GDLSeriesFloat64)
		}

		if sD.Get(i+10) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
		}
	}
}

func Test_GDLSeriesFloat64_Filter(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewGDLSeriesFloat64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []float64{1.0, 3.0, 4.0, 6.0, 7.0, 9.0, 10.0, 11.0, 13.0, 14.0, 16.0, 17.0, 19.0, 20.0}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the FilterByMask() method.
	filtered := s.FilterByMask(filterMask)

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]float64) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range filtered.GetNullMask() {
		if v != resultMask[i] {
			t.Errorf("Expected nullMask of %v, got %v at index %d", resultMask[i], v, i)
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the FilterByIndeces() method.
	filtered = s.FilterByIndeces(filterIndeces)

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]float64) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range filtered.GetNullMask() {
		if v != resultMask[i] {
			t.Errorf("Expected nullMask of %v, got %v at index %d", resultMask[i], v, i)
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////

	// try to filter by a series with a different length.
	filtered = filtered.FilterByMask(filterMask)

	if e, ok := filtered.(GDLSeriesError); !ok || e.Error() != "GDLSeriesFloat64.FilterByMask: mask length (20) does not match series length (14)" {
		t.Errorf("Expected GDLSeriesError, got %v", filtered)
	}

	// Another test.
	data = []float64{2.0, 323, 42, 4.1, 9, 674.0, 42, 48, 9811, 79, 3, 12, 492.3, 47005, -173.4, -28, 323, 42.5, 4, 9.0, 31, 425, 2}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewGDLSeriesFloat64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []float64{2.0, -28, 2}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the FilterByMask() method.
	filtered = s.FilterByMask(filterMask)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]float64) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range filtered.GetNullMask() {
		if v != true {
			t.Errorf("Expected nullMask of %v, got %v at index %d", true, v, i)
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the FilterByIndeces() method.
	filtered = s.FilterByIndeces(filterIndeces)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]float64) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range filtered.GetNullMask() {
		if v != true {
			t.Errorf("Expected nullMask of %v, got %v at index %d", true, v, i)
		}
	}
}

func TestGDLSeriesFloat64_Multiplication(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// s * 1.5
	res := NewGDLSeriesFloat64("test", true, true, data).Mul(NewGDLSeriesFloat64("test", true, true, []float64{1.5}))
	if e, ok := res.(GDLSeriesError); ok {
		t.Errorf("Got error: %v", e)
	}

	// Check the length.
	if res.Len() != 20 {
		t.Errorf("Expected length of 20, got %d", res.Len())
	}

	// Check the data.
	for i, v := range res.Data().([]float64) {
		if v != data[i]*1.5 {
			t.Errorf("Expected %v, got %v at index %d", data[i]*1.5, v, i)
		}
	}

	// 1.5 * s
	res = NewGDLSeriesFloat64("test", true, true, []float64{1.5}).Mul(NewGDLSeriesFloat64("test", true, true, data))
	if e, ok := res.(GDLSeriesError); ok {
		t.Errorf("Got error: %v", e)
	}

	// Check the length.
	if res.Len() != 20 {
		t.Errorf("Expected length of 20, got %d", res.Len())
	}

	// Check the data.
	for i, v := range res.Data().([]float64) {
		if v != data[i]*1.5 {
			t.Errorf("Expected %v, got %v at index %d", data[i]*1.5, v, i)
		}
	}
}

func BenchmarkGDLSeriesFloat64_Mul_SerScal_Perf(b *testing.B) {

	N := 1_000_000
	data := make([]float64, N)
	for i := 0; i < N; i++ {
		data[i] = float64(i)
	}

	ser := NewGDLSeriesFloat64("test", true, false, data)
	scal := NewGDLSeriesFloat64("test", true, false, []float64{1.5})

	// s * 1.5
	var res GDLSeries
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ser.Mul(scal)
	}

	if e, ok := res.(GDLSeriesError); ok {
		b.Errorf("Got error: %v", e)
	}

	// Check the length.
	if res.Len() != N {
		b.Errorf("Expected length of %d, got %d", N, res.Len())
	}
}

func BenchmarkGDLSeriesFloat64_Mul_SerSer_Perf(b *testing.B) {

	N := 1_000_000
	data1 := make([]float64, N)
	data2 := make([]float64, N)
	for i := 0; i < N; i++ {
		data1[i] = float64(i)
		data2[i] = float64(N - i - 1)
	}

	ser1 := NewGDLSeriesFloat64("test", true, false, data1)
	ser2 := NewGDLSeriesFloat64("test", true, false, data2)

	// s * 1.5
	var res GDLSeries
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ser1.Mul(ser2)
	}

	if e, ok := res.(GDLSeriesError); ok {
		b.Errorf("Got error: %v", e)
	}

	// Check the length.
	if res.Len() != N {
		b.Errorf("Expected length of %d, got %d", N, res.Len())
	}
}
