function Validate(state) {
  if (state.account.status !== 'login_true') {
    return 'User must be logged in before a bet can be placed';
  }

  return '';
}

export default {
  Validate,
};

