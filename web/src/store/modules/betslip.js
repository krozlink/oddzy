import Vue from 'vue';
import Util from '../../api/util';

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
  updateBetAmount({ commit }, data) {
    commit('updateBetAmount', data);
  },
  submit({ commit, state }) {
    let isValid = true;
    Object.values(state.bets).forEach((b) => {
      const msg = b.validate();
      if (msg !== '') {
        isValid = false;
        commit('updateBetMessage', {
          betId: b.bet_id,
          message: msg,
        });
      }
    });

    if (!isValid) return;

    commit('updateStatus', 'submitting');
    setTimeout(() => {
      commit('betSuccessful');

      const ref = Util.Random(100000, 999999);
      commit('setMessage', {
        lines: ['Bet successfully placed', `Reference No. ${ref}`],
        type: 'success',
      });
    }, 1000);
  },
};

const mutations = {
  toggle(state) {
    state.show = !state.show;
  },
  addToBetslip(state, bet) {
    state.show = true;
    if (state.status === 'submitting') return;

    // prevent duplicates
    if (state.bets[bet.bet_id]) return;
    Vue.set(state.bets, bet.bet_id, bet);
    Vue.set(state.message, 'type', '');
    Vue.set(state.message, 'lines', []);
  },
  removeFromBetslip(state, betId) {
    if (state.status === 'submitting') return;
    if (state.bets[betId]) {
      Vue.delete(state.bets, betId);
    }
  },
  updateBetAmount(state, { betId, amount }) {
    if (state.status === 'submitting') return;
    Vue.set(state.bets[betId], 'amount', amount);
  },
  updateStatus(state, status) {
    state.status = status;
  },
  updateBetMessage(state, { betId, message }) {
    Vue.set(state.bets[betId], 'message', message);
  },
  betSuccessful(state) {
    state.status = '';
    state.bets = {};
  },
  setMessage(state, { lines, type }) {
    Vue.set(state.message, 'lines', lines);
    Vue.set(state.message, 'type', type);
  },
};

const state = {
  show: false,
  status: '',
  bets: {

  },

  message: {
    lines: [],
    type: '',
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
