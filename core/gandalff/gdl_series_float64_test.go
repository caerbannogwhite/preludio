package gandalff

import (
	"math"
	"math/rand"
	"testing"
	"typesys"
)

func Test_SeriesFloat64_Base(t *testing.T) {

	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesFloat64.
	s := NewSeriesFloat64("test", true, true, data)

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
	p := NewSeriesFloat64("test", false, true, data)

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
	p = p.MakeNullable().(SeriesFloat64)

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

func Test_SeriesFloat64_Append(t *testing.T) {
	dataA := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	dataB := []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	dataC := []float64{21.0, 22.0, 23.0, 24.0, 25.0, 26.0, 27.0, 28.0, 29.0, 30.0}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesFloat64("testA", true, true, dataA)
	sB := NewSeriesFloat64("testB", true, true, dataB)
	sC := NewSeriesFloat64("testC", true, true, dataC)

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
	sD := NewSeriesFloat64("testD", true, true, dataD)

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
			sD = sD.Append(r).(SeriesFloat64)
		case 1:
			sD = sD.Append([]float64{r}).(SeriesFloat64)
		case 2:
			sD = sD.Append(NullableFloat64{true, r}).(SeriesFloat64)
		case 3:
			sD = sD.Append([]NullableFloat64{{false, r}}).(SeriesFloat64)
		}

		if sD.Get(i+10) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
		}
	}
}

func Test_SeriesFloat64_Cast(t *testing.T) {
	data := []float64{0.0, 1.0, 0.0, 3.0, 4.0, 5.0, -6.0, 7.0, 8.0, 9.0}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesFloat64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Cast to bool.
	resBool := s.Cast(typesys.BoolType, nil)

	// Check the data.
	for i, v := range resBool.Data().([]bool) {
		if data[i] == 0.0 && v != false {
			t.Errorf("Expected %t, got %t at index %d", false, v, i)
		} else if data[i] != 0.0 && v != true {
			t.Errorf("Expected %t, got %t at index %d", true, v, i)
		}
	}

	// Check the null mask.
	for i, v := range resBool.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to int32.
	resInt := s.Cast(typesys.Int32Type, nil)

	// Check the data.
	for i, v := range resInt.Data().([]int32) {
		if v != int32(data[i]) {
			t.Errorf("Expected %d, got %d at index %d", int32(data[i]), v, i)
		}
	}

	// Check the null mask.
	for i, v := range resInt.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to int64.
	resInt64 := s.Cast(typesys.Int64Type, nil)

	// Check the data.
	for i, v := range resInt64.Data().([]int64) {
		if v != int64(data[i]) {
			t.Errorf("Expected %d, got %d at index %d", int64(data[i]), v, i)
		}
	}

	// Check the null mask.
	for i, v := range resInt64.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to string.
	resString := s.Cast(typesys.StringType, NewStringPool())

	// Check the data.
	for i, v := range resString.Data().([]string) {
		if mask[i] && v != NULL_STRING {
			t.Errorf("Expected %s, got %s at index %d", NULL_STRING, v, i)
		} else if !mask[i] && v != floatToString(data[i]) {
			t.Errorf("Expected %s, got %s at index %d", floatToString(data[i]), v, i)
		}
	}

	// Check the null mask.
	for i, v := range resString.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to error.
	castError := s.Cast(typesys.ErrorType, nil)

	// Check the message.
	if castError.(SeriesError).msg != "SeriesFloat64.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_SeriesFloat64_Filter(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewSeriesFloat64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []float64{1.0, 3.0, 4.0, 6.0, 7.0, 9.0, 10.0, 11.0, 13.0, 14.0, 16.0, 17.0, 19.0, 20.0}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool("mask", false, true, filterMask).(SeriesBool))

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
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

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
	// 							Check the Filter() method.
	filtered = s.Filter(filterIndeces)

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
	filtered = filtered.Filter(filterMask)

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesFloat64.FilterByMask: mask length (20) does not match series length (14)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []float64{2.0, 323, 42, 4.1, 9, 674.0, 42, 48, 9811, 79, 3, 12, 492.3, 47005, -173.4, -28, 323, 42.5, 4, 9.0, 31, 425, 2}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesFloat64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []float64{2.0, -28, 2}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

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
	// 							Check the Filter() method.
	filtered = s.Filter(filterIndeces)

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

func Test_SeriesFloat64_Map(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -2, 323, 24, -23, 4, 42, 5, -6, 7}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true}

	// Create a new series.
	s := NewSeriesFloat64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(nullMask)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		if v.(float64) >= 7 && v.(float64) <= 100 {
			return true
		}
		return false
	}, nil)

	expectedBool := []bool{false, false, false, false, false, false, true, true, true, true, false, false, true, false, false, true, false, false, true}
	for i, v := range resBool.Data().([]bool) {
		if v != expectedBool[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedBool[i], v, i)
		}
	}

	// Map the series to int32.
	resInt := s.Map(func(v any) any {
		if v.(float64) < 0 {
			return (-int32(v.(float64))) % 7
		}
		return int32(v.(float64)) % 7
	}, nil)

	expectedInt32 := []int32{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt.Data().([]int32) {
		if v != expectedInt32[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt32[i], v, i)
		}
	}

	// Map the series to int64.
	resInt64 := s.Map(func(v any) any {
		if v.(float64) < 0 {
			return (-int64(v.(float64))) % 7
		}
		return int64(v.(float64)) % 7
	}, nil)

	expectedInt64 := []int64{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt64.Data().([]int64) {
		if v != expectedInt64[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt64[i], v, i)
		}
	}

	// Map the series to float64.
	resFloat64 := s.Map(func(v any) any {
		if v.(float64) >= 0 {
			return -v.(float64)
		}
		return v.(float64)
	}, nil)

	expectedFloat64 := []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -2, -323, -24, -23, -4, -42, -5, -6, -7}
	for i, v := range resFloat64.Data().([]float64) {
		if v != expectedFloat64[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedFloat64[i], v, i)
		}
	}

	// Map the series to string.
	resString := s.Map(func(v any) any {
		if v.(float64) >= 0 {
			return "pos"
		}
		return "neg"
	}, NewStringPool())

	expectedString := []string{"pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "neg", "pos", "pos", "neg", "pos", "pos", "pos", "neg", "pos"}
	for i, v := range resString.Data().([]string) {
		if v != expectedString[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedString[i], v, i)
		}
	}
}

