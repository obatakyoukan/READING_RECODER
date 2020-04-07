import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import { UPDATE_CURRENT, UPDATE_BOOKS, UPDATE_BOOK, DELETE_BOOK } from './mutation-types'

Vue.use(Vuex)

export default new Vuex.Store({
  strict: true,
  state: {
   books: [],
   current: null
  },
  getters: {
   bookCount( state ) {
    return state.books.length
   },
   allBooks( state ) {
    return state.books
   },
   getRangeByPage( state ) {
    return page => {
     const SIZE = 3
     return state.books.slice( (page-1)*SIZE , (page-1)*SIZE+SIZE )
    }
   },
   getBookById( state ) {
    return id => {
     return state.books.find( book => book.id == id )
    }
   },
   current( state ) {
    return state.current;
   }
  },
  mutations: {
   [ UPDATE_CURRENT ]( state , payload ) {
    state.current = payload
   },
   //追加
   [ UPDATE_BOOKS ]( state , payload ) {
    state.books = []
    state.books = payload
   },
   [ UPDATE_BOOK ]( state , payload ){
    let b = this.getters.getBookById( payload.id )
    if (b) {
     Object.assign(b, payload)
    } else {
     state.books.push(payload)
    }
   },
   [ DELETE_BOOK ]( state , payload ){
    let b = this.getters.getBookById( payload.id )
    if (b) {
     state.books = state.books.filter( book => book.id !== b.id )
    }
   }
  },
  actions: {
   [UPDATE_CURRENT]( { commit }, payload ) {
    commit( UPDATE_CURRENT, payload)
   },
   //追加
   [UPDATE_BOOKS]( {commit}, payload){
    commit( UPDATE_BOOKS , payload )
   },
   [UPDATE_BOOK]( {commit} , payload){
    commit( UPDATE_BOOK , payload )
   },
   [DELETE_BOOK]( {commit} , payload){
    commit( DELETE_BOOK , payload )
   }
  },
  plugins: [createPersistedState({
   key: 'reading-recorder',
   storage: localStorage
  })],
  modules: {
  }
})
