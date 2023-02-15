// import { EditorView, basicSetup } from "codemirror";
import { ParseCsv } from "../wailsjs/go/main/App";
import { AppState } from "./AppState";
import { PreludioPipeline } from "./PreludioPipeline";
import { TopMenuPaneElement } from "./TopMenuPaneElement";
// import { MainButton } from "./utils/MainButton";
// import * as monaco from "monaco-editor";

export class App extends HTMLDivElement {
  private topMenuPaneElement: HTMLDivElement;
  private mainPaneElement: HTMLDivElement;
  private leftPaneElement: HTMLDivElement;
  private codeEditorPaneElement: HTMLDivElement;
  private logErrorPaneElement: HTMLDivElement;
  private tableEditorPaneElement: HTMLDivElement;

  private state: AppState;

  constructor() {
    super();
    this.id = "app";

    this.topMenuPaneElement = TopMenuPaneElement("top-menu-pane");
    this.mainPaneElement = document.createElement("div");
    this.leftPaneElement = document.createElement("div");
    this.codeEditorPaneElement = document.createElement("div");
    this.logErrorPaneElement = document.createElement("div");
    this.tableEditorPaneElement = document.createElement("div");

    this.state = new AppState();

    this._initHTMLElement();
  }

  private _initHTMLElement() {
    this.mainPaneElement.id = "main-pane";

    this.leftPaneElement.id = "left-pane";

    // Code Editor Pane
    this.codeEditorPaneElement.id = "code-editor-pane";

    // Log Error Pane
    this.logErrorPaneElement.id = "log-error-pane";

    // Table Editor Pane
    this.tableEditorPaneElement.id = "table-editor-pane";

    // TODO: NATIVE select input file feature
    // const filePathInput = new DropDownOptionsMenu(
    //   "file-path-input",
    //   [],
    //   (): DropDownOption[] => {
    //     let options = new Array<DropDownOption>();
    //     LookUpPath("C:").then((d) => {
    //       console.log(d);
    //       for (let o of d) {
    //         options.push({ value: o, name: o });
    //       }
    //     });

    //     return options;
    //   }
    // );

    // const formData = new FormData();
    const fileInput = document.createElement("input");
    fileInput.id = "import-table-input";
    fileInput.className = "file-input";
    fileInput.type = "file";
    fileInput.addEventListener("change", async (e: any) => {
      if (e.target !== null && e.target.files !== null && e.target.files.length > 0) {
        const file = e.target.files[0];
        Promise.resolve(file.arrayBuffer()).then((ab: ArrayBuffer) => {
          const u = new Uint8Array(ab);
          ParseCsv(`[${u.toString()}]`).then((d) => {
            console.log(d);
          });
          // SendFileFromJStoGo(`[${u.toString()}]`);
        });

        // const blob = new Blob([e.target.files[0]]);

        // const buffer = await blob.arrayBuffer();
        // console.log(buffer);

        // ParseCsv(buffer).then((d) => {
        //   console.log(d);
        // });
      }
    });

    // const importTableButton = new MainButton("import-table", "import-svgrepo-com", fileInput);

    // this.tableEditorPaneElement.appendChild(filePathInput);
    // this.tableEditorPaneElement.appendChild(importTableButton);

    this.leftPaneElement.appendChild(this.codeEditorPaneElement);
    this.leftPaneElement.appendChild(this.logErrorPaneElement);

    // Main Panel
    this.mainPaneElement.appendChild(this.leftPaneElement);
    this.mainPaneElement.appendChild(this.tableEditorPaneElement);

    // App
    this.appendChild(this.topMenuPaneElement);
    this.appendChild(this.mainPaneElement);
  }

  addNewPipeline(pipelineName?: string) {
    AppState.pipelineCounter++;
    let name = `Pipeline ${AppState.pipelineCounter}`;
    if (pipelineName !== undefined) {
      name = pipelineName;
    }

    const pipelineEditorPane = document.getElementById("pipeline-editor-pane");
    if (pipelineEditorPane !== null) {
      this.state.pipelines.push(new PreludioPipeline(pipelineEditorPane, name));
    }
  }

  importTable() {}

  getNumberOfPipelines(): number {
    return this.state.pipelines.length;
  }

  getCodeEditorPaneElement(): HTMLDivElement {
    return this.codeEditorPaneElement;
  }
}

window.customElements.define("app-element", App, {
  extends: "div",
});
