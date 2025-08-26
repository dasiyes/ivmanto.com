<template>
  <section class="h-full grid place-items-center bg-light-gray hero-bg-pattern">
    <div class="container mx-auto px-6 text-center relative z-10">
      <span class="text-primary font-semibold tracking-wider"
        >GOOGLE CLOUD PLATFORM SPECIALIST</span
      >
      <h1 class="text-4xl md:text-6xl font-bold mt-4 text-dark-slate leading-tight">
        Expert Cloud Data Architecture &amp; AI Solutions
      </h1>
      <p class="text-lg md:text-xl text-gray-600 mt-6 max-w-3xl mx-auto">
        Transforming your data into actionable insights with robust, scalable, and intelligent cloud
        platforms built on GCP.
      </p>
      <div class="mt-10 flex justify-center gap-4">
        <RouterLink
          to="/contact"
          class="bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-all text-lg"
          >Book a Consultation</RouterLink
        >
        <RouterLink
          to="/services"
          class="bg-white text-primary font-bold py-3 px-8 rounded-lg border border-gray-200 hover:bg-gray-50 transition-all text-lg"
          >Learn More</RouterLink
        >
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
