package gandalff

type GDataFrame struct {
	isGrouped bool
	series    []GSeries
	pool      *StringPool
}

func NewGDataFrame() *GDataFrame {
	return &GDataFrame{
		series: make([]GSeries, 0),
		pool:   NewStringPool(),
	}
}

func (df *GDataFrame) IsGrouped() bool {
	return df.isGrouped
}

func (df *GDataFrame) GetPool() *StringPool {
	return df.pool
}

func (df *GDataFrame) AddSeries(series GSeries) {
	df.series = append(df.series, series)
}

func (df *GDataFrame) AddSeriesFromBools(name string, isNullable bool, makeCopy bool, data []bool) {
	series := NewGSeriesBool(name, isNullable, makeCopy, data)
	df.AddSeries(series)
}

func (df *GDataFrame) AddSeriesFromInts(name string, isNullable bool, makeCopy bool, data []int) {
	series := NewGSeriesInt(name, isNullable, makeCopy, data)
	df.AddSeries(series)
}

func (df *GDataFrame) AddSeriesFromFloats(name string, isNullable bool, makeCopy bool, data []float64) {
	series := NewGSeriesFloat(name, isNullable, makeCopy, data)
	df.AddSeries(series)
}

func (df *GDataFrame) AddSeriesFromStrings(name string, isNullable bool, data []string) {
	series := NewGSeriesString(name, isNullable, data, df.pool)
	df.AddSeries(series)
}

// Names returns the names of the series in the dataframe.
func (df *GDataFrame) Names() []string {
	names := make([]string, len(df.series))
	for i, series := range df.series {
		names[i] = series.Name()
	}
	return names
}

// Types returns the types of the series in the dataframe.
func (df *GDataFrame) Types() []GSeriesType {
	types := make([]GSeriesType, len(df.series))
	for i, series := range df.series {
		types[i] = series.Type()
	}
	return types
}

// Series returns the series with the given name.
func (df *GDataFrame) Series(name string) GSeries {
	for _, series := range df.series {
		if series.Name() == name {
			return series
		}
	}
	return nil
}

func (df *GDataFrame) SeriesAt(index int) GSeries {
	if index < 0 || index >= len(df.series) {
		return nil
	}
	return df.series[index]
}

func (df *GDataFrame) NCols() int {
	return len(df.series)
}

func (df *GDataFrame) NRows() int {
	if len(df.series) == 0 {
		return 0
	}
	return df.series[0].Len()
}

func (df *GDataFrame) Select(names ...string) *GDataFrame {
	selected := NewGDataFrame()
	for _, name := range names {
		series := df.Series(name)
		if series != nil {
			selected.AddSeries(series)
		}
	}
	return selected
}

func (df *GDataFrame) InPlaceSelect(names ...string) {
	selected := df.Select(names...)
	df.series = selected.series
}

func (df *GDataFrame) SelectAt(indices ...int) *GDataFrame {
	selected := NewGDataFrame()
	for _, index := range indices {
		series := df.SeriesAt(index)
		if series != nil {
			selected.AddSeries(series)
		}
	}
	return selected
}

func (df *GDataFrame) InPlaceSelectAt(indices ...int) {
	selected := df.SelectAt(indices...)
	df.series = selected.series
}

func (df *GDataFrame) Filter() *GDataFrame {
	filtered := NewGDataFrame()

	return filtered
}

func (df *GDataFrame) GroupBy(by ...string) *GDataFrame {
	grouped := NewGDataFrame()

	return grouped
}
