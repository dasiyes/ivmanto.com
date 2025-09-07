<template>
  <div class="flex justify-center" aria-hidden="true">
    <!-- Frame with soft-focus background -->
    <div
      class="relative flex items-center justify-center w-[440px] h-[440px] overflow-hidden rounded-3xl bg-white/10 shadow-lg"
    >
      <!-- Blurred background -->
      <img
        src="/cloud-pic-2.webp"
        alt="blured-frame"
        class="absolute inset-0 h-full w-full scale-125 object-cover blur-2xl"
        width="263"
        height="263"
      />
      <!-- Foreground image -->
      <img
        src="/cloud-pic-2.webp"
        alt="front-pic"
        class="relative w-[350px] drop-shadow-lg"
        width="413"
        height="413"
      />
    </div>
  </div>
</template>
