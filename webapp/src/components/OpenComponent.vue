<template>
  <div id="#open">
  <!--- Content --->
  <div>
    <div v-if="!isLoggedIn && !loginClicked && !registrationClicked">
      <buttonComponent
          :selected="false"
          :label="'Войти'"
          @btnClick="()=>{$router.push('/login'); loginClicked = true}"
      />
      <buttonComponent
          :selected="false"
          :label="'Регистрация'"
          @btnClick="()=>{$router.push('/registration')}"
      />
    </div>
  </div>
  </div>
</template>

<script>
import Vue from "vue";
import ButtonComponent from "@/components/ButtonComponent";



export default {
  name: "OpenComponent",
  components: {
    ButtonComponent
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
      this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", { params: { ObjectID: Number(val) } })
          .then(response=>{
            this.name = response.data.find(u=>u.ID===val).Data.Name
            this.$router.push(`/content`)
          })
    }
  }
}
</script>

<style>
#open {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #363537;
  height: 100vh;
  max-height: 100vh;
}
</style>
