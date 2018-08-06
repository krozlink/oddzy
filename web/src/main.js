import Vue from 'vue';
// import Buefy from 'buefy';
import 'buefy/lib/buefy.css';
import './styles/default.scss'

import App from './App.vue';
import router from './router';
import store from './store';

Vue.config.productionTip = false;
// Vue.use(Buefy);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
