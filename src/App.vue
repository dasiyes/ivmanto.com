<script setup lang="ts">
import { RouterView } from 'vue-router'
import TheHeader from './components/layout/TheHeader.vue'
import TheFooter from './components/layout/TheFooter.vue'
</script>

<template>
  <div class="flex flex-col min-h-screen bg-gray-100">
    <TheHeader />
    <main class="flex-grow">
      <RouterView />
    </main>
    <TheFooter />
  </div>
</template>
