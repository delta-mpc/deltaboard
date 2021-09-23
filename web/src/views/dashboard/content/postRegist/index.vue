<template>
   <div class="post-page">
      <transition appear name="slide-fade">
         <div class="content-bg">
            <iframe style="width:100%;height:100%;border:none" :src="`${config.post_registry_url}#${language}`"></iframe>
         </div>
      </transition>
   </div>
</template>

<script>
import { mapState } from "vuex";
import UserModel from "@/model/user";
export default {
   name: "postRegist",
   mounted() {
      UserModel.getMyUserInfo();
   },
   computed: {
      ...mapState({
         user: (state) => state.user,
         config: (state) => {
            return state.config.config;
         },
      }),
      language() {
         var language = (
            navigator.browserLanguage || navigator.language
         ).toLowerCase();
         var localeLang = "";
         if (language.indexOf("zh") > -1) {
            localeLang = "zh";
         } else {
            localeLang = "en";
         }
         return localeLang
      },
   },
   methods: {},
};
</script>

<style lang="stylus" scoped>
.post-page {
   background: page-bg-color;
   height: 100%;
   -webkit-overflow-scrolling: touch;
}

.upload-button {
   color: #607185;
   background: #dfe5ec;
   border: 1px solid #607185;
}

.content-bg {
   padding-top: 20px;
   height: 100%;
   background: white;
   border: 1px #e4e4e4 solid;
}
</style>