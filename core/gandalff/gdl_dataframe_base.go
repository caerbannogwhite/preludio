package gandalff

import (
	"fmt"
	"sort"
	"typesys"
)

type BaseDataFramePartitionEntry struct {
	index     int
	name      string
	partition GDLSeriesPartition
}

type BaseDataFrame struct {
	isGrouped  bool
	err        error
	series     []GDLSeries
	pool       *StringPool
	partitions []BaseDataFramePartitionEntry
}

func NewBaseDataFrame() DataFrame {
	return &BaseDataFrame{
		series: make([]GDLSeries, 0),
		pool:   NewStringPool(),
	}
}

///////////////////////			BASIC ACCESSORS     	/////////////////////////////////

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

func (df BaseDataFrame) AddSeries(series GDLSeries) DataFrame {
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

func (df BaseDataFrame) AddSeriesFromBools(name string, isNullable bool, data []bool) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromBools: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromBools: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	series := NewGDLSeriesBool(name, isNullable, data)
	df.AddSeries(series)
	return df
}

func (df BaseDataFrame) AddSeriesFromInts(name string, isNullable bool, makeCopy bool, data []int) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromInts: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromInts: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	series := NewGDLSeriesInt32(name, isNullable, makeCopy, data)
	df.AddSeries(series)
	return df
}

func (df BaseDataFrame) AddSeriesFromFloats(name string, isNullable bool, makeCopy bool, data []float64) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromFloats: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromFloats: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	series := NewGDLSeriesFloat64(name, isNullable, makeCopy, data)
	df.AddSeries(series)
	return df
}

func (df BaseDataFrame) AddSeriesFromStrings(name string, isNullable bool, data []string) DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromStrings: cannot add series to a grouped dataframe")
		return df
	}

	if df.NCols() > 0 && len(data) != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.AddSeriesFromStrings: series length (%d) does not match dataframe length (%d)", len(data), df.NRows())
		return df
	}

	series := NewGDLSeriesString(name, isNullable, data, df.pool)
	df.AddSeries(series)
	return df
}

// Returns the series with the given name.
func (df BaseDataFrame) Series(name string) GDLSeries {
	for _, series := range df.series {
		if series.Name() == name {
			return series
		}
	}
	return GDLSeriesError{msg: fmt.Sprintf("BaseDataFrame.Series: series \"%s\" not found", name)}
}

// Returns the series at the given index.
func (df BaseDataFrame) SeriesAt(index int) GDLSeries {
	if index < 0 || index >= len(df.series) {
		return nil
	}
	return df.series[index]
}

func (df BaseDataFrame) Select(names ...string) DataFrame {
	if df.err != nil {
		return df
	}

	selected := NewBaseDataFrame()
	for _, name := range names {
		series := df.Series(name)
		if series != nil {
			selected.AddSeries(series)
		} else {
			return BaseDataFrame{
				err: fmt.Errorf("BaseDataFrame.Select: series \"%s\" not found", name),
			}
		}
	}

	return selected
}

func (df BaseDataFrame) SelectAt(indices ...int) DataFrame {
	if df.err != nil {
		return df
	}

	selected := NewBaseDataFrame()
	for _, index := range indices {
		series := df.SeriesAt(index)
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

func (df BaseDataFrame) Filter(mask GDLSeriesBool) DataFrame {
	if df.err != nil {
		return df
	}

	if mask.Len() != df.NRows() {
		df.err = fmt.Errorf("BaseDataFrame.Filter: mask length (%d) does not match dataframe length (%d)", mask.Len(), df.NRows())
		return df
	}

	filtered := NewBaseDataFrame()
	for _, series := range df.series {
		filtered.AddSeries(series.Filter(mask))
	}

	return filtered
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

func (df BaseDataFrame) groupHelper() (DataFrame, *[][]int, *[]int) {

	// Keep track of which series are not grouped
	seriesIndices := make(map[int]bool)
	for i := 0; i < df.NCols(); i++ {
		seriesIndices[i] = true
	}

	result := NewBaseDataFrame()

	// The last partition tells us how many groups there are
	// and how many rows are in each group
	indeces := df.partitions[len(df.partitions)-1].partition.GetIndices()

	// Keep only the grouped series
	for _, partition := range df.partitions {
		seriesIndices[partition.index] = false
		old := df.SeriesAt(partition.index)

		switch old.(type) {
		case GDLSeriesBool:
			values := make([]bool, len(indeces))
			for i, group := range indeces {
				values[i] = old.Get(group[0]).(bool)
			}
			result.AddSeries(NewGDLSeriesBool(old.Name(), old.IsNullable(), values))

		case GDLSeriesInt32:
			values := make([]int, len(indeces))
			for i, group := range indeces {
				values[i] = old.Get(group[0]).(int)
			}
			result.AddSeries(NewGDLSeriesInt32(old.Name(), old.IsNullable(), false, values))

		case GDLSeriesFloat64:
			values := make([]float64, len(indeces))
			for i, group := range indeces {
				values[i] = old.Get(group[0]).(float64)
			}
			result.AddSeries(NewGDLSeriesFloat64(old.Name(), old.IsNullable(), false, values))

		case GDLSeriesString:
			values := make([]string, len(indeces))
			for i, group := range indeces {
				values[i] = old.Get(group[0]).(string)
			}
			result.AddSeries(NewGDLSeriesString(old.Name(), old.IsNullable(), values, df.pool))
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

func (df BaseDataFrame) Take(start, end, step int) DataFrame {
	if df.err != nil {
		return df
	}

	taken := NewBaseDataFrame()
	for _, series := range df.series {
		taken.AddSeries(series.Take(start, end, step))
	}

	return taken
}

///////////////////////////////		SUMMARY		/////////////////////////////////////////

func (df BaseDataFrame) Count(name string) DataFrame {
	if df.err != nil {
		return df
	}

	var result DataFrame
	if df.isGrouped {

		var indeces *[][]int
		result, indeces, _ = df.groupHelper()

		// Add the count series
		counts := make([]int, len(*indeces))
		for i, group := range *indeces {
			counts[i] = len(group)
		}

		result.AddSeries(NewGDLSeriesInt32(name, false, false, counts))

	} else {
		result := NewBaseDataFrame()
		result.AddSeries(NewGDLSeriesInt32(name, false, false, []int{df.NRows()}))
	}

	return result
}

func (df BaseDataFrame) Sum() DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {

		var result DataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_sum_grouped__(series, *indeces)))
		}

		return result

	} else {
		result := NewBaseDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_sum__(series)}))
		}
		return result
	}
}

func (df BaseDataFrame) Min() DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result DataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_min_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewBaseDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_min__(series)}))
		}
		return result
	}
}

func (df BaseDataFrame) Max() DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result DataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_max_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewBaseDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_max__(series)}))
		}
		return result
	}
}

func (df BaseDataFrame) Mean() DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result DataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_mean_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewBaseDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_mean__(series)}))
		}
		return result
	}
}

func (df BaseDataFrame) Std() DataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result DataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_std_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewBaseDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_std__(series)}))
		}
		return result
	}
}

///////////////////////////////		PRINTING	/////////////////////////////////////////

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

///////////////////////////////		IO		/////////////////////////////////////////////

func (df BaseDataFrame) FromCSV() *GDLCsvReader {
	return NewGDLCsvReader()
}
