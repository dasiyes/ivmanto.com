/**
 * Analytics Client Plugin
 *
 * Handles SPA-style pageview tracking on route changes.
 * GTM is initialized by the CookieBanner component after cookie consent.
 * This plugin fires virtual pageview events on each client-side navigation
 * so GA4 (via GTM) tracks them correctly.
 */
import { trackEvent } from '~/services/analytics'

export default defineNuxtPlugin(() => {
  const router = useRouter()

  router.afterEach((to) => {
    // Push a virtual pageview event to the dataLayer on every route change.
    // GTM should be configured to fire a GA4 page_view tag on this event.
    trackEvent('page_view', {
      page_path: to.fullPath,
      page_title: document.title,
    })
  })
})
