# Development Plan: v0.0.16

## Context & Motivation

v0.0.15 delivered the Nuxt 3 SSG migration, JSON-LD structured data, Open Graph/Twitter Card meta tags, and Pub/Sub push notifications for instant blog cache refresh. The site now scores 71/100 on SEO tools, up from near-zero crawlability as an SPA.

v0.0.16 focuses on **closing the remaining SEO gaps**, **hardening the blog pipeline**, and **improving developer/operator experience**. These are the highest-ROI items identified during v0.0.15 work and a follow-up SEO tool scan (March 2026).

### SEO Tool Scan Results (March 2026)

| Recommendation | Status | Action |
|---|---|---|
| FAQPage Schema | ✅ Already live (v0.0.15) | None — tool scan likely stale |
| OG & Twitter Card meta tags | ✅ Already live (v0.0.15) | None — tool scan likely stale |
| Organization & Author Schema | ⚠️ Partial | Have `ProfessionalService` but no explicit `Organization`. Author info invisible on blog (client-rendered). → **Phase 7** |
| Meta Description + Keyword Alignment | ❌ Needs fix | 189 chars (should be 150-160). H1 doesn't match primary keyword. → **Phase 8** |
| Content with Stats & Citations | ❌ New | No tables, data points, or citations on pages. → **Phase 9** |

---

## Phase 1: Dynamic Sitemap Generation
> Priority: HIGH | Estimated effort: ~1.5 hours

