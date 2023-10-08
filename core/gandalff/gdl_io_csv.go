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

type CsvReader struct {
	header           bool
	delimiter        rune
	guessDataTypeLen int
	path             string
	nullValues       bool
	reader           io.Reader
	schema           *typesys.Schema
	pool             *StringPool
}

func NewCsvReader(pool *StringPool) *CsvReader {
	return &CsvReader{
		header:           CSV_READER_DEFAULT_HEADER,
		delimiter:        CSV_READER_DEFAULT_DELIMITER,
		guessDataTypeLen: CSV_READER_DEFAULT_GUESS_DATA_TYPE_LEN,
		path:             "",
		nullValues:       false,
		reader:           nil,
		schema:           nil,
		pool:             pool,
	}
}

func (r *CsvReader) SetHeader(header bool) *CsvReader {
	r.header = header
	return r
}

func (r *CsvReader) SetDelimiter(delimiter rune) *CsvReader {
	r.delimiter = delimiter
	return r
}

func (r *CsvReader) SetGuessDataTypeLen(guessDataTypeLen int) *CsvReader {
	r.guessDataTypeLen = guessDataTypeLen
	return r
}

func (r *CsvReader) SetPath(path string) *CsvReader {
	r.path = path
	return r
}

func (r *CsvReader) SetNullValues(nullValues bool) *CsvReader {
	r.nullValues = nullValues
	return r
}

func (r *CsvReader) SetReader(reader io.Reader) *CsvReader {
	r.reader = reader
	return r
}

func (r *CsvReader) SetSchema(schema *typesys.Schema) *CsvReader {
	r.schema = schema
	return r
}

func (r *CsvReader) Read() DataFrame {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return BaseDataFrame{err: fmt.Errorf("CsvReader: no reader specified")}
	}

	if r.pool == nil {
		r.pool = NewStringPool()
	}

	names, series, err := readCSV(r.reader, r.delimiter, r.header, r.nullValues, r.guessDataTypeLen, r.schema, r.pool)
	if err != nil {
		return BaseDataFrame{err: err}
	}

	df := NewBaseDataFrame().SetStringPool(r.pool)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
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
		return typesys.Int64Type
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
func readCSV(reader io.Reader, delimiter rune, header bool, nullValues bool, guessDataTypeLen int, schema *typesys.Schema, stringPool *StringPool) ([]string, []Series, error) {

	// TODO: Add support for Time and Duration types (defined in a schema)
	// TODO: Optimize null masks (use bit vectors)?
	// TODO: Try to optimize this function by using goroutines: read the rows (like 1000)
	//		and guess the data types in parallel

	if stringPool == nil {
		return nil, nil, fmt.Errorf("readCSV: string pool cannot be nil")
	}

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
			return nil, nil, err
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

	nullMasks := make([][]bool, len(dataTypes))
	if nullValues {
		for i := range nullMasks {
			nullMasks[i] = make([]bool, 0)
		}
	}

	values := make([]interface{}, len(dataTypes))
	for i := range values {
		switch dataTypes[i] {
		case typesys.BoolType:
			values[i] = make([]bool, 0)
		case typesys.IntType:
			values[i] = make([]int, 0)
		case typesys.Int64Type:
			values[i] = make([]int64, 0)
		case typesys.Float64Type:
			values[i] = make([]float64, 0)
		case typesys.StringType:
			values[i] = make([]*string, 0)
		}
	}

	// If no schema: add records for guessing to values
	if schema == nil {
		if nullValues {
			for _, record := range recordsForGuessing {
				for i, v := range record {
					switch dataTypes[i] {
					case typesys.BoolType:
						if b, err := tg.atoBool(v); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]bool), b)
						}

					case typesys.IntType:
						if d, err := strconv.Atoi(v); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int), int(d))
						}

					case typesys.Int64Type:
						if d, err := strconv.ParseInt(v, 10, 64); err != nil {
							return nil, nil, err
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int64), d)
						}

					case typesys.Float64Type:
						if f, err := strconv.ParseFloat(v, 64); err != nil {
							return nil, nil, err
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]float64), f)
						}

					case typesys.StringType:
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]*string), stringPool.Put(v))
					}
				}
			}
		} else {
			for _, record := range recordsForGuessing {
				for i, v := range record {
					switch dataTypes[i] {
					case typesys.BoolType:
						b, err := tg.atoBool(v)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]bool), b)

					case typesys.IntType:
						d, err := strconv.Atoi(v)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]int), int(d))

					case typesys.Int64Type:
						d, err := strconv.ParseInt(v, 10, 64)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]int64), d)

					case typesys.Float64Type:
						f, err := strconv.ParseFloat(v, 64)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]float64), f)

					case typesys.StringType:
						values[i] = append(values[i].([]*string), stringPool.Put(v))
					}
				}
			}
		}
	}

	if nullValues {
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
					if b, err := tg.atoBool(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]bool), b)
					}

				case typesys.IntType:
					if d, err := strconv.Atoi(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int), int(d))
					}

				case typesys.Int64Type:
					if d, err := strconv.ParseInt(v, 10, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int64), d)
					}

				case typesys.Float64Type:
					if f, err := strconv.ParseFloat(v, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]float64), f)
					}

				case typesys.StringType:
					nullMasks[i] = append(nullMasks[i], false)
					values[i] = append(values[i].([]*string), stringPool.Put(v))
				}
			}
		}
	} else {
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
						return nil, nil, err
					}
					values[i] = append(values[i].([]bool), b)

				case typesys.IntType:
					d, err := strconv.Atoi(v)
					if err != nil {
						return nil, nil, err
					}
					values[i] = append(values[i].([]int), int(d))

				case typesys.Int64Type:
					d, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return nil, nil, err
					}
					values[i] = append(values[i].([]int64), d)

				case typesys.Float64Type:
					f, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return nil, nil, err
					}
					values[i] = append(values[i].([]float64), f)

				case typesys.StringType:
					values[i] = append(values[i].([]*string), stringPool.Put(v))
				}
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
	series := make([]Series, len(names))
	for i := range names {
		switch dataTypes[i] {
		case typesys.BoolType:
			series[i] = NewSeriesBool(values[i].([]bool), nullMasks[i], false, stringPool)

		case typesys.IntType:
			series[i] = NewSeriesInt(values[i].([]int), nullMasks[i], false, stringPool)

		case typesys.Int64Type:
			series[i] = NewSeriesInt64(values[i].([]int64), nullMasks[i], false, stringPool)

		case typesys.Float64Type:
			series[i] = NewSeriesFloat64(values[i].([]float64), nullMasks[i], false, stringPool)

		case typesys.StringType:
			series[i] = SeriesString{
				isNullable: nullValues,
				data:       values[i].([]*string),
				nullMask:   __binVecFromBools(nullMasks[i]),
				pool:       stringPool,
			}
		}
	}

	return names, series, nil
}

