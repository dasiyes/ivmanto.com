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

// Typewriter cycling text
const typewriterWords = ['Data Architecture', 'AI Pipelines', 'Cloud Migration', 'Real-time Analytics']
const currentWordIndex = ref(0)
const displayedText = ref('')
const isDeleting = ref(false)
let typewriterInterval: ReturnType<typeof setTimeout> | null = null

function typewriterStep() {
  const currentWord = typewriterWords[currentWordIndex.value]

  if (!isDeleting.value) {
    // Typing
    displayedText.value = currentWord.substring(0, displayedText.value.length + 1)
    if (displayedText.value === currentWord) {
      // Pause, then start deleting
      typewriterInterval = setTimeout(() => {
        isDeleting.value = true
        typewriterStep()
      }, 2200)
      return
    }
    typewriterInterval = setTimeout(typewriterStep, 80)
  } else {
    // Deleting
    displayedText.value = currentWord.substring(0, displayedText.value.length - 1)
    if (displayedText.value === '') {
      isDeleting.value = false
      currentWordIndex.value = (currentWordIndex.value + 1) % typewriterWords.length
      typewriterInterval = setTimeout(typewriterStep, 400)
      return
    }
    typewriterInterval = setTimeout(typewriterStep, 40)
  }
}

// Scroll-reveal observers
const articlesRef = ref<HTMLElement | null>(null)
const contactRef = ref<HTMLElement | null>(null)
const articlesVisible = ref(false)
const contactVisible = ref(false)

// Scroll indicator fade
const scrolled = ref(false)

let observers: IntersectionObserver[] = []

function handleScroll() {
  scrolled.value = window.scrollY > 100
}

// Horizontal scroll drag
function initDragScroll(el: HTMLElement | null) {
  if (!el) return
  let isDown = false
  let startX = 0
  let scrollLeft = 0

  el.addEventListener('mousedown', (e) => {
    isDown = true
    startX = e.pageX - el.offsetLeft
    scrollLeft = el.scrollLeft
  })
  el.addEventListener('mouseleave', () => { isDown = false })
  el.addEventListener('mouseup', () => { isDown = false })
  el.addEventListener('mousemove', (e) => {
    if (!isDown) return
    e.preventDefault()
    const x = e.pageX - el.offsetLeft
    el.scrollLeft = scrollLeft - (x - startX) * 1.5
  })
}

const scrollContainerRef = ref<HTMLElement | null>(null)

onMounted(() => {
  // Trigger hero animations after a brief delay
  setTimeout(() => { heroLoaded.value = true }, 100)
  // Start typewriter after hero loads
  setTimeout(() => typewriterStep(), 1200)

  window.addEventListener('scroll', handleScroll, { passive: true })

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

  // Init horizontal drag scroll
  initDragScroll(scrollContainerRef.value)
})

onUnmounted(() => {
  observers.forEach((o) => o.disconnect())
  window.removeEventListener('scroll', handleScroll)
  if (typewriterInterval) clearTimeout(typewriterInterval)
})
</script>

