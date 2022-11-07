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
  width: 100%;
  height: 100%;
  color: #94d1be;
  text-align: left;
}
.clicker:hover {
  color: #363537;
}
.clicker.red {
  color: red;
}
.clicker.red:hover {
  color: #fefefe;
}
</style>