<template>
  <div class="header">
    <img v-bind:src="require('../assets/logoRetelling-short.png')" class="header-logo" alt="logo"/>
    <div v-if="!isMobile" class="header-btns">
      <ButtonComponent
          :selected="($route.path==='/content/recommendations') || ($route.path==='/content/search')"
          :label="'Лента'"
          :icon="'lenta.svg'"
          :text-disappear="true"
          @btnClick="()=>{$router.push('/content/recommendations')}"
          class="header-btn"
          animated
          shadowed
      ></ButtonComponent>
      <ButtonComponent
          :selected="($route.path==='/profile')"
          :label="'Профиль'"
          :icon="'profile.svg'"
          :text-disappear="true"
          @btnClick="()=>{$router.push('/profile')}"
          class="header-btn"
          animated
          shadowed
      ></ButtonComponent>
      <ButtonComponent
          :selected="($route.path==='/settings')"
          :label="'Настройки'"
          :icon="'settings.svg'"
          :text-disappear="true"
          @btnClick="$router.push('/settings')"
          class="header-btn"
          animated
          shadowed
      ></ButtonComponent>
    </div>
    <div v-if="isMobile" class="mobile">
      <ButtonComponent
          :selected="($route.path==='/content/recommendations') || ($route.path==='/content/search')"
          :label="'Лента'"
          :icon="'lenta.svg'"
          :text-disappear="true"
          @btnClick="()=>{$router.push('/content/recommendations')}"
          class="header-btn"
          animated
          shadow
      ></ButtonComponent>
      <ButtonComponent
          :selected="($route.path==='/profile')"
          :label="'Профиль'"
          :icon="'profile.svg'"
          :text-disappear="true"
          @btnClick="()=>{$router.push('/profile')}"
          class="header-btn"
          animated
          shadow
      ></ButtonComponent>
      <ButtonComponent
          :selected="($route.path==='/settings')"
          :label="'Настройки'"
          :icon="'settings.svg'"
          :text-disappear="true"
          @btnClick="$router.push('/settings')"
          class="header-btn"
          animated
          shadow
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
  },
  data(){
    return {
      width: 0
    }
  },
  computed: {
    isMobile() {
      return this.width<500
    }
  },
  created() {
    this.width = window.innerWidth;
    window.addEventListener('resize', this.updateWidth);
  },
  methods: {
    updateWidth() {
        this.width = window.innerWidth;
    }
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
  width: 100%;
}
.header-btns {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-left: auto;
}
.header-btn {
  margin-right: 1vw;
  border: 0px;
}
.header-logo {
  height: 4vw;
}
.mobile{
  display: flex;
  position: fixed;
  bottom: 0;
  justify-content: space-between;
  z-index: 100;
  padding: 0;
  margin: 0;
  align-content: center;
  background-color: #94d1be;
  width: 100%;
}
</style>