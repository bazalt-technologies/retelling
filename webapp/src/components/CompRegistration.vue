<template>
  <div class="screen">
    <div class="regBox">
      Регистрация
      <div>
        <input type="text"
               placeholder="Имя"
               title="name"
               v-model="name"
               class="stdInput"
        />
      </div>
      <div>
        <input type="text"
               placeholder="Логин"
               title="login"
               v-model="login"
               class="stdInput"
        />
      </div>
      <div>
        <input type="text"
               placeholder="Пароль"
               title="password"
               v-model="password"
               class="stdInput"
        />
      </div>
      <div>
        <input type="number"
               placeholder="Возраст"
               title="password"
               v-model="age"
               class="stdInput"
        />
      </div>
      <div>
      <button-component @btnClick="onRegister"
                        :label="'Зарегестрироваться'"
                        :selected="false"
                        class="btn"
      />
      </div>
      <div class="subText">
        <a>Есть аккаунт? Войти</a>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue"
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "CompRegistration",
  components: {ButtonComponent},
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
.screen {
  display: flex;
  flex-direction: column;
  align-items: center;
  align-content: center;
  justify-content: center;
  height: 80vh;
}
.regBox {
  display: flex;
  flex-direction: column;
  align-items: center;
  align-content: center;
  justify-content: center;
  width: 20vw;
  height: 40vh;
  min-width: 300px;
  min-height: 300px;
  background: #fefefe;
  border-radius: 25px;
}
.stdInput {
  border-color: #363537;
  border-width: 2px;
  border-top: none; border-left: none; border-right: none;
  min-width: 250px;
  width: 17vw;
  margin-top: 10px;
}
.btn {
  margin-top: 20px;
}
.subText {
  margin-top: 15px;
  font-size: 12px;
  cursor: pointer;
}
</style>