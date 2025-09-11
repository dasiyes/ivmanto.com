<script lang="ts">
// This <script> block is for module-level exports like metadata.
export const metadata = {
  title: 'Mastering Vue',
  summary: 'How to be proficient in Vue',
  date: '2024-07-20',
}
</script>

<script setup lang="ts">
// This <script setup> block is for the component's Composition API logic.
</script>

<template>
  <article class="prose lg:prose-xl max-w-none">
    <h1>Mastering Vue 3 Composition API</h1>
    <p class="lead">
      A deep dive into the Vue 3 Composition API. Learn how to write cleaner, more reusable, and
      more scalable Vue components.
    </p>
    <p>
      The Composition API is one of the most significant new features in Vue 3. It provides a
      flexible, powerful way to organize and reuse component logic. This article will guide you
      through its core concepts.
    </p>
    <h2>Why the Composition API?</h2>
    <p>
      The Options API, while great for small to medium-sized components, can become cumbersome as
      components grow. Logic related to a single feature can be scattered across `data`, `methods`,
      `computed`, and `watch`. The Composition API solves this by allowing you to group related
      logic together.
    </p>
  </article>
</template>
