package gandalff

import (
	"math"
	"math/rand"
	"preludiometa"
	"testing"
)

func Test_SeriesFloat64_Base(t *testing.T) {

	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesFloat64.
	s := NewSeriesFloat64(data, mask, true, ctx)

	// Check the length.
	if s.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", s.Len())
	}

	// Check the type.
	if s.Type() != preludiometa.Float64Type {
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

	// Check the Set() with a null value.
	for i := range s.Data().([]float64) {
		s.Set(i, nil)
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
	p := NewSeriesFloat64(data, nil, true, ctx)

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
	p = p.MakeNullable().(SeriesFloat64)

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

func Test_SeriesFloat64_Append(t *testing.T) {
	dataA := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	dataB := []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	dataC := []float64{21.0, 22.0, 23.0, 24.0, 25.0, 26.0, 27.0, 28.0, 29.0, 30.0}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesFloat64(dataA, maskA, true, ctx)
	sB := NewSeriesFloat64(dataB, maskB, true, ctx)
	sC := NewSeriesFloat64(dataC, maskC, true, ctx)

	// Append the series.
	result := sA.Append(sB).Append(sC)

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

	// Append float64, []float64, NullableFloat64, []NullableFloat64
	s := NewSeriesFloat64([]float64{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		r := rand.Float64()
		switch i % 4 {
		case 0:
			s = s.Append(r).(SeriesFloat64)
		case 1:
			s = s.Append([]float64{r}).(SeriesFloat64)
		case 2:
			s = s.Append(NullableFloat64{true, r}).(SeriesFloat64)
		case 3:
			s = s.Append([]NullableFloat64{{false, r}}).(SeriesFloat64)
		}

		if s.Get(i) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, s.Get(i), i, i%4)
		}
	}

	// Append nil
	s = NewSeriesFloat64([]float64{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(nil).(SeriesFloat64)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesNA
	s = NewSeriesFloat64([]float64{}, nil, true, ctx)
	na := NewSeriesNA(10, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(na).(SeriesFloat64)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], na.GetNullMask(), nil, "SeriesFloat64.Append") {
			t.Errorf("Expected %v, got %v at index %d", na.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}

	// Append NullableFloat64
	s = NewSeriesFloat64([]float64{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(NullableFloat64{false, 1}).(SeriesFloat64)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append []NullableFloat64
	s = NewSeriesFloat64([]float64{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append([]NullableFloat64{{false, 1}}).(SeriesFloat64)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesFloat64
	s = NewSeriesFloat64([]float64{}, nil, true, ctx)
	b := NewSeriesFloat64(dataA, []bool{true, true, true, true, true, true, true, true, true, true}, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(b).(SeriesFloat64)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], b.GetNullMask(), nil, "SeriesFloat64.Append") {
			t.Errorf("Expected %v, got %v at index %d", b.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}
}

func Test_SeriesFloat64_Cast(t *testing.T) {
	data := []float64{0.0, 1.0, 0.0, 3.0, 4.0, 5.0, -6.0, 7.0, 8.0, 9.0}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesFloat64(data, mask, true, ctx)

	// Cast to bool.
	resBool := s.Cast(preludiometa.BoolType)

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

	// Cast to int.
	resInt := s.Cast(preludiometa.IntType)

	// Check the data.
	for i, v := range resInt.Data().([]int) {
		if v != int(data[i]) {
			t.Errorf("Expected %d, got %d at index %d", int(data[i]), v, i)
		}
	}

	// Check the null mask.
	for i, v := range resInt.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to int64.
	resInt64 := s.Cast(preludiometa.Int64Type)

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
	resString := s.Cast(preludiometa.StringType)

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
	castError := s.Cast(preludiometa.ErrorType)

	// Check the message.
	if castError.(SeriesError).msg != "SeriesFloat64.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_SeriesFloat64_Filter(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewSeriesFloat64(data, mask, true, ctx)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []float64{1.0, 3.0, 4.0, 6.0, 7.0, 9.0, 10.0, 11.0, 13.0, 14.0, 16.0, 17.0, 19.0, 20.0}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool(filterMask, nil, true, ctx))

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

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesFloat64.Filter: mask length (20) does not match series length (14)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []float64{2.0, 323, 42, 4.1, 9, 674.0, 42, 48, 9811, 79, 3, 12, 492.3, 47005, -173.4, -28, 323, 42.5, 4, 9.0, 31, 425, 2}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesFloat64(data, mask, true, ctx)

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
	s := NewSeriesFloat64(data, nullMask, true, ctx)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		if v.(float64) >= 7 && v.(float64) <= 100 {
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
		if v.(float64) < 0 {
			return (-int(v.(float64))) % 7
		}
		return int(v.(float64)) % 7
	})

	expectedInt := []int{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 2, 1, 3, 2, 4, 0, 5, 6, 0}
	for i, v := range resInt.Data().([]int) {
		if v != expectedInt[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedInt[i], v, i)
		}
	}

	// Map the series to int64.
	resInt64 := s.Map(func(v any) any {
		if v.(float64) < 0 {
			return (-int64(v.(float64))) % 7
		}
		return int64(v.(float64)) % 7
	})

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
	})

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
	})

	expectedString := []string{"pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "pos", "neg", "pos", "pos", "neg", "pos", "pos", "pos", "neg", "pos"}
	for i, v := range resString.Data().([]string) {
		if v != expectedString[i] {
			t.Errorf("Expected %v, got %v at index %d", expectedString[i], v, i)
		}
	}
}

func Test_SeriesFloat64_Group(t *testing.T) {
	var partMap map[int64][]int

	data1 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	data1Mask := []bool{false, false, false, false, false, true, true, true, true, true}
	data2 := []float64{1, 1, 2, 2, 1, 1, 2, 2, 1, 1}
	data2Mask := []bool{false, true, false, true, false, true, false, true, false, true}
	data3 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data3Mask := []bool{false, false, false, false, false, true, true, true, true, true}

	// Test 1
	s1 := NewSeriesFloat64(data1, data1Mask, true, ctx).
		group()

	p1 := s1.GetPartition().getMap()
	if len(p1) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(p1))
	}

	partMap = map[int64][]int{
		0: {0, 1, 2, 3, 4},
		1: {5, 6, 7, 8, 9},
	}
	if !checkEqPartitionMap(p1, partMap, nil, "Float64 Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p1)
	}

	// Test 2
	s2 := NewSeriesFloat64(data2, data2Mask, true, ctx).
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
	if !checkEqPartitionMap(p2, partMap, nil, "Float64 Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p2)
	}

	// Test 3
	s3 := NewSeriesFloat64(data3, data3Mask, true, ctx).
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
	if !checkEqPartitionMap(p3, partMap, nil, "Float64 Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p3)
	}

	// debugPrintPartition(s1.GetPartition(), s1)
	// debugPrintPartition(s2.GetPartition(), s1, s2)
	// debugPrintPartition(s3.GetPartition(), s1, s2, s3)

	partMap = nil
}

