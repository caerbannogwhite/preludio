package gandalff

import (
	"testing"
)

func Test_GSeriesBool_Base(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GSeriesBool.
	s := NewGSeriesBool("test", true, true, data)

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
	if s.Type() != BoolType {
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
	for i, _ := range s.Data().([]bool) {
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
	for i, _ := range s.Data().([]bool) {
		s.SetNull(i)
	}

	// Check the null values.
	for i, _ := range s.Data().([]bool) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get() method.
	for i, _ := range s.Data().([]bool) {
		if s.Get(i) != data[i] {
			t.Errorf("Expected Get(%d) to be %t, got %t", i, data[i], s.Get(i))
		}
	}

	// Check the Set() method.
	for i, _ := range s.Data().([]bool) {
		s.Set(i, !data[i])
	}

	// Check the data.
	for i, v := range s.Data().([]bool) {
		if v != !data[i] {
			t.Errorf("Expected data of []bool{false, true, false, true, false, true, false, true, false, true}, got %v", s.Data())
		}
	}
}

func Test_GSeriesBool_LogicOperators(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}

	// Create two new series.
	sA := NewGSeriesBool("testA", true, true, dataA)
	sB := NewGSeriesBool("testB", true, true, dataB)

	// Set the null masks.
	sA.SetNullMask(maskA)
	sB.SetNullMask(maskB)

	// Check the And() method.
	and := sA.And(sB)
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
}
