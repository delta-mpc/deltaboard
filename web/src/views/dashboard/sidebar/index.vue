<template>
  <el-aside :width="isCollapse?'64px':'220px'">
  <el-menu
      ref="menu"
      class="sidebar-menu"
      :default-active="barActiveIndex"
      background-color="#f5f8fa"
      text-color="#606266"
      active-text-color="#7DD0D7"
      :collapse="isCollapse"
    >
    <el-menu-item class="sidebar-menu-item" ref="1" index="1" @click="sideItemClicked">
      <font-awesome-icon fixed-width  size="50px" icon="lightbulb" class="side-icon"></font-awesome-icon>
      <span slot="title">{{$t('Playground')}}</span>
    </el-menu-item>
   <el-menu-item class="sidebar-menu-item" ref="2" index="2" @click="sideItemClicked">
     <font-awesome-icon fixed-width size="50px" icon="tasks" class="side-icon"></font-awesome-icon>
     <span>{{$t('dashboard.sidebar.tasklist')}}</span>
    </el-menu-item>
    <el-menu-item class="sidebar-menu-item" ref="3" index="3" @click="sideItemClicked">
      <font-awesome-icon fixed-width  size="50px" icon="network-wired" class="side-icon"></font-awesome-icon>
      <span>{{$t('dashboard.sidebar.nodelist')}}</span>
    </el-menu-item>
    <el-menu-item class="sidebar-menu-item" ref="4" index="4" v-if="user.role == 1" @click="sideItemClicked">
      <font-awesome-icon fixed-width  size="50px" icon="user-friends" class="side-icon"></font-awesome-icon>
      <span>{{$t('dashboard.sidebar.user_management')}}</span>
    </el-menu-item>
    <el-menu-item class="sidebar-menu-item" ref="5" index="5" @click="sideItemClicked">
      <font-awesome-icon fixed-width  size="50px" icon="user-alt" class="side-icon"></font-awesome-icon>
      <span>{{$t('dashboard.sidebar.personal_profile')}}</span>
    </el-menu-item>
    <div class="collapse-toggle" @click="isCollapse = !isCollapse">
      <font-awesome-icon fixed-width  v-if="!isCollapse" size="50px" icon="angle-double-left" class="side-icon"></font-awesome-icon>
      <font-awesome-icon fixed-width  v-if="isCollapse" size="50px" icon="angle-double-right" class="side-icon"></font-awesome-icon>
    </div>
  </el-menu>
  </el-aside>
</template>
<script>
import { mapState } from 'vuex'
const INDEX_MAP = {
   1:'playground',
   2:['myTasks','taskDetail'],
   3:'listnodes',
   4:'userlist',
   5:'profile'
}
export default {
  name: "sideBar",
  data() {
    return {
      isCollapse: false
    };
  },
  computed: {
    ...mapState(['user']),
    barActiveIndex(){
       let routerName = this.$route.name,
           actIdx = Object.keys(INDEX_MAP).filter((itm)=>INDEX_MAP[itm] == routerName || INDEX_MAP[itm].indexOf(routerName) >= 0)
       if(actIdx.length > 0)
         return actIdx[0]
       return "1"
    }
  },
  methods: {
    sideItemClicked(itm) {
      this.$router.push({ name: Array.isArray(INDEX_MAP[itm.index]) ? INDEX_MAP[itm.index][0] : INDEX_MAP[itm.index]});
    }
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
