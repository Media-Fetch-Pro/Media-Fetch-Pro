import { fileURLToPath, URL } from 'node:url'
import VueI18n from "@intlify/vite-plugin-vue-i18n";
import { resolve } from "path";

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:7789/',
        changeOrigin: true,
      },
    }
  },
  plugins: [
    vue(),
    VueI18n({
      include: resolve(__dirname, "./src/locales/**"),
      runtimeOnly: false,
    }),],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
