<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { ref, onMounted, computed, watch } from 'vue'

type TimeSlot = {
  startTime: string
  endTime: string
}

const selectedDate = ref(new Date())
const bookingWindowDays = 30 // The booking window should ideally come from a shared config
const availableSlots = ref<TimeSlot[]>([])
const selectedSlot = ref<TimeSlot | null>(null)
const isLoading = ref(false)
const error = ref<string | null>(null)

const bookingDetails = ref({
  name: '',
  email: '',
  notes: '',
})

const isBookingConfirmed = ref(false)

const formattedDate = computed(() => {
  return selectedDate.value.toLocaleDateString('en-US', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
})

const isPreviousDayDisabled = computed(() => {
  const today = new Date()
  today.setHours(0, 0, 0, 0) // Normalize to the start of the day
  return selectedDate.value <= today
})

const isNextDayDisabled = computed(() => {
  const today = new Date()
  const limitDate = new Date(today.setDate(today.getDate() + bookingWindowDays))
  limitDate.setHours(0, 0, 0, 0) // Normalize
  return selectedDate.value >= limitDate
})

const timezone = computed(() => {
  return Intl.DateTimeFormat().resolvedOptions().timeZone.replace(/_/g, ' ')
})

function toYYYYMMDD(date: Date) {
  return date.toISOString().split('T')[0]
}

async function fetchAvailability(date: Date) {
  isLoading.value = true
  error.value = null
  selectedSlot.value = null
  availableSlots.value = [] // Clear previous slots immediately
  try {
    const dateStr = toYYYYMMDD(date)
    const response = await fetch(`/api/booking/availability?date=${dateStr}`)
    if (!response.ok) {
      // Provide a more specific error for easier debugging
      throw new Error(`Server responded with status ${response.status}`)
    }
    availableSlots.value = await response.json()
  } catch (e: any) {
    console.error('Failed to fetch availability:', e)
    error.value = 'Could not load available time slots. Please try again later.'
    availableSlots.value = []
  } finally {
    isLoading.value = false
  }
}

function changeDay(amount: number) {
  const newDate = new Date(selectedDate.value)
  newDate.setDate(newDate.getDate() + amount)
  newDate.setHours(12, 0, 0, 0) // Avoid timezone-related date shifts
  selectedDate.value = newDate
  // A watcher will automatically fetch availability
}

function selectSlot(slot: TimeSlot) {
  selectedSlot.value = slot
}

async function handleBookingSubmit() {
  if (!selectedSlot.value) return

  isLoading.value = true
  error.value = null
  try {
    const response = await fetch('/api/booking/book', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        startTime: selectedSlot.value.startTime,
        ...bookingDetails.value,
      }),
    })

    if (!response.ok) {
      const errorData = await response.text()
      // The server provides a user-friendly error message (e.g., "time slot is already booked").
      throw new Error(errorData || 'Booking failed. Please try again.')
    }

    isBookingConfirmed.value = true
  } catch (e: any) {
    console.error('Booking submission failed:', e)
    // Show the error message from the server or our fallback.
    error.value = e.message || 'An unexpected error occurred during booking.'
  } finally {
    isLoading.value = false
  }
}

function resetBookingProcess() {
  isBookingConfirmed.value = false
  selectedSlot.value = null
  bookingDetails.value = { name: '', email: '', notes: '' }
  // Re-fetch availability for the current date
  fetchAvailability(selectedDate.value)
}

function formatTime(dateString: string) {
  return new Date(dateString).toLocaleTimeString('en-US', {
    hour: 'numeric',
    minute: '2-digit',
    hour12: true,
  })
}

watch(selectedDate, (newDate) => {
  fetchAvailability(newDate)
})

onMounted(() => {
  fetchAvailability(selectedDate.value)
})
</script>

