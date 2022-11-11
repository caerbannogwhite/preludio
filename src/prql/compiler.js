import antlr4 from "antlr4";

import prqlListener from "../grammar/prqlListener.js";
import prqlLexer from "../grammar/prqlLexer.js";
import prqlParser from "../grammar/prqlParser.js";
import { Blob } from "buffer";
import { TextEncoder } from "util";

export const TYPE_NULL = 0;
export const TYPE_BOOL = 1;
export const TYPE_NUMERIC = 2;
export const TYPE_STRING = 3;
export const TYPE_INTERVAL = 5;
export const TYPE_RANGE = 6;
export const TYPE_LIST = 7;
export const TYPE_PIPELINE = 8;
export const TERM_IDENT = 10;

export const OP_BEGIN_PIPELINE = 0;
export const OP_END_PIPELINE = 1;
export const OP_ASSIGN_STMT = 2;
// export const OP_BEGIN_FUNC_CALL = 3;
export const OP_MAKE_FUNC_CALL = 4;
export const OP_BEGIN_LIST = 5;
export const OP_END_LIST = 6;
export const OP_ADD_FUNC_PARAM = 7;
export const OP_ADD_EXPR_TERM = 8;
export const OP_PUSH_NAMED_PARAM = 9;
export const OP_PUSH_ASSIGN_IDENT = 10;
export const OP_PUSH_TERM = 11;
export const OP_END_FUNC_CALL_PARAM = 12;
export const OP_GOTO = 50;

export const OP_BINARY_MUL = 100;
export const OP_BINARY_DIV = 101;
export const OP_BINARY_MOD = 102;
export const OP_BINARY_PLUS = 103;
export const OP_BINARY_MINUS = 104;

export const OP_BINARY_EQ = 110;
export const OP_BINARY_NE = 111;
export const OP_BINARY_GE = 112;
export const OP_BINARY_LE = 113;
export const OP_BINARY_GT = 114;
export const OP_BINARY_LT = 115;

export const OP_BINARY_AND = 120;
export const OP_BINARY_OR = 121;
export const OP_BINARY_COALESCE = 122;

export default class PrqlCompiler extends prqlListener {
  constructor(params) {
    super();

    let debugLevel = 0;
    let verbosity = false;
    if (params && params.debugLevel) {
      debugLevel = params.debugLevel;
    }
    if (params && params.verbosity) {
      verbosity = params.verbosity;
    }

    this.__debug_level = debugLevel;
    this.__verbosity_level = verbosity;

    this.__indent_symbol = "  ";

    this.__symbol_table_str = [];
    this.__instructions = [];

    if (this.__verbosity_level) {
      console.log(`  ****  PRQL Compiler  ****`);
      console.log(`    Debug Level: ${this.__debug_level}`);
      console.log(`    Verbose:     ${this.__verbosity_level}`);
      console.log(`  ****                 ****\n`);
    }

    this.terms = [];
  }

  _toByteArray2(n) {
    const b = [0, 0];
    let i = 1;
    while (i < 2) {
      let r = n % 256 ** i;
      n -= r;
      b[2 - i] = Math.floor(r / 256 ** (i - 1));
      i++;
    }
    return b;
  }

  _toByteArray8(n) {
    const b = [0, 0, 0, 0, 0, 0, 0, 0];
    let i = 1;
    while (i < 8) {
      let r = n % 256 ** i;
      n -= r;
      b[8 - i] = Math.floor(r / 256 ** (i - 1));
      i++;
    }
    return b;
  }

  getByteCode() {
    const encoder = new TextEncoder();
    const stringSymbols = [];
    for (let s of this.__symbol_table_str) {
      const v = encoder.encode(s);
      stringSymbols.push(...this._toByteArray8(v.length), ...v);
    }

    const instructions = [];
    for (let i = 0; i < this.__instructions.length; i += 3) {
      if (
        this.__instructions[i] === OP_PUSH_TERM &&
        this.__instructions[i + 1] === TYPE_NUMERIC
      ) {
        instructions.push(
          ...this._toByteArray2(this.__instructions[i]),
          ...this._toByteArray2(this.__instructions[i + 1]),
          ...new Uint8Array(
            new Float64Array([this.__instructions[i + 2]]).buffer
          )
        );
      } else {
        instructions.push(
          ...this._toByteArray2(this.__instructions[i]),
          ...this._toByteArray2(this.__instructions[i + 1]),
          ...this._toByteArray8(this.__instructions[i + 2])
        );
      }
    }

    return new Blob([
      new Uint8Array([
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

        // for each element in the symbol table, get the lenght of
        // the encoded string and the uint8 encoded array
        ...this._toByteArray8(this.__symbol_table_str.length),
        ...stringSymbols,

        // instructions as uint 64 array
        ...instructions,
      ]),
    ]);
  }

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

