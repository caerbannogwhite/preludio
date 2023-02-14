import { ImportCsv, ImportExcel } from "../../wailsjs/go/main/App";

type SeriesType = boolean | number | string | DataFrame;
type IndexElement = { i: number; v: SeriesType[] };

export class Schema {
  public name: string;
  public type: SeriesType;

  constructor(name: string, type: SeriesType) {
    this.name = name;
    this.type = type;
  }
}

export class Series extends Schema {
  private _values: SeriesType[];

  constructor(name: string, type: SeriesType, values: SeriesType[]) {
    super(name, type);
    this._values = values;
  }

  getLength(): number {
    return this._values.length;
  }

  getValues(): SeriesType[] {
    return this._values;
  }

  setValues(vals: SeriesType[]) {
    this._values = vals;
  }

  take(num: number) {
    if (num > 0) {
      this._values = this._values.slice(0, num);
    } else if (num < 0) {
      this._values = this._values.slice(num);
    }
  }

  sort(): number[] {
    let valuesWithIndex = new Array<IndexElement>(this._values.length);
    for (let i = 0; i < this._values.length; ++i) {
      valuesWithIndex[i] = { i: i, v: [this._values[i]] };
    }

    valuesWithIndex = valuesWithIndex.sort(
      (a: IndexElement, b: IndexElement) => {
        return a.v[0] > b.v[0] ? 1 : -1;
      }
    );

    const indeces = new Array<number>(this._values.length);
    for (let i = 0; i < this._values.length; ++i) {
      indeces[i] = valuesWithIndex[i].i;
    }

    return indeces;
  }
}

export class DataFrame extends Array<Series> {
  private _parentElement: HTMLElement;
  private _names: string[];
  private _index: IndexElement[];

  constructor(parent: HTMLElement) {
    super();
    this._parentElement = parent;
    this._names = new Array<string>();
    this._index = new Array<IndexElement>();
  }

  empty() {
    this.length = 0;
    this._names.length = 0;
    this._index.length = 0;
  }

  getNumCols(): number {
    return this.length;
  }

  getNumRows(): number {
    return this[0].getLength();
  }

  fromCsv(path: string) {
    ImportCsv(path).then((d) => {
      console.log(d);
    });
  }

  fromExcel(path: string) {
    ImportExcel(path);
  }

  fromSeries(...series: Series[]) {
    this.empty();
    if (series.length > 0) {
      const l = series[0].getLength();
      for (let s of series) {
        if (s.getLength() !== l || this._names.includes(s.name)) {
          this.empty();
          return;
        }

        this._names.push(s.name);
        this.push(s);
      }

      // populate index
      this._index = new Array<IndexElement>(this.length);
      for (let i = 0; i < this.length; ++i) {
        this._index[i] = { i: i, v: [] };
      }
    }
  }

  withSchema(names: string[], types: SeriesType[]) {
    this.empty();
    if (names.length === types.length) {
      for (let i = 0; i < names.length; ++i) {
        this._names.push(names[i]);
        switch (types[i]) {
          case "boolean":
            this.push(new Series(names[i], types[i], new Array<boolean>()));
            break;
          case "number":
            this.push(new Series(names[i], types[i], new Array<number>()));
            break;
          case "string":
            this.push(new Series(names[i], types[i], new Array<string>()));
            break;
          case "DataFrame":
            this.push(new Series(names[i], types[i], new Array<DataFrame>()));
            break;
          default:
            this.push(new Series(names[i], types[i], new Array<any>()));
            break;
        }
      }
    }
  }

  addSeries(series: Series) {
    if (
      this[0].getLength() === series.getLength() ||
      !this._names.includes(series.name)
    ) {
      this._names.push(series.name);
      this.push(series);
    }
  }

  // VERBS

  groupBy(...names: string[]): DataFrame {
    for (let name of names) {
      const idx = this._names.indexOf(name);
      const values = this[idx].getValues();
      for (let i = 0; i < values.length; ++i) {
        this._index[i].v.push(values[i]);
      }
    }

    return this;
  }

  ungroup(): DataFrame {
    for (let i = 0; i < this.length; ++i) {
      this._index[i] = { i: i, v: [] };
    }

    return this;
  }

  orderBy(...names: string[]): DataFrame {
    this.groupBy(...names);
    const sorted = this._index.sort((a: IndexElement, b: IndexElement) => {
      for (let i = 0; i < a.v.length; ++i) {
        if (a.v[i] > b.v[i]) {
          return 1;
        } else if (a.v[i] < b.v[i]) {
          return -1;
        }
      }
      return 0;
    });

    // swap elements in all the series
    for (let s of this) {
      const vals = s.getValues();
      let tmp = vals[0];
      for (let i = 0; i < s.getLength(); ++i) {
        if (sorted[i].i !== i) {
          tmp = vals[sorted[i].i];
          vals[sorted[i].i] = vals[i];
          vals[i] = tmp;
        }
      }
      s.setValues(vals);
    }
    this.ungroup();

    return this;
  }

  derive(): DataFrame {
    return this;
  }

  select(...names: string[]): DataFrame {
    const newDataFrame = new DataFrame(this._parentElement);

    for (let name of names) {
      if (this._names.includes(name)) {
        const idx = this._names.indexOf(name);
        newDataFrame.addSeries(this[idx]);
      }
    }

    return newDataFrame;
  }

  take(num: number): DataFrame {
    for (let idx in this) {
      this[idx].take(num);
    }
    return this;
  }

  hasSameSchema(other: DataFrame): boolean {
    if (this.length === other.length) {
      for (let i = 0; i < this.length; ++i) {
        if (this[i].name !== other[i].name || this[i].type !== other[i].type) {
          return false;
        }
      }
      return true;
    }
    return false;
  }

  display() {
    const constainer = document.createElement("div");
    constainer.id = "dataframe-container";

    const table = document.createElement("table");

    constainer.appendChild(table);
    this._parentElement.appendChild(constainer);
  }
}
