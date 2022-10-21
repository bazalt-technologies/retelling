<template>
<div class="contentShell">
  <div class="contentTitle">
    <div class="contentTitleText">
      {{title}}
    </div>
    <div class="contentTitleExtra">
      Тип: {{type}}<br/>
      Жанр: {{genre}}
      <like-btn  :liked="liked" :likes="likes" @likeBtnClick="$emit('likeBtnClick')"/>
    </div>
  </div>
  <div class="contentDescription">
    <div class="contentDescriptionText">
      {{description}}
    </div>
  </div>
</div>
</template>

<script>
import LikeBtn from "@/components/LikeBtn";
export default {
  name: "ContentComponent",
  components: {LikeBtn},
  props: {
    title: String,
    type: String,
    genre: String,
    likes: Number,
    liked: Boolean,
    description: String,
  },
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
</style>