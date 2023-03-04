package main

import (
	"compiler"
	"preludio"
)

func main() {

	warnings := true
	debugLevel := 20

	be := new(preludio.ByteEater).
		InitVM().
		SetPrintWarning(warnings).
		SetDebugLevel(debugLevel)

	bytecode := compiler.Compile("C:\\Users\\massi\\source\\repos\\preludio\\tests\\00_test_min.prql")
	be.ReadBytecode(bytecode)

	be.PrintLog()
}
