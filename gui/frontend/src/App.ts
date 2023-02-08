// import { EditorView, basicSetup } from "codemirror";
import { ParseCsv } from "../wailsjs/go/main/App";
import { AppState } from "./AppState";
import { PreludioPipeline } from "./PreludioPipeline";
import { MainButton } from "./utils/MainButton";
// import * as monaco from "monaco-editor";

export class App extends HTMLDivElement {
  private topMenuPaneElement: HTMLDivElement;
  private mainPaneElement: HTMLDivElement;
  private codeEditorPaneElement: HTMLDivElement;
  // private pipelineEditorPaneElement: HTMLDivElement;
  private tableEditorPaneElement: HTMLDivElement;

  private state: AppState;

  constructor() {
    super();
    this.id = "app";

    this.topMenuPaneElement = document.createElement("div");
    this.mainPaneElement = document.createElement("div");
    this.codeEditorPaneElement = document.createElement("div");
    // this.pipelineEditorPaneElement = document.createElement("div");
    this.tableEditorPaneElement = document.createElement("div");

    this.state = new AppState();

    this._initHTMLElement();
  }

  private _initHTMLElement() {
    this.topMenuPaneElement.id = "top-menu-pane";
    this.topMenuPaneElement.className = "menu-pane";

    this.mainPaneElement.id = "main-pane";

    // Code Editor Pane
    this.codeEditorPaneElement.id = "code-editor-pane";

    // monaco.editor.create(this.codeEditorPaneElement, {
    //   value: "function hello() {\n\talert('Hello world!');\n}",
    //   language: "javascript",
    // });

    // const editor = new EditorView({
    //   extensions: [basicSetup],
    //   parent: this.codeEditorPaneElement,
    // });

    // console.log(editor);

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
      console.log("hello");
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

    // formData.append("file", fileInput.files);

    // const xhr = new XMLHttpRequest();
    // xhr.open("POST", "your-server-url");
    // xhr.send(formData);

    const importTableButton = new MainButton("import-table", "import-svgrepo-com", fileInput);

    // this.tableEditorPaneElement.appendChild(filePathInput);
    this.tableEditorPaneElement.appendChild(importTableButton);

    // Main Panel
    this.mainPaneElement.appendChild(this.codeEditorPaneElement);
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
