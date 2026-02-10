package blog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// Storage defines the interface for blog article storage operations.
type Storage interface {
	ListMarkdownFiles(ctx context.Context) ([]string, error)
	ReadFile(ctx context.Context, name string) ([]byte, error)
	WriteMetadataCache(ctx context.Context, cache *MetadataCache) error
}

type gcsStorage struct {
	client *storage.Client
	bucket string
	logger *slog.Logger
}

// NewStorage creates a new GCS-backed storage service.
func NewStorage(client *storage.Client, bucket string, logger *slog.Logger) Storage {
	return &gcsStorage{
		client: client,
		bucket: bucket,
		logger: logger.With("service", "blog_gcs"),
	}
}

// ListMarkdownFiles returns the names of all .md files in the bucket root.
func (s *gcsStorage) ListMarkdownFiles(ctx context.Context) ([]string, error) {
	bkt := s.client.Bucket(s.bucket)
	it := bkt.Objects(ctx, &storage.Query{
		Prefix:    "",
		Delimiter: "/",
	})

	var files []string
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("listing bucket objects: %w", err)
		}
		if strings.HasSuffix(attrs.Name, ".md") {
			files = append(files, attrs.Name)
		}
	}
	return files, nil
}

// ReadFile reads the full contents of a file from the bucket.
func (s *gcsStorage) ReadFile(ctx context.Context, name string) ([]byte, error) {
	rc, err := s.client.Bucket(s.bucket).Object(name).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("opening object %q: %w", name, err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("reading object %q: %w", name, err)
	}
	return data, nil
}

// WriteMetadataCache writes the metadata cache as metadata.json to the bucket.
func (s *gcsStorage) WriteMetadataCache(ctx context.Context, cache *MetadataCache) error {
	w := s.client.Bucket(s.bucket).Object("metadata.json").NewWriter(ctx)
	w.ContentType = "application/json"

	if err := json.NewEncoder(w).Encode(cache); err != nil {
		w.Close()
		return fmt.Errorf("encoding metadata cache: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("closing metadata.json writer: %w", err)
	}

	s.logger.Info("wrote metadata.json to GCS", "article_count", len(cache.Articles))
	return nil
}
