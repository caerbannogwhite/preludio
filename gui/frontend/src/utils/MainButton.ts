import { CURRENT_DICTIONARY } from "../declarations";
import { Icon } from "./Icon";
import "./MainButton.css";

export class MainButton extends HTMLDivElement {
  private elementId: string;
  private fileInput: HTMLInputElement | undefined;

  constructor(id: string, iconName?: string, fileInput?: HTMLInputElement) {
    super();
    this.elementId = id;
    this.id = `${id}-button`;
    this.fileInput = fileInput;

    this._initHTMLElement(iconName);
  }

  private _initHTMLElement(iconName?: string) {
    this.className = "main-button";

    if (iconName !== undefined) {
      this.appendChild(new Icon(iconName));
    }

    const label = document.createElement("label");
    label.id = `${this.elementId}-label`;
    label.className = "main-button-label";

    const text = CURRENT_DICTIONARY[this.elementId];
    if (text !== undefined) {
      label.innerHTML = text;
    }

    if (this.fileInput !== undefined) {
      label.htmlFor = this.fileInput.id;
      this.appendChild(this.fileInput);
    }

    this.appendChild(label);
  }
}

window.customElements.define("main-button", MainButton, {
  extends: "div",
});
