<template>
   <div class="container">
      <transition appear name="slide-fade">
      <div class="table-ctn">
         <div class="tbl-wrapper">
            <el-table :data="tableData" @row-click="openDetail" row-class-name="clickable">
               <el-table-column :label="$t('dashboard.tasklist.task_id')" prop="id"></el-table-column>
               <el-table-column :label="$t('dashboard.tasklist.task_name')" prop="name"></el-table-column>
               <el-table-column :label="$t('dashboard.tasklist.task_type')" prop="type"></el-table-column>
               <el-table-column :label="$t('dashboard.tasklist.creator_node')" prop="creator">
                  <template v-slot="{ row }">
                        {{row.creator}}({{nodes[row.creator]}})
                  </template>
               </el-table-column>
               <el-table-column  :label="$t('dashboard.tasklist.creator_account')" prop="creator_name">
                     <template v-slot="{ row }">
                          {{row.creator_name || "--" }}
                     </template>
               </el-table-column>
               <el-table-column :label="$t('dashboard.tasklist.status')" prop="status">
                  <template v-slot="{ row }">
                     <el-tag type="warning" v-if="row.status == $appGlobal.constants.TASK_STATUS_PENDING">
                        {{$t('dashboard.tasklist.pending')}}
                     </el-tag>
                     <el-tag type="success" v-else-if="row.status == $appGlobal.constants.TASK_STATUS_RUNNING">
                        {{$t('dashboard.tasklist.running')}}
                     </el-tag>
                     <el-tag type="danger" v-else-if="row.status == $appGlobal.constants.TASK_STATUS_ERROR">
                        {{$t('dashboard.tasklist.error')}}
                     </el-tag>
                     <el-tag type="info" v-else-if="row.status == $appGlobal.constants.TASK_STATUS_FINISHED">
                        {{$t('dashboard.tasklist.finished')}}
                     </el-tag>
                     <el-tag type="info" v-else-if="row.status == $appGlobal.constants.TASK_STATUS_CONFIRMED">
                        {{$t('dashboard.tasklist.confirmed')}}
                     </el-tag>
                  </template>
               </el-table-column>
               <el-table-column prop="created_at" :label="$t('common.created_at')">{{  }}
                  <template v-slot="{ row }">{{ (row.created_at / 1000) | second2Date }}</template>
               </el-table-column>
            </el-table> 
         </div>  
         <div class="paging">
         <el-pagination background :current-page.sync="currentPage" :total="totalCount" layout="prev, pager, next" 
               :page-size="taskPageSize" 
               @current-change="load" :hide-on-single-page="true"></el-pagination>
         </div> 
      </div>  
      </transition>
    </div>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'
import V1TaskAPI from "@/api/v1/tasks"
import V1NodeAPI from "@/api/v1/node";
export default {
  name: "asset",
  data() {
    return {
      tableData:[],
      showTaskDetail:false,
      currentPage:1,
      taskPageSize:10,
      totalCount:0,
      nodes:{},
      taskLogPage:{
         taskLogMetaData:{},
         taskLogData:[],
         currentTaskId:null,
         start: 0,
         limit: 60
      }
    };
  },
  mounted() {
   V1NodeAPI.listNodes(1,999).then((res)=>{
      this.nodes = (res.list || []).reduce((pre,cur)=>{
         pre[cur.id] = cur.name
         return pre
      },{})
      this.load()
   })
  },
  computed:{
     ...mapState({
        user:state => state.user
     })
  },
  methods: {
     openDetail(row){
        this.$router.push({path:`task/${row.id}`})
     },
     onOpened(){
        this.taskLogPage.start = 0
        this.loadTaskLog()
        this.loadTaskMeta()
     },  
     viewTaskDetail(id) {
        this.taskLogPage.currentTaskId = id
        this.taskLogPage.taskLogData = []
        this.showTaskDetail = true
     },  
     load(){
        if(this.user.role != 1) {
            V1TaskAPI.getUserTasks(this.user.id,this.currentPage,this.taskPageSize).then((res)=>{
               this.tableData = res.tasks
               // total_pages is total_count
               this.totalCount = res.total_pages
            })
        } else {
            V1TaskAPI.getAllTasks(this.currentPage,this.taskPageSize).then((res)=>{
               (res.user_tasks ||[]).forEach((itm)=>{
                  res.tasks.forEach((tItm)=>{
                     if(tItm.id == itm.nodeTaskId) {
                        tItm.creator_name = itm.name
                     }
                  })
               })
               this.tableData = res.tasks
               this.totalCount = res.total_pages * this.taskPageSize
            })
        }
     },
     loadTaskMeta(){
        V1TaskAPI.getTaskMeta(this.taskLogPage.currentTaskId).then((res)=>{
           this.taskLogPage.taskLogMetaData = res.tasks ? res.tasks[0] :[]
        })
     },
     loadTaskLog(){
        V1TaskAPI.getTaskLogs(this.taskLogPage.currentTaskId,
            this.taskLogPage.start,
            this.taskLogPage.limit).then((res)=>{
            this.taskLogPage.taskLogData = this.taskLogPage.taskLogData.concat(res);
            console.log(this.taskLogPage.taskLogData);
            let finalLogData = res[res.length - 1];
            this.taskLogPage.start = finalLogData.id + 1;
      })
     }
  }
};
</script>

<style lang="stylus" scoped>
.paging {
   margin-top 20px
   display flex
   justify-content center
}
.container {
  padding-top 40px
  padding-left 20px
  height 100%
  .btn-ctn {
     width 150px
     height 50px
     padding-left 20px
  }
  .table-ctn {
     padding 20px
     height 100%
     .tbl-wrapper {
         height calc(100% - 200px)
         overflow hidden
     }
  }
}
.task-log-title {
   margin-top 20px
   margin-bottom 20px
}
.task-logs {
   height 600px
   background rgb(24, 48, 85)
   &::-webkit-scrollbar {
        width: 6px;
        background-color: #e5e5e5;
    }
   &::-webkit-scrollbar-thumb {
        background-color: #b7b7b7;
        border-radius: 3px;
   }
   li {
      list-style none
      line-height 22px
      color rgb(230, 236, 241)
      overflow: auto;
      tab-size: 2;
      direction: ltr;
      font-size: 14px;
      background: rgb(24, 48, 85);
      text-align: left;
      font-family: "Source Code Pro", Consolas, Menlo, Monaco, Courier, monospace;
      line-height: 1.4;
      white-space: pre;
      word-spacing: normal;
      border-radius: 3px;
      .date {
         width 100px
         margin-right 10px
      }
   }
}
.task-drawer {
   padding 20px
}
.reset-password {
  text-align right
}

.action-btn {
   width 100px
   font-size 10px
}

a {
  text-decoration underline
  color #409EFF
}
</style>