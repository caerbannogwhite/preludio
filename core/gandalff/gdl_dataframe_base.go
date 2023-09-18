package gandalff

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
	"typesys"
)

type BaseDataFramePartitionEntry struct {
	index     int
	name      string
	partition SeriesPartition
}

type BaseDataFrame struct {
	isGrouped  bool
	err        error
	series     []Series
	pool       *StringPool
	partitions []BaseDataFramePartitionEntry
	sortParams []SortParam
}

func NewBaseDataFrame() DataFrame {
	return &BaseDataFrame{
		series: make([]Series, 0),
		pool:   NewStringPool(),
	}
}

////////////////////////			BASIC ACCESSORS

// Names returns the names of the series in the dataframe.
func (df BaseDataFrame) Names() []string {
	names := make([]string, len(df.series))
	for i, series := range df.series {
		names[i] = series.Name()
	}
	return names
}

// Types returns the types of the series in the dataframe.
func (df BaseDataFrame) Types() []typesys.BaseType {
	types := make([]typesys.BaseType, len(df.series))
	for i, series := range df.series {
		types[i] = series.Type()
	}
	return types
}

// NCols returns the number of columns in the dataframe.
func (df BaseDataFrame) NCols() int {
	return len(df.series)
}

// NRows returns the number of rows in the dataframe.
func (df BaseDataFrame) NRows() int {
	if len(df.series) == 0 {
		return 0
	}
	return df.series[0].Len()
}

func (df BaseDataFrame) IsErrored() bool {
	return df.err != nil
}

func (df BaseDataFrame) IsGrouped() bool {
	return df.isGrouped
}

func (df BaseDataFrame) GetError() error {
	return df.err
}

func (df BaseDataFrame) GetStringPool() *StringPool {
	return df.pool
}

func (df BaseDataFrame) SetStringPool(pool *StringPool) DataFrame {
	for i, series := range df.series {
		if s, ok := series.(SeriesString); ok {
			df.series[i] = s.SetStringPool(pool)
		}
	}

	df.pool = pool
	return df
}

func (df BaseDataFrame) GetSeriesIndex(name string) int {
	for i, series := range df.series {
		if series.Name() == name {
			return i
		}
	}
	return -1
}

func (df BaseDataFrame) AddSeries(series ...Series) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeries: cannot add series to a grouped dataframe")
		return df
	}

	for _, series_ := range series {
		if df.NCols() > 0 && series_.Len() != df.NRows() {
			df.err = fmt.Errorf("BaseDataFrame.AddSeries: series length (%d) does not match dataframe length (%d)", series_.Len(), df.NRows())
			return df
		}

		df.series = append(df.series, series_)
	}
	return df
}

func (df BaseDataFrame) AddSeriesFromBool(name string, isNullable, makeCopy bool, data []bool) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromBool: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromBool: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	return df.AddSeries(NewSeriesBool(name, isNullable, makeCopy, data, df.pool))
}

func (df BaseDataFrame) AddSeriesFromInt32(name string, isNullable, makeCopy bool, data []int32) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromInt32: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromInt32: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	return df.AddSeries(NewSeriesInt32(name, isNullable, makeCopy, data, df.pool))
}

func (df BaseDataFrame) AddSeriesFromInt64(name string, isNullable, makeCopy bool, data []int64) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromInt64: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromInt64: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	return df.AddSeries(NewSeriesInt64(name, isNullable, makeCopy, data, df.pool))
}

func (df BaseDataFrame) AddSeriesFromFloat64(name string, isNullable, makeCopy bool, data []float64) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromFloat64: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromFloat64: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	return df.AddSeries(NewSeriesFloat64(name, isNullable, makeCopy, data, df.pool))
}

func (df BaseDataFrame) AddSeriesFromString(name string, isNullable bool, makeCopy bool, data []string) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromString: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromString: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	return df.AddSeries(NewSeriesString(name, isNullable, makeCopy, data, df.pool))
}

func (df BaseDataFrame) AddSeriesFromTime(name string, isNullable, makeCopy bool, data []time.Time) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromTime: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromTime: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	return df.AddSeries(NewSeriesTime(name, isNullable, makeCopy, data, df.pool))
}

