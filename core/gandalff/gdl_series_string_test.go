package gandalff

import (
	"math/rand"
	"preludiometa"
	"testing"
	"time"
)

func Test_SeriesString_Base(t *testing.T) {
	var data, expected []string

	data = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesString.
	s := NewSeriesString(data, mask, true, ctx)

	// Check the length.
	if s.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", s.Len())
	}

	// Check the type.
	if s.Type() != preludiometa.StringType {
		t.Errorf("Expected type of StringType, got %s", s.Type().ToString())
	}

	// Check the data.
	expected = []string{"a", "b", NULL_STRING, "d", "e", NULL_STRING, "g", "h", NULL_STRING, "j"}
	if !checkEqSliceString(s.Data().([]string), expected, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", expected, s.Data())
	}

	// Check the nullability.
	if !s.IsNullable() {
		t.Errorf("Expected IsNullable() to be true, got false")
	}

	// Check the null mask.
	if !checkEqSliceBool(s.GetNullMask(), mask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", mask, s.GetNullMask())
	}

	// Check the null values.
	for i := range s.Data().([]string) {
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

	// Check the Set() with null values.
	for i := range s.Data().([]string) {
		s.Set(i, nil)
	}

	// Check the null values.
	for i := range s.Data().([]string) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get method.
	for i, v := range s.Data().([]string) {
		if s.Get(i) != v {
			t.Errorf("Expected Get(%d) to be %s, got %s", i, v, s.Get(i))
		}
	}

	s = NewSeriesString(data, mask, true, ctx)

	// Check the Set method.
	for i, v := range s.Data().([]string) {
		s.Set(i, v+"x")
	}

	// Check the data.
	expected = []string{"ax", "bx", NULL_STRING + "x", "dx", "ex", NULL_STRING + "x", "gx", "hx", NULL_STRING + "x", "jx"}
	if !checkEqSliceString(s.Data().([]string), expected, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", expected, s.Data())
	}

	// Check the MakeNullable() method.
	p := NewSeriesString(data, nil, true, ctx)

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
	p = p.MakeNullable().(SeriesString)

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

func Test_SeriesString_Append(t *testing.T) {
	var dataA, dataB, dataC, expected []string

	dataA = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	dataB = []string{"k", "l", "m", "n", "o", "p", "q", "r", "s", "t"}
	dataC = []string{"u", "v", "w", "x", "y", "z", "1", "2", "3", "4"}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesString(dataA, maskA, true, ctx)
	sB := NewSeriesString(dataB, maskB, true, ctx)
	sC := NewSeriesString(dataC, maskC, true, ctx)

	// Append the series.
	result := sA.Append(sB).Append(sC)

	// Check the length.
	if result.Len() != 30 {
		t.Errorf("Expected length of 30, got %d", result.Len())
	}

	// Check the data.
	expected = []string{
		"a", "b", NULL_STRING, "d", "e", NULL_STRING, "g", "h", NULL_STRING, "j",
		"k", "l", "m", "n", NULL_STRING, "p", NULL_STRING, "r", "s", NULL_STRING,
		NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING,
	}
	if !checkEqSliceString(result.Data().([]string), expected, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", expected, result.Data())
	}

	// Check the null mask.
	if !checkEqSliceBool(result.GetNullMask(), append(maskA, append(maskB, maskC...)...), nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", append(maskA, append(maskB, maskC...)...), result.GetNullMask())
	}

	// Append random values.
	s := NewSeriesString([]string{}, nil, true, ctx)

	alpha := "abcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < 100; i++ {
		r := string(alpha[rand.Intn(len(alpha))])
		switch i % 4 {
		case 0:
			s = s.Append(r).(SeriesString)
		case 1:
			s = s.Append([]string{r}).(SeriesString)
		case 2:
			s = s.Append(NullableString{true, r}).(SeriesString)
		case 3:
			s = s.Append([]NullableString{{false, r}}).(SeriesString)
		}

		if s.Get(i) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, s.Get(i), i, i%4)
		}
	}

	// Append nil
	s = NewSeriesString([]string{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(nil).(SeriesString)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesNA
	s = NewSeriesString([]string{}, nil, true, ctx)
	na := NewSeriesNA(10, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(na).(SeriesString)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], na.GetNullMask(), nil, "SeriesString.Append") {
			t.Errorf("Expected %v, got %v at index %d", na.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}

	// Append NullableString
	s = NewSeriesString([]string{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(NullableString{false, "a"}).(SeriesString)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append []NullableString
	s = NewSeriesString([]string{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append([]NullableString{{false, "a"}}).(SeriesString)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesString
	s = NewSeriesString([]string{}, nil, true, ctx)
	b := NewSeriesString(dataA, []bool{true, true, true, true, true, true, true, true, true, true}, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(b).(SeriesString)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], b.GetNullMask(), nil, "SeriesString.Append") {
			t.Errorf("Expected %v, got %v at index %d", b.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}
}

func Test_SeriesString_Cast(t *testing.T) {
	data := []string{"true", "false", "0", "3", "4", "5", "hello", "7", "8", "world"}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesString(data, mask, true, ctx)

	// Cast to bool.
	resBool := s.Cast(preludiometa.BoolType)
	expectedMask := []bool{false, false, true, true, true, true, true, true, true, true}

	// Check the data.
	if !checkEqSlice(resBool.Data().([]bool), []bool{true, false, false, false, false, false, false, false, false, false}, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", []bool{true, false, false, false, false, false, false, false, false, false}, resBool.Data())
	}

	// Check the null mask.
	if !checkEqSlice(resBool.GetNullMask(), expectedMask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", mask, resBool.GetNullMask())
	}

	// Cast to int.
	resInt := s.Cast(preludiometa.IntType)
	expectedInt := []int{0, 0, 0, 3, 4, 0, 0, 7, 0, 0}
	expectedMask = []bool{true, true, true, false, false, true, true, false, true, true}

	// Check the data.
	if !checkEqSlice(resInt.Data().([]int), expectedInt, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", expectedInt, resInt.Data())
	}

	// Check the null mask.
	if !checkEqSlice(resInt.GetNullMask(), expectedMask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", expectedMask, resInt.GetNullMask())
	}

	// Cast to int64.
	resInt64 := s.Cast(preludiometa.Int64Type)
	expectedInt64 := []int64{0, 0, 0, 3, 4, 0, 0, 7, 0, 0}

	// Check the data.
	if !checkEqSlice(resInt64.Data().([]int64), expectedInt64, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", expectedInt64, resInt64.Data())
	}

	// Check the null mask.
	if !checkEqSlice(resInt64.GetNullMask(), expectedMask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", expectedMask, resInt64.GetNullMask())
	}

	// Cast to float64.
	resFloat64 := s.Cast(preludiometa.Float64Type)
	expectedFloat64 := []float64{0, 0, 0, 3, 4, 0, 0, 7, 0, 0}

	// Check the data.
	if !checkEqSlice(resFloat64.Data().([]float64), expectedFloat64, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", expectedFloat64, resFloat64.Data())
	}

	// Check the null mask.
	if !checkEqSlice(resFloat64.GetNullMask(), expectedMask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", expectedMask, resFloat64.GetNullMask())
	}

	// Cast to string.
	resString := s.Cast(preludiometa.StringType)
	expectedString := []string{"true", "false", NULL_STRING, "3", "4", NULL_STRING, "hello", "7", NULL_STRING, "world"}

	// Check the data.
	if !checkEqSlice(resString.Data().([]string), expectedString, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", expectedString, resString.Data())
	}

	// Check the null mask.
	if !checkEqSlice(resString.GetNullMask(), mask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", mask, resString.GetNullMask())
	}

	// Cast to time.
	if s.Cast(preludiometa.TimeType).(SeriesError).GetError() != "SeriesString.Cast: cannot cast to Time, use SeriesString.ParseTime(layout) instead" {
		t.Errorf("Expected error, got %v", s.Cast(preludiometa.TimeType))
	}

	// Cast to error.
	castError := s.Cast(preludiometa.ErrorType)

	// Check the message.
	if castError.(SeriesError).msg != "SeriesString.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}

	// Parse time.
	s = NewSeriesString([]string{"2019-01-01", "2019-01-02", "2019-01-03", "2019-01-04"}, nil, true, ctx)
	expectedTime := []time.Time{
		time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
	}

	if !checkEqSlice(s.ParseTime("2006-01-02").Data().([]time.Time), expectedTime, nil, "SeriesString.ParseTime") {
		t.Errorf("Expected data of %v, got %v", expectedTime, s.ParseTime("2006-01-02").Data())
	}
}

func Test_SeriesString_Filter(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t"}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewSeriesString(data, mask, true, ctx)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []string{"a", "c", "d", "f", "g", "i", "j", NULL_STRING, "m", NULL_STRING, "p", NULL_STRING, "s", NULL_STRING}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool(filterMask, nil, true, ctx))

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	if !checkEqSliceString(filtered.Data().([]string), result, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", result, filtered.Data())
	}

	// Check the null mask.
	if !checkEqSliceBool(filtered.GetNullMask(), resultMask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", resultMask, filtered.GetNullMask())
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	if !checkEqSliceString(filtered.Data().([]string), result, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", result, filtered.Data())
	}

	// Check the null mask.
	if !checkEqSliceBool(filtered.GetNullMask(), resultMask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", resultMask, filtered.GetNullMask())
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterIndeces)

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	if !checkEqSliceString(filtered.Data().([]string), result, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", result, filtered.Data())
	}

	// Check the null mask.
	if !checkEqSliceBool(filtered.GetNullMask(), resultMask, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", resultMask, filtered.GetNullMask())
	}

	/////////////////////////////////////////////////////////////////////////////////////

	// try to filter by a series with a different length.
	filtered = filtered.Filter(filterMask)

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesString.Filter: mask length (20) does not match series length (14)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w"}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesString(data, mask, true, ctx)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []string{NULL_STRING, NULL_STRING, NULL_STRING}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	if !checkEqSliceString(filtered.Data().([]string), result, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", result, filtered.Data())
	}

	// Check the null mask.
	if !checkEqSliceBool(filtered.GetNullMask(), []bool{true, true, true}, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", []bool{true, true, true}, filtered.GetNullMask())
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterIndeces)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	if !checkEqSliceString(filtered.Data().([]string), result, nil, "SeriesString.Data") {
		t.Errorf("Expected data of %v, got %v", result, filtered.Data())
	}

	// Check the null mask.
	if !checkEqSliceBool(filtered.GetNullMask(), []bool{true, true, true}, nil, "SeriesString.GetNullMask") {
		t.Errorf("Expected null mask of %v, got %v", []bool{true, true, true}, filtered.GetNullMask())
	}
}

func Test_SeriesString_Map(t *testing.T) {
	data := []string{"", "hello", "world", "this", "is", "a", "test", "of", "the", "map", "function", "in", "the", "series", "", "this", "is", "a", "", "test"}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, true, true, false, false, true}

	// Create a new series.
	s := NewSeriesString(data, nullMask, true, ctx)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		return v.(string) == ""
	})

	expectedBool := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false}
	if !checkEqSliceBool(resBool.Data().([]bool), expectedBool, nil, "SeriesString.Map") {
		t.Errorf("Expected data of %v, got %v", expectedBool, resBool.Data())
	}

	// Map the series to int.
	resInt := s.Map(func(v any) any {
		return int(len(v.(string)))
	})

	expectedInt := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 6, 0, 2, 2, 1, 0, 2}
	if !checkEqSliceInt(resInt.Data().([]int), expectedInt, nil, "SeriesString.Map") {
		t.Errorf("Expected data of %v, got %v", expectedInt, resInt.Data())
	}

	// Map the series to int64.
	resInt64 := s.Map(func(v any) any {
		return int64(len(v.(string)))
	})

	expectedInt64 := []int64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 6, 0, 2, 2, 1, 0, 2}
	if !checkEqSliceInt64(resInt64.Data().([]int64), expectedInt64, nil, "SeriesString.Map") {
		t.Errorf("Expected data of %v, got %v", expectedInt64, resInt64.Data())
	}

	// Map the series to float64.
	resFloat64 := s.Map(func(v any) any {
		return -float64(len(v.(string)))
	})

	expectedFloat64 := []float64{-2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -6, -0, -2, -2, -1, -0, -2}
	if !checkEqSliceFloat64(resFloat64.Data().([]float64), expectedFloat64, nil, "SeriesString.Map") {
		t.Errorf("Expected data of %v, got %v", expectedFloat64, resFloat64.Data())
	}

	// Map the series to string.
	resString := s.Map(func(v any) any {
		if v.(string) == "" {
			return "empty"
		}
		return ""
	})

	expectedString := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "empty", "", "", "", "empty", ""}
	if !checkEqSliceString(resString.Data().([]string), expectedString, nil, "SeriesString.Map") {
		t.Errorf("Expected data of %v, got %v", expectedString, resString.Data())
	}
}

