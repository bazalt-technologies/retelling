<template>
<button
    class="stdBtn"
    :class="{
  selected: selected
}"
    @click="$emit('btnClick')"
>
  <img class="btnIcon"  v-if="icon !== undefined" v-bind:src="require('../assets/'+`${icon}`)" alt="">
  {{label}}
</button>
</template>

<script>
export default {
  name: "ButtonComponent",
  props: {
    label: String,
    icon: String,
    selected: Boolean
  },

}
</script>

<style scoped>
.stdBtn {
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 5px 10px 5px 5px;
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
  margin-right: 5px;
}
</style>