import './style.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router' // This line must be correct

const COOKIE_CONSENT_KEY = 'cookie_consent'
const consent: string | null = localStorage.getItem(COOKIE_CONSENT_KEY)

if (consent === 'accepted') {
  // User has accepted cookies.
  // This is where you would initialize analytics, tracking scripts, etc.
  console.log('Cookie consent is "accepted". Initializing analytics...')
  // e.g., initializeGoogleAnalytics();
} else {
  // User has either declined or not made a choice yet.
  // Do not run any non-essential, cookie-setting scripts.
  console.log(`Cookie consent is "${consent}". Analytics will not be initialized.`)
}

const app = createApp(App)
app.use(router)

// Wait for the router to be ready before mounting the app
router.isReady().then(() => {
  app.mount('#app')
})
