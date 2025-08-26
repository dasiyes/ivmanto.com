import './style.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router' // This line must be correct

const app = createApp(App)
app.use(router)

// Wait for the router to be ready before mounting the app
router.isReady().then(() => {
  app.mount('#app')
})
