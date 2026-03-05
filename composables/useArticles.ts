import type { ArticleMeta, Article } from '~/types/article'

/**
 * Composable for fetching and managing articles.
 *
 * Uses Nuxt's $fetch (universal) so it works both server-side (during nuxi generate)
 * and client-side (in the browser). This is critical for SSG pre-rendering of blog pages.
 */
export function useArticles()
{
  const articles = useState<ArticleMeta[]>('articles', () => [])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  async function fetchArticles(): Promise<ArticleMeta[]>
  {
    if (articles.value.length > 0) return articles.value
    isLoading.value = true
    error.value = null
    try
    {
      const data = await $fetch<ArticleMeta[]>('/api/articles')
      articles.value = data
      return data
    } catch (e)
    {
      error.value = e instanceof Error ? e.message : 'Unknown error'
      return []
    } finally
    {
      isLoading.value = false
    }
  }

  async function fetchArticle(slug: string): Promise<Article | null>
  {
    try
    {
      const data = await $fetch<Article>(`/api/articles/${slug}`)
      return data
    } catch (e: any)
    {
      if (e?.statusCode === 404 || e?.status === 404) return null
      error.value = e instanceof Error ? e.message : 'Unknown error'
      return null
    }
  }

  const sortedArticles = computed(() =>
    [...articles.value].sort(
      (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime(),
    ),
  )

  function getArticleBySlug(slug: string): ArticleMeta | undefined
  {
    return articles.value.find((a) => a.slug === slug)
  }

  return {
    articles,
    sortedArticles,
    isLoading,
    error,
    fetchArticles,
    fetchArticle,
    getArticleBySlug,
  }
}
