package bytefeeder

import (
	"encoding/binary"
	"fmt"
	"preludiometa"
	"strings"

	antlr "github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func CompileSource(source string) ([]byte, []preludiometa.LogEnty, error) {
	inputStream := antlr.NewInputStream(source)

	lexer := NewpreludioLexer(inputStream)
	tokStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewpreludioParser(tokStream)

	compiler := new(ByteFeeder)
	antlr.ParseTreeWalkerDefault.Walk(compiler, parser.Program())

	return compiler.GetBytecode()
}

func CompileFile(path string) ([]byte, []preludiometa.LogEnty, error) {
	fileStream, err := antlr.NewFileStream(path)
	if err != nil {
		panic(err)
	}

	lexer := NewpreludioLexer(fileStream)
	tokStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewpreludioParser(tokStream)

	compiler := new(ByteFeeder)
	antlr.ParseTreeWalkerDefault.Walk(compiler, parser.Program())

	return compiler.GetBytecode()
}

type SymbolTable []string

func (t *SymbolTable) Add(symbol string) int {
	pos := -1
	for idx, s := range *t {
		if s == symbol {
			pos = idx
			break
		}
	}

	if pos == -1 {
		pos = len(*t)
		*t = append(*t, symbol)
	}

	return pos
}

func (t *SymbolTable) IndexOf(symbol string) int {
	for idx, s := range *t {
		if s == symbol {
			return idx
		}
	}
	return -1
}

type ByteFeeder struct {
	*BasepreludioParserListener

	verbose         bool
	err             error
	lastInstruction preludiometa.OPCODE

	symbolTable  SymbolTable
	instructions []byte
	logs         []preludiometa.LogEnty
}

func (bf *ByteFeeder) Init() *ByteFeeder {
	bf.instructions = make([]byte, 0)
	bf.symbolTable = make([]string, 0)
	bf.logs = make([]preludiometa.LogEnty, 0)

	return bf
}

func (bf *ByteFeeder) SetVerbose(flag bool) *ByteFeeder {
	bf.verbose = flag
	return bf
}

func (bf *ByteFeeder) AppendInstruction(opcode preludiometa.OPCODE, param1 preludiometa.PARAM1, param2 int) {
	bf.lastInstruction = opcode

	bf.instructions = append(bf.instructions, byte(opcode))
	bf.instructions = append(bf.instructions, byte(param1))
	bf.instructions = binary.BigEndian.AppendUint32(bf.instructions, uint32(param2))
}

func (bf *ByteFeeder) RemoveLastInstruction() {
	bf.instructions = bf.instructions[:len(bf.instructions)-preludiometa.PRELUDIO_INSTRUCTION_SIZE]
	bf.lastInstruction = preludiometa.OPCODE(bf.instructions[len(bf.instructions)-preludiometa.PRELUDIO_INSTRUCTION_SIZE])
}

func (bf *ByteFeeder) GetBytecode() ([]byte, []preludiometa.LogEnty, error) {

	// incipit: 4 bytes mark, 4 empty bytes
	// and the number of elements in the symbol table
	bytecode := []byte{
		0x11, 0x01, 0x19, 0x93,
	}

	bytecode = binary.BigEndian.AppendUint32(bytecode, uint32(len(bf.symbolTable)))

	for _, symbol := range bf.symbolTable {
		enc := []byte(symbol)
		bytecode = binary.BigEndian.AppendUint32(bytecode, uint32(len(enc)))
		bytecode = append(bytecode, enc...)
	}

	if bf.err == nil {
		bf.logs = append(bf.logs, preludiometa.LogEnty{
			LogType: preludiometa.LOG_DEBUG,
			Level:   5,
			Message: "Bytecode generated successfully"})
	}

	return append(bytecode, bf.instructions...), bf.logs, bf.err
}

func (bf *ByteFeeder) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (bf *ByteFeeder) VisitErrorNode(node antlr.ErrorNode) {
	bf.err = fmt.Errorf("error node: %s", node.GetText())
}

// EnterEveryRule is called when any rule is entered.
func (bf *ByteFeeder) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (bf *ByteFeeder) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNl is called when production nl is entered.
// func (bf *ByteFeeder) EnterNl(ctx *NlContext) {}

// // ExitNl is called when production nl is exited.
// func (bf *ByteFeeder) ExitNl(ctx *NlContext) {}

// EnterProgram is called when production program is entered.
func (bf *ByteFeeder) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (bf *ByteFeeder) ExitProgram(ctx *ProgramContext) {}

// // EnterProgramIntro is called when production programIntro is entered.
// func (bf *ByteFeeder) EnterProgramIntro(ctx *ProgramIntroContext) {}

