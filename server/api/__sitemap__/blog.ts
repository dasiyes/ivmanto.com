import { defineSitemapEventHandler } from '#imports'

interface ArticleMeta {
  slug: string
  title: string
  summary: string
  date: string
  published: boolean
}

export default defineSitemapEventHandler(async () => {
  const backendUrl =
    process.env.BACKEND_URL || 'https://ivmanto.com'

  try {
    const articles = await $fetch<ArticleMeta[]>(`${backendUrl}/api/articles`)

    return articles.map((a) => ({
      loc: `/blog/${a.slug}`,
      lastmod: a.date,
      priority: 0.8,
    }))
  } catch (e) {
    // If the backend is unreachable at build time, return empty list.
    // The dynamic /api/sitemap-blog.xml from the Go backend covers this gap.
    console.warn('[sitemap] Failed to fetch articles from backend:', e)
    return []
  }
})
