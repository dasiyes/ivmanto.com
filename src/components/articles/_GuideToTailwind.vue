<script lang="ts">
// This <script> block is for module-level exports like metadata.
export const metadata = {
  title: 'A Practicle guide to ... (series)',
  summary: 'Learn how to on specific topic ...',
  date: '2024-07-20',
}
</script>

<script setup lang="ts">
// This <script setup> block is for the component's Composition API logic.
</script>

<template>
  <article class="prose lg:prose-xl max-w-none">
    <h1>A Practical Guide to Tailwind CSS</h1>
    <p class="lead">
      Learn how to speed up your development workflow with Tailwind CSS. We cover the basics,
      customization, and best practices for building modern UIs.
    </p>
    <p>
      Tailwind CSS is a utility-first CSS framework that provides low-level utility classes to build
      custom designs without ever leaving your HTML. It's a different approach compared to
      component-based frameworks like Bootstrap or Foundation.
    </p>
    <h2>The Utility-First Approach</h2>
    <p>
      Instead of pre-styled components like `.card` or `.button`, Tailwind provides utility classes
      like `flex`, `pt-4`, `text-center`, and `rotate-90` that can be composed to build any design,
      directly in your markup.
    </p>
    <p>
      This approach gives you complete control over the styling of your components and avoids the
      need to write custom CSS for every little thing.
    </p>
  </article>
</template>
