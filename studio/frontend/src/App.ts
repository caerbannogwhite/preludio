// import { EditorView, basicSetup } from "codemirror";
import { ParseCsv } from "../wailsjs/go/main/App";
import { AppState } from "./AppState";
import { PreludioPipeline } from "./PreludioPipeline";
import { TopMenuPaneElement } from "./TopMenuPaneElement";
// import { MainButton } from "./utils/MainButton";
import * as monaco from "monaco-editor";

const STARTING_CODE_SAMPLE = [
  `importCSV "C:\\\\Users\\\\massi\\\\Downloads\\\\Cars.csv" delimiter: ";" header:true`,
  `derive [`,
  `  [MPG, Displacement] = ([MPG, Displacement] | strReplace "," "." | asFloat),`,
  `  num = Cylinders * 2 - 1.123e-1,`,
  `  Car_Origin = Car + " - " + Origin`,
  `]`,
  `describe`,
  `take 20`,
  `select [Car_Origin, MPG, Cylinders, num]`,
  `exportCSV "C:\\\\Users\\\\massi\\\\Downloads\\\\Cars1.csv" delimiter: '\\t'`,
];

export enum Theme {
  ligth = "vs-light",
  dark = "vs-dark",
}

type AppSettings = {
  theme: Theme;
};

const DEFAULT_SETTINGS = {
  theme: Theme.dark,
};

export class App extends HTMLDivElement {
  private topMenuPaneElement: HTMLDivElement;
  private mainPaneElement: HTMLDivElement;
  private leftPaneElement: HTMLDivElement;
  private codeEditorPaneElement: HTMLDivElement;
  private logErrorPaneElement: HTMLDivElement;
  private tableEditorPaneElement: HTMLDivElement;

  private settings: AppSettings;
  private state: AppState;
  private loaded: boolean;

  constructor() {
    super();
    this.id = "app";
    this.loaded = false;

    this.settings = DEFAULT_SETTINGS;
    this.state = new AppState();

    this.topMenuPaneElement = TopMenuPaneElement("top-menu-pane", this);
    this.mainPaneElement = document.createElement("div");
    this.leftPaneElement = document.createElement("div");
    this.codeEditorPaneElement = document.createElement("div");
    this.logErrorPaneElement = document.createElement("div");
    this.tableEditorPaneElement = document.createElement("div");

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

    this.loaded = true;
    this.setTheme(this.settings.theme);
  }

  // Switch theme between Light theme and Dark theme
  switchTheme() {
    if (!this.loaded) {
      return;
    }

    if (this.settings.theme !== Theme.ligth) {
      this.setTheme(Theme.ligth);
    } else {
      this.setTheme(Theme.dark);
    }
  }

  // Set the them for the whole app
  setTheme(theme: Theme) {
    if (!this.loaded) {
      return;
    }

    switch (theme) {
      case Theme.dark:
        this.className = "dark-mode";
        this.logErrorPaneElement.className = "dark-mode";
        this.tableEditorPaneElement.className = "dark-mode";
        monaco.editor.setTheme(theme);
        break;
      // case Theme.ligth:
      default:
        this.className = "light-mode";
        this.logErrorPaneElement.className = "ligtht-mode";
        this.tableEditorPaneElement.className = "light-mode";
        monaco.editor.setTheme(theme);
        break;
    }

    this.settings.theme = theme;
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

  pageLoaded() {
    monaco.editor.create(this.codeEditorPaneElement, {
      value: STARTING_CODE_SAMPLE.join("\n"),
      language: "python",
      minimap: { enabled: false },
    });

    monaco.editor.setTheme(this.settings.theme);
  }
}

window.customElements.define("app-element", App, {
  extends: "div",
});
