package gandalff

import (
	"preludiometa"
	"testing"
	"time"
)

func Test_SeriesTime_Append(t *testing.T) {
	dataA := []time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}
	dataB := []time.Time{time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 13, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 14, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 17, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 19, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC)}
	dataC := []time.Time{time.Date(2020, 1, 21, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 23, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 24, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 25, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 27, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 28, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 30, 0, 0, 0, 0, time.UTC)}

	maskA := []bool{false, false, true, false, false, true, false, false, true, false}
	maskB := []bool{false, false, false, false, true, false, true, false, false, true}
	maskC := []bool{true, true, true, true, true, true, true, true, true, true}

	// Create two new series.
	sA := NewSeriesTime(dataA, maskA, true, ctx)
	sB := NewSeriesTime(dataB, maskB, true, ctx)
	sC := NewSeriesTime(dataC, maskC, true, ctx)

	// Append the series.
	result := sA.Append(sB).Append(sC)

	// Check the length.
	if result.Len() != 30 {
		t.Errorf("Expected length of 30, got %d", result.Len())
	}

	// Check the data.
	for i, v := range result.Data().([]time.Time) {
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

	// Append time.Time, []time.Time, NullableTime, []NullableTime
	s := NewSeriesTime([]time.Time{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		tm := time.Date(2020, 1, i, 0, 0, 0, 0, time.UTC)
		switch i % 4 {
		case 0:
			s = s.Append(tm).(SeriesTime)
		case 1:
			s = s.Append([]time.Time{tm}).(SeriesTime)
		case 2:
			s = s.Append(NullableTime{true, tm}).(SeriesTime)
		case 3:
			s = s.Append([]NullableTime{{false, tm}}).(SeriesTime)
		}

		if s.Get(i) != tm {
			t.Errorf("Expected %v, got %t at index %d (case %d)", tm, s.Get(i), i, i%4)
		}
	}

	// Append nil
	s = NewSeriesTime([]time.Time{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(nil).(SeriesTime)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesNA
	s = NewSeriesTime([]time.Time{}, nil, true, ctx)
	na := NewSeriesNA(10, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(na).(SeriesTime)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], na.GetNullMask(), nil, "SeriesTime.Append") {
			t.Errorf("Expected %v, got %v at index %d", na.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}

	// Append NullableTime
	s = NewSeriesTime([]time.Time{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(NullableTime{false, time.Now()}).(SeriesTime)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append []NullableTime
	s = NewSeriesTime([]time.Time{}, nil, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append([]NullableTime{{false, time.Now()}}).(SeriesTime)
		if !s.IsNull(i) {
			t.Errorf("Expected %t, got %t at index %d", true, s.IsNull(i), i)
		}
	}

	// Append SeriesTime
	s = NewSeriesTime([]time.Time{}, nil, true, ctx)
	b := NewSeriesTime(dataA, []bool{true, true, true, true, true, true, true, true, true, true}, true, ctx)

	for i := 0; i < 100; i++ {
		s = s.Append(b).(SeriesTime)
		if !checkEqSlice(s.GetNullMask()[s.Len()-10:], b.GetNullMask(), nil, "SeriesTime.Append") {
			t.Errorf("Expected %v, got %v at index %d", b.GetNullMask(), s.GetNullMask()[s.Len()-10:], i)
		}
	}
}

func Test_SeriesTime_Cast(t *testing.T) {
	data := []time.Time{
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC),
	}

	// Create a new series.
	s := NewSeriesTime(data, nil, true, ctx)

	// Cast to bool.
	if s.Cast(preludiometa.BoolType).GetError() != "SeriesTime.Cast: cannot cast to Bool" {
		t.Errorf("Expected an error, got %v", s.Cast(preludiometa.BoolType))
	}

	// Cast to int.
	resInt := s.Cast(preludiometa.IntType).(SeriesInt)
	expectedInt := []int{1577836800000000000, 1577923200000000000, 1578009600000000000, 1578096000000000000, 1578182400000000000, 1578268800000000000, 1578355200000000000, 1578441600000000000, 1578528000000000000, 1578614400000000000}

	if !checkEqSliceInt(resInt.Data().([]int), expectedInt, nil, "") {
		t.Errorf("SeriesTime.Cast: expected %v, got %v", expectedInt, resInt.Data())
	}

	// Cast to int64.
	resInt64 := s.Cast(preludiometa.Int64Type).(SeriesInt64)
	expectedInt64 := []int64{1577836800000000000, 1577923200000000000, 1578009600000000000, 1578096000000000000, 1578182400000000000, 1578268800000000000, 1578355200000000000, 1578441600000000000, 1578528000000000000, 1578614400000000000}

	if !checkEqSliceInt64(resInt64.Data().([]int64), expectedInt64, nil, "") {
		t.Errorf("SeriesTime.Cast: expected %v, got %v", expectedInt64, resInt64.Data())
	}

	// Cast to float64.
	resFloat64 := s.Cast(preludiometa.Float64Type).(SeriesFloat64)
	expectedFloat64 := []float64{1577836800000000000, 1577923200000000000, 1578009600000000000, 1578096000000000000, 1578182400000000000, 1578268800000000000, 1578355200000000000, 1578441600000000000, 1578528000000000000, 1578614400000000000}

	if !checkEqSliceFloat64(resFloat64.Data().([]float64), expectedFloat64, nil, "") {
		t.Errorf("SeriesTime.Cast: expected %v, got %v", expectedFloat64, resFloat64.Data())
	}

	// Cast to string.
	resString := s.Cast(preludiometa.StringType).(SeriesString)

	expectedString := []string{"2020-01-01 00:00:00", "2020-01-02 00:00:00", "2020-01-03 00:00:00", "2020-01-04 00:00:00", "2020-01-05 00:00:00", "2020-01-06 00:00:00", "2020-01-07 00:00:00", "2020-01-08 00:00:00", "2020-01-09 00:00:00", "2020-01-10 00:00:00"}
	if !checkEqSliceString(resString.Data().([]string), expectedString, nil, "") {
		t.Errorf("SeriesTime.Cast: expected %v, got %v", expectedString, resString.Data())
	}

	resString = s.SetTimeFormat("2006-01-02").Cast(preludiometa.StringType).(SeriesString)

	expectedString = []string{"2020-01-01", "2020-01-02", "2020-01-03", "2020-01-04", "2020-01-05", "2020-01-06", "2020-01-07", "2020-01-08", "2020-01-09", "2020-01-10"}
	if !checkEqSliceString(resString.Data().([]string), expectedString, nil, "") {
		t.Errorf("SeriesTime.Cast: expected %v, got %v", expectedString, resString.Data())
	}
}

func Test_SeriesTime_Map(t *testing.T) {

	baseTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	data := []time.Time{
		baseTime, baseTime.AddDate(0, 0, 1), baseTime.AddDate(0, 0, 2), baseTime.AddDate(0, 0, 3), baseTime.AddDate(0, 0, 4), baseTime.AddDate(0, 0, 5), baseTime.AddDate(0, 0, 6), baseTime.AddDate(0, 0, 7), baseTime.AddDate(0, 0, 8), baseTime.AddDate(0, 0, 9),
		baseTime.AddDate(0, 0, 10), baseTime.AddDate(0, 0, 11), baseTime.AddDate(0, 0, 12), baseTime.AddDate(0, 0, 13), baseTime.AddDate(0, 0, 14), baseTime.AddDate(0, 0, 15), baseTime.AddDate(0, 0, 16), baseTime.AddDate(0, 0, 17), baseTime.AddDate(0, 0, 18), baseTime.AddDate(0, 0, 19),
	}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true, false}

	// Create a new series.
	s := NewSeriesTime(data, nullMask, true, ctx)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		return v.(time.Time).Day()%2 == 0
	})

	expectedBool := []bool{false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true}
	if !checkEqSliceBool(resBool.Data().([]bool), expectedBool, nil, "") {
		t.Errorf("SeriesTime.Map: expected %v, got %v", expectedBool, resBool.Data())
	}

	// Map the series to int.
	resInt := s.Map(func(v any) any {
		return int(v.(time.Time).Day())
	})

	expectedInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	if !checkEqSliceInt(resInt.Data().([]int), expectedInt, nil, "") {
		t.Errorf("SeriesTime.Map: expected %v, got %v", expectedInt, resInt.Data())
	}

	// Map the series to int64.
	resInt64 := s.Map(func(v any) any {
		return int64(v.(time.Time).Day())
	})

	expectedInt64 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	if !checkEqSliceInt64(resInt64.Data().([]int64), expectedInt64, nil, "") {
		t.Errorf("SeriesTime.Map: expected %v, got %v", expectedInt64, resInt64.Data())
	}

	// Map the series to float64.
	resFloat64 := s.Map(func(v any) any {
		return float64(v.(time.Time).Day())
	})

	expectedFloat64 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10.0, 11.0, 12.0, 13.0, 14, 15.0, 16.0, 17.0, 18, 19.0, 20.0}
	if !checkEqSliceFloat64(resFloat64.Data().([]float64), expectedFloat64, nil, "") {
		t.Errorf("SeriesTime.Map: expected %v, got %v", expectedFloat64, resFloat64.Data())
	}

	// Map the series to string.
	resString := s.Map(func(v any) any {
		return v.(time.Time).Format("2006-01-02")
	})

	expectedString := []string{"2020-01-01", "2020-01-02", "2020-01-03", "2020-01-04", "2020-01-05", "2020-01-06", "2020-01-07", "2020-01-08", "2020-01-09", "2020-01-10", "2020-01-11", "2020-01-12", "2020-01-13", "2020-01-14", "2020-01-15", "2020-01-16", "2020-01-17", "2020-01-18", "2020-01-19", "2020-01-20"}
	if !checkEqSliceString(resString.Data().([]string), expectedString, nil, "") {
		t.Errorf("SeriesTime.Map: expected %v, got %v", expectedString, resString.Data())
	}
}

