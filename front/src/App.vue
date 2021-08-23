<template>
  <div id="app">
    <router-view/>
  </div>
</template>
<script>
import ConfigAPI from '@/api/v1/config'
export default {
   created(){
      ConfigAPI.getConfig().then((res)=>{
         this.$store.commit('config/SET_CONFIG',res)
      })
   }
}
</script>
<style>

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

</style>
