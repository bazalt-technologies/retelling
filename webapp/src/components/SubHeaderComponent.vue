<template>
  <div class="subHdrCnt">
    <sub-header-button-component :label="'Рекомендации'" :selected="true"/>
    <sub-header-button-component :label="'Поиск'" :selected="true"/>
  </div>
</template>

<script>
import SubHeaderButtonComponent from "@/components/SubHeaderButtonComponent";
export default {
  name: "SubHeaderComponent",
  components: {SubHeaderButtonComponent}
}
</script>

<style scoped>
.subHdrCnt {
  display: flex;
  flex-direction: row;
  align-content: space-between;
  align-self: auto;
  width: 100%;
}
</style>