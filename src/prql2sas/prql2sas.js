import antlr4 from "antlr4";

import prqlListener from "../grammar/prqlListener.js";
import prqlLexer from "../grammar/prqlLexer.js";
import prqlParser from "../grammar/prqlParser.js";

class FuncCall {
  constructor() {
    this.name = "";
    this.params = [];
  }
}

class Pipeline {
  constructor() {
    this.funcCallStack = [];
    this.currFuncCall = null;
  }

  addFuncCall(funcCall) {
    if (this.currFuncCall !== null) {
      this.funcCallStack.push(this.currFuncCall);
    }
    this.currFuncCall = funcCall;
  }

  removeFuncCall() {
    if (this.funcCallStack.length > 0) {
      this.currFuncCall = this.funcCallStack.pop();
    } else {
      this.currFuncCall = null;
    }
  }
}

class ExprTree {
  constructor() {}
  resolve() {}
}

class Term {}

export default class Prql2SASTranspiler extends prqlListener {
  constructor() {
    super();
    this.currTempTableId = 0;

    this.pipelineStack = [];
    this.currPipeline = null;

    this.currExprTree = null;

    this.variableStack = [];
  }

  getSASCode() {
    return "";
  }

  // Enter a parse tree produced by prqlParser#nl.
  enterNl(ctx) {}

  // Exit a parse tree produced by prqlParser#nl.
  exitNl(ctx) {}

  // Enter a parse tree produced by prqlParser#query.
  enterQuery(ctx) {
    if (this.currPipeline !== null) {
      this.pipelineStack.push(this.currPipeline);
    }
    this.currPipeline = new Pipeline();
  }

  // Exit a parse tree produced by prqlParser#query.
  exitQuery(ctx) {
    if (this.pipelineStack.length > 0) {
      this.currPipeline = this.pipelineStack.pop();
    } else {
      this.currPipeline = null;
    }
  }

  // Enter a parse tree produced by prqlParser#queryDef.
  enterQueryDef(ctx) {}

  // Exit a parse tree produced by prqlParser#queryDef.
  exitQueryDef(ctx) {}

  // Enter a parse tree produced by prqlParser#funcDef.
  enterFuncDef(ctx) {}

  // Exit a parse tree produced by prqlParser#funcDef.
  exitFuncDef(ctx) {}

  // Enter a parse tree produced by prqlParser#funcDefName.
  enterFuncDefName(ctx) {}

  // Exit a parse tree produced by prqlParser#funcDefName.
  exitFuncDefName(ctx) {}

  // Enter a parse tree produced by prqlParser#funcDefParams.
  enterFuncDefParams(ctx) {}

  // Exit a parse tree produced by prqlParser#funcDefParams.
  exitFuncDefParams(ctx) {}

  // Enter a parse tree produced by prqlParser#funcDefParam.
  enterFuncDefParam(ctx) {}

  // Exit a parse tree produced by prqlParser#funcDefParam.
  exitFuncDefParam(ctx) {}

  // Enter a parse tree produced by prqlParser#typeDef.
  enterTypeDef(ctx) {}

  // Exit a parse tree produced by prqlParser#typeDef.
  exitTypeDef(ctx) {}

  // Enter a parse tree produced by prqlParser#typeTerm.
  enterTypeTerm(ctx) {}

  // Exit a parse tree produced by prqlParser#typeTerm.
  exitTypeTerm(ctx) {}

  // Enter a parse tree produced by prqlParser#table.
  enterTable(ctx) {}

  // Exit a parse tree produced by prqlParser#table.
  exitTable(ctx) {}

  // Enter a parse tree produced by prqlParser#pipe.
  enterPipe(ctx) {}

  // Exit a parse tree produced by prqlParser#pipe.
  exitPipe(ctx) {}

  // Enter a parse tree produced by prqlParser#pipeline.
  enterPipeline(ctx) {}

  // Exit a parse tree produced by prqlParser#pipeline.
  exitPipeline(ctx) {}

  // Enter a parse tree produced by prqlParser#identBackticks.
  enterIdentBackticks(ctx) {}

  // Exit a parse tree produced by prqlParser#identBackticks.
  exitIdentBackticks(ctx) {}

  // Enter a parse tree produced by prqlParser#signedIdent.
  enterSignedIdent(ctx) {}

  // Exit a parse tree produced by prqlParser#signedIdent.
  exitSignedIdent(ctx) {}

  // Enter a parse tree produced by prqlParser#keyword.
  enterKeyword(ctx) {}

  // Exit a parse tree produced by prqlParser#keyword.
  exitKeyword(ctx) {}

  // Enter a parse tree produced by prqlParser#funcCall.
  enterFuncCall(ctx) {
    this.currPipeline.addFuncCall(new FuncCall());
  }

  // Exit a parse tree produced by prqlParser#funcCall.
  exitFuncCall(ctx) {
    const name = ctx.IDENT().symbol.text;
    console.log("func name:", name);
  }

  // Enter a parse tree produced by prqlParser#namedArg.
  enterNamedArg(ctx) {}

