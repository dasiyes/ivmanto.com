<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const sectionRef = ref<HTMLElement | null>(null)
const isVisible = ref(false)

let observer: IntersectionObserver | null = null

onMounted(() => {
  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          isVisible.value = true
        }
      })
    },
    { threshold: 0.15 }
  )
  if (sectionRef.value) {
    observer.observe(sectionRef.value)
  }
})

onUnmounted(() => {
  observer?.disconnect()
})
</script>

<template>
  <section ref="sectionRef" class="py-20 md:py-28 bg-white relative overflow-hidden">
    <!-- Subtle background decoration -->
    <div class="absolute top-10 right-0 w-72 h-72 rounded-full bg-primary/5 blur-3xl pointer-events-none"></div>
    <div class="absolute bottom-10 left-0 w-96 h-96 rounded-full bg-primary-dark/5 blur-3xl pointer-events-none"></div>

    <div class="container mx-auto px-6 relative z-10">
      <div
        class="text-center mb-16 transition-all duration-700"
        :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
      >
        <span class="text-primary font-semibold tracking-widest text-sm uppercase">My Approach</span>
        <h2 class="text-3xl md:text-5xl font-bold text-dark-slate mt-3">How I Work</h2>
        <p class="text-lg text-gray-500 mt-4 max-w-2xl mx-auto">
          A proven, 3-step approach to turning your data challenges into scalable solutions.
        </p>
      </div>

      <div class="grid md:grid-cols-3 gap-8 relative">
        <!-- Connecting line (desktop only) -->
        <div class="hidden md:block absolute top-[60px] left-[16.66%] right-[16.66%] h-0.5 bg-gradient-to-r from-primary/20 via-primary/40 to-primary/20"></div>

        <!-- Step 1 -->
        <div
          class="glass-card-light p-8 text-center group transition-all duration-500 hover:-translate-y-2"
          :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-12'"
          :style="{ transitionDelay: isVisible ? '0.2s' : '0s' }"
        >
          <div class="w-[72px] h-[72px] rounded-2xl flex items-center justify-center mx-auto mb-6 transition-all duration-300 relative z-10"
            style="background: var(--gradient-primary); box-shadow: 0 4px 20px rgba(0,168,150,.25);">
            <span class="text-2xl font-bold text-white">01</span>
          </div>
          <h3 class="text-xl font-bold text-dark-slate mb-3">Discovery & Audit</h3>
          <p class="text-gray-500 leading-relaxed">
            We start by understanding your business goals and auditing your current data landscape to identify gaps and opportunities.
          </p>
        </div>

        <!-- Step 2 -->
        <div
          class="glass-card-light p-8 text-center group transition-all duration-500 hover:-translate-y-2"
          :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-12'"
          :style="{ transitionDelay: isVisible ? '0.4s' : '0s' }"
        >
          <div class="w-[72px] h-[72px] rounded-2xl flex items-center justify-center mx-auto mb-6 transition-all duration-300 relative z-10"
            style="background: var(--gradient-primary); box-shadow: 0 4px 20px rgba(0,168,150,.25);">
            <span class="text-2xl font-bold text-white">02</span>
          </div>
          <h3 class="text-xl font-bold text-dark-slate mb-3">Strategy & Architecture</h3>
          <p class="text-gray-500 leading-relaxed">
            I design a custom, organic blueprint tailored to your needs — prioritizing security, scalability, and cost-efficiency on GCP.
          </p>
        </div>

        <!-- Step 3 -->
        <div
          class="glass-card-light p-8 text-center group transition-all duration-500 hover:-translate-y-2"
          :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-12'"
          :style="{ transitionDelay: isVisible ? '0.6s' : '0s' }"
        >
          <div class="w-[72px] h-[72px] rounded-2xl flex items-center justify-center mx-auto mb-6 transition-all duration-300 relative z-10"
            style="background: var(--gradient-primary); box-shadow: 0 4px 20px rgba(0,168,150,.25);">
            <span class="text-2xl font-bold text-white">03</span>
          </div>
          <h3 class="text-xl font-bold text-dark-slate mb-3">Implementation & Scale</h3>
          <p class="text-gray-500 leading-relaxed">
            From coding pipelines to deploying AI models, I lead the technical execution and ensure your team is empowered to take the lead.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
