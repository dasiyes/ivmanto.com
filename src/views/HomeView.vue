<script setup lang="ts">
// This view component will act as a container for your home page sections.
// We can add other sections like Hero, About, etc., here later.
import ArticlesSection from '../components/sections/ArticlesSection.vue'
import ContactSection from '../components/sections/ContactSection.vue'
</script>

<template>
  <div>
    <ArticlesSection />
    <ContactSection />
  </div>
</template>
