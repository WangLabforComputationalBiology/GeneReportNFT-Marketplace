<template>

    <body>
        <div class="banner">
            <h1 class="banner-title">User</h1>
        </div>
        <span>{{ address }}</span><br>
        <span>Wallet balance: </span> <span>{{ balance }}</span>
        <router-link to="/verify" style="color: #fff;">
            <el-button>
                verify
            </el-button>
        </router-link>
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
.banner {
    display: flex;

    .banner-title {
        font-size: 70px;
        color: #67C23A;
    }

}

.dropPage {
    height: 700px;
}

</style>