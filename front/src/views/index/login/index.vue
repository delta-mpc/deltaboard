<template>
  <div class="login-page">
    <Navbar />
    <div class="container">
      <el-tabs v-model="activateName" class="login-panel" type="border-card">
        <el-tab-pane label="账户登录" name="login">
          <el-form ref="loginForm" :model="loginForm" :rules="loginRules" 
          class="login-form" autocomplete="on" label-position="left" @submit.native.prevent>
            <el-form-item prop="username">
              <el-input ref="email" v-model="loginForm.username" placeholder="请输入用户名" 
              name="phonenumber" type="text" tabindex="1" autocomplete="on" @keyup.enter.native="login"/>
            </el-form-item>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
              <el-form-item prop="password">
                <el-input :key="loginForm.passwordType" ref="loginpassword" v-model="loginForm.password" 
                :type="loginForm.passwordType" placeholder="请输入登录密码" name="password" 
                tabindex="2" autocomplete="on" @keyup.enter.native="login"/>
              </el-form-item>
            </el-tooltip>
            <el-button :loading="loading" type="primary" @click="login">登录</el-button>
            <div class="reset-password">
              <a href="/reset-pass" class="medium-secondary-text">忘记密码？</a>
            </div>
          </el-form>
        </el-tab-pane>
        <el-tab-pane v-if="false" label="账户注册" name="register">
          <el-form ref="registForm" :model="registForm" :rules="loginRules" 
          class="login-form" autocomplete="on" label-position="left">
           <el-form-item prop="email">
              <el-input ref="email" v-model="registForm.email" placeholder="请输入邮箱" name="email" type="text" tabindex="4" autocomplete="on"/>
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
            @click.native.prevent="register">注册用户</el-button>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>


<script>
import V1UserAPI from "@/api/v1/users"
import UserModel from '@/model/user'
import Navbar from '../../navbar.vue'
export default {
  name: "login",
  components: {
    Navbar
  },
  data() {
    const self = this;
    const validatePassword = (rule, value, callback) => {
      if (value.length < 5) {
        callback(new Error("密码不能小于5位"));
      } else {
        callback();
      }
    };
    const validateRepeatPassword = (rule, value, callback) => {
      if (value !== self.registForm.password) {
        callback(new Error("两次密码输入不一致"));
      } else {
        callback();
      }
    };
    const validateCode = (rule, value, callback) => {
      if (!value) {
        callback(new Error("请输入验证码"));
      } else {
        callback();
      }
    };
    const validateEmail = (rule, value, callback) => {
      if (!/^[\w\d_-]+@[\w\d_-]+(\.[\w\d_-]+)+$/.test(value)) {
        callback(new Error("请输入正确的邮箱"));
      } else {
        callback();
      }
    };
    return {
      loginForm: {
        password: "",
        passwordType: "password",
        phonenumber: "",
        username:""
      },
      activateName: "login",
      registForm: {
        password: "",
        repeatPassword: "",
        repeatPasswordType: "password",
        passwordType: "password",
        noticed: false,
        validation_code:"",
        phonenumber:""
      },
      loginRules: {
        password: [ { required: true, trigger: "blur", validator: validatePassword },],
        repeatPassword: [ { required: true, trigger: "blur", validator: validateRepeatPassword, }, ],
        validation_code: [ { required: true, trigger: "blur", validator: validateCode }, ],
        email:[ {required: true, trigger: "blur",validator:validateEmail} ]
      },
      capsTooltip: false,
      loading: false,
      redirect: undefined,
      otherQuery: {},
    };
  },
  methods: {
    showInformation() {
      this.$router.push({ name: 'information' })
    },
    register() {
      this.$refs.registForm.validate((valid) => {
        if (valid) {
          this.loading = true;
          V1UserAPI.register(
            this.registForm.email,
            this.registForm.password
          ).then((response) => {
            this.$store.commit("user/SET_USER", response);
            this.$router.push({ name: "playground" });
          }).finally(() => {
            this.loading = false;
          });
        }
      });
    },
    login() {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.loading = true;
          V1UserAPI.login(this.loginForm.username, this.loginForm.password).then((response) => {
            this.$store.commit("user/SET_USER", response);
            this.$store.commit("sidebar/SET_PLAYGROUND_PAGE");
            this.$router.push({ name: "playground" });
            localStorage.setItem('visibilitychange', 'changed')
          }).finally(() => {
            this.loading = false;
          });
        }
      });
    },
    logout() {
      window.confidentialNotified = false
      this.$router.push({ name: "login" });
    },
  },
  beforeRouteEnter(to, from, next) {
    UserModel.getMyUserInfo().then((user) => {
      if (from.name) {
        if (from.name === 'rootIndex' || from.name === 'verifyFile') {
          next({ name: 'playground' })
        }
        else {
          next({ name: from.name })
        }
      }
      else {
        next({ name: 'playground' })
      }
    }).catch((error) => {
      next()
    })
  },
};
</script>

<style lang="stylus" scoped>
/deep/.el-input-group__append {
   background transparent
   border none
}
.login-page {
  position relative
  width 100%
  height 100%
  
  background url("../login/image/loginin-bg.png") center bottom no-repeat
}

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