package compiler

import (
	"encoding/binary"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type OPCODE uint16
type PARAM1 uint16

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

type Instruction struct {
	opcode OPCODE
	param1 PARAM1
	param2 uint64
}

type SymbolTable []string

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
	instructions []Instruction
}

func (bf *ByteFeeder) Init() *ByteFeeder {
	bf.instructions = make([]Instruction, 0)
	bf.symbolTable = make([]string, 0)

	return bf
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

	// TODO: make this more efficient: allocate an array first
	for _, instruction := range bf.instructions {
		bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(instruction.opcode))
		bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(instruction.param1))
		bytecode = binary.BigEndian.AppendUint64(bytecode, uint64(instruction.param2))
	}

	return bytecode
}

func (bf *ByteFeeder) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (bf *ByteFeeder) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (bf *ByteFeeder) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (bf *ByteFeeder) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNl is called when production nl is entered.
func (bf *ByteFeeder) EnterNl(ctx *NlContext) {}

// ExitNl is called when production nl is exited.
func (bf *ByteFeeder) ExitNl(ctx *NlContext) {}

// EnterProgram is called when production program is entered.
func (bf *ByteFeeder) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (bf *ByteFeeder) ExitProgram(ctx *ProgramContext) {}

// EnterProgramIntro is called when production programIntro is entered.
func (bf *ByteFeeder) EnterProgramIntro(ctx *ProgramIntroContext) {}

