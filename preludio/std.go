package preludio

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type PreludioFunction func(funcName string, vm *ByteEater)

func PreludioFunc_Derive(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	var df dataframe.DataFrame
	var series_ []series.Series
	positional, _, err := vm.GetFunctionParams(funcName, nil, true)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	if list, err := positional[1].getList(); err == nil {
		series_ = make([]series.Series, len(list))

		for i, val := range list {
			switch col := val.getValue().(type) {
			case []bool:
				series_[i] = series.New(col, series.Bool, val.Name)
			case []int:
				series_[i] = series.New(col, series.Int, val.Name)
			case []float64:
				series_[i] = series.New(col, series.Float, val.Name)
			case []string:
				series_[i] = series.New(col, series.String, val.Name)
			}
		}
	} else {
		val := positional[1].getValue()
		series_ = make([]series.Series, 1)
		switch col := val.(type) {
		case []bool:
			series_ = append(series_, series.New(col, series.Bool, positional[1].Name))
		case []int:
			series_ = append(series_, series.New(col, series.Int, positional[1].Name))
		case []float64:
			series_ = append(series_, series.New(col, series.Float, positional[1].Name))
		case []string:
			series_ = append([]series.Series{}, series.New(col, series.String, positional[1].Name))
		}

	}

	df = df.CBind(dataframe.New(series_...))
	vm.StackPush(NewPreludioInternalTerm(df))
}

// Describe a Dataframe
func PreludioFunc_Describe(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// expecting a Dataframe
	if len(positional) == 0 {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, "expecting at least one positional parameter.")))
	} else {
		// var symbol PreludioSymbol
		// var list PreludioList
		var df dataframe.DataFrame

		// Describe all
		if len(positional) == 1 {
			switch v := positional[0].getValue().(type) {
			case []bool:
			case []int:
			case []float64:
			case []string:
			case PreludioList:
			case dataframe.DataFrame:
				df = v
			}

			vm.__output.log = append(vm.__output.log, fmt.Sprintln(df.Describe()))
			vm.StackPush(NewPreludioInternalTerm(df))
		} else

		// Describe a subset
		if len(positional) == 2 {
			// names := make([]string, 0)
			// switch v := positional[1].getValue().(type) {
			// case PreludioSymbol:
			// case PreludioList:
			// }

			fmt.Println(positional[1])

			// switch v := positional[0].getValue().(type) {
			// case []bool:
			// case []int:
			// case []float64:
			// case []string:
			// case PreludioList:
			// case dataframe.DataFrame:
			// 	fmt.Println(v.Select().Describe())
			// }
		}
	}
}

// Export a Dataframe into a CSV file
func PreludioFunc_ExportCsv(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PreludioInternal{
		// "delimiter": NewPreludioInternalTerm([]string{","}),
		"header": NewPreludioInternalTerm([]bool{true}),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, &named, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	var header bool
	var path string
	var df dataframe.DataFrame
	var outputFile *os.File

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	path, err = positional[1].getStringScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	outputFile, err = os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	err = outputFile.Truncate(int64(0))
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// delimiter, err = named["delimiter"].getStringScalar()
	// if err != nil {
	// 	vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
	// 	return
	// }

	// if len(delimiter) > 1 {
	// 	vm.PrintWarning("delimiter length greater than 1, ignoring remaining characters")
	// }

	header, err = named["header"].getBoolScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	err = df.WriteCSV(outputFile, dataframe.WriteHeader(header))
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, df.Error())))
		return
	}
}

// Load a Dataframe from the Name Space
func PreludioFunc_From(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PreludioInternal{}

	var err error
	var df dataframe.DataFrame

	positional, _, err := vm.GetFunctionParams(funcName, &named, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, df.Error())))
		return
	}

	vm.StackPush(NewPreludioInternalTerm(df))
	vm.SetCurrentDataFrame()
}

