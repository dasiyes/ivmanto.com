<template>
  <div v-if="showBanner" class="cookie-banner">
    <p class="cookie-text">
      This website uses cookies to enhance your browsing experience and analyze site traffic. By
      clicking "Accept", you agree to our use of cookies.
      <router-link to="/privacy-policy">Full Privacy Policy</router-link>.
    </p>
    <div class="cookie-actions">
      <button @click="declineCookies" class="btn btn-secondary">Decline</button>
      <button @click="acceptCookies" class="btn btn-primary">Accept</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { ref, onMounted } from 'vue'

const showBanner = ref(false)

const COOKIE_CONSENT_KEY = 'cookie_consent'

onMounted(() => {
  // Check if consent has already been given
  if (!localStorage.getItem(COOKIE_CONSENT_KEY)) {
    showBanner.value = true
  }
})

const acceptCookies = () => {
  localStorage.setItem(COOKIE_CONSENT_KEY, 'accepted')
  showBanner.value = false
}

const declineCookies = () => {
  localStorage.setItem(COOKIE_CONSENT_KEY, 'declined')
  showBanner.value = false
}
</script>

<style scoped>
.cookie-banner {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: #2d3748; /* A dark background, adjust to your theme */
  color: #f7fafc; /* Light text color */
  padding: 1rem 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  z-index: 1000;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.2);
  flex-wrap: wrap;
  gap: 1rem;
}

.cookie-text {
  margin: 0;
  flex-grow: 1;
  font-size: 0.9rem;
}

.cookie-text a {
  color: #63b3ed; /* A light blue for links, adjust to your theme */
  text-decoration: underline;
}

.cookie-actions {
  display: flex;
  gap: 0.75rem;
}

.btn {
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.2s ease-in-out;
}

.btn-primary {
  background-color: #2b6cb0; /* A primary blue, adjust to your theme */
  color: white;
}

.btn-primary:hover {
  background-color: #3182ce;
}

.btn-secondary {
  background-color: #4a5568; /* A secondary gray, adjust to your theme */
  color: white;
}

.btn-secondary:hover {
  background-color: #2d3748;
}
</style>
