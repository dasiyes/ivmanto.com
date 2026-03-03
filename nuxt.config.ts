// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,

  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/sitemap', '@nuxtjs/google-fonts'],

  googleFonts: {
    families: {
      Montserrat: [400, 700],
    },
    display: 'swap',
    download: true,
    preload: true,
  },

  site: {
    url: 'https://ivmanto.com',
  },

  sitemap: {
    sources: ['/api/__sitemap__/blog'],
    exclude: ['/login', '/booking-demo'],
  },

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
      ],
    },
  },

  typescript: {
    strict: true,
  },

  compatibilityDate: '2025-02-28',
})
