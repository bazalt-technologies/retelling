import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import storeData from "@/store";
import Vuex from 'vuex'
import {BootstrapVue, IconsPlugin} from "bootstrap-vue";
import 'bootstrap/dist/css/bootstrap.css'

import 'bootstrap-vue/dist/bootstrap-vue.css'
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
Vue.use(Vuex)
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
const store = new Vuex.Store(storeData)
new Vue({
  router: router,
  store: store,
  render: h => h(App),
}).$mount('#app')
