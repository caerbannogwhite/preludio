package gandalff

import (
	"testing"
	"typesys"
)

func Test_GDLSeriesFloat64_Base(t *testing.T) {

	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GDLSeriesFloat64.
	s := NewGDLSeriesFloat64("test", true, true, data)

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
	if s.Type() != typesys.Float64Type {
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

	// Check the SetNull() method.
	for i := range s.Data().([]float64) {
		s.SetNull(i)
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
}
