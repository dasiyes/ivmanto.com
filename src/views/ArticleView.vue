<template>
  <div class="bg-light-gray py-12 sm:py-16" @click="closeDropdown">
    <div v-if="article" class="mx-auto max-w-4xl px-4 sm:px-6 lg:px-8">
      <!-- Top Navigation Bar -->
      <div
        class="mb-8 p-2 bg-white rounded-lg shadow-sm flex items-center justify-between gap-2 text-sm font-medium"
      >
        <!-- Previous Article Button -->
        <RouterLink
          v-if="previousArticle"
          :to="`/blog/${previousArticle.slug}`"
          class="nav-button"
          :title="`Previous: ${previousArticle.title}`"
        >
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </RouterLink>
        <span v-else class="nav-button-disabled">
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </span>

        <!-- Custom Dropdown - "Google Translate" style -->
        <div class="relative flex-grow flex justify-center">
          <button @click.stop="toggleDropdown" class="nav-button-dropdown">
            <svg
              class="h-5 w-5"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1z"
                clip-rule="evenodd"
              />
            </svg>
            <span class="hidden sm:inline ml-2">Articles</span>
          </button>
          <transition
            enter-active-class="transition ease-out duration-100"
            enter-from-class="transform opacity-0 scale-95"
            enter-to-class="transform opacity-100 scale-100"
            leave-active-class="transition ease-in duration-75"
            leave-from-class="transform opacity-100 scale-100"
            leave-to-class="transform opacity-0 scale-95"
          >
            <div
              v-if="isDropdownOpen"
              class="absolute top-full mt-2 w-72 max-h-80 overflow-y-auto bg-white rounded-lg shadow-xl ring-1 ring-black ring-opacity-5 z-50"
            >
              <div class="py-1">
                <button
                  v-for="a in otherArticles"
                  :key="a.slug"
                  @click="navigateTo(a.slug)"
                  class="dropdown-item"
                >
                  {{ a.title }}
                </button>
              </div>
            </div>
          </transition>
        </div>

        <!-- Next Article Button -->
        <RouterLink
          v-if="nextArticle"
          :to="`/blog/${nextArticle.slug}`"
          class="nav-button"
          :title="`Next: ${nextArticle.title}`"
        >
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </RouterLink>
        <span v-else class="nav-button-disabled">
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </span>

        <!-- Share Button -->
        <div class="border-l border-gray-200 pl-2">
          <button @click="shareArticle" class="nav-button" title="Share article">
            <svg
              v-if="!copied"
              class="h-5 w-5"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z"
              />
            </svg>
            <svg
              v-else
              class="h-5 w-5 text-green-500"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clip-rule="evenodd"
              />
            </svg>
          </button>
        </div>
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
import { computed, ref, watch } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { articles, getArticleBySlug, type Article } from '@/data/articles'

const props = defineProps<{
  slug: string
}>()

const router = useRouter()

const isDropdownOpen = ref(false)
const copied = ref(false)

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

function navigateTo(slug: string) {
  router.push(`/blog/${slug}`)
}

function toggleDropdown() {
  isDropdownOpen.value = !isDropdownOpen.value
}

function closeDropdown() {
  isDropdownOpen.value = false
}

async function shareArticle() {
  if (!article.value) return

  const shareData = {
    title: article.value.title,
    text: article.value.summary,
    url: window.location.href,
  }

  if (navigator.share) {
    try {
      await navigator.share(shareData)
    } catch (err) {
      console.error('Error sharing:', err)
    }
  } else {
    // Fallback for browsers that don't support Web Share API
    await navigator.clipboard.writeText(window.location.href)
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  }
}

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

// Close dropdown on route change
watch(
  () => props.slug,
  () => {
    closeDropdown()
  },
)
</script>

<style scoped>
.nav-button {
  @apply h-10 w-10 flex items-center justify-center bg-gray-100 hover:bg-gray-200 text-gray-600 hover:text-primary rounded-md transition-all duration-200 hover:scale-110;
}
.nav-button-dropdown {
  @apply h-10 px-4 flex items-center justify-center bg-gray-100 hover:bg-gray-200 text-gray-600 hover:text-primary rounded-md transition-all duration-200;
}
.nav-button-disabled {
  @apply h-10 w-10 flex items-center justify-center bg-gray-100 text-gray-400 rounded-md cursor-not-allowed opacity-75;
}
.dropdown-item {
  @apply block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900;
}
</style>
