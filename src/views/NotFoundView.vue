<template>
  <div class="min-h-[60vh] flex flex-col items-center justify-center px-4 py-16 sm:px-6 lg:px-8 bg-gray-50">
    <div class="max-w-max mx-auto text-center">
      <p class="text-sm font-semibold text-primary uppercase tracking-wide">404 error</p>
      <h1 class="mt-2 text-4xl font-extrabold text-gray-900 tracking-tight sm:text-5xl">Page not found</h1>
      <p class="mt-4 text-lg text-gray-500">Sorry, we couldn’t find the page you’re looking for.</p>
      <div class="mt-6">
        <RouterLink to="/" class="text-base font-medium text-primary hover:text-primary-dark">
          Go back home<span aria-hidden="true"> &rarr;</span>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useHead } from '@vueuse/head'

useHead({
  title: 'Page Not Found - Ivmanto',
  meta: [
    {
      name: 'robots',
      content: 'noindex',
    },
  ],
})
</script>
