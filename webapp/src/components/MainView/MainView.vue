<template>
<div>
  <div>
    <sub-header-component ref="subheader"/>
  </div>
  <content-component v-for="c in content"
      :key="c.ID"
      :content="c"
      @likeBtnClick="likeClicked(c)"
      @contentClicked="()=>{$router.push({name: 'contentReviews', params:{id:c.ID, c}})}"
  />
</div>
</template>

<script>
import ContentComponent from "@/components/ContentComponent";
import SubHeaderComponent from "@/components/SubHeaderComponent";
import Vue from "vue";
export default {
  name: "MainView",
  components: {
    ContentComponent,
    SubHeaderComponent
  },
  data() {
    return {
      liked1: Boolean,
      content: [],
      likes: []
    }
  },
  beforeCreate() {
    let user = this.$store.getters.getUser

    if (!user) {
      this.$router.push('/login')
    }
  },
  created() {
    let user = this.$store.getters.getUser

    if (!user) {
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
      }) : []
      this.$store.commit('setContent', this.content)
    })
  },
  methods: {
    likeClicked(val) {
      let user = this.$store.getters.getUser

      if (!user) {
        this.$router.push('/login')
      }
      if(!val.usersLiked.includes(user.ID)){
        const data = {
          UserID: user.ID,
          ObjectID: val.ID
        }
        this.$http.post(Vue.prototype.$baseUrl+"/api/v1/likes", data)
            .then(()=>{
              val.usersLiked.push(this.user.ID)
            })
        return
      } else {
        const data = {
          UserID: user.ID,
          ObjectID: val.ID
        }
        this.$http.delete(Vue.prototype.$baseUrl+"/api/v1/likes", {data}).then(()=>{
          let id =val.usersLiked.indexOf(user.ID)
          val.usersLiked.splice(id,1)
        })
      }
    }
  }
}
</script>

<style scoped>

</style>