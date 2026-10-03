<template>
  <div class="container mx-auto px-6 py-12">
    <div class="flex flex-col md:flex-row gap-12">
      <!-- Left Column: Sidebar Navigation -->
      <aside class="w-full md:w-1/3 lg:w-1/4 flex-shrink-0">
        <h2 class="text-xl font-bold text-dark-slate mb-4 border-b pb-2">Services</h2>
        <nav class="space-y-2">
          <NuxtLink
            v-for="s in offerServices"
            :key="s.id"
            :to="`/services/${s.id}`"
            @click="trackServiceClick(s)"
            class="block p-3 -m-3 rounded-lg transition-colors"
            :class="{
              'bg-light-gray text-primary': s.id === id,
              'hover:bg-gray-50': s.id !== id,
            }"
          >
            <p v-if="s.eyebrow" class="text-xs font-bold uppercase tracking-wider text-primary">
              {{ s.eyebrow }}
            </p>
            <p class="font-semibold text-dark-slate">{{ s.menuTitle }}</p>
            <p v-if="s.price" class="text-sm font-semibold text-amber-dark mt-1">
              {{ s.price.label }}
            </p>
          </NuxtLink>
        </nav>

        <h2 class="text-sm font-bold uppercase tracking-wider text-gray-500 mt-10 mb-3">
          Expertise
        </h2>
        <nav class="space-y-1">
          <NuxtLink
            v-for="s in expertiseServices"
            :key="s.id"
            :to="`/services/${s.id}`"
            @click="trackServiceClick(s)"
            class="block px-3 py-1.5 -mx-3 rounded-lg text-sm transition-colors"
            :class="{
              'bg-light-gray text-primary font-semibold': s.id === id,
              'text-gray-700 hover:bg-gray-50': s.id !== id,
            }"
          >
            {{ s.menuTitle }}
          </NuxtLink>
        </nav>
      </aside>

      <!-- Right Column: Main content area -->
      <div class="w-full md:w-2/3 lg:w-3/4">
        <div v-if="service" class="space-y-8">
          <!-- Top Bar: Industries -->
          <div class="p-4 bg-light-gray rounded-lg">
            <h3 class="font-semibold text-dark-slate mb-2">Relevant Industries</h3>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="industry in service.industries"
                :key="industry"
                class="bg-white text-primary text-sm font-medium px-3 py-1 rounded-full border border-gray-200"
              >
                {{ industry }}
              </span>
            </div>
          </div>

          <!-- Main content with right sidebar for keywords -->
          <div class="flex flex-col lg:flex-row gap-8">
            <main
              class="w-full lg:w-2/3 bg-white border border-gray-200 rounded-xl shadow-sm overflow-hidden"
            >
              <component :is="service.detailsComponent" v-if="service.detailsComponent" />
              <ServiceOffer v-else :service="service" />
            </main>

            <!-- Right Sidebar: Key Concepts & Topics -->
            <aside
              v-if="
                (service.tagDetails && Object.keys(service.tagDetails).length) ||
                service.keywords?.length
              "
              class="w-full lg:w-1/3"
            >
              <div class="p-4 bg-light-gray rounded-lg sticky top-24 space-y-6">
                <div v-if="service.tagDetails && Object.keys(service.tagDetails).length">
                  <h3 class="font-bold text-dark-slate mb-4">Key Concepts</h3>
                  <div class="space-y-4">
                    <div v-for="(desc, tag) in service.tagDetails" :key="tag">
                      <p class="font-semibold text-primary">{{ tag }}</p>
                      <p class="text-sm text-gray-600">{{ desc }}</p>
                    </div>
                  </div>
                </div>
                <div v-if="service.keywords?.length">
                  <h3 class="font-bold text-dark-slate mb-3">Topics</h3>
                  <div class="flex flex-wrap gap-2">
                    <span
                      v-for="keyword in service.keywords"
                      :key="keyword"
                      class="text-xs font-medium text-primary-dark bg-white border border-gray-200 rounded-full px-2.5 py-1"
                    >
                      {{ keyword }}
                    </span>
                  </div>
                </div>
              </div>
            </aside>
          </div>

          <!-- Bottom Bar: CTA -->
          <div
            class="p-6 bg-primary text-white rounded-lg flex flex-col sm:flex-row gap-4 justify-between sm:items-center"
          >
            <div>
              <h3 class="font-bold text-xl">{{ cta.title }}</h3>
              <p>{{ cta.text }}</p>
            </div>
            <NuxtLink
              :to="{ name: 'booking' }"
              @click="trackBookConsultationClick"
              class="bg-white text-primary font-bold py-2 px-5 rounded-lg hover:bg-gray-100 transition-colors whitespace-nowrap text-center"
            >
              {{ cta.button }}
            </NuxtLink>
          </div>

          <!-- Related Articles -->
          <div v-if="relatedArticles.length > 0" class="mt-8">
            <h3 class="text-xl font-bold text-dark-slate mb-4">Related Articles</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <NuxtLink
                v-for="article in relatedArticles"
                :key="article.slug"
                :to="`/blog/${article.slug}`"
                class="block p-4 bg-light-gray rounded-lg hover:shadow-md transition-all hover:-translate-y-0.5"
              >
                <h4 class="font-semibold text-dark-slate">{{ article.title }}</h4>
                <p class="text-sm text-gray-600 mt-1 line-clamp-2">{{ article.summary }}</p>
                <span class="mt-2 inline-block text-primary font-semibold text-sm"
                  >Read: {{ article.title }} &rarr;</span
                >
              </NuxtLink>
            </div>
          </div>
        </div>
        <div v-else class="text-center p-12">
          <h1 class="text-2xl font-bold">Service Not Found</h1>
          <p class="mt-4">The service you are looking for does not exist.</p>
          <NuxtLink to="/services" class="text-primary mt-6 inline-block"
            >See all services</NuxtLink
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ServiceOffer from '~/components/services/ServiceOffer.vue'
import { offerServices, expertiseServices, getServiceById, type Service } from '~/data/services'
import { trackEvent } from '~/services/analytics'

