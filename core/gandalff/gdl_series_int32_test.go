package gandalff

import (
	"math/rand"
	"testing"
	"typesys"
)

func Test_GDLSeriesInt32_Base(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GDLSeriesInt32.
	s := NewGDLSeriesInt32("test", true, true, data)

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
	if s.Type() != typesys.Int32Type {
		t.Errorf("Expected type of Int32Type, got %s", s.Type().ToString())
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

	// Check the SetNull method.
	for i := range s.Data().([]int) {
		s.SetNull(i)
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
		if v != i+10 {
			t.Errorf("Expected data of []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, got %v", s.Data())
		}
	}

	// Check the MakeNullable() method.
	p := NewGDLSeriesInt32("test", false, true, data)

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
	p = p.MakeNullable().(GDLSeriesInt32)

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

func Test_GDLSeriesInt32_Take(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GDLSeriesInt32.
	s := NewGDLSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

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
	result = s.Take(0, 5, 2)

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
	result = s.Take(5, 10, 2)

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

func Test_GDLSeriesInt32_Append(t *testing.T) {
	dataA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataB := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	dataC := []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewGDLSeriesInt32("testA", true, true, dataA)
	sB := NewGDLSeriesInt32("testB", true, true, dataB)
	sC := NewGDLSeriesInt32("testC", true, true, dataC)

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

	// Append random values.
	dataD := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sD := NewGDLSeriesInt32("testD", true, true, dataD)

	// Check the original data.
	for i, v := range sD.Data().([]int) {
		if v != dataD[i] {
			t.Errorf("Expected %d, got %d at index %d", dataD[i], v, i)
		}
	}

	for i := 0; i < 100; i++ {
		r := rand.Intn(100)
		switch i % 4 {
		case 0:
			sD = sD.Append(r).(GDLSeriesInt32)
		case 1:
			sD = sD.Append([]int{r}).(GDLSeriesInt32)
		case 2:
			sD = sD.Append(NullableInt32{true, r}).(GDLSeriesInt32)
		case 3:
			sD = sD.Append([]NullableInt32{{false, r}}).(GDLSeriesInt32)
		}

		if sD.Get(i+10) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
		}
	}
}

func Test_GDLSeriesInt32_Cast(t *testing.T) {
	data := []int{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewGDLSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Cast to bool.
	result := s.Cast(typesys.BoolType, nil)

	// Check the length.
	if result.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", result.Len())
	}

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

	// Cast to float64.
	result = s.Cast(typesys.Float64Type, nil)

	// Check the length.
	if result.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", result.Len())
	}

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

	// Check the length.
	if result.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", result.Len())
	}

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
	if castError.(GDLSeriesError).msg != "GDLSeriesInt32.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_GDLSeriesInt32_Filter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewGDLSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []int{1, 3, 4, 6, 7, 9, 10, 11, 13, 14, 16, 17, 19, 20}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewGDLSeriesBool("filter", false, filterMask).(GDLSeriesBool))

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
	// 							Check the FilterByMask() method.
	filtered = s.FilterByMask(filterMask)

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
	// 							Check the FilterByIndeces() method.
	filtered = s.FilterByIndeces(filterIndeces)

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
	filtered = filtered.FilterByMask(filterMask)

	if e, ok := filtered.(GDLSeriesError); !ok || e.Error() != "GDLSeriesInt32.FilterByMask: mask length (20) does not match series length (14)" {
		t.Errorf("Expected GDLSeriesError, got %v", filtered)
	}

	// Another test.
	data = []int{2, 323, 42, 4, 9, 674, 42, 48, 9811, 79, 3, 12, 492, 47005, -173, -28, 323, 42, 4, 9, 31, 425, 2}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewGDLSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []int{2, -28, 2}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the FilterByMask() method.
	filtered = s.FilterByMask(filterMask)

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
	// 							Check the FilterByIndeces() method.
	filtered = s.FilterByIndeces(filterIndeces)

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

func Test_GDLSeriesInt32_Map(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -2, 323, 24, -23, 4, 42, 5, -6, 7}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true}

	// Create a new series.
	s := NewGDLSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(nullMask)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		if v.(int) >= 7 && v.(int) <= 100 {
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

	// Map the series to int.
	resInt := s.Map(func(v any) any {
		if v.(int) < 0 {
			return -(v.(int)) % 7
		}
		return v.(int) % 7
	}, nil)

	expectedInt := []int{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt.Data().([]int) {
		if v != expectedInt[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt[i], v, i)
		}
	}

	// Map the series to float64.
	resFloat64 := s.Map(func(v any) any {
		if v.(int) >= 0 {
			return float64(-v.(int))
		}
		return float64(v.(int))
	}, nil)

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
	}, NewStringPool())

	expectedString := []string{"pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "neg", "pos", "pos", "neg", "pos", "pos", "pos", "neg", "pos"}
	for i, v := range resString.Data().([]string) {
		if v != expectedString[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedString[i], v, i)
		}
	}
}

