<template>
  <div v-if="article" class="py-8">
    <!-- The dynamic component for the article content will be rendered here -->
    <component :is="article.component" />
  </div>
  <div v-else class="text-center py-16">
    <h1 class="text-3xl font-bold">Article not found</h1>
    <p class="mt-4">Sorry, we couldn't find the article you're looking for.</p>
    <RouterLink to="/blog" class="mt-6 inline-block text-blue-600 hover:underline">
      &larr; Back to all articles
    </RouterLink>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { getArticleBySlug, type Article } from '@/data/articles'
import { RouterLink } from 'vue-router'

const props = defineProps<{
  slug: string
}>()

const article = computed<Article | undefined>(() => {
  // The slug is passed as a prop from the router
  return getArticleBySlug(props.slug)
})
</script>
