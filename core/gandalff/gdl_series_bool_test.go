package gandalff

import (
	"testing"
	"typesys"
)

func Test_GDLSeriesBool_Base(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GDLSeriesBool.
	s := NewGDLSeriesBool("test", true, true, data)

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
}

func Test_GDLSeriesBool_LogicOperators(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}

	// Create two new series.
	sA := NewGDLSeriesBool("testA", true, true, dataA)
	sB := NewGDLSeriesBool("testB", true, true, dataB)

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

	// Check the Or() method.
	or := sA.Or(sB)
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
	not := sA.Not()
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

func Test_GDLSeriesBool_Filter(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new series.
	s := NewGDLSeriesBool("test", true, true, data)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filter := []bool{true, false, true, true, false, true, true, false, true, true}

	result := []bool{true, true, false, false, true, true, false}
	resultMask := []bool{false, true, false, true, false, true, false}

	// Check the Filter() method.
	filtered := s.Filter(filter)

	// Check the data.
	for i, v := range filtered.Data().([]bool) {
		if v != result[i] {
			t.Errorf("Expected data of []bool{true, true, false, false, true, true, false}, got %v", filtered.Data())
		}
	}

	// Check the null mask.
	for i, v := range filtered.GetNullMask() {
		if v != resultMask[i] {
			t.Errorf("Expected nullMask of []bool{false, true, false, true, false, true, false}, got %v", filtered.GetNullMask())
		}
	}
}
