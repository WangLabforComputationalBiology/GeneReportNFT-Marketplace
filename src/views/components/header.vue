<template>
    <div class="header">
        <div class="side" @click="toHome">
            <img src="@/icons/BioChainer.svg" width="175px" fetchpriority="high" loading="eager" alt="logo">
        </div>
        <div class="header-wrapper">
            <!-- 路由菜单 -->
            <div class="routers">
                <router-link to="/index"
                    :class="['router-selection', { 'active': $route.path === '/index' }]">Home</router-link>
                <router-link to="/plaza" :class="['router-selection', { 'active': $route.path === '/plaza' }]">Data
                    Plaza</router-link>
                <router-link to="/publish" :class="['router-selection', { 'active': $route.path === '/publish' }]">Data
                    Publish</router-link>
                <router-link to="/about"
                    :class="['router-selection', { 'active': $route.path === '/about' }]">About</router-link>
                <router-link to="/login" class="router-selection" v-if="!account">Login</router-link>
            </div>
        </div>
        <span class="side" v-if="account">
            <div class="account" @click="toAccount">{{ fixedAccount }}</div>
        </span>
    </div>
</template>

<script setup>
import { useWalletStore } from '@/stores/account'
import { computed } from 'vue'
import { useRouter } from 'vue-router'

const wallet = useWalletStore()
const account = computed(() => wallet.address)
const router = useRouter()

const fixedAccount = computed(() => {
    if (!account.value) return ''
    return account.value.slice(0, 8) + '...' + account.value.slice(-6)
})

function toHome() {
    router.push('/index')
}

function toAccount() {
    router.push('/account')
}
</script>

<style lang="scss" scoped>
.header {
    display: flex;
    z-index: 1000;
    width: 100vw;
    height: 5vh;
    min-height: 45px;
    min-width: 700px;
    position: sticky;
    top: 0;
    background: #ffffff40;
    @include blur;

    a {
        text-decoration: none;
    }

    .side {
        height: 5vh;
        min-height: 45px !important;
        flex: 1;
        min-width: 220px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: #16952d;

        .account {
            font-size: 20px;
            font-weight: 700;
            cursor: pointer;
        }
    }

    .header-wrapper {
        align-items: center;
        justify-content: center;
        position: relative;
        width: 85vw;
        min-width: 800px;
        min-height: 45px;
        display: flex;

        .routers {
            display: flex;
            position: absolute;
            right: 0;
        }

        .router-selection {
            position: relative;
            color: #333;
            height: 5vh;
            gap: 15px;
            margin: 0 10px 0 20px;
            font-size: 20px;
            align-content: center;
            font-weight: bold;

            &:after {
                content: '';
                position: absolute;
                background: #169608;
                bottom: 0;
                left: 0;
                width: 0;
                height: 3px;
                /* 边框厚度 */
                transition: width 0.4s ease;
                /* 动画持续0.4秒 */
            }

            &:hover:after {
                width: 100%;
                /* 鼠标悬停时边框宽度变为100% */
            }
        }

        .active {
            color: #169608;

            &::after {
                width: 100%;
            }
        }

    }
}
</style>