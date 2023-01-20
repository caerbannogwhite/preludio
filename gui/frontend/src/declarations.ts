import PreludioPipeline from "./PreludioPipeline";

export type LangDictType = { [key: string]: string };

export type SavedAppState = {
  pipelines: Array<PreludioPipeline>;
};
