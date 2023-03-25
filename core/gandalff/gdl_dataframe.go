package gandalff

import (
	"fmt"
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

func (df *GDLDataFrame) Filter() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	filtered := NewGDLDataFrame()

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

		grouped := NewGDLDataFrame()
		grouped.isGrouped = true
		grouped.partitions = make([]GDLDataFramePartitionEntry, len(by))

		partitionsIndex := 0

		for i, series := range df.series {
			for _, name := range by {
				if series.Name() == name {

					// First partition: group the series
					if partitionsIndex == 0 {
						grouped.partitions[partitionsIndex] = GDLDataFramePartitionEntry{
							index:     i,
							name:      name,
							partition: series.Group(),
						}
					} else

					// Subsequent partitions: sub-group the series
					{
						grouped.partitions[partitionsIndex] = GDLDataFramePartitionEntry{
							index:     i,
							name:      name,
							partition: series.SubGroup(grouped.partitions[partitionsIndex-1].partition),
						}
					}

					// Series found, increment the partitions index
					partitionsIndex++
					break
				}
			}

			// Add the series to the grouped dataframe
			grouped.AddSeries(series)
		}

		return grouped
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

///////////////////////////////		SUMMARY		/////////////////////////////////////////

func (df *GDLDataFrame) Count() *GDLDataFrame {
	if df.err != nil {
		return df
	}

	result := NewGDLDataFrame()

	if df.isGrouped {

		for _, partition := range df.partitions {
			series := df.SeriesAt(partition.index)
			result.AddSeries(series)
		}

		// Add the count series
		counts := make([]int, len(df.partitions[len(df.partitions)-1].partition.GetNonNullGroups()))
		for i, group := range df.partitions[len(df.partitions)-1].partition.GetNonNullGroups() {
			counts[i] = len(group)
		}

		if df.partitions[len(df.partitions)-1].partition.GetNullGroup() != nil {
			counts = append(counts, len(df.partitions[len(df.partitions)-1].partition.GetNullGroup()))
		}

		result.AddSeries(NewGDLSeriesInt32("count", false, false, counts))

	} else {
		result.AddSeries(NewGDLSeriesInt32("count", false, false, []int{df.NRows()}))
	}

	return result
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
		fmt.Println("GROUPED")
	} else {
		fmt.Println("NOT GROUPED")
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
			switch c.Type() {
			case typesys.BoolType:
				fmt.Printf(fmtString, truncate(fmt.Sprintf("%t", c.Data().([]bool)[i]), colSize))
			case typesys.Int32Type:
				fmt.Printf(fmtString, truncate(fmt.Sprintf("%d", c.Data().([]int)[i]), colSize))
			case typesys.Float64Type:
				fmt.Printf(fmtString, truncate(fmt.Sprintf("%f", c.Data().([]float64)[i]), colSize))
			case typesys.StringType:
				fmt.Printf(fmtString, truncate(c.Data().([]string)[i], colSize))
			}
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
