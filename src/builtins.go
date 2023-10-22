package preludiocore

import (
	"fmt"
	"os"
	"preludiometa"

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

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	switch v := positional[1].getValue().(type) {

	// Derive: paramenter is list, multiple columns
	case pList:
		for _, val := range v {
			switch col := val.getValue().(type) {
			case gandalff.SeriesBool:
				df = df.AddSeries(val.name, col)
			case gandalff.SeriesInt64:
				df = df.AddSeries(val.name, col)
			case gandalff.SeriesFloat64:
				df = df.AddSeries(val.name, col)
			case gandalff.SeriesString:
				df = df.AddSeries(val.name, col)
			default:
				vm.setPanicMode(fmt.Sprintf("%s: expecting a list of Series, got %T", funcName, val))
				return
			}
		}

	// Derive: single column
	case gandalff.SeriesBool:
		df = df.AddSeries(positional[1].name, v)
	case gandalff.SeriesInt64:
		df = df.AddSeries(positional[1].name, v)
	case gandalff.SeriesFloat64:
		df = df.AddSeries(positional[1].name, v)
	case gandalff.SeriesString:
		df = df.AddSeries(positional[1].name, v)

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting a Series, got %T", funcName, v))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
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
		// var symbol pSymbol
		// var list pList
		var df gandalff.DataFrame

		// Describe all
		if len(positional) == 1 {
			switch v := positional[0].getValue().(type) {
			case []bool:
			case []int64:
			case []float64:
			case []string:
			case pList:
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
			// case pSymbol:
			// case pList:
			// }

			fmt.Println(positional[1])

			// switch v := positional[0].getValue().(type) {
			// case []bool:
			// case []int:
			// case []float64:
			// case []string:
			// case pList:
			// case gandalff.DataFrame:
			// 	fmt.Println(v.Select().Describe())
			// }
		}
	}
}

// Write a Dataframe into a CSV file
func PreludioFunc_WriteCSV(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*pIntern{
		"del":  vm.newPInternTerm([]string{","}),
		"head": vm.newPInternTerm([]bool{true}),
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

	if df, err = positional[0].getDataframe(); err != nil {
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
	defer outputFile.Close()

	err = outputFile.Truncate(int64(0))
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	delimiter, err := named["del"].getStringScalar()
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

	header, err = named["head"].getBoolScalar()
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

	named := map[string]*pIntern{}

	var err error
	var df gandalff.DataFrame

	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
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

	named := map[string]*pIntern{}

	var err error
	var df gandalff.DataFrame

	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Read a Dataframe form CSV file
func PreludioFunc_ReadCSV(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*pIntern{
		"del":  vm.newPInternTerm([]string{","}),
		"head": vm.newPInternTerm([]bool{true}),
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

	delimiter, err = named["del"].getStringScalar()
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

	header, err := named["head"].getBoolScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df := gandalff.NewBaseDataFrame(vm.__context).
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

// Get the names of the columns of a Dataframe
func PreludioFunc_Names(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var df gandalff.DataFrame
	var err error

	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	vm.stackPush(vm.newPInternTerm(df.Names()))
}

// Create a new Dataframe
func PreludioFunc_New(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var list pList
	var err error
	var df gandalff.DataFrame

	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	switch len(positional) {
	case 1:
		// New: parameter is a list of assignments
		if positional[0].isList() {
			list, err = positional[0].getList()
			if err != nil {
				vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
				return
			}

			df = gandalff.NewBaseDataFrame(vm.__context)
			for _, p := range list {
				switch v := p.value.(type) {
				case gandalff.Series:
					df = df.AddSeries(p.name, v)
				default:
					vm.setPanicMode(fmt.Sprintf("%s: exprecting list of assignments for building a new dataframe, got %T", funcName, p.value))
					return
				}
			}
		} else {
			vm.setPanicMode(fmt.Sprintf("%s: expecting assignment for building a new dataframe, got %T", funcName, positional[0].getValue()))
			return
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, "expecting exactly one positional parameter."))
		return
	}

	if df.IsErrored() {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, df.GetError()))
		return
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

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// The first value can be both a symbol or a list of symbols
	switch v := positional[1].getValue().(type) {
	case pSymbol:
		vm.stackPush(vm.newPInternTerm(df.Select(string(v))))
		vm.setCurrentDataFrame()

	case pList:
		list, err := positional[1].listToStringSlice()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		vm.stackPush(vm.newPInternTerm(df.Select(list...)))
		vm.setCurrentDataFrame()

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting symbol or list of symbols, got %T", funcName, v))
		return
	}
}

// Group a Dataframe
func PreludioFunc_GroupBy(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, false)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// The first value can be both a symbol or a list of symbols
	switch v := positional[1].getValue().(type) {
	case pSymbol:
		vm.stackPush(vm.newPInternTerm(df.GroupBy(string(v))))
		vm.setCurrentDataFrame()

	case pList:
		list, err := positional[1].listToStringSlice()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}

		vm.stackPush(vm.newPInternTerm(df.GroupBy(list...)))
		vm.setCurrentDataFrame()

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting symbol or list of symbols, got %T", funcName, v))
		return
	}
}

// Ungroup a Dataframe
func PreludioFunc_Ungroup(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, false)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	vm.stackPush(vm.newPInternTerm(df.Ungroup()))
	vm.setCurrentDataFrame()
}

// Join two Dataframes
func PreludioFunc_Join(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*pIntern{
		"on": nil,
	}

	var ok bool
	var err error
	var how pSymbol
	var df1, df2 gandalff.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, &named, false, false)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if len(positional) != 3 {
		vm.setPanicMode(fmt.Sprintf("%s: expecting 3 positional arguments, got %d", funcName, len(positional)))
		return
	}

	// Left dataframe
	if df1, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// How
	if how, err = positional[1].getSymbol(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// Right dataframe
	if symb, err := positional[2].getSymbol(); err == nil {
		if df2, ok = vm.symbolResolution(symb).(gandalff.DataFrame); !ok {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, "expecting dataframe"))
			return
		}
	} else {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// By
	on := make([]string, 0)
	if named["on"] != nil {
		if on, err = named["on"].listToStringSlice(); err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
	}

	var res gandalff.DataFrame
	switch how {
	case "inner":
		res = df1.Join(gandalff.INNER_JOIN, df2, on...)
	case "outer":
		res = df1.Join(gandalff.OUTER_JOIN, df2, on...)
	case "left":
		res = df1.Join(gandalff.LEFT_JOIN, df2, on...)
	case "right":
		res = df1.Join(gandalff.RIGHT_JOIN, df2, on...)
	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting one of 'inner', 'outer', 'left', 'right', got %s", funcName, how))
		return
	}

	if res.IsErrored() {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, res.GetError()))
		return
	}

	vm.stackPush(vm.newPInternTerm(res))
	vm.setCurrentDataFrame()
}

