<script setup lang="ts">
import { RouterView } from 'vue-router'
import Footer from './components/layout/TheFooter.vue'
import Header from './components/layout/TheHeader.vue'
import CookieBanner from './components/CookieBanner.vue'
</script>

<template>
  <div class="flex flex-col min-h-screen bg-gray-100">
    <Header />
    <!-- Main Content -->
    <main class="flex-grow flex flex-col">
      <RouterView />
    </main>
    <Footer />
    <CookieBanner />
  </div>
</template>
