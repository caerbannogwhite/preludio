package gandalff

import (
	"fmt"
	"time"
)

func NewSeries(isNullable, makeCopy bool, memOpt bool, data interface{}, pool *StringPool) Series {
	switch data := data.(type) {
	case []bool:
		// if memOpt {
		// 	return NewSeriesBoolMemOpt(isNullable, makeCopy, data, pool)
		// } else {
		return NewSeriesBool(isNullable, makeCopy, data, pool)
		// }
	case []int32:
		return NewSeriesInt32(isNullable, makeCopy, data, pool)
	case []int64:
		return NewSeriesInt64(isNullable, makeCopy, data, pool)
	case []float64:
		return NewSeriesFloat64(isNullable, makeCopy, data, pool)
	case []string:
		return NewSeriesString(isNullable, makeCopy, data, pool)
	case []time.Time:
		return NewSeriesTime(isNullable, makeCopy, data, pool)
	case []time.Duration:
		return NewSeriesDuration(isNullable, makeCopy, data, pool)
	default:
		return SeriesError{fmt.Sprintf("NewSeries: unsupported type %T", data)}
	}
}

func NewSeriesError(err string) SeriesError {
	return SeriesError{msg: err}
}

func NewSeriesBool(isNullable, makeCopy bool, data []bool, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]bool, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesBool{isNullable: isNullable, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesBoolMemOpt(isNullable bool, makeCopy bool, data []bool, pool *StringPool) Series {
	size := len(data)
	actualData := __binVecInit(size)
	for i := 0; i < size; i++ {
		if data[i] {
			actualData[i>>3] |= 1 << uint(i%8)
		}
	}

	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	return SeriesBoolMemOpt{isNullable: isNullable, size: size, data: actualData, nullMask: nullMask, pool: pool}
}

func NewSeriesInt32(isNullable, makeCopy bool, data []int32, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int32, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesInt32{isNullable: isNullable, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesInt64(isNullable, makeCopy bool, data []int64, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesInt64{isNullable: isNullable, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesFloat64(isNullable, makeCopy bool, data []float64, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesFloat64{isNullable: isNullable, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesString(isNullable bool, makeCopy bool, data []string, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	actualData := make([]*string, len(data))
	for i, v := range data {
		actualData[i] = pool.Put(v)
	}

	return SeriesString{isNullable: isNullable, data: actualData, nullMask: nullMask, pool: pool}
}

func NewSeriesTime(isNullable, makeCopy bool, data []time.Time, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]time.Time, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesTime{isNullable: isNullable, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesDuration(isNullable, makeCopy bool, data []time.Duration, pool *StringPool) Series {
	var nullMask []uint8
	if isNullable {
		nullMask = __binVecInit(len(data))
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]time.Duration, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesDuration{isNullable: isNullable, data: data, nullMask: nullMask, pool: pool}
}
