import {
  __std_derive__,
  __std_filter__,
  __std_from__,
  __std_import__,
  __std_select__,
} from "./std";

export const TYPE_NULL = 0;
export const TYPE_BOOL = 1;
export const TYPE_NUMERIC = 2;
export const TYPE_STRING = 3;
export const TYPE_IDENT = 4;
export const TYPE_INTERVAL = 5;
export const TYPE_RANGE = 6;
export const TYPE_LIST = 7;
export const TYPE_PIPELINE = 8;

export const OP_BEGIN_PIPELINE = 0;
export const OP_END_PIPELINE = 1;
export const OP_ASSIGN_TABLE = 2;
export const OP_BEGIN_FUNC_CALL = 3;
export const OP_END_FUNC_CALL = 4;
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

const STATUS_INIT = 0;
const STATUS_READING_EXPR = 1;

export class PrqlVM {
  constructor(params) {
    let debugLevel = 0;
    if (params && params.debugLevel) {
      debugLevel = params.debugLevel;
    }

    this.__debug_level__ = debugLevel;

    this.__current_directory__ = "";
    this.__current_table__ = null;
    this.__current_tables_list__ = [];
    this.__current_function_name__ = null;
    this.__current_status__ = STATUS_INIT;

    this.__functions__ = {
      derive: __std_derive__,
      filter: __std_filter__,
      from: __std_from__,
      import: __std_import__,
      select: __std_select__,
    };

    this.__stack_pointer__ = 0;
    this.__stacks__ = [];
  }

  push(opCode, param1 = null, param2 = null, param3 = null) {
    switch (opCode) {
      case OP_BEGIN_PIPELINE:
        if (this.__debug_level__ > 10) {
          console.log(`OP_BEGIN_PIPELINE`);
        }
        this.__stacks__.unshift([]);
        break;

      case OP_END_PIPELINE:
        if (this.__debug_level__ > 10) {
          console.log(`OP_END_PIPELINE`);
        }
        this.__stack_pointer__++;
        break;

      case OP_BEGIN_LIST:
        if (this.__debug_level__ > 15) {
          console.log(`OP_BEGIN_LIST`);
        }
        break;

      case OP_END_LIST:
        if (this.__debug_level__ > 15) {
          console.log(`OP_END_LIST`);
        }
        break;

      case OP_ADD_EXPR_TERM:
        if (this.__debug_level__ > 15) {
          console.log(`OP_ADD_EXPR_TERM: ${param1}, ${param2}, ${param3}`);
        }
        break;

      // ADD FUNC PARAM: param name, assign ident, expr
      case OP_ADD_FUNC_PARAM:
        if (this.__debug_level__ > 15) {
          console.log(`OP_ADD_FUNC_PARAM: ${param1}, ${param2}, ${param3}`);
        }
        break;

      case OP_CALL_FUNC:
        if (this.__debug_level__ > 10) {
          console.log(`OP_CALL_FUNC: ${param3}`);
        }
        break;
    }
    this.__stacks__[this.__stack_pointer__].push([
      opCode,
      param1,
      param2,
      param3,
    ]);
  }

  printByteCode() {
    for (let stack of this.__stacks__) {
      for (let op of stack) {
        console.log(op);
      }
    }
  }
}
