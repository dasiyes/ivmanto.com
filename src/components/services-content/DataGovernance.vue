<script setup lang="ts">
import { ref } from 'vue'
const title = ref('Data Governance & Strategy')
</script>

<template>
  <article class="prose prose-sm lg:prose-base max-w-none p-6">
    <h2>{{ title }}</h2>
    <p class="italic lg:text-l">
      Implementing DAMA-aligned principles for data quality and security.
    </p>
    <p class="italic lg:text-sm">(A 0-minute read)</p>
    <hr class="border-gray-200 my-2" />

    <p>
      We help you establish a strong data governance framework based on
      <strong>DAMA</strong> principles. This ensures your data is accurate, consistent, and secure,
      transforming it into a trustworthy asset for decision-making.
    </p>
  </article>
</template>
