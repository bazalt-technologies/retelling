<template>
  <div class="contentShell">
    <div class="contentTitle">
      <div class="contentTitleText">
        Имя пользователя: {{user.Data.Name}}
      </div>
      <div class="contentTitleExtra">
        </div>
    </div>
    <div class="contentDescription">
      <div class="contentDescriptionText">
        Возраст: {{user.Data.Age}}<br/>
        {{ `${user.Data.Profession ? 'Профессия : '+user.Data.Profession : ''}` }}
        Количество ревью : {{user.Data.ReviewCount}}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProfileComponent",
  data () {
    return {
      user: null,
      data: null
    }
  },
  created() {
    this.user = JSON.parse(localStorage.getItem('User'))
  }
}
</script>

<style scoped>
.contentShell {
  width: 100vw;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 25px;
}
.contentTitle {
  width: 35vw;
  background: #fefefe;
  border-bottom-left-radius: 20px;
  border-top-left-radius: 20px;
  display: initial;
}
.contentDescription {
  width: 60vw;
  background: #efefef;
  border-bottom-right-radius: 20px;
  border-top-right-radius: 20px;
}
.contentTitleText {
  text-wrap: normal;
  font-size: 25px;
  text-align: initial;
  margin-left: 10px;
  margin-top: 10px;
}
.contentTitleExtra {
  text-wrap: normal;
  font-size: 20px;
  text-align: initial;
  margin-left: 10px;
  margin-top: 10px;
}
.contentDescriptionText {
  text-align: initial;
  font-size: 16px;
  margin: 10px;
}
</style>