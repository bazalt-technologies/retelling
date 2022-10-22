import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'

Vue.config.productionTip = false
Vue.prototype.$baseUrl="http://localhost:8081"

import http from "./plugins/http"
import routes from "@/routes";
Vue.use(http, {
  baseUrl: Vue.prototype.$baseUrl
})

const router = new VueRouter({
  history: 'history',
  routes, // short for `routes: routes`
})
Vue.use(VueRouter)
new Vue({
  router: router,
  render: h => h(App),
}).$mount('#app')
