<template>
<div>
  <div>
    <sub-header-component ref="subheader"/>
  </div>
  <content-component v-for="c in content"
      :key="c.ID"
      :title="c.title"
      :type="c.type"
      :genre="`${c.genre1}, ${c.genre2 ? c.genre2 : ''}, ${c.genre3 ? c.genre3 : ''}`"
      :likes="likes1"
      :liked="liked1"
      @likeBtnClick="liked1 = !liked1; likes1 += liked1*2 -1"
      :description="c.description"
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
      likes1: 93,
      content: [],
      genres: [],
      types: []
    }
  },
  created() {

    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/genres").then(response =>{
      this.genres = response && response.data ? response.data : []
    })
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/types").then(response =>{
      this.types = response && response.data ? response.data : []
    })
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/content").then(response =>{
      this.content = response && response.data ? response.data.map(c=>{
        return {
          title: c.Title,
          description: c.Description,
          type: this.types.find(t => t.ID === c.TypeID).Type,
          genre1: this.genres.find(g => g.ID === c.GenreID1).Genre,
          genre2: c.GenreID2 ? this.genres.find(g => g.ID ===c.GenreID2).Genre : 0,
          genre3: c.GenreID3 ? this.genres.find(g => g.ID ===c.GenreID3).Genre : 0,

        }
      }) : []
    })
    localStorage.setItem('Genres', JSON.stringify(this.genres))
    localStorage.setItem('Types', JSON.stringify(this.types))
  }
}
</script>

<style scoped>

</style>