class User {
  constructor(username, userAttributes) {
    this.username = username;

    this.first_name = userAttributes.find(a => a.Name === 'given_name').Value;
    this.last_name = userAttributes.find(a => a.Name === 'family_name').Value;

    this.email_address = userAttributes.find(a => a.Name === 'email').Value;
    this.email_verified = userAttributes.find(a => a.Name === 'email_verified').Value;

    this.mobile = userAttributes.find(a => a.Name === 'phone_number').Value;
    this.mobile_verified = userAttributes.find(a => a.Name === 'phone_number_verified').Value;

    this.address = userAttributes.find(a => a.Name === 'address').Value;
    this.date_of_birth = userAttributes.find(a => a.Name === 'birthdate').Value;
  }
}

export default User;
