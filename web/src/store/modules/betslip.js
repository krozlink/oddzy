// import Vue from 'vue';

const getters = {

};


const actions = {
  toggle({ commit }) {
    commit('toggle');
  },
  addToBetslip({ commit }, bet) {
    commit('addToBetslip', bet);
  },
  removeFromBetslip({ commit }, betId) {
    commit('removeFromBetslip', betId);
  },
};

const mutations = {
  toggle(state) {
    state.show = !state.show;
  },
  addToBetslip(state, bet) {
    state.bets.append(bet);
  },
  removeFromBetslip(state, betId) {
    state.bets = state.bets.filter(b => b.bet_id !== betId);
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
