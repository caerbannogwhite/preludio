package preludiocore

import (
	"fmt"
	"os"

	"github.com/caerbannogwhite/enchanter/dataframe"
	"github.com/caerbannogwhite/enchanter/meta"
	"github.com/caerbannogwhite/enchanter/series"
)

type PreludioFunction func(funcName string, vm *ByteEater)

func PreludioFunc_Derive(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df dataframe.DataFrame
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
	case __p_list__:
		for _, val := range v {
			switch col := val.getValue().(type) {
			case series.Bools:
				df = df.AddSeries(val.name, col)
			case series.Int64s:
				df = df.AddSeries(val.name, col)
			case series.Float64s:
				df = df.AddSeries(val.name, col)
			case series.Strings:
				df = df.AddSeries(val.name, col)
			default:
				vm.setPanicMode(fmt.Sprintf("%s: expecting a list of Series, got %T", funcName, val))
				return
			}
		}

	// Derive: single column
	case series.Bools:
		df = df.AddSeries(positional[1].name, v)
	case series.Int64s:
		df = df.AddSeries(positional[1].name, v)
	case series.Float64s:
		df = df.AddSeries(positional[1].name, v)
	case series.Strings:
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
		// var symbol __p_symbol__
		// var list __p_list__
		var df dataframe.DataFrame

		// Describe all
		if len(positional) == 1 {
			switch v := positional[0].getValue().(type) {
			case []bool:
			case []int64:
			case []float64:
			case []string:
			case __p_list__:
			case dataframe.DataFrame:
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
			// case dataframe.DataFrame:
			// 	fmt.Println(v.Select().Describe())
			// }
		}
	}
}

// Write a Dataframe into a CSV file
func PreludioFunc_WriteCSV(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
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
	var df dataframe.DataFrame
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

	// O_WRONLY is required: without an access mode os.OpenFile opens the file
	// read-only (O_RDONLY is 0), and every write to it fails. That went
	// unnoticed because it only fails on some platforms — on Windows Go's
	// syscall.Open adds GENERIC_WRITE whenever O_CREAT is set, so the same
	// call happened to be writable there — and because the CSV writer does not
	// check its write errors, leaving an empty file behind and no error.
	// O_TRUNC replaces the separate Truncate call.
	outputFile, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}
	defer outputFile.Close()

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

	res := df.WriteCsv().
		SetDelimiter(del).
		SetHeader(header).
		SetWriter(outputFile).
		Write()

	if res != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, res.Error()))
		return
	}
}

// Filter rows of a Dataframe
func PreludioFunc_Filter(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{}

	var err error
	var df dataframe.DataFrame

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
	case series.Bools:
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
	var df dataframe.DataFrame

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

	named := map[string]*__p_intern__{
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

	df := dataframe.ReadCsv(vm.__context).
		SetReader(inputFile).
		SetDelimiter(del).
		SetHeader(header).
		Read()

	if df.Err() != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, df.Err()))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Get the names of the columns of a Dataframe
