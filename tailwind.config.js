/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './components/**/*.{vue,js,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './composables/**/*.ts',
    './plugins/**/*.ts',
    './app.vue',
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Montserrat', 'sans-serif'],
        cursive: ['"Brush Script MT"', 'cursive'],
      },
      colors: {
        primary: '#00a896',
        'dark-slate': '#1e293b',
        accent: '#00a896',
        'light-gray': '#f8f9fa',
        'accent-light': '#02c39a',
        secondary: '#028090',
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
}
