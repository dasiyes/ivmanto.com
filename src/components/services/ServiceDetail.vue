<template>
  <div class="p-8 h-full flex flex-col overflow-y-auto" ref="scrollContainer">
    <div v-if="service" class="bg-white/80 backdrop-blur-sm p-8 rounded-lg shadow-md">
      <h3 class="text-2xl font-bold text-dark-slate">{{ service.title }}</h3>
      <p class="mt-4 text-gray-700 leading-relaxed">
        <template v-for="(part, index) in parsedDetails" :key="index">
          <span
            v-if="part.isTag"
            :data-tag-name="part.tagName"
            class="text-primary font-semibold cursor-pointer hover:underline"
            @mouseover="updateDetailView(part.tagName)"
            >{{ part.text }}</span
          >
          <span v-else>{{ part.text }}</span>
        </template>
      </p>
    </div>
    <div v-else class="flex-grow flex items-center justify-center text-center text-gray-500">
      <p class="text-lg">Select a service from the left to see the details.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch, onBeforeUnmount, nextTick } from 'vue'
import type { Service } from '@/data/services'

const props = defineProps<{
  service: Service | undefined
}>()

const emit = defineEmits(['update-right-column'])

const scrollContainer = ref<HTMLElement | null>(null)
let observer: IntersectionObserver | null = null

const parsedDetails = computed(() => {
  if (!props.service?.details) return []
  const regex = /(#\w+)/g
  return props.service.details.split(regex).map((part) => {
    if (part.startsWith('#')) {
      const tagName = part.substring(1)
      if (props.service?.tagDetails && tagName in props.service.tagDetails) {
        return { isTag: true, text: part, tagName }
      }
    }
    return { isTag: false, text: part, tagName: '' }
  })
})

function setupObserver() {
  if (observer) {
    observer.disconnect()
  }

  const options = {
    root: scrollContainer.value,
    rootMargin: '-20% 0px -70% 0px', // A "trigger zone" in the upper part of the container
    threshold: 0,
  }

  observer = new IntersectionObserver((entries) => {
    const intersectingEntry = entries.find((entry) => entry.isIntersecting)
    if (intersectingEntry) {
      const tagName = (intersectingEntry.target as HTMLElement).dataset.tagName
      updateDetailView(tagName)
    }
  }, options)

  nextTick(() => {
    if (scrollContainer.value) {
      const tags = scrollContainer.value.querySelectorAll('[data-tag-name]')
      tags.forEach((tag) => observer!.observe(tag))
    }
  })
}

function updateDetailView(tagName: string | undefined) {
  // If we receive an invalid tag name or there are no details, do nothing.
  // This makes the content "stick" until a new valid event occurs.
  if (!tagName || !props.service?.tagDetails || !props.service.tagDetails[tagName]) {
    return
  }
  // Otherwise, emit the new content to the parent.
  emit('update-right-column', props.service.tagDetails[tagName])
}

watch(
  () => props.service,
  (newService) => {
    emit('update-right-column', undefined)
    if (scrollContainer.value) {
      scrollContainer.value.scrollTop = 0
    }
    if (newService) {
      setupObserver()
    }
  },
  { immediate: true },
)

onBeforeUnmount(() => {
  if (observer) {
    observer.disconnect()
  }
})
</script>
