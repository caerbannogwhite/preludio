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

export const BINARY_OP_MUL = 100;
export const BINARY_OP_DIV = 101;
export const BINARY_OP_MOD = 102;
export const BINARY_OP_PLUS = 103;
export const BINARY_OP_MINUS = 104;

export const BINARY_OP_EQ = 110;
export const BINARY_OP_NE = 111;
export const BINARY_OP_GE = 112;
export const BINARY_OP_LE = 113;
export const BINARY_OP_GT = 114;
export const BINARY_OP_LT = 115;

export const BINARY_OP_AND = 120;
export const BINARY_OP_OR = 121;
export const BINARY_OP_COALESCE = 122;

export const LANG_QUERY = 500;
export const LANG_PIPELINE = 501;
export const LANG_EXPR_CALL = 502;
export const LANG_FUNC_CALL = 503;
export const LANG_EXPR = 504;
export const LANG_ASSIGN = 505;
export const LANG_ASSIGN_CALL = 506;
export const LANG_NAMED_ARG = 507;

export const OP_BEGIN_PIPELINE = 0;
export const OP_END_PIPELINE = 1;
// export const OP_BEGIN_FUNC_CALL = 2;
// export const OP_END_FUNC_CALL = 3;
export const OP_BEGIN_LIST = 4;
export const OP_END_LIST = 5;
export const OP_ADD_FUNC_PARAM = 6;
export const OP_CALL_FUNC = 8;
// export const OP_BEGIN_EXPR = 6;
// export const OP_END_EXPR = 7;

const STATUS_INIT = 0;
const STATUS_READING_EXPR = 1;

export class PrqlVM {
  constructor(debugLevel = 0) {
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

      case OP_BEGIN_FUNC_CALL:
        if (this.__debug_level__ > 10) {
          console.log(`OP_BEGIN_FUNC_CALL, ${param1}`);
        }
        this.__current_function_name__ = param1;
        this.__stacks__[this.__stack_pointer__].push([
          OP_BEGIN_FUNC_CALL,
          param1,
        ]);
        break;

      case OP_END_FUNC_CALL:
        if (this.__debug_level__ > 10) {
          console.log(`OP_END_FUNC_CALL`);
        }
        this.__stacks__[this.__stack_pointer__].push([OP_END_FUNC_CALL]);
        break;

      case OP_BEGIN_LIST:
        if (this.__debug_level__ > 10) {
          console.log(`OP_BEGIN_LIST`);
        }
        this.__current_function_name__ = param1;
        this.__stacks__[this.__stack_pointer__].push([OP_BEGIN_LIST, param1]);
        break;

      case OP_END_LIST:
        if (this.__debug_level__ > 10) {
          console.log(`OP_END_LIST`);
        }
        this.__stacks__[this.__stack_pointer__].push([OP_END_LIST]);
        break;

      // ADD FUNC PARAM: param name, assign ident, expr
      case OP_ADD_FUNC_PARAM:
        if (this.__debug_level__ > 10) {
          console.log(`OP_ADD_FUNC_PARAM, ${param1}, ${param2}, ${param3}`);
        }
        this.__stacks__[this.__stack_pointer__].push([
          OP_ADD_FUNC_PARAM,
          param1,
          param2,
          param3,
        ]);
        break;
    }
  }

  printByteCode() {
    for (let stack of this.__stacks__) {
      for (let op of stack) {
        console.log(op);
      }
    }
  }
}
