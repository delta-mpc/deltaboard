<template>
  <div class="index-div">
    <keep-alive>
      <router-view v-if="$route.meta.keepAlive" />
    </keep-alive>
    <router-view v-if="!$route.meta.keepAlive"></router-view>
  </div>
</template>

<script>
export default {
  name: "index",
  data() {
    return {};
  },
  methods: {},
};
</script>

<style lang="stylus" scoped>

.index-div {
  height 100%
}

</style>