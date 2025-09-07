<template>
  <div class="bg-white px-6 py-24 sm:py-32 lg:px-8">
    <article v-if="article" class="mx-auto max-w-3xl text-base leading-7 text-gray-700">
      <p class="text-base font-semibold leading-7 text-primary">
        <RouterLink to="/blog" class="hover:underline">&larr; Back to all articles</RouterLink>
      </p>
      <h1 class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
        {{ article.title }}
      </h1>
      <div class="mt-10">
        <component :is="article.component" />
      </div>
    </article>
    <div v-else class="text-center">
      <h1 class="text-2xl font-bold">Article not found</h1>
      <p class="mt-4">The article you are looking for does not exist.</p>
      <RouterLink to="/blog" class="mt-6 inline-block text-primary hover:underline"
        >Go back to the blog</RouterLink
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { useHead } from '@vueuse/head'
import { getArticleBySlug, type Article } from '@/data/articles'

const props = defineProps<{
  slug: string
}>()

const route = useRoute()
const siteUrl = 'https://ivmanto.com'

const article = computed<Article | undefined>(() => {
  // The slug from the route is passed as a prop
  return getArticleBySlug(props.slug)
})

useHead({
  title: () => (article.value ? `${article.value.title} | ivmanto.com` : 'Blog Post'),
  meta: [
    {
      name: 'description',
      content: () => article.value?.summary || 'Read insights on data and AI.',
    },
  ],
  script: [
    {
      id: 'article-schema',
      type: 'application/ld+json',
      children: computed(() => {
        if (!article.value) {
          return ''
        }
        return JSON.stringify(
          {
            '@context': 'https://schema.org',
            '@type': 'BlogPosting',
            mainEntityOfPage: {
              '@type': 'WebPage',
              '@id': `${siteUrl}${route.path}`,
            },
            headline: article.value.title,
            description: article.value.summary,
            // image: `${siteUrl}/images/blog/${article.value.slug}.png`, // Recommended: Add an image for each article
            author: { '@id': `${siteUrl}/about#person` },
            publisher: { '@id': `${siteUrl}/#organization` },
            datePublished: article.value.date,
            // dateModified: article.value.dateModified || article.value.date // If you add a 'dateModified' property to your articles, include it here
          },
          null,
          2,
        )
      }),
      key: () => `article-schema-${props.slug}`,
    },
  ],
})
</script>
