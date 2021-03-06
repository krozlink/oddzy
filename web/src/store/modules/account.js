import Vue from 'vue';
import Auth from '../../api/auth';

function userLoggedIn(state, user) {
  state.authenticated = true;
  state.status = 'login_true';
  state.status_message = '';

  Vue.set(state.user_details, 'username', user.username);
  Vue.set(state.user_details, 'first_name', user.first_name);
  Vue.set(state.user_details, 'last_name', user.last_name);
  Vue.set(state.user_details, 'address', user.address);
  Vue.set(state.user_details, 'date_of_birth', user.date_of_birth);
  Vue.set(state.user_details, 'email_address', user.email_address);
  Vue.set(state.user_details, 'email_verified', user.email_verified);
  Vue.set(state.user_details, 'phone_number', user.mobile);
  Vue.set(state.user_details, 'phone_verified', user.mobile_verified);
}
function userLoggedOut(state) {
  state.authenticated = false;
  state.status = 'login_false';
  state.status_message = '';

  Vue.set(state.user_details, 'username', '');
  Vue.set(state.user_details, 'first_name', '');
  Vue.set(state.user_details, 'last_name', '');
  Vue.set(state.user_details, 'address', '');
  Vue.set(state.user_details, 'date_of_birth', '');
  Vue.set(state.user_details, 'email_address', '');
  Vue.set(state.user_details, 'email_verified', false);
  Vue.set(state.user_details, 'phone_number', '');
  Vue.set(state.user_details, 'phone_verified', false);
}

const getters = {

};


const actions = {
  displayLogin({ commit }, display) {
    commit('displayLoginScreen', display);
  },
  displayRegister({ commit }, display) {
    commit('displayRegistrationScreen', display);
  },

  async logout({ commit }) {
    Auth.Logout();
    commit('logout');
  },

  async register({ commit }, fields) {
    commit('registrationSubmitted');
    try {
      const username = fields.user_name.getValue();
      const password = fields.password.getValue();

      const register = await Auth.Register(username, password, fields);
      console.log(register);

      const auth = await Auth.Login(username, password);
      console.log(auth);

      const user = await Auth.GetCurrentUser();
      console.log(user);

      commit('registrationSuccessful', user);
    } catch (ex) {
      console.log(ex);
      commit('registrationFailed', 'An unexpected error occurred');
    }
  },

  async userLogin({ commit }, { username, password }) {
    try {
      commit('userLoginSubmitted');
      const auth = await Auth.Login(username, password);
      console.log(auth);

      const user = await Auth.GetCurrentUser();
      commit('userLoginSuccessful', user);
    } catch (ex) {
      console.log(ex);
      let message = '';
      if (ex.code === 'UserNotFoundException' || ex.code === 'NotAuthorizedException') {
        message = 'Incorrect username or password';
      } else {
        message = 'Unexpected error occurred';
      }
      commit('userLoginFailed', message);
    }
  },

  async autoLogin({ commit }) {
    try {
      const user = await Auth.GetCurrentUser();
      commit('autoLoginSuccessful', user);
    } catch (ex) {
      commit('autoLoginFailed');
    }
  },
};

const mutations = {
  registrationSuccessful(state, user) {
    state.display_register = false;
    userLoggedIn(state, user);
  },
  registrationFailed(state, message) {
    state.authenticated = false;
    state.status = 'registration_failed';
    state.registration_message = message;
  },
  registrationSubmitted(state) {
    state.status = 'registration_submitted';
    state.status_message = 'Registering...';
  },
  displayRegistrationScreen(state, display) {
    state.display_register = display;
  },
  logout(state) {
    userLoggedOut(state);
  },
  autoLoginFailed(state) {
    state.authenticated = false;
    state.status = 'login_false';
    state.status_message = '';
  },
  autoLoginSuccessful(state, user) {
    userLoggedIn(state, user);
  },

  userLoginSuccessful(state, user) {
    state.display_login = false;
    userLoggedIn(state, user);
  },
  userLoginFailed(state, message) {
    state.authenticated = false;
    state.status = 'login_failed';
    state.status_message = message;
  },
  userLoginSubmitted(state) {
    state.status = 'login_submitted';
    state.status_message = '';
  },
  displayLoginScreen(state, display) {
    state.display_login = display;
  },
};

const state = {
  display_login: false,
  display_register: false,

  authenticated: false,

  user_details: {
    username: '',
    first_name: '',
    last_name: '',

    date_of_birth: '',
    address: '',

    email_address: '',
    email_verified: false,

    phone_number: '',
    phone_verified: false,
  },

  status: 'login_checking',
  status_message: '',
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
