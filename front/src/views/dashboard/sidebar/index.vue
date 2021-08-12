<template>
  <el-menu ref="menu" class="sidebar" :default-active="barActiveIndex" @select="handleSelect">
    <el-menu-item ref="1" index="1" @click="sideItemClicked('playground')">
       <div class="item-div"><i class="el-icon-notebook-2 side-icon"></i><span>Playground</span></div>
    </el-menu-item>
    <el-menu-item ref="2" index="2" v-if="user.name == 'admin'" @click="sideItemClicked('userlist')">
      <div class="item-div"><i class="el-icon-user side-icon"></i><span>Add User</span></div>
    </el-menu-item>
    <el-menu-item ref="3" index="3" @click="sideItemClicked('changepassword')">
      <div class="item-div"><i class="el-icon-setting side-icon"></i><span>Change Password</span></div>
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
.maskClz {
   background page-bg-color
}
.el-menu {
  border-right none
  margin-top 30px
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
    border-left 3px solid white
    height 50px
  }

  &:hover, &:focus {
    background white
    .item-div {
      background page-bg-color
      border-left 3px solid black

      span {
        color #053e5d
      }
    }
  }
}

.el-menu-item.is-active {
  .item-div {
    background page-bg-color
    border-left 3px solid black

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