import "./TableEditorPane.css";

export class TableEditorPane extends HTMLDivElement {
  private table: HTMLTableElement;

  constructor() {
    super();
    this.id = `table-editor-pane`;
    this.className = "";

    this.table = document.createElement("table");

    this._initHTMLElement();
  }

  private _initHTMLElement() {
    this.table.className = "table-view";
  }

  setTableValues(values: string[][]) {
    this.table.innerHTML = "";
    values.forEach((row) => {
      const tr = document.createElement("tr");
      row.forEach((value) => {
        const td = document.createElement("td");
        td.innerText = value;
        tr.appendChild(td);
      });
      this.table.appendChild(tr);
    });
  }
}

window.customElements.define("table-editor-pane", TableEditorPane, {
  extends: "div",
});
