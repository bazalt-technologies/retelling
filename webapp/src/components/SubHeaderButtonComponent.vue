<template>
  <button
      class="subBtn"
      @click="$emit('btnClick')"
      :class="{
        selected: selected
      }"
  >
    {{label}}
  </button>
</template>

<script>
export default {
  name: "SubHeaderButtonComponent",
  props: {
    label: String,
    selected: Boolean
  }
}
</script>

<style scoped>
.subBtn {
  border: none;
  border-bottom-left-radius: 25px;
  border-bottom-right-radius: 25px;
  background: #fefefe;
  width: 50vw;
  padding: 4px;
  cursor: pointer;
}
.subBtn.selected {
  background: #efefef;
}
</style>