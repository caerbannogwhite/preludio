package gandalff

import (
	"testing"
	"time"
)

func Test_SeriesTime_Map(t *testing.T) {

	baseTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	data := []time.Time{
		baseTime, baseTime.AddDate(0, 0, 1), baseTime.AddDate(0, 0, 2), baseTime.AddDate(0, 0, 3), baseTime.AddDate(0, 0, 4), baseTime.AddDate(0, 0, 5), baseTime.AddDate(0, 0, 6), baseTime.AddDate(0, 0, 7), baseTime.AddDate(0, 0, 8), baseTime.AddDate(0, 0, 9),
		baseTime.AddDate(0, 0, 10), baseTime.AddDate(0, 0, 11), baseTime.AddDate(0, 0, 12), baseTime.AddDate(0, 0, 13), baseTime.AddDate(0, 0, 14), baseTime.AddDate(0, 0, 15), baseTime.AddDate(0, 0, 16), baseTime.AddDate(0, 0, 17), baseTime.AddDate(0, 0, 18), baseTime.AddDate(0, 0, 19),
	}
	nullMask := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true, false}

	// Create a new series.
	s := NewSeriesTime("test", true, false, data, NewStringPool()).
		SetNullMask(nullMask)

	// Map the series to bool.
	resBool := s.Map(func(v any) any {
		return v.(time.Time).Day()%2 == 0
	})

	expectedBool := []bool{false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true}
	if !checkEqSliceBool(resBool.Data().([]bool), expectedBool, nil, "") {
		t.Errorf("SeriesTime.Map: expected %v, got %v", expectedBool, resBool.Data())
	}

	// Map the series to int32.
	resInt := s.Map(func(v any) any {
		return int32(v.(time.Time).Day())
	})

	expectedInt := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	if !checkEqSliceInt32(resInt.Data().([]int32), expectedInt, nil, "") {
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
	pool := NewStringPool()

	times := NewSeriesTime("test", true, false, []time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, pool).(SeriesTime)
	timev := NewSeriesTime("test", true, false, []time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, pool).(SeriesTime)
	times_ := NewSeriesTime("test", true, false, []time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, pool).SetNullMask([]bool{true}).(SeriesTime)
	timev_ := NewSeriesTime("test", true, false, []time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}, pool).(SeriesTime)

	durations := NewSeriesDuration("test", true, false, []time.Duration{time.Duration(1 * dayNano)}, pool).(SeriesDuration)
	durationv := NewSeriesDuration("test", true, false, []time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, pool).(SeriesDuration)
	durations_ := NewSeriesDuration("test", true, false, []time.Duration{time.Duration(1 * dayNano)}, pool).SetNullMask([]bool{true}).(SeriesDuration)
	durationv_ := NewSeriesDuration("test", true, false, []time.Duration{time.Duration(1 * dayNano), time.Duration(2 * dayNano), time.Duration(3 * dayNano), time.Duration(4 * dayNano), time.Duration(5 * dayNano), time.Duration(6 * dayNano), time.Duration(7 * dayNano), time.Duration(8 * dayNano), time.Duration(9 * dayNano), time.Duration(10 * dayNano)}, pool).SetNullMask([]bool{true, true, true, true, true, true, true, true, true, true}).(SeriesDuration)

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
}
