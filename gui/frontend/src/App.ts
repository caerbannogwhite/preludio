import { SavedAppState } from "./declarations";
import PreludioPipeline from "./PreludioPipeline";
import { MainButton } from "./utils/MainButton";

export class App extends HTMLDivElement {
  private topMenuPaneElement: HTMLDivElement;
  private mainPaneElement: HTMLDivElement;
  private codeEditorPaneElement: HTMLDivElement;
  private pipelineEditorPaneElement: HTMLDivElement;
  private tableEditorPaneElement: HTMLDivElement;

  private static pipelineCounter: number = 0;
  private pipelines: Array<PreludioPipeline>;

  constructor(savedStatus?: SavedAppState) {
    super();
    this.id = "app";

    this.topMenuPaneElement = document.createElement("div");
    this.mainPaneElement = document.createElement("div");
    this.codeEditorPaneElement = document.createElement("div");
    this.pipelineEditorPaneElement = document.createElement("div");
    this.tableEditorPaneElement = document.createElement("div");

    this._initHTMLElement();

    this.pipelines = new Array<PreludioPipeline>();

    if (savedStatus !== undefined) {
      this._loadSavedStatus(savedStatus);
    }
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
      "src/assets/icons/plus-square-svgrepo-com.svg"
    );
    addNewPipelineButton.addEventListener("click", () => this.addNewPipeline());

    this.codeEditorPaneElement.appendChild(addNewPipelineButton);
    this.codeEditorPaneElement.appendChild(this.pipelineEditorPaneElement);

    // Table Editor Pane
    this.tableEditorPaneElement.id = "table-editor-pane";
    this.tableEditorPaneElement.appendChild(
      new MainButton("import-table", "src/assets/icons/import-svgrepo-com.svg")
    );

    // Main Panel
    this.mainPaneElement.appendChild(this.codeEditorPaneElement);
    this.mainPaneElement.appendChild(this.tableEditorPaneElement);

    // App
    this.appendChild(this.topMenuPaneElement);
    this.appendChild(this.mainPaneElement);
  }

  private _loadSavedStatus(savedStatus: SavedAppState) {
    this.pipelines = savedStatus.pipelines;
  }

  addNewPipeline(pipelineName?: string) {
    App.pipelineCounter++;
    let name = `Pipeline ${App.pipelineCounter}`;
    if (pipelineName !== undefined) {
      name = pipelineName;
    }

    const pipelineEditorPane = document.getElementById("pipeline-editor-pane");
    if (pipelineEditorPane !== null) {
      this.pipelines.push(new PreludioPipeline(pipelineEditorPane, name));
    }
  }

  getNumberOfPipelines(): number {
    return this.pipelines.length;
  }

  exportStatus(): SavedAppState {
    return {
      pipelines: this.pipelines,
    };
  }
}

window.customElements.define("app-element", App, {
  extends: "div",
});
