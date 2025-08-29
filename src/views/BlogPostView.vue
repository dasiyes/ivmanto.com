<template>
  <div class="bg-white px-6 py-24 sm:py-32 lg:px-8">
    <div v-if="article" class="mx-auto max-w-3xl text-base leading-7 text-gray-700">
      <p class="text-base font-semibold leading-7 text-primary">
        <RouterLink to="/blog" class="hover:underline">&larr; Back to all articles</RouterLink>
      </p>
      <h1 class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
        {{ article.title }}
      </h1>
      <div class="mt-10">
        <component :is="article.component" />
      </div>
    </div>
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
import { RouterLink } from 'vue-router'
import { getArticleBySlug, type Article } from '@/data/articles'

const props = defineProps<{
  slug: string
}>()

const article = computed<Article | undefined>(() => {
  // The slug from the route is passed as a prop
  return getArticleBySlug(props.slug)
})
</script>