type CsvWriter struct {
	delimiter  rune
	header     bool
	path       string
	nullString string
	writer     io.Writer
	dataframe  DataFrame
}

func NewCsvWriter() *CsvWriter {
	return &CsvWriter{
		delimiter:  CSV_READER_DEFAULT_DELIMITER,
		header:     CSV_READER_DEFAULT_HEADER,
		path:       "",
		nullString: NULL_STRING,
		writer:     nil,
		dataframe:  nil,
	}
}

func (w *CsvWriter) SetDelimiter(delimiter rune) *CsvWriter {
	w.delimiter = delimiter
	return w
}

func (w *CsvWriter) SetHeader(header bool) *CsvWriter {
	w.header = header
	return w
}

func (w *CsvWriter) SetPath(path string) *CsvWriter {
	w.path = path
	return w
}

func (w *CsvWriter) SetNullString(nullString string) *CsvWriter {
	w.nullString = nullString
	return w
}

func (w *CsvWriter) SetWriter(writer io.Writer) *CsvWriter {
	w.writer = writer
	return w
}

func (w *CsvWriter) SetDataFrame(dataframe DataFrame) *CsvWriter {
	w.dataframe = dataframe
	return w
}

func (w *CsvWriter) Write() DataFrame {
	err := writeCSV(w.dataframe, w.writer, w.delimiter, w.header, w.nullString)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err}
	}

	return w.dataframe
}

func writeCSV(df DataFrame, writer io.Writer, delimiter rune, header bool, nullString string) error {
	series := make([]Series, df.NCols())
	for i := 0; i < df.NCols(); i++ {
		series[i] = df.SeriesAt(i)
	}

	if writer == nil {
		return fmt.Errorf("writeCSV: no writer specified")
	}

	if header {
		for i, name := range df.Names() {
			if i > 0 {
				fmt.Fprintf(writer, "%c", delimiter)
			}
			fmt.Fprintf(writer, "%s", name)
		}

		fmt.Fprintf(writer, "\n")
	}

	for i := 0; i < df.NRows(); i++ {
		for j, s := range series {
			if j > 0 {
				fmt.Fprintf(writer, "%c", delimiter)
			}

			if s.IsNull(i) {
				fmt.Fprintf(writer, "%s", nullString)
			} else {
				fmt.Fprintf(writer, "%s", s.GetString(i))
			}
		}

		fmt.Fprintf(writer, "\n")
	}

	return nil
}
