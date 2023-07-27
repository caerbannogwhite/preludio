package preludiocore

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gandalff"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type PreludioFunction func(funcName string, vm *ByteEater)

func PreludioFunc_Derive(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
	var series_ []gandalff.Series
	positional, _, err := vm.GetFunctionParams(funcName, nil, true, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if list, err := positional[1].getList(); err == nil {
		series_ = make([]gandalff.Series, len(list))

		for i, val := range list {
			switch col := val.getValue().(type) {
			case []bool:
				series_[i] = gandalff.NewSeriesBool(val.name, true, false, col)
			case []int64:
				series_[i] = gandalff.NewSeriesInt64(val.name, true, false, col)
			case []float64:
				series_[i] = gandalff.NewSeriesFloat64(val.name, true, false, col)
			case []string:
				series_[i] = gandalff.NewSeriesString(val.name, true, col, vm.__stringPool)
			}
		}
	} else {
		val := positional[1].getValue()
		series_ = make([]gandalff.Series, 1)
		switch col := val.(type) {
		case []bool:
			series_ = append(series_, gandalff.NewSeriesBool(positional[1].name, true, false, col))
		case []int64:
			series_ = append(series_, gandalff.NewSeriesInt64(positional[1].name, true, false, col))
		case []float64:
			series_ = append(series_, gandalff.NewSeriesFloat64(positional[1].name, true, false, col))
		case []string:
			series_ = append(series_, gandalff.NewSeriesString(positional[1].name, true, col, vm.__stringPool))
		}

	}

	for _, s := range series_ {
		df = df.AddSeries(s)
	}

	vm.stackPush(newPInternTerm(df))
}

// Describe a Dataframe
func PreludioFunc_Describe(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// expecting a Dataframe
	if len(positional) == 0 {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, "expecting at least one positional parameter."))
	} else {
		// var symbol __p_symbol__
		// var list __p_list__
		var df gandalff.DataFrame

		// Describe all
		if len(positional) == 1 {
			switch v := positional[0].getValue().(type) {
			case []bool:
			case []int64:
			case []float64:
			case []string:
			case __p_list__:
			case gandalff.DataFrame:
				df = v
			}

			vm.printInfo(0, fmt.Sprintln(df.Describe()))
			vm.stackPush(newPInternTerm(df))
		} else

		// Describe a subset
		if len(positional) == 2 {
			// names := make([]string, 0)
			// switch v := positional[1].getValue().(type) {
			// case __p_symbol__:
			// case __p_list__:
			// }

			fmt.Println(positional[1])

			// switch v := positional[0].getValue().(type) {
			// case []bool:
			// case []int:
			// case []float64:
			// case []string:
			// case __p_list__:
			// case gandalff.DataFrame:
			// 	fmt.Println(v.Select().Describe())
			// }
		}
	}
}

// Write a Dataframe into a CSV file
func PreludioFunc_WriteCSV(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
		"delimiter": newPInternTerm([]string{","}),
		"header":    newPInternTerm([]bool{true}),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	var header bool
	var path string
	var df gandalff.DataFrame
	var outputFile *os.File

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	path, err = positional[1].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	outputFile, err = os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	err = outputFile.Truncate(int64(0))
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	delimiter, err := named["delimiter"].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if len(delimiter) > 1 {
		vm.printWarning("delimiter length greater than 1, ignoring remaining characters")
	}

	header, err = named["header"].getBoolScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	res := df.ToCSV().
		SetDelimiter(rune(delimiter[0])).
		SetHeader(header).
		SetWriter(outputFile).
		Write()

	if res.IsErrored() {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, res.GetError()))
		return
	}
}