// // ExitProgramIntro is called when production programIntro is exited.
// func (bf *ByteFeeder) ExitProgramIntro(ctx *ProgramIntroContext) {}

// // EnterFuncDef is called when production funcDef is entered.
// func (bf *ByteFeeder) EnterFuncDef(ctx *FuncDefContext) {}

// // ExitFuncDef is called when production funcDef is exited.
// func (bf *ByteFeeder) ExitFuncDef(ctx *FuncDefContext) {}

// // EnterFuncDefName is called when production funcDefName is entered.
// func (bf *ByteFeeder) EnterFuncDefName(ctx *FuncDefNameContext) {}

// // ExitFuncDefName is called when production funcDefName is exited.
// func (bf *ByteFeeder) ExitFuncDefName(ctx *FuncDefNameContext) {}

// // EnterFuncDefParams is called when production funcDefParams is entered.
// func (bf *ByteFeeder) EnterFuncDefParams(ctx *FuncDefParamsContext) {}

// // ExitFuncDefParams is called when production funcDefParams is exited.
// func (bf *ByteFeeder) ExitFuncDefParams(ctx *FuncDefParamsContext) {}

// // EnterFuncDefParam is called when production funcDefParam is entered.
// func (bf *ByteFeeder) EnterFuncDefParam(ctx *FuncDefParamContext) {}

// // ExitFuncDefParam is called when production funcDefParam is exited.
// func (bf *ByteFeeder) ExitFuncDefParam(ctx *FuncDefParamContext) {}

// // EnterTypeDef is called when production typeDef is entered.
// func (bf *ByteFeeder) EnterTypeDef(ctx *TypeDefContext) {}

// // ExitTypeDef is called when production typeDef is exited.
// func (bf *ByteFeeder) ExitTypeDef(ctx *TypeDefContext) {}

// // EnterTypeTerm is called when production typeTerm is entered.
// func (bf *ByteFeeder) EnterTypeTerm(ctx *TypeTermContext) {}

// // ExitTypeTerm is called when production typeTerm is exited.
// func (bf *ByteFeeder) ExitTypeTerm(ctx *TypeTermContext) {}

// // EnterStmt is called when production stmt is entered.
func (bf *ByteFeeder) EnterStmt(ctx *StmtContext) {
	bf.AppendInstruction(preludiometa.OP_START_STMT, 0, 0)
}

// // ExitStmt is called when production stmt is exited.
func (bf *ByteFeeder) ExitStmt(ctx *StmtContext) {
	bf.AppendInstruction(preludiometa.OP_END_STMT, 0, 0)
}

// EnterVarAssignStmt is called when production assignStmt is entered.
func (bf *ByteFeeder) EnterVarAssignStmt(ctx *VarAssignStmtContext) {}

// ExitVarAssignStmt is called when production assignStmt is exited.
func (bf *ByteFeeder) ExitVarAssignStmt(ctx *VarAssignStmtContext) {
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(preludiometa.OP_VAR_ASSIGN, 0, pos)
}

// EnterVarDefStmt is called when production varDeclStmt is entered.
func (bf *ByteFeeder) EnterVarDeclStmt(ctx *VarDeclStmtContext) {}

// ExitVarDeclStmt is called when production varDeclStmt is exited.
func (bf *ByteFeeder) ExitVarDeclStmt(ctx *VarDeclStmtContext) {
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(preludiometa.OP_VAR_DECL, 0, pos)
}

// EnterIfElseStmt is called when production ifElseStmt is entered.
func (bf *ByteFeeder) EnterIfElseStmt(ctx *IfElseStmtContext) {}

// ExitIfElseStmt is called when production ifElseStmt is exited.
func (bf *ByteFeeder) ExitIfElseStmt(ctx *IfElseStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (bf *ByteFeeder) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (bf *ByteFeeder) ExitForStmt(ctx *ForStmtContext) {}

// EnterHelpStmt is called when production helpStmt is entered.
func (bf *ByteFeeder) EnterHelpStmt(ctx *HelpStmtContext) {}

// ExitHelpStmt is called when production helpStmt is exited.
func (bf *ByteFeeder) ExitHelpStmt(ctx *HelpStmtContext) {
	bf.AppendInstruction(preludiometa.OP_HELP, 0, bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText()))
}

// EnterPipeline is called when production pipeline is entered.
func (bf *ByteFeeder) EnterPipeline(ctx *PipelineContext) {
	bf.AppendInstruction(preludiometa.OP_START_PIPELINE, 0, 0)
}

