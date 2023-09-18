package gandalff

import (
	"time"
)

func NewSeries(name string, isNullable, makeCopy bool, memOpt bool, data interface{}, pool *StringPool) Series {
	switch data.(type) {
	case []bool:
		if memOpt {
			return NewSeriesBoolMemOpt(name, isNullable, makeCopy, data.([]bool), pool)
		} else {
			return NewSeriesBool(name, isNullable, makeCopy, data.([]bool), pool)
		}
	case []int32:
		return NewSeriesInt32(name, isNullable, makeCopy, data.([]int32), pool)
	case []int64:
		return NewSeriesInt64(name, isNullable, makeCopy, data.([]int64), pool)
	case []float64:
		return NewSeriesFloat64(name, isNullable, makeCopy, data.([]float64), pool)
	case []string:
		return NewSeriesString(name, isNullable, makeCopy, data.([]string), pool)
	case []time.Time:
		return NewSeriesTime(name, isNullable, makeCopy, data.([]time.Time), pool)
	default:
		return SeriesError{msg: "NewSeries: unknown type"}
	}
}

func NewSeriesError(err string) SeriesError {
	return SeriesError{msg: err}
}

func NewSeriesBool(name string, isNullable, makeCopy bool, data []bool, pool *StringPool) Series {
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

	return SeriesBool{isNullable: isNullable, name: name, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesBoolMemOpt(name string, isNullable bool, makeCopy bool, data []bool, pool *StringPool) Series {
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

	return SeriesBoolMemOpt{isNullable: isNullable, name: name, size: size, data: actualData, nullMask: nullMask, pool: pool}
}

func NewSeriesInt32(name string, isNullable, makeCopy bool, data []int32, pool *StringPool) Series {
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

	return SeriesInt32{isNullable: isNullable, name: name, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesInt64(name string, isNullable, makeCopy bool, data []int64, pool *StringPool) Series {
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

	return SeriesInt64{isNullable: isNullable, name: name, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesFloat64(name string, isNullable, makeCopy bool, data []float64, pool *StringPool) Series {
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

	return SeriesFloat64{isNullable: isNullable, name: name, data: data, nullMask: nullMask, pool: pool}
}

func NewSeriesString(name string, isNullable bool, makeCopy bool, data []string, pool *StringPool) Series {
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

	return SeriesString{isNullable: isNullable, name: name, data: actualData, nullMask: nullMask, pool: pool}
}

func NewSeriesTime(name string, isNullable, makeCopy bool, data []time.Time, pool *StringPool) Series {
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

	return SeriesTime{isNullable: isNullable, name: name, data: data, nullMask: nullMask, pool: pool}
}
