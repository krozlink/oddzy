import { Update } from '../../api/tracking/index';

const getters = {

};

const actions = {
  push({ commit }, action) {
    commit('addAction', action);
  },
};

const mutations = {
  addAction(state, action) {
    state.actions.push(action);
    Update();
  },
  clearActions(state) {
    state.actions = [];
  },
};

const state = {
  actions: [],
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
