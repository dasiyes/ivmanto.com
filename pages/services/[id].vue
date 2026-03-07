<template>
  <div class="container mx-auto px-6 py-12">
    <div class="flex flex-col md:flex-row gap-12">
      <!-- Left Column: Sidebar Navigation -->
      <aside class="w-full md:w-1/3 lg:w-1/4 flex-shrink-0">
        <h2 class="text-xl font-bold text-dark-slate mb-4 border-b pb-2">Our Services</h2>
        <nav class="space-y-2">
          <NuxtLink
            v-for="s in allServices"
            :key="s.id"
            :to="`/services/${s.id}`"
            @click="trackServiceClick(s)"
            class="block p-3 -m-3 rounded-lg transition-colors"
            :class="{
              'bg-light-gray text-primary': s.id === id,
              'hover:bg-gray-50': s.id !== id,
            }"
          >
            <p class="font-semibold text-dark-slate">{{ s.menuTitle }}</p>
            <p class="text-sm text-gray-600 mt-1">{{ s.summary }}</p>
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
            <!-- Core Service Article -->
            <main
              class="w-full lg:w-2/3 bg-white border border-gray-200 rounded-xl shadow-sm overflow-hidden"
            >
              <component :is="service.detailsComponent" />
            </main>

            <!-- Right Sidebar: Keywords -->
            <aside
              v-if="service.tagDetails && Object.keys(service.tagDetails).length"
              class="w-full lg:w-1/3"
            >
              <div class="p-4 bg-light-gray rounded-lg sticky top-24">
                <h3 class="font-bold text-dark-slate mb-4">Key Concepts</h3>
                <div class="space-y-4">
                  <div v-for="(desc, tag) in service.tagDetails" :key="tag">
                    <p class="font-semibold text-primary">{{ tag }}</p>
                    <p class="text-sm text-gray-600">{{ desc }}</p>
                  </div>
                </div>
              </div>
            </aside>
          </div>

          <!-- Bottom Bar: CTA -->
          <div class="p-6 bg-primary text-white rounded-lg flex justify-between items-center">
            <div>
              <h3 class="font-bold text-xl">Ready to build your data foundation?</h3>
              <p>Let's discuss how these services can be tailored to your business.</p>
            </div>
            <NuxtLink
              :to="{ name: 'booking' }"
              @click="trackBookConsultationClick"
              class="bg-white text-primary font-bold py-2 px-5 rounded-lg hover:bg-gray-100 transition-colors whitespace-nowrap"
            >
              Book a Consultation
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
                <span class="mt-2 inline-block text-primary font-semibold text-sm">Read Article &rarr;</span>
              </NuxtLink>
            </div>
          </div>
        </div>
        <div v-else class="text-center p-12">
          <h1 class="text-2xl font-bold">Service Not Found</h1>
          <p class="mt-4">The service you are looking for does not exist.</p>
          <NuxtLink to="/" class="text-primary mt-6 inline-block">Go back to Home</NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { services as allServices, getServiceById, type Service } from '~/data/services'
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

// Page-level SEO metadata
const routeMetadata: Record<string, { title: string; description: string }> = {
  'data-strategy-and-governance': {
    title: 'Data Strategy & Governance | ivmanto.com',
    description:
      'Develop a clear data strategy and robust governance framework. We align your data initiatives with business goals for maximum impact and compliance.',
  },
  'data-architecture': {
    title: 'Data Architecture on GCP | ivmanto.com',
    description:
      'Design and build scalable, secure data architectures on Google Cloud Platform (GCP). We leverage BigQuery, Cloud Storage, and modern data engineering practices.',
  },
  sovereigncloud: {
    title: 'Sovereign Cloud Solutions | ivmanto.com',
    description:
      'Explore architectural perspectives on Data, Operations, and AI Sovereignty to meet your compliance and security needs in the cloud.',
  },
  'ml-engineering': {
    title: 'AI & ML Solutions | ivmanto.com',
    description:
      'Leverage the power of AI and Machine Learning on GCP. We build custom solutions, from predictive analytics to generative AI, to solve your toughest challenges.',
  },
  principles: {
    title: 'Guiding Principles | ivmanto.com',
    description:
      'Our DAMA-aligned principles for data strategy, governance, and architecture ensure your data becomes a reliable, valuable asset for decision-making and AI.',
  },
}

const meta = computed(() => routeMetadata[id.value])

useSeoMeta({
  title: computed(() => meta.value?.title ?? 'Services | ivmanto.com'),
  description: computed(() => meta.value?.description ?? ''),
})

const serviceSchema = computed(() => {
  if (!service.value) return null
  return {
    '@context': 'https://schema.org',
    '@type': 'Service',
    serviceType: service.value.menuTitle,
    name: service.value.menuTitle,
    description: service.value.summary,
    provider: { '@id': 'https://ivmanto.com/#organization' },
    areaServed: { '@type': 'Country', name: 'Global' },
    url: `https://ivmanto.com${route.path}`,
  }
})

useHead({
  script: [
    {
      id: 'service-schema',
      type: 'application/ld+json',
      children: computed(() =>
        serviceSchema.value ? JSON.stringify(serviceSchema.value, null, 2) : '',
      ),
    },
  ],
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
