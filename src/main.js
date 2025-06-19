import './assets/main.css'
import 'element-plus/dist/index.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'//持久化存储
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue' //导入 ElementPlus 组件库中的所有图标


/*全局组件*/
import Header from './views/components/header.vue'



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

app.component('Header', Header)

app.config.productionTip = false; // 关闭生产提示

app.mount('#app')








