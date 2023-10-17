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

func Test_SeriesDuration_Arithmetic_Add(t *testing.T) {
	dayNano := int64(24 * time.Hour.Nanoseconds())

	durations := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx)
	durationv := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx)
	durations_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx).SetNullMask([]bool{true})
	durationv_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	times := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	timev := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	times_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).SetNullMask([]bool{true})
	timev_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).
		SetNullMask([]bool{true, false, true, false, true, false, true, false, true, false})

	ss := NewSeriesString([]string{"2"}, nil, true, ctx)
	sv := NewSeriesString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, true, ctx)
	ss_ := NewSeriesString([]string{"2"}, nil, true, ctx).SetNullMask([]bool{true})
	sv_ := NewSeriesString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | duration
	if !checkEqSlice(durations.Add(durations).Data(), []time.Duration{time.Duration(2 * dayNano)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(2 * dayNano)}, durations.Add(durations).Data())
	}
	if !checkEqSlice(durations.Add(durationv).Data(), []time.Duration{time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano), time.Duration(11 * dayNano)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano), time.Duration(11 * dayNano)}, durations.Add(durationv).Data())
	}
	if !checkEqSlice(durations_.Add(durations_).GetNullMask(), []bool{true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, durations_.Add(durations_).GetNullMask())
	}
	if !checkEqSlice(durations_.Add(durationv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durations_.Add(durationv_).GetNullMask())
	}

	// scalar | time
	if !checkEqSlice(durations.Add(times).Data(), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)}, durations.Add(times).Data())
	}
	if !checkEqSlice(durations.Add(timev).Data(), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, durations.Add(timev).Data())
	}
	if !checkEqSlice(durations_.Add(times_).GetNullMask(), []bool{true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, durations_.Add(times_).GetNullMask())
	}
	if !checkEqSlice(durations_.Add(timev_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durations_.Add(timev_).GetNullMask())
	}

	// scalar | string
	if !checkEqSlice(durations.Add(ss).Data(), []string{"24h0m0s2"}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []string{"24h0m0s2"}, durations.Add(ss).Data())
	}
	if !checkEqSlice(durations.Add(sv).Data(), []string{"24h0m0s1", "24h0m0s2", "24h0m0s3", "24h0m0s4", "24h0m0s5", "24h0m0s6", "24h0m0s7", "24h0m0s8", "24h0m0s9", "24h0m0s10"}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []string{"24h0m0s1", "24h0m0s2", "24h0m0s3", "24h0m0s4", "24h0m0s5", "24h0m0s6", "24h0m0s7", "24h0m0s8", "24h0m0s9", "24h0m0s10"}, durations.Add(sv).Data())
	}
	if !checkEqSlice(durations_.Add(ss_).GetNullMask(), []bool{true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true}, durations_.Add(ss_).GetNullMask())
	}
	if !checkEqSlice(durations_.Add(sv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durations_.Add(sv_).GetNullMask())
	}

	// vector | duration
	if !checkEqSlice(durationv.Add(durations).Data(), []time.Duration{time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano), time.Duration(11 * dayNano)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano), time.Duration(11 * dayNano)}, durationv.Add(durations).Data())
	}
	if !checkEqSlice(durationv.Add(durationv).Data(), []time.Duration{time.Duration(2 * dayNano), time.Duration(4 * dayNano), time.Duration(6 * dayNano), time.Duration(8 * dayNano), time.Duration(10 * dayNano), time.Duration(12 * dayNano), time.Duration(14 * dayNano), time.Duration(16 * dayNano), time.Duration(18 * dayNano), time.Duration(20 * dayNano)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(2 * dayNano), time.Duration(4 * dayNano), time.Duration(6 * dayNano), time.Duration(8 * dayNano), time.Duration(10 * dayNano), time.Duration(12 * dayNano), time.Duration(14 * dayNano), time.Duration(16 * dayNano), time.Duration(18 * dayNano), time.Duration(20 * dayNano)}, durationv.Add(durationv).Data())
	}
	if !checkEqSlice(durationv_.Add(durations_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durationv_.Add(durations_).GetNullMask())
	}
	if !checkEqSlice(durationv_.Add(durationv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, durationv_.Add(durationv_).GetNullMask())
	}

	// vector | time
	if !checkEqSlice(durationv.Add(times).Data(), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, durationv.Add(times).Data())
	}
	if !checkEqSlice(durationv.Add(timev).Data(), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 14, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC)}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 14, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC)}, durationv.Add(timev).Data())
	}
	if !checkEqSlice(durationv_.Add(times_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durationv_.Add(times_).GetNullMask())
	}
	if !checkEqSlice(durationv_.Add(timev_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durationv_.Add(timev_).GetNullMask())
	}

	// vector | string
	if !checkEqSlice(durationv.Add(ss).Data(), []string{"24h0m0s2", "48h0m0s2", "72h0m0s2", "96h0m0s2", "120h0m0s2", "144h0m0s2", "168h0m0s2", "192h0m0s2", "216h0m0s2", "240h0m0s2"}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []string{"24h0m0s2", "48h0m0s2", "72h0m0s2", "96h0m0s2", "120h0m0s2", "144h0m0s2", "168h0m0s2", "192h0m0s2", "216h0m0s2", "240h0m0s2"}, durationv.Add(ss).Data())
	}
	if !checkEqSlice(durationv.Add(sv).Data(), []string{"24h0m0s1", "48h0m0s2", "72h0m0s3", "96h0m0s4", "120h0m0s5", "144h0m0s6", "168h0m0s7", "192h0m0s8", "216h0m0s9", "240h0m0s10"}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []string{"24h0m0s1", "48h0m0s2", "72h0m0s3", "96h0m0s4", "120h0m0s5", "144h0m0s6", "168h0m0s7", "192h0m0s8", "216h0m0s9", "240h0m0s10"}, durationv.Add(sv).Data())
	}
	if !checkEqSlice(durationv_.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durationv_.Add(ss_).GetNullMask())
	}
	if !checkEqSlice(durationv_.Add(sv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "SeriesDuration.Add") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, durationv_.Add(sv_).GetNullMask())
	}
}