func (df BaseDataFrame) Replace(name string, s Series) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.Replace: cannot replace series in a grouped dataframe")
		return df
	}

	index := df.GetSeriesIndex(name)
	if index == -1 {
		df.err = fmt.Errorf("BaseDataFrame.Replace: series \"%s\" not found", name)
		return df
	}

	if s.Len() != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.Replace: series length (%d) does not match dataframe length (%d)", s.Len(), df.NRows())
		return df
	}

	df.series[index] = s.SetName(name)
	return df
}

// Returns the series with the given name.
func (df BaseDataFrame) Series(name string) Series {
	for _, series := range df.series {
		if series.Name() == name {
			return series
		}
	}
	return SeriesError{msg: fmt.Sprintf("BaseDataFrame.Series: series \"%s\" not found", name)}
}

// Returns the series with the given name.
// For internal use only: returns nil if the series is not found.
func (df BaseDataFrame) __series(name string) Series {
	for _, series := range df.series {
		if series.Name() == name {
			return series
		}
	}
	return nil
}

// Returns the series at the given index.
func (df BaseDataFrame) SeriesAt(index int) Series {
	if index < 0 || index >= len(df.series) {
		return SeriesError{msg: fmt.Sprintf("BaseDataFrame.SeriesAt: index %d out of bounds", index)}
	}
	return df.series[index]
}

func (df BaseDataFrame) Select(names ...string) DataFrame {
	if df.err != nil {
		return df
	}

	seriesList := make([]Series, 0)
	for _, name := range names {
		series := df.__series(name)
		if series != nil {
			seriesList = append(seriesList, series)
		} else {
			return BaseDataFrame{
				err: fmt.Errorf("BaseDataFrame.Select: series \"%s\" not found", name),
			}
		}
	}

	return BaseDataFrame{
		series: seriesList,
		pool:   df.pool,
	}
}

func (df BaseDataFrame) SelectAt(indices ...int) DataFrame {
	if df.err != nil {
		return df
	}

	selected := NewBaseDataFrame()
	for _, index := range indices {
		if index < 0 || index >= len(df.series) {
			selected.AddSeries(df.series[index])
		} else {
			return BaseDataFrame{err: fmt.Errorf("BaseDataFrame.SelectAt: index %d out of bounds", index)}
		}
	}

	return selected
}

func (df BaseDataFrame) Filter(mask SeriesBool) DataFrame {
	if df.err != nil {
		return df
	}

	if mask.Len() != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.Filter: mask length (%d) does not match dataframe length (%d)", mask.Len(), df.NRows())
		return df
	}

	seriesList := make([]Series, 0)
	for _, series := range df.series {
		seriesList = append(seriesList, series.Filter(mask))
	}

	return BaseDataFrame{
		series: seriesList,
		pool:   df.pool,
	}
}

func (df BaseDataFrame) GroupBy(by ...string) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		// TODO: figure out what to do here
		return df
	} else {

		// Check that all the group by columns exist
		for _, name := range by {
			found := false
			for _, series := range df.series {
				if series.Name() == name {
					found = true
					break
				}
			}
			if !found {
				df.err = fmt.Errorf("BaseDataFrame.GroupBy: column \"%s\" not found", name)
				return df
			}
		}

		df.isGrouped = true
		df.partitions = make([]BaseDataFramePartitionEntry, len(by))

		for partitionsIndex, name := range by {

			i := df.GetSeriesIndex(name)
			series := df.series[i]

			// First partition: group the series
			if partitionsIndex == 0 {
				df.partitions[partitionsIndex] = BaseDataFramePartitionEntry{
					index:     i,
					name:      name,
					partition: series.group().GetPartition(),
				}
			} else

			// Subsequent partitions: sub-group the series
			{
				df.partitions[partitionsIndex] = BaseDataFramePartitionEntry{
					index:     i,
					name:      name,
					partition: series.GroupBy(df.partitions[partitionsIndex-1].partition).GetPartition(),
				}
			}
		}

		return df
	}
}

func (df BaseDataFrame) Ungroup() DataFrame {
	if df.err != nil {
		return df
	}

	df.isGrouped = false
	df.partitions = nil
	return df
}

func (df BaseDataFrame) getPartitions() []SeriesPartition {
	if df.err != nil {
		return nil
	}

	if df.isGrouped {
		partitions := make([]SeriesPartition, len(df.partitions))
		for i, partition := range df.partitions {
			partitions[i] = partition.partition
		}
		return partitions
	} else {
		return nil
	}
}

