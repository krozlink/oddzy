import { CognitoUser, CognitoUserPool, AuthenticationDetails, CognitoUserAttribute } from 'amazon-cognito-identity-js';
import AWSConfig from '../config/cognito';
import User from './user';

const PoolData = {
  UserPoolId: AWSConfig.UserPoolId,
  ClientId: AWSConfig.ClientId,
};

function Register(user, password, fields) {
  return new Promise((resolve, reject) => {
    const userPool = new CognitoUserPool(PoolData);

    const attributes = Object.values(fields).filter(f => f.attribute_name !== null).map(f => new CognitoUserAttribute({
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

function Login(user, password) {
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
      mfaSetup(challengeName, challengeParameters) {
        console.log('mfaSetup');
      },
      // @ts-ignore
      associateSecretCode(secretCode) {
        console.log('associateSecretCode');
      },
      selectMFAType(challengeName, challengeParameters) {
        console.log('selectMFAType');
      },
      totpRequired(secretCode) {
        console.log('totpRequired');
      },
      mfaRequired(codeDeliveryDetails) {
        console.log('mfaRequired');
      },
    });
  });
}

function GetCurrentUser() {
  return new Promise((resolve, reject) => {
    const userPool = new CognitoUserPool(PoolData);
    const cognitoUser = userPool.getCurrentUser();

    if (cognitoUser !== null) {
      cognitoUser.getSession((sessionErr, session) => {
        if (sessionErr) {
          reject(sessionErr);
          return;
        }

        console.log(`session validity: ${session.isValid()}`);

        cognitoUser.getUserData((err, data) => {
          if (err) {
            reject(err);
          } else {
            resolve(new User(data.Username, data.UserAttributes));
          }
        });
      });
    } else {
      reject(new Error('No current user'));
    }
  });
}

function Logout() {
  const userPool = new CognitoUserPool(PoolData);
  const cognitoUser = userPool.getCurrentUser();

  if (cognitoUser != null) {
    cognitoUser.signOut();
  }
}

function IsLoggedIn() {
  return new Promise((resolve, reject) => {
    const userPool = new CognitoUserPool(PoolData);
    const cognitoUser = userPool.getCurrentUser();
    if (cognitoUser === null) {
      resolve(false);
      return;
    }

    cognitoUser.getSession((err, session) => {
      if (err) {
        reject(err);
        return;
      }

      resolve(session.isValid());
    });
  });
}

export default {
  Register,
  Login,
  Logout,
  GetCurrentUser,
  IsLoggedIn,
};
