# Migration Plan: Vue 3 SPA → Nuxt 3 SSG

## Goal
Migrate ivmanto.com from a Vue 3 SPA (Vite) to Nuxt 3 with Static Site Generation (SSG) to solve the core SEO problem: crawlers and AI search engines cannot see content rendered only by JavaScript.

## Constraints
- **Zero content regression** — all text, styling, and user-facing behavior remains identical
- **Blog system preserved** — articles fetched at runtime from Go backend → GCS bucket pipeline
- **Booking system preserved** — real-time calendar availability from Go backend
- **Analytics preserved** — GTM/GA4 with cookie consent gate, GA session stitching
- **Backend unchanged** — Go backend continues to serve `/api/*` routes

---

## Phase 0: Preparation & Scaffolding
> Estimated effort: ~1 hour

- [ ] **0.1** Create a new git branch: `feat/nuxt-ssg-migration`
- [ ] **0.2** Scaffold a Nuxt 3 project alongside existing code:
  ```bash
  npx nuxi@latest init ivmanto-nuxt --template v3
  ```
  Then merge the generated config files (`nuxt.config.ts`, `tsconfig.json`, `app.vue`) into the existing repo root. Do NOT delete existing `src/` yet.
- [ ] **0.3** Update `package.json`:
  - Remove: `vue-router`, `@vueuse/head`, `serve`, `@vitejs/plugin-vue`, `vite`
  - Add: `nuxt` (^3.x), `@nuxtjs/tailwindcss`, `@nuxt/content` (if needed later)
  - Keep: `vue`, `axios`, `dompurify`, `tailwindcss`, `@tailwindcss/typography`, all dev deps (eslint, prettier, typescript, etc.)
  - Update scripts:
    ```json
    "dev": "nuxt dev",
    "build": "nuxt build",
    "generate": "nuxt generate",
    "preview": "nuxt preview",
    "postinstall": "nuxt prepare",
    "lint": "eslint . --fix",
    "format": "prettier --write ."
    ```
- [ ] **0.4** Create `nuxt.config.ts`:
  ```ts
  export default defineNuxtConfig({
    ssr: true,
    modules: ['@nuxtjs/tailwindcss'],
    css: ['~/assets/css/main.css'],
    runtimeConfig: {
      public: {
        geminiApiKey: '',
      }
    },
    nitro: {
      devProxy: {
        '/api': { target: 'http://localhost:8080', changeOrigin: true }
      }
    },
    routeRules: {
      // Pre-render static marketing pages
      '/': { prerender: true },
      '/about': { prerender: true },
      '/services': { prerender: true },
      '/services/**': { prerender: true },
      '/booking-demo': { prerender: true },
      '/login': { prerender: true },
      '/privacy-policy': { prerender: true },
      // Client-only for dynamic pages
      '/booking': { ssr: false },
      '/booking/cancel': { ssr: false },
      // Blog: client-rendered (articles come from runtime API)
      '/blog': { ssr: false },
      '/blog/**': { ssr: false },
    },
    app: {
      head: {
        htmlAttrs: { lang: 'en' },
        charset: 'utf-8',
        viewport: 'width=device-width, initial-scale=1',
        link: [
          { rel: 'icon', href: '/favicon.ico' },
          { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
          { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        ],
      }
    },
    typescript: {
      strict: true,
    },
  })
  ```
- [ ] **0.5** Move `tailwind.config.js` to project root (already there). Update `content` paths:
  ```js
  content: [
    './components/**/*.{vue,js,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './composables/**/*.ts',
    './plugins/**/*.ts',
    './app.vue',
  ]
  ```
- [ ] **0.6** Create `assets/css/main.css` from existing `src/style.css` (copy contents, same Tailwind directives + custom styles)
- [ ] **0.7** Run `npm install` and `npx nuxt prepare` to verify the scaffold compiles

---

## Phase 1: Layout & Core Structure
> Estimated effort: ~1 hour

- [ ] **1.1** Create `app.vue` (Nuxt entry point):
  ```vue
  <template>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </template>
  ```
