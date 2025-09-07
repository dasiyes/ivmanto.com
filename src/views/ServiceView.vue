<template>
  <div class="container mx-auto px-6 py-12">
    <div class="flex flex-col md:flex-row gap-12">
      <!-- Left Column: Sidebar Navigation -->
      <aside class="w-full md:w-1/3 lg:w-1/4 flex-shrink-0">
        <h2 class="text-xl font-bold text-dark-slate mb-4 border-b pb-2">Our Services</h2>
        <nav class="space-y-2">
          <!-- The sidebar is dynamically generated from services.ts -->
          <RouterLink
            v-for="s in allServices"
            :key="s.id"
            :to="`/services/${s.id}`"
            class="block p-3 -m-3 rounded-lg transition-colors"
            :class="{
              'bg-light-gray text-primary': s.id === id,
              'hover:bg-gray-50': s.id !== id,
            }"
          >
            <p class="font-semibold text-dark-slate">{{ s.menuTitle }}</p>
            <p class="text-sm text-gray-600 mt-1">{{ s.summary }}</p>
          </RouterLink>
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
            <router-link
              :to="{ name: 'booking' }"
              class="bg-white text-primary font-bold py-2 px-5 rounded-lg hover:bg-gray-100 transition-colors whitespace-nowrap"
            >
              Book a Consultation
            </router-link>
          </div>
        </div>
        <div v-else class="text-center p-12">
          <h1 class="text-2xl font-bold">Service Not Found</h1>
          <p class="mt-4">The service you are looking for does not exist.</p>
          <RouterLink to="/" class="text-primary mt-6 inline-block">Go back to Home</RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { useHead } from '@vueuse/head'
import { services as allServices, getServiceById } from '@/data/services'

const props = defineProps<{
  // This `id` is passed automatically by the router because of `props: true`
  id: string
}>()

const route = useRoute()
const service = computed(() => getServiceById(props.id))

const serviceSchema = computed(() => {
  if (!service.value) {
    return null
  }
  return {
    '@context': 'https://schema.org',
    '@type': 'Service',
    serviceType: service.value.menuTitle,
    name: service.value.menuTitle,
    description: service.value.summary,
    provider: {
      '@id': 'https://ivmanto.com/#organization',
    },
    areaServed: {
      '@type': 'Country',
      name: 'Global', // You can change this if you serve specific regions
    },
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
      // Use a key to ensure the script tag is updated when the route changes
      key: () => `service-schema-${props.id}`,
    },
  ],
})
</script>

<style scoped>
/* Add active styles for the sidebar navigation */
.router-link-exact-active {
  background-color: #f8f9fa; /* bg-light-gray */
  color: #00a896; /* text-primary */
}
</style>
