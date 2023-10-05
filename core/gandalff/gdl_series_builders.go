package gandalff

import (
	"fmt"
	"time"
)

func NewSeries(data interface{}, nullMask []bool, makeCopy bool, memOpt bool, pool *StringPool) Series {
	switch data := data.(type) {
	case []bool:
		// if memOpt {
		// 	return NewSeriesBoolMemOpt(isNullable, makeCopy, data, pool)
		// } else {
		return NewSeriesBool(data, nullMask, makeCopy, pool)
		// }
	case []int32:
		return NewSeriesInt32(data, nullMask, makeCopy, pool)
	case []int64:
		return NewSeriesInt64(data, nullMask, makeCopy, pool)
	case []float64:
		return NewSeriesFloat64(data, nullMask, makeCopy, pool)
	case []string:
		return NewSeriesString(data, nullMask, makeCopy, pool)
	case []time.Time:
		return NewSeriesTime(data, nullMask, makeCopy, pool)
	case []time.Duration:
		return NewSeriesDuration(data, nullMask, makeCopy, pool)
	default:
		return SeriesError{fmt.Sprintf("NewSeries: unsupported type %T", data)}
	}
}

func NewSeriesError(err string) SeriesError {
	return SeriesError{msg: err}
}

func NewSeriesBool(data []bool, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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

func NewSeriesBoolMemOpt(data []bool, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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

func NewSeriesInt32(data []int32, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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

func NewSeriesInt64(data []int64, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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

func NewSeriesFloat64(data []float64, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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

func NewSeriesString(data []string, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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
		actualData[i] = pool.Put(v)
	}

	return SeriesString{isNullable: isNullable, data: actualData, nullMask: nullMask_, pool: pool}
}

func NewSeriesTime(data []time.Time, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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

func NewSeriesDuration(data []time.Duration, nullMask []bool, makeCopy bool, pool *StringPool) Series {
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
