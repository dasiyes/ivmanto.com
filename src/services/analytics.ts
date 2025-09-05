/**
 * @file Centralized analytics service for tracking events with Google Tag Manager.
 */

/**
 * Defines the structure for the window.dataLayer object.
 */
declare global {
  interface Window {
    dataLayer: Array<Record<string, any>>
  }
}

/**
 * Pushes a custom event to the Google Tag Manager data layer.
 * @param eventName The name of the event (e.g., 'click_book_consultation').
 * @param params Additional data to send with the event.
 */
export const trackEvent = (eventName: string, params: Record<string, any> = {}): void => {
  window.dataLayer = window.dataLayer || []
  window.dataLayer.push({
    event: eventName,
    ...params,
  })
}
