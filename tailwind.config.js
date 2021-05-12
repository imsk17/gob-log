const colors = require("tailwindcss/colors")
module.exports = {
  purge: false,
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      boxShadow: {
        'sharp-xl': '8px 8px 0px #000',
        'sharp-lg': '5px 5px 0px #000',
        'sharp-md': '3px 3px 0px #000'
      },
      colors: {
        lime: colors.lime
      }
    },
    
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
