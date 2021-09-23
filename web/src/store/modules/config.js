const SET_CONFIG = 'SET_CONFIG'

const state = {
   config:{}
}

const mutations = {
    [SET_CONFIG](state, payload) {
      state.config = payload
    }
}

export default {
    namespaced: true,
    state,
    mutations,
}