import Vue from 'vue';
import Router from 'vue-router';
import Posts from './components/Posts';
import Login from './components/Login';
import Register from './components/Register';
import { store } from './stores/store';

Vue.use(Router);

export const router = new Router({
  mode: 'history',
  base: '/',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Posts,
    },

    {
      path: '/login',
      name: 'login',
      component: Login,
    },

    {
      path: '/register',
      name: 'register',
      component: Register,
    },

    { path: '*', redirect: '/' },
  ],
});

router.beforeEach((to, from, next) => {
  const publicPages = ['/register', '/login'];
  const authRequired = !publicPages.includes(to.path);
  const isLoggedIn = store.getters.isLoggedIn;

  if (authRequired && !isLoggedIn) {
    return next('/login');
  }

  next();
});