func PreludioFunc_Names(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var df dataframe.DataFrame
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

	var list __p_list__
	var err error
	var df dataframe.DataFrame

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

			df = dataframe.NewDataFrame(vm.__context)
			for _, p := range list {
				switch v := p.expr[0].(type) {
				case series.Series:
					df = df.AddSeries(p.name, v)
				default:
					vm.setPanicMode(fmt.Sprintf("%s: exprecting list of assignments for building a new dataframe, got %T", funcName, p.expr[0]))
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

	if df.Err() != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, df.Err()))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// anchorSelectors anchors each column selector.
//
// A preludio selector may be a pattern (an IDENT may contain '*', e.g.
// "a.*"), so selection goes through enchanter's SelectMatching, which
// matches regular expressions unanchored. Anchoring keeps a plain name
// selecting only the column with that name while patterns keep working.
func anchorSelectors(selectors []string) []string {
	anchored := make([]string, len(selectors))
	for i, s := range selectors {
		anchored[i] = "^(?:" + s + ")$"
	}
	return anchored
}

// Select a subset of the Dataframe's columns
func PreludioFunc_Select(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df dataframe.DataFrame
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
	case __p_symbol__:
		vm.stackPush(vm.newPInternTerm(df.SelectMatching(anchorSelectors([]string{string(v)})...)))
		vm.setCurrentDataFrame()

	case __p_list__:
		list, err := positional[1].listToStringSlice()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		vm.stackPush(vm.newPInternTerm(df.SelectMatching(anchorSelectors(list)...)))
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
	var df dataframe.DataFrame
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
	case __p_symbol__:
		vm.stackPush(vm.newPInternTerm(df.GroupBy(string(v))))
		vm.setCurrentDataFrame()

	case __p_list__:
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
	var df dataframe.DataFrame
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

	named := map[string]*__p_intern__{
		"on": nil,
	}

	var ok bool
	var err error
	var how __p_symbol__
	var df1, df2 dataframe.DataFrame
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
		if df2, ok = vm.symbolResolution(symb).(dataframe.DataFrame); !ok {
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

	var res dataframe.DataFrame
	switch how {
	case "inner":
		res = df1.Join(dataframe.JoinInner, df2, on...)
	case "outer":
		res = df1.Join(dataframe.JoinOuter, df2, on...)
	case "left":
		res = df1.Join(dataframe.JoinLeft, df2, on...)
	case "right":
		res = df1.Join(dataframe.JoinRight, df2, on...)
	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting one of 'inner', 'outer', 'left', 'right', got %s", funcName, how))
		return
	}

	if res.Err() != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, res.Err()))
		return
	}

	vm.stackPush(vm.newPInternTerm(res))
	vm.setCurrentDataFrame()
}

// Sort all the values in the Dataframe
func PreludioFunc_OrderBy(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df dataframe.DataFrame
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
	sortParams := make([]dataframe.SortParam, 0)
	switch v := positional[1].getValue().(type) {
	case __p_symbol__:
		if positional[1].isNeg() {
			sortParams = append(sortParams, dataframe.Desc(string(v)))
		} else {
			sortParams = append(sortParams, dataframe.Asc(string(v)))
		}

	case __p_list__:
		for _, v1 := range positional[1].getValue().(__p_list__) {
			switch v2 := v1.expr[0].(type) {
			case __p_symbol__:
				if v1.isNeg() {
					sortParams = append(sortParams, dataframe.Desc(string(v2)))
				} else {
					sortParams = append(sortParams, dataframe.Asc(string(v2)))
				}
			default:
				vm.setPanicMode(fmt.Sprintf("%s: expecting symbol, got %T", funcName, v))
				return
			}
		}

	default:
		vm.setPanicMode(fmt.Sprintf("%s: expecting symbol or list of symbols, got %T", funcName, v))
		return
	}

	vm.stackPush(vm.newPInternTerm(df.OrderBy(sortParams...)))
	vm.setCurrentDataFrame()
}

