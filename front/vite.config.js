import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

const __dirname = import.meta.dirname;

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@':  __dirname + '/src',
    }
  }
})
