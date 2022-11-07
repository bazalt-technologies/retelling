<template>
<button class="confirmBtn">{{this.label}}</button>
</template>

<script>
export default {
  name: "ConfirmBtn",
  props: {
    label: String,
  }
}
</script>

<style scoped>
.confirmBtn {
  background: none;
  border: none;
  background: #94d1be;
  border-radius: 25px;
  height: 35px;
}
</style>