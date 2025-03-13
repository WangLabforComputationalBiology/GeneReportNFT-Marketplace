//路由页面
import create from "@/views/create.vue";
import drop from "@/views/drop.vue";
import stats from "@/views/stats.vue";
import index from "@/views/index.vue";
import login from "@/views/login.vue";
import market from "@/views/market.vue";
import { createWebHistory } from "vue-router";
import { createRouter } from "vue-router";

const routes = [
    {
        path: '/', // 根路径
        redirect: '/index', // 重定向到 /home
    },
    {
        path:'/create',
        component:create
    },
    {
        path:"/drop",
        component:drop
    },
    {
        path:"/stats",
        component:stats
    },
    {
        path:"/index",
        component:index
    },
    {
        path:"/market",
        component:market
    },
    {
        path:"/login",
        component:login
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})

export default router