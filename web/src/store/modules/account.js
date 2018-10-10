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
      const user = fields.user_name.getValue();
      const password = fields.password.getValue();

      const register = await Auth.Register(user, password, fields);
      console.log(register);
      const auth = await Auth.Authenticate(user, password);
      console.log(auth);
      commit('registrationSuccessful');
    } catch (ex) {
      console.log(ex);
      commit('registrationFailed');
    }
  },

  async authenticate({ commit }, user, password) {
    try {
      commit('loginSubmitted');
      const auth = await Auth.Authenticate(user, password);
      console.log(auth);
      commit('loginSuccessful');
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
  registrationSuccessful(state) {
    state.display_register = false;
    state.registation_status = '';
    state.registration_message = '';
    state.authenticated = true;
  },
  registrationFailed(state) {
    state.authenticated = false;
    state.registation_status = 'failed';
    state.registration_message = 'An unexpected error occurred';
  },
  registrationSubmitted(state) {
    state.registation_status = 'submitted';
    state.registration_message = 'Registering...';
  },
  displayRegistrationScreen(state, display) {
    state.display_register = display;
  },

  loginSuccessful(state) {
    state.display_login = false;
    state.login_status = '';
    state.login_message = '';
    state.authenticated = true;
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
    given_name: '',
    family_name: '',
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
