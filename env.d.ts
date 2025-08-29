/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  // Replace the final 'any' with 'unknown' for full type safety
  const component: DefineComponent<Record<string, unknown>, Record<string, unknown>, unknown>
  export default component
}
