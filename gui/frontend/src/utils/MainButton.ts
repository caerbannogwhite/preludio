import { CURRENT_DICTIONARY } from "../declarations";
import "./MainButton.css";

export class MainButton extends HTMLDivElement {
  private elementId: string;
  constructor(id: string, iconSrc?: string) {
    super();
    this.elementId = id;
    this.id = `${id}-button`;

    this._initHTMLElement(iconSrc);
  }

  private _initHTMLElement(iconSrc?: string) {
    this.className = "main-button";

    if (iconSrc !== undefined) {
      const img = document.createElement("img");
      img.className = "icon";
      img.src = iconSrc;

      this.appendChild(img);
    }

    const label = document.createElement("label");
    label.id = `${this.elementId}-label`;
    label.className = "main-button-label";

    const text = CURRENT_DICTIONARY[this.id];
    if (text !== undefined) {
      label.innerHTML = text;
    }

    this.appendChild(label);
  }
}

window.customElements.define("main-button", MainButton, {
  extends: "div",
});