// Take a subset of the Dataframe's rows
func PreludioFunc_Take(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df dataframe.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	// A negative bound counts from the end, as the old Take did.
	norm := func(v int) int {
		if v < 0 {
			return df.NRows() + v
		}
		return v
	}

	switch len(positional) {
	case 2:
		a, err := positional[1].getInt64Scalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
			return
		}
		df = df.Slice(0, norm(int(a)))

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
		df = df.Slice(norm(int(a)), norm(int(b)))

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
		start, stop, step := norm(int(a)), norm(int(b)), int(c)
		if step < 0 {
			step = -step
		}
		if step == 0 {
			vm.setPanicMode(fmt.Sprintf("%s: step cannot be zero", funcName))
			return
		}
		indices := make([]int, 0, (stop-start+step-1)/step)
		for x := start; x < stop; x += step {
			indices = append(indices, x)
		}
		df = df.TakeIndices(indices)

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
	series_ := make(map[string]series.Series)
	switch len(positional) {

	// 1 PARAM
	case 1:
		switch v := positional[0].getValue().(type) {

		// BASE TYPES
		case []bool:
			series_[positional[0].name] = series.NewSeriesBool(v, nil, false, vm.__context)
		case []int64:
			series_[positional[0].name] = series.NewSeriesInt64(v, nil, false, vm.__context)
		case []float64:
			series_[positional[0].name] = series.NewSeriesFloat64(v, nil, false, vm.__context)
		case []string:
			series_[positional[0].name] = series.NewSeriesString(v, nil, false, vm.__context)

		// LIST
		case __p_list__:
			for _, e := range v {
				switch t := e.getValue().(type) {
				case []bool:
					series_[e.name] = series.NewSeriesBool(t, nil, false, vm.__context)
				case []int64:
					series_[e.name] = series.NewSeriesInt64(t, nil, false, vm.__context)
				case []float64:
					series_[e.name] = series.NewSeriesFloat64(t, nil, false, vm.__context)
				case []string:
					series_[e.name] = series.NewSeriesString(t, nil, false, vm.__context)
				default:
					vm.setPanicMode(fmt.Sprintf("%s: expected string, got %T.", funcName, t))
					return
				}
			}
			vm.stackPush(vm.newPInternTerm(v))

		// DATAFRAME
		case dataframe.DataFrame:
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

// Coerce series to a given type
func preludioAsType(funcName string, vm *ByteEater, coerceType meta.BaseType) {
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
		case dataframe.DataFrame:

			var names []string
			var series_ []series.Series

			switch t := positional[1].getValue().(type) {
			case series.Series:
				names = []string{positional[1].name}
				series_ = []series.Series{t}
			case __p_list__:
				names = make([]string, len(t))
				series_ = make([]series.Series, len(t))
				for i, e := range t {
					switch s := e.getValue().(type) {
					case series.Series:
						names[i] = e.name
						series_[i] = s
					default:
						vm.setPanicMode(fmt.Sprintf("%s: expecting series, got %T", funcName, s))
						return
					}
				}
			default:
				vm.setPanicMode(fmt.Sprintf("%s: expecting series, got %T", funcName, t))
				return
			}

			for i, s := range series_ {
				v = v.Replace(names[i], s.Cast(coerceType))
			}

			vm.stackPush(vm.newPInternTerm(v))

		case __p_list__:

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
			fmt.Println("TODO: StrReplace: []string")

		// LIST
		case __p_list__:
			fmt.Println("TODO: StrReplace: list")

		// DATAFRAME
		case dataframe.DataFrame:
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

		switch v := positional[1].expr[0].(type) {
		case series.Strings:
			df = df.Replace(positional[1].name, v.Replace(strOld, strNew, int(num)))
			vm.stackPush(vm.newPInternTerm(df))

		case series.Errors:
			vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, v.Err()))
			return

		case __p_list__:
			for _, e := range v {
				switch t := e.getValue().(type) {
				case series.Strings:
					df = df.Replace(e.name, t.Replace(strOld, strNew, int(num)))
				case series.Errors:
					vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, t.Err()))
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

///////////////////////////////////////////////////////////////////////////////
///////						AGGREGATION

// aggSlice builds a slice of enchanter aggregators. The aggregator type is
// unexported in enchanter, so its slice type cannot be written out here;
// type inference names it through this helper.
func aggSlice[T any](xs ...T) []T {
	return xs
}

// Aggregate a grouped Dataframe: agg! count:true mean:[MPG] sum:[Weight]
func PreludioFunc_Agg(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	listAggs := []string{"sum", "mean", "min", "max", "std", "variance", "median", "any", "all"}

	named := map[string]*__p_intern__{"count": nil}
	for _, name := range listAggs {
		named[name] = nil
	}

	var err error
	var df dataframe.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, &named, false, false)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	aggs := aggSlice(dataframe.Count())[:0]

	if named["count"] != nil {
		count, err := named["count"].getBoolScalar()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: count: %s", funcName, err))
			return
		}
		if count {
			aggs = append(aggs, dataframe.Count())
		}
	}

	for _, name := range listAggs {
		if named[name] == nil {
			continue
		}
		cols, err := named[name].listToStringSlice()
		if err != nil {
			vm.setPanicMode(fmt.Sprintf("%s: %s expects a list of columns: %s", funcName, name, err))
			return
		}
		for _, col := range cols {
			switch name {
			case "sum":
				aggs = append(aggs, dataframe.Sum(col))
			case "mean":
				aggs = append(aggs, dataframe.Mean(col))
			case "min":
				aggs = append(aggs, dataframe.Min(col))
			case "max":
				aggs = append(aggs, dataframe.Max(col))
			case "std":
				aggs = append(aggs, dataframe.Std(col))
			case "variance":
				aggs = append(aggs, dataframe.Variance(col))
			case "median":
				aggs = append(aggs, dataframe.Median(col))
			case "any":
				aggs = append(aggs, dataframe.Any(col))
			case "all":
				aggs = append(aggs, dataframe.All(col))
			}
		}
	}

	if len(aggs) == 0 {
		vm.setPanicMode(fmt.Sprintf("%s: expecting at least one aggregator, e.g. count:true or mean:[a]", funcName))
		return
	}

	res := df.Agg(aggs...).Run()
	if res.Err() != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, res.Err()))
		return
	}

	vm.stackPush(vm.newPInternTerm(res))
	vm.setCurrentDataFrame()
}

