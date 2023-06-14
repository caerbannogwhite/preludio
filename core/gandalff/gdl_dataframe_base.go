package gandalff

import (
	"fmt"
	"sort"
	"sync"
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

func (df BaseDataFrame) GetPool() *StringPool {
	return df.pool
}

func (df BaseDataFrame) GetSeriesIndex(name string) int {
	for i, series := range df.series {
		if series.Name() == name {
			return i
		}
	}
	return -1
}

func (df BaseDataFrame) AddSeries(series Series) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeries: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && series.Len() != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeries: series length (%d) does not match dataframe length (%d)", series.Len(), df.NRows())
		return df
	}

	df.series = append(df.series, series)
	return df
}

func (df BaseDataFrame) AddSeriesFromBool(name string, isNullable bool, data []bool) DataFrame {
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

	return df.AddSeries(NewSeriesBool(name, isNullable, data))
}

func (df BaseDataFrame) AddSeriesFromInt32(name string, isNullable bool, makeCopy bool, data []int) DataFrame {
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

	return df.AddSeries(NewSeriesInt32(name, isNullable, makeCopy, data))
}

func (df BaseDataFrame) AddSeriesFromInt64(name string, isNullable bool, makeCopy bool, data []int64) DataFrame {
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

	return df.AddSeries(NewSeriesInt64(name, isNullable, makeCopy, data))
}

func (df BaseDataFrame) AddSeriesFromFloat64(name string, isNullable bool, makeCopy bool, data []float64) DataFrame {
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

	return df.AddSeries(NewSeriesFloat64(name, isNullable, makeCopy, data))
}

