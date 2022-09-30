import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import '@/styles/element-variables.scss'
import '@/styles/abstracts/index.styl'
import {ApiPromise} from "@/api/api-promise";
import filters from '@/filters/index.js'
import ErrorMessage from '@/model/errorMessage'
import VueClipboard from 'vue-clipboard2'
import { library } from '@fortawesome/fontawesome-svg-core'
import {
    faSignOutAlt,
    faLightbulb,
    faTasks,
    faNetworkWired,
    faUserFriends,
    faUserAlt,
    faAngleDoubleLeft, faAngleDoubleRight
} from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import i18n from '@/i18n/i18n.js'
Vue.prototype.$t = (key, value) => i18n.t(key, value);
(function () {
   File.prototype.arrayBuffer = File.prototype.arrayBuffer || myArrayBuffer;
   Blob.prototype.arrayBuffer = Blob.prototype.arrayBuffer || myArrayBuffer;

   function myArrayBuffer() {
     return new Promise((resolve) => {
       let fr = new FileReader();
       fr.onload = () => {
         resolve(fr.result);
       };
       fr.readAsArrayBuffer(this);
     })
   }
 })();
window.BASE_API = (process.env.VUE_APP_BASE_API && process.env.VUE_APP_BASE_API != '') ?
                  process.env.VUE_APP_BASE_API : window.location.protocol + '//' + window.location.host
Vue.use(VueClipboard)

library.add(faSignOutAlt,faLightbulb,faTasks,faNetworkWired,faUserFriends,faUserAlt,faAngleDoubleLeft, faAngleDoubleRight)

Vue.component('font-awesome-icon', FontAwesomeIcon)
console.log('initializing')
Vue.config.productionTip = false

Vue.use(ElementUI)

Object.keys(filters).forEach((key) => {
  Vue.filter(key, filters[key])
})

Vue.prototype.$message = ElementUI.Message
Vue.prototype.$appGlobal = {
   constants:{
      TASK_STATUS_PENDING:'PENDING',
      TASK_STATUS_RUNNING:'RUNNING',
      TASK_STATUS_FINISHED:'FINISHED',
      TASK_STATUS_ERROR:'ERROR',
      TASK_STATUS_CONFIRMED: 'CONFIRMED',
      USER_APPROVE_STATUS_REGISTED:1,
      USER_APPROV_STATUS_APPROVED:2
   }
}

Vue.prototype.$errorMessage = (error,cb)=>{
   try {
      let errorMessage = error.response.data.message || (error.response.data.data && error.response.data.data.message)
      let showMessage = ErrorMessage.getLocalMessage(errorMessage)
      if (showMessage) {
         ElementUI.Message({ message: showMessage, type: 'error', duration: 3 * 1000})
      }
      else {
         if(cb) {
            cb(error)
         } else {
            ApiPromise.defaultOnReject(error)
         }
      }
   } catch(e){
      ApiPromise.defaultOnReject(error)
   }
}
new Vue({
  router,
  store,
  i18n,
  render: h => h(App),
}).$mount('#app')

