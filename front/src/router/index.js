import { createWebHistory, createRouter } from 'vue-router'

import Login from '@/pages/Login.vue'
import Default from '@/layouts/Default.vue'
import SoundBoxView from '@/views/SoundBoxView.vue'
import SettingsView from '@/views/SettingsView.vue'

import { userProfileStore } from '@/stores/userProfile'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  { 
    path: '/', 
    component: Default,
    redirect: '/home',
    children: [
      {
        path: 'home',
        component: SoundBoxView
      },
      {
        path: 'settings',
        component: SettingsView
      }
    ],  
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from) => {
  // replace it with a dedicated logic
  const userStore = userProfileStore()

  const isAuthenticated = userStore.hasSession
  if (
    // make sure the user is authenticated
    !isAuthenticated &&
    // ❗️ Avoid an infinite redirect
    to.name !== 'Login'
  ) {
    // redirect the user to the login page
    return { name: 'Login' }
  }
})

export default router