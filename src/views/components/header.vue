<template>
    <div class="header">
        <div class="side" @click="toHome">
            L
        </div>
        <div class="header-wrapper">
            <!-- 路由菜单 -->
            <div class="routers">
                <router-link to="/index" class="router-selection ">Home</router-link>
                <router-link to="/market" class="router-selection ">Market</router-link>
                <router-link to="/create" class="router-selection">Create</router-link>
                <router-link to="/login" class="router-selection" v-if="!account">Login</router-link>
                <router-link to="/about" class="router-selection">About</router-link>
                <!-- <el-icon class="userIcon" v-if="account">
                    <router-link to="/user" style="color: #169608;">
                        <UserFilled />
                    </router-link>
                </el-icon> -->
                <!-- <div v-if="account" >{{ fixedAccount }}</div> -->
            </div>

            <!-- <button  @click="redirectToOAuth">获取微基因数据</button> -->

        </div>
        <span class="side">
            <div class="account" v-if="account">{{ fixedAccount }}</div>
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
    }
}
</script>

<style lang="scss" scoped>
.header {
    z-index: 1000;
    width: 99vw;
    height: 60px;
    display: flex;
    position: sticky;
    top: 0;
    transition: background-color 0.5s;

    a {
        text-decoration: none;
    }

    .side {
        height: 60px;
        flex: 1;
        min-width: 120px;
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        color: #16952d;

        .account {
            font-size: 20px;
            font-weight: 700;
        }
    }

    .header-wrapper {
        align-items: center;
        justify-content: center;
        position: relative;
        width: 85vw;
        min-width: 1200px;
        display: flex;

        .routers {
            height: 60px;
            display: flex;
            position: absolute;
            right: 0;
        }

        .router-selection {
            color: #169608;
            margin: 0 25px;
            font-size: 20px;
            align-content: center;
            font-weight: bold;

            &:hover {
                color: #67C23A;
            }

            &:first-child {
                margin-left: 0;
            }
        }

    }


}

.userIcon {
    position: absolute;
    right: 0;
    font-size: 30px;
    color: #169608;
    cursor: pointer;
}
</style>