const getters = {

};

const actions = {
  displayLogin({ commit }, display) {
    commit('displayLogin', display);
  },
  displayRegister({ commit }, display) {
    commit('displayRegister', display);
  },
};

const mutations = {
  displayLogin(state, display) {
    state.display_login = display;
  },
  displayRegister(state, display) {
    state.display_register = display;
  },
};


const state = {
  display_login: false,
  display_register: false,
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
