package gandalff

import (
	"math"
	"math/rand"
	"testing"
	"typesys"
)

func Test_SeriesInt64_Base(t *testing.T) {

	data := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesInt64.
	s := NewSeriesInt64("test", true, true, data)

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
	if s.Type() != typesys.Int64Type {
		t.Errorf("Expected type of Int64Type, got %s", s.Type().ToString())
	}

	// Check the data.
	for i, v := range s.Data().([]int64) {
		if v != data[i] {
			t.Errorf("Expected data of []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, got %v", s.Data())
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
	for i := range s.Data().([]int64) {
		if s.IsNull(i) != mask[i] {
			t.Errorf("Expected IsNull(%d) to be %t, got %t", i, mask[i], s.IsNull(i))
		}
	}

	// Check the null count.
	if s.NullCount() != 3 {
		t.Errorf("Expected NullCount() to be 3, got %d", s.NullCount())
	}

	// Check the HasNull method.
	if !s.HasNull() {
		t.Errorf("Expected HasNull() to be true, got false")
	}

	// Check the SetNull method.
	for i := range s.Data().([]int64) {
		s.SetNull(i)
	}

	// Check the null values.
	for i := range s.Data().([]int64) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get method.
	for i := range s.Data().([]int64) {
		if s.Get(i).(int64) != data[i] {
			t.Errorf("Expected Get(%d) to be %d, got %d", i, data[i], s.Get(i).(int64))
		}
	}

	// Check the Set method.
	for i := range s.Data().([]int64) {
		s.Set(i, int64(i+10))
	}

	// Check the data.
	for i, v := range s.Data().([]int64) {
		if v != int64(i+10) {
			t.Errorf("Expected data of []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, got %v", s.Data())
		}
	}

	// Check the MakeNullable() method.
	p := NewSeriesInt64("test", false, true, data)

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
	p = p.MakeNullable().(SeriesInt64)

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

func Test_SeriesInt64_Take(t *testing.T) {

	data := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesInt64.
	s := NewSeriesInt64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Take the first 5 values.
	result := s.Take(0, 5, 1)

	// Check the length.
	if result.Len() != 5 {
		t.Errorf("Expected length of 5, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int64) {
		if v != data[i] {
			t.Errorf("Expected data of []int{1, 2, 3, 4, 5}, got %v", result.Data())
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of []bool{false, false, false, false, true}, got %v", result.GetNullMask())
		}
	}

	// Take the last 5 values.
	result = s.Take(5, 10, 1)

	// Check the length.
	if result.Len() != 5 {
		t.Errorf("Expected length of 5, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int64) {
		if v != data[i+5] {
			t.Errorf("Expected data of []int{6, 7, 8, 9, 10}, got %v", result.Data())
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i+5] {
			t.Errorf("Expected nullMask of []bool{true, false, false, true, false}, got %v", result.GetNullMask())
		}
	}

	// Take the first 5 values in steps of 2.
	result = s.Take(0, 5, 2)

	// Check the length.
	if result.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int64) {
		if v != data[i*2] {
			t.Errorf("Expected data of []int{1, 3, 5}, got %v", result.Data())
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i*2] {
			t.Errorf("Expected nullMask of []bool{false, false, true}, got %v", result.GetNullMask())
		}
	}

	// Take the last 5 values in steps of 2.
	result = s.Take(5, 10, 2)

	// Check the length.
	if result.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int64) {
		if v != data[i*2+5] {
			t.Errorf("Expected data of []int{6, 8, 10}, got %v", result.Data())
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i*2+5] {
			t.Errorf("Expected nullMask of []bool{true, false, false}, got %v", result.GetNullMask())
		}
	}
}

func Test_SeriesInt64_Append(t *testing.T) {
	dataA := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataB := []int64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	dataC := []int64{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesInt64("testA", true, true, dataA)
	sB := NewSeriesInt64("testB", true, true, dataB)
	sC := NewSeriesInt64("testC", true, true, dataC)

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
	for i, v := range result.Data().([]int64) {
		if i < 10 {
			if v != dataA[i] {
				t.Errorf("Expected %d, got %d at index %d", dataA[i], v, i)
			}
		} else if i < 20 {
			if v != dataB[i-10] {
				t.Errorf("Expected %d, got %d at index %d", dataB[i-10], v, i)
			}
		} else {
			if v != dataC[i-20] {
				t.Errorf("Expected %d, got %d at index %d", dataC[i-20], v, i)
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
	dataD := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sD := NewSeriesInt64("testD", true, true, dataD)

	// Check the original data.
	for i, v := range sD.Data().([]int64) {
		if v != dataD[i] {
			t.Errorf("Expected %d, got %d at index %d", dataD[i], v, i)
		}
	}

	for i := 0; i < 100; i++ {
		r := int64(rand.Intn(100))
		switch i % 4 {
		case 0:
			sD = sD.Append(r).(SeriesInt64)
		case 1:
			sD = sD.Append([]int64{r}).(SeriesInt64)
		case 2:
			sD = sD.Append(NullableInt64{true, r}).(SeriesInt64)
		case 3:
			sD = sD.Append([]NullableInt64{{false, r}}).(SeriesInt64)
		}

		if sD.Get(i+10) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
		}
	}
}

func Test_SeriesInt64_Cast(t *testing.T) {
	data := []int64{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesInt64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Cast to bool.
	result := s.Cast(typesys.BoolType, nil)

	// Check the data.
	expected := []bool{false, true, false, true, true, true, true, true, true, true}
	for i, v := range result.Data().([]bool) {
		if v != expected[i] {
			t.Errorf("Expected %t, got %t at index %d", expected[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to int32.
	result = s.Cast(typesys.Int32Type, nil)

	// Check the data.
	expectedInt32 := []int32{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range result.Data().([]int32) {
		if v != expectedInt32[i] {
			t.Errorf("Expected %d, got %d at index %d", expectedInt32[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to float64.
	result = s.Cast(typesys.Float64Type, nil)

	// Check the data.
	expectedFloat := []float64{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range result.Data().([]float64) {
		if v != expectedFloat[i] {
			t.Errorf("Expected %f, got %f at index %d", expectedFloat[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to string.
	result = s.Cast(typesys.StringType, NewStringPool())

	// Check the data.
	expectedString := []string{"0", "1", NULL_STRING, "3", "4", NULL_STRING, "6", "7", NULL_STRING, "9"}

	for i, v := range result.Data().([]string) {
		if v != expectedString[i] {
			t.Errorf("Expected %s, got %s at index %d", expectedString[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to error.
	castError := s.Cast(typesys.ErrorType, nil)

	// Check the message.
	if castError.(SeriesError).msg != "SeriesInt64.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_SeriesInt64_Filter(t *testing.T) {
	data := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewSeriesInt64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []int64{1, 3, 4, 6, 7, 9, 10, 11, 13, 14, 16, 17, 19, 20}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool("filter", false, true, filterMask).(SeriesBool))

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]int64) {
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
	for i, v := range filtered.Data().([]int64) {
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
	for i, v := range filtered.Data().([]int64) {
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

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesInt64.FilterByMask: mask length (20) does not match series length (14)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []int64{2, 323, 42, 4, 9, 674, 42, 48, 9811, 79, 3, 12, 492, 47005, -173, -28, 323, 42, 4, 9, 31, 425, 2}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesInt64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []int64{2, -28, 2}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]int64) {
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
	for i, v := range filtered.Data().([]int64) {
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

func Test_SeriesInt64_Map(t *testing.T) {
	data := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -2, 323, 24, -23, 4, 42, 5, -6, 7}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true}

	// Create a new series.
	s := NewSeriesInt64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(nullMask)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		if v.(int64) >= 7 && v.(int64) <= 100 {
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
		if v.(int64) < 0 {
			return int32(-(v.(int64)) % 7)
		}
		return int32(v.(int64) % 7)
	}, nil)

	expectedInt32 := []int32{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt.Data().([]int32) {
		if v != expectedInt32[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt32[i], v, i)
		}
	}

	// Map the series to float64.
	resFloat64 := s.Map(func(v any) any {
		if v.(int64) >= 0 {
			return float64(-v.(int64))
		}
		return float64(v.(int64))
	}, nil)

	expectedFloat64 := []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -2, -323, -24, -23, -4, -42, -5, -6, -7}
	for i, v := range resFloat64.Data().([]float64) {
		if v != expectedFloat64[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedFloat64[i], v, i)
		}
	}

	// Map the series to string.
	resString := s.Map(func(v any) any {
		if v.(int64) >= 0 {
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

func Test_SeriesInt64_Sort(t *testing.T) {

	data := []int64{2, 323, 42, 4, 9, 674, 42, 48, 9811, 79, 3, 12, 492, 47005, -173, -28, 323, 42, 4, 9, 31, 425, 2}
	mask := []bool{false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false}

	// Create a new series.
	s := NewSeriesInt64("test", true, true, data)

	// Sort the series.
	sorted := s.Sort()

	// Check the length.
	if sorted.Len() != 23 {
		t.Errorf("Expected length of 23, got %d", sorted.Len())
	}

	// Check the data.
	result := []int64{-173, -28, 2, 2, 3, 4, 4, 9, 9, 12, 31, 42, 42, 42, 48, 79, 323, 323, 425, 492, 674, 9811, 47005}
	for i, v := range sorted.Data().([]int64) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////

	// Create a new series.
	s = NewSeriesInt64("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Sort the series.
	sorted = s.Sort()

	// Check the length.
	if sorted.Len() != 23 {
		t.Errorf("Expected length of 23, got %d", sorted.Len())
	}

	// Check the data.
	result = []int64{-28, 2, 2, 3, 4, 4, 9, 9, 42, 48, 79, 323, 323, 425, 492, 47005, 42, 674, 9811, 12, -173, 42, 31}
	for i, v := range sorted.Data().([]int64) {
		if i < 16 && v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range sorted.GetNullMask() {
		if i < 16 && v != false {
			t.Errorf("Expected nullMask of %v, got %v at index %d", false, v, i)
		} else if i >= 16 && v != true {
			t.Errorf("Expected nullMask of %v, got %v at index %d", true, v, i)
		}
	}
}

func Test_SeriesInt64_GroupedSort(t *testing.T) {
	data := []int64{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// mask := []bool{false, true, false, false, false, false, true, false, false, false, false, true, false, false, false}

	partData := []int64{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	p := NewSeriesInt64("part", true, true, partData).Group()

	// Create a new series.
	s := NewSeriesInt64("test", true, true, data).
		SubGroup(p.GetPartition()).
		Sort()

	// Check the length.
	if s.Len() != 15 {
		t.Errorf("Expected length of 15, got %d", s.Len())
	}

	// Check the data.
	result := []int64{6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 11, 12, 13, 14, 15}
	for i, v := range s.Data().([]int64) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////

	// s = NewSeriesInt64("test", true, true, data).
	// 	SetNullMask(mask).
	// 	SubGroup(p.GetPartition()).
	// 	Sort()

	// // Check the length.
	// if s.Len() != 15 {
	// 	t.Errorf("Expected length of 15, got %d", s.Len())
	// }

	// // Check the data.
	// result = []int64{6, 7, 8, 10, 1, 2, 3, 5, 11, 12, 13, 15, 9, 4, 14}
	// for i, v := range s.Data().([]int64) {
	// 	if i < 14 && v != result[i] {
	// 		t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
	// 	}
	// }

	// // Check the null mask.
	// for i, v := range s.GetNullMask() {
	// 	if i < 12 && v != false {
	// 		t.Errorf("Expected nullMask of %v, got %v at index %d", false, v, i)
	// 	} else if i >= 12 && v != true {
	// 		t.Errorf("Expected nullMask of %v, got %v at index %d", true, v, i)
	// 	}
	// }

	/////////////////////////////////////////////////////////////////////////////////////
	// 								Reverse sort.

	dataRev := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// maskRev := []bool{false, true, false, false, false, false, true, false, false, false, false, true, false, false, false}

	s = NewSeriesInt64("test", true, true, dataRev).
		SubGroup(p.GetPartition()).
		SortRev()

	// Check the length.
	if s.Len() != 15 {
		t.Errorf("Expected length of 15, got %d", s.Len())
	}

	// Check the data.
	result = []int64{5, 4, 3, 2, 1, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6}
	for i, v := range s.Data().([]int64) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	///////////////////////////////////////////////////////////////////////////////////

	// s = NewSeriesInt64("test", true, true, dataRev).
	// 	SetNullMask(maskRev).
	// 	SubGroup(p.GetPartition()).
	// 	SortRev()

	// // Check the length.
	// if s.Len() != 15 {
	// 	t.Errorf("Expected length of 15, got %d", s.Len())
	// }

	// // Check the data.
	// result = []int64{5, 4, 3, 1, 10, 9, 8, 6, 15, 14, 13, 11, 2, 7, 12}
	// for i, v := range s.Data().([]int64) {
	// 	if i < 14 && v != result[i] {
	// 		t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
	// 	}
	// }

	// // Check the null mask.
	// for i, v := range s.GetNullMask() {
	// 	if i < 12 && v != false {
	// 		t.Errorf("Expected nullMask of %v, got %v at index %d", false, v, i)
	// 	} else if i >= 12 && v != true {
	// 		t.Errorf("Expected nullMask of %v, got %v at index %d", true, v, i)
	// 	}
	// }
}

// func Test_SeriesInt64_Multiplication(t *testing.T) {

// 	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

// 	// s * 2
// 	res := NewSeriesInt64("test", true, true, &data).Mul(NewSeriesInt64("test", true, true, &[]int{2}))
// 	if e, ok := res.(SeriesError); ok {
// 		t.Errorf("Got error: %v", e)
// 	}

// 	// Check the length.
// 	if res.Len() != 20 {
// 		t.Errorf("Expected length of 20, got %d", res.Len())
// 	}

// 	// Check the data.
// 	for i, v := range res.Data().([]int64) {
// 		if v != data[i]*2 {
// 			t.Errorf("Expected %v, got %v at index %d", data[i]*2, v, i)
// 		}
// 	}

// 	// 2 * s
// 	res = NewSeriesInt64("test", true, true, &[]int{2}).Mul(NewSeriesInt64("test", true, true, &data))
// 	if e, ok := res.(SeriesError); ok {
// 		t.Errorf("Got error: %v", e)
// 	}

// 	// Check the length.
// 	if res.Len() != 20 {
// 		t.Errorf("Expected length of 20, got %d", res.Len())
// 	}

// 	// Check the data.
// 	for i, v := range res.Data().([]int64) {
// 		if v != data[i]*2 {
// 			t.Errorf("Expected %v, got %v at index %d", data[i]*2, v, i)
// 		}
// 	}
// }

func Test_SeriesInt64_Arithmetic_Mul(t *testing.T) {
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
	if !checkEqSlice(i64s.Mul(bools).Data().([]int64), []int64{2}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(boolv).Data().([]int64), []int64{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(bools_).GetNullMask(), []bool{true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}

	// scalar | int32
	if !checkEqSlice(i64s.Mul(i32s).Data().([]int64), []int64{4}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(i32v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(i32s_).GetNullMask(), []bool{true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}

	// scalar | int64
	if !checkEqSlice(i64s.Mul(i64s).Data().([]int64), []int64{4}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(i64v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(i64s_).GetNullMask(), []bool{true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}

	// scalar | float64
	if !checkEqSlice(i64s.Mul(f64s).Data().([]float64), []float64{4}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(f64s_).GetNullMask(), []bool{true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64s.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}

	// vector | bool
	if !checkEqSlice(i64v.Mul(bools).Data().([]int64), []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(boolv).Data().([]int64), []int64{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}

	// vector | int32
	if !checkEqSlice(i64v.Mul(i32s).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(i32v).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}

	// vector | int64
	if !checkEqSlice(i64v.Mul(i64s).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(i64v).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}

	// vector | float64
	if !checkEqSlice(i64v.Mul(f64s).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(f64v).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
	if !checkEqSlice(i64v.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mul") {
		t.Errorf("Error in Int64 Mul")
	}
}

func Test_SeriesInt64_Arithmetic_Div(t *testing.T) {
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
	if !checkEqSlice(i64s.Div(bools).Data().([]float64), []float64{2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(boolv).Data().([]float64), []float64{2, math.Inf(1), 2, math.Inf(1), 2, math.Inf(1), 2, 2, math.Inf(1), math.Inf(1)}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(bools_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// scalar | int32
	if !checkEqSlice(i64s.Div(i32s).Data().([]float64), []float64{1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(i32v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(i32s_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// scalar | int64
	if !checkEqSlice(i64s.Div(i64s).Data().([]float64), []float64{1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(i64v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(i64s_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// scalar | float64
	if !checkEqSlice(i64s.Div(f64s).Data().([]float64), []float64{1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(f64v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(f64s_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64s.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | bool
	if !checkEqSlice(i64v.Div(bools).Data().([]float64), []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(boolv).Data().([]float64), []float64{1, math.Inf(1), 3, math.Inf(1), 5, math.Inf(1), 7, 8, math.Inf(1), math.Inf(1)}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | int32
	if !checkEqSlice(i64v.Div(i32s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(i32v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | int64
	if !checkEqSlice(i64v.Div(i64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(i64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | float64
	if !checkEqSlice(i64v.Div(f64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(f64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i64v.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
}

func Test_SeriesInt64_Arithmetic_Mod(t *testing.T) {
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
	if !checkEqSlice(i64s.Mod(bools).Data().([]float64), []float64{0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(bools_).GetNullMask(), []bool{true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}

	// scalar | int32
	if !checkEqSlice(i64s.Mod(i32s).Data().([]float64), []float64{0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(i32v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(i32s_).GetNullMask(), []bool{true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}

	// scalar | int64
	if !checkEqSlice(i64s.Mod(i64s).Data().([]float64), []float64{0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(i64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(i64s_).GetNullMask(), []bool{true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}

	// scalar | float64
	if !checkEqSlice(i64s.Mod(f64s).Data().([]float64), []float64{0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(f64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(f64s_).GetNullMask(), []bool{true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64s.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}

	// vector | bool
	if !checkEqSlice(i64v.Mod(bools).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}

	// vector | int32
	if !checkEqSlice(i64v.Mod(i32s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(i32v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}

	// vector | int64
	if !checkEqSlice(i64v.Mod(i64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(i64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}

	// vector | float64
	if !checkEqSlice(i64v.Mod(f64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(f64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
	if !checkEqSlice(i64v.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Mod") {
		t.Errorf("Error in Int64 Mod")
	}
}

func Test_SeriesInt64_Arithmetic_Pow(t *testing.T) {
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
	if !checkEqSlice(i64s.Pow(bools).Data().([]int64), []int64{2}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(boolv).Data().([]int64), []int64{2, 1, 2, 1, 2, 1, 2, 2, 1, 1}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(bools_).GetNullMask(), []bool{true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}

	// scalar | int32
	if !checkEqSlice(i64s.Pow(i32s).Data().([]int64), []int64{4}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(i32v).Data().([]int64), []int64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(i32s_).GetNullMask(), []bool{true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}

	// scalar | int64
	if !checkEqSlice(i64s.Pow(i64s).Data().([]int64), []int64{4}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(i64v).Data().([]int64), []int64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(i64s_).GetNullMask(), []bool{true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}

	// scalar | float64
	if !checkEqSlice(i64s.Pow(f64s).Data().([]float64), []float64{4}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(f64v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(f64s_).GetNullMask(), []bool{true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64s.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}

	// vector | bool
	if !checkEqSlice(i64v.Pow(bools).Data().([]int64), []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(boolv).Data().([]int64), []int64{1, 1, 3, 1, 5, 1, 7, 8, 1, 1}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}

	// vector | int32
	if !checkEqSlice(i64v.Pow(i32s).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(i32v).Data().([]int64), []int64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}

	// vector | int64
	if !checkEqSlice(i64v.Pow(i64s).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(i64v).Data().([]int64), []int64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}

	// vector | float64
	if !checkEqSlice(i64v.Pow(f64s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(f64v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
	if !checkEqSlice(i64v.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Pow") {
		t.Errorf("Error in Int64 Pow")
	}
}

func Test_SeriesInt64_Arithmetic_Add(t *testing.T) {
	pool := NewStringPool()

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

	ss := NewSeriesString("test", true, []string{"2"}, pool).(SeriesString)
	sv := NewSeriesString("test", true, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, pool).(SeriesString)
	ss_ := NewSeriesString("test", true, []string{"2"}, pool).SetNullMask([]bool{true}).(SeriesString)
	sv_ := NewSeriesString("test", true, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, pool).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesString)

	// scalar | bool
	if !checkEqSlice(i64s.Add(bools).Data().([]int64), []int64{3}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(boolv).Data().([]int64), []int64{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(bools_).GetNullMask(), []bool{true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// scalar | int32
	if !checkEqSlice(i64s.Add(i32s).Data().([]int64), []int64{4}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(i32v).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(i32s_).GetNullMask(), []bool{true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// scalar | int64
	if !checkEqSlice(i64s.Add(i64s).Data().([]int64), []int64{4}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(i64v).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(i64s_).GetNullMask(), []bool{true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// scalar | float64
	if !checkEqSlice(i64s.Add(f64s).Data().([]float64), []float64{4}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(f64v).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(f64s_).GetNullMask(), []bool{true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// scalar | string
	if !checkEqSlice(i64s.Add(ss).Data().([]string), []string{"22"}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(sv).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(ss_).GetNullMask(), []bool{true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64s.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// vector | bool
	if !checkEqSlice(i64v.Add(bools).Data().([]int64), []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(boolv).Data().([]int64), []int64{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// vector | int32
	if !checkEqSlice(i64v.Add(i32s).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(i32v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// vector | int64
	if !checkEqSlice(i64v.Add(i64s).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(i64v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// vector | float64
	if !checkEqSlice(i64v.Add(f64s).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}

	// vector | string
	if !checkEqSlice(i64v.Add(ss).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(sv).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
	if !checkEqSlice(i64v.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Add") {
		t.Errorf("Error in Int64 Add")
	}
}

func Test_SeriesInt64_Arithmetic_Sub(t *testing.T) {
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
	res = i64v.Sub(i32s)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 1 || res.Data().([]int64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []int64{0, 1, 2}, res.Data().([]int64))
	}

	res = i64v.Sub(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v.Sub(i32v)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 0 || res.Data().([]int64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []int64{0, 0, 0}, res.Data().([]int64))
	}

	res = i64v.Sub(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | int64
	res = i64v.Sub(i64s)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 1 || res.Data().([]int64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []int64{0, 1, 2}, res.Data().([]int64))
	}

	res = i64v.Sub(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v.Sub(i64v)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 0 || res.Data().([]int64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []int64{0, 0, 0}, res.Data().([]int64))
	}

	res = i64v.Sub(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | float64
	res = i64v.Sub(f64s)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 1 || res.Data().([]float64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []float64{0, 1, 2}, res.Data().([]float64))
	}

	res = i64v.Sub(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v.Sub(f64v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 0 || res.Data().([]float64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0, 0, 0}, res.Data().([]float64))
	}

	res = i64v.Sub(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int32
	res = i64v_.Sub(i32s)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 1 || res.Data().([]int64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []int64{0, 1, 2}, res.Data().([]int64))
	}

	res = i64v_.Sub(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v_.Sub(i32v)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 0 || res.Data().([]int64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []int64{0, 0, 0}, res.Data().([]int64))
	}

	res = i64v_.Sub(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int64
	res = i64v_.Sub(i64s)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 1 || res.Data().([]int64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []int64{0, 1, 2}, res.Data().([]int64))
	}

	res = i64v_.Sub(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v_.Sub(i64v)
	if res.Data().([]int64)[0] != 0 || res.Data().([]int64)[1] != 0 || res.Data().([]int64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []int64{0, 0, 0}, res.Data().([]int64))
	}

	res = i64v_.Sub(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | float64
	res = i64v_.Sub(f64s)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 1 || res.Data().([]float64)[2] != 2 {
		t.Errorf("Expected %v, got %v", []float64{0, 1, 2}, res.Data().([]float64))
	}

	res = i64v_.Sub(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v_.Sub(f64v)
	if res.Data().([]float64)[0] != 0 || res.Data().([]float64)[1] != 0 || res.Data().([]float64)[2] != 0 {
		t.Errorf("Expected %v, got %v", []float64{0, 0, 0}, res.Data().([]int64))
	}

	res = i64v_.Sub(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}
}

func Test_SeriesInt64_Logical_Eq(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt64_Logical_Ne(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt64_Logical_Lt(t *testing.T) {
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
	res = i64s.Lt(i32s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = i64s.Lt(i32s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i64s.Lt(i32v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i64s.Lt(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | int64
	res = i64s.Lt(i64s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = i64s.Lt(i64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i64s.Lt(i64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i64s.Lt(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | float64
	res = i64s.Lt(f64s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = i64s.Lt(f64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i64s.Lt(f64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i64s.Lt(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int32
	res = i64v.Lt(i32s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i64v.Lt(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v.Lt(i32v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i64v.Lt(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int64
	res = i64v.Lt(i64s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i64v.Lt(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v.Lt(i64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i64v.Lt(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | float64
	res = i64v.Lt(f64s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i64v.Lt(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i64v.Lt(f64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i64v.Lt(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}
}

func Test_SeriesInt64_Logical_Le(t *testing.T) {
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
	res = i64s.Le(i32s)
	if res.Data().([]bool)[0] != true {
		t.Errorf("Expected %v, got %v", []bool{true}, res.Data().([]bool))
	}

	res = i64s.Le(i32s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i64s.Le(i32v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i64s.Le(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | int64
	res = i64s.Le(i64s)
	if res.Data().([]bool)[0] != true {
		t.Errorf("Expected %v, got %v", []bool{true}, res.Data().([]bool))
	}

	res = i64s.Le(i64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i64s.Le(i64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i64s.Le(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | float64
	res = i64s.Le(f64s)
	if res.Data().([]bool)[0] != true {
		t.Errorf("Expected %v, got %v", []bool{true}, res.Data().([]bool))
	}

	res = i64s.Le(f64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i64s.Le(f64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i64s.Le(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int32
	res = i64v.Le(i32s)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i64v.Le(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = i64v.Le(i32v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i64v.Le(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int64
	res = i64v.Le(i64s)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i64v.Le(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.GetNullMask())
	}

	res = i64v.Le(i64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i64v.Le(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | float64
	res = i64v.Le(f64s)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i64v.Le(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.GetNullMask())
	}

	res = i64v.Le(f64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i64v.Le(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}
}

func Test_SeriesInt64_Logical_Gt(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt64_Logical_Ge(t *testing.T) {
	// TODO: add tests for all types
}
