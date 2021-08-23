const SET_USER = 'SET_USER'
const CLEAR_USER = 'CLEAR_USER'

const state = {
    name: '',
    certificte: '',
    real_name: '',
    card_no: '',
    phonenumber:'',
    delta_token:'',
    role:0
}

const mutations = {
    [SET_USER](state, payload) {
        Object.assign(state, payload)
    },
    [CLEAR_USER] () {
        Object.assign(state, {
            phonenumber: null,
            name: null,
        })
    }
}

export default {
    namespaced: true,
    state,
    mutations,
}