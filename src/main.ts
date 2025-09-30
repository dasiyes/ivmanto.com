// Initialize the dataLayer for Google Tag Manager/Analytics.
// This should be one of the very first things your app does to prevent race conditions.
window.dataLayer = window.dataLayer || []

import { createApp } from 'vue'
import { createHead } from '@vueuse/head'
import App from './App.vue'
import router from './router'
import { initGtm, trackEvent } from './services/analytics'
import './style.css'

const COOKIE_CONSENT_KEY = 'cookie_consent'
const consent: string | null = localStorage.getItem(COOKIE_CONSENT_KEY)

if (consent === 'accepted') {
  console.log('Cookie consent is "accepted". Initializing Google Tag Manager...')
  initGtm()

  // Track initial pageview
  trackEvent('page_view', {
    page_path: window.location.pathname + window.location.search,
    page_title: document.title,
  })

  // Track SPA pageviews on route change
  router.afterEach((to) => {
    trackEvent('page_view', {
      page_path: to.fullPath,
      page_title: document.title,
    })
  })
} else {
  // User has either declined or not made a choice yet.
  // Do not run any non-essential, cookie-setting scripts.
  console.log(`Cookie consent is "${consent}". Analytics will not be initialized.`)
}

const app = createApp(App)
const head = createHead()
app.use(router)
app.use(head)

app.mount('#app')
