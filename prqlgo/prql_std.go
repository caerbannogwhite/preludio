package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

type PrqlFunction func(funcName string, vm *PrqlVirtualMachine)

func PrqlFunc_Derive(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

}

func PrqlFunc_Describe(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

}

func PrqlFunc_ExportCsv(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PrqlExpr{
		// "delimiter": NewPrqlExpr(","),
		"header": NewPrqlExpr(true),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, 1, 1, &named, false)
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	var header bool
	var path string
	var df dataframe.DataFrame
	var outputFile *os.File

	df, err = positional[0].GetValueDataframe()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	path, err = positional[1].GetValueString()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	outputFile, err = os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	// delimiter, err = named["delimiter"].GetValueString()
	// if err != nil {
	// 	vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
	// 	return
	// }

	// if len(delimiter) > 1 {
	// 	vm.PrintWarning("delimiter length greater than 1, ignoring remaining characters")
	// }

	header, err = named["header"].GetValueBool()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	err = df.WriteCSV(outputFile, dataframe.WriteHeader(header))
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, df.Error())))
		return
	}
}

func PrqlFunc_From(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PrqlExpr{}

	var err error
	var df dataframe.DataFrame

	positional, _, err := vm.GetFunctionParams(funcName, 0, 1, &named, false)
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].GetValueDataframe()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, df.Error())))
		return
	}

	vm.StackPush(NewPrqlInternalTerm(df))
}

func PrqlFunc_ImportCsv(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PrqlExpr{
		"delimiter": NewPrqlExpr(","),
		"header":    NewPrqlExpr(true),
	}

	var err error
	positional, _, err := vm.GetFunctionParams(funcName, 0, 1, &named, false)
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	var path, delimiter string
	var inputFile *os.File

	path, err = positional[0].GetValueString()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	inputFile, err = os.Open(path)
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	delimiter, err = named["delimiter"].GetValueString()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	if len(delimiter) > 1 {
		vm.PrintWarning("delimiter length greater than 1, ignoring remaining characters")
	}

	df := dataframe.ReadCSV(inputFile, dataframe.WithDelimiter(rune(delimiter[0])))
	if df.Error() != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, df.Error())))
		return
	}

	vm.StackPush(NewPrqlInternalTerm(df))
}

func PrqlFunc_New(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	// named := map[string]*PrqlExpr{}

	// var err error
	// _, assignements, err := vm.GetFunctionParams(funcName, 0, &named, false, true)
	// if err != nil {
	// 	vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
	// 	return
	// }

	// ser := make([]series.Series, len(assignements))
	// idx := 0
	// for k, v := range assignements {
	// 	ser[idx] = series.New(v, series.Bool, k)
	// 	idx++
	// }
	// df := dataframe.New(ser...)

	// vm.StackPush(NewPrqlInternalTerm(df))
}

func PrqlFunc_Select(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

func PrqlFunc_Sort(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

func PrqlFunc_Take(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PrqlExpr{}

	var err error
	var df dataframe.DataFrame
	var num int64
	positional, _, err := vm.GetFunctionParams(funcName, 1, 1, &named, false)
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	df, err = positional[0].GetValueDataframe()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	num, err = positional[1].GetValueInteger()
	if err != nil {
		vm.StackPush(NewPrqlInternalError(fmt.Sprintf("function %s: %s", funcName, err)))
		return
	}

	rows := make([]int, num)
	for i, _ := range rows {
		rows[i] = i
	}

	df = df.Subset(rows)

	vm.StackPush(NewPrqlInternalTerm(df))
}
