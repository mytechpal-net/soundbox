import { defineStore } from 'pinia'

export const userProfileStore = defineStore("sessionId", {
  state: () => ({
    userId: sessionStorage.getItem("userId"),
    sessionExpiration: sessionStorage.getItem("sessionExp")
  }),
  getters: {
    hasSession: (state) => state.sessionExpiration !== null && state.sessionExpiration > Date.now()
  },
  actions: {
    saveSession({ UserId, SessionExp }) {
      this.userId = UserId
      this.sessionExpiration = SessionExp

      // save to session storage as well
      sessionStorage.setItem("userId", UserId)
      sessionStorage.setItem("sessionExp", SessionExp)
    }
  }
})