<script setup lang="ts">
import { ref } from 'vue'
import { useGemini } from '@/composables/useGemini'

const { isLoading, error, callGemini } = useGemini()
const topic = ref('')
const ideas = ref([])

async function generateIdeas() {
  if (!topic.value) {
    error.value = 'Please enter a topic.'
    ideas.value = []
    return
  }

  error.value = null
  ideas.value = []

  const prompt = `Generate 3 creative and professional article titles based on the topic: "${topic.value}". The articles should be relevant for a blog about data architecture, cloud computing, and AI. Format the response as a JSON object with a single key "ideas" which is an array of strings. Example: {"ideas": ["Title 1", "Title 2", "Title 3"]}`

  try {
    const responseText = await callGemini(prompt, true)
    const responseJson = JSON.parse(responseText)
    if (responseJson.ideas) {
      ideas.value = responseJson.ideas
    }
  } catch (e) {
    // Error is already set in the composable
    console.error('Failed to generate article ideas:', e)
  }
}
</script>

<template>
  <section id="articles" class="py-20 md:py-28">
    <div class="container mx-auto px-6">
      <div class="text-center mb-16">
        <h2 class="text-3xl md:text-4xl font-bold text-dark-slate">Insights & Articles</h2>
        <p class="text-lg text-gray-600 mt-4 max-w-2xl mx-auto">
          Sharing knowledge from real-world projects and my vision for the future of data.
        </p>
      </div>
      <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div class="border border-gray-200 rounded-xl overflow-hidden group">
          <div class="p-6">
            <span class="text-sm text-gray-500">Data Architecture • 10 min read</span>
            <h3
              class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors"
            >
              Common Architectural Patterns for Modern Data Platforms
            </h3>
            <p class="mt-3 text-gray-600">
              A deep dive into the pros and cons of Data Lakes, Warehouses, and the modern Lakehouse
              architecture on GCP.
            </p>
            <a href="#" class="text-primary font-semibold mt-4 inline-block">Read More &rarr;</a>
          </div>
        </div>
        <div class="border border-gray-200 rounded-xl overflow-hidden group">
          <div class="p-6">
            <span class="text-sm text-gray-500">AI/ML • 8 min read</span>
            <h3
              class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors"
            >
              From Notebook to Production with Vertex AI
            </h3>
            <p class="mt-3 text-gray-600">
              A practical guide to operationalizing your machine learning models using Google
              Cloud's unified AI platform.
            </p>
            <a href="#" class="text-primary font-semibold mt-4 inline-block">Read More &rarr;</a>
          </div>
        </div>
        <div class="border border-gray-200 rounded-xl overflow-hidden group">
          <div class="p-6">
            <span class="text-sm text-gray-500">Data Governance • 12 min read</span>
            <h3
              class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors"
            >
              Why DAMA Principles Matter for Your Business
            </h3>
            <p class="mt-3 text-gray-600">
              Exploring how standardized data management practices can reduce risk and increase the
              value of your data assets.
            </p>
            <a href="#" class="text-primary font-semibold mt-4 inline-block">Read More &rarr;</a>
          </div>
        </div>
      </div>

      <div class="mt-20 bg-light-gray p-8 md:p-12 rounded-xl text-center">
        <h3 class="text-2xl md:text-3xl font-bold text-dark-slate">Need Inspiration?</h3>
        <p class="text-lg text-gray-600 mt-3 max-w-xl mx-auto">
          Enter a topic below and our AI will generate some creative article ideas for you.
        </p>
        <div class="mt-6 max-w-lg mx-auto flex flex-col sm:flex-row gap-4">
          <input
            v-model="topic"
            @keyup.enter="generateIdeas"
            type="text"
            placeholder="e.g., 'AI in retail'"
            class="w-full bg-white border-gray-300 rounded-md py-3 px-4 focus:ring-accent focus:border-accent text-lg"
          />
          <button
            @click="generateIdeas"
            :disabled="isLoading"
            class="bg-accent text-white font-bold py-3 px-6 rounded-lg hover:bg-opacity-90 transition-all text-lg whitespace-nowrap flex items-center justify-center"
          >
            <span v-if="!isLoading">✨ Generate Ideas</span>
            <div v-else class="loader border-t-white w-6 h-6"></div>
          </button>
        </div>
        <div class="mt-8 text-left">
          <div v-if="isLoading" class="flex justify-center"></div>
          <div v-else-if="error" class="text-red-500 text-center">
            <p>{{ error }}</p>
          </div>
          <div v-else-if="ideas.length > 0">
            <ul class="space-y-3 list-disc list-inside text-gray-700">
              <li v-for="(idea, index) in ideas" :key="index">{{ idea }}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
