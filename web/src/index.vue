<template>
  <div class="index-div">
    <Navbar/>
    <keep-alive>
      <router-view v-if="$route.meta.keepAlive" />
    </keep-alive>
      <router-view v-if="!$route.meta.keepAlive"></router-view>
  </div>
</template>

<script>
import Navbar from '@/components/navbar.vue'
export default {
  name: "index",
  data() {
    return {};
  },
  components:{
     Navbar
  },
  methods: {},
};
</script>

<style lang="stylus" scoped>

.index-div {
  height 100%
  display flex
  flex-direction column
}
</style>