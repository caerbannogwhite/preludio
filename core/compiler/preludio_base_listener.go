// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package preludiocompiler // preludio
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BasepreludioListener is a complete listener for a parse tree produced by preludioParser.
type BasepreludioListener struct{}

var _ preludioListener = &BasepreludioListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepreludioListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepreludioListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepreludioListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepreludioListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNl is called when production nl is entered.
func (s *BasepreludioListener) EnterNl(ctx *NlContext) {}

// ExitNl is called when production nl is exited.
func (s *BasepreludioListener) ExitNl(ctx *NlContext) {}

// EnterProgram is called when production program is entered.
func (s *BasepreludioListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasepreludioListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgramIntro is called when production programIntro is entered.
func (s *BasepreludioListener) EnterProgramIntro(ctx *ProgramIntroContext) {}

// ExitProgramIntro is called when production programIntro is exited.
func (s *BasepreludioListener) ExitProgramIntro(ctx *ProgramIntroContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BasepreludioListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BasepreludioListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterFuncDefName is called when production funcDefName is entered.
func (s *BasepreludioListener) EnterFuncDefName(ctx *FuncDefNameContext) {}

// ExitFuncDefName is called when production funcDefName is exited.
func (s *BasepreludioListener) ExitFuncDefName(ctx *FuncDefNameContext) {}

// EnterFuncDefParams is called when production funcDefParams is entered.
func (s *BasepreludioListener) EnterFuncDefParams(ctx *FuncDefParamsContext) {}

// ExitFuncDefParams is called when production funcDefParams is exited.
func (s *BasepreludioListener) ExitFuncDefParams(ctx *FuncDefParamsContext) {}

// EnterFuncDefParam is called when production funcDefParam is entered.
func (s *BasepreludioListener) EnterFuncDefParam(ctx *FuncDefParamContext) {}

// ExitFuncDefParam is called when production funcDefParam is exited.
func (s *BasepreludioListener) ExitFuncDefParam(ctx *FuncDefParamContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BasepreludioListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BasepreludioListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterTypeTerm is called when production typeTerm is entered.
func (s *BasepreludioListener) EnterTypeTerm(ctx *TypeTermContext) {}

// ExitTypeTerm is called when production typeTerm is exited.
func (s *BasepreludioListener) ExitTypeTerm(ctx *TypeTermContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasepreludioListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasepreludioListener) ExitStmt(ctx *StmtContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BasepreludioListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BasepreludioListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterVarDefStmt is called when production varDefStmt is entered.
func (s *BasepreludioListener) EnterVarDefStmt(ctx *VarDefStmtContext) {}

// ExitVarDefStmt is called when production varDefStmt is exited.
func (s *BasepreludioListener) ExitVarDefStmt(ctx *VarDefStmtContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BasepreludioListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BasepreludioListener) ExitPipeline(ctx *PipelineContext) {}

// EnterInlinePipeline is called when production inlinePipeline is entered.
func (s *BasepreludioListener) EnterInlinePipeline(ctx *InlinePipelineContext) {}

// ExitInlinePipeline is called when production inlinePipeline is exited.
func (s *BasepreludioListener) ExitInlinePipeline(ctx *InlinePipelineContext) {}

// EnterIdentBacktick is called when production identBacktick is entered.
func (s *BasepreludioListener) EnterIdentBacktick(ctx *IdentBacktickContext) {}

// ExitIdentBacktick is called when production identBacktick is exited.
func (s *BasepreludioListener) ExitIdentBacktick(ctx *IdentBacktickContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BasepreludioListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BasepreludioListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncCallParam is called when production funcCallParam is entered.
func (s *BasepreludioListener) EnterFuncCallParam(ctx *FuncCallParamContext) {}

// ExitFuncCallParam is called when production funcCallParam is exited.
func (s *BasepreludioListener) ExitFuncCallParam(ctx *FuncCallParamContext) {}

// EnterNamedArg is called when production namedArg is entered.
func (s *BasepreludioListener) EnterNamedArg(ctx *NamedArgContext) {}

// ExitNamedArg is called when production namedArg is exited.
func (s *BasepreludioListener) ExitNamedArg(ctx *NamedArgContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasepreludioListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasepreludioListener) ExitAssign(ctx *AssignContext) {}

// EnterMultiAssign is called when production multiAssign is entered.
func (s *BasepreludioListener) EnterMultiAssign(ctx *MultiAssignContext) {}

// ExitMultiAssign is called when production multiAssign is exited.
func (s *BasepreludioListener) ExitMultiAssign(ctx *MultiAssignContext) {}

// EnterExprCall is called when production exprCall is entered.
func (s *BasepreludioListener) EnterExprCall(ctx *ExprCallContext) {}

// ExitExprCall is called when production exprCall is exited.
func (s *BasepreludioListener) ExitExprCall(ctx *ExprCallContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasepreludioListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasepreludioListener) ExitExpr(ctx *ExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BasepreludioListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasepreludioListener) ExitTerm(ctx *TermContext) {}

// EnterExprUnary is called when production exprUnary is entered.
func (s *BasepreludioListener) EnterExprUnary(ctx *ExprUnaryContext) {}

// ExitExprUnary is called when production exprUnary is exited.
func (s *BasepreludioListener) ExitExprUnary(ctx *ExprUnaryContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasepreludioListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasepreludioListener) ExitLiteral(ctx *LiteralContext) {}

// EnterList is called when production list is entered.
func (s *BasepreludioListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BasepreludioListener) ExitList(ctx *ListContext) {}

// EnterNestedPipeline is called when production nestedPipeline is entered.
func (s *BasepreludioListener) EnterNestedPipeline(ctx *NestedPipelineContext) {}

// ExitNestedPipeline is called when production nestedPipeline is exited.
func (s *BasepreludioListener) ExitNestedPipeline(ctx *NestedPipelineContext) {}