func Test_SeriesTime_Arithmetic_Add(t *testing.T) {
	dayNano := int64(24 * time.Hour.Nanoseconds())

	times := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	timev := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	times_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).SetNullMask([]bool{true})
	timev_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).
		SetNullMask([]bool{true, false, true, false, true, false, true, false, true, false})

	durations := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx)
	durationv := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx)
	durations_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx).SetNullMask([]bool{true})
	durationv_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	ss := NewSeriesString([]string{"2"}, nil, true, ctx)
	sv := NewSeriesString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, true, ctx)
	ss_ := NewSeriesString([]string{"2"}, nil, true, ctx).SetNullMask([]bool{true})
	sv_ := NewSeriesString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | time
	if !checkEqSlice(times.Add(times).Data().([]time.Time), []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC)}, times.Add(times).Data())
	}
	if !checkEqSlice(times.Add(timev).Data().([]time.Time), []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 3, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 4, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 5, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 6, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 7, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 8, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 9, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 10, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 11, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 3, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 4, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 5, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 6, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 7, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 8, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 9, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 10, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 11, 0, 0, 0, 0, time.UTC)}, times.Add(timev).Data())
	}
	if !checkEqSlice(times_.Add(times_).GetNullMask(), []bool{true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true}, times_.Add(times_).GetNullMask())
	}
	if !checkEqSlice(times_.Add(timev_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, times_.Add(timev_).GetNullMask())
	}

	// scalar | duration
	if !checkEqSlice(times.Add(durations).Data().([]time.Time), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)}, times.Add(durations).Data())
	}
	if !checkEqSlice(times.Add(durationv).Data().([]time.Time), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, times.Add(durationv).Data())
	}
	if !checkEqSlice(times_.Add(durations_).GetNullMask(), []bool{true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true}, times_.Add(durations_).GetNullMask())
	}
	if !checkEqSlice(times_.Add(durationv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, times_.Add(durationv_).GetNullMask())
	}

	// scalar | string
	if !checkEqSlice(times.Add(ss).Data().([]string), []string{"2020-01-01 00:00:00 +0000 UTC2"}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []string{"2020-01-01 00:00:00 +0000 UTC2"}, times.Add(ss).Data())
	}
	if !checkEqSlice(times.Add(sv).Data().([]string), []string{"2020-01-01 00:00:00 +0000 UTC1", "2020-01-01 00:00:00 +0000 UTC2", "2020-01-01 00:00:00 +0000 UTC3", "2020-01-01 00:00:00 +0000 UTC4", "2020-01-01 00:00:00 +0000 UTC5", "2020-01-01 00:00:00 +0000 UTC6", "2020-01-01 00:00:00 +0000 UTC7", "2020-01-01 00:00:00 +0000 UTC8", "2020-01-01 00:00:00 +0000 UTC9", "2020-01-01 00:00:00 +0000 UTC10"}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []string{"2020-01-01 00:00:00 +0000 UTC1", "2020-01-01 00:00:00 +0000 UTC2", "2020-01-01 00:00:00 +0000 UTC3", "2020-01-01 00:00:00 +0000 UTC4", "2020-01-01 00:00:00 +0000 UTC5", "2020-01-01 00:00:00 +0000 UTC6", "2020-01-01 00:00:00 +0000 UTC7", "2020-01-01 00:00:00 +0000 UTC8", "2020-01-01 00:00:00 +0000 UTC9", "2020-01-01 00:00:00 +0000 UTC10"}, times.Add(sv).Data())
	}
	if !checkEqSlice(times_.Add(ss_).GetNullMask(), []bool{true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true}, times_.Add(ss_).GetNullMask())
	}
	if !checkEqSlice(times_.Add(sv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, times_.Add(sv_).GetNullMask())
	}

	// vector | time
	if !checkEqSlice(timev.Add(times).Data().([]time.Time), []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 3, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 4, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 5, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 6, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 7, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 8, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 9, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 10, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 11, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 3, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 4, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 5, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 6, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 7, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 8, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 9, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 10, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 11, 0, 0, 0, 0, time.UTC)}, timev.Add(times).Data())
	}
	if !checkEqSlice(timev.Add(timev).Data().([]time.Time), []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 4, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 6, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 8, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 10, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 12, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 14, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 16, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 18, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 20, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(4040, 2, 2, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 4, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 6, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 8, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 10, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 12, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 14, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 16, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 18, 0, 0, 0, 0, time.UTC), time.Date(4040, 2, 20, 0, 0, 0, 0, time.UTC)}, timev.Add(timev).Data())
	}
	if !checkEqSlice(timev_.Add(times_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Add(times_).GetNullMask())
	}
	if !checkEqSlice(timev_.Add(timev_).GetNullMask(), []bool{true, false, true, false, true, false, true, false, true, false}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, false, true, false, true, false, true, false, true, false}, timev_.Add(timev_).GetNullMask())
	}

	// vector | duration
	if !checkEqSlice(timev.Add(durations).Data().([]time.Time), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)}, timev.Add(durations).Data())
	}
	if !checkEqSlice(timev.Add(durationv).Data().([]time.Time), []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 14, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []time.Time{time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 14, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC)}, timev.Add(durationv).Data())
	}
	if !checkEqSlice(timev_.Add(durations_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Add(durations_).GetNullMask())
	}
	if !checkEqSlice(timev_.Add(durationv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Add(durationv_).GetNullMask())
	}

	// vector | string
	if !checkEqSlice(timev.Add(ss).Data().([]string), []string{"2020-01-01 00:00:00 +0000 UTC2", "2020-01-02 00:00:00 +0000 UTC2", "2020-01-03 00:00:00 +0000 UTC2", "2020-01-04 00:00:00 +0000 UTC2", "2020-01-05 00:00:00 +0000 UTC2", "2020-01-06 00:00:00 +0000 UTC2", "2020-01-07 00:00:00 +0000 UTC2", "2020-01-08 00:00:00 +0000 UTC2", "2020-01-09 00:00:00 +0000 UTC2", "2020-01-10 00:00:00 +0000 UTC2"}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []string{"2020-01-01 00:00:00 +0000 UTC2", "2020-01-02 00:00:00 +0000 UTC2", "2020-01-03 00:00:00 +0000 UTC2", "2020-01-04 00:00:00 +0000 UTC2", "2020-01-05 00:00:00 +0000 UTC2", "2020-01-06 00:00:00 +0000 UTC2", "2020-01-07 00:00:00 +0000 UTC2", "2020-01-08 00:00:00 +0000 UTC2", "2020-01-09 00:00:00 +0000 UTC2", "2020-01-10 00:00:00 +0000 UTC2"}, timev.Add(ss).Data())
	}
	if !checkEqSlice(timev.Add(sv).Data().([]string), []string{"2020-01-01 00:00:00 +0000 UTC1", "2020-01-02 00:00:00 +0000 UTC2", "2020-01-03 00:00:00 +0000 UTC3", "2020-01-04 00:00:00 +0000 UTC4", "2020-01-05 00:00:00 +0000 UTC5", "2020-01-06 00:00:00 +0000 UTC6", "2020-01-07 00:00:00 +0000 UTC7", "2020-01-08 00:00:00 +0000 UTC8", "2020-01-09 00:00:00 +0000 UTC9", "2020-01-10 00:00:00 +0000 UTC10"}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []string{"2020-01-01 00:00:00 +0000 UTC1", "2020-01-02 00:00:00 +0000 UTC2", "2020-01-03 00:00:00 +0000 UTC3", "2020-01-04 00:00:00 +0000 UTC4", "2020-01-05 00:00:00 +0000 UTC5", "2020-01-06 00:00:00 +0000 UTC6", "2020-01-07 00:00:00 +0000 UTC7", "2020-01-08 00:00:00 +0000 UTC8", "2020-01-09 00:00:00 +0000 UTC9", "2020-01-10 00:00:00 +0000 UTC10"}, timev.Add(sv).Data())
	}
	if !checkEqSlice(timev_.Add(ss_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Add(ss_).GetNullMask())
	}
	if !checkEqSlice(timev_.Add(sv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Add: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Add(sv_).GetNullMask())
	}
}

