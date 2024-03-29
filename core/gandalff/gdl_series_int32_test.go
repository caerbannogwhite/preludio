package gandalff

import (
	"math"
	"math/rand"
	"testing"
	"typesys"
)

func Test_SeriesInt32_Base(t *testing.T) {

	data := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesInt32.
	s := NewSeriesInt32("test", true, true, data)

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
	for i, v := range s.Data().([]int32) {
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
	for i := range s.Data().([]int32) {
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
	for i := range s.Data().([]int32) {
		s.SetNull(i)
	}

	// Check the null values.
	for i := range s.Data().([]int32) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get method.
	for i := range s.Data().([]int32) {
		if s.Get(i).(int32) != data[i] {
			t.Errorf("Expected Get(%d) to be %d, got %d", i, data[i], s.Get(i).(int))
		}
	}

	// Check the Set method.
	for i := range s.Data().([]int32) {
		s.Set(i, i+10)
	}

	// Check the data.
	for i, v := range s.Data().([]int32) {
		if v != int32(i+10) {
			t.Errorf("Expected data of []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, got %v", s.Data())
		}
	}

	// Check the MakeNullable() method.
	p := NewSeriesInt32("test", false, true, data)

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
	p = p.MakeNullable().(SeriesInt32)

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

func Test_SeriesInt32_Take(t *testing.T) {

	data := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesInt32.
	s := NewSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Take the first 5 values.
	result := s.Take(0, 5, 1)

	// Check the length.
	if result.Len() != 5 {
		t.Errorf("Expected length of 5, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int32) {
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
	for i, v := range result.Data().([]int32) {
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
	for i, v := range result.Data().([]int32) {
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
	for i, v := range result.Data().([]int32) {
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

func Test_SeriesInt32_Append(t *testing.T) {
	dataA := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataB := []int32{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	dataC := []int32{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesInt32("testA", true, true, dataA)
	sB := NewSeriesInt32("testB", true, true, dataB)
	sC := NewSeriesInt32("testC", true, true, dataC)

	// Set the null masks.
	sA.SetNullMask(maskA)
	sB.SetNullMask(maskB)
	sC.SetNullMask(maskC)

	// Append the series.
	result := sA.Append(sB).Append(sC)

	// Check the name.
	if result.Name() != "testA" {
		t.Errorf("Expected name of \"testA\", got %s", result.Name())
	}

	// Check the length.
	if result.Len() != 30 {
		t.Errorf("Expected length of 30, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]int32) {
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
	dataD := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sD := NewSeriesInt32("testD", true, true, dataD)

	// Check the original data.
	for i, v := range sD.Data().([]int32) {
		if v != dataD[i] {
			t.Errorf("Expected %d, got %d at index %d", dataD[i], v, i)
		}
	}

	for i := 0; i < 100; i++ {
		r := int32(rand.Intn(100))
		switch i % 4 {
		case 0:
			sD = sD.Append(r).(SeriesInt32)
		case 1:
			sD = sD.Append([]int32{r}).(SeriesInt32)
		case 2:
			sD = sD.Append(NullableInt32{true, r}).(SeriesInt32)
		case 3:
			sD = sD.Append([]NullableInt32{{false, r}}).(SeriesInt32)
		}

		if sD.Get(i+10) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
		}
	}
}

func Test_SeriesInt32_Cast(t *testing.T) {
	data := []int32{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesInt32("test", true, true, data)

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

	// Cast to int64.
	result = s.Cast(typesys.Int64Type, nil)

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
	if castError.(SeriesError).msg != "SeriesInt32.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_SeriesInt32_Filter(t *testing.T) {
	data := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []int32{1, 3, 4, 6, 7, 9, 10, 11, 13, 14, 16, 17, 19, 20}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool("filter", false, true, filterMask).(SeriesBool))

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]int32) {
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
	for i, v := range filtered.Data().([]int32) {
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
	for i, v := range filtered.Data().([]int32) {
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

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesInt32.FilterByMask: mask length (20) does not match series length (14)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []int32{2, 323, 42, 4, 9, 674, 42, 48, 9811, 79, 3, 12, 492, 47005, -173, -28, 323, 42, 4, 9, 31, 425, 2}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []int32{2, -28, 2}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]int32) {
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
	for i, v := range filtered.Data().([]int32) {
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

func Test_SeriesInt32_Map(t *testing.T) {
	data := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -2, 323, 24, -23, 4, 42, 5, -6, 7}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true}

	// Create a new series.
	s := NewSeriesInt32("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(nullMask)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		if v.(int32) >= 7 && v.(int32) <= 100 {
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
		if v.(int32) < 0 {
			return -(v.(int32)) % 7
		}
		return v.(int32) % 7
	}, nil)

	expectedInt32 := []int32{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt.Data().([]int32) {
		if v != expectedInt32[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt32[i], v, i)
		}
	}

	// Map the series to int64.
	resInt64 := s.Map(func(v any) any {
		if v.(int32) >= 0 {
			return int64(v.(int32) % 7)
		}
		return int64(-v.(int32) % 7)
	}, nil)

	expectedInt64 := []int64{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt64.Data().([]int64) {
		if v != expectedInt64[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt64[i], v, i)
		}
	}

	// Map the series to float64.
	resFloat64 := s.Map(func(v any) any {
		if v.(int32) >= 0 {
			return float64(-v.(int32))
		}
		return float64(v.(int32))
	}, nil)

	expectedFloat64 := []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -2, -323, -24, -23, -4, -42, -5, -6, -7}
	for i, v := range resFloat64.Data().([]float64) {
		if v != expectedFloat64[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedFloat64[i], v, i)
		}
	}

	// Map the series to string.
	resString := s.Map(func(v any) any {
		if v.(int32) >= 0 {
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

func Test_SeriesInt32_Group(t *testing.T) {
	var partMap map[int64][]int

	data1 := []int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	data1Mask := []bool{false, false, false, false, false, true, true, true, true, true}
	data2 := []int32{1, 1, 2, 2, 1, 1, 2, 2, 1, 1}
	data2Mask := []bool{false, true, false, true, false, true, false, true, false, true}
	data3 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data3Mask := []bool{false, false, false, false, false, true, true, true, true, true}

	// Test 1
	s1 := NewSeriesInt32("s1", true, true, data1).
		SetNullMask(data1Mask).
		group()

	p1 := s1.GetPartition().getMap()
	if len(p1) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(p1))
	}

	partMap = map[int64][]int{
		0: {0, 1, 2, 3, 4},
		1: {5, 6, 7, 8, 9},
	}
	if !checkEqPartitionMap(p1, partMap, nil, "Int32 Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p1)
	}

	// Test 2
	s2 := NewSeriesInt32("s2", true, true, data2).
		SetNullMask(data2Mask).
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
	if !checkEqPartitionMap(p2, partMap, nil, "Int32 Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p2)
	}

	// Test 3
	s3 := NewSeriesInt32("test", true, true, data3).
		SetNullMask(data3Mask).
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
	if !checkEqPartitionMap(p3, partMap, nil, "Int32 Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p3)
	}

	// debugPrintPartition(s1.GetPartition(), s1)
	// debugPrintPartition(s2.GetPartition(), s1, s2)
	// debugPrintPartition(s3.GetPartition(), s1, s2, s3)

	partMap = nil
}

func Test_SeriesInt32_Sort(t *testing.T) {
	data := []int32{-195, -27, 33, 679, -67, 920, -352, -674, 250, 767, 697, 873, -802, -123, 308, -558, -518, 169, 313, 593}
	mask := []bool{false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true}

	// Create a new series.
	s := NewSeriesInt32("test", false, true, data)

	// Sort the series.
	sorted := s.Sort()

	// Check the data.
	expected := []int32{-802, -674, -558, -518, -352, -195, -123, -67, -27, 33, 169, 250, 308, 313, 593, 679, 697, 767, 873, 920}
	if !checkEqSliceInt32(sorted.Data().([]int32), expected, nil, "") {
		t.Errorf("SeriesInt32.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]int32))
	}

	// Create a new series.
	s = NewSeriesInt32("test", true, true, data).
		SetNullMask(mask)

	// Sort the series.
	sorted = s.Sort()

	// Check the data.
	expected = []int32{-802, -518, -352, -195, -67, 33, 250, 308, 313, 697, 920, 873, 767, -123, -674, -558, 679, 169, -27, 593}
	if !checkEqSliceInt32(sorted.Data().([]int32), expected, nil, "") {
		t.Errorf("SeriesInt32.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]int32))
	}

	// Check the null mask.
	expectedMask := []bool{false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true}
	if !checkEqSliceBool(sorted.GetNullMask(), expectedMask, nil, "") {
		t.Errorf("SeriesInt32.Sort() failed, expecting %v, got %v", expectedMask, sorted.GetNullMask())
	}
}

func Test_SeriesInt32_Arithmetic_Mul(t *testing.T) {
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
	if !checkEqSlice(i32s.Mul(bools).Data().([]int32), []int32{2}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(boolv).Data().([]int32), []int32{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(bools_).GetNullMask(), []bool{true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}

	// scalar | int32
	if !checkEqSlice(i32s.Mul(i32s).Data().([]int32), []int32{4}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(i32v).Data().([]int32), []int32{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(i32s_).GetNullMask(), []bool{true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Mul(i64s).Data().([]int64), []int64{4}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(i64v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(i64s_).GetNullMask(), []bool{true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Mul(f64s).Data().([]float64), []float64{4}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(f64s_).GetNullMask(), []bool{true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32s.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}

	// vector | bool
	if !checkEqSlice(i32v.Mul(bools).Data().([]int32), []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(boolv).Data().([]int32), []int32{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}

	// vector | int32
	if !checkEqSlice(i32v.Mul(i32s).Data().([]int32), []int32{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(i32v).Data().([]int32), []int32{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}

	// vector | int64
	if !checkEqSlice(i32v.Mul(i64s).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(i64v).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}

	// vector | float64
	if !checkEqSlice(i32v.Mul(f64s).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(f64v).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
	if !checkEqSlice(i32v.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mul") {
		t.Errorf("Error in Int32 Mul")
	}
}

func Test_SeriesInt32_Arithmetic_Div(t *testing.T) {
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
	if !checkEqSlice(i32s.Div(bools).Data().([]float64), []float64{2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(boolv).Data().([]float64), []float64{2, math.Inf(1), 2, math.Inf(1), 2, math.Inf(1), 2, 2, math.Inf(1), math.Inf(1)}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(bools_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// scalar | int32
	if !checkEqSlice(i32s.Div(i32s).Data().([]float64), []float64{1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(i32v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(i32s_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Div(i64s).Data().([]float64), []float64{1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(i64v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(i64s_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Div(f64s).Data().([]float64), []float64{1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(f64v).Data().([]float64), []float64{2, 1, 0.6666666666666666, 0.5, 0.4, 0.3333333333333333, 0.2857142857142857, 0.25, 0.2222222222222222, 0.2}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(f64s_).GetNullMask(), []bool{true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32s.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | bool
	if !checkEqSlice(i32v.Div(bools).Data().([]float64), []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(boolv).Data().([]float64), []float64{1, math.Inf(1), 3, math.Inf(1), 5, math.Inf(1), 7, 8, math.Inf(1), math.Inf(1)}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | int32
	if !checkEqSlice(i32v.Div(i32s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(i32v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | int64
	if !checkEqSlice(i32v.Div(i64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(i64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}

	// vector | float64
	if !checkEqSlice(i32v.Div(f64s).Data().([]float64), []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(f64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
	if !checkEqSlice(i32v.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Div") {
		t.Errorf("Error in Int32 Div")
	}
}

func Test_SeriesInt32_Arithmetic_Mod(t *testing.T) {
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
	if !checkEqSlice(i32s.Mod(bools).Data().([]float64), []float64{0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(bools_).GetNullMask(), []bool{true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}

	// scalar | int32
	if !checkEqSlice(i32s.Mod(i32s).Data().([]float64), []float64{0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(i32v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(i32s_).GetNullMask(), []bool{true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Mod(i64s).Data().([]float64), []float64{0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(i64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(i64s_).GetNullMask(), []bool{true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Mod(f64s).Data().([]float64), []float64{0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(f64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(f64s_).GetNullMask(), []bool{true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32s.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}

	// vector | bool
	if !checkEqSlice(i32v.Mod(bools).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}

	// vector | int32
	if !checkEqSlice(i32v.Mod(i32s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(i32v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}

	// vector | int64
	if !checkEqSlice(i32v.Mod(i64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(i64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}

	// vector | float64
	if !checkEqSlice(i32v.Mod(f64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(f64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
	if !checkEqSlice(i32v.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Mod") {
		t.Errorf("Error in Int32 Mod")
	}
}

func Test_SeriesInt32_Arithmetic_Pow(t *testing.T) {
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
	if !checkEqSlice(i32s.Pow(bools).Data().([]int64), []int64{2}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(boolv).Data().([]int64), []int64{2, 1, 2, 1, 2, 1, 2, 2, 1, 1}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(bools_).GetNullMask(), []bool{true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}

	// scalar | int32
	if !checkEqSlice(i32s.Pow(i32s).Data().([]int64), []int64{4}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(i32v).Data().([]int64), []int64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(i32s_).GetNullMask(), []bool{true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Pow(i64s).Data().([]int64), []int64{4}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(i64v).Data().([]int64), []int64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(i64s_).GetNullMask(), []bool{true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Pow(f64s).Data().([]float64), []float64{4}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(f64v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(f64s_).GetNullMask(), []bool{true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32s.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}

	// vector | bool
	if !checkEqSlice(i32v.Pow(bools).Data().([]int64), []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(boolv).Data().([]int64), []int64{1, 1, 3, 1, 5, 1, 7, 8, 1, 1}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}

	// vector | int32
	if !checkEqSlice(i32v.Pow(i32s).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(i32v).Data().([]int64), []int64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}

	// vector | int64
	if !checkEqSlice(i32v.Pow(i64s).Data().([]int64), []int64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(i64v).Data().([]int64), []int64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}

	// vector | float64
	if !checkEqSlice(i32v.Pow(f64s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(f64v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
	if !checkEqSlice(i32v.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Pow") {
		t.Errorf("Error in Int32 Pow")
	}
}

func Test_SeriesInt32_Arithmetic_Add(t *testing.T) {
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
	if !checkEqSlice(i32s.Add(bools).Data().([]int32), []int32{3}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(boolv).Data().([]int32), []int32{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(bools_).GetNullMask(), []bool{true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// scalar | int32
	if !checkEqSlice(i32s.Add(i32s).Data().([]int32), []int32{4}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(i32v).Data().([]int32), []int32{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(i32s_).GetNullMask(), []bool{true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// scalar | int64
	if !checkEqSlice(i32s.Add(i64s).Data().([]int64), []int64{4}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(i64v).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(i64s_).GetNullMask(), []bool{true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// scalar | float64
	if !checkEqSlice(i32s.Add(f64s).Data().([]float64), []float64{4}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(f64v).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(f64s_).GetNullMask(), []bool{true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// scalar | string
	if !checkEqSlice(i32s.Add(ss).Data().([]string), []string{"22"}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(sv).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(ss_).GetNullMask(), []bool{true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32s.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// vector | bool
	if !checkEqSlice(i32v.Add(bools).Data().([]int32), []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(boolv).Data().([]int32), []int32{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// vector | int32
	if !checkEqSlice(i32v.Add(i32s).Data().([]int32), []int32{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(i32v).Data().([]int32), []int32{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// vector | int64
	if !checkEqSlice(i32v.Add(i64s).Data().([]int64), []int64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(i64v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// vector | float64
	if !checkEqSlice(i32v.Add(f64s).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}

	// vector | string
	if !checkEqSlice(i32v.Add(ss).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(sv).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
	if !checkEqSlice(i32v.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int32 Add") {
		t.Errorf("Error in Int32 Add")
	}
}

func Test_SeriesInt32_Arithmetic_Sub(t *testing.T) {
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
	if !checkEqSlice(i32s.Sub(bools).Data().([]int32), []int32{1}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(boolv).Data().([]int32), []int32{1, 2, 1, 2, 1, 2, 1, 1, 2, 2}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(bools_).GetNullMask(), []bool{true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// scalar | int32
	if !checkEqSlice(i32s.Sub(i32s).Data().([]int32), []int32{0}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32s.Sub(i32v).Data().([]int32), []int32{1, 0, -1, -2, -3, -4, -5, -6, -7, -8}, nil, "Int64 Sub") {
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
	if !checkEqSlice(i32v.Sub(bools).Data().([]int32), []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(boolv).Data().([]int32), []int32{0, 2, 2, 4, 4, 6, 6, 7, 9, 10}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}

	// vector | int32
	if !checkEqSlice(i32v.Sub(i32s).Data().([]int32), []int32{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, nil, "Int64 Sub") {
		t.Errorf("Error in Int64 Sub")
	}
	if !checkEqSlice(i32v.Sub(i32v).Data().([]int32), []int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Int64 Sub") {
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

func Test_SeriesInt32_Logical_Eq(t *testing.T) {
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

	// vector | int32
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

func Test_SeriesInt32_Logical_Ne(t *testing.T) {
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

	// vector | int32
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

func Test_SeriesInt32_Logical_Lt(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt32_Logical_Le(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt32_Logical_Gt(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesInt32_Logical_Ge(t *testing.T) {
	// TODO: add tests for all types
}
