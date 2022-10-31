<template>
<div class="newReviewShell">
  <div class="newReviewBox">
  <div>
  <select v-model="selectedContent" class="newReviewPick">
    <option disabled value="">Выберите контент для ревью</option>
    <option v-for="(c) in content" :key="c.ID" :value="c.ID">{{ c.title }}</option>
  </select>
  </div>
  <div>
    <textarea class="newReviewField" v-model="text"/>
  </div>
  <ButtonComponent
      :label="'Сохранить'"
      :icon="'publish.svg'"
      @btnClick="saveReview"
      class="saveNewReview"
  />
  </div>
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
      this.saveReview()
    },
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
.newReviewShell {
  width: 100vw;
  display: flex;
  align-items: center;
  flex-direction: column;
  margin-top: 25px;
}
.newReviewBox {
  width: 300px;
  height: fit-content;
  background: #efefef;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  padding: 5px;
  border-radius: 20px;
}
.saveNewReview {
  align-self: center;
  margin: 0;
}
.newReviewPick {
  border: none;
  background: none;
  border-bottom: #363537 2px solid;
  width: 75%;
  margin: 10px;
}
.newReviewField {
  border: none;
  background: none;
  border-bottom: #363537 2px solid;
  width: 75%;
  height: fit-content;
  margin: 10px;
}
</style>