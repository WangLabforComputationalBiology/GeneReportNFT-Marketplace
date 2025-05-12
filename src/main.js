import './assets/main.css'
import '../theme/index.css'
import 'element-plus/dist/index.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'//持久化存储
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue' //导入 ElementPlus 组件库中的所有图标


import Header from './views/components/header.vue'
import Footer from './views/components/footer.vue'


const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
const app = createApp(App)
    .use(router)
    .use(ElementPlus)
    .use(pinia)

    
//注册 ElementPlus 组件库中的所有图标到全局 Vue 应用中
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
//注册全局组件
app.component('Header', Header) 
app.component('Footer', Footer)

app.config.productionTip = false; // 关闭生产提示

app.mount('#app')








