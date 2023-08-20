package gandalff

import "typesys"

type DataFrameJoinType int8

const (
	INNER_JOIN DataFrameJoinType = iota
	LEFT_JOIN
	RIGHT_JOIN
	OUTER_JOIN
)

type DataFrame interface {

	// Basic accessors.

	// Names returns the names of the series in the dataframe.
	Names() []string
	// Types returns the types of the series in the dataframe.
	Types() []typesys.BaseType
	// NCols returns the number of columns in the dataframe.
	NCols() int
	// NRows returns the number of rows in the dataframe.
	NRows() int

	IsErrored() bool

	IsGrouped() bool

	GetError() error

	GetStringPool() *StringPool
	SetStringPool(pool *StringPool) DataFrame

	GetSeriesIndex(name string) int

	// Add new series to the dataframe.

	// AddSeries adds a generic series to the dataframe.
	AddSeries(series ...Series) DataFrame
	// AddSeriesFromBool adds a series of bools to the dataframe.
	AddSeriesFromBool(name string, isNullable, makeCopy bool, data []bool) DataFrame
	// AddSeriesFromInt32 adds a series of ints to the dataframe.
	AddSeriesFromInt32(name string, isNullable, makeCopy bool, data []int32) DataFrame
	// AddSeriesFromInt64 adds a series of ints to the dataframe.
	AddSeriesFromInt64(name string, isNullable, makeCopy bool, data []int64) DataFrame
	// AddSeriesFromFloat adds a series of floats to the dataframe.
	AddSeriesFromFloat64(name string, isNullable, makeCopy bool, data []float64) DataFrame
	// AddSeriesFromString adds a series of strings to the dataframe.
	AddSeriesFromString(name string, isNullable bool, data []string) DataFrame

	// Replace the series with the given name.
	Replace(name string, s Series) DataFrame

	// Returns the series with the given name.
	Series(name string) Series

	// Returns the series at the given index.
	SeriesAt(index int) Series

	Select(names ...string) DataFrame

	SelectAt(indices ...int) DataFrame

	Filter(mask SeriesBool) DataFrame

	GroupBy(by ...string) DataFrame

	Ungroup() DataFrame

	getPartitions() []SeriesPartition

	Join(how DataFrameJoinType, other DataFrame, on ...string) DataFrame

	Take(params ...int) DataFrame

	Agg(...aggregator) DataFrame

	// Sort the dataframe.
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	OrderBy(params ...SortParam) DataFrame

	// IO

	Describe() string
	Records(header bool) [][]string
	PrettyPrint(nrows ...int) DataFrame

	FromCSV() *CsvReader
	ToCSV() *CsvWriter
}

////////////////////////			AGGREGATORS

type AggregateType int8

const (
	AGGREGATE_COUNT AggregateType = iota
	AGGREGATE_SUM
	AGGREGATE_MEAN
	AGGREGATE_MEDIAN
	AGGREGATE_MIN
	AGGREGATE_MAX
	AGGREGATE_STD
)

const DEFAULT_COUNT_NAME = "n"

type aggregator interface {
	getSeriesName() string
	getAggregateType() AggregateType
}

type countAggregator struct {
	name  string
	type_ AggregateType
}

func (agg countAggregator) getSeriesName() string {
	return agg.name
}

func (agg countAggregator) getAggregateType() AggregateType {
	return agg.type_
}

func Count() aggregator {

	return countAggregator{DEFAULT_COUNT_NAME, AGGREGATE_COUNT}
}

type sumAggregator struct {
	name  string
	type_ AggregateType
}

func (agg sumAggregator) getSeriesName() string {
	return agg.name
}

func (agg sumAggregator) getAggregateType() AggregateType {
	return agg.type_
}

func Sum(name string) aggregator {
	return sumAggregator{name, AGGREGATE_SUM}
}

type meanAggregator struct {
	name  string
	type_ AggregateType
}

func (agg meanAggregator) getSeriesName() string {
	return agg.name
}

func (agg meanAggregator) getAggregateType() AggregateType {
	return agg.type_
}

func Mean(name string) aggregator {
	return meanAggregator{name, AGGREGATE_MEAN}
}

type medianAggregator struct {
	name  string
	type_ AggregateType
}

func (agg medianAggregator) getSeriesName() string {
	return agg.name
}

func (agg medianAggregator) getAggregateType() AggregateType {
	return agg.type_
}

func Median(name string) aggregator {
	return medianAggregator{name, AGGREGATE_MEDIAN}
}

type minAggregator struct {
	name  string
	type_ AggregateType
}

func (agg minAggregator) getSeriesName() string {
	return agg.name
}

func (agg minAggregator) getAggregateType() AggregateType {
	return agg.type_
}

func Min(name string) aggregator {
	return minAggregator{name, AGGREGATE_MIN}
}

type maxAggregator struct {
	name  string
	type_ AggregateType
}

func (agg maxAggregator) getSeriesName() string {
	return agg.name
}

func (agg maxAggregator) getAggregateType() AggregateType {
	return agg.type_
}

func Max(name string) aggregator {
	return maxAggregator{name, AGGREGATE_MAX}
}

type stdAggregator struct {
	name  string
	type_ AggregateType
}

func (agg stdAggregator) getSeriesName() string {
	return agg.name
}

func (agg stdAggregator) getAggregateType() AggregateType {
	return agg.type_
}

func Std(name string) aggregator {
	return stdAggregator{name, AGGREGATE_STD}
}

////////////////////////			SORT

type SortParam struct {
	asc    bool
	name   string
	series Series
}

func Asc(name string) SortParam {
	return SortParam{asc: true, name: name}
}

func Desc(name string) SortParam {
	return SortParam{asc: false, name: name}
}
