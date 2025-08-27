<script setup lang="ts">
import HeroInfographic from '@/components/sections/HeroInfographicSection.vue'
</script>

<template>
  <main>
    <section class="py-20 md:py-28 bg-light-gray hero-bg-pattern relative">
      <HeroInfographic />
      <div class="container mx-auto px-6 text-center relative z-10">
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
          <a
            href="#contact"
            class="bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-all text-lg"
            >Book a Consultation</a
          ><a
            href="#services"
            class="bg-white text-primary font-bold py-3 px-8 rounded-lg border border-gray-200 hover:bg-gray-50 transition-all text-lg"
            >Learn More</a
          >
        </div>
      </div>
    </section>

    <section id="about" class="py-20 md:py-28 bg-light-gray">
      <div class="container mx-auto px-6 grid md:grid-cols-2 gap-12 items-center">
        <div>
          <img
            src="/src/assets/nt.jpg"
            alt="IVMANTO"
            class="rounded-xl shadow-lg w-full max-w-md mx-auto"
          />
        </div>
        <div>
          <h2 class="text-3xl md:text-4xl font-bold text-dark-slate">
            A Partner in Your Data Journey
          </h2>
          <p class="text-lg text-gray-600 mt-6">
            I am a dedicated Cloud Data Architect with extensive experience on the Google Cloud
            Platform. My passion is helping businesses unlock the true value of their data. I focus
            on building solutions that are not just technologically sound, but also align perfectly
            with your business goals.
          </p>
          <p class="text-lg text-gray-600 mt-4">
            I'm currently on the path to becoming a DAMA Certified Data Management Professional,
            ensuring that my work adheres to the highest industry standards for data governance and
            quality.
          </p>
          <div class="mt-8 flex flex-wrap gap-4">
            <a
              href="https://linkedin.com/in/nikolaytonev"
              class="bg-white border border-gray-300 text-dark-slate font-medium py-2 px-4 rounded-lg hover:bg-gray-50 transition-all flex items-center gap-2"
              target="_blank"
              rel="noopener noreferrer"
              >LinkedIn</a
            ><a
              href="https://github.com/dasiyes"
              class="bg-white border border-gray-300 text-dark-slate font-medium py-2 px-4 rounded-lg hover:bg-gray-50 transition-all flex items-center gap-2"
              target="_blank"
              rel="noopener noreferrer"
              >GitHub</a
            ><a
              href="https://www.credly.com/users/nikolay-tonev.a60be769/badges#credly"
              class="bg-white border border-gray-300 text-dark-slate font-medium py-2 px-4 rounded-lg hover:bg-gray-50 transition-all flex items-center gap-2"
              target="_blank"
              rel="noopener noreferrer"
              >Credly Profile</a
            ><a
              href="https://cdmp.info/"
              class="bg-white border border-gray-300 text-dark-slate font-medium py-2 px-4 rounded-lg hover:bg-gray-50 transition-all flex items-center gap-2"
              target="_blank"
              rel="noopener noreferrer"
              >DAMA Professionals</a
            >
          </div>
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
          <div class="border border-gray-200 rounded-xl overflow-hidden group">
            <div class="p-6">
              <span class="text-sm text-gray-500">Data Architecture • 10 min read</span>
              <h3
                class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors"
              >
                Common Architectural Patterns for Modern Data Platforms
              </h3>
              <p class="mt-3 text-gray-600">
                A deep dive into the pros and cons of Data Lakes, Warehouses, and the modern
                Lakehouse architecture on GCP.
              </p>
              <a href="#" class="text-primary font-semibold mt-4 inline-block">Read More →</a>
            </div>
          </div>
          <div class="border border-gray-200 rounded-xl overflow-hidden group">
            <div class="p-6">
              <span class="text-sm text-gray-500">AI/ML • 8 min read</span>
              <h3
                class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors"
              >
                From Notebook to Production with Vertex AI
              </h3>
              <p class="mt-3 text-gray-600">
                A practical guide to operationalizing your machine learning models using Google
                Cloud's unified AI platform.
              </p>
              <a href="#" class="text-primary font-semibold mt-4 inline-block">Read More →</a>
            </div>
          </div>
          <div class="border border-gray-200 rounded-xl overflow-hidden group">
            <div class="p-6">
              <span class="text-sm text-gray-500">Data Governance • 12 min read</span>
              <h3
                class="text-xl font-bold text-dark-slate mt-2 group-hover:text-primary transition-colors"
              >
                Why DAMA Principles Matter for Your Business
              </h3>
              <p class="mt-3 text-gray-600">
                Exploring how standardized data management practices can reduce risk and increase
                the value of your data assets.
              </p>
              <a href="#" class="text-primary font-semibold mt-4 inline-block">Read More →</a>
            </div>
          </div>
        </div>
        <div class="mt-20 bg-light-gray p-8 md:p-12 rounded-xl text-center">
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
            Have a project in mind? Let my AI assistant help you draft your inquiry.
          </p>
        </div>
        <div class="max-w-xl mx-auto bg-white/10 p-6 md:p-8 rounded-xl backdrop-blur-sm">
          <div>
            <div class="h-80 overflow-y-auto pr-2 space-y-4 mb-4">
              <div class="flex justify-start">
                <div class="p-3 rounded-lg max-w-xs bg-secondary text-white">
                  Hello! To help me understand your needs, please briefly describe the project or
                  challenge you're facing.
                </div>
              </div>
            </div>
            <div class="flex gap-4">
              <input
                type="text"
                placeholder="Type your response..."
                class="flex-grow bg-white/20 border-transparent rounded-md py-2 px-3 text-white focus:ring-accent focus:border-accent"
              /><button
                class="bg-accent text-white font-bold py-2 px-5 rounded-lg hover:bg-opacity-90 transition-all flex items-center justify-center w-24"
              >
                <span>Send</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>

<style>
.hero-bg-pattern {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' width='100' height='100' viewBox='0 0 200 200'%3e%3crect fill='%23F8F9FA' width='200' height='200'/%3e%3cg fill='none' stroke='%239CA3AF' stroke-width='1'%3e%3cpath d='M100 0L200 100 100 200 0 100z'/%3e%3cpath d='M100 50L150 100 100 150 50 100z'/%3e%3c/g%3e%3c/svg%3e");
}
</style>
