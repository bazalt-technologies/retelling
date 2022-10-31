<template>
<div>
  <button-component :label="'Назад'" @btnClick="()=>{$router.push($route.params.route)}"/>
  <content-component
      :key="content.ID"
      :content="content"
      @likeBtnClick="likeClicked(content)"
  />
  <review-component v-for="r in reviews"
      :key="r.ID"
      :review="r"
      />
  <div v-if="!reviews.length">пока пусто..</div>
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
      content: null,
      reviews: []
    }
  },
  created() {
    this.content = this.$route.params.c
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
    }
  }
}
</script>

<style scoped>

</style>