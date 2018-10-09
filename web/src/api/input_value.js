class InputValue {
  constructor(name, validator, options) {
    this.name = name;
    this.value = '';
    this.validate = validator;
    this.active = false;
    this.error = '';

    if (options !== undefined && options.placeholder !== undefined) {
      this.placeholder = options.placeholder;
    } else {
      this.placeholder = name;
    }

    if (options !== undefined && options.type !== undefined) {
      if (options.type === 'checkbox') {
        this.value = false;
      }
      this.type = options.type;
    } else {
      this.type = 'text';
    }
  }

  setError(err) {
    this.error = err;
  }

  clearError() {
    this.error = '';
  }

  isValid() {
    return this.error === '';
  }

  reset() {
    this.value = '';
    this.active = false;
    this.error = '';
  }

  showError() {
    return this.active && !this.isValid();
  }

  activate() {
    this.active = true;
  }
}

export default InputValue;
