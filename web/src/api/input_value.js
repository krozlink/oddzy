class InputValue {
  constructor(name, validator, options) {
    this.name = name;
    this.value = '';
    this.validate = validator;
    this.active = false;
    this.error = '';
    this.isValid = true;

    if (options !== undefined && options.placeholder !== undefined) {
      this.placeholder = options.placeholder;
    } else {
      this.placeholder = name;
    }

    if (options !== undefined && options.type !== undefined) {
      this.type = options.type;
    } else {
      this.type = 'text';
    }
  }

  reset() {
    this.value = '';
    this.active = false;
    this.error = '';
    this.isValid = true;
  }

  showError() {
    return this.active && this.error.trim() !== '';
  }

  activate() {
    this.active = true;
  }
}

export default InputValue;
