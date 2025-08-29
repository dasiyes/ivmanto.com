import type { Component } from 'vue'

export type Article = {
  slug: string
  title: string
  summary: string
  date: string
  // The component is no longer async here because we are using an eager glob import.
  // This is beneficial for accessing metadata synchronously.
  component: Component
}

// Use Vite's glob import to get all article modules.
// `eager: true` imports the modules synchronously, which allows us to
// access their metadata immediately for building the article list.
const articleModules = import.meta.glob<// Type definition for the imported modules
{
  default: Component
  metadata: { title: string; summary: string; date: string }
}>('../components/articles/*.vue', { eager: true })

export const articles: Article[] = Object.entries(articleModules)
  .filter(([, module]) => {
    // Add a defensive check to ensure metadata exists. This prevents errors
    // if a .vue file is accidentally created in the folder without metadata.
    return module.metadata && module.metadata.title
  })
  .map(([path, module]) => {
    // Derive the slug from the file path, e.g., '../components/articles/dama-principles.vue' -> 'dama-principles'
    const slug = path.split('/').pop()?.replace('.vue', '') ?? 'unknown-slug'

    return {
      slug: slug,
      title: module.metadata.title,
      summary: module.metadata.summary,
      date: module.metadata.date,
      // The component is the default export from the .vue file
      component: module.default,
    }
  })

export function getArticleBySlug(slug: string): Article | undefined {
  return articles.find((article) => article.slug === slug)
}
