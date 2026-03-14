<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

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
    isOpen: false
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

const sectionRef = ref<HTMLElement | null>(null)
const isVisible = ref(false)
let observer: IntersectionObserver | null = null

onMounted(() => {
  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) isVisible.value = true
      })
    },
    { threshold: 0.15 }
  )
  if (sectionRef.value) observer.observe(sectionRef.value)
})

onUnmounted(() => {
  observer?.disconnect()
})
</script>

<template>
  <section ref="sectionRef" class="py-20 md:py-28 bg-gray-50 relative overflow-hidden">
    <!-- Subtle background accent -->
    <div class="absolute -top-20 -right-20 w-80 h-80 rounded-full bg-amber/5 blur-3xl pointer-events-none"></div>

    <div class="container mx-auto px-6 max-w-4xl relative z-10">
      <div
        class="text-center mb-12 transition-all duration-700"
        :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
      >
        <span class="text-amber font-semibold tracking-widest text-sm uppercase">FAQ</span>
        <h2 class="text-3xl md:text-5xl font-bold text-dark-slate mt-3">Common Questions</h2>
        <p class="text-gray-500 mt-4">Everything you need to know about partnering with me.</p>
      </div>

      <div class="space-y-4">
        <div 
          v-for="(faq, index) in faqs" 
          :key="index"
          class="glass-card-light transition-all duration-500 hover:shadow-lg"
          :class="[
            isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8',
            faq.isOpen ? 'ring-1 ring-primary/20' : ''
          ]"
          :style="{ transitionDelay: isVisible ? `${0.1 * index}s` : '0s' }"
        >
          <button 
            @click="toggleFaq(index)"
            class="w-full px-6 py-5 text-left flex justify-between items-center transition-colors"
          >
            <span class="font-semibold text-dark-slate pr-4">{{ faq.question }}</span>
            <div
              class="w-8 h-8 rounded-lg flex items-center justify-center flex-shrink-0 transition-all duration-300"
              :class="faq.isOpen ? 'bg-primary text-white rotate-180' : 'bg-gray-100 text-gray-400'"
            >
              <svg 
                class="w-4 h-4 transition-transform duration-300"
                fill="none" 
                stroke="currentColor" 
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </div>
          </button>
          <div 
            class="accordion-content px-6 text-gray-500 leading-relaxed"
            :class="{ 'is-open': faq.isOpen }"
          >
            {{ faq.answer }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
