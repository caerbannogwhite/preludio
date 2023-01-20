import "./style.css";
import "./app.css";

import { Greet } from "../wailsjs/go/main/App";

// Setup the greet function
window.greet = function () {
  // Get name
  let name = nameElement!.value;

  // Check if the input is empty
  if (name === "") return;

  // Call App.Greet(name)
  try {
    Greet(name)
      .then((result) => {
        // Update result with data back from App.Greet()
        resultElement!.innerText = result;
      })
      .catch((err) => {
        console.error(err);
      });
  } catch (err) {
    console.error(err);
  }
};

let nameElement = document.getElementById("name") as HTMLInputElement;
nameElement.focus();
let resultElement = document.getElementById("result");

declare global {
  interface Window {
    greet: () => void;
  }
}

class PreludioPipeline {
  private body: HTMLElement;
  private funcionList: Array<PipelineFunctionCall>;

  constructor() {
    this.funcionList = Array<PipelineFunctionCall>();

    this.body = document.createElement("div");
    this.body.className = `pipeline-body`;
  }

  appendFunction(funcCall: PipelineFunctionCall) {
    this.funcionList.push(funcCall);
    this.body.appendChild(funcCall.getBody());
  }
}

class PipelineFunctionCall {
  private body: HTMLElement;

  constructor(funcName: string) {
    const label = document.createElement("label");
    label.id = `${funcName}-label-name`;
    label.className = `function-call-label-name`;
    label.innerHTML = funcName;

    const param1 = document.createElement("input");
    param1.id = `${funcName}-param1-input`;
    param1.className = `function-call-param-input`;

    this.body = document.createElement("div");
    this.body.className = `function-call-body`;
    this.body.appendChild(label);
    this.body.appendChild(param1);
  }

  getBody(): HTMLElement {
    return this.body;
  }
}