// Sort all the values in the Dataframe
func PreludioFunc_OrderBy(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df gandalff.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, false)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if len(positional) != 2 {
		vm.setPanicMode(fmt.Sprintf("%s: expecting 2 parameters, got %d", funcName, len(positional)))
		return
	}

	// The first value can be both a symbol or a list of symbols
	sortParams := make([]gandalff.SortParam, 0)
	// switch v := positional[1].getValue().(type) {
	// case pSymbol:
	// 	if positional[1].isNeg() {
	// 		sortParams = append(sortParams, gandalff.Desc(string(v)))
	// 	} else {
	// 		sortParams = append(sortParams, gandalff.Asc(string(v)))
	// 	}

	// case pList:
	// 	for _, v1 := range positional[1].getValue().(pList) {
	// 		switch v2 := v1.value.(type) {
	// 		case pSymbol:
	// 			if v1.isNeg() {
	// 				sortParams = append(sortParams, gandalff.Desc(string(v2)))
	// 			} else {
	// 				sortParams = append(sortParams, gandalff.Asc(string(v2)))
	// 			}
	// 		default:
	// 			vm.setPanicMode(fmt.Sprintf("%s: expecting symbol, got %T", funcName, v))
	// 			return
	// 		}
	// 	}

	// default:
	// 	vm.setPanicMode(fmt.Sprintf("%s: expecting symbol or list of symbols, got %T", funcName, v))
	// 	return
	// }

	vm.stackPush(vm.newPInternTerm(df.OrderBy(sortParams...)))
	vm.setCurrentDataFrame()
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

	if df, err = positional[0].getDataframe(); err != nil {
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
			series_[positional[0].name] = gandalff.NewSeriesBool(v, nil, false, vm.__context)
		case []int64:
			series_[positional[0].name] = gandalff.NewSeriesInt64(v, nil, false, vm.__context)
		case []float64:
			series_[positional[0].name] = gandalff.NewSeriesFloat64(v, nil, false, vm.__context)
		case []string:
			series_[positional[0].name] = gandalff.NewSeriesString(v, nil, false, vm.__context)

		// LIST
		case pList:
			for _, e := range v {
				switch t := e.getValue().(type) {
				case []bool:
					series_[e.name] = gandalff.NewSeriesBool(t, nil, false, vm.__context)
				case []int64:
					series_[e.name] = gandalff.NewSeriesInt64(t, nil, false, vm.__context)
				case []float64:
					series_[e.name] = gandalff.NewSeriesFloat64(t, nil, false, vm.__context)
				case []string:
					series_[e.name] = gandalff.NewSeriesString(t, nil, false, vm.__context)
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

// Coerce series to a given type
func preludioAsType(funcName string, vm *ByteEater, coerceType preludiometa.BaseType) {
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

			var names []string
			var series []gandalff.Series

			switch t := positional[1].getValue().(type) {
			case gandalff.Series:
				names = []string{positional[1].name}
				series = []gandalff.Series{t}
			case pList:
				names = make([]string, len(t))
				series = make([]gandalff.Series, len(t))
				for i, e := range t {
					switch s := e.getValue().(type) {
					case gandalff.Series:
						names[i] = e.name
						series[i] = s
					default:
						vm.setPanicMode(fmt.Sprintf("%s: expecting series, got %T", funcName, s))
						return
					}
				}
			default:
				vm.setPanicMode(fmt.Sprintf("%s: expecting series, got %T", funcName, t))
				return
			}

			for i, s := range series {
				v = v.Replace(names[i], s.Cast(coerceType))
			}

			vm.stackPush(vm.newPInternTerm(v))

		case pList:

		default:
			vm.setPanicMode(fmt.Sprintf("%s: expecting dataframe or list, got %T", funcName, v))
			return
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting 1 or 2 parameters, got %d", funcName, len(positional)))
		return
	}
}

///////////////////////////////////////////////////////////////////////////////
///////						STRING FUNCTIONS

func PreludioFunc_StrReplace(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*pIntern{
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
			fmt.Println("TODO: StrReplace: []string")

		// LIST
		case pList:
			fmt.Println("TODO: StrReplace: list")

		// DATAFRAME
		case gandalff.DataFrame:
			fmt.Println("TODO: StrReplace: dataframe")

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

		switch v := positional[1].value.(type) {
		case gandalff.SeriesString:
			df = df.Replace(positional[1].name, v.Replace(strOld, strNew, int(num)))
			vm.stackPush(vm.newPInternTerm(df))

		case gandalff.SeriesError:
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, v.GetError()))
			return

		case pList:
			for _, e := range v {
				switch t := e.getValue().(type) {
				case gandalff.SeriesString:
					df = df.Replace(e.name, t.Replace(strOld, strNew, int(num)))
				case gandalff.SeriesError:
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, t.GetError()))
					return
				default:
					vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, t))
					return
				}
			}
			vm.stackPush(vm.newPInternTerm(df))

		default:
			vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, v))
			return
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting 1 or 2 positional parameters, received %d.", funcName, len(positional)))
		return
	}
}
