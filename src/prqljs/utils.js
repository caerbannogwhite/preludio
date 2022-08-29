const PRQL_META = {
  table: 'table',
  column: 'column',
  scalar: 'scalar'
};

const PRQL_TYPE = {
  bool: 'bool',
  num: 'number',
  str: 'string'
};

class Table {
  constructor(columns = []) {
    this.names = [];
    this.types = [];
    for (let c in columns) {
      this.names.push(c.getName());
      this.types.push(c.getType());
    }
  }

  addColumn(col) {
    this.names.push(col.getName());
    this.types.push(col.getType());
  }

  addColumn(name, type) {
    this.names.push(name);
    this.types.push(type);
  }

  getColumnNames() {
    return this.names;
  }

  getColumnTypes() {
    return this.types;
  }

  getMatchingColumnNames(e) {

  }
}

class Column {
  constructor(name, type) {
    this.name = name;
    this.type = type;
  }

  getName() {
    return this.name;
  }

  getType() {
    return this.type;
  }

  setName(name) {
    this.name = name;
  }

  setType(type) {
    this.type = type;
  }
}

