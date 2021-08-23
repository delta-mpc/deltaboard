<template>
<div class="section">
   <h2 class="title">{{title}}</h2>
   <slot></slot>
</div>
</template>
<script>
export default {
  name: "setting-section",
  props:{
     title:{
        type :String,
        default : ''
     }
  }
};
</script>

<style lang="stylus">
.section {
   width:100%
   padding-bottom:100px
}
.title {
   margin-top:-10px;
   margin-bottom:40px
}
</style>