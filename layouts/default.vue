<script setup lang="ts">
const route = useRoute()

// Global SEO and Schema
useHead({
  title: 'ivmanto.com | Data & AI Consultancy',
  meta: [
    {
      name: 'description',
      content:
        'Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset.',
    },
  ],
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
    {
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'WebSite',
        name: 'IVMANTO',
        url: 'https://ivmanto.com',
        potentialAction: {
          '@type': 'SearchAction',
          target: 'https://ivmanto.com/blog?q={search_term_string}',
          'query-input': 'required name=search_term_string',
        },
      }),
    },
  ],
  link: [
    {
      rel: 'canonical',
      href: computed(() => {
        const baseUrl = 'https://ivmanto.com'
        let path = route.path === '/' ? '' : route.path
        if (path.endsWith('/') && path.length > 1) {
          path = path.slice(0, -1)
        }
        return `${baseUrl}${path}`
      }),
    },
  ],
})
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <LayoutAppHeader />
    <slot />
    <LayoutTheFooter />
    <CookieBanner />
  </div>
</template>
