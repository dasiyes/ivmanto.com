import type { ArticleMeta } from '~/types/article'

/**
 * Server API route: /api/articles
 * Proxies article-list requests to the Go backend.
 * During `nuxi generate`, this runs server-side so blog pages can be pre-rendered.
 */
export default defineEventHandler(async (event) =>
{
    const config = useRuntimeConfig()
    const baseUrl = config.apiBaseUrl || 'https://ivmanto.com'

    const articles = await $fetch<ArticleMeta[]>(`${baseUrl}/api/articles`)
    return articles
})
