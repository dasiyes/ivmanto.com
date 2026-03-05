<script setup lang="ts">
const route = useRoute()

const siteTitle = 'ivmanto.com | Data & AI Consultancy'
const siteDescription =
  'Data & AI Consultancy on Google Cloud Platform. Expert data architecture, governance, and AI solutions that turn your data into a strategic asset.'
const siteUrl = 'https://ivmanto.com'
const ogImage = `${siteUrl}/social-sharing-card.webp`

useHead({
  titleTemplate: (titleChunk) => {
    if (!titleChunk) return siteTitle;
    return titleChunk.includes('ivmanto.com') ? titleChunk : `${titleChunk} | ivmanto.com`;
  }
})

useSeoMeta({
  title: siteTitle,
  description: siteDescription,
  author: 'Nikolay Tonev',
  ogType: 'website',
  ogSiteName: 'ivmanto.com',
  ogLocale: 'en_US',
  ogTitle: siteTitle,
  ogDescription: siteDescription,
  ogUrl: computed(() => {
    let path = route.path === '/' ? '' : route.path
    if (path.endsWith('/') && path.length > 1) path = path.slice(0, -1)
    return `${siteUrl}${path}`
  }),
  ogImage: ogImage,
  twitterCard: 'summary_large_image',
  twitterTitle: siteTitle,
  twitterDescription: siteDescription,
  twitterImage: ogImage,
})
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
