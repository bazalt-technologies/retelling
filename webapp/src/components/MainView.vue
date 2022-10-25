<template>
<div>
  <div>
    <sub-header-component ref="subheader"/>
  </div>
  <content-component v-for="c in content"
      :key="c.ID"
      :content="c"
      @likeBtnClick="likeClicked(c)"
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
      likes: [],
      user: null
    }
  },
  created() {
    this.user = JSON.parse(localStorage.getItem('User'))
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/genres").then(response =>{
      let genres = response && response.data ? response.data : []
      localStorage.setItem('Genres', JSON.stringify(genres))
    })
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/types").then(response =>{
      let types = response && response.data ? response.data : []
      localStorage.setItem('Types', JSON.stringify(types))
    })
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/content").then(response =>{
      this.content = response && response.data ? response.data.map(c=>{
        return {
          ID: c.ID,
          title: c.Title,
          description: c.Description,
          type: JSON.parse(localStorage.getItem('Types')).find(t => t.ID === c.TypeID).Type,
          genre1: JSON.parse(localStorage.getItem('Genres')).find(g => g.ID === c.GenreID1).Genre,
          genre2: c.GenreID2 ? JSON.parse(localStorage.getItem('Genres')).find(g => g.ID ===c.GenreID2).Genre : 0,
          genre3: c.GenreID3 ? JSON.parse(localStorage.getItem('Genres')).find(g => g.ID ===c.GenreID3).Genre : 0,
          usersLiked: c.UsersLiked || []
        }
      }) : []
    })
  },
  methods: {
    likeClicked(val) {
      if(!val.usersLiked.includes(this.user.ID)){
        const data = {
          UserID: this.user.ID,
          ContentID: val.ID
        }
        this.$http.post(Vue.prototype.$baseUrl+"/api/v1/likes", data)
            .then(()=>{
              val.usersLiked.push(this.user.ID)
            })
        return
      } else {
        const data = {
          UserID: this.user.ID,
          ContentID: val.ID
        }
        this.$http.delete(Vue.prototype.$baseUrl+"/api/v1/likes", {data}).then(()=>{
          let id =val.usersLiked.indexOf(this.user.ID)
          val.usersLiked.splice(id,1)
        })
      }
    }
  }
}
</script>

<style scoped>

</style>