import PreludioPipeline from "./PreludioPipeline";

export type LangDictType = { [key: string]: string };

export type SavedAppStatus = {
  pipelines: Array<PreludioPipeline>;
};
