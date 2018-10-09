function Mandatory(field) {
  field.clearError();
  if (field.value.trim() === '') {
    field.setError(`${field.name} is mandatory`);
  }

  return field.isValid();
}

function Match(password, confirm) {
  confirm.clearError();
  if (confirm.value !== password.value) {
    confirm.setError('Passwords do not match');
  }

  return confirm.isValid();
}

function IsTrue(field, message) {
  field.clearError();
  if (!field.value) {
    field.setError(message);
  }
  return field.isValid();
}

function EmailAddress(field) {
  field.clearError();

  const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  if (!re.test(field.value.trim().toLowerCase())) {
    field.setError('Not a valid email address');
  }

  return field.isValid();
}

function MobileNumber(field) {
  field.clearError();

  const re = /^04[0-9]{8}$/;
  if (!re.test(field.value.trim())) {
    field.setError('Invalid mobile (04xxxxxxxx)');
  }

  return field.isValid();
}

function Date(field) {
  field.clearError();

  const re = /^[0-9]{2}\/[0-9]{2}\/[0-9]{4}$/;
  if (!re.test(field.value.trim())) {
    field.setError('Invalid date (DD/MM/YYYY)');
  }

  return field.isValid();
}

function MinimumAge(field, age) {
  field.clearError();

  field.setError('Must be at least 18 years old');

  return field.isValid();
}

function Password(field) {
  field.clearError();

  if (field.value.trim().length < 8) {
    field.setError('Must be at least 8 letters');
  }

  return field.isValid();
}

export default {
  Mandatory,
  Match,
  EmailAddress,
  MobileNumber,
  IsTrue,
  Date,
  Password,
  MinimumAge,
};