### Problem
The current `public/sitemap.xml` is a static file that was manually written. New blog articles (e.g., Ivmo's recent post) do not appear in the sitemap. The `lastmod` dates are stale (2025-09-06 for pages, 2026-02-27 for all articles).

### Solution
Replace the static file with auto-generated sitemaps using `@nuxtjs/sitemap`.

- [ ] **1.1** Install `@nuxtjs/sitemap`:
  ```bash
  npm install -D @nuxtjs/sitemap
  ```

- [ ] **1.2** Add to `nuxt.config.ts` modules and configure:
  ```ts
  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/sitemap'],
  site: {
    url: 'https://ivmanto.com',
  },
  sitemap: {
    sources: ['/api/__sitemap__/blog'],
    exclude: ['/login', '/booking-demo'],
  },
  ```

- [ ] **1.3** Create `server/api/__sitemap__/blog.ts` — a Nitro server route that fetches article slugs from the backend API at generate-time and returns them as sitemap entries:
  ```ts
  export default defineSitemapEventHandler(async () => {
    const articles = await $fetch('https://ivmanto-backend-service-ffdh2bu2iq-ey.a.run.app/api/articles')
    return articles.map((a) => ({
      loc: `/blog/${a.slug}`,
      lastmod: a.date,
      priority: 0.8,
    }))
  })
  ```

- [ ] **1.4** Delete the static `public/sitemap.xml` (will be auto-generated)

- [ ] **1.5** Update `public/robots.txt` to point to the new sitemap URL (should still be `/sitemap.xml`)

- [ ] **1.6** Verify with `npx nuxi generate` that `.output/public/sitemap.xml` is generated correctly with all pages and blog articles

---

## Phase 2: Blog Frontmatter Validation & Resilience
> Priority: HIGH | Estimated effort: ~1.5 hours

### Problem
Ivmo published an article missing `published: true` and `summary` fields. The article was silently skipped by the cache — no error, no warning in the API response. The digital assistant needs clear requirements, and the backend should be more resilient and informative.

### Solution

#### 2A: Backend — Frontmatter Validation Logging

- [ ] **2.1** Add validation warnings in `cache.go` `refresh()` method. After parsing, log specific warnings for missing required fields:
  ```go
  // After parsing, before the Published check:
  if article.Title == "" {
      c.logger.Warn("article missing title", "slug", slug, "file", filename)
  }
  if article.Summary == "" {
      c.logger.Warn("article missing summary", "slug", slug, "file", filename)
  }
  if article.Date == "" {
      c.logger.Warn("article missing date", "slug", slug, "file", filename)
  }
  if !article.Published {
      c.logger.Info("skipping unpublished article (published != true)", "slug", slug)
      continue
  }
  ```

- [ ] **2.2** Add a `GET /api/_internal/articles/status` diagnostic endpoint to `handler.go` that returns all articles (including unpublished) with their validation status. Protect with the same push token:
  ```json
  {
    "total_files": 18,
    "published": 17,
    "skipped": [
      {
        "slug": "the-shift-to-agentic-ai",
        "reason": "published field is false or missing",
        "has_title": true,
        "has_summary": false,
        "has_date": true
      }
    ]
  }
  ```
  This gives Ivmo (and operators) a way to diagnose why an article isn't appearing.

#### 2B: Ivmo Frontmatter Contract

- [ ] **2.3** Create `docs/blog-article-spec.md` — a specification document for the blog article frontmatter contract that Ivmo must follow:
  - Required fields: `title`, `summary`, `date`, `published`
  - Optional fields: `author`, `tags`
  - Filename convention: `kebab-case-title.md`
  - Example template with all fields
  - Validation rules (e.g., `published` must be `true` for the article to appear)

---

## Phase 3: Dynamic Sitemap for Blog (Backend Endpoint)
> Priority: MEDIUM | Estimated effort: ~1 hour

### Problem
Even with `@nuxtjs/sitemap` at build time, new articles published between site builds won't appear in the sitemap. Since the site is SSG (static), the sitemap is frozen at build time.

### Solution
Add a backend endpoint that serves a dynamic sitemap fragment for blog articles.

- [ ] **3.1** Add `GET /api/sitemap-blog.xml` endpoint to the Go backend that generates a sitemap XML from the current article cache:
  ```go
  func (h *Handler) handleBlogSitemap(w http.ResponseWriter, r *http.Request) {
      articles := h.cache.GetAllPublished()
      w.Header().Set("Content-Type", "application/xml")
      // Generate <urlset> with all published articles
  }
  ```

- [ ] **3.2** Register the route in `handler.go`:
  ```go
  mux.HandleFunc("GET /api/sitemap-blog.xml", h.handleBlogSitemap)
  ```

- [ ] **3.3** Update `public/robots.txt` to include both sitemaps:
  ```
  Sitemap: https://ivmanto.com/sitemap.xml
  Sitemap: https://ivmanto.com/api/sitemap-blog.xml
  ```

---

## Phase 4: Image & Font Performance Optimization
> Priority: MEDIUM | Estimated effort: ~1 hour

### Problem
- Google Fonts are loaded via external `<link>` tags, blocking the critical rendering path
- The `nt.jpg` photo and `cloud-pic-2.webp` are not optimized for responsive loading
- No explicit width/height on images can cause layout shifts (CLS)

### Solution

- [ ] **4.1** Self-host Google Fonts to eliminate the external request chain:
  - Install `@nuxtjs/google-fonts`:
    ```bash
    npm install -D @nuxtjs/google-fonts
    ```
  - Add to `nuxt.config.ts`:
    ```ts
    modules: ['@nuxtjs/tailwindcss', '@nuxtjs/sitemap', '@nuxtjs/google-fonts'],
    googleFonts: {
      families: {
        Montserrat: [400, 700],
      },
      display: 'swap',
      download: true,
      preload: true,
    },
    ```
  - Remove the manual `<link>` tags from `nuxt.config.ts` → `app.head.link`

- [ ] **4.2** Add explicit `width` and `height` attributes to all `<img>` tags across pages/components to prevent Cumulative Layout Shift (CLS):
  - `pages/about.vue` — `nt.jpg`
  - `layouts/default.vue` or `components/layout/AppLogo.vue` — logo
  - `components/sections/HeroInfographicSection.vue` — hero image
  - Any other images discovered during audit

- [ ] **4.3** Add `loading="lazy"` to below-the-fold images (about page photo, blog images, etc.)

---

## Phase 5: Update `robots.txt` & Security Headers
> Priority: LOW | Estimated effort: ~30 minutes

- [ ] **5.1** Update `robots.txt`:
  ```
  User-agent: *
  Allow: /
  Disallow: /login
  Disallow: /booking-demo
  Disallow: /api/_internal/

  Sitemap: https://ivmanto.com/sitemap.xml
  Sitemap: https://ivmanto.com/api/sitemap-blog.xml
  ```

- [ ] **5.2** Remove the stale `Disallow: /search` rule (no `/search` route exists)

---

## Phase 6: Version Bump & Cleanup
> Priority: LOW | Estimated effort: ~15 minutes

- [ ] **6.1** Bump `package.json` version from `0.0.15` to `0.0.16`

- [ ] **6.2** Remove legacy dev dependencies that are no longer needed:
  - Audit `package.json` for any leftover Vue SPA deps (e.g., `@rushstack/eslint-patch`, ESLint/Prettier configs that may not align with Nuxt)

- [ ] **6.3** Run `npx nuxi generate` and verify full build passes with no warnings

- [ ] **6.4** Test locally with `npx serve .output/public -s` to verify all routes work

---

## Phase 7: Strengthen Organization & Author Schema (SEO Tool Fix)
> Priority: HIGH | Estimated effort: ~1 hour

### Problem
The SEO tool reports "no Organization schema" and "no author information." We have `ProfessionalService` (a subtype of `Organization`) but the tool doesn't recognize it as a standalone Organization. Google's Dec 2025 update shows sites with author credentials gained +2.3 ranking positions. This is critical for E-E-A-T (Experience, Expertise, Authoritativeness, Trustworthiness) signals.

Additionally, blog articles have `BlogPosting` schema with author references, but since blog pages are client-rendered (`ssr: false`), crawlers see empty HTML and can't read the JSON-LD.

### Solution

- [ ] **7.1** Add explicit `Organization` schema alongside `ProfessionalService` in `layouts/default.vue`:
  ```json
  {
    "@context": "https://schema.org",
    "@type": ["Organization", "ProfessionalService"],
    "@id": "https://ivmanto.com/#organization",
    "name": "IVMANTO",
    "alternateName": "IVMANTO - Nikolay Tonev Data & Cloud Solutions",
    "url": "https://ivmanto.com",
    "logo": "https://ivmanto.com/logo.png",
    "image": "https://ivmanto.com/social-sharing-card.webp",
    "description": "Expert Cloud Data Architecture & AI Solutions on Google Cloud Platform.",
    "founder": { "@id": "https://ivmanto.com/about#person" },
    "address": {
      "@type": "PostalAddress",
      "addressCountry": "DE"
    },
    "sameAs": [
      "https://linkedin.com/in/nikolaytonev",
      "https://github.com/dasiyes",
      "https://g.dev/ivmanto-nikolaytonev"
    ]
  }
  ```
  Using `@type: ["Organization", "ProfessionalService"]` satisfies both the generic Organization check and the specific ProfessionalService type.

- [ ] **7.2** Enrich the `Person` schema on the about page (`pages/about.vue`) with author credentials:
  ```json
  {
    "@type": "Person",
    "@id": "https://ivmanto.com/about#person",
    "name": "Nikolay Tonev",
    "jobTitle": "Cloud Data Architect & AI Consultant",
    "worksFor": { "@id": "https://ivmanto.com/#organization" },
    "url": "https://ivmanto.com/about",
    "sameAs": [
      "https://linkedin.com/in/nikolaytonev",
      "https://github.com/dasiyes",
      "https://g.dev/ivmanto-nikolaytonev"
    ],
    "knowsAbout": [
      "Google Cloud Platform",
      "Data Architecture",
      "AI & Machine Learning",
      "Data Governance",
      "DAMA-DMBOK"
    ]
  }
  ```

- [ ] **7.3** Add a global `author` meta tag in `layouts/default.vue`:
  ```ts
  useSeoMeta({
    author: 'Nikolay Tonev',
  })
  ```

---

## Phase 8: Meta Description & Primary Keyword Alignment (SEO Tool Fix)
> Priority: HIGH | Estimated effort: ~45 minutes

### Problem
The SEO tool reports:
- Meta description is **189 characters** (should be 150-160)
- Primary keyword **"Data AI Consultancy"** is missing from title, description, AND H1
- The H1 currently says "Expert Cloud Data Architecture & AI Solutions" — doesn't align with the primary keyword

Current state:
- **Title**: `ivmanto.com | Data & AI Consultancy` ← has keyword variant but with `&`
- **Meta description**: `Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset.` ← 189 chars
- **H1**: `Expert Cloud Data Architecture & AI Solutions` ← no "consultancy" keyword

### Solution

- [ ] **8.1** Shorten meta description to 150-160 characters and include the primary keyword in `layouts/default.vue`:
  ```
  Before (189 chars):
  "Expert Data & AI consultancy specializing in Google Cloud Platform (GCP). We help businesses with data architecture, governance, and AI-driven solutions to turn data into a strategic asset."

  After (~155 chars):
  "Data & AI Consultancy on Google Cloud Platform. Expert data architecture, governance, and AI solutions that turn your data into a strategic asset."
  ```

- [ ] **8.2** Align the H1 on the homepage (`pages/index.vue`) to include the primary keyword:
  ```
  Before: "Expert Cloud Data Architecture & AI Solutions"
  After:  "Data & AI Consultancy — Expert Cloud Architecture & Solutions"
  ```
  This ensures the H1, title, and description all reinforce the same keyword cluster.

- [ ] **8.3** Update the OG title and Twitter title defaults in `layouts/default.vue` to match the new title format

- [ ] **8.4** Review and tighten per-page meta descriptions on `pages/about.vue`, `pages/services/index.vue`, and `pages/blog/index.vue` — each should be unique, 150-160 chars, and include relevant keywords

---

## Phase 9: Content Depth Signals — Statistics & Citations (SEO Tool Fix)
> Priority: MEDIUM | Estimated effort: ~2 hours

### Problem
The SEO tool found "content lacks depth signals — no tables, structured data points, or authoritative quotes." Princeton GEO research shows adding statistics increases AI visibility by 37% and citations by 40%.

### Solution
Add structured content elements to key marketing pages to boost both traditional and AI search visibility.

- [ ] **9.1** Add a **"Key Facts & Figures"** section or inline statistics to the homepage (`pages/index.vue`):
  - Example data points: years of experience, number of GCP certifications, number of successful projects, industries served
  - Use a visually distinct format (e.g., a stats bar or card grid) that also renders as semantic HTML

- [ ] **9.2** Add a **credentials/certifications table** to `pages/about.vue`:
  - Google Cloud certifications with dates
  - DAMA/DMBOK qualifications
  - Use `<table>` HTML element for structured data

- [ ] **9.3** Add **authoritative citations** to service pages where applicable:
  - Reference Gartner, Forrester, or Google Cloud documentation for industry claims
  - Use `<blockquote cite="...">` for proper semantic markup
  - Example: On the Data Architecture page, cite Google's BigQuery benchmarks or Gartner's data management quadrant

- [ ] **9.4** Consider adding an `ItemList` schema to the services landing page (`pages/services/index.vue`) to give search engines a structured view of all offered services

---

## Phase 10: Version Verification & Final Build
> Priority: LOW | Estimated effort: ~30 minutes

- [ ] **10.1** Run `npx nuxi generate` and verify full build passes with no warnings
- [ ] **10.2** Validate all schemas using https://validator.schema.org/
- [ ] **10.3** Verify meta descriptions with `curl` for all pre-rendered pages
- [ ] **10.4** Test locally with `npx serve .output/public -s` to verify all routes work
- [ ] **10.5** Re-run the SEO tool to verify improvements

---

## Summary

| Phase | Description | Priority | Effort | Source |
|-------|-------------|----------|--------|--------|
| 1 | Dynamic sitemap generation | HIGH | ~1.5h | v0.0.15 gap |
| 2 | Blog frontmatter validation & spec | HIGH | ~1.5h | v0.0.15 incident |
| 3 | Backend dynamic blog sitemap | MEDIUM | ~1h | v0.0.15 gap |
| 4 | Image & font performance | MEDIUM | ~1h | Performance |
| 5 | robots.txt & security headers | LOW | ~30m | Housekeeping |
| 6 | Version bump & cleanup | LOW | ~15m | Housekeeping |
| 7 | Organization & Author Schema | HIGH | ~1h | SEO tool scan |
| 8 | Meta Description & Keyword Alignment | HIGH | ~45m | SEO tool scan |
| 9 | Content Depth Signals (Stats & Citations) | MEDIUM | ~2h | SEO tool scan |
| 10 | Final verification & build | LOW | ~30m | QA |

**Total estimated effort: ~10 hours**

---

## Out of Scope (Future Versions)

The following items were considered but deferred to keep v0.0.16 focused:

- **Blog SSR/ISR** (Option B from v0.0.15): Switch blog pages from `ssr: false` to server-rendered with ISR. This would make blog content visible to crawlers in the HTML, but requires switching from `nuxt generate` (pure SSG) to `nuxt build` (hybrid server mode) and changing the Dockerfile to run a Node.js server instead of `serve`. Significant architectural change — better suited for a dedicated version.

- **New service pages**: Data Quality, Data Engineering, AI Solutions, Agentic AI, Backend Development, Microservices Design (from `TODO_in_services.md`). Content-heavy work that should be its own initiative.

- **Customer Authentication & "My Space"**: Per the roadmap, this is a "Version 2.0" feature that builds on top of the backend integration.

- **Internal linking strategy**: Cross-linking between service pages and blog articles. Requires content audit and strategy — better as a dedicated SEO sprint.

- **Nuxt Image module (`@nuxt/image`)**: Full responsive image pipeline with automatic format conversion (WebP/AVIF). Adds build complexity — evaluate when image count grows.
