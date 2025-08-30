<script setup lang="ts">
import HeroInfographic from '@/components/sections/HeroInfographicSection.vue'
import ContactForm from '@/components/ContactForm.vue'
import { RouterLink } from 'vue-router'
import { articles } from '@/data/articles'
import { computed } from 'vue'

// Sort articles by date to ensure the latest are featured, then take the top 3.
const featuredArticles = computed(() =>
  [...articles].sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime()).slice(0, 3),
)
</script>

<template>
  <div>
    <section class="py-16 md:py-24 bg-light-gray hero-bg-pattern overflow-hidden">
      <!-- Infographic is now first, centered, and scaled down -->
      <div
        class="flex justify-center transform scale-50 origin-top transition-transform duration-300 md:scale-75 -mb-56 md:-mb-28"
      >
        <HeroInfographic />
      </div>

      <div class="container mx-auto px-6 text-center">
        <span class="text-primary font-semibold tracking-wider"
          >GOOGLE CLOUD PLATFORM SPECIALIST</span
        >
        <h1 class="text-4xl md:text-6xl font-bold mt-4 text-dark-slate leading-tight">
          Expert Cloud Data Architecture & AI Solutions
        </h1>
        <p class="text-lg md:text-xl text-gray-600 mt-6 max-w-3xl mx-auto">
          Transforming your data into actionable insights with robust, scalable, and intelligent
          cloud platforms built on GCP.
        </p>
        <div class="mt-10 flex justify-center gap-4">
          <RouterLink
            :to="{ name: 'booking' }"
            class="bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-all text-lg"
            >Book a Consultation</RouterLink
          >
          <RouterLink
            :to="{ name: 'services' }"
            class="bg-white text-primary font-bold py-3 px-8 rounded-lg border border-gray-200 hover:bg-gray-50 transition-all text-lg"
            >Learn More</RouterLink
          >
        </div>
      </div>
    </section>

    <section id="articles" class="py-20 md:py-28">
      <div class="container mx-auto px-6">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-dark-slate">Insights & Articles</h2>
          <p class="text-lg text-gray-600 mt-4 max-w-2xl mx-auto">
            Sharing knowledge from real-world projects and my vision for the future of data.
          </p>
        </div>
        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="article in featuredArticles"
            :key="article.slug"
            class="border border-gray-200 rounded-xl overflow-hidden group"
          >
            <div class="p-6">
              <span class="text-sm text-gray-500">{{
                new Date(article.date).toLocaleDateString('en-US', {
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric',
                })
              }}</span>
              <h3
                class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors"
              >
                {{ article.title }}
              </h3>
              <p class="mt-3 text-gray-600">{{ article.summary }}</p>
              <RouterLink
                :to="`/blog/${article.slug}`"
                class="text-primary font-semibold mt-4 inline-block"
                >Read More →</RouterLink
              >
            </div>
          </div>
        </div>
        <div class="text-center mt-16">
          <RouterLink
            to="/blog"
            class="bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-all text-lg"
            >View All Articles</RouterLink
          >
        </div>
        <div class="mt-24 bg-light-gray p-8 md:p-12 rounded-xl text-center">
          <h3 class="text-2xl md:text-3xl font-bold text-dark-slate">Need Inspiration?</h3>
          <p class="text-lg text-gray-600 mt-3 max-w-xl mx-auto">
            Enter a topic below and our AI will generate some creative article ideas for you.
          </p>
          <div class="mt-6 max-w-lg mx-auto flex flex-col sm:flex-row gap-4">
            <input
              type="text"
              placeholder="e.g., 'AI in retail'"
              class="w-full bg-white border-gray-300 rounded-md py-3 px-4 focus:ring-accent focus:border-accent text-lg"
            /><button
              class="bg-accent text-white font-bold py-3 px-6 rounded-lg hover:bg-opacity-90 transition-all text-lg whitespace-nowrap flex items-center justify-center"
            >
              <span>✨ Generate Ideas</span>
            </button>
          </div>
          <div class="mt-8 text-left"><!--v-if--></div>
        </div>
      </div>
    </section>

    <section id="contact" class="py-20 md:py-28 bg-dark-slate text-white">
      <div class="container mx-auto px-6">
        <div class="text-center mb-12">
          <h2 class="text-3xl md:text-4xl font-bold">Let's Build Something Great Together</h2>
          <p class="text-lg text-gray-300 mt-4 max-w-2xl mx-auto">
            Have a project in mind? Let me know the details and I'll draft your inquiry.
          </p>
        </div>
        <ContactForm />
      </div>
    </section>
  </div>
</template>

<style>
.hero-bg-pattern {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' width='100' height='100' viewBox='0 0 200 200'%3e%3crect fill='%23F8F9FA' width='200' height='200'/%3e%3cg fill='none' stroke='%239CA3AF' stroke-width='1'%3e%3cpath d='M100 0L200 100 100 200 0 100z'/%3e%3cpath d='M100 50L150 100 100 150 50 100z'/%3e%3c/g%3e%3c/svg%3e");
}
</style>
