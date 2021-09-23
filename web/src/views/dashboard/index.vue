<template>
  <div class="app-wrapper">
    <el-container class="dashboard">
      <el-container class="main-container">
        <SideBar/>
        <el-main>
         <PlayGround :key="'playground'" :visible="$route.name === 'playground'"  />
         <router-view :key="'otherpage'" v-if="$route.name !== 'playground'"/>
        </el-main>
      </el-container>
      <div class="side-shadow"></div>
    </el-container>
  </div>
</template>

<script>
import SideBar from "./sidebar"
import { mapState } from 'vuex'
import UserModel from '@/model/user'
import PlayGround from '@/views/dashboard/content/playground/index.vue'
export default {
  name: "dashboard",
  components: {
    SideBar,
    PlayGround
  },
  data() {
    return {
    }
  },
  computed: {
    ...mapState(['user']),
  },
  beforeRouteEnter(to, from, next) {
    UserModel.getMyUserInfo().then((user) => {
        if(user.approve_status != 2) {
           next({name:'post-regist'})
        }
        next()
    }).catch((err) => {
      next({ name: 'login' })
    })
  },
};
</script>

<style lang="stylus" scoped>
.main-container {
   height 100%
}
.app-wrapper {
  overflow:hidden
  flex 1
}

.dashboard {
  height: 100%;
}

.el-header {
  padding: 0;
  height: 60px;
  border-bottom 1px solid #f5f8fa
}

.el-main {
  padding 0px
  margin-bottom 2px
  overflow hidden
  position relative
}
.userauth-dilaog {
   /deep/&.el-dialog__wrapper {
      background white
   }
  /deep/.el-dialog__body {
    display flex
    justify-content center
  }
  .log-content {
    width 520px
  }
  .el-steps {
    margin-top 48px
  }
  .el-form {
    width 386px
    margin 24px auto
  }
  /deep/.el-dialog__header {
    height 32px
  }
  .form-footer {
    text-align center
  }
  .form-item-btn {
    text-align center
  }

  /deep/.el-icon-connection,
  /deep/.el-icon-office-building {
    font-size 40px
  }

  /deep/.el-form-item {
    margin-bottom 28px
  }
  .back-button {
    border none
  }
}

</style>
