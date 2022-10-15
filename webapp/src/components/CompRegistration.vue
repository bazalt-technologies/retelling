<template>
  <div>
    <div>
      <input type="text" title="login" v-model="name" :width="100" :height="50">
    </div>
    <div>
      <input type="text" title="login" v-model="login" :width="100" :height="50">
    </div>
    <div>
      <input type="text" title="password" v-model="password" :width="100" :height="50">
    </div>
    <div>
      <input type="number" title="password" v-model="age" :width="100" :height="50">
    </div>
    <div>
    <button @click="onRegister" class="button">ЗАРЕГИСТРИРОВАТЬСЯ</button>
    </div>
  </div>
</template>

<script>
import Vue from "vue"
export default {
  name: "CompRegistration",
  data() {
    return{
      name:null,
      login: null,
      password: null,
      age: null,
    }
  },
  methods: {
    onRegister() {
      if (!this.login || !this.password || !this.age) {
        return
      }
      if (this.password.length<8) {
        return
      }
      const json = {
        login:this.login,
        password:this.password,
        Data:{
            Name:this.name,
            Age: Number(this.age)
        }
      }
      this.$http.post(Vue.prototype.$baseUrl+"/api/v1/users", json)
          .then(response=>{
            console.log(typeof response.data)
            this.$emit("user_id", response.data)
            this.$emit("registered", true)
          })
    }
  }
}
</script>

<style scoped>

</style>