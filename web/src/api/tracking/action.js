class Action {
  constructor(area, type, actionData, metaData) {
    this.timestamp = new Date();
    this.area = area;
    this.type = type;
    this.data = actionData;
    this.metadata = metaData;
  }
}

export default Action;
