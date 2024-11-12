/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/**/*.{go,js,templ,html,txt}"
    ],
    theme: {
      extend: {},
    },
    plugins: [
      require('daisyui'),
    ],
  }