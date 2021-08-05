<template>
  <div class="asset-page">
    <div class="content-bg">
       <iframe style="width:100%;height:100%;border:none" :src="`https://localhost:8000/hub/external/login?next=/hub/&username=${user.email}&token=${user.user_token}`"></iframe>
    </div>
  </div>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'

export default {
  name: "asset",
  data() {
    return {
      cacheList: {},
      assetList: [],
      page: 1,
      totalCount: 0,
    };
  },
  mounted() {
  },
  computed:{
     ...mapState({
        user:state => state.user
     }),
  },
  methods: {
  },
  beforeRouteEnter(to, from, next) {
    store.commit('sidebar/SET_ASSET_PAGE')
    next()
  },
};
</script>

<style lang="stylus" scoped>

.asset-page {
  padding page-padding-top
  background page-bg-color
  height 100%
  -webkit-overflow-scrolling touch
}

.asset-table {
  /deep/.el-table__body {
    // 使表格兼容safari，不错位
    width: 100%;
  }

  /deep/ thead {      
    font-size 16px
    color #053e5d
  }
  /deep/th {
    background #dfe5ec
  }
  th.is-leaf {
    border-bottom 2px solid #053E5D
  }

  /deep/td {
    border-top 1px solid #EBEEF5
    border-bottom none
    cursor pointer
  }

  /deep/.cell {
    font-size 16px
    color theme-color
    padding-left 37px
    white-space nowrap
    text-overflow ellipsis
  }
}

.asset-table::before {
  height 0px
}

.header {
  height 70px
}

.upload-button {
  color #607185
  background #dfe5ec
  border 1px solid #607185
}

.content-bg {
  height 100%
  background white
  border 1px #e4e4e4 solid
}

.el-pagination {
  text-align center
}

</style>