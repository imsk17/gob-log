const colors = require("tailwindcss/colors")
module.exports = {
  purge: false,
  darkMode: false, // or 'media' or 'class'
  theme: {
    colors: {
      primarybg: '#1E1F21',
      secondarybg: '#292A2E',
      pink: '#8D6F95',
      yellow: '#F7D365',
      blue: '#819DB5',
      white: '#FFFFFF',
      green: '#B9C36A',
      red: '#B8505B'
    }
  },
  corePlugins: {
    preflight: false
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
