export interface ArticleMeta {
  slug: string
  title: string
  summary: string
  date: string
  published: boolean
}

export interface Article extends ArticleMeta {
  content: string
}
