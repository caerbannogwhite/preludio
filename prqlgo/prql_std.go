package main

import (
	"fmt"
)

type PrqlFunction func(vm *PrqlVirtualMachine)

func PrqlFunc_Derive(vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling PrqlFunc_Derive")
	}
}

func PrqlFunc_ExportCsv(vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling PrqlFunc_ExportCsv")
	}
}

func PrqlFunc_From(vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling PrqlFunc_From")
	}
}

func PrqlFunc_ImportCsv(vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling PrqlFunc_ImportCsv")
	}

	// dataframe.ReadCSV()
}

func PrqlFunc_Select(vm *PrqlVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "", "", "", "Calling PrqlFunc_Select")
	}
}
