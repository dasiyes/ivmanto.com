import { ref, computed } from 'vue'
import type { ArticleMeta, Article } from '@/types/article'

const articlesCache = ref<ArticleMeta[]>([])
const isLoading = ref(false)
const error = ref<string | null>(null)
let hasFetched = false

export function useArticles() {
  async function fetchArticles(): Promise<void> {
    if (hasFetched && articlesCache.value.length > 0) return
    isLoading.value = true
    error.value = null
    try {
      const res = await fetch('/api/articles')
      if (!res.ok) throw new Error(`Failed to fetch articles: ${res.status}`)
      articlesCache.value = await res.json()
      hasFetched = true
    } catch (e) {
      error.value = e instanceof Error ? e.message : 'Unknown error'
    } finally {
      isLoading.value = false
    }
  }

  async function fetchArticle(slug: string): Promise<Article | null> {
    try {
      const res = await fetch(`/api/articles/${slug}`)
      if (!res.ok) {
        if (res.status === 404) return null
        throw new Error(`Failed to fetch article: ${res.status}`)
      }
      return await res.json()
    } catch (e) {
      error.value = e instanceof Error ? e.message : 'Unknown error'
      return null
    }
  }

  const sortedArticles = computed(() =>
    [...articlesCache.value].sort(
      (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime(),
    ),
  )

  function getArticleBySlug(slug: string): ArticleMeta | undefined {
    return articlesCache.value.find((a) => a.slug === slug)
  }

  return {
    articles: articlesCache,
    sortedArticles,
    isLoading,
    error,
    fetchArticles,
    fetchArticle,
    getArticleBySlug,
  }
}
