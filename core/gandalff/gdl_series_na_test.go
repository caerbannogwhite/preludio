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
