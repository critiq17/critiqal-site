/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  darkMode: 'class',  // Для :root.dark в вашем app.css
  theme: {
    extend: {}
  },
  plugins: [
    require('@tailwindcss/forms')  // Для вашего @plugin в app.css
  ],
}
