package pseries

import (
	"regexp"
	"strings"
	"time"
)

type PSeriesType uint8

const (
	BoolType PSeriesType = iota
	IntType
	FloatType
	StringType
	TimeType
	DurationType
	ComplexType
	InterfaceType
)

func PSeriesTypeToString(t PSeriesType) string {
	switch t {
	case BoolType:
		return "bool"
	case IntType:
		return "int"
	case FloatType:
		return "float"
	case StringType:
		return "string"
	case TimeType:
		return "time"
	case DurationType:
		return "duration"
	case ComplexType:
		return "complex"
	case InterfaceType:
		return "interface"
	}
	return "unknown"
}

type PSeries interface {
	// Returns the length of the series.
	Len() int
	// Returns the name of the series.
	Name() string
	// Returns the type of the series.
	Type() PSeriesType
	// Returns the data of the series.
	Data() interface{}
	// Returns the data of the series as a slice of strings.
	StringData() []string
	// Returns the data of the series as a slice of bools.
	BoolData() []bool
	// Returns the data of the series as a slice of ints.
	IntData() []int
	// Returns the data of the series as a slice of floats.
	FloatData() []float64
	// Returns the data of the series as a slice of time.Time.
	TimeData() []time.Time
	// Returns the data of the series as a slice of time.Duration.
	DurationData() []time.Duration
	// Returns the data of the series as a slice of complex128.
	ComplexData() []complex128
	// Returns the data of the series as a slice of interface{}.
}

// PSeriesBool represents a series of bools.
type PSeriesBool struct {
	name string
	data []bool
}

func NewPSeriesBool(name string, data []bool) PSeriesBool {
	return PSeriesBool{name: name, data: data}
}

func (s PSeriesBool) Len() int {
	return len(s.data)
}

func (s PSeriesBool) Name() string {
	return s.name
}

func (s PSeriesBool) Type() PSeriesType {
	return BoolType
}

func (s PSeriesBool) Data() interface{} {
	return s.data
}

// PseriesString represents a series of strings.
// The data is stored as pointers to strings to save memory.
// The actual strings are stored in a separate string pool.
type PSeriesString struct {
	name string
	data []*string
}

func NewPSeriesString(name string, data []string) PSeriesString {
	var s PSeriesString
	s.name = name
	s.data = make([]*string, len(data))
	for i, v := range data {
		s.data[i] = &v
	}
	return s
}

func (s PSeriesString) Len() int {
	return len(s.data)
}

func (s PSeriesString) Name() string {
	return s.name
}

func (s PSeriesString) Type() PSeriesType {
	return StringType
}

func (s PSeriesString) Data() interface{} {
	return s.data
}

func (s PSeriesString) StringData() []string {
	var data []string
	for _, v := range s.data {
		data = append(data, *v)
	}
	return data
}

// StringPool returns the string pool of the series.
func (s PSeriesString) StringPool() []string {
	var pool []string
	for _, v := range s.data {
		pool = append(pool, *v)
	}
	return pool
}

type PDataFrame struct {
	series []PSeries
}

func NewPDataFrame(series []PSeries) PDataFrame {
	return PDataFrame{series: series}
}

func (df PDataFrame) Len() int {
	return len(df.series)
}

func (df PDataFrame) Series() []PSeries {
	return df.series
}

func (df PDataFrame) SeriesByName(name string) PSeries {
	for _, s := range df.series {
		if s.Name() == name {
			return s
		}
	}
	return nil
}

func (df PDataFrame) SeriesByType(t PSeriesType) []PSeries {
	var series []PSeries
	for _, s := range df.series {
		if s.Type() == t {
			series = append(series, s)
		}
	}
	return series
}

func (df PDataFrame) SeriesByTypeAndName(t PSeriesType, name string) PSeries {
	for _, s := range df.series {
		if s.Type() == t && s.Name() == name {
			return s
		}
	}
	return nil
}

func (df PDataFrame) SeriesByTypeAndNamePrefix(t PSeriesType, prefix string) []PSeries {
	var series []PSeries
	for _, s := range df.series {
		if s.Type() == t && strings.HasPrefix(s.Name(), prefix) {
			series = append(series, s)
		}
	}
	return series
}

func (df PDataFrame) SeriesByTypeAndNameSuffix(t PSeriesType, suffix string) []PSeries {
	var series []PSeries
	for _, s := range df.series {
		if s.Type() == t && strings.HasSuffix(s.Name(), suffix) {
			series = append(series, s)
		}
	}
	return series
}

func (df PDataFrame) SeriesByTypeAndNameContains(t PSeriesType, contains string) []PSeries {
	var series []PSeries
	for _, s := range df.series {
		if s.Type() == t && strings.Contains(s.Name(), contains) {
			series = append(series, s)
		}
	}
	return series
}

func (df PDataFrame) SeriesByTypeAndNameRegex(t PSeriesType, regex string) []PSeries {
	var series []PSeries
	for _, s := range df.series {
		if s.Type() == t && regexp.MustCompile(regex).MatchString(s.Name()) {
			series = append(series, s)
		}
	}
	return series
}

func (df PDataFrame) SeriesByTypeAndNameRegexPrefix(t PSeriesType, regex string) []PSeries {
	var series []PSeries
	for _, s := range df.series {
		if s.Type() == t && regexp.MustCompile("^"+regex).MatchString(s.Name()) {
			series = append(series, s)
		}
	}
	return series
}

// Join joins the series of the dataframe with the series of the other dataframe.
func (df PDataFrame) Join(other PDataFrame) PDataFrame {
	var series []PSeries
	for _, s := range df.series {
		series = append(series, s)
	}
	for _, s := range other.series {
		series = append(series, s)
	}
	return PDataFrame{series: series}
}

// Join joins the series of the dataframe with the series of the other dataframe.
func (df PDataFrame) JoinSeries(series []PSeries) PDataFrame {
	for _, s := range series {
		df.series = append(df.series, s)
	}
	return df
}

// Join joins the series of the dataframe with the series of the other dataframe.
func (df PDataFrame) JoinSeriesByName(name string, series []PSeries) PDataFrame {
	for _, s := range series {
		df.series = append(df.series, NewPSeriesBool(name, s.BoolData()))
	}
	return df
}

// Join joins the series of the dataframe with the series of the other dataframe.
func (df PDataFrame) JoinSeriesByType(t PSeriesType, series []PSeries) PDataFrame {
	for _, s := range series {
		df.series = append(df.series, NewPSeriesBool(s.Name(), s.BoolData()))
	}
	return df
}

// Join joins the series of the dataframe with the series of the other dataframe.
func (df PDataFrame) JoinSeriesByTypeAndName(t PSeriesType, name string, series []PSeries) PDataFrame {
	for _, s := range series {
		df.series = append(df.series, NewPSeriesBool(name, s.BoolData()))
	}
	return df
}

// The string pool is a struct that
type StringPool struct {
	pool map[string]string
}
