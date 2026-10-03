<template>
  <article class="p-6 md:p-8 space-y-10">
    <!-- Header -->
    <header>
      <p v-if="service.eyebrow" class="text-xs font-bold uppercase tracking-wider text-primary">
        {{ service.eyebrow }}
      </p>
      <h1 class="mt-2 text-2xl md:text-3xl font-bold text-dark-slate leading-tight">
        {{ service.headline ?? service.menuTitle }}
      </h1>
      <p class="mt-4 text-gray-600 leading-relaxed">{{ service.intro ?? service.summary }}</p>

      <div
        v-if="service.price"
        class="mt-6 inline-flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-4 rounded-lg border-2 border-amber bg-amber/5 px-5 py-3"
      >
        <span class="text-2xl font-bold text-dark-slate whitespace-nowrap">{{ service.price.label }}</span>
        <span class="text-sm text-gray-700">{{ service.price.note }}</span>
      </div>
    </header>

    <!-- Problems -->
    <section v-if="service.problems?.length">
      <h2 class="text-lg font-bold text-dark-slate">Sounds familiar?</h2>
      <ul class="mt-4 space-y-2">
        <li v-for="problem in service.problems" :key="problem" class="flex gap-3 text-gray-700">
          <span class="mt-2 h-1.5 w-1.5 flex-shrink-0 rounded-full bg-amber" aria-hidden="true" />
          <span>{{ problem }}</span>
        </li>
      </ul>
    </section>

    <!-- Deliverables -->
    <section v-if="service.deliverables?.length">
      <h2 class="text-lg font-bold text-dark-slate">What I deliver</h2>
      <div class="mt-4 grid grid-cols-1 xl:grid-cols-2 gap-4">
        <div
          v-for="item in service.deliverables"
          :key="item.title"
          class="rounded-lg border border-gray-200 bg-light-gray p-4"
        >
          <h3 class="font-semibold text-dark-slate">{{ item.title }}</h3>
          <p class="mt-1 text-sm text-gray-600">{{ item.text }}</p>
        </div>
      </div>
    </section>

    <!-- Use cases -->
    <section v-if="service.useCases?.length">
      <h2 class="text-lg font-bold text-dark-slate">Typical projects</h2>
      <ul class="mt-4 space-y-2">
        <li v-for="useCase in service.useCases" :key="useCase" class="flex gap-3 text-gray-700">
          <svg
            class="mt-1 h-4 w-4 flex-shrink-0 text-primary"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 13l4 4L19 7"
            />
          </svg>
          <span>{{ useCase }}</span>
        </li>
      </ul>
    </section>

    <!-- Stack -->
    <section v-if="service.stack?.length">
      <h2 class="text-lg font-bold text-dark-slate">Tools I use</h2>
      <div class="mt-4 flex flex-wrap gap-2">
        <span
          v-for="tool in service.stack"
          :key="tool"
          class="text-sm font-medium text-primary-dark bg-white border border-gray-200 rounded-full px-3 py-1"
        >
          {{ tool }}
        </span>
      </div>
    </section>

    <!-- FAQs -->
    <section v-if="service.faqs?.length">
      <h2 class="text-lg font-bold text-dark-slate">Questions</h2>
      <div class="mt-4 divide-y divide-gray-200 border-y border-gray-200">
        <details v-for="faq in service.faqs" :key="faq.question" class="group py-3">
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

    <!-- Next step up the ladder -->
    <NuxtLink
      v-if="nextStep"
      :to="`/services/${nextStep.id}`"
      class="block rounded-lg border border-primary/30 bg-primary/5 p-5 hover:shadow-md transition-shadow"
    >
      <p class="text-xs font-bold uppercase tracking-wider text-primary">Next step</p>
      <p class="mt-1 font-semibold text-dark-slate">{{ nextStep.menuTitle }} &rarr;</p>
      <p class="mt-1 text-sm text-gray-600">{{ nextStep.summary }}</p>
    </NuxtLink>
  </article>
</template>

<script setup lang="ts">
import { getServiceById, type Service } from '~/data/services'

const props = defineProps<{ service: Service }>()

const nextStep = computed(() => getServiceById(props.service.nextStepId))
</script>
