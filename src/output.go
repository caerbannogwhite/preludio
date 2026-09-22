package preludiocore

import (
	"github.com/caerbannogwhite/enchanter/dataframe"
	"github.com/caerbannogwhite/enchanter/meta"
	"github.com/caerbannogwhite/enchanter/series"
)

func seriesToColumnar(fullOutput bool, outputSnippetLength int, name string, series series.Series) meta.Columnar {
	col := meta.Columnar{}
	col.Name = name
	col.Type = series.Type().String()
	col.ActualLength = series.Len()

	if !fullOutput && series.Len() > outputSnippetLength {
		series = series.Slice(0, outputSnippetLength)
	}
	col.Data = series.DataAsString()
	col.Nulls = series.NullMask()

	return col
}

func dataFrameToColumnar(fullOutput bool, outputSnippetLength int, df *dataframe.DataFrame) []meta.Columnar {
	columns := make([]meta.Columnar, (*df).NCols())
	for i, name := range (*df).Names() {
		columns[i] = seriesToColumnar(fullOutput, outputSnippetLength, name, (*df).Col(name))
	}
	return columns
}