func Test_SeriesFloat64_Arithmetic_Mul(t *testing.T) {
	bools := NewSeriesBool("test", true, false, []bool{true}).(SeriesBool)
	boolv := NewSeriesBool("test", true, false, []bool{true, false, true, false, true, false, true, true, false, false}).(SeriesBool)
	bools_ := NewSeriesBool("test", true, false, []bool{true}).SetNullMask([]bool{true}).(SeriesBool)
	boolv_ := NewSeriesBool("test", true, false, []bool{true, false, true, false, true, false, true, true, false, false}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesBool)

	i32s := NewSeriesInt32("test", true, false, []int32{2}).(SeriesInt32)
	i32v := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt32)
	i32s_ := NewSeriesInt32("test", true, false, []int32{2}).SetNullMask([]bool{true}).(SeriesInt32)
	i32v_ := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt32)

	i64s := NewSeriesInt64("test", true, false, []int64{2}).(SeriesInt64)
	i64v := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt64)
	i64s_ := NewSeriesInt64("test", true, false, []int64{2}).SetNullMask([]bool{true}).(SeriesInt64)
	i64v_ := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt64)

	f64s := NewSeriesFloat64("test", true, false, []float64{2}).(SeriesFloat64)
	f64v := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesFloat64)
	f64s_ := NewSeriesFloat64("test", true, false, []float64{2}).SetNullMask([]bool{true}).(SeriesFloat64)
	f64v_ := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesFloat64)

	// scalar | bool
	if !checkEqSlice(f64s.Mul(bools).Data().([]float64), []float64{2}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(boolv).Data().([]float64), []float64{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(bools_).GetNullMask(), []bool{true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}

	// scalar | int32
	if !checkEqSlice(f64s.Mul(i32s).Data().([]float64), []float64{4}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(i32v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Mul(i64s).Data().([]float64), []float64{4}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(i64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Mul(f64s).Data().([]float64), []float64{4}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64s.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}

	// vector | bool
	if !checkEqSlice(f64v.Mul(bools).Data().([]float64), []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(boolv).Data().([]float64), []float64{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}

	// vector | int32
	if !checkEqSlice(f64v.Mul(i32s).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(i32v).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}

	// vector | int64
	if !checkEqSlice(f64v.Mul(i64s).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(i64v).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}

	// vector | float64
	if !checkEqSlice(f64v.Mul(f64s).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(f64v).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
	if !checkEqSlice(f64v.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mul") {
		t.Errorf("Error in Float64 Mul")
	}
}

func Test_SeriesFloat64_Arithmetic_Div(t *testing.T) {
	bools := NewSeriesBool("test", true, false, []bool{true}).(SeriesBool)
	boolv := NewSeriesBool("test", true, false, []bool{true, false, true, false, true, false, true, true, false, false}).(SeriesBool)
	bools_ := NewSeriesBool("test", true, false, []bool{true}).SetNullMask([]bool{true}).(SeriesBool)
	boolv_ := NewSeriesBool("test", true, false, []bool{true, false, true, false, true, false, true, true, false, false}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesBool)

	i32s := NewSeriesInt32("test", true, false, []int32{2}).(SeriesInt32)
	i32v := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt32)
	i32s_ := NewSeriesInt32("test", true, false, []int32{2}).SetNullMask([]bool{true}).(SeriesInt32)
	i32v_ := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt32)

	i64s := NewSeriesInt64("test", true, false, []int64{2}).(SeriesInt64)
	i64v := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt64)
	i64s_ := NewSeriesInt64("test", true, false, []int64{2}).SetNullMask([]bool{true}).(SeriesInt64)
	i64v_ := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt64)

	f64s := NewSeriesFloat64("test", true, false, []float64{2}).(SeriesFloat64)
	f64v := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesFloat64)
	f64s_ := NewSeriesFloat64("test", true, false, []float64{2}).SetNullMask([]bool{true}).(SeriesFloat64)
	f64v_ := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesFloat64)

	// scalar | bool
	if !checkEqSlice(f64s.Div(bools).Data().([]float64), []float64{2}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(boolv).Data().([]float64), []float64{2, math.Inf(1), 2, math.Inf(1), 2, math.Inf(1), 2, 2, math.Inf(1), math.Inf(1)}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(bools_).GetNullMask(), []bool{true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}

	// scalar | int32
	if !checkEqSlice(f64s.Div(i32s).Data().([]float64), []float64{1}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(i32v).Data().([]float64), []float64{2, 1, 2.0 / 3, 0.5, 0.4, 1.0 / 3, 2.0 / 7, 0.25, 2.0 / 9, 0.2}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Div(i64s).Data().([]float64), []float64{1}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(i64v).Data().([]float64), []float64{2, 1, 2.0 / 3, 0.5, 0.4, 1.0 / 3, 2.0 / 7, 0.25, 2.0 / 9, 0.2}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Div(f64s).Data().([]float64), []float64{1}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(f64v).Data().([]float64), []float64{2, 1, 2.0 / 3, 0.5, 0.4, 1.0 / 3, 2.0 / 7, 0.25, 2.0 / 9, 0.2}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64s.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}

	// vector | bool
	if !checkEqSlice(f64v.Div(bools).Data().([]float64), []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(boolv).Data().([]float64), []float64{1, math.Inf(1), 3, math.Inf(1), 5, math.Inf(1), 7, 8, math.Inf(1), math.Inf(1)}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}

	// vector | int32
	if !checkEqSlice(f64v.Div(i32s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(i32v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}

	// vector | int64
	if !checkEqSlice(f64v.Div(i64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(i64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}

	// vector | float64
	if !checkEqSlice(f64v.Div(f64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(f64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
	if !checkEqSlice(f64v.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Div") {
		t.Errorf("Error in Float64 Div")
	}
}

func Test_SeriesFloat64_Arithmetic_Mod(t *testing.T) {
	i32s := NewSeriesInt32("test", true, false, []int32{2}).(SeriesInt32)
	i32v := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt32)
	i32s_ := NewSeriesInt32("test", true, false, []int32{2}).SetNullMask([]bool{true}).(SeriesInt32)
	i32v_ := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt32)

	i64s := NewSeriesInt64("test", true, false, []int64{2}).(SeriesInt64)
	i64v := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt64)
	i64s_ := NewSeriesInt64("test", true, false, []int64{2}).SetNullMask([]bool{true}).(SeriesInt64)
	i64v_ := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt64)

	f64s := NewSeriesFloat64("test", true, false, []float64{2}).(SeriesFloat64)
	f64v := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesFloat64)
	f64s_ := NewSeriesFloat64("test", true, false, []float64{2}).SetNullMask([]bool{true}).(SeriesFloat64)
	f64v_ := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesFloat64)

	// scalar | int32
	if !checkEqSlice(f64s.Mod(i32s).Data().([]int64), []int64{0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i32v).Data().([]int64), []int64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Mod(i64s).Data().([]int64), []int64{0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i64v).Data().([]int64), []int64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Mod(f64s).Data().([]int64), []int64{0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(f64v).Data().([]int64), []int64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// vector | int32
	if !checkEqSlice(f64v.Mod(i32s).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i32v).Data().([]int64), []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// vector | int64
	if !checkEqSlice(f64v.Mod(i64s).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i64v).Data().([]int64), []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// vector | float64
	if !checkEqSlice(f64v.Mod(f64s).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(f64v).Data().([]int64), []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
}

func Test_SeriesFloat64_Arithmetic_Pow(t *testing.T) {
	i32s := NewSeriesInt32("test", true, false, []int32{2}).(SeriesInt32)
	i32v := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt32)
	i32s_ := NewSeriesInt32("test", true, false, []int32{2}).SetNullMask([]bool{true}).(SeriesInt32)
	i32v_ := NewSeriesInt32("test", true, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt32)

	i64s := NewSeriesInt64("test", true, false, []int64{2}).(SeriesInt64)
	i64v := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesInt64)
	i64s_ := NewSeriesInt64("test", true, false, []int64{2}).SetNullMask([]bool{true}).(SeriesInt64)
	i64v_ := NewSeriesInt64("test", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesInt64)

	f64s := NewSeriesFloat64("test", true, false, []float64{2}).(SeriesFloat64)
	f64v := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(SeriesFloat64)
	f64s_ := NewSeriesFloat64("test", true, false, []float64{2}).SetNullMask([]bool{true}).(SeriesFloat64)
	f64v_ := NewSeriesFloat64("test", true, false, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesFloat64)

	// scalar | int32
	if !checkEqSlice(f64s.Pow(i32s).Data().([]float64), []float64{4}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(i32v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Pow(i64s).Data().([]float64), []float64{4}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(i64v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Pow(f64s).Data().([]float64), []float64{4}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(f64v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64s.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}

	// vector | int32
	if !checkEqSlice(f64v.Pow(i32s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(i32v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}

	// vector | int64
	if !checkEqSlice(f64v.Pow(i64s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(i64v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}

	// vector | float64
	if !checkEqSlice(f64v.Pow(f64s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(f64v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
	if !checkEqSlice(f64v.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Pow") {
		t.Errorf("Error in Float64 Pow")
	}
}

func Test_SeriesFloat64_Arithmetic_Sub(t *testing.T) {
	var res Series

	i32s := NewSeriesInt32("test", true, false, []int32{1}).(SeriesInt32)
	i32v := NewSeriesInt32("test", true, false, []int32{1, 2, 3}).(SeriesInt32)
	i32s_ := NewSeriesInt32("test", true, false, []int32{1}).SetNullMask([]bool{true}).(SeriesInt32)
	i32v_ := NewSeriesInt32("test", true, false, []int32{1, 2, 3}).SetNullMask([]bool{true, true, false}).(SeriesInt32)

	i64s := NewSeriesInt64("test", true, false, []int64{1}).(SeriesInt64)
	i64v := NewSeriesInt64("test", true, false, []int64{1, 2, 3}).(SeriesInt64)
	i64s_ := NewSeriesInt64("test", true, false, []int64{1}).SetNullMask([]bool{true}).(SeriesInt64)
	i64v_ := NewSeriesInt64("test", true, false, []int64{1, 2, 3}).SetNullMask([]bool{true, true, false}).(SeriesInt64)

	f64s := NewSeriesFloat64("test", true, false, []float64{1}).(SeriesFloat64)
	f64v := NewSeriesFloat64("test", true, false, []float64{1, 2, 3}).(SeriesFloat64)
	f64s_ := NewSeriesFloat64("test", true, false, []float64{1}).SetNullMask([]bool{true}).(SeriesFloat64)
	f64v_ := NewSeriesFloat64("test", true, false, []float64{1, 2, 3}).SetNullMask([]bool{true, true, false}).(SeriesFloat64)

	// scalar | int32
	res = f64s.Sub(i32s)
	if res.Data().([]float64)[0] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0}, res.Data().([]float64))
	}

	res = f64s.Sub(i32v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != -1 || res.Data().([]float64)[2] != -2 {
		t.Errorf("Expected %v, got %v", []float64{0, -1, -2}, res.Data().([]float64))
	}

	res = f64s.Sub(i32s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Sub(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | int64
	res = f64s.Sub(i64s)
	if res.Data().([]float64)[0] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0}, res.Data().([]float64))
	}

	res = f64s.Sub(i64v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != -1 || res.Data().([]float64)[2] != -2 {
		t.Errorf("Expected %v, got %v", []float64{0, -1, -2}, res.Data().([]float64))
	}

	res = f64s.Sub(i64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Sub(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | float64
	res = f64s.Sub(f64s)
	if res.Data().([]float64)[0] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0}, res.Data().([]float64))
	}

	res = f64s.Sub(f64v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != -1 || res.Data().([]float64)[2] != -2 {
		t.Errorf("Expected %v, got %v", []float64{0, -1, -2}, res.Data().([]float64))
	}

	res = f64s.Sub(f64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Sub(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int32
	res = f64v.Sub(i32s)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 1 || res.Data().([]float64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []float64{0, 1, 2}, res.Data().([]float64))
	}

	res = f64v.Sub(i32v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 0 || res.Data().([]float64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0, 0, 0}, res.Data().([]float64))
	}

	res = f64v.Sub(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = f64v.Sub(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int64
	res = f64v.Sub(i64s)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 1 || res.Data().([]float64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []float64{0, 1, 2}, res.Data().([]float64))
	}

	res = f64v.Sub(i64v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 0 || res.Data().([]float64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0, 0, 0}, res.Data().([]float64))
	}

	res = f64v.Sub(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = f64v.Sub(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | float64
	res = f64v.Sub(f64s)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 1 || res.Data().([]float64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []float64{0, 1, 2}, res.Data().([]float64))
	}

	res = f64v.Sub(f64v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 0 || res.Data().([]float64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0, 0, 0}, res.Data().([]float64))
	}

	res = f64v.Sub(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = f64v.Sub(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}
}

func Test_SeriesFloat64_Arithmetic_Add(t *testing.T) {
	var res Series

	i32s := NewSeriesInt32("test", true, false, []int32{1}).(SeriesInt32)
	i32v := NewSeriesInt32("test", true, false, []int32{1, 2, 3}).(SeriesInt32)
	i32s_ := NewSeriesInt32("test", true, false, []int32{1}).SetNullMask([]bool{true}).(SeriesInt32)
	i32v_ := NewSeriesInt32("test", true, false, []int32{1, 2, 3}).SetNullMask([]bool{true, true, false}).(SeriesInt32)

	i64s := NewSeriesInt64("test", true, false, []int64{1}).(SeriesInt64)
	i64v := NewSeriesInt64("test", true, false, []int64{1, 2, 3}).(SeriesInt64)
	i64s_ := NewSeriesInt64("test", true, false, []int64{1}).SetNullMask([]bool{true}).(SeriesInt64)
	i64v_ := NewSeriesInt64("test", true, false, []int64{1, 2, 3}).SetNullMask([]bool{true, true, false}).(SeriesInt64)

	f64s := NewSeriesFloat64("test", true, false, []float64{1.0}).(SeriesFloat64)
	f64v := NewSeriesFloat64("test", true, false, []float64{1.0, 2.0, 3.0}).(SeriesFloat64)
	f64s_ := NewSeriesFloat64("test", true, false, []float64{1.0}).SetNullMask([]bool{true}).(SeriesFloat64)
	f64v_ := NewSeriesFloat64("test", true, false, []float64{1.0, 2.0, 3.0}).SetNullMask([]bool{true, true, false}).(SeriesFloat64)

	pool := NewStringPool()
	ss := NewSeriesString("test", true, []string{"1"}, pool).(SeriesString)
	ss_ := NewSeriesString("test", true, []string{"1"}, pool).SetNullMask([]bool{true}).(SeriesString)
	sv := NewSeriesString("test", true, []string{"1", "2", "3"}, pool).(SeriesString)
	sv_ := NewSeriesString("test", true, []string{"1", "2", "3"}, pool).SetNullMask([]bool{true, true, false}).(SeriesString)

	// scalar | int32
	res = i32s.Add(i32s)
	if res.Data().([]int32)[0] != 2 {
		t.Errorf("Expected %v, got %v", []int32{2}, res.Data().([]int32))
	}

	res = i32s.Add(i32v)
	if res.Data().([]int32)[0] != 2 || res.Data().([]int32)[1] != 3 || res.Data().([]int32)[2] != 4 {
		t.Errorf("Expected %v, got %v", []int32{2, 3, 4}, res.Data().([]int32))
	}

	res = i32s.Add(i32s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i32s.Add(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | int64
	res = i64s.Add(i64s)
	if res.Data().([]int64)[0] != 2 {
		t.Errorf("Expected %v, got %v", []int64{2}, res.Data().([]int64))
	}

	res = i64s.Add(i64v)
	if res.Data().([]int64)[0] != 2 || res.Data().([]int64)[1] != 3 || res.Data().([]int64)[2] != 4 {
		t.Errorf("Expected %v, got %v", []int64{2, 3, 4}, res.Data().([]int64))
	}

	res = i64s.Add(i64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i64s.Add(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | float64
	res = f64s.Add(f64s)
	if res.Data().([]float64)[0] != 2.0 {
		t.Errorf("Expected %v, got %v", []float64{2.0}, res.Data().([]float64))
	}

	res = f64s.Add(f64v)
	if res.Data().([]float64)[0] != 2.0 || res.Data().([]float64)[1] != 3.0 || res.Data().([]float64)[2] != 4.0 {
		t.Errorf("Expected %v, got %v", []float64{2.0, 3.0, 4.0}, res.Data().([]float64))
	}

	res = f64s.Add(f64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Add(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | string
	res = f64s.Add(ss)
	if res.Data().([]string)[0] != "11" {
		t.Errorf("Expected %v, got %v", []string{"11"}, res.Data().([]string))
	}

	res = f64s.Add(sv)
	if res.Data().([]string)[0] != "11" || res.Data().([]string)[1] != "12" || res.Data().([]string)[2] != "13" {
		t.Errorf("Expected %v, got %v", []string{"11", "12", "13"}, res.Data().([]string))
	}

	res = f64s.Add(ss_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Add(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int32
	res = i32v.Add(i32s)
	if res.Data().([]int32)[0] != 2 || res.Data().([]int32)[1] != 3 || res.Data().([]int32)[2] != 4 {
		t.Errorf("Expected %v, got %v", []int32{2, 3, 4}, res.Data().([]int32))
	}

	res = i32v.Add(i32v)
	if res.Data().([]int32)[0] != 2 || res.Data().([]int32)[1] != 4 || res.Data().([]int32)[2] != 6 {
		t.Errorf("Expected %v, got %v", []int32{2, 4, 6}, res.Data().([]int32))
	}

	res = i32v.Add(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = i32v.Add(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | int64
	res = i64v.Add(i64s)
	if res.Data().([]int64)[0] != 2 || res.Data().([]int64)[1] != 3 || res.Data().([]int64)[2] != 4 {
		t.Errorf("Expected %v, got %v", []int64{2, 3, 4}, res.Data().([]int64))
	}

	res = i64v.Add(i64v)
	if res.Data().([]int64)[0] != 2 || res.Data().([]int64)[1] != 4 || res.Data().([]int64)[2] != 6 {
		t.Errorf("Expected %v, got %v", []int64{2, 4, 6}, res.Data().([]int64))
	}

	res = i64v.Add(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = i64v.Add(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | float64
	res = f64v.Add(f64s)
	if res.Data().([]float64)[0] != 2.0 || res.Data().([]float64)[1] != 3.0 || res.Data().([]float64)[2] != 4.0 {
		t.Errorf("Expected %v, got %v", []float64{2.0, 3.0, 4.0}, res.Data().([]float64))
	}

	res = f64v.Add(f64v)
	if res.Data().([]float64)[0] != 2.0 || res.Data().([]float64)[1] != 4.0 || res.Data().([]float64)[2] != 6.0 {
		t.Errorf("Expected %v, got %v", []float64{2.0, 4.0, 6.0}, res.Data().([]float64))
	}

	res = f64v.Add(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = f64v.Add(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | string
	res = f64v.Add(ss)
	if res.Data().([]string)[0] != "11" || res.Data().([]string)[1] != "21" || res.Data().([]string)[2] != "31" {
		t.Errorf("Expected %v, got %v", []string{"11", "21", "31"}, res.Data().([]string))
	}

	res = f64v.Add(ss_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = f64v.Add(sv)
	if res.Data().([]string)[0] != "11" || res.Data().([]string)[1] != "22" || res.Data().([]string)[2] != "33" {
		t.Errorf("Expected %v, got %v", []string{"11", "22", "33"}, res.Data().([]string))
	}

	res = f64v.Add(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}
}

func Benchmark_SeriesFloat64_Mul_SerScal_Perf(b *testing.B) {

	N := 1_000_000
	data := make([]float64, N)
	for i := 0; i < N; i++ {
		data[i] = float64(i)
	}

	ser := NewSeriesFloat64("test", true, false, data).(SeriesFloat64)
	scal := NewSeriesFloat64("test", true, false, []float64{1.5})

	// s * 1.5
	var res Series
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ser.Mul(scal)
	}

	if e, ok := res.(SeriesError); ok {
		b.Errorf("Got error: %v", e)
	}

	// Check the length.
	if res.Len() != N {
		b.Errorf("Expected length of %d, got %d", N, res.Len())
	}
}

func Benchmark_SeriesFloat64_Mul_SerSer_Perf(b *testing.B) {

	N := 1_000_000
	data1 := make([]float64, N)
	data2 := make([]float64, N)
	for i := 0; i < N; i++ {
		data1[i] = float64(i)
		data2[i] = float64(N - i - 1)
	}

	ser1 := NewSeriesFloat64("test", true, false, data1).(SeriesFloat64)
	ser2 := NewSeriesFloat64("test", true, false, data2)

	// s * 1.5
	var res Series
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ser1.Mul(ser2)
	}

	if e, ok := res.(SeriesError); ok {
		b.Errorf("Got error: %v", e)
	}

	// Check the length.
	if res.Len() != N {
		b.Errorf("Expected length of %d, got %d", N, res.Len())
	}
}

func Test_SeriesFloat64_Logical_Eq(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesFloat64_Logical_Ne(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesFloat64_Logical_Lt(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesFloat64_Logical_Le(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesFloat64_Logical_Gt(t *testing.T) {
	var res Series

	i32s := NewSeriesInt32("test", true, false, []int32{1}).(SeriesInt32)
	i32v := NewSeriesInt32("test", true, false, []int32{1, 2, 3}).(SeriesInt32)
	i32s_ := NewSeriesInt32("test", true, false, []int32{1}).SetNullMask([]bool{true}).(SeriesInt32)
	i32v_ := NewSeriesInt32("test", true, false, []int32{1, 2, 3}).SetNullMask([]bool{true, true, false}).(SeriesInt32)

	i64s := NewSeriesInt64("test", true, false, []int64{1}).(SeriesInt64)
	i64v := NewSeriesInt64("test", true, false, []int64{1, 2, 3}).(SeriesInt64)
	i64s_ := NewSeriesInt64("test", true, false, []int64{1}).SetNullMask([]bool{true}).(SeriesInt64)
	i64v_ := NewSeriesInt64("test", true, false, []int64{1, 2, 3}).SetNullMask([]bool{true, true, false}).(SeriesInt64)

	f64s := NewSeriesFloat64("test", true, false, []float64{1.0}).(SeriesFloat64)
	f64v := NewSeriesFloat64("test", true, false, []float64{1.0, 2.0, 3.0}).(SeriesFloat64)
	f64s_ := NewSeriesFloat64("test", true, false, []float64{1.0}).SetNullMask([]bool{true}).(SeriesFloat64)
	f64v_ := NewSeriesFloat64("test", true, false, []float64{1.0, 2.0, 3.0}).SetNullMask([]bool{true, true, false}).(SeriesFloat64)

	// scalar | int32
	res = f64s.Gt(i32s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = f64s.Gt(i32s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Gt(i32v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = f64s.Gt(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// scalar | int64
	res = f64s.Gt(i64s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = f64s.Gt(i64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Gt(i64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = f64s.Gt(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// scalar | float64
	res = f64s.Gt(f64s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = f64s.Gt(f64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = f64s.Gt(f64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = f64s.Gt(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | int32
	res = f64v.Gt(i32s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = f64v.Gt(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	res = f64v.Gt(i32v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = f64v.Gt(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | int64
	res = f64v.Gt(i64s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = f64v.Gt(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	res = f64v.Gt(i64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = f64v.Gt(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | float64
	res = f64v.Gt(f64s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = f64v.Gt(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	res = f64v.Gt(f64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = f64v.Gt(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}
}

func Test_SeriesFloat64_Logical_Ge(t *testing.T) {
	// TODO: add tests for all types
}
