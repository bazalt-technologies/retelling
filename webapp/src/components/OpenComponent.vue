<template>
  <div id="#open">
  <!--- Content --->
  <div>
    <div>
      <buttonComponent
          :selected="false"
          :label="'Войти'"
          @btnClick="()=>{$router.push('/login')}"
          animated
      />
      <buttonComponent
          :selected="false"
          :label="'Регистрация'"
          @btnClick="()=>{$router.push('/registration')}"
          animated
      />
    </div>
  </div>
  </div>
</template>

<script>
import ButtonComponent from "@/components/ButtonComponent";



export default {
  name: "OpenComponent",
  components: {
    ButtonComponent
  },
  created() {
    let user = this.$store.getters.getUser;
    if (user) {this.$router.push('/content/recommendations')}
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Comfortaa&display=swap');
#open {
  font-family: 'Comfortaa', cursive;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #363537;
  height: 100vh;
  max-height: 100vh;
}
</style>
