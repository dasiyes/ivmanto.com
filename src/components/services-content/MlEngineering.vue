<script setup lang="ts">
// 1. Import your composable
import { usePageMetadata } from '@/composables/usePageMetadata'
// 2. Get the reactive title
const { cleanTitle } = usePageMetadata()
</script>

<template>
  <article class="prose prose-sm lg:prose-base max-w-none p-6">
    <h2>{{ cleanTitle }}</h2>
    <p class="italic lg:text-l">
      Operationalizing machine learning models from prototype to production.
    </p>
    <p class="italic lg:text-sm">(A 0-minute read)</p>
    <hr class="border-gray-200 my-2" />
    <p>
      Our MLOps services streamline the machine learning lifecycle. We use
      <strong>VertexAI</strong> to build automated <strong>CI/CD</strong> pipelines for model
      training, deployment, and monitoring, ensuring your models deliver continuous value.
    </p>
  </article>
</template>
