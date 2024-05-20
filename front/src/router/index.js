import { createWebHistory, createRouter } from 'vue-router'

import Login from '@/pages/Login.vue'
import Default from '@/layouts/Default.vue'
import HomeView from '@/views/HomeView.vue'
import SoundBoxView from '@/views/SoundBoxView.vue'
import SettingsView from '@/views/SettingsView.vue'

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
        component: HomeView
      },
      {
        path: 'soundbox',
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
  const isAuthenticated = false
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