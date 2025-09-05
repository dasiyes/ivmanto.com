<template>
  <div v-if="service" class="h-full flex flex-col p-4 md:p-6">
    <!-- Main content card that sits on top of the patterned background -->
    <div
      class="flex-grow flex flex-col bg-white/90 backdrop-blur-sm rounded-xl shadow-lg ring-1 ring-black ring-opacity-5 overflow-hidden"
    >
      <!-- Header Section -->
      <div class="p-6 border-b border-gray-200 flex-shrink-0">
        <h3 class="text-xl font-bold text-dark-slate">{{ service.title }}</h3>
        <p class="mt-2 italic text text-gray-600">{{ service.summary }}</p>
      </div>

      <!-- Dynamic Content Section -->
      <div class="flex-grow overflow-y-auto">
        <component :is="service.detailsComponent" />
      </div>

      <!-- Footer/Tags Section -->
      <div class="p-4 border-t border-gray-200 bg-gray-50/80 flex-shrink-0 space-y-4">
        <h4 class="font-semibold text-sm text-gray-700 mb-2">Related Technologies & Concepts</h4>
        <div class="flex flex-wrap gap-2">
          <button
            v-for="(detail, tag) in service.tagDetails"
            :key="tag"
            @click="$emit('update-right-column', detail)"
            class="px-2 py-1 text-xs rounded-md bg-blue-100 text-blue-800 hover:bg-blue-200 transition-colors"
          >
            {{ tag }}
          </button>
        </div>
        <!-- Service-specific CTA -->
        <div class="pt-4 border-t border-gray-200/60">
          <a
            :href="consultationLink"
            class="w-full inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-dark transition-colors"
          >
            Book a Consultation for this Service
          </a>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="p-6 text-center text-gray-500">
    <p>Select a service to see the details.</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Service } from '@/data/services'

const props = defineProps<{
  service: Service | undefined
}>()

defineEmits<{
  (e: 'update-right-column', content: string | undefined): void
}>()

const consultationLink = computed(() => {
  if (!props.service) {
    return '/booking'
  }
  const subject = `Consultation about: ${props.service.title}`
  return `/contact?subject=${encodeURIComponent(subject)}`
})
</script>
