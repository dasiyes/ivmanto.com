<template>
  <header class="bg-white/80 backdrop-blur-lg sticky top-0 z-50 border-b border-gray-200">
    <nav class="container mx-auto px-6 py-3 flex justify-between items-center relative">
      <router-link to="/"
        ><img src="@/assets/mockup/logo.svg" alt="IVMANTO Logo" class="h-16 w-auto"
      /></router-link>
      <div class="hidden md:flex items-center space-x-8">
        <!-- Services Menu with Dropdown -->
        <div class="group py-4 -my-4">
          <router-link
            :to="{ name: 'services' }"
            class="text-gray-600 hover:text-primary transition-colors flex items-center gap-1 cursor-pointer"
          >
            <span>Services</span>
            <svg
              class="w-4 h-4 text-gray-500 group-hover:text-primary transition-transform duration-300 group-hover:rotate-180"
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
      <button class="md:hidden">
        <!-- Mobile menu icon -->
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16m-7 6h7"
          ></path>
        </svg>
      </button>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { services } from '@/data/services'
import { RouterLink } from 'vue-router'
</script>
