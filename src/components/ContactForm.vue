<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { trackEvent, getGaClientId } from '@/services/analytics'

const props = defineProps<{
  // The source prop helps us identify where the form was submitted from,
  // as per the analytics plan (e.g., 'home_page_form', 'contact_page_form').
  source: string
}>()

const route = useRoute()

const formData = ref({
  name: '',
  email: '',
  message: '',
  sendCopyToSelf: false,
})

// 3. Use the onMounted lifecycle hook to read the URL when the component loads.
onMounted(() => {
  const subject = route.query.subject

  if (subject && typeof subject === 'string') {
    // If a subject exists, populate the message field.
    // Vue Router automatically decodes the value for you.
    formData.value.message = subject

    // 4. (Recommended) Clean the URL to remove the query parameter.
    // This prevents the message from being re-filled if the user refreshes the page.
    // It correctly preserves the #contact anchor.
    const cleanUrl = window.location.pathname + window.location.hash
    window.history.replaceState({}, document.title, cleanUrl)
  }
})

const formState = ref<'idle' | 'submitting' | 'success' | 'error'>('idle')
const errorMessage = ref('')

async function handleSubmit() {
  formState.value = 'submitting'
  errorMessage.value = ''

  // As per the analytics plan, capture the GA Client ID for server-side tracking.
  const clientId = getGaClientId()

  try {
    const response = await fetch('/api/contact', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        ...formData.value,
        ga_client_id: clientId, // Send the client ID to the backend.
      }),
    })

    if (!response.ok) {
      const errorText = await response.text()
      throw new Error(errorText || 'An unknown error occurred.')
    }

    formState.value = 'success'
    // As per the analytics plan, track this secondary conversion goal.
    trackEvent('contact_form_submit', {
      // Use the source prop to provide context on where the submission happened.
      source: props.source,
    })

    // Optionally reset form after a delay
    setTimeout(() => {
      formData.value = { name: '', email: '', message: '', sendCopyToSelf: false }
      formState.value = 'idle'
    }, 3000)
  } catch (err: any) {
    formState.value = 'error'
    errorMessage.value = err.message
  }
}
</script>

<template>
  <div class="max-w-xl mx-auto bg-white/10 p-6 md:p-8 rounded-xl backdrop-blur-sm">
    <div v-if="formState === 'success'" class="text-center p-8">
      <h3 class="text-2xl font-bold text-green-400">Thank You!</h3>
      <p class="text-gray-300 mt-2">
        Your message has been sent successfully. I'll get back to you shortly.
      </p>
    </div>
    <form v-else @submit.prevent="handleSubmit">
      <div class="space-y-6">
        <div>
          <label for="name" class="block text-gray-300 font-medium mb-2">Full Name</label>
          <input
            type="text"
            id="name"
            v-model="formData.name"
            required
            class="w-full bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent"
          />
        </div>
        <div>
          <label for="email" class="block text-gray-300 font-medium mb-2">Email</label>
          <input
            type="email"
            id="email"
            v-model="formData.email"
            required
            class="w-full bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent"
          />
        </div>
        <div>
          <label for="message" class="block text-gray-300 font-medium mb-2"
            >How can I help you?</label
          >
          <textarea
            id="message"
            v-model="formData.message"
            required
            rows="5"
            class="w-full bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent"
          ></textarea>
        </div>
      </div>
      <div class="mt-6">
        <label class="flex items-center cursor-pointer">
          <input
            type="checkbox"
            v-model="formData.sendCopyToSelf"
            class="h-4 w-4 rounded border-gray-400 bg-white/20 text-accent focus:ring-accent"
          />
          <span class="ml-2 text-gray-300">Send a copy of this message to myself</span>
        </label>
      </div>
      <div class="mt-6">
        <button
          type="submit"
          :disabled="formState === 'submitting'"
          class="w-full bg-accent text-white font-bold py-3 px-6 rounded-lg hover:bg-opacity-90 transition-all text-lg disabled:bg-gray-500 disabled:cursor-not-allowed"
        >
          <span v-if="formState === 'submitting'">Sending...</span>
          <span v-else>Send Message</span>
        </button>
      </div>
      <div v-if="formState === 'error'" class="mt-4 text-center text-red-400">
        <p>Error: {{ errorMessage }}</p>
      </div>
    </form>
  </div>
</template>