  // Enter a parse tree produced by prqlParser#assignStmt.
  enterAssignStmt(ctx) {
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `-> AssignStmt`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#assignStmt.
  exitAssignStmt(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `<- AssignStmt`
      );
    }

    const identName = ctx.IDENT().symbol.text;
    let pos = this.__symbol_table_str.indexOf(identName);
    if (pos === -1) {
      pos = this.__symbol_table_str.length;
      this.__symbol_table_str.push(identName);
    }

    this.__instructions.push(OP_ASSIGN_STMT, 0, pos);
  }

  // Enter a parse tree produced by prqlParser#pipeline.
  enterPipeline(ctx) {
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `-> Pipeline`
      );
    }

    this.__instructions.push(OP_BEGIN_PIPELINE, 0, 0);

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#pipeline.
  exitPipeline(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `<- Pipeline`
      );
    }

    this.__instructions.push(OP_END_PIPELINE, 0, 0);
  }

  // Enter a parse tree produced by prqlParser#identBackticks.
  // enterIdentBackticks(ctx) {}

  // Exit a parse tree produced by prqlParser#identBackticks.
  // exitIdentBackticks(ctx) {}

  // Enter a parse tree produced by prqlParser#signedIdent.
  enterSignedIdent(ctx) {}

  // Exit a parse tree produced by prqlParser#signedIdent.
  exitSignedIdent(ctx) {}

  // Enter a parse tree produced by prqlParser#keyword.
  // enterKeyword(ctx) {}

  // Exit a parse tree produced by prqlParser#keyword.
  // exitKeyword(ctx) {}

  // Enter a parse tree produced by prqlParser#funcCall.
  enterFuncCall(ctx) {
    if (this.__debug_level > 10) {
      const funcName = ctx.IDENT().symbol.text;
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) +
          `-> FuncCall: ${funcName}`
      );
    }

    // this.__instructions.push(OP_BEGIN_FUNC_CALL, 0, 0);

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#funcCall.
  exitFuncCall(ctx) {
    this.__rec_depth__--;
    const funcName = ctx.IDENT().symbol.text;
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) +
          `<- FuncCall: ${funcName}`
      );
    }

    let pos = this.__symbol_table_str.indexOf(funcName);
    if (pos === -1) {
      pos = this.__symbol_table_str.length;
      this.__symbol_table_str.push(funcName);
    }

    this.__instructions.push(OP_MAKE_FUNC_CALL, 0, pos);
  }

  // Enter a parse tree produced by prqlParser#funcCallParam.
  enterFuncCallParam(ctx) {
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `-> FuncCallParam`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#funcCallParam.
  exitFuncCallParam(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `<- FuncCallParam`
      );
    }

    this.__instructions.push(OP_END_FUNC_CALL_PARAM, 0, 0);
  }

  // Enter a parse tree produced by prqlParser#namedArg.
  enterNamedArg(ctx) {
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `-> NamedArg`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#namedArg.
  exitNamedArg(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `<- NamedArg`
      );
    }

    const paramName = ctx.IDENT().symbol.text;
    let pos = this.__symbol_table_str.indexOf(paramName);
    if (pos === -1) {
      pos = this.__symbol_table_str.length;
      this.__symbol_table_str.push(paramName);
    }

    this.__instructions.push(OP_PUSH_NAMED_PARAM, 0, pos);
  }

  // Enter a parse tree produced by prqlParser#assign.
  enterAssign(ctx) {
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `-> Assign`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#assign.
  exitAssign(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 10) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `<- Assign`
      );
    }

    const identName = ctx.IDENT().symbol.text;
    let pos = this.__symbol_table_str.indexOf(identName);
    if (pos === -1) {
      pos = this.__symbol_table_str.length;
      this.__symbol_table_str.push(identName);
    }

    this.__instructions.push(OP_PUSH_ASSIGN_IDENT, 0, pos);
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
    if (this.__debug_level > 10) {
      console.log(this.__indent_symbol.repeat(this.__rec_depth__) + `-> Expr`);
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#expr.
  exitExpr(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 10) {
      console.log(this.__indent_symbol.repeat(this.__rec_depth__) + `<- Expr`);
    }

    // operation or nested expression
    if (ctx.children.length === 3) {
      if (ctx.children[0].symbol && ctx.children[0].symbol.text === "(") {
        // console.log(ctx.children[0].symbol.text);
      } else {
        switch (ctx.children[1].getText()) {
          case "*":
            this.__instructions.push(OP_BINARY_MUL, 0, 0);
            break;
          case "/":
            this.__instructions.push(OP_BINARY_DIV, 0, 0);
            break;
          case "%":
            this.__instructions.push(OP_BINARY_MOD, 0, 0);
            break;
          case "+":
            this.__instructions.push(OP_BINARY_PLUS, 0, 0);
            break;
          case "-":
            this.__instructions.push(OP_BINARY_MINUS, 0, 0);
            break;
          case "==":
            this.__instructions.push(OP_BINARY_EQ, 0, 0);
            break;
          case "!=":
            this.__instructions.push(OP_BINARY_NE, 0, 0);
            break;
          case ">=":
            this.__instructions.push(OP_BINARY_GE, 0, 0);
            break;
          case "<=":
            this.__instructions.push(OP_BINARY_LE, 0, 0);
            break;
          case ">":
            this.__instructions.push(OP_BINARY_GT, 0, 0);
            break;
          case "<":
            this.__instructions.push(OP_BINARY_LT, 0, 0);
            break;
        }
      }
    }
  }

  // Enter a parse tree produced by prqlParser#term.
  enterTerm(ctx) {
    if (this.__debug_level > 15) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) +
          `-> Term: ${ctx.getText()}`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#term.
  exitTerm(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 15) {
      console.log(this.__indent_symbol.repeat(this.__rec_depth__) + `<- Term`);
    }
  }

  // Enter a parse tree produced by prqlParser#exprUnary.
  enterExprUnary(ctx) {}

  // Exit a parse tree produced by prqlParser#exprUnary.
  exitExprUnary(ctx) {}

  // Enter a parse tree produced by prqlParser#literal.
  enterLiteral(ctx) {
    if (this.__debug_level > 15) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) +
          `-> Literal: ${ctx.getText()}`
      );
    }

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#literal.
  exitLiteral(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 15) {
      console.log(
        this.__indent_symbol.repeat(this.__rec_depth__) + `<- Literal`
      );
    }

    switch (ctx.children.length) {
      case 1:
        // NULL
        if (ctx.NULL_() !== null) {
          this.__instructions.push(OP_PUSH_TERM, TYPE_NULL, 0);
        }

        // BOOLEAN
        else if (ctx.BOOLEAN() !== null) {
          this.__instructions.push(
            OP_PUSH_TERM,
            TYPE_BOOL,
            ctx.BOOLEAN().getText() === "true" ? 1 : 0
          );
        }

        // NUMBER
        else if (ctx.NUMBER() !== null && ctx.NUMBER().length > 0) {
          this.__instructions.push(
            OP_PUSH_TERM,
            TYPE_NUMERIC,
            parseFloat(ctx.NUMBER()[0].getText())
          );
        }

        // STRING
        else if (ctx.STRING() !== null) {
          const str = ctx.STRING().getText().replace(/['"]+/g, "");
          let pos = this.__symbol_table_str.indexOf(str);
          if (pos === -1) {
            pos = this.__symbol_table_str.length;
            this.__symbol_table_str.push(str);
          }

          this.__instructions.push(OP_PUSH_TERM, TYPE_STRING, pos);
        }

        // IDENT
        else if (ctx.IDENT() !== null && ctx.IDENT().length > 0) {
          const id = ctx.IDENT()[0].getText();
          let pos = this.__symbol_table_str.indexOf(id);
          if (pos === -1) {
            pos = this.__symbol_table_str.length;
            this.__symbol_table_str.push(id);
          }

          this.__instructions.push(OP_PUSH_TERM, TERM_IDENT, pos);
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
        //   const start = { type: TERM_IDENT, value: ctx.children[0].getText() };
        // } else {
        //   const start = { type: TYPE_NUMERIC, value: s };
        // }

        // const e = ctx.children[2].getText();
        // if (end === NaN) {
        //   const end = { type: TERM_IDENT, value: ctx.children[2].getText() };
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
    if (this.__debug_level > 15) {
      console.log(this.__indent_symbol.repeat(this.__rec_depth__) + `-> List`);
    }

    this.__instructions.push(OP_BEGIN_LIST, 0, 0);

    this.__rec_depth__++;
  }

  // Exit a parse tree produced by prqlParser#list.
  exitList(ctx) {
    this.__rec_depth__--;
    if (this.__debug_level > 15) {
      console.log(this.__indent_symbol.repeat(this.__rec_depth__) + `<- List`);
    }

    this.__instructions.push(OP_END_LIST, 0, 0);
  }

  // Enter a parse tree produced by prqlParser#nestedPipeline.
  // enterNestedPipeline(ctx) {}

  // Exit a parse tree produced by prqlParser#nestedPipeline.
  // exitNestedPipeline(ctx) {}
}

export function getByteCode(source) {
  const { CommonTokenStream, InputStream } = antlr4;

  const chars = new InputStream(source, true);
  const lexer = new prqlLexer(chars);
  const tokens = new CommonTokenStream(lexer);
  const parser = new prqlParser(tokens);

  parser.buildParseTrees = true;
  const tree = parser.program();
  const compiler = new PrqlCompiler({ debugLevel: 5, verbosity: true });
  antlr4.tree.ParseTreeWalker.DEFAULT.walk(compiler, tree);

  return compiler.getByteCode();
}
