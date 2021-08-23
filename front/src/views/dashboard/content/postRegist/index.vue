<template>
  <div class="post-page">
     <Navbar :showLogin="false"/>
     <transition appear name="slide-fade">
     <div class="content-bg">
         <iframe style="width:100%;height:100%;border:none" :src="`${config.post_registry_url}`"></iframe>
     </div>
    </transition>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Navbar from '@/views/navbar.vue'
import UserModel from '@/model/user'
export default {
  name: "playground",
  components:{
     Navbar
  },
  mounted() {
     UserModel.getMyUserInfo()
  },
  computed:{
     ...mapState({
        user:state => state.user,
        config:(state)=>{return state.config.config}
     }),
  },
  methods: {
  }
};
</script>

<style lang="stylus" scoped>

.post-page {
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
  padding-top 20px
  height 100%
  background white
  border 1px #e4e4e4 solid
}


</style>