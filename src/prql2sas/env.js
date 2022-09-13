import { fromCSV, loadJSON } from "arquero";

const funcDerive = {
  name: "derive",
  implementation: function (env, params) {},
};

const funcFilter = {
  name: "filter",
  implementation: function (env, params) {},
};

const funcFrom = {
  name: "from",
  implementation: function (env, params) {},
};

// params:
//  - file path
//  - [type]: "csv", "json"
const funcImport = {
  name: "import",
  implementation: function (env, params) {
    if (params.length < 1) {
      console.error("function import: expecting at least one parameter.");
      return;
    }

    let fileType = "csv";
    for (let p of params) {
      if (p.name === "type") {
        fileType = p.value;
      }
    }

    // file path
    const filePath = params[1].value;
    switch (fileType) {
      case "csv":
        env.currentTable = fromCSV(FileAttachment(filePath).text());
        break;

      case "json":
        env.currentTable = loadJSON(filePath);
        break;

      default:
        console.error(`function import: file type not supported (${fileType})`);
        break;
    }
  },
};

const funcSelect = {
  name: "select",
  implementation: function (env, params) {},
};

export const PRQL_ENVIRONMENT = {
  variables: {
    currentDirectory: "",
    currentTable: null,
  },
  functions: {
    derive: funcDerive,
    filter: funcFilter,
    from: funcFrom,
    import: funcImport,
    select: funcSelect,
  },
};
