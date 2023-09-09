package preludiocore

import (
	"gandalff"
	"typesys"
)

func seriesToColumnar(fullOutput bool, outputSnippetLength int, series gandalff.Series) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = series.Name()
	col.Type = series.Type().ToString()
	col.ActualLength = series.Len()

	if !fullOutput && series.Len() > outputSnippetLength {
		series = series.Take(outputSnippetLength)
	}
	col.Data = series.DataAsString()
	col.Nulls = series.GetNullMask()

	return col
}

func dataFrameToColumnar(fullOutput bool, outputSnippetLength int, df *gandalff.DataFrame) []typesys.Columnar {
	columns := make([]typesys.Columnar, (*df).NCols())
	for i, name := range (*df).Names() {
		columns[i] = seriesToColumnar(fullOutput, outputSnippetLength, (*df).Series(name))
	}
	return columns
}
