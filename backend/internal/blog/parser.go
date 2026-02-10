package blog

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	gmparser "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
)

// Parser converts markdown content to HTML and extracts frontmatter metadata.
type Parser struct {
	md goldmark.Markdown
}

// NewParser creates a new markdown parser with frontmatter and GFM support.
func NewParser() *Parser {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&frontmatter.Extender{},
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	return &Parser{md: md}
}

// Parse takes raw markdown bytes and returns extracted metadata and rendered HTML.
// The slug is derived from the filename, not from frontmatter.
func (p *Parser) Parse(source []byte, slug string) (*Article, error) {
	var buf bytes.Buffer

	ctx := gmparser.NewContext()
	if err := p.md.Convert(source, &buf, gmparser.WithContext(ctx)); err != nil {
		return nil, fmt.Errorf("goldmark convert: %w", err)
	}

	var meta ArticleMeta
	d := frontmatter.Get(ctx)
	if d != nil {
		if err := d.Decode(&meta); err != nil {
			return nil, fmt.Errorf("frontmatter decode: %w", err)
		}
	}

	meta.Slug = slug

	return &Article{
		ArticleMeta: meta,
		Content:     buf.String(),
	}, nil
}
