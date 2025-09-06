<template>
  <div class="container mx-auto px-6 py-12">
    <div class="flex flex-col md:flex-row gap-12">
      <!-- Left Column: Sidebar Navigation -->
      <aside class="w-full md:w-1/3 lg:w-1/4 flex-shrink-0">
        <h2 class="text-xl font-bold text-dark-slate mb-4 border-b pb-2">Our Services</h2>
        <nav class="space-y-2">
          <!-- The sidebar is dynamically generated from services.ts -->
          <RouterLink
            v-for="service in allServices"
            :key="service.id"
            :to="`/services/${service.id}`"
            class="block p-3 -m-3 rounded-lg transition-colors"
            :class="{
              'bg-light-gray text-primary': service.id === id,
              'hover:bg-gray-50': service.id !== id,
            }"
          >
            <p class="font-semibold text-dark-slate">{{ service.menuTitle }}</p>
            <p class="text-sm text-gray-600 mt-1">{{ service.summary }}</p>
          </RouterLink>
        </nav>
      </aside>

      <!-- Right Column: Dynamic Service Content -->
      <main class="w-full md:w-2/3 lg:w-3/4">
        <div v-if="serviceComponent" class="bg-white border border-gray-200 rounded-xl shadow-sm">
          <!-- This renders the correct component, e.g., DataArchitecture.vue -->
          <component :is="serviceComponent" />
        </div>
        <div v-else class="text-center p-12">
          <h1 class="text-2xl font-bold">Service Not Found</h1>
          <p class="mt-4">The service you are looking for does not exist.</p>
          <RouterLink to="/" class="text-primary mt-6 inline-block">Go back to Home</RouterLink>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink } from 'vue-router'
import { services as allServices, getServiceById } from '@/data/services'

const props = defineProps<{
  // This `id` is passed automatically by the router because of `props: true`
  id: string
}>()

/**
 * Finds the specific service to display based on the ID from the URL.
 * Its `detailsComponent` will be rendered in the main content area.
 */
const serviceComponent = computed(() => {
  const service = getServiceById(props.id)
  return service ? service.detailsComponent : null
})
</script>

<style scoped>
/* Add active styles for the sidebar navigation */
.router-link-exact-active {
  background-color: #f8f9fa; /* bg-light-gray */
  color: #00a896; /* text-primary */
}
</style>
