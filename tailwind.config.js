/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/pages/*.html"],
  theme: {
    extend: {},
  },
  plugins: [require("tailwindcss-animation-delay")],
};
