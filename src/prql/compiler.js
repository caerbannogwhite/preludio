import antlr4 from "antlr4";

import prqlListener from "../grammar/prqlListener.js";
import prqlLexer from "../grammar/prqlLexer.js";
import prqlParser from "../grammar/prqlParser.js";
import {
  OP_ASSIGN_TABLE,
  OP_BEGIN_FUNC_CALL,
  OP_BEGIN_LIST,
  OP_BEGIN_PIPELINE,
  OP_BINARY_DIV,
  OP_BINARY_EQ,
  OP_BINARY_GE,
  OP_BINARY_GT,
  OP_BINARY_LE,
  OP_BINARY_LT,
  OP_BINARY_MINUS,
  OP_BINARY_MOD,
  OP_BINARY_MUL,
  OP_BINARY_NE,
  OP_BINARY_PLUS,
  OP_END_FUNC_CALL,
  OP_END_FUNC_CALL_PARAM,
  OP_END_LIST,
  OP_END_PIPELINE,
  OP_PUSH_ASSIGN_IDENT,
  OP_PUSH_NAMED_PARAM,
  OP_PUSH_TERM,
  PrqlVM,
  TYPE_BOOL,
  TYPE_IDENT,
  TYPE_NULL,
  TYPE_NUMERIC,
  TYPE_STRING,
} from "./vm.js";
import { Blob } from "buffer";
import { TextEncoder } from "util";

export default class PrqlCompiler extends prqlListener {
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

    this.__debug_level__ = debugLevel;
    this.__verbose__ = verbose;

    this.__rec_depth__ = 0;
    this.__indent__ = "  ";

    this.__symbol_table__ = [];
    this.__instructions__ = [];

    if (this.__verbose__) {
      console.log(`  ****  PRQL Compiler  ****`);
      console.log(`    Debug Level: ${this.__debug_level__}`);
      console.log(`    Verbose:     ${this.__verbose__}`);
      console.log(`  ****                 ****\n`);
    }

    this.terms = [];

