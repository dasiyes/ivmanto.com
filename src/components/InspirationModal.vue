<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { ref, computed } from 'vue'

type Idea = {
  title: string
  summary: string
}

defineProps<{
  isOpen: boolean
  isLoading: boolean
  ideas: Idea[]
  error: string | null
  topic: string
}>()

const emit = defineEmits(['close'])

const selectedIdea = ref<Idea | null>(null)

function closeModal() {
  emit('close')
  selectedIdea.value = null // Reset on close
}

function selectIdea(idea: Idea) {
  selectedIdea.value = idea
}

const bookingLink = computed(() => {
  if (selectedIdea.value) {
    return {
      name: 'booking',
      query: {
        topic: selectedIdea.value.title,
        summary: selectedIdea.value.summary,
      },
    }
  }
  // Default link if no idea is selected
  return { name: 'booking' }
})
</script>

<template>
  <teleport to="body">
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition ease-in duration-200"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="isOpen"
        class="fixed inset-0 z-50 bg-black bg-opacity-50 backdrop-blur-sm flex items-center justify-center p-4"
        @click.self="closeModal"
      >
        <transition
          enter-active-class="transition ease-out duration-300"
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
          leave-active-class="transition ease-in duration-200"
          leave-from-class="opacity-100 scale-100"
          leave-to-class="opacity-0 scale-95"
        >
          <div class="bg-white rounded-2xl shadow-2xl w-full max-w-2xl max-h-[90vh] flex flex-col">
            <!-- Header -->
            <div class="p-6 border-b border-gray-200">
              <h3 class="text-2xl font-bold text-dark-slate">
                Ideas for: <span class="text-primary">{{ topic }}</span>
              </h3>
            </div>

            <!-- Content -->
            <div class="p-6 overflow-y-auto flex-grow">
              <div v-if="isLoading" class="space-y-4">
                <div v-for="n in 3" :key="n" class="animate-pulse">
                  <div class="h-6 bg-gray-200 rounded w-3/4 mb-2"></div>
                  <div class="h-4 bg-gray-200 rounded w-full"></div>
                </div>
              </div>
              <div v-else-if="error" class="text-red-600 bg-red-50 p-4 rounded-lg">
                <p class="font-bold">An error occurred</p>
                <p>{{ error }}</p>
              </div>
              <div v-else-if="ideas.length > 0" class="space-y-4">
                <div
                  v-for="(idea, index) in ideas"
                  :key="index"
                  @click="selectIdea(idea)"
                  class="p-4 rounded-lg cursor-pointer transition-all duration-200"
                  :class="[
                    selectedIdea?.title === idea.title
                      ? 'bg-primary/10 border-primary border ring-2 ring-primary/20'
                      : 'hover:bg-gray-100 border border-transparent',
                  ]"
                >
                  <h4 class="text-xl font-semibold text-dark-slate">{{ idea.title }}</h4>
                  <p class="mt-1 text-gray-600">{{ idea.summary }}</p>
                </div>
              </div>
            </div>

            <!-- Toolbar -->
            <div
              class="p-4 bg-gray-50 border-t border-gray-200 flex justify-between items-center rounded-b-2xl"
            >
              <RouterLink
                :to="bookingLink"
                class="bg-accent text-white font-bold py-2 px-5 rounded-lg hover:bg-opacity-90 transition-all text-base"
                @click="closeModal"
                >Book a Consultation</RouterLink
              >
              <button
                @click="closeModal"
                class="bg-white text-primary font-bold py-2 px-5 rounded-lg border border-gray-200 hover:bg-gray-100 transition-all text-base"
              >
                Close
              </button>
            </div>
          </div>
        </transition>
      </div>
    </transition>
  </teleport>
</template>
