<template>
  <footer class="bg-dark-slate text-gray-400 border-t border-white/5">
    <div class="container mx-auto px-6 py-12">
      <!-- Main footer grid -->
      <div class="grid md:grid-cols-3 gap-8 mb-8">
        <!-- Brand -->
        <div>
          <p class="text-white font-bold text-lg mb-2">IVMANTO</p>
          <p class="text-sm leading-relaxed">
            Expert Cloud Data Architecture & AI Solutions built on Google Cloud Platform.
          </p>
        </div>

        <!-- Quick Links -->
        <div>
          <p class="text-white font-semibold text-sm tracking-widest uppercase mb-4">Quick Links</p>
          <div class="flex flex-col space-y-2 text-sm">
            <NuxtLink to="/about" class="hover:text-primary transition-colors">About</NuxtLink>
            <NuxtLink to="/services" class="hover:text-primary transition-colors">Services</NuxtLink>
            <NuxtLink to="/assessment" class="hover:text-primary transition-colors">Assessment</NuxtLink>
            <NuxtLink to="/blog" class="hover:text-primary transition-colors">Blog</NuxtLink>
            <NuxtLink to="/#contact" class="hover:text-primary transition-colors">Contact</NuxtLink>
          </div>
        </div>

        <!-- Connect & Friends -->
        <div>
          <p class="text-white font-semibold text-sm tracking-widest uppercase mb-4">Connect</p>
          <div class="flex flex-col space-y-2 text-sm">
            <a href="https://linkedin.com/in/nikolaytonev" target="_blank" rel="noopener noreferrer" class="hover:text-primary transition-colors">LinkedIn</a>
            <a href="https://github.com/dasiyes" target="_blank" rel="noopener noreferrer" class="hover:text-primary transition-colors">GitHub</a>
            <a href="https://x.com/IvmantoSol" target="_blank" rel="noopener noreferrer" class="hover:text-primary transition-colors">X / Twitter</a>
            <a href="https://g.dev/ivmanto-nikolaytonev" target="_blank" rel="noopener noreferrer" class="hover:text-primary transition-colors">Google Developer Profile</a>
          </div>

          <p class="text-white font-semibold text-sm tracking-widest uppercase mt-6 mb-3">Friends</p>
          <div class="flex flex-col space-y-2 text-sm">
            <a href="https://vivily.de/" target="_blank" rel="noopener noreferrer" class="hover:text-amber transition-colors">Vivily.de</a>
          </div>
        </div>
      </div>

      <!-- Divider -->
      <div class="border-t border-white/10 pt-8 flex flex-col md:flex-row justify-between items-center gap-4">
        <p class="text-sm">&copy; 2025 "IVMANTO" Nikolay Tonev Data & Cloud Solutions. All Rights Reserved.</p>
        <div class="flex space-x-6 text-sm">
          <NuxtLink to="/privacy-policy" class="hover:text-primary transition-colors">Privacy Policy</NuxtLink>
          <NuxtLink to="/terms" class="hover:text-primary transition-colors">Terms of Service</NuxtLink>
        </div>
      </div>
    </div>
  </footer>
</template>