- [ ] **1.2** Create `layouts/default.vue`:
  - Move global `useHead()` from existing `App.vue` (ProfessionalService schema, canonical URL, title, description)
  - Include `<AppHeader />`, `<slot />`, `<AppFooter />`, `<CookieBanner />`
  - The canonical URL logic should use `useRoute()` just as it does now
  - Replace `@vueuse/head`'s `useHead` with Nuxt's built-in `useHead`
  - Replace `usePageMetadata()` composable calls with Nuxt's `useSeoMeta()` where appropriate
- [ ] **1.3** Move layout components (no changes needed to templates/styles):
  - `src/components/layout/AppHeader.vue` → `components/layout/AppHeader.vue`
  - `src/components/layout/AppFooter.vue` → `components/layout/AppFooter.vue`
  - `src/components/layout/AppLogo.vue` → `components/layout/AppLogo.vue`
  - `src/components/layout/TheHeader.vue` → `components/layout/TheHeader.vue`
  - `src/components/layout/TheFooter.vue` → `components/layout/TheFooter.vue`
  - `src/components/CookieBanner.vue` → `components/CookieBanner.vue`
  - **In each component**: Remove explicit `import { RouterLink } from 'vue-router'` — Nuxt auto-imports `<NuxtLink>`. Replace all `<router-link>` and `<RouterLink>` with `<NuxtLink>`.
  - **In AppHeader.vue**: Replace `useRoute()` import with auto-imported `useRoute()`

### 1.4 NuxtLink Migration Rules (applies to ALL components)
When migrating any component:
1. Remove `import { RouterLink } from 'vue-router'` or `import { RouterLink, useRoute } from 'vue-router'`
2. Replace `<RouterLink>` / `<router-link>` with `<NuxtLink>`
3. `useRoute()` and `useRouter()` are auto-imported in Nuxt — remove their imports
4. Named routes (`{ name: 'booking' }`) work the same way in Nuxt
5. Hash links (`/#contact`) work the same way

---

## Phase 2: Move Components
> Estimated effort: ~1.5 hours

All components move from `src/components/` → `components/`. The key changes are:
- Remove `RouterLink` imports (Nuxt auto-imports)
- Replace `<RouterLink>` with `<NuxtLink>`
- Remove `useRoute`/`useRouter` imports (Nuxt auto-imports)
- Replace `@/` path aliases with `~/` (Nuxt convention) or rely on auto-imports

### Section Components
- [ ] **2.1** Move `src/components/sections/` → `components/sections/`:
  - `HeroInfographicSection.vue` — no changes needed (pure template)
  - `ProcessSection.vue` — no changes needed (pure template)
  - `FAQSection.vue` — no changes needed (self-contained logic)
  - `ContactSection.vue` — if exists, move as-is
  - `HeroSection.vue` — if exists, move as-is
  - `AboutSection.vue` — if exists, move as-is
  - `ArticlesSection.vue` — if exists, move as-is
  - `ServicesIndex.vue` — if exists, move as-is
  - `ServicesSection.vue` — if exists, move as-is

### About Components
- [ ] **2.2** Move `src/components/about/` → `components/about/`:
  - `ApproachSection.vue` — no changes
  - `ExpertiseSection.vue` — no changes
  - `CertificatesSection.vue` — no changes

### Service Content Components
- [ ] **2.3** Move `src/components/services-content/` → `components/services-content/`:
  - `DataArchitecture.vue` — update: replace `usePageMetadata()` import with Nuxt auto-import or remove if unused
  - `DataGovernance.vue` — same
  - `MlEngineering.vue` — same
  - `Principles.vue` — same
  - `SovereignCloudDE.vue` — same

### Other Components
- [ ] **2.4** Move remaining components:
  - `ContactForm.vue` — replace `useRoute` import, replace `RouterLink` → `NuxtLink`
  - `InspirationModal.vue` — replace `RouterLink` → `NuxtLink`, `<teleport to="body">` works in Nuxt
  - `PrivacyPolicy.vue` — if used as a component, move as-is
  - `services/RightColumnContent.vue` — move as-is
  - `services/ServiceDetail.vue` — move as-is

---

## Phase 3: Move Composables, Services, Data, Types
> Estimated effort: ~30 minutes

