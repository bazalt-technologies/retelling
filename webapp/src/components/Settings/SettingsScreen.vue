<template>
  <div class="setShell">
    <settings-btn :label="'Изменить логин'" @SetBtnClicked="()=>{this.loginChange = !this.loginChange}"/>
    <div class="textAreaShell" v-if="!this.loginChange">
      Новый логин
      <textarea class="setTextArea"/>
      <confirm-btn :label="'Сохранить'"/>
    </div>
    <settings-btn :label="'Изменить пароль'" @SetBtnClicked="()=>{this.passChange = !this.passChange}"/>
    <div class="textAreaShell" v-if="!this.passChange">
      Старый пароль:
      <textarea class="setTextArea"/>
      Новый пароль:
      <textarea class="setTextArea"/>
      <confirm-btn :label="'Сохранить'"/>
    </div>
    <settings-btn :label="'Изменить имя'" @SetBtnClicked="()=>{this.nameChange = !this.nameChange}"/>
    <div class="textAreaShell" v-if="!this.nameChange">
      Новое имя:
      <textarea class="setTextArea"/>
      <confirm-btn :label="'Сохранить'"/>
    </div>
    <settings-btn :label="'Изменить возраст'" @SetBtnClicked="()=>{this.ageChange = !this.ageChange}"/>
    <div class="textAreaShell" v-if="!this.ageChange">
      Новый возраст:
      <textarea class="setTextArea"/>
      <confirm-btn :label="'Сохранить'"/>
    </div>
    <settings-btn :label="'Изменить профессию'" @SetBtnClicked="()=>{this.profChange = !this.profChange}"/>
    <div class="textAreaShell" v-if="!this.profChange">
      Новая профессия:
      <textarea class="setTextArea"/>
      <confirm-btn :label="'Сохранить'"/>
    </div>
    <settings-btn :label="'Выйти'" :red="true" @SetBtnClicked="quit"/>

  </div>
</template>

<script>
import SettingsBtn from "@/components/Settings/SettingsBtn";
import ConfirmBtn from "@/components/ConfirmBtn";
export default {
  name: "ProfileComponent",
  components: {ConfirmBtn, SettingsBtn},
  data () {
    return {
      user: null,
      data: null,
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