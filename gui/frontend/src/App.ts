import { LangDictType, SavedAppState } from "./declarations";
import PreludioPipeline from "./PreludioPipeline";

class App {
  private static pipelineCounter: number = 0;
  private pipelines: Array<PreludioPipeline>;

  constructor(dictionary: LangDictType, savedStatus?: SavedAppState) {
    this.pipelines = new Array<PreludioPipeline>();

    this._loadDictionary(dictionary);
    if (savedStatus !== undefined) {
      this._loadSavedStatus(savedStatus);
    }
    this._initCodeEditorPane();
  }

  private _loadDictionary(dictionary: LangDictType) {
    for (let entry of Object.entries(dictionary)) {
      const e = document.getElementById(entry[0]);
      if (e !== null) {
        e.innerHTML = entry[1];
      }
    }
  }

  private _loadSavedStatus(savedStatus: SavedAppState) {
    this.pipelines = savedStatus.pipelines;
  }

  // Initialize the code editor panel
  private _initCodeEditorPane() {
    if (this.pipelines.length === 0) {
    }
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

export default App;