  // Exit a parse tree produced by prqlParser#namedArg.
  exitNamedArg(ctx) {}

  // Enter a parse tree produced by prqlParser#assign.
  enterAssign(ctx) {}

  // Exit a parse tree produced by prqlParser#assign.
  exitAssign(ctx) {}

  // Enter a parse tree produced by prqlParser#assignCall.
  enterAssignCall(ctx) {}

  // Exit a parse tree produced by prqlParser#assignCall.
  exitAssignCall(ctx) {}

  // Enter a parse tree produced by prqlParser#exprCall.
  enterExprCall(ctx) {}

  // Exit a parse tree produced by prqlParser#exprCall.
  exitExprCall(ctx) {}

  // Enter a parse tree produced by prqlParser#expr.
  enterExpr(ctx) {
    if (this.currExprTree === null) {
      this.currExprTree = new ExprTree();
    }
  }

  // Exit a parse tree produced by prqlParser#expr.
  exitExpr(ctx) {
    // term
    if (ctx.children.length === 1) {
    } else {
      console.log(ctx.children[1]);

      switch (typeof ctx.children[1]) {
        case prqlParser.OperatorAddContext:
          break;
        case prqlParser.OperatorCoalesceContext:
          break;
        case prqlParser.OperatorCompareContext:
          break;
        case prqlParser.OperatorLogicalContext:
          break;
        case prqlParser.OperatorMulContext:
          break;
      }
    }
  }

  // Enter a parse tree produced by prqlParser#term.
  enterTerm(ctx) {}

  // Exit a parse tree produced by prqlParser#term.
  exitTerm(ctx) {}

  // Enter a parse tree produced by prqlParser#exprUnary.
  enterExprUnary(ctx) {}

  // Exit a parse tree produced by prqlParser#exprUnary.
  exitExprUnary(ctx) {}

  // Enter a parse tree produced by prqlParser#literal.
  enterLiteral(ctx) {}

  // Exit a parse tree produced by prqlParser#literal.
  exitLiteral(ctx) {}

  // Enter a parse tree produced by prqlParser#list.
  enterList(ctx) {}

  // Exit a parse tree produced by prqlParser#list.
  exitList(ctx) {}

  // Enter a parse tree produced by prqlParser#nestedPipeline.
  enterNestedPipeline(ctx) {}

  // Exit a parse tree produced by prqlParser#nestedPipeline.
  exitNestedPipeline(ctx) {}

  // Enter a parse tree produced by prqlParser#range.
  enterRange(ctx) {}

  // Exit a parse tree produced by prqlParser#range.
  exitRange(ctx) {}

  // Enter a parse tree produced by prqlParser#operator.
  enterOperator(ctx) {}

  // Exit a parse tree produced by prqlParser#operator.
  exitOperator(ctx) {}

  // Enter a parse tree produced by prqlParser#operatorUnary.
  enterOperatorUnary(ctx) {}

  // Exit a parse tree produced by prqlParser#operatorUnary.
  exitOperatorUnary(ctx) {}

  // Enter a parse tree produced by prqlParser#operatorMul.
  enterOperatorMul(ctx) {}

  // Exit a parse tree produced by prqlParser#operatorMul.
  exitOperatorMul(ctx) {}

  // Enter a parse tree produced by prqlParser#operatorAdd.
  enterOperatorAdd(ctx) {}

  // Exit a parse tree produced by prqlParser#operatorAdd.
  exitOperatorAdd(ctx) {}

  // Enter a parse tree produced by prqlParser#operatorCompare.
  enterOperatorCompare(ctx) {}

  // Exit a parse tree produced by prqlParser#operatorCompare.
  exitOperatorCompare(ctx) {}

  // Enter a parse tree produced by prqlParser#operatorLogical.
  enterOperatorLogical(ctx) {}

  // Exit a parse tree produced by prqlParser#operatorLogical.
  exitOperatorLogical(ctx) {}

  // Enter a parse tree produced by prqlParser#operatorCoalesce.
  enterOperatorCoalesce(ctx) {}

  // Exit a parse tree produced by prqlParser#operatorCoalesce.
  exitOperatorCoalesce(ctx) {}

  // Enter a parse tree produced by prqlParser#intervalKind.
  enterIntervalKind(ctx) {}

  // Exit a parse tree produced by prqlParser#intervalKind.
  exitIntervalKind(ctx) {}

  // Enter a parse tree produced by prqlParser#interval.
  enterInterval(ctx) {}

  // Exit a parse tree produced by prqlParser#interval.
  exitInterval(ctx) {}
}

export function transpile(source) {
  const { CommonTokenStream, InputStream } = antlr4;

  var chars = new InputStream(source, true);
  var lexer = new prqlLexer(chars);
  var tokens = new CommonTokenStream(lexer);
  var parser = new prqlParser(tokens);

  parser.buildParseTrees = true;
  var tree = parser.query();
  var transpiler = new Prql2SASTranspiler();
  antlr4.tree.ParseTreeWalker.DEFAULT.walk(transpiler, tree);

  return transpiler.getSASCode();
}
