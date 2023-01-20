import PreludioPipeline from "./PreludioPipeline";

export type LangDictType = { [key: string]: string };

export type PreludioFunctionParam = { name: string; type: string; num: number };
export type PreludioFunction = { name: string; docs: string; params: Array<PreludioFunctionParam> };
export type PreludioFunctionsList = { [key: string]: PreludioFunction };

export type SavedAppState = {
  pipelines: Array<PreludioPipeline>;
};
