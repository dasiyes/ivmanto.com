# Agent Collaboration Rules

## Core Principles
* Simplicity First: Make every change as simple as possible. Impact minimal code.
* No Laziness: Find root causes. No temporary fixes. Senior developer standards.
* Minimal Impact: Changes should only touch what's necessary. Avoid introducing bugs.

## Task Management
1. Plan First: Write plan to tasks/todo.md with checkable items
2. Verify Plan: Check in before starting implementation
3. Track Progress: Mark items complete as you go
4. Explain Changes: High-level summary at each step
5. Document Results: Add review section to tasks/todo.md
6. Capture Lessons: Update tasks/lessons.md after corrections

---

# Project Architecture & Guidelines

## Project Overview

**IVMANTO — Cloud Data Architecture & AI Consultancy Website**
A freelancer's professional website serving as a sales tool. Built to showcase data architecture, AI/ML, and cloud consulting services on Google Cloud Platform.

- **Framework**: Nuxt 3 (Vue 3) with SSG (Static Site Generation)
- **Styling**: TailwindCSS 3.x with `@tailwindcss/typography` plugin
- **Language**: TypeScript (strict)
- **Deployment**: Docker → Cloud Run (port 8080)
- **Backend API**: Go backend (separate service), proxied via `/api` routes
- **Analytics**: Google Tag Manager (GTM-PXRF8FPQ) + GA4 (G-W1TJ3KMZ6V)
- **AI Features**: Google Gemini API (inspiration ideas generation)

---

## ⚠️ Critical Rules (STRICT)

1. **CONTENT PRESERVATION**: All existing page content, copy, styling, and user-facing behavior MUST remain identical after any migration or refactoring. Zero content regressions.
2. **NO UNAUTHORIZED CHANGES**: Never apply code changes unless specifically requested.
3. **BLOG SYSTEM INTEGRITY**: Blog articles are served dynamically from the Go backend (`/api/articles`), which reads from a GCS bucket. This runtime data-fetching pattern MUST be preserved — articles are NOT static and must be fetched at request time or via ISR/SWR.
4. **BOOKING SYSTEM INTEGRITY**: The booking calendar (`/api/booking/availability`, `/api/booking/book`, `/api/booking/cancel`) reads real-time availability from the Go backend. This MUST remain a fully client-side dynamic feature — never pre-rendered.
5. **COOKIE CONSENT FIRST**: Google Tag Manager and all analytics scripts MUST only initialize AFTER explicit user consent via the cookie banner. No analytics on page load without consent.
6. **API PROXY PATTERN**: All `/api/*` requests proxy to the Go backend. In development, this is handled by the Nuxt dev server proxy. In production, the Go backend serves both the API and the static frontend.

---

## Architecture Overview

### Rendering Strategy (Hybrid SSG)

The site uses **Nuxt 3 Hybrid Rendering** — a mix of pre-rendered static pages and client-side dynamic features:

| Route Pattern | Rendering Mode | Reason |
|---|---|---|
| `/` (Home) | **SSG** (pre-rendered at build) | Static marketing content, SEO-critical |
| `/about` | **SSG** | Static content, SEO-critical |
| `/services` | **SSG** | Static content from `services.ts` |
| `/services/:id` | **SSG** | Static content, all IDs known at build time |
| `/blog` | **SSR** or **ISR** | Article list fetched from API at runtime |
| `/blog/:slug` | **SSR** or **ISR** | Article content fetched from API at runtime |
| `/booking` | **SPA/CSR** | Fully dynamic (real-time availability) |
| `/booking/cancel` | **SPA/CSR** | Dynamic (token-based cancellation) |
| `/booking-demo` | **SSG** | Static iframe embed |
| `/login` | **SSG** | Static placeholder page (noindex) |
| `/privacy-policy` | **SSG** | Static legal content |

**Key Concept**: Pages that need real-time data (booking, blog) use `ssr: false` (client-only) or ISR. Marketing pages are fully pre-rendered for maximum SEO benefit.

### Blog Articles — GCS Bucket Pipeline

**How it works today (MUST be preserved):**
1. Author publishes a markdown/HTML article to a GCS bucket
2. The Go backend serves articles via REST API:
   - `GET /api/articles` → returns `ArticleMeta[]` (slug, title, summary, date, published)
   - `GET /api/articles/:slug` → returns full `Article` (includes HTML content)
   - `GET /api/articles/:slug/likes` → returns like count
   - `POST/DELETE /api/articles/:slug/like` → toggle like
