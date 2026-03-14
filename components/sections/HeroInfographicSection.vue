<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface Article {
  slug: string
  title: string
  summary: string
  date: string
}

const props = defineProps<{
  articles: Article[]
}>()

const isVisible = ref(false)

onMounted(() => {
  setTimeout(() => { isVisible.value = true }, 600)
})

// Staggered offsets for "randomly distributed" feel
const cardStyles = [
  { top: '0%', right: '0', transform: 'rotate(-1.5deg)', delay: '0.3s' },
  { top: '34%', right: '40px', transform: 'rotate(1deg)', delay: '0.5s' },
  { top: '66%', right: '10px', transform: 'rotate(-0.5deg)', delay: '0.7s' },
]
</script>

<template>
  <div class="relative w-full" style="height: 650px;">
    <!-- Particle field background -->
    <div class="particle-field">
      <div
        v-for="i in 8"
        :key="i"
        class="particle"
        :style="{
          left: `${(i * 12) % 100}%`,
          top: `${(i * 20 + 10) % 100}%`,
          width: '2px',
          height: '2px',
          animationDuration: `${7 + (i % 4) * 2}s`,
          animationDelay: `${(i % 5) * -2}s`,
        }"
      ></div>
    </div>

    <!-- Article Cards — staggered layout -->
    <div
      v-for="(article, index) in (articles || []).slice(0, 3)"
      :key="article.slug"
      class="absolute w-[350px] transition-all duration-700 group cursor-pointer"
      :class="isVisible ? 'opacity-100' : 'opacity-0 translate-y-8'"
      :style="{
        top: cardStyles[index]?.top,
        right: cardStyles[index]?.right,
        transform: isVisible ? cardStyles[index]?.transform : 'translateY(32px)',
        transitionDelay: cardStyles[index]?.delay,
      }"
    >
      <NuxtLink :to="`/blog/${article.slug}`" class="block">
        <div class="rounded-2xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300 hover:-translate-y-1"
          style="background: rgba(255,255,255,0.92); border: 1px solid rgba(0,168,150,0.12);">
          <!-- Teal title strip -->
          <div class="px-5 py-3" style="background: var(--gradient-primary);">
            <span class="text-white text-sm font-semibold tracking-wider uppercase">Latest Article</span>
          </div>
          <!-- Card body -->
          <div class="px-5 py-5">
            <h3 class="text-base font-bold text-gray-800 leading-snug group-hover:text-primary transition-colors duration-300 line-clamp-2">
              {{ article.title }}
            </h3>
            <p class="text-sm text-gray-500 mt-2.5 leading-relaxed line-clamp-2">
              {{ article.summary }}
            </p>
            <div class="flex items-center justify-between mt-4">
              <span class="text-sm text-gray-400">{{
                new Date(article.date).toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
              }}</span>
              <span class="text-sm text-primary font-semibold group-hover:translate-x-1 transition-transform duration-300">Read →</span>
            </div>
          </div>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>
