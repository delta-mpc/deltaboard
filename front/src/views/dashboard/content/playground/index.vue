<template>
  <div class="playground-page">
     <transition appear name="slide-fade">
    <div class="content-bg">
       <iframe style="width:100%;height:100%;border:none" :src="`${localUrl}/hub/external/login?next=/hub/&username=${user.name}&token=${user.user_token}`"></iframe>
    </div>
    </transition>
  </div>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'

export default {
  name: "playground",
  mounted() {
  },
  computed:{
     ...mapState({
        user:state => state.user
     }),
     localUrl(){
        return "https://api.board.deltampc.com"
         // return window.location.protocol + "//" + window.location.host;
     }
  },
  methods: {
  },
  beforeRouteEnter(to, from, next) {
    store.commit('sidebar/SET_PLAYGROUND_PAGE')
    next()
  },
};
</script>

<style lang="stylus" scoped>

.playground-page {
  background page-bg-color
  height 100%
  -webkit-overflow-scrolling touch
}

.upload-button {
  color #607185
  background #dfe5ec
  border 1px solid #607185
}

.content-bg {
  height 100%
  background white
  border 1px #e4e4e4 solid
}


</style>