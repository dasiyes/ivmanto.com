<template>
  <div class="relative flex justify-center" aria-hidden="true">
    <!-- Animated Abstract Orb -->
    <div class="relative w-[300px] h-[300px] md:w-[400px] md:h-[400px]">
      <!-- Outer ring -->
      <div
        class="absolute inset-0 rounded-full animate-pulse-glow"
        style="background: conic-gradient(from 0deg, transparent 0%, rgba(0,168,150,.3) 25%, transparent 50%, rgba(2,195,154,.3) 75%, transparent 100%); animation: orbSpin 12s linear infinite, pulseGlow 3s ease-in-out infinite;"
      ></div>

      <!-- Mid glow layer -->
      <div
        class="absolute inset-6 rounded-full"
        style="background: radial-gradient(circle, rgba(0,168,150,.15) 0%, transparent 70%); animation: orbSpin 8s linear infinite reverse;"
      ></div>

      <!-- Core gradient orb -->
      <div
        class="absolute inset-12 rounded-full animate-float"
        style="background: radial-gradient(circle at 40% 40%, rgba(2,195,154,.5), rgba(0,168,150,.3) 40%, rgba(2,128,144,.2) 70%, transparent); box-shadow: 0 0 80px rgba(0,168,150,.3), inset 0 0 60px rgba(2,195,154,.1);"
      ></div>

      <!-- Floating inner accent dots -->
      <div
        class="absolute top-1/4 left-1/3 w-3 h-3 rounded-full bg-primary-light/60"
        style="animation: float 4s ease-in-out infinite;"
      ></div>
      <div
        class="absolute bottom-1/3 right-1/4 w-2 h-2 rounded-full bg-primary/50"
        style="animation: float 5s ease-in-out 1s infinite;"
      ></div>
      <div
        class="absolute top-1/2 right-1/3 w-1.5 h-1.5 rounded-full bg-accent-light/40"
        style="animation: float 6s ease-in-out 0.5s infinite;"
      ></div>
    </div>
  </div>
</template>

<style scoped>
@keyframes orbSpin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>
