<template>
  <div id="app">
    <!--- Header --->
    <div class="">
      <HeaderComponent ref="header"/>
      <div v-if="isLoggedIn">
        <sub-header-component ref="subheader"/>
      </div>
    </div>

    <!--- Content --->
    <div>
      <div v-if="!isLoggedIn && !loginClicked && !registrationClicked">
        <buttonComponent
            :selected=false
            :label="'Войти'"
            @btnClick="()=>{loginClicked=!loginClicked}"
        />
        <buttonComponent
            :selected=false
            :label="'Регистрация'"
            @btnClick="()=>{registrationClicked=!registrationClicked}"
        />
      </div>
      <LoginComponent v-if="loginClicked" @user_id="setUser" @login="(val)=>{isLoggedIn=val}" class="logReg"/>
      <CompRegistration v-if="registrationClicked" @user_id="setUser" @registered="(val)=>{isRegistered=val}"/>
      <MainView v-if="isRegistered || isLoggedIn" :name="name"/>
    </div>
  </div>
</template>

<script>
import LoginComponent from './components/LoginComponent.vue'
import CompRegistration from "@/components/CompRegistration";
import MainView from "@/components/MainView";
import Vue from "vue";
import ButtonComponent from "@/components/ButtonComponent";
import HeaderComponent from "@/components/HeaderComponent";
import SubHeaderComponent from "@/components/SubHeaderComponent";

export default {
  name: 'App',
  components: {
    SubHeaderComponent,
    HeaderComponent,
    ButtonComponent,
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
      name: null,
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
  background-color: #363537;
  height: 100vh;
  max-height: 100vh;
}
</style>