3. The frontend fetches and renders articles at runtime
4. Article content is sanitized with DOMPurify before rendering (`v-html`)

**In Nuxt SSG context**: Blog pages CANNOT be pre-rendered at build time because articles are added dynamically to the GCS bucket. Options:
- **Option A (Recommended)**: Use `routeRules` with ISR (Incremental Static Regeneration) — pages are generated on first visit and cached with a TTL
- **Option B**: Mark `/blog` routes as `ssr: false` (pure client-side rendering, like current SPA behavior)

### Booking System — Real-Time Calendar

**How it works today (MUST be preserved):**
1. User selects a date → `GET /api/booking/availability?date=YYYY-MM-DD` returns available `TimeSlot[]`
2. User selects a slot and fills form → `POST /api/booking/book` with slot ID + user details
3. Confirmation email sent by backend
4. Cancellation via `POST /api/booking/cancel` with JWT token from email link
5. GA session info (clientId, sessionId) is sent with booking for server-side analytics stitching

**In Nuxt SSG context**: Booking pages MUST be client-only rendered. Use `<ClientOnly>` wrapper or `ssr: false` route rule.

### Contact Form & Ideas Generation

- **Contact form** (`POST /api/contact`): Sends form data + GA clientId to backend for email dispatch
- **Ideas generation** (`POST /api/generate-ideas`): Backend proxies to Gemini API, returns `Idea[]`
- **Ideas email** (`POST /api/ideas/email`): Sends generated ideas to user's email via backend
- **Gemini direct** (`useGemini.ts`): Client-side Gemini API calls (currently used but may be superseded by backend proxy)

---

## Structured Data / Schema Markup Strategy

### Current Schema (preserve and enhance):

1. **ProfessionalService** (global, in layout) — `@id: https://ivmanto.com/#organization`
2. **Person** (on `/about` page) — `@id: https://ivmanto.com/about#person`, linked to org via `worksFor`
3. **Service** (on each `/services/:id` page) — dynamic per-service, linked to org via `provider`

### New Schema to Add:

4. **WebSite** (global, in layout) — with `SearchAction` pointing to `/blog` search
5. **FAQPage** (on home page) — from the 4 FAQ items in `FAQSection.vue`
6. **Article/BlogPosting** (on each `/blog/:slug` page) — for blog article SEO

### Implementation Rule:
In Nuxt, schema markup is injected via `useHead()` or `useSchemaOrg()`. Since pages are pre-rendered (SSG), the JSON-LD will be present in the static HTML — solving the original SEO problem of crawlers not seeing schema in the SPA.

---

## Directory Structure (Nuxt 3 Conventions)

