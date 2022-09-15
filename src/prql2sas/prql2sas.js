import antlr4 from "antlr4";

import prqlListener from "../grammar/prqlListener.js";
import prqlLexer from "../grammar/prqlLexer.js";
import prqlParser from "../grammar/prqlParser.js";
import { PRQL_ENVIRONMENT } from "./env.js";

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

class PrqlExpression {
  constructor() {
    this.stack = [];
  }

  push(t) {
    this.stack.push(t);
  }

  resolve() {}
}

class Term {}

export default class Prql2SASTranspiler extends prqlListener {
  constructor(env) {
    super();
    this.currentEnv = env;

    this.currTempTableId = 0;

    this.pipelineStack = [];
    this.currPipeline = null;

    this.currExpr = null;

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
    // console.log("func name:", name);
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
    if (this.currExpr === null) {
      this.currExpr = new PrqlExpression();
    }
  }

  // Exit a parse tree produced by prqlParser#expr.
  exitExpr(ctx) {
    // operation or nested expression
    if (ctx.children.length === 3) {
      if (ctx.children[0].symbol && ctx.children[0].symbol.text === "(") {
        // console.log(ctx.children[0].symbol.text);
      } else {
        this.currExpr.push(ctx.children[1].getText());
      }
    }

    console.log(this.currExpr.stack);
  }

  // Enter a parse tree produced by prqlParser#term.
  enterTerm(ctx) {}

  // Exit a parse tree produced by prqlParser#term.
  exitTerm(ctx) {
    this.currExpr.push(ctx.getText());
  }

  // Enter a parse tree produced by prqlParser#exprUnary.
  enterExprUnary(ctx) {}

  // Exit a parse tree produced by prqlParser#exprUnary.
  exitExprUnary(ctx) {}

  // Enter a parse tree produced by prqlParser#literal.
  enterLiteral(ctx) {}

  // Exit a parse tree produced by prqlParser#literal.
  exitLiteral(ctx) {
    if (ctx.children[0].NUMBER()) {
      if (ctx.children.length === 2) {
      } else {
      }
    } else if (ctx.children[0].BOOLEAN()) {
    } else if (ctx.children[0].NULL_()) {
    } else if (ctx.children[0].STRING()) {
    } else {
    }
  }

  // Enter a parse tree produced by prqlParser#list.
  enterList(ctx) {}

  // Exit a parse tree produced by prqlParser#list.
  exitList(ctx) {}

  // Enter a parse tree produced by prqlParser#nestedPipeline.
  enterNestedPipeline(ctx) {}

  // Exit a parse tree produced by prqlParser#nestedPipeline.
  exitNestedPipeline(ctx) {}
}

export function transpile(source) {
  const { CommonTokenStream, InputStream } = antlr4;

  const chars = new InputStream(source, true);
  const lexer = new prqlLexer(chars);
  const tokens = new CommonTokenStream(lexer);
  const parser = new prqlParser(tokens);

  parser.buildParseTrees = true;
  const tree = parser.query();
  const transpiler = new Prql2SASTranspiler(PRQL_ENVIRONMENT);
  antlr4.tree.ParseTreeWalker.DEFAULT.walk(transpiler, tree);

  return transpiler.getSASCode();
}
