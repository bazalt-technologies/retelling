<template>
  <div>
    <ReviewComponent v-for="r in reviews"
      :key="r.ID"
      :review="r"
    />
  </div>
</template>

<script>
import Vue from "vue";
import ReviewComponent from "@/components/ReviewComponent";
export default {
  name: "ProfileReviewsComponent",
  components: {ReviewComponent},
  data() {
    return {
      reviews: []
    }
  },
  created() {
    let usr = JSON.parse(localStorage.getItem('User'))
    if (!usr.Data.ReviewCount){
      this.reviews = []
      return
    }
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
    let content = []
    this.$http.get(Vue.prototype.$baseUrl+"/api/v1/content").then(response => {
      content = response && response.data ? response.data : []
    })
      this.$http.get(Vue.prototype.$baseUrl+"/api/v1/reviews", {UserID: usr.ID}).then(response => {
      this.reviews = response.data ? response.data.map(r=>{
        return {
          title: content.find(c=>c.ID===response.data.ContentID).Title,
          user: usr.Data.Name,
          text: r.Review,
          date: new Date(r.Date * 1000).toLocaleString("ru", dateOptions)
        }
      }) : []

    })
  }
}
</script>

<style scoped>

</style>