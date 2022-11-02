<template>
<div class="contentShell">
  <div class="contentTitle">
    <div>
    <div class="contentTitleText">
      {{content.title}}
    </div>
    <div class="contentTitleExtra">
      Тип: {{content.type}}<br/>
      Жанр: {{`${content.genre1}, ${content.genre2 ? content.genre2 : ''}, ${content.genre3 ? content.genre3 : ''}`}}
    </div>
    </div>
    <div class="contentBtns">
      <like-btn  :liked="liked" :likes="content.usersLiked.length" @likeBtnClick="$emit('likeBtnClick')"/>
      <review-btn
          :icon="'rate.svg'"
          :label="'Оставить отзыв'"
          :emits="'reviewBtnClick'"
          :text-disappear="true"
          @reviewBtnClick="$emit('addYourReview')"/>
    </div>
  </div>
  <div class="contentDescription">
    <div class="contentDescriptionText">
      {{content.description}}
    </div>
    <review-btn
        v-if="$route.path==='/content/recommendations' || $route.path==='/content/search'"
        :label="'Посмотреть отзывы'"
        :emits="'showReviewClick'"
        :text-disappear="false"
        class="showReviewBtn"
        @showReviewClick="$emit('showReviewClick')"/>
  </div>
</div>
</template>

<script>
import LikeBtn from "@/components/LikeBtn";
import ReviewBtn from "@/components/ReviewBtn";
export default {
  name: "ContentComponent",
  components: {ReviewBtn, LikeBtn},
  props: {
    content: Object,
  },
  data() {
    return{
      liked: false,
      likes: [],
      user: null
    }
  },
  created() {
    this.user = this.$store.getters.getUser
    this.liked = this.content.usersLiked.includes(this.user.ID)
    this.likes = this.content.usersLiked
  },
  watch: {
    likes:{
      handler(val) {
        this.liked = val.includes(this.user.ID)
      },
      deep: true
    }
  },
  updated() {
    this.likes = this.content.usersLiked
  }
}
</script>

<style scoped>
.contentShell {
  width: 100vw;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 25px;
}
.contentTitle {
  width: 35vw;
  background: #fefefe;
  border-bottom-left-radius: 20px;
  border-top-left-radius: 20px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding-left: 10px;
  padding-right: 10px;
}
.contentDescription {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 60vw;
  background: #efefef;
  border-bottom-right-radius: 20px;
  border-top-right-radius: 20px;
}
.contentTitleText {
  text-wrap: normal;
  font-size: 25px;
  font-weight: bold;
  text-align: initial;
  margin-left: 10px;
  margin-top: 10px;
}
.contentTitleExtra {
  text-wrap: normal;
  font-size: 16px;
  color: #363537;

  text-align: initial;
  margin-left: 10px;
  margin-top: 10px;
}
.contentDescriptionText {
  text-align: initial;
  font-size: 16px;
  margin: 10px;
}
.contentBtns {
  display: flex;
  flex-direction: row;
}
.showReviewBtn {
  width: 57vw;
  align-self: center;
}
</style>