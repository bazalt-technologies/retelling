<template>

      <button class="likeBtn" :class="{liked: liked, notLiked: !liked}" @click="$emit('likeBtnClick')">
      <img class="likeImg" v-if="!liked" src="../assets/not-liked.svg"/>
      <img class="likeImg" :class="{liked: liked}" v-if="liked" src="../assets/liked.svg"/>
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
  margin-right: 5px;
}
.likeImg.liked{
  animation-name: like;
  animation-iteration-count: 2;
  animation-duration: 500ms;
  animation-fill-mode: backwards;
}
.likeBtn.liked {
  border: 2px solid rgba(148, 209, 190, 1);
  background: rgba(148, 209, 190, 0.2);

}
.likeBtn.notLiked {
  border: 2px solid #363537;
  background: none;
}
.likeBtn:hover{
  animation-name: upper;
  animation-duration: 300ms;
  animation-iteration-count: 1;
  animation-fill-mode: forwards;
}
.likeCntr {
  font-size: 16px;
  align-self: center;
}
@keyframes like {
  from {
    transform: scale3d(1,1,1);
  }
  50% {
    transform: scale3d(1.2,1.2,1.2);
  }
  to {
    transform: scale3d(1,1,1);
  }
}
</style>