import { Config, CognitoIdentityCredentials } from 'aws-sdk';
import { CognitoUser, CognitoUserPool, AuthenticationDetails, CognitoUserAttribute } from 'amazon-cognito-identity-js';
import AWSConfig from '../config/cognito';

// isAuthenticated
// getCurrentUser
// logout
// authenticate
// register
// forgotPassword

function Register(fields) {
  return new Promise((resolve, reject) => {
    const poolData = {
      UserPoolId: AWSConfig.UserPoolId,
      ClientId: AWSConfig.ClientId,
    };
    const userPool = new CognitoUserPool(poolData);

    const attributes = Object.values(fields).filter(f => f.attribute_name !== null && f.attribute_name !== 'preferred_username').map(f => ({
      Name: f.attribute_name,
      Value: f.getValue(),
    }));

    userPool.signUp(fields.user_name.getValue(), fields.password.getValue(), attributes, null, (err, result) => {
      if (err) {
        reject(err);
      } else {
        resolve(result);
      }
    });
  });
}

export default {
  Register,
};
