package preludio

import (
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type LogEnty struct {
	LogType LOG_TYPE `json:"logType"`
	Level   uint8    `json:"level"`
	Message string   `json:"message"`
}

type Columnar struct {
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	ActualLength int      `json:"actualLength"` // actual length of the column
	Data         []string `json:"data"`
}

type PreludioOutput struct {
	Log  []LogEnty    `json:"log"`
	Data [][]Columnar `json:"data"`
}

func NewColumnarBool(name string, fullOutput bool, data []bool) Columnar {
	col := Columnar{}
	col.Name = name
	col.Type = "bool"
	col.ActualLength = len(data)
	if !fullOutput {
		col.Data = make([]string, 5)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == 4 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = strconv.FormatBool(v)
	}
	return col
}

func NewColumnarInt(name string, fullOutput bool, data []int) Columnar {
	col := Columnar{}
	col.Name = name
	col.Type = "int"
	col.ActualLength = len(data)
	if !fullOutput {
		col.Data = make([]string, 5)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == 4 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = strconv.Itoa(v)
	}
	return col
}

func NewColumnarFloat(name string, fullOutput bool, data []float64) Columnar {
	col := Columnar{}
	col.Name = name
	col.Type = "float"
	col.ActualLength = len(data)
	if !fullOutput {
		col.Data = make([]string, 5)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == 4 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return col
}

func NewColumnarString(name string, fullOutput bool, data []string) Columnar {
	col := Columnar{}
	col.Name = name
	col.Type = "string"
	col.ActualLength = len(data)
	if !fullOutput {
		col.Data = make([]string, 5)
	} else {
		col.Data = make([]string, len(data))
	}
	for i, v := range data {
		if !fullOutput && i == 4 {
			col.Data[i] = "..."
			break
		}
		col.Data[i] = v
	}
	return col
}

func DataFrameToColumnar(fullOutput bool, df *dataframe.DataFrame) ([]Columnar, error) {
	columns := make([]Columnar, df.Ncol())
	for i, name := range df.Names() {
		col := df.Col(name)
		switch col.Type() {
		case series.Bool:
			val, err := col.Bool()
			if err != nil {
				return nil, err
			}
			columns[i] = NewColumnarBool(col.Name, fullOutput, val)
		case series.Int:
			val, err := col.Int()
			if err != nil {
				return nil, err
			}
			columns[i] = NewColumnarInt(col.Name, fullOutput, val)
		case series.Float:
			columns[i] = NewColumnarFloat(col.Name, fullOutput, col.Float())
		case series.String:
			columns[i] = NewColumnarString(col.Name, fullOutput, col.Records())
		}
	}
	return columns, nil
}
