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

function parseDate(value) {
  // format D(D)/M(M)/(YY)YY
  const dateFormat = /^\d{1,4}[.|/|-]\d{1,2}[.|/|-]\d{1,4}$/;

  if (dateFormat.test(value)) {
    // remove any leading zeros from date values
    const s = value.replace(/0*(\d*)/gi, '$1');
    const dateArray = s.split(/[.|/|-]/);

    // correct month value
    dateArray[1] = parseInt(dateArray[1], 10) - 1;

    // convert day / year to integers
    dateArray[0] = parseInt(dateArray[0], 10);
    dateArray[2] = parseInt(dateArray[2], 10);

    // correct year value
    if (dateArray[2].length < 4) {
      // correct year value
      dateArray[2] = (parseInt(dateArray[2], 10) < 50) ? 2000 + parseInt(dateArray[2], 10) : 1900 + parseInt(dateArray[2], 10);
    }

    const testDate = new Date(dateArray[2], dateArray[1], dateArray[0]);
    if (testDate.getDate() !== dateArray[0] || testDate.getMonth() !== dateArray[1] || testDate.getFullYear() !== dateArray[2]) {
      return null;
    }
    return testDate;
  }
  return null;
}

function IsDate(field) {
  field.clearError();

  if (parseDate(field.value) == null) {
    field.setError('Invalid date (DD/MM/YYYY)');
  }

  return field.isValid();
}

function MinimumAge(field, age) {
  field.clearError();

  const date = parseDate(field.value);
  const now = new Date();
  const minAge = new Date().setFullYear(now.getFullYear() - age);

  if (date > minAge) {
    field.setError('Must be at least 18 years old');
  }

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
  IsDate,
  Password,
  MinimumAge,
};
