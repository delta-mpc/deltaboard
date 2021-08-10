<template>
      <div class="container">
      <el-tabs v-model="activateName" class="login-panel" type="border-card">
        <el-tab-pane label="添加用户" name="register">
          <el-form ref="registForm" :model="registForm" 
          class="login-form" autocomplete="on" label-position="left">
            <el-form-item prop="username">
              <el-input ref="username" v-model="registForm.username" placeholder="请输入用户名" name="username" type="text" tabindex="4" autocomplete="on"/>
            </el-form-item>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
              <el-form-item prop="password">
                <el-input :key="registForm.passwordType" ref="registpassword" 
                v-model="registForm.password" :type="registForm.passwordType" 
                placeholder="请输入密码" name="password" tabindex="2" autocomplete="on"/>
              </el-form-item>
            </el-tooltip>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
              <el-form-item prop="repeatPassword">
                <el-input :key="registForm.repeatPasswordType" ref="repeatpassword" v-model="registForm.repeatPassword" :type="registForm.repeatPasswordType"
                placeholder="请再次输入密码" name="repeatpassword" tabindex="2" autocomplete="on"/>
              </el-form-item>
            </el-tooltip>
            <el-button :loading="loading" type="primary" style="width: 100%; margin-bottom: 30px" 
                  @click.native.prevent="register">添加用户</el-button>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'
import V1UserAPI from "@/api/v1/users"
export default {
  name: "asset",
  data() {
    return {
      loading: false,
      capsTooltip: false,
      activateName:'register',
      registForm: {
        password: "",
        repeatPassword: "",
        repeatPasswordType: "password",
        passwordType: "password",
        username:""
      },
    };
  },
  mounted() {
  },
  computed:{
     ...mapState({
        user:state => state.user
     }),
  },
  methods: {
     register() {
      this.loading = true;
         V1UserAPI.register(
            this.registForm.username,
            this.registForm.password
          ).then((response) => {
             this.$message(`已添加用户`)
          }).finally(() => {
            this.loading = false;
      });
    },
  },
  beforeRouteEnter(to, from, next) {
    store.commit('sidebar/SET_ADD_USR_PAGE')
    next()
  },
};
</script>

<style lang="stylus" scoped>
.container {
  padding-top 140px
}
.login-panel {
  width 450px
  margin 0 auto
  
  background white
  border 1px solid #DCDFE6
  box-shadow 0px 0px 17px 3px rgba(176, 193, 213, 0.66)
  border-radius border-radius

  /deep/.el-tabs__header {
    border none
    border-top-left-radius border-radius
    border-top-right-radius border-radius
    overflow hidden
  }

  /deep/.el-tabs__content {
    padding 60px
  }

  /deep/.el-tabs__nav {
    width 100%
    display flex
  }
  
  /deep/.el-tabs__item {
    text-align center
    flex 1
    height 50px
    line-height 50px
    border none
    background #dcdfe6
    color #606266
    font-size 20px
  }

  ::v-deep .el-input {
    input {
      color black
    }
    textarea {
      color black
    }
  }
  /deep/.el-input__inner {
    padding 0px
    border none
    border-bottom 1px solid #143654
    border-radius 0px
  }

  .el-form-item {
    margin-bottom 44px
  }
}

.login-form {
  position relative
  width 520px
  max-width 100%
  padding 10px 0 35px 0
  margin 0 auto
  overflow hidden
}

.reset-password {
  text-align right
}

.el-button {
  width 100%
  margin-bottom 20px
  font-size 18px  
}

a {
  text-decoration underline
  color #409EFF
}
</style>