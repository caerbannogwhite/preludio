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

}

func PreludioFunc_Describe(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	var df, tmpDf dataframe.DataFrame
	var symbol PreludioSymbol
	var list PreludioList
	positional, _, err := vm.GetFunctionParams(funcName, 1, 1, nil, false)
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

		var names []string
		if len(list) == 0 {
			names = df.Names()
		} else {
			names = make([]string, len(list))
			for i, v := range list {
				symbol, err = v.GetSymbol()
				if err != nil {
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
					return
				}
				names[i] = string(symbol)
			}
		}

		tmpDf = df.Select(names)
		fmt.Println(tmpDf.Describe())
	} else {
		tmpDf = df.Select([]string{string(symbol)})
		fmt.Println(tmpDf.Describe())
	}

	vm.StackPush(NewPreludioInternalTerm(df))
}

func PreludioFunc_ExportCsv(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PreludioInternal{
		// "delimiter": NewPreludioInternalTerm([]string{","}),
		"header": NewPreludioInternalTerm([]bool{true}),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, 1, 1, &named, false)
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

func PreludioFunc_From(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PreludioInternal{}

	var err error
	var df dataframe.DataFrame

	positional, _, err := vm.GetFunctionParams(funcName, 0, 1, &named, false)
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
}

func PreludioFunc_ImportCsv(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PreludioInternal{
		"delimiter": NewPreludioInternalTerm([]string{","}),
		"header":    NewPreludioInternalTerm([]bool{true}),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, 0, 1, &named, false)
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
}

func PreludioFunc_New(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var list PreludioList
	var err error
	positional, _, err := vm.GetFunctionParams(funcName, 0, 1, nil, false)
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
}

func PreludioFunc_Select(funcName string, vm *PreludioVM) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	var err error
	var df dataframe.DataFrame
	var symbol PreludioSymbol
	var list PreludioList
	positional, _, err := vm.GetFunctionParams(funcName, 1, 1, nil, false)
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
	positional, _, err := vm.GetFunctionParams(funcName, 1, 1, nil, false)
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
