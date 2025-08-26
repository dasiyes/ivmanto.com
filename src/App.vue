<script setup lang="ts">
import AppHeader from './components/AppHeader.vue'
import HeroSection from './components/HeroSection.vue'
import ServicesSection from './components/ServicesSection.vue'
import AboutSection from './components/AboutSection.vue'
import ArticlesSection from './components/ArticlesSection.vue'
import ContactSection from './components/ContactSection.vue'
import AppFooter from './components/AppFooter.vue'
</script>

<template>
  <AppHeader />
  <main>
    <HeroSection />
    <ServicesSection />
    <AboutSection />
    <ArticlesSection />
    <ContactSection />
  </main>
  <AppFooter />
</template>

<style>
/* Global styles are now in src/assets/main.css */
</style>