func (df BaseDataFrame) groupHelper() (DataFrame, *[][]int, *[]int) {

	// Keep track of which series are not grouped
	seriesIndices := make(map[int]bool)
	for i := 0; i < df.NCols(); i++ {
		seriesIndices[i] = true
	}

	result := NewBaseDataFrame().(*BaseDataFrame)

	// The last partition tells us how many groups there are
	// and how many rows are in each group
	indeces := make([][]int, 0, df.partitions[len(df.partitions)-1].partition.getSize())
	for _, group := range df.partitions[len(df.partitions)-1].partition.getMap() {
		indeces = append(indeces, group)
	}

	// Keep only the grouped series
	for _, partition := range df.partitions {
		seriesIndices[partition.index] = false
		old := df.series[partition.index]

		// TODO: null masks, null values are all mapped to the same group

		switch series := old.(type) {
		case SeriesBool:
			values := make([]bool, len(indeces))
			for i, group := range indeces {
				values[i] = series.data[group[0]]
			}
			result.series = append(result.series, SeriesBool{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
			})

		case SeriesInt32:
			values := make([]int32, len(indeces))
			for i, group := range indeces {
				values[i] = series.data[group[0]]
			}

			result.series = append(result.series, SeriesInt32{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
			})

		case SeriesInt64:
			values := make([]int64, len(indeces))
			for i, group := range indeces {
				values[i] = series.data[group[0]]
			}

			result.series = append(result.series, SeriesInt64{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
			})

		case SeriesFloat64:
			values := make([]float64, len(indeces))
			for i, group := range indeces {
				values[i] = series.data[group[0]]
			}

			result.series = append(result.series, SeriesFloat64{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
			})

		case SeriesString:
			values := make([]*string, len(indeces))
			for i, group := range indeces {
				values[i] = series.data[group[0]]
			}

			result.series = append(result.series, SeriesString{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
				pool:       series.pool,
			})
		}
	}

	// Get the indices of the ungrouped series
	ungroupedSeriesIndices := make([]int, 0)
	for index, isGrouped := range seriesIndices {
		if isGrouped {
			ungroupedSeriesIndices = append(ungroupedSeriesIndices, index)
		}
	}

	// sort the indices
	sort.Ints(ungroupedSeriesIndices)

	return result, &indeces, &ungroupedSeriesIndices
}

