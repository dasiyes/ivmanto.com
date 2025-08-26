<script setup>
import TheHeader from './components/layout/TheHeader.vue'
import TheFooter from './components/layout/TheFooter.vue'
import HeroSection from './components/sections/HeroSection.vue'
import ServicesSection from './components/sections/ServicesSection.vue'
import AboutSection from './components/sections/AboutSection.vue'
import ArticlesSection from './components/sections/ArticlesSection.vue'
import ContactSection from './components/sections/ContactSection.vue'
</script>

<template>
  <TheHeader />
  <main>
    <HeroSection />
    <ServicesSection />
    <AboutSection />
    <ArticlesSection />
    <ContactSection />
  </main>
  <TheFooter />
</template>
