<script setup lang="ts">
import { ref, defineComponent } from 'vue'
import { RouterLink } from 'vue-router'
import AppLogo from './AppLogo.vue'

// Define icons locally to bypass import issues
const HamburgerIcon = defineComponent({
  template: `
    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
    </svg>
  `,
})

const CloseIcon = defineComponent({
  template: `
    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
    </svg>
  `,
})

const isMobileMenuOpen = ref(false)

const navLinks = [
  { text: 'Services', to: '/#services' },
  { text: 'About', to: '/#about' },
  { text: 'Articles', to: '/#articles' },
  { text: 'Contact', to: '/#contact' },
]

const actionLinks = [
  { text: 'Get In Touch', to: '/#contact', primary: true },
  { text: 'Client Login', to: '/login', primary: false },
]

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

      <!-- Desktop Menu -->
      <div class="hidden md:flex items-center space-x-8">
        <RouterLink
          v-for="link in navLinks"
          :key="link.to"
          :to="link.to"
          class="text-gray-600 hover:text-primary transition-colors"
          >{{ link.text }}</RouterLink
        >
      </div>
      <div class="hidden md:flex items-center space-x-4">
        <RouterLink
          v-for="link in actionLinks"
          :key="link.to"
          :to="link.to"
          :class="[
            'font-medium py-2 px-4 rounded-lg transition-all',
            link.primary
              ? 'bg-primary text-white hover:bg-opacity-90'
              : 'bg-light-gray text-primary hover:bg-gray-200',
          ]"
          >{{ link.text }}</RouterLink
        >
      </div>

      <!-- Mobile Menu Button -->
      <button @click="toggleMobileMenu" class="md:hidden">
        <CloseIcon v-if="isMobileMenuOpen" />
        <HamburgerIcon v-else />
      </button>
    </nav>

    <!-- Mobile Menu -->
    <div
      :class="{ block: isMobileMenuOpen, hidden: !isMobileMenuOpen }"
      class="md:hidden px-6 pb-4 space-y-2"
    >
      <RouterLink
        v-for="link in navLinks"
        :key="link.to"
        :to="link.to"
        @click="closeMobileMenu"
        class="block text-gray-600 hover:text-primary transition-colors py-2"
        >{{ link.text }}</RouterLink
      >
      <div class="border-t border-gray-200 pt-4 mt-4 space-y-2">
        <RouterLink
          v-for="link in actionLinks"
          :key="link.to"
          :to="link.to"
          @click="closeMobileMenu"
          :class="[
            'block text-center font-medium py-2 px-4 rounded-lg transition-all',
            link.primary
              ? 'bg-primary text-white hover:bg-opacity-90'
              : 'bg-light-gray text-primary hover:bg-gray-200',
          ]"
          >{{ link.text }}</RouterLink
        >
      </div>
    </div>
  </header>
</template>
