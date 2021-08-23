<template>
  <el-menu ref="menu" class="sidebar-menu" :default-active="barActiveIndex" @select="handleSelect">
    <el-menu-item class="sidebar-menu" ref="1" index="1" @click="sideItemClicked('playground')">
       <div class="item-div">
          <font-awesome-icon size="50px" icon="lightbulb" class="side-icon">
          </font-awesome-icon><span>Playground</span></div>
    </el-menu-item>
   <el-menu-item class="sidebar-menu" ref="2" index="2" @click="sideItemClicked('myTasks')">
      <div class="item-div">
         <font-awesome-icon size="50px" icon="tasks" class="side-icon">
         </font-awesome-icon><span>任务列表</span></div>
    </el-menu-item>
    <el-menu-item class="sidebar-menu" ref="3" index="3" @click="sideItemClicked('listnodes')">
      <div class="item-div">
         <font-awesome-icon size="50px" icon="network-wired" class="side-icon">
          </font-awesome-icon><span>节点列表</span></div>
    </el-menu-item>
    <el-menu-item class="sidebar-menu" ref="4" index="4" v-if="user.role == 1" @click="sideItemClicked('userlist')">
      <div class="item-div"><font-awesome-icon size="50px" icon="user-friends" class="side-icon">
          </font-awesome-icon><span>用户管理</span></div>
    </el-menu-item>
    <el-menu-item class="sidebar-menu" ref="5" index="5" @click="sideItemClicked('profile')">
      <div class="item-div"><font-awesome-icon size="50px" icon="user-alt" class="side-icon">
          </font-awesome-icon><span>个人中心</span></div>
    </el-menu-item>
  </el-menu>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: "sideBar",
  data() {
    return {};
  },
  computed: {
    ...mapState({
      barActiveIndex: state => { return state.sidebar.sidebarActiveIndex },
    }),
    ...mapState(['user']),
    maskCss(){
       return {
         width:'100%',
         pointerEvents:'none',
         height:'50px',
         marginLeft:'9px',
         position:'absolute',
         left:'0px',
         top:50 * (this.barActiveIndex - 1) + 'px'
       }
    }
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

<style lang="stylus" scoped>
.sidebar-menu {
   border-right none !important
}
.maskClz {
   background page-bg-color
}
.el-menu-item {
  padding-left 9px !important
  padding-right 0px
  height 50px
  line-height 50px

  border-right 1px solid #f5f8fa
  .side-icon, {
    margin-left 30px
    margin-right 30px
  }

  .item-div {
    height 50px
  }

  &:hover, &:focus {
    background white
    .item-div {
      background page-bg-color

      span {
        color #053e5d
      }
    }
  }
}

.el-menu-item.is-active {
  .item-div {
    background page-bg-color

    span {
      color #053e5d
    }
  }
}

span {
  color #b6b7b9
}

</style>


<style>

.tool-item-1 {
  padding: 5px 10px;
  top: 120px !important;
}

.tool-item-2 {
  padding: 5px 10px;
  top: 170px !important;
}

.tool-item-3 {
  padding: 5px 10px;
  top: 270px !important;
}

</style>