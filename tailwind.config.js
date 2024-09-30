const { fontFamily } = require("tailwindcss/defaultTheme")
import franken from 'franken-ui/shadcn-ui/preset-quick';

/** @type {import('tailwindcss').Config} */
module.exports = {
  presets: [franken()],
  content: ["**/*.html", "**/*.templ"],
  safelist: [
    {
      pattern: /^uk-/
    },
    'ProseMirror',
    'ProseMirror-focused',
    'tiptap'
  ],
  theme: {
    extend: {}
  },
  plugins: [],
}

