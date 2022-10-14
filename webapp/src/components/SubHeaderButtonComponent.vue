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
  name: "SubHeaderButtonComponent.vue",
  props: {
    label: String,
    selected: Boolean
  },
}
</script>

<style scoped>

</style>