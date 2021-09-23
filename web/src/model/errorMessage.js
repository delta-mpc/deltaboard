import ErrorMessage from '@/model/errorMessage'
import ElementUI from 'element-ui'
import Vue from 'vue'
import {ApiPromise} from "@/api/api-promise";
export const errorMessages = [
   {
       server: 'user_already_exist',
       local: 'dashboard.errormessage.user_already_exist'
   }, {
       server: 'user_not_found',
       local:  'dashboard.errormessage.user_not_found'
   }, {
       server: 'invalid_password',
       local:  'dashboard.errormessage.invalid_password'
   }, {
      server:'user login status expired',
      local: 'dashboard.errormessage.login_status_expired'
   }
]

export const UserLoginStatuExpired = "user login status expired"

ApiPromise.defaultOnReject = (err) => {
   let errorMessage = (err.response && err.response.data && err.response.data.message) || err
   let showMessage = ErrorMessage.getLocalMessage(errorMessage)
   if (showMessage) {
      ElementUI.Message({ message: showMessage, type: 'error', duration: 3 * 1000})
   }
   else {
      ElementUI.Message.error(err.toString());
   }
};

export const $errorMessage = (error,cb)=>{
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

export default {
    getMessageTypes() {
        return errorMessages
    },
    
    getLocalMessage(server) {
        let label = null
        errorMessages.forEach((message) => {
            if (message.server === server) { label = Vue.prototype.$t(message.local) }
        })
        return label
    },
}