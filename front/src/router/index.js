import { createWebHistory, createRouter } from 'vue-router'

import Login from '@/pages/Login.vue'
import Default from '@/layouts/Default.vue'
import SoundBoxView from '@/views/SoundBoxView.vue'
import ConfigView from '@/views/ConfigView.vue'
import UserView from '@/views/UserView.vue'

import { userProfileStore } from '@/stores/userProfile'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
    beforeEnter: (to, from) => {
      // https://pinia.vuejs.org/core-concepts/outside-component-usage.html
      const userStore = userProfileStore()
      if (userStore.hasSession) {
        return { path: '/app' }
      }
    }
  },
  { 
    path: '/', 
    component: Default,
    redirect: '/app',
    children: [
      {
        path: 'app',
        component: SoundBoxView
      },
      {
        path: 'config',
        component: ConfigView
      },
    ],  
  },
  {
    path: '/user',
    component: Default,
    children: [
      {
        path: '',
        component: UserView
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from) => {
  // https://pinia.vuejs.org/core-concepts/outside-component-usage.html
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