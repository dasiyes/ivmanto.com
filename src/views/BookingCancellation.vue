<template>
  <div class="cancellation-container">
    <div v-if="isLoading" class="status-box">
      <h2>Processing Cancellation...</h2>
      <p>Please wait while we process your request.</p>
    </div>
    <div v-else-if="error" class="status-box error">
      <h2>Cancellation Failed</h2>
      <p>{{ error }}</p>
      <router-link to="/">Go to Homepage</router-link>
    </div>
    <div v-else class="status-box success">
      <h2>Booking Cancelled</h2>
      <p>{{ message }}</p>
      <router-link to="/booking">Book a new consultation</router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const isLoading = ref(true)
const error = ref(null)
const message = ref('')

onMounted(async () => {
  const token = route.query.token

  if (!token) {
    error.value = 'No cancellation token found. The link may be invalid.'
    isLoading.value = false
    return
  }

  try {
    const response = await fetch(`/api/booking/cancel?token=${token}`)

    if (!response.ok) {
      // Try to get a meaningful error message from the backend response body
      const errorText = await response.text()
      throw new Error(errorText || `Server responded with status: ${response.status}`)
    }

    const data = await response.json()
    message.value = data.message || 'Your booking has been successfully cancelled.'
  } catch (err) {
    console.error('Cancellation error:', err)
    error.value = err.message || 'An unexpected error occurred. Please contact support.'
  } finally {
    isLoading.value = false
  }
})
</script>

<style scoped>
.cancellation-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
  padding: 2rem;
}

.status-box {
  text-align: center;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  max-width: 500px;
  width: 100%;
}

.status-box.success {
  background-color: #e8f5e9; /* Light green */
  border: 1px solid #4caf50; /* Green */
}

.status-box.error {
  background-color: #ffebee; /* Light red */
  border: 1px solid #f44336; /* Red */
}

h2 {
  margin-top: 0;
  color: #333;
}

p {
  color: #555;
}
</style>