///////////////////////////////////////////////////////////////////////////////
///////						FILE FORMATS

// Read an Excel file: rxlsx! p'file.xlsx' sheet:'Sheet1'
func PreludioFunc_ReadXlsx(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
		"sheet": vm.newPInternTerm([]string{"Sheet1"}),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	path, err := positional[0].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	sheet, err := named["sheet"].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df := dataframe.ReadXlsx(vm.__context).SetPath(path).SetSheet(sheet).Read()
	if df.Err() != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, df.Err()))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Write an Excel file: wxlsx! p'file.xlsx' sheet:'Sheet1'
func PreludioFunc_WriteXlsx(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	named := map[string]*__p_intern__{
		"sheet": vm.newPInternTerm([]string{"Sheet1"}),
	}

	var err error
	var df dataframe.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, &named, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	path, err := positional[1].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	sheet, err := named["sheet"].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if err = df.WriteXlsx().SetPath(path).SetSheet(sheet).Write(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
}

// Read an XPT (SAS transport) file: rxpt! p'file.xpt'
func PreludioFunc_ReadXpt(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	path, err := positional[0].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df := dataframe.ReadXpt(vm.__context).SetPath(path).Read()
	if df.Err() != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, df.Err()))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}

// Write an XPT (SAS transport) file: wxpt! p'file.xpt'
func PreludioFunc_WriteXpt(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	var df dataframe.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if df, err = positional[0].getDataframe(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	path, err := positional[1].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	if err = df.WriteXpt().SetPath(path).Write(); err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
}

// Read a SAS7BDAT file (read-only, there is no writer): rsas! p'file.sas7bdat'
func PreludioFunc_ReadSas(funcName string, vm *ByteEater) {
	vm.printDebug(5, "STARTING", funcName, "")

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false, true)
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	path, err := positional[0].getStringScalar()
	if err != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, err))
		return
	}

	df := dataframe.ReadSas7bdat(vm.__context).SetPath(path).Read()
	if df.Err() != nil {
		vm.setPanicMode(fmt.Sprintf("%s: %s", funcName, df.Err()))
		return
	}

	vm.stackPush(vm.newPInternTerm(df))
	vm.setCurrentDataFrame()
}
