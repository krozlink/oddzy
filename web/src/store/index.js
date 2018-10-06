import Vue from 'vue';
import Vuex from 'vuex';
import racing from './modules/racing';
import account from './modules/account';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    racing,
    account,
  },
  strict: debug,
  state: {
  },
  mutations: {

  },
  actions: {

  },
});
