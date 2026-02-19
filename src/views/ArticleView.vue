<template>
  <div class="bg-light-gray py-12 sm:py-16" @click="closeDropdown">
    <!-- Loading State -->
    <div v-if="isLoadingContent" class="text-center py-16">
      <div
        class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-current border-r-transparent align-[-0.125em] motion-reduce:animate-[spin_1.5s_linear_infinite]"
        role="status"
      >
        <span class="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
          >Loading...</span
        >
      </div>
    </div>

    <!-- Article Content -->
    <div v-else-if="article" class="mx-auto max-w-4xl px-4 sm:px-6 lg:px-8">
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
        <div class="prose prose-lg max-w-none" v-html="sanitizedContent"></div>

        <!-- Like button section -->
        <div class="mt-12 pt-8 border-t border-gray-200 flex justify-end items-center">
          <button
            @click="handleLike"
            :disabled="isLoadingLikes"
            class="flex items-center gap-2 text-gray-500 hover:text-primary disabled:text-primary disabled:cursor-not-allowed transition-colors"
            title="Like this article"
          >
            <svg
              class="h-6 w-6"
              :class="{ 'text-red-500 fill-current': isLiked }"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4.318 6.318a4.5 4.5 0 016.364 0L12 7.5l1.318-1.182a4.5 4.5 0 116.364 6.364L12 20.364l-7.682-7.682a4.5 4.5 0 010-6.364z"
              />
            </svg>
            <span v-if="!isLoadingLikes" class="font-semibold text-lg tabular-nums">{{
              likeCount
            }}</span>
            <span v-else class="font-semibold text-lg">...</span>
          </button>
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
import { computed, ref, watch, onMounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { trackEvent } from '@/services/analytics'
import { useArticles } from '@/composables/useArticles'
import DOMPurify from 'dompurify'
import { useHead } from '@vueuse/head'
import type { Article } from '@/types/article'

const props = defineProps<{
  slug: string
}>()

const router = useRouter()
const { sortedArticles, fetchArticles, fetchArticle } = useArticles()

const isDropdownOpen = ref(false)
const copied = ref(false)
const likeCount = ref(0)
const isLiked = ref(false)
const isLoadingLikes = ref(true)
const isLoadingContent = ref(true)

const article = ref<Article | null>(null)

// Head Management for SEO (Fix Soft 404)
useHead({
  meta: computed(() => {
    // If not loading AND article is null => NoIndex
    if (!isLoadingContent.value && !article.value) {
      return [{ name: 'robots', content: 'noindex' }]
    }
    return []
  }),
})

const sanitizedContent = computed(() => {
  if (!article.value?.content) return ''
  return DOMPurify.sanitize(article.value.content)
})

const currentArticleIndex = computed(() => {
  if (!article.value) return -1
  return sortedArticles.value.findIndex((a) => a.slug === article.value?.slug)
})

const previousArticle = computed(() => {
  if (currentArticleIndex.value > 0) {
    return sortedArticles.value[currentArticleIndex.value - 1]
  }
  return undefined
})

const nextArticle = computed(() => {
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

async function fetchLikes() {
  if (!article.value) return
  isLoadingLikes.value = true
  try {
    // This endpoint will need to be created in the backend.
    const response = await fetch(`/api/articles/${article.value.slug}/likes`)
    if (response.ok) {
      const data = await response.json()
      likeCount.value = data.likes || 0
    } else {
      console.warn(`Could not fetch likes for ${article.value.slug}. Defaulting to 0.`)
      likeCount.value = 0
    }
  } catch (error) {
    console.error('Error fetching likes:', error)
    likeCount.value = 0 // Default to 0 on error
  } finally {
    isLoadingLikes.value = false
  }
}

async function handleLike() {
  if (!article.value) return

  // Toggle the like state locally first for an instant UI response.
  const wasLiked = isLiked.value
  if (wasLiked) {
    isLiked.value = false
    likeCount.value--
    localStorage.removeItem(`liked-${article.value.slug}`)
    // Track the 'unlike' action for more detailed engagement analysis.
    trackEvent('like_insight', {
      insight_id: article.value.slug,
      insight_title: article.value.title,
      like_status: 'unliked', // Add status parameter
    })
  } else {
    isLiked.value = true
    likeCount.value++
    localStorage.setItem(`liked-${article.value.slug}`, 'true')

    // Track the 'like' event for analytics, as per our plan.
    trackEvent('like_insight', {
      insight_id: article.value.slug,
      insight_title: article.value.title,
      like_status: 'liked', // Add status parameter
    })
  }

  // Then, attempt to sync this change with the backend.
  // The UI will not be reverted if this call fails.
  try {
    // Use DELETE for un-liking and POST for liking.
    const method = wasLiked ? 'DELETE' : 'POST'
    const response = await fetch(`/api/articles/${article.value.slug}/like`, { method })

    if (response.ok) {
      // If the backend call succeeds, re-sync the count with the server's "source of truth".
      const data = await response.json()
      likeCount.value = data.likes
    } else {
      console.error(`Failed to sync like status. API responded with status: ${response.status}`)
    }
  } catch (error) {
    // Also log network errors without reverting the UI state.
    console.error('Network error while syncing like status:', error)
  }
}

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

async function loadArticle(slug: string) {
  isLoadingContent.value = true
  article.value = await fetchArticle(slug)
  isLoadingContent.value = false
}

onMounted(async () => {
  await fetchArticles()
  await loadArticle(props.slug)
  fetchLikes()
  if (localStorage.getItem(`liked-${props.slug}`)) {
    isLiked.value = true
  }
})

// Close dropdown on route change and load new article
watch(
  () => props.slug,
  async (newSlug) => {
    closeDropdown()
    likeCount.value = 0
    isLiked.value = false
    isLoadingLikes.value = true
    await loadArticle(newSlug)
    if (localStorage.getItem(`liked-${newSlug}`)) {
      isLiked.value = true
    }
    fetchLikes()
  },
)
</script>

<style lang="postcss" scoped>
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
.tabular-nums {
  font-variant-numeric: tabular-nums;
}
</style>
