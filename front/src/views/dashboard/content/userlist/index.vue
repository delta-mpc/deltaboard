<template>
   <div class="container">
      <transition appear name="slide-fade">
         <el-tabs v-model="activeName">
            <el-tab-pane label="管理" name="management">
               <div class="btn-ctn">
                  <el-button size="medium" type="primary" @click="showAddUser = true">添加用户</el-button>
               </div>
               <div class="table-ctn">
                  <el-table :data="management.tableData">
                     <el-table-column label="用户名" prop="name"></el-table-column>
                     <el-table-column prop="created_at" label="日期">
                        <template v-slot="{ row }">{{ row.created_at | second2Date }}</template>
                     </el-table-column>
                     <el-table-column label="操作" prop="name">
                        <template v-slot="{ row }">
                           <el-tag class="clickable" type="danger" @click="delUser(row)">删除</el-tag>
                        </template>
                     </el-table-column>
                  </el-table>
               </div>
            </el-tab-pane>
            <el-tab-pane v-if="config.public_registration" label="审核" name="approval">
               <div class="table-ctn">
                  <el-table :data="approval.tableData">
                     <el-table-column label="用户名" prop="name"></el-table-column>
                     <el-table-column prop="created_at" label="日期">
                        <template v-slot="{ row }">{{ row.created_at | second2Date }}</template>
                     </el-table-column>
                     <el-table-column label="操作" prop="name">
                        <template v-slot="{ row }">
                           <el-row>
                              <el-col span="4">
                                 <el-tag class="clickable" type="success" @click="approveUser(row)">通过</el-tag>
                              </el-col>
                              <el-col span="4">
                                 <el-tag class="clickable" type="danger" @click="rejectUser(row)">拒绝</el-tag>
                              </el-col>
                           </el-row>
                        </template>
                     </el-table-column>
                  </el-table>
               </div>
            </el-tab-pane>
         </el-tabs>
      </transition>
      <el-dialog :visible.sync="showAddUser" :title="'添加用户'" width="500px">
         <el-form ref="registForm" :model="registForm" class="login-form" autocomplete="on" label-position="left">
            <el-form-item prop="username">
               <el-input ref="username" v-model="registForm.username" placeholder="请输入用户名" name="username" type="text"
                  tabindex="4" autocomplete="on" />
            </el-form-item>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
               <el-form-item prop="password">
                  <el-input :key="registForm.passwordType" ref="registpassword" v-model="registForm.password"
                     :type="registForm.passwordType" placeholder="请输入密码" name="password" tabindex="2"
                     autocomplete="on" />
               </el-form-item>
            </el-tooltip>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
               <el-form-item prop="repeatPassword">
                  <el-input :key="registForm.repeatPasswordType" ref="repeatpassword"
                     v-model="registForm.repeatPassword" :type="registForm.repeatPasswordType" placeholder="请再次输入密码"
                     name="repeatpassword" tabindex="2" autocomplete="on" />
               </el-form-item>
            </el-tooltip>
            <el-button :loading="loading" type="primary"
               @click.native.prevent="register">添加用户</el-button>
         </el-form>
      </el-dialog>

   </div>
</template>

<script>
import store from "@/store";
import { mapState } from "vuex";
import V1UserAPI from "@/api/v1/users";
export default {
   name: "asset",
   data() {
      return {
         management:{
            tableData:[]
         },
         approval:{
            tableData:[]
         },
         capsTooltip: false,
         loading: false,
         showAddUser: false,
         activeName: "management",
         registForm: {
            password: "",
            repeatPassword: "",
            repeatPasswordType: "password",
            passwordType: "password",
            username: "",
         },
      };
   },
   watch:{
      'activeName':'load'
   },
   mounted() {
      this.load();
   },
   computed: {
      ...mapState({
         user: (state) => state.user,
         config:(state)=> state.config.config
      })
   },
   methods: {
      delUser(user) {
         this.$confirm(`确认删除用户${user.name} ? `).then((res) => {
            V1UserAPI.delUser(user.id).then((res) => {
               this.$message(`用户已删除`);
               this.load();
            });
         });
      },
      approveUser(user){
         this.$confirm(`确认通过用户${user.name}申请 ? `).then((res) => {
            V1UserAPI.approveUser(user.id).then((res) => {
               this.$message(`用户申请已通过`);
               this.load();
            });
         });
      },
      rejectUser(user) {
         this.$confirm(`确认拒绝用户${user.name}申请 ? `).then((res) => {
            V1UserAPI.rejectUser(user.id).then((res) => {
               this.$message(`用户申请已拒绝`);
               this.load();
            });
         });
      },
      load() {
         if(this.activeName == 'management') {
            V1UserAPI.list(1, 999).then((res) => {
               this['management']['tableData'] = res.list;
            });
         }
         if(this.activeName == 'approval') {
            V1UserAPI.fetchUser(this.$appGlobal.constants.USER_APPROVE_STATUS_REGISTED,1,999).then((res) => {
               this['approval']['tableData'] = res.list;
            });
         }
      },
      register() {
         this.loading = true;
         V1UserAPI.register(this.registForm.username, this.registForm.password)
            .then((response) => {
               this.$message(`已添加用户`);
               this.showAddUser = false;
               this.load();
            })
            .finally(() => {
               this.loading = false;
            });
      },
   },
   beforeRouteEnter(to, from, next) {
      store.commit("sidebar/SET_ADD_USR_PAGE");
      next();
   },
};
</script>

<style lang="stylus" scoped>
.container {
   padding-top: 40px;
   padding-left: 20px;

   .btn-ctn {
      width: 150px;
      height: 50px;
      padding-left: 20px;
   }

   .table-ctn {
      padding: 20px;
   }
}


.reset-password {
   text-align: right;
}

a {
   text-decoration: underline;
   color: #409EFF;
}
</style>