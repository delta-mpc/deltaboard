<template>
  <div class="app-wrapper">
    <el-container class="dashboard">
      <el-header>
        <Navbar :showLogin="false"/>
      </el-header>
      <el-container class="main-container">
        <el-aside width="263px">
          <SideBar/>
          <div class="side-blank"></div>
        </el-aside>
        <el-main>
         <PlayGround :style="{visibility:$route.name == 'playground' ? 'visible':'hidden',overflow:'hidden',position:'absolute',width: 'calc(100vw - 263px)',height: 'calc(100vh - 60px)'}"/>
         <router-view v-if="$route.name != 'playground'"/>
        </el-main>
      </el-container> 
      <div class="side-shadow"></div>
    </el-container>
  </div>
</template>

<script>
import SideBar from "./sidebar"
import { mapState } from 'vuex'
import UserModel from '@/model/user'
import PlayGround from '@/views/dashboard/content/playground/index.vue'
import Navbar from '@/views/navbar'
export default {
  name: "dashboard",
  components: {
    SideBar,
    Navbar,
    PlayGround
  },
  data() {
    const validateName = (rule, value, callback) => {
      if (value.length === 0) {
        callback(new Error('请输入真实姓名'))
      }
      else {
        callback()
      }
    };
    
    const validateIDNum = (rule, value, callback) => {
      var str = value.trim()
			var regex = /^\d{6}(18|19|20)?\d{2}(0[1-9]|1[12])(0[1-9]|[12]\d|3[01])\d{3}(\d|X)$/
			if(!regex.test(str)){
				callback(new Error("请输入正确的身份证号码!"))
			}
      else {
        callback()
      }
    };

    const self = this;
    const validateCreateCompanyName = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入正确的企业名称"));
      } else {
        callback();
      }
    };
    const validateLicense = (rule, value, callback) => {
      var pattern = /(^(?:(?![IOZSV])[\dA-Z]){2}\d{6}(?:(?![IOZSV])[\dA-Z]){10}$)|(^\d{15}$)/
      if (!pattern.test(value)) {
        callback(new Error("请输入正确的营业执照"));
      } else {
        callback();
      }
    };
    const validateJoinCompanyName = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入正确的企业名称"));
      } else {
        callback();
      }
    };
    return {
      loading: false,
      authStep: 1,
      real_auth_disable: true,
      authForm: {
        name: '',
        idNum: '',
      },
      authRules: {
        name: { required: true, trigger: 'all', validator: validateName},
        idNum: { required: true, trigger: "all", validator: validateIDNum },
      },
      showCreate: false,
      showJoin: false,
      showCreateOrBind: true,

      labelPosition: "top",
      createForm: {
        companyName: "",
        license: "",
      },
      createRules: {
        companyName: [
          { required: true, trigger: "blur", validator: validateCreateCompanyName },
        ],
        license: [
          { required: true, trigger: "blur", validator: validateLicense },
        ],
      },

      joinForm: {
        companyName: "",
      },
      joinRules: {
        companyName: [
          { required: true, trigger: "blur", validator: validateJoinCompanyName },
        ],
      }
     }
  },
  watch: {
    authForm: {
      handler: function (newValue) {
        this.$refs.authForm.validate((valid) => {
        this.real_auth_disable = !valid
      })
      },
      deep: true,
    },
  },
  computed: {
    ...mapState(['dashboard']),
    ...mapState(['user']),
    ...mapState(['company']),
  },
  mounted() {
    if (this.user.card_no.length === 0) {
      this.authStep = 1
    }
    else {
      this.authStep = 2
    }
    window.addEventListener('visibilitychange', function() {
      if (document.visibilityState !== 'hidden') {
      let changed = localStorage.getItem('visibilitychange')
      if (changed === 'changed') {
        window.location.reload()
      }
      localStorage.removeItem('visibilitychange')
      } 
    })
  },
  methods: {
  },
  beforeRouteEnter(to, from, next) {
    UserModel.getMyUserInfo().then((user) => {
        if(user.approve_status != 2) {
           next({name:'post-regist'})
        }
        next()
    }).catch((err) => {
      next({ name: 'login' })
    })
  },
};
</script>

<style lang="stylus" scoped>

.app-wrapper {
  height: 100%;
}

.dashboard {
  height: 100%;
}

.el-header {
  padding: 0;
  height: 60px;
  border-bottom 1px solid #f5f8fa
}

.el-main {
  padding 0px
  margin-bottom 2px
  overflow hidden
}
.userauth-dilaog {
   /deep/&.el-dialog__wrapper {
      background white
   }
  /deep/.el-dialog__body {
    display flex
    justify-content center
  }
  .log-content {
    width 520px
  }
  .el-steps {
    margin-top 48px
  }
  .el-form {
    width 386px
    margin 24px auto 
  }
  /deep/.el-dialog__header {
    height 32px
  }
  .form-footer {
    text-align center
  }
  .form-item-btn {
    text-align center
  }
  .first-step {
    margin-bottom 105px
  }
    
  .second-step,
  .third-step {
    margin 30px auto
    margin-bottom 105px
    display flex
    justify-content center
  }

  .second-step-select {
    display flex
    justify-content center
  }
  
  .bind-operation {
    display flex
    flex-direction column
    width 116px
    height 96px
    justify-content center
    align-items center
    background #E1E7EE
    border 2px solid theme-color
    border-radius 5px
  }

  /deep/.el-icon-connection,
  /deep/.el-icon-office-building {
    font-size 40px
  }

  /deep/.el-form-item {
    margin-bottom 28px
  }
  .back-button {
    border none
  }
}

</style>
