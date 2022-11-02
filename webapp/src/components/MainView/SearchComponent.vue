<template>
  <div>
  <div>
    <sub-header-component ref="subheader"/>
  </div>
    <div>
      <textarea v-model="searchStr"/>
      <ButtonComponent
          :icon="'filter.svg'"
          class="header-btn"
          @btnClick="()=>{$refs['filter-modal'].show()}"
      />
    </div>
  <content-component v-for="c in filteredContent"
                     :key="c.ID"
                     :content="c"
                     @likeBtnClick="likeClicked(c)"
                     @showReviewClick="()=>{$router.push({name: 'contentReviews', params:{id:c.ID, c, route}})}"
                     @addYourReview="addReview(c)"
  />
    <b-modal ref="filter-modal" scrollable title="Фильтр">
      <template #modal-header>
        <b-button size="sm" variant="outline-danger" @click="closeFilterModal">
          Закрыть
        </b-button>
        <!-- Эмулировать встроенное модальное действие кнопки закрытия заголовка -->
      </template>
     <b-form-group label="Типы контента:">
      <b-form-checkbox-group
          id="checkbox-types"
          v-model="selectedTypes"
          :options="types"
          name="flavour-1"
      ></b-form-checkbox-group>
    </b-form-group>
      <b-form-group label="Жанры:">
        <b-form-checkbox-group
            id="checkbox-genres"
            v-model="selectedGenres"
            :options="genres"
            name="flavour-1"
        ></b-form-checkbox-group>
      </b-form-group>
      <template #modal-footer>
        <div>
          <ButtonComponent
              selected
              :label="'Применить'"
              class="saveNewReview"
              @btnClick="()=>{$refs['filter-modal'].hide()}"
          />
        </div>
      </template>
    </b-modal>
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
import SubHeaderComponent from "@/components/SubHeaderComponent";
import Vue from "vue";
import ButtonComponent from "@/components/ButtonComponent";

export default {
  name: "searchComponent",
  components: {
    ButtonComponent,
    ContentComponent,
    SubHeaderComponent
  },
  data() {
    return {
      liked1: Boolean,
      content: [],
      likes: [],
      user: null,
      searchStr: "",
      route: "",
      selectedTypes: [],
      selectedGenres: [],
      genres: [],
      types: [],
      selectedContent: null,
      text: ""
    }
  },
  computed: {
    filteredContent() {
      let fContent = this.content
      if (this.searchStr!=="") {
        return fContent.filter(c=>(`${c.title.toLowerCase()}`+`${c.type.toLowerCase()}`+
            `${c.genre1.toLowerCase()}`+`${c.genre2?c.genre2.toLowerCase():''}`+`${c.genre3?c.genre3.toLowerCase():''}`+`${c.description}`)
            .includes(this.searchStr.toLowerCase()))
      }
      if (this.selectedTypes.length!==0) {
        return fContent.filter(c=>this.selectedTypes.includes(c.type))
      }
      if (this.selectedGenres.length!==0) {
        return fContent.filter(c=>this.selectedGenres.includes(c.genre1) || this.selectedGenres.includes(c.genre2) || this.selectedGenres.includes(c.genre3))
      }
      return this.content
      }
  },
  beforeCreate() {
    this.user = this.$store.getters.getUser

    if (!this.user) {
      this.$router.push('/login')
    }
  },
  created() {
    this.route = this.$route.path
    this.user = this.$store.getters.getUser
    if (!this.user) {
      this.$router.push('/login')
    }
    this.content = this.$store.getters.getContent || []
    this.filteredContent = this.$store.getters.getContent || []
    this.types = this.$store.getters.getTypes.map(t=>{
      return{
        text: t.Type,
        value: t.Type
      }
    }) || []
    this.genres = this.$store.getters.getGenres.map(g=>{
      return{
        text: g.Genre,
        value: g.Genre
      }
    })  || []
    this.selectedContent = this.content[0]
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
    },
    closeFilterModal() {
      this.selectedGenres=[]
      this.selectedTypes=[]
      this.$refs['filter-modal'].hide()
    }
  }
}
</script>

<style scoped>

</style>