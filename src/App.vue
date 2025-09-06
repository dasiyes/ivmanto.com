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
import Header from '@/components/layout/TheHeader.vue'
import Footer from '@/components/layout/TheFooter.vue'
import CookieConsentBanner from '@/components/CookieBanner.vue'
import { useHead } from '@vueuse/head'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const siteUrl = 'https://ivmanto.com'

// SEO metadata mapping for specific routes
const routeMetadata: Record<string, { title: string; description: string }> = {
  '/services/our-principles': {
    title: 'Guiding Principles | ivmanto.com',
    description:
      'Our DAMA-aligned principles for data strategy, governance, and architecture ensure your data becomes a reliable, valuable asset.',
  },
  '/services/data-architecture-on-gcp': {
    title: 'Data Architecture on GCP | ivmanto.com',
    description:
      'We design robust, scalable data architectures on Google Cloud (GCP) using BigQuery, GCS, and more, turning your AI strategy into reality.',
  },
  // We can add more routes here later
}

// Default metadata for other pages
const defaultTitle = 'ivmanto.com | Data & AI Consultancy'
const defaultDescription =
  'Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset.'

// Dynamically computed metadata based on the current route
const pageTitle = computed(() => routeMetadata[route.path]?.title ?? defaultTitle)
const pageDescription = computed(() => routeMetadata[route.path]?.description ?? defaultDescription)

useHead({
  title: pageTitle,
  meta: [
    { name: 'description', content: pageDescription },
    // Open Graph
    { property: 'og:type', content: 'website' },
    { property: 'og:title', content: pageTitle },
    { property: 'og:description', content: pageDescription },
    { property: 'og:url', content: computed(() => `${siteUrl}${route.path}`) },
    { property: 'og:image', content: `${siteUrl}/social-sharing-card.png` },
    // Twitter Card
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:title', content: pageTitle },
    { name: 'twitter:description', content: pageDescription },
    { name: 'twitter:image', content: `${siteUrl}/social-sharing-card.png` },
  ],
  link: [{ rel: 'canonical', href: computed(() => `${siteUrl}${route.path}`) }],
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
