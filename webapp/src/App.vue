<template>
  <div id="app">
    <div v-if="!isLoggedIn && !loginClicked && !registrationClicked">
      <button class="button" @click="()=>{loginClicked=!loginClicked}">войти</button>
      <button class="button" @click="()=>{registrationClicked=!registrationClicked}">регистрация</button>
    </div>
    <LoginComponent v-if="loginClicked" :isLoggedIn="isLoggedIn" @user_id="setUser" @login="(val)=>{isLoggedIn=val}"/>
    <CompRegistration v-if="registrationClicked" @user_id="setUser" @registered="(val)=>{isRegistered=val}"/>
    <MainView v-if="isRegistered || isLoggedIn" :name="name"/>
  </div>
</template>

<script>
import LoginComponent from './components/Login.vue'
import CompRegistration from "@/components/CompRegistration";
import MainView from "@/components/MainView";
import Vue from "vue";

export default {
  name: 'App',
  components: {
    MainView,
    CompRegistration,
    LoginComponent
  },
  data() {
    return {
      isLoggedIn: false,
      isRegistered: false,
      loginClicked: false,
      registrationClicked: false,
      name: null
    }
  },
  methods: {
    setUser(val){
      console.log(val)
      this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", { params: { ObjectID: Number(val) } })
          .then(response=>{
            this.name = response.data.find(u=>u.ID===val).Data.Name
          })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
