<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { trackEvent } from '~/services/analytics'
import { generateInspirationIdeas, type Idea } from '~/services/api'

useSeoMeta({
  title: 'ivmanto.com | Data Consultancy & AI Solutions',
  description:
    'Expert data consultancy on Google Cloud Platform. Providing data architecture, governance, and AI solutions that turn your data into a strategic asset.',
  ogTitle: 'ivmanto.com | Data Consultancy & AI Solutions',
  ogDescription:
    'Expert data consultancy on Google Cloud Platform. Providing data architecture, governance, and AI solutions that turn your data into a strategic asset.',
  twitterTitle: 'ivmanto.com | Data Consultancy & AI Solutions',
  twitterDescription:
    'Expert data consultancy on Google Cloud Platform. Providing data architecture, governance, and AI solutions.',
})

// FAQPage schema for SEO
useHead({
  script: [
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'FAQPage',
        mainEntity: [
          {
            '@type': 'Question',
            name: 'Do you work with existing on-premise infrastructure?',
            acceptedAnswer: {
              '@type': 'Answer',
              text: 'Yes. I specialize in hybrid cloud architectures and can help you seamlessly bridge your on-premise systems with Google Cloud Platform, or lead a full migration strategy.',
            },
          },
          {
            '@type': 'Question',
            name: 'How long does a typical architecture audit take?',
            acceptedAnswer: {
              '@type': 'Answer',
              text: 'A standard initial assessment usually takes 1 to 2 weeks, depending on the complexity of your current landscape and the availability of your technical team for discovery sessions.',
            },
          },
          {
            '@type': 'Question',
            name: 'Do you provide hands-on engineering or just strategy?',
            acceptedAnswer: {
              '@type': 'Answer',
              text: "I do both. I believe the best architects are those who still know how to build. I provide the high-level strategic blueprint and can lead the hands-on implementation of pipelines and infrastructure.",
            },
          },
          {
            '@type': 'Question',
            name: 'How do you ensure data security and compliance?',
            acceptedAnswer: {
              '@type': 'Answer',
              text: "Security is 'baked in' from day one. I implement 'Secure by Design' principles, following the latest CIS benchmarks for GCP and ensuring alignment with GDPR or industry-specific regulations.",
            },
          },
        ],
      }),
    },
  ],
})

const { sortedArticles, fetchArticles } = useArticles()

await useAsyncData('home-articles', () => fetchArticles())

const featuredArticles = computed(() => sortedArticles.value.slice(0, 3))

// State for the "Need Inspiration" feature
const topic = ref('')
const isLoading = ref(false)
const generatedIdeas = ref<Idea[]>([])
const error = ref<string | null>(null)
const isModalOpen = ref(false)

async function handleGenerateIdeas() {
  if (!topic.value.trim()) return
  trackEvent('generate_inspiration_ideas', { topic: topic.value })
  isModalOpen.value = true
  isLoading.value = true
  error.value = null
  generatedIdeas.value = []
  try {
    generatedIdeas.value = await generateInspirationIdeas(topic.value)
  } catch (e: any) {
    console.error('Failed to generate ideas:', e)
    error.value = e.message || 'An unexpected error occurred. Please try again later.'
  } finally {
    isLoading.value = false
  }
}

function trackBookConsultationClick() {
  trackEvent('click_book_consultation', { source: 'home_hero_cta' })
}

// Hero animation state
const heroLoaded = ref(false)

// Scroll-reveal observer
const articlesRef = ref<HTMLElement | null>(null)
const contactRef = ref<HTMLElement | null>(null)
const articlesVisible = ref(false)
const contactVisible = ref(false)

let observers: IntersectionObserver[] = []

onMounted(() => {
  // Trigger hero animations after a brief delay
  setTimeout(() => { heroLoaded.value = true }, 100)

  // Scroll observers
  const createObserver = (el: HTMLElement | null, flag: { value: boolean }) => {
    if (!el) return null
    const obs = new IntersectionObserver(
      (entries) => entries.forEach((e) => { if (e.isIntersecting) flag.value = true }),
      { threshold: 0.1 }
    )
    obs.observe(el)
    return obs
  }

  const o1 = createObserver(articlesRef.value, articlesVisible)
  const o2 = createObserver(contactRef.value, contactVisible)
  if (o1) observers.push(o1)
  if (o2) observers.push(o2)
})

onUnmounted(() => {
  observers.forEach((o) => o.disconnect())
})
</script>

