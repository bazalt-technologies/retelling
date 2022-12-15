<template>
  <div class="setShell">
    <settings-btn :label="'Изменить пароль'" @SetBtnClicked="()=>{this.passChange = !this.passChange}"/>
    <div class="textAreaShell" v-if="!this.passChange">
      Новый пароль:
      <textarea class="setTextArea" v-model="password"/>
      <confirm-btn :label="'Сохранить'" @click="saveChanges"/>
    </div>
    <settings-btn :label="'Изменить имя'" @SetBtnClicked="()=>{this.nameChange = !this.nameChange}"/>
    <div class="textAreaShell" v-if="!this.nameChange">
      Новое имя:
      <textarea class="setTextArea" v-model="name"/>
      <confirm-btn :label="'Сохранить'" @click="saveChanges"/>
    </div>
    <settings-btn :label="'Изменить возраст'" @SetBtnClicked="()=>{this.ageChange = !this.ageChange}"/>
    <div class="textAreaShell" v-if="!this.ageChange">
      Новый возраст:
      <textarea class="setTextArea" v-model="age"/>
      <confirm-btn :label="'Сохранить'" @click="saveChanges"/>
    </div>
    <settings-btn :label="'Изменить профессию'" @SetBtnClicked="()=>{this.profChange = !this.profChange}"/>
    <div class="textAreaShell" v-if="!this.profChange">
      Новая профессия:
      <textarea class="setTextArea" v-model="profession"/>
      <confirm-btn :label="'Сохранить'" @click="saveChanges"/>
    </div>
    <settings-btn :label="'Выйти'" :red="true" @SetBtnClicked="quit"/>

  </div>
</template>

<script>
import SettingsBtn from "@/components/Settings/SettingsBtn";
import ConfirmBtn from "@/components/ConfirmBtn";
import Vue from "vue";
export default {
  name: "ProfileComponent",
  components: {ConfirmBtn, SettingsBtn},
  data () {
    return {
      user: null,
      data: null,
      login: null,
      password: null,
      profession: null,
      name: null,
      oldPwd: null,
      age:null,
      loginChange: Boolean,
      passChange: Boolean,
      nameChange: Boolean,
      ageChange: Boolean,
      profChange: Boolean,
    }
  },
  beforeMount() {
    this.user = this.$store.getters.getUser
    if (!this.user) {
      this.$router.push('/login')
    }
  },
  created() {
    this.user = this.$store.getters.getUser
    console.log(this.user)
    if (!this.user) {
      this.$router.push('/login')
    }
  },
  updated() {
    if (!this.user) {
      this.$router.push('/login')
    }
  },
  methods: {
    quit() {
      this.$store.commit('setUser', null)
      this.$router.push('/login')
    },
    saveChanges() {
      console.log(this.login)
      if (!this.name) {
        this.name = this.user.Data.Name
      }
      if (!this.login) {
        this.login = this.user.Login
      }
      if (!this.profession) {
        this.profession = this.user.Data.Profession
      }
      if (!this.age) {
        this.age = this.user.Data.Age
      }
      if (!this.password) {
        const data ={
                ID: this.user.ID,
                login: this.login,
                password: this.user.Password,
                data: {
                  reviewCount: this.user.Data.ReviewCount,
                  name: this.name,
                  age: Number(this.age),
                  likes: this.user.Data.Likes,
                  profession: this.profession
                }
              }
              console.log(data)
              this.$http.patch(Vue.prototype.$baseUrl+"/api/v1/users", data).then(response=>{
                this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {ObjectID:Number(response.data)}}).then(r=>{
                  this.user = r && r.data ? r.data[0] : null
                  this.$store.commit('setUser', this.user)
                  this.$router.push('/profile')
                })
              })
            } else {
        if (this.password.length < 8) {
          alert("новый пароль слишком короткий")
          return
        } else {
          let data = {
            UserID: this.user.ID,
            password: this.newPwd
          }
          this.$http.post(Vue.prototype.$baseUrl + "/api/v1/updatePassword", data).then(() => {
            const userData = {
              ID: this.user.ID,
              login: this.user.Login,
              password: this.user.Password,
              Data: {
                reviewCount: Number(this.user.Data.ReviewCount),
                name: this.name,
                age: Number(this.age),
                likes: this.user.Data.Likes,
                profession: this.profession
              }
            }
            this.$http.patch(Vue.prototype.$baseUrl + "/api/v1/users", userData).then(response => {
              this.$http.get(Vue.prototype.$baseUrl + "/api/v1/users", {params: {ObjectID: Number(response.data)}}).then(r => {
                this.user = r && r.data ? r.data[0] : null
                this.$store.commit('setUser', this.user)
                this.$router.push('/profile')
              })
            })
            this.$router.push('/profile')
          })
        }
      }
    }
  }
}
</script>

<style scoped>
.setShell {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.setTextArea {
  width: 300px;
  height: 50px;
  padding-top:10px;
  border: #94d1be 3px solid;
  border-radius: 20px;
  resize: none;
  margin-top: 5px;
  margin-bottom: 10px;
}
.textAreaShell {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: #fefefe;
}
</style>