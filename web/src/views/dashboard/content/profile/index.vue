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
                  <setting-section :title="$t('dashboard.profile.personal_information')">
                     <template>
                        <el-form label-position="left" label-width="70px">
                           <el-form-item :label="$t('dashboard.profile.name')">
                              <strong>{{user.name}}</strong>
                           </el-form-item>
                           <el-form-item :label="$t('common.password')">
                              <el-button size="mini" class="clickable" @click="showChangePassword = true">
                                 {{$t('dashboard.profile.change_pwd')}}
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
                           <el-form-item :label="$t('dashboard.profile.dashboard_api')">

                              <el-row gutter="15" type="flex" align="middle">
                                 <el-col span="12">
                                    <el-input :value="`${localUrl}/submit/${user.delta_token}`" readonly
                                       disabled></el-input>
                                 </el-col>
                                 <el-col span="6">
                                    <el-button size="mini" :title="$t('dashboard.profile.copy')" icon="el-icon-copy-document" @click="copyLink">
                                    </el-button>
                                    <el-button size="mini" :title="$t('dashboard.profile.renew')" icon="el-icon-refresh" @click="renewLink">
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
      <el-dialog :title="$t('dashboard.profile.change_pwd')" :width="'500px'" :visible.sync="showChangePassword" :custom-class="'dialog'">
         <el-row>
            <el-col :span="12" :offset="6">
               <el-form class="change-pwd-form" ref="passForm" :model="passForm" :rules="passRules">
                  <el-form-item class="form-item" prop="oldPass">
                     <el-input class="pass-input" :placeholder="$t('dashboard.profile.input_old_password')" v-model="passForm.oldPass" show-password
                        @keyup.enter.native="modifyLoginPass"></el-input>
                  </el-form-item>
                  <el-form-item class="form-item" prop="newPass">
                     <el-input class="pass-input" :placeholder="$t('dashboard.profile.input_new_password')" v-model="passForm.newPass" show-password
                        @keyup.enter.native="modifyLoginPass"></el-input>
                  </el-form-item>
                  <el-form-item class="form-item" prop="checkPass">
                     <el-input class="pass-input" :placeholder="$t('dashboard.profile.repeat_new_password')" v-model="passForm.checkPass" show-password
                        @keyup.enter.native="modifyLoginPass"></el-input>
                  </el-form-item>
                  <el-form-item class="form-button-item">
                     <el-button class="change-button" type="primary" @click="modifyLoginPass">{{$t('dashboard.profile.change')}}</el-button>
                  </el-form-item>
               </el-form>
            </el-col>
         </el-row>
      </el-dialog>
   </div>
</template>

<script>
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
            callback(new Error(this.$t('common.password_no_less_than_5_digits')));
         } else if (value !== this.passForm.newPass) {
            callback(new Error(this.$t('common.password_mismatch')));
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
      },
   },
   methods: {
      copyLink() {
         this.$copyText(
            `${this.localUrl}/submit/${this.user.delta_token}`
         ).then((res) => {
            this.$message(this.$t('dashboard.profile.address_copied'));
         });
      },
      renewLink() {
         this.$confirm(
            this.$t('dashboard.profile.confirm_renew_api')
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
               ).then((response) => {
                     this.$router.push({ name: "login" });
                  })
                  .finally(() => {
                     this.loading = false;
                  });
            }
         });
      },
   }
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
