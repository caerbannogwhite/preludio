import { fromCSV, loadJSON } from "arquero";

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

const __std_derive__ = (env, params) => {
  const evaluatedParams = evaluateParams(env, params);
};

const __std_filter__ = (env, params) => {};

const __std_from__ = (env, params) => {};

// params:
//  - file path
//  - [type]: "csv", "json"
const __std_import__ = (env, params) => {
  if (params.length < 1) {
    console.error(`[import] 🧐 Expecting at least one parameter.`);
    return;
  }

  const evaluatedParams = evaluateParams(env, params);

  let fileType = "csv";
  for (let p of evaluatedParams) {
    if (p.name === "type") {
      fileType = p.value;
    }
  }

  // file path
  const filePath = evaluatedParams[1].value;
  switch (fileType) {
    case "csv":
      env.__current_table__ = fromCSV(FileAttachment(filePath).text());
      break;

    case "json":
      env.__current_table__ = loadJSON(filePath);
      break;

    default:
      console.error(
        `[import] 😣 File type not supported (entered type: "${fileType}").`
      );
      break;
  }
};

const __std_select__ = (env, params) => {};

export const PRQL_ENVIRONMENT = {
  variables: {
    __current_directory__: "",
    __current_table__: null,
  },
  functions: {
    derive: __std_derive__,
    filter: __std_filter__,
    from: __std_from__,
    import: __std_import__,
    select: __std_select__,
  },
};

const evaluateParams = (env, params) => {
  const evaluatedParams = [];
  for (let p of params) {
    let v = null;
    switch (p.type) {
      //
      // Expression
      case LANG_EXPR:
        if (p.value.length === undefined) {
          // do nothing ?
        } else if (p.value.length === 0) {
          // do nothing ?
        } else if (p.value.length === 1) {
          v = p.value[0].value;
        } else {
          const left = [];
          for (let e of p.value) {
          }
        }
        break;

      // Default
      default:
        console.error(`[PRQL] Parameter type unknown: ${p.type}.`);
        break;
    }
    evaluatedParams.push({ name: p.name, value: v });
  }

  return evaluateParams;
};
