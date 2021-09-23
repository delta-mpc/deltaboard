<template>
  <div class="login-page">
    <div class="container">
      <transition appear name="slide-fade">
       <el-tabs v-model="activateName" type="border-card" class="login-panel">
        <el-tab-pane :label="$t('common.account_login')" name="login">
            <el-form ref="loginForm" :model="loginForm" :rules="loginRules" 
          class="login-form" autocomplete="on" label-position="left" @submit.native.prevent>
            <el-form-item prop="username">
              <el-input ref="email" value="" v-model="loginForm.username" :placeholder="$t('common.please_input_your_name')" 
              name="phonenumber" type="text" tabindex="1" autocomplete="on" @keyup.enter.native="login"/>
            </el-form-item>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
              <el-form-item prop="password">
                <el-input :key="loginForm.passwordType" ref="loginpassword" v-model="loginForm.password" 
                :type="loginForm.passwordType" :placeholder="$t('common.please_input_your_pwd')" name="password" 
                tabindex="2" autocomplete="off" @keyup.enter.native="login"/>
              </el-form-item>
            </el-tooltip>
            <el-button v-show="showLogin" :loading="loading" type="primary" @click="login">{{$t('common.login')}}</el-button>
            </el-form>
        </el-tab-pane>
        <el-tab-pane v-if="config.public_registration" :label="$t('common.regist_account')" name="register">
            <el-form ref="registForm" :model="registForm" :rules="loginRules" 
          class="login-form" autocomplete="on" label-position="left">
           <el-form-item prop="username">
              <el-input ref="username" v-model="registForm.username" :placeholder="$t('common.please_input_your_name')" name="username" type="text" tabindex="4" autocomplete="on"/>
            </el-form-item>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
              <el-form-item prop="password">
                <el-input :key="registForm.passwordType" ref="registpassword" 
                v-model="registForm.password" :type="registForm.passwordType" 
                :placeholder="$t('common.please_input_your_pwd')" name="password" tabindex="2" autocomplete="on"/>
              </el-form-item>
            </el-tooltip>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
              <el-form-item prop="repeatPassword">
                <el-input :key="registForm.repeatPasswordType" ref="repeatpassword" v-model="registForm.repeatPassword" :type="registForm.repeatPasswordType"
                :placeholder="$t('common.please_repeat_your_pwd')" name="repeatpassword" tabindex="2" autocomplete="on"/>
              </el-form-item>
            </el-tooltip>
            <el-button :loading="loading" type="primary" @click.native.prevent="register">{{$t('common.regist')}}</el-button>
            </el-form>
        </el-tab-pane>
         </el-tabs>
      </transition>
    </div>
  </div>
</template>


<script>
import V1UserAPI from "@/api/v1/users"
import UserModel from '@/model/user'
import {mapState} from 'vuex'
export default {
  name: "login",
  computed:{
     ...mapState({config:(state)=> state.config.config})
  },
  data() {
    const self = this;
    const validatePassword = (rule, value, callback) => {
      if (value.length < 5) {
        callback(new Error(this.$t('common.password_no_less_than_5_digits')));
      } else {
        callback();
      }
    };
    const validateRepeatPassword = (rule, value, callback) => {
      if (value !== self.registForm.password) {
        callback(new Error(this.$t('common.password_mismatch')));
      } else {
        callback();
      }
    };
    const validateCode = (rule, value, callback) => {
      if (!value) {
        callback(new Error(this.$t('common.please_input_validation_code')));
      } else {
        callback();
      }
    };
    const validateEmail = (rule, value, callback) => {
      if (!/^[\w\d_-]+@[\w\d_-]+(\.[\w\d_-]+)+$/.test(value)) {
        callback(new Error(this.$t('common.please_input_the_correct_email_addr')));
      } else {
        callback();
      }
    };
    return {
      showLogin:true,
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
            this.registForm.username,
            this.registForm.password
          ).then((response) => {
            this.$router.push({ name: "post-regist",query:{user:this.registForm.username}});
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
            if(response.approve_status == this.$appGlobal.constants.USER_APPROVE_STATUS_REGISTED) {
                this.$router.push({ name: "post-regist" });
            } else {
               this.$store.commit("user/SET_USER", response);
               this.$router.push({ name: "playground" });
            }
          }).finally(() => {
            this.loading = false;
          });
        }
      });
    },
    logout() {
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
      if(error == 'user not approved') {
         next({name:'post-regist'})
      } else {
         next()
      }
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
}

.container {
  padding-top 140px
}

.login-panel {
  width 450px
  margin 0 auto
  border-radius border-radius
}

.reset-password {
  text-align right
}

a {
  text-decoration underline
  color #409EFF
}

</style>