### Composables
- [ ] **3.1** Move `src/composables/` → `composables/`:
  - `useArticles.ts` — keep as-is. The `fetch('/api/articles')` calls will work from client-side. For SSR-compatible version, consider wrapping with `useFetch()` later.
  - `useGemini.ts` — update: replace `import.meta.env.VITE_GEMINI_API_KEY` with `useRuntimeConfig().public.geminiApiKey`
  - `usePageMetadata.ts` — **KEEP for now** as a compatibility layer. Eventually replace with `useSeoMeta()` per-page, but during migration, keeping it working prevents regressions.

### Services
- [ ] **3.2** Move `src/services/` → `services/` (NOT in composables — these are plain utility modules, not composables):
  - `analytics.ts` — no changes needed (already has `typeof window === 'undefined'` guards)
  - `api.ts` — no changes needed
  - `gemini.ts` — update: replace `import.meta.env.VITE_GEMINI_API_KEY` with runtime config access. **Note**: Since this is not a composable, it cannot call `useRuntimeConfig()` directly. Either:
    - Convert to accept `apiKey` as a parameter, or
    - Move the API key lookup to the calling component

### Data & Types
- [ ] **3.3** Move `src/data/services.ts` → `data/services.ts`
  - Update `defineAsyncComponent` imports: In Nuxt, components in `components/` are auto-imported. The `defineAsyncComponent(() => import(...))` pattern still works, but paths change from `@/components/` to `~/components/` or use `resolveComponent()`.
- [ ] **3.4** Move `src/types/article.ts` → `types/article.ts`

---

## Phase 4: Convert Views → Pages
> Estimated effort: ~2 hours (the core of the migration)

### 4.1 Home Page
- [ ] **4.1.1** Create `pages/index.vue` from `src/views/HomeView.vue`:
  - Move all content as-is
  - Replace imports: `RouterLink` → `NuxtLink`, remove `useRoute`/`useRouter` imports
  - Replace `@/` paths with `~/` paths
  - Add page-level `useSeoMeta()`:
    ```ts
    useSeoMeta({
      title: 'ivmanto.com | Data & AI Consultancy',
      description: 'Expert Data & AI consultancy...',
    })
    ```
  - **Add FAQPage schema** (NEW — from SEO recommendation):
    ```ts
    useHead({
      script: [{
        type: 'application/ld+json',
        children: JSON.stringify({
          '@context': 'https://schema.org',
          '@type': 'FAQPage',
          mainEntity: [
            { '@type': 'Question', name: '...', acceptedAnswer: { '@type': 'Answer', text: '...' } },
            // ... all 4 FAQ items
          ]
        })
      }]
    })
    ```

### 4.2 About Page
- [ ] **4.2.1** Create `pages/about.vue` from `src/views/AboutView.vue`:
  - Keep Person schema as-is
  - Replace `useHead` import → use Nuxt auto-imported version
  - Replace `RouterLink` → `NuxtLink`
  - Fix image path: `src="/src/assets/nt.jpg"` → `src="~/assets/nt.jpg"` (or move image to `public/` and use `/nt.jpg`)

### 4.3 Services Pages
- [ ] **4.3.1** Create `pages/services/index.vue` from `src/views/ServicesLanding.vue`:
  - Replace `RouterLink` → `NuxtLink`
  - Replace `@/data/services` → `~/data/services`
- [ ] **4.3.2** Create `pages/services/[id].vue` from `src/views/ServiceView.vue`:
  - Replace `props.id` (from vue-router `props: true`) with `useRoute().params.id`
  - Keep the dynamic Service schema markup
  - Replace `@vueuse/head` → Nuxt `useHead`
  - Replace `RouterLink` → `NuxtLink`
  - Update `defineAsyncComponent` paths

### 4.4 Blog Pages
- [ ] **4.4.1** Create `pages/blog/index.vue` from `src/views/ArticleListView.vue`:
  - This page is client-rendered (`ssr: false` in routeRules)
  - Replace `RouterLink` → `NuxtLink`
  - Keep `useArticles()` composable call as-is
