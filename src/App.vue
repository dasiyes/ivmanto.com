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

const route = useRoute()

const siteUrl = 'https://ivmanto.com'

//   '/services/go-backend-development': {
//   title: 'Go Backend Development | ivmanto.com',
//   description:
//     'High-performance Go (Golang) backend development for data-intensive applications. We build scalable, concurrent, and efficient cloud-native services.',
// },

// SEO metadata mapping for specific routes
const routeMetadata: Record<string, { title: string; description: string }> = {
  '/': {
    title: 'ivmanto.com | Data & AI Consultancy',
    description:
      'Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset.',
  },
  '/services': {
    title: 'Services | ivmanto.com',
    description:
      'Explore our Data & AI services. From data strategy and GCP architecture to custom AI/ML solutions and Go backend development, we empower your business with data.',
  },
  '/services/data-strategy-and-governance': {
    title: 'Data Strategy & Governance | ivmanto.com',
    description:
      'Develop a clear data strategy and robust governance framework. We align your data initiatives with business goals for maximum impact and compliance.',
  },
  '/services/data-architecture': {
    title: 'Data Architecture on GCP | ivmanto.com',
    description:
      'Design and build scalable, secure data architectures on Google Cloud Platform (GCP). We leverage BigQuery, Cloud Storage, and modern data engineering practices.',
  },
  '/services/ml-engineering': {
    title: 'AI & ML Solutions | ivmanto.com',
    description:
      'Leverage the power of AI and Machine Learning on GCP. We build custom solutions, from predictive analytics to generative AI, to solve your toughest challenges.',
  },
  '/services/principles': {
    title: 'Guiding Principles | ivmanto.com',
    description:
      'Our DAMA-aligned principles for data strategy, governance, and architecture ensure your data becomes a reliable, valuable asset for decision-making and AI.',
  },
  '/blog': {
    title: 'Insights & Articles | ivmanto.com',
    description:
      'Read our latest articles and insights on data strategy, cloud architecture, AI/ML, and software engineering. Stay ahead of the curve with expert analysis.',
  },
  '/about': {
    title: 'About | ivmanto.com',
    description:
      'Learn about IVMANTO and our mission to help businesses harness the power of data. Meet the experts behind our innovative data and AI solutions.',
  },
  '/booking': {
    title: 'Contact us | ivmanto.com',
    description:
      'Get in touch with IVMANTO to discuss your data and AI challenges. Book a free consultation or send us a message to start your data transformation journey.',
  },
  '/privacy-policy': {
    title: 'Privacy Policy | ivmanto.com',
    description:
      'Read the IVMANTO Privacy Policy to understand how we collect, use, and protect your personal data in accordance with GDPR and other regulations.',
  },
}

// Default metadata for other pages
const defaultTitle = 'ivmanto.com | Data & AI Consultancy'
const defaultDescription =
  'Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset.'

// Dynamically computed metadata based on the current route
const pageTitle = computed(() => routeMetadata[route.path]?.title ?? defaultTitle)
const pageDescription = computed(() => routeMetadata[route.path]?.description ?? defaultDescription)

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
