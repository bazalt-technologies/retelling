<template>
<button
    class="stdBtn"
    :class="{
  selected: selected
}"
    @click="$emit('btnClick')"
>
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
  }
}
</script>

<style scoped>
.stdBtn {
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