//路由页面
import index from "@/views/index.vue";
// import create from "@/views/create.vue";
// import drop from "@/views/drop.vue";
// import stats from "@/views/stats.vue";
// import login from "@/views/login.vue";
// import market from "@/views/market.vue";
import { createWebHistory } from "vue-router";
import { createRouter } from "vue-router";

const routes = [
    {
        path: '/', // 根路径
        redirect: '/index', // 重定向到 /home
    },
    {
        path: "/index",
        component: index
    },
    {
        path: '/create',
        component: () => import('@/views/create.vue')
    },
    {
        path: "/drop",
        component: () => import('@/views/drop.vue')
    },
    {
        path: "/stats",
        component: () => import('@/views/stats.vue')
    },
    {
        path: "/market",
        component: () => import('@/views/market.vue')
    },
    {
        path:'/market/purchase',
        component: () => import('@/views/market/purchase.vue')
    },
    
    {
        path: "/login",
        component: () => import('@/views/login.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return { top: 0, left: 0 }; // 每次切换路由时回到顶部,这是vue3写法
    }

})

export default router