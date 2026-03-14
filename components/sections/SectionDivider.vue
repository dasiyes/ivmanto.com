<script setup lang="ts">
defineProps<{
  flip?: boolean
}>()
</script>

<template>
  <div class="relative w-full overflow-hidden" :class="flip ? 'rotate-180' : ''" style="height: 80px; margin-top: -1px; margin-bottom: -1px;">
    <svg
      class="absolute bottom-0 w-full h-full"
      viewBox="0 0 1440 80"
      preserveAspectRatio="none"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M0 40C240 80 480 0 720 40C960 80 1200 0 1440 40V80H0Z"
        fill="currentColor"
      />
    </svg>
  </div>
</template>
