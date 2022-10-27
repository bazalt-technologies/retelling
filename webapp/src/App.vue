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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #363537;
  height: max-content;
  min-height: calc(100vh - 10px);
  padding-bottom: 10px;
}
</style>
