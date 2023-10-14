package gandalff

import (
	"math/rand"
	"preludiometa"
	"testing"
)

func Test_SeriesBoolMemOpt_Base(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesBoolMemOpt.
	s := newSeriesBoolMemOpt(data, mask, true, ctx)

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

	// Check the Set() with null values.
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
	p := newSeriesBoolMemOpt(data, nil, true, ctx)

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
	p = p.MakeNullable().(SeriesBoolMemOpt)

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

func Test_SeriesBoolMemOpt_Append(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}
	dataC := []bool{true, true, true, true, true, true, true, true, true, true}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := newSeriesBoolMemOpt(dataA, maskA, true, ctx)
	sB := newSeriesBoolMemOpt(dataB, maskB, true, ctx)
	sC := newSeriesBoolMemOpt(dataC, maskC, true, ctx)

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

	// Append random values.
	dataD := []bool{true, false, true, false, true, false, true, false, true, false}
	sD := newSeriesBoolMemOpt(dataD, nil, true, ctx).MakeNullable().(SeriesBoolMemOpt)

	// Check the original data.
	for i, v := range sD.Data().([]bool) {
		if v != dataD[i] {
			t.Errorf("Expected %t, got %t at index %d", dataD[i], v, i)
		}
	}

	for i := 0; i < 100; i++ {
		if rand.Float32() > 0.5 {
			switch i % 4 {
			case 0:
				sD = sD.Append(true).(SeriesBoolMemOpt)
			case 1:
				sD = sD.Append([]bool{true}).(SeriesBoolMemOpt)
			case 2:
				sD = sD.Append(NullableBool{true, true}).(SeriesBoolMemOpt)
			case 3:
				sD = sD.Append([]NullableBool{{false, true}}).(SeriesBoolMemOpt)
			}

			if sD.Get(i+10) != true {
				t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
			}
		} else {
			switch i % 4 {
			case 0:
				sD = sD.Append(false).(SeriesBoolMemOpt)
			case 1:
				sD = sD.Append([]bool{false}).(SeriesBoolMemOpt)
			case 2:
				sD = sD.Append(NullableBool{true, false}).(SeriesBoolMemOpt)
			case 3:
				sD = sD.Append([]NullableBool{{false, false}}).(SeriesBoolMemOpt)
			}

			if sD.Get(i+10) != false {
				t.Errorf("Expected %t, got %t at index %d (case %d)", false, sD.Get(i+10), i+10, i%4)
			}
		}
	}
}

func Test_SeriesBoolMemOpt_Cast(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := newSeriesBoolMemOpt(data, mask, true, ctx)

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
	if castError.(SeriesError).msg != "SeriesBoolMemOpt.Cast: invalid type Error" {
		t.Errorf("Expected error, got %v", castError)
	}
}

func Test_SeriesBoolMemOpt_LogicOperators(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}

	// Create two new series.
	sA := newSeriesBoolMemOpt(dataA, maskA, true, ctx)
	sB := newSeriesBoolMemOpt(dataB, maskB, true, ctx)

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
	sA = newSeriesBoolMemOpt(dataA, maskA, true, ctx)
	sB = newSeriesBoolMemOpt(dataB, maskB, true, ctx)

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
	not := newSeriesBoolMemOpt(dataA, maskA, true, ctx).
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

func Test_SeriesBoolMemOpt_Filter(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false, false, true, true}
	mask := []bool{false, false, true, false, false, true, false, false, true, false, false, true, true}

	// Create a new series.
	s := newSeriesBoolMemOpt(data, mask, true, ctx)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12}

	result := []bool{true, true, false, false, true, true, false, false, true}
	resultMask := []bool{false, true, false, true, false, true, false, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(newSeriesBoolMemOpt(filterMask, nil, true, ctx))

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

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesBoolMemOpt.FilterByMask: mask length (13) does not match series length (9)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = newSeriesBoolMemOpt(data, mask, true, ctx)

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

func Test_SeriesBoolMemOpt_Map(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false, false, true, true}
	mask := []bool{false, false, true, false, false, true, false, false, true, false, false, true, true}

	// Create a new series.
	s := newSeriesBoolMemOpt(data, mask, true, ctx)

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
