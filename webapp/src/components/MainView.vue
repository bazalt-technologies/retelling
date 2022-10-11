<template>
<div>hello, {{name}}</div>
</template>

<script>
export default {
  name: "MainView",
  props:{
    name: {
      type: String,
      required: true
    }
  }
}
</script>

<style scoped>

</style>