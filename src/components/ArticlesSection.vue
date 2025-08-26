//
<script setup lang="ts">
import { ref } from 'vue'
import { callGeminiAPI } from '../services/gemini'

const topicInput = ref('')
const isLoading = ref(false)
// ... import your new callGeminiAPI service

// Replace `ideasOutput` with reactive state for ideas and errors
const ideas = ref<string[]>([])
const error = ref<string | null>(null)

// ...

interface IdeasResponse {
  ideas: string[]
}

async function generateIdeas(): Promise<void> {
  const topic = topicInput.value.trim()
  if (!topic) {
    error.value = 'Please enter a topic.'
    return
  }

  isLoading.value = true
  ideas.value = []
  error.value = null

  const prompt = `Generate 3 creative and professional article titles based on the topic: "${topic}". The articles should be relevant for a blog about data architecture, cloud computing, and AI.`

  try {
    // Use your new service
    const responseJson = await callGeminiAPI<IdeasResponse>(prompt, {
      responseMimeType: 'application/json',
      responseSchema: {
        type: 'OBJECT',
        properties: { ideas: { type: 'ARRAY', items: { type: 'STRING' } } },
      },
    })
    if (responseJson?.ideas?.length > 0) {
      ideas.value = responseJson.ideas
    } else {
      throw new Error('No ideas found in AI response.')
    }
  } catch (e) {
    error.value =
      e instanceof Error
        ? e.message
        : "Sorry, I couldn't generate ideas at the moment. Please try again later."
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <section id="articles" class="py-20 md:py-28">
    <div class="container mx-auto px-6 text-center">
      <h2 class="text-3xl md:text-4xl font-bold text-dark-slate">Need Blog Ideas?</h2>
      <p class="text-lg text-gray-600 mt-4 max-w-2xl mx-auto">
        Enter a topic and let my AI assistant generate some creative article titles for you.
      </p>

      <div class="mt-10 max-w-xl mx-auto">
        <div class="flex flex-col sm:flex-row gap-4">
          <input
            v-model="topicInput"
            @keyup.enter="generateIdeas"
            :disabled="isLoading"
            type="text"
            placeholder="e.g., 'Data Mesh implementation'"
            class="flex-grow bg-white border border-gray-300 rounded-md py-3 px-4 text-gray-700 focus:ring-accent focus:border-accent disabled:opacity-50"
          />
          <button
            @click="generateIdeas"
            :disabled="isLoading"
            class="bg-primary text-white font-bold py-3 px-6 rounded-lg hover:bg-opacity-90 transition-all text-lg disabled:opacity-50"
          >
            <span v-if="isLoading">Generating...</span>
            <span v-else>Generate Ideas</span>
          </button>
        </div>
      </div>

      <div v-if="error" class="mt-8 text-left text-red-500">
        <p>{{ error }}</p>
      </div>
      <div v-else-if="ideas.length" class="mt-8 text-left">
        <ul class="space-y-3 list-disc list-inside text-gray-700">
          <li v-for="(idea, index) in ideas" :key="index">{{ idea }}</li>
        </ul>
      </div>
    </div>
  </section>
</template>
