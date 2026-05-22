<script setup lang="ts">
const { cleanTitle } = usePageMetadata()
</script>

<template>
  <article class="prose prose-sm lg:prose-base max-w-none p-6">
    <h1>{{ cleanTitle }}</h1>
    <p class="italic lg:text-l">
      Beyond chatbots — autonomous agents that reason, act, and collaborate.
    </p>
    <hr class="border-gray-200 my-2" />
    <p>
      Go beyond basic chatbots. We specialize in designing and implementing next-generation
      <strong>Agentic AI solutions</strong> — autonomous agents capable of executing complex tasks,
      reasoning, and collaborating. Working side-by-side with your internal engineering or product
      teams, we provide the <strong>architectural guidance</strong>, best practices, and
      implementation support needed to deploy <strong>intelligent workflows</strong> that deliver
      tangible business outcomes.
    </p>
    <h4>Key Deliverables</h4>
    <ul>
      <li><strong>Multi-agent system architecture</strong></li>
      <li><strong>Technical implementation guidance</strong></li>
      <li><strong>Team workshops</strong></li>
    </ul>
  </article>
</template>
