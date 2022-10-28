<template>
<div>
  {{"Имя:"}}<textarea v-model="name"/>
  {{"Возраст:"}}<textarea v-model="age"/>
  <ButtonComponent v-if="!changePwd"
      :label="'сменить пароль'"
      @btnClick="changePwd=true"
  />
  <div v-if="changePwd">
    <textarea v-model="newPwd"/>
    <textarea v-model="confirmedPwd"/>
  </div>
  <div>
    <ButtonComponent
                     :label="'сохранить'"
                     @btnClick="save"
    />
  </div>
</div>
</template>

<script>
import ButtonComponent from "@/components/ButtonComponent";
import Vue from "vue";
export default {
  name: "SettingsComponent",
  components: {ButtonComponent},
  data() {
    return {
      user: null,
      name: "",
      age: 0,
      changePwd: false,
      newPwd: "",
      confirmedPwd: ""
    }
  },
  created() {
    this.user = this.$store.getters.getUser
  },
  methods: {
    save() {
      if (!this.name) {
        this.name = this.user.Data.Name
      }
      if (!this.age) {
        this.age = this.user.Data.Age
      }
      if (!this.changePwd) {
        const data ={
          ID: this.user.ID,
          login: this.user.Login,
          password: this.user.Password,
          data: {
            reviewCount: this.user.Data.ReviewCount,
            name: this.name,
            age: this.age,
            likes: this.user.Data.Likes
          }
        }
        this.$http.patch(Vue.prototype.$baseUrl+"/api/v1/users", data).then(response=>{
          this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {ObjectID:Number(response.data)}}).then(r=>{
            this.user = r && r.data ? r.data[0] : null
            this.$store.commit('setUser', this.user)
            this.$router.push('/profile')
          })
        })
      } else {
        if (this.newPwd!==this.confirmedPwd || this.newPwd.length<8) {
          alert("пароли не совпадают или новый пароль слишком короткий")
          return
        } else {
          let data = {
            UserID: this.user.ID,
            password: this.newPwd
          }
          this.$http.post(Vue.prototype.$baseUrl+"/api/v1/updatePassword", data).then(()=>{
            const userData ={
              ID: this.user.ID,
              login: this.user.Login,
              password: this.user.Password,
              Data: {
                reviewCount: Number(this.user.Data.ReviewCount),
                name: this.name,
                age: Number(this.age),
                likes: this.user.Data.Likes
              }
            }
              this.$http.patch(Vue.prototype.$baseUrl+"/api/v1/users", userData).then(response=>{
                this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {ObjectID:Number(response.data)}}).then(r=>{
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

</style>