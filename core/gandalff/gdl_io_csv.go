package gandalff

import (
	"encoding/csv"
	"fmt"
	"os"

	"io"
	"preludiometa"
	"regexp"
	"strconv"
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
	schema           *preludiometa.Schema
	ctx              *Context
}

func NewCsvReader(ctx *Context) *CsvReader {
	return &CsvReader{
		header:           CSV_READER_DEFAULT_HEADER,
		delimiter:        CSV_READER_DEFAULT_DELIMITER,
		guessDataTypeLen: CSV_READER_DEFAULT_GUESS_DATA_TYPE_LEN,
		path:             "",
		nullValues:       false,
		reader:           nil,
		schema:           nil,
		ctx:              ctx,
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

func (r *CsvReader) SetSchema(schema *preludiometa.Schema) *CsvReader {
	r.schema = schema
	return r
}

func (r *CsvReader) SetContext(ctx *Context) *CsvReader {
	r.ctx = ctx
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

	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("CsvReader: no context specified")}
	}

	names, series, err := readCSV(r.reader, r.delimiter, r.header, r.nullValues, r.guessDataTypeLen, r.schema, r.ctx)
	if err != nil {
		return BaseDataFrame{err: err}
	}

	df := NewBaseDataFrame(r.ctx)
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

func (tg typeGuesser) guessType(s string) preludiometa.BaseType {
	if tg.boolRegex.MatchString(s) {
		return preludiometa.BoolType
	} else if tg.intRegex.MatchString(s) {
		return preludiometa.Int64Type
	} else if tg.floatRegex.MatchString(s) {
		return preludiometa.Float64Type
	}
	return preludiometa.StringType
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
func readCSV(reader io.Reader, delimiter rune, header bool, nullValues bool, guessDataTypeLen int, schema *preludiometa.Schema, ctx *Context) ([]string, []Series, error) {

	// TODO: Add support for Time and Duration types (defined in a schema)
	// TODO: Optimize null masks (use bit vectors)?
	// TODO: Try to optimize this function by using goroutines: read the rows (like 1000)
	//		and guess the data types in parallel

	if ctx == nil {
		return nil, nil, fmt.Errorf("readCSV: no context specified")
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

	var dataTypes []preludiometa.BaseType
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
					if dataTypes[j] == preludiometa.StringType {
						continue
					}
					if tg.guessType(v) != dataTypes[j] {
						dataTypes[j] = preludiometa.StringType
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
		case preludiometa.BoolType:
			values[i] = make([]bool, 0)
		case preludiometa.IntType:
			values[i] = make([]int, 0)
		case preludiometa.Int64Type:
			values[i] = make([]int64, 0)
		case preludiometa.Float64Type:
			values[i] = make([]float64, 0)
		case preludiometa.StringType:
			values[i] = make([]*string, 0)
		}
	}

	// If no schema: add records for guessing to values
	if schema == nil {
		if nullValues {
			for _, record := range recordsForGuessing {
				for i, v := range record {
					switch dataTypes[i] {
					case preludiometa.BoolType:
						if b, err := tg.atoBool(v); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]bool), b)
						}

					case preludiometa.IntType:
						if d, err := strconv.Atoi(v); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int), int(d))
						}

					case preludiometa.Int64Type:
						if d, err := strconv.ParseInt(v, 10, 64); err != nil {
							return nil, nil, err
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int64), d)
						}

					case preludiometa.Float64Type:
						if f, err := strconv.ParseFloat(v, 64); err != nil {
							return nil, nil, err
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]float64), f)
						}

					case preludiometa.StringType:
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
					}
				}
			}
		} else {
			for _, record := range recordsForGuessing {
				for i, v := range record {
					switch dataTypes[i] {
					case preludiometa.BoolType:
						b, err := tg.atoBool(v)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]bool), b)

					case preludiometa.IntType:
						d, err := strconv.Atoi(v)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]int), int(d))

					case preludiometa.Int64Type:
						d, err := strconv.ParseInt(v, 10, 64)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]int64), d)

					case preludiometa.Float64Type:
						f, err := strconv.ParseFloat(v, 64)
						if err != nil {
							return nil, nil, err
						}
						values[i] = append(values[i].([]float64), f)

					case preludiometa.StringType:
						values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
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
				case preludiometa.BoolType:
					if b, err := tg.atoBool(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]bool), b)
					}

				case preludiometa.IntType:
					if d, err := strconv.Atoi(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int), int(d))
					}

				case preludiometa.Int64Type:
					if d, err := strconv.ParseInt(v, 10, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int64), d)
					}

				case preludiometa.Float64Type:
					if f, err := strconv.ParseFloat(v, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]float64), f)
					}

				case preludiometa.StringType:
					nullMasks[i] = append(nullMasks[i], false)
					values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
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
				case preludiometa.BoolType:
					b, err := tg.atoBool(v)
					if err != nil {
						return nil, nil, err
					}
					values[i] = append(values[i].([]bool), b)

				case preludiometa.IntType:
					d, err := strconv.Atoi(v)
					if err != nil {
						return nil, nil, err
					}
					values[i] = append(values[i].([]int), int(d))

				case preludiometa.Int64Type:
					d, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return nil, nil, err
					}
					values[i] = append(values[i].([]int64), d)

				case preludiometa.Float64Type:
					f, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return nil, nil, err
					}
					values[i] = append(values[i].([]float64), f)

				case preludiometa.StringType:
					values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
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
		case preludiometa.BoolType:
			series[i] = NewSeriesBool(values[i].([]bool), nullMasks[i], false, ctx)

		case preludiometa.IntType:
			series[i] = NewSeriesInt(values[i].([]int), nullMasks[i], false, ctx)

		case preludiometa.Int64Type:
			series[i] = NewSeriesInt64(values[i].([]int64), nullMasks[i], false, ctx)

		case preludiometa.Float64Type:
			series[i] = NewSeriesFloat64(values[i].([]float64), nullMasks[i], false, ctx)

		case preludiometa.StringType:
			series[i] = SeriesString{
				isNullable: nullValues,
				data:       values[i].([]*string),
				nullMask:   __binVecFromBools(nullMasks[i]),
				ctx:        ctx,
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
				fmt.Fprintf(writer, "%s", s.GetAsString(i))
			}
		}

		fmt.Fprintf(writer, "\n")
	}

	return nil
}