func (df BaseDataFrame) AddSeriesFromString(name string, isNullable bool, data []string) DataFrame {
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

	return df.AddSeries(NewSeriesString(name, isNullable, data, df.pool))
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

// Returns the series at the given index.
// For internal use only: returns nil if the series is not found.
func (df BaseDataFrame) __seriesAt(index int) Series {
	if index < 0 || index >= len(df.series) {
		return nil
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
		series := df.__seriesAt(index)
		if series != nil {
			selected.AddSeries(series)
		} else {
			return BaseDataFrame{
				err: fmt.Errorf("BaseDataFrame.SelectAt: series at index %d not found", index),
			}
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
					partition: series.Group().GetPartition(),
				}
			} else

			// Subsequent partitions: sub-group the series
			{
				df.partitions[partitionsIndex] = BaseDataFramePartitionEntry{
					index:     i,
					name:      name,
					partition: series.SubGroup(df.partitions[partitionsIndex-1].partition).GetPartition(),
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

func (df BaseDataFrame) GetPartitions() []SeriesPartition {
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
	indeces := make([][]int, 0, df.partitions[len(df.partitions)-1].partition.GetSize())
	for _, group := range df.partitions[len(df.partitions)-1].partition.GetMap() {
		indeces = append(indeces, group)
	}

	// Keep only the grouped series
	for _, partition := range df.partitions {
		seriesIndices[partition.index] = false
		old := df.__seriesAt(partition.index)

		// TODO: null masks, null values are all mapped to the same group

		switch series := old.(type) {
		case SeriesBool:
			values := make([]bool, len(indeces))
			for i, group := range indeces {
				values[i] = old.Get(group[0]).(bool)
			}
			result.series = append(result.series, NewSeriesBool(old.Name(), old.IsNullable(), values))

		case SeriesInt32:
			values := make([]int, len(indeces))
			data := series.__getDataPtr()
			for i, group := range indeces {
				values[i] = (*data)[group[0]]
			}

			result.series = append(result.series, SeriesInt32{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
			})

		case SeriesInt64:
			values := make([]int64, len(indeces))
			data := series.__getDataPtr()
			for i, group := range indeces {
				values[i] = (*data)[group[0]]
			}

			result.series = append(result.series, SeriesInt64{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
			})

		case SeriesFloat64:
			values := make([]float64, len(indeces))
			data := series.__getDataPtr()
			for i, group := range indeces {
				values[i] = (*data)[group[0]]
			}

			result.series = append(result.series, SeriesFloat64{
				name:       series.name,
				isNullable: series.isNullable,
				nullMask:   __binVecInit(len(indeces)),
				data:       values,
			})

		case SeriesString:
			values := make([]*string, len(indeces))
			data := series.__getDataPtr()
			for i, group := range indeces {
				values[i] = (*data)[group[0]]
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

	joined := NewBaseDataFrame()

	switch how {
	case INNER_JOIN:
		// TODO: implement
		pA := dfGrouped.GetPartitions()
		pB := otherGrouped.GetPartitions()

		// Get the maps, keys and sort them
		mapA := pA[len(pA)-1].GetMap()
		mapB := pB[len(pB)-1].GetMap()

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
		keysIntersection := make([]int64, 0, len(keysA))
		for i, j := 0, 0; i < len(keysA) && j < len(keysB); {
			if keysA[i] < keysB[j] {
				i++
			} else if keysA[i] > keysB[j] {
				j++
			} else {
				keysIntersection = append(keysIntersection, keysA[i])
				i++
				j++
			}
		}

		// Get indices of the intersection
		indices := make([]int, 0, len(keysIntersection))
		for _, key := range keysIntersection {
			indices = append(indices, mapA[key][0])
		}

		// Join columns
		for i := range on {
			joined = joined.AddSeries(dfGrouped.Series(on[i]).FilterByIndeces(indices))
		}

		// A columns
		for _, name := range colsDiffA {
			joined = joined.AddSeries(df.Series(name).FilterByIndeces(indices))
		}

		// B columns
		for _, name := range colsDiffB {
			joined = joined.AddSeries(other.Series(name).FilterByIndeces(indices))
		}

	case LEFT_JOIN:
		// TODO: implement

	case RIGHT_JOIN:
		// TODO: implement

	case OUTER_JOIN:
		// TODO: implement
	}

	return joined
}

func (df BaseDataFrame) Take(start, end, step int) DataFrame {
	if df.err != nil {
		return df
	}

	taken := NewBaseDataFrame()
	for _, series := range df.series {
		taken = taken.AddSeries(series.Take(start, end, step))
	}

	return taken
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
					counts := make([]int, len(*indeces))
					for i, group := range *indeces {
						counts[i] = len(group)
					}
					result = result.AddSeries(NewSeriesInt32(agg.getSeriesName(), false, false, counts))

				case AGGREGATE_SUM:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_sum_grouped__(series, *indeces)))

				case AGGREGATE_MIN:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_min_grouped__(series, *indeces)))

				case AGGREGATE_MAX:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_max_grouped__(series, *indeces)))

				case AGGREGATE_MEAN:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_mean_grouped__(series, *indeces)))

				case AGGREGATE_MEDIAN:
					// TODO: implement

				case AGGREGATE_STD:
					result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, __gdl_std_grouped__(series, *indeces)))
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
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, resultData))
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
				result = result.AddSeries(NewSeriesInt32(agg.getSeriesName(), false, false, []int{df.NRows()}))

			case AGGREGATE_SUM:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_sum__(series)}))

			case AGGREGATE_MIN:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_min__(series)}))

			case AGGREGATE_MAX:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_max__(series)}))

			case AGGREGATE_MEAN:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_mean__(series)}))

			case AGGREGATE_MEDIAN:
				// TODO: implement

			case AGGREGATE_STD:
				result = result.AddSeries(NewSeriesFloat64(series.Name(), false, false, []float64{__gdl_std__(series)}))
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

func (df BaseDataFrame) PrettyPrint() {
	if df.err != nil {
		fmt.Println(df.err)
		return
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

	// data
	for i := 0; i < df.NRows(); i++ {
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
}

////////////////////////			IO

func (df BaseDataFrame) FromCSV() *GDLCsvReader {
	return NewGDLCsvReader()
}
