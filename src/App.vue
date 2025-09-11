<template>
  <div id="app-container">
    <Header />
    <main>
      <router-view />
    </main>
    <Footer />
    <CookieConsentBanner />
  </div>
</template>

<script setup lang="ts">
import Header from '@/components/layout/AppHeader.vue'
import Footer from '@/components/layout/AppFooter.vue'
import CookieConsentBanner from '@/components/CookieBanner.vue'
import { trackEvent } from '@/services/analytics'
import { useHead } from '@vueuse/head'
import { computed, watch } from 'vue'
import { useRoute } from 'vue-router'
import { usePageMetadata } from '@/composables/usePageMetadata' // 2. Import your new composable

const route = useRoute()

const siteUrl = 'https://ivmanto.com'

// 3. Get the reactive title and description from your composable
const { pageTitle, pageDescription } = usePageMetadata()

// Watch for route changes to manually trigger a page_view event for GTM/GA4.
// While GA4's Enhanced Measurement for SPAs can be brittle, manual tracking is more robust.
// This ensures every route change is captured as a page view. The initial page view
// is handled by the GTM's GA4 config tag, so this watcher only handles subsequent navigations.
watch(
  () => route.path,
  (path) => {
    trackEvent('page_view', {
      // By passing the computed title directly, we avoid race conditions
      // with document.title updates.
      page_title: pageTitle.value,
      page_path: path,
    })
  },
)

useHead({
  title: pageTitle,
  meta: [
    { name: 'description', content: pageDescription },
    // Open Graph
    { property: 'og:type', content: 'website' },
    { property: 'og:title', content: pageTitle },
    { property: 'og:description', content: pageDescription },
    { property: 'og:url', content: computed(() => `${siteUrl}${route.path}`) },
    { property: 'og:image', content: `${siteUrl}/social-sharing-card.webp` },
    // Twitter Card
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:title', content: pageTitle },
    { name: 'twitter:description', content: pageDescription },
    { name: 'twitter:image', content: `${siteUrl}/social-sharing-card.webp` },
  ],
  link: [{ rel: 'canonical', href: computed(() => `${siteUrl}${route.path}`) }],
  script: [
    {
      id: 'organization-schema',
      type: 'application/ld+json',
      children: JSON.stringify(
        {
          '@context': 'https://schema.org',
          '@type': 'Organization',
          '@id': `${siteUrl}/#organization`,
          name: 'ivmanto.com | Data & AI Consultancy',
          url: siteUrl,
          logo: `${siteUrl}/logo.webp`, // Make sure you have a logo.png in your /public directory
          contactPoint: {
            '@type': 'ContactPoint',
            email: 'nikolay.tonev@ivmanto.com',
            contactType: 'customer service',
          },
          sameAs: ['https://linkedin.com/in/nikolaytonev', 'https://github.com/dasiyes'],
        },
        null,
        2,
      ),
    },
    {
      id: 'website-schema',
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'WebSite',
        url: siteUrl,
        potentialAction: {
          '@type': 'SearchAction',
          target: `${siteUrl}/search?q={search_term_string}`,
          'query-input': 'required name=search_term_string',
        },
      }),
    },
  ],
})
</script>

<style>
#app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

main {
  flex: 1;
}
</style>
