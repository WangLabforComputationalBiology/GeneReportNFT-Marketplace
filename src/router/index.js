//路由页面
import { createWebHistory } from "vue-router";
import { createRouter } from "vue-router";
import walletAuthGuard from "./accountGuard";
import homeIndex from "@/views/home/index.vue";

const routes = [
    {
        path: '/',
        redirect: '/index', // 重定向
    }, 
    {
        path: "/index",
        component: homeIndex
    },
    {
        path: '/publish',
        component: () => import('@/views/publish/publish.vue')
    },
    {
        path: "/publish/selectProfile/:uuid",//最后面的动态参数不能改
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
        component: () => import('@/views/verify/verify.vue'),
        meta: { requiresAuth: true },
        beforeEnter: walletAuthGuard,
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
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('@/views/components/404.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return { top: 0, left: 0 }; // 每次切换路由时回到顶部
    }
})

export default router