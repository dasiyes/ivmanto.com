<template>
  <div class="p-6 h-full">
    <Transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="transform opacity-0 translate-y-2"
      enter-to-class="transform opacity-100 translate-y-0"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="transform opacity-100 translate-y-0"
      leave-to-class="transform opacity-0 translate-y-2"
    >
      <div v-if="content">
        <h4 class="font-bold text-dark-slate text-lg">Keyword Details</h4>
        <p class="mt-2 text-gray-600 text-sm leading-relaxed">{{ content }}</p>
      </div>
      <div v-else class="h-full flex items-center justify-center text-center text-gray-400">
        <p>Hover over a #tag in the service description for more details.</p>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  content: string | undefined
}>()
</script>
