<template>
  <div class="reviewWithTrash">
  <div class="reviewShell">
    <div class="reviewTitleShell">
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
    <ButtonComponent v-if="isUser"
                     :icon="'delete.svg'"
                     @btnClick="$emit('deleteReview')"
                     class="deleteBtn"/>
    </div>
    <div class="reviewDescription">
      <div class="reviewDescriptionText">
        {{review.text}}
      </div>
    </div>
  </div>

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
  justify-content: center;
}
.reviewTitleShell {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  width: 90vw;
  background: #94d1be;
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
}
.deleteBtn {
  background: #94d1be;
}
</style>