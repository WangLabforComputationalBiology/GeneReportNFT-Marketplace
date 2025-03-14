//路由页面
import index from "@/views/index.vue";
// import create from "@/views/create.vue";
// import drop from "@/views/drop.vue";
// import stats from "@/views/stats.vue";
// import login from "@/views/login.vue";
// import market from "@/views/market.vue";
import { createWebHistory } from "vue-router";
import { createRouter } from "vue-router";

const create = () => import('@/views/create.vue');
const drop = () => import('@/views/drop.vue');
const stats = () => import('@/views/stats.vue');
const login = () => import('@/views/login.vue');
const market = () => import('@/views/market.vue');
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
        component: create
    },
    {
        path: "/drop",
        component: drop
    },
    {
        path: "/stats",
        component: stats
    },
    {
        path: "/market",
        component: market
    },
    {
        path: "/login",
        component: login
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