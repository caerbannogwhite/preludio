package gandalff

import (
	"math"
	"math/rand"
	"preludiometa"
	"testing"
)

func Test_SeriesInt_Base(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesInt.
	s := NewSeriesInt(data, mask, true, ctx)

	// Check the length.
	if s.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", s.Len())
	}

	// Check the type.
	if s.Type() != preludiometa.IntType {
		t.Errorf("Expected type of IntType, got %s", s.Type().ToString())
	}

	// Check the data.
	for i, v := range s.Data().([]int) {
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
	for i := range s.Data().([]int) {
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

	// Check the Set() with a null value.
	for i := range s.Data().([]int) {
		s.Set(i, nil)
	}

	// Check the null values.
	for i := range s.Data().([]int) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get method.
	for i := range s.Data().([]int) {
		if s.Get(i).(int) != data[i] {
			t.Errorf("Expected Get(%d) to be %d, got %d", i, data[i], s.Get(i).(int))
		}
	}

	// Check the Set method.
	for i := range s.Data().([]int) {
		s.Set(i, i+10)
	}

	// Check the data.
	for i, v := range s.Data().([]int) {
		if v != int(i+10) {
			t.Errorf("Expected data of []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, got %v", s.Data())
		}
	}

	// Check the MakeNullable() method.
	p := NewSeriesInt(data, nil, true, ctx)

	// Check the nullability.
	if p.IsNullable() {
		t.Errorf("Expected IsNullable() to be false, got true")
	}

	// Set values to null.
	p.Set(1, nil)
	p.Set(3, nil)
	p.Set(5, nil)

	// Check the null count.
	if p.NullCount() != 0 {
		t.Errorf("Expected NullCount() to be 0, got %d", p.NullCount())
	}

	// Make the series nullable.
	p = p.MakeNullable().(SeriesInt)

	// Check the nullability.
	if !p.IsNullable() {
		t.Errorf("Expected IsNullable() to be true, got false")
	}

	// Check the null count.
	if p.NullCount() != 0 {
		t.Errorf("Expected NullCount() to be 0, got %d", p.NullCount())
	}

	p.Set(1, nil)
	p.Set(3, nil)
	p.Set(5, nil)

	// Check the null count.
	if p.NullCount() != 3 {
		t.Errorf("Expected NullCount() to be 3, got %d", p.NullCount())
	}
}

func Test_SeriesInt_Take(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesInt.
	s := NewSeriesInt(data, mask, true, ctx)

	// Take the first 5 values.
	result := s.Take(0, 5, 1)

	// Check the length.
	if result.Len() != 5 {
		t.Errorf("Expected length of 5, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int) {
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
	for i, v := range result.Data().([]int) {
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
	result = s.Take(0, 6, 2)

	// Check the length.
	if result.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int) {
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
	result = s.Take(5, 11, 2)

	// Check the length.
	if result.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int) {
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

func Test_SeriesInt_Append(t *testing.T) {
	dataA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataB := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	dataC := []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesInt(dataA, maskA, true, ctx)
	sB := NewSeriesInt(dataB, maskB, true, ctx)
	sC := NewSeriesInt(dataC, maskC, true, ctx)

	// Append the series.
	result := sA.Append(sB).Append(sC)

	// Check the length.
	if result.Len() != 30 {
		t.Errorf("Expected length of 30, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int) {
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

	// Append int, []int, NullableInt, []NullableInt
	s := NewSeriesInt([]int{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		r := int(rand.Intn(100))
		switch i % 4 {
		case 0:
			s = s.Append(r).(SeriesInt)
		case 1:
			s = s.Append([]int{r}).(SeriesInt)
		case 2:
			s = s.Append(NullableInt{true, r}).(SeriesInt)
		case 3:
			s = s.Append([]NullableInt{{false, r}}).(SeriesInt)
		}

		if s.Get(i) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, s.Get(i), i, i%4)
		}
	}

	// Append nil
	s = NewSeriesInt([]int{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(nil).(SeriesInt)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesNA
	s = NewSeriesInt([]int{}, nil, true, ctx)
	na := NewSeriesNA(10, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(na).(SeriesInt)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], na.GetNullMask(), nil, "SeriesInt.Append") {
			t.Errorf("Expected %v, got %v at index %d", na.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}

	// Append NullableInt
	s = NewSeriesInt([]int{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(NullableInt{false, 1}).(SeriesInt)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append []NullableInt
	s = NewSeriesInt([]int{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append([]NullableInt{{false, 1}}).(SeriesInt)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesInt
	s = NewSeriesInt([]int{}, nil, true, ctx)
	b := NewSeriesInt(dataA, []bool{true, true, true, true, true, true, true, true, true, true}, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(b).(SeriesInt)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], b.GetNullMask(), nil, "SeriesInt.Append") {
			t.Errorf("Expected %v, got %v at index %d", b.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}
}

func Test_SeriesInt_Cast(t *testing.T) {
	data := []int{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesInt(data, mask, true, ctx)

	// Cast to bool.
	result := s.Cast(preludiometa.BoolType)

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

	// Cast to int64.
	result = s.Cast(preludiometa.Int64Type)

	// Check the data.
	expectedInt := []int64{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range result.Data().([]int64) {
		if v != expectedInt[i] {
			t.Errorf("Expected %d, got %d at index %d", expectedInt[i], v, i)
		}
	}

	// Check the null mask.
	for i, v := range result.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to float64.
	result = s.Cast(preludiometa.Float64Type)

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
	result = s.Cast(preludiometa.StringType)

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
	castError := s.Cast(preludiometa.ErrorType)

	// Check the message.
	if castError.(SeriesError).msg != "SeriesInt.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_SeriesInt_Filter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewSeriesInt(data, mask, true, ctx)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []int{1, 3, 4, 6, 7, 9, 10, 11, 13, 14, 16, 17, 19, 20}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool(filterMask, nil, true, ctx))

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]int) {
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
	for i, v := range filtered.Data().([]int) {
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
	for i, v := range filtered.Data().([]int) {
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

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesInt.Filter: mask length (20) does not match series length (14)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []int{2, 323, 42, 4, 9, 674, 42, 48, 9811, 79, 3, 12, 492, 47005, -173, -28, 323, 42, 4, 9, 31, 425, 2}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesInt(data, mask, true, ctx)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []int{2, -28, 2}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]int) {
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
	for i, v := range filtered.Data().([]int) {
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

func Test_SeriesInt_Map(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -2, 323, 24, -23, 4, 42, 5, -6, 7}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true}

	// Create a new series.
	s := NewSeriesInt(data, nullMask, true, ctx)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		if v.(int) >= 7 && v.(int) <= 100 {
			return true
		}
		return false
	})

	expectedBool := []bool{false, false, false, false, false, false, true, true, true, true, false, false, true, false, false, true, false, false, true}
	for i, v := range resBool.Data().([]bool) {
		if v != expectedBool[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedBool[i], v, i)
		}
	}

	// Map the series to int.
	resInt := s.Map(func(v any) any {
		if v.(int) < 0 {
			return -(v.(int)) % 7
		}
		return v.(int) % 7
	})

	expectedInt := []int{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt.Data().([]int) {
		if v != expectedInt[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt[i], v, i)
		}
	}

	// Map the series to int64.
	resInt64 := s.Map(func(v any) any {
		if v.(int) >= 0 {
			return int64(v.(int) % 7)
		}
		return int64(-v.(int) % 7)
	})

	expectedInt64 := []int64{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt64.Data().([]int64) {
		if v != expectedInt64[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt64[i], v, i)
		}
	}

	// Map the series to float64.
	resFloat64 := s.Map(func(v any) any {
		if v.(int) >= 0 {
			return float64(-v.(int))
		}
		return float64(v.(int))
	})

	expectedFloat64 := []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -2, -323, -24, -23, -4, -42, -5, -6, -7}
	for i, v := range resFloat64.Data().([]float64) {
		if v != expectedFloat64[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedFloat64[i], v, i)
		}
	}

	// Map the series to string.
	resString := s.Map(func(v any) any {
		if v.(int) >= 0 {
			return "pos"
		}
		return "neg"
	})

	expectedString := []string{"pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "neg", "pos", "pos", "neg", "pos", "pos", "pos", "neg", "pos"}
	for i, v := range resString.Data().([]string) {
		if v != expectedString[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedString[i], v, i)
		}
	}
}

func Test_SeriesInt_Group(t *testing.T) {
	var partMap map[int64][]int

	data1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	data1Mask := []bool{false, false, false, false, false, true, true, true, true, true}
	data2 := []int{1, 1, 2, 2, 1, 1, 2, 2, 1, 1}
	data2Mask := []bool{false, true, false, true, false, true, false, true, false, true}
	data3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data3Mask := []bool{false, false, false, false, false, true, true, true, true, true}

	// Test 1
	s1 := NewSeriesInt(data1, data1Mask, true, ctx).
		group()

	p1 := s1.GetPartition().getMap()
	if len(p1) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(p1))
	}

	partMap = map[int64][]int{
		0: {0, 1, 2, 3, 4},
		1: {5, 6, 7, 8, 9},
	}
	if !checkEqPartitionMap(p1, partMap, nil, "Int Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p1)
	}

	// Test 2
	s2 := NewSeriesInt(data2, data2Mask, true, ctx).
		GroupBy(s1.GetPartition())

	p2 := s2.GetPartition().getMap()
	if len(p2) != 6 {
		t.Errorf("Expected 6 groups, got %d", len(p2))
	}

	partMap = map[int64][]int{
		0: {0, 4},
		1: {1, 3},
		2: {2},
		3: {5, 7, 9},
		4: {6},
		5: {8},
	}
	if !checkEqPartitionMap(p2, partMap, nil, "Int Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p2)
	}

	// Test 3
	s3 := NewSeriesInt(data3, data3Mask, true, ctx).
		GroupBy(s2.GetPartition())

	p3 := s3.GetPartition().getMap()
	if len(p3) != 8 {
		t.Errorf("Expected 8 groups, got %d", len(p3))
	}

	partMap = map[int64][]int{
		0: {0},
		1: {1},
		2: {2},
		3: {3},
		4: {4},
		5: {5, 7, 9},
		6: {6},
		7: {8},
	}
	if !checkEqPartitionMap(p3, partMap, nil, "Int Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p3)
	}

	// debugPrintPartition(s1.GetPartition(), s1)
	// debugPrintPartition(s2.GetPartition(), s1, s2)
	// debugPrintPartition(s3.GetPartition(), s1, s2, s3)

	partMap = nil
}

func Test_SeriesInt_Sort(t *testing.T) {
	data := []int{-195, -27, 33, 679, -67, 920, -352, -674, 250, 767, 697, 873, -802, -123, 308, -558, -518, 169, 313, 593}
	mask := []bool{false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true}

	// Create a new series.
	s := NewSeriesInt(data, nil, true, ctx)

	// Sort the series.
	sorted := s.Sort()

	// Check the data.
	expected := []int{-802, -674, -558, -518, -352, -195, -123, -67, -27, 33, 169, 250, 308, 313, 593, 679, 697, 767, 873, 920}
	if !checkEqSlice(sorted.Data().([]int), expected, nil, "") {
		t.Errorf("SeriesInt.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]int))
	}

	// Create a new series.
	s = NewSeriesInt(data, mask, true, ctx)

	// Sort the series.
	sorted = s.Sort()

	// Check the data.
	expected = []int{-802, -518, -352, -195, -67, 33, 250, 308, 313, 697, 920, 873, 767, -123, -674, -558, 679, 169, -27, 593}
	if !checkEqSlice(sorted.Data().([]int), expected, nil, "") {
		t.Errorf("SeriesInt.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]int))
	}

	// Check the null mask.
	expectedMask := []bool{false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true}
	if !checkEqSliceBool(sorted.GetNullMask(), expectedMask, nil, "") {
		t.Errorf("SeriesInt.Sort() failed, expecting %v, got %v", expectedMask, sorted.GetNullMask())
	}
}

func Test_SeriesInt_Arithmetic_Mul(t *testing.T) {
	bools := NewSeriesBool([]bool{true}, nil, true, ctx)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx)
	bools_ := NewSeriesBool([]bool{true}, nil, true, ctx).SetNullMask([]bool{true})
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i32s := NewSeriesInt([]int{2}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{2}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i64s := NewSeriesInt64([]int64{2}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{2}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	f64s := NewSeriesFloat64([]float64{2}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{2}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | bool
	if !checkEqSlice(i32s.Mul(bools).Data().([]int), []int{2}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(boolv).Data().([]int), []int{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(bools_).GetNullMask(), []bool{true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}

	// scalar | int
	if !checkEqSlice(i32s.Mul(i32s).Data().([]int), []int{4}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(i32v).Data().([]int), []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(i32s_).GetNullMask(), []bool{true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Mul(i64s).Data().([]int64), []int64{4}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(i64v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(i64s_).GetNullMask(), []bool{true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Mul(f64s).Data().([]float64), []float64{4}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(f64s_).GetNullMask(), []bool{true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32s.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}

	// vector | bool
	if !checkEqSlice(i32v.Mul(bools).Data().([]int), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(boolv).Data().([]int), []int{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}

	// vector | int
	if !checkEqSlice(i32v.Mul(i32s).Data().([]int), []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(i32v).Data().([]int), []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}

	// vector | int64
	if !checkEqSlice(i32v.Mul(i64s).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(i64v).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}

	// vector | float64
	if !checkEqSlice(i32v.Mul(f64s).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(f64v).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
	if !checkEqSlice(i32v.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mul") {
		t.Errorf("Error in Int Mul")
	}
}

func Test_SeriesInt_Arithmetic_Div(t *testing.T) {
	bools := NewSeriesBool([]bool{true}, nil, true, ctx)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx)
	bools_ := NewSeriesBool([]bool{true}, nil, true, ctx).SetNullMask([]bool{true})
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i32s := NewSeriesInt([]int{2}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{2}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i64s := NewSeriesInt64([]int64{2}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{2}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	f64s := NewSeriesFloat64([]float64{2}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{2}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | bool
	if !checkEqSlice(i32s.Div(bools).Data().([]float64), []float64{2}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(boolv).Data().([]float64), []float64{2, math.Inf(1), 2, math.Inf(1), 2, math.Inf(1), 2, 2, math.Inf(1), math.Inf(1)}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(bools_).GetNullMask(), []bool{true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}

	// scalar | int
	if !checkEqSlice(i32s.Div(i32s).Data().([]float64), []float64{1}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(i32v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(i32s_).GetNullMask(), []bool{true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Div(i64s).Data().([]float64), []float64{1}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(i64v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(i64s_).GetNullMask(), []bool{true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Div(f64s).Data().([]float64), []float64{1}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(f64v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(f64s_).GetNullMask(), []bool{true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32s.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}

	// vector | bool
	if !checkEqSlice(i32v.Div(bools).Data().([]float64), []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(boolv).Data().([]float64), []float64{1, math.Inf(1), 3, math.Inf(1), 5, math.Inf(1), 7, 8, math.Inf(1), math.Inf(1)}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}

	// vector | int
	if !checkEqSlice(i32v.Div(i32s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(i32v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}

	// vector | int64
	if !checkEqSlice(i32v.Div(i64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(i64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}

	// vector | float64
	if !checkEqSlice(i32v.Div(f64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(f64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
	if !checkEqSlice(i32v.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Div") {
		t.Errorf("Error in Int Div")
	}
}

func Test_SeriesInt_Arithmetic_Mod(t *testing.T) {
	bools := NewSeriesBool([]bool{true}, nil, true, ctx)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx)
	bools_ := NewSeriesBool([]bool{true}, nil, true, ctx).SetNullMask([]bool{true})
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i32s := NewSeriesInt([]int{2}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{2}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i64s := NewSeriesInt64([]int64{2}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{2}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	f64s := NewSeriesFloat64([]float64{2}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{2}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | bool
	if !checkEqSlice(i32s.Mod(bools).Data().([]float64), []float64{0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(bools_).GetNullMask(), []bool{true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}

	// scalar | int
	if !checkEqSlice(i32s.Mod(i32s).Data().([]float64), []float64{0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(i32v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(i32s_).GetNullMask(), []bool{true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Mod(i64s).Data().([]float64), []float64{0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(i64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(i64s_).GetNullMask(), []bool{true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Mod(f64s).Data().([]float64), []float64{0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(f64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(f64s_).GetNullMask(), []bool{true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32s.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}

	// vector | bool
	if !checkEqSlice(i32v.Mod(bools).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}

	// vector | int
	if !checkEqSlice(i32v.Mod(i32s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(i32v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}

	// vector | int64
	if !checkEqSlice(i32v.Mod(i64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(i64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}

	// vector | float64
	if !checkEqSlice(i32v.Mod(f64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(f64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
	if !checkEqSlice(i32v.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Mod") {
		t.Errorf("Error in Int Mod")
	}
}

func Test_SeriesInt_Arithmetic_Exp(t *testing.T) {
	bools := NewSeriesBool([]bool{true}, nil, true, ctx)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx)
	bools_ := NewSeriesBool([]bool{true}, nil, true, ctx).SetNullMask([]bool{true})
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i32s := NewSeriesInt([]int{2}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{2}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i64s := NewSeriesInt64([]int64{2}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{2}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	f64s := NewSeriesFloat64([]float64{2}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{2}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | bool
	if !checkEqSlice(i32s.Exp(bools).Data().([]int64), []int64{2}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(boolv).Data().([]int64), []int64{2, 1, 2, 1, 2, 1, 2, 2, 1, 1}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(bools_).GetNullMask(), []bool{true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}

	// scalar | int
	if !checkEqSlice(i32s.Exp(i32s).Data().([]int64), []int64{4}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(i32v).Data().([]int64), []int64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(i32s_).GetNullMask(), []bool{true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Exp(i64s).Data().([]int64), []int64{4}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(i64v).Data().([]int64), []int64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(i64s_).GetNullMask(), []bool{true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Exp(f64s).Data().([]float64), []float64{4}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(f64v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(f64s_).GetNullMask(), []bool{true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32s.Exp(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}

	// vector | bool
	if !checkEqSlice(i32v.Exp(bools).Data().([]int64), []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(boolv).Data().([]int64), []int64{1, 1, 3, 1, 5, 1, 7, 8, 1, 1}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}

	// vector | int
	if !checkEqSlice(i32v.Exp(i32s).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(i32v).Data().([]int64), []int64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}

	// vector | int64
	if !checkEqSlice(i32v.Exp(i64s).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(i64v).Data().([]int64), []int64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}

	// vector | float64
	if !checkEqSlice(i32v.Exp(f64s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(f64v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
	if !checkEqSlice(i32v.Exp(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Exp") {
		t.Errorf("Error in Int Exp")
	}
}

func Test_SeriesInt_Arithmetic_Add(t *testing.T) {
	bools := NewSeriesBool([]bool{true}, nil, true, ctx)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx)
	bools_ := NewSeriesBool([]bool{true}, nil, true, ctx).SetNullMask([]bool{true})
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i32s := NewSeriesInt([]int{2}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{2}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i64s := NewSeriesInt64([]int64{2}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{2}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	f64s := NewSeriesFloat64([]float64{2}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{2}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	ss := NewSeriesString([]string{"2"}, nil, true, ctx)
	sv := NewSeriesString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, true, ctx)
	ss_ := NewSeriesString([]string{"2"}, nil, true, ctx).SetNullMask([]bool{true})
	sv_ := NewSeriesString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | bool
	if !checkEqSlice(i32s.Add(bools).Data().([]int), []int{3}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(boolv).Data().([]int), []int{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(bools_).GetNullMask(), []bool{true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// scalar | int
	if !checkEqSlice(i32s.Add(i32s).Data().([]int), []int{4}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(i32v).Data().([]int), []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(i32s_).GetNullMask(), []bool{true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Add(i64s).Data().([]int64), []int64{4}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(i64v).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(i64s_).GetNullMask(), []bool{true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Add(f64s).Data().([]float64), []float64{4}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(f64v).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(f64s_).GetNullMask(), []bool{true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// scalar | string
	if !checkEqSlice(i32s.Add(ss).Data().([]string), []string{"22"}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(sv).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(ss_).GetNullMask(), []bool{true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32s.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// vector | bool
	if !checkEqSlice(i32v.Add(bools).Data().([]int), []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(boolv).Data().([]int), []int{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// vector | int
	if !checkEqSlice(i32v.Add(i32s).Data().([]int), []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(i32v).Data().([]int), []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// vector | int64
	if !checkEqSlice(i32v.Add(i64s).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(i64v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// vector | float64
	if !checkEqSlice(i32v.Add(f64s).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}

	// vector | string
	if !checkEqSlice(i32v.Add(ss).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(sv).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
	if !checkEqSlice(i32v.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int Add") {
		t.Errorf("Error in Int Add")
	}
}

func Test_SeriesInt_Arithmetic_Sub(t *testing.T) {
	bools := NewSeriesBool([]bool{true}, nil, true, ctx)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx)
	bools_ := NewSeriesBool([]bool{true}, nil, true, ctx).SetNullMask([]bool{true})
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i32s := NewSeriesInt([]int{2}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{2}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	i64s := NewSeriesInt64([]int64{2}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{2}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	f64s := NewSeriesFloat64([]float64{2}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{2}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | bool
	if !checkEqSlice(i32s.Sub(bools).Data().([]int), []int{1}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(boolv).Data().([]int), []int{1, 2, 1, 2, 1, 2, 1, 1, 2, 2}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(bools_).GetNullMask(), []bool{true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// scalar | int
	if !checkEqSlice(i32s.Sub(i32s).Data().([]int), []int{0}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(i32v).Data().([]int), []int{1, 0, -1, -2, -3, -4, -5, -6, -7, -8}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(i32s_).GetNullMask(), []bool{true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Sub(i64s).Data().([]int64), []int64{0}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(i64v).Data().([]int64), []int64{1, 0, -1, -2, -3, -4, -5, -6, -7, -8}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(i64s_).GetNullMask(), []bool{true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Sub(f64s).Data().([]float64), []float64{0}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(f64v).Data().([]float64), []float64{1, 0, -1, -2, -3, -4, -5, -6, -7, -8}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(f64s_).GetNullMask(), []bool{true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// vector | bool
	if !checkEqSlice(i32v.Sub(bools).Data().([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(boolv).Data().([]int), []int{0, 2, 2, 4, 4, 6, 6, 7, 9, 10}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// vector | int
	if !checkEqSlice(i32v.Sub(i32s).Data().([]int), []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(i32v).Data().([]int), []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// vector | int64
	if !checkEqSlice(i32v.Sub(i64s).Data().([]int64), []int64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(i64v).Data().([]int64), []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// vector | float64
	if !checkEqSlice(i32v.Sub(f64s).Data().([]float64), []float64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(f64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
}

func Test_SeriesInt_Logical_Eq(t *testing.T) {
	var res Series

	i32s := NewSeriesInt([]int{1}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{1}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	i64s := NewSeriesInt64([]int64{1}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{1}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	f64s := NewSeriesFloat64([]float64{1.0}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1.0, 2.0, 3.0}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{1.0}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1.0, 2.0, 3.0}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	// scalar | int
	res = i32s.Eq(i32s)
	if res.Data().([]bool)[0] != true {
		t.Errorf("Expected %v, got %v", true, res.Data().([]bool)[0])
	}

	res = i32s.Eq(i32s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i32s.Eq(i32v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i32s.Eq(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | int64
	res = i32s.Eq(i64s)
	if res.Data().([]bool)[0] != true {
		t.Errorf("Expected %v, got %v", true, res.Data().([]bool)[0])
	}

	res = i32s.Eq(i64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i32s.Eq(i64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i32s.Eq(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// scalar | float64
	res = i32s.Eq(f64s)
	if res.Data().([]bool)[0] != true {
		t.Errorf("Expected %v, got %v", true, res.Data().([]bool)[0])
	}

	res = i32s.Eq(f64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{true}, res.GetNullMask())
	}

	res = i32s.Eq(f64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i32s.Eq(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int
	res = i32v.Eq(i32s)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i32v.Eq(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	res = i32v.Eq(i32v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i32v.Eq(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | int64
	res = i32v.Eq(i64s)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i32v.Eq(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i32v.Eq(i64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i32v.Eq(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}

	// vector | float64
	res = i32v.Eq(f64s)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = i32v.Eq(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.GetNullMask())
	}

	res = i32v.Eq(f64v)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = i32v.Eq(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, res.GetNullMask())
	}
}

func Test_SeriesInt_Logical_Ne(t *testing.T) {
	var res Series

	i32s := NewSeriesInt([]int{1}, nil, true, ctx)
	i32v := NewSeriesInt([]int{1, 2, 3}, nil, true, ctx)
	i32s_ := NewSeriesInt([]int{1}, nil, true, ctx).SetNullMask([]bool{true})
	i32v_ := NewSeriesInt([]int{1, 2, 3}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	i64s := NewSeriesInt64([]int64{1}, nil, true, ctx)
	i64v := NewSeriesInt64([]int64{1, 2, 3}, nil, true, ctx)
	i64s_ := NewSeriesInt64([]int64{1}, nil, true, ctx).SetNullMask([]bool{true})
	i64v_ := NewSeriesInt64([]int64{1, 2, 3}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	f64s := NewSeriesFloat64([]float64{1.0}, nil, true, ctx)
	f64v := NewSeriesFloat64([]float64{1.0, 2.0, 3.0}, nil, true, ctx)
	f64s_ := NewSeriesFloat64([]float64{1.0}, nil, true, ctx).SetNullMask([]bool{true})
	f64v_ := NewSeriesFloat64([]float64{1.0, 2.0, 3.0}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	// scalar | int
	res = i32s.Ne(i32s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = i32s.Ne(i32s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.GetNullMask())
	}

	res = i32s.Ne(i32v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i32s.Ne(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// scalar | int64
	res = i32s.Ne(i64s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = i32s.Ne(i64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.GetNullMask())
	}

	res = i32s.Ne(i64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i32s.Ne(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// scalar | float64
	res = i32s.Ne(f64s)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.Data().([]bool))
	}

	res = i32s.Ne(f64s_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", []bool{false}, res.GetNullMask())
	}

	res = i32s.Ne(f64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i32s.Ne(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | int
	res = i32v.Ne(i32s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i32v.Ne(i32s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = i32v.Ne(i32v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i32v.Ne(i32v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | int64
	res = i32v.Ne(i64s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i32v.Ne(i64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = i32v.Ne(i64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i32v.Ne(i64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}

	// vector | float64
	res = i32v.Ne(f64s)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = i32v.Ne(f64s_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.GetNullMask())
	}

	res = i32v.Ne(f64v)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = i32v.Ne(f64v_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, true}, res.GetNullMask())
	}
}

func Test_SeriesInt_Logical_Lt(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt_Logical_Le(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt_Logical_Gt(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt_Logical_Ge(t *testing.T) {
	// TODO: add tests for all types
}