// ExitProgramIntro is called when production programIntro is exited.
func (bf *ByteFeeder) ExitProgramIntro(ctx *ProgramIntroContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (bf *ByteFeeder) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (bf *ByteFeeder) ExitFuncDef(ctx *FuncDefContext) {}

// EnterFuncDefName is called when production funcDefName is entered.
func (bf *ByteFeeder) EnterFuncDefName(ctx *FuncDefNameContext) {}

// ExitFuncDefName is called when production funcDefName is exited.
func (bf *ByteFeeder) ExitFuncDefName(ctx *FuncDefNameContext) {}

// EnterFuncDefParams is called when production funcDefParams is entered.
func (bf *ByteFeeder) EnterFuncDefParams(ctx *FuncDefParamsContext) {}

// ExitFuncDefParams is called when production funcDefParams is exited.
func (bf *ByteFeeder) ExitFuncDefParams(ctx *FuncDefParamsContext) {}

// EnterFuncDefParam is called when production funcDefParam is entered.
func (bf *ByteFeeder) EnterFuncDefParam(ctx *FuncDefParamContext) {}

// ExitFuncDefParam is called when production funcDefParam is exited.
func (bf *ByteFeeder) ExitFuncDefParam(ctx *FuncDefParamContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (bf *ByteFeeder) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (bf *ByteFeeder) ExitTypeDef(ctx *TypeDefContext) {}

// EnterTypeTerm is called when production typeTerm is entered.
func (bf *ByteFeeder) EnterTypeTerm(ctx *TypeTermContext) {}

// ExitTypeTerm is called when production typeTerm is exited.
func (bf *ByteFeeder) ExitTypeTerm(ctx *TypeTermContext) {}

// EnterStmt is called when production stmt is entered.
func (bf *ByteFeeder) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (bf *ByteFeeder) ExitStmt(ctx *StmtContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (bf *ByteFeeder) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (bf *ByteFeeder) ExitAssignStmt(ctx *AssignStmtContext) {
	identName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.IndexOf(identName)
	if pos == -1 {
		pos = len(bf.symbolTable)
		bf.symbolTable = append(bf.symbolTable, identName)
	}

	bf.instructions = append(bf.instructions, Instruction{OP_ASSIGN_STMT, 0, uint64(pos)})
}

// EnterPipeline is called when production pipeline is entered.
func (bf *ByteFeeder) EnterPipeline(ctx *PipelineContext) {
	bf.instructions = append(bf.instructions, Instruction{OP_START_PIPELINE, 0, 0})
}

// ExitPipeline is called when production pipeline is exited.
func (bf *ByteFeeder) ExitPipeline(ctx *PipelineContext) {
	bf.instructions = append(bf.instructions, Instruction{OP_END_PIPELINE, 0, 0})
}

// EnterInlinePipeline is called when production inlinePipeline is entered.
func (bf *ByteFeeder) EnterInlinePipeline(ctx *InlinePipelineContext) {
	bf.instructions = append(bf.instructions, Instruction{OP_START_PIPELINE, 0, 0})
}

// ExitInlinePipeline is called when production inlinePipeline is exited.
func (bf *ByteFeeder) ExitInlinePipeline(ctx *InlinePipelineContext) {
	bf.instructions = append(bf.instructions, Instruction{OP_END_PIPELINE, 0, 0})
}

// EnterIdentBacktick is called when production identBacktick is entered.
func (bf *ByteFeeder) EnterIdentBacktick(ctx *IdentBacktickContext) {}

// ExitIdentBacktick is called when production identBacktick is exited.
func (bf *ByteFeeder) ExitIdentBacktick(ctx *IdentBacktickContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (bf *ByteFeeder) EnterFuncCall(ctx *FuncCallContext) {
	bf.instructions = append(bf.instructions, Instruction{OP_START_FUNC_CALL, 0, 0})
}

// ExitFuncCall is called when production funcCall is exited.
func (bf *ByteFeeder) ExitFuncCall(ctx *FuncCallContext) {
	funcName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.IndexOf(funcName)
	if pos == -1 {
		pos = len(bf.symbolTable)
		bf.symbolTable = append(bf.symbolTable, funcName)
	}

	bf.instructions = append(bf.instructions, Instruction{OP_MAKE_FUNC_CALL, 0, uint64(pos)})
}

// EnterFuncCallParam is called when production funcCallParam is entered.
func (bf *ByteFeeder) EnterFuncCallParam(ctx *FuncCallParamContext) {}

// ExitFuncCallParam is called when production funcCallParam is exited.
func (bf *ByteFeeder) ExitFuncCallParam(ctx *FuncCallParamContext) {
	bf.instructions = append(bf.instructions, Instruction{OP_END_CHUNCK, 0, 0})
}

// EnterNamedArg is called when production namedArg is entered.
func (bf *ByteFeeder) EnterNamedArg(ctx *NamedArgContext) {}

// ExitNamedArg is called when production namedArg is exited.
func (bf *ByteFeeder) ExitNamedArg(ctx *NamedArgContext) {
	paramName := ctx.IDENT().GetSymbol().GetText()
	pos := bf.symbolTable.IndexOf(paramName)
	if pos == -1 {
		pos = len(bf.symbolTable)
		bf.symbolTable = append(bf.symbolTable, paramName)
	}

	bf.instructions = append(bf.instructions, Instruction{OP_PUSH_NAMED_PARAM, 0, uint64(pos)})
}

// EnterAssign is called when production assign is entered.
func (bf *ByteFeeder) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (bf *ByteFeeder) ExitAssign(ctx *AssignContext) {}

// EnterMultiAssign is called when production multiAssign is entered.
func (bf *ByteFeeder) EnterMultiAssign(ctx *MultiAssignContext) {}

// ExitMultiAssign is called when production multiAssign is exited.
func (bf *ByteFeeder) ExitMultiAssign(ctx *MultiAssignContext) {}

// EnterExprCall is called when production exprCall is entered.
func (bf *ByteFeeder) EnterExprCall(ctx *ExprCallContext) {}

// ExitExprCall is called when production exprCall is exited.
func (bf *ByteFeeder) ExitExprCall(ctx *ExprCallContext) {}

// EnterExpr is called when production expr is entered.
func (bf *ByteFeeder) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (bf *ByteFeeder) ExitExpr(ctx *ExprContext) {}

// EnterTerm is called when production term is entered.
func (bf *ByteFeeder) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (bf *ByteFeeder) ExitTerm(ctx *TermContext) {}

// EnterExprUnary is called when production exprUnary is entered.
func (bf *ByteFeeder) EnterExprUnary(ctx *ExprUnaryContext) {}

// ExitExprUnary is called when production exprUnary is exited.
func (bf *ByteFeeder) ExitExprUnary(ctx *ExprUnaryContext) {}

// EnterLiteral is called when production literal is entered.
func (bf *ByteFeeder) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (bf *ByteFeeder) ExitLiteral(ctx *LiteralContext) {
	switch ctx.GetChildCount() {
	case 1:
		// NULL
		if ctx.NULL_() == nil {
			bf.instructions = append(bf.instructions, Instruction{OP_PUSH_TERM, TERM_NULL, 0})
		} else

		// BOOLEAN
		if ctx.BOOLEAN() != nil {
			if ctx.BOOLEAN().GetText() == "true" {
				bf.instructions = append(bf.instructions, Instruction{OP_PUSH_TERM, TERM_BOOL, 1})
			} else {
				bf.instructions = append(bf.instructions, Instruction{OP_PUSH_TERM, TERM_BOOL, 0})
			}
		} else

		// INTEGER
		// if ctx.INTEGER() != nil && ctx.INTEGER().length > 0 {
		// 	n, _ := strconv.ParseInt(ctx.INTEGER()[0].getText(), 10, 64)
		// 	bf.instructions = append(bf.instructions, Instruction{OP_PUSH_TERM, TERM_INTEGER, n})
		// } else

		// // FLOAT
		// if ctx.FLOAT() != nil && ctx.FLOAT().length > 0 {
		// 	bf.instructions = append(bf.instructions, OP_PUSH_TERM, TERM_FLOAT, parseFloat(ctx.FLOAT()[0].getText()))
		// } else

		// STRING
		if ctx.STRING() != nil {
			replacer := strings.NewReplacer("\"", "", "'", "")
			str := replacer.Replace(ctx.STRING().GetText())
			pos := bf.symbolTable.IndexOf(str)
			if pos == -1 {
				pos = len(bf.symbolTable)
				bf.symbolTable = append(bf.symbolTable, str)
			}

			bf.instructions = append(bf.instructions, Instruction{OP_PUSH_TERM, TERM_STRING, uint64(pos)})
		}

		// IDENT
		// if ctx.IDENT() != nil && ctx.IDENT() > 0 {
		// 	id := ctx.IDENT()[0].getText()
		// 	pos := bf.symbolTable.IndexOf(id)
		// 	if pos == -1 {
		// 		pos = len(bf.symbolTable)
		// 		bf.symbolTable = append(bf.symbolTable, id)
		// 	}

		// 	bf.instructions = append(bf.instructions, Instruction{OP_PUSH_TERM, TERM_SYMBOL, uint64(pos)})
		// }

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
	bf.instructions = append(bf.instructions, Instruction{OP_START_LIST, 0, 0})
}

// ExitList is called when production list is exited.
func (bf *ByteFeeder) ExitList(ctx *ListContext) {
	bf.instructions = append(bf.instructions, Instruction{OP_END_LIST, 0, 0})
}

// EnterNestedPipeline is called when production nestedPipeline is entered.
// func (bf *ByteFeeder) EnterNestedPipeline(ctx *NestedPipelineContext) {}

// ExitNestedPipeline is called when production nestedPipeline is exited.
// func (bf *ByteFeeder) ExitNestedPipeline(ctx *NestedPipelineContext) {}
