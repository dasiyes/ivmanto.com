/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Montserrat', 'sans-serif'],
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
  plugins: [],
}
