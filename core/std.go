package preludiocore

import (
	"fmt"
	"os"
	"strings"
	"typesys"

	"gandalff"
)

type PreludioFunction func(funcName string, vm *ByteEater)

func PreludioFunc_Derive(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
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

	switch v := positional[1].getValue().(type) {

	// Derive: paramenter is list, multiple columns
	case __p_list__:
		for _, val := range v {
			switch col := val.getValue().(type) {
			case gandalff.SeriesBool:
				df = df.AddSeries(col.SetName(val.name))
			case gandalff.SeriesInt64:
				df = df.AddSeries(col.SetName(val.name))
			case gandalff.SeriesFloat64:
				df = df.AddSeries(col.SetName(val.name))
			case gandalff.SeriesString:
				df = df.AddSeries(col.SetName(val.name))
			default:
				vm.setPanicMode(fmt.Sprintf("%s: expecting a list of Series, got %T", funcName, val))
			}
		}

	// Derive: single column
	case gandalff.SeriesBool:
		df = df.AddSeries(v.SetName(positional[1].name))
	case gandalff.SeriesInt64:
		df = df.AddSeries(v.SetName(positional[1].name))
	case gandalff.SeriesFloat64:
		df = df.AddSeries(v.SetName(positional[1].name))
	case gandalff.SeriesString:
		df = df.AddSeries(v.SetName(positional[1].name))

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting a Series, got %T", funcName, v))
	}

	vm.stackPush(vm.newPInternTerm(df))
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
			vm.stackPush(vm.newPInternTerm(df))
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
		"delimiter": vm.newPInternTerm([]string{","}),
		"header":    vm.newPInternTerm([]bool{true}),
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

	// delimiter has to be a single character
	del := rune(delimiter[0])
	if len(delimiter) > 1 {
		if delimiter == "\\t" {
			del = '\t'
		} else {
			vm.printWarning("delimiter length greater than 1, ignoring remaining characters")
		}
	}

	header, err = named["header"].getBoolScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	res := df.ToCSV().
		SetDelimiter(del).
		SetHeader(header).
		SetWriter(outputFile).
		Write()

	if res.IsErrored() {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, res.GetError()))
		return
	}
}