func (df BaseDataFrame) Join(how DataFrameJoinType, other DataFrame, on ...string) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.Join: cannot join a grouped dataframe")
		return df
	}

	if other.IsGrouped() {
		df.err = fmt.Errorf("BaseDataFrame.Join: cannot join with a grouped dataframe")
		return df
	}

	// CHECK: all the join columns must exist
	// CHECK: all the join columns must have the same type
	types := make([]typesys.BaseType, len(on))
	for _, name := range on {

		// Series A
		found := false
		for _, series := range df.series {
			if series.Name() == name {
				found = true

				// keep track of the types
				types = append(types, series.Type())
				break
			}
		}
		if !found {
			df.err = fmt.Errorf("BaseDataFrame.Join: column \"%s\" not found in left dataframe", name)
			return df
		}

		// Series B
		found = false
		for _, series := range other.(BaseDataFrame).series {
			if series.Name() == name {
				found = true

				// CHECK: the types must match
				if types[len(types)-1] != series.Type() {
					df.err = fmt.Errorf("BaseDataFrame.Join: columns \"%s\" have different types", name)
					return df
				}
				break
			}
		}
		if !found {
			df.err = fmt.Errorf("BaseDataFrame.Join: column \"%s\" not found in right dataframe", name)
			return df
		}
	}

	// CASE: on is empty -> use all columns with the same name
	if len(on) == 0 {
		for _, name := range df.Names() {
			if other.GetSeriesIndex(name) != -1 {
				on = append(on, name)
			}
		}
	}

	// CASE: on is still empty -> error
	if len(on) == 0 {
		df.err = fmt.Errorf("BaseDataFrame.Join: no columns to join on")
		return df
	}

	// CHECK: all columns in on must have the same type
	for _, name := range on {
		if df.Series(name).Type() != other.Series(name).Type() {
			df.err = fmt.Errorf("BaseDataFrame.Join: columns \"%s\" have different types", name)
			return df
		}
	}

	// CASE: the dataframes have different string pools
	if df.GetStringPool() != other.GetStringPool() {
		if df.NRows() < other.NRows() {
			df = df.SetStringPool(other.GetStringPool()).(BaseDataFrame)
		} else {
			other = other.SetStringPool(df.GetStringPool())
		}
	}

	// Group the dataframes by the join columns
	dfGrouped := df.GroupBy(on...).(BaseDataFrame)
	otherGrouped := other.GroupBy(on...).(BaseDataFrame)

	colsDiffA := make([]string, 0)
	colsDiffB := make([]string, 0)

	// Get the columns that are not in the join columns
	for _, name := range df.Names() {
		found := false
		for _, joinName := range on {
			if name == joinName {
				found = true
				break
			}
		}
		if !found {
			colsDiffA = append(colsDiffA, name)
		}
	}

	for _, name := range other.Names() {
		found := false
		for _, joinName := range on {
			if name == joinName {
				found = true
				break
			}
		}
		if !found {
			colsDiffB = append(colsDiffB, name)
		}
	}

	// Get the columns that are in both dataframes
	commonCols := make(map[string]bool)
	for _, name := range df.Names() {
		for _, otherName := range other.Names() {
			if name == otherName {
				commonCols[name] = true
				break
			}
		}
	}

	joined := NewBaseDataFrame()

	pA := dfGrouped.getPartitions()
	pB := otherGrouped.getPartitions()

	// Get the maps, keys and sort them
	mapA := pA[len(pA)-1].getMap()
	mapB := pB[len(pB)-1].getMap()

	keysA := make([]int64, 0, len(mapA))
	keysB := make([]int64, 0, len(mapB))

	for key := range mapA {
		keysA = append(keysA, key)
	}

	for key := range mapB {
		keysB = append(keysB, key)
	}

	sort.Slice(keysA, func(i, j int) bool { return keysA[i] < keysA[j] })
	sort.Slice(keysB, func(i, j int) bool { return keysB[i] < keysB[j] })

	// Find the intersection
	keysAOnly := make([]int64, 0, len(keysA))
	keysBOnly := make([]int64, 0, len(keysB))
	keysIntersection := make([]int64, 0, len(keysA))

	var i, j int = 0, 0
	for i < len(keysA) && j < len(keysB) {
		if keysA[i] < keysB[j] {
			keysAOnly = append(keysAOnly, keysA[i])
			i++
		} else if keysA[i] > keysB[j] {
			keysBOnly = append(keysBOnly, keysB[j])
			j++
		} else {
			keysIntersection = append(keysIntersection, keysA[i])
			i++
			j++
		}
	}

	for i < len(keysA) {
		keysAOnly = append(keysAOnly, keysA[i])
		i++
	}

	for j < len(keysB) {
		keysBOnly = append(keysBOnly, keysB[j])
		j++
	}

	switch how {
	case INNER_JOIN:
		// Get indices of the intersection
		indicesA := make([]int, 0, len(keysIntersection))
		indicesB := make([]int, 0, len(keysIntersection))

		for _, key := range keysIntersection {
			for _, indexA := range mapA[key] {
				for _, indexB := range mapB[key] {
					indicesA = append(indicesA, indexA)
					indicesB = append(indicesB, indexB)
				}
			}
		}

		// Join columns
		for i := range on {
			joined = joined.AddSeries(dfGrouped.Series(on[i]).Filter(indicesA))
		}

		// A columns
		var ser_ Series
		for _, name := range colsDiffA {
			ser_ = df.Series(name).Filter(indicesA)
			if commonCols[name] {
				ser_ = ser_.SetName(name + "_x")
			}
			joined = joined.AddSeries(ser_)
		}

		// B columns
		for _, name := range colsDiffB {
			ser_ = other.Series(name).Filter(indicesB)
			if commonCols[name] {
				ser_ = ser_.SetName(name + "_y")
			}
			joined = joined.AddSeries(ser_.Filter(indicesB))
		}

	case LEFT_JOIN:
		indicesA := make([]int, 0, len(keysA))
		indicesB := make([]int, 0, len(keysIntersection))

		for _, key := range keysAOnly {
			indicesA = append(indicesA, mapA[key]...)
		}

		for _, key := range keysIntersection {
			for _, indexA := range mapA[key] {
				for _, indexB := range mapB[key] {
					indicesA = append(indicesA, indexA)
					indicesB = append(indicesB, indexB)
				}
			}
		}

		// Join columns
		for i := range on {
			joined = joined.AddSeries(dfGrouped.Series(on[i]).Filter(indicesA))
		}

		// A columns
		var ser_ Series
		for _, name := range colsDiffA {
			ser_ = df.Series(name).Filter(indicesA)
			if commonCols[name] {
				ser_ = ser_.SetName(name + "_x")
			}
			joined = joined.AddSeries(ser_)
		}

		padBlen := len(indicesA) - len(indicesB)
		nullMask := make([]bool, padBlen)
		for i := range nullMask {
			nullMask[i] = true
		}

		// B columns
		for _, name := range colsDiffB {
			ser_ = other.Series(name).Filter(indicesB)
			switch ser_.Type() {
			case typesys.BoolType:
				ser_ = NewSeriesBool(ser_.Name(), true, false, make([]bool, padBlen), df.pool).
					SetNullMask(nullMask).(SeriesBool).
					Append(ser_)

			case typesys.Int32Type:
				ser_ = NewSeriesInt32(ser_.Name(), true, false, make([]int32, padBlen), df.pool).
					SetNullMask(nullMask).(SeriesInt32).
					Append(ser_)

			case typesys.Int64Type:
				ser_ = NewSeriesInt64(ser_.Name(), true, false, make([]int64, padBlen), df.pool).
					SetNullMask(nullMask).(SeriesInt64).
					Append(ser_)

			case typesys.Float64Type:
				ser_ = NewSeriesFloat64(ser_.Name(), true, false, make([]float64, padBlen), df.pool).
					SetNullMask(nullMask).(SeriesFloat64).
					Append(ser_)

			case typesys.StringType:
				ser_ = NewSeriesString(ser_.Name(), true, false, make([]string, padBlen), df.pool).
					SetNullMask(nullMask).(SeriesString).
					Append(ser_)
			}

			if commonCols[name] {
				ser_ = ser_.SetName(name + "_y")
			}
			joined = joined.AddSeries(ser_)
		}

	case RIGHT_JOIN:
		indicesA := make([]int, 0, len(keysIntersection))
		indicesB := make([]int, 0, len(keysB))

		for _, key := range keysIntersection {
			for _, indexA := range mapA[key] {
				for _, indexB := range mapB[key] {
					indicesA = append(indicesA, indexA)
					indicesB = append(indicesB, indexB)
				}
			}
		}

		for _, key := range keysBOnly {
			indicesB = append(indicesB, mapB[key]...)
		}

		// Join columns
		for i := range on {
			joined = joined.AddSeries(otherGrouped.Series(on[i]).Filter(indicesB))
		}

		padAlen := len(indicesB) - len(indicesA)
		nullMask := make([]bool, padAlen)
		for i := range nullMask {
			nullMask[i] = true
		}

		// A columns
		var ser_ Series
		for _, name := range colsDiffA {
			ser_ = df.Series(name).Filter(indicesA)
			switch ser_.Type() {
			case typesys.BoolType:
				ser_ = ser_.(SeriesBool).Append(NewSeriesBool(ser_.Name(), true, false, make([]bool, padAlen), df.pool).SetNullMask(nullMask))

			case typesys.Int32Type:
				ser_ = ser_.(SeriesInt32).Append(NewSeriesInt32(ser_.Name(), true, false, make([]int32, padAlen), df.pool).SetNullMask(nullMask))

			case typesys.Int64Type:
				ser_ = ser_.(SeriesInt64).Append(NewSeriesInt64(ser_.Name(), true, false, make([]int64, padAlen), df.pool).SetNullMask(nullMask))

			case typesys.Float64Type:
				ser_ = ser_.(SeriesFloat64).Append(NewSeriesFloat64(ser_.Name(), true, false, make([]float64, padAlen), df.pool).SetNullMask(nullMask))

			case typesys.StringType:
				ser_ = ser_.(SeriesString).Append(NewSeriesString(ser_.Name(), true, false, make([]string, padAlen), df.pool).SetNullMask(nullMask))
			}

			if commonCols[name] {
				ser_ = ser_.SetName(name + "_x")
			}
			joined = joined.AddSeries(ser_)
		}

		// B columns
		for _, name := range colsDiffB {
			ser_ = other.Series(name).Filter(indicesB)
			if commonCols[name] {
				ser_ = ser_.SetName(name + "_y")
			}
			joined = joined.AddSeries(ser_)
		}

	case OUTER_JOIN:
		indicesA := make([]int, 0, len(keysA))
		indicesB := make([]int, 0, len(keysB))

		padAlen := 0
		padBlen := 0

		for _, key := range keysAOnly {
			indicesA = append(indicesA, mapA[key]...)
			padBlen += len(mapA[key])
		}

		intersectionLen := 0
		for _, key := range keysIntersection {
			for _, indexA := range mapA[key] {
				for _, indexB := range mapB[key] {
					indicesA = append(indicesA, indexA)
					indicesB = append(indicesB, indexB)
					intersectionLen++
				}
			}
		}

		for _, key := range keysBOnly {
			indicesB = append(indicesB, mapB[key]...)
			padAlen += len(mapB[key])
		}

		// Join columns
		indicesBOnly := indicesB[intersectionLen:]
		for i := range on {
			joined = joined.AddSeries(
				dfGrouped.Series(on[i]).
					Filter(indicesA).Append(
					otherGrouped.Series(on[i]).
						Filter(indicesBOnly)))
		}

		nullMaskA := make([]bool, padAlen)
		for i := range nullMaskA {
			nullMaskA[i] = true
		}

		nullMaskB := make([]bool, padBlen)
		for i := range nullMaskB {
			nullMaskB[i] = true
		}

		// A columns
		var ser_ Series
		for _, name := range colsDiffA {
			ser_ = df.Series(name).Filter(indicesA)
			switch ser_.Type() {
			case typesys.BoolType:
				ser_ = ser_.(SeriesBool).Append(NewSeriesBool(ser_.Name(), true, false, make([]bool, padAlen), df.pool).SetNullMask(nullMaskA))

			case typesys.Int32Type:
				ser_ = ser_.(SeriesInt32).Append(NewSeriesInt32(ser_.Name(), true, false, make([]int32, padAlen), df.pool).SetNullMask(nullMaskA))

			case typesys.Int64Type:
				ser_ = ser_.(SeriesInt64).Append(NewSeriesInt64(ser_.Name(), true, false, make([]int64, padAlen), df.pool).SetNullMask(nullMaskA))

			case typesys.Float64Type:
				ser_ = ser_.(SeriesFloat64).Append(NewSeriesFloat64(ser_.Name(), true, false, make([]float64, padAlen), df.pool).SetNullMask(nullMaskA))

			case typesys.StringType:
				ser_ = ser_.(SeriesString).Append(NewSeriesString(ser_.Name(), true, false, make([]string, padAlen), df.pool).SetNullMask(nullMaskA))
			}

			if commonCols[name] {
				ser_ = ser_.SetName(name + "_x")
			}
			joined = joined.AddSeries(ser_)
		}

		// B columns
		for _, name := range colsDiffB {
			ser_ = other.Series(name).Filter(indicesB)
			switch ser_.Type() {
			case typesys.BoolType:
				ser_ = NewSeriesBool(ser_.Name(), true, false, make([]bool, padBlen), df.pool).
					SetNullMask(nullMaskB).(SeriesBool).
					Append(ser_)

			case typesys.Int32Type:
				ser_ = NewSeriesInt32(ser_.Name(), true, false, make([]int32, padBlen), df.pool).
					SetNullMask(nullMaskB).(SeriesInt32).
					Append(ser_)

			case typesys.Int64Type:
				ser_ = NewSeriesInt64(ser_.Name(), true, false, make([]int64, padBlen), df.pool).
					SetNullMask(nullMaskB).(SeriesInt64).
					Append(ser_)

			case typesys.Float64Type:
				ser_ = NewSeriesFloat64(ser_.Name(), true, false, make([]float64, padBlen), df.pool).
					SetNullMask(nullMaskB).(SeriesFloat64).
					Append(ser_)

			case typesys.StringType:
				ser_ = NewSeriesString(ser_.Name(), true, false, make([]string, padBlen), df.pool).
					SetNullMask(nullMaskB).(SeriesString).
					Append(ser_)
			}

			if commonCols[name] {
				ser_ = ser_.SetName(name + "_y")
			}
			joined = joined.AddSeries(ser_)
		}
	}

	return joined
}

