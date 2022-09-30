<template>
  <div class="task-drawer">
    <el-page-header
      @back="$router.go(-1)"
      :content="
        this.$t('dashboard.taskdetail.task {id}', { id: currentTaskId })
      "
    ></el-page-header>
    <div class="drawer-wrapper">
      <el-descriptions title="Metadata">
        <el-descriptions-item
          :label="$t('common.created_at')"
          :labelStyle="labelStyle"
          >{{ taskLogMetaData.created_at / 1000 || 0 | second2Date }}
        </el-descriptions-item>
        <el-descriptions-item
          :label="$t('dashboard.taskdetail.creator')"
          :labelStyle="labelStyle"
          >{{ taskLogMetaData.creator }}</el-descriptions-item
        >
        <el-descriptions-item label="ID" :labelStyle="labelStyle">{{
          taskLogMetaData.id
        }}</el-descriptions-item>
        <el-descriptions-item
          :label="$t('dashboard.taskdetail.task_name')"
          :labelStyle="labelStyle"
        >
          {{ taskLogMetaData.name }}
        </el-descriptions-item>
        <el-descriptions-item
          :label="$t('dashboard.taskdetail.status')"
          :labelStyle="labelStyle"
        >
          <el-tag
            type="warning"
            v-if="
              taskLogMetaData.status == $appGlobal.constants.TASK_STATUS_PENDING
            "
          >
            {{ $t("dashboard.taskdetail.pending") }}
          </el-tag>
          <el-tag
            type="success"
            v-else-if="
              taskLogMetaData.status == $appGlobal.constants.TASK_STATUS_RUNNING
            "
          >
            {{ $t("dashboard.taskdetail.running") }}
          </el-tag>
          <el-tag
            type="danger"
            v-else-if="
              taskLogMetaData.status == $appGlobal.constants.TASK_STATUS_ERROR
            "
          >
            {{ $t("dashboard.taskdetail.error") }}
          </el-tag>
          <el-tag
            type="info"
            v-else-if="taskLogMetaData.status == $appGlobal.constants.TASK_STATUS_FINISHED"
          >
            {{ $t("dashboard.taskdetail.finished") }}
          </el-tag>
          <el-tag
            type="info"
            v-else-if="taskLogMetaData.status == $appGlobal.constants.TASK_STATUS_CONFIRMED"
          >
            {{ $t("dashboard.taskdetail.confirmed") }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item
          :label="$t('dashboard.tasklist.task_type')"
          :labelStyle="labelStyle"
          >{{ taskLogMetaData.type }}</el-descriptions-item
        >
      </el-descriptions>
      <el-button
        size="medium"
        type="primary"
        v-if="taskLogMetaData.status == 'FINISHED' || taskLogMetaData.status == 'CONFIRMED'"
        @click.stop="downloadWeights"
      >
        {{ $t("dashboard.taskdetail.download_result") }}
      </el-button>
      <div class="el-descriptions__title task-log-title">
        {{ $t("dashboard.taskdetail.task_logs") }}
      </div>
      <div
        class="task-logs"
        v-infinite-scroll="loadTaskLog"
        style="overflow: auto"
      >
        <div v-for="(itm, index) in taskLogData" :key="index">
          <div class="content">
            {{ (itm.created_at / 1000) | second2Date }} &nbsp;&nbsp;
            {{ itm.message }}
          </div>
          <div class="content" v-if="itm.tx_hash">
            tx hash:
            <a target="_blank" :href="txUrl(itm.tx_hash)">{{ itm.tx_hash }}</a>
          </div>
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
      logStart: 0,
      logLimit: 60,
      labelStyle: {
        display: "flex",
        alignItems: "center",
      },
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
      this.logStart = 0;
      // this.loadTaskLog();
      this.loadTaskMeta();
    },
    loadTaskMeta() {
      V1TaskAPI.getTaskMeta(this.currentTaskId).then((res) => {
        this.taskLogMetaData = res.tasks ? res.tasks[0] : [];
      });
    },
    loadTaskLog() {
      console.log(`load task log page ${this.page}`);
      V1TaskAPI.getTaskLogs(
        this.currentTaskId,
        this.logStart,
        this.logLimit
      ).then((res) => {
        this.taskLogData = this.taskLogData.concat(res);
        let finalLogData = res[res.length - 1];
        this.logStart = finalLogData.id + 1;
        console.log(`load task log complete, page increase to ${this.page}`);
      });
    },
    txUrl(txHash) {
      return `https://explorer.deltampc.com/tx/${txHash}/internal-transactions`;
    },
  },
  beforeRouteEnter(to, from, next) {
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
  overflow: auto;
  padding: 20px;
  box-sizing: border-box;

  &::-webkit-scrollbar {
    width: 6px;
    height: 6px;
    background-color: #e5e5e5;
  }

  &::-webkit-scrollbar-thumb {
    background-color: #b7b7b7;
    border-radius: 3px;
  }

  div {
    background: rgb(24, 48, 85);
    border-radius: 3px;

    .content {
      width: 100%;
      font-size: 14px;
      direction: ltr;
      color: rgb(230, 236, 241);
      tab-size: 2;
      list-style: none;
      text-align: left;
      font-family: 'Source Code Pro', Consolas, Menlo, Monaco, Courier, monospace;
      line-height: 1.4;
      word-spacing: normal;
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
