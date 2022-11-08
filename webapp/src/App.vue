<template>
  <div id="app" :class="{mobile: isMobile}">
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
  components:{HeaderComponent},
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
}
#app.mobile {
  padding-bottom: 100px;
}
</style>
