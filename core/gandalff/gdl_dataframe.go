package gandalff

import (
	"fmt"
	"sort"
	"typesys"
)

type GDLDataFramePartitionEntry struct {
	index     int
	name      string
	partition GDLSeriesPartition
}

type GDLDataFrame struct {
	isGrouped  bool
	err        error
	series     []GDLSeries
	pool       *StringPool
	partitions []GDLDataFramePartitionEntry
}

func NewGDLDataFrame() *GDLDataFrame {
	return &GDLDataFrame{
		series: make([]GDLSeries, 0),
		pool:   NewStringPool(),
	}
}

func (df *GDLDataFrame) IsErrored() bool {
	return df.err != nil
}

func (df *GDLDataFrame) IsGrouped() bool {
	return df.isGrouped
}

func (df *GDLDataFrame) GetError() error {
	return df.err
}

func (df *GDLDataFrame) GetPool() *StringPool {
	return df.pool
}

func (df *GDLDataFrame) GetSeriesIndex(name string) int {
	for i, series := range df.series {
		if series.Name() == name {
			return i
		}
	}
	return -1
}

func (df *GDLDataFrame) AddSeries(series GDLSeries) {
	df.series = append(df.series, series)
}

func (df *GDLDataFrame) AddSeriesFromBools(name string, isNullable bool, data []bool) {
	series := NewGDLSeriesBool(name, isNullable, data)
	df.AddSeries(series)
}

func (df *GDLDataFrame) AddSeriesFromInts(name string, isNullable bool, makeCopy bool, data []int) {
	series := NewGDLSeriesInt32(name, isNullable, makeCopy, data)
	df.AddSeries(series)
}

func (df *GDLDataFrame) AddSeriesFromFloats(name string, isNullable bool, makeCopy bool, data []float64) {
	series := NewGDLSeriesFloat64(name, isNullable, makeCopy, data)
	df.AddSeries(series)
}

func (df *GDLDataFrame) AddSeriesFromStrings(name string, isNullable bool, data []string) {
	series := NewGDLSeriesString(name, isNullable, data, df.pool)
	df.AddSeries(series)
}

// Names returns the names of the series in the dataframe.
func (df *GDLDataFrame) Names() []string {
	names := make([]string, len(df.series))
	for i, series := range df.series {
		names[i] = series.Name()
	}
	return names
}

// Types returns the types of the series in the dataframe.
func (df *GDLDataFrame) Types() []typesys.BaseType {
	types := make([]typesys.BaseType, len(df.series))
	for i, series := range df.series {
		types[i] = series.Type()
	}
	return types
}

func (df *GDLDataFrame) NCols() int {
	return len(df.series)
}

func (df *GDLDataFrame) NRows() int {
	if len(df.series) == 0 {
		return 0
	}
	return df.series[0].Len()
}

// Returns the series with the given name.
func (df *GDLDataFrame) Series(name string) GDLSeries {
	for _, series := range df.series {
		if series.Name() == name {
			return series
		}
	}
	return nil
}

// Returns the series at the given index.
func (df *GDLDataFrame) SeriesAt(index int) GDLSeries {
	if index < 0 || index >= len(df.series) {
		return nil
	}
	return df.series[index]
}

func (df *GDLDataFrame) Select(names ...string) *GDLDataFrame {
	if df.err != nil {
		return df
	}

	selected := NewGDLDataFrame()
	for _, name := range names {
		series := df.Series(name)
		if series != nil {
			selected.AddSeries(series)
		} else {
			selected.err = fmt.Errorf("GDLDataFrame.Select: series \"%s\" not found", name)
			return selected
		}
	}

	return selected
}

func (df *GDLDataFrame) InPlaceSelect(names ...string) error {
	if df.err != nil {
		return df.err
	}

	selected := df.Select(names...)
	if selected.err != nil {
		return selected.err
	}

	df.series = selected.series
	return nil
}

func (df *GDLDataFrame) SelectAt(indices ...int) *GDLDataFrame {
	if df.err != nil {
		return df
	}

	selected := NewGDLDataFrame()
	for _, index := range indices {
		series := df.SeriesAt(index)
		if series != nil {
			selected.AddSeries(series)
		} else {
			selected.err = fmt.Errorf("GDLDataFrame.SelectAt: series at index %d not found", index)
			return selected
		}
	}

	return selected
}