```
ivmanto.com/
├── .agents/                     # Agent rules & workflows (unchanged)
├── assets/                      # Unprocessed assets (images, fonts)
│   ├── css/
│   │   └── main.css             # Tailwind directives + custom styles
│   ├── mockup/
│   │   └── logo.svg
│   └── nt.jpg
├── components/                  # Auto-imported Vue components
│   ├── layout/
│   │   ├── AppHeader.vue
│   │   ├── AppFooter.vue
│   │   └── AppLogo.vue
│   ├── sections/
│   │   ├── HeroInfographicSection.vue
│   │   ├── ProcessSection.vue
│   │   ├── FAQSection.vue
│   │   ├── ContactSection.vue
│   │   └── ...
│   ├── about/
│   │   ├── ApproachSection.vue
│   │   ├── ExpertiseSection.vue
│   │   └── CertificatesSection.vue
│   ├── services/
│   │   └── ServiceDetail.vue
│   ├── services-content/
│   │   ├── DataArchitecture.vue
│   │   ├── DataGovernance.vue
│   │   ├── MlEngineering.vue
│   │   ├── Principles.vue
│   │   └── SovereignCloudDE.vue
│   ├── ContactForm.vue
│   ├── CookieBanner.vue
│   ├── InspirationModal.vue
│   └── PrivacyPolicy.vue
├── composables/                 # Auto-imported composables
│   ├── useArticles.ts
│   ├── useGemini.ts
│   └── usePageMetadata.ts       # May be replaced by Nuxt's built-in useSeoMeta()
├── data/                        # Static data files
│   └── services.ts
├── layouts/                     # Nuxt layouts
│   └── default.vue              # Header + Footer + global schema + cookie banner
├── pages/                       # File-based routing (replaces vue-router)
│   ├── index.vue                # / (Home)
│   ├── about.vue                # /about
│   ├── booking/
│   │   ├── index.vue            # /booking
│   │   └── cancel.vue           # /booking/cancel
│   ├── booking-demo.vue         # /booking-demo
│   ├── blog/
│   │   ├── index.vue            # /blog (article list)
│   │   └── [slug].vue           # /blog/:slug (article detail)
│   ├── services/
│   │   ├── index.vue            # /services (landing)
│   │   └── [id].vue             # /services/:id (detail)
│   ├── login.vue                # /login
│   ├── privacy-policy.vue       # /privacy-policy
│   └── [...slug].vue            # Catch-all 404
├── plugins/                     # Nuxt plugins
│   └── analytics.client.ts      # GTM/GA4 initialization (client-only)
├── public/                      # Static files served as-is
│   ├── certs/                   # Certificate badge images
│   ├── cloud-pic-2.webp
│   ├── favicon.ico
│   ├── logo.png
│   └── social-sharing-card.webp
├── server/                      # Nuxt server routes (if needed for API proxy)
│   └── api/                     # Optional: server-side API proxy routes
├── services/                    # Non-auto-imported service modules
│   ├── analytics.ts
│   ├── api.ts
│   └── gemini.ts
├── types/                       # TypeScript type definitions
│   └── article.ts
├── nuxt.config.ts               # Nuxt configuration (replaces vite.config.ts)
├── tailwind.config.js           # TailwindCSS configuration
├── tsconfig.json                # TypeScript configuration
├── Dockerfile                   # Docker build for Cloud Run
├── package.json
└── .env                         # Environment variables (NUXT_PUBLIC_* prefix)
```

---

## Key Migration Patterns

### 1. Routing: vue-router → File-based

**Before** (vue-router):
```ts
{ path: '/services/:id', component: ServiceView, props: true }
```

**After** (Nuxt file-based):
```
pages/services/[id].vue
```
- Route params become file/folder names in brackets
- `props: true` is replaced by `useRoute().params.id`

### 2. Head Management: @vueuse/head → Nuxt useHead/useSeoMeta

**Before**:
```ts
import { useHead } from '@vueuse/head'
useHead({ title: 'Page Title', meta: [...] })
```

**After**:
```ts
useHead({ title: 'Page Title' })
useSeoMeta({ description: '...' })
```
- `useHead()` and `useSeoMeta()` are auto-imported in Nuxt
- No need for `createHead()` plugin registration

### 3. Data Fetching: fetch in onMounted → useFetch/useAsyncData

**Before** (SPA pattern):
```ts
onMounted(async () => {
  const res = await fetch('/api/articles')
  articles.value = await res.json()
})
```

**After** (Nuxt SSR-aware):
```ts
const { data: articles } = await useFetch('/api/articles')
```
- `useFetch` is SSR-aware — runs on server during SSR, provides cached data on client
- For client-only pages (booking), continue using `onMounted` + `fetch` or wrap in `<ClientOnly>`

### 4. Layout: App.vue → layouts/default.vue

**Before** (`App.vue`):
```vue
<template>
  <AppHeader />
  <router-view />
  <AppFooter />
  <CookieBanner />
</template>
```

**After** (`layouts/default.vue`):
```vue
<template>
  <AppHeader />
  <slot />
  <AppFooter />
  <CookieBanner />
</template>
```

### 5. Environment Variables: VITE_* → NUXT_PUBLIC_*

**Before**: `import.meta.env.VITE_GEMINI_API_KEY`
**After**: `useRuntimeConfig().public.geminiApiKey`

### 6. Analytics: main.ts initialization → Plugin

**Before** (in `main.ts`):
```ts
if (consent === 'accepted') { initGtm() }
```

**After** (`plugins/analytics.client.ts`):
```ts
export default defineNuxtPlugin(() => {
  const consent = localStorage.getItem('cookie_consent')
  if (consent === 'accepted') { initGtm() }
})
```

### 7. Dynamic Component Loading: defineAsyncComponent → Nuxt auto-imports

**Before**:
```ts
detailsComponent: defineAsyncComponent(
  () => import('@/components/services-content/DataArchitecture.vue')
)
```

