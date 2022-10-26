<template>
<div>
  <button-component :label="'Назад'" @btnClick="()=>{$router.push('/content')}"/>
  <content-component
      :key="content.ID"
      :content="content"
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
      localStorage.setItem('AllUsers', JSON.stringify(users))
    })
    let dateOptions = {
      era: 'long',
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
              user: JSON.parse(localStorage.getItem('AllUsers')).find(u=>u.ID===r.UserID).Data.Name
          }
          }) : []
        })
  }
}
</script>

<style scoped>

</style>