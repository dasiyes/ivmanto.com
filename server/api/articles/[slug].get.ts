import type { Article } from '~/types/article'

/**
 * Server API route: /api/articles/:slug
 * Proxies individual article requests to the Go backend.
 * During `nuxi generate`, this runs server-side so blog pages can be pre-rendered.
 */
export default defineEventHandler(async (event) =>
{
    const slug = getRouterParam(event, 'slug')
    if (!slug)
    {
        throw createError({ statusCode: 400, statusMessage: 'Missing slug parameter' })
    }

    const config = useRuntimeConfig()
    const baseUrl = config.apiBaseUrl || 'https://ivmanto.com'

    try
    {
        const article = await $fetch<Article>(`${baseUrl}/api/articles/${slug}`)
        return article
    } catch (error: any)
    {
        if (error?.statusCode === 404 || error?.status === 404)
        {
            throw createError({ statusCode: 404, statusMessage: 'Article not found' })
        }
        throw createError({ statusCode: 500, statusMessage: 'Failed to fetch article' })
    }
})
