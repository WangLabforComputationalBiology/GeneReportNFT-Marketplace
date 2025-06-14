//路由页面
import { createWebHistory } from "vue-router";
import { createRouter } from "vue-router";
import walletAuthGuard from "./accountGuard";

const routes = [
    {
        path: '/', // 根路径
        redirect: '/index', // 重定向到 /home
    },
    {
        path: "/index",
        component: () => import('@/views/home/index.vue')
    },
    {
        path: '/publish',
        component: () => import('@/views/publish/publish.vue')
    },
    {
        path: "/create/selectProfile/:lastSegment",//最后面的动态参数不能改
        component: () => import('@/views/publish/selectProfile.vue')
    },
    {
        path: "/plaza",
        component: () => import('@/views/plaza/plaza.vue')
    },
    {
        path: '/plaza/confirm',
        component: () => import('@/views/plaza/confirm.vue')
    },
    {
        path: '/verify',
        component: () => import('@/views/verify/verify.vue')
    },
    {
        path: "/login",
        component: () => import('@/views/account/login.vue')
    },
    {
        path: "/account",
        component: () => import('@/views/account/account.vue'),
        meta: { requiresAuth: true },
        beforeEnter: walletAuthGuard,
    },
    {
        path: "/about",
        component: () => import('@/views/about.vue')
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return { top: 0, left: 0 }; // 每次切换路由时回到顶部
    }
})

export default router