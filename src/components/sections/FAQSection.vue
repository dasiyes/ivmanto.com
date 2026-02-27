<script setup lang="ts">
import { ref } from 'vue'

const faqs = ref([
  {
    question: "Do you work with existing on-premise infrastructure?",
    answer: "Yes. I specialize in hybrid cloud architectures and can help you seamlessly bridge your on-premise systems with Google Cloud Platform, or lead a full migration strategy.",
    isOpen: false
  },
  {
    question: "How long does a typical architecture audit take?",
    answer: "A standard initial assessment usually takes 1 to 2 weeks, depending on the complexity of your current landscape and the availability of your technical team for discovery sessions.",
    isOpen: false
  },
  {
    question: "Do you provide hands-on engineering or just strategy?",
    answer: "I do both. I believe the best architects are those who still know how to build. I provide the high-level strategic blueprint and can lead the hands-on implementation of pipelines and infrastructure.",
  },
  {
    question: "How do you ensure data security and compliance?",
    answer: "Security is 'baked in' from day one. I implement 'Secure by Design' principles, following the latest CIS benchmarks for GCP and ensuring alignment with GDPR or industry-specific regulations.",
    isOpen: false
  }
])

function toggleFaq(index: number) {
  faqs.value[index].isOpen = !faqs.value[index].isOpen
}
</script>

<template>
  <section class="py-20 bg-light-gray">
    <div class="container mx-auto px-6 max-w-4xl">
      <div class="text-center mb-12">
        <h2 class="text-3xl font-bold text-dark-slate">Common Questions</h2>
        <p class="text-gray-600 mt-4">Everything you need to know about partnering with me.</p>
      </div>

      <div class="space-y-4">
        <div 
          v-for="(faq, index) in faqs" 
          :key="index"
          class="bg-white rounded-lg shadow-sm overflow-hidden border border-gray-100"
        >
          <button 
            @click="toggleFaq(index)"
            class="w-full px-6 py-5 text-left flex justify-between items-center hover:bg-gray-50 transition-colors"
          >
            <span class="font-bold text-dark-slate">{{ faq.question }}</span>
            <svg 
              class="w-5 h-5 text-primary transition-transform duration-300"
              :class="{ 'rotate-180': faq.isOpen }"
              fill="none" 
              stroke="currentColor" 
              viewBox="0 0 24 24"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          <div 
            v-if="faq.isOpen" 
            class="px-6 pb-5 text-gray-600 animate-fadeIn"
          >
            {{ faq.answer }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fadeIn {
  animation: fadeIn 0.3s ease-out;
}
</style>
