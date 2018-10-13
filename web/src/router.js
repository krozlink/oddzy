import Vue from 'vue';
import Router from 'vue-router';
import Account from './views/Account.vue';
import Home from './views/Home.vue';
import Racing from './views/Racing.vue';
import RaceCard from './views/RaceCard.vue';

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
      path: '/account',
      name: 'account',
      component: Account,
    },
    {
      path: '/racing',
      name: 'racing-today',
      component: Racing,
    },
    {
      path: '/racing/:date',
      name: 'racing-date',
      component: Racing,
    },
    {
      path: '/racing/:location/:id',
      name: 'race-card',
      component: RaceCard,
    },
  ],
});
