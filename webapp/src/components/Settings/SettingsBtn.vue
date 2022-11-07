<template>
<div class="setBtnShell" :class="{red:red}">
  <button class="clicker" :class="{red:red}" @click="$emit('SetBtnClicked')">{{this.label}}</button>
</div>
</template>

<script>
export default {
  name: "SettingsBtn",
  props: {
    icon: String,
    label: String,
    red: Boolean,
  }
}
</script>

<style scoped>
.setBtnShell {
  margin-top: 10px;
  width: 90vw;
  height: 50px;
  background: #363537;
  border-top: #94d1be 1px solid;
  border-bottom: #94d1be 1px solid;
}
.setBtnShell:hover {
  background: #94d1be;
}
.setBtnShell.red {
  border-top: red 1px solid;
  border-bottom: red 1px solid;
}
.setBtnShell.red:hover {
  background: red;
}
.clicker {
  background: none;
  border: none;
  width: 90%;
  height: 100%;
  color: #94d1be;
  text-align: left;
  animation: none;
}
.clicker:hover {
  animation-name: slide;
  animation-duration: 500ms;
  animation-iteration-count: 1;
  animation-fill-mode: forwards;
}
.clicker.red {
  color: red;
}
.clicker.red:hover {
  color: #fefefe;
}
@keyframes slide {
  from{
    transform: translateX(0px);
    background: none;
  }
  to{
    transform: translateX(10px);
    background-color: #363537;
  }
}
</style>