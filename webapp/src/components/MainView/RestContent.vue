<template>
  <div>
    <div>
      <sub-header-component ref="subheader"/>
    </div>
    <div class="animated">
      <content-component v-for="c in content"
                         :key="c.ID"
                         :content="c"
                         @likeBtnClick="likeClicked(c)"
                         @showReviewClick="()=>{$router.push({name: 'contentReviews', params:{id:c.ID, c, route}})}"
                         @addYourReview="addReview(c)"
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
              animated
          />
        </div>
      </template>
    </b-modal>

  </div>
</template>

<script>
import ContentComponent from "@/components/ContentComponent";
import SubHeaderComponent from "@/components/SubHeaderComponent";
import Vue from "vue";
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "RestContent",
  components: {
    ContentComponent,
    SubHeaderComponent,
    ButtonComponent
  },
  data() {
    return {
      liked1: Boolean,
      content: [],
      likes: [],
      user: null,
      route: "",
      selectedContent: null,
      text: "",
      render: true
    }
  },
  beforeCreate() {
    this.user = this.$store.getters.getUser

    if (!this.user) {
      this.$router.push('/login')
    }
  },
  created() {
    this.user = this.$store.getters.getUser
    this.route = this.$route.path
    if (!this.user) {
      this.$router.push('/login')
    }
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/genres").then(response =>{
      let genres = response && response.data ? response.data : []

      this.$store.commit('setGenres', genres)
    })
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/types").then(response =>{
      let types = response && response.data ? response.data : []
      this.$store.commit('setTypes', types)
    })
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/content").then(response =>{
      this.content = response && response.data ? response.data.map(c=>{
        return {
          ID: c.ID,
          title: c.Title,
          description: c.Description,
          type: this.$store.getters.getTypes.find(t => t.ID === c.TypeID).Type,
          genre1: this.$store.getters.getGenres.find(g => g.ID === c.GenreID1).Genre,
          genre2: c.GenreID2 ? this.$store.getters.getGenres.find(g => g.ID ===c.GenreID2).Genre : 0,
          genre3: c.GenreID3 ? this.$store.getters.getGenres.find(g => g.ID ===c.GenreID3).Genre : 0,
          usersLiked: c.UsersLiked || []
        }
      }).filter(c=>c.usersLiked.length<=5) : []
      this.selectedContent = this.content[0]
    })
    this.render = false
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

<style>
.animated{
  animation-name: slide-content;
  animation-duration: 500ms;
  animation-iteration-count: 1;
}

@keyframes slide-content {
  from{
    transform: translateY(100%);
  }
  to {
    transform: translateY(0%);
  }
}
</style>