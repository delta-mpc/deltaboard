const freezingInstances = []
export default {
   data(){
      return {
         freezed:false,
         freezeMessage:'',
         freezeInstance:null
      }
   },
   beforeDestroy(){
      freezingInstances.forEach((instance)=>instance.close())
      freezingInstances.length = 0
   },
   watch:{
      'freezed':function(newV,oldV){
         if(newV) {
            freezingInstances.forEach((instance)=>instance.close())
            freezingInstances.length = 0
            freezingInstances.push(this.$loading({fullscreen:true,text:this.freezeMessage}))
         } else {
            freezingInstances.forEach((instance)=>instance.close())
            freezingInstances.length = 0
         }
      }
   }
}