// ExitPipeline is called when production pipeline is exited.
func (bf *ByteFeeder) ExitPipeline(ctx *PipelineContext) {
	bf.AppendInstruction(preludiometa.OP_END_PIPELINE, 0, 0)
}

// EnterInlinePipeline is called when production inlinePipeline is entered.
func (bf *ByteFeeder) EnterInlinePipeline(ctx *InlinePipelineContext) {
	bf.AppendInstruction(preludiometa.OP_START_PIPELINE, 0, 0)
}

// ExitInlinePipeline is called when production inlinePipeline is exited.
func (bf *ByteFeeder) ExitInlinePipeline(ctx *InlinePipelineContext) {
	bf.AppendInstruction(preludiometa.OP_END_PIPELINE, 0, 0)
}

// EnterFuncCall is called when production funcCall is entered.
func (bf *ByteFeeder) EnterFuncCall(ctx *FuncCallContext) {
	bf.AppendInstruction(preludiometa.OP_START_FUNC_CALL, 0, 0)
}

// ExitFuncCall is called when production funcCall is exited.
func (bf *ByteFeeder) ExitFuncCall(ctx *FuncCallContext) {
	// funcName := ctx.IDENT().GetSymbol().GetText()
	bf.AppendInstruction(preludiometa.OP_MAKE_FUNC_CALL, 0, bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText()))
}

// EnterFuncCallParam is called when production funcCallParam is entered.
func (bf *ByteFeeder) EnterFuncCallParam(ctx *FuncCallParamContext) {}

// ExitFuncCallParam is called when production funcCallParam is exited.
func (bf *ByteFeeder) ExitFuncCallParam(ctx *FuncCallParamContext) {
	bf.AppendInstruction(preludiometa.OP_END_CHUNCK, 0, 0)
}

// EnterNamedArg is called when production namedArg is entered.
func (bf *ByteFeeder) EnterNamedArg(ctx *NamedArgContext) {}

// ExitNamedArg is called when production namedArg is exited.
func (bf *ByteFeeder) ExitNamedArg(ctx *NamedArgContext) {
	// paramName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(preludiometa.OP_PUSH_NAMED_PARAM, 0, pos)
}

// EnterAssign is called when production assign is entered.
func (bf *ByteFeeder) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (bf *ByteFeeder) ExitAssign(ctx *AssignContext) {
	// identName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(preludiometa.OP_PUSH_ASSIGN_IDENT, 0, pos)
}

// EnterMultiAssign is called when production multiAssign is entered.
func (bf *ByteFeeder) EnterMultiAssign(ctx *MultiAssignContext) {}

// ExitMultiAssign is called when production multiAssign is exited.
func (bf *ByteFeeder) ExitMultiAssign(ctx *MultiAssignContext) {}

