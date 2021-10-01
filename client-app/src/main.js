import Vue from 'vue';
import App from './App.vue';
import { router } from './router';
import vuetify from './plugins/vuetify';
import { store } from './stores/store';
import '@mdi/font/css/materialdesignicons.css';

// Vue.config.productionTip = false;

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');