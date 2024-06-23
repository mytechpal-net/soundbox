import { ref } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios';

const apiProtocol = import.meta.env.VITE_APP_BACKEND_URL === 'prod' ? 'https' : 'http'
const apiUrl = `${apiProtocol}://${import.meta.env.VITE_APP_BACKEND_URL}`

export const userProfileStore = defineStore("sessionId", {
  state: () => ({
    userId: sessionStorage.getItem("userId"),
    sessionExpiration: sessionStorage.getItem("sessionExp"),
    soundBox: ref(null),
    isAdmin: false
  }),
  getters: {
    hasSession: (state) => state.sessionExpiration !== null && state.sessionExpiration > Date.now()
  },
  actions: {
    async createSession(loginData) {
      try {
        const { data } = await axios.post(apiUrl + "/login", loginData, { withCredentials: true })
        this.saveSession(data)
      } catch (error) {
        console.error("Unable to get a session")
        return error
      }
    },
    saveSession({ UserId, Role, SessionExp }) {
      this.userId = UserId
      this.sessionExpiration = SessionExp

      // save to session storage as well
      sessionStorage.setItem("userId", UserId)
      sessionStorage.setItem("sessionExp", SessionExp)
    },
    cleanSession() {
      this.saveSession({ UserId: null, SessionExp: null})

      // clean sessions
      this.soundBox = null
    },
    async logout() {
      try {
        await axios.get(apiUrl + "/app/user/logout", { withCredentials: true })
        this.cleanSession()
      } catch (error) {
        console.log('unable to logout')
        return error
      }
    }
  }
})