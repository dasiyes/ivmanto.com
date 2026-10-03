<script setup lang="ts">
import { offerServices, expertiseServices, coreStack, getServiceById } from '~/data/services'
import { trackEvent } from '~/services/analytics'

const pageTitle = 'Freelance Cloud Data & AI Agent Engineer for SMBs | ivmanto.com'
const pageDescription =
  'Hands-on cloud data engineering, custom AI agents and private RAG knowledge systems for small and medium businesses. Senior independent engineer based in Germany, serving the EU and clients worldwide.'

useSeoMeta({
  title: pageTitle,
  description: pageDescription,
  ogTitle: pageTitle,
  ogDescription: pageDescription,
  twitterTitle: pageTitle,
  twitterDescription: pageDescription,
})

const audit = getServiceById('ai-readiness-audit')
const pillars = offerServices.filter((s) => s.tier === 'core')
const retainer = offerServices.find((s) => s.tier === 'ongoing')

const reasons = [
  {
    title: 'Direct senior access',
    text: 'You work with the engineer who designs and writes the code from day one. No account managers, no hand-off to juniors.',
  },
  {
    title: 'You own everything',
    text: 'Code, pipelines, cloud project and keys stay with you. No proprietary platform and no monthly license markup.',
  },
  {
    title: 'Data first',
    text: 'Agents fail on messy data. I fix the foundation before building on it, so what goes live keeps working.',
  },
  {
    title: 'Guardrails built in',
    text: 'Narrow permissions, human approval for sensitive actions, logging and tests against your real cases.',
  },
  {
    title: 'EU hosting by default',
    text: 'Based in Germany. Systems run in EU regions and are designed around GDPR. Clients outside the EU are welcome too.',
  },
]

const steps = [
  {
    title: 'Readiness audit',
    duration: '1 to 2 days',
    text: 'Review your workflows and data, rank use cases by ROI and scope the first build.',
  },
  {
    title: 'Prototype and evaluation',
    duration: '1 to 2 weeks',
    text: 'A working end-to-end prototype on your real data, measured against real cases.',
  },
  {
    title: 'Production deployment',
    duration: '2 to 4 weeks',
    text: 'Integration with your stack, CI/CD, guardrails, monitoring and logging.',
  },
  {
    title: 'Handoff and support',
    duration: 'Ongoing',
    text: 'Full documentation and repository handoff, plus an optional monthly retainer.',
  },
]

const faqs = [
  {
    question: 'Why hire an independent engineer instead of an AI automation agency?',
    answer:
      'With an agency you often pay for account management and get junior developers on the build. With me you talk directly to the senior engineer doing the work, iterate faster and keep full ownership of the code and infrastructure.',
  },
  {
    question: 'What does it cost to get started?',
    answer:
      'Every engagement starts with the Data & AI Readiness Audit, from €100. If we sign the project, the audit fee is fully credited, so the audit is free. You then get a fixed-price proposal for the first build.',
  },
  {
    question: 'Our data is messy. Can we still start with AI?',
    answer:
      'Yes, and it is the most common starting point. The audit shows what needs cleaning first, and Cloud Data Engineering makes the data reliable enough for agents to use.',
  },
  {
    question: 'Where is our data hosted?',
    answer:
      'In your own Google Cloud project, in an EU region by default. Your data is not used to train public models.',
  },
  {
    question: 'Do you work with companies outside Germany and the EU?',
    answer: 'Yes. I focus on the EU and DACH region and work remotely with clients worldwide.',
  },
]

useHead({
  script: [
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'ItemList',
        name: 'Cloud Data & AI Engineering Services',
        itemListElement: offerServices.map((s, i) => ({
          '@type': 'ListItem',
          position: i + 1,
          name: s.menuTitle,
          description: s.summary,
          url: `https://ivmanto.com/services/${s.id}`,
        })),
      }),
    },
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'FAQPage',
        mainEntity: faqs.map((f) => ({
          '@type': 'Question',
          name: f.question,
          acceptedAnswer: { '@type': 'Answer', text: f.answer },
        })),
      }),
    },
  ],
})

function trackCta(source: string) {
  trackEvent('click_book_consultation', { source })
}
</script>

