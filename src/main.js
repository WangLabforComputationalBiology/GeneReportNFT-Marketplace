import './assets/main.css'
import 'element-plus/dist/index.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'//持久化存储
import ElementPlus from 'element-plus'

/*全局组件*/
import Header from './views/components/header.vue'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
const app = createApp(App)
    .use(router)
    .use(ElementPlus)
    .use(pinia)

app.component('Header', Header)

app.config.productionTip = false; // 关闭生产提示

app.mount('#app')








