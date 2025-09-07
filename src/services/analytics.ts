declare global {
  interface Window {
    dataLayer: Record<string, any>[]
  }
}

/**
 * Pushes a custom event to the Google Tag Manager data layer.
 * @param {string} eventName The name of the event.
 * @param {Record<string, any>} params Additional data to send with the event.
 */
export const trackEvent = (eventName: string, params: Record<string, any> = {}) => {
  window.dataLayer = window.dataLayer || []
  window.dataLayer.push({
    event: eventName,
    ...params,
  })
}

type GaSessionInfo = {
  clientId: string | null
  sessionId: string | null
}

/**
 * Retrieves the Google Analytics Client ID and Session ID from browser cookies.
 * This is crucial for stitching frontend user sessions with backend (server-to-server) events.
 * @param {string} measurementId The GA4 Measurement ID (e.g., 'G-XXXXXXXXXX').
 * @returns {GaSessionInfo} An object containing the clientId and sessionId.
 */
export const getGaSessionInfo = (measurementId: string): GaSessionInfo => {
  let clientId: string | null = null
  let sessionId: string | null = null
  const measurementCookieName = `_ga_${measurementId.replace('G-', '')}`

  const cookies = document.cookie.split(';')
  for (const cookie of cookies) {
    const cookieStr = cookie.trim()

    if (cookieStr.startsWith('_ga=')) {
      const value = cookieStr.substring(4) // length of '_ga='
      clientId = value.split('.').slice(2).join('.')
    }

    if (cookieStr.startsWith(measurementCookieName + '=')) {
      const value = cookieStr.substring(measurementCookieName.length + 1)
      const parts = value.split('.')
      if (parts.length >= 3) sessionId = parts[2]
    }
  }

  return { clientId, sessionId }
}
