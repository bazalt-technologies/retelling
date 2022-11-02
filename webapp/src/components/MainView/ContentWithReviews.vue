<template>
<div>
  <button-component :label="'Назад'" @btnClick="()=>{$router.push($route.params.route)}"/>
  <content-component
      :key="content.ID"
      :content="content"
      @likeBtnClick="likeClicked(content)"
      @addYourReview="addReview(content)"
  />
  <review-component v-for="r in reviews"
      :key="r.ID"
      :review="r"
      :content-show="false"
      />
  <div v-if="!reviews.length">пока пусто..</div>
  <b-modal ref="new-review-modal" scrollable title="Новое ревью">
    <template #modal-header="{ close }">
      <b-button size="sm" variant="outline-danger" @click="close()">
        Закрыть
      </b-button>
      <!-- Эмулировать встроенное модальное действие кнопки закрытия заголовка -->
    </template>
    <div class="newReviewShell">
      <div class="newReviewBox">
        <div>Выберите контент</div>
        <div>
          <select class="newReviewPick">
            <option value="" selected>{{selectedContent.title}}</option>
          </select>
        </div>
        <div>
          <textarea class="newReviewField" v-model="text"/>
        </div>
      </div>
    </div>
    <template #modal-footer>
      <div>
        <ButtonComponent
            selected
            :label="'Сохранить'"
            :icon="'publish.svg'"
            @btnClick="saveReview"
            class="saveNewReview"
        />
      </div>
    </template>
  </b-modal>
</div>
</template>

<script>
import ContentComponent from "@/components/ContentComponent";
import Vue from "vue";
import ReviewComponent from "@/components/ReviewComponent";
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "ContentWithReviews",
  components: {ButtonComponent, ReviewComponent, ContentComponent},
  props: {
    route: String
  },
  data() {
    return {
      user: null,
      content: null,
      reviews: [],
      text: "",
      selectedContent: null
    }
  },
  created() {
    this.content = this.$route.params.c
    this.selectedContent = this.content
    let users
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users").then(resp=>{
      users = resp && resp.data ? resp.data : []
      this.$store.commit('setAllUsers', users)
    })
    let dateOptions = {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      weekday: 'long',
      timezone: 'UTC',
      hour: 'numeric',
      minute: 'numeric',
      second: 'numeric'
    };
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/reviews", {params: {ObjectID: this.content.ID}})
        .then(response=>{
          this.reviews = response && response.data? response.data.map(r=>{
            return{
              text: r.Review,
              date: new Date(r.Date * 1000).toLocaleString("ru", dateOptions),
              user: this.$store.getters.getAllUsers.find(u=>u.ID===r.UserID).Data.Name
          }
          }) : []
        })
  },
  methods: {
    likeClicked(val) {
      this.user = this.$store.getters.getUser

      if (!this.user) {
        this.$router.push('/login')
      }
      if(!val.usersLiked.includes(this.user.ID)){
        const data = {
          UserID: this.user.ID,
          ObjectID: val.ID
        }
        this.$http.post(Vue.prototype.$baseUrl+"/api/v1/likes", data)
            .then(()=>{
              val.usersLiked.push(this.user.ID)
            })
      } else {
        const data = {
          UserID: this.user.ID,
          ObjectID: val.ID
        }
        this.$http.delete(Vue.prototype.$baseUrl+"/api/v1/likes", {data}).then(()=>{
          let id =val.usersLiked.indexOf(this.user.ID)
          val.usersLiked.splice(id,1)
        })
      }
    },
    addReview(val) {
      this.selectedContent = val
      this.$refs["new-review-modal"].show()
    },
    saveReview() {
      this.user = this.$store.getters.getUser
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
        ContentID: this.selectedContent.ID,
        Date: parseInt(Date.now()/1000)
      }
      this.$http.post(Vue.prototype.$baseUrl+"/api/v1/reviews", data).then(()=>{
        this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {ObjectID:Number(this.user.ID)}}).then(r=>{
          this.user = r && r.data ? r.data[0] : null
          this.$store.commit('setUser', this.user)
          this.$router.push('/content/recommendations')
        })
      })
      this.$refs['new-review-modal'].hide()
    }
  }
}
</script>

<style scoped>

</style>