<script setup lang="ts">
const { cleanTitle } = usePageMetadata()
</script>

<template>
  <article class="prose prose-sm lg:prose-base max-w-none p-6">
    <h1>{{ cleanTitle }}</h1>
    <p class="italic lg:text-l">
      Cutting through the hype to find AI and automation that actually pays off.
    </p>
    <hr class="border-gray-200 my-2" />
    <p>
      Many small and medium-sized businesses want to leverage AI but don't know where to start. We
      help SMBs bypass the hype and identify
      <strong>high-impact, practical use-cases</strong> for automation. By auditing your current
      workflows, we pinpoint bottlenecks and design a clear <strong>automation roadmap</strong> that
      directly reduces operational costs and scales your capabilities.
    </p>
    <h4>Key Deliverables</h4>
    <ul>
      <li><strong>Workflow assessments</strong></li>
      <li><strong>ROI prioritization</strong></li>
      <li><strong>Actionable automation roadmaps</strong></li>
    </ul>
  </article>
</template>
