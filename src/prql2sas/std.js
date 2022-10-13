import { loadCSV, loadJSON } from "arquero";

export const __std_derive__ = (env, params) => {
  const evaluatedParams = evaluateParams(env, params);
};

export const __std_filter__ = (env, params) => {};

export const __std_from__ = (env, params) => {};

// params:
//  - file path
//  - [type]: "csv", "json"
export const __std_import__ = async (env, params) => {
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
  const filePath = evaluatedParams[0].value;
  switch (fileType) {
    case "csv":
      env.__current_table__ = loadCSV(filePath);
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

export const __std_select__ = (env, params) => {};

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

      case LANG_ASSIGN:
        break;

      // Default
      default:
        console.error(`[PRQL] Parameter type unknown: ${p.type}.`);
        break;
    }
    evaluatedParams.push({ name: p.name, value: v });
  }

  return evaluatedParams;
};
