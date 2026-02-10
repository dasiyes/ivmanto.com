<template>
  <div class="space-y-12 py-12 px-4 sm:px-6 lg:px-8">
    <header class="text-center">
      <h1 class="font-mono text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
        &lt;ivm <span class="text-red-500">/</span>&gt;
      </h1>
      <p class="mt-4 max-w-3xl mx-auto text-xl text-gray-600">
        Insights on data, architecture, and cloud technology.
      </p>
    </header>

    <!-- Search Bar -->
    <div class="max-w-2xl mx-auto">
      <input
        v-model="searchQuery"
        type="text"
        placeholder="Search articles by title or keyword..."
        class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg shadow-sm focus:ring-primary focus:border-primary transition"
      />
    </div>

    <main class="max-w-7xl mx-auto">
      <!-- Loading State -->
      <div v-if="isLoading" class="text-center py-16">
        <p class="text-xl text-gray-500">Loading articles...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-16">
        <h3 class="text-2xl font-semibold text-red-600">Failed to load articles</h3>
        <p class="text-gray-500 mt-2">{{ error }}</p>
      </div>

      <div v-else-if="filteredArticles.length > 0">
        <!-- Featured Articles -->
        <section v-if="featuredArticles.length > 0">
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
          <div class="grid md:grid-cols-2 gap-8">
            <RouterLink
              v-for="article in olderArticles"
              :key="article.slug"
              :to="`/blog/${article.slug}`"
              class="block p-6 bg-white rounded-lg border border-gray-200 hover:border-primary hover:shadow-md transition-all duration-200"
            >
              <p class="text-sm text-gray-500 mb-2">{{ formatDate(article.date) }}</p>
              <h3 class="mb-2 text-xl font-bold tracking-tight text-gray-900">
                {{ article.title }}
              </h3>
              <p class="font-normal text-gray-700 text-sm line-clamp-2">
                {{ article.summary }}
              </p>
            </RouterLink>
          </div>
        </section>
      </div>

      <!-- No Results Message -->
      <div v-else class="text-center py-16">
        <h3 class="text-2xl font-semibold text-gray-700">No articles found</h3>
        <p class="text-gray-500 mt-2">Try adjusting your search query.</p>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useArticles } from '@/composables/useArticles'
import { RouterLink } from 'vue-router'
import { computed, ref, onMounted } from 'vue'

const { sortedArticles, isLoading, error, fetchArticles } = useArticles()

onMounted(() => {
  fetchArticles()
})

const searchQuery = ref('')

// Filter articles based on search query
const filteredArticles = computed(() => {
  if (!searchQuery.value) {
    return sortedArticles.value
  }
  return sortedArticles.value.filter(
    (article) =>
      article.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      article.summary.toLowerCase().includes(searchQuery.value.toLowerCase()),
  )
})

const featuredArticles = computed(() => filteredArticles.value.slice(0, 3))
const olderArticles = computed(() => filteredArticles.value.slice(3))

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}
</script>