// EnterExpr is called when production expr is entered.
func (bf *ByteFeeder) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (bf *ByteFeeder) ExitExpr(ctx *ExprContext) {
	// operation or nested expression
	switch ctx.GetChildCount() {

	// UNARY OPERATIONS
	case 2:
		switch ctx.GetChild(0).GetPayload().(*antlr.CommonToken).GetText() {
		case preludiometa.SYMBOL_REV:
			bf.AppendInstruction(preludiometa.OP_UNARY_REV, 0, 0)

		case preludiometa.SYMBOL_UNARY_SUB:
			if bf.lastInstruction == preludiometa.OP_UNARY_SUB {
				bf.RemoveLastInstruction()
				return
			}
			bf.AppendInstruction(preludiometa.OP_UNARY_SUB, 0, 0)

		case preludiometa.SYMBOL_UNARY_ADD:
			if bf.lastInstruction == preludiometa.OP_UNARY_ADD {
				return
			}
			bf.AppendInstruction(preludiometa.OP_UNARY_ADD, 0, 0)

		case preludiometa.SYMBOL_UNARY_NOT:
			if bf.lastInstruction == preludiometa.OP_UNARY_NOT {
				bf.RemoveLastInstruction()
				return
			}
			bf.AppendInstruction(preludiometa.OP_UNARY_NOT, 0, 0)
		}

	// BINARY OPERATIONS
	case 3:

		// parenthesis
		if ctx.RPAREN() != nil {
			return
		}

		switch ctx.GetChild(1).GetPayload().(*antlr.CommonToken).GetText() {
		case preludiometa.SYMBOL_INDEXING:
			bf.AppendInstruction(preludiometa.OP_INDEXING, 0, 0)

		case preludiometa.SYMBOL_IN:
			bf.AppendInstruction(preludiometa.OP_BINARY_IN, 0, 0)

		case preludiometa.SYMBOL_BINARY_MUL:
			bf.AppendInstruction(preludiometa.OP_BINARY_MUL, 0, 0)

		case preludiometa.SYMBOL_BINARY_DIV:
			bf.AppendInstruction(preludiometa.OP_BINARY_DIV, 0, 0)

		case preludiometa.SYMBOL_BINARY_MOD:
			bf.AppendInstruction(preludiometa.OP_BINARY_MOD, 0, 0)

		case preludiometa.SYMBOL_BINARY_EXP:
			bf.AppendInstruction(preludiometa.OP_BINARY_EXP, 0, 0)

		case preludiometa.SYMBOL_BINARY_ADD:
			bf.AppendInstruction(preludiometa.OP_BINARY_ADD, 0, 0)

		case preludiometa.SYMBOL_BINARY_SUB:
			bf.AppendInstruction(preludiometa.OP_BINARY_SUB, 0, 0)

		case preludiometa.SYMBOL_BINARY_EQ:
			bf.AppendInstruction(preludiometa.OP_BINARY_EQ, 0, 0)

		case preludiometa.SYMBOL_BINARY_NE:
			bf.AppendInstruction(preludiometa.OP_BINARY_NE, 0, 0)

		case preludiometa.SYMBOL_BINARY_GE:
			bf.AppendInstruction(preludiometa.OP_BINARY_GE, 0, 0)

		case preludiometa.SYMBOL_BINARY_LE:
			bf.AppendInstruction(preludiometa.OP_BINARY_LE, 0, 0)

		case preludiometa.SYMBOL_BINARY_GT:
			bf.AppendInstruction(preludiometa.OP_BINARY_GT, 0, 0)

		case preludiometa.SYMBOL_BINARY_LT:
			bf.AppendInstruction(preludiometa.OP_BINARY_LT, 0, 0)

		case preludiometa.SYMBOL_BINARY_AND:
			bf.AppendInstruction(preludiometa.OP_BINARY_AND, 0, 0)

		case preludiometa.SYMBOL_BINARY_OR:
			bf.AppendInstruction(preludiometa.OP_BINARY_OR, 0, 0)

		case preludiometa.SYMBOL_BINARY_MODEL:
			bf.AppendInstruction(preludiometa.OP_BINARY_MODEL, 0, 0)

		case preludiometa.SYMBOL_BINARY_COALESCE:
			bf.AppendInstruction(preludiometa.OP_BINARY_COALESCE, 0, 0)
		}
	}
}

