import antlr4 from "antlr4";

import prqlListener from "../grammar/prqlListener.js";
import prqlLexer from "../grammar/prqlLexer.js";
import prqlParser from "../grammar/prqlParser.js";
import {
  BINARY_OP_DIV,
  BINARY_OP_EQ,
  BINARY_OP_GE,
  BINARY_OP_GT,
  BINARY_OP_LT,
  BINARY_OP_MINUS,
  BINARY_OP_MOD,
  BINARY_OP_MUL,
  BINARY_OP_NE,
  BINARY_OP_PLUS,
  LANG_ASSIGN,
  LANG_EXPR,
  PRQL_ENVIRONMENT,
  TYPE_BOOL,
  TYPE_IDENT,
  TYPE_INTERVAL,
  TYPE_LIST,
  TYPE_NULL,
  TYPE_NUMERIC,
  TYPE_RANGE,
  TYPE_STRING,
} from "./env.js";

export default class Prql2SASTranspiler extends prqlListener {
  constructor(env) {
    super();
    this.currentEnv = env;

    this.currTempTableId = 0;

    this.pipelineStack = [];
    this.currPipeline = null;

    this.funcCallParams = null;

    this.expr = null;
    this.term = null;

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
  enterQuery(ctx) {}

  // Exit a parse tree produced by prqlParser#query.
  exitQuery(ctx) {}

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
  enterPipeline(ctx) {
    this.currPipeline = [];
  }

  // Exit a parse tree produced by prqlParser#pipeline.
  exitPipeline(ctx) {}

  // Enter a parse tree produced by prqlParser#identBackticks.
  enterIdentBackticks(ctx) {}

  // Exit a parse tree produced by prqlParser#identBackticks.
  exitIdentBackticks(ctx) {
    this.term = {
      type: TYPE_IDENT,
      value: ctx.children[0].getText(),
    };
  }

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
    this.funcCallParams = [];
  }

  // Exit a parse tree produced by prqlParser#funcCall.
  exitFuncCall(ctx) {
    this.funcCall = {
      name: ctx.IDENT().symbol.text,
      params: this.funcCallParams,
    };
    this.funcCallParams = null;
  }

  // Enter a parse tree produced by prqlParser#namedArg.
  enterNamedArg(ctx) {}

  // Exit a parse tree produced by prqlParser#namedArg.
  exitNamedArg(ctx) {
    if (this.expr === null) {
      this.namedArg = {
        type: LANG_ASSIGN,
        name: ctx.IDENT().symbol.text,
        value: this.assign,
      };
      this.assign = null;
    } else {
      this.namedArg = {
        type: LANG_EXPR,
        name: ctx.IDENT().symbol.text,
        value: this.expr,
      };
      this.expr = null;
    }
  }

  // Enter a parse tree produced by prqlParser#assign.
  enterAssign(ctx) {}

  // Exit a parse tree produced by prqlParser#assign.
  exitAssign(ctx) {
    this.assign = {
      ident: ctx.IDENT().symbol.text,
      expr: this.expr,
    };
    this.expr = null;
  }

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
    if (this.expr === null) {
      this.expr = [];
    }
  }

  // Exit a parse tree produced by prqlParser#expr.
  exitExpr(ctx) {
    // operation or nested expression
    if (ctx.children.length === 3) {
      if (ctx.children[0].symbol && ctx.children[0].symbol.text === "(") {
        // console.log(ctx.children[0].symbol.text);
      } else {
        switch (ctx.children[1].getText()) {
          case "*":
            this.expr.push({ type: BINARY_OP_MUL });
            break;
          case "/":
            this.expr.push({ type: BINARY_OP_DIV });
            break;
          case "%":
            this.expr.push({ type: BINARY_OP_MOD });
            break;
          case "+":
            this.expr.push({ type: BINARY_OP_PLUS });
            break;
          case "-":
            this.expr.push({ type: BINARY_OP_MINUS });
            break;
          case "==":
            this.expr.push({ type: BINARY_OP_EQ });
            break;
          case "!=":
            this.expr.push({ type: BINARY_OP_NE });
            break;
          case ">=":
            this.expr.push({ type: BINARY_OP_GE });
            break;
          case "<=":
            this.expr.push({ type: BINARY_OP_LE });
            break;
          case ">":
            this.expr.push({ type: BINARY_OP_GT });
            break;
          case "<":
            this.expr.push({ type: BINARY_OP_LT });
            break;
        }
      }
    }
  }

  // Enter a parse tree produced by prqlParser#term.
  enterTerm(ctx) {}

  // Exit a parse tree produced by prqlParser#term.
  exitTerm(ctx) {
    this.expr.push(this.term);
  }

  // Enter a parse tree produced by prqlParser#exprUnary.
  enterExprUnary(ctx) {}

  // Exit a parse tree produced by prqlParser#exprUnary.
  exitExprUnary(ctx) {}

  // Enter a parse tree produced by prqlParser#literal.
  enterLiteral(ctx) {}

  // Exit a parse tree produced by prqlParser#literal.
  exitLiteral(ctx) {
    switch (ctx.children.length) {
      case 1:
        if (ctx.NULL_() !== null) {
          this.term = { type: TYPE_NULL };
        } else if (ctx.BOOLEAN() !== null) {
          this.term = {
            type: TYPE_BOOL,
            value: ctx.BOOLEAN().getText() === "true",
          };
        } else if (ctx.NUMBER() !== null && ctx.NUMBER().length > 0) {
          this.term = {
            type: TYPE_NUMERIC,
            value: parseFloat(ctx.NUMBER()[0].getText()),
          };
        } else if (ctx.STRING() !== null) {
          this.term = {
            type: TYPE_STRING,
            value: ctx.STRING().getText().replace(/['"]+/g, ""),
          };
        } else if (ctx.IDENT() !== null && ctx.IDENT().length > 0) {
          this.term = {
            type: TYPE_IDENT,
            value: ctx.IDENT()[0].getText(),
          };
        }
        break;

      // time interval
      case 2:
        this.term = {
          type: TYPE_INTERVAL,
          value: {
            num: parseFloat(ctx.children[0].getText()),
            kind: ctx.children[1].getText(),
          },
        };
        break;

      // range
      case 3:
        const s = ctx.children[0].getText();
        if (s === NaN) {
          const start = { type: TYPE_IDENT, value: ctx.children[0].getText() };
        } else {
          const start = { type: TYPE_NUMERIC, value: s };
        }

        const e = ctx.children[2].getText();
        if (end === NaN) {
          const end = { type: TYPE_IDENT, value: ctx.children[2].getText() };
        } else {
          const end = { type: TYPE_NUMERIC, value: e };
        }

        this.term = {
          type: TYPE_RANGE,
          value: {
            start: start,
            end: end,
          },
        };
        break;
    }
  }

  // Enter a parse tree produced by prqlParser#list.
  enterList(ctx) {}

  // Exit a parse tree produced by prqlParser#list.
  exitList(ctx) {
    this.term = {
      type: TYPE_LIST,
      value: null,
    };
  }

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
