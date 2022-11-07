<template>
<button
    class="stdBtn"
    :class="{
  selected: selected
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
.btnIcon {
  width: 32px;
  height: 32px;
  margin-right: 0;
}
.btnIcon.withText {
  margin-right: 5px;
}
</style>