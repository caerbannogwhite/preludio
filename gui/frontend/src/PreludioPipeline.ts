class PreludioPipeline {
  private editorPane: HTMLElement;
  private body: HTMLElement;
  private name: string;
  private funcionList: Array<PipelineFunctionCall>;

  constructor(editorPane: HTMLElement, name: string) {
    this.editorPane = editorPane;
    this.name = name;

    this.funcionList = Array<PipelineFunctionCall>();

    this._initHTMLElement();
  }

  private _initHTMLElement() {
    this.body = document.createElement("div");
    this.body.id = `${this.name}-pipeline-body`;
    this.body.className = `pipeline-body`;

    const nameDiv = document.createElement("div");
    nameDiv.id = `${this.name}-name-div`;
    nameDiv.className = "pipeline-name";

    const icon = document.createElement("img");
    icon.className = "icon";
    icon.src = "src/assets/icons/pipe-svgrepo-com.svg";

    const label = document.createElement("label");
    label.id = `${this.name}-name-label`;
    label.className = "pipeline-name";
    label.innerHTML = this.name;

    nameDiv.appendChild(icon);
    nameDiv.appendChild(label);

    this.body.appendChild(nameDiv);

    this.editorPane.appendChild(this.body);
  }

  appendFunction(funcCall: PipelineFunctionCall) {
    this.funcionList.push(funcCall);
    this.body.appendChild(funcCall.getBody());
  }

  getName(): string {
    return this.name;
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

export default PreludioPipeline;
