export default defineEventHandler((event) =>
{
    const host = getRequestHeader(event, 'host')

    if (host && host.startsWith('www.'))
    {
        const newHost = host.replace('www.', '')
        const url = getRequestURL(event)
        const redirectUrl = `https://${newHost}${url.pathname}${url.search}`

        return sendRedirect(event, redirectUrl, 301)
    }
})
