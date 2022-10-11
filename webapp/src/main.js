import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false
Vue.prototype.$baseUrl="http://localhost:8081"

import http from "./plugins/http"
Vue.use(http, {
  baseUrl: Vue.prototype.$baseUrl
})
new Vue({
  render: h => h(App),
}).$mount('#app')
