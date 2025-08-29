<template>
  <div class="relative bg-gray-50 px-6 pt-16 pb-20 lg:px-8 lg:pt-24 lg:pb-28">
    <div class="absolute inset-0">
      <div class="h-1/3 bg-white sm:h-2/3" />
    </div>
    <div class="relative mx-auto max-w-7xl">
      <div class="text-center">
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">`<ivm />`</h2>
        <p class="mx-auto mt-3 max-w-2xl text-xl text-gray-500 sm:mt-4">
          Insights, articles, and deep dives into the world of modern web development.
        </p>
      </div>
      <div class="mx-auto mt-12 grid max-w-lg gap-5 lg:max-w-none lg:grid-cols-3">
        <div
          v-for="article in articles"
          :key="article.slug"
          class="flex flex-col overflow-hidden rounded-lg shadow-lg"
        >
          <div class="flex flex-1 flex-col justify-between bg-white p-6">
            <div class="flex-1">
              <p class="text-sm font-medium text-primary">
                <span class="hover:underline">Article</span>
              </p>
              <RouterLink :to="`/blog/${article.slug}`" class="mt-2 block">
                <p class="text-xl font-semibold text-gray-900">{{ article.title }}</p>
                <p class="mt-3 text-base text-gray-500">{{ article.summary }}</p>
              </RouterLink>
            </div>
            <div class="mt-6 flex items-center">
              <div class="text-sm text-gray-500">
                <time :datetime="article.date">{{
                  new Date(article.date).toLocaleDateString('en-US', {
                    year: 'numeric',
                    month: 'long',
                    day: 'numeric',
                  })
                }}</time>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { articles } from '@/data/articles'
</script>