// EnterLiteral is called when production literal is entered.
func (bf *ByteFeeder) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (bf *ByteFeeder) ExitLiteral(ctx *LiteralContext) {
	// IDENT
	if ctx.IDENT() != nil {
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_SYMBOL,
			bf.symbolTable.Add(
				ctx.IDENT().GetSymbol().GetText()))
	} else

	// NULL
	if ctx.NA() != nil {
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_NULL, 0)
	} else

	// BOOLEAN
	if ctx.BOOLEAN_LIT() != nil {
		if ctx.BOOLEAN_LIT().GetSymbol().GetText() == preludiometa.SYMBOL_TRUE {
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_BOOLEAN, 1)
		} else {
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_BOOLEAN, 0)
		}
	} else

	// INTEGER
	if ctx.INTEGER_LIT() != nil {
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_INTEGER,
			bf.symbolTable.Add(
				ctx.INTEGER_LIT().GetSymbol().GetText()))
	} else

	// RANGE
	if ctx.RANGE_LIT() != nil {
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_RANGE,
			bf.symbolTable.Add(
				ctx.RANGE_LIT().GetSymbol().GetText()))
	} else

	// FLOAT
	if ctx.FLOAT_LIT() != nil {
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_FLOAT,
			bf.symbolTable.Add(
				ctx.FLOAT_LIT().GetSymbol().GetText()))
	} else

	// STRING
	if ctx.STRING_LIT() != nil {
		val := ctx.STRING_LIT().GetText()
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_STRING,
			bf.symbolTable.Add(val[1:len(val)-1]))
	} else

	// STRING INTERP
	if ctx.STRING_INTERP_LIT() != nil {
		var expr, val string
		var stream *antlr.InputStream

		val = ctx.STRING_INTERP_LIT().GetText()
		val = val[2 : len(val)-1]

		strBegin := 0
		exprBegin := strings.Index(val, preludiometa.SYMBOL_LBRACE)
		exprEnd := -1
		exprNum := 0
		for exprBegin != -1 {

			exprEnd = strings.Index(val, preludiometa.SYMBOL_RBRACE)
			if exprEnd == -1 {
				bf.err = fmt.Errorf("invalid string interpolation: %s", val)
				return
			}

			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_STRING,
				bf.symbolTable.Add(val[strBegin:exprBegin]))

			// START Parse expression
			expr = val[exprBegin+1 : exprEnd]
			stream = antlr.NewInputStream(expr)

			antlr.ParseTreeWalkerDefault.Walk(bf, NewpreludioParser(
				antlr.NewCommonTokenStream(NewpreludioLexer(stream), 0)).Expr())
			// END Parse expression

			bf.AppendInstruction(preludiometa.OP_BINARY_ADD, 0, 0)

			val = val[exprEnd+1:]
			exprBegin = strings.Index(val, preludiometa.SYMBOL_LBRACE)
			exprNum++
		}

		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_STRING,
			bf.symbolTable.Add(val))

		// add binary add instruction for each expression
		for i := 0; i < exprNum; i++ {
			bf.AppendInstruction(preludiometa.OP_BINARY_ADD, 0, 0)
		}
	} else

	// STRING RAW
	if ctx.STRING_RAW_LIT() != nil {
		val := ctx.STRING_RAW_LIT().GetText()
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_STRING_RAW,
			bf.symbolTable.Add(val[2:len(val)-1]))
	} else

	// STRING PATH
	if ctx.STRING_PATH_LIT() != nil {
		val := ctx.STRING_PATH_LIT().GetText()
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_STRING_PATH,
			bf.symbolTable.Add(val[2:len(val)-1]))
	} else

	// REGEX
	if ctx.REGEX_LIT() != nil {
		val := ctx.REGEX_LIT().GetText()
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_REGEX,
			bf.symbolTable.Add(val[2:len(val)-1]))
	} else

	// DATE
	if ctx.DATE_LIT() != nil {
		val := ctx.DATE_LIT().GetText()
		bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DATE,
			bf.symbolTable.Add(val[2:len(val)-1]))
	} else

	// DURATION
	if ctx.DURATION_LIT() != nil {
		val := strings.Split(ctx.DURATION_LIT().GetText(), preludiometa.SYMBOL_COLON)
		switch val[1] {
		case preludiometa.SYMBOL_DURATION_MICROSECOND, preludiometa.SYMBOL_DURATION_MICROSECOND_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_MICROSECOND,
				bf.symbolTable.Add(val[0]))

		case preludiometa.SYMBOL_DURATION_MILLISECOND, preludiometa.SYMBOL_DURATION_MILLISECOND_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_MILLISECOND,
				bf.symbolTable.Add(val[0]))

		case preludiometa.SYMBOL_DURATION_SECOND, preludiometa.SYMBOL_DURATION_SECOND_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_SECOND,
				bf.symbolTable.Add(val[0]))

		case preludiometa.SYMBOL_DURATION_MINUTE, preludiometa.SYMBOL_DURATION_MINUTE_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_MINUTE,
				bf.symbolTable.Add(val[0]))

		case preludiometa.SYMBOL_DURATION_HOUR, preludiometa.SYMBOL_DURATION_HOUR_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_HOUR,
				bf.symbolTable.Add(val[0]))

		case preludiometa.SYMBOL_DURATION_DAY, preludiometa.SYMBOL_DURATION_DAY_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_DAY,
				bf.symbolTable.Add(val[0]))

		case preludiometa.SYMBOL_DURATION_MONTH, preludiometa.SYMBOL_DURATION_MONTH_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_MONTH,
				bf.symbolTable.Add(val[0]))

		case preludiometa.SYMBOL_DURATION_YEAR, preludiometa.SYMBOL_DURATION_YEAR_SHORT:
			bf.AppendInstruction(preludiometa.OP_PUSH_TERM, preludiometa.TERM_DURATION_YEAR,
				bf.symbolTable.Add(val[0]))
		}
	}
}

// EnterList is called when production list is entered.
func (bf *ByteFeeder) EnterList(ctx *ListContext) {
	bf.AppendInstruction(preludiometa.OP_START_LIST, 0, 0)
}

// ExitList is called when production list is exited.
func (bf *ByteFeeder) ExitList(ctx *ListContext) {
	bf.AppendInstruction(preludiometa.OP_END_LIST, 0, 0)
}

// EnterNestedPipeline is called when production nestedPipeline is entered.
// func (bf *ByteFeeder) EnterNestedPipeline(ctx *NestedPipelineContext) {}

// ExitNestedPipeline is called when production nestedPipeline is exited.
// func (bf *ByteFeeder) ExitNestedPipeline(ctx *NestedPipelineContext) {}