// Load a Dataframe from the Name Space
func PreludioFunc_From(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{}

	var err error
	var df gandalff.DataFrame

	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	vm.stackPush(newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Read a Dataframe form CSV file
func PreludioFunc_ReadCSV(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
		"delimiter": newPInternTerm([]string{","}),
		"header":    newPInternTerm([]bool{true}),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	var path, delimiter string
	var inputFile *os.File

	path, err = positional[0].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	inputFile, err = os.Open(path)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}
	defer inputFile.Close()

	delimiter, err = named["delimiter"].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// delimiter has to be a single character
	del := rune(delimiter[0])
	if len(delimiter) > 1 {
		if delimiter == "\\t" {
			del = '\t'
		} else {
			vm.printWarning("delimiter length greater than 1, ignoring remaining characters")
		}
	}

	header, err := named["header"].getBoolScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df := gandalff.NewBaseDataFrame().
		FromCSV().
		SetReader(inputFile).
		SetDelimiter(del).
		SetHeader(header).
		Read()

	if df.IsErrored() {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, df.GetError()))
		return
	}

	vm.stackPush(newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Create a new Dataframe
func PreludioFunc_New(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var list __p_list__
	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	list, err = positional[0].getList()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	s := make([]series.Series, len(list))
	for i, e := range list {
		if l, ok := e.getValue().(__p_list__); ok {

			switch l[0].getValue().(type) {
			case []bool:
				var vals []bool
				vals, err = e.listToBoolVector()
				if err != nil {
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
					return
				}
				s[i] = series.New(vals, series.Bool, e.name)
			case []int64:
				var vals []int64
				vals, err = e.listToInt64Vector()
				if err != nil {
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
					return
				}
				s[i] = series.New(vals, series.Int, e.name)
			case []float64:
				var vals []float64
				vals, err = e.listToFloat64Vector()
				if err != nil {
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
					return
				}
				s[i] = series.New(vals, series.Int, e.name)
			case []string:
				var vals []string
				vals, err = e.listToStringVector()
				if err != nil {
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
					return
				}
				s[i] = series.New(vals, series.String, e.name)
			}
		} else {
			vm.setPanicMode(fmt.Sprintf("%s: exprecting list for building dataframe, got %T", funcName, l))
			return
		}
	}

	vm.stackPush(newPInternTerm(dataframe.New(s...)))
	vm.setCurrentDataFrame()
}

// Select a subset of the Dataframe's columns
func PreludioFunc_Select(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
	var symbol __p_symbol__
	var list __p_list__
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, false)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// The first value can be both a symbol or a list of symbols
	symbol, err = positional[1].getSymbol()
	if err != nil {
		list, err = positional[1].getList()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}

		names := make([]string, len(list))
		for i, v := range list {
			symbol, err = v.getSymbol()
			if err != nil {
				vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
				return
			}
			names[i] = string(symbol)
		}
		df = df.Select(names...)
	} else {
		df = df.Select([]string{string(symbol)}...)
	}

	vm.stackPush(newPInternTerm(df))
}

// Sort all the values in the Dataframe
func PreludioFunc_Sort(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")
}

// Take a subset of the Dataframe's rows
func PreludioFunc_Take(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
	var num int64
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	num, err = positional[1].getInt64Scalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df = df.Take(0, int(num), 1)

	vm.stackPush(newPInternTerm(df))
}

///////////////////////////////////////////////////////////////////////////////
///////						ENVIRONMENT FUNCTIONS

// Set what's left in the stack to the current Dataframe
func PreludioFunc_ToCurrent(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// POSITIONAL PARAMETERS
	series_ := make(map[string]series.Series)
	switch len(positional) {

	// 1 PARAM
	case 1:
		switch v := positional[0].getValue().(type) {

		// BASE TYPES
		case []bool:
			series_[positional[0].name] = series.New(v, series.Bool, positional[0].name)
		case []int:
			series_[positional[0].name] = series.New(v, series.Int, positional[0].name)
		case []float64:
			series_[positional[0].name] = series.New(v, series.Float, positional[0].name)
		case []string:
			series_[positional[0].name] = series.New(v, series.String, positional[0].name)

		// LIST
		case __p_list__:
			for _, e := range v {
				switch t := e.getValue().(type) {
				case []bool:
					series_[e.name] = series.New(t, series.Bool, e.name)
				case []int:
					series_[e.name] = series.New(t, series.Int, e.name)
				case []float64:
					series_[e.name] = series.New(t, series.Float, e.name)
				case []string:
					series_[e.name] = series.New(t, series.String, e.name)
				default:
					vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, t))
					return
				}
			}
			vm.stackPush(newPInternTerm(v))

		// DATAFRAME
		case gandalff.DataFrame:
			// TODO

		default:
			vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, v))
			return
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting one positional parameter, received %d.", funcName, len(positional)))
		return
	}

	df := vm.__currentDataFrame
	names := make([]string, 0)
	for name := range series_ {
		names = append(names, name)
	}

	vals := make([]series.Series, len(series_))
	i := 0
	for _, s := range series_ {
		vals[i] = s
		i++
	}

	// TODO: fix this
	// df = df.Drop(names).CBind(dataframe.New(vals...))
	vm.__currentDataFrame = df
}

///////////////////////////////////////////////////////////////////////////////
///////						COERCION FUNCTIONS

// Coerce variables to bool
func PreludioFunc_AsBool(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")
}

// Coerce variables to integer
func PreludioFunc_AsInteger(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")
}

