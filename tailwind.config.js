const colors = require("tailwindcss/colors")
module.exports = {
  purge: [
    './templates/*.html',
  ],
  darkMode: false, // or 'media' or 'class'
  theme: {},
  corePlugins: {
    preflight: false
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
