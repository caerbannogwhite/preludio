package main

import (
	"compiler"
	"fmt"
	"preludio"
)

func main() {

	code := `
readCSV "C:\\Users\\massi\\Downloads\\Cars.csv" delimiter: ";"
take 20
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

	bytecode := compiler.CompileSource(code)
	be.RunBytecode(bytecode)

	res := be.GetOutput()

	for _, log := range res.Log {
		fmt.Println(log.Message)
	}

	fmt.Println(res.Data)
}
