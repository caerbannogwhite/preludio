import "./Icon.css";

export class Icon extends HTMLImageElement {
  constructor(iconName: string) {
    super();
    this.className = "icon";
    this.src = `src/assets/icons/${iconName}.svg`;
  }
}

window.customElements.define("icon-element", Icon, {
  extends: "img",
});
