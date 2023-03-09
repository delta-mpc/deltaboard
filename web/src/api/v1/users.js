
import v1 from './v1';

export default {
   register(username, password) {
      return v1.post('/users/tokens', {'user_name': username, 'password': password})
   },
   fetchUser(apprv_status,pageID,pageSize,sort) {
      return v1.get('/users/approve_status', {params:{approve_status:apprv_status,page: pageID, page_size: pageSize, sort: sort}})
   },
   approveUser(userId) {
      return v1.post(`/users/approve/${userId}`)
   },
   rejectUser(userId) {
      return v1.post(`/users/reject/${userId}`)
   },
   delUser(userId) {
      return v1.del(`/users/del/${userId}`)
   },
   renewDeltaToken(userId) {
      return v1.post(`/users/delta_token/renew`)
   },
   login(username, password) {
      return v1.post('/users', {'user_name': username, 'password': password})
   },
   list(pageID,pageSize){
      return v1.get('/users/list', {params:{page_id: pageID, page_size: pageSize}})
   },
    resetLoginPass(token, password) {
        return v1.put('/users/password/' + token, { 'password': password})
    },
    changeLoginPass(oldPass, newPass) {
        return v1.put('/users/password', {'new_password': newPass, 'old_password': oldPass})
    },
    logout() {
        return v1.del('/users/tokens')
    },
    getMyUserInfo() {
        return v1.get('/users', { 'withCredentials': true })
    },
    updatePhoneNumber(new_phone_number,validation_code) {
      return v1.post('/vault/authorization',{new_phone_number:new_phone_number,validation_code:validation_code})
    }
};
