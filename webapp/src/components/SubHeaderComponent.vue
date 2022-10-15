<template>
  <div class="subHdrCnt">
    <sub-header-button-component
        :label="'Рекомендации'"
        :selected="recSelected"
        @btnClick="recSelected = true"
    />
    <sub-header-button-component
        :label="'Поиск'"
        :selected="!recSelected"
        @btnClick="recSelected = false"
    />
  </div>
</template>

<script>
import SubHeaderButtonComponent from "@/components/SubHeaderButtonComponent";
export default {
  name: "SubHeaderComponent",
  components: {SubHeaderButtonComponent},
  data() {
    return {
      recSelected: true
    }
  }
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