// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package bytefeeder // preludioParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BasepreludioParserListener is a complete listener for a parse tree produced by preludioParser.
type BasepreludioParserListener struct{}

var _ preludioParserListener = &BasepreludioParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepreludioParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepreludioParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepreludioParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepreludioParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNl is called when production nl is entered.
func (s *BasepreludioParserListener) EnterNl(ctx *NlContext) {}

// ExitNl is called when production nl is exited.
func (s *BasepreludioParserListener) ExitNl(ctx *NlContext) {}

// EnterProgram is called when production program is entered.
func (s *BasepreludioParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasepreludioParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgramIntro is called when production programIntro is entered.
func (s *BasepreludioParserListener) EnterProgramIntro(ctx *ProgramIntroContext) {}

// ExitProgramIntro is called when production programIntro is exited.
func (s *BasepreludioParserListener) ExitProgramIntro(ctx *ProgramIntroContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BasepreludioParserListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BasepreludioParserListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterFuncDefName is called when production funcDefName is entered.
func (s *BasepreludioParserListener) EnterFuncDefName(ctx *FuncDefNameContext) {}

// ExitFuncDefName is called when production funcDefName is exited.
func (s *BasepreludioParserListener) ExitFuncDefName(ctx *FuncDefNameContext) {}

// EnterFuncDefParams is called when production funcDefParams is entered.
func (s *BasepreludioParserListener) EnterFuncDefParams(ctx *FuncDefParamsContext) {}

// ExitFuncDefParams is called when production funcDefParams is exited.
func (s *BasepreludioParserListener) ExitFuncDefParams(ctx *FuncDefParamsContext) {}

// EnterFuncDefParam is called when production funcDefParam is entered.
func (s *BasepreludioParserListener) EnterFuncDefParam(ctx *FuncDefParamContext) {}

// ExitFuncDefParam is called when production funcDefParam is exited.
func (s *BasepreludioParserListener) ExitFuncDefParam(ctx *FuncDefParamContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BasepreludioParserListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BasepreludioParserListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterTypeTerm is called when production typeTerm is entered.
func (s *BasepreludioParserListener) EnterTypeTerm(ctx *TypeTermContext) {}

// ExitTypeTerm is called when production typeTerm is exited.
func (s *BasepreludioParserListener) ExitTypeTerm(ctx *TypeTermContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasepreludioParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasepreludioParserListener) ExitStmt(ctx *StmtContext) {}

// EnterVarAssignStmt is called when production varAssignStmt is entered.
func (s *BasepreludioParserListener) EnterVarAssignStmt(ctx *VarAssignStmtContext) {}

// ExitVarAssignStmt is called when production varAssignStmt is exited.
func (s *BasepreludioParserListener) ExitVarAssignStmt(ctx *VarAssignStmtContext) {}

// EnterVarDeclStmt is called when production varDeclStmt is entered.
func (s *BasepreludioParserListener) EnterVarDeclStmt(ctx *VarDeclStmtContext) {}

// ExitVarDeclStmt is called when production varDeclStmt is exited.
func (s *BasepreludioParserListener) ExitVarDeclStmt(ctx *VarDeclStmtContext) {}

// EnterRetStmt is called when production retStmt is entered.
func (s *BasepreludioParserListener) EnterRetStmt(ctx *RetStmtContext) {}

// ExitRetStmt is called when production retStmt is exited.
func (s *BasepreludioParserListener) ExitRetStmt(ctx *RetStmtContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BasepreludioParserListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BasepreludioParserListener) ExitPipeline(ctx *PipelineContext) {}

// EnterInlinePipeline is called when production inlinePipeline is entered.
func (s *BasepreludioParserListener) EnterInlinePipeline(ctx *InlinePipelineContext) {}

// ExitInlinePipeline is called when production inlinePipeline is exited.
func (s *BasepreludioParserListener) ExitInlinePipeline(ctx *InlinePipelineContext) {}

// EnterNestedPipeline is called when production nestedPipeline is entered.
func (s *BasepreludioParserListener) EnterNestedPipeline(ctx *NestedPipelineContext) {}

// ExitNestedPipeline is called when production nestedPipeline is exited.
func (s *BasepreludioParserListener) ExitNestedPipeline(ctx *NestedPipelineContext) {}

// EnterIdentBacktick is called when production identBacktick is entered.
func (s *BasepreludioParserListener) EnterIdentBacktick(ctx *IdentBacktickContext) {}

// ExitIdentBacktick is called when production identBacktick is exited.
func (s *BasepreludioParserListener) ExitIdentBacktick(ctx *IdentBacktickContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BasepreludioParserListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BasepreludioParserListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncCallParam is called when production funcCallParam is entered.
func (s *BasepreludioParserListener) EnterFuncCallParam(ctx *FuncCallParamContext) {}

// ExitFuncCallParam is called when production funcCallParam is exited.
func (s *BasepreludioParserListener) ExitFuncCallParam(ctx *FuncCallParamContext) {}

// EnterNamedArg is called when production namedArg is entered.
func (s *BasepreludioParserListener) EnterNamedArg(ctx *NamedArgContext) {}

// ExitNamedArg is called when production namedArg is exited.
func (s *BasepreludioParserListener) ExitNamedArg(ctx *NamedArgContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasepreludioParserListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasepreludioParserListener) ExitAssign(ctx *AssignContext) {}

// EnterMultiAssign is called when production multiAssign is entered.
func (s *BasepreludioParserListener) EnterMultiAssign(ctx *MultiAssignContext) {}

// ExitMultiAssign is called when production multiAssign is exited.
func (s *BasepreludioParserListener) ExitMultiAssign(ctx *MultiAssignContext) {}

// EnterExprCall is called when production exprCall is entered.
func (s *BasepreludioParserListener) EnterExprCall(ctx *ExprCallContext) {}

// ExitExprCall is called when production exprCall is exited.
func (s *BasepreludioParserListener) ExitExprCall(ctx *ExprCallContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasepreludioParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasepreludioParserListener) ExitExpr(ctx *ExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BasepreludioParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasepreludioParserListener) ExitTerm(ctx *TermContext) {}

// EnterExprUnary is called when production exprUnary is entered.
func (s *BasepreludioParserListener) EnterExprUnary(ctx *ExprUnaryContext) {}

// ExitExprUnary is called when production exprUnary is exited.
func (s *BasepreludioParserListener) ExitExprUnary(ctx *ExprUnaryContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasepreludioParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasepreludioParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterList is called when production list is entered.
func (s *BasepreludioParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BasepreludioParserListener) ExitList(ctx *ListContext) {}
