package gandalff

import (
	"math"
	"math/rand"
	"preludiometa"
	"testing"
)

func Test_SeriesBool_Base(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesBool.
	s := NewSeriesBool(data, mask, true, ctx)

	// Check the length.
	if s.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", s.Len())
	}

	// Check the type.
	if s.Type() != preludiometa.BoolType {
		t.Errorf("Expected type of BoolType, got %s", s.Type().ToString())
	}

	// Check the data.
	for i, v := range s.Data().([]bool) {
		if v != data[i] {
			t.Errorf("Expected data of []bool{true, false, true, false, true, false, true, false, true, false}, got %v", s.Data())
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
	for i := range s.Data().([]bool) {
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
	for i := range s.Data().([]bool) {
		s.Set(i, nil)
	}

	// Check the null values.
	for i := range s.Data().([]bool) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get() method.
	for i := range s.Data().([]bool) {
		if s.Get(i) != data[i] {
			t.Errorf("Expected Get(%d) to be %t, got %t", i, data[i], s.Get(i))
		}
	}

	// Check the Set() method.
	for i := range s.Data().([]bool) {
		s.Set(i, !data[i])
	}

	// Check the data.
	for i, v := range s.Data().([]bool) {
		if v != !data[i] {
			t.Errorf("Expected data of []bool{false, true, false, true, false, true, false, true, false, true}, got %v", s.Data())
		}
	}

	// Check the MakeNullable() method.
	p := NewSeriesBool(data, nil, true, ctx)

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
	p = p.MakeNullable().(SeriesBool)

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

func Test_SeriesBool_Append(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}
	dataC := []bool{true, true, true, true, true, true, true, true, true, true}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesBool(dataA, maskA, true, ctx)
	sB := NewSeriesBool(dataB, maskB, true, ctx)
	sC := NewSeriesBool(dataC, maskC, true, ctx)

	// Append the series.
	result := sA.Append(sB).Append(sC)

	// Check the length.
	if result.Len() != 30 {
		t.Errorf("Expected length of 30, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]bool) {
		if i < 10 {
			if v != dataA[i] {
				t.Errorf("Expected %t, got %t at index %d", dataA[i], v, i)
			}
		} else if i < 20 {
			if v != dataB[i-10] {
				t.Errorf("Expected %t, got %t at index %d", dataB[i-10], v, i)
			}
		} else {
			if v != dataC[i-20] {
				t.Errorf("Expected %t, got %t at index %d", dataC[i-20], v, i)
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

	// Append bool, []bool, NullableBool, []NullableBool
	s := NewSeriesBool([]bool{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		if rand.Float32() > 0.5 {
			switch i % 4 {
			case 0:
				s = s.Append(true).(SeriesBool)
			case 1:
				s = s.Append([]bool{true}).(SeriesBool)
			case 2:
				s = s.Append(NullableBool{true, true}).(SeriesBool)
			case 3:
				s = s.Append([]NullableBool{{false, true}}).(SeriesBool)
			}

			if s.Get(i) != true {
				t.Errorf("Expected %t, got %t at index %d (case %d)", true, s.Get(i), i, i%4)
			}
		} else {
			switch i % 4 {
			case 0:
				s = s.Append(false).(SeriesBool)
			case 1:
				s = s.Append([]bool{false}).(SeriesBool)
			case 2:
				s = s.Append(NullableBool{true, false}).(SeriesBool)
			case 3:
				s = s.Append([]NullableBool{{false, false}}).(SeriesBool)
			}

			if s.Get(i) != false {
				t.Errorf("Expected %t, got %t at index %d (case %d)", false, s.Get(i), i, i%4)
			}
		}
	}

	// Append nil
	s = NewSeriesBool([]bool{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(nil).(SeriesBool)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesNA
	s = NewSeriesBool([]bool{}, nil, true, ctx)
	na := NewSeriesNA(10, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(na).(SeriesBool)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], na.GetNullMask(), nil, "SeriesBool.Append") {
			t.Errorf("Expected %v, got %v at index %d", na.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}

	// Append NullableBool
	s = NewSeriesBool([]bool{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(NullableBool{false, true}).(SeriesBool)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append []NullableBool
	s = NewSeriesBool([]bool{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append([]NullableBool{{false, true}}).(SeriesBool)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesBool
	s = NewSeriesBool([]bool{}, nil, true, ctx)
	b := NewSeriesBool([]bool{true, false, true, false, true, false, true, false, true, false}, []bool{true, true, true, true, true, true, true, true, true, true}, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(b).(SeriesBool)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], b.GetNullMask(), nil, "SeriesBool.Append") {
			t.Errorf("Expected %v, got %v at index %d", b.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}
}

func Test_SeriesBool_Cast(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesBool(data, mask, true, ctx)

	// Cast to int.
	castInt := s.Cast(preludiometa.IntType)

	// Check the data.
	for i, v := range castInt.Data().([]int) {
		if data[i] && v != 1 {
			t.Errorf("Expected %d, got %d at index %d", 1, v, i)
		} else if !data[i] && v != 0 {
			t.Errorf("Expected %d, got %d at index %d", 0, v, i)
		}
	}

	// Check the null mask.
	for i, v := range castInt.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to int64.
	castInt64 := s.Cast(preludiometa.Int64Type)

	// Check the data.
	for i, v := range castInt64.Data().([]int64) {
		if data[i] && v != 1 {
			t.Errorf("Expected %d, got %d at index %d", 1, v, i)
		} else if !data[i] && v != 0 {
			t.Errorf("Expected %d, got %d at index %d", 0, v, i)
		}
	}

	// Cast to float64.
	castFloat64 := s.Cast(preludiometa.Float64Type)

	// Check the data.
	for i, v := range castFloat64.Data().([]float64) {
		if data[i] && v != 1.0 {
			t.Errorf("Expected %f, got %f at index %d", 1.0, v, i)
		} else if !data[i] && v != 0.0 {
			t.Errorf("Expected %f, got %f at index %d", 0.0, v, i)
		}
	}

	// Check the null mask.
	for i, v := range castFloat64.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to string.
	castString := s.Cast(preludiometa.StringType)

	// Check the data.
	for i, v := range castString.Data().([]string) {
		if mask[i] && v != NULL_STRING {
			t.Errorf("Expected %s, got %s at index %d", NULL_STRING, v, i)
		} else if !mask[i] && data[i] && v != "true" {
			t.Errorf("Expected %s, got %s at index %d", "true", v, i)
		} else if !mask[i] && !data[i] && v != "false" {
			t.Errorf("Expected %s, got %s at index %d", "false", v, i)
		}

	}

	// Check the null mask.
	for i, v := range castString.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to error.
	castError := s.Cast(preludiometa.ErrorType)

	// Check the message.
	if castError.(SeriesError).msg != "SeriesBool.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_SeriesBool_LogicOperators(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}

	// Create two new series.
	sA := NewSeriesBool(dataA, maskA, true, ctx)
	sB := NewSeriesBool(dataB, maskB, true, ctx)

	// Check the And() method.
	and := sA.And(sB)

	// Check the size.
	if and.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", and.Len())
	}

	// Check the result data.
	for i, v := range and.Data().([]bool) {
		if v != (dataA[i] && dataB[i]) {
			t.Errorf("Expected data of []bool{false, false, false, false, true, false, false, false, false, false}, got %v", and.Data())
		}
	}

	// Check the result null mask.
	for i, v := range and.GetNullMask() {
		if v != (maskA[i] || maskB[i]) {
			t.Errorf("Expected nullMask of []bool{false, false, true, false, true, true, true, false, true, true}, got %v", and.GetNullMask())
		}
	}

	// Check the Or() method.
	// Create two new series.
	sA = NewSeriesBool(dataA, maskA, true, ctx)
	sB = NewSeriesBool(dataB, maskB, true, ctx)

	or := sA.Or(sB)

	// Check the size.
	if or.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", or.Len())
	}

	// Check the result data.
	for i, v := range or.Data().([]bool) {
		if v != (dataA[i] || dataB[i]) {
			t.Errorf("Expected data of []bool{true, true, true, false, true, false, true, true, true, false}, got %v", or.Data())
		}
	}

	// Check the result null mask.
	for i, v := range or.GetNullMask() {
		if v != (maskA[i] || maskB[i]) {
			t.Errorf("Expected nullMask of []bool{false, false, true, false, true, true, true, false, true, true}, got %v", or.GetNullMask())
		}
	}

	// Check the Not() method.
	not := NewSeriesBool(dataA, maskA, true, ctx).
		Not()

	// Check the size.
	if not.Len() != 10 {
		t.Errorf("Expected length of 10, got %d", not.Len())
	}

	// Check the result data.
	for i, v := range not.Data().([]bool) {
		if v != !dataA[i] {
			t.Errorf("Expected data of []bool{false, true, false, true, false, true, false, true, false, true}, got %v", not.Data())
		}
	}

	// Check the result null mask.
	for i, v := range not.GetNullMask() {
		if v != maskA[i] {
			t.Errorf("Expected nullMask of []bool{false, false, true, false, false, true, false, false, true, false}, got %v", not.GetNullMask())
		}
	}
}

func Test_SeriesBool_Filter(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false, false, true, true}
	mask := []bool{false, false, true, false, false, true, false, false, true, false, false, true, true}

	// Create a new series.
	s := NewSeriesBool(data, mask, true, ctx)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12}

	result := []bool{true, true, false, false, true, true, false, false, true}
	resultMask := []bool{false, true, false, true, false, true, false, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool(filterMask, nil, true, ctx))

	// Check the length.
	if filtered.Len() != 9 {
		t.Errorf("Expected length of 7, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]bool) {
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
	if filtered.Len() != 9 {
		t.Errorf("Expected length of 7, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]bool) {
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
	if filtered.Len() != 9 {
		t.Errorf("Expected length of 9, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]bool) {
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

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesBool.Filter: mask length (13) does not match series length (9)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesBool(data, mask, true, ctx)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []bool{true, true, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered = s.Filter(filterMask)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]bool) {
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
	for i, v := range filtered.Data().([]bool) {
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

func Test_SeriesBool_Map(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false, false, true, true}
	mask := []bool{false, false, true, false, false, true, false, false, true, false, false, true, true}

	// Create a new series.
	s := NewSeriesBool(data, mask, true, ctx)

	// MAP TO BOOL

	mappedBool := s.Map(func(v any) any {
		return !v.(bool)
	})

	resultBool := []bool{false, true, false, true, false, true, false, true, false, true, true, false, false}

	// Check the data.
	for i, v := range mappedBool.Data().([]bool) {
		if v != resultBool[i] {
			t.Errorf("Expected %v, got %v at index %d", resultBool[i], v, i)
		}
	}

	// Map the series to int.
	mappedInt := s.Map(func(v any) any {
		if v.(bool) {
			return int(1)
		}
		return int(0)
	})

	resultInt := []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1}

	// Check the data.
	for i, v := range mappedInt.Data().([]int) {
		if v != resultInt[i] {
			t.Errorf("Expected %v, got %v at index %d", resultInt[i], v, i)
		}
	}

	// Map the series to int64.
	mappedInt64 := s.Map(func(v any) any {
		if v.(bool) {
			return int64(1)
		}
		return int64(0)
	})

	resultInt64 := []int64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1}

	// Check the data.
	for i, v := range mappedInt64.Data().([]int64) {
		if v != resultInt64[i] {
			t.Errorf("Expected %v, got %v at index %d", resultInt64[i], v, i)
		}
	}

	// Map the series to float64.
	mappedFloat := s.Map(func(v any) any {
		if v.(bool) {
			return 1.0
		}
		return 0.0
	})

	resultFloat := []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1}

	// Check the data.
	for i, v := range mappedFloat.Data().([]float64) {
		if v != resultFloat[i] {
			t.Errorf("Expected %v, got %v at index %d", resultFloat[i], v, i)
		}
	}

	// Map the series to string.
	mappedString := s.Map(func(v any) any {
		if v.(bool) {
			return "true"
		}
		return "false"
	})

	resultString := []string{"true", "false", "true", "false", "true", "false", "true", "false", "true", "false", "false", "true", "true"}

	// Check the data.
	for i, v := range mappedString.Data().([]string) {
		if v != resultString[i] {
			t.Errorf("Expected %v, got %v at index %d", resultString[i], v, i)
		}
	}
}

func Test_SeriesBool_Group(t *testing.T) {
	var partMap map[int64][]int

	data1 := []bool{true, true, true, true, true, true, true, true, true, true}
	data1Mask := []bool{false, false, false, false, false, true, true, true, true, true}
	data2 := []bool{true, true, false, false, true, true, false, false, true, true}
	data2Mask := []bool{false, true, false, true, false, true, false, true, false, true}
	data3 := []bool{true, false, true, false, true, false, true, false, true, false}
	data3Mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Test 1
	s1 := NewSeriesBool(data1, data1Mask, true, ctx).
		group()

	p1 := s1.GetPartition().getMap()
	if len(p1) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(p1))
	}

	partMap = map[int64][]int{
		0: {0, 1, 2, 3, 4},
		1: {5, 6, 7, 8, 9},
	}
	if !checkEqPartitionMap(p1, partMap, nil, "Bool Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p1)
	}

	// Test 2
	s2 := NewSeriesBool(data2, data2Mask, true, ctx).
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
	if !checkEqPartitionMap(p2, partMap, nil, "Bool Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p2)
	}

	// Test 3
	s3 := NewSeriesBool(data3, data3Mask, true, ctx).
		GroupBy(s2.GetPartition())

	p3 := s3.GetPartition().getMap()
	if len(p3) != 7 {
		t.Errorf("Expected 7 groups, got %d", len(p3))
	}

	partMap = map[int64][]int{
		0: {0, 4},
		1: {1, 3},
		2: {2},
		3: {5},
		4: {6},
		5: {7, 9},
		6: {8},
	}
	if !checkEqPartitionMap(p3, partMap, nil, "Bool Group") {
		t.Errorf("Expected partition map of %v, got %v", partMap, p3)
	}

	// debugPrintPartition(s1.GetPartition(), s1)
	// debugPrintPartition(s2.GetPartition(), s1, s2)
	// debugPrintPartition(s3.GetPartition(), s1, s2, s3)

	partMap = nil
}

func Test_SeriesBool_Sort(t *testing.T) {
	data := []bool{false, false, false, true, true, true, false, true, false, false, true, true, true, false, true, true, true, false, true, false}
	mask := []bool{false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true}

	// Create a new series.
	s := NewSeriesBool(data, nil, true, ctx)

	// Sort the series.
	sorted := s.Sort()

	// Check the data.
	expected := []bool{false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true}
	if !checkEqSliceBool(sorted.Data().([]bool), expected, nil, "") {
		t.Errorf("SeriesBool.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]bool))
	}

	// Create a new series.
	s = NewSeriesBool(data, mask, true, ctx)

	// Sort the series.
	sorted = s.Sort()

	// Check the data.
	expected = []bool{false, false, false, false, true, true, true, true, true, true, true, true, false, false, true, true, true, false, false, false}
	if !checkEqSliceBool(sorted.Data().([]bool), expected, nil, "") {
		t.Errorf("SeriesBool.Sort() failed, expecting %v, got %v", expected, sorted.Data().([]bool))
	}

	// Check the null mask.
	expectedMask := []bool{false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true}
	if !checkEqSliceBool(sorted.GetNullMask(), expectedMask, nil, "") {
		t.Errorf("SeriesBool.Sort() failed, expecting %v, got %v", expectedMask, sorted.GetNullMask())
	}
}

func Test_SeriesBool_Arithmetic_Mul(t *testing.T) {
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
	if !checkEqSlice(bools.Mul(bools).Data().([]int64), []int64{1}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(bools.Mul(boolv).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(bools.Mul(bools_).GetNullMask(), []bool{true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(bools.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}

	// scalar | int
	if !checkEqSlice(i32s.Mul(i32s).Data().([]int), []int{4}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(i32s.Mul(i32v).Data().([]int), []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(i32s.Mul(i32s_).GetNullMask(), []bool{true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(i32s.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}

	// scalar | int64
	if !checkEqSlice(i64s.Mul(i64s).Data().([]int64), []int64{4}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(i64s.Mul(i64v).Data().([]int64), []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(i64s.Mul(i64s_).GetNullMask(), []bool{true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(i64s.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}

	// scalar | float64
	if !checkEqSlice(f64s.Mul(f64s).Data().([]float64), []float64{4}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(f64s.Mul(f64v).Data().([]float64), []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(f64s.Mul(f64s_).GetNullMask(), []bool{true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(f64s.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}

	// vector | bool
	if !checkEqSlice(boolv.Mul(bools).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(boolv).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}

	// vector | int
	if !checkEqSlice(boolv.Mul(i32s).Data().([]int), []int{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(i32v).Data().([]int), []int{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}

	// vector | int64
	if !checkEqSlice(boolv.Mul(i64s).Data().([]int64), []int64{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(i64v).Data().([]int64), []int64{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}

	// vector | float64
	if !checkEqSlice(boolv.Mul(f64s).Data().([]float64), []float64{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(f64v).Data().([]float64), []float64{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
}

func Test_SeriesBool_Arithmetic_Div(t *testing.T) {
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
	if !checkEqSlice(bools.Div(bools).Data().([]float64), []float64{1}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(boolv).Data().([]float64), []float64{1, math.Inf(1), 1, math.Inf(1), 1, math.Inf(1), 1, 1, math.Inf(1), math.Inf(1)}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(bools_).GetNullMask(), []bool{true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}

	// scalar | int
	if !checkEqSlice(bools.Div(i32s).Data().([]float64), []float64{0.5}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(i32v).Data().([]float64), []float64{1, 0.5, 0.3333333333333333, 0.25, 0.2, 0.16666666666666666, 0.14285714285714285, 0.125, 0.1111111111111111, 0.1}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(i32s_).GetNullMask(), []bool{true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}

	// scalar | int64
	if !checkEqSlice(bools.Div(i64s).Data().([]float64), []float64{0.5}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(i64v).Data().([]float64), []float64{1, 0.5, 0.3333333333333333, 0.25, 0.2, 0.16666666666666666, 0.14285714285714285, 0.125, 0.1111111111111111, 0.1}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(i64s_).GetNullMask(), []bool{true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}

	// scalar | float64
	if !checkEqSlice(bools.Div(f64s).Data().([]float64), []float64{0.5}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(f64v).Data().([]float64), []float64{1, 0.5, 0.3333333333333333, 0.25, 0.2, 0.16666666666666666, 0.14285714285714285, 0.125, 0.1111111111111111, 0.1}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(f64s_).GetNullMask(), []bool{true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(bools.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}

	// vector | bool
	if !checkEqSlice(boolv.Div(bools).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(boolv).Data().([]float64), []float64{1, math.NaN(), 1, math.NaN(), 1, math.NaN(), 1, 1, math.NaN(), math.NaN()}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}

	// vector | int
	if !checkEqSlice(boolv.Div(i32s).Data().([]float64), []float64{0.5, 0, 0.5, 0, 0.5, 0, 0.5, 0.5, 0, 0}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(i32v).Data().([]float64), []float64{1, 0, 0.3333333333333333, 0, 0.2, 0, 0.14285714285714285, 0.125, 0, 0}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}

	// vector | int64
	if !checkEqSlice(boolv.Div(i64s).Data().([]float64), []float64{0.5, 0, 0.5, 0, 0.5, 0, 0.5, 0.5, 0, 0}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(i64v).Data().([]float64), []float64{1, 0, 0.3333333333333333, 0, 0.2, 0, 0.14285714285714285, 0.125, 0, 0}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}

	// vector | float64
	if !checkEqSlice(boolv.Div(f64s).Data().([]float64), []float64{0.5, 0, 0.5, 0, 0.5, 0, 0.5, 0.5, 0, 0}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(f64v).Data().([]float64), []float64{1, 0, 0.3333333333333333, 0, 0.2, 0, 0.14285714285714285, 0.125, 0, 0}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
	if !checkEqSlice(boolv.Div(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Div") {
		t.Errorf("Error in Bool Div")
	}
}

func Test_SeriesBool_Arithmetic_Mod(t *testing.T) {
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
	if !checkEqSlice(bools.Mod(bools).Data().([]float64), []float64{0}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(boolv).Data().([]float64), []float64{0, math.NaN(), 0, math.NaN(), 0, math.NaN(), 0, 0, math.NaN(), math.NaN()}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(bools_).GetNullMask(), []bool{true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}

	// scalar | int
	if !checkEqSlice(bools.Mod(i32s).Data().([]float64), []float64{1}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(i32v).Data().([]float64), []float64{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(i32s_).GetNullMask(), []bool{true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}

	// scalar | int64
	if !checkEqSlice(bools.Mod(i64s).Data().([]float64), []float64{1}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(i64v).Data().([]float64), []float64{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(i64s_).GetNullMask(), []bool{true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}

	// scalar | float64
	if !checkEqSlice(bools.Mod(f64s).Data().([]float64), []float64{1}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(f64v).Data().([]float64), []float64{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(f64s_).GetNullMask(), []bool{true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
	if !checkEqSlice(bools.Mod(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Mod") {
		t.Errorf("Error in Bool Mod")
	}
}

func Test_SeriesBool_Arithmetic_Exp(t *testing.T) {
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
	if !checkEqSlice(bools.Exp(bools).Data().([]int64), []int64{1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(boolv).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(bools_).GetNullMask(), []bool{true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}

	// scalar | int
	if !checkEqSlice(bools.Exp(i32s).Data().([]int64), []int64{1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(i32v).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(i32s_).GetNullMask(), []bool{true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}

	// scalar | int64
	if !checkEqSlice(bools.Exp(i64s).Data().([]int64), []int64{1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(i64v).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(i64s_).GetNullMask(), []bool{true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}

	// scalar | float64
	if !checkEqSlice(bools.Exp(f64s).Data().([]float64), []float64{1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(f64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(f64s_).GetNullMask(), []bool{true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(bools.Exp(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}

	// vector | bool
	if !checkEqSlice(boolv.Exp(bools).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(boolv).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}

	// vector | int
	if !checkEqSlice(boolv.Exp(i32s).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(i32v).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}

	// vector | int64
	if !checkEqSlice(boolv.Exp(i64s).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(i64v).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}

	// vector | float64
	if !checkEqSlice(boolv.Exp(f64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(f64v).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
	if !checkEqSlice(boolv.Exp(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Exp") {
		t.Errorf("Error in Bool Exp")
	}
}

func Test_SeriesBool_Arithmetic_Add(t *testing.T) {
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
	if !checkEqSlice(bools.Add(bools).Data().([]int64), []int64{2}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(boolv).Data().([]int64), []int64{2, 1, 2, 1, 2, 1, 2, 2, 1, 1}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(bools_).GetNullMask(), []bool{true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// scalar | int
	if !checkEqSlice(bools.Add(i32s).Data().([]int), []int{3}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(i32v).Data().([]int), []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(i32s_).GetNullMask(), []bool{true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// scalar | int64
	if !checkEqSlice(bools.Add(i64s).Data().([]int64), []int64{3}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(i64v).Data().([]int64), []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(i64s_).GetNullMask(), []bool{true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// scalar | float64
	if !checkEqSlice(bools.Add(f64s).Data().([]float64), []float64{3}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(f64v).Data().([]float64), []float64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(f64s_).GetNullMask(), []bool{true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// scalar | string
	if !checkEqSlice(bools.Add(ss).Data().([]string), []string{"true2"}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(sv).Data().([]string), []string{"true1", "true2", "true3", "true4", "true5", "true6", "true7", "true8", "true9", "true10"}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(ss_).GetNullMask(), []bool{true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// vector | bool
	if !checkEqSlice(boolv.Add(bools).Data().([]int64), []int64{2, 1, 2, 1, 2, 1, 2, 2, 1, 1}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(boolv).Data().([]int64), []int64{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// vector | int
	if !checkEqSlice(boolv.Add(i32s).Data().([]int), []int{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(i32v).Data().([]int), []int{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// vector | int64
	if !checkEqSlice(boolv.Add(i64s).Data().([]int64), []int64{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(i64v).Data().([]int64), []int64{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// vector | float64
	if !checkEqSlice(boolv.Add(f64s).Data().([]float64), []float64{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(f64v).Data().([]float64), []float64{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}

	// vector | string
	if !checkEqSlice(boolv.Add(ss).Data().([]string), []string{"true2", "false2", "true2", "false2", "true2", "false2", "true2", "true2", "false2", "false2"}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(sv).Data().([]string), []string{"true1", "false2", "true3", "false4", "true5", "false6", "true7", "true8", "false9", "false10"}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
}

func Test_SeriesBool_Arithmetic_Sub(t *testing.T) {
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
	if !checkEqSlice(bools.Sub(bools).Data().([]int64), []int64{0}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(boolv).Data().([]int64), []int64{0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(bools_).GetNullMask(), []bool{true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}

	// scalar | int
	if !checkEqSlice(bools.Sub(i32s).Data().([]int), []int{-1}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(i32v).Data().([]int), []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(i32s_).GetNullMask(), []bool{true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}

	// scalar | int64
	if !checkEqSlice(bools.Sub(i64s).Data().([]int64), []int64{-1}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(i64v).Data().([]int64), []int64{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(i64s_).GetNullMask(), []bool{true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}

	// scalar | float64
	if !checkEqSlice(bools.Sub(f64s).Data().([]float64), []float64{-1}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(f64v).Data().([]float64), []float64{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(f64s_).GetNullMask(), []bool{true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}

	// vector | bool
	if !checkEqSlice(boolv.Sub(bools).Data().([]int64), []int64{0, -1, 0, -1, 0, -1, 0, 0, -1, -1}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(boolv).Data().([]int64), []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}

	// vector | int
	if !checkEqSlice(boolv.Sub(i32s).Data().([]int), []int{-1, -2, -1, -2, -1, -2, -1, -1, -2, -2}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(i32v).Data().([]int), []int{0, -2, -2, -4, -4, -6, -6, -7, -9, -10}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}

	// vector | int64
	if !checkEqSlice(boolv.Sub(i64s).Data().([]int64), []int64{-1, -2, -1, -2, -1, -2, -1, -1, -2, -2}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(i64v).Data().([]int64), []int64{0, -2, -2, -4, -4, -6, -6, -7, -9, -10}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}

	// vector | float64
	if !checkEqSlice(boolv.Sub(f64s).Data().([]float64), []float64{-1, -2, -1, -2, -1, -2, -1, -1, -2, -2}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(f64v).Data().([]float64), []float64{0, -2, -2, -4, -4, -6, -6, -7, -9, -10}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
}

func Test_SeriesBool_Boolean_Or(t *testing.T) {
	nas := NewSeriesNA(1, ctx)
	nav := NewSeriesNA(10, ctx)

	bools := NewSeriesBool([]bool{true}, nil, true, ctx)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx)
	bools_ := NewSeriesBool([]bool{true}, nil, true, ctx).SetNullMask([]bool{true}).(SeriesBool)
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true}).(SeriesBool)

	// scalar | bool
	if !checkEqSlice(bools.Or(bools).Data().([]bool), []bool{true}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true}, bools.Or(bools).Data())
	}
	if !checkEqSlice(bools.Or(boolv).Data().([]bool), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, bools.Or(boolv).Data())
	}
	if !checkEqSlice(bools.Or(bools_).GetNullMask(), []bool{true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true}, bools.Or(bools_).GetNullMask())
	}
	if !checkEqSlice(bools.Or(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, bools.Or(boolv_).GetNullMask())
	}

	// scalar | NA
	if !checkEqSlice(bools.Or(nas).Data().([]bool), []bool{true}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true}, bools.Or(nas).Data())
	}
	if !checkEqSlice(bools.Or(nav).Data().([]bool), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, bools.Or(nav).Data())
	}
	if !checkEqSlice(bools.Or(nas).GetNullMask(), []bool{false}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false}, bools.Or(nas).GetNullMask())
	}
	if !checkEqSlice(bools.Or(nav).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, bools.Or(nav).GetNullMask())
	}
	if !checkEqSlice(bools_.Or(nas).GetNullMask(), []bool{true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true}, bools_.Or(nas).GetNullMask())
	}
	if !checkEqSlice(bools_.Or(nav).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, bools_.Or(nav).GetNullMask())
	}

	// vector | bool
	if !checkEqSlice(boolv.Or(bools).Data().([]bool), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, boolv.Or(bools).Data())
	}
	if !checkEqSlice(boolv.Or(boolv).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, boolv.Or(boolv).Data())
	}
	if !checkEqSlice(boolv.Or(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, boolv.Or(bools_).GetNullMask())
	}
	if !checkEqSlice(boolv.Or(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, boolv.Or(boolv_).GetNullMask())
	}

	// vector | NA
	if !checkEqSlice(boolv.Or(nas).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, boolv.Or(nas).Data())
	}
	if !checkEqSlice(boolv.Or(nav).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "Bool Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, boolv.Or(nav).Data())
	}
	if !checkEqSlice(boolv.Or(nas).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, boolv.Or(nas).GetNullMask())
	}
	if !checkEqSlice(boolv.Or(nav).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, boolv.Or(nav).GetNullMask())
	}
	if !checkEqSlice(boolv_.Or(nas).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, boolv_.Or(nas).GetNullMask())
	}
	if !checkEqSlice(boolv_.Or(nav).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, boolv_.Or(nav).GetNullMask())
	}
}
