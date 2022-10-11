import axios from "axios"

export default {
    install(Vue, options) {
        Vue.prototype.$http = axios.create({
            baseUrl: options.baseUrl || "http://localhost:8081",
            headers: options.headers || null
        })
    }
}
