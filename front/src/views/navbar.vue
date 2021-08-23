<template>
  <div class="header">
    <router-link :to="{ name: 'playground' }">
      <img class="logo" :src="require('@/assets/logo.svg')"/> 
    </router-link>
    <div class="right-menu">
      <router-link :to="{ name: 'login' }" v-if="showLogin">登录/注册</router-link>
      <el-menu v-else :default-active="1" class="menu" mode="horizontal" @select="handleSelect">
         <el-submenu index="1">
            <template slot="title"><span class="user">{{ navUserInfo }}</span></template>
               <el-menu-item index="profile">
                  <font-awesome-icon size="lg" icon="user-alt" style="margin-left:30px"></font-awesome-icon>
                  <span style="margin-left:20px">个人中心</span>
               </el-menu-item>
               <div class="divider"></div>
               <el-menu-item index="logout">
                  <font-awesome-icon size="lg" icon="sign-out-alt" style="margin-left:30px"></font-awesome-icon>
                  <span style="margin-left:20px">退出账户</span>
               </el-menu-item>
         </el-submenu>
      </el-menu>
    </div>
  </div>
</template>

<script>

import UserAPI from '@/api/v1/users'
import { mapState } from 'vuex'
import ErrorMessage from '@/model/errorMessage'
import { Message } from "element-ui"
export default {
  name: "",
  props: {
    showLogin: {
      type: Boolean,
      default: () => {
        return true
      }
    },
  },
  data() {
    return {};
  },
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
    handleCommand(command) {
      if (command === 'change') {
        this.change()
      }
      else if (command === 'logout') {
        this.logout()
      }
      else if (command === 'digitalID') {
        this.digitalID()
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
    digitalID() {
      if (this.name.length === 0 || this.company.name.length === 0) {
        this.$store.commit('dashboard/SET_SHOW_AUTH')
        return;
      }
      this.$store.commit('setting/SET_DIGITAL_ACTIVE')
      this.$router.push({ name: 'setting' })
    },
    logout() {
      if (!confirm('是否确认退出？')) return;
      let iframe = document.createElement('iframe');
      iframe.src = window.location.protocol + '//' + window.location.host + '/hub/logout'
      iframe.style.width = '0px'
      iframe.style.height = '0px'
      document.body.appendChild(iframe)
      setTimeout(()=>{
         iframe.remove()
         UserAPI.logout().then(() => {
        window.confidentialNotified = false
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
  background page-bg-color
  .logo {
    height 40px
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