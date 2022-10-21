<template>
<div>
  <content-component
    :title="'Побег из Шоушенка'"
    :type="'Фильм'"
    :genre="'Драма'"
    :likes="likes1"
    :liked="liked1"
    @likeBtnClick="liked1 = !liked1; likes1 += liked1*2 -1"
    :description="'Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. ' +
     'Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе '+
'стороны решётки. Каждый, кто попадает в эти стены, становится '+
'их рабом до конца жизни. Но Энди, обладающий живым умом '+
'и доброй душой, находит подход как к заключённым, '+
'так и к охранникам, добиваясь их особого к себе расположения.'"
  />
</div>
</template>

<script>
import ContentComponent from "@/components/ContentComponent";
export default {
  name: "MainView",
  components: {ContentComponent},
  data() {
    return {
      liked1: Boolean,
      likes1: 93,
    }
  }
}
</script>

<style scoped>

</style>