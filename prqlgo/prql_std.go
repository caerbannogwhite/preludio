package main

import "fmt"

type PRQLFunction func(vm *PRQLVirtualMachine)

func prql_derive(vm *PRQLVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "", "", "", "Calling prql_derive")
	}
}

func prql_export_csv(vm *PRQLVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "", "", "", "Calling prql_export_csv")
	}
}

func prql_from(vm *PRQLVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "", "", "", "Calling prql_from")
	}
}

func prql_import_csv(vm *PRQLVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "", "", "", "Calling prql_import_csv")
	}
}

func prql_select(vm *PRQLVirtualMachine) {
	if vm.__debugLevel > 5 {
		fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "", "", "", "Calling prql_select")
	}
}
