<template>
  <div class="p-8 h-full flex flex-col">
    <div v-if="service" class="bg-white/80 backdrop-blur-sm p-8 rounded-lg shadow-md">
      <h3 class="text-2xl font-bold text-dark-slate">{{ service.title }}</h3>
      <p class="mt-4 text-gray-700 leading-relaxed">{{ service.details }}</p>
    </div>
    <div v-else class="flex-grow flex items-center justify-center text-center text-gray-500">
      <p class="text-lg">Select a service from the left to see the details.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Service } from '@/data/services'

defineProps<{
  service: Service | undefined
}>()
</script>