- [ ] **4.4.2** Create `pages/blog/[slug].vue` from `src/views/ArticleView.vue`:
  - This page is client-rendered (`ssr: false` in routeRules)
  - Replace `props.slug` with `useRoute().params.slug`
  - Replace `useRouter().push()` with `navigateTo()`
  - Replace `RouterLink` → `NuxtLink`
  - Keep DOMPurify sanitization
  - Keep like functionality and article navigation
  - **Add Article/BlogPosting schema** (NEW):
    ```ts
    useHead({
      script: computed(() => article.value ? [{
        type: 'application/ld+json',
        children: JSON.stringify({
          '@context': 'https://schema.org',
          '@type': 'BlogPosting',
          headline: article.value.title,
          description: article.value.summary,
          datePublished: article.value.date,
          author: { '@id': 'https://ivmanto.com/about#person' },
          publisher: { '@id': 'https://ivmanto.com/#organization' },
          url: `https://ivmanto.com/blog/${article.value.slug}`,
        })
      }] : [])
    })
    ```
- [ ] **4.4.3** The `BlogView.vue` wrapper (just `<RouterView />`) is no longer needed — Nuxt handles nested routing automatically via the `pages/blog/` directory structure.

### 4.5 Booking Pages
- [ ] **4.5.1** Create `pages/booking/index.vue` from `src/views/BookingCalendar.vue`:
  - Client-rendered (`ssr: false`)
  - Replace `useRoute`/`RouterLink` imports
  - Keep all booking logic exactly as-is
  - The `Intl.DateTimeFormat` guard for timezone is already SSR-safe
- [ ] **4.5.2** Create `pages/booking/cancel.vue` from `src/views/BookingCancellation.vue`:
  - Client-rendered (`ssr: false`)
  - Replace imports

### 4.6 Other Pages
- [ ] **4.6.1** Create `pages/booking-demo.vue` from `src/views/BookingGoogleDemo.vue`
- [ ] **4.6.2** Create `pages/login.vue` from `src/views/LoginView.vue`:
  - Replace `@vueuse/head` → Nuxt `useHead`
  - Keep `noindex` meta
- [ ] **4.6.3** Create `pages/privacy-policy.vue`:
  - Import `PrivacyPolicy` component and render it
  - Or inline the content directly
- [ ] **4.6.4** Create `pages/[...slug].vue` (catch-all 404) from `src/views/NotFoundView.vue`:
  - Replace `@vueuse/head` → Nuxt `useHead`
  - Keep `noindex` meta

### 4.7 Scroll Behavior
- [ ] **4.7.1** Replicate the scroll behavior from vue-router config. In Nuxt, create `app/router.options.ts`:
  ```ts
  import type { RouterConfig } from '@nuxt/schema'

  export default <RouterConfig>{
    scrollBehavior(to) {
      if (to.hash) {
        return { el: to.hash, behavior: 'smooth' }
      }
      return { top: 0 }
    }
  }
  ```

---

## Phase 5: Plugins & Analytics
> Estimated effort: ~30 minutes

- [ ] **5.1** Create `plugins/analytics.client.ts`:
  ```ts
  export default defineNuxtPlugin(() => {
    // Initialize dataLayer
    window.dataLayer = window.dataLayer || []

    const consent = localStorage.getItem('cookie_consent')
    if (consent === 'accepted') {
      const { initGtm, trackEvent } = await import('~/services/analytics')
      initGtm()

      // Track initial pageview
      trackEvent('page_view', {
        page_path: window.location.pathname + window.location.search,
        page_title: document.title,
      })

      // Track SPA pageviews on route change
      const router = useRouter()
      router.afterEach((to) => {
        trackEvent('page_view', {
          page_path: to.fullPath,
          page_title: document.title,
        })
      })
    }
  })
  ```
  The `.client.ts` suffix ensures this only runs in the browser.

- [ ] **5.2** Remove analytics initialization from what was `main.ts` — it's now handled by the plugin.

- [ ] **5.3** Verify `window.dataLayer` type declaration still works. Move `declare global` block from `analytics.ts` to a `types/global.d.ts` if needed.

---

## Phase 6: Structured Data Enhancement (SEO Recommendation)
> Estimated effort: ~30 minutes

- [ ] **6.1** Add **WebSite schema** to `layouts/default.vue`:
  ```ts
  useHead({
    script: [{
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'WebSite',
        name: 'IVMANTO',
        url: 'https://ivmanto.com',
        potentialAction: {
          '@type': 'SearchAction',
          target: 'https://ivmanto.com/blog?q={search_term_string}',
          'query-input': 'required name=search_term_string'
        }
      })
    }]
  })
  ```
  This tells search engines the site has a search feature pointing to the blog's search input.

- [ ] **6.2** Add **FAQPage schema** to `pages/index.vue` (home page) — as described in Phase 4.1.

- [ ] **6.3** Add **BlogPosting schema** to `pages/blog/[slug].vue` — as described in Phase 4.4.2.

- [ ] **6.4** Verify existing schemas still work:
  - ProfessionalService (global layout)
  - Person (about page)
  - Service (service detail pages)

---

## Phase 7: Static Assets & Public Files
> Estimated effort: ~15 minutes

- [ ] **7.1** Move/verify public assets:
  - `public/certs/` — keep as-is
  - `public/favicon.ico` — keep (or copy from `src/favicon.ico`)
  - `public/cloud-pic-2.webp` — verify exists
  - `public/logo.png` — verify exists (referenced in schema)
  - `public/social-sharing-card.webp` — verify exists (referenced in schema)
- [ ] **7.2** Move `src/assets/nt.jpg` → `public/nt.jpg` (simpler) or keep in `assets/` and use `~/assets/nt.jpg` in the about page template
- [ ] **7.3** Move `src/assets/mockup/logo.svg` → `assets/mockup/logo.svg`
- [ ] **7.4** Verify Google Fonts loading. The `<link>` tags from `index.html` should be moved to `nuxt.config.ts` → `app.head.link` (already covered in Phase 0.4)

---

## Phase 8: Dockerfile Update
> Estimated effort: ~15 minutes

- [ ] **8.1** Update `Dockerfile` for Nuxt SSG:
  ```dockerfile
  # Stage 1: Build
  FROM node:20-bookworm-slim AS build
  RUN apt-get update && apt-get upgrade -y && \
      apt-get clean && rm -rf /var/lib/apt/lists/*
  WORKDIR /app
  COPY package.json package-lock.json ./
  RUN npm install
  COPY . .
  RUN npx nuxi generate

  # Stage 2: Serve
  FROM node:20-bookworm-slim
  RUN apt-get update && apt-get upgrade -y && \
      apt-get clean && rm -rf /var/lib/apt/lists/*
  RUN npm install -g serve
  WORKDIR /app
  COPY --from=build /app/.output/public .
  EXPOSE 8080
  CMD ["serve", "-s", "."]
  ```
  **Key change**: `npm run build` → `npx nuxi generate`, `/app/dist` → `/app/.output/public`

---

## Phase 9: Cleanup & Verification
> Estimated effort: ~1 hour

- [ ] **9.1** Delete old files:
  - `src/` directory (all files have been migrated)
  - `src/main.ts` (replaced by Nuxt entry)
  - `src/App.vue` (replaced by `app.vue` + `layouts/default.vue`)
  - `src/router/index.ts` (replaced by file-based routing)
  - `vite.config.ts` (replaced by `nuxt.config.ts`)
  - `env.d.ts` (Nuxt generates its own)
- [ ] **9.2** Remove unused dependencies from `package.json`:
  - `vue-router` (Nuxt includes it)
  - `@vueuse/head` (Nuxt has built-in head management)
  - `serve` (keep — still used in Docker)
  - `@vitejs/plugin-vue` (Nuxt includes Vite + Vue plugin)
- [ ] **9.3** Run `npm run generate` and verify output:
  - Check `.output/public/` contains HTML files for all pre-rendered routes
  - Verify each HTML file has proper `<title>`, `<meta>`, and JSON-LD schema
- [ ] **9.4** Run `npm run preview` and test every page manually:
  - [ ] Home page — hero, process section, articles, FAQ, contact form, inspiration modal
  - [ ] About page — photo, content, certifications, Person schema
  - [ ] Services landing — all 5 service cards
  - [ ] Each service detail page — sidebar, content, keywords, CTA, Service schema
  - [ ] Blog list — search, article cards, loading states
  - [ ] Blog article — content rendering, navigation, share, likes
  - [ ] Booking calendar — date selection, time slots, form submission
  - [ ] Booking cancellation — token-based cancellation
  - [ ] Login — noindex, placeholder
  - [ ] Privacy policy — full legal text
  - [ ] 404 page — catch-all, noindex
  - [ ] Cookie banner — shows on first visit, consent gating
  - [ ] Mobile responsive — header menu, all pages
- [ ] **9.5** Validate structured data:
  - Use https://validator.schema.org/ on the generated HTML files
  - Verify ProfessionalService, WebSite, Person, Service, FAQPage schemas
- [ ] **9.6** Test the API proxy in dev mode:
  - `npm run dev` with Go backend running
  - Verify `/api/articles`, `/api/booking/availability`, `/api/contact` all work
- [ ] **9.7** Build Docker image and test:
  ```bash
  docker build -t ivmanto-nuxt .
  docker run -p 8080:8080 ivmanto-nuxt
  ```

---

## Phase 10: Blog Route Strategy Decision
> This is a decision point — choose before starting Phase 4.4

### Option A: Pure Client-Side Blog (Simpler, matches current behavior)
- Blog pages use `ssr: false` in `routeRules`
- Articles are fetched client-side exactly like today
- **Pro**: Simplest migration, no behavior change
- **Con**: Blog pages still serve empty HTML to crawlers (same SEO limitation as SPA)
- **Mitigation**: Blog content is behind a dynamic API, so search engines may not index individual articles regardless

### Option B: Nuxt Server Mode for Blog (Better SEO, more complex)
- Switch from `nuxt generate` (pure SSG) to `nuxt build` (server mode) for the whole site
- Blog pages use ISR: `'/blog/**': { isr: 3600 }` (regenerate every hour)
- Requires changing Dockerfile to run `node .output/server/index.mjs` instead of `serve`
- **Pro**: Blog articles get full server-rendered HTML with schema markup
- **Con**: Requires a running Node.js server (not just static files), more complex deployment

### Recommendation
Start with **Option A** (pure client-side blog). This achieves the primary goal — making all marketing pages (home, about, services) crawlable with full HTML and schema. The blog is secondary for SEO since its content is dynamic. Option B can be pursued later as an enhancement.

---

## Summary of New Files Created

| File | Purpose |
|---|---|
| `nuxt.config.ts` | Nuxt configuration (replaces vite.config.ts) |
| `app.vue` | Nuxt entry point |
| `layouts/default.vue` | Global layout (header, footer, schema, cookie banner) |
| `pages/index.vue` | Home page |
| `pages/about.vue` | About page |
| `pages/services/index.vue` | Services landing |
| `pages/services/[id].vue` | Service detail |
| `pages/blog/index.vue` | Blog article list |
| `pages/blog/[slug].vue` | Blog article detail |
| `pages/booking/index.vue` | Booking calendar |
| `pages/booking/cancel.vue` | Booking cancellation |
| `pages/booking-demo.vue` | Google Calendar embed demo |
| `pages/login.vue` | Client login placeholder |
| `pages/privacy-policy.vue` | Privacy policy |
| `pages/[...slug].vue` | Catch-all 404 |
| `plugins/analytics.client.ts` | GTM/GA4 initialization |
| `app/router.options.ts` | Scroll behavior |
| `assets/css/main.css` | Tailwind + custom styles |

## Files to Delete After Migration

| File | Replaced By |
|---|---|
| `src/main.ts` | Nuxt auto-initialization + `plugins/analytics.client.ts` |
| `src/App.vue` | `app.vue` + `layouts/default.vue` |
| `src/router/index.ts` | File-based routing in `pages/` |
| `vite.config.ts` | `nuxt.config.ts` |
| `env.d.ts` | Nuxt auto-generated types |
| All `src/views/*.vue` | `pages/*.vue` |
| All `src/components/` | `components/` (moved, not deleted) |
| All `src/composables/` | `composables/` (moved, not deleted) |
| All `src/services/` | `services/` (moved, not deleted) |
| All `src/data/` | `data/` (moved, not deleted) |
| All `src/types/` | `types/` (moved, not deleted) |
