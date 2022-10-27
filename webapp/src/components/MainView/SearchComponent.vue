<template>
  <div>
  <div>
    <sub-header-component ref="subheader"/>
  </div>
    <div>
      <textarea v-model="searchStr"/>
    </div>
  <content-component v-for="c in filteredContent"
                     :key="c.ID"
                     :content="c"
                     @likeBtnClick="likeClicked(c)"
                     @showReviewClick="()=>{$router.push({name: 'contentReviews', params:{id:c.ID, c, route}})}"
  />
  </div>
</template>

<script>
import ContentComponent from "@/components/ContentComponent";
import SubHeaderComponent from "@/components/SubHeaderComponent";
import Vue from "vue";

export default {
  name: "searchComponent",
  components: {
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
      route: ""
    }
  },
  computed: {
    filteredContent() {
      return this.searchStr==="" ? this.content
          : this.content.filter(c=>c.title.toLowerCase().includes(this.searchStr.toLowerCase()))
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
    }
  }
}
</script>

<style scoped>

</style>