import Vue from 'vue';
import Util from '../../api/util';
import Betslip from '../../api/betslip';

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

  confirm({ commit }) {
    commit('updateStatus', Betslip.STATUS.SUBMITTING);
    setTimeout(() => {
      commit('betSuccessful');

      const ref = Util.Random(100000, 999999);
      commit('setMessage', {
        lines: ['Bet successfully placed', `Reference No. ${ref}`],
        type: 'success',
      });
    }, 1000);
  },

  cancel({ commit }) {
    commit('updateStatus', Betslip.STATUS.UNPLACED);
  },


  place({ commit, state, rootState }) {
    let isValid = true;

    const error = Betslip.Validate(rootState);
    if (error !== '') {
      isValid = false;
      commit('setMessage', {
        lines: [error],
        type: 'error',
      });
    }

    Object.values(state.bets).forEach((b) => {
      const msg = b.validate();
      commit('updateBetMessage', {
        betId: b.bet_id,
        message: msg,
      });

      if (msg) isValid = false;
    });

    if (!isValid) return;

    commit('updateStatus', Betslip.STATUS.UNCONFIRMED);
  },
};

const mutations = {
  toggle(state) {
    state.show = !state.show;
  },
  addToBetslip(state, bet) {
    state.show = true;
    if (state.status === Betslip.STATUS.SUBMITTING) return;

    // prevent duplicates
    if (state.bets[bet.bet_id]) return;
    Vue.set(state.bets, bet.bet_id, bet);
    Vue.set(state.message, 'type', '');
    Vue.set(state.message, 'lines', []);
  },
  removeFromBetslip(state, betId) {
    if (state.status === Betslip.STATUS.SUBMITTING) return;
    if (state.bets[betId]) {
      Vue.delete(state.bets, betId);
    }
  },
  updateBetAmount(state, { betId, amount }) {
    if (state.status === Betslip.STATUS.SUBMITTING) return;
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
