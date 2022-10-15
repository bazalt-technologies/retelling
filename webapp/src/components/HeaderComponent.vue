<template>
  <div class="header">
    <img v-bind:src="require('../assets/logoRetelling-short.png')" class="header-logo" alt="logo"/>
    <div class="header-btns">
      <ButtonComponent
          :selected="(modeSelected===1)"
          :label="'Лента'"
          :icon="'lenta.png'"
          @btnClick="modeSelected = 1"
          class="header-btn"
      ></ButtonComponent>
      <ButtonComponent
          :selected="(modeSelected===2)"
          :label="'Профиль'"
          :icon="'profile.png'"
          @btnClick="modeSelected = 2"
          class="header-btn"
      ></ButtonComponent>
      <ButtonComponent
          :selected="(modeSelected===3)"
          :label="'Настройки'"
          :icon="'settings.png'"
          @btnClick="modeSelected = 3"
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
  },
  data() {
    return {
      modeSelected: 1
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
  width: 15vw;
  min-width: 125px;
}
</style>