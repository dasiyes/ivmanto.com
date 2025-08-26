<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import AppLogo from './AppLogo.vue'

const isMobileMenuOpen = ref(false)

function toggleMobileMenu() {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

function closeMobileMenu() {
  isMobileMenuOpen.value = false
}
</script>

<template>
  <header class="bg-white/80 backdrop-blur-lg sticky top-0 z-50 border-b border-gray-200">
    <nav class="container mx-auto px-6 py-3 flex justify-between items-center">
      <RouterLink to="/" @click="closeMobileMenu">
        <AppLogo />
      </RouterLink>

      <div class="hidden md:flex items-center space-x-8">
        <RouterLink to="/#services" class="text-gray-600 hover:text-primary transition-colors"
          >Services</RouterLink
        >
        <RouterLink to="/#about" class="text-gray-600 hover:text-primary transition-colors"
          >About</RouterLink
        >
        <RouterLink to="/#articles" class="text-gray-600 hover:text-primary transition-colors"
          >Articles</RouterLink
        >
        <RouterLink to="/#contact" class="text-gray-600 hover:text-primary transition-colors"
          >Contact</RouterLink
        >
      </div>
      <div class="hidden md:flex items-center space-x-4">
        <RouterLink
          to="/#contact"
          class="bg-primary text-white font-medium py-2 px-4 rounded-lg hover:bg-opacity-90 transition-all"
          >Get In Touch</RouterLink
        >
        <RouterLink
          to="/login"
          class="bg-light-gray text-primary font-medium py-2 px-4 rounded-lg hover:bg-gray-200 transition-all"
          >Client Login</RouterLink
        >
      </div>

      <button @click="toggleMobileMenu" class="md:hidden">
        <svg
          class="w-6 h-6"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16m-7 6h7"
          ></path>
        </svg>
      </button>
    </nav>

    <div :class="{ hidden: !isMobileMenuOpen }" class="md:hidden px-6 pb-4 space-y-2">
      <RouterLink
        to="/#services"
        @click="closeMobileMenu"
        class="block text-gray-600 hover:text-primary transition-colors"
        >Services</RouterLink
      >
      <RouterLink
        to="/#about"
        @click="closeMobileMenu"
        class="block text-gray-600 hover:text-primary transition-colors"
        >About</RouterLink
      >
      <RouterLink
        to="/#articles"
        @click="closeMobileMenu"
        class="block text-gray-600 hover:text-primary transition-colors"
        >Articles</RouterLink
      >
      <RouterLink
        to="/#contact"
        @click="closeMobileMenu"
        class="block text-gray-600 hover:text-primary transition-colors"
        >Contact</RouterLink
      >
      <RouterLink
        to="/#contact"
        @click="closeMobileMenu"
        class="block bg-primary text-white text-center font-medium mt-2 py-2 px-4 rounded-lg hover:bg-opacity-90 transition-all"
        >Get In Touch</RouterLink
      >
      <RouterLink
        to="/login"
        @click="closeMobileMenu"
        class="block bg-light-gray text-primary text-center font-medium mt-2 py-2 px-4 rounded-lg hover:bg-gray-200 transition-all"
        >Client Login</RouterLink
      >
    </div>
  </header>
</template>
