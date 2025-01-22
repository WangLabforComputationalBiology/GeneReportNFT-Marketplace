//路由页面
import create from "@/components/create.vue";
import drop from "@/components/drop.vue";
import stats from "@/components/stats.vue";
import { createWebHistory } from "vue-router";
import { createRouter } from "vue-router";

const routes = [
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
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})

export default router