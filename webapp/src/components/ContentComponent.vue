<template>
<div class="contentShell" @click="$emit('contentClicked')">
  <div class="contentTitle">
    <div class="contentTitleText">
      {{content.title}}
    </div>
    <div class="contentTitleExtra">
      Тип: {{content.type}}<br/>
      Жанр: {{`${content.genre1}, ${content.genre2 ? content.genre2 : ''}, ${content.genre3 ? content.genre3 : ''}`}}
      <div class="contentBtns">
        <like-btn  :liked="liked" :likes="content.usersLiked.length" @likeBtnClick="$emit('likeBtnClick')"/>
        <review-btn/>
      </div>
    </div>
  </div>
  <div class="contentDescription">
    <div class="contentDescriptionText">
      {{content.description}}
    </div>
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
      user: JSON.parse(localStorage.getItem('User'))
    }
  },
  created() {
    this.liked = this.content.usersLiked.includes(this.user.ID)
  },
  watch: {
    likes:{
      handler(val) {
        this.liked = val.includes(this.user.ID)
      },
      immediate: true,
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
  display: initial;
}
.contentDescription {
  width: 60vw;
  background: #efefef;
  border-bottom-right-radius: 20px;
  border-top-right-radius: 20px;
}
.contentTitleText {
  text-wrap: normal;
  font-size: 25px;
  text-align: initial;
  margin-left: 10px;
  margin-top: 10px;
}
.contentTitleExtra {
  text-wrap: normal;
  font-size: 20px;
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
</style>