// Import a Dataframe form CSV file
func PreludioFunc_ImportCsv(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PreludioInternal{
		"delimiter": NewPreludioInternalTerm([]string{","}),
		"header":    NewPreludioInternalTerm([]bool{true}),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, &named, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	var path, delimiter string
	var inputFile *os.File

	path, err = positional[0].getStringScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	inputFile, err = os.Open(path)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	delimiter, err = named["delimiter"].getStringScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	if len(delimiter) > 1 {
		vm.PrintWarning("delimiter length greater than 1, ignoring remaining characters")
	}

	df := dataframe.ReadCSV(inputFile, dataframe.WithDelimiter(rune(delimiter[0])))
	if df.Error() != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, df.Error())))
		return
	}

	vm.StackPush(NewPreludioInternalTerm(df))
	vm.SetCurrentDataFrame()
}

// Create a new Dataframe
func PreludioFunc_New(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var list PreludioList
	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	list, err = positional[0].getList()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	s := make([]series.Series, len(list))
	for i, e := range list {
		if l, ok := e.getValue().(PreludioList); ok {

			switch l[0].getValue().(type) {
			case []bool:
				var vals []bool
				vals, err = e.ListToBoolVector()
				if err != nil {
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
					return
				}
				s[i] = series.New(vals, series.Bool, e.Name)
			case []int:
				var vals []int
				vals, err = e.ListToIntegerVector()
				if err != nil {
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
					return
				}
				s[i] = series.New(vals, series.Int, e.Name)
			case []float64:
				var vals []float64
				vals, err = e.ListToFloatVector()
				if err != nil {
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
					return
				}
				s[i] = series.New(vals, series.Int, e.Name)
			case []string:
				var vals []string
				vals, err = e.ListToStringVector()
				if err != nil {
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
					return
				}
				s[i] = series.New(vals, series.String, e.Name)
			}
		} else {
			vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: exprecting list for building dataframe, got %T", funcName, l)))
			return
		}
	}

	vm.StackPush(NewPreludioInternalTerm(dataframe.New(s...)))
	vm.SetCurrentDataFrame()
}

// Select a subset of the Dataframe's columns
func PreludioFunc_Select(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	var df dataframe.DataFrame
	var symbol PreludioSymbol
	var list PreludioList
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// The first value can be both a symbol or la list of symbols
	symbol, err = positional[1].getSymbol()
	if err != nil {
		list, err = positional[1].getList()
		if err != nil {
			vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
			return
		}

		names := make([]string, len(list))
		for i, v := range list {
			symbol, err = v.getSymbol()
			if err != nil {
				vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
				return
			}
			names[i] = string(symbol)
		}
		df = df.Select(names)
	} else {
		df = df.Select([]string{string(symbol)})
	}

	vm.StackPush(NewPreludioInternalTerm(df))
}

// Sort all the values in the Dataframe
func PreludioFunc_Sort(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

// Take a subset of the Dataframe's rows
func PreludioFunc_Take(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	var df dataframe.DataFrame
	var num int
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].getDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	num, err = positional[1].getIntegerScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	rows := make([]int, num)
	for i, _ := range rows {
		rows[i] = i
	}

	df = df.Subset(rows)

	vm.StackPush(NewPreludioInternalTerm(df))
}

///////////////////////////////////////////////////////////////////////////////
///////						ENVIRONMENT FUNCTIONS

// Set what's left in the stack to the current Dataframe
func PreludioFunc_ToCurrent(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
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
			series_[positional[0].Name] = series.New(v, series.Bool, positional[0].Name)
		case []int:
			series_[positional[0].Name] = series.New(v, series.Int, positional[0].Name)
		case []float64:
			series_[positional[0].Name] = series.New(v, series.Float, positional[0].Name)
		case []string:
			series_[positional[0].Name] = series.New(v, series.String, positional[0].Name)

		// LIST
		case PreludioList:
			for _, e := range v {
				switch t := e.getValue().(type) {
				case []bool:
					series_[e.Name] = series.New(t, series.Bool, e.Name)
				case []int:
					series_[e.Name] = series.New(t, series.Int, e.Name)
				case []float64:
					series_[e.Name] = series.New(t, series.Float, e.Name)
				case []string:
					series_[e.Name] = series.New(t, series.String, e.Name)
				default:
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: expected string, got %T.", funcName, t)))
					return
				}
			}
			vm.StackPush(NewPreludioInternalTerm(v))

		// DATAFRAME
		case dataframe.DataFrame:
			// TODO

		default:
			vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: expected string, got %T.", funcName, v)))
			return
		}

	default:
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: expecting one positional parameter, received %d.", funcName, len(positional))))
		return
	}

	df := *vm.__currentDataFrame
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

	df = df.Drop(names).CBind(dataframe.New(vals...))
	vm.__currentDataFrame = &df
}

