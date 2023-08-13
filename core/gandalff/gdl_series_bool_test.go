package gandalff

import (
	"math"
	"math/rand"
	"testing"
	"typesys"
)

func Test_SeriesBool_Base(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new SeriesBool.
	s := NewSeriesBool("test", true, true, data)

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
	if s.Type() != typesys.BoolType {
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

	// Check the SetNull() method.
	for i := range s.Data().([]bool) {
		s.SetNull(i)
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
	p := NewSeriesBool("test", false, true, data)

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
	p = p.MakeNullable().(SeriesBool)

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

func Test_SeriesBool_Append(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}
	dataC := []bool{true, true, true, true, true, true, true, true, true, true}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesBool("testA", true, true, dataA)
	sB := NewSeriesBool("testB", true, true, dataB)
	sC := NewSeriesBool("testC", true, true, dataC)

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
	sD := NewSeriesBool("testD", true, true, dataD)

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
				sD = sD.Append(true).(SeriesBool)
			case 1:
				sD = sD.Append([]bool{true}).(SeriesBool)
			case 2:
				sD = sD.Append(NullableBool{true, true}).(SeriesBool)
			case 3:
				sD = sD.Append([]NullableBool{{false, true}}).(SeriesBool)
			}

			if sD.Get(i+10) != true {
				t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
			}
		} else {
			switch i % 4 {
			case 0:
				sD = sD.Append(false).(SeriesBool)
			case 1:
				sD = sD.Append([]bool{false}).(SeriesBool)
			case 2:
				sD = sD.Append(NullableBool{true, false}).(SeriesBool)
			case 3:
				sD = sD.Append([]NullableBool{{false, false}}).(SeriesBool)
			}

			if sD.Get(i+10) != false {
				t.Errorf("Expected %t, got %t at index %d (case %d)", false, sD.Get(i+10), i+10, i%4)
			}
		}
	}
}

func Test_SeriesBool_Cast(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewSeriesBool("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Cast to int32.
	castInt32 := s.Cast(typesys.Int32Type, nil)

	// Check the data.
	for i, v := range castInt32.Data().([]int32) {
		if data[i] && v != 1 {
			t.Errorf("Expected %d, got %d at index %d", 1, v, i)
		} else if !data[i] && v != 0 {
			t.Errorf("Expected %d, got %d at index %d", 0, v, i)
		}
	}

	// Check the null mask.
	for i, v := range castInt32.GetNullMask() {
		if v != mask[i] {
			t.Errorf("Expected nullMask of %t, got %t at index %d", mask[i], v, i)
		}
	}

	// Cast to int64.
	castInt64 := s.Cast(typesys.Int64Type, nil)

	// Check the data.
	for i, v := range castInt64.Data().([]int64) {
		if data[i] && v != 1 {
			t.Errorf("Expected %d, got %d at index %d", 1, v, i)
		} else if !data[i] && v != 0 {
			t.Errorf("Expected %d, got %d at index %d", 0, v, i)
		}
	}

	// Cast to float64.
	castFloat64 := s.Cast(typesys.Float64Type, nil)

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
	castString := s.Cast(typesys.StringType, NewStringPool())

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
	castError := s.Cast(typesys.ErrorType, nil)

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
	sA := NewSeriesBool("testA", true, true, dataA)
	sB := NewSeriesBool("testB", true, true, dataB)

	// Set the null masks.
	sA.SetNullMask(maskA)
	sB.SetNullMask(maskB)

	sbA := sA.(SeriesBool)
	sbB := sB.(SeriesBool)

	// Check the And() method.
	and := sbA.And(sbB)

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
	sA = NewSeriesBool("testA", true, true, dataA)
	sB = NewSeriesBool("testB", true, true, dataB)

	// Set the null masks.
	sA.SetNullMask(maskA)
	sB.SetNullMask(maskB)

	sbA = sA.(SeriesBool)
	sbB = sB.(SeriesBool)
	or := sbA.Or(sbB)

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
	not := NewSeriesBool("test", true, true, dataA).
		SetNullMask(maskA).(SeriesBool).
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
	s := NewSeriesBool("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12}

	result := []bool{true, true, false, false, true, true, false, false, true}
	resultMask := []bool{false, true, false, true, false, true, false, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewSeriesBool("mask", false, true, filterMask).(SeriesBool))

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

	if e, ok := filtered.(SeriesError); !ok || e.GetError() != "SeriesBool.FilterByMask: mask length (13) does not match series length (9)" {
		t.Errorf("Expected SeriesError, got %v", filtered)
	}

	// Another test.
	data = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewSeriesBool("test", true, true, data)

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
	s := NewSeriesBool("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// MAP TO BOOL

	mappedBool := s.Map(func(v any) any {
		return !v.(bool)
	}, nil)

	resultBool := []bool{false, true, false, true, false, true, false, true, false, true, true, false, false}

	// Check the data.
	for i, v := range mappedBool.Data().([]bool) {
		if v != resultBool[i] {
			t.Errorf("Expected %v, got %v at index %d", resultBool[i], v, i)
		}
	}

	// Map the series to int32.
	mappedInt := s.Map(func(v any) any {
		if v.(bool) {
			return int32(1)
		}
		return int32(0)
	}, nil)

	resultInt := []int32{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1}

	// Check the data.
	for i, v := range mappedInt.Data().([]int32) {
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
	}, nil)

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
	}, nil)

	resultFloat := []float64{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1}

	// Check the data.
	for i, v := range mappedFloat.Data().([]float64) {
		if v != resultFloat[i] {
			t.Errorf("Expected %v, got %v at index %d", resultFloat[i], v, i)
		}
	}

	// Map the series to string.
	pool := NewStringPool()
	mappedString := s.Map(func(v any) any {
		if v.(bool) {
			return "true"
		}
		return "false"
	}, pool)

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
	s1 := NewSeriesBool("s1", true, true, data1).
		SetNullMask(data1Mask).
		Group()

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
	s2 := NewSeriesBool("s2", true, true, data2).
		SetNullMask(data2Mask).
		SubGroup(s1.GetPartition())

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
	s3 := NewSeriesBool("test", true, true, data3).
		SetNullMask(data3Mask).
		SubGroup(s2.GetPartition())

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
}

