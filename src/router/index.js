//路由页面
import Test1 from "@/components/test1.vue";
import Test2 from "@/components/test2.vue";
import Router from "vue-router";

export default new Router({
    routes:[
        {
            path:'/test1',
            name:'test1',
            component:Test1

        },
        {
            path:"/test2",
            name:'test2',
            component:Test2
        }
    ]
})