<template>
  <div>
    <!-- ═══════════════════════════════════════════════ -->
    <!-- HERO SECTION — Immersive w/ typewriter + badges -->
    <!-- ═══════════════════════════════════════════════ -->
    <section class="relative min-h-[92vh] flex items-center overflow-hidden" style="background: var(--gradient-hero);">
      <!-- Dot grid overlay -->
      <div class="absolute inset-0 dot-grid opacity-30"></div>

      <!-- Animated background blobs -->
      <div class="blob w-[600px] h-[600px] bg-primary/10 -top-40 -right-40 animate-blob-float"></div>
      <div class="blob w-[500px] h-[500px] bg-primary-dark/8 -bottom-24 -left-24 animate-blob-float-reverse"></div>
      <div class="blob w-[300px] h-[300px] bg-amber/5 top-1/3 right-1/4 animate-blob-float" style="animation-delay: -3s;"></div>

      <div class="container mx-auto px-6 relative z-10">
        <div class="grid lg:grid-cols-2 gap-16 items-center">
          <!-- Text Content -->
          <div class="text-center lg:text-left">
            <!-- Eyebrow -->
            <span
              class="inline-block text-amber font-semibold tracking-[0.2em] text-sm uppercase mb-6 transition-all duration-700"
              :class="heroLoaded ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'"
            >
              Google Cloud Platform Specialist
            </span>

            <!-- Main headline — static part -->
            <h1 class="text-4xl md:text-5xl lg:text-[3.5rem] font-extrabold leading-tight">
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

            <!-- Typewriter cycling line -->
            <div
              class="mt-5 h-10 flex items-center transition-all duration-700"
              :class="heroLoaded ? 'opacity-100' : 'opacity-0'"
              :style="{ transitionDelay: '0.55s' }"
            >
              <span class="text-xl md:text-2xl font-light text-primary-light tracking-wide">
                {{ displayedText }}<span class="typewriter-cursor"></span>
              </span>
            </div>

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

          <!-- Floating Badges Visual -->
          <div
            class="hidden lg:block transition-all duration-1000"
            :class="heroLoaded ? 'opacity-100 scale-100' : 'opacity-0 scale-90'"
            :style="{ transitionDelay: '0.4s' }"
          >
            <SectionsHeroInfographicSection :articles="featuredArticles" />
          </div>
        </div>
      </div>

      <!-- Scroll indicator -->
      <div
        class="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center gap-2 transition-opacity duration-500"
        :class="scrolled ? 'opacity-0' : 'opacity-60'"
      >
        <span class="text-xs text-gray-400 tracking-widest uppercase">Scroll to explore</span>
        <svg class="w-5 h-5 text-gray-400 scroll-indicator" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7" />
        </svg>
      </div>
    </section>

    <!-- Wave divider: hero → stats -->
    <SectionsSectionDivider class="text-hero-to" style="background: var(--gradient-hero);" />

    <!-- Stats Counter Section -->
    <SectionsStatsSection />

    <!-- Wave divider: stats → process -->
    <SectionsSectionDivider class="text-white" style="background: var(--gradient-hero);" />

    <!-- Process Section -->
    <SectionsProcessSection />

    <!-- Wave divider: process → articles -->
    <SectionsSectionDivider class="text-gray-50 bg-white" />

    <!-- ═══════════════════════════════════════════════ -->
    <!-- ARTICLES SECTION — Horizontal scroll            -->
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

        <!-- Horizontal Scroll Cards (desktop) / Grid (mobile) -->
        <div
          ref="scrollContainerRef"
          class="horizontal-scroll lg:flex lg:overflow-x-auto md:grid md:grid-cols-2 grid grid-cols-1 gap-6"
        >
          <div
            v-for="(article, index) in featuredArticles"
            :key="article.slug"
            class="gradient-border-spin glass-card-light p-6 transition-all duration-500 hover:-translate-y-2 hover:shadow-xl group lg:min-w-[340px] lg:max-w-[380px]"
            :class="articlesVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-12'"
            :style="{ transitionDelay: articlesVisible ? `${0.15 * index}s` : '0s' }"
          >
            <div class="flex items-center gap-3 mb-3">
              <span class="text-sm text-gray-400 font-medium">{{
                new Date(article.date).toLocaleDateString('en-US', {
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric',
                })
              }}</span>
              <span class="text-xs text-primary/60 bg-primary/10 px-2 py-0.5 rounded-full">~3 min read</span>
            </div>
            <h3
              class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors duration-300"
            >
              {{ article.title }}
            </h3>
            <p class="mt-3 text-gray-500 leading-relaxed line-clamp-3">{{ article.summary }}</p>
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
    <!-- CONTACT SECTION — Split layout                  -->
    <!-- ═══════════════════════════════════════════════ -->
    <section id="contact" ref="contactRef" class="py-20 md:py-28 text-white relative overflow-hidden" style="background: var(--gradient-hero);">
      <!-- Dot grid -->
      <div class="absolute inset-0 dot-grid opacity-20"></div>

      <!-- Floating blobs -->
      <div class="blob w-[400px] h-[400px] bg-primary/15 -top-20 -left-20 animate-blob-float"></div>
      <div class="blob w-[300px] h-[300px] bg-amber/10 -bottom-16 -right-16 animate-blob-float-reverse"></div>

      <div class="container mx-auto px-6 relative z-10">
        <div class="grid lg:grid-cols-2 gap-16 items-start">
          <!-- Left: Compelling copy + quick facts -->
          <div
            class="transition-all duration-700"
            :class="contactVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
          >
            <span class="text-amber font-semibold tracking-widest text-sm uppercase">Get Started</span>
            <h2 class="text-3xl md:text-5xl font-bold mt-3 leading-tight">Let's Build<br/>Something <span class="gradient-text">Great</span></h2>
            <p class="text-lg text-gray-400 mt-6 max-w-lg">
              Have a project in mind? Let's discuss how I can help transform your data landscape
              into a competitive advantage.
            </p>

            <!-- Quick facts -->
            <div class="mt-10 space-y-5">
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-xl flex items-center justify-center bg-primary/20 text-primary flex-shrink-0">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
                </div>
                <div>
                  <div class="font-semibold text-white">Fast Response</div>
                  <div class="text-sm text-gray-400">Usually within 24 hours</div>
                </div>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-xl flex items-center justify-center bg-primary/20 text-primary flex-shrink-0">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/></svg>
                </div>
                <div>
                  <div class="font-semibold text-white">NDA Available</div>
                  <div class="text-sm text-gray-400">Your data stays confidential</div>
                </div>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-xl flex items-center justify-center bg-primary/20 text-primary flex-shrink-0">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z"/></svg>
                </div>
                <div>
                  <div class="font-semibold text-white">Free Intro Call</div>
                  <div class="text-sm text-gray-400">30-minute discovery session</div>
                </div>
              </div>
            </div>
          </div>

          <!-- Right: Contact Form -->
          <div
            class="transition-all duration-700"
            :class="contactVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-8'"
            :style="{ transitionDelay: '0.2s' }"
          >
            <ContactForm source="home_page_form" />
          </div>
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
