<script setup lang="ts">
import { useHead } from '@vueuse/head'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

// Global Canonical Tag
// Forces all URLs to be treated as https://ivmanto.com/path
// Strips www, query params, and trailing slashes if needed (though Google usually handles trailing slashes fine)
useHead({
  script: [
    {
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'ProfessionalService',
        '@id': 'https://ivmanto.com/#organization',
        name: 'IVMANTO',
        url: 'https://ivmanto.com',
        logo: 'https://ivmanto.com/logo.png',
        image: 'https://ivmanto.com/social-sharing-card.webp',
        description:
          'Expert Cloud Data Architecture & AI Solutions built on Google Cloud Platform.',
        address: {
          '@type': 'PostalAddress',
          addressCountry: 'Germany',
        },
        founder: {
          '@id': 'https://ivmanto.com/about#person',
        },
      }),
    },
  ],
  link: [
    {
      rel: 'canonical',
      href: computed(() => {
        // Base URL is strictly non-www HTTPS
        const baseUrl = 'https://ivmanto.com'
        // Get the current path from the router, ensuring no trailing slash for consistency
        let path = route.path === '/' ? '' : route.path
        if (path.endsWith('/') && path.length > 1) {
          path = path.slice(0, -1)
        }
        // Construct the full canonical URL
        return `${baseUrl}${path}`
      }),
    },
  ],
})
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <router-view />
  </div>
</template>

<style>
/* Global Styles */
@import './style.css';
</style>
