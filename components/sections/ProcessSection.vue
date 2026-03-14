<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const steps = [
  {
    icon: 'search',
    title: 'Discovery & Audit',
    description: 'We start by understanding your business goals and auditing your current data landscape to identify gaps and opportunities.',
  },
  {
    icon: 'blueprint',
    title: 'Strategy & Architecture',
    description: 'I design a custom, organic blueprint tailored to your needs — prioritizing security, scalability, and cost-efficiency on GCP.',
  },
  {
    icon: 'rocket',
    title: 'Implementation & Scale',
    description: 'From coding pipelines to deploying AI models, I lead the technical execution and ensure your team is empowered to take the lead.',
  },
]

const sectionRef = ref<HTMLElement | null>(null)
const isVisible = ref(false)
let observer: IntersectionObserver | null = null

// Mouse tilt tracking per card
const tilts = ref(steps.map(() => ({ x: 0, y: 0 })))

function handleMouseMove(index: number, event: MouseEvent) {
  const card = (event.currentTarget as HTMLElement)
  const rect = card.getBoundingClientRect()
  const centerX = rect.left + rect.width / 2
  const centerY = rect.top + rect.height / 2
  const rotateY = ((event.clientX - centerX) / (rect.width / 2)) * 6
  const rotateX = -((event.clientY - centerY) / (rect.height / 2)) * 6
  tilts.value[index] = { x: rotateX, y: rotateY }
}

function handleMouseLeave(index: number) {
  tilts.value[index] = { x: 0, y: 0 }
}

// Cursor glow tracking
const glowPositions = ref(steps.map(() => ({ x: 0, y: 0 })))

function handleGlowMove(index: number, event: MouseEvent) {
  const card = (event.currentTarget as HTMLElement)
  const rect = card.getBoundingClientRect()
  glowPositions.value[index] = {
    x: event.clientX - rect.left,
    y: event.clientY - rect.top,
  }
}

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
  <section ref="sectionRef" class="py-24 md:py-32 bg-white relative overflow-hidden">
    <!-- Subtle background decoration -->
    <div class="absolute top-10 right-0 w-72 h-72 rounded-full bg-primary/5 blur-3xl pointer-events-none"></div>
    <div class="absolute bottom-10 left-0 w-96 h-96 rounded-full bg-primary-dark/5 blur-3xl pointer-events-none"></div>

    <div class="container mx-auto px-6 relative z-10">
      <div
        class="text-center mb-20 transition-all duration-700"
        :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
      >
        <span class="text-primary font-semibold tracking-widest text-sm uppercase">My Approach</span>
        <h2 class="text-3xl md:text-5xl font-bold text-dark-slate mt-3">How I Work</h2>
        <p class="text-lg text-gray-500 mt-4 max-w-2xl mx-auto">
          A proven, 3-step approach to turning your data challenges into scalable solutions.
        </p>
      </div>

      <div class="grid md:grid-cols-3 gap-10 relative">
        <!-- Animated SVG connecting line (desktop only) -->
        <svg class="hidden md:block absolute top-[56px] left-[16%] right-[16%] w-[68%] h-[3px] overflow-visible" preserveAspectRatio="none">
          <line
            x1="0" y1="1" x2="100%" y2="1"
            stroke="url(#lineGrad)"
            stroke-width="2"
            class="draw-line"
            :class="{ 'is-drawn': isVisible }"
          />
          <defs>
            <linearGradient id="lineGrad" x1="0" y1="0" x2="1" y2="0">
              <stop offset="0%" stop-color="rgba(0,168,150,0.1)" />
              <stop offset="50%" stop-color="rgba(0,168,150,0.5)" />
              <stop offset="100%" stop-color="rgba(0,168,150,0.1)" />
            </linearGradient>
          </defs>
        </svg>

        <!-- Step Cards -->
        <div
          v-for="(step, index) in steps"
          :key="index"
          class="card-glow glass-card-light p-8 text-center group transition-all duration-500 cursor-default"
          :class="isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-12'"
          :style="{
            transitionDelay: isVisible ? `${0.3 + index * 0.2}s` : '0s',
            transform: `perspective(800px) rotateX(${tilts[index].x}deg) rotateY(${tilts[index].y}deg)${isVisible ? '' : ' translateY(48px)'}`,
            transition: 'transform 0.15s ease-out, opacity 0.5s var(--transition-smooth)',
          }"
          @mousemove="(e) => { handleMouseMove(index, e); handleGlowMove(index, e); }"
          @mouseleave="handleMouseLeave(index)"
        >
          <!-- Cursor glow -->
          <div
            class="absolute pointer-events-none w-[200px] h-[200px] rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-300 z-0"
            :style="{
              background: 'radial-gradient(circle, rgba(0,168,150,0.12), transparent 70%)',
              left: `${glowPositions[index].x - 100}px`,
              top: `${glowPositions[index].y - 100}px`,
            }"
          ></div>

          <!-- Animated SVG Icon -->
          <div class="w-[72px] h-[72px] rounded-2xl flex items-center justify-center mx-auto mb-6 relative z-10"
            style="background: var(--gradient-primary); box-shadow: 0 4px 20px rgba(0,168,150,.25);">
            <!-- Search icon -->
            <svg v-if="step.icon === 'search'" class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
              <circle cx="11" cy="11" r="7" class="draw-line" :class="{ 'is-drawn': isVisible }" :style="{ transitionDelay: `${0.5 + index * 0.3}s` }" />
              <line x1="16.5" y1="16.5" x2="21" y2="21" class="draw-line" :class="{ 'is-drawn': isVisible }" :style="{ transitionDelay: `${0.8 + index * 0.3}s` }" />
            </svg>
            <!-- Blueprint icon -->
            <svg v-else-if="step.icon === 'blueprint'" class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
              <rect x="3" y="3" width="18" height="18" rx="2" class="draw-line" :class="{ 'is-drawn': isVisible }" :style="{ transitionDelay: `${0.5 + index * 0.3}s` }" />
              <line x1="3" y1="9" x2="21" y2="9" class="draw-line" :class="{ 'is-drawn': isVisible }" :style="{ transitionDelay: `${0.7 + index * 0.3}s` }" />
              <line x1="9" y1="9" x2="9" y2="21" class="draw-line" :class="{ 'is-drawn': isVisible }" :style="{ transitionDelay: `${0.9 + index * 0.3}s` }" />
            </svg>
            <!-- Rocket icon -->
            <svg v-else class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
              <path d="M12 2C6.48 8 4 12 4 16l4-2 4 6 4-6 4 2c0-4-2.48-8-8-14z" class="draw-line" :class="{ 'is-drawn': isVisible }" :style="{ transitionDelay: `${0.5 + index * 0.3}s` }" />
            </svg>
          </div>

          <h3 class="text-xl font-bold text-dark-slate mb-3 relative z-10">{{ step.title }}</h3>
          <p class="text-gray-500 leading-relaxed relative z-10">
            {{ step.description }}
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
