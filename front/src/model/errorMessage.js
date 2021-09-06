import {ApiPromise} from "@/api/api-promise";
import ErrorMessage from '@/model/errorMessage'
import ElementUI from 'element-ui'
export const errorMessages = [
{
    server: 'user_already_exist',
    local: '用户已存在'
}, {
    server: 'user_not_found',
    local: '用户不存在'
}, {
    server: 'invalid_password',
    local: '密码不正确'
}, {
    server: 'user_already_exist',
    local: '用户已存在'
}, {
    server: 'user_already_exist',
    local: '用户已存在'
}, {
    server: 'user_already_exist',
    local: '用户已存在'
}, {
    server: 'user_already_exist',
    local: '用户已存在'
},{
   server:'user login status expired',
   local:'登陆状态已过期'
}
]

export const UserLoginStatuExpired = "user login status expired"

ApiPromise.defaultOnReject = (err) => {
   let errorMessage = err.response.data.message
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
            if (message.server === server) { label = message.local }
        })
        return label
    },
}