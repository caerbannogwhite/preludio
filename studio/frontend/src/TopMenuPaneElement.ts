import { App, Theme } from "./App";
import "./TopMenuPaneElement.css";
import { Icon } from "./utils/Icon";

export class ButtonBar extends HTMLDivElement {
  constructor(id: string) {
    super();
    this.id = `${id}-button-bar`;
    this.className = "button-bar";
  }
}

window.customElements.define("button-bar-element", ButtonBar, {
  extends: "div",
});

export const TopMenuPaneElement = (id: string, app: App) => {
  // app.switchTheme();
  const menu = document.createElement("div");
  menu.id = id;
  menu.className = "menu-pane";

  const buttonBar1 = new ButtonBar("1");

  const runAllButton = document.createElement("button");
  runAllButton.id = "run-all-button";
  runAllButton.appendChild(new Icon("run-all-svgrepo-com"));

  const decreaseTextSizeButton = document.createElement("button");
  decreaseTextSizeButton.id = "run-all-button";
  decreaseTextSizeButton.appendChild(new Icon("zoom-out-svgrepo-com"));

  const increaseTextSizeButton = document.createElement("button");
  increaseTextSizeButton.id = "run-all-button";
  increaseTextSizeButton.appendChild(new Icon("zoom-in-svgrepo-com"));

  const switchThemeButton = document.createElement("button");
  switchThemeButton.id = "run-all-button";
  switchThemeButton.appendChild(new Icon("symbol-color-svgrepo-com"));
  switchThemeButton.addEventListener("click", () => {
    app.switchTheme();
  });

  const squirrelButton = document.createElement("button");
  squirrelButton.id = "run-all-button";
  squirrelButton.appendChild(new Icon("squirrel-svgrepo-com"));

  buttonBar1.appendChild(runAllButton);
  buttonBar1.appendChild(decreaseTextSizeButton);
  buttonBar1.appendChild(increaseTextSizeButton);
  buttonBar1.appendChild(switchThemeButton);
  buttonBar1.appendChild(squirrelButton);

  menu.appendChild(buttonBar1);

  return menu;
};
