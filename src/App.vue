<script setup lang="ts">
import { useHead } from '@vueuse/head'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

// Global Canonical Tag
// Forces all URLs to be treated as https://ivmanto.com/path
// Strips www, query params, and trailing slashes if needed (though Google usually handles trailing slashes fine)
useHead({
  link: [
    {
      rel: 'canonical',
      href: computed(() => {
        // Base URL is strictly non-www HTTPS
        const baseUrl = 'https://ivmanto.com'
        // Get the current path from the router (e.g., /services/principles)
        const path = route.path === '/' ? '' : route.path
        // Construct the full canonical URL
        return `${baseUrl}${path}`
      }),
    },
  ],
})
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <router-view />
  </div>
</template>

<style>
/* Global Styles */
@import './style.css';
</style>
