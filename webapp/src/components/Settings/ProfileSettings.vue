<template>
  <div class="reviewShell">
    <div class="reviewTitle">
      <div class="reviewTitleText">
        Имя пользователя: {{user.Data.Name}}
      </div>
      <ButtonComponent
          :label="'Выйти'"
          :icon="'logout.svg'"
          @btnClick="quit"
          class="quit-btn"
      />
    </div>
    <div class="reviewDescription">
      <div class="reviewDescriptionText">
        Возраст: {{user.Data.Age}}<br/>
        {{ `${user.Data.Profession ? 'Профессия : '+user.Data.Profession : ''}` }}
        Количество ревью : {{user.Data.ReviewCount}}
      </div>
    </div>
    <div>
      <ButtonComponent
          :label="'изменить'"
          @btnClick="()=>{$router.push('/settings/profile')}"
          :icon="'edit.svg'"
          class="header-btn"
      />
    </div>
  </div>
</template>

<script>
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "ProfileComponent",
  components: {ButtonComponent},
  data () {
    return {
      user: null,
      data: null
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
.reviewShell {
  width: 100vw;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 25px;
}
.reviewTitle {
  width: 35vw;
  background: #fefefe;
  border-bottom-left-radius: 20px;
  border-top-left-radius: 20px;
  display: initial;
}
.reviewDescription {
  width: 60vw;
  background: #efefef;
  border-bottom-right-radius: 20px;
  border-top-right-radius: 20px;
}
.reviewTitleText {
  text-wrap: normal;
  font-size: 25px;
  text-align: initial;
  margin-left: 10px;
  margin-top: 10px;
}
.reviewDescriptionText {
  text-align: initial;
  font-size: 16px;
  margin: 10px;
}
.quit-btn {
  margin: 10px;
}
</style>