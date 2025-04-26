<template>
    <div class="header">
        <div class="side" @click="toHome">
            L
        </div>
        <div class="headerWrapper">
            <!-- 路由菜单 -->
            <div class="routers">
                <router-link to="/market" class="routerSelection ">Market</router-link>
                <router-link to="/create" class="routerSelection">Create</router-link>
                <!-- <router-link to="/drop" class="routerSelection">Drop</router-link> -->
                <router-link to="/stats" class="routerSelection">Stats</router-link>
            </div>

            <!-- 搜索框 -->
            <span class="navigation">
                <input id="navigationInput" type="text" placeholder="search..." />
                <div type="primary" class="searchBtn">
                    <el-icon class="searchIcon" size="large">
                        <Search />
                    </el-icon>
                </div>
            </span>
            <!-- 登录路由 -->
            <router-link to="/login" class="routerSelection"
                style="position:absolute;right:20px;width:50px" v-if="!account">Login</router-link>
            <!-- 账户面板 -->
            <el-icon class="userIcon" v-if="account">
                <router-link to="/user" style="color: #169608;">
                    <UserFilled />
                </router-link>
            </el-icon>
            <div v-if="account" style="color: #909399;position: absolute; right: 50px">{{ fixedAccount }}</div>
        </div>
        <span class="side">
        </span>
    </div>
</template>

<script>
import { useWalletStore } from '@/stores/account';
import { Search, UserFilled } from '@element-plus/icons-vue';
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
    background-color: #ffffffee;
    border-bottom: 1px solid #E4E7ED;

    a {
        text-decoration: none;
    }

    .side {
        flex: 1;
        min-width: 120px;
        display: flex;
        align-items: center;
        cursor: pointer;
        font-size: 40px;
        color: #16952d;
        font-weight: bold;
    }

    .headerWrapper {
        align-items: center;
        position: relative;
        width: 85vw;
        min-width: 1200px;
        display: flex;
        margin: 0 auto;

        .navigation {
            width: 300px;
            display: flex;
            align-items: center;
            margin: 4px 0 0 20px;

            input {
                width: 250px;
                height: 35px;
                border: 2px solid #E4E7ED;
                border-radius: 10px;
            }

            .searchBtn {
                background-color: #169608;
                border-radius: 12px;
                margin-left: 4px;

                .searchIcon {
                    height: 34px;
                    width: 45px;
                    color: #fff;
                }

                &:hover {
                    background-color: #67C23A;
                    color: #FFF;
                    cursor: pointer;
                }
            }
        }

        .routers {
            height: 80px;
            display: flex;
        }

        .routerSelection {
            color: #169608;
            margin: 0 25px;
            font-size: 22px;
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