const route = useRoute()
const id = computed(() => route.params.id as string)
const service = computed(() => getServiceById(id.value))

// Fetch articles for related articles section
const { fetchArticles, getArticleBySlug } = useArticles()
await fetchArticles()

const relatedArticles = computed(() => {
  if (!service.value) return []
  return service.value.relatedBlogSlugs
    .map((slug) => getArticleBySlug(slug))
    .filter((a): a is NonNullable<typeof a> => a != null)
})

const cta = computed(() =>
  service.value?.tier === 'entry'
    ? {
        title: 'Book your readiness audit',
        text: 'From €100, fully credited when we sign the project. Start with a free 30-minute call.',
        button: 'Book a Scoping Call',
      }
    : {
        title: 'Not sure where to start?',
        text: 'Book a free 30-minute scoping call. We talk directly about feasibility, timeline and cost.',
        button: 'Book a Scoping Call',
      },
)

// Page-level SEO metadata
useSeoMeta({
  title: computed(() => service.value?.seoTitle ?? 'Services | ivmanto.com'),
  description: computed(() => service.value?.seoDescription ?? ''),
  ogTitle: computed(() => service.value?.seoTitle ?? 'Services | ivmanto.com'),
  ogDescription: computed(() => service.value?.seoDescription ?? ''),
})

const areaServed = [
  { '@type': 'Place', name: 'European Union' },
  { '@type': 'Country', name: 'Germany' },
  { '@type': 'Country', name: 'Austria' },
  { '@type': 'Country', name: 'Switzerland' },
  { '@type': 'Place', name: 'Worldwide' },
]

const serviceSchema = computed(() => {
  const s = service.value
  if (!s) return null
  return {
    '@context': 'https://schema.org',
    '@type': 'Service',
    serviceType: s.menuTitle,
    name: s.menuTitle,
    description: s.seoDescription,
    category: 'Data & AI Engineering',
    provider: { '@id': 'https://ivmanto.com/#organization' },
    areaServed,
    url: `https://ivmanto.com${route.path}`,
    ...(s.keywords?.length ? { keywords: s.keywords.join(', ') } : {}),
    ...(s.price
      ? {
          offers: {
            '@type': 'Offer',
            description: s.price.note,
            priceSpecification: {
              '@type': 'PriceSpecification',
              minPrice: s.price.amount,
              priceCurrency: s.price.currency,
            },
          },
        }
      : {}),
  }
})

const faqSchema = computed(() => {
  const faqs = service.value?.faqs
  if (!faqs?.length) return null
  return {
    '@context': 'https://schema.org',
    '@type': 'FAQPage',
    mainEntity: faqs.map((f) => ({
      '@type': 'Question',
      name: f.question,
      acceptedAnswer: { '@type': 'Answer', text: f.answer },
    })),
  }
})

useHead({
  script: computed(() => {
    const scripts = []
    if (serviceSchema.value) {
      scripts.push({
        id: 'service-schema',
        type: 'application/ld+json',
        innerHTML: JSON.stringify(serviceSchema.value, null, 2),
      })
    }
    if (faqSchema.value) {
      scripts.push({
        id: 'service-faq-schema',
        type: 'application/ld+json',
        innerHTML: JSON.stringify(faqSchema.value, null, 2),
      })
    }
    return scripts
  }),
})

function trackServiceClick(service: Service) {
  trackEvent('view_service_details', {
    service_id: service.id,
    service_name: service.menuTitle,
  })
}

function trackBookConsultationClick() {
  trackEvent('click_book_consultation', {
    source: 'service_page_cta',
    service_id: id.value,
  })
}
</script>

<style scoped>
.router-link-exact-active {
  background-color: #f8f9fa;
  color: #00a896;
}
</style>