func (df *GDLDataFrame) InPlaceSelectAt(indices ...int) error {
	if df.err != nil {
		return df.err
	}

	selected := df.SelectAt(indices...)
	if selected.err != nil {
		return selected.err
	}

	df.series = selected.series
	return nil
}

func (df *GDLDataFrame) Filter(mask GDLSeriesBool) *GDLDataFrame {
	if df.err != nil {
		return df
	}

	if mask.Len() != df.NRows() {
		df.err = fmt.Errorf("GDLDataFrame.Filter: mask length (%d) does not match dataframe length (%d)", mask.Len(), df.NRows())
		return df
	}

	filtered := NewGDLDataFrame()
	m := mask.Data().([]bool)
	for _, series := range df.series {
		filtered.AddSeries(series.FilterByMask(m))
	}

	return filtered
}

func (df *GDLDataFrame) GroupBy(by ...string) *GDLDataFrame {
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
				df.err = fmt.Errorf("GDLDataFrame.GroupBy: column \"%s\" not found", name)
				return df
			}
		}

		df.isGrouped = true
		df.partitions = make([]GDLDataFramePartitionEntry, len(by))

		for partitionsIndex, name := range by {

			i := df.GetSeriesIndex(name)
			series := df.series[i]

			// First partition: group the series
			if partitionsIndex == 0 {
				df.partitions[partitionsIndex] = GDLDataFramePartitionEntry{
					index:     i,
					name:      name,
					partition: series.Group().GetPartition(),
				}
			} else

			// Subsequent partitions: sub-group the series
			{
				df.partitions[partitionsIndex] = GDLDataFramePartitionEntry{
					index:     i,
					name:      name,
					partition: series.SubGroup(df.partitions[partitionsIndex-1].partition).GetPartition(),
				}
			}
		}

		return df
	}
}

func (df *GDLDataFrame) Ungroup() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	df.isGrouped = false
	df.partitions = nil
	return df
}

func (df *GDLDataFrame) groupHelper() (*GDLDataFrame, *[][]int, *[]int) {

	// Keep track of which series are not grouped
	seriesIndices := make(map[int]bool)
	for i := 0; i < df.NCols(); i++ {
		seriesIndices[i] = true
	}

	result := NewGDLDataFrame()

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

///////////////////////////////		SUMMARY		/////////////////////////////////////////

func (df *GDLDataFrame) Count(name string) *GDLDataFrame {
	if df.err != nil {
		return df
	}

	var result *GDLDataFrame
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
		result := NewGDLDataFrame()
		result.AddSeries(NewGDLSeriesInt32(name, false, false, []int{df.NRows()}))
	}

	return result
}

func (df *GDLDataFrame) Sum() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {

		var result *GDLDataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_sum_grouped__(series, *indeces)))
		}

		return result

	} else {
		result := NewGDLDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_sum__(series)}))
		}
		return result
	}
}

func (df *GDLDataFrame) Min() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result *GDLDataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_min_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewGDLDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_min__(series)}))
		}
		return result
	}
}

func (df *GDLDataFrame) Max() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result *GDLDataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_max_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewGDLDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_max__(series)}))
		}
		return result
	}
}

func (df *GDLDataFrame) Mean() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result *GDLDataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_mean_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewGDLDataFrame()
		for _, series := range df.series {
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, []float64{__gdl_mean__(series)}))
		}
		return result
	}
}

func (df *GDLDataFrame) Std() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	if df.isGrouped {
		var result *GDLDataFrame
		var indeces *[][]int
		var ungroupedSeriesIndices *[]int

		result, indeces, ungroupedSeriesIndices = df.groupHelper()

		for _, index := range *ungroupedSeriesIndices {
			series := df.series[index]
			result.AddSeries(NewGDLSeriesFloat64(series.Name(), false, false, __gdl_std_grouped__(series, *indeces)))
		}

		return result
	} else {
		result := NewGDLDataFrame()
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

func (df *GDLDataFrame) PrettyPrint() {
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
