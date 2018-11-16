import Util from '../../api/util';

const getters = {

};

/* eslint-disable no-restricted-globals  */

const actions = {
  push({ commit, state }, action) {
    if (state.session.id === '') {
      commit('newSession', {
        id: Util.UUID(),
        user_agent: navigator.userAgent,
        device_width: screen.width,
        device_height: screen.height,
        timezone: new Date().getTimezoneOffset(),
        ip_address: '',
      });
    }

    commit('addAction', action);
  },
};

const mutations = {
  addAction(state, action) {
    state.actions.push(action);
  },
  newSession(state, session) {
    state.session = session;
  },
  clearActions(state) {
    state.actions = [];
  },
};

const state = {
  actions: [],
  session: {
    id: '',
    user_agent: '',
    device_width: 0,
    device_height: 0,
    timezone: 0,
    ip_address: '',
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