func Test_SeriesFloat64_Sort(t *testing.T) {
	data := []float64{3.8, 5.7, -2.3, -0.2, 6.6, -0.5, -6.4, -2.4, 0.2, -2.8, -7.1, -1.7, -4.2, 1.3, -6.2, -2.8, -4.4, -0.6, 0.0, 9.3}
	mask := []bool{false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true}

	// Create a new series.
	s := NewSeriesFloat64(data, nil, true, ctx)

	// Sort the series.
	sorted := s.Sort()

	// Check the data.
	expected := []float64{-7.1, -6.4, -6.2, -4.4, -4.2, -2.8, -2.8, -2.4, -2.3, -1.7, -0.6, -0.5, -0.2, 0.0, 0.2, 1.3, 3.8, 5.7, 6.6, 9.3}
	if !checkEqSliceFloat64(sorted.Data().([]float64), expected, nil, "") {
		t.Errorf("SeriesFloat64.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]float64))
	}

	// Create a new series.
	s = NewSeriesFloat64(data, mask, true, ctx)

	// Sort the series.
	sorted = s.Sort()

	// Check the data.
	expected = []float64{-7.1, -6.4, -6.2, -4.4, -4.2, -2.3, 0.0, 0.2, 3.8, 6.6, -0.5, -1.7, -2.8, 1.3, -2.4, -2.8, -0.2, -0.6, 5.7, 9.3}
	if !checkEqSliceFloat64(sorted.Data().([]float64), expected, nil, "") {
		t.Errorf("SeriesFloat64.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]float64))
	}

	// Check the null mask.
	expectedMask := []bool{false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true}
	if !checkEqSliceBool(sorted.GetNullMask(), expectedMask, nil, "") {
		t.Errorf("SeriesFloat64.Sort() failed, expecting %v, got %v", expectedMask, sorted.GetNullMask())
	}
}

func Test_SeriesFloat64_Arithmetic_Mul(t *testing.T) {
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

	// scalar | int
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

	// vector | int
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

	// scalar | int
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

	// vector | int
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
	if !checkEqSlice(f64s.Mod(bools).Data().([]float64), []float64{0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(bools_).GetNullMask(), []bool{true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// scalar | int
	if !checkEqSlice(f64s.Mod(i32s).Data().([]float64), []float64{0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i32v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Mod(i64s).Data().([]float64), []float64{0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Mod(f64s).Data().([]float64), []float64{0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(f64v).Data().([]float64), []float64{0, 0, 2, 2, 2, 2, 2, 2, 2, 2}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64s.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// vector | bool
	if !checkEqSlice(f64v.Mod(bools).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// vector | int
	if !checkEqSlice(f64v.Mod(i32s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i32v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// vector | int64
	if !checkEqSlice(f64v.Mod(i64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}

	// vector | float64
	if !checkEqSlice(f64v.Mod(f64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(f64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
	if !checkEqSlice(f64v.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Mod") {
		t.Errorf("Error in Float64 Mod")
	}
}

func Test_SeriesFloat64_Arithmetic_Exp(t *testing.T) {
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
	if !checkEqSlice(f64s.Exp(bools).Data().([]float64), []float64{2}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(boolv).Data().([]float64), []float64{2, 1, 2, 1, 2, 1, 2, 2, 1, 1}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(bools_).GetNullMask(), []bool{true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}

	// scalar | int
	if !checkEqSlice(f64s.Exp(i32s).Data().([]float64), []float64{4}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(i32v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Exp(i64s).Data().([]float64), []float64{4}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(i64v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Exp(f64s).Data().([]float64), []float64{4}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(f64v).Data().([]float64), []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64s.Exp(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}

	// vector | bool
	if !checkEqSlice(f64v.Exp(bools).Data().([]float64), []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(boolv).Data().([]float64), []float64{1, 1, 3, 1, 5, 1, 7, 8, 1, 1}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}

	// vector | int
	if !checkEqSlice(f64v.Exp(i32s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(i32v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}

	// vector | int64
	if !checkEqSlice(f64v.Exp(i64s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(i64v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}

	// vector | float64
	if !checkEqSlice(f64v.Exp(f64s).Data().([]float64), []float64{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(f64v).Data().([]float64), []float64{1, 4, 27, 256, 3125, 46656, 823543, 16777216, 387420489, 10000000000}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
	if !checkEqSlice(f64v.Exp(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Exp") {
		t.Errorf("Error in Float64 Exp")
	}
}

func Test_SeriesFloat64_Arithmetic_Add(t *testing.T) {
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
	if !checkEqSlice(f64s.Add(bools).Data().([]float64), []float64{3}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(boolv).Data().([]float64), []float64{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(bools_).GetNullMask(), []bool{true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// scalar | int
	if !checkEqSlice(f64s.Add(i32s).Data().([]float64), []float64{4}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(i32v).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Add(i64s).Data().([]float64), []float64{4}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(i64v).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Add(f64s).Data().([]float64), []float64{4}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(f64v).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// scalar | string
	if !checkEqSlice(f64s.Add(ss).Data().([]string), []string{"22"}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(sv).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(ss_).GetNullMask(), []bool{true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64s.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// vector | bool
	if !checkEqSlice(f64v.Add(bools).Data().([]float64), []float64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(boolv).Data().([]float64), []float64{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// vector | int
	if !checkEqSlice(f64v.Add(i32s).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(i32v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// vector | int64
	if !checkEqSlice(f64v.Add(i64s).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(i64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// vector | float64
	if !checkEqSlice(f64v.Add(f64s).Data().([]float64), []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}

	// vector | string
	if !checkEqSlice(f64v.Add(ss).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(sv).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
	if !checkEqSlice(f64v.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Add") {
		t.Errorf("Error in Float64 Add")
	}
}

func Test_SeriesFloat64_Arithmetic_Sub(t *testing.T) {
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
	if !checkEqSlice(f64s.Sub(bools).Data().([]float64), []float64{1}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(boolv).Data().([]float64), []float64{1, 2, 1, 2, 1, 2, 1, 1, 2, 2}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(bools_).GetNullMask(), []bool{true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}

	// scalar | int
	if !checkEqSlice(f64s.Sub(i32s).Data().([]float64), []float64{0}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(i32v).Data().([]float64), []float64{1, 0, -1, -2, -3, -4, -5, -6, -7, -8}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(i32s_).GetNullMask(), []bool{true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}

	// scalar | int64
	if !checkEqSlice(f64s.Sub(i64s).Data().([]float64), []float64{0}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(i64v).Data().([]float64), []float64{1, 0, -1, -2, -3, -4, -5, -6, -7, -8}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(i64s_).GetNullMask(), []bool{true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Sub(f64s).Data().([]float64), []float64{0}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(f64v).Data().([]float64), []float64{1, 0, -1, -2, -3, -4, -5, -6, -7, -8}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(f64s_).GetNullMask(), []bool{true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64s.Sub(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}

	// vector | bool
	if !checkEqSlice(f64v.Sub(bools).Data().([]float64), []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(boolv).Data().([]float64), []float64{0, 2, 2, 4, 4, 6, 6, 7, 9, 10}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}

	// vector | int
	if !checkEqSlice(f64v.Sub(i32s).Data().([]float64), []float64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(i32v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}

	// vector | int64
	if !checkEqSlice(f64v.Sub(i64s).Data().([]float64), []float64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(i64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}

	// vector | float64
	if !checkEqSlice(f64v.Sub(f64s).Data().([]float64), []float64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(f64v).Data().([]float64), []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
	if !checkEqSlice(f64v.Sub(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Float64 Sub") {
		t.Errorf("Error in Float64 Sub")
	}
}

func Benchmark_SeriesFloat64_Mul_SerScal_Perf(b *testing.B) {

	N := 1_000_000
	data := make([]float64, N)
	for i := 0; i < N; i++ {
		data[i] = float64(i)
	}

	ser := NewSeriesFloat64(data, nil, true, ctx)
	scal := NewSeriesFloat64([]float64{1.5}, nil, true, ctx)

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

	ser1 := NewSeriesFloat64(data1, nil, true, ctx)
	ser2 := NewSeriesFloat64(data2, nil, true, ctx)

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

	// vector | int
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