func (df BaseDataFrame) Take(params ...int) DataFrame {
	if df.err != nil {
		return df
	}

	indeces, err := seriesTakePreprocess("BaseDataFrame", df.NRows(), params...)
	if err != nil {
		df.err = err
		return df
	}

	taken := NewBaseDataFrame()
	for _, series := range df.series {
		taken = taken.AddSeries(series.filterIntSlice(indeces, false))
	}

	return taken
}

func (df BaseDataFrame) Len() int {
	if df.err != nil || len(df.series) < 1 {
		return 0
	}

	return df.series[0].Len()
}

func (df BaseDataFrame) Less(i, j int) bool {
	for _, param := range df.sortParams {
		if !param.series.equal(i, j) {
			return (param.asc && param.series.Less(i, j)) || (!param.asc && param.series.Less(j, i))
		}
	}

	return false
}

func (df BaseDataFrame) Swap(i, j int) {
	for _, series := range df.series {
		series.Swap(i, j)
	}
}

func (df BaseDataFrame) OrderBy(params ...SortParam) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.OrderBy: cannot order grouped DataFrame")
		return df
	}

	// CHECK: params must have unique names and names must be valid
	paramNames := make(map[string]bool)
	for i, param := range params {
		if paramNames[param.name] {
			df.err = fmt.Errorf("BaseDataFrame.OrderBy: series names must be unique")
			return df
		}
		paramNames[param.name] = true

		if series := df.__series(param.name); series != nil {
			params[i].series = series
		} else {
			df.err = fmt.Errorf("BaseDataFrame.OrderBy: series \"%s\" not found", param.name)
			return df
		}
	}

	df.sortParams = params
	sort.Sort(df)
	df.sortParams = nil

	return df
}

