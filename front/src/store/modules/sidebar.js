
const SET_PLAYGROUND_PAGE = 'SET_PLAYGROUND_PAGE'
const SET_ADD_USR_PAGE = 'SET_ADD_USR_PAGE'
const SET_SETTING_PAGE = 'SET_SETTING_PAGE'
const SET_PAGE_INDEX = 'SET_PAGE_INDEX'
const SET_CHGPWD_PAGE = 'SET_CHGPWD_PAGE'
const state = {
    sidebarActiveIndex: null,
}

const mutations = {
   [SET_PLAYGROUND_PAGE](state) {
      state.sidebarActiveIndex = '1'
   },
    [SET_ADD_USR_PAGE](state) {
        state.sidebarActiveIndex = '2'
    },
    [SET_CHGPWD_PAGE](state) {
      state.sidebarActiveIndex = '3'
    },
    [SET_PAGE_INDEX](state, payload) {
        Object.assign(state, payload)
    },
}

export default {
    namespaced: true,
    state,
    mutations,
}