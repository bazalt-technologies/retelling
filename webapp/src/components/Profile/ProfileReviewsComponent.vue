<template>
  <div class="profileWithReviewsShell">
    <ReviewComponent v-for="r in reviews"
      :key="r.ID"
      :review="r"
      :is-user="true"
      :content-show="true"
      @deleteReview="deleteR(r)"
    />
    <div>
      <ButtonComponent
          :label="'Новое ревью'"
          :icon="'add.svg'"
          @btnClick="newReview"
          class="header-btn"
      />
    </div>
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
            <select v-model="selectedContent" class="newReviewPick">
              <option disabled selected="selected" value="">Выберите контент для ревью</option>
              <option v-for="(c) in content" :key="c.ID" :value="c.ID">{{ c.title }}</option>
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
import Vue from "vue";
import ReviewComponent from "@/components/ReviewComponent";
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "ProfileReviewsComponent",
  components: {ButtonComponent, ReviewComponent},
  data() {
    return {
      reviews: [],
      user: null,
      content: [],
      text: "",
      selectedContent: 0
    }
  },
  created() {
    this.user = this.$store.getters.getUser
    if (!this.user) {
      this.$router.push('/login')
    }
    if (!this.user.Data.ReviewCount){
      this.reviews = []
      return
    }
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
    let content = []
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/content").then(response => {
      content = response && response.data ? response.data : []
      this.$http.get(Vue.prototype.$baseUrl+"/api/v1/reviews", {params:
      {
        UserID: this.user.ID
      }
    }).then(response => {
        this.reviews = response.data ? response.data.map(r=>{
          return {
            id: r.ID,
            title: content.find(c=>c.ID===r.ContentID).Title,
            user: this.user.Data.Name,
            text: r.Review,
            date: new Date(r.Date * 1000).toLocaleString("ru", dateOptions)
          }
        }) : []

      })
    })
    this.content = this.$store.getters.getContent
  },
  beforeMount() {
    document.addEventListener("keydown", this.onKeyDown)
  },
  beforeDestroy() {
    document.removeEventListener("keydown", this.onKeyDown)
  },
  updated() {
    let usr = this.$store.getters.getUser
    if (!usr) {
      this.$router.push('/login')
    }
  },
  methods: {
    newReview() {
      this.$refs['new-review-modal'].show()
    },
    deleteR(val) {
      const data = {ID: val.id, UserID: this.user.ID}
      let idx = this.reviews.indexOf(this.reviews.find(r=>r.id===val.id)[0])
      this.reviews.splice(idx,1)
      this.$http.delete(Vue.prototype.$baseUrl+"/api/v1/reviews", {data}).then(()=>{
        this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {ObjectID:Number(this.user.ID)}}).then(r=>{
          this.user = r && r.data ? r.data[0] : null
          this.$store.commit('setUser', this.user)
        })
      })
    },
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
          this.$router.push('/content/recommendations')
        })
      })
      this.$refs['new-review-modal'].hide()
    }
  }
}
</script>

<style>

.newReviewShell {
  display: flex;
  align-items: center;
  flex-direction: column;
  background: #94d1be;
  border-radius: 20px;
}
.newReviewBox {
  width: 400px;
  height: 300px;
  background: white;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 5px;
  border-radius: 20px;
  margin: 35px;
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
  height: 100px;
  border: none;
  background: none;
  border-bottom: #363537 2px solid;
  width: 75%;
  margin: 10px;
}
</style>