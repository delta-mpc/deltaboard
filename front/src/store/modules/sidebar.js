const SET_VAULT_PAGE = 'SET_VAULT_PAGE'
const SET_PLAYGROUND_PAGE = 'SET_PLAYGROUND_PAGE'
const SET_COMPANY_PAGE = 'SET_COMPANY_PAGE'
const SET_CREDIT_PAGE = 'SET_CREDIT_PAGE'
const SET_SETTING_PAGE = 'SET_SETTING_PAGE'
const SET_PAGE_INDEX = 'SET_PAGE_INDEX'
const SET_APP_PAGE = 'SET_APP_PAGE'
const state = {
    sidebarActiveIndex: null,
}

const mutations = {
    [SET_VAULT_PAGE](state) {
        state.sidebarActiveIndex = '2'
    },
    [SET_PLAYGROUND_PAGE](state) {
        state.sidebarActiveIndex = '1'
    },
    [SET_COMPANY_PAGE](state) {
        state.sidebarActiveIndex = '3'
    },
    [SET_CREDIT_PAGE](state) {
        state.sidebarActiveIndex = '4'
    },
    [SET_SETTING_PAGE](state) {
        state.sidebarActiveIndex = '5'
    },
    [SET_APP_PAGE](state) {
      state.sidebarActiveIndex = '6'
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