<template>
  <div class="header">
    <img v-bind:src="require('../assets/logoRetelling-short.png')" class="header-logo" alt="logo"/>
    <div class="header-btns">
      <ButtonComponent
          :selected="($route.path==='/content/recommendations') || ($route.path==='/content/search')"
          :label="'Лента'"
          :icon="'lenta.svg'"
          :text-disappear="true"
          @btnClick="()=>{$router.push('/content/recommendations')}"
          class="header-btn"
      ></ButtonComponent>
      <ButtonComponent
          :selected="($route.path==='/profile')"
          :label="'Профиль'"
          :icon="'profile.svg'"
          :text-disappear="true"
          @btnClick="()=>{$router.push('/profile')}"
          class="header-btn"
      ></ButtonComponent>
      <ButtonComponent
          :selected="($route.path==='/settings')"
          :label="'Настройки'"
          :icon="'settings.svg'"
          :text-disappear="true"
          @btnClick="$router.push('/settings')"
          class="header-btn"
      ></ButtonComponent>
    </div>
  </div>
</template>

<script>
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "HeaderComponent",
  components: {
    ButtonComponent,
  }
}
</script>

<style scoped>
.header {
  display: flex;
  flex-direction: row;
  align-content: space-between;
  background-color: #94d1be;
  margin: 0;
  padding: 5px;
}
.header-btns {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-left: auto;
}
.header-btn {
  margin-right: 1vw;
}
.header-logo {
  height: 4vw;
  min-height: 40px;
}
</style>