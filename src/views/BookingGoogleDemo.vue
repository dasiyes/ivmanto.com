<script setup lang="ts">
import { ref } from 'vue'

// IMPORTANT: Replace this with your actual Google Calendar Booking Page URL
// To get this URL:
// 1. Open Google Calendar
// 2. Under "Booking pages", hover over your booking page
// 3. Click Options → Sharing options → Website embed
// 4. Choose "Inline booking page" and copy the URL from the iframe src
const bookingPageUrl = ref(
  'https://calendar.google.com/calendar/appointments/schedules/YOUR_SCHEDULE_ID?gv=true',
)

// Example of what a real URL looks like:
// https://calendar.google.com/calendar/appointments/schedules/AcZssZ1234abcd...?gv=true
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold text-dark-slate mb-2 text-center">Book a Consultation</h1>
    <p class="text-gray-600 text-center mb-8">
      Select a time that works for you using the calendar below.
    </p>

    <!-- Embedded Google Booking Page -->
    <div class="bg-white rounded-lg shadow-lg overflow-hidden">
      <iframe
        :src="bookingPageUrl"
        style="border: 0"
        width="100%"
        height="600"
        frameborder="0"
        title="Schedule an appointment"
      ></iframe>
    </div>

    <!-- Fallback / Instructions -->
    <div class="mt-6 p-4 bg-blue-50 border border-blue-200 rounded-lg">
      <p class="text-sm text-blue-800">
        <strong>Note:</strong> If the calendar doesn't load, you can
        <a
          :href="bookingPageUrl"
          target="_blank"
          rel="noopener noreferrer"
          class="underline hover:text-blue-600"
        >
          open the booking page directly </a
        >.
      </p>
    </div>

    <!-- Demo notice - remove this in production -->
    <div class="mt-4 p-4 bg-yellow-50 border border-yellow-300 rounded-lg">
      <p class="text-sm text-yellow-800">
        <strong>Demo Notice:</strong> This is a demonstration of the Google Calendar Booking Page
        embed. Replace the <code class="bg-yellow-100 px-1 rounded">bookingPageUrl</code> in the
        component with your actual booking page URL from Google Calendar.
      </p>
    </div>
  </div>
</template>
