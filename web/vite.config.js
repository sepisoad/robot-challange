import { defineConfig } from 'vite'
import { plugin } from 'vite-plugin-elm'
import { viteSingleFile } from "vite-plugin-singlefile"

export default defineConfig({
  plugins: [plugin(), viteSingleFile()],
  server: {
    proxy: {
      '/api/robots': 'http://0.0.0.0:8080',
    },
  },
})

