import "./style.css";

import { App } from "./App";

document.body.appendChild(new App());

// Default Wails code

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

// declare global {
//   interface Window {
//     greet: () => void;
//   }
// }
