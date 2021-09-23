<template>
  <div class="header">
    <router-link :to="{ name: 'playground' }">
      <img class="logo" :src="require('@/assets/logo.svg')"/>
    </router-link>
    <div class="right-menu">
      <el-menu
          default-active="1"
          class="menu"
          mode="horizontal"
          @select="handleSelect">
         <el-submenu index="1" v-if="!showLogin">
            <template slot="title">
                <span class="user">{{ navUserInfo }}</span>
            </template>
           <el-menu-item index="profile">
              <font-awesome-icon size="lg" icon="user-alt" style="margin-left:30px"></font-awesome-icon>
              <span style="margin-left:20px">{{$t('dashboard.navbar.personal_profile')}}</span>
           </el-menu-item>
           <div class="divider"></div>
           <el-menu-item index="logout">
              <font-awesome-icon size="lg" icon="sign-out-alt" style="margin-left:30px"></font-awesome-icon>
              <span style="margin-left:20px">{{$t('dashboard.navbar.logout')}}</span>
           </el-menu-item>
         </el-submenu>
         <el-menu-item index="1" v-else>
            {{$t('dashboard.navbar.login/regist')}}
         </el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<script>

import UserAPI from '@/api/v1/users'
import { mapState } from 'vuex'
import ErrorMessage from '@/model/errorMessage'
import { Message } from "element-ui"
import freezable from '@/mixins/freezable'
export default {
  name: "",
  props: {
  },
  data() {
    return {};
  },
  mixins:[freezable],
  computed: {
    ...mapState({
      phoneNum: state => state.user.phonenumber,
      name: state => state.user.name,
      user: state => state.user
    }),
    ...mapState(['company']),
    navUserInfo() {
      return this.name ? this.name : this.hiddenPhone
    },
    hiddenName() {
      let name = this.name
      if(null != name && name != undefined){
        if(name.length==2){
          return '*' + name.substring(1,2) //截取name 字符串截取第一个字符，
        } else if(name.length==3){
          return "*"+name.substring(1,3)//截取第一个和第三个字符
        } else if(name.length>3){
          return "**"+name.substring(2,name.length)//截取第一个和大于第4个字符
        }
      }
      return null;
    },
    hiddenPhone() {
      var reg = /^(\d{3})\d{4}(\d{4})$/
      return this.phoneNum.replace(reg, "$1****$2");
    },
    localUrl(){
        return window.BASE_API
    },
    showLogin(){
       return this.user.name == null || this.user.name.length == 0
    }
  },
  methods: {
    handleSelect(key,keyPath) {
       switch (key) {
          case 'logout':
             this.logout()
             break;
          case 'profile':
             this.$router.push({name:'profile'})
             break;
       }
    },
    change() {
      if (this.name.length === 0 || this.company.name.length === 0) {
        this.$store.commit('dashboard/SET_SHOW_AUTH')
        return;
      }
      this.$store.commit('setting/SET_CHANGEPASS_ACTIVE')
      this.$router.push({ name: 'setting' })
    },
    logout() {
      if (!confirm(this.$t('dashboard.navbar.confirm_quit'))) return;
      let iframe = document.createElement('iframe');
      iframe.src = this.localUrl + '/hub/logout'
      iframe.style.width = '0px'
      iframe.style.height = '0px'
      this.freezeMessage = `${this.$t('dashboard.navbar.quiting')}...`
      this.freezed = true
      document.body.appendChild(iframe)
      setTimeout(()=>{
         iframe.remove()
         this.freezed = false
         UserAPI.logout().then(() => {
         this.$store.commit('user/CLEAR_USER')   
         this.$router.push({ name: "login" });
      }).catch((error) => {
        this.$errorMessage(error,(error)=>{
           let errorMessage = error.response.data.message
           let showMessage = ErrorMessage.getLocalMessage(errorMessage)
           if (errorMessage === 'user login status expired') {
               this.$router.push({ name: 'login' });
               return
           }
           if (showMessage) {
               Message({ message: showMessage, type: 'error', duration: 3 * 1000})
           }
        })
        this.$router.push({ name: "login" });
      })
      },3000)
    },
  },
};
</script>

<style lang="stylus" scoped>
.divider {
   margin 10px 30px
   height 1px
   background black
   opacity 0.1
}
.menu {
   background transparent
   border none
}
.header {
  height 60px
  font-size 22px
  display flex
  justify-content space-between
  align-items center
  flex-shrink 0
  background-color white
  border-bottom 1px solid #EBEEF5
  .logo {
    height 28px
    font-size 20px
    line-height 50px
    margin-left 30px
    color black
  }

  .right-menu {
    padding-right 93px
    .el-dropdown {
      cursor pointer
    }
  }

  a,
  .user {
    font-size 16px
    margin 0px 10px
    text-decoration none
  }

  .menu-link {
    margin-right 30px
  }
}
</style>
<style lang="stylus">
.el-header {

}
</style>
