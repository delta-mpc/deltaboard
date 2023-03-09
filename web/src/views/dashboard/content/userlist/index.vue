<template>
   <div class="container">
      <transition appear name="slide-fade">
         <el-tabs v-model="activeName">
            <el-tab-pane :label="$t('dashboard.userlist.management')" name="management">
               <div class="btn-ctn">
                  <el-button size="medium" type="primary" @click="showAddUser = true">{{$t('dashboard.userlist.add_user')}}</el-button>
               </div>
               <div class="table-ctn" v-infinite-scroll="load">
                  <el-table :data="management.tableData" header-row-class-name='header-row'>
                     <el-table-column :label="$t('dashboard.userlist.user_name')" prop="name"></el-table-column>
                     <el-table-column prop="created_at" :label="$t('common.created_at')">
                        <template v-slot="{ row }">{{ row.created_at | second2Date }}</template>
                     </el-table-column>
                     <el-table-column :label="$t('dashboard.userlist.commands')" prop="name">
                        <template v-slot="{ row }">
                           <el-tag class="clickable" type="danger" @click="delUser(row)">{{$t('dashboard.userlist.delete')}}</el-tag>
                        </template>
                     </el-table-column>
                  </el-table>
               </div>
            </el-tab-pane>
            <el-tab-pane v-if="config.public_registration" :label="$t('dashboard.userlist.check&approve')" name="approval">
               <div class="table-ctn" v-infinite-scroll="load">
                  <el-table :data="approval.tableData" header-row-class-name='header-row'>
                     <el-table-column :label="$t('dashboard.userlist.user_name')" prop="name"></el-table-column>
                     <el-table-column prop="created_at" :label="$t('common.created_at')">
                        <template v-slot="{ row }">{{ row.created_at | second2Date }}</template>
                     </el-table-column>
                     <el-table-column :label="$t('dashboard.userlist.commands')" prop="name">
                        <template v-slot="{ row }">
                           <el-row>
                              <el-col span="4">
                                 <el-tag class="clickable" type="success" @click="approveUser(row)">{{$t('dashboard.userlist.approve')}}</el-tag>
                              </el-col>
                              <el-col span="4">
                                 <el-tag class="clickable" type="danger" @click="rejectUser(row)">{{$t('dashboard.userlist.reject')}}</el-tag>
                              </el-col>
                           </el-row>
                        </template>
                     </el-table-column>
                  </el-table>
               </div>
            </el-tab-pane>
         </el-tabs>
      </transition>
      <el-dialog :visible.sync="showAddUser" :title="$t('dashboard.userlist.add_user')" width="500px">
         <el-form ref="registForm" :model="registForm" class="login-form" autocomplete="on" label-position="left">
            <el-form-item prop="username">
               <el-input ref="username" v-model="registForm.username" :placeholder="$t('common.please_input_your_name')" name="username" type="text"
                  tabindex="4" autocomplete="on" />
            </el-form-item>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
               <el-form-item prop="password">
                  <el-input :key="registForm.passwordType" ref="registpassword" v-model="registForm.password"
                     :type="registForm.passwordType" :placeholder="$t('common.please_input_your_pwd')" name="password" tabindex="2"
                     autocomplete="on" />
               </el-form-item>
            </el-tooltip>
            <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
               <el-form-item prop="repeatPassword">
                  <el-input :key="registForm.repeatPasswordType" ref="repeatpassword"
                     v-model="registForm.repeatPassword" :type="registForm.repeatPasswordType" 
                     :placeholder="$t('common.please_repeat_your_pwd')"
                     name="repeatpassword" tabindex="2" autocomplete="on" />
               </el-form-item>
            </el-tooltip>
            <el-button :loading="loading" type="primary"
               @click.native.prevent="register">{{$t('dashboard.userlist.add_user')}}</el-button>
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
            tableData:[],
            page_size:12,
            page:1,
            sort: 0,
         },
         approval:{
            tableData:[],
            page_size:12,
            page:1,
            sort: 1,
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
      'activeName':function(newV,oldV){
         if(newV != oldV) {
            this.management.page = 1
            this.management.tableData = []
            this.approval.page = 1
            this.approval.tableData = []
            this.load()
         }
      }
   },
   computed: {
      ...mapState({
         user: (state) => state.user,
         config:(state)=> state.config.config
      })
   },
   methods: {
      delUser(user) {
         this.$confirm(this.$t('dashboard.userlist.confirm_delete',{name:user.name})).then((res) => {
            V1UserAPI.delUser(user.id).then((res) => {
               this.$message(this.$t('dashboard.userlist.user_deleted'));
               this.management.tableData = this.management.tableData.filter((itm)=>itm.id != user.id)
               // this.load();
            });
         });
      },
      approveUser(user){
         this.$confirm(this.$t('dashboard.userlist.approve_user',{name:user.name})).then((res) => {
            V1UserAPI.approveUser(user.id).then((res) => {
               this.$message(this.$t('dashboard.userlist.user_approved'));
               this.approval.tableData = this.approval.tableData.filter((itm)=>itm.id != user.id)
               // this.load();
            });
         });
      },
      rejectUser(user) {
         this.$confirm(this.$t('dashboard.userlist.reject_user',{name:user.name})).then((res) => {
            V1UserAPI.rejectUser(user.id).then((res) => {
               this.$message(this.$t('dashboard.userlist.user_rejected'));
               this.approval.tableData = this.approval.tableData.filter((itm)=>itm.id != user.id)
               // this.load();
            });
         });
      },
      load() {
         if(this.activeName == 'management') {
            V1UserAPI.fetchUser(this.$appGlobal.constants.USER_APPROV_STATUS_APPROVED,this.management.page,this.management.page_size, this.management.sort).then((res) => {
               let lst = res.list.filter((itm)=>this['management']['tableData'].findIndex((data)=>data.id == itm.id) < 0)
               this['management']['tableData'].push(...lst);
               this.management.page += 1
            });
         }
         if(this.activeName == 'approval') {
            V1UserAPI.fetchUser(this.$appGlobal.constants.USER_APPROVE_STATUS_REGISTED,this.approval.page,this.approval.page_size, this.approval.sort).then((res) => {
               let lst = res.list.filter((itm)=>this['approval']['tableData'].findIndex((data)=>data.id == itm.id) < 0)
               this['approval']['tableData'].push(...lst);
               this.approval.page += 1
            });
         }
      },
      register() {
         this.loading = true;
         V1UserAPI.register(this.registForm.username, this.registForm.password)
            .then((response) => {
               this.$message(this.$t('dashboard.userlist.user_added'));
               this.showAddUser = false;
               this.load();
            })
            .finally(() => {
               this.loading = false;
            });
      },
   }
};
</script>
<style>
.header-class{
   position: absolute;
   top :0px
}
</style>
<style lang="stylus" scoped>
.container {
   padding-top: 40px;
   padding-left: 20px;
   height calc(100vh - 105px)
   overflow hidden
   .btn-ctn {
      width: 150px;
      height: 50px;
      padding-left: 20px;
   }
   .header-row {
      background #ccc
   }
   .table-ctn {
      padding: 20px;
      max-height calc(100vh - 240px)
      overflow:auto
      &::-webkit-scrollbar {
         width: 0px;
         background-color: #e5e5e5;
      }
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