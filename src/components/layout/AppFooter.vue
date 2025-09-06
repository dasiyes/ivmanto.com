<template>
  <footer class="bg-light-gray">
    <div class="container mx-auto px-6 py-8 text-center text-gray-600">
      <p>© 2024 IVMANTO. All Rights Reserved.</p>
      <div class="mt-4 flex justify-center space-x-6">
        <a
          href="https://linkedin.com/in/nikolaytonev"
          target="_blank"
          rel="noopener noreferrer"
          class="hover:text-primary"
          >LinkedIn</a
        >
        <a
          href="https://github.com/dasiyes"
          target="_blank"
          rel="noopener noreferrer"
          class="hover:text-primary"
          >GitHub</a
        >
        <router-link to="/privacy-policy" class="hover:text-primary">Privacy Policy</router-link>
      </div>
    </div>
  </footer>
</template>
<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
