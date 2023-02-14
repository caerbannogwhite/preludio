"use strict";
import "./style.css";

import { App } from "./App";
import * as monaco from "monaco-editor/esm/vs/editor/editor.main";

// https://github.com/microsoft/monaco-editor/blob/main/docs/integrate-esm.md#using-vite
// https://github.com/microsoft/monaco-editor-webpack-plugin/issues/135
self.MonacoEnvironment = {
  getWorker: function (workerId, label) {
    const getWorkerModule = (moduleUrl, label) => {
      return new Worker(self.MonacoEnvironment.getWorkerUrl(moduleUrl), {
        name: label,
        type: "module",
      });
    };

    switch (label) {
      case "json":
        return getWorkerModule("/monaco-editor/esm/vs/language/json/json.worker?worker", label);
      case "css":
      case "scss":
      case "less":
        return getWorkerModule("/monaco-editor/esm/vs/language/css/css.worker?worker", label);
      case "html":
      case "handlebars":
      case "razor":
        return getWorkerModule("/monaco-editor/esm/vs/language/html/html.worker?worker", label);
      case "typescript":
      case "javascript":
        return getWorkerModule("/monaco-editor/esm/vs/language/typescript/ts.worker?worker", label);
      default:
        return getWorkerModule("/monaco-editor/esm/vs/editor/editor.worker?worker", label);
    }
  },
};

const app = new App();
document.body.appendChild(app);

monaco.editor.create(app.getCodeEditorPaneElement(), {
  value: "function hello() {\n\talert('Hello world!');\n}",
  language: "python",
  minimap: { enabled: false },
});

// Default Wails code

// Setup the greet function
// window.greet = function () {
//   // Get name
//   let name = nameElement!.value;

//   // Check if the input is empty
//   if (name === "") return;

//   // Call App.Greet(name)
//   try {
//     Greet(name)
//       .then((result) => {
//         // Update result with data back from App.Greet()
//         resultElement!.innerText = result;
//       })
//       .catch((err) => {
//         console.error(err);
//       });
//   } catch (err) {
//     console.error(err);
//   }
// };

// declare global {
//   interface Window {
//     greet: () => void;
//   }
// }
