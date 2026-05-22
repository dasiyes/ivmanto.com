<script setup lang="ts">
const { cleanTitle } = usePageMetadata()
</script>

<template>
  <article class="prose prose-sm lg:prose-base max-w-none p-6">
    <h1>{{ cleanTitle }}</h1>
    <p class="italic lg:text-l">
      Reliable, scalable pipelines that turn raw data into real business value.
    </p>
    <hr class="border-gray-200 my-2" />
    <p>
      AI and automation are only as good as the data powering them. We design and architect
      <strong>robust, scalable data pipelines</strong> that ingest raw information and transform it
      into <strong>high-quality, curated datasets</strong>. Whether preparing for advanced analytics
      or feeding AI models, we ensure your <strong>data infrastructure</strong> is reliable, secure,
      and ready to deliver real business value.
    </p>
    <h4>Key Deliverables</h4>
    <ul>
      <li><strong>Architecture design</strong></li>
      <li><strong>Data integration</strong></li>
      <li><strong>Production-ready curated data environments</strong></li>
    </ul>
  </article>
</template>
