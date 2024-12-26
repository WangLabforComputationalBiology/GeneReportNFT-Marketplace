import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import Header from './components/Header.vue'


createApp(App).mount('#app')



export default{
    component:{
        Header
    }
}
