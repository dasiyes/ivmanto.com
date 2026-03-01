// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,

  modules: ['@nuxtjs/tailwindcss'],

  css: ['~/assets/css/main.css'],

  runtimeConfig: {
    public: {
      geminiApiKey: '',
    },
  },

  nitro: {
    devProxy: {
      '/api': { target: 'http://localhost:8080', changeOrigin: true },
    },
  },

  routeRules: {
    // Pre-render static marketing pages
    '/': { prerender: true },
    '/about': { prerender: true },
    '/services': { prerender: true },
    '/services/**': { prerender: true },
    '/booking-demo': { prerender: true },
    '/login': { prerender: true },
    '/privacy-policy': { prerender: true },
    // Client-only for dynamic pages
    '/booking': { ssr: false },
    '/booking/cancel': { ssr: false },
    // Blog: client-rendered (articles come from runtime API)
    '/blog': { ssr: false },
    '/blog/**': { ssr: false },
  },

  app: {
    head: {
      htmlAttrs: { lang: 'en' },
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      link: [
        { rel: 'icon', href: '/favicon.ico' },
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        {
          rel: 'preconnect',
          href: 'https://fonts.gstatic.com',
          crossorigin: '',
        },
        {
          rel: 'preload',
          as: 'style',
          href: 'https://fonts.googleapis.com/css2?family=Montserrat:wght@400;700&display=swap',
        },
        {
          rel: 'stylesheet',
          href: 'https://fonts.googleapis.com/css2?family=Montserrat:wght@400;700&display=swap',
        },
      ],
    },
  },

  typescript: {
    strict: true,
  },

  compatibilityDate: '2025-02-28',
})
