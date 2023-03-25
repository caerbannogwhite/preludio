package gandalff

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"typesys"
)

type TypeGuesser struct {
	boolRegex      *regexp.Regexp
	boolTrueRegex  *regexp.Regexp
	boolFalseRegex *regexp.Regexp
	intRegex       *regexp.Regexp
	floatRegex     *regexp.Regexp
}

// Get the regexes for guessing data types
func NewTypeGuesser() TypeGuesser {
	boolRegex := regexp.MustCompile(`^([Tt]([Rr][Uu][Ee])?)|([Ff]([Aa][Ll][Ss][Ee])?)$`)
	intRegex := regexp.MustCompile(`^[-+]?[0-9]+$`)
	floatRegex := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?$`)

	boolTrueRegex := regexp.MustCompile(`^[Tt]([Rr][Uu][Ee])?$`)
	boolFalseRegex := regexp.MustCompile(`^[Ff]([Aa][Ll][Ss][Ee])?$`)

	return TypeGuesser{boolRegex, boolTrueRegex, boolFalseRegex, intRegex, floatRegex}
}

func (tg TypeGuesser) GuessType(s string) typesys.BaseType {
	if tg.boolRegex.MatchString(s) {
		return typesys.BoolType
	} else if tg.intRegex.MatchString(s) {
		return typesys.Int32Type
	} else if tg.floatRegex.MatchString(s) {
		return typesys.Float64Type
	}
	return typesys.StringType
}

func (tg TypeGuesser) AtoBool(s string) (bool, error) {
	if tg.boolTrueRegex.MatchString(s) {
		return true, nil
	} else if tg.boolFalseRegex.MatchString(s) {
		return false, nil
	}
	return false, fmt.Errorf("cannot convert \"%s\" to bool", s)
}

// FromCSV reads a CSV file and returns a GDLDataFrame.
func FromCSV(reader io.Reader, delimiter rune, header bool, guessDataTypeLen int) *GDLDataFrame {

	// TODO: Add support for reading CSV files with missing values
	// TODO: Try to optimize this function by using goroutines: read the rows (like 1000)
	//		and guess the data types in parallel

	isNullable := false
	stringPool := NewStringPool()

	// Initialize TypeGuesser
	tg := NewTypeGuesser()

	// Initialize GDLDataFrame
	df := NewGDLDataFrame()

	// Initialize CSV reader
	csvReader := csv.NewReader(reader)
	csvReader.Comma = delimiter
	csvReader.FieldsPerRecord = -1

	// Read header if present
	var names []string
	var err error
	if header {
		names, err = csvReader.Read()
		if err != nil {
			df.err = err
			return df
		}
	}

	// Guess data types
	var dataTypes []typesys.BaseType
	recordsForGuessing := make([][]string, 0)

	for i := 0; i < guessDataTypeLen; i++ {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		recordsForGuessing = append(recordsForGuessing, record)

		for j, v := range record {
			if i == 0 {
				dataTypes = append(dataTypes, tg.GuessType(v))
			} else {
				if dataTypes[j] == typesys.StringType {
					continue
				}
				if tg.GuessType(v) != dataTypes[j] {
					dataTypes[j] = typesys.StringType
				}
			}
		}
	}

	values := make([]interface{}, len(dataTypes))
	for i := range values {
		switch dataTypes[i] {
		case typesys.BoolType:
			values[i] = make([]bool, 0)
		case typesys.Int32Type:
			values[i] = make([]int, 0)
		case typesys.Float64Type:
			values[i] = make([]float64, 0)
		case typesys.StringType:
			values[i] = make([]string, 0)
		}
	}

	// Add records for guessing to values
	for _, record := range recordsForGuessing {
		for i, v := range record {
			switch dataTypes[i] {
			case typesys.BoolType:
				b, err := tg.AtoBool(v)
				if err != nil {
					df.err = err
					return df
				}
				values[i] = append(values[i].([]bool), b)
			case typesys.Int32Type:
				d, err := strconv.Atoi(v)
				if err != nil {
					df.err = err
					return df
				}
				values[i] = append(values[i].([]int), d)
			case typesys.Float64Type:
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					df.err = err
					return df
				}
				values[i] = append(values[i].([]float64), f)
			case typesys.StringType:
				values[i] = append(values[i].([]string), v)
			}
		}
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		for i, v := range record {
			switch dataTypes[i] {
			case typesys.BoolType:
				b, err := tg.AtoBool(v)
				if err != nil {
					df.err = err
					return df
				}
				values[i] = append(values[i].([]bool), b)
			case typesys.Int32Type:
				d, err := strconv.Atoi(v)
				if err != nil {
					df.err = err
					return df
				}
				values[i] = append(values[i].([]int), d)
			case typesys.Float64Type:
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					df.err = err
					return df
				}
				values[i] = append(values[i].([]float64), f)
			case typesys.StringType:
				values[i] = append(values[i].([]string), v)
			}
		}
	}

	// Generate names if not present
	if !header {
		for i := 0; i < len(dataTypes); i++ {
			names = append(names, fmt.Sprintf("Column %d", i+1))
		}
	}

	// Create series
	for i, name := range names {
		switch dataTypes[i] {
		case typesys.BoolType:
			df.AddSeries(NewGDLSeriesBool(name, isNullable, false, values[i].([]bool)))
		case typesys.Int32Type:
			df.AddSeries(NewGDLSeriesInt32(name, isNullable, false, values[i].([]int)))
		case typesys.Float64Type:
			df.AddSeries(NewGDLSeriesFloat64(name, isNullable, false, values[i].([]float64)))
		case typesys.StringType:
			df.AddSeries(NewGDLSeriesString(name, isNullable, values[i].([]string), stringPool))
		}
	}

	return df
}
