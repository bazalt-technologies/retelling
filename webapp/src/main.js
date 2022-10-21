import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'

Vue.config.productionTip = false
Vue.prototype.$baseUrl="http://localhost:8081"

import http from "./plugins/http"
import ContentComponent from "@/components/ContentComponent";
import MainView from "@/components/MainView";
Vue.use(http, {
  baseUrl: Vue.prototype.$baseUrl
})

const routes = [
  {
    path: "/user/:userId/content",
    component: MainView
  }
]
const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes, // short for `routes: routes`
})
const app = new Vue({
  render: h => h(App),
})
app.$mount('#app')
app.use(router)