**After**: Components in `components/` are auto-imported. For the dynamic service detail pattern, use `resolveComponent()` or keep `defineAsyncComponent` (both work in Nuxt).

---

## Environment Configuration

### Required Environment Variables
| Variable | Usage | Nuxt Convention |
|---|---|---|
| `VITE_GEMINI_API_KEY` | Client-side Gemini API calls | `NUXT_PUBLIC_GEMINI_API_KEY` |
| (Go backend vars) | Backend-only | Not in frontend config |

### Nuxt Runtime Config
```ts
// nuxt.config.ts
export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      geminiApiKey: '', // overridden by NUXT_PUBLIC_GEMINI_API_KEY
    }
  }
})
```

---

## Deployment Architecture

### Current (SPA):
```
Docker Build → npm run build → /dist (static) → serve -s .
Go Backend: separate service, frontend proxies /api to it
```

### Target (Nuxt SSG):
```
Docker Build → npx nuxi generate → .output/public (static) → serve -s .
Go Backend: unchanged, frontend proxies /api to it
```

**Key difference**: `nuxt generate` produces one HTML file per pre-rendered route. Dynamic routes (`/blog/:slug`) either:
- Are excluded from generation and served as client-rendered SPAs
- Use ISR when running with a Nuxt server (requires switching from `serve` to `node .output/server/index.mjs`)

### Dockerfile Update
The Dockerfile currently uses `serve -s .` for the SPA. For pure SSG (no ISR), this remains the same — just the build command and output directory change. For ISR support, the Dockerfile would need to run the Nuxt server instead.

---

## Code Quality Standards

### Vue/Nuxt Frontend
- **Components**: Vue 3 Composition API with `<script setup lang="ts">`
- **Styling**: TailwindCSS utility classes; scoped `<style>` for component-specific styles
- **Formatting**: Prettier (100 char width, single quotes, no semicolons)
- **Linting**: ESLint with Vue + TypeScript rules
- **Type Safety**: Strict TypeScript — no `any` unless unavoidable
- **Content Sanitization**: Always use DOMPurify for user/API-provided HTML (`v-html`)

### Backend (Go)
- **Style**: Standard Go formatting (`gofmt`). No custom linter beyond `go vet`.
- **Module path**: `ivmanto.com/backend`. Internal packages live under `backend/internal/`.
- **Handlers**: Each domain (booking, blog, contact, ideas) has its own package in `backend/internal/` exporting an HTTP handler. Routes are registered in `backend/cmd/server/main.go`.
- **Env vars**: All env vars are loaded by `internal/config` at startup. Missing required vars cause a fatal exit logged as `"missing required environment variables"`. When you add a new env var:
  1. Wire it into `internal/config/config.go` (load + validate)
  2. Add a placeholder to `backend/.env.example`
  3. Add the substitution to `cloudbuild.yaml` OR add it to Secret Manager — state which in the PR
- **Error handling**: Wrap errors with context (`fmt.Errorf("doing X: %w", err)`). HTTP handlers respond with non-2xx + JSON body `{"error": "..."}` — do not panic.
- **Logging**: Use the structured logger middleware (`internal/middleware`). For ad-hoc logs use `log.Printf` with `key=value` pairs. No `fmt.Println` in handler code.
- **External clients**: GCS, Vertex AI, Calendar, SMTP clients are initialised once in `main.go` and injected into handlers. Do not create new clients per request.
- **Calendar / Meet**: All Google Calendar work goes through `internal/gcal`. Meet conferences require `ConferenceDataVersion(1)` on the Calendar update call. DWD impersonation is non-negotiable — do not switch to a different auth strategy without owner approval.
- **Tests**: Place tests next to the package (`foo_test.go`). Mock GCP clients via interfaces where already abstracted (see `internal/gcal` for the pattern). Don't add new test frameworks.

### SEO Standards
- Every page MUST have a unique `<title>` and `<meta name="description">`
- Pages not meant for indexing MUST have `<meta name="robots" content="noindex">`
- Canonical URLs MUST be set on every page
- Structured data (JSON-LD) MUST be present on all public pages
- Images MUST have `alt` attributes, `width`, and `height`

### Analytics Standards
- GTM initialization MUST be gated behind cookie consent
- All conversion events (booking, contact form, service views) MUST use `trackEvent()`
- GA session info MUST be forwarded to backend for server-side event stitching
- `window.dataLayer` MUST be initialized before any GTM code runs
