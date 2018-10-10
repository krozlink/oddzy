import { Config, CognitoIdentityCredentials } from 'aws-sdk';
import { CognitoUser, CognitoUserPool, AuthenticationDetails, CognitoUserAttribute } from 'amazon-cognito-identity-js';
import AWSConfig from '../config/cognito';

const PoolData = {
  UserPoolId: AWSConfig.UserPoolId,
  ClientId: AWSConfig.ClientId,
};

// isAuthenticated
// getCurrentUser
// logout
// authenticate
// register
// forgotPassword


function Register(user, password, fields) {
  return new Promise((resolve, reject) => {
    const userPool = new CognitoUserPool(PoolData);

    const attributes = Object.values(fields).filter(f => f.attribute_name !== null).map(f => ({
      Name: f.attribute_name,
      Value: f.getValue(),
    }));

    userPool.signUp(user, password, attributes, null, (err, result) => {
      if (err) {
        reject(err);
      } else {
        resolve(result);
      }
    });
  });
}

function Authenticate(user, password) {
  return new Promise((resolve, reject) => {
    const authData = new AuthenticationDetails({
      Username: user,
      Password: password,
    });

    const userPool = new CognitoUserPool(PoolData);

    const cognitoUser = new CognitoUser({
      Username: user,
      Pool: userPool,
    });

    cognitoUser.authenticateUser(authData, {
      onSuccess: (result) => {
        resolve(result);
      },
      onFailure: (err) => {
        reject(err);
      },
    });
  });
}

export default {
  Register,
  Authenticate,
};
