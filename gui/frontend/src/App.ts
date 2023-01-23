import { AppState } from "./AppState";
import { PreludioPipeline } from "./PreludioPipeline";
import { MainButton } from "./utils/MainButton";

export class App extends HTMLDivElement {
  private topMenuPaneElement: HTMLDivElement;
  private mainPaneElement: HTMLDivElement;
  private codeEditorPaneElement: HTMLDivElement;
  private pipelineEditorPaneElement: HTMLDivElement;
  private tableEditorPaneElement: HTMLDivElement;

  private state: AppState;

  constructor() {
    super();
    this.id = "app";

    this.topMenuPaneElement = document.createElement("div");
    this.mainPaneElement = document.createElement("div");
    this.codeEditorPaneElement = document.createElement("div");
    this.pipelineEditorPaneElement = document.createElement("div");
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
    this.pipelineEditorPaneElement.id = "pipeline-editor-pane";

    const addNewPipelineButton = new MainButton(
      "add-new-pipeline",
      "plus-square-svgrepo-com"
    );
    addNewPipelineButton.addEventListener("click", () => this.addNewPipeline());

    this.codeEditorPaneElement.appendChild(addNewPipelineButton);
    this.codeEditorPaneElement.appendChild(this.pipelineEditorPaneElement);

    // Table Editor Pane
    this.tableEditorPaneElement.id = "table-editor-pane";
    this.tableEditorPaneElement.appendChild(
      new MainButton("import-table", "import-svgrepo-com")
    );

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

  getNumberOfPipelines(): number {
    return this.state.pipelines.length;
  }
}

window.customElements.define("app-element", App, {
  extends: "div",
});
