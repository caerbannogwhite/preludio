class PreludioPipeline {
  private body: HTMLElement;
  private name: string;
  private funcionList: Array<PipelineFunctionCall>;

  constructor(name: string) {
    this.funcionList = Array<PipelineFunctionCall>();

    this.body = document.createElement("div");
    this.body.className = `pipeline-body`;

    this.name = name;
    const label = document.createElement("label");
    label.innerHTML = this.name;

    this.body.appendChild(label);
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
