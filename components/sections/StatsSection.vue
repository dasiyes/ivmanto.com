<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const stats = [
  { target: 10, suffix: '+', label: 'Years Experience' },
  { target: 50, suffix: '+', label: 'Projects Delivered' },
  { target: 100, suffix: '%', label: 'Google Cloud Platform' },
]

const currentValues = ref(stats.map(() => 0))
const isVisible = ref(false)
const sectionRef = ref<HTMLElement | null>(null)
const completed = ref(false)
let observer: IntersectionObserver | null = null

function animateCounter(index: number, target: number, duration: number) {
  const start = performance.now()
  const step = (now: number) => {
    const elapsed = now - start
    const progress = Math.min(elapsed / duration, 1)
    // Ease-out cubic
    const eased = 1 - Math.pow(1 - progress, 3)
    currentValues.value[index] = Math.round(eased * target)
    if (progress < 1) {
      requestAnimationFrame(step)
    } else {
      currentValues.value[index] = target
    }
  }
  requestAnimationFrame(step)
}

onMounted(() => {
  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting && !completed.value) {
          isVisible.value = true
          completed.value = true
          // Stagger counter starts
          stats.forEach((stat, i) => {
            setTimeout(() => animateCounter(i, stat.target, 1800), i * 200)
          })
        }
      })
    },
    { threshold: 0.3 }
  )
  if (sectionRef.value) observer.observe(sectionRef.value)
})

onUnmounted(() => {
  observer?.disconnect()
})
</script>

<template>
  <section
    ref="sectionRef"
    class="py-16 md:py-20 relative overflow-hidden"
    style="background: var(--gradient-hero);"
  >
    <!-- Subtle dot grid -->
    <div class="absolute inset-0 dot-grid opacity-20"></div>

    <div class="container mx-auto px-6 relative z-10">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 md:gap-12">
        <div
          v-for="(stat, index) in stats"
          :key="index"
          class="text-center transition-all duration-700"
          :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
          :style="{ transitionDelay: `${index * 0.15}s` }"
        >
          <div
            class="counter-value text-5xl md:text-6xl font-extrabold text-white mb-2"
            :class="{ 'counter-glow': currentValues[index] === stat.target }"
          >
            {{ currentValues[index] }}{{ stat.suffix }}
          </div>
          <div class="text-gray-400 text-lg font-medium">{{ stat.label }}</div>
        </div>
      </div>
    </div>
  </section>
</template>
