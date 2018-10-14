import Vue from 'vue';
import Vuex from 'vuex';
import racing from './modules/racing';
import account from './modules/account';
import betslip from './modules/betslip';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    racing,
    account,
    betslip,
  },
  strict: debug,
  state: {
  },
  mutations: {

  },
  actions: {

  },
});
