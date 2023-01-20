import { LangDictType, SavedAppStatus } from "./declarations";
import PreludioPipeline from "./PreludioPipeline";

class AppStatus {
  private static pipelineCounter: number = 0;
  private pipelines: Array<PreludioPipeline>;

  constructor(dictionary: LangDictType, savedStatus?: SavedAppStatus) {
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

  private _loadSavedStatus(savedStatus: SavedAppStatus) {
    this.pipelines = savedStatus.pipelines;
  }

  // Initialize the code editor panel
  private _initCodeEditorPane() {
    if (this.pipelines.length === 0) {
    }
  }

  addNewPipeline(pipelineName?: string) {
    AppStatus.pipelineCounter++;
    let name = `Pipeline #${AppStatus.pipelineCounter}`;
    if (pipelineName !== undefined) {
      name = pipelineName;
    }
    
    this.pipelines.push(new PreludioPipeline(name));
  }

  getNumberOfPipelines(): number {
    return this.pipelines.length;
  }

  exportStatus(): SavedAppStatus {
    return {
      pipelines: this.pipelines,
    };
  }
}

export default AppStatus;
