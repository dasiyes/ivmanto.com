# Release dev-v0.1.3 — Featured Services & AI Search

## Overview

This release adds three new, prominently featured services to the Services
section and makes them discoverable by traditional search engines and AI search
engines through structured keywords and an expanded AI-search file set.

## What shipped

### 1. Three new featured services

Added to `data/services.ts` and rendered first on `/services` with a "Featured"
badge and amber accent styling:

| Service | URL |
|---|---|
| AI & Automation Strategic Discovery | `/services/ai-automation-discovery` |
| Data Pipeline Design & Architecture | `/services/data-pipeline-engineering` |
| Agentic AI Solution Design & Team Enablement | `/services/agentic-ai-solutions` |

Each has a full detail page (`/services/[id]`) with its own content component,
relevant-industries list, per-page SEO metadata, and `Service` JSON-LD.

The `Service` type gained two optional fields:

- `featured` — drives the badge + accent treatment on the index grid.
- `keywords` — a per-service keyword set (see below).

### 2. Keyword indexing

Each featured service has a curated keyword set, surfaced two ways so both
traditional crawlers and AI search engines can pick it up:

- In the `Service` JSON-LD `keywords` property (structured data).
- As visible "Topics" tag pills on the service cards and detail pages.

### 3. AI-search semantics

- `Service` JSON-LD enriched with `keywords` and `category`.
- `public/llms.txt` refreshed — the three new services are listed first.
- New `public/llms-full.txt` — extended descriptions, key deliverables, and
  topic keywords for all eight services, aimed at AI crawlers.

## Files changed

- `data/services.ts` — `Service` type fields + 3 new service entries
- `components/services-content/` — 3 new detail content components
- `pages/services/index.vue` — featured badge/accent + keyword tags
- `pages/services/[id].vue` — per-service SEO, enriched JSON-LD, Topics sidebar
- `composables/usePageMetadata.ts` — routes for the 3 new services
- `public/llms.txt`, `public/llms-full.txt` — AI-search files

## Verification

- `npm run generate` — 102 routes prerendered, exit 0, no errors; all three
  new service pages build.
- Generated output checked: `sitemap.xml` includes all 8 service URLs;
  detail-page JSON-LD carries `keywords` + `category`; featured CSS compiled.
- Visual check of `/services` and a service detail page in the browser.

## Known issues / follow-ups

These are pre-existing and were not introduced by this release:

- `npm run lint` fails — `eslint.config.ts` (flat config) requires ESLint 9,
  but `package.json` pins ESLint 8.57. Formatting is handled by Prettier
  (`npm run format`), which works.
- Pages set no explicit `background-color`, so visitors with OS dark mode see
  content pages rendered dark-on-dark. A single global background rule on
  `html`/`body` resolves it.
