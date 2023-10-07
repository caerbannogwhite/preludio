package gandalff

import (
	"fmt"
	"time"
)

func NewSeries(data interface{}, nullMask []bool, makeCopy bool, memOpt bool, pool *StringPool) Series {
	switch data := data.(type) {
	case nil:
		return NewSeriesNA(1, pool)
	case []bool:
		// if memOpt {
		// 	return NewSeriesBoolMemOpt(isNullable, makeCopy, data, pool)
		// } else {
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesBool(data, nullMask, makeCopy, pool)
		// }

	case []int32:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesInt32(data, nullMask, makeCopy, pool)

	case []int64:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesInt64(data, nullMask, makeCopy, pool)

	case []float64:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesFloat64(data, nullMask, makeCopy, pool)

	case []string:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesString(data, nullMask, makeCopy, pool)

	case []time.Time:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesTime(data, nullMask, makeCopy, pool)

	case []time.Duration:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesDuration(data, nullMask, makeCopy, pool)

	default:
		return SeriesError{fmt.Sprintf("NewSeries: unsupported type %T", data)}
	}
}

func NewSeriesError(err string) SeriesError {
	return SeriesError{msg: err}
}

func NewSeriesNA(size int, pool *StringPool) SeriesNA {
	return SeriesNA{size: size, pool: pool}
}

func NewSeriesBool(data []bool, nullMask []bool, makeCopy bool, pool *StringPool) SeriesBool {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]bool, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesBool{isNullable: isNullable, data: data, nullMask: nullMask_, pool: pool}
}

func NewSeriesBoolMemOpt(data []bool, nullMask []bool, makeCopy bool, pool *StringPool) SeriesBoolMemOpt {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	size := len(data)
	actualData := __binVecFromBools(data)

	return SeriesBoolMemOpt{isNullable: isNullable, size: size, data: actualData, nullMask: nullMask_, pool: pool}
}

func NewSeriesInt32(data []int32, nullMask []bool, makeCopy bool, pool *StringPool) SeriesInt32 {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int32, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesInt32{isNullable: isNullable, data: data, nullMask: nullMask_, pool: pool}
}

func NewSeriesInt64(data []int64, nullMask []bool, makeCopy bool, pool *StringPool) SeriesInt64 {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesInt64{isNullable: isNullable, data: data, nullMask: nullMask_, pool: pool}
}

func NewSeriesFloat64(data []float64, nullMask []bool, makeCopy bool, pool *StringPool) SeriesFloat64 {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesFloat64{isNullable: isNullable, data: data, nullMask: nullMask_, pool: pool}
}

func NewSeriesString(data []string, nullMask []bool, makeCopy bool, pool *StringPool) SeriesString {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	actualData := make([]*string, len(data))
	for i, v := range data {
		if nullMask[i] {
			actualData[i] = pool.nullStringPtr
			continue
		}
		actualData[i] = pool.Put(v)
	}

	return SeriesString{isNullable: isNullable, data: actualData, nullMask: nullMask_, pool: pool}
}

func NewSeriesTime(data []time.Time, nullMask []bool, makeCopy bool, pool *StringPool) SeriesTime {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]time.Time, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesTime{isNullable: isNullable, data: data, nullMask: nullMask_, pool: pool}
}

func NewSeriesDuration(data []time.Duration, nullMask []bool, makeCopy bool, pool *StringPool) SeriesDuration {
	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]time.Duration, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesDuration{isNullable: isNullable, data: data, nullMask: nullMask_, pool: pool}
}
