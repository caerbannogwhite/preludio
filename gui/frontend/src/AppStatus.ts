import { LangDictType } from "./declarations";
import PreludioPipeline from "./PreludioPipeline";

class AppStatus {
  private static pipelineCounter: number = 0;
  private pipelines: Array<PreludioPipeline>;

  constructor(dictionary: LangDictType) {
    this._loadDictionary(dictionary);
    this._initCodeEditorPane();
    this.pipelines = new Array<PreludioPipeline>();
  }

  private _loadDictionary(dictionary: LangDictType) {
    for (let entry of Object.entries(dictionary)) {
      const e = document.getElementById(entry[0]);
      if (e !== null) {
        e.innerHTML = entry[1];
      }
    }
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
}

export default AppStatus;
