import DateHelper from './date-helper';

function Mandatory(field) {
  field.clearError();
  if (field.raw_value.trim() === '') {
    field.setError(`${field.description} is mandatory`);
  }

  return field.isValid();
}

function Match(password, confirm) {
  confirm.clearError();
  if (confirm.raw_value !== password.raw_value) {
    confirm.setError('Passwords do not match');
  }

  return confirm.isValid();
}

function IsTrue(field, message) {
  field.clearError();
  if (!field.raw_value) {
    field.setError(message);
  }
  return field.isValid();
}

function EmailAddress(field) {
  field.clearError();

  const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  if (!re.test(field.raw_value.trim().toLowerCase())) {
    field.setError('Not a valid email address');
  }

  return field.isValid();
}

function MobileNumber(field) {
  field.clearError();

  const re = /^04[0-9]{8}$/;
  if (!re.test(field.raw_value.trim())) {
    field.setError('Invalid mobile (04xxxxxxxx)');
  }

  return field.isValid();
}


function IsDate(field) {
  field.clearError();

  if (DateHelper.parseDateString(field.raw_value) == null) {
    field.setError('Invalid date (DD/MM/YYYY)');
  }

  return field.isValid();
}

function MinimumAge(field, age) {
  field.clearError();

  const date = DateHelper.parseDateString(field.raw_value);
  const now = new Date();
  const minAge = new Date().setFullYear(now.getFullYear() - age);

  if (date > minAge) {
    field.setError('Must be at least 18 years old');
  }

  return field.isValid();
}

function Password(field) {
  field.clearError();

  if (field.raw_value.trim().length < 8) {
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
  IsDate,
  Password,
  MinimumAge,
};