<template>
  <div class="max-w-2xl mx-auto bg-white p-8 rounded-lg shadow-lg">
    <div v-if="!isBookingConfirmed">
      <h2 class="text-2xl font-bold text-dark-slate mb-6 text-center">Book a Consultation</h2>

      <!-- Date Selector -->
      <div class="flex items-center justify-between mb-6">
        <button
          @click="changeDay(-1)"
          class="p-2 rounded-full hover:bg-gray-200 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="isPreviousDayDisabled || isLoading"
        >
          &lt;
        </button>
        <div class="text-center">
          <h3 class="text-lg font-semibold text-gray-700">{{ formattedDate }}</h3>
          <p class="text-xs text-gray-500">Timezone: {{ timezone }}</p>
        </div>
        <button
          @click="changeDay(1)"
          class="p-2 rounded-full hover:bg-gray-200 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="isNextDayDisabled || isLoading"
        >
          &gt;
        </button>
      </div>

      <!-- Error Message -->
      <div
        v-if="error"
        class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4"
        role="alert"
      >
        <span class="block sm:inline">{{ error }}</span>
      </div>

      <!-- Loading Spinner -->
      <div v-if="isLoading" class="grid grid-cols-2 sm:grid-cols-3 gap-4">
        <!-- Skeleton Loaders -->
        <div v-for="n in 6" :key="n" class="p-3 h-12 bg-gray-200 rounded-lg animate-pulse"></div>
      </div>

      <!-- Time Slots -->
      <div v-else-if="!selectedSlot" class="min-h-[100px]">
        <!--
          The min-h-[100px] prevents the layout from jumping when slots
          are loaded or when the "no slots" message appears.
        -->
        <div v-if="availableSlots.length > 0" class="grid grid-cols-2 sm:grid-cols-3 gap-4">
          <button
            v-for="slot in availableSlots"
            :key="slot.startTime"
            @click="selectSlot(slot)"
            class="p-3 border rounded-lg text-center text-primary hover:bg-primary hover:text-white transition-colors"
          >
            {{ formatTime(slot.startTime) }}
          </button>
        </div>
        <div v-else-if="!error" class="text-center text-gray-500 py-8">
          <p>No available slots for this day. Please try another date.</p>
        </div>
      </div>

      <!-- Booking Form -->
      <form v-else @submit.prevent="handleBookingSubmit">
        <h4 class="text-xl font-semibold mb-4">
          Confirming for: {{ formatTime(selectedSlot.startTime) }}
        </h4>
        <div class="space-y-4">
          <div>
            <label for="name" class="block text-gray-700 font-medium mb-1">Full Name</label>
            <input
              type="text"
              id="name"
              v-model="bookingDetails.name"
              required
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary"
            />
          </div>
          <div>
            <label for="email" class="block text-gray-700 font-medium mb-1">Email</label>
            <input
              type="email"
              id="email"
              v-model="bookingDetails.email"
              required
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary"
            />
          </div>
          <div>
            <label for="notes" class="block text-gray-700 font-medium mb-1"
              >Project Notes (Optional)</label
            >
            <textarea
              id="notes"
              v-model="bookingDetails.notes"
              rows="3"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary"
            ></textarea>
          </div>
        </div>
        <div class="flex justify-end gap-4 mt-6">
          <button
            type="button"
            @click="selectedSlot = null"
            class="px-6 py-2 bg-gray-200 text-gray-800 rounded-lg hover:bg-gray-300"
          >
            Back
          </button>
          <button
            type="submit"
            :disabled="isLoading"
            class="px-6 py-2 bg-primary text-white rounded-lg hover:bg-primary-dark disabled:bg-gray-400 disabled:cursor-not-allowed"
          >
            <span v-if="isLoading">Booking...</span>
            <span v-else>Confirm Booking</span>
          </button>
        </div>
      </form>
    </div>

    <!-- Confirmation Message -->
    <div v-else class="text-center py-12">
      <h2 class="text-2xl font-bold text-green-600 mb-4">Booking Confirmed!</h2>
      <p class="text-gray-700">
        Thank you, {{ bookingDetails.name }}. A confirmation email has been sent to
        {{ bookingDetails.email }}.
      </p>
      <div class="flex justify-center gap-4 mt-8">
        <RouterLink to="/" class="px-6 py-2 bg-gray-200 text-gray-800 rounded-lg hover:bg-gray-300">
          Home
        </RouterLink>
        <button
          @click="resetBookingProcess"
          class="px-6 py-2 bg-primary text-white rounded-lg hover:bg-primary-dark"
        >
          Book Another
        </button>
      </div>
    </div>
  </div>
</template>
