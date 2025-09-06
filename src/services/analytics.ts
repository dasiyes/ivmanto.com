// Augment the Window interface to include dataLayer for Google Tag Manager.
declare global {
  interface Window {
    dataLayer: any[]
  }
}

const GTM_ID = 'GTM-PXRF8FPQ'

let isGtmInitialized = false

/**
 * Initializes the Google Tag Manager script.
 * This function should only be called once the user has given cookie consent.
 */
export const initGtm = () => {
  if (isGtmInitialized || !GTM_ID) {
    return
  }

  // Standard GTM script injection
  ;(function (w: Window, d: Document, s: string, l: 'dataLayer', i: string) {
    w[l] = w[l] || []
    w[l].push({ 'gtm.start': new Date().getTime(), event: 'gtm.js' })
    const f = d.getElementsByTagName(s)[0]
    const j = d.createElement(s) as HTMLScriptElement
    const dl = l !== 'dataLayer' ? '&l=' + l : ''
    j.async = true
    j.src = 'https://www.googletagmanager.com/gtm.js?id=' + i + dl
    f.parentNode?.insertBefore(j, f)
  })(window, document, 'script', 'dataLayer', GTM_ID)

  isGtmInitialized = true
  console.log('Google Tag Manager initialized.')
}

export const trackEvent = (eventName: string, params: Record<string, any> = {}) => {
  if (typeof window.dataLayer === 'undefined') {
    console.warn('DataLayer not available. Event not tracked:', eventName, params)
    return
  }
  window.dataLayer.push({
    event: eventName,
    ...params,
  })
}
