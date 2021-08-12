<template>
      <div class="container">
      <div class="btn-ctn">
         <el-button @click="showAddUser = true">添加用户</el-button>   
      </div>
      <transition appear name="slide-fade">
      <div class="table-ctn">
         <el-table :data="tableData">
            <el-table-column label="用户名" prop="name"></el-table-column>
            <el-table-column prop="created_at" label="日期">{{  }}
              <template v-slot="{ row }">{{ row.created_at | second2Date }}</template>
            </el-table-column>
         </el-table>   
      </div>   
      </transition>
      <el-dialog :visible.sync="showAddUser" :title="'添加用户'">
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
      </el-dialog>

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
      tableData:[],
      loading: false,
      capsTooltip: false,
      showAddUser:false,
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
     this.load()
  },
  computed:{
     ...mapState({
        user:state => state.user
     }),
  },
  methods: {
     load(){
        V1UserAPI.list(0,999).then((res)=>{
            this.tableData = res.list
        })
     },
     register() {
      this.loading = true;
         V1UserAPI.register(
            this.registForm.username,
            this.registForm.password
          ).then((response) => {
             this.$message(`已添加用户`)
             this.showAddUser = false
             this.load()
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
  padding-top 40px
  padding-left 20px
  .btn-ctn {
     width 150px
     height 50px
     padding-left 20px
  }
  .table-ctn {
     padding 20px
  }
}
/deep/.el-dialog {
   width 600px
}

.login-form {
  position relative
  width 320px
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