///////////////////////////////////////////////////////////////////////////////
///////						COERCION FUNCTIONS

// Coerce variables to bool
func PreludioFunc_AsBool(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

// Coerce variables to integer
func PreludioFunc_AsInteger(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

// Coerce variables to float
func PreludioFunc_AsFloat(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
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
			vm.StackPush(NewPreludioInternalTerm(res))
			return

		case []int:
			res := make([]float64, len(v))
			for i := range v {
				res[i] = float64(v[i])
			}
			vm.StackPush(NewPreludioInternalTerm(res))
			return

		case []float64:
			vm.StackPush(NewPreludioInternalTerm(v))
			return

		case []string:
			res := make([]float64, len(v))
			for i := range v {
				res[i], err = strconv.ParseFloat(v[i], 64)
				if err != nil {
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
					return
				}
			}
			vm.StackPush(NewPreludioInternalTerm(v))
			return

		// LIST
		case PreludioList:
			for i, e := range v {
				switch t := e.getValue().(type) {
				case []bool:
					res := make([]float64, len(t))
					for j := range t {
						if t[j] {
							res[j] = 1.0
						}
					}
					v[i] = *NewPreludioInternalTerm(res)

				case []int:
					res := make([]float64, len(t))
					for j := range t {
						res[j] = float64(t[j])
					}
					v[i] = *NewPreludioInternalTerm(res)

				case []float64:

				case []string:
					res := make([]float64, len(t))
					for j := range v {
						res[j], err = strconv.ParseFloat(t[j], 64)
						if err != nil {
							vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
							return
						}
					}
					v[i] = *NewPreludioInternalTerm(res)

				}
			}
			vm.StackPush(NewPreludioInternalTerm(v))

		// DATAFRAME
		case dataframe.DataFrame:
			// TODO
		}

	default:
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: expecting one positional parameter, received %d.", funcName, len(positional))))
		return
	}
}

// Coerce variables to string
func PreludioFunc_AsString(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

///////////////////////////////////////////////////////////////////////////////
///////						STRING FUNCTIONS

func PreludioFunc_StrReplace(funcName string, vm *ByteEater) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PreludioInternal{
		"old": nil,
		"new": nil,
		"n":   NewPreludioInternalTerm([]int{-1}),
	}

	var err error
	var num int
	var strOld, strNew string
	positional, _, err := vm.GetFunctionParams(funcName, &named, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// NAMED PARAMETERS
	// GET old
	if named["old"] == nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: nammed parameter 'old' is required since it has no default value.", funcName)))
		return
	}
	strOld, err = named["old"].getStringScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// GET new
	if named["new"] == nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: nammed parameter 'new' is required since it has no default value.", funcName)))
		return
	}
	strNew, err = named["new"].getStringScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// GET num
	num, err = named["n"].getIntegerScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// POSITIONAL PARAMETERS
	switch len(positional) {

	// 1 PARAM
	case 1:
		switch v := positional[0].getValue().(type) {

		// BASE TYPES
		case []string:
			for i := range v {
				v[i] = strings.Replace(v[i], strOld, strNew, num)
			}
			vm.StackPush(NewPreludioInternalTerm(v))

		// LIST
		case PreludioList:
			for i, e := range v {
				switch t := e.getValue().(type) {
				case []string:
					for j := range v {
						t[j] = strings.Replace(t[j], strOld, strNew, num)
						if err != nil {
							vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
							return
						}
					}
					v[i] = *NewPreludioInternalTerm(t)

				default:
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: expected string, got %T.", funcName, t)))
					return
				}
			}
			vm.StackPush(NewPreludioInternalTerm(v))

		// DATAFRAME
		case dataframe.DataFrame:
			// TODO

		default:
			vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: expected string, got %T.", funcName, v)))
			return
		}

	default:
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: expecting one positional parameter, received %d.", funcName, len(positional))))
		return
	}
}