func Test_SeriesTime_Sub(t *testing.T) {
	dayNano := int64(24 * time.Hour.Nanoseconds())

	times := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	timev := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	times_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).SetNullMask([]bool{true})
	timev_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).
		SetNullMask([]bool{true, false, true, false, true, false, true, false, true, false})

	durations := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx)
	durationv := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx)
	durations_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano)}, nil, true, ctx).SetNullMask([]bool{true})
	durationv_ := NewSeriesDuration([]time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, nil, true, ctx).
		SetNullMask([]bool{false, true, false, true, false, true, false, true, false, true})

	// scalar | time
	if !checkEqSlice(times.Sub(times).Data().([]time.Duration), []time.Duration{0}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Duration{0}, times.Sub(times).Data())
	}
	if !checkEqSlice(times.Sub(timev).Data().([]time.Duration), []time.Duration{time.Duration(0), time.Duration(-1 * dayNano), time.Duration(-2 * dayNano), time.Duration(-3 * dayNano), time.Duration(-4 * dayNano), time.Duration(-5 * dayNano), time.Duration(-6 * dayNano), time.Duration(-7 * dayNano), time.Duration(-8 * dayNano), time.Duration(-9 * dayNano)}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Duration{time.Duration(0), time.Duration(-1 * dayNano), time.Duration(-2 * dayNano), time.Duration(-3 * dayNano), time.Duration(-4 * dayNano), time.Duration(-5 * dayNano), time.Duration(-6 * dayNano), time.Duration(-7 * dayNano), time.Duration(-8 * dayNano), time.Duration(-9 * dayNano)}, times.Sub(timev).Data())
	}
	if !checkEqSlice(times_.Sub(times_).GetNullMask(), []bool{true}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true}, times_.Sub(times_).GetNullMask())
	}
	if !checkEqSlice(times_.Sub(timev_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, times_.Sub(timev_).GetNullMask())
	}

	// scalar | duration
	if !checkEqSlice(times.Sub(durations).Data().([]time.Time), []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)}, times.Sub(durations).Data())
	}
	if !checkEqSlice(times.Sub(durationv).Data().([]time.Time), []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 30, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 29, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 28, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 27, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 26, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 24, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 23, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 22, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 30, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 29, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 28, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 27, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 26, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 24, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 23, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 22, 0, 0, 0, 0, time.UTC)}, times.Sub(durationv).Data())
	}
	if !checkEqSlice(times_.Sub(durations_).GetNullMask(), []bool{true}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true}, times_.Sub(durations_).GetNullMask())
	}
	if !checkEqSlice(times_.Sub(durationv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, times_.Sub(durationv_).GetNullMask())
	}

	// vector | time
	if !checkEqSlice(timev.Sub(times).Data().([]time.Duration), []time.Duration{0, time.Duration(dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano)}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Duration{0, time.Duration(dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano)}, timev.Sub(times).Data())
	}
	if !checkEqSlice(timev.Sub(timev).Data().([]time.Duration), []time.Duration{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Duration{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, timev.Sub(timev).Data())
	}
	if !checkEqSlice(timev_.Sub(times_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Sub(times_).GetNullMask())
	}
	if !checkEqSlice(timev_.Sub(timev_).GetNullMask(), []bool{true, false, true, false, true, false, true, false, true, false}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true, false, true, false, true, false, true, false, true, false}, timev_.Sub(timev_).GetNullMask())
	}

	// vector | duration
	if !checkEqSlice(timev.Sub(durations).Data().([]time.Time), []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC)}, timev.Sub(durations).Data())
	}
	if !checkEqSlice(timev.Sub(durationv).Data().([]time.Time), []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC), time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)}, timev.Sub(durationv).Data())
	}
	if !checkEqSlice(timev_.Sub(durations_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Sub(durations_).GetNullMask())
	}
	if !checkEqSlice(timev_.Sub(durationv_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Sub: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Sub(durationv_).GetNullMask())
	}
}

