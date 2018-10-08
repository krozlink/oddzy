const getters = {

};

const actions = {
  displayLogin({ commit }, display) {
    commit('displayLogin', display);
  },
  displayRegister({ commit }, display) {
    commit('displayRegister', display);
  },

  register({ commit }) {
    commit('setRegisterStatus', 'registering');
  },
};

const mutations = {
  displayLogin(state, display) {
    state.display_login = display;
  },
  displayRegister(state, display) {
    state.display_register = display;
  },

  setRegisterStatus(state, status) {
    state.register_status = status;
  },
};


const state = {
  display_login: false,
  display_register: false,

  register_status: '',
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
