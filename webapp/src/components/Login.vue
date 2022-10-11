<template>
  <div>
    <div>
      <input type="text" title="login" v-model="login" :width="100" :height="50">
    </div>
    <div>
      <input type="text" title="password" v-model="password" :width="100" :height="50">
    </div>
    <button @click="onLogin" class="button">ВОЙТИ</button>
  </div>
</template>

<script>
import Vue from "vue"
export default {
  props: {
    isLoggedIn:{
      type: Boolean,
      required: true
    }
  },
  name: "LoginComponent",
  data() {
    return{
      login: null,
      password: null
    }
  },
  methods: {
    onLogin() {
      console.log(Vue.prototype.$baseUrl)
      const data = {
        login: this.login,
        password: this.password
      }
      this.$http.post(Vue.prototype.$baseUrl+"/api/v1/authUser", data)
          .then(response=>{
            console.log(response.data)
            this.$emit("login", true)
            this.$emit("user_id", response.data)
          })
    }
  }
}
</script>

<style>
.button {
  background-color: #4CAF50;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
}
</style>