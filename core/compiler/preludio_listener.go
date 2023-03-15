// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package preludiocompiler // preludio
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// preludioListener is a complete listener for a parse tree produced by preludioParser.
type preludioListener interface {
	antlr.ParseTreeListener

	// EnterNl is called when entering the nl production.
	EnterNl(c *NlContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterProgramIntro is called when entering the programIntro production.
	EnterProgramIntro(c *ProgramIntroContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterFuncDefName is called when entering the funcDefName production.
	EnterFuncDefName(c *FuncDefNameContext)

	// EnterFuncDefParams is called when entering the funcDefParams production.
	EnterFuncDefParams(c *FuncDefParamsContext)

	// EnterFuncDefParam is called when entering the funcDefParam production.
	EnterFuncDefParam(c *FuncDefParamContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterTypeTerm is called when entering the typeTerm production.
	EnterTypeTerm(c *TypeTermContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterPipeline is called when entering the pipeline production.
	EnterPipeline(c *PipelineContext)

	// EnterInlinePipeline is called when entering the inlinePipeline production.
	EnterInlinePipeline(c *InlinePipelineContext)

	// EnterIdentBacktick is called when entering the identBacktick production.
	EnterIdentBacktick(c *IdentBacktickContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncCallParam is called when entering the funcCallParam production.
	EnterFuncCallParam(c *FuncCallParamContext)

	// EnterNamedArg is called when entering the namedArg production.
	EnterNamedArg(c *NamedArgContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterMultiAssign is called when entering the multiAssign production.
	EnterMultiAssign(c *MultiAssignContext)

	// EnterExprCall is called when entering the exprCall production.
	EnterExprCall(c *ExprCallContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterExprUnary is called when entering the exprUnary production.
	EnterExprUnary(c *ExprUnaryContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterNestedPipeline is called when entering the nestedPipeline production.
	EnterNestedPipeline(c *NestedPipelineContext)

	// ExitNl is called when exiting the nl production.
	ExitNl(c *NlContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitProgramIntro is called when exiting the programIntro production.
	ExitProgramIntro(c *ProgramIntroContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitFuncDefName is called when exiting the funcDefName production.
	ExitFuncDefName(c *FuncDefNameContext)

	// ExitFuncDefParams is called when exiting the funcDefParams production.
	ExitFuncDefParams(c *FuncDefParamsContext)

	// ExitFuncDefParam is called when exiting the funcDefParam production.
	ExitFuncDefParam(c *FuncDefParamContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitTypeTerm is called when exiting the typeTerm production.
	ExitTypeTerm(c *TypeTermContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitPipeline is called when exiting the pipeline production.
	ExitPipeline(c *PipelineContext)

	// ExitInlinePipeline is called when exiting the inlinePipeline production.
	ExitInlinePipeline(c *InlinePipelineContext)

	// ExitIdentBacktick is called when exiting the identBacktick production.
	ExitIdentBacktick(c *IdentBacktickContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncCallParam is called when exiting the funcCallParam production.
	ExitFuncCallParam(c *FuncCallParamContext)

	// ExitNamedArg is called when exiting the namedArg production.
	ExitNamedArg(c *NamedArgContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitMultiAssign is called when exiting the multiAssign production.
	ExitMultiAssign(c *MultiAssignContext)

	// ExitExprCall is called when exiting the exprCall production.
	ExitExprCall(c *ExprCallContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitExprUnary is called when exiting the exprUnary production.
	ExitExprUnary(c *ExprUnaryContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitNestedPipeline is called when exiting the nestedPipeline production.
	ExitNestedPipeline(c *NestedPipelineContext)
}
