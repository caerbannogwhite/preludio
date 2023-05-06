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

	GetPool() *StringPool

	GetSeriesIndex(name string) int

	// Add new series to the dataframe.

	// AddSeries adds a generic series to the dataframe.
	AddSeries(series GDLSeries) DataFrame
	// AddSeriesFromBools adds a series of bools to the dataframe.
	AddSeriesFromBools(name string, isNullable bool, data []bool) DataFrame
	// AddSeriesFromInts adds a series of ints to the dataframe.
	AddSeriesFromInts(name string, isNullable bool, makeCopy bool, data []int) DataFrame
	// AddSeriesFromFloats adds a series of floats to the dataframe.
	AddSeriesFromFloats(name string, isNullable bool, makeCopy bool, data []float64) DataFrame
	// AddSeriesFromStrings adds a series of strings to the dataframe.
	AddSeriesFromStrings(name string, isNullable bool, data []string) DataFrame

	// Returns the series with the given name.
	Series(name string) GDLSeries

	// Returns the series at the given index.
	SeriesAt(index int) GDLSeries

	Select(names ...string) DataFrame

	SelectAt(indices ...int) DataFrame

	Filter(mask GDLSeriesBool) DataFrame

	GroupBy(by ...string) DataFrame

	Ungroup() DataFrame

	Join(how DataFrameJoinType, other DataFrame, on ...string) DataFrame

	Take(start, end, step int) DataFrame

	Count(name string) DataFrame
	Sum() DataFrame
	Mean() DataFrame
	Min() DataFrame
	Max() DataFrame
	Std() DataFrame

	PrettyPrint()

	FromCSV() *GDLCsvReader
}
