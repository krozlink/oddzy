function Mandatory(field) {
  field.clearError();
  if (field.value.trim() === '') {
    field.setError(`${field.name} is mandatory`);
  }

  return field.isValid();
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
  return '';
}

export default {
  Mandatory,
  EmailAddress,
  IsTrue,
};
