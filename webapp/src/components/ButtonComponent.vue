<template>
<button
    class="stdBtn"
    :class="{
  selected: selected,
  animation: animated,
  shadowed: shadowed
}"
    @click="$emit('btnClick')"
>
  <img class="btnIcon"
       :class="{withText: (this.winW > 500 || !this.textDisappear) && label !== undefined}"
       v-if="icon !== undefined" v-bind:src="require('../assets/'+`${icon}`)" alt="">
  <div v-if="this.winW > 500 || !this.textDisappear">{{label}}</div>
</button>
</template>

<script>
export default {
  name: "ButtonComponent",
  props: {
    label: String,
    icon: String,
    selected: Boolean,
    textDisappear: Boolean,
    animated: Boolean,
    shadowed: Boolean
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

<style>
.stdBtn {
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 5px 10px 5px 10px;
  text-align: center;
  font-size: 15px;
  cursor: pointer;
  color: #94d1be;
  border: none;
  border-radius: 25px;
  background: #363537;
}
.stdBtn.selected{
  background: #fefefe;
}
.stdBtn.shadowed{
  box-shadow: 5px 5px 10px black;
}
.stdBtn.animation:hover{
  animation-name: upper;
  animation-duration: 300ms;
  animation-iteration-count: 1;
  animation-fill-mode: forwards;
}
.btnIcon {
  width: 32px;
  height: 32px;
  margin-right: 0;
}
.btnIcon.withText {
  margin-right: 5px;
}
@keyframes upper {
  from{
    transform: translateY(0px);
  }
  to{
    transform: translateY(-5px);
  }
}
</style>