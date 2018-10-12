import Vue from 'vue';
import Auth from '../../api/auth';

const getters = {

};

const actions = {
  displayLogin({ commit }, display) {
    commit('displayLoginScreen', display);
  },
  displayRegister({ commit }, display) {
    commit('displayRegistrationScreen', display);
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

  async login({ commit }, username, password) {
    try {
      commit('loginSubmitted');
      const auth = await Auth.Login(username, password);
      console.log(auth);

      const user = await Auth.GetCurrentUser();
      commit('loginSuccessful', user);
    } catch (ex) {
      console.log(ex);
      let message = '';
      if (ex.message === 'Invalid password') {
        message = 'Incorrect username or password';
      } else {
        message = 'Unexpected error occurred';
      }
      commit('loginFailed', message);
    }
  },
};

const mutations = {
  userLoggedIn(state, user) {
    state.authenticated = true;
    state.login_status = '';
    state.login_message = '';
    Vue.set(state.user_details, 'username', user.username);
    Vue.set(state.user_details, 'first_name', user.first_name);
    Vue.set(state.user_details, 'last_name', user.last_name);
    Vue.set(state.user_details, 'email_address', user.email_address);
  },
  registrationSuccessful(state, user) {
    state.display_register = false;
    state.registation_status = '';
    state.registration_message = '';
    this.userLoggedIn(user);
  },
  registrationFailed(state, message) {
    state.authenticated = false;
    state.registation_status = 'failed';
    state.registration_message = message;
  },
  registrationSubmitted(state) {
    state.registation_status = 'submitted';
    state.registration_message = 'Registering...';
  },
  displayRegistrationScreen(state, display) {
    state.display_register = display;
  },

  loginSuccessful(state, user) {
    state.display_login = false;
    this.userLoggedIn(user);
  },
  loginFailed(state, message) {
    state.authenticated = false;
    state.login_status = 'failed';
    state.login_message = message;
  },
  loginSubmitted(state) {
    state.login_status = 'submitted';
    state.login_message = 'Logging in...';
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
    email_address: '',
  },

  registation_status: '',
  registration_message: '',

  login_status: '',
  login_message: '',
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
