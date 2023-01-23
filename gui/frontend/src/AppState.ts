import { PreludioPipeline } from "./PreludioPipeline";

export class AppState {
  static pipelineCounter: number = 0;
  pipelines: Array<PreludioPipeline>;

  constructor() {
    this.pipelines = new Array<PreludioPipeline>();
  }

  serialize(): string {
    const state = {
      pipelineCounter: AppState.pipelineCounter,
      pipelines: this.pipelines,
    };

    return JSON.stringify(state);
  }

  deserialize(json: string) {
    const obj = JSON.parse(json);

    AppState.pipelineCounter = obj.pipelineCounter;
    this.pipelines = obj.pipelines;
  }
}
