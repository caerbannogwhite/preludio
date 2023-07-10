package gandalff

func NewSeriesBool(name string, isNullable bool, data []bool) Series {
	size := len(data)
	var actualData []uint8
	if size%8 == 0 {
		actualData = make([]uint8, (size >> 3))
		for i := 0; i < size; i++ {
			if data[i] {
				actualData[i>>3] |= 1 << uint(i%8)
			}
		}
	} else {
		actualData = make([]uint8, (size>>3)+1)
		for i := 0; i < size; i++ {
			if data[i] {
				actualData[i>>3] |= 1 << uint(i%8)
			}
		}
	}

	var nullMask []uint8
	if isNullable {
		nullMask = make([]uint8, len(actualData))

	} else {
		nullMask = make([]uint8, 0)
	}

	return SeriesBool{isNullable: isNullable, name: name, size: size, data: actualData, nullMask: nullMask}
}

func NewSeriesInt32(name string, isNullable bool, makeCopy bool, data []int32) Series {
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

func NewSeriesInt64(name string, isNullable bool, makeCopy bool, data []int64) Series {
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

func NewSeriesFloat64(name string, isNullable bool, makeCopy bool, data []float64) Series {
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
