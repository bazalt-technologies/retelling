<template>
  <div class="regScreen">
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
      <div v-if="noData" class="errMsg">
        Указаны не все данные
      </div>
      <div v-if="weakPassword" class="errMsg">
        Пароль должен быть длины не менее 8 символов
      </div>
      <div>
      <button-component @btnClick="()=>{onRegister();}"
                        :label="'Зарегистрироваться'"
                        :selected="false"
                        class="btn"
      />
      </div>
      <div class="subText">
        <div @click="$router.push('/login')">Есть аккаунт? Войти</div>
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
      noData: false,
      weakPassword: false,
    }
  },
  methods: {
    onRegister() {
      if (!this.login || !this.password || !this.age) {
        this.noData = true;
        return
      } else {
        this.noData = false;
      }
      if (this.password.length<8) {
        this.weakPassword = true;
        return
      } else {
        this.weakPassword = false;
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
            this.$store.commit('setUser', json)
            this.$emit("user_id", response.data)
            this.$emit("registered", true)
            this.$router.push('/content')
          })
    }
  }
}
</script>

<style scoped>
.regScreen {
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
  align-content: center;
  justify-content: space-between;
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
  margin-top: 15px;
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
</style>