    this.vm = new PrqlVM({ debugLevel: this.__debug_level__ });
  }

  getByteCode() {
    const toByteArray8 = (n) => {
      const b = [0, 0, 0, 0, 0, 0, 0, 0];
      let i = 1;
      while (i < 8) {
        let r = n % 256 ** i;
        n -= r;
        b[8 - i] = Math.floor(r / 256 ** (i - 1));
        i++;
      }
      return b;
    };

    const encoder = new TextEncoder();
    const symbols = [];
    for (let s of this.__symbol_table__) {
      const v = encoder.encode(s);
      symbols.push(...toByteArray8(v.length), ...v);
    }

    const a = new Uint8Array([
      // incipit: 4 bytes mark, 4 empty bytes
      // and the number of elements in the symbol table
      0x11,
      0x01,
      0x19,
      0x93,
      0x00,
      0x00,
      0x00,
      0x00,
      ...toByteArray8(this.__symbol_table__.length),

      // for each element in the symbol table, get the lenght of
      // the encoded string and the uint8 encoded array
      ...symbols,

      // instructions as uint 64 array
      ...this.__instructions__.map((i) => toByteArray8(i)).flat(),
    ]);

    return new Blob([a]);
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
  enterTable(ctx) {
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> Table`);
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#table.
  exitTable(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- Table`);
    }

    const identName = ctx.IDENT().symbol.text;
    let pos = this.__symbol_table__.indexOf(identName);
    if (pos === -1) {
      pos = this.__symbol_table__.length;
      this.__symbol_table__.push(identName);
    }

    this.__instructions__.push(OP_ASSIGN_TABLE, pos, 0);
  }

  // Enter a parse tree produced by prqlParser#pipe.
  enterPipe(ctx) {}

  // Exit a parse tree produced by prqlParser#pipe.
  exitPipe(ctx) {}

  // Enter a parse tree produced by prqlParser#pipeline.
  enterPipeline(ctx) {
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> Pipeline`);
    }

    this.__instructions__.push(OP_BEGIN_PIPELINE, 0, 0);

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#pipeline.
  exitPipeline(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 10) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- Pipeline`);
    }

    this.__instructions__.push(OP_END_PIPELINE, 0, 0);
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
    const funcName = ctx.IDENT().symbol.text;
    if (this.__debug_level__ > 10) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) +
          `-> FuncCall: ${funcName}`
      );
    }

    let pos = this.__symbol_table__.indexOf(funcName);
    if (pos === -1) {
      pos = this.__symbol_table__.length;
      this.__symbol_table__.push(funcName);
    }

    this.__instructions__.push(OP_BEGIN_FUNC_CALL, pos, 0);

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

    this.__instructions__.push(OP_END_FUNC_CALL, 0, 0);
  }

  // Enter a parse tree produced by prqlParser#funcCallParam.
  enterFuncCallParam(ctx) {
    if (this.__debug_level__ > 10) {
      console.log(
        this.__indent__.repeat(this.__rec_depth__) + `-> FuncCallParam`
      );
    }

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

    this.__instructions__.push(OP_END_FUNC_CALL_PARAM, 0, 0);
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

    const paramName = ctx.IDENT().symbol.text;
    let pos = this.__symbol_table__.indexOf(paramName);
    if (pos === -1) {
      pos = this.__symbol_table__.length;
      this.__symbol_table__.push(paramName);
    }

    this.__instructions__.push(OP_PUSH_NAMED_PARAM, pos, 0);
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

    const identName = ctx.IDENT().symbol.text;
    let pos = this.__symbol_table__.indexOf(identName);
    if (pos === -1) {
      pos = this.__symbol_table__.length;
      this.__symbol_table__.push(identName);
    }

    this.__instructions__.push(OP_PUSH_ASSIGN_IDENT, pos, 0);
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

    // operation or nested expression
    if (ctx.children.length === 3) {
      if (ctx.children[0].symbol && ctx.children[0].symbol.text === "(") {
        // console.log(ctx.children[0].symbol.text);
      } else {
        switch (ctx.children[1].getText()) {
          case "*":
            this.__instructions__.push(OP_BINARY_MUL, 0, 0);
            break;
          case "/":
            this.__instructions__.push(OP_BINARY_DIV, 0, 0);
            break;
          case "%":
            this.__instructions__.push(OP_BINARY_MOD, 0, 0);
            break;
          case "+":
            this.__instructions__.push(OP_BINARY_PLUS, 0, 0);
            break;
          case "-":
            this.__instructions__.push(OP_BINARY_MINUS, 0, 0);
            break;
          case "==":
            this.__instructions__.push(OP_BINARY_EQ, 0, 0);
            break;
          case "!=":
            this.__instructions__.push(OP_BINARY_NE, 0, 0);
            break;
          case ">=":
            this.__instructions__.push(OP_BINARY_GE, 0, 0);
            break;
          case "<=":
            this.__instructions__.push(OP_BINARY_LE, 0, 0);
            break;
          case ">":
            this.__instructions__.push(OP_BINARY_GT, 0, 0);
            break;
          case "<":
            this.__instructions__.push(OP_BINARY_LT, 0, 0);
            break;
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
        // NULL
        if (ctx.NULL_() !== null) {
          this.__instructions__.push(OP_PUSH_TERM, TYPE_NULL, 0);
        }

        // BOOLEAN
        else if (ctx.BOOLEAN() !== null) {
          this.__instructions__.push(
            OP_PUSH_TERM,
            TYPE_BOOL,
            ctx.BOOLEAN().getText() === "true" ? 1 : 0
          );
        }

        // NUMBER
        else if (ctx.NUMBER() !== null && ctx.NUMBER().length > 0) {
          const num = ctx.NUMBER()[0].getText();
          let pos = this.__symbol_table__.indexOf(num);
          if (pos === -1) {
            pos = this.__symbol_table__.length;
            this.__symbol_table__.push(num);
          }

          this.__instructions__.push(OP_PUSH_TERM, TYPE_NUMERIC, pos);
        }

        // STRING
        else if (ctx.STRING() !== null) {
          const str = ctx.STRING().getText().replace(/['"]+/g, "");
          let pos = this.__symbol_table__.indexOf(str);
          if (pos === -1) {
            pos = this.__symbol_table__.length;
            this.__symbol_table__.push(str);
          }

          this.__instructions__.push(OP_PUSH_TERM, TYPE_STRING, pos);
        }

        // IDENT
        else if (ctx.IDENT() !== null && ctx.IDENT().length > 0) {
          const id = ctx.IDENT()[0].getText();
          let pos = this.__symbol_table__.indexOf(id);
          if (pos === -1) {
            pos = this.__symbol_table__.length;
            this.__symbol_table__.push(id);
          }

          this.__instructions__.push(OP_PUSH_TERM, TYPE_IDENT, pos);
        }
        break;

      // time interval
      case 2:
        // this.term = {
        //   type: TYPE_INTERVAL,
        //   num: parseFloat(ctx.children[0].getText()),
        //   kind: ctx.children[1].getText(),
        // };
        break;

      // range
      case 3:
        // const s = ctx.children[0].getText();
        // if (s === NaN) {
        //   const start = { type: TYPE_IDENT, value: ctx.children[0].getText() };
        // } else {
        //   const start = { type: TYPE_NUMERIC, value: s };
        // }

        // const e = ctx.children[2].getText();
        // if (end === NaN) {
        //   const end = { type: TYPE_IDENT, value: ctx.children[2].getText() };
        // } else {
        //   const end = { type: TYPE_NUMERIC, value: e };
        // }

        // this.term = {
        //   type: TYPE_RANGE,
        //   start: start,
        //   end: end,
        // };
        break;
    }
  }

  // Enter a parse tree produced by prqlParser#list.
  enterList(ctx) {
    if (this.__debug_level__ > 15) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `-> List`);
    }

    this.__instructions__.push(OP_BEGIN_LIST, 0, 0);

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#list.
  exitList(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level__ > 15) {
      console.log(this.__indent__.repeat(this.__rec_depth__) + `<- List`);
    }

    this.__instructions__.push(OP_END_LIST, 0, 0);
  }

  // Enter a parse tree produced by prqlParser#nestedPipeline.
  enterNestedPipeline(ctx) {}

  // Exit a parse tree produced by prqlParser#nestedPipeline.
  exitNestedPipeline(ctx) {}
}

export function getByteCode(source) {
  const { CommonTokenStream, InputStream } = antlr4;

  const chars = new InputStream(source, true);
  const lexer = new prqlLexer(chars);
  const tokens = new CommonTokenStream(lexer);
  const parser = new prqlParser(tokens);

  parser.buildParseTrees = true;
  const tree = parser.query();
  const compiler = new PrqlCompiler({ debugLevel: 5, verbose: true });
  antlr4.tree.ParseTreeWalker.DEFAULT.walk(compiler, tree);

  return compiler.getByteCode();
}
