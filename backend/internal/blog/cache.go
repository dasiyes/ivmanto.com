package blog

import (
	"context"
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"sync"
	"time"
)

const defaultRefreshInterval = 5 * time.Minute

// Cache holds the in-memory article cache with periodic refresh from GCS.
type Cache struct {
	storage Storage
	parser  *Parser
	logger  *slog.Logger

	mu       sync.RWMutex
	articles map[string]*Article // keyed by slug
	metaList []ArticleMeta       // published only, sorted by date desc

	stopCh chan struct{}
}

// NewCache creates a new blog cache, performs the initial load, and starts
// a background goroutine that refreshes the cache periodically.
func NewCache(ctx context.Context, storage Storage, parser *Parser, logger *slog.Logger) (*Cache, error) {
	c := &Cache{
		storage:  storage,
		parser:   parser,
		logger:   logger.With("service", "blog_cache"),
		articles: make(map[string]*Article),
		stopCh:   make(chan struct{}),
	}

	if err := c.refresh(ctx); err != nil {
		return nil, fmt.Errorf("initial blog cache load: %w", err)
	}

	go c.backgroundRefresh()

	return c, nil
}

// GetAllPublished returns metadata for all published articles, sorted newest first.
func (c *Cache) GetAllPublished() []ArticleMeta {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]ArticleMeta, len(c.metaList))
	copy(result, c.metaList)
	return result
}

// GetBySlug returns a full article (metadata + HTML) by slug, or nil if not found.
func (c *Cache) GetBySlug(slug string) *Article {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.articles[slug]
}

// refresh re-reads all .md files from GCS, parses them, and updates the cache.
func (c *Cache) refresh(ctx context.Context) error {
	files, err := c.storage.ListMarkdownFiles(ctx)
	if err != nil {
		return fmt.Errorf("listing markdown files: %w", err)
	}

	newArticles := make(map[string]*Article, len(files))
	var newMeta []ArticleMeta

	for _, filename := range files {
		data, err := c.storage.ReadFile(ctx, filename)
		if err != nil {
			c.logger.Error("failed to read article file", "file", filename, "error", err)
			continue
		}

		slug := strings.TrimSuffix(filename, ".md")
		article, err := c.parser.Parse(data, slug)
		if err != nil {
			c.logger.Error("failed to parse article", "file", filename, "error", err)
			continue
		}

		if !article.Published {
			c.logger.Debug("skipping unpublished article", "slug", slug)
			continue
		}

		newArticles[slug] = article
		newMeta = append(newMeta, article.ArticleMeta)
	}

	// Sort by date descending (newest first).
	sort.Slice(newMeta, func(i, j int) bool {
		return newMeta[i].Date > newMeta[j].Date
	})

	c.mu.Lock()
	c.articles = newArticles
	c.metaList = newMeta
	c.mu.Unlock()

	c.logger.Info("blog cache refreshed", "article_count", len(newMeta))

	// Write derived metadata.json to GCS (best-effort).
	cache := &MetadataCache{
		GeneratedAt: time.Now().UTC(),
		Articles:    newMeta,
	}
	if err := c.storage.WriteMetadataCache(ctx, cache); err != nil {
		c.logger.Error("failed to write metadata cache to GCS", "error", err)
	}

	return nil
}

// backgroundRefresh runs refresh on a ticker until Stop is called.
func (c *Cache) backgroundRefresh() {
	ticker := time.NewTicker(defaultRefreshInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			if err := c.refresh(ctx); err != nil {
				c.logger.Error("background cache refresh failed", "error", err)
			}
			cancel()
		case <-c.stopCh:
			return
		}
	}
}

// Stop shuts down the background refresh goroutine.
func (c *Cache) Stop() {
	close(c.stopCh)
}
