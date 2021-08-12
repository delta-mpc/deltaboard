<template>
      <div class="container">
      <transition appear name="slide-fade">
         <el-tabs v-model="activateName" class="login-panel" type="border-card">
        <el-tab-pane label="修改密码" name="changepassword">
          <el-form class="change-form" ref="passForm" :label-position="labelPosition" :rules="passRules" :model="passForm" @submit.native.prevent>
            <el-form-item class="form-item" prop="oldPass">
               <el-input class="pass-input" placeholder="请输入旧密码" v-model="passForm.oldPass" show-password @keyup.enter.native="modifyLoginPass"></el-input>
            </el-form-item>
            <el-form-item class="form-item" prop="newPass">
               <el-input class="pass-input" placeholder="请输入新密码" v-model="passForm.newPass" show-password @keyup.enter.native="modifyLoginPass"></el-input>
            </el-form-item>
            <el-form-item class="form-item" prop="checkPass">
               <el-input class="pass-input" placeholder="请再次输入新密码" v-model="passForm.checkPass" show-password @keyup.enter.native="modifyLoginPass"></el-input>
            </el-form-item>
            <el-form-item class="form-button-item">
               <el-button class="change-button" type="primary" @click="modifyLoginPass">修改</el-button>
            </el-form-item>
        </el-form>
        </el-tab-pane>
         </el-tabs>
      </transition>
    </div>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'
import V1UserAPI from "@/api/v1/users"
export default {
  name: "asset",
  data() {
     const checkNewPass = (rule, value, callback) => {
      if (value.length < 5) {
        callback(new Error("新密码不能小于5位"));
      } else if (value !== this.passForm.newPass) {
        callback(new Error("两次输入密码不一致"));
      } else {
        callback();
      }
    };
    return {
      page: 1,
      totalCount: 0,
      loading: false,
      capsTooltip: false,
      activateName:'changepassword',
      passForm: {
        oldPass: "",
        newPass: "",
        checkPass: ""
      },
      passRules:{
         checkPass: { validator: checkNewPass, trigger: "blur"}
      }
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
     modifyLoginPass() {
      this.$refs.passForm.validate((valid) => {
        if (valid) {
          V1UserAPI.changeLoginPass(this.passForm.oldPass, this.passForm.newPass).then((response) => {
              this.$router.push({ name: 'login' })
          }).finally(() => {
            this.loading = false;
          })
        }
      });
    },
  },
  beforeRouteEnter(to, from, next) {
    store.commit('sidebar/SET_CHGPWD_PAGE')
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