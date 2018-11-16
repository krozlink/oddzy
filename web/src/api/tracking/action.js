class Action {
  constructor(area, type, actionData) {
    this.timestamp = new Date().getTime();
    this.area = area; // e.g Betslip, Register, Login
    this.type = type; // e.g Username field
    this.data = actionData;
  }
}

export default Action;
