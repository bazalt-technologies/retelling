<template>
  <button class="reviewBtn" @click="$emit(emits)">
    <img class="reviewImg"  v-if="icon !== undefined" :src="require('../assets/'+`${icon}`)"/>
    <div class="reviewTxt" v-if="this.winW > 600 || !this.textDisappear" >{{label}}</div>
  </button>
</template>

<script>
export default {
  name: "ReviewBtn",
  props: {
    label: String,
    emits: String,
    icon: String,
    textDisappear: Boolean,
  },
  data() {
    return {
      winW: null
    }
  },
  created()  {
    window.addEventListener("resize", this.resizeHandler);
  },
  destroyed()  {
    window.removeEventListener("resize", this.resizeHandler);
  },
  mounted()  {
    this.winW =  window.innerWidth;
  },
  methods: {
    resizeHandler() {
      this.winW = window.innerWidth
    }
  }
}
</script>

<style scoped>
.reviewBtn {
  display: flex;
  justify-content: center;
  align-content: center;
  margin: 5px 5px 10px 0;
  width: fit-content;
  border-radius: 10px;
  padding: 5px;
  border: 2px solid #363537;
  background: none;
  cursor: pointer;
}
.reviewImg {
  margin-right: 5px;
}
.reviewTxt {
  font-size: 16px;
  align-self: center;
}
.reviewBtn:hover{
  animation-name: upper;
  animation-duration: 300ms;
  animation-iteration-count: 1;
  animation-fill-mode: forwards;
}
</style>