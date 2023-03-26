package gandalff

import (
	"math/rand"
	"testing"
	"typesys"
)

func Test_GDLSeriesBool_Base(t *testing.T) {
	data := []bool{true, false, true, false, true, false, true, false, true, false}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GDLSeriesBool.
	s := NewGDLSeriesBool("test", true, data)

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
	p := NewGDLSeriesBool("test", false, data)

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
	p = p.MakeNullable().(GDLSeriesBool)

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

func Test_GDLSeriesBool_Append(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}
	dataC := []bool{true, true, true, true, true, true, true, true, true, true}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewGDLSeriesBool("testA", true, dataA)
	sB := NewGDLSeriesBool("testB", true, dataB)
	sC := NewGDLSeriesBool("testC", true, dataC)

	// Set the null masks.
	sA.SetNullMask(maskA)
	sB.SetNullMask(maskB)
	sC.SetNullMask(maskC)

	// Append the series.
	result := sA.AppendSeries(sB).AppendSeries(sC)

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
	sD := NewGDLSeriesBool("testD", true, dataD)

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
				sD = sD.Append(true).(GDLSeriesBool)
			case 1:
				sD = sD.Append([]bool{true}).(GDLSeriesBool)
			case 2:
				sD = sD.Append(NullableBool{true, true}).(GDLSeriesBool)
			case 3:
				sD = sD.Append([]NullableBool{{false, true}}).(GDLSeriesBool)
			}

			if sD.Get(i+10) != true {
				t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
			}
		} else {
			switch i % 4 {
			case 0:
				sD = sD.Append(false).(GDLSeriesBool)
			case 1:
				sD = sD.Append([]bool{false}).(GDLSeriesBool)
			case 2:
				sD = sD.Append(NullableBool{true, false}).(GDLSeriesBool)
			case 3:
				sD = sD.Append([]NullableBool{{false, false}}).(GDLSeriesBool)
			}

			if sD.Get(i+10) != false {
				t.Errorf("Expected %t, got %t at index %d (case %d)", false, sD.Get(i+10), i+10, i%4)
			}
		}
	}
}

func Test_GDLSeriesBool_LogicOperators(t *testing.T) {
	dataA := []bool{true, false, true, false, true, false, true, false, true, false}
	dataB := []bool{false, true, false, false, true, false, false, true, false, false}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}

	// Create two new series.
	sA := NewGDLSeriesBool("testA", true, dataA)
	sB := NewGDLSeriesBool("testB", true, dataB)

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
	s := NewGDLSeriesBool("test", true, data)

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
