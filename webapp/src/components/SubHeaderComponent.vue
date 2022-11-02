<template>
  <div class="subHdrCnt">
    <sub-header-button-component
        :label="'Рекомендации'"
        :selected="$route.path === '/content/recommendations'"
        @btnClick="()=>{$router.push('/content/recommendations')}"
    />
    <sub-header-button-component
        :label="'Поиск'"
        :selected="$route.path === '/content/search'"
        @btnClick="()=>{$router.push('/content/search')}"
    />
  </div>
</template>

<script>
import SubHeaderButtonComponent from "@/components/SubHeaderButtonComponent";
export default {
  name: "SubHeaderComponent",
  components: {SubHeaderButtonComponent},
}
</script>

<style scoped>
.subHdrCnt {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-left: auto;
}
</style>