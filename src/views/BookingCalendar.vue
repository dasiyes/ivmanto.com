<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { ref, onMounted, computed, watch } from 'vue'

type TimeSlot = {
  id: string
  start: string
  end: string
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
  const startOfToday = new Date(today.getFullYear(), today.getMonth(), today.getDate())

  const selected = selectedDate.value
  const startOfSelected = new Date(selected.getFullYear(), selected.getMonth(), selected.getDate())

  return startOfSelected.getTime() <= startOfToday.getTime()
})

const isNextDayDisabled = computed(() => {
  const today = new Date()
  // new Date() constructor correctly handles month/year rollovers
  const limitDate = new Date(
    today.getFullYear(),
    today.getMonth(),
    today.getDate() + bookingWindowDays,
  )

  const selected = selectedDate.value
  const startOfSelected = new Date(selected.getFullYear(), selected.getMonth(), selected.getDate())

  return startOfSelected.getTime() >= limitDate.getTime()
})

const timezone = computed(() => {
  return Intl.DateTimeFormat().resolvedOptions().timeZone.replace(/_/g, ' ')
})

function toYYYYMMDD(date: Date) {
  const year = date.getFullYear()
  const month = (date.getMonth() + 1).toString().padStart(2, '0')
  const day = date.getDate().toString().padStart(2, '0')
  return `${year}-${month}-${day}`
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
      throw new Error(`Server responded with status ${response.status}`)
    }
    // response.json() can fail on an empty body, so we catch that case.
    const data = await response.json().catch(() => null)
    // Ensure availableSlots is always an array to prevent template errors on `.length`.
    let slots = data || []

    // Check if the fetched date is today.
    const now = new Date()
    const isFetchingForToday =
      date.getFullYear() === now.getFullYear() &&
      date.getMonth() === now.getMonth() &&
      date.getDate() === now.getDate()

    if (isFetchingForToday) {
      console.log('[BookingCalendar] Filtering slots for today. Current time:', now.toISOString())
      console.log('[BookingCalendar] Original slots received:', JSON.parse(JSON.stringify(slots)))
      // Filter out slots where the start time has already passed.
      slots = slots.filter((slot: TimeSlot) => {
        const slotTime = new Date(slot.start)
        const isFuture = slotTime.getTime() > now.getTime()
        console.log(`[BookingCalendar] Checking slot ${slot.start}: is in future? ${isFuture}`)
        return isFuture
      })
      console.log('[BookingCalendar] Filtered slots:', JSON.parse(JSON.stringify(slots)))
    }

    availableSlots.value = slots
  } catch (e: any) {
    console.error('Failed to fetch availability:', e)
    error.value = 'Could not load available time slots. Please try again later.'
    availableSlots.value = []
  } finally {
    isLoading.value = false
  }
}

function changeDay(amount: number) {
  const currentDate = selectedDate.value
  // This "functional" approach of creating a new Date from primitives is the
  // most robust way to avoid mutation side-effects with Vue's reactivity.
  const newDate = new Date(
    currentDate.getFullYear(),
    currentDate.getMonth(),
    currentDate.getDate() + amount,
  )
  selectedDate.value = newDate
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
        eventId: selectedSlot.value.id,
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
          class="h-10 w-10 flex items-center justify-center bg-gray-100 text-gray-600 rounded-md transition-all duration-200 hover:bg-gray-200 hover:text-primary hover:scale-110 disabled:cursor-not-allowed disabled:opacity-75 disabled:scale-100 disabled:text-gray-400"
          :disabled="isPreviousDayDisabled || isLoading"
          aria-label="Previous day"
        >
          <svg
            class="w-6 h-6 text-gray-600"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 19l-7-7 7-7"
            ></path>
          </svg>
        </button>
        <div class="text-center">
          <h3 class="text-lg font-semibold text-gray-700">{{ formattedDate }}</h3>
          <p class="text-xs text-gray-500">Timezone: {{ timezone }}</p>
        </div>
        <button
          @click="changeDay(1)"
          class="h-10 w-10 flex items-center justify-center bg-gray-100 text-gray-600 rounded-md transition-all duration-200 hover:bg-gray-200 hover:text-primary hover:scale-110 disabled:cursor-not-allowed disabled:opacity-75 disabled:scale-100 disabled:text-gray-400"
          :disabled="isNextDayDisabled || isLoading"
          aria-label="Next day"
        >
          <svg
            class="w-6 h-6 text-gray-600"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
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
            :key="slot.start"
            @click="selectSlot(slot)"
            class="p-3 border rounded-lg text-center text-primary hover:bg-primary hover:text-white transition-colors"
          >
            {{ formatTime(slot.start) }}
          </button>
        </div>
        <div
          v-else-if="!error"
          class="flex items-center justify-center bg-blue-50 border border-blue-200 text-blue-800 px-4 py-3 rounded-lg"
          role="status"
        >
          <svg
            class="w-5 h-5 mr-3 text-blue-600"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
              clip-rule="evenodd"
            ></path>
          </svg>
          <span class="block sm:inline"
            >There are no available slots for this day. Please try another date.</span
          >
        </div>
      </div>

      <!-- Booking Form -->
      <form v-else @submit.prevent="handleBookingSubmit">
        <h4 class="text-xl font-semibold mb-4">
          Confirming for: {{ formatTime(selectedSlot.start) }}
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