////////////////////////			SUMMARY

func (df BaseDataFrame) Agg(aggregators ...aggregator) DataFrame {
	if df.err != nil {
		return df
	}

	// CHECK: aggregators must have unique names and names must be valid
	aggNames := make(map[string]bool)
	for _, agg := range aggregators {

		// CASE: aggregator count has a default name
		if agg.getAggregateType() != AGGREGATE_COUNT {
			if aggNames[agg.getSeriesName()] {
				df.err = fmt.Errorf("BaseDataFrame.Agg: aggregator names must be unique")
				return df
			}
			aggNames[agg.getSeriesName()] = true

			if df.__series(agg.getSeriesName()) == nil {
				df.err = fmt.Errorf("BaseDataFrame.Agg: series \"%s\" not found", agg.getSeriesName())
				return df
			}
		}
	}

	var result DataFrame
	if df.isGrouped {

		var indeces *[][]int
		result, indeces, _ = df.groupHelper()

		if df.NRows() < MINIMUM_PARALLEL_SIZE_2 {
			for _, agg := range aggregators {
				series := df.__series(agg.getSeriesName())

				switch agg.getAggregateType() {
				case AGGREGATE_COUNT:
					counts := make([]int64, len(*indeces))
					for i, group := range *indeces {
						counts[i] = int64(len(group))
					}
					result = result.AddSeries(NewSeriesInt64(agg.getSeriesName(), false, false, counts, df.pool))

				case AGGREGATE_SUM:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_sum_grouped__(series, *indeces), df.pool))

				case AGGREGATE_MIN:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_min_grouped__(series, *indeces), df.pool))

				case AGGREGATE_MAX:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_max_grouped__(series, *indeces), df.pool))

				case AGGREGATE_MEAN:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_mean_grouped__(series, *indeces), df.pool))

				case AGGREGATE_MEDIAN:
					// TODO: implement

				case AGGREGATE_STD:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_std_grouped__(series, *indeces), df.pool))
				}
			}
		} else {

			var wg sync.WaitGroup
			wg.Add(THREADS_NUMBER)

			buffer := make(chan __stats_thread_data)
			for i := 0; i < THREADS_NUMBER; i++ {
				go __stats_worker(&wg, buffer)
			}

			for _, agg := range aggregators {
				series := df.__series(agg.getSeriesName())

				resultData := make([]float64, len(*indeces))
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, resultData, df.pool))
				for gi, group := range *indeces {
					buffer <- __stats_thread_data{
						op:      agg.getAggregateType(),
						gi:      gi,
						indeces: group,
						series:  series,
						res:     resultData,
					}
				}
			}

			close(buffer)
			wg.Wait()
		}

	} else {

		result = NewBaseDataFrame()

		for _, agg := range aggregators {
			series := df.__series(agg.getSeriesName())

			switch agg.getAggregateType() {
			case AGGREGATE_COUNT:
				result = result.AddSeries(NewSeriesInt64(agg.getSeriesName(), false, false, []int64{int64(df.NRows())}, df.pool))

			case AGGREGATE_SUM:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_sum__(series)}, df.pool))

			case AGGREGATE_MIN:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_min__(series)}, df.pool))

			case AGGREGATE_MAX:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_max__(series)}, df.pool))

			case AGGREGATE_MEAN:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_mean__(series)}, df.pool))

			case AGGREGATE_MEDIAN:
				// TODO: implement

			case AGGREGATE_STD:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_std__(series)}, df.pool))
			}
		}
	}

	return result
}