func Test_SeriesString_Group(t *testing.T) {
	var partMap map[int64][]int

	data1 := []string{"foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo"}
	data1Mask := []bool{false, false, false, false, false, true, true, true, true, true}
	data2 := []string{"foo", "foo", "bar", "bar", "foo", "foo", "bar", "bar", "foo", "foo"}
	data2Mask := []bool{false, true, false, true, false, true, false, true, false, true}
	data3 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	data3Mask := []bool{false, false, false, false, false, true, true, true, true, true}

	// Test 1
	s1 := NewSeriesString(data1, data1Mask, true, ctx).
		group()

	p1 := s1.GetPartition().getMap()
	if len(p1) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(p1))
	}

	partMap = map[int64][]int{
		0: {0, 1, 2, 3, 4},
		1: {5, 6, 7, 8, 9},
	}
	if !checkEqPartitionMap(p1, partMap, nil, "String Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p1)
	}

	// Test 2
	s2 := NewSeriesString(data2, data2Mask, true, ctx).
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
	if !checkEqPartitionMap(p2, partMap, nil, "String Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p2)
	}

	// Test 3
	s3 := NewSeriesString(data3, data3Mask, true, ctx).
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
	if !checkEqPartitionMap(p3, partMap, nil, "String Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p3)
	}

	// debugPrintPartition(s1.GetPartition(), s1)
	// debugPrintPartition(s2.GetPartition(), s1, s2)
	// debugPrintPartition(s3.GetPartition(), s1, s2, s3)

	partMap = nil
}

