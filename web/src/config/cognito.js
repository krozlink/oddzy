export default {
  region: process.env.VUE_APP_AWS_REGION,
  UserPoolId: process.env.VUE_APP_AWS_COGNITO_USER_POOL_ID,
  ClientId: process.env.VUE_APP_AWS_COGNITO_CLIENT_ID,
};
