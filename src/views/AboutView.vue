<script setup lang="ts">
import { RouterLink } from 'vue-router'
import ExpertiseSection from '@/components/about/ExpertiseSection.vue'
import ApproachSection from '@/components/about/ApproachSection.vue'
import CertificatesSection from '@/components/about/CertificatesSection.vue'
</script>

<template>
  <div>
    <section class="py-20 md:py-28 bg-light-gray">
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

    <ApproachSection />

    <ExpertiseSection />

    <CertificatesSection />

    <!-- Call to Action Section -->
    <section class="bg-primary text-white">
      <div class="container mx-auto px-6 py-20 text-center">
        <h2 class="text-3xl md:text-4xl font-bold">Ready to Unlock Your Data's Potential?</h2>
        <p class="text-lg mt-4 max-w-2xl mx-auto">
          Let's discuss how we can build a data strategy that drives growth and innovation for your
          business.
        </p>
        <RouterLink
          to="/#contact"
          class="mt-8 inline-block bg-white text-primary font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-all"
          >Get In Touch</RouterLink
        >
      </div>
    </section>
  </div>
</template>
