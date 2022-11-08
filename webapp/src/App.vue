<template>
  <div id="app">
    <div>
      <HeaderComponent ref="header"/>
    </div>
    <router-view/>
  </div>
</template>

<script>
import HeaderComponent from "@/components/HeaderComponent";
export default {
  name: 'App',
  components:{HeaderComponent}
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Comfortaa&family=Press+Start+2P&family=Scada&display=swap');
#app {
  font-family: 'Scada', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #363537;
  height: max-content;
  min-height: calc(100vh);
  padding-bottom: 10px;
}
</style>
