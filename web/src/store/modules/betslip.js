// import Vue from 'vue';

const getters = {

};


const actions = {
  toggle({ commit }) {
    commit('toggle');
  },
};

const mutations = {
  toggle(state) {
    state.show = !state.show;
  },
};

const state = {
  show: false,
  bets: [],
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
