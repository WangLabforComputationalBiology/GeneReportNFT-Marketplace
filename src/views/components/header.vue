<template>
    <div class="header">
        <div class="side" @click="toHome">
            <div class="logo"><span style="color: #169608;">Bio</span>Chainer</div>
        </div>
        <div class="header-wrapper">
            <!-- 路由菜单 -->
            <div class="routers">
                <router-link to="/index" class="router-selection ">Home</router-link>
                <router-link to="/plaza" class="router-selection ">Data Plaza</router-link>
                <router-link to="/create" class="router-selection">Data Create</router-link>
                <router-link to="/about" class="router-selection">About</router-link>
                <router-link to="/login" class="router-selection" v-if="!account">Login</router-link>
            </div>

            <!-- <button  @click="redirectToOAuth">获取微基因数据</button> -->

        </div>
        <span class="side" v-if="account">
            <div class="account" @click="toUser">{{ fixedAccount }}</div>
        </span>
    </div>
</template>

<script>
import { useWalletStore } from '@/stores/account';

export default {
    name: "Header",
    data() {
        return {
            account: null,
        }
    },
    created() {
        const wallet = useWalletStore();
        this.account = wallet.address;
    },
    computed: {
        fixedAccount() {
            return this.account.slice(0, 5) + '...' + this.account.slice(-4);
        }
    }
    ,
    methods: {
        toHome() {
            this.$router.push('/index');
        }
        ,
        // redirectToOAuth() {
        //     window.location.href = import.meta.env.VITE_APP_BASE_URL+'/user/oauth2Wegene';
        // }
        toUser() {
            this.$router.push('/user');
        }
    }
}
</script>

<style lang="scss" scoped>
.header {
    z-index: 1000;
    width: 100vw;
    height: 5vh;
    min-width: 1200px;
    display: flex;
    position: sticky;
    top: 0;
    transition: background-color 0.5s;
    background-color: #fff;

    a {
        text-decoration: none;
    }

    .side {

        height: 5vh;

        flex: 1;
        min-width: 200px;
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        color: #16952d;

        .account {
            font-size: 20px;
            font-weight: 700;
        }

        .logo {
            height: 100%;
            padding-left: 15px;
            font-size: 36px;
            font-weight: 700;
            color: #333;
        }
    }

    .header-wrapper {
        align-items: center;
        justify-content: center;
        position: relative;
        width: 85vw;
        min-width: 800px;
        display: flex;

        .routers {
            // height: 60px;
            display: flex;
            position: absolute;
            right: 0;
        }

        .router-selection {
            position: relative;
            color: #333;
            height: 5vh;
            margin: 0 25px;
            font-size: 20px;
            align-content: center;
            font-weight: bold;

            &:after {
                content: '';
                position: absolute;
                bottom: 0;
                left: 0;
                width: 0;
                height: 3px;
                /* 边框厚度 */
                background: #333;
                /* 蓝色边框 */
                transition: width 0.4s ease;
                /* 动画持续0.4秒 */
            }

            &:hover:after {
                width: 100%;
                /* 鼠标悬停时边框宽度变为100% */
            }
        }

    }


}
</style>