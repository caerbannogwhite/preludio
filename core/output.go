package preludiocore

import (
	"strconv"
	"typesys"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func NewColumnarBool(name string, fullOutput bool, outputSnippetLength int, data []bool) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = name
	col.Type = "bool"
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

func NewColumnarInt(name string, fullOutput bool, outputSnippetLength int, data []int) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = name
	col.Type = "int"
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
		col.Data[i] = strconv.Itoa(v)
	}
	return col
}

func NewColumnarFloat(name string, fullOutput bool, outputSnippetLength int, data []float64) typesys.Columnar {
	col := typesys.Columnar{}
	col.Name = name
	col.Type = "float"
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
	col.Type = "string"
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

func DataFrameToColumnar(fullOutput bool, outputSnippetLength int, df *dataframe.DataFrame) ([]typesys.Columnar, error) {
	columns := make([]typesys.Columnar, df.Ncol())
	for i, name := range df.Names() {
		col := df.Col(name)
		switch col.Type() {
		case series.Bool:
			val, err := col.Bool()
			if err != nil {
				return nil, err
			}
			columns[i] = NewColumnarBool(col.Name, fullOutput, outputSnippetLength, val)
		case series.Int:
			val, err := col.Int()
			if err != nil {
				return nil, err
			}
			columns[i] = NewColumnarInt(col.Name, fullOutput, outputSnippetLength, val)
		case series.Float:
			columns[i] = NewColumnarFloat(col.Name, fullOutput, outputSnippetLength, col.Float())
		case series.String:
			columns[i] = NewColumnarString(col.Name, fullOutput, outputSnippetLength, col.Records())
		}
	}
	return columns, nil
}
