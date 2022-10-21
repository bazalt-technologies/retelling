<template>

      <button class="likeBtn" :class="{liked: liked, notLiked: !liked}" @click="$emit('likeBtnClick')">
      <img class="likeImg" v-if="!liked" src="../assets/not-liked.png"/>
      <img class="likeImg" v-if="liked" src="../assets/liked.png"/>
      <div class="likeCntr">{{likes}}</div>
      </button>

</template>

<script>

export default {
  name: "LikeBtn",
  props: {
    likes: Number,
    liked: Boolean,
  },
}
</script>

<style scoped>
.likeBtn {
  display: flex;
  justify-content: center;
  align-content: center;
  margin: 5px 5px 10px 0;
  width: 75px;
  border-radius: 10px;
  padding: 5px;
}
.likeImg {
  margin-right: 10px;
}
.likeBtn.liked {
  border: 2px solid rgba(148, 209, 190, 1);
  background: rgba(148, 209, 190, 0.2);
}
.likeBtn.notLiked {
  border: 2px solid #363537;
  background: none;
}
.likeCntr {
  font-size: 16px;
  align-self: center;
}
</style>