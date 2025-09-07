<template>
  <header class="bg-white/80 backdrop-blur-lg sticky top-0 z-50 border-b border-gray-200 shadow-sm">
    <nav class="container mx-auto px-6 py-3 flex justify-between items-center relative">
      <router-link to="/"
        ><img
          src="@/assets/mockup/logo.svg"
          alt="IVMANTO Logo"
          class="h-16 w-auto"
          width="64"
          height="64"
      /></router-link>
      <div class="hidden md:flex items-center space-x-8">
        <!-- Services Menu with Dropdown -->
        <!-- The py-4 -my-4 trick creates an invisible vertical padding area to bridge the gap between the link and the submenu, preventing it from closing prematurely. -->
        <div :class="{ group: !isServicesPage }" class="py-4 -my-4">
          <router-link
            :to="{ name: 'services' }"
            class="text-gray-600 transition-colors flex items-center gap-1 cursor-pointer"
            :class="{
              'hover:text-primary': !isServicesPage,
              'text-primary font-semibold': isServicesPage,
            }"
          >
            <span>Services</span>
            <svg
              class="w-4 h-4 text-gray-500 transition-transform duration-300"
              :class="{ 'group-hover:text-primary group-hover:rotate-180': !isServicesPage }"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M19 9l-7 7-7-7"
              ></path>
            </svg>
          </router-link>

          <!-- Mega Menu -->
          <div
            v-if="!isServicesPage"
            class="absolute top-full left-1/2 -translate-x-1/2 w-[80%] max-w-5xl opacity-0 invisible group-hover:opacity-100 group-hover:visible pointer-events-none group-hover:pointer-events-auto transition-opacity duration-300"
          >
            <div class="bg-white shadow-lg rounded-lg p-8 border border-gray-100 mt-2">
              <div class="grid grid-cols-2 gap-x-10 gap-y-6">
                <!-- Dynamic Service Items from services.ts -->
                <router-link
                  v-for="service in services"
                  :key="service.id"
                  :to="`/services/${service.id}`"
                  class="block p-3 -m-3 rounded-lg hover:bg-gray-50 transition-colors"
                >
                  <p class="font-bold text-dark-slate">{{ service.menuTitle }}</p>
                  <p class="text-sm text-gray-600 mt-1">{{ service.summary }}</p>
                </router-link>
              </div>
            </div>
          </div>
        </div>
        <router-link to="/about" class="text-gray-600 hover:text-primary transition-colors"
          >About</router-link
        >
        <router-link to="/blog" class="text-gray-600 hover:text-primary transition-colors"
          >Articles</router-link
        >
        <router-link to="/#contact" class="text-gray-600 hover:text-primary transition-colors"
          >Contact</router-link
        >
      </div>
      <div class="hidden md:flex items-center space-x-4">
        <router-link
          to="/#contact"
          class="bg-primary text-white font-medium py-2 px-4 rounded-lg hover:bg-opacity-90 transition-all"
          >Get In Touch</router-link
        >
        <router-link
          :to="{ name: 'login' }"
          class="bg-light-gray text-primary font-medium py-2 px-4 rounded-lg hover:bg-gray-200 transition-all"
          >Client Login</router-link
        >
      </div>
      <!-- Mobile Menu Button -->
      <button
        @click="isMobileMenuOpen = !isMobileMenuOpen"
        class="md:hidden"
        aria-label="Toggle menu"
      >
        <!-- Hamburger Icon -->
        <svg
          v-if="!isMobileMenuOpen"
          class="w-6 h-6"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16m-7 6h7"
          ></path>
        </svg>
        <!-- Close Icon -->
        <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          ></path>
        </svg>
      </button>
    </nav>

    <!-- Mobile Menu Panel -->
    <div v-if="isMobileMenuOpen" class="md:hidden bg-white border-t border-gray-200">
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
        <router-link
          to="/services"
          @click="isMobileMenuOpen = false"
          class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-primary hover:bg-gray-50"
          >Services</router-link
        >
        <router-link
          to="/about"
          @click="isMobileMenuOpen = false"
          class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-primary hover:bg-gray-50"
          >About</router-link
        >
        <router-link
          to="/blog"
          @click="isMobileMenuOpen = false"
          class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-primary hover:bg-gray-50"
          >Articles</router-link
        >
        <router-link
          to="/#contact"
          @click="isMobileMenuOpen = false"
          class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-primary hover:bg-gray-50"
          >Contact</router-link
        >
      </div>
      <div class="pt-4 pb-3 border-t border-gray-200">
        <div class="px-5 flex items-center gap-4">
          <router-link
            to="/#contact"
            @click="isMobileMenuOpen = false"
            class="flex-1 bg-primary text-white font-medium py-2 px-4 rounded-lg hover:bg-opacity-90 transition-all text-center"
            >Get In Touch</router-link
          >
          <router-link
            :to="{ name: 'login' }"
            @click="isMobileMenuOpen = false"
            class="flex-1 bg-light-gray text-primary font-medium py-2 px-4 rounded-lg hover:bg-gray-200 transition-all text-center"
            >Client Login</router-link
          >
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { services } from '@/data/services'
import { RouterLink, useRoute } from 'vue-router'
import { computed, ref } from 'vue'

const isMobileMenuOpen = ref(false)
const route = useRoute()
const isServicesPage = computed(() => route.path.startsWith('/services'))
</script>
