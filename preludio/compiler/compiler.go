package compiler

import (
	"encoding/binary"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type OPCODE uint8
type PARAM1 uint8

const (
	TERM_NULL     PARAM1 = 0
	TERM_BOOL     PARAM1 = 1
	TERM_INTEGER  PARAM1 = 2
	TERM_FLOAT    PARAM1 = 3
	TERM_STRING   PARAM1 = 4
	TERM_INTERVAL PARAM1 = 5
	TERM_RANGE    PARAM1 = 6
	TERM_LIST     PARAM1 = 7
	TERM_PIPELINE PARAM1 = 8
	TERM_SYMBOL   PARAM1 = 10
)

const (
	OP_START_PIPELINE    OPCODE = 0
	OP_END_PIPELINE      OPCODE = 1
	OP_ASSIGN_STMT       OPCODE = 2
	OP_START_FUNC_CALL   OPCODE = 3
	OP_MAKE_FUNC_CALL    OPCODE = 4
	OP_START_LIST        OPCODE = 5
	OP_END_LIST          OPCODE = 6
	OP_ADD_FUNC_PARAM    OPCODE = 7
	OP_ADD_EXPR_TERM     OPCODE = 8
	OP_PUSH_NAMED_PARAM  OPCODE = 9
	OP_PUSH_ASSIGN_IDENT OPCODE = 10
	OP_PUSH_TERM         OPCODE = 11
	OP_END_CHUNCK        OPCODE = 12
	OP_GOTO              OPCODE = 50

	OP_BINARY_MUL OPCODE = 100
	OP_BINARY_DIV OPCODE = 101
	OP_BINARY_MOD OPCODE = 102
	OP_BINARY_ADD OPCODE = 103
	OP_BINARY_SUB OPCODE = 104
	OP_BINARY_POW OPCODE = 105

	OP_BINARY_EQ OPCODE = 110
	OP_BINARY_NE OPCODE = 111
	OP_BINARY_GE OPCODE = 112
	OP_BINARY_LE OPCODE = 113
	OP_BINARY_GT OPCODE = 114
	OP_BINARY_LT OPCODE = 115

	OP_BINARY_AND      OPCODE = 120
	OP_BINARY_OR       OPCODE = 121
	OP_BINARY_COALESCE OPCODE = 122
	OP_BINARY_MODEL    OPCODE = 123

	OP_UNARY_ADD OPCODE = 130
	OP_UNARY_SUB OPCODE = 131
	OP_UNARY_NOT OPCODE = 132
)

func Compile(source string) []byte {
	input, _ := antlr.NewFileStream(source)
	lexer := NewpreludioLexer(input)
	tokStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewpreludioParser(tokStream)

	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Program()

	compiler := new(ByteFeeder)
	antlr.ParseTreeWalkerDefault.Walk(compiler, tree)

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
	*BasepreludioListener

	symbolTable  SymbolTable
	instructions []byte
}

func (bf *ByteFeeder) Init() *ByteFeeder {
	bf.instructions = make([]byte, 0)
	bf.symbolTable = make([]string, 0)

	return bf
}

func (bf *ByteFeeder) AppendInstruction(opcode OPCODE, param1 PARAM1, param2 int) {
	bf.instructions = append(bf.instructions, byte(opcode))
	bf.instructions = append(bf.instructions, byte(param1))
	bf.instructions = binary.BigEndian.AppendUint32(bf.instructions, uint32(param2))
}

func (bf *ByteFeeder) GetBytecode() []byte {

	// incipit: 4 bytes mark, 4 empty bytes
	// and the number of elements in the symbol table
	bytecode := []byte{
		0x11, 0x01, 0x19, 0x93,
		0x00, 0x00, 0x00, 0x00,
	}

	bytecode = binary.BigEndian.AppendUint64(bytecode, uint64(len(bf.symbolTable)))

	for _, symbol := range bf.symbolTable {
		enc := []byte(symbol)
		bytecode = binary.BigEndian.AppendUint64(bytecode, uint64(len(enc)))
		bytecode = append(bytecode, enc...)
	}

	return append(bytecode, bf.instructions...)
}

func (bf *ByteFeeder) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (bf *ByteFeeder) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (bf *ByteFeeder) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (bf *ByteFeeder) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNl is called when production nl is entered.
// func (bf *ByteFeeder) EnterNl(ctx *NlContext) {}

// // ExitNl is called when production nl is exited.
// func (bf *ByteFeeder) ExitNl(ctx *NlContext) {}

// // EnterProgram is called when production program is entered.
// func (bf *ByteFeeder) EnterProgram(ctx *ProgramContext) {}

// // ExitProgram is called when production program is exited.
// func (bf *ByteFeeder) ExitProgram(ctx *ProgramContext) {}

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
// func (bf *ByteFeeder) EnterStmt(ctx *StmtContext) {}

// // ExitStmt is called when production stmt is exited.
// func (bf *ByteFeeder) ExitStmt(ctx *StmtContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (bf *ByteFeeder) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (bf *ByteFeeder) ExitAssignStmt(ctx *AssignStmtContext) {
	// identName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(OP_ASSIGN_STMT, 0, pos)
}

