<template>
  <div class="header">
    <router-link :to="{ name: 'asset' }">
      <div class="logo">Delta - Dashboard</div>
    </router-link>
    <div class="right-menu">
      <router-link :to="{ name: 'login' }" v-if="showLogin">登录/注册</router-link>
      <el-dropdown class="dropdown-class" @command="handleCommand" v-else>
        <span class="user">{{ navUserInfo }}<i class="el-icon-arrow-down el-icon--right"></i></span>
        <el-dropdown-menu class="dropdown-class" slot="dropdown">
          <el-dropdown-item command="change">个人中心</el-dropdown-item>
          <el-dropdown-item command="logout">退出账户</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>

import UserAPI from '@/api/v1/users'
import { mapState } from 'vuex'
import ErrorMessage from '@/model/errorMessage'
import { Message } from "element-ui"
import axios from 'axios';
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
      iframe.src = 'https://192.168.1.46:8000/hub/logout'
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

.header {
  height 60px
  font-size 22px
  background theme-color

  display flex
  justify-content space-between
  align-items center
  flex-shrink 0

  .logo {
    text-align center
    font-size 20px
    color white
    padding-left 56px
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
    color white
    margin 0px 10px
    text-decoration none
  }

  .menu-link {
    margin-right 30px
  }

}

.el-dropdown-menu__item.is-disabled  {
  color #606266
  background #dde6ed
}

.dropdown-class {
  padding-top 0px
  padding-left 0px
  padding-bottom 0px
}

</style>