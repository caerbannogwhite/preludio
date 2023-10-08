package gandalff

import (
	"testing"
)

func Test_SeriesNA_Append(t *testing.T) {
	var res Series
	var baseMask, expectedMask []bool
	pool := NewStringPool()

	nas := NewSeriesNA(10, pool)

	baseMask = []bool{true, true, true, true, true, true, true, true, true, true}
	int64s := NewSeriesInt64([]int64{1, 2, 3, 4, 5}, []bool{false, true, false, true, false}, false, pool)
	strings := NewSeriesString([]string{"a", "b", "c", "d", "e"}, []bool{false, true, false, true, false}, false, pool)

	// Append nil
	res = nas.Append(nil)
	expectedMask = append(baseMask, true)
	if res.Len() != 11 {
		t.Errorf("Expected length 11, got %d", res.Len())
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append SeriesNA
	res = nas.Append(NewSeriesNA(5, pool))
	expectedMask = append(baseMask, true, true, true, true, true)
	if res.Len() != 15 {
		t.Errorf("Expected length 15, got %d", res.Len())
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append int64
	res = nas.Append(int64(1))
	expectedMask = append(baseMask, false)
	if res.Len() != 11 {
		t.Errorf("Expected length 11, got %d", res.Len())
	}
	if res.Get(10).(int64) != 1 {
		t.Errorf("Expected last element to be 1, got %v", res.Get(10))
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append NullableInt64
	res = nas.Append(NullableInt64{Value: 1, Valid: true})
	expectedMask = append(baseMask, false)
	if res.Len() != 11 {
		t.Errorf("Expected length 11, got %d", res.Len())
	}
	if res.Get(10).(int64) != 1 {
		t.Errorf("Expected last element to be 1, got %v", res.Get(10))
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append NullableInt64
	res = nas.Append(NullableInt64{Value: 1, Valid: false})
	expectedMask = append(baseMask, true)
	if res.Len() != 11 {
		t.Errorf("Expected length 11, got %d", res.Len())
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append []int64
	res = nas.Append([]int64{1, 2, 3, 4, 5})
	expectedMask = append(baseMask, false, false, false, false, false)
	if res.Len() != 15 {
		t.Errorf("Expected length 15, got %d", res.Len())
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append []NullableInt64
	res = nas.Append([]NullableInt64{
		{Value: 1, Valid: true},
		{Value: 2, Valid: false},
		{Value: 3, Valid: true},
		{Value: 4, Valid: false},
		{Value: 5, Valid: true}})
	expectedMask = append(baseMask, false, true, false, true, false)
	if res.Len() != 15 {
		t.Errorf("Expected length 15, got %d", res.Len())
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append SeriesInt64
	res = nas.Append(int64s)
	expectedMask = append(baseMask, false, true, false, true, false)
	if res.Len() != 15 {
		t.Errorf("Expected length 15, got %d", res.Len())
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append string
	res = nas.Append("a")
	expectedMask = append(baseMask, false)
	if res.Len() != 11 {
		t.Errorf("Expected length 11, got %d", res.Len())
	}
	if res.Get(10).(string) != "a" {
		t.Errorf("Expected last element to be a, got %v", res.Get(10))
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append NullableString
	res = nas.Append(NullableString{Value: "a", Valid: true})
	expectedMask = append(baseMask, false)
	if res.Len() != 11 {
		t.Errorf("Expected length 11, got %d", res.Len())
	}
	if res.Get(10).(string) != "a" {
		t.Errorf("Expected last element to be a, got %v", res.Get(10))
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append NullableString
	res = nas.Append(NullableString{Value: "a", Valid: false})
	expectedMask = append(baseMask, true)
	if res.Len() != 11 {
		t.Errorf("Expected length 11, got %d", res.Len())
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append []string
	res = nas.Append([]string{"a", "b", "c", "d", "e"})
	expectedMask = append(baseMask, false, false, false, false, false)
	if res.Len() != 15 {
		t.Errorf("Expected length 15, got %d", res.Len())
	}
	if !checkEqSliceString(res.Data().([]string), []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "a", "b", "c", "d", "e"}, nil, "Append") {
		t.Errorf("Expecting %v, got %v", []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "a", "b", "c", "d", "e"}, res.Data().([]string))
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append []NullableString
	res = nas.Append([]NullableString{
		{Value: "a", Valid: true},
		{Value: "b", Valid: false},
		{Value: "c", Valid: true},
		{Value: "d", Valid: false},
		{Value: "e", Valid: true}})
	expectedMask = append(baseMask, false, true, false, true, false)
	if res.Len() != 15 {
		t.Errorf("Expected length 15, got %d", res.Len())
	}
	if !checkEqSliceString(res.Data().([]string), []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "a", NULL_STRING, "c", NULL_STRING, "e"}, nil, "Append") {
		t.Errorf("Expecting %v, got %v", []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "a", NULL_STRING, "c", NULL_STRING, "e"}, res.Data().([]string))
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}

	// Append SeriesString
	res = nas.Append(strings)
	expectedMask = append(baseMask, false, true, false, true, false)
	if res.Len() != 15 {
		t.Errorf("Expected length 15, got %d", res.Len())
	}
	if !checkEqSliceString(res.Data().([]string), []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "a", NULL_STRING, "c", NULL_STRING, "e"}, nil, "Append") {
		t.Errorf("Expecting %v, got %v", []string{NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, NULL_STRING, "a", NULL_STRING, "c", NULL_STRING, "e"}, res.Data().([]string))
	}
	if !checkEqSliceBool(res.GetNullMask(), expectedMask, nil, "Append") {
		t.Errorf("Expected null mask to be %v, got %v", expectedMask, res.GetNullMask())
	}
}

func Test_SeriesNA_Arithmetic_Add(t *testing.T) {
	pool := NewStringPool()

	nas := NewSeriesNA(1, pool)
	nav := NewSeriesNA(10, pool)

	ints := NewSeriesInt64([]int64{1}, nil, false, pool)
	intv := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, false, pool)
	ints_ := NewSeriesInt64([]int64{1}, nil, false, pool).SetNullMask([]bool{true})
	intv_ := NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil, false, pool).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	strings := NewSeriesString([]string{"a"}, nil, false, pool)
	stringv := NewSeriesString([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, nil, false, pool)
	strings_ := NewSeriesString([]string{"a"}, nil, false, pool).SetNullMask([]bool{true})
	stringv_ := NewSeriesString([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, nil, false, pool).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | na
	if !checkEqSliceBool(nas.Add(nas).GetNullMask(), []bool{true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true}, nas.Add(nas).GetNullMask())
	}
	if !checkEqSliceBool(nas.Add(nav).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nas.Add(nav).GetNullMask())
	}

	// scalar | int64
	if res, ok := nas.Add(ints).(SeriesNA); !ok || res.Len() != 1 {
		t.Errorf("Expected SeriesNA of length 1, got %v", res)
	}
	if res, ok := nas.Add(intv).(SeriesNA); !ok || res.Len() != 10 {
		t.Errorf("Expected SeriesNA of length 10, got %v", res)
	}
	if !checkEqSliceBool(nas.Add(ints).GetNullMask(), []bool{true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true}, nas.Add(ints).GetNullMask())
	}
	if !checkEqSliceBool(nas.Add(intv).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nas.Add(intv).GetNullMask())
	}
	if !checkEqSliceBool(nas.Add(ints_).GetNullMask(), []bool{true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true}, nas.Add(ints_).GetNullMask())
	}
	if !checkEqSliceBool(nas.Add(intv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nas.Add(intv_).GetNullMask())
	}

	// scalar | string
	if !checkEqSlice(nas.Add(strings).Data(), []string{"NAa"}, nil, "NA Add") {
		t.Errorf("Expected data to be %v, got %v", []string{"NAa"}, nas.Add(strings).Data())
	}
	if !checkEqSlice(nas.Add(stringv).Data(), []string{"NAa", "NAb", "NAc", "NAd", "NAe", "NAf", "NAg", "NAh", "NAi", "NAj"}, nil, "NA Add") {
		t.Errorf("Expected data to be %v, got %v", []string{"NAa", "NAb", "NAc", "NAd", "NAe", "NAf", "NAg", "NAh", "NAi", "NAj"}, nas.Add(stringv).Data())
	}
	if !checkEqSliceBool(nas.Add(strings).GetNullMask(), []bool{false}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false}, nas.Add(strings).GetNullMask())
	}
	if !checkEqSliceBool(nas.Add(stringv).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, nas.Add(stringv).GetNullMask())
	}
	if !checkEqSliceBool(nas.Add(strings_).GetNullMask(), []bool{true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true}, nas.Add(strings_).GetNullMask())
	}
	if !checkEqSliceBool(nas.Add(stringv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, nas.Add(stringv_).GetNullMask())
	}

	// vector | na
	if !checkEqSliceBool(nav.Add(nas).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Add(nas).GetNullMask())
	}
	if !checkEqSliceBool(nav.Add(nav).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Add(nav).GetNullMask())
	}

	// vector | int64
	if res, ok := nav.Add(ints).(SeriesNA); !ok || res.Len() != 10 {
		t.Errorf("Expected SeriesNA of length 10, got %v", res)
	}
	if res, ok := nav.Add(intv).(SeriesNA); !ok || res.Len() != 10 {
		t.Errorf("Expected SeriesNA of length 10, got %v", res)
	}
	if !checkEqSliceBool(nav.Add(ints).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Add(ints).GetNullMask())
	}
	if !checkEqSliceBool(nav.Add(intv).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Add(intv).GetNullMask())
	}
	if !checkEqSliceBool(nav.Add(ints_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Add(ints_).GetNullMask())
	}
	if !checkEqSliceBool(nav.Add(intv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Add(intv_).GetNullMask())
	}

	// vector | string
	if !checkEqSlice(nav.Add(strings).Data(), []string{"NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa"}, nil, "NA Add") {
		t.Errorf("Expected data to be %v, got %v", []string{"NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa", "NAa"}, nav.Add(strings).Data())
	}
	if !checkEqSlice(nav.Add(stringv).Data(), []string{"NAa", "NAb", "NAc", "NAd", "NAe", "NAf", "NAg", "NAh", "NAi", "NAj"}, nil, "NA Add") {
		t.Errorf("Expected data to be %v, got %v", []string{"NAa", "NAb", "NAc", "NAd", "NAe", "NAf", "NAg", "NAh", "NAi", "NAj"}, nav.Add(stringv).Data())
	}
	if !checkEqSliceBool(nav.Add(strings).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, nav.Add(strings).GetNullMask())
	}
	if !checkEqSliceBool(nav.Add(stringv).GetNullMask(), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, nav.Add(stringv).GetNullMask())
	}
	if !checkEqSliceBool(nav.Add(strings_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Add(strings_).GetNullMask())
	}
	if !checkEqSliceBool(nav.Add(stringv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "NA Add") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, nav.Add(stringv_).GetNullMask())
	}
}