// EnterPipeline is called when production pipeline is entered.
func (bf *ByteFeeder) EnterPipeline(ctx *PipelineContext) {
	bf.AppendInstruction(OP_START_PIPELINE, 0, 0)
}

// ExitPipeline is called when production pipeline is exited.
func (bf *ByteFeeder) ExitPipeline(ctx *PipelineContext) {
	bf.AppendInstruction(OP_END_PIPELINE, 0, 0)
}

// EnterInlinePipeline is called when production inlinePipeline is entered.
func (bf *ByteFeeder) EnterInlinePipeline(ctx *InlinePipelineContext) {
	bf.AppendInstruction(OP_START_PIPELINE, 0, 0)
}

// ExitInlinePipeline is called when production inlinePipeline is exited.
func (bf *ByteFeeder) ExitInlinePipeline(ctx *InlinePipelineContext) {
	bf.AppendInstruction(OP_END_PIPELINE, 0, 0)
}

// EnterIdentBacktick is called when production identBacktick is entered.
func (bf *ByteFeeder) EnterIdentBacktick(ctx *IdentBacktickContext) {}

// ExitIdentBacktick is called when production identBacktick is exited.
func (bf *ByteFeeder) ExitIdentBacktick(ctx *IdentBacktickContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (bf *ByteFeeder) EnterFuncCall(ctx *FuncCallContext) {
	bf.AppendInstruction(OP_START_FUNC_CALL, 0, 0)
}

// ExitFuncCall is called when production funcCall is exited.
func (bf *ByteFeeder) ExitFuncCall(ctx *FuncCallContext) {
	// funcName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(OP_MAKE_FUNC_CALL, 0, pos)
}

// EnterFuncCallParam is called when production funcCallParam is entered.
func (bf *ByteFeeder) EnterFuncCallParam(ctx *FuncCallParamContext) {}

// ExitFuncCallParam is called when production funcCallParam is exited.
func (bf *ByteFeeder) ExitFuncCallParam(ctx *FuncCallParamContext) {
	bf.AppendInstruction(OP_END_CHUNCK, 0, 0)
}

// EnterNamedArg is called when production namedArg is entered.
func (bf *ByteFeeder) EnterNamedArg(ctx *NamedArgContext) {}

// ExitNamedArg is called when production namedArg is exited.
func (bf *ByteFeeder) ExitNamedArg(ctx *NamedArgContext) {
	// paramName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(OP_PUSH_NAMED_PARAM, 0, pos)
}

// EnterAssign is called when production assign is entered.
func (bf *ByteFeeder) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (bf *ByteFeeder) ExitAssign(ctx *AssignContext) {
	// identName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.Add(ctx.IDENT().GetSymbol().GetText())
	bf.AppendInstruction(OP_PUSH_ASSIGN_IDENT, 0, pos)
}

// EnterMultiAssign is called when production multiAssign is entered.
func (bf *ByteFeeder) EnterMultiAssign(ctx *MultiAssignContext) {}

// ExitMultiAssign is called when production multiAssign is exited.
func (bf *ByteFeeder) ExitMultiAssign(ctx *MultiAssignContext) {}

// EnterExprCall is called when production exprCall is entered.
func (bf *ByteFeeder) EnterExprCall(ctx *ExprCallContext) {
	bf.AppendInstruction(OP_END_CHUNCK, 0, 0)
}

// ExitExprCall is called when production exprCall is exited.
func (bf *ByteFeeder) ExitExprCall(ctx *ExprCallContext) {}

// EnterExpr is called when production expr is entered.
func (bf *ByteFeeder) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (bf *ByteFeeder) ExitExpr(ctx *ExprContext) {
	// operation or nested expression
	if ctx.GetChildCount() == 3 {
		if ctx.GetChild(0).GetPayload().(*antlr.CommonToken).GetText() == "(" {
			// console.log(ctx.children[0].symbol.text);
		} else {
			switch ctx.GetChild(1).GetPayload().(*antlr.CommonToken).GetText() {
			case "*":
				bf.AppendInstruction(OP_BINARY_MUL, 0, 0)

			case "/":
				bf.AppendInstruction(OP_BINARY_DIV, 0, 0)

			case "%":
				bf.AppendInstruction(OP_BINARY_MOD, 0, 0)

			case "+":
				bf.AppendInstruction(OP_BINARY_ADD, 0, 0)

			case "-":
				bf.AppendInstruction(OP_BINARY_SUB, 0, 0)

			case "==":
				bf.AppendInstruction(OP_BINARY_EQ, 0, 0)

			case "!=":
				bf.AppendInstruction(OP_BINARY_NE, 0, 0)

			case ">=":
				bf.AppendInstruction(OP_BINARY_GE, 0, 0)

			case "<=":
				bf.AppendInstruction(OP_BINARY_LE, 0, 0)

			case ">":
				bf.AppendInstruction(OP_BINARY_GT, 0, 0)

			case "<":
				bf.AppendInstruction(OP_BINARY_LT, 0, 0)

			case "**":
				bf.AppendInstruction(OP_BINARY_POW, 0, 0)

			case "~":
				bf.AppendInstruction(OP_BINARY_MODEL, 0, 0)

			}
		}
	}
}

// EnterTerm is called when production term is entered.
func (bf *ByteFeeder) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (bf *ByteFeeder) ExitTerm(ctx *TermContext) {}

// EnterExprUnary is called when production exprUnary is entered.
func (bf *ByteFeeder) EnterExprUnary(ctx *ExprUnaryContext) {}

// ExitExprUnary is called when production exprUnary is exited.
func (bf *ByteFeeder) ExitExprUnary(ctx *ExprUnaryContext) {
	if ctx.GetChildCount() == 2 {
		switch ctx.GetChild(0).GetPayload().(*antlr.CommonToken).GetText() {
		case "-":
			bf.AppendInstruction(OP_UNARY_SUB, 0, 0)

		case "+":
			bf.AppendInstruction(OP_UNARY_ADD, 0, 0)

		case "not":
			bf.AppendInstruction(OP_UNARY_NOT, 0, 0)

		}
	}
}

// EnterLiteral is called when production literal is entered.
func (bf *ByteFeeder) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (bf *ByteFeeder) ExitLiteral(ctx *LiteralContext) {
	switch ctx.GetChildCount() {
	case 1:
		// NULL
		if ctx.NULL_() == nil {
			bf.AppendInstruction(OP_PUSH_TERM, TERM_NULL, 0)
		} else

		// BOOLEAN
		if ctx.BOOLEAN() != nil {
			if ctx.BOOLEAN().GetSymbol().GetText() == "true" {
				bf.AppendInstruction(OP_PUSH_TERM, TERM_BOOL, 1)
			} else {
				bf.AppendInstruction(OP_PUSH_TERM, TERM_BOOL, 0)
			}
		} else

		// INTEGER
		if ctx.INTEGER(0) != nil {
			num := ctx.INTEGER(0).GetSymbol().GetText()
			pos := bf.symbolTable.Add(num)
			bf.AppendInstruction(OP_PUSH_TERM, TERM_INTEGER, pos)
		} else

		// // FLOAT
		if ctx.FLOAT(0) != nil {
			num := ctx.FLOAT(0).GetSymbol().GetText()
			pos := bf.symbolTable.Add(num)
			bf.AppendInstruction(OP_PUSH_TERM, TERM_FLOAT, pos)
		} else

		// STRING
		if ctx.STRING() != nil {
			replacer := strings.NewReplacer("\"", "", "'", "")
			str := replacer.Replace(ctx.STRING().GetText())
			pos := bf.symbolTable.Add(str)
			bf.AppendInstruction(OP_PUSH_TERM, TERM_STRING, pos)
		}

		// IDENT
		if ctx.IDENT(0) != nil {
			id := ctx.IDENT(0).GetSymbol().GetText()
			pos := bf.symbolTable.Add(id)
			bf.AppendInstruction(OP_PUSH_TERM, TERM_SYMBOL, pos)
		}

	// time interval
	case 2:
		// this.term = {
		//   type: TERM_INTERVAL,
		//   num: parseFloat(ctx.children[0].getText()),
		//   kind: ctx.children[1].getText(),
		// };

	// range
	case 3:
		// const s = ctx.children[0].getText();
		// if (s === NaN) {
		//   const start = { type: TERM_SYMBOL, value: ctx.children[0].getText() };
		// } else {
		//   const start = { type: TERM_FLOAT, value: s };
		// }

		// const e = ctx.children[2].getText();
		// if (end === NaN) {
		//   const end = { type: TERM_SYMBOL, value: ctx.children[2].getText() };
		// } else {
		//   const end = { type: TERM_FLOAT, value: e };
		// }

		// this.term = {
		//   type: TERM_RANGE,
		//   start: start,
		//   end: end,
		// };

	}
}

// EnterList is called when production list is entered.
func (bf *ByteFeeder) EnterList(ctx *ListContext) {
	bf.AppendInstruction(OP_START_LIST, 0, 0)
}

// ExitList is called when production list is exited.
func (bf *ByteFeeder) ExitList(ctx *ListContext) {
	bf.AppendInstruction(OP_END_LIST, 0, 0)
}

// EnterNestedPipeline is called when production nestedPipeline is entered.
// func (bf *ByteFeeder) EnterNestedPipeline(ctx *NestedPipelineContext) {}

// ExitNestedPipeline is called when production nestedPipeline is exited.
// func (bf *ByteFeeder) ExitNestedPipeline(ctx *NestedPipelineContext) {}
