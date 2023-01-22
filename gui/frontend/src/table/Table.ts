import "./table.css";

export class Table extends HTMLDivElement {
  constructor() {
    super();

    this._initHTMLElement();
  }

  private _initHTMLElement() {
    const tab = document.createElement("table");

    const header = document.createElement("tr");

    tab.appendChild(header);

    this.appendChild(tab);
  }
}
