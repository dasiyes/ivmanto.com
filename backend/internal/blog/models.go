package blog

import "time"

// ArticleMeta holds metadata for a single blog article.
// Fields are parsed from YAML frontmatter in .md files.
type ArticleMeta struct {
	Slug      string `json:"slug"      yaml:"-"`
	Title     string `json:"title"     yaml:"title"`
	Summary   string `json:"summary"   yaml:"summary"`
	Date      string `json:"date"      yaml:"date"`
	Published bool   `json:"published" yaml:"published"`
}

// Article is the full article with rendered HTML content.
type Article struct {
	ArticleMeta
	Content string `json:"content"`
}

// MetadataCache is the structure persisted to metadata.json in GCS.
type MetadataCache struct {
	GeneratedAt time.Time     `json:"generated_at"`
	Articles    []ArticleMeta `json:"articles"`
}
