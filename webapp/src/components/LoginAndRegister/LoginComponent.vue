<template>
  <div class="screen">
    <div class="regBox">
      Вход
      <div class="divStdInput">
        <input type="text" placeholder="Логин" title="login" v-model="login" class="stdInputAll">
      </div>
      <div class="divStdInput">
        <input id="password"
               placeholder="Пароль"
               title="password"
               v-model="password"
               class="stdInputPassword"
               :class="{animated:wrongPasswd}"
               :type="pwdShown ? 'text' : 'password'">
        <button class="eye"
                @click="()=>{pwdShown=!pwdShown}"
                :class="{animated:wrongPasswd}"
        >
          <img :src="require('../../assets/'+`${pwdShown ? 'eye-closed.svg':'eye-opened.svg'}`)" alt=""/>
        </button>
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
  beforeMount() {
    document.addEventListener("keydown", this.onKeyDown)
  },
  beforeDestroy() {
    document.removeEventListener("keydown", this.onKeyDown)
  },
  methods: {
    onKeyDown(e) {
      if (!["Enter"].includes(e.code)) {
        return
      }
      e.preventDefault()
      this.onLogin()
    },
    onLogin() {
      this.wrongPasswd = false;
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
              this.wrongPasswd = true;
              let el = document.getElementById("password")
              el.className=".stdInputPassword.animated"
              console.log(el)
            } else {
              this.wrongPasswd = true;
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

.stdInputAll {
  border-color: #363537;
  border-width: 2px;
  border-top: none; border-left: none; border-right: none;
  min-width: 150px;
  width: 200px;
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
.divStdInput {
  display: flex;
  flex-direction: row;
}
.stdInputPassword {
  border-color: #363537;
  border-width: 2px;
  border-top: none; border-left: none; border-right: none;
  min-width: 114px;
  width: calc(200px - 36px);
  height: 28px;
}
.eye{
  height: 28px;
  background: none;
  border: none;
  border-bottom: #363537 2px solid;
}

.stdInputPassword.animated {
  animation-name: shake;
  animation-duration: 1s;
  animation-fill-mode: both;
}
.eye.animated {
  animation-name: shake;
  animation-duration: 1s;
  animation-fill-mode: both;
}
@keyframes shake {
  0%, 100% {transform: translateX(0); border-color: red}
  10%, 30%, 50%, 70%, 90% {transform: translateX(-10px);}
  20%, 40%, 60%, 80% {transform: translateX(10px);}
}

</style>