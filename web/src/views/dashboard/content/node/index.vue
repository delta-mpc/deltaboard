<template>
   <div class="container">
      <transition appear name="slide-fade">
        <div class="table-ctn">
            <el-table :data="tableData">
               <el-table-column :label="$t('dashboard.nodelist.node_name')" prop="name" width="200px"></el-table-column>
               <el-table-column :label="$t('dashboard.nodelist.node_id')" prop="id"></el-table-column>
               <el-table-column prop="url" :label="$t('dashboard.nodelist.node_addr')"></el-table-column>
            </el-table>
         </div>
         <el-pagination background :current-page.sync="page" :total="totalCount" layout="prev, pager, next" 
               :page-size="page_size" 
               @current-change="load" :hide-on-single-page="true">
         </el-pagination>
      </transition>
   </div>
</template>

<script>
import { mapState } from "vuex";
import V1NodeAPI from "@/api/v1/node";
export default {
   name: "asset",
   data() {
      return {
         tableData:[],
         capsTooltip: false,
         loading: false,
         page:1,
         totalCount:0,
         page_size:20
      };
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
      load() {
         V1NodeAPI.listNodes(this.page,this.page_size).then((res)=>{
            this.tableData = res.list
            this.totalCount = res.total * this.page_size
         })
      },
   }
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