class InputValue {
  constructor(description, name, autocomplete, attributeName, validators, options) {
    this.name = name;
    this.autocomplete = autocomplete;
    this.description = description;
    this.attribute_name = attributeName;
    this.raw_value = '';
    this.validators = validators;
    this.active = false;
    this.error = '';
    this.parser = null;
    this.placeholder = description;
    this.type = 'text';

    if (options !== undefined && options.parser !== undefined) {
      this.parser = options.parser;
    }

    if (options !== undefined && options.placeholder !== undefined) {
      this.placeholder = options.placeholder;
    }

    if (options !== undefined && options.type !== undefined) {
      if (options.type === 'checkbox') {
        this.raw_value = false;
      }
      this.type = options.type;
    }
  }

  validate() {
    if (!this.validators) return true;
    for (let i = 0; i < this.validators.length; i += 1) {
      this.validators[i](this);

      // Only one error message can be displayed so stop validating as soon as an error is found
      if (!this.isValid()) return false;
    }
    return this.isValid();
  }

  getValue() {
    if (!this.parser) {
      return this.raw_value;
    }

    return this.parser(this.raw_value);
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
    this.raw_value = '';
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
