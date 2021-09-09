<template>
   <div class="task-drawer">
      <el-page-header @back="$router.go(-1)" :content="`任务 ${currentTaskId}`"></el-page-header>
      <div class="drawer-wrapper">
         <el-descriptions title="Metadata">
            <el-descriptions-item label="创建日期" :labelStyle="labelStyle">{{taskLogMetaData.created_at || 0 | second2Date }}
            </el-descriptions-item>
            <el-descriptions-item label="创建者" :labelStyle="labelStyle">{{taskLogMetaData.creator}}</el-descriptions-item>
            <el-descriptions-item label="ID" :labelStyle="labelStyle">{{taskLogMetaData.id}}</el-descriptions-item>
            <el-descriptions-item label="任务名" :labelStyle="labelStyle">
               {{taskLogMetaData.name}}
            </el-descriptions-item>
            <el-descriptions-item label="任务状态" :labelStyle="labelStyle">
               <el-tag type="warning" v-if="taskLogMetaData.status == $appGlobal.constants.TASK_STATUS_PENDING">
                  等待运行
               </el-tag>
               <el-tag type="success" v-else-if="taskLogMetaData.status == $appGlobal.constants.TASK_STATUS_RUNNING">
                  运行中
               </el-tag>
               <el-tag type="info" v-else>
                  完成
               </el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="任务类型" :labelStyle="labelStyle">{{taskLogMetaData.type}}</el-descriptions-item>
         </el-descriptions>
         <el-button size="medium" type="primary" v-if="taskLogMetaData.status == 'FINISHED'"
            @click.stop="downloadWeights">
            下载权重</el-button>
         <div class="el-descriptions__title task-log-title">任务日志</div>
         <div class="task-logs" v-infinite-scroll="loadTaskLog" style="overflow:auto">
             <div v-for="itm,index in taskLogData" :key="index">
                  <span class="date">{{itm.created_at / 1000 | second2Date}}</span>
                  <span>{{itm.message}}</span>
               </div>
         </div>
      </div>
   </div>
</template>

<script>
import { mapState } from "vuex";
import V1TaskAPI from "@/api/v1/tasks";
export default {
   name: "playground",
   mounted() {},
   computed: {
      ...mapState({
         user: (state) => state.user,
      }),
      currentTaskId() {
         return this.$route.params.task_id;
      },
   },
   data() {
      return {
         taskLogMetaData: {},
         taskLogData: [],
         page: 1,
         page_size: 60,
         labelStyle:{
            display:'flex',
            alignItems:'center'
         }
      };
   },
   methods: {
      downloadWeights() {
         window.open(
            `${window.BASE_API}/v1/tasks/result/${this.taskLogMetaData.id}`,
            "_blank"
         );
      },
      init() {
         this.page = 1;
         this.loadTaskLog();
         this.loadTaskMeta();
      },
      loadTaskMeta() {
         V1TaskAPI.getTaskMeta(this.currentTaskId).then((res) => {
            this.taskLogMetaData = res.tasks ? res.tasks[0] : [];
         });
      },
      loadTaskLog() {
         V1TaskAPI.getTaskLogs(
            this.currentTaskId,
            this.page,
            this.page_size
         ).then((res) => {
            this.taskLogData = this.taskLogData.concat(res);
            this.page += 1;
         });
      },
   },
   beforeRouteEnter(to, from, next) {
      // store.commit("sidebar/SET_PLAYGROUND_PAGE");
      next((vm) => {
         vm.init();
      });
   },
};
</script>

<style lang="stylus" scoped>
.back {
   position: absolute;
   left: 5px;
   font-size: 18px;
}

.upload-button {
   color: #607185;
   background: #dfe5ec;
   border: 1px solid #607185;
}

.content-bg {
   height: 100%;
   background: white;
   border: 1px #e4e4e4 solid;
}

.task-log-title {
   margin-top: 20px;
   margin-bottom: 20px;
}

.task-logs {
   height: 600px;
   background: rgb(24, 48, 85);
   overflow auto
   padding 20px
   box-sizing border-box
   &::-webkit-scrollbar {
      width: 6px;
      height: 6px
      background-color: #e5e5e5;
   }

   &::-webkit-scrollbar-thumb {
      background-color: #b7b7b7;
      border-radius: 3px;
   }

   div {
      list-style: none;
      line-height: 22px;
      color: rgb(230, 236, 241);
      tab-size: 2;
      direction: ltr;
      font-size: 14px;
      background: rgb(24, 48, 85);
      text-align: left;
      font-family: 'Source Code Pro', Consolas, Menlo, Monaco, Courier, monospace;
      line-height: 1.4;
      white-space: pre;
      word-spacing: normal;
      border-radius: 3px;
      .date {
         width: 100px;
         margin-right: 10px;
      }
   }
}

.task-drawer {
   padding: 30px;
   position: relative;
}

.drawer-wrapper {
   padding: 30px 0px;
}
</style>