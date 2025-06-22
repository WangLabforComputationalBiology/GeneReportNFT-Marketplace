import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  css: {
    preprocessorOptions: {
      scss: {
        api: 'modern-compiler' // "modern"代替"legacy"
      },
    },
  },
  server: {
    host: '0.0.0.0', // 允许所有IP访问
    port: 5173,      // 确保端口未被占用
    strictPort: true // 如果端口被占用则报错
  }

})
