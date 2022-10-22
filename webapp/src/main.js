import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'

Vue.config.productionTip = false
Vue.prototype.$baseUrl="http://localhost:8081"

import http from "./plugins/http"
import MainView from "@/components/MainView";
import LoginComponent from "@/components/LoginComponent";
import OpenComponent from "@/components/OpenComponent";
import CompRegistration from "@/components/CompRegistration";
Vue.use(http, {
  baseUrl: Vue.prototype.$baseUrl
})

const routes = [
  {
    path: "/registration",
    component: CompRegistration
  },
  {
    path: "/login",
    component: LoginComponent
  },
  {
    path: "/content",
    component: MainView
  },
  {
    path: "/",
    component: OpenComponent
  }
]
const router = new VueRouter({
  history: 'history',
  routes, // short for `routes: routes`
})
Vue.use(VueRouter)
new Vue({
  router: router,
  render: h => h(App),
}).$mount('#app')
