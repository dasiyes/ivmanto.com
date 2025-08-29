<template>
  <div class="space-y-16 py-8 px-4 sm:px-6 lg:px-8">
    <header class="text-center">
      <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">From the Blog</h1>
      <p class="mt-4 max-w-2xl mx-auto text-xl text-gray-600">
        Insights on data, architecture, and cloud technology.
      </p>
    </header>

    <main class="max-w-7xl mx-auto">
      <!-- Featured Articles -->
      <section>
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 mb-8">Featured Articles</h2>
        <div class="grid gap-8 lg:grid-cols-3">
          <RouterLink
            v-for="article in featuredArticles"
            :key="article.slug"
            :to="`/blog/${article.slug}`"
            class="block p-6 bg-white rounded-lg border border-gray-200 shadow-md hover:shadow-xl hover:-translate-y-1 transition-all duration-300"
          >
            <p class="text-sm text-gray-500 mb-2">{{ formatDate(article.date) }}</p>
            <h3 class="mb-2 text-2xl font-bold tracking-tight text-gray-900">
              {{ article.title }}
            </h3>
            <p class="font-normal text-gray-700">{{ article.summary }}</p>
          </RouterLink>
        </div>
      </section>

      <!-- All Other Articles -->
      <section v-if="olderArticles.length > 0" class="mt-16">
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 mb-8">More Articles</h2>
        <div class="space-y-4">
          <RouterLink
            v-for="article in olderArticles"
            :key="article.slug"
            :to="`/blog/${article.slug}`"
            class="block p-6 bg-white rounded-lg border border-gray-200 hover:border-blue-500 transition-colors duration-200"
          >
            <div class="flex justify-between items-baseline">
              <h3 class="text-xl font-semibold text-gray-800">{{ article.title }}</h3>
              <p class="text-sm text-gray-500 flex-shrink-0 ml-4">
                {{ formatDate(article.date) }}
              </p>
            </div>
          </RouterLink>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { articles } from '@/data/articles'
import { RouterLink } from 'vue-router'
import { computed } from 'vue'

// Sort articles by date, newest first
const sortedArticles = computed(() =>
  [...articles].sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime()),
)

const featuredArticles = computed(() => sortedArticles.value.slice(0, 3))
const olderArticles = computed(() => sortedArticles.value.slice(3))

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}
</script>
