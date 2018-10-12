class User {
  constructor(username, userAttributes) {
    this.username = username;
    this.firstName = userAttributes.find(a => a.Name === 'given_name').Value;
    this.lastName = userAttributes.find(a => a.Name === 'family_name').Value;
    this.email = userAttributes.find(a => a.Name === 'email').Value;
  }
}

export default User;
