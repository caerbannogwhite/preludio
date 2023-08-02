package preludiocore

import (
	"gandalff"
	"strconv"
	"typesys"
)

func NewColumnarBool(name string, fullOutput bool, outputSnippetLength int, data []bool) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = name
	col.Type = typesys.BoolType.ToString()
	col.ActualLength = len(data)
	if !fullOutput && len(data) > outputSnippetLength {
		col.Data = make([]string, outputSnippetLength)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == outputSnippetLength-1 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = strconv.FormatBool(v)
	}
	return col
}

func NewColumnarInt64(name string, fullOutput bool, outputSnippetLength int, data []int64) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = name
	col.Type = typesys.Int64Type.ToString()
	col.ActualLength = len(data)
	if !fullOutput && len(data) > outputSnippetLength {
		col.Data = make([]string, outputSnippetLength)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == outputSnippetLength-1 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = strconv.FormatInt(v, 10)
	}
	return col
}

func NewColumnarFloat(name string, fullOutput bool, outputSnippetLength int, data []float64) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = name
	col.Type = typesys.Float64Type.ToString()
	col.ActualLength = len(data)
	if !fullOutput && len(data) > outputSnippetLength {
		col.Data = make([]string, outputSnippetLength)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == outputSnippetLength-1 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return col
}

func NewColumnarString(name string, fullOutput bool, outputSnippetLength int, data []string) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = name
	col.Type = typesys.StringType.ToString()
	col.ActualLength = len(data)
	if !fullOutput && len(data) > outputSnippetLength {
		col.Data = make([]string, outputSnippetLength)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == outputSnippetLength-1 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = v
	}
	return col
}

func DataFrameToColumnar(fullOutput bool, outputSnippetLength int, df *gandalff.DataFrame) ([]typesys.Columnar, error) {
	columns := make([]typesys.Columnar, (*df).NCols())

	for i, name := range (*df).Names() {
		col := (*df).Series(name)
		switch ser := col.(type) {
		case gandalff.SeriesBool:
			columns[i] = NewColumnarBool(ser.Name(), fullOutput, outputSnippetLength, ser.Bools())
		case gandalff.SeriesInt64:
			columns[i] = NewColumnarInt64(ser.Name(), fullOutput, outputSnippetLength, ser.Int64s())
		case gandalff.SeriesFloat64:
			columns[i] = NewColumnarFloat(ser.Name(), fullOutput, outputSnippetLength, ser.Float64s())
		case gandalff.SeriesString:
			columns[i] = NewColumnarString(ser.Name(), fullOutput, outputSnippetLength, ser.Strings())
		}
	}
	return columns, nil
}