<template>
  <div class="container mx-auto px-6 py-12">
    <!-- Hero -->
    <section class="text-center max-w-3xl mx-auto">
      <p class="text-sm font-bold uppercase tracking-wider text-primary">
        Independent senior engineer for small and medium businesses
      </p>
      <h1 class="mt-3 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
        Hands-on Cloud Data & AI Agent Engineering
      </h1>
      <p class="mt-6 text-lg leading-8 text-gray-600">
        From reliable data pipelines to AI agents that do real work in your tools. Built and
        deployed in your own cloud, with no agency overhead and no vendor lock-in.
      </p>
      <div class="mt-8 flex flex-col sm:flex-row gap-4 justify-center">
        <NuxtLink
          to="/services/ai-readiness-audit"
          @click="
            trackEvent('view_service_details', {
              service_id: 'ai-readiness-audit',
              source: 'services_hero',
            })
          "
          class="bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-colors"
        >
          Start with the Readiness Audit
        </NuxtLink>
        <NuxtLink
          :to="{ name: 'booking' }"
          @click="trackCta('services_hero')"
          class="border border-primary text-primary font-bold py-3 px-8 rounded-lg hover:bg-primary/5 transition-colors"
        >
          Book a 30-minute call
        </NuxtLink>
      </div>
    </section>

    <!-- Entry offer -->
    <section v-if="audit" class="mt-16 max-w-4xl mx-auto">
      <NuxtLink
        :to="`/services/${audit.id}`"
        class="relative block rounded-xl border-2 border-amber bg-amber/5 p-6 md:p-8 shadow-md hover:shadow-lg transition-shadow"
      >
        <span
          class="absolute -top-3 left-6 inline-flex items-center bg-amber text-white text-xs font-bold uppercase tracking-wide px-3 py-1 rounded-full shadow-sm"
        >
          {{ audit.eyebrow }}
        </span>
        <div class="flex flex-col md:flex-row md:items-center gap-6 justify-between">
          <div>
            <h2 class="text-2xl font-bold text-dark-slate">{{ audit.menuTitle }}</h2>
            <p class="mt-2 text-gray-700">{{ audit.summary }}</p>
          </div>
          <div v-if="audit.price" class="md:text-right flex-shrink-0 md:max-w-[14rem]">
            <p class="text-3xl font-bold text-dark-slate whitespace-nowrap">{{ audit.price.label }}</p>
            <p class="mt-1 text-sm text-gray-600">{{ audit.price.note }}</p>
          </div>
        </div>
      </NuxtLink>
    </section>

    <!-- Pillars -->
    <section class="mt-16">
      <h2 class="text-center text-3xl font-bold text-dark-slate">What I build</h2>
      <div class="mt-8 grid grid-cols-1 md:grid-cols-3 gap-8">
        <div
          v-for="pillar in pillars"
          :key="pillar.id"
          class="flex flex-col rounded-lg border border-gray-200 p-6 shadow-sm"
        >
          <p class="text-xs font-bold uppercase tracking-wider text-primary">
            {{ pillar.eyebrow }}
          </p>
          <h3 class="mt-2 text-lg font-semibold text-dark-slate">{{ pillar.menuTitle }}</h3>
          <p class="mt-2 text-sm text-gray-600 flex-grow">{{ pillar.summary }}</p>
          <div v-if="pillar.useCases?.length" class="mt-4">
            <p class="text-xs font-semibold uppercase tracking-wide text-gray-500">For example</p>
            <ul class="mt-2 space-y-1">
              <li
                v-for="useCase in pillar.useCases.slice(0, 2)"
                :key="useCase"
                class="text-sm text-gray-700"
              >
                {{ useCase }}
              </li>
            </ul>
          </div>
          <NuxtLink
            :to="`/services/${pillar.id}`"
            class="mt-4 inline-block text-primary font-semibold hover:underline"
          >
            Learn More &rarr;
          </NuxtLink>
        </div>
      </div>

      <NuxtLink
        v-if="retainer"
        :to="`/services/${retainer.id}`"
        class="mt-8 block rounded-lg border border-primary/30 bg-primary/5 p-6 hover:shadow-md transition-shadow"
      >
        <p class="text-xs font-bold uppercase tracking-wider text-primary">
          {{ retainer.eyebrow }}
        </p>
        <h3 class="mt-1 text-lg font-semibold text-dark-slate">{{ retainer.menuTitle }} &rarr;</h3>
        <p class="mt-1 text-sm text-gray-600">{{ retainer.summary }}</p>
      </NuxtLink>
    </section>

    <!-- Why independent -->
    <section class="mt-20 max-w-5xl mx-auto">
      <h2 class="text-center text-3xl font-bold text-dark-slate">
        Why work with an independent senior engineer
      </h2>
      <p class="mt-4 text-center text-gray-600 max-w-3xl mx-auto">
        Most AI initiatives stall because the data underneath is fragmented, or because a demo
        prototype breaks on real edge cases. I design, build and deploy the whole system myself.
      </p>
      <div class="mt-8 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="reason in reasons" :key="reason.title" class="rounded-lg bg-light-gray p-5">
          <h3 class="font-semibold text-dark-slate">{{ reason.title }}</h3>
          <p class="mt-1 text-sm text-gray-600">{{ reason.text }}</p>
        </div>
      </div>
    </section>

    <!-- Process -->
    <section class="mt-20 max-w-5xl mx-auto">
      <h2 class="text-center text-3xl font-bold text-dark-slate">How an engagement works</h2>
      <ol class="mt-8 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <li
          v-for="(step, i) in steps"
          :key="step.title"
          class="rounded-lg border border-gray-200 p-5"
        >
          <p class="text-3xl font-bold text-primary/40">{{ String(i + 1).padStart(2, '0') }}</p>
          <h3 class="mt-2 font-semibold text-dark-slate">{{ step.title }}</h3>
          <p class="text-xs font-semibold uppercase tracking-wide text-amber-dark">
            {{ step.duration }}
          </p>
          <p class="mt-2 text-sm text-gray-600">{{ step.text }}</p>
        </li>
      </ol>
    </section>

    <!-- Stack -->
    <section class="mt-20 max-w-4xl mx-auto text-center">
      <h2 class="text-2xl font-bold text-dark-slate">Tools I work with</h2>
      <div class="mt-6 flex flex-wrap justify-center gap-2">
        <span
          v-for="tool in coreStack"
          :key="tool"
          class="text-sm font-medium text-primary-dark bg-light-gray border border-gray-200 rounded-full px-3 py-1"
        >
          {{ tool }}
        </span>
      </div>
    </section>

    <!-- FAQ -->
    <section class="mt-20 max-w-3xl mx-auto">
      <h2 class="text-2xl font-bold text-dark-slate">Frequently asked questions</h2>
      <div class="mt-6 divide-y divide-gray-200 border-y border-gray-200">
        <details v-for="faq in faqs" :key="faq.question" class="group py-4">
          <summary
            class="flex cursor-pointer list-none items-center justify-between gap-4 font-semibold text-dark-slate"
          >
            {{ faq.question }}
            <span class="text-primary transition-transform group-open:rotate-45" aria-hidden="true"
              >+</span
            >
          </summary>
          <p class="mt-2 text-gray-600">{{ faq.answer }}</p>
        </details>
      </div>
    </section>

    <!-- Expertise -->
    <section class="mt-20 max-w-4xl mx-auto">
      <h2 class="text-2xl font-bold text-dark-slate mb-6">Background and expertise</h2>
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <NuxtLink
          v-for="s in expertiseServices"
          :key="s.id"
          :to="`/services/${s.id}`"
          class="block p-4 bg-light-gray rounded-lg hover:shadow-md transition-shadow"
        >
          <h3 class="font-semibold text-dark-slate">{{ s.menuTitle }}</h3>
          <p class="text-sm text-gray-600 mt-1">{{ s.summary }}</p>
        </NuxtLink>
      </div>
    </section>

    <!-- Related reading -->
    <section class="mt-20 max-w-4xl mx-auto">
      <h2 class="text-2xl font-bold text-dark-slate mb-6">Related Reading</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <NuxtLink
          to="/blog/the-shift-to-agentic-ai-and-autonomous-workflows"
          class="block p-4 bg-light-gray rounded-lg hover:shadow-md transition-shadow"
        >
          <h3 class="font-semibold text-dark-slate">The Shift to Agentic AI</h3>
          <p class="text-sm text-gray-600 mt-1">From conversational AI to autonomous workflows</p>
        </NuxtLink>
        <NuxtLink
          to="/blog/DataMeshGovernance"
          class="block p-4 bg-light-gray rounded-lg hover:shadow-md transition-shadow"
        >
          <h3 class="font-semibold text-dark-slate">Data Mesh Governance</h3>
          <p class="text-sm text-gray-600 mt-1">Enforcing policies without becoming a bottleneck</p>
        </NuxtLink>
        <NuxtLink
          to="/blog/VisionaryDataArchitecture"
          class="block p-4 bg-light-gray rounded-lg hover:shadow-md transition-shadow"
        >
          <h3 class="font-semibold text-dark-slate">Visionary Data Architecture</h3>
          <p class="text-sm text-gray-600 mt-1">Stop building data museums</p>
        </NuxtLink>
      </div>
    </section>

    <!-- CTA -->
    <section class="mt-20 text-center">
      <h2 class="text-3xl font-bold">Ready to streamline your workflows?</h2>
      <p class="mt-4 text-lg text-gray-600 max-w-2xl mx-auto">
        Whether you need to fix a data pipeline, connect scattered systems or build your first
        production AI agent, start with a 30-minute call. No salespeople, just a direct technical
        conversation about feasibility, timeline and cost.
      </p>
      <NuxtLink
        :to="{ name: 'booking' }"
        @click="trackCta('services_footer')"
        class="mt-6 inline-block bg-primary text-white font-bold py-3 px-8 rounded-lg hover:bg-opacity-90 transition-colors"
      >
        Book a Scoping Call
      </NuxtLink>
    </section>
  </div>
</template>
