<script setup lang="ts">
// Global Schema (JSON-LD structured data)
useHead({
  script: [
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': ['Organization', 'ProfessionalService'],
        '@id': 'https://ivmanto.com/#organization',
        name: 'IVMANTO',
        alternateName: 'IVMANTO - Nikolay Tonev Data & Cloud Solutions',
        url: 'https://ivmanto.com',
        logo: 'https://ivmanto.com/logo.png',
        image: 'https://ivmanto.com/social-sharing-card.webp',
        description:
          'Expert Cloud Data Architecture & AI Solutions built on Google Cloud Platform.',
        address: {
          '@type': 'PostalAddress',
          addressCountry: 'DE',
        },
        founder: {
          '@id': 'https://ivmanto.com/about#person',
        },
        sameAs: [
          'https://linkedin.com/in/nikolaytonev',
          'https://github.com/dasiyes',
          'https://g.dev/ivmanto-nikolaytonev',
        ],
      }),
    },
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
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
})
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <LayoutAppHeader />
    <main class="flex-1">
      <slot />
    </main>
    <LayoutTheFooter />
    <CookieBanner />
  </div>
</template>
