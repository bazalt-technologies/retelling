<template>
  <div>
    <ReviewComponent v-for="r in reviews"
      :key="r.ID"
      :review="r"
      :is-user="true"
      @deleteReview="deleteR(r)"
    />
    <div>
      <ButtonComponent
          :selected="($route.path==='/content')"
          :label="'Новое ревью'"
          :icon="'add.svg'"
          @btnClick="newReview"
          class="header-btn"
      />
    </div>
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
      user: null
    }
  },
  beforeMount() {
    this.user = JSON.parse(localStorage.getItem('User'))
    if (!this.user) {
      this.$router.push('/login')
    }
  },
  created() {
    this.user = JSON.parse(localStorage.getItem('User'))
    if (!this.user.Data.ReviewCount){
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
      this.$http.get(Vue.prototype.$baseUrl+"/api/v1/reviews", {UserID: this.user.ID}).then(response => {
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

  },
  updated() {
    let usr = JSON.parse(localStorage.getItem('User'))
    if (!usr) {
      this.$router.push('/login')
    }
  },
  methods: {
    newReview() {
      this.$router.push({name:"newReview"})
    },
    deleteR(val) {
      const data = {ID: val.id, UserID: this.user.ID}
      this.$http.delete(Vue.prototype.$baseUrl+"/api/v1/reviews", {data}).then(()=>{
        this.$http.get(Vue.prototype.$baseUrl+"/api/v1/users", {params: {ObjectID:Number(this.user.ID)}}).then(r=>{
          this.user = r && r.data ? r.data[0] : null
          localStorage.setItem('User', JSON.stringify(this.user))
          this.$router.push('/content')
        })
      })
    }
  }
}
</script>

<style scoped>

</style>