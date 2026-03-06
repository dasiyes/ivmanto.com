# How the Process of Publishing a New Article Works

## Overview

Publishing a new article requires only uploading a `.md` file with proper YAML frontmatter to the GCS bucket. Everything else — cache refresh, metadata generation, and frontend rebuild — happens automatically.

## End-to-End Flow

```
Agent uploads article.md to GCS
  → GCS notification triggers Pub/Sub
    → Backend receives push, sees it's a .md file
      → Debounced cache refresh reads ALL .md files
        → metadata.json is auto-regenerated (best-effort)
        → Frontend rebuild webhook is triggered (Cloud Build)
```

## What the Publishing Agent Needs to Do

1. **Upload a single `.md` file** to the GCS articles bucket.
2. The file must include valid **YAML frontmatter** with these fields:

```yaml
---
title: "Your Article Title"
summary: "A short description of the article."
date: "2026-03-06"
published: true
---
```

3. The markdown body follows the frontmatter as usual.

> **Important:** The agent does **not** need to edit `metadata.json`. It is auto-generated.

## Key Backend Components

| Component | File | Role |
|---|---|---|
| Cache | `backend/internal/blog/cache.go` | Reads all `.md` files from GCS, parses them, builds in-memory cache, writes `metadata.json` |
| Handler | `backend/internal/blog/handler.go` | Serves `GET /api/articles` and `GET /api/articles/{slug}` from the in-memory cache |
| Pub/Sub push | `backend/internal/blog/handler.go` | Receives GCS notifications, filters for `.md` files only, triggers debounced cache refresh |
| Frontend rebuild | `backend/internal/blog/handler.go` | On `OBJECT_FINALIZE` events for `.md` files, calls the Cloud Build webhook to trigger a frontend rebuild |
| Parser | `backend/internal/blog/parser.go` | Parses YAML frontmatter + markdown body into `Article` structs with rendered HTML |

## About `metadata.json`

- **Generated automatically** by the backend after every cache refresh (lines 168–175 in `cache.go`).
- **Never read** by the backend or frontend — the API serves from the in-memory cache built directly from `.md` files.
- **Ignored by Pub/Sub** — the handler filters for `.md` suffix only (line 99 in `handler.go`), so writing `metadata.json` does not cause a refresh loop.
- It exists as a **diagnostic/convenience snapshot** and could be removed without breaking anything.

## Duplicate Title Fix (2026-03-06)

The blog article page (`pages/blog/[slug].vue`) was showing the article title twice:
1. Once from the template's `<h1>{{ article.title }}</h1>` (metadata-driven).
2. Once from the `<h1>` inside the rendered markdown content (`v-html`).

**Fix:** The `sanitizedContent` computed property now strips the leading `<h1>` from the HTML content before rendering, since the title is already displayed from structured metadata. Each page retains exactly one `<h1>` for proper SEO.
