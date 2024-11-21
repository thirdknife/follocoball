import franken from 'franken-ui/shadcn-ui/preset-quick';

/** @type {import('tailwindcss').Config} */
module.exports = {
  presets: [franken()],
  content: ["./templa/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue,templ}"],
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

