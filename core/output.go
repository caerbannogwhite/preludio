package preludiocore

import (
	"gandalff"
	"preludiometa"
)

func seriesToColumnar(fullOutput bool, outputSnippetLength int, name string, series gandalff.Series) preludiometa.Columnar {
	col := preludiometa.Columnar{}
	col.Name = name
	col.Type = series.Type().ToString()
	col.ActualLength = series.Len()

	if !fullOutput && series.Len() > outputSnippetLength {
		series = series.Take(outputSnippetLength)
	}
	col.Data = series.DataAsString()
	col.Nulls = series.GetNullMask()

	return col
}

func dataFrameToColumnar(fullOutput bool, outputSnippetLength int, df *gandalff.DataFrame) []preludiometa.Columnar {
	columns := make([]preludiometa.Columnar, (*df).NCols())
	for i, name := range (*df).Names() {
		columns[i] = seriesToColumnar(fullOutput, outputSnippetLength, name, (*df).Series(name))
	}
	return columns
}
