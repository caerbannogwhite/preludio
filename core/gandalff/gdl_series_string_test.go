package gandalff

import (
	"math/rand"
	"testing"
	"typesys"
)

var stringPool *StringPool

func init() {
	stringPool = NewStringPool()
}

func Test_GDLSeriesString_Base(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	mask := []bool{false, false, true, false, false, true, false, false, true, false}

	// Create a new GDLSeriesString.
	s := NewGDLSeriesString("test", true, data, stringPool)

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
	if s.Type() != typesys.StringType {
		t.Errorf("Expected type of StringType, got %s", s.Type().ToString())
	}

	// Check the data.
	for i, v := range s.Data().([]string) {
		if v != data[i] {
			t.Errorf("Expected data of []string{\"a\", \"b\", \"c\", \"d\", \"e\", \"f\", \"g\", \"h\", \"i\", \"j\"}, got %v", s.Data())
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
	for i := range s.Data().([]string) {
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

	// Check the SetNull() method.
	for i := range s.Data().([]string) {
		s.SetNull(i)
	}

	// Check the null values.
	for i := range s.Data().([]string) {
		if !s.IsNull(i) {
			t.Errorf("Expected IsNull(%d) to be true, got false", i)
		}
	}

	// Check the null count.
	if s.NullCount() != 10 {
		t.Errorf("Expected NullCount() to be 10, got %d", s.NullCount())
	}

	// Check the Get method.
	for i, v := range s.Data().([]string) {
		if s.Get(i) != v {
			t.Errorf("Expected Get(%d) to be %s, got %s", i, v, s.Get(i))
		}
	}

	// Check the Set method.
	for i, v := range s.Data().([]string) {
		s.Set(i, v+"x")
	}

	// Check the data.
	for i, v := range s.Data().([]string) {
		if v != data[i]+"x" {
			t.Errorf("Expected data of []string{\"ax\", \"bx\", \"cx\", \"dx\", \"ex\", \"fx\", \"gx\", \"hx\", \"ix\", \"jx\"}, got %v", s.Data())
		}
	}

	// Check the MakeNullable() method.
	p := NewGDLSeriesString("test", false, data, stringPool)

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
	p = p.MakeNullable().(GDLSeriesString)

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

func Test_GDLSeriesString_Append(t *testing.T) {
	dataA := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	dataB := []string{"k", "l", "m", "n", "o", "p", "q", "r", "s", "t"}
	dataC := []string{"u", "v", "w", "x", "y", "z", "1", "2", "3", "4"}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewGDLSeriesString("testA", true, dataA, stringPool)
	sB := NewGDLSeriesString("testB", true, dataB, stringPool)
	sC := NewGDLSeriesString("testC", true, dataC, stringPool)

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
	for i, v := range result.Data().([]string) {
		if i < 10 {
			if v != dataA[i] {
				t.Errorf("Expected %s, got %s at index %d", dataA[i], v, i)
			}
		} else if i < 20 {
			if v != dataB[i-10] {
				t.Errorf("Expected %s, got %s at index %d", dataB[i-10], v, i)
			}
		} else {
			if v != dataC[i-20] {
				t.Errorf("Expected %s, got %s at index %d", dataC[i-20], v, i)
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
	dataD := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	sD := NewGDLSeriesString("testD", true, dataD, stringPool)

	// Check the original data.
	for i, v := range sD.Data().([]string) {
		if v != dataD[i] {
			t.Errorf("Expected %s, got %s at index %d", dataD[i], v, i)
		}
	}

	alpha := "abcdefghijklmnopqrstuvwxyz0123456789"

	for i := 0; i < 100; i++ {
		r := string(alpha[rand.Intn(len(alpha))])
		switch i % 4 {
		case 0:
			sD = sD.Append(r).(GDLSeriesString)
		case 1:
			sD = sD.Append([]string{r}).(GDLSeriesString)
		case 2:
			sD = sD.Append(NullableString{true, r}).(GDLSeriesString)
		case 3:
			sD = sD.Append([]NullableString{{false, r}}).(GDLSeriesString)
		}

		if sD.Get(i+10) != r {
			t.Errorf("Expected %t, got %t at index %d (case %d)", true, sD.Get(i+10), i+10, i%4)
		}
	}
}

func Test_GDLSeriesString_Filter(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t"}
	mask := []bool{false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, true}

	// Create a new series.
	s := NewGDLSeriesString("test", true, data, stringPool)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask := []bool{true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, true}
	filterIndeces := []int{0, 2, 3, 5, 6, 8, 9, 10, 12, 13, 15, 16, 18, 19}

	result := []string{"a", "c", "d", "f", "g", "i", "j", "k", "m", "n", "p", "q", "s", "t"}
	resultMask := []bool{false, false, false, false, false, false, false, true, false, true, false, true, false, true}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the Filter() method.
	filtered := s.Filter(NewGDLSeriesBool("mask", false, filterMask).(GDLSeriesBool))

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]string) {
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
	// 							Check the FilterByMask() method.
	filtered = s.FilterByMask(filterMask)

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]string) {
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
	// 							Check the FilterByIndeces() method.
	filtered = s.FilterByIndeces(filterIndeces)

	// Check the length.
	if filtered.Len() != len(result) {
		t.Errorf("Expected length of %d, got %d", len(result), filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]string) {
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
	filtered = filtered.FilterByMask(filterMask)

	if e, ok := filtered.(GDLSeriesError); !ok || e.Error() != "GDLSeriesString.FilterByMask: mask length (20) does not match series length (14)" {
		t.Errorf("Expected GDLSeriesError, got %v", filtered)
	}

	// Another test.
	data = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w"}
	mask = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	// Create a new series.
	s = NewGDLSeriesString("test", true, data, stringPool)

	// Set the null mask.
	s.SetNullMask(mask)

	// Filter mask.
	filterMask = []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true}
	filterIndeces = []int{0, 15, 22}

	result = []string{"a", "p", "w"}

	/////////////////////////////////////////////////////////////////////////////////////
	// 							Check the FilterByMask() method.
	filtered = s.FilterByMask(filterMask)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]string) {
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
	// 							Check the FilterByIndeces() method.
	filtered = s.FilterByIndeces(filterIndeces)

	// Check the length.
	if filtered.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", filtered.Len())
	}

	// Check the data.
	for i, v := range filtered.Data().([]string) {
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
