<template>
  <div class="py-12 sm:py-16">
    <div v-if="article" class="mx-auto max-w-4xl px-4 sm:px-6 lg:px-8">
      <!-- Top Navigation Bar -->
      <div
        class="mb-8 p-3 bg-white rounded-lg shadow-sm flex flex-wrap items-center justify-center gap-4 text-sm font-medium"
      >
        <!-- Previous Article Button -->
        <RouterLink v-if="previousArticle" :to="`/blog/${previousArticle.slug}`" class="nav-button"
          >Previous</RouterLink
        >
        <span v-else class="nav-button-disabled">Previous</span>

        <!-- Article Selector Dropdown -->
        <div class="order-last sm:order-none w-full sm:w-64">
          <select
            @change="navigateToArticle"
            class="w-full border-gray-300 rounded-md shadow-sm focus:ring-primary focus:border-primary"
          >
            <option value="">Jump to another article...</option>
            <option v-for="a in otherArticles" :key="a.slug" :value="a.slug">
              {{ a.title }}
            </option>
          </select>
        </div>

        <!-- Next Article Button -->
        <RouterLink v-if="nextArticle" :to="`/blog/${nextArticle.slug}`" class="nav-button"
          >Next</RouterLink
        >
        <span v-else class="nav-button-disabled">Next</span>
      </div>

      <!-- Article Content -->
      <div class="bg-white rounded-lg shadow-lg p-8 md:p-12">
        <div class="mb-8 border-b pb-8 border-gray-200">
          <p class="text-base text-gray-500">Published on {{ formatDate(article.date) }}</p>
          <h1 class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
            {{ article.title }}
          </h1>
          <p class="mt-6 text-xl leading-8 text-gray-700">{{ article.summary }}</p>
        </div>
        <div class="prose prose-lg max-w-none">
          <component :is="article.component" />
        </div>
      </div>
    </div>

    <!-- Article Not Found State -->
    <div v-else class="text-center py-16">
      <h1 class="text-3xl font-bold">Article not found</h1>
      <p class="mt-4">Sorry, we couldn't find the article you're looking for.</p>
      <RouterLink to="/blog" class="mt-6 inline-block text-primary hover:underline">
        &larr; Back to all articles
      </RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { articles, getArticleBySlug, type Article } from '@/data/articles'

const props = defineProps<{
  slug: string
}>()

const router = useRouter()

const article = computed<Article | undefined>(() => getArticleBySlug(props.slug))

// Sort articles by date (newest first) to determine previous/next
const sortedArticles = computed(() =>
  [...articles].sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime()),
)

const currentArticleIndex = computed(() => {
  if (!article.value) return -1
  return sortedArticles.value.findIndex((a) => a.slug === article.value?.slug)
})

const previousArticle = computed<Article | undefined>(() => {
  if (currentArticleIndex.value > 0) {
    return sortedArticles.value[currentArticleIndex.value - 1]
  }
  return undefined
})

const nextArticle = computed<Article | undefined>(() => {
  if (
    currentArticleIndex.value < sortedArticles.value.length - 1 &&
    currentArticleIndex.value !== -1
  ) {
    return sortedArticles.value[currentArticleIndex.value + 1]
  }
  return undefined
})

const otherArticles = computed(() => {
  if (!article.value) return []
  return sortedArticles.value.filter((a) => a.slug !== article.value?.slug)
})

function navigateToArticle(event: Event) {
  const slug = (event.target as HTMLSelectElement).value
  if (slug) {
    router.push(`/blog/${slug}`)
  }
}

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}
</script>

<style scoped>
.nav-button {
  @apply bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 px-4 rounded-md transition-colors w-28 text-center;
}
.nav-button-disabled {
  @apply bg-gray-100 text-gray-400 py-2 px-4 rounded-md cursor-not-allowed w-28 text-center;
}
</style>
