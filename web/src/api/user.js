class User {
  constructor(username, userAttributes) {
    this.username = username;
    this.firstName = userAttributes.given_name.value;
  }
}

export default User;
