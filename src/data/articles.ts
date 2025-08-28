import { defineAsyncComponent, type Component } from 'vue'

export type Article = {
  slug: string
  title: string
  summary: string
  date: string
  component: Component
}

export const articles: Article[] = [
  {
    slug: 'common-architectural-patterns',
    title: 'Common Architectural Patterns for Modern Data Platforms',
    summary:
      'A deep dive into the pros and cons of Data Lakes, Warehouses, and the modern Lakehouse architecture on GCP.',
    date: '2024-07-20',
    // NOTE: You will need to create this component file.
    component: defineAsyncComponent(
      () => import('@/components/articles/CommonArchitecturalPatterns.vue'),
    ),
  },
  {
    slug: 'vertex-ai-production',
    title: 'From Notebook to Production with Vertex AI',
    summary:
      "A practical guide to operationalizing your machine learning models using Google Cloud's unified AI platform.",
    date: '2024-07-15',
    // NOTE: You will need to create this component file.
    component: defineAsyncComponent(() => import('@/components/articles/VertexAiProduction.vue')),
  },
  {
    slug: 'dama-principles',
    title: 'Why DAMA Principles Matter for Your Business',
    summary:
      'Exploring how standardized data management practices can reduce risk and increase the value of your data assets.',
    date: '2024-07-10',
    // NOTE: You will need to create this component file.
    component: defineAsyncComponent(() => import('@/components/articles/DamaPrinciples.vue')),
  },
]

export function getArticleBySlug(slug: string): Article | undefined {
  return articles.find((article) => article.slug === slug)
}