func Test_SeriesTime_Eq(t *testing.T) {
	// TODO: implement
}

func Test_SeriesTime_Ne(t *testing.T) {
	times := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	timev := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx)
	times_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).SetNullMask([]bool{true})
	timev_ := NewSeriesTime([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, nil, true, ctx).
		SetNullMask([]bool{true, false, true, false, true, false, true, false, true, false})

	// scalar | time
	if !checkEqSlice(times.Ne(times).Data().([]bool), []bool{false}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{false}, times.Ne(times).Data())
	}
	if !checkEqSlice(times.Ne(timev).Data().([]bool), []bool{false, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{false, true, true, true, true, true, true, true, true, true}, times.Ne(timev).Data())
	}
	if !checkEqSlice(times_.Ne(times_).GetNullMask(), []bool{true}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{true}, times_.Ne(times_).GetNullMask())
	}
	if !checkEqSlice(times_.Ne(timev_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, times_.Ne(timev_).GetNullMask())
	}

	// vector | time
	if !checkEqSlice(timev.Ne(times).Data().([]bool), []bool{false, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{false, true, true, true, true, true, true, true, true, true}, timev.Ne(times).Data())
	}
	if !checkEqSlice(timev.Ne(timev).Data().([]bool), []bool{false, false, false, false, false, false, false, false, false, false}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{false, false, false, false, false, false, false, false, false, false}, timev.Ne(timev).Data())
	}
	if !checkEqSlice(timev_.Ne(times_).GetNullMask(), []bool{true, true, true, true, true, true, true, true, true, true}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{true, true, true, true, true, true, true, true, true, true}, timev_.Ne(times_).GetNullMask())
	}
	if !checkEqSlice(timev_.Ne(timev_).GetNullMask(), []bool{true, false, true, false, true, false, true, false, true, false}, nil, "") {
		t.Errorf("SeriesTime.Eq: expected %v, got %v", []bool{true, false, true, false, true, false, true, false, true, false}, timev_.Ne(timev_).GetNullMask())
	}
}

func Test_SeriesTime_Lt(t *testing.T) {
	// TODO: implement
}

func Test_SeriesTime_Le(t *testing.T) {
	// TODO: implement
}

func Test_SeriesTime_Gt(t *testing.T) {
	// TODO: implement
}

func Test_SeriesTime_Ge(t *testing.T) {
	// TODO: implement
}
