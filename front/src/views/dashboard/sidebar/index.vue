<template>
  <el-aside :width="isCollapse?'64px':'220px'">
  <el-menu
      ref="menu"
      class="sidebar-menu"
      :default-active="barActiveIndex"
      background-color="#f5f8fa"
      text-color="#606266"
      active-text-color="#7DD0D7"
      @select="handleSelect"
      :collapse="isCollapse"
    >
    <el-menu-item class="sidebar-menu-item" ref="1" index="1" @click="sideItemClicked('playground')">
      <font-awesome-icon size="50px" icon="lightbulb" class="side-icon"></font-awesome-icon>
      <span slot="title">Playground</span>
    </el-menu-item>
   <el-menu-item class="sidebar-menu-item" ref="2" index="2" @click="sideItemClicked('myTasks')">
     <font-awesome-icon size="50px" icon="tasks" class="side-icon"></font-awesome-icon>
     <span>任务列表</span>
    </el-menu-item>
    <el-menu-item class="sidebar-menu-item" ref="3" index="3" @click="sideItemClicked('listnodes')">
      <font-awesome-icon size="50px" icon="network-wired" class="side-icon"></font-awesome-icon>
      <span>节点列表</span>
    </el-menu-item>
    <el-menu-item class="sidebar-menu-item" ref="4" index="4" v-if="user.role == 1" @click="sideItemClicked('userlist')">
      <font-awesome-icon size="50px" icon="user-friends" class="side-icon"></font-awesome-icon>
      <span>用户管理</span>
    </el-menu-item>
    <el-menu-item class="sidebar-menu-item" ref="5" index="5" @click="sideItemClicked('profile')">
      <font-awesome-icon size="50px" icon="user-alt" class="side-icon"></font-awesome-icon>
      <span>个人中心</span>
    </el-menu-item>
    <div class="collapse-toggle" @click="isCollapse = !isCollapse">
      <font-awesome-icon v-if="!isCollapse" size="50px" icon="angle-double-left" class="side-icon"></font-awesome-icon>
      <font-awesome-icon v-if="isCollapse" size="50px" icon="angle-double-right" class="side-icon"></font-awesome-icon>
    </div>
  </el-menu>
  </el-aside>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: "sideBar",
  data() {
    return {
      isCollapse: false
    };
  },
  computed: {
    ...mapState({
      barActiveIndex: state => { return state.sidebar.sidebarActiveIndex },
    }),
    ...mapState(['user'])
  },
  mounted() {
    window.addEventListener('beforeunload', () => {
      localStorage.setItem('sidebarActiveIndex', this.barActiveIndex)
    })
    if (this.barActiveIndex === null || this.barActiveIndex === undefined) {
      let index = localStorage.getItem('sidebarActiveIndex')
      let pageIndex = { sidebarActiveIndex: index }
      this.$store.commit('sidebar/SET_PAGE_INDEX', pageIndex)
    }
    localStorage.removeItem('sidebarActiveIndex')
  },
  methods: {
    sideItemClicked(name) {
      this.$router.push({ name: name});
    },
    handleSelect(key, keyPath) {
      if (this.user.name.length === 0) {
        return;
      }
      let pageIndex = { sidebarActiveIndex: key }
      this.$store.commit('sidebar/SET_PAGE_INDEX', pageIndex)
    },
  },
};
</script>
<style lang="stylus">
.sidebar-menu {
  height 100%
  padding-top 24px
  box-sizing border-box
  border-right 0

  .sidebar-menu-item:hover,
    .sidebar-menu-item.is-active {
    background-color: #ffffff !important
  }

  .side-icon {
    margin-right 12px
  }

  .collapse-toggle {
    position absolute
    right 8px
    bottom 5px
    width 38px
    height 44px
    line-height 44px
    text-align center
    color #606266
    cursor pointer
    border-radius 5px
    padding-left 10px
    padding-top 4px

    &:hover {
      background-color white
    }
  }
}

</style>
