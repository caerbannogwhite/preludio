import "./DropDownOptionsMenu.css";

export type DropDownOption = { value: string; name: string };

export class DropDownOptionsMenu extends HTMLDivElement {
  private options: Array<HTMLOptionElement>;
  private textInputElement: HTMLInputElement;
  private optionsListElement: HTMLDivElement;
  private visibleOptionsNum: number;

  constructor(id: string, options: Array<DropDownOption>) {
    super();
    this.id = id;
    this.visibleOptionsNum = 0;
    this.options = new Array<HTMLOptionElement>();

    this.textInputElement = document.createElement("input");
    this.optionsListElement = document.createElement("div");

    this._initHTMLElement(options);
  }

  private _initHTMLElement(options: Array<DropDownOption>) {
    this.textInputElement.id = `${this.id}-text-input`;
    this.textInputElement.className = "drop-down-input";
    this.textInputElement.addEventListener("keydown", () => {
      this._cleanList();
    });
    this.textInputElement.addEventListener("keyup", () => {
      this._checkTextInput();
    });

    this.optionsListElement.id = `${this.id}-drop-down-options-list`;
    this.optionsListElement.className = "drop-down-options-list";
    this.optionsListElement.style.display = "none";

    for (let option of options) {
      const optionElement = document.createElement("option");
      optionElement.value = option.value;
      optionElement.innerHTML = option.name;
      optionElement.className = "drop-down-option";
      optionElement.style.display = "none";

      this.optionsListElement.appendChild(optionElement);
      this.options.push(optionElement);
    }

    this.appendChild(this.textInputElement);
    this.appendChild(this.optionsListElement);
  }

  private _checkTextInput() {
    const searchText = this.textInputElement.value.trim().toLocaleLowerCase();
    this.visibleOptionsNum = 0;
    if (searchText === "") {
      for (let option of this.options) {
        option.style.display = "none";
      }
    } else {
      for (let option of this.options) {
        if (option.innerHTML.toLocaleLowerCase().startsWith(searchText)) {
          this.visibleOptionsNum++;
          option.style.display = "block";
        } else {
          option.style.display = "none";
        }
      }
    }

    if (this.visibleOptionsNum > 0) {
      this._resetListElementPosition();
      this.optionsListElement.style.display = "block";
    } else {
      this.optionsListElement.style.display = "none";
    }
  }

  private _cleanList() {
    if (this.textInputElement.value.trim().toLocaleLowerCase() === "") {
      this.visibleOptionsNum = 0;
      for (let option of this.options) {
        option.style.display = "none";
      }
    }

    if (this.visibleOptionsNum > 0) {
      this._resetListElementPosition();
      this.optionsListElement.style.display = "block";
    } else {
      this.optionsListElement.style.display = "none";
    }
  }

  private _resetListElementPosition() {
    this.optionsListElement.style.top = `${
      this.textInputElement.getBoundingClientRect().bottom // + window.screenY
    }px`;
    this.optionsListElement.style.left = `${
      this.textInputElement.getBoundingClientRect().left // + window.screenX
    }px`;
    this.optionsListElement.style.width = `${
      this.textInputElement.getBoundingClientRect().width - 40 // 40 is the magic number!!
    }px`;
  }
}

window.customElements.define("drop-down-options-menu", DropDownOptionsMenu, {
  extends: "div",
});
