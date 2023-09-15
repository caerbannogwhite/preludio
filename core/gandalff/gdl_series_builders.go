package gandalff

import "time"

func NewSeriesError(err string) SeriesError {
	return SeriesError{msg: err}
}

func NewSeriesBool(name string, isNullable, makeCopy bool, data []bool) Series {
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

	return SeriesBool{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

func NewSeriesBoolMemOpt(name string, isNullable bool, data []bool) Series {
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

	return SeriesBoolMemOpt{isNullable: isNullable, name: name, size: size, data: actualData, nullMask: nullMask}
}

func NewSeriesInt32(name string, isNullable, makeCopy bool, data []int32) Series {
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

	return SeriesInt32{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

func NewSeriesInt64(name string, isNullable, makeCopy bool, data []int64) Series {
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

	return SeriesInt64{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

func NewSeriesFloat64(name string, isNullable, makeCopy bool, data []float64) Series {
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

	return SeriesFloat64{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

func NewSeriesString(name string, isNullable bool, data []string, pool *StringPool) Series {
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

func NewSeriesTime(name string, isNullable, makeCopy bool, data []time.Time) Series {
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

	return SeriesTime{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}
