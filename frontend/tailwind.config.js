/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'dark-gradient-start': '#1a202c',
        'dark-gradient-end': '#2d3748',
      },
    },
  },
  plugins: [],
}