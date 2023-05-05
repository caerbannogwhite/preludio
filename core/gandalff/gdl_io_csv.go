package gandalff

import (
	"encoding/csv"
	"fmt"
	"os"

	"io"
	"regexp"
	"strconv"
	"typesys"
)

const CSV_READER_DEFAULT_DELIMITER = ','
const CSV_READER_DEFAULT_HEADER = true
const CSV_READER_DEFAULT_GUESS_DATA_TYPE_LEN = 1000

type GDLCsvReader struct {
	header           bool
	delimiter        rune
	guessDataTypeLen int
	path             string
	reader           io.Reader
	schema           *typesys.Schema
}

func NewGDLCsvReader() *GDLCsvReader {
	return &GDLCsvReader{
		header:           CSV_READER_DEFAULT_HEADER,
		delimiter:        CSV_READER_DEFAULT_DELIMITER,
		guessDataTypeLen: CSV_READER_DEFAULT_GUESS_DATA_TYPE_LEN,
		path:             "",
		reader:           nil,
		schema:           nil,
	}
}

func (r *GDLCsvReader) SetHeader(header bool) *GDLCsvReader {
	r.header = header
	return r
}

func (r *GDLCsvReader) SetDelimiter(delimiter rune) *GDLCsvReader {
	r.delimiter = delimiter
	return r
}

func (r *GDLCsvReader) SetGuessDataTypeLen(guessDataTypeLen int) *GDLCsvReader {
	r.guessDataTypeLen = guessDataTypeLen
	return r
}

func (r *GDLCsvReader) SetPath(path string) *GDLCsvReader {
	r.path = path
	return r
}

func (r *GDLCsvReader) SetReader(reader io.Reader) *GDLCsvReader {
	r.reader = reader
	return r
}

func (r *GDLCsvReader) SetSchema(schema *typesys.Schema) *GDLCsvReader {
	r.schema = schema
	return r
}

func (r *GDLCsvReader) Read() *GDLDataFrame {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return &GDLDataFrame{err: err}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return &GDLDataFrame{err: fmt.Errorf("GDLCsvReader: no reader specified")}
	}

	series, err := readCSV(r.reader, r.delimiter, r.header, r.guessDataTypeLen, r.schema)
	if err != nil {
		return &GDLDataFrame{err: err}
	}

	df := NewGDLDataFrame()
	for _, s := range series {
		df.AddSeries(s)
	}

	return df
}

type typeGuesser struct {
	boolRegex      *regexp.Regexp
	boolTrueRegex  *regexp.Regexp
	boolFalseRegex *regexp.Regexp
	intRegex       *regexp.Regexp
	floatRegex     *regexp.Regexp
}

// Get the regexes for guessing data types
func newTypeGuesser() typeGuesser {
	boolRegex := regexp.MustCompile(`^([Tt]([Rr][Uu][Ee])?)|([Ff]([Aa][Ll][Ss][Ee])?)$`)
	intRegex := regexp.MustCompile(`^[-+]?[0-9]+$`)
	floatRegex := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?$`)

	boolTrueRegex := regexp.MustCompile(`^[Tt]([Rr][Uu][Ee])?$`)
	boolFalseRegex := regexp.MustCompile(`^[Ff]([Aa][Ll][Ss][Ee])?$`)

	return typeGuesser{boolRegex, boolTrueRegex, boolFalseRegex, intRegex, floatRegex}
}

func (tg typeGuesser) guessType(s string) typesys.BaseType {
	if tg.boolRegex.MatchString(s) {
		return typesys.BoolType
	} else if tg.intRegex.MatchString(s) {
		return typesys.Int32Type
	} else if tg.floatRegex.MatchString(s) {
		return typesys.Float64Type
	}
	return typesys.StringType
}

func (tg typeGuesser) atoBool(s string) (bool, error) {
	if tg.boolTrueRegex.MatchString(s) {
		return true, nil
	} else if tg.boolFalseRegex.MatchString(s) {
		return false, nil
	}
	return false, fmt.Errorf("cannot convert \"%s\" to bool", s)
}

// ReadCSV reads a CSV file and returns a GDLDataFrame.
func readCSV(reader io.Reader, delimiter rune, header bool, guessDataTypeLen int, schema *typesys.Schema) ([]GDLSeries, error) {

	// TODO: Add support for reading CSV files with missing values
	// TODO: Try to optimize this function by using goroutines: read the rows (like 1000)
	//		and guess the data types in parallel

	isNullable := false
	stringPool := NewStringPool()

	// Initialize TypeGuesser
	tg := newTypeGuesser()

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
			return nil, err
		}
	}

	var dataTypes []typesys.BaseType
	var recordsForGuessing [][]string

	// Guess data types
	if schema == nil {
		recordsForGuessing = make([][]string, guessDataTypeLen)

		for i := 0; i < guessDataTypeLen; i++ {
			record, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			recordsForGuessing[i] = record

			for j, v := range record {
				if i == 0 {
					dataTypes = append(dataTypes, tg.guessType(v))
				} else {
					if dataTypes[j] == typesys.StringType {
						continue
					}
					if tg.guessType(v) != dataTypes[j] {
						dataTypes[j] = typesys.StringType
					}
				}
			}
		}
	} else

	// Use schema
	{
		dataTypes = schema.GetDataTypes()
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

	// If no schema: add records for guessing to values
	if schema == nil {
		for _, record := range recordsForGuessing {
			for i, v := range record {
				switch dataTypes[i] {
				case typesys.BoolType:
					b, err := tg.atoBool(v)
					if err != nil {
						return nil, err
					}
					values[i] = append(values[i].([]bool), b)
				case typesys.Int32Type:
					d, err := strconv.Atoi(v)
					if err != nil {
						return nil, err
					}
					values[i] = append(values[i].([]int), d)
				case typesys.Float64Type:
					f, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return nil, err
					}
					values[i] = append(values[i].([]float64), f)
				case typesys.StringType:
					values[i] = append(values[i].([]string), v)
				}
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
				b, err := tg.atoBool(v)
				if err != nil {
					return nil, err
				}
				values[i] = append(values[i].([]bool), b)
			case typesys.Int32Type:
				d, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				values[i] = append(values[i].([]int), d)
			case typesys.Float64Type:
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, err
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
	series := make([]GDLSeries, len(dataTypes))
	for i, name := range names {
		switch dataTypes[i] {
		case typesys.BoolType:
			series[i] = NewGDLSeriesBool(name, isNullable, values[i].([]bool))
		case typesys.Int32Type:
			series[i] = NewGDLSeriesInt32(name, isNullable, false, values[i].([]int))
		case typesys.Float64Type:
			series[i] = NewGDLSeriesFloat64(name, isNullable, false, values[i].([]float64))
		case typesys.StringType:
			series[i] = NewGDLSeriesString(name, isNullable, values[i].([]string), stringPool)
		}
	}

	return series, nil
}
