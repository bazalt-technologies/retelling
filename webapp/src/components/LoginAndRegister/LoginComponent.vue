<template>
  <div class="screen">
    <div class="regBox">
      Вход
      <div>
        <input type="text" placeholder="Логин" title="login" v-model="login" class="stdInput">
      </div>
      <div>
        <input placeholder="Пароль" title="password" v-model="password" class="stdInput" :type="pwdShown ? 'text' : 'password'">
        <button class="eye" @click="()=>{pwdShown=!pwdShown}"><img :src="require('../../assets/'+`${pwdShown ? 'eye-closed.svg':'eye-opened.svg'}`)" alt=""/></button>
      </div>
      <div v-if="wrongPasswd" class="errMsg">Неверный логин или пароль</div>
      <button-component @btnClick="() => {onLogin();}"
                        :label="'Войти'"
                        :selected="false"
                        class="btn"
      />
      <div class="subText">
        <div @click="$router.push('/registration')">Нет аккаунта? Зарегистрироваться</div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue"
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "LoginComponent",
  components: {ButtonComponent},
  data() {
    return{
      login: null,
      password: null,
      wrongPasswd: false,
      pwdShown: false
    }
  },
  mounted() {
    let usr = this.$store.getters.getUser
    if( usr){
      this.$router.push(`this.user`)
    }
  },
  methods: {
    onLogin() {
      const data = {
        login: this.login,
        password: this.password
      }
      this.$http.post(Vue.prototype.$baseUrl+"/api/v1/authUser", data)
          .then(response=>{
            if (response.status !== 401) {
              this.wrongPasswd = false;
              this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {"ObjectID": Number(response.data)}}).then(resp=>{
                let usr = resp.data ? resp.data.find(u=>u.ID===response.data) : {}
                this.$store.commit('setUser', usr)
                this.$router.push("/content/recommendations")
              })
            } else {
              this.wrongPasswd = true;
            }
          }).catch(err=> {
            if (err.response.statusText && err.response.statusText === 'Unauthorized') {
              alert("Неверный логин или пароль")
            } else {
              alert("Ошибка, повторите позже")
            }
          })
    }
  }
}
</script>

<style>
.screen {
  display: flex;
  flex-direction: column;
  align-items: center;
  align-content: center;
  justify-content: center;
  background: url("../../assets/officebg.png") center/cover;
  height: calc(100vh - 4vw - 10px);
  max-height: calc(100vh - 50px);
}
.regBox {
  display: flex;
  flex-direction: column;
  align-items: center;
  align-content: space-between;
  justify-content: space-between;
  padding: 20px;
  width: 20vw;
  height: 25vh;
  min-width: 300px;
  min-height: 150px;
  background: #fefefe;
  border-radius: 25px;
}
.stdInput {
  border-color: #363537;
  border-width: 2px;
  border-top: none; border-left: none; border-right: none;
  min-width: 250px;
  width: 17vw;
}
.btn {
}
.subText {
  margin-top: 15px;
  font-size: 12px;
  cursor: pointer;
}
.errMsg {
  margin-top: 5px;
  font-size: 10px;
  color: red;
}

.eye{
  height: 20px;
  background: none;
  border: none;
}
</style>