func Test_SeriesNA_Boolean_Or(t *testing.T) {
	nas := NewSeriesNA(1, nil)
	nav := NewSeriesNA(10, nil)

	bools := NewSeriesBool([]bool{true}, nil, true, nil)
	boolv := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, nil)
	bools_ := NewSeriesBool([]bool{true}, nil, true, nil).SetNullMask([]bool{true})
	boolv_ := NewSeriesBool([]bool{true, false, true, false, true, false, true, true, false, false}, nil, true, nil).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | na
	if !checkEqSliceBool(nas.Or(nas).GetNullMask(), []bool{true}, nil, "NA Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true}, nas.Or(nas).GetNullMask())
	}
	if !checkEqSliceBool(nas.Or(nav).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nas.Or(nav).GetNullMask())
	}

	// scalar | bool
	if !checkEqSliceBool(nas.Or(bools).Data().([]bool), []bool{true}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true}, nas.Or(bools).Data().([]bool))
	}
	if !checkEqSliceBool(nas.Or(boolv).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, nas.Or(boolv).Data().([]bool))
	}
	if !checkEqSliceBool(nas.Or(bools_).Data().([]bool), []bool{true}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true}, nas.Or(bools_).Data().([]bool))
	}
	if !checkEqSliceBool(nas.Or(boolv_).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, nas.Or(boolv_).Data().([]bool))
	}

	// vector | na
	if !checkEqSliceBool(nav.Or(nas).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Or(nas).GetNullMask())
	}
	if !checkEqSliceBool(nav.Or(nav).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "NA Or") {
		t.Errorf("Expected null mask to be %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, nav.Or(nav).GetNullMask())
	}

	// vector | bool
	if !checkEqSliceBool(nav.Or(bools).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, nav.Or(bools).Data().([]bool))
	}
	if !checkEqSliceBool(nav.Or(boolv).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, nav.Or(boolv).Data().([]bool))
	}
	if !checkEqSliceBool(nav.Or(bools_).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, nav.Or(bools_).Data().([]bool))
	}
	if !checkEqSliceBool(nav.Or(boolv_).Data().([]bool), []bool{true, false, true, false, true, false, true, true, false, false}, nil, "NA Or") {
		t.Errorf("Expected data to be %v, got %v", []bool{true, false, true, false, true, false, true, true, false, false}, nav.Or(boolv_).Data().([]bool))
	}

}
