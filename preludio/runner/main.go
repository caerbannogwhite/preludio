package main

import (
	"compiler"
	"preludio"
)

func main() {

	code := `
importCSV "C:\\Users\\massi\\Downloads\\Cars.csv" delimiter: ";" header:true
derive [
  num = Cylinders * 2 - 1.123e-1,
  Car_Origin = Car + " - " + Origin
]
take 20
select [num, Car_Origin, MPG, Cylinders]
describe
exportCSV "C:\\Users\\massi\\Downloads\\Cars1.csv" delimiter: '\t'
`

	// inputStream := antlr.NewInputStream(code)

	// lexer := compiler.NewpreludioLexer(inputStream)
	// tokStream := antlr.NewCommonTokenStream(lexer, 0)
	// parser := compiler.NewpreludioParser(tokStream)

	// interp := parser.GetInterpreter()
	// for {
	// 	tok := tokStream.LT(1)
	// 	fmt.Println(tok.GetTokenType(), tok.GetText())
	// 	if tok.GetTokenType() == antlr.TokenEOF {
	// 		break
	// 	}
	// 	tokStream.Consume()
	// }

	warnings := true
	debugLevel := 20

	be := new(preludio.ByteEater).
		InitVM().
		SetPrintWarning(warnings).
		SetDebugLevel(debugLevel)

	// bytecode := compiler.Compile("C:\\Users\\massi\\source\\repos\\preludio\\tests\\00_test_min.prql")
	bytecode := compiler.CompileSource(code)
	be.RunBytecode(bytecode)

	// be.PrintLog()
}
