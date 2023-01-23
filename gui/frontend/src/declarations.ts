import dictionary from "./assets/dictionary.json";
import functionsList from "./assets/preludio-functions.json";

export type LangDictType = { [key: string]: string };

export type PreludioFunctionParam = { name: string; type: string; num: number };
export type PreludioFunction = {
  name: string;
  docs: string;
  params: Array<PreludioFunctionParam>;
};
export type PreludioFunctionsList = { [key: string]: PreludioFunction };

export const CURRENT_DICTIONARY: LangDictType = dictionary;
export const PRELUDIO_FUNCTIONS_LIST: PreludioFunctionsList = functionsList;