// Filter rows of a Dataframe
func PreludioFunc_Filter(funcName string, vm *ByteEater) {
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

	switch v := positional[1].getValue().(type) {
	case gandalff.SeriesBool:
		vm.stackPush(vm.newPInternTerm(df.Filter(v)))

	default:
		vm.setPanicMode(fmt.Sprintf("%s: invalid type %T", funcName, v))
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

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Read a Dataframe form CSV file
func PreludioFunc_ReadCSV(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
		"delimiter": vm.newPInternTerm([]string{","}),
		"header":    vm.newPInternTerm([]bool{true}),
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

	vm.stackPush(vm.newPInternTerm(df))
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

	df := gandalff.NewBaseDataFrame()

	var ser gandalff.Series
	for _, e := range list {
		if l, ok := e.getValue().(__p_list__); ok {
			switch l[0].getValue().(type) {
			case []bool:
				ser, err = e.listToSeriesBool()
			case []int64:
				ser, err = e.listToSeriesInt64()
			case []float64:
				ser, err = e.listToSeriesFloat64()
			case []string:
				ser, err = e.listToSeriesString()
			}

			if err != nil {
				vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
				return
			}
			df.AddSeries(ser)

		} else {
			vm.setPanicMode(fmt.Sprintf("%s: exprecting list for building dataframe, got %T", funcName, l))
			return
		}
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Select a subset of the Dataframe's columns
func PreludioFunc_Select(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
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
	switch v := positional[1].getValue().(type) {
	case __p_symbol__:
		vm.stackPush(vm.newPInternTerm(df.Select(string(v))))

	case __p_list__:
		list, err := positional[1].listToStringSlice()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		vm.stackPush(vm.newPInternTerm(df.Select(list...)))

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting symbol or list of symbols, got %T", funcName, v))
	}
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

	switch len(positional) {
	case 2:
		a, err := positional[1].getInt64Scalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		df = df.Take(int(a))

	case 3:
		a, err := positional[1].getInt64Scalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		b, err := positional[2].getInt64Scalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		df = df.Take(int(a), int(b))

	case 4:
		a, err := positional[1].getInt64Scalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		b, err := positional[2].getInt64Scalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		c, err := positional[3].getInt64Scalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		df = df.Take(int(a), int(b), int(c))

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting 2, 3 or 4 parameters, got %d", funcName, len(positional)))
	}

	vm.stackPush(vm.newPInternTerm(df))
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
	series_ := make(map[string]gandalff.Series)
	switch len(positional) {

	// 1 PARAM
	case 1:
		switch v := positional[0].getValue().(type) {

		// BASE TYPES
		case []bool:
			series_[positional[0].name] = gandalff.NewSeriesBool(positional[0].name, true, false, v)
		case []int64:
			series_[positional[0].name] = gandalff.NewSeriesInt64(positional[0].name, true, false, v)
		case []float64:
			series_[positional[0].name] = gandalff.NewSeriesFloat64(positional[0].name, true, false, v)
		case []string:
			series_[positional[0].name] = gandalff.NewSeriesString(positional[0].name, true, v, vm.__stringPool)

		// LIST
		case __p_list__:
			for _, e := range v {
				switch t := e.getValue().(type) {
				case []bool:
					series_[e.name] = gandalff.NewSeriesBool(e.name, true, false, t)
				case []int64:
					series_[e.name] = gandalff.NewSeriesInt64(e.name, true, false, t)
				case []float64:
					series_[e.name] = gandalff.NewSeriesFloat64(e.name, true, false, t)
				case []string:
					series_[e.name] = gandalff.NewSeriesString(e.name, true, t, vm.__stringPool)
				default:
					vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, t))
					return
				}
			}
			vm.stackPush(vm.newPInternTerm(v))

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

	vals := make([]gandalff.Series, len(series_))
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

func preludioAsType(funcName string, vm *ByteEater, coerceType typesys.BaseType) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// POSITIONAL PARAMETERS
	switch len(positional) {
	case 1:

	case 2:
		switch v := positional[0].getValue().(type) {
		case gandalff.DataFrame:

			var cols []string
			switch t := positional[1].getValue().(type) {
			case string:
				cols = []string{t}
			case __p_symbol__:
				cols = []string{string(t)}
			case __p_list__:
				cols, err = positional[1].listToStringSlice()
				if err != nil {
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
					return
				}
			default:
				vm.setPanicMode(fmt.Sprintf("%s: expecting string or list of strings, got %T", funcName, t))
			}

			for _, col := range cols {
				v = v.Replace(col, v.Series(col).Cast(coerceType, vm.__stringPool))
			}

		case __p_list__:

		default:
			vm.setPanicMode(fmt.Sprintf("%s: expecting dataframe or list, got %T", funcName, v))
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting 1 or 2 parameters, got %d", funcName, len(positional)))
	}
}

// Coerce variables to bool
func PreludioFunc_AsBool(funcName string, vm *ByteEater) {
	preludioAsType(funcName, vm, typesys.BoolType)
}

// Coerce variables to integer
func PreludioFunc_AsInteger(funcName string, vm *ByteEater) {
	preludioAsType(funcName, vm, typesys.Int64Type)
}

// Coerce variables to float
func PreludioFunc_AsFloat(funcName string, vm *ByteEater) {
	preludioAsType(funcName, vm, typesys.Float64Type)
}

// Coerce variables to string
func PreludioFunc_AsString(funcName string, vm *ByteEater) {
	preludioAsType(funcName, vm, typesys.StringType)
}

///////////////////////////////////////////////////////////////////////////////
///////						STRING FUNCTIONS

func PreludioFunc_StrReplace(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
		"old": nil,
		"new": nil,
		"n":   vm.newPInternTerm([]int64{-1}),
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
			vm.stackPush(vm.newPInternTerm(v))

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
					v[i] = *vm.newPInternTerm(t)

				default:
					vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, t))
					return
				}
			}
			vm.stackPush(vm.newPInternTerm(v))

		// DATAFRAME
		case gandalff.DataFrame:
			// TODO

		default:
			vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, v))
			return
		}

	// 2 PARAMS: dataframe, column name
	case 2:
		df, err := positional[0].getDataframe()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}

		switch v := positional[1].getValue().(type) {
		case gandalff.SeriesString:
			df = df.Replace(v.Name(), v.Replace(strOld, strNew, int(num)))
			vm.stackPush(vm.newPInternTerm(df))

		case gandalff.SeriesError:
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, v.GetError()))
			return

		case __p_list__:
			fmt.Println("TODO: strReplace list of series")

		default:
			vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, v))
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting 1 or 2 positional parameters, received %d.", funcName, len(positional)))
		return
	}
}
