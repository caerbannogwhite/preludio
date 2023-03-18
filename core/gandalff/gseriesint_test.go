package gandalff

import (
	"testing"
)

func Test_GSeriesInt_Base(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GSeriesInt.
	s := NewGSeriesInt("test", true, true, data)

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
	if s.Type() != IntType {
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
	for i, _ := range s.Data().([]int) {
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
	for i, _ := range s.Data().([]int) {
		s.SetNull(i)
	}

	// Check the null values.
	for i, _ := range s.Data().([]int) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get method.
	for i, _ := range s.Data().([]int) {
		if s.Get(i).(int) != data[i] {
			t.Errorf("Expected Get(%d) to be %d, got %d", i, data[i], s.Get(i).(int))
		}
	}

	// Check the Set method.
	for i, _ := range s.Data().([]int) {
		s.Set(i, i+10)
	}

	// Check the data.
	for i, v := range s.Data().([]int) {
		if v != i+10 {
			t.Errorf("Expected data of []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, got %v", s.Data())
		}
	}
}
