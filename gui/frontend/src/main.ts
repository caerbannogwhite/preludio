import "./style.css";
import "./app.css";

import AppStatus from "./AppStatus";
import dictionary from "./assets/dictionary.json";
import { LangDictType } from "./declarations";

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

const currentDictionary: LangDictType = dictionary;

const appStatus = new AppStatus(currentDictionary);

const createNewPipelineButton = document.getElementById(
  "add-new-pipeline-label"
);
if (createNewPipelineButton !== null) {
  createNewPipelineButton.addEventListener("click", () => {
    appStatus.addNewPipeline();
  });
}

declare global {
  interface Window {
    // greet: () => void;
    createNewPipeline: () => void;
  }
}
