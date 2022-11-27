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

func PrqlFunc_ExportCsv(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

func PrqlFunc_From(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}

func PrqlFunc_ImportCsv(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}

	named := map[string]*PrqlExpr{
		"delimiter": NewPrqlExpr(","),
		"header":    NewPrqlExpr(true),
	}
	positional, _ := vm.GetFunctionParams(funcName, 1, &named, false)

	var path, delimiter string
	var inputFile *os.File
	var err error

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

func PrqlFunc_Select(funcName string, vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling "+funcName)
	}
}
