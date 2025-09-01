import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    proxy: {
      // Proxy all requests starting with /api to the Go backend
      '/api': {
        target: 'http://localhost:8080', // The address of your running Go server
        changeOrigin: true,
        // DO NOT rewrite the path. The Go backend's routes are already prefixed with /api.
        // The absence of a 'rewrite' rule here is the fix.
      },
    },
  },
})
