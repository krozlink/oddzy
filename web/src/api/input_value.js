class InputValue {
  constructor(name, validators, options) {
    this.name = name;
    this.value = '';
    this.validators = validators;
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

  validate() {
    for (let i = 0; i < this.validators.length; i += 1) {
      this.validators[i](this);

      // Only one error message can be displayed so stop validating as soon as an error is found
      if (!this.isValid()) return false;
    }
    return this.isValid();
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