// Coerce variables to float
func PreludioFunc_AsFloat(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// POSITIONAL PARAMETERS
	switch len(positional) {

	// 1 PARAM
	case 1:
		switch v := positional[0].getValue().(type) {

		// BASE TYPES
		case []bool:
			res := make([]float64, len(v))
			for i := range v {
				if v[i] {
					res[i] = 1.0
				}
			}
			vm.stackPush(newPInternTerm(res))
			return

		case []int:
			res := make([]float64, len(v))
			for i := range v {
				res[i] = float64(v[i])
			}
			vm.stackPush(newPInternTerm(res))
			return

		case []float64:
			vm.stackPush(newPInternTerm(v))
			return

		case []string:
			res := make([]float64, len(v))
			for i := range v {
				res[i], err = strconv.ParseFloat(v[i], 64)
				if err != nil {
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
					return
				}
			}
			vm.stackPush(newPInternTerm(v))
			return

		// LIST
		case __p_list__:
			for i, e := range v {
				switch t := e.getValue().(type) {
				case []bool:
					res := make([]float64, len(t))
					for j := range t {
						if t[j] {
							res[j] = 1.0
						}
					}
					v[i] = *newPInternTerm(res)

				case []int:
					res := make([]float64, len(t))
					for j := range t {
						res[j] = float64(t[j])
					}
					v[i] = *newPInternTerm(res)

				case []float64:

				case []string:
					res := make([]float64, len(t))
					for j := range v {
						res[j], err = strconv.ParseFloat(t[j], 64)
						if err != nil {
							vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
							return
						}
					}
					v[i] = *newPInternTerm(res)

				}
			}
			vm.stackPush(newPInternTerm(v))

		// DATAFRAME
		case gandalff.DataFrame:
			// TODO
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting one positional parameter, received %d.", funcName, len(positional)))
		return
	}
}

// Coerce variables to string
func PreludioFunc_AsString(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")
}

///////////////////////////////////////////////////////////////////////////////
///////						STRING FUNCTIONS

func PreludioFunc_StrReplace(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
		"old": nil,
		"new": nil,
		"n":   newPInternTerm([]int{-1}),
	}

	var err error
	var num int64
	var strOld, strNew string
	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// NAMED PARAMETERS
	// GET old
	if named["old"] == nil {
		vm.setPanicMode(fmt.Sprintf("%s: nammed parameter 'old' is required since it has no default value.", funcName))
		return
	}
	strOld, err = named["old"].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// GET new
	if named["new"] == nil {
		vm.setPanicMode(fmt.Sprintf("%s: nammed parameter 'new' is required since it has no default value.", funcName))
		return
	}
	strNew, err = named["new"].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// GET num
	num, err = named["n"].getInt64Scalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// POSITIONAL PARAMETERS
	switch len(positional) {

	// 1 PARAM: string series or list of string series
	case 1:
		switch v := positional[0].getValue().(type) {

		// BASE TYPES
		case []string:
			for i := range v {
				v[i] = strings.Replace(v[i], strOld, strNew, int(num))
			}
			vm.stackPush(newPInternTerm(v))

		// LIST
		case __p_list__:
			for i, e := range v {
				switch t := e.getValue().(type) {
				case []string:
					for j := range v {
						t[j] = strings.Replace(t[j], strOld, strNew, int(num))
						if err != nil {
							vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
							return
						}
					}
					v[i] = *newPInternTerm(t)

				default:
					vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, t))
					return
				}
			}
			vm.stackPush(newPInternTerm(v))

		// DATAFRAME
		case gandalff.DataFrame:
			// TODO

		default:
			vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, v))
			return
		}

	// 2 PARAMS: dataframe, column name
	case 2:
		// df, err := positional[0].getDataframe()
		// if err != nil {
		// 	vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		// 	return
		// }

		switch v := positional[1].getValue().(type) {
		// case []string:
		// 	for i := range v {
		// 		df[v[i]] = strings.Replace(df[v[i]].([]string), strOld, strNew, num)
		// 	}
		// 	vm.stackPush(newPInternTerm(df))

		// case __p_list__:
		// 	for i, e := range v {
		// 		switch t := e.getValue().(type) {
		// 		case []string:
		// 			for j := range v {
		// 				df[t[j]] = strings.Replace(df[t[j]].([]string), strOld, strNew, num)
		// 			}
		// 			v[i] = *newPInternTerm(t)

		// 		default:
		// 			vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, t))
		// 			return
		// 		}
		// 	}
		// 	vm.stackPush(newPInternTerm(v))
		default:
			vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, v))
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting 1 or 2 positional parameters, received %d.", funcName, len(positional)))
		return
	}
}
