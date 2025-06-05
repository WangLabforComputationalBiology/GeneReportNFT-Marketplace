//路由页面
import { createWebHistory } from "vue-router";
import { createRouter } from "vue-router";

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
        path: '/create',
        component: () => import('@/views/create/create.vue')
    },
    {
        path: "/create/selectProfile/:lastSegment",//最后面的动态参数不能改
        component: () => import('@/views/create/selectProfile.vue')
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
        path: "/plaza",
        component: () => import('@/views/plaza/plaza.vue')
    },
    {
        path: '/market/purchase',
        component: () => import('@/views/plaza/purchase.vue')
    },
    {
        path: '/verify',
        component: () => import('@/views/verify.vue')
    },
    {
        path: "/login",
        component: () => import('@/views/login.vue')
    },
    {
        path: "/user",
        component: () => import('@/views/user.vue')
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