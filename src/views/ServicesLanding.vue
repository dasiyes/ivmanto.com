<template>
  <div class="container mx-auto px-6 py-12">
    <div class="text-center max-w-3xl mx-auto mb-12">
      <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
        Our Data & AI Services
      </h1>
      <p class="mt-6 text-lg leading-8 text-gray-600">
        We help businesses with data architecture, governance, and AI-driven solutions to turn data
        into a strategic asset. Explore our offerings to find the right solution for you.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      <div
        v-for="service in allServices"
        :key="service.id"
        class="flex flex-col rounded-lg border border-gray-200 shadow-sm p-6"
      >
        <h3 class="text-lg font-semibold text-dark-slate">{{ service.menuTitle }}</h3>
        <p class="mt-2 text-sm text-gray-600 flex-grow">{{ service.summary }}</p>
        <RouterLink
          :to="`/services/${service.id}`"
          class="mt-4 inline-block text-primary font-semibold hover:underline"
        >
          Learn More &rarr;
        </RouterLink>
      </div>
    </div>

    <div class="mt-20 text-center">
      <h2 class="text-3xl font-bold">Ready to build your data foundation?</h2>
      <p class="mt-4 text-lg text-gray-600">
        Let's discuss how these services can be tailored to your business.
      </p>
      <RouterLink
        :to="{ name: 'booking' }"
        class="mt-6 inline-block bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-colors"
      >
        Book a Consultation
      </RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { services as allServices } from '@/data/services'
</script>
