<template>
  <div class="reviewWithTrash">
  <div class="reviewShell">
    <div class="reviewTitle">
      <div class="reviewTitleText">
        {{review.user}}
      </div>
      <div class="reviewTitleExtra">
        <div v-if="isUser">
        {{review.title}}<br/>
        </div>
        {{review.date}}<br/>
      </div>
    </div>
    <div class="reviewDescription">
      <div class="reviewDescriptionText">
        {{review.text}}
      </div>
    </div>
  </div>
    <ButtonComponent v-if="isUser"
        :icon="'delete.svg'"
        @btnClick="$emit('deleteReview')"
        class="header-btn"/>
  </div>
</template>

<script>
import ButtonComponent from "@/components/ButtonComponent";
export default {
  name: "ReviewComponent",
  components: {ButtonComponent},
  props: {
    isUser:Boolean,
    review: Object,
    contentShow: Boolean,
  }
}
</script>

<style scoped>
.reviewShell {
  display: flex;
  flex-direction: column;
  justify-content: center;
  margin-top: 25px;
  align-items: center;
}
.reviewTitle {
  width: 90vw;
  background: #94d1be;
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
  display: initial;
}
.reviewDescription {
  width: 90vw;
  background: #fefefe;
  border-top: #94d1be 2px solid;
  border-bottom-right-radius: 20px;
  border-bottom-left-radius: 20px;
}
.reviewTitleText {
  text-wrap: normal;
  font-size: 25px;
  font-weight: bold;
  text-align: initial;
  margin-left: 10px;
  margin-top: 10px;
}
.reviewTitleExtra {
  text-wrap: normal;
  font-size: 14px;
  color: #363537;
  text-align: initial;
  margin-left: 10px;
}
.reviewDescriptionText {
  text-align: initial;
  font-size: 16px;
  margin: 10px;
}
.reviewWithTrash {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  align-self: center;
  align-content: center;
  margin-left: 15px;
}
</style>