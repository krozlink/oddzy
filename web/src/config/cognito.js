export default {
  region: process.env.VUE_APP_AWS_REGION,
  IdentityPoolId: process.env.VUE_APP_AWS_COGNITO_IDENTITY_POOL_ID,
  UserPoolId: process.env.VUE_APP_AWS_COGNITO_USER_POOL_ID,
  ClientId: process.env.VUE_APP_AWS_COGNITO_CLIENT_ID,
};
