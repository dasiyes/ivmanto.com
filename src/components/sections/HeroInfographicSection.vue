<template>
  <div class="flex justify-center" aria-hidden="true">
    <!-- Frame with soft-focus background -->
    <div
      class="relative flex items-center justify-center w-[440px] h-[440px] overflow-hidden rounded-3xl bg-white/10 shadow-lg"
    >
      <!-- Blurred background -->
      <img
        src="@/assets/cloud-pic-2.png"
        alt=""
        class="absolute inset-0 h-full w-full scale-125 object-cover blur-2xl"
      />
      <!-- Foreground image -->
      <img src="@/assets/cloud-pic-2.png" alt="" class="relative w-[350px] drop-shadow-lg" />
    </div>
  </div>
</template>
