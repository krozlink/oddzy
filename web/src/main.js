import Vue from 'vue';
import VueSocketio from 'vue-socket.io';

import 'buefy/lib/buefy.css';
import './styles/default.scss';

import App from './App.vue';
import router from './router';
import store from './store';

if (!process.env.VUE_APP_SERVERLESS_ONLY) {
  Vue.use(VueSocketio, `${process.env.VUE_APP_SOCKETIO}/prices`, store);
}

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');

window.addEventListener('resize', () => {
  console.log(`window dimensions: ${document.body.clientWidth}x${document.body.clientHeight}`);
  console.log(`screen dimensions: ${window.screen.width}x${window.screen.height}`);
});
