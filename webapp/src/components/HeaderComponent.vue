<template>
  <div class="header">
    <img v-bind:src="require('../assets/logoRetelling.png')" class="header-logo" alt="logo"/>
    <div class="header-btn">
      <ButtonComponent
          :selected=false
          :label="'Лента'"
          :icon="'lenta.png'"
          style="margin-left: 5px; margin-right: 5px"
      ></ButtonComponent>
      <ButtonComponent
          :selected=true
          :label="'Профиль'"
          :icon="'profile.png'"
          style="margin-left: 5px; margin-right: 5px"
      ></ButtonComponent>
      <ButtonComponent
          :selected=false
          :label="'Настройки'"
          :icon="'settings.png'"
          style="margin-left: 5px; margin-right: 5px"
      ></ButtonComponent>
    </div>
  </div>
</template>

<script>
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "HeaderComponent.vue",
  components: {
    ButtonComponent,
  },
}
</script>

<style scoped>
.header {
  display: flex;
  flex-direction: row;
  align-content: space-between;
  background-color: #94d1be;
  margin: 0;
}
.header-btn {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-left: auto;
}
.header-logo {
  height: 100px
}
</style>