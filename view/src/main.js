import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import GAuth from 'vue-google-oauth2'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

const googleAuthOption = {
  clientId: '98250998013-tbbgns4o192oci849p3e8pf7ntfdgl8g.apps.googleusercontent.com',
  scope: 'profile email',
  prompt: 'consent'
}

Vue.use(GAuth, googleAuthOption)
Vue.use(BootstrapVue)
Vue.prototype.$axios = axios; // add
Vue.config.productionTip = true;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