func Test_SeriesBool_Arithmetic_Mul(t *testing.T) {
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

	// scalar | int32
	if !checkEqSlice(i32s.Mul(i32s).Data().([]int32), []int32{4}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(i32s.Mul(i32v).Data().([]int32), []int32{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, nil, "Bool Mul") {
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

	// vector | int32
	if !checkEqSlice(boolv.Mul(i32s).Data().([]int32), []int32{2, 0, 2, 0, 2, 0, 2, 2, 0, 0}, nil, "Bool Mul") {
		t.Errorf("Error in Bool Mul")
	}
	if !checkEqSlice(boolv.Mul(i32v).Data().([]int32), []int32{1, 0, 3, 0, 5, 0, 7, 8, 0, 0}, nil, "Bool Mul") {
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

	// scalar | int32
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

	// vector | int32
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

	// scalar | int32
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

func Test_SeriesBool_Arithmetic_Pow(t *testing.T) {
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
	if !checkEqSlice(bools.Pow(bools).Data().([]int64), []int64{1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(boolv).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(bools_).GetNullMask(), []bool{true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}

	// scalar | int32
	if !checkEqSlice(bools.Pow(i32s).Data().([]int64), []int64{1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(i32v).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(i32s_).GetNullMask(), []bool{true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}

	// scalar | int64
	if !checkEqSlice(bools.Pow(i64s).Data().([]int64), []int64{1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(i64v).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(i64s_).GetNullMask(), []bool{true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}

	// scalar | float64
	if !checkEqSlice(bools.Pow(f64s).Data().([]float64), []float64{1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(f64v).Data().([]float64), []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(f64s_).GetNullMask(), []bool{true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(bools.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}

	// vector | bool
	if !checkEqSlice(boolv.Pow(bools).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(boolv).Data().([]int64), []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(bools_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(boolv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}

	// vector | int32
	if !checkEqSlice(boolv.Pow(i32s).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(i32v).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(i32s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(i32v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}

	// vector | int64
	if !checkEqSlice(boolv.Pow(i64s).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(i64v).Data().([]int64), []int64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(i64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(i64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}

	// vector | float64
	if !checkEqSlice(boolv.Pow(f64s).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(f64v).Data().([]float64), []float64{1, 0, 1, 0, 1, 0, 1, 1, 0, 0}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(f64s_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
	if !checkEqSlice(boolv.Pow(f64v_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "Bool Pow") {
		t.Errorf("Error in Bool Pow")
	}
}

func Test_SeriesBool_Arithmetic_Add(t *testing.T) {
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

	// scalar | int32
	if !checkEqSlice(bools.Add(i32s).Data().([]int32), []int32{3}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(bools.Add(i32v).Data().([]int32), []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, nil, "Bool Add") {
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

	// vector | int32
	if !checkEqSlice(boolv.Add(i32s).Data().([]int32), []int32{3, 2, 3, 2, 3, 2, 3, 3, 2, 2}, nil, "Bool Add") {
		t.Errorf("Error in Bool Add")
	}
	if !checkEqSlice(boolv.Add(i32v).Data().([]int32), []int32{2, 2, 4, 4, 6, 6, 8, 9, 9, 10}, nil, "Bool Add") {
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

	// scalar | int32
	if !checkEqSlice(bools.Sub(i32s).Data().([]int32), []int32{-1}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(bools.Sub(i32v).Data().([]int32), []int32{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, nil, "Bool Sub") {
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

	// vector | int32
	if !checkEqSlice(boolv.Sub(i32s).Data().([]int32), []int32{-1, -2, -1, -2, -1, -2, -1, -1, -2, -2}, nil, "Bool Sub") {
		t.Errorf("Error in Bool Sub")
	}
	if !checkEqSlice(boolv.Sub(i32v).Data().([]int32), []int32{0, -2, -2, -4, -4, -6, -6, -7, -9, -10}, nil, "Bool Sub") {
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
