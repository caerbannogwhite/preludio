package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type PreludioFunction func(funcName string, vm *PreludioVM)

func PreludioFunc_Derive(funcName string, vm *PreludioVM) {
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

	df, err = positional[0].GetDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	if positional[1].IsList() {
		list, _ := positional[1].GetList()
		series_ = make([]series.Series, len(list))

		for i, val := range list {
			switch col := val.GetValue().(type) {
			case []bool:
				series_[i] = series.New(col, series.Bool, val.Name)
			case []int:
				series_[i] = series.New(col, series.Int, val.Name)
			case []float64:
				series_[i] = series.New(col, series.Float, (*val).Name)
			case []string:
				series_[i] = series.New(col, series.String, val.Name)
			}
		}
	} else {
		val := positional[1].GetValue()
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
func PreludioFunc_Describe(funcName string, vm *PreludioVM) {
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
			switch v := positional[0].GetValue().(type) {
			case []bool:
			case []int:
			case []float64:
			case []string:
			case PreludioList:
			case dataframe.DataFrame:
				df = v
			}

			fmt.Println(df.Describe())
			vm.StackPush(NewPreludioInternalTerm(df))
		} else

		// Describe a subset
		if len(positional) == 2 {
			// names := make([]string, 0)
			// switch v := positional[1].GetValue().(type) {
			// case PreludioSymbol:
			// case PreludioList:
			// }

			fmt.Println(positional[1])

			// switch v := positional[0].GetValue().(type) {
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
func PreludioFunc_ExportCsv(funcName string, vm *PreludioVM) {
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

	df, err = positional[0].GetDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	path, err = positional[1].GetStringScalar()
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

	// delimiter, err = named["delimiter"].GetStringScalar()
	// if err != nil {
	// 	vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
	// 	return
	// }

	// if len(delimiter) > 1 {
	// 	vm.PrintWarning("delimiter length greater than 1, ignoring remaining characters")
	// }

	header, err = named["header"].GetBoolScalar()
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
func PreludioFunc_From(funcName string, vm *PreludioVM) {
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

	df, err = positional[0].GetDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, df.Error())))
		return
	}

	vm.StackPush(NewPreludioInternalTerm(df))
	vm.SetCurrentDataFrame()
}

// Import a Dataframe form CSV file
func PreludioFunc_ImportCsv(funcName string, vm *PreludioVM) {
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

	path, err = positional[0].GetStringScalar()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	inputFile, err = os.Open(path)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	delimiter, err = named["delimiter"].GetStringScalar()
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
func PreludioFunc_New(funcName string, vm *PreludioVM) {
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

	list, err = positional[0].GetList()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	s := make([]series.Series, len(list))
	for i, e := range list {
		if l, ok := e.GetValue().(PreludioList); ok {

			switch l[0].GetValue().(type) {
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

func PreludioFunc_Select(funcName string, vm *PreludioVM) {
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

	df, err = positional[0].GetDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// The first value can be both a symbol or la list of symbols
	symbol, err = positional[1].GetSymbol()
	if err != nil {
		list, err = positional[1].GetList()
		if err != nil {
			vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
			return
		}

		names := make([]string, len(list))
		for i, v := range list {
			symbol, err = v.GetSymbol()
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

func PreludioFunc_Sort(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

func PreludioFunc_Take(funcName string, vm *PreludioVM) {
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

	df, err = positional[0].GetDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	num, err = positional[1].GetIntegerScalar()
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
///////						COERCION FUNCTIONS

func PreludioFunc_ToFloat(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	var df dataframe.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].GetDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	vm.StackPush(NewPreludioInternalTerm(df))
}

///////////////////////////////////////////////////////////////////////////////
///////						STRING FUNCTIONS

func PreludioFunc_StrReplace(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	var df dataframe.DataFrame
	positional, _, err := vm.GetFunctionParams(funcName, nil, false)
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].GetDataframe()
	if err != nil {
		vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	vm.StackPush(NewPreludioInternalTerm(df))
}
