import UserAPI from '@/api/v1/users'
import store from '@/store'

export default {
    getMyUserInfo() {
        return UserAPI.getMyUserInfo().then((response) => {
            store.commit('user/SET_USER', response)
            return response
        }).catch((err) => {
            return Promise.reject(err)
        })
    },
    updateUserPhone(newPhoneNum,validation_code){
       return UserAPI.updatePhoneNumber(newPhoneNum,validation_code).then((res)=>{
         return this.getMyUserInfo()
       })
    },
}