func Test_SeriesString_Sort(t *testing.T) {
	data := []string{"w", "w", "d", "y", "b", "e", "a", "e", "e", "b", "l", "u", "a", "g", "w", "u", "{", "x", "t", "h"}
	mask := []bool{false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true}

	// Create a new series.
	s := NewSeriesString(data, nil, true, ctx)

	// Sort the series.
	sorted := s.Sort()

	// Check the data.
	expected := []string{"a", "a", "b", "b", "d", "e", "e", "e", "g", "h", "l", "t", "u", "u", "w", "w", "w", "x", "y", "{"}
	if !checkEqSliceString(sorted.Data().([]string), expected, nil, "") {
		t.Errorf("SeriesString.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]string))
	}

	// Create a new series.
	s = NewSeriesString(data, mask, true, ctx)

	// Sort the series.
	sorted = s.Sort()

	// Check the data.
	expected = []string{"a", "a", "b", "d", "e", "l", "t", "w", "w", "{", NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING}
	if !checkEqSliceString(sorted.Data().([]string), expected, nil, "") {
		t.Errorf("SeriesString.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]string))
	}

	// Check the null mask.
	expectedMask := []bool{false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true}
	if !checkEqSliceBool(sorted.GetNullMask(), expectedMask, nil, "") {
		t.Errorf("SeriesString.Sort() failed, expecting %v, got %v", expectedMask, sorted.GetNullMask())
	}
}

func Test_SeriesString_Arithmetic_Add(t *testing.T) {
	nas := NewSeriesNA(1, ctx)
	nav := NewSeriesNA(10, ctx)

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

	// scalar | NA
	if !checkEqSlice(ss.Add(nas).Data().([]string), []string{"2" + NULL_STRING}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"2" + NULL_STRING}, ss.Add(nas).Data().([]string))
	}
	if !checkEqSlice(ss.Add(nav).Data().([]string), []string{"2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING, "2" + NULL_STRING}, ss.Add(nav).Data().([]string))
	}
	if !checkEqSlice(ss.Add(nas).GetNullMask(), []bool{false}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false}, ss.Add(nas).GetNullMask())
	}
	if !checkEqSlice(ss.Add(nav).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, ss.Add(nav).GetNullMask())
	}
	if !checkEqSlice(ss_.Add(nas).GetNullMask(), []bool{true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, ss_.Add(nas).GetNullMask())
	}
	if !checkEqSlice(ss_.Add(nav).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, ss_.Add(nav).GetNullMask())
	}

	// scalar | bool
	if !checkEqSlice(ss.Add(bools).Data().([]string), []string{"2true"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"2true"}, ss.Add(bools).Data().([]string))
	}
	if !checkEqSlice(ss.Add(boolv).Data().([]string), []string{"2true", "2false", "2true", "2false", "2true", "2false", "2true", "2true", "2false", "2false"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"2true", "2false", "2true", "2false", "2true", "2false", "2true", "2true", "2false", "2false"}, ss.Add(boolv).Data().([]string))
	}
	if !checkEqSlice(ss.Add(bools_).GetNullMask(), []bool{true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, ss.Add(bools_).GetNullMask())
	}
	if !checkEqSlice(ss.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, ss.Add(boolv_).GetNullMask())
	}

	// scalar | int
	if !checkEqSlice(ss.Add(i32s).Data().([]string), []string{"22"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"22"}, ss.Add(i32s).Data().([]string))
	}
	if !checkEqSlice(ss.Add(i32v).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, ss.Add(i32v).Data().([]string))
	}
	if !checkEqSlice(ss.Add(i32s_).GetNullMask(), []bool{true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, ss.Add(i32s_).GetNullMask())
	}
	if !checkEqSlice(ss.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, ss.Add(i32v_).GetNullMask())
	}

	// scalar | int64
	if !checkEqSlice(ss.Add(i64s).Data().([]string), []string{"22"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"22"}, ss.Add(i64s).Data().([]string))
	}
	if !checkEqSlice(ss.Add(i64v).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, ss.Add(i64v).Data().([]string))
	}
	if !checkEqSlice(ss.Add(i64s_).GetNullMask(), []bool{true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, ss.Add(i64s_).GetNullMask())
	}
	if !checkEqSlice(ss.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, ss.Add(i64v_).GetNullMask())
	}

	// scalar | float64
	if !checkEqSlice(ss.Add(f64s).Data().([]string), []string{"22"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"22"}, ss.Add(f64s).Data().([]string))
	}
	if !checkEqSlice(ss.Add(f64v).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, ss.Add(f64v).Data().([]string))
	}
	if !checkEqSlice(ss.Add(f64s_).GetNullMask(), []bool{true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, ss.Add(f64s_).GetNullMask())
	}
	if !checkEqSlice(ss.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, ss.Add(f64v_).GetNullMask())
	}

	// scalar | string
	if !checkEqSlice(ss.Add(ss).Data().([]string), []string{"22"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"22"}, ss.Add(ss).Data().([]string))
	}
	if !checkEqSlice(ss.Add(sv).Data().([]string), []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "210"}, ss.Add(sv).Data().([]string))
	}
	if !checkEqSlice(ss.Add(ss_).GetNullMask(), []bool{true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, ss.Add(ss_).GetNullMask())
	}
	if !checkEqSlice(ss.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, ss.Add(sv_).GetNullMask())
	}

	// vector | NA
	if !checkEqSlice(sv.Add(nas).Data().([]string), []string{"1" + NULL_STRING, "2" + NULL_STRING, "3" + NULL_STRING, "4" + NULL_STRING, "5" + NULL_STRING, "6" + NULL_STRING, "7" + NULL_STRING, "8" + NULL_STRING, "9" + NULL_STRING, "10" + NULL_STRING}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"1" + NULL_STRING, "2" + NULL_STRING, "3" + NULL_STRING, "4" + NULL_STRING, "5" + NULL_STRING, "6" + NULL_STRING, "7" + NULL_STRING, "8" + NULL_STRING, "9" + NULL_STRING, "10" + NULL_STRING}, sv.Add(nas).Data().([]string))
	}
	if !checkEqSlice(sv.Add(nav).Data().([]string), []string{"1" + NULL_STRING, "2" + NULL_STRING, "3" + NULL_STRING, "4" + NULL_STRING, "5" + NULL_STRING, "6" + NULL_STRING, "7" + NULL_STRING, "8" + NULL_STRING, "9" + NULL_STRING, "10" + NULL_STRING}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"1" + NULL_STRING, "2" + NULL_STRING, "3" + NULL_STRING, "4" + NULL_STRING, "5" + NULL_STRING, "6" + NULL_STRING, "7" + NULL_STRING, "8" + NULL_STRING, "9" + NULL_STRING, "10" + NULL_STRING}, sv.Add(nav).Data().([]string))
	}
	if !checkEqSlice(sv.Add(nas).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, sv.Add(nas).GetNullMask())
	}
	if !checkEqSlice(sv.Add(nav).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, sv.Add(nav).GetNullMask())
	}
	if !checkEqSlice(sv_.Add(nas).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, sv_.Add(nas).GetNullMask())
	}
	if !checkEqSlice(sv_.Add(nav).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, sv_.Add(nav).GetNullMask())
	}

	// vector | bool
	if !checkEqSlice(sv.Add(bools).Data().([]string), []string{"1true", "2true", "3true", "4true", "5true", "6true", "7true", "8true", "9true", "10true"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"1true", "2true", "3true", "4true", "5true", "6true", "7true", "8true", "9true", "10true"}, sv.Add(bools).Data().([]string))
	}
	if !checkEqSlice(sv.Add(boolv).Data().([]string), []string{"1true", "2false", "3true", "4false", "5true", "6false", "7true", "8true", "9false", "10false"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"1true", "2false", "3true", "4false", "5true", "6false", "7true", "8true", "9false", "10false"}, sv.Add(boolv).Data().([]string))
	}
	if !checkEqSlice(sv.Add(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, sv.Add(bools_).GetNullMask())
	}
	if !checkEqSlice(sv.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, sv.Add(boolv_).GetNullMask())
	}

	// vector | int
	if !checkEqSlice(sv.Add(i32s).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, sv.Add(i32s).Data().([]string))
	}
	if !checkEqSlice(sv.Add(i32v).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, sv.Add(i32v).Data().([]string))
	}
	if !checkEqSlice(sv.Add(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, sv.Add(i32s_).GetNullMask())
	}
	if !checkEqSlice(sv.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, sv.Add(i32v_).GetNullMask())
	}

	// vector | int64
	if !checkEqSlice(sv.Add(i64s).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, sv.Add(i64s).Data().([]string))
	}
	if !checkEqSlice(sv.Add(i64v).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, sv.Add(i64v).Data().([]string))
	}
	if !checkEqSlice(sv.Add(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, sv.Add(i64s_).GetNullMask())
	}
	if !checkEqSlice(sv.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, sv.Add(i64v_).GetNullMask())
	}

	// vector | float64
	if !checkEqSlice(sv.Add(f64s).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, sv.Add(f64s).Data().([]string))
	}
	if !checkEqSlice(sv.Add(f64v).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, sv.Add(f64v).Data().([]string))
	}
	if !checkEqSlice(sv.Add(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, sv.Add(f64s_).GetNullMask())
	}
	if !checkEqSlice(sv.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, sv.Add(f64v_).GetNullMask())
	}

	// vector | string
	if !checkEqSlice(sv.Add(ss).Data().([]string), []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"12", "22", "32", "42", "52", "62", "72", "82", "92", "102"}, sv.Add(ss).Data().([]string))
	}
	if !checkEqSlice(sv.Add(sv).Data().([]string), []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "1010"}, sv.Add(sv).Data().([]string))
	}
	if !checkEqSlice(sv.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, sv.Add(ss_).GetNullMask())
	}
	if !checkEqSlice(sv.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "String Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, sv.Add(sv_).GetNullMask())
	}
}

