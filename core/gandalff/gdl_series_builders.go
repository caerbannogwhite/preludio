package gandalff

import (
	"fmt"
	"time"
)

func NewSeries(data interface{}, nullMask []bool, makeCopy bool, memOpt bool, ctx *Context) Series {
	if ctx == nil {
		return SeriesError{"NewSeries: context is nil"}
	}

	switch data := data.(type) {
	case nil:
		return NewSeriesNA(1, ctx)
	case []bool:
		// if memOpt {
		// 	return NewSeriesBoolMemOpt(isNullable, makeCopy, data, ctx)
		// } else {
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesBool(data, nullMask, makeCopy, ctx)
		// }

	case []int:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesInt(data, nullMask, makeCopy, ctx)

	case []int64:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesInt64(data, nullMask, makeCopy, ctx)

	case []float64:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesFloat64(data, nullMask, makeCopy, ctx)

	case []string:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesString(data, nullMask, makeCopy, ctx)

	case []time.Time:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesTime(data, nullMask, makeCopy, ctx)

	case []time.Duration:
		if nullMask != nil && len(nullMask) != len(data) {
			return SeriesError{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesDuration(data, nullMask, makeCopy, ctx)

	default:
		return SeriesError{fmt.Sprintf("NewSeries: unsupported type %T", data)}
	}
}

// Build an Error Series
func NewSeriesError(err string) SeriesError {
	return SeriesError{msg: err}
}

// Build an NA Series
func NewSeriesNA(size int, ctx *Context) SeriesNA {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesNA: context is nil")
	}

	if size < 0 {
		size = 0
	}

	return SeriesNA{size: size, ctx: ctx}
}

// Build a Bool Series, if nullMask is nil then the series is not nullable
func NewSeriesBool(data []bool, nullMask []bool, makeCopy bool, ctx *Context) SeriesBool {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesBool: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
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

	return SeriesBool{
		isNullable: isNullable,
		data:       data,
		nullMask:   nullMask_,
		ctx:        ctx,
	}
}

// Build a Bool Series, if nullMask is nil then the series is not nullable
func newSeriesBoolMemOpt(data []bool, nullMask []bool, makeCopy bool, ctx *Context) SeriesBoolMemOpt {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesBoolMemOpt: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	size := len(data)
	actualData := __binVecFromBools(data)

	return SeriesBoolMemOpt{
		isNullable: isNullable,
		size:       size,
		data:       actualData,
		nullMask:   nullMask_,
		ctx:        ctx,
	}
}

// Build a Int Series, if nullMask is nil then the series is not nullable
func NewSeriesInt(data []int, nullMask []bool, makeCopy bool, ctx *Context) SeriesInt {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesInt: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int, len(data))
		copy(actualData, data)
		data = actualData
	}

	return SeriesInt{
		isNullable: isNullable,
		data:       data,
		nullMask:   nullMask_,
		ctx:        ctx,
	}
}

// Build a Int64 Series, if nullMask is nil then the series is not nullable
func NewSeriesInt64(data []int64, nullMask []bool, makeCopy bool, ctx *Context) SeriesInt64 {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesInt64: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
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

	return SeriesInt64{
		isNullable: isNullable,
		data:       data,
		nullMask:   nullMask_,
		ctx:        ctx,
	}
}

// Build a Float64 Series, if nullMask is nil then the series is not nullable
func NewSeriesFloat64(data []float64, nullMask []bool, makeCopy bool, ctx *Context) SeriesFloat64 {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesFloat64: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
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

	return SeriesFloat64{
		isNullable: isNullable,
		data:       data,
		nullMask:   nullMask_,
		ctx:        ctx,
	}
}

// Build a String Series, if nullMask is nil then the series is not nullable
func NewSeriesString(data []string, nullMask []bool, makeCopy bool, ctx *Context) SeriesString {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesString: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = __binVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	actualData := make([]*string, len(data))
	if nullMask != nil {
		for i, v := range data {
			if nullMask[i] {
				actualData[i] = ctx.stringPool.nullStringPtr
				continue
			}
			actualData[i] = ctx.stringPool.Put(v)
		}
	} else {
		for i, v := range data {
			actualData[i] = ctx.stringPool.Put(v)
		}
	}

	return SeriesString{
		isNullable: isNullable,
		data:       actualData,
		nullMask:   nullMask_,
		ctx:        ctx,
	}
}

// Build a Time Series, if nullMask is nil then the series is not nullable
func NewSeriesTime(data []time.Time, nullMask []bool, makeCopy bool, ctx *Context) SeriesTime {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesTime: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
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

	return SeriesTime{
		isNullable: isNullable,
		data:       data,
		nullMask:   nullMask_,
		ctx:        ctx,
		timeFormat: ctx.timeFormat,
	}
}

// Build a Duration Series, if nullMask is nil then the series is not nullable
func NewSeriesDuration(data []time.Duration, nullMask []bool, makeCopy bool, ctx *Context) SeriesDuration {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesDuration: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
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

	return SeriesDuration{
		isNullable: isNullable,
		data:       data,
		nullMask:   nullMask_,
		ctx:        ctx,
	}
}
