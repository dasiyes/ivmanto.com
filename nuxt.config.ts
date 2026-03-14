// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,

  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/sitemap', '@nuxtjs/google-fonts', '@nuxt/image'],

  googleFonts: {
    families: {
      Inter: '300..800',
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
    exclude: ['/login', '/booking-demo', '/booking', '/booking/cancel'],
  },

  css: ['~/assets/css/main.css'],

  runtimeConfig: {
    apiBaseUrl: process.env.NUXT_API_BASE_URL || 'https://ivmanto.com',
    public: {
      geminiApiKey: '',
    },
  },

  nitro: {
    devProxy: {
      '/api': { target: 'http://localhost:8080', changeOrigin: true },
    },
    prerender: {
      // Crawl links to discover blog pages from the sitemap source
      crawlLinks: true,
    },
  },

  features: {
    inlineStyles: true
  },

  routeRules: {
    // Global Security Headers
    '/**': {
      headers: {
        'Strict-Transport-Security': 'max-age=31536000; includeSubDomains; preload',
      },
    },
    // Pre-render static marketing pages
    '/': { prerender: true },
    '/about': { prerender: true },
    '/services': { prerender: true },
    '/services/**': { prerender: true },
    '/booking-demo': { prerender: true },
    '/login': { prerender: true },
    '/privacy-policy': { prerender: true },
    '/terms': { prerender: true },
    // Client-only for dynamic pages
    '/booking': { prerender: true },
    '/booking/cancel': { prerender: true },
    '/assessment': { prerender: true },
    // Blog: pre-rendered at generate time for SEO
    '/blog': { prerender: true },
    '/blog/**': { prerender: true },
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

  hooks: {
    async 'prerender:routes'(ctx)
    {
      const backendUrl =
        process.env.NUXT_API_BASE_URL || 'https://ivmanto.com'
      try
      {
        const articles = await (await fetch(`${backendUrl}/api/articles`)).json()
        for (const article of articles)
        {
          ctx.routes.add(`/blog/${article.slug}`)
        }
        console.log(`[prerender] Discovered ${articles.length} blog routes`)
      } catch (e)
      {
        console.warn('[prerender] Could not fetch blog routes:', e)
      }
    },
  },

  compatibilityDate: '2025-02-28',
})
