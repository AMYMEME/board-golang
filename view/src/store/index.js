import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const backEndHost = 'http://localhost:8090' //backend

const enhanceAccessToken = () => {
  const accessToken = localStorage.accessToken
  if (!accessToken) return
  axios.defaults.headers.common['Authorization'] = `Bearer ${accessToken}`;
}
enhanceAccessToken()

export default new Vuex.Store({
  state: {
    accessToken: localStorage.accessToken
  },
  getters: {
  },
  mutations: {
    login (state, accessToken) {
      state.accessToken = accessToken
      localStorage.accessToken = accessToken
      alert('로그인이 완료되었습니다')
    },
    logout (state) {
      state.accessToken = null
      delete localStorage.accessToken
      alert('로그아웃이 완료되었습니다')
    },
    backRedirect() {
      var params = {};
      window.location.search.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(str, key, value) { params[key] = value; });
      var redirectURI = decodeURIComponent(params.redirect);
      if (redirectURI === 'undefined') {
          location.replace('/');
      } else {
          location.replace(redirectURI);
      }
    },
  },
  actions: {
      loginWithGoogle ({commit}, result) {
          return axios.post(`${backEndHost}/auth/naver`, {
              name: result.givenName,
              uniq_id: result.metadata.source.id
          })
              .then((res) => {
                  axios.defaults.headers.common['Authorization'] = `Bearer ${res.data}`;
                  commit('login', res.data)
              })
              .catch(err => {
                  alert('에러가 발생했습니다: ' + err + '\n다시 시도해 주세요.')
              })
      },
      loginWithNaver ({commit}, result) {
          return axios.post(`${backEndHost}/auth/naver`, {
              name: result.name,
              uniq_id: result.uniqId
          })
              .then((res) => {
                  axios.defaults.headers.common['Authorization'] = `Bearer ${res.data}`;
                  commit('login', res.data)
              })
              .catch(err => {
                  alert('에러가 발생했습니다: ' + err + '\n다시 시도해 주세요.')
              })
      },
      loginWithKakao ({commit}, result) {
          return axios.post(`${backEndHost}/auth/kakao`, {
              name: result.kakao_account.profile.nickname,
              uniq_id: result.id.toString()
          })
              .then((res) => {
                  axios.defaults.headers.common['Authorization'] = `Bearer ${res.data}`;
                  commit('login', res.data)
              })
              .catch(err => {
                  alert('에러가 발생했습니다: ' + err + '\n다시 시도해 주세요.')
              })
      },
      logout ({commit}) {
          axios.defaults.headers.common['Authorization'] = undefined
          commit('logout')
      },
  }
})