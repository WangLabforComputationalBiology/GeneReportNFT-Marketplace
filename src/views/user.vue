<template>

    <body>
        <div class="banner">
            <h1 class="banner-title">User</h1>
        </div>
        <span>{{ address }}</span><br>
        <span>Wallet balance: </span> <span>{{ balance }}</span>
    </body>
</template>

<script>
import { useWalletStore } from '@/stores/account';
const wallet = useWalletStore();
export default {
    name: "drop",
    data() {
        return {
            dropList: [],
            address: wallet.address,
            balance: wallet.balance,
        }
    },

    created() {
    },

    methods: {

    },
}
</script>

<style lang="scss" scoped>
// 全局样式
body {
    margin: auto;

    width: 80vw;
    min-width: 1200px;
    min-height: calc(100vh - 80px);
}
</style>

<style lang="scss" scoped>
.banner {
    display: flex;
    // border-bottom: 1px solid #E4E7ED;

    .banner-title {
        font-size: 70px;
        color: #67C23A;
    }

}

.dropPage {
    height: 700px;
}
</style>