////////////////////////			PRINTING

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func (df BaseDataFrame) Describe() string {
	return ""
}

func (df BaseDataFrame) Records(header bool) [][]string {
	var out [][]string
	if header {
		out = make([][]string, df.NRows()+1)
	} else {
		out = make([][]string, df.NRows())
	}

	h := 0
	if header {
		out[0] = make([]string, df.NCols())
		for j := 0; j < df.NCols(); j++ {
			out[0][j] = df.series[j].Name()
		}

		h = 1
	}

	for i := 0 + h; i < df.NRows()+h; i++ {
		out[i] = make([]string, df.NCols())
		for j := 0; j < df.NCols(); j++ {
			out[i][j] = df.series[j].GetString(i - h)
		}
	}

	return out
}

func (df BaseDataFrame) PrettyPrint(nrowsParam ...int) DataFrame {
	if df.err != nil {
		fmt.Println(df.err)
		return df
	}

	if df.isGrouped {
		fmt.Printf("    GROUPED BY")
		for i, p := range df.partitions {
			if i != len(df.partitions)-1 {
				fmt.Printf(" %s,", p.name)
			} else {
				fmt.Printf(" %s", p.name)
			}
		}
		fmt.Printf("\n\n")
	}

	colSize := 10
	actualColSize := colSize + 3
	fmtString := fmt.Sprintf("| %%%ds ", colSize)

	// header
	fmt.Printf("    ")
	for i := 0; i < len(df.series)*actualColSize; i++ {
		if i%actualColSize == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println("+")

	// column names
	// check if there are any column names
	colNames := false
	for _, c := range df.series {
		if c.Name() != "" {
			colNames = true
			break
		}
	}

	// only print column names if there are any
	if colNames {
		fmt.Printf("    ")
		for _, c := range df.series {
			fmt.Printf(fmtString, truncate(c.Name(), colSize))
		}
		fmt.Println("|")

		// separator
		fmt.Printf("    ")
		for i := 0; i < len(df.series)*actualColSize; i++ {
			if i%actualColSize == 0 {
				fmt.Print("+")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println("+")
	}

	// column types
	fmt.Printf("    ")
	for _, c := range df.series {
		fmt.Printf(fmtString, truncate(c.Type().ToString(), colSize))
	}
	fmt.Println("|")

	// separator
	fmt.Printf("    ")
	for i := 0; i < len(df.series)*actualColSize; i++ {
		if i%actualColSize == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println("+")

	var nrows int
	if len(nrowsParam) == 0 {
		if df.NRows() < 20 {
			nrows = df.NRows()
		} else {
			nrows = 10
		}
	} else {
		nrows = nrowsParam[0]
	}

	// data
	if nrows >= 0 {
		nrows = int(math.Min(float64(nrows), float64(df.NRows())))
	} else {
		nrows = df.NRows()
	}

	for i := 0; i < nrows; i++ {
		fmt.Printf("    ")
		for _, c := range df.series {
			fmt.Printf(fmtString, truncate(c.GetString(i), colSize))
		}
		fmt.Println("|")
	}

	// separator
	fmt.Printf("    ")
	for i := 0; i < len(df.series)*actualColSize; i++ {
		if i%actualColSize == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println("+")

	return df
}

////////////////////////			IO

func (df BaseDataFrame) FromCSV() *CsvReader {
	return NewCsvReader(df.pool)
}

func (df BaseDataFrame) ToCSV() *CsvWriter {
	return NewCsvWriter().SetDataFrame(df)
}
