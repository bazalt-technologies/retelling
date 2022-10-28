<template>
<div class="reviewShell">
  <div>
  <select v-model="selectedContent">
    <option disabled value="">Выберите контент для ревью</option>
    <option v-for="(c) in content" :key="c.ID" :value="c.ID">{{ c.title }}</option>
  </select>
  </div>
  <div>
    <textarea v-model="text"/>
  </div>
  <ButtonComponent
      :label="'Сохранить'"
      :icon="'publish.svg'"
      @btnClick="saveReview"
      class="header-btn"
  />
</div>
</template>

<script>
import ButtonComponent from "@/components/ButtonComponent";
import Vue from "vue";
export default {
  name: "NewReviewComponent",
  components: {ButtonComponent},
  data () {
    return {
      user: null,
      content: [],
      text: "",
      selectedContent: 0
    }
  },
  created() {
    this.content = this.$store.getters.getContent
    this.user = this.$store.getters.getUser
  },
  methods: {
    saveReview() {
      if(!this.selectedContent) {
        alert("Ошибка! Выберите контент!")
        return
      }
      if(!this.text) {
        alert("Ошибка! Введите текст ревью!")
        return
      }
      const data = {
        UserID: this.user.ID,
        Review: this.text,
        ContentID: this.selectedContent,
        Date: parseInt(Date.now()/1000)
      }
      this.$http.post(Vue.prototype.$baseUrl+"/api/v1/reviews", data).then(()=>{
        this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {ObjectID:Number(this.user.ID)}}).then(r=>{
          this.user = r && r.data ? r.data[0] : null
          this.$store.commit('setUser', this.user)
          this.$router.push('/profile')
        })
      })
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
</style>