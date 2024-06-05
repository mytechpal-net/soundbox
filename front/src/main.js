import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import './styles/app.css'
import App from './App.vue'

const pinia = createPinia()
const app = createApp(App)

/**
 * declare backend service url
 * Global properties can be called with this.propName in a vue context
 */
app.config.globalProperties.apiUrl = import.meta.env.VITE_APP_BACKEND_URL
app.config.globalProperties.appVersion = import.meta.env.VITE_APP_VERSION

app.use(router)
app.use(pinia)
app.mount('#app')
