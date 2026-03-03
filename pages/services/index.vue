<script setup lang="ts">
import { services as allServices } from '~/data/services'

useSeoMeta({
  title: 'Services | ivmanto.com',
  description:
    'Data & AI consulting services on GCP: data architecture, governance, AI/ML solutions, and sovereign cloud. Turn your data into a strategic asset.',
  ogTitle: 'Services | ivmanto.com',
  ogDescription:
    'Data strategy, GCP architecture, AI/ML solutions, and Go backend development — explore our full service offering.',
  twitterTitle: 'Services | ivmanto.com',
  twitterDescription:
    'Data strategy, GCP architecture, AI/ML solutions, and Go backend development — explore our full service offering.',
})

// ItemList schema for structured service listing
useHead({
  script: [
    {
      type: 'application/ld+json',
      children: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'ItemList',
        name: 'Data & AI Consulting Services',
        itemListElement: allServices.map((s, i) => ({
          '@type': 'ListItem',
          position: i + 1,
          name: s.menuTitle,
          description: s.summary,
          url: `https://ivmanto.com/services/${s.id}`,
        })),
      }),
    },
  ],
})
</script>

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
        <NuxtLink
          :to="`/services/${service.id}`"
          class="mt-4 inline-block text-primary font-semibold hover:underline"
        >
          Learn More &rarr;
        </NuxtLink>
      </div>
    </div>

    <div class="mt-20 text-center">
      <h2 class="text-3xl font-bold">Ready to build your data foundation?</h2>
      <p class="mt-4 text-lg text-gray-600">
        Let's discuss how these services can be tailored to your business.
      </p>
      <NuxtLink
        :to="{ name: 'booking' }"
        class="mt-6 inline-block bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-colors"
      >
        Book a Consultation
      </NuxtLink>
    </div>
  </div>
</template>
