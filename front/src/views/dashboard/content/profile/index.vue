<template>
   <div class="container">
      <transition appear name="slide-fade">
         <div class="profile">
            <el-row>
               <el-col span="24">
                  <el-divider></el-divider>
               </el-col>
            </el-row>
            <el-row type="flex" gutter="15">
               <el-col span="10">
                  <setting-section title="个人信息">
                     <template>
                        <el-form label-position="left" label-width="70px">
                           <el-form-item label="用户名">
                              <strong>{{user.name}}</strong>
                           </el-form-item>
                           <el-form-item label="密码">
                              <el-button size="mini" class="clickable" @click="showChangePassword = true">修改密码
                              </el-button>
                           </el-form-item>
                        </el-form>
                     </template>
                  </setting-section>
               </el-col>
               <el-col span="2">
                  <div class="divider">
                     <el-divider direction="vertical"></el-divider>
                  </div>
               </el-col>
               <el-col span="10">
                  <div class="face">
                     <i class="el-icon-user"></i>
                  </div>
               </el-col>
            </el-row>
            <el-row>
               <el-col span="24">
                  <el-divider></el-divider>
               </el-col>
            </el-row>
            <el-row>
               <el-col>
                  <setting-section title="Deltaboard API">
                     <template>
                        <el-form>
                           <el-form-item label="API 地址">

                              <el-row gutter="15" type="flex" align="middle">
                                 <el-col span="12">
                                    <el-input :value="`${localUrl}/v1/tasks/${user.delta_token}}`" readonly
                                       disabled></el-input>
                                 </el-col>
                                 <el-col span="6">
                                    <el-button size="mini" title="复制" icon="el-icon-copy-document" @click="copyLink">
                                    </el-button>
                                    <el-button size="mini" title="刷新" icon="el-icon-refresh" @click="renewLink">
                                    </el-button>
                                 </el-col>
                              </el-row>
                           </el-form-item>
                        </el-form>
                     </template>
                  </setting-section>
               </el-col>
            </el-row>
         </div>
      </transition>
      <el-dialog :title="'修改密码'" :width="'500px'" :visible.sync="showChangePassword" :custom-class="'dialog'">
         <el-row>
            <el-col :span="12" :offset="6">
               <el-form class="change-pwd-form" ref="passForm" :model="passForm" :rules="passRules">
                  <el-form-item class="form-item" prop="oldPass">
                     <el-input class="pass-input" placeholder="请输入旧密码" v-model="passForm.oldPass" show-password
                        @keyup.enter.native="modifyLoginPass"></el-input>
                  </el-form-item>
                  <el-form-item class="form-item" prop="newPass">
                     <el-input class="pass-input" placeholder="请输入新密码" v-model="passForm.newPass" show-password
                        @keyup.enter.native="modifyLoginPass"></el-input>
                  </el-form-item>
                  <el-form-item class="form-item" prop="checkPass">
                     <el-input class="pass-input" placeholder="请再次输入新密码" v-model="passForm.checkPass" show-password
                        @keyup.enter.native="modifyLoginPass"></el-input>
                  </el-form-item>
                  <el-form-item class="form-button-item">
                     <el-button class="change-button" type="primary" @click="modifyLoginPass">修改</el-button>
                  </el-form-item>
               </el-form>
            </el-col>
         </el-row>
      </el-dialog>
   </div>
</template>

<script>
import store from "@/store";
import { mapState } from "vuex";
import V1UserAPI from "@/api/v1/users";
import UserModel from "@/model/user";
import SettingSection from "@/views/components/settingsection.vue";
export default {
   name: "asset",
   components: {
      "setting-section": SettingSection,
   },
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
         showChangePassword: false,
         activateName: "profile",
         passForm: {
            oldPass: "",
            newPass: "",
            checkPass: "",
         },
         passRules: {
            checkPass: { validator: checkNewPass, trigger: "manual" },
         },
      };
   },
   mounted() {
      console.log(this.user);
   },
   computed: {
      ...mapState({
         user: (state) => state.user,
      }),
      localUrl() {
         return window.BASE_API;
         // return 'https://localhost:8090'
         // return window.location.protocol + "//" + window.location.host;
      },
   },
   methods: {
      copyLink() {
         this.$copyText(
            `https://${this.localUrl}/v1/tasks/${this.user.delta_token}`
         ).then((res) => {
            this.$message("链接已复制");
         });
      },
      renewLink() {
         this.$confirm(
            `更新API地址后，旧的API地址将失效，原有的使用旧API地址的服务将无法正常运行，确定要更新吗？`
         ).then((res) => {
            V1UserAPI.renewDeltaToken(this.user.userId).then((res) => {
               UserModel.getMyUserInfo();
            });
         });
      },
      modifyLoginPass() {
         this.$refs.passForm.validate((valid) => {
            if (valid) {
               V1UserAPI.changeLoginPass(
                  this.passForm.oldPass,
                  this.passForm.newPass
               )
                  .then((response) => {
                     this.$router.push({ name: "login" });
                  })
                  .finally(() => {
                     this.loading = false;
                  });
            }
         });
      },
   },
   beforeRouteEnter(to, from, next) {
      store.commit("sidebar/SET_CHGPWD_PAGE");
      next();
   },
};
</script>

<style lang="stylus" scoped>
.divider {
   font-size: 200px;
}

.face {
   width: 150px;
   height: 150px;
   background: #f7fcfc;
   border-radius: 50%;
   display: flex;
   align-items: center;
   margin: 50px auto;
   font-size: 70px;

   i {
      margin: 0px auto;
   }
}

.change-pwd-form {
   margin: 20px auto;
}

.profile {
   width: 650px;
   margin-left: 50px;

   .form {
      width: 650px;
      margin: 20px auto;
   }
}

.container {
   padding-top: 40px;
}

.profile-item {
   width: 100%;
   display: flex;
}

.reset-password {
   text-align: right;
}

a {
   text-decoration: underline;
   color: #409EFF;
}
</style>