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
  LANG_FUNC_CALL,
  LANG_PIPELINE,
  OP_ADD_FUNC_PARAM,
  OP_BEGIN_LIST,
  OP_BEGIN_PIPELINE,
  OP_CALL_FUNC,
  OP_END_LIST,
  OP_END_PIPELINE,
  PrqlVM,
  TYPE_BOOL,
  TYPE_IDENT,
  TYPE_INTERVAL,
  TYPE_LIST,
  TYPE_NULL,
  TYPE_NUMERIC,
  TYPE_RANGE,
  TYPE_STRING,
} from "./vm.js";

export default class PrqlFrontEnd extends prqlListener {
  constructor(params) {
    super();

    let debugLevel = 0;
    let verbose = false;
    if (params && params.debugLevel) {
      debugLevel = params.debugLevel;
    }
    if (params && params.verbose) {
      verbose = params.verbose;
    }

    this.__debug_level__ = params.debugLevel;
    this.__verbose__ = params.verbose;

    this.__rec_depth__ = 0;
    this.__indent__ = "  ";

    if (this.__verbose__) {
      console.log(`  ****  PRQL Listener  ****  `);
      console.log(`    Debug Level: ${this.__debug_level__}`);
      console.log(`    Verbose:     ${this.__verbose__}`);
    }

    this.term = null;
    this.param = null;

    this.vm = new PrqlVM({ debugLevel: this.__debug_level__ });
  }

  printByteCode() {
    this.vm.printByteCode();
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
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> Pipeline`);
    }
    this.vm.push(OP_BEGIN_PIPELINE);

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#pipeline.
  exitPipeline(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- Pipeline`);
    }
    this.vm.push(OP_END_PIPELINE);
  }

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
    if (this.__debug_level__ > 10) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) +
          `-> FuncCall: ${ctx.IDENT().symbol.text}`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#funcCall.
  exitFuncCall(ctx) {
    this.__rec_depth__--;
    const funcName = ctx.IDENT().symbol.text;
    if (this.__debug_level__ > 10) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) + `<- FuncCall: ${funcName}`
      );
    }

    this.vm.push(OP_CALL_FUNC, funcName);
  }

  // Enter a parse tree produced by prqlParser#funcCallParam.
  enterFuncCallParam(ctx) {
    if (this.__debug_level__ > 10) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) + `-> FuncCallParam`
      );
    }

    this.param = { name: null, ident: null, expr: [] };

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#funcCallParam.
  exitFuncCallParam(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 10) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) + `<- FuncCallParam`
      );
    }

    this.vm.push(
      OP_ADD_FUNC_PARAM,
      this.param.name,
      this.param.ident,
      this.param.expr
    );
    this.param = null;
  }

  // Enter a parse tree produced by prqlParser#namedArg.
  enterNamedArg(ctx) {
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> NamedArg`);
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#namedArg.
  exitNamedArg(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- NamedArg`);
    }

    if (this.param !== null) {
      this.param.name = ctx.IDENT().symbol.text;
    }
  }

  // Enter a parse tree produced by prqlParser#assign.
  enterAssign(ctx) {
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> Assign`);
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#assign.
  exitAssign(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- Assign`);
    }

    if (this.param !== null) {
      this.param.ident = ctx.IDENT().symbol.text;
    }
  }

  // Enter a parse tree produced by prqlParser#assignCall.
  enterAssignCall(ctx) {}

  // Exit a parse tree produced by prqlParser#assignCall.
  exitAssignCall(ctx) {}

  // Enter a parse tree produced by prqlParser#exprCall.
  enterExprCall(ctx) {}

  // Exit a parse tree produced by prqlParser#exprCall.
  exitExprCall(ctx) {
    // if (this.funcCall !== null) {
    //   this.pipeline.push({
    //     type: LANG_FUNC_CALL,
    //     name: this.funcCall.name,
    //     params: this.funcCall.params,
    //   });
    //   this.funcCall = null;
    // } else if (this.expr !== null) {
    //   this.pipeline.push({
    //     type: LANG_EXPR,
    //     value: this.expr,
    //   });
    //   this.expr = null;
    // } else {
    //   console.error("exitExprCall: no func call nor expression available.");
    // }
  }

  // Enter a parse tree produced by prqlParser#expr.
  enterExpr(ctx) {
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> Expr`);
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#expr.
  exitExpr(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- Expr`);
    }

    if (this.param !== null) {
      // operation or nested expression
      if (ctx.children.length === 3) {
        if (ctx.children[0].symbol && ctx.children[0].symbol.text === "(") {
          // console.log(ctx.children[0].symbol.text);
        } else {
          switch (ctx.children[1].getText()) {
            case "*":
              this.param.expr.push({ type: BINARY_OP_MUL });
              break;
            case "/":
              this.param.expr.push({ type: BINARY_OP_DIV });
              break;
            case "%":
              this.param.expr.push({ type: BINARY_OP_MOD });
              break;
            case "+":
              this.param.expr.push({ type: BINARY_OP_PLUS });
              break;
            case "-":
              this.param.expr.push({ type: BINARY_OP_MINUS });
              break;
            case "==":
              this.param.expr.push({ type: BINARY_OP_EQ });
              break;
            case "!=":
              this.param.expr.push({ type: BINARY_OP_NE });
              break;
            case ">=":
              this.param.expr.push({ type: BINARY_OP_GE });
              break;
            case "<=":
              this.param.expr.push({ type: BINARY_OP_LE });
              break;
            case ">":
              this.param.expr.push({ type: BINARY_OP_GT });
              break;
            case "<":
              this.param.expr.push({ type: BINARY_OP_LT });
              break;
          }
        }
      }
    }
  }

  // Enter a parse tree produced by prqlParser#term.
  enterTerm(ctx) {
    if (this.__debug_level__ > 15) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) + `-> Term: ${ctx.getText()}`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#term.
  exitTerm(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 15) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- Term`);
    }

    if (this.param !== null) {
      this.param.expr.push(this.term);
    }
  }

  // Enter a parse tree produced by prqlParser#exprUnary.
  enterExprUnary(ctx) {}

  // Exit a parse tree produced by prqlParser#exprUnary.
  exitExprUnary(ctx) {}

  // Enter a parse tree produced by prqlParser#literal.
  enterLiteral(ctx) {
    if (this.__debug_level__ > 15) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) +
          `-> Literal: ${ctx.getText()}`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#literal.
  exitLiteral(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 15) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- Literal`);
    }

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
          num: parseFloat(ctx.children[0].getText()),
          kind: ctx.children[1].getText(),
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
          start: start,
          end: end,
        };
        break;
    }
  }

  // Enter a parse tree produced by prqlParser#list.
  enterList(ctx) {
    if (this.__debug_level__ > 15) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> List`);
    }

    this.vm.push(OP_BEGIN_LIST);

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#list.
  exitList(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 15) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- List`);
    }

    this.vm.push(OP_END_LIST);
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
  const fontend = new PrqlFrontEnd({ debugLevel: 20, verbose: true });
  antlr4.tree.ParseTreeWalker.DEFAULT.walk(fontend, tree);

  fontend.printByteCode();
}