func Test_SeriesDuration_Arithmetic_Sub(t *testing.T) {
	dayNano := int64(24 * time.Hour.Nanoseconds())

	durations := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx)
	durationv := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx)
	durations_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx).SetNullMask([]bool{true})
	durationv_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | duration
	if !checkEqSlice(durations.Sub(durations).Data(), []time.Duration{time.Duration(0)}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(0)}, durations.Sub(durations).Data())
	}
	if !checkEqSlice(durations.Sub(durationv).Data(), []time.Duration{time.Duration(0), time.Duration(-1 * dayNano), time.Duration(-2 * dayNano), time.Duration(-3 * dayNano), time.Duration(-4 * dayNano), time.Duration(-5 * dayNano), time.Duration(-6 * dayNano), time.Duration(-7 * dayNano), time.Duration(-8 * dayNano), time.Duration(-9 * dayNano)}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(0), time.Duration(-1 * dayNano), time.Duration(-2 * dayNano), time.Duration(-3 * dayNano), time.Duration(-4 * dayNano), time.Duration(-5 * dayNano), time.Duration(-6 * dayNano), time.Duration(-7 * dayNano), time.Duration(-8 * dayNano), time.Duration(-9 * dayNano)}, durations.Sub(durationv).Data())
	}
	if !checkEqSlice(durations_.Sub(durations_).GetNullMask(), []bool{true}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []bool{true}, durations_.Sub(durations_).GetNullMask())
	}
	if !checkEqSlice(durations_.Sub(durationv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durations_.Sub(durationv_).GetNullMask())
	}

	// duration | scalar
	if !checkEqSlice(durationv.Sub(durations).Data(), []time.Duration{time.Duration(0), time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano)}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(0), time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano)}, durationv.Sub(durations).Data())
	}
	if !checkEqSlice(durationv.Sub(durationv).Data(), []time.Duration{time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0)}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0)}, durationv.Sub(durationv).Data())
	}
	if !checkEqSlice(durationv_.Sub(durations_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, durationv_.Sub(durations_).GetNullMask())
	}
	if !checkEqSlice(durationv_.Sub(durationv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, durationv_.Sub(durationv_).GetNullMask())
	}

	// duration | vector
	if !checkEqSlice(durations.Sub(durationv).Data(), []time.Duration{time.Duration(0), time.Duration(-1 * dayNano), time.Duration(-2 * dayNano), time.Duration(-3 * dayNano), time.Duration(-4 * dayNano), time.Duration(-5 * dayNano), time.Duration(-6 * dayNano), time.Duration(-7 * dayNano), time.Duration(-8 * dayNano), time.Duration(-9 * dayNano)}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(0), time.Duration(-1 * dayNano), time.Duration(-2 * dayNano), time.Duration(-3 * dayNano), time.Duration(-4 * dayNano), time.Duration(-5 * dayNano), time.Duration(-6 * dayNano), time.Duration(-7 * dayNano), time.Duration(-8 * dayNano), time.Duration(-9 * dayNano)}, durations.Sub(durationv).Data())
	}
	if !checkEqSlice(durationv.Sub(durationv).Data(), []time.Duration{time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0)}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []time.Duration{time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0), time.Duration(0)}, durationv.Sub(durationv).Data())
	}
	if !checkEqSlice(durationv_.Sub(durationv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, durationv_.Sub(durationv_).GetNullMask())
	}
	if !checkEqSlice(durationv_.Sub(durationv_).GetNullMask(), []bool{false, true, false, true, false, true, false, true, false, true}, nil, "SeriesDuration.Sub") {
		t.Errorf("Expected %v, got %v", []bool{false, true, false, true, false, true, false, true, false, true}, durationv_.Sub(durationv_).GetNullMask())
	}
}