func Test_SeriesString_Logical_Eq(t *testing.T) {
	ss := NewSeriesString([]string{"1"}, nil, true, ctx)
	ss_ := NewSeriesString([]string{"1"}, nil, true, ctx).SetNullMask([]bool{true})
	sv := NewSeriesString([]string{"1", "2", "3"}, nil, true, ctx)
	sv_ := NewSeriesString([]string{"1", "2", "3"}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	// scalar | scalar
	res := ss.Eq(ss)
	if res.Data().([]bool)[0] != true {
		t.Errorf("Expected %v, got %v", true, res.Data().([]bool)[0])
	}

	res = ss.Eq(ss_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", true, res.IsNull(0))
	}

	// scalar | vector
	res = ss.Eq(sv)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = ss.Eq(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	// vector | scalar
	res = sv.Eq(ss)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{true, false, false}, res.Data().([]bool))
	}

	res = sv_.Eq(ss)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	// vector | vector
	res = sv.Eq(sv)
	if res.Data().([]bool)[0] != true || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{true, true, true}, res.Data().([]bool))
	}

	res = sv.Eq(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	res = sv_.Eq(sv)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	res = sv_.Eq(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{true, true, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}
}

func Test_SeriesString_Logical_Ne(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesString_Logical_Lt(t *testing.T) {
	ss := NewSeriesString([]string{"1"}, nil, true, ctx)
	ss_ := NewSeriesString([]string{"1"}, nil, true, ctx).SetNullMask([]bool{true})
	sv := NewSeriesString([]string{"1", "2", "3"}, nil, true, ctx)
	sv_ := NewSeriesString([]string{"1", "2", "3"}, nil, true, ctx).SetNullMask([]bool{true, true, false})

	// scalar | scalar
	res := ss.Lt(ss)
	if res.Data().([]bool)[0] != false {
		t.Errorf("Expected %v, got %v", false, res.Data().([]bool)[0])
	}

	res = ss.Lt(ss_)
	if res.IsNull(0) == false {
		t.Errorf("Expected %v, got %v", true, res.IsNull(0))
	}

	// scalar | vector
	res = ss.Lt(sv)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != true || res.Data().([]bool)[2] != true {
		t.Errorf("Expected %v, got %v", []bool{false, true, true}, res.Data().([]bool))
	}

	res = ss.Lt(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	// vector | scalar
	res = sv.Lt(ss)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = sv_.Lt(ss)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	// vector | vector
	res = sv.Lt(sv)
	if res.Data().([]bool)[0] != false || res.Data().([]bool)[1] != false || res.Data().([]bool)[2] != false {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, res.Data().([]bool))
	}

	res = sv.Lt(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	res = sv_.Lt(sv)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}

	res = sv_.Lt(sv_)
	if res.IsNull(0) == false || res.IsNull(1) == false || res.IsNull(2) == true {
		t.Errorf("Expected %v, got %v", []bool{false, false, false}, []bool{res.IsNull(0), res.IsNull(1), res.IsNull(2)})
	}
}

func Test_SeriesString_Logical_Le(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesString_Logical_Gt(t *testing.T) {
	// TODO: add tests for all types
}

func Test_SeriesString_Logical_Ge(t *testing.T) {
	// TODO: add tests for all types
}
