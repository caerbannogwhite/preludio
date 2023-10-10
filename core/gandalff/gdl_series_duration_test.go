package gandalff

import (
	"testing"
	"time"
)

func Test_SeriesDuration_Append(t *testing.T) {
	dataA := []time.Duration{time.Second, time.Second * 2, time.Second * 3, time.Second * 4, time.Second * 5, time.Second * 6, time.Second * 7, time.Second * 8, time.Second * 9, time.Second * 10}
	dataB := []time.Duration{time.Second * 11, time.Second * 12, time.Second * 13, time.Second * 14, time.Second * 15, time.Second * 16, time.Second * 17, time.Second * 18, time.Second * 19, time.Second * 20}
	dataC := []time.Duration{time.Second * 21, time.Second * 22, time.Second * 23, time.Second * 24, time.Second * 25, time.Second * 26, time.Second * 27, time.Second * 28, time.Second * 29, time.Second * 30}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesDuration(dataA, maskA, true, ctx)
	sB := NewSeriesDuration(dataB, maskB, true, ctx)
	sC := NewSeriesDuration(dataC, maskC, true, ctx)

	// Append the series.
	result := sA.Append(sB).Append(sC)

	// Check the length.
	if result.Len() != 30 {
		t.Errorf("Expected length of 30, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]time.Duration) {
		if i < 10 {
			if v != dataA[i] {
				t.Errorf("Expected %v, got %v at index %d", dataA[i], v, i)
			}
		} else if i < 20 {
			if v != dataB[i-10] {
				t.Errorf("Expected %v, got %v at index %d", dataB[i-10], v, i)
			}
		} else {
			if v != dataC[i-20] {
				t.Errorf("Expected %v, got %v at index %d", dataC[i-20], v, i)
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

	// Append time.Duration, []time.Duration, NullableDuration, []NullableDuration
	s := NewSeriesDuration([]time.Duration{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		switch i % 4 {
		case 0:
			s = s.Append(time.Duration(i)).(SeriesDuration)
		case 1:
			s = s.Append([]time.Duration{time.Duration(i)}).(SeriesDuration)
		case 2:
			s = s.Append(NullableDuration{true, time.Duration(i)}).(SeriesDuration)
		case 3:
			s = s.Append([]NullableDuration{{false, time.Duration(i)}}).(SeriesDuration)
		}

		if s.Get(i) != time.Duration(i) {
			t.Errorf("Expected %v, got %t at index %d (case %d)", time.Duration(i), s.Get(i), i, i%4)
		}
	}

	// Append nil
	s = NewSeriesDuration([]time.Duration{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(nil).(SeriesDuration)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesNA
	s = NewSeriesDuration([]time.Duration{}, nil, true, ctx)
	na := NewSeriesNA(10, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(na).(SeriesDuration)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], na.GetNullMask(), nil, "SeriesDuration.Append") {
			t.Errorf("Expected %v, got %v at index %d", na.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}

	// Append NullableDuration
	s = NewSeriesDuration([]time.Duration{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(NullableDuration{false, time.Second}).(SeriesDuration)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append []NullableDuration
	s = NewSeriesDuration([]time.Duration{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append([]NullableDuration{{false, time.Second}}).(SeriesDuration)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesDuration
	s = NewSeriesDuration([]time.Duration{}, nil, true, ctx)
	b := NewSeriesDuration(dataA, []bool{true, true, true, true, true, true, true, true, true, true}, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(b).(SeriesDuration)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], b.GetNullMask(), nil, "SeriesDuration.Append") {
			t.Errorf("Expected %v, got %v at index %d", b.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}
}
