import "./style.css";
import "./app.css";

import App from "./App";
import { CURRENT_DICTIONARY } from "./declarations";
// import dictionary from "./assets/dictionary.json";
// import functionsList from "./assets/preludio-functions.json";
// import { LangDictType, PreludioFunctionsList } from "./declarations";

// import { Greet } from "../wailsjs/go/main/App";

// Setup the greet function
// window.greet = function () {
//   // Get name
//   let name = nameElement!.value;

//   // Check if the input is empty
//   if (name === "") return;

//   // Call App.Greet(name)
//   try {
//     Greet(name)
//       .then((result) => {
//         // Update result with data back from App.Greet()
//         resultElement!.innerText = result;
//       })
//       .catch((err) => {
//         console.error(err);
//       });
//   } catch (err) {
//     console.error(err);
//   }
// };

// let nameElement = document.getElementById("name") as HTMLInputElement;
// nameElement.focus();
// let resultElement = document.getElementById("result");

// export const CURRENT_DICTIONARY: LangDictType = dictionary;
// export const PRELUDIO_FUNCTIONS_LIST: PreludioFunctionsList = functionsList;

function loadDictionary() {
  for (let entry of Object.entries(CURRENT_DICTIONARY)) {
    const e = document.getElementById(entry[0]);
    if (e !== null) {
      e.innerHTML = entry[1];
    }
  }
}

loadDictionary();

const app = new App();

const createNewPipelineButton = document.getElementById(
  "add-new-pipeline-button"
);
if (createNewPipelineButton !== null) {
  createNewPipelineButton.addEventListener("click", () => app.addNewPipeline());
}

declare global {
  interface Window {
    // greet: () => void;
  }
}
