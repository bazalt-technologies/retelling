<template>
<button
    class="stdBtn"
    :class="{
  selected: selected
}"
    @click="$emit('btnClick')"
>
  <img v-bind:src="require('../assets/'+`${icon}`)" alt="">
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
  created() {
    console.log(this.icon)
  }
}

</script>

<style scoped>
.stdBtn {
  display: flex;
  flex-direction: row;
  padding: 10px;
  text-align: center;
  text-transform: uppercase;
  cursor: pointer;
  color: #94d1be;
  border: none;
  border-radius: 25px;
  background: #363537;
}
.stdBtn.selected{
  background: #fefefe;
}
</style>