func Test_GDLSeriesInt32_Sort(t *testing.T) {

	data := []int{2, 323, 42, 4, 9, 674, 42, 48, 9811, 79, 3, 12, 492, 47005, -173, -28, 323, 42, 4, 9, 31, 425, 2}
	mask := []bool{false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false}

	// Create a new series.
	s := NewGDLSeriesInt32("test", true, true, data)

	// Sort the series.
	sorted := s.Sort()

	// Check the length.
	if sorted.Len() != 23 {
		t.Errorf("Expected length of 23, got %d", sorted.Len())
	}

	// Check the data.
	result := []int{-173, -28, 2, 2, 3, 4, 4, 9, 9, 12, 31, 42, 42, 42, 48, 79, 323, 323, 425, 492, 674, 9811, 47005}
	for i, v := range sorted.Data().([]int) {
		if v != result[i] {
			t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////

	// Create a new series.
	s = NewGDLSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Sort the series.
	sorted = s.Sort()

	// Check the length.
	if sorted.Len() != 23 {
		t.Errorf("Expected length of 23, got %d", sorted.Len())
	}

	// Check the data.
	result = []int{-28, 2, 2, 3, 4, 4, 9, 9, 42, 48, 79, 323, 323, 425, 492, 47005, 42, 674, 9811, 12, -173, 42, 31}
	for i, v := range sorted.Data().([]int) {
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

func Test_GDLSeriesInt32_GroupedSort(t *testing.T) {
	// data := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// mask := []bool{false, true, false, false, false, false, true, false, false, false, false, true, false, false, false}

	// partData := []int{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	// p := NewGDLSeriesInt32("part", true, true, partData).Group()

	// // Create a new series.
	// s := NewGDLSeriesInt32("test", true, true, data).
	// 	SubGroup(p.GetPartition()).
	// 	Sort()

	// // Check the length.
	// if s.Len() != 15 {
	// 	t.Errorf("Expected length of 15, got %d", s.Len())
	// }

	// // Check the data.
	// result := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 11, 12, 13, 14, 15}
	// for i, v := range s.Data().([]int) {
	// 	if v != result[i] {
	// 		t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
	// 	}
	// }

	// /////////////////////////////////////////////////////////////////////////////////////

	// s = NewGDLSeriesInt32("test", true, true, data).
	// 	SetNullMask(mask).
	// 	SubGroup(p.GetPartition()).
	// 	Sort()

	// // Check the length.
	// if s.Len() != 15 {
	// 	t.Errorf("Expected length of 15, got %d", s.Len())
	// }

	// // Check the data.
	// result = []int{6, 7, 8, 10, 1, 2, 3, 5, 11, 12, 13, 15, 9, 4, 14}
	// for i, v := range s.Data().([]int) {
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

	// /////////////////////////////////////////////////////////////////////////////////////
	// // 								Reverse sort.

	// dataRev := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// // maskRev := []bool{false, true, false, false, false, false, true, false, false, false, false, true, false, false, false}

	// s = NewGDLSeriesInt32("test", true, true, dataRev).
	// 	SubGroup(p.GetPartition()).
	// 	SortRev()

	// // Check the length.
	// if s.Len() != 15 {
	// 	t.Errorf("Expected length of 15, got %d", s.Len())
	// }

	// // Check the data.
	// result = []int{5, 4, 3, 2, 1, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6}
	// for i, v := range s.Data().([]int) {
	// 	if v != result[i] {
	// 		t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
	// 	}
	// }

	/////////////////////////////////////////////////////////////////////////////////////

	// s = NewGDLSeriesInt32("test", true, true, dataRev).
	// 	SetNullMask(maskRev).
	// 	SubGroup(p.GetPartition()).
	// 	SortRev()

	// // Check the length.
	// if s.Len() != 15 {
	// 	t.Errorf("Expected length of 15, got %d", s.Len())
	// }

	// // Check the data.
	// result = []int{5, 4, 3, 1, 10, 9, 8, 6, 15, 14, 13, 11, 2, 7, 12}
	// for i, v := range s.Data().([]int) {
	// 	if i < 14 && v != result[i] {
	// 		t.Errorf("Expected %v, got %v at index %d", result[i], v, i)
	// 	}
	// }

	// Check the null mask.
	// for i, v := range s.GetNullMask() {
	// 	if i < 12 && v != false {
	// 		t.Errorf("Expected nullMask of %v, got %v at index %d", false, v, i)
	// 	} else if i >= 12 && v != true {
	// 		t.Errorf("Expected nullMask of %v, got %v at index %d", true, v, i)
	// 	}
	// }
}

// func Test_GDLSeriesInt32_Multiplication(t *testing.T) {

// 	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

// 	// s * 2
// 	res := NewGDLSeriesInt32("test", true, true, &data).Mul(NewGDLSeriesInt32("test", true, true, &[]int{2}))
// 	if e, ok := res.(GDLSeriesError); ok {
// 		t.Errorf("Got error: %v", e)
// 	}

// 	// Check the length.
// 	if res.Len() != 20 {
// 		t.Errorf("Expected length of 20, got %d", res.Len())
// 	}

// 	// Check the data.
// 	for i, v := range res.Data().([]int) {
// 		if v != data[i]*2 {
// 			t.Errorf("Expected %v, got %v at index %d", data[i]*2, v, i)
// 		}
// 	}

// 	// 2 * s
// 	res = NewGDLSeriesInt32("test", true, true, &[]int{2}).Mul(NewGDLSeriesInt32("test", true, true, &data))
// 	if e, ok := res.(GDLSeriesError); ok {
// 		t.Errorf("Got error: %v", e)
// 	}

// 	// Check the length.
// 	if res.Len() != 20 {
// 		t.Errorf("Expected length of 20, got %d", res.Len())
// 	}

// 	// Check the data.
// 	for i, v := range res.Data().([]int) {
// 		if v != data[i]*2 {
// 			t.Errorf("Expected %v, got %v at index %d", data[i]*2, v, i)
// 		}
// 	}
// }
