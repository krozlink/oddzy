import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import Racing from './views/Racing.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      
    },
    {
      path: '/racing',
      name: 'racing',
      component: Racing, 
    },
  ],
});