<template>
  <div>
    <!-- ═══════════════════════════════════════════════ -->
    <!-- HERO SECTION — Dark gradient w/ kinetic text   -->
    <!-- ═══════════════════════════════════════════════ -->
    <section class="relative min-h-[90vh] flex items-center overflow-hidden" style="background: var(--gradient-hero);">
      <!-- Dot grid overlay -->
      <div class="absolute inset-0 dot-grid opacity-40"></div>

      <!-- Animated background blobs -->
      <div class="blob w-[500px] h-[500px] bg-primary/20 -top-32 -right-32 animate-blob-float"></div>
      <div class="blob w-[400px] h-[400px] bg-primary-dark/15 -bottom-20 -left-20 animate-blob-float-reverse"></div>
      <div class="blob w-[250px] h-[250px] bg-amber/10 top-1/3 right-1/4 animate-blob-float" style="animation-delay: -3s;"></div>

      <div class="container mx-auto px-6 relative z-10">
        <div class="grid lg:grid-cols-2 gap-12 items-center">
          <!-- Text Content -->
          <div class="text-center lg:text-left">
            <!-- Eyebrow -->
            <span
              class="inline-block text-amber font-semibold tracking-[0.2em] text-sm uppercase mb-6 transition-all duration-700"
              :class="heroLoaded ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'"
            >
              Google Cloud Platform Specialist
            </span>

            <!-- Main headline with kinetic reveal -->
            <h1 class="text-4xl md:text-5xl lg:text-6xl font-extrabold leading-tight">
              <span
                class="block text-white transition-all duration-700"
                :class="heroLoaded ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'"
                :style="{ transitionDelay: '0.15s' }"
              >
                <span class="gradient-text">Data</span> Consultancy
              </span>
              <span
                class="block text-white mt-2 transition-all duration-700"
                :class="heroLoaded ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'"
                :style="{ transitionDelay: '0.35s' }"
              >
                & <span class="gradient-text">AI</span> Solutions
              </span>
            </h1>

            <!-- Subheadline -->
            <p
              class="text-lg md:text-xl text-gray-400 mt-6 max-w-xl transition-all duration-700"
              :class="heroLoaded ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'"
              :style="{ transitionDelay: '0.5s' }"
            >
              Transforming your data into actionable insights with robust, scalable, and intelligent
              cloud platforms built on GCP.
            </p>

            <!-- CTA Buttons -->
            <div
              class="mt-10 flex flex-col sm:flex-row gap-4 justify-center lg:justify-start transition-all duration-700"
              :class="heroLoaded ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'"
              :style="{ transitionDelay: '0.65s' }"
            >
              <NuxtLink
                :to="{ name: 'booking' }"
                @click="trackBookConsultationClick"
                class="btn-glow text-center"
              >Book a Consultation</NuxtLink>
              <NuxtLink
                :to="{ name: 'services' }"
                class="btn-outline-glow text-center"
              >Explore Services</NuxtLink>
            </div>
          </div>

          <!-- Animated Visual -->
          <div
            class="hidden lg:block transition-all duration-1000"
            :class="heroLoaded ? 'opacity-100 scale-100' : 'opacity-0 scale-90'"
            :style="{ transitionDelay: '0.4s' }"
          >
            <SectionsHeroInfographicSection />
          </div>
        </div>
      </div>
    </section>

    <!-- Wave divider: hero → process -->
    <SectionsSectionDivider class="text-white bg-hero-to" />

    <!-- Process Section -->
    <SectionsProcessSection />

    <!-- Wave divider: process → articles -->
    <SectionsSectionDivider class="text-gray-50 bg-white" />

    <!-- ═══════════════════════════════════════════════ -->
    <!-- ARTICLES SECTION                                -->
    <!-- ═══════════════════════════════════════════════ -->
    <section id="articles" ref="articlesRef" class="py-20 md:py-28 bg-gray-50 relative overflow-hidden">
      <!-- Subtle background accents -->
      <div class="absolute top-20 left-0 w-96 h-96 rounded-full bg-primary/5 blur-3xl pointer-events-none"></div>
      <div class="absolute bottom-20 right-0 w-72 h-72 rounded-full bg-amber/5 blur-3xl pointer-events-none"></div>

      <div class="container mx-auto px-6 relative z-10">
        <div
          class="text-center mb-16 transition-all duration-700"
          :class="articlesVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
        >
          <span class="text-primary font-semibold tracking-widest text-sm uppercase">Blog</span>
          <h2 class="text-3xl md:text-5xl font-bold text-dark-slate mt-3">Insights & Articles</h2>
          <p class="text-lg text-gray-500 mt-4 max-w-2xl mx-auto">
            Sharing knowledge from real-world projects and my vision for the future of data.
          </p>
        </div>

        <!-- Article Cards -->
        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="(article, index) in featuredArticles"
            :key="article.slug"
            class="glass-card-light p-6 transition-all duration-500 hover:-translate-y-2 hover:shadow-xl group"
            :class="articlesVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-12'"
            :style="{ transitionDelay: articlesVisible ? `${0.15 * index}s` : '0s' }"
          >
            <span class="text-sm text-gray-400 font-medium">{{
              new Date(article.date).toLocaleDateString('en-US', {
                year: 'numeric',
                month: 'long',
                day: 'numeric',
              })
            }}</span>
            <h3
              class="text-xl font-bold text-dark-slate mt-3 group-hover:text-primary transition-colors duration-300"
            >
              {{ article.title }}
            </h3>
            <p class="mt-3 text-gray-500 leading-relaxed">{{ article.summary }}</p>
            <NuxtLink
              :to="`/blog/${article.slug}`"
              class="text-primary font-semibold mt-4 inline-flex items-center gap-1 group/link"
            >
              Read More
              <span class="transition-transform duration-300 group-hover/link:translate-x-1">→</span>
            </NuxtLink>
          </div>
        </div>

        <div
          class="text-center mt-16 transition-all duration-700"
          :class="articlesVisible ? 'opacity-100 translate-y-0 delay-500' : 'opacity-0 translate-y-8'"
        >
          <NuxtLink
            to="/blog"
            class="inline-block font-bold py-3 px-8 rounded-xl text-lg transition-all duration-300 text-primary border-2 border-primary/20 hover:border-primary hover:bg-primary hover:text-white"
          >View All Articles</NuxtLink>
        </div>

        <!-- Need Inspiration Section -->
        <div
          class="mt-24 glass-card-light p-8 md:p-12 text-center transition-all duration-700"
          :class="articlesVisible ? 'opacity-100 translate-y-0 delay-700' : 'opacity-0 translate-y-8'"
        >
          <span class="text-amber font-semibold tracking-widest text-sm uppercase">AI-Powered</span>
          <h3 class="text-2xl md:text-3xl font-bold text-dark-slate mt-2">Need Inspiration?</h3>
          <p class="text-lg text-gray-500 mt-3 max-w-xl mx-auto">
            Enter a topic below and our AI will generate some creative article ideas for you.
          </p>
          <form
            @submit.prevent="handleGenerateIdeas"
            class="mt-6 max-w-lg mx-auto flex flex-col sm:flex-row gap-4"
          >
            <input
              v-model="topic"
              type="text"
              placeholder="e.g., 'AI in retail'"
              class="w-full bg-white border border-gray-200 rounded-xl py-3 px-4 focus:ring-2 focus:ring-primary/30 focus:border-primary text-lg transition-all duration-300 focus:shadow-lg focus:shadow-primary/5"
            />
            <button
              type="submit"
              class="btn-glow whitespace-nowrap flex items-center justify-center"
              :disabled="isLoading"
            >
              <span v-if="!isLoading">Generate Ideas</span>
              <span v-else>Generating...</span>
            </button>
          </form>
        </div>
      </div>
    </section>

    <!-- FAQ Section -->
    <SectionsFAQSection />

    <!-- ═══════════════════════════════════════════════ -->
    <!-- CONTACT SECTION — Gradient with floating blobs  -->
    <!-- ═══════════════════════════════════════════════ -->
    <section id="contact" ref="contactRef" class="py-20 md:py-28 text-white relative overflow-hidden" style="background: var(--gradient-hero);">
      <!-- Dot grid -->
      <div class="absolute inset-0 dot-grid opacity-20"></div>

      <!-- Floating blobs -->
      <div class="blob w-[400px] h-[400px] bg-primary/15 -top-20 -left-20 animate-blob-float"></div>
      <div class="blob w-[300px] h-[300px] bg-amber/10 -bottom-16 -right-16 animate-blob-float-reverse"></div>

      <div class="container mx-auto px-6 relative z-10">
        <div
          class="text-center mb-12 transition-all duration-700"
          :class="contactVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
        >
          <span class="text-amber font-semibold tracking-widest text-sm uppercase">Get Started</span>
          <h2 class="text-3xl md:text-5xl font-bold mt-3">Let's Build Something Great Together</h2>
          <p class="text-lg text-gray-400 mt-4 max-w-2xl mx-auto">
            Have a project in mind? Let me know the details and I'll draft your inquiry.
          </p>
        </div>
        <div
          class="transition-all duration-700"
          :class="contactVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
          :style="{ transitionDelay: '0.2s' }"
        >
          <ContactForm source="home_page_form" />
        </div>
      </div>
    </section>

    <InspirationModal
      :is-open="isModalOpen"
      :is-loading="isLoading"
      :ideas="generatedIdeas"
      :error="error"
      :topic="topic"
      @close="isModalOpen = false"
    />